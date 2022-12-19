package version1

import (
	"time"

	"github.com/pip-services3-gox/pip-services3-commons-gox/config"
)

type SettingsSectionV1 struct {
	Id         string               `json:"id"`
	Parameters *config.ConfigParams `json:"parameters"`
	UpdateTime time.Time            `json:"update_time"`
}

func EmptySettingsSectionV1() *SettingsSectionV1 {
	return &SettingsSectionV1{}
}

func NewSettingsSectionV1(id string, parameters config.ConfigParams) *SettingsSectionV1 {
	return &SettingsSectionV1{
		Id:         id,
		Parameters: config.NewConfigParamsFromMaps(parameters.Value()),
		UpdateTime: time.Now(),
	}
}
