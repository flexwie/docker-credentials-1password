package main

import (
	"github.com/docker/docker-credential-helpers/credentials"
	onepassword "github.com/flexwie/docker-credentials-1password"
)

func main() {
	credentials.Serve(onepassword.Onepassword{})
}
