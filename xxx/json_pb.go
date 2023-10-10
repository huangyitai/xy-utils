package xxx

import (
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/proto"
)

// JSONOrProtojsonUnmarshal ...
func JSONOrProtojsonUnmarshal(data []byte, v interface{}) error {
	if m, ok := v.(proto.Message); ok {
		return ProtojsonUnmarshalOptions.Unmarshal(data, m)
	}
	return JsonUnmarshal(data, v)
}

// JSONOrProtojsonMarshal ...
func JSONOrProtojsonMarshal(v interface{}) ([]byte, error) {
	if m, ok := v.(proto.Message); ok {
		return ProtojsonOptions.Marshal(m)
	}
	return JsonMarshal(v)
}

// JSONOrProtojsonToJSONStr ...
func JSONOrProtojsonToJSONStr(v interface{}) string {
	bs, err := JSONOrProtojsonMarshal(v)
	if err != nil {
		log.Err(err).Caller(1).Send()
	}
	return UnsafeToString(bs)
}

// JSONOrProtojsonToJSONBytes ...
func JSONOrProtojsonToJSONBytes(v interface{}) []byte {
	bs, err := JSONOrProtojsonMarshal(v)
	if err != nil {
		log.Err(err).Caller(1).Send()
	}
	return bs
}
