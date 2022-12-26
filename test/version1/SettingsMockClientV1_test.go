package test_version1

import (
	"testing"

	"github.com/pip-services-infrastructure2/client-settings-go/version1"
)

type settingsMockClientV1Test struct {
	client  *version1.SettingsMockClientV1
	fixture *SettingsClientFixtureV1
}

func newSettingsMockClientV1Test() *settingsMockClientV1Test {
	return &settingsMockClientV1Test{}
}

func (c *settingsMockClientV1Test) setup(t *testing.T) *SettingsClientFixtureV1 {
	c.client = version1.NewSettingsMockClientV1()

	c.fixture = NewSettingsClientFixtureV1(c.client)

	return c.fixture
}

func (c *settingsMockClientV1Test) teardown(t *testing.T) {
	c.client = nil
}

func TestMockCrudOperations(t *testing.T) {
	c := newSettingsMockClientV1Test()
	fixture := c.setup(t)
	defer c.teardown(t)

	fixture.TestCrudOperations(t)
}
