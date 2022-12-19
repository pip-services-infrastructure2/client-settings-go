package version1

import (
	"context"

	"github.com/pip-services3-gox/pip-services3-commons-gox/config"
	"github.com/pip-services3-gox/pip-services3-commons-gox/data"
)

type SettingsNullClientV1 struct {
}

func NewSettingsNullClientV1() *SettingsNullClientV1 {
	return &SettingsNullClientV1{}
}

func (c *SettingsNullClientV1) GetSectionIds(ctx context.Context, correlationId string, filter *data.FilterParams, paging *data.PagingParams) (result data.DataPage[string], err error) {
	return *data.NewEmptyDataPage[string](), nil
}

func (c *SettingsNullClientV1) GetSections(ctx context.Context, correlationId string, filter *data.FilterParams, paging *data.PagingParams) (result data.DataPage[*SettingsSectionV1], err error) {
	return *data.NewEmptyDataPage[*SettingsSectionV1](), nil
}

func (c *SettingsNullClientV1) GetSectionById(ctx context.Context, correlationId string, id string) (result *config.ConfigParams, err error) {
	return config.NewEmptyConfigParams(), nil
}

func (c *SettingsNullClientV1) SetSection(ctx context.Context, correlationId string, id string, params *config.ConfigParams) (result *config.ConfigParams, err error) {
	return params, nil
}

func (c *SettingsNullClientV1) ModifySection(ctx context.Context, correlationId string, id string, updateParams *config.ConfigParams, incrementParams *config.ConfigParams) (result *config.ConfigParams, err error) {
	if updateParams == nil {
		updateParams = config.NewEmptyConfigParams()
	}

	if incrementParams == nil {
		incrementParams = config.NewEmptyConfigParams()
	}

	updateParams = updateParams.Override(incrementParams)
	return updateParams, nil
}
