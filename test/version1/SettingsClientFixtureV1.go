package test_version1

import (
	"context"
	"testing"

	"github.com/pip-services-infrastructure2/client-settings-go/version1"
	"github.com/pip-services3-gox/pip-services3-commons-gox/config"
	"github.com/stretchr/testify/assert"
)

type SettingsClientFixtureV1 struct {
	Client version1.ISettingsClientV1
}

func NewSettingsClientFixtureV1(client version1.ISettingsClientV1) *SettingsClientFixtureV1 {
	return &SettingsClientFixtureV1{
		Client: client,
	}
}

func (c *SettingsClientFixtureV1) clear() {
	page, _ := c.Client.GetSections(context.Background(), "", nil, nil)

	for _, v := range page.Data {
		section := v
		c.Client.SetSection(context.Background(), "", section.Id, config.NewEmptyConfigParams())
	}
}

func (c *SettingsClientFixtureV1) TestCrudOperations(t *testing.T) {
	c.clear()
	defer c.clear()

	// Create one section
	params, err := c.Client.SetSection(context.Background(), "", "test.1", config.NewConfigParamsFromTuples(
		"key1", "value11",
		"key2", "value12",
	))
	assert.Nil(t, err)

	assert.NotNil(t, params)
	assert.Equal(t, "value11", params.GetAsString("key1"))

	// Create another section
	params, err = c.Client.ModifySection(context.Background(), "", "test.2", config.NewConfigParamsFromTuples(
		"key1", "value21",
	), config.NewConfigParamsFromTuples(
		"key2", 1,
	))
	assert.Nil(t, err)

	assert.NotNil(t, params)
	assert.Equal(t, "value21", params.GetAsString("key1"))
	assert.Equal(t, "1", params.GetAsString("key2"))

	// Get second section
	params, err = c.Client.GetSectionById(context.Background(), "", "test.2")
	assert.Nil(t, err)

	assert.NotNil(t, params)
	assert.Equal(t, "value21", params.GetAsString("key1"))
	assert.Equal(t, "1", params.GetAsString("key2"))

	// Get all sections
	page, err1 := c.Client.GetSections(context.Background(), "", nil, nil)
	assert.Nil(t, err1)

	assert.NotNil(t, page)
	assert.True(t, len(page.Data) >= 2)

	// Get all section ids
	pageIds, err1 := c.Client.GetSectionIds(context.Background(), "", nil, nil)
	assert.Nil(t, err1)

	assert.NotNil(t, pageIds)
	assert.True(t, len(pageIds.Data) >= 2)
}
