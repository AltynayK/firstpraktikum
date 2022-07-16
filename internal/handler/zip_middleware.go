package handler

import (
	"compress/gzip"
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"database/sql"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	uuid "github.com/satori/go.uuid"
)

type gzipResponseWriter struct {
	*gzip.Writer
	http.ResponseWriter
}

func (w *gzipResponseWriter) Header() http.Header {
	return w.ResponseWriter.Header()
}

func (w *gzipResponseWriter) Write(b []byte) (int, error) {
	h := w.ResponseWriter.Header()
	if h.Get("Content-Type") == "" {
		h.Set("Content-Type", http.DetectContentType(b))
	}
	return w.Writer.Write(b)
}

func GzipHandler(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			h.ServeHTTP(w, r)
			return
		}
		w.Header().Set("Content-Encoding", "gzip")
		w.Header().Set("Vary", "Accept-Encoding")
		gw := gzip.NewWriter(w)
		defer gw.Close()
		w = &gzipResponseWriter{gw, w}
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

func Decompress(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get(`Content-Encoding`) == `gzip` {
			gz, err := gzip.NewReader(r.Body)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			r.Body = gz
		}
		next.ServeHTTP(w, r)
	})
}

var Id uuid.UUID
var Key []byte
var userCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

func SetCookie(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session")
		Id = uuid.NewV4()
		if err != nil {
			cookie = &http.Cookie{
				Name:       "session",
				Value:      Id.String(),
				Path:       "",
				Domain:     "",
				Expires:    time.Time{},
				RawExpires: "",
				MaxAge:     0,
				Secure:     false,
				HttpOnly:   true,
				SameSite:   0,
				Raw:        "",
				Unparsed:   []string{},
			}
			// константа aes.BlockSize определяет размер блока и равна 16 байтам
			Key, err := generateRandom(aes.BlockSize) // ключ шифрования
			if err != nil {
				fmt.Printf("error: %v\n", err)
				return
			}
			encrypt([]byte(cookie.Value), Key)
			http.SetCookie(w, cookie)
		}
		ctx := context.WithValue(r.Context(), userCtxKey, cookie.Value)
		r = r.WithContext(ctx)
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

func CheckCookie(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session")
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
		}
		decrypt([]byte(cookie.Value), Key)
		if cookie.Value != r.Context().Value("session") {
			fmt.Print("sdfsdf")
		}
		next.ServeHTTP(w, r)
	})

}

func encrypt(plaintext []byte, key []byte) ([]byte, error) {
	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}
	return gcm.Seal(nonce, nonce, plaintext, nil), nil
}

func decrypt(ciphertext []byte, key []byte) ([]byte, error) {
	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}
	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, errors.New("ciphertext too short")
	}
	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	return gcm.Open(nil, nonce, ciphertext, nil)
}

func CreateTable(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		db, err := sql.Open("postgres", *DBdns)
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()
		_, err = db.Exec("CREATE TABLE IF NOT EXISTS data (id serial primary key, short_url varchar, original_url varchar, user_id varchar)")
		if err != nil {
			panic(err)
		}
		next.ServeHTTP(w, r)
	})
}
