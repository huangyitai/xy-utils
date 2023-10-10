package confx

import (
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"reflect"
)

// UnmarshalWithPath ...
func (r ReadFunc) UnmarshalWithPath(key string, ptr interface{}, format, path string,
	opts ...viper.DecoderConfigOption) error {

	bs, err := r(key)
	if err != nil {
		return err
	}
	return UnmarshalAnyWithPath(bs, ptr, format, path, opts...)
}

// UnmarshalJSONWithPath ...
func (r ReadFunc) UnmarshalJSONWithPath(key string, ptr interface{}, path string) error {
	return r.UnmarshalWithPath(key, ptr, JSON, path)
}

// UnmarshalYAMLWithPath ...
func (r ReadFunc) UnmarshalYAMLWithPath(key string, ptr interface{}, path string) error {
	return r.UnmarshalWithPath(key, ptr, YAML, path)
}

// UnmarshalTOMLWithPath ...
func (r ReadFunc) UnmarshalTOMLWithPath(key string, ptr interface{}, path string) error {
	return r.UnmarshalWithPath(key, ptr, TOML, path)
}

// ReadValueWithPath ...
func (r ReadFunc) ReadValueWithPath(key string, value reflect.Value, format string, path string,
	opts ...viper.DecoderConfigOption) error {
	bs, err := r(key)
	if err != nil {
		return errors.WithStack(err)
	}

	log.Trace().
		Bytes("sBytes", bs).Str("sKey", key).Str("sFormat", format).
		Msg("[confx]read value")

	return ReadValueWithPath(bs, value, format, path, opts...)
}
