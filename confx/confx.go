package confx

import (
	jsoniter "github.com/json-iterator/go"
)

// YAML ...
const (
	YAML = "yaml"

	JSON = "json"

	TOML = "toml"

	INI = "ini"

	HCL = "hcl"

	PROPERTIES = "properties"

	ENV = "env"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary
