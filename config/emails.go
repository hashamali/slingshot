package config

// EmailsConfig contains config for emails to use.
type EmailsConfig struct {
	From string   `required:"true" envconfig:"EMAIL_FROM" mapstructure:"emailFrom"`
	To   []string `required:"true" envconfig:"EMAIL_TO" mapstructure:"emailTo"`
}
