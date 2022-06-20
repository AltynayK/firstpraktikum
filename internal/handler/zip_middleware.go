package handler

import (
	"compress/gzip"
	"crypto/aes"
	"crypto/rand"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"

	uuid "github.com/satori/go.uuid"
)

// type gzipBodyWriter struct {
// 	http.ResponseWriter
// 	writer io.Writer
// }

// func (gz gzipBodyWriter) Write(b []byte) (int, error) {
// 	return gz.writer.Write(b)
// }
var gzPool = sync.Pool{
	New: func() interface{} {
		w := gzip.NewWriter(ioutil.Discard)
		gzip.NewWriterLevel(w, gzip.BestCompression)
		return w
	},
}

type gzipResponseWriter struct {
	io.Writer
	http.ResponseWriter
}

func (w *gzipResponseWriter) WriteHeader(status int) {
	w.Header().Del("Content-Length")
	w.ResponseWriter.WriteHeader(status)
}

func (w *gzipResponseWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}

// Gzip func handler
func CompressGzip(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			next.ServeHTTP(w, r)
			return
		}

		w.Header().Set("Content-Encoding", "gzip")

		gz := gzPool.Get().(*gzip.Writer)
		defer gzPool.Put(gz)

		gz.Reset(w)
		defer gz.Close()

		next.ServeHTTP(&gzipResponseWriter{ResponseWriter: w, Writer: gz}, r)
	})
}
func Cookie(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session")
		id := uuid.NewV4()
		if err != nil {

			cookie = &http.Cookie{
				Name:     "session",
				Value:    id.String(),
				HttpOnly: true,
			}
			// константа aes.BlockSize определяет размер блока и равна 16 байтам
			key, err := generateRandom(aes.BlockSize) // ключ шифрования
			if err != nil {
				fmt.Printf("error: %v\n", err)
				return
			}

			// получаем cipher.Block
			aesblock, err := aes.NewCipher(key)
			if err != nil {
				fmt.Printf("error: %v\n", err)
				return
			}

			dst := make([]byte, aes.BlockSize) // зашифровываем
			aesblock.Encrypt(dst, []byte(cookie.Value))
			//fmt.Printf("encrypted: %x\n", dst)

			http.SetCookie(w, cookie)

		}

		next.ServeHTTP(w, r)
	})
}

func generateRandom(size int) ([]byte, error) {
	b := make([]byte, size)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func deShifr(dst []byte) {
	// константа aes.BlockSize определяет размер блока и равна 16 байтам
	key, err := generateRandom(aes.BlockSize) // ключ шифрования
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

	// получаем cipher.Block
	aesblock, err := aes.NewCipher(key)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	src2 := make([]byte, aes.BlockSize) // расшифровываем
	aesblock.Decrypt(src2, dst)
	//fmt.Printf("decrypted: %s\n", src2)
	//fmt.Println(cookie)
}
