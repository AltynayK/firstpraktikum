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
	if u, f := os.LookupEnv("SERVER_ADDRESS"); f {
		ServerAddress = u
	}
	if u, f := os.LookupEnv("BASE_URL"); f {
		BaseURL = u
	}
	if u, flg := os.LookupEnv("FILE_STORAGE_PATH"); flg {
		FileStoragePath = u
	}
	if u, f := os.LookupEnv("DatabaseDNS"); f {

		DatabaseDNS = u
	}
	return &Config{
		ServerAddress:   ServerAddress,
		BaseURL:         BaseURL,
		FileStoragePath: FileStoragePath,
		DatabaseDNS:     DatabaseDNS,
	}
}
