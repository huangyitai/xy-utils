package xxx

import (
	"encoding/json"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ProtojsonOptions ...
var ProtojsonOptions = protojson.MarshalOptions{
	EmitUnpopulated: true,
	UseProtoNames:   true,
	UseEnumNumbers:  true,
}

// ProtojsonUnmarshalOptions ...
var ProtojsonUnmarshalOptions = protojson.UnmarshalOptions{}

// ProtoToJSONStr ...
func ProtoToJSONStr(m proto.Message) string {
	return ProtojsonOptions.Format(m)
}

// ProtoToJSONBytes ...
func ProtoToJSONBytes(m proto.Message) []byte {
	bs, err := ProtojsonOptions.Marshal(m)
	if err != nil {
		log.Err(err).Send()
	}
	return bs
}

// ProtoToJSONRaw ...
func ProtoToJSONRaw(m proto.Message) json.RawMessage {
	bs, err := ProtojsonOptions.Marshal(m)
	if err != nil {
		log.Err(err).Send()
	}
	return bs
}
