package app

import (
	"flag"
	"os"
)

type Config struct {
	ServerAddress   string
	BaseURL         string
	FileStoragePath string
	DatabaseDNS     string
}

func NewConfig() *Config {
	ServerAddress := os.Getenv("SERVER_ADDRESS")
	BaseURL := os.Getenv("BASE_URL")
	FileStoragePath := os.Getenv("FILE_STORAGE_PATH")
	DatabaseDNS := os.Getenv("DatabaseDNS")

	flag.StringVar(&ServerAddress, "a", "127.0.0.1:8080", "ServerAddress - адрес запуска HTTP-сервера")
	flag.StringVar(&BaseURL, "b", "http://"+ServerAddress, "BaseURL")
	//increment#
	flag.StringVar(&FileStoragePath, "f", "texts.txt", "FileStoragePath - путь до файла LongURL")
	//flag.StringVar(&DatabaseDNS, "d", "host=localhost port=5432 user=altynay password=password dbname=somedb sslmode=disable", "DatabaseDNS")
	flag.StringVar(&DatabaseDNS, "d", "", "DatabaseDNS")
	flag.Parse()
	return &Config{
		ServerAddress:   ServerAddress,
		BaseURL:         BaseURL,
		FileStoragePath: FileStoragePath,
		DatabaseDNS:     DatabaseDNS,
	}
}
