package xxx

import (
	"reflect"
	"testing"
)

func TestUnsafeToBytes(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{"空串", args{""}, nil},
		{"普通字符串", args{"123123123"}, []byte("123123123")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UnsafeToBytes(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UnsafeToBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnsafeToString(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "空切片", args: args{nil}, want: ""},
		{"普通切片", args{[]byte("123123123")}, "123123123"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UnsafeToString(tt.args.b); got != tt.want {
				t.Errorf("UnsafeToString() = %v, want %v", got, tt.want)
			}
		})
	}
}
