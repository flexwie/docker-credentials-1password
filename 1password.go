package onepassword

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"os/exec"

	"github.com/docker/docker-credential-helpers/credentials"
)

type Onepassword struct{}

func (o Onepassword) Add(creds *credentials.Credentials) error {
	if err := evalOp(); err != nil {
		return errors.New("could not find op cli")
	}

	json_content, err := json.Marshal(creds)
	if err != nil {
		return err
	}

	enc := base64.StdEncoding.EncodeToString(json_content)

	cmd := exec.Command("op", "item", "edit", "'Docker'", fmt.Sprintf("'credentials.%s[password]=%s'", creds.ServerURL, enc))
	if err = cmd.Run(); err != nil {
		return err
	}

	return nil
}

func (o Onepassword) Delete(serverUrl string) error {
	if err := evalOp(); err != nil {
		return errors.New("could not find op cli")
	}

	return nil
}

func (o Onepassword) Get(serverUrl string) (string, string, error) {
	if err := evalOp(); err != nil {
		return "", "", errors.New("could not find op cli")
	}

	return "", "", nil
}

func (o Onepassword) List() (map[string]string, error) {
	if err := evalOp(); err != nil {
		return nil, errors.New("could not find op cli")
	}

	return nil, nil
}

func evalOp() error {
	cmd := exec.Command("op", "--version")
	return cmd.Run()
}
