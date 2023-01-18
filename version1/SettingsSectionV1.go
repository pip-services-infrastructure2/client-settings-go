package version1

import (
	"time"

	"github.com/pip-services3-gox/pip-services3-commons-gox/config"
)

type SettingsSectionV1 struct {
	Id         string            `json:"id"`
	Parameters map[string]string `json:"parameters"`
	UpdateTime time.Time         `json:"update_time"`
}

func (c *SettingsSectionV1) GetParametersAsConfigParams() *config.ConfigParams {
	return config.NewConfigParamsFromValue(c.Parameters)
}

func (c *SettingsSectionV1) SetParametersAsConfigParams(params config.ConfigParams) {
	c.Parameters = params.Value()
}

func NewEmptySettingsSectionV1() *SettingsSectionV1 {
	return &SettingsSectionV1{}
}

func NewSettingsSectionV1(id string, parameters config.ConfigParams) *SettingsSectionV1 {
	return &SettingsSectionV1{
		Id:         id,
		Parameters: parameters.Value(),
		UpdateTime: time.Now(),
	}
}
