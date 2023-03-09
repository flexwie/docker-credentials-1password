package main

import (
	"github.com/charmbracelet/log"
	"github.com/docker/docker-credential-helpers/credentials"
	onepassword "github.com/flexwie/docker-credentials-1password"
	"github.com/flexwie/docker-credentials-1password/pkg/config"
)

func main() {
	logger := log.New()
	config := config.Config{}
	if err := config.Read(); err != nil {
		logger.Fatal(err)
	}

	credentials.Serve(onepassword.Onepassword{Log: logger, Config: config})
}
