package config

import (
	"gopkg.in/yaml.v3"
)

type PermissionsConfig struct {
	Permissions []PermissionConfig
	Roles       []RoleConfig
	Groups      []GroupConfig
}

func NewPermissionsConfig(str string) (PermissionsConfig, error) {
	permissionsConfig := PermissionsConfig{}
	err := yaml.Unmarshal([]byte(str), &permissionsConfig)
	return permissionsConfig, err
}

