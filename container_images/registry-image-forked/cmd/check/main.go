package main

import (
	"os"

	"github.com/GoogleCloudPlatform/guest-test-infra/container_images/registry-image-forked/commands"
	"github.com/fatih/color"
	"github.com/sirupsen/logrus"
)

func main() {
	color.NoColor = false

	command := commands.NewCheck(
		os.Stdin,
		os.Stderr,
		os.Stdout,
		os.Args,
	)

	err := command.Execute()
	if err != nil {
		logrus.Errorf("%s", err)
		os.Exit(1)
	}
}
