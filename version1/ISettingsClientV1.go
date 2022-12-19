package version1

import (
	"context"

	"github.com/pip-services3-gox/pip-services3-commons-gox/config"
	"github.com/pip-services3-gox/pip-services3-commons-gox/data"
)

type ISettingsClientV1 interface {
	GetSectionIds(ctx context.Context, correlationId string, filter *data.FilterParams,
		paging *data.PagingParams) (result data.DataPage[string], err error)

	GetSections(ctx context.Context, correlationId string, filter *data.FilterParams,
		paging *data.PagingParams) (result data.DataPage[*SettingsSectionV1], err error)

	GetSectionById(ctx context.Context, correlationId string, id string) (result *config.ConfigParams, err error)

	SetSection(ctx context.Context, correlationId string, id string,
		params *config.ConfigParams) (result *config.ConfigParams, err error)

	ModifySection(ctx context.Context, correlationId string, id string, updateParams *config.ConfigParams,
		incrementParams *config.ConfigParams) (result *config.ConfigParams, err error)
}
