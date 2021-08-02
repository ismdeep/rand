package rand

import (
	"fmt"
	"testing"
)

func TestIntBetween(t *testing.T) {
	type args struct {
		left  int
		right int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "",
			args: args{
				left:  1,
				right: 3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println(IntBetween(tt.args.left, tt.args.right))
		})
	}
}

func TestInt(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println(Int())
		})
	}
}
