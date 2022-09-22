package app

import (
	"flag"
)

type Config struct {
	ServerAddress   string
	BaseURL         string
	FileStoragePath string
	DatabaseDNS     string
}

var (
	ServerAddress   string
	BaseURL         string
	FileStoragePath string
	DatabaseDNS     string
)

func init() {
	//increment#5
	flag.StringVar(&ServerAddress, "a", "127.0.0.1:8080", "ServerAddress - адрес запуска HTTP-сервера")
	flag.StringVar(&BaseURL, "b", "http://"+ServerAddress, "BaseURL")
	//increment#
	flag.StringVar(&FileStoragePath, "f", "texts.txt", "FileStoragePath - путь до файла LongURL")
	//flag.StringVar(&DatabaseDNS, "d", "host=localhost port=5432 user=altynay password=password dbname=somedb sslmode=disable", "DatabaseDNS")
	flag.StringVar(&DatabaseDNS, "d", "", "DatabaseDNS")
}
func NewConfig() *Config {

	flag.Parse()
	return &Config{
		ServerAddress:   ServerAddress,
		BaseURL:         BaseURL,
		FileStoragePath: FileStoragePath,
		DatabaseDNS:     DatabaseDNS,
	}
}
