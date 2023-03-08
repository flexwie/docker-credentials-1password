package onepassword

import (
	"testing"

	"github.com/charmbracelet/log"
	"github.com/docker/docker-credential-helpers/credentials"
)

func TestOpStore(t *testing.T) {
	creds := &credentials.Credentials{
		ServerURL: "https://test.cr.io:1234/abc1",
		Username:  "foobar",
		Secret:    "barfoo",
	}

	logger := log.New()
	logger.SetLevel(log.DebugLevel)

	store := Onepassword{Log: logger}

	if err := store.Add(creds); err != nil {
		t.Fatal(err)
	}

	username, secret, err := store.Get(creds.ServerURL)
	if err != nil {
		t.Fatal(err)
	}

	if username != creds.Username {
		t.Fatalf("expected %s, got %s", creds.Username, username)
	}
	if secret != creds.Secret {
		t.Fatalf("expected %s, got %s", creds.Secret, secret)
	}

	if err = store.Delete(creds.ServerURL); err != nil {
		t.Fatal(err)
	}
}
