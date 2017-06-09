package tools

import "fmt"

func (conf *Config) validate(errors []error) []error {

	if len(conf.Name) == 0 {
		errors = append(errors, fmt.Errorf("name parameter not provided in config file"))
	}
	if len(conf.Channels) == 0 {
		errors = append(errors, fmt.Errorf("no channels parameter specified in config file"))
	}
	if len(conf.Messages) == 0 {
		errors = append(errors, fmt.Errorf("no messages parameter specified in config file"))
	}

	return errors
}
