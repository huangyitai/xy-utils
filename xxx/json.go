package xxx

import (
	"encoding/json"
	jsoniter "github.com/json-iterator/go"
	"github.com/rs/zerolog/log"
)

// JsonMarshaller ...
var JsonMarshaller marshaller = jsoniter.ConfigDefault

// JsonMarshal ...
var JsonMarshal = jsoniter.ConfigDefault.Marshal

// JsonUnmarshal ...
var JsonUnmarshal = jsoniter.ConfigDefault.Unmarshal

// marshaller ...
type marshaller interface {
	Marshal(v interface{}) ([]byte, error)
}

// ToJSONBytes ...
func ToJSONBytes(v interface{}) []byte {
	bs, err := JsonMarshaller.Marshal(v)
	if err != nil {
		log.Err(err).Send()
	}
	return bs
}

// ToJSONStr ...
func ToJSONStr(v interface{}) string {
	bs, err := JsonMarshaller.Marshal(v)
	if err != nil {
		log.Err(err).Send()
	}
	return string(bs)
}

// ToJSONRaw ...
func ToJSONRaw(v interface{}) json.RawMessage {
	return ToJSONBytes(v)
}

func errorsToInterfaces(errs []error) []interface{} {
	var a []interface{}
	for _, err := range errs {
		if err == nil {
			a = append(a, nil)
			continue
		}
		a = append(a, err.Error())
	}
	return a
}

// ErrorsToJSONBytes ...
func ErrorsToJSONBytes(errs []error) []byte {
	return ToJSONBytes(errorsToInterfaces(errs))
}

// ErrorsToJSONStr ...
func ErrorsToJSONStr(errs []error) string {
	return ToJSONStr(errorsToInterfaces(errs))
}

// ErrorsToJSONRaw ...
func ErrorsToJSONRaw(errs []error) json.RawMessage {
	return ToJSONRaw(errorsToInterfaces(errs))
}

// JSONCopy 使用json序列化和反序列化实现数据复制
func JSONCopy(src, dst interface{}) error {
	bs, err := JsonMarshal(src)
	if err != nil {
		return err
	}
	err = JsonUnmarshal(bs, dst)
	if err != nil {
		return err
	}
	return nil
}
