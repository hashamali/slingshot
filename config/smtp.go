package config

// SMTPConfig contains config for SMTP communication.
type SMTPConfig struct {
	Host     string `required:"true" envconfig:"SMTP_HOST" mapstructure:"smtpHost"`
	Port     int    `required:"true" envconfig:"SMTP_PORT" mapstructure:"smtpPort"`
	Password string `required:"true" envconfig:"SMTP_PASSWORD" mapstructure:"smtpPassword"`
}
