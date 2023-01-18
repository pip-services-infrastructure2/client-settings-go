package version1

import (
	"encoding/json"

	"github.com/pip-services-infrastructure2/client-settings-go/protos"
	"github.com/pip-services3-gox/pip-services3-commons-gox/convert"
	"github.com/pip-services3-gox/pip-services3-commons-gox/data"
	"github.com/pip-services3-gox/pip-services3-commons-gox/errors"
)

func fromError(err error) *protos.ErrorDescription {
	if err == nil {
		return nil
	}

	desc := errors.ErrorDescriptionFactory.Create(err)
	obj := &protos.ErrorDescription{
		Type:          desc.Type,
		Category:      desc.Category,
		Code:          desc.Code,
		CorrelationId: desc.CorrelationId,
		Status:        convert.StringConverter.ToString(desc.Status),
		Message:       desc.Message,
		Cause:         desc.Cause,
		StackTrace:    desc.StackTrace,
		Details:       fromMap(desc.Details),
	}

	return obj
}

func toError(obj *protos.ErrorDescription) error {
	if obj == nil || (obj.Category == "" && obj.Message == "") {
		return nil
	}

	description := &errors.ErrorDescription{
		Type:          obj.Type,
		Category:      obj.Category,
		Code:          obj.Code,
		CorrelationId: obj.CorrelationId,
		Status:        convert.IntegerConverter.ToInteger(obj.Status),
		Message:       obj.Message,
		Cause:         obj.Cause,
		StackTrace:    obj.StackTrace,
		Details:       toMap(obj.Details),
	}

	return errors.ApplicationErrorFactory.Create(description)
}

func fromMap(val map[string]any) map[string]string {
	r := map[string]string{}

	for k, v := range val {
		r[k] = convert.StringConverter.ToString(v)
	}

	return r
}

func toMap(val map[string]string) map[string]any {
	r := map[string]any{}

	for k, v := range val {
		r[k] = v
	}

	return r
}

func toJson(value any) string {
	if value == nil {
		return ""
	}

	b, err := json.Marshal(value)
	if err != nil {
		return ""
	}
	return string(b[:])
}

func fromJson(value string) any {
	if value == "" {
		return nil
	}

	var m any
	json.Unmarshal([]byte(value), &m)
	return m
}

func fromSettingsSection(section *SettingsSectionV1) *protos.SettingsSection {
	if section == nil {
		return nil
	}

	obj := &protos.SettingsSection{
		Id:         section.Id,
		Parameters: section.Parameters,
		UpdateTime: convert.StringConverter.ToString(section.UpdateTime),
	}

	return obj
}

func toSettingsSection(obj *protos.SettingsSection) *SettingsSectionV1 {
	if obj == nil {
		return nil
	}

	settings := &SettingsSectionV1{
		Id:         obj.Id,
		Parameters: obj.Parameters,
		UpdateTime: convert.DateTimeConverter.ToDateTime(obj.UpdateTime),
	}

	return settings
}

func fromSettingsSectionPage(page data.DataPage[*SettingsSectionV1]) *protos.SettingsSectionPage {

	obj := &protos.SettingsSectionPage{
		Total: int64(page.Total),
		Data:  make([]*protos.SettingsSection, len(page.Data)),
	}

	for i, v := range page.Data {
		section := v
		obj.Data[i] = fromSettingsSection(section)
	}

	return obj
}

func toSettingsSectionPage(obj *protos.SettingsSectionPage) data.DataPage[*SettingsSectionV1] {
	if obj == nil {
		return *data.NewEmptyDataPage[*SettingsSectionV1]()
	}

	sections := make([]*SettingsSectionV1, len(obj.Data))

	for i, v := range obj.Data {
		sections[i] = toSettingsSection(v)
	}

	page := *data.NewDataPage(sections, int(obj.Total))

	return page
}

func fromSettingsIdPage(page data.DataPage[string]) *protos.SettingsIdPage {
	obj := &protos.SettingsIdPage{
		Total: int64(page.Total),
		Data:  make([]string, len(page.Data)),
	}

	copy(obj.Data, page.Data)

	return obj
}

func toSettingsIdPage(obj *protos.SettingsIdPage) data.DataPage[string] {
	ids := make([]string, len(obj.Data))

	copy(ids, obj.Data)

	page := *data.NewDataPage(ids, int(obj.Total))

	return page
}
