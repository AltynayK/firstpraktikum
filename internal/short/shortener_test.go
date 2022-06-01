package short

import "testing"

func TestWriteShortURL(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "simple test #1", // описывается каждый тест
			args: args{
				url: "https://practicum.yandex.ru/",
			}, // значения, которые будет принимать функция
			want: "http://127.0.0.1:8080/1", // ожидаемое значение
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WriteShortURL(tt.args.url); got != tt.want {
				t.Errorf("WriteShortURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
