package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/hashamali/ghttp"
	"github.com/hashamali/slingshot/config"
	"github.com/hashamali/slingshot/email"
	"github.com/hashamali/slingshot/recaptcha"
)

func main() {
	var projectConfig *config.Config
	var err error

	useEnvVars := os.Getenv("USE_ENV_VARS")
	if useEnvVars != "" && strings.ToLower(useEnvVars) != "false" {
		projectConfig, err = config.GetConfigFromEnvVars()
	} else {
		projectConfig, err = config.GetConfigFromFile()
	}

	if err != nil {
		panic(err.Error())
	}

	emailSender := email.SMTPEmail{
		Config: projectConfig.SMTPConfig,
	}

	router, err := ghttp.GenerateRouter(nil)
	if err != nil {
		panic(err.Error())
	}

	router.RegisterRoute(ghttp.Route{
		Path:   "/ping",
		Method: http.MethodGet,
		Handle: func(w http.ResponseWriter, r *http.Request) {
			ghttp.RespondNoContent(w)
		},
	})

	router.RegisterRoute(ghttp.Route{
		Path:   "/form",
		Method: http.MethodPost,
		Handle: func(w http.ResponseWriter, r *http.Request) {
			// Parse form data.
			err := r.ParseForm()
			if err != nil {
				ghttp.RespondError(w, ghttp.CreateErrorWithError(err, http.StatusBadRequest, "Unable to read body.", nil))
				return
			}

			// Optional reCAPTCHA validation.
			responseToken := r.PostForm.Get("g-recaptcha-response")
			if !recaptcha.Validate(responseToken, projectConfig.RecaptchaConfig) {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			// Prepare message.
			message := "Slingshot Message:\r\n\r\n"
			for paramName, paramValue := range r.PostForm {
				message = fmt.Sprintf("%s%s: %s\r\n", message, paramName, paramValue)
			}

			// Send message.
			to := strings.Join(projectConfig.EmailsConfig.To, ", ")
			body := fmt.Sprintf("To: %s\r\nSubject: New Slingshot Message\r\n\r\n%s\r\n", to, message)
			err = emailSender.SendEmail(projectConfig.EmailsConfig.From, projectConfig.EmailsConfig.To, []byte(body))
			if err != nil {
				ghttp.RespondError(w, ghttp.CreateErrorWithError(err, http.StatusBadGateway, "Unable to send email.", nil))
				return
			}

			// Return 204.
			ghttp.RespondNoContent(w)
		},
	})

	router.Serve()
}
