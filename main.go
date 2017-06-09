package main

import (
	"fmt"
	"os"

	"github.com/45cali/slacky/tools"
)

func main() {

	filename := "example_config.yaml"
	conf, errors := tools.GetConfig(filename)
	if len(errors) > 0 {
		for _, e := range errors {
			fmt.Println(e)
		}
		os.Exit(0)
	}
	fmt.Println(conf)
}
