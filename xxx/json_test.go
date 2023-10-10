package xxx

import (
	"fmt"
	"testing"

	"github.com/huangyitai/xy-utils/testx"
)

func TestToJSONBytes(t *testing.T) {
	m := map[string][]byte{}
	m["123"] = []byte("sssss")
	m["456"] = []byte("s123")
	println(ToJSONStr(m))

	es := []error{fmt.Errorf("123 error"), fmt.Errorf("23123 error"), fmt.Errorf("1111, errr"), nil}
	println(ErrorsToJSONStr(es))

}

func TestJSONCopy(t *testing.T) {
	type Apple struct {
		Name  string
		Price int
	}

	type args struct {
		src interface{}
		dst interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "1",
			args: args{
				src: map[string]interface{}{"a": map[string]interface{}{"Name": "Amy", "Price": 1}},
				dst: &map[string]*Apple{},
			},
			wantErr: false,
		},
		{
			name: "2",
			args: args{
				src: map[string]interface{}{"a": map[string]interface{}{"Name": "Amy", "Price": "1"}},
				dst: &map[string]*Apple{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := JSONCopy(tt.args.src, tt.args.dst); (err != nil) != tt.wantErr {
				t.Errorf("JSONCopy() error = %v, wantErr %v", err, tt.wantErr)
			}
			testx.PrintJSONPretty(tt.args.dst)
		})
	}
}
