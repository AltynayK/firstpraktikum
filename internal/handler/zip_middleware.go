package handler

import (
	"compress/gzip"
	"crypto/aes"
	"crypto/rand"
	"fmt"
	"io"
	"net/http"
	"strings"

	uuid "github.com/satori/go.uuid"
)

type gzipBodyWriter struct {
	http.ResponseWriter
	writer io.Writer
}

func (gz gzipBodyWriter) Write(b []byte) (int, error) {
	return gz.writer.Write(b)
}

func CompressGzip(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			next.ServeHTTP(w, r)
			return
		}

		gz, err := gzip.NewWriterLevel(w, gzip.BestCompression)
		if err != nil {
			http.Error(w, "", http.StatusInternalServerError)
			return
		}
		defer gz.Close()
		w.Header().Set("Content-Encoding", "gzip")
		//w.Header().Set("Vary", "Accept-Encoding")
		w.Header().Del("Content-Length")
		next.ServeHTTP(gzipBodyWriter{
			ResponseWriter: w,
			writer:         gz,
		}, r)
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
