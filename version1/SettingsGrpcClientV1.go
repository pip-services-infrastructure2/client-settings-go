package version1

import (
	"context"

	"github.com/pip-services3-gox/pip-services3-commons-gox/config"
	"github.com/pip-services3-gox/pip-services3-commons-gox/data"
	"github.com/pip-services3-gox/pip-services3-grpc-gox/clients"
	"github.com/service-infrastructure2/client-settings-go/protos"
)

type SettingsGrpcClientV1 struct {
	*clients.GrpcClient
}

func NewSettingsGrpcClientV1() *SettingsGrpcClientV1 {
	return &SettingsGrpcClientV1{
		GrpcClient: clients.NewGrpcClient("settings_v1.Settings"),
	}
}

func (c *SettingsGrpcClientV1) GetSectionIds(ctx context.Context, correlationId string, filter *data.FilterParams,
	paging *data.PagingParams) (result data.DataPage[string], err error) {
	timing := c.Instrument(ctx, correlationId, "settings_v1.get_section_ids")
	defer timing.EndTiming(ctx, err)

	req := &protos.SettingsPageRequest{
		CorrelationId: correlationId,
	}
	if filter != nil {
		req.Filter = filter.Value()
	}
	if paging != nil {
		req.Paging = &protos.PagingParams{
			Skip:  paging.GetSkip(0),
			Take:  (int32)(paging.GetTake(100)),
			Total: paging.Total,
		}
	}

	reply := new(protos.SettingsIdPageReply)
	err = c.CallWithContext(ctx, "get_section_ids", correlationId, req, reply)
	if err != nil {
		return *data.NewEmptyDataPage[string](), err
	}

	if reply.Error != nil {
		err = toError(reply.Error)
		return *data.NewEmptyDataPage[string](), err
	}

	result = toSettingsIdPage(reply.Page)

	return result, nil
}

func (c *SettingsGrpcClientV1) GetSections(ctx context.Context, correlationId string, filter *data.FilterParams,
	paging *data.PagingParams) (result data.DataPage[*SettingsSectionV1], err error) {
	timing := c.Instrument(ctx, correlationId, "settings_v1.get_sections")
	defer timing.EndTiming(ctx, err)

	req := &protos.SettingsPageRequest{
		CorrelationId: correlationId,
	}
	if filter != nil {
		req.Filter = filter.Value()
	}
	if paging != nil {
		req.Paging = &protos.PagingParams{
			Skip:  paging.GetSkip(0),
			Take:  (int32)(paging.GetTake(100)),
			Total: paging.Total,
		}
	}

	reply := new(protos.SettingsSectionPageReply)
	err = c.CallWithContext(ctx, "get_sections", correlationId, req, reply)
	if err != nil {
		return *data.NewEmptyDataPage[*SettingsSectionV1](), err
	}

	if reply.Error != nil {
		err = toError(reply.Error)
		return *data.NewEmptyDataPage[*SettingsSectionV1](), err
	}

	result = toSettingsSectionPage(reply.Page)

	return result, nil
}

func (c *SettingsGrpcClientV1) GetSectionById(ctx context.Context, correlationId string, id string) (result *config.ConfigParams, err error) {
	timing := c.Instrument(ctx, correlationId, "settings_v1.get_section_by_id")
	defer timing.EndTiming(ctx, err)

	req := &protos.SettingsIdRequest{
		CorrelationId: correlationId,
		Id:            id,
	}

	reply := new(protos.SettingsParamsReply)
	err = c.CallWithContext(ctx, "get_section_by_id", correlationId, req, reply)
	if err != nil {
		return nil, err
	}

	if reply.Error != nil {
		err = toError(reply.Error)
		return nil, err
	}

	result = config.NewConfigParamsFromMaps(reply.Parameters)

	return result, nil
}

func (c *SettingsGrpcClientV1) SetSection(ctx context.Context, correlationId string, id string, params *config.ConfigParams) (result *config.ConfigParams, err error) {
	timing := c.Instrument(ctx, correlationId, "settings_v1.set_section")
	defer timing.EndTiming(ctx, err)

	req := &protos.SettingsParamsRequest{
		CorrelationId: correlationId,
		Id:            id,
		Parameters:    params.Value(),
	}

	reply := new(protos.SettingsParamsReply)
	err = c.CallWithContext(ctx, "set_section", correlationId, req, reply)
	if err != nil {
		return nil, err
	}

	if reply.Error != nil {
		err = toError(reply.Error)
		return nil, err
	}

	result = config.NewConfigParamsFromMaps(reply.Parameters)

	return result, nil
}

func (c *SettingsGrpcClientV1) ModifySection(ctx context.Context, correlationId string, id string, updateParams *config.ConfigParams, incrementParams *config.ConfigParams) (result *config.ConfigParams, err error) {
	timing := c.Instrument(ctx, correlationId, "settings_v1.modify_section")
	defer timing.EndTiming(ctx, err)

	req := &protos.SettingsModifyParamsRequest{
		CorrelationId: correlationId,
		Id:            id,
	}

	if updateParams != nil {
		req.UpdateParameters = updateParams.Value()
	}
	if incrementParams != nil {
		req.IncrementParameters = incrementParams.Value()
	}

	reply := new(protos.SettingsParamsReply)
	err = c.CallWithContext(ctx, "modify_section", correlationId, req, reply)
	if err != nil {
		return nil, err
	}

	if reply.Error != nil {
		err = toError(reply.Error)
		return nil, err
	}

	result = config.NewConfigParamsFromMaps(reply.Parameters)

	return result, nil
}
