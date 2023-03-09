package config

import (
	"encoding/json"
	"os"
	"testing"
)

func TestRead(t *testing.T) {
	// arrange
	path := "./test.json"
	os.Setenv("CRED_FILE", path)
	config := &Config{
		Docker: DockerConfig{
			Item:    "test_item",
			Section: "test_section",
		},
	}

	content, err := json.Marshal(config)
	if err != nil {
		t.Fatal(err)
	}
	if err = os.WriteFile("./test.json", content, 0777); err != nil {
		t.Fatal(err)
	}

	// act
	testConfig := &Config{}
	if err = testConfig.Read(); err != nil {
		t.Fatal(err)
	}

	// assert
	if testConfig.Docker.Item != config.Docker.Item {
		t.Fail()
	}

	// cleanup
	os.Remove(path)
	os.Unsetenv("CRED_FILE")
}
