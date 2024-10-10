package confx

import (
	"io/ioutil"
	"os"
	"reflect"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

// ReadFunc ...
type ReadFunc func(string) ([]byte, error)

// ReadFile ...
var ReadFile ReadFunc = ioutil.ReadFile

// ReadEnv ...
var ReadEnv ReadFunc = func(s string) ([]byte, error) {
	env := os.Getenv(s)
	return []byte(env), nil
}

var readFuncTable = map[string]ReadFunc{
	"file": ReadFile,
	"env":  ReadEnv,
}

// Register ...
func Register(name string, read ReadFunc) {
	readFuncTable[name] = read
}

// Use ...
func Use(read ReadFunc) ReadFunc {
	return read
}

// Unmarshal ...
func (r ReadFunc) Unmarshal(key string, ptr interface{}, format string, opts ...viper.DecoderConfigOption) error {
	bytes, err := r(key)
	if err != nil {
		return err
	}
	return UnmarshalAny(bytes, ptr, format, opts...)
}

// UnmarshalJSON ...
func (r ReadFunc) UnmarshalJSON(key string, ptr interface{}, opts ...viper.DecoderConfigOption) error {
	return r.Unmarshal(key, ptr, JSON, opts...)
}

// UnmarshalYAML ...
func (r ReadFunc) UnmarshalYAML(key string, ptr interface{}, opts ...viper.DecoderConfigOption) error {
	return r.Unmarshal(key, ptr, YAML, opts...)
}

// UnmarshalTOML ...
func (r ReadFunc) UnmarshalTOML(key string, ptr interface{}, opts ...viper.DecoderConfigOption) error {
	return r.Unmarshal(key, ptr, TOML, opts...)
}

// ReadValue ...
func (r ReadFunc) ReadValue(key string, value reflect.Value, format string, opts ...viper.DecoderConfigOption) error {
	//读取config字符串
	bytes, err := r(key)
	if err != nil {
		return errors.WithStack(err)
	}

	log.Trace().
		Bytes("sBytes", bytes).Str("sKey", key).Str("sFormat", format).
		Msg("[confx]read value")

	return ReadValue(bytes, value, format, opts...)
}
