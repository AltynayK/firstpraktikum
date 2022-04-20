package service

import "testing"

func TestWriteURLByID(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "simple test #1",
			args: args{
				url: "https://practicum.yandex.ru/",
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WriteURLByID(tt.args.url); got != tt.want {
				t.Errorf("WriteURLByID() = %v, want %v", got, tt.want)
			}
		})
	}
}
