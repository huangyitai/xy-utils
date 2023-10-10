package testx

import (
	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigDefault

// PrintJSONPretty ...
func PrintJSONPretty(i interface{}) {
	bs, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		panic(err)
	}
	println(string(bs))
}
