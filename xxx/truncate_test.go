package xxx

import (
	"reflect"
	"testing"
)

func TestTruncateToMultiLine(t *testing.T) {
	type args struct {
		str       string
		prefixLen int
		suffixLen int
	}
	var tests = []struct {
		name string
		args args
		want string
	}{
		{
			name: "case1",
			args: args{
				str:       "1234567890",
				prefixLen: 3,
				suffixLen: 3,
			},
			want: "123\n…\n890",
		},
		{
			name: "case2",
			args: args{
				str:       "1234567890",
				prefixLen: 3,
				suffixLen: 0,
			},
			want: "123\n…",
		},
		{
			name: "case3",
			args: args{
				str:       "1234567890",
				prefixLen: 0,
				suffixLen: 3,
			},
			want: "…\n890",
		},
		{
			name: "case4",
			args: args{
				str:       "1234567890",
				prefixLen: 5,
				suffixLen: 5,
			},
			want: "1234567890",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TruncateToMultiLine(tt.args.str, tt.args.prefixLen, tt.args.suffixLen); got != tt.want {
				t.Errorf("TruncateToMultiLine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTruncateBytes(t *testing.T) {
	type args struct {
		bs        []byte
		prefixLen int
		suffixLen int
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "case1",
			args: args{
				bs:        []byte("1234567890"),
				prefixLen: 7,
				suffixLen: 0,
			},
			want: []byte("1234567..."),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TruncateBytes(tt.args.bs, tt.args.prefixLen, tt.args.suffixLen); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TruncateBytes() = %v, want %v", string(got), string(tt.want))
			}
		})
	}
}
