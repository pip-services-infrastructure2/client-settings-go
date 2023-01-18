package version1

import (
	"context"
	"reflect"
	"time"

	"github.com/pip-services3-gox/pip-services3-commons-gox/config"
	"github.com/pip-services3-gox/pip-services3-commons-gox/data"
)

type SettingsMockClientV1 struct {
	settings    []*SettingsSectionV1
	maxPageSize int
	proto       reflect.Type
}

func NewSettingsMockClientV1() *SettingsMockClientV1 {
	return &SettingsMockClientV1{
		settings:    make([]*SettingsSectionV1, 0),
		maxPageSize: 100,
		proto:       reflect.TypeOf(SettingsSectionV1{}),
	}
}

func (c *SettingsMockClientV1) GetSectionIds(ctx context.Context, correlationId string, filter *data.FilterParams, paging *data.PagingParams) (result data.DataPage[string], err error) {
	items := make([]string, 0)
	for _, v := range c.settings {
		item := v
		items = append(items, item.Id)
	}

	return *data.NewDataPage(items, len(c.settings)), nil
}

func (c *SettingsMockClientV1) GetSections(ctx context.Context, correlationId string, filter *data.FilterParams, paging *data.PagingParams) (result data.DataPage[*SettingsSectionV1], err error) {
	items := make([]*SettingsSectionV1, 0)
	for _, v := range c.settings {
		item := v
		items = append(items, item)
	}
	return *data.NewDataPage(items, len(c.settings)), nil
}

func (c *SettingsMockClientV1) GetSectionById(ctx context.Context, correlationId string, id string) (result *config.ConfigParams, err error) {
	for _, v := range c.settings {
		if v.Id == id {
			result = config.NewConfigParams(v.Parameters)
			break
		}
	}

	if result == nil {
		result = config.NewEmptyConfigParams()
	}

	return result, nil
}

func (c *SettingsMockClientV1) SetSection(ctx context.Context, correlationId string, id string, params *config.ConfigParams) (result *config.ConfigParams, err error) {
	item := NewSettingsSectionV1(id, *params)

	c.settings = append(c.settings, item)
	return config.NewConfigParams(item.Parameters), nil
}

func (c *SettingsMockClientV1) ModifySection(ctx context.Context, correlationId string, id string, updateParams *config.ConfigParams, incrementParams *config.ConfigParams) (result *config.ConfigParams, err error) {
	var setting *SettingsSectionV1

	index := -1

	for i, item := range c.settings {
		if item.Id == id {
			index = i
			break
		}
	}

	if index >= 0 {
		setting = c.settings[index]
	} else {
		setting = NewSettingsSectionV1(id, *config.NewEmptyConfigParams())
	}

	params := config.NewConfigParams(setting.Parameters)

	// Update parameters
	if updateParams != nil {
		for _, key := range updateParams.Keys() {
			if val, ok := updateParams.Get(key); ok {
				params.SetAsObject(key, val)
			}
		}
	}

	// Increment parameters
	if incrementParams != nil {
		for _, key := range incrementParams.Keys() {
			if _, ok := incrementParams.Get(key); ok {
				increment := incrementParams.GetAsLongWithDefault(key, 0)
				value := params.GetAsLongWithDefault(key, 0)
				value = value + increment
				params.SetAsObject(key, value)
			}
		}
	}

	// Update time
	setting.UpdateTime = time.Now()

	setting.Parameters = params.Value()

	if index < 0 {
		c.settings = append(c.settings, setting)
	}

	return params, nil
}
