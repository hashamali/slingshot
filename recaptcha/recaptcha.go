package recaptcha

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/hashamali/slingshot/config"
)

const recaptchaURL = "https://www.google.com/recaptcha/api/siteverify"

type recaptchaResponse struct {
	Success bool `json:"success"`
}

// Validate will validate a reCAPTCHA token if applicable.
func Validate(responseToken string, c config.RecaptchaConfig) bool {
	if c.Secret != "" {
		if responseToken == "" {
			return false
		}

		r, err := http.PostForm(recaptchaURL, url.Values{"secret": {c.Secret}, "response": {responseToken}})
		if err != nil {
			return false
		}

		defer r.Body.Close()
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return false
		}

		response := recaptchaResponse{}
		err = json.Unmarshal(body, &response)
		if err != nil {
			return false
		}

		return response.Success
	}

	return true
}
