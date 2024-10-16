package model

import "yoomall/core"

type SystemConfig struct {
	*core.Model
	Key     string `gorm:"column:key;type:varchar(255)" json:"key"`
	Value   string `gorm:"column:value;type:text" json:"value"`
	Group   string `gorm:"column:group;type:varchar(255)" json:"group"`
	Comment string `gorm:"column:comment;type:text" json:"comment"`
}

// TableName implements core.IModel.
func (s *SystemConfig) TableName() string {
	return "system_config"
}

type SystemConfigGroup struct {
	*core.Model
	Name        string `gorm:"column:name;type:varchar(255)" json:"name"`
	Description string `gorm:"column:description;type:text" json:"description"`
	Key         string `gorm:"column:key;type:varchar(255)" json:"key"`
}

// TableName implements core.IModel.
func (s *SystemConfigGroup) TableName() string {
	return "system_config_group"
}

var _ core.IModel = (*SystemConfig)(nil)

var _ core.IModel = (*SystemConfigGroup)(nil)
