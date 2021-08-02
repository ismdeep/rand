package rand

import (
	"fmt"
	"testing"
)

func TestHex(t *testing.T) {
	type args struct {
		len int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "",
			args: args{
				len: 10,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_ = Hex(tt.args.len)
		})
	}
}

func TestHexStr(t *testing.T) {
	type args struct {
		len int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "",
			args: args{
				len: 10,
			},
		},
		{
			name: "",
			args: args{
				len: 9,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println(HexStr(tt.args.len))
		})
	}
}
