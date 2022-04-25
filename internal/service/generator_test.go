package service

import "testing"

func TestGetURLFromID(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "simple test #1", // описывается каждый тест
			args: args{
				id: 1,
			}, // значения, которые будет принимать функция
			want: "https://practicum.yandex.ru/", // ожидаемое значение
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetURLFromID(tt.args.id); got != tt.want {
				t.Errorf("GetURLFromID() = %v, want %v", got, tt.want)
			}
		})
	}
}
