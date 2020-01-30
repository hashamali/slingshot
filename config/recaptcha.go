package config

// RecaptchaConfig contains config for reCAPTCHA.
type RecaptchaConfig struct {
	Secret string `envconfig:"RECAPTCHA_SECRET" mapstructure:"recaptchaSecret"`
}
