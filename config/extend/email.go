package config

type emailConfig struct {
	EmailAddress string
	EmailName    string
	Password     string
	SMTPServer   string
	Port         int
}

var EmailConfig = emailConfig{}

func (ec *emailConfig) Check() {

}
