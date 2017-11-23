package utils_test

import (
	"github.com/penguinn/penguin/utils"
	"testing"
)

func TestISHan(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "han",
			args: args{
				str: "你好",
			},
			want: true,
		},
		{
			name: "number",
			args: args{
				str: "123456",
			},
			want: false,
		},
		{
			name: "letter",
			args: args{
				str: "abc",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := utils.ISHan(tt.args.str); got != tt.want {
				t.Errorf("ISHan() = %v, want %v", got, tt.want)
			}
		})
	}
}
