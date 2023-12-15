package srvx

import (
	"os"
	"strings"

	"github.com/huangyitai/xy-utils/confx"
	"github.com/huangyitai/xy-utils/xxx"
)

var global Config

func init() {
	global.Service = os.Getenv("SERVICE_NAME")
	global.Instance = os.Getenv("HOSTNAME")
	global.Platform = os.Getenv("PLATFORM_NAME")
	global.Project = os.Getenv("PROJ_NAME")
	global.Environment = os.Getenv("DEPLOY_ENV")
	global.Binary = os.Getenv("BIN_NAME")
}

// GetService 获取service名称
func GetService() string {
	return global.Service
}

// GetInstance 获取实例名称
func GetInstance() string {
	return global.Instance
}

// GetPlatform 获取部署平台名称
func GetPlatform() string {
	return global.Platform
}

// GetProject 获取项目名称
func GetProject() string {
	return global.Project
}

// GetEnvironment 获取环境名称
func GetEnvironment() string {
	return global.Environment
}

// GetBinary 获取服务二进制文件名称
func GetBinary() string {
	return global.Binary
}

// IsDev 判断是否为开发环境
func IsDev() bool {
	return !IsPre() && !IsProd()
}

// IsPre 判断是否为预发布环境
func IsPre() bool {
	env := strings.ToLower(GetEnvironment())
	switch env {
	case "pre", "pre-release", "prerelease":
		return true
	default:
		return strings.HasPrefix(env, "pre")
	}
}

// IsProd 判断是否为正式环境
func IsProd() bool {
	env := strings.ToLower(GetEnvironment())
	switch env {
	case "prod", "production", "release":
		return true
	default:
		return strings.HasPrefix(env, "prod")
	}
}

// SetupMergeConfig ...
func SetupMergeConfig(cfg *Config) {
	if cfg == nil {
		return
	}
	if cfg.Service != "" {
		global.Service = cfg.Service
	}
	if cfg.Instance != "" {
		global.Instance = cfg.Instance
	}
	if cfg.Platform != "" {
		global.Platform = cfg.Platform
	}
	if cfg.Project != "" {
		global.Project = cfg.Project
	}
	if cfg.Environment != "" {
		global.Environment = cfg.Environment
	}
	if cfg.Binary != "" {
		global.Binary = cfg.Binary
	}
}

// SetupMergeStringWithPath ...
func SetupMergeStringWithPath(str, format, path string) error {
	cfg := new(Config)
	if path == "" {
		err := confx.UnmarshalAny([]byte(os.ExpandEnv(str)), cfg, format)
		if err != nil {
			return err
		}
	} else {
		err := confx.UnmarshalAnyWithPath([]byte(os.ExpandEnv(str)), cfg, format, path)
		if err != nil {
			return err
		}
	}
	SetupMergeConfig(cfg)
	return nil
}

// SetupMergeReadWithPath ...
func SetupMergeReadWithPath(read confx.ReadFunc, key, format, path string) error {
	bs, err := read(key)
	if err != nil {
		return err
	}
	return SetupMergeStringWithPath(xxx.UnsafeToString(bs), format, path)
}
