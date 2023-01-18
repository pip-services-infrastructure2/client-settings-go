package version1

import (
	"context"

	cconf "github.com/pip-services3-gox/pip-services3-commons-gox/config"
	cdata "github.com/pip-services3-gox/pip-services3-commons-gox/data"
	"github.com/pip-services3-gox/pip-services3-grpc-gox/clients"
)

type SettingsCommandableGrpcClientV1 struct {
	*clients.CommandableGrpcClient
}

func NewSettingsCommandableGrpcClientV1() *SettingsCommandableGrpcClientV1 {
	return NewSettingsCommandableGrpcClientV1WithConfig(nil)
}

func NewSettingsCommandableGrpcClientV1WithConfig(config *cconf.ConfigParams) *SettingsCommandableGrpcClientV1 {
	c := &SettingsCommandableGrpcClientV1{
		CommandableGrpcClient: clients.NewCommandableGrpcClient("v1/settings"),
	}

	if config != nil {
		c.Configure(context.Background(), config)
	}

	return c
}

func (c *SettingsCommandableGrpcClientV1) GetSectionIds(ctx context.Context, correlationId string, filter *cdata.FilterParams, paging *cdata.PagingParams) (result cdata.DataPage[string], err error) {
	res, err := c.CallCommand(ctx, "get_section_ids", correlationId, cdata.NewAnyValueMapFromTuples(
		"filter", filter,
		"paging", paging,
	))

	if err != nil {
		return *cdata.NewEmptyDataPage[string](), err
	}

	return clients.HandleHttpResponse[cdata.DataPage[string]](res, correlationId)
}

func (c *SettingsCommandableGrpcClientV1) GetSections(ctx context.Context, correlationId string, filter *cdata.FilterParams, paging *cdata.PagingParams) (result cdata.DataPage[*SettingsSectionV1], err error) {
	res, err := c.CallCommand(ctx, "get_sections", correlationId, cdata.NewAnyValueMapFromTuples(
		"filter", filter,
		"paging", paging,
	))

	if err != nil {
		return *cdata.NewEmptyDataPage[*SettingsSectionV1](), err
	}

	page, convErr := clients.HandleHttpResponse[cdata.DataPage[*SettingsSectionV1]](res, correlationId)

	if convErr != nil {
		return page, convErr
	}

	return page, nil
}

func (c *SettingsCommandableGrpcClientV1) GetSectionById(ctx context.Context, correlationId string, id string) (result *cconf.ConfigParams, err error) {
	res, err := c.CallCommand(ctx, "get_section_by_id", correlationId, cdata.NewAnyValueMapFromTuples(
		"id", id,
	))

	if err != nil {
		return nil, err
	}

	mapParams, convErr := clients.HandleHttpResponse[map[string]any](res, correlationId)
	parameters := cconf.NewConfigParamsFromValue(mapParams)

	if convErr != nil {
		return parameters, convErr
	}

	return parameters, convErr
}

func (c *SettingsCommandableGrpcClientV1) SetSection(ctx context.Context, correlationId string, id string, params *cconf.ConfigParams) (result *cconf.ConfigParams, err error) {
	res, err := c.CallCommand(ctx, "set_section", correlationId, cdata.NewAnyValueMapFromTuples(
		"id", id,
		"parameters", params.Value(),
	))

	if err != nil {
		return nil, err
	}

	mapParams, convErr := clients.HandleHttpResponse[map[string]any](res, correlationId)
	parameters := cconf.NewConfigParamsFromValue(mapParams)

	if convErr != nil {
		return parameters, convErr
	}

	return parameters, convErr
}

func (c *SettingsCommandableGrpcClientV1) ModifySection(ctx context.Context, correlationId string, id string, updateParams *cconf.ConfigParams, incrementParams *cconf.ConfigParams) (result *cconf.ConfigParams, err error) {
	res, err := c.CallCommand(ctx, "modify_section", correlationId, cdata.NewAnyValueMapFromTuples(
		"id", id,
		"update_parameters", updateParams.Value(),
		"increment_parameters", incrementParams.Value(),
	))

	if err != nil {
		return nil, err
	}

	mapParams, convErr := clients.HandleHttpResponse[map[string]any](res, correlationId)
	parameters := cconf.NewConfigParamsFromValue(mapParams)

	if convErr != nil {
		return parameters, convErr
	}

	return parameters, convErr
}
