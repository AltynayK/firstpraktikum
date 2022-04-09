package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPost(t *testing.T) {

	type want struct {
		code int

		contentType string
	}
	tests := []struct {
		name string
		want want
	}{
		{
			name: "negative test #1",
			want: want{
				code: 400,

				//Location: "plain/text",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodGet, "/{id}", nil)

			// создаём новый Recorder
			w := httptest.NewRecorder()
			// определяем хендлер
			h := http.HandlerFunc(Get)
			// запускаем сервер
			h.ServeHTTP(w, request)
			res := w.Result()

			// проверяем код ответа
			if res.StatusCode != tt.want.code {
				t.Errorf("Expected status code %d, got %d", tt.want.code, w.Code)
			}

			// заголовок ответа
			// if res.Header.Get("Content-Type") != tt.want.contentType {
			// 	t.Errorf("Expected Content-Type %s, got %s", tt.want.contentType, res.Header.Get("Content-Type"))
			// }
		})
	}
}
