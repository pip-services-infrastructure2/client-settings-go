package test_version1

import (
	"context"
	"os"
	"testing"

	"github.com/pip-services-infrastructure2/client-settings-go/version1"
	"github.com/pip-services3-gox/pip-services3-commons-gox/config"
)

type settingsCommandableHttpClientV1Test struct {
	client  *version1.SettingsCommandableHttpClientV1
	fixture *SettingsClientFixtureV1
}

func newSettingsCommandableHttpClientV1Test() *settingsCommandableHttpClientV1Test {
	return &settingsCommandableHttpClientV1Test{}
}

func (c *settingsCommandableHttpClientV1Test) setup(t *testing.T) *SettingsClientFixtureV1 {
	var HTTP_HOST = os.Getenv("HTTP_HOST")
	if HTTP_HOST == "" {
		HTTP_HOST = "localhost"
	}
	var HTTP_PORT = os.Getenv("HTTP_PORT")
	if HTTP_PORT == "" {
		HTTP_PORT = "8080"
	}

	var httpConfig = config.NewConfigParamsFromTuples(
		"connection.protocol", "http",
		"connection.host", HTTP_HOST,
		"connection.port", HTTP_PORT,
	)

	c.client = version1.NewSettingsCommandableHttpClientV1()
	c.client.Configure(context.Background(), httpConfig)
	c.client.Open(context.Background(), "")

	c.fixture = NewSettingsClientFixtureV1(c.client)

	return c.fixture
}

func (c *settingsCommandableHttpClientV1Test) teardown(t *testing.T) {
	c.client.Close(context.Background(), "")
}

func TestCommandableHttpCrudOperations(t *testing.T) {
	c := newSettingsCommandableHttpClientV1Test()
	fixture := c.setup(t)
	defer c.teardown(t)

	fixture.TestCrudOperations(t)
}
