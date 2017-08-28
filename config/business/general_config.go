package business

import (
	"fmt"
)

type GeneralConfig struct {
	DefaultLanguage    string   `json:"default_language"`
	SupportedLanguages []string `json:"supported_languages"`
}

func (c GeneralConfig) String() string {
	return fmt.Sprintf("gen:lang=%s;langs=%s", c.DefaultLanguage, c.SupportedLanguages)
}

func (c *GeneralConfig) Check() {
	if c.DefaultLanguage == "" {
		c.DefaultLanguage = "en"
	}
	if len(c.SupportedLanguages) < 1 {
		c.SupportedLanguages = append(c.SupportedLanguages, "en")
	}
}
