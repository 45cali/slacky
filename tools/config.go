package tools

import (
	"fmt"
	"io/ioutil"
	"os"

	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	slack    slackConfig
	Name     string
	Channels []string
	Messages []Message
}

type Message struct {
	Contains string
	Response string
}

type slackConfig struct {
	token string
	url   string
	botID string
}

func (conf *Config) parseYaml(filename string, errors []error) []error {

	source, err := ioutil.ReadFile(filename)
	if err != nil {
		errors = append(errors, fmt.Errorf("could not read file: %s", filename))
	}
	err = yaml.Unmarshal(source, &conf)
	if err != nil {
		errors = append(errors, fmt.Errorf("could not decode yaml file: %s", filename))
	}

	return errors
}

func (conf *Config) getEnvVars(errors []error) []error {
	conf.slack.url = os.Getenv("SLACK_URL")
	conf.slack.token = os.Getenv("SLACK_TOKEN")
	conf.slack.botID = os.Getenv("SLACK_BOT_ID")

	if len(conf.slack.url) == 0 {
		errors = append(errors, fmt.Errorf("environment variable SLACK_URL not set"))
	}
	if len(conf.slack.token) == 0 {
		errors = append(errors, fmt.Errorf("environment variable SLACK_TOKEN not set"))
	}
	if len(conf.slack.botID) == 0 {
		errors = append(errors, fmt.Errorf("environment variable SLACK_BOT_ID not set"))
	}
	return errors
}

func GetConfig(f string) (conf Config, errors []error) {

	errors = conf.parseYaml(f, errors)
	errors = conf.getEnvVars(errors)
	errors = conf.validate(errors)
	return
}
