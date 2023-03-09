package onepassword

import (
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"os/exec"

	"github.com/charmbracelet/log"
	"github.com/docker/docker-credential-helpers/credentials"
	"github.com/flexwie/docker-credentials-1password/pkg/config"
)

type Onepassword struct {
	Log    log.Logger
	Config config.Config
}

func (o Onepassword) Add(creds *credentials.Credentials) error {
	logger := o.Log.With(log.WithPrefix("ADD"))

	if err := evalOp(); err != nil {
		return errors.New("could not find op cli")
	}

	json_content, err := json.Marshal(creds)
	if err != nil {
		return err
	}
	logger.Debug("unmarshaled", "content", string(json_content))

	enc := base64.StdEncoding.EncodeToString(json_content)
	logger.Debug("encoded", "content", enc)

	urlHash, _ := hashUrl(creds.ServerURL)
	logger.Debug("created url hash", "content", urlHash)

	path := fmt.Sprintf("credentials.%s[password]=%s", urlHash, enc)
	logger.Debug("created path", "content", path)

	out, err := exec.Command("op", "item", "edit", "Docker", path).CombinedOutput()
	if err != nil {
		return errors.New(string(out))
	}

	return nil
}

func (o Onepassword) Delete(serverUrl string) error {
	logger := o.Log.With(log.WithPrefix("ADD"))

	if err := evalOp(); err != nil {
		return errors.New("could not find op cli")
	}

	urlHash, _ := hashUrl(serverUrl)
	logger.Debug("created url hash", "content", urlHash)

	cmd := exec.Command("op", "item", "edit", "Docker", fmt.Sprintf("credentials.%s[delete]", urlHash))
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

func (o Onepassword) Get(serverUrl string) (string, string, error) {
	logger := o.Log.With(log.WithPrefix("ADD"))

	if err := evalOp(); err != nil {
		return "", "", errors.New("could not find op cli")
	}

	urlHash, _ := hashUrl(serverUrl)
	logger.Debug("created url hash", "content", urlHash)

	path := fmt.Sprintf("credentials.%s", urlHash)
	logger.Debug("created path", "content", path)

	out, err := exec.Command("op", "item", "get", "Docker", "--fields", path).CombinedOutput()
	if err != nil {
		return "", "", errors.New(string(out))
	}

	dec := make([]byte, base64.StdEncoding.DecodedLen(len(out)))
	n, err := base64.StdEncoding.Decode(dec, out)
	if err != nil {
		return "", "", err
	}
	dec = dec[:n]

	logger.Debug("decoded base64", "content", string(dec))

	var result *credentials.Credentials
	err = json.Unmarshal(dec, &result)
	if err != nil {
		return "", "", err
	}

	logger.Debug("unmarshaled json", "content", result)

	return result.Username, result.Secret, nil
}

func (o Onepassword) List() (map[string]string, error) {
	return nil, errors.New("method not implemented")
}

func evalOp() error {
	cmd := exec.Command("op", "--version")
	return cmd.Run()
}

func hashUrl(url string) (string, error) {
	h := sha1.New()
	h.Write([]byte(url))

	sha := h.Sum(nil)
	return hex.EncodeToString(sha), nil
}
