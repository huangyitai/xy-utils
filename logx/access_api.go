package logx

import "github.com/huangyitai/xy-utils/xxx"

// ToJSONStr 将对象转换为日志中出现的json串，会按照Payload长度限制进行截断
func ToJSONStr(v interface{}) string {
	return xxx.JSONOrProtojsonTruncateToJSONStr(v, MaxLenPayloadPrefix, MaxLenPayloadSuffix)
}

// TruncateBytes 将字节数组按Payload长度限制截断
func TruncateBytes(bs []byte) []byte {
	return xxx.TruncateBytes(bs, MaxLenPayloadPrefix, MaxLenPayloadSuffix)
}
