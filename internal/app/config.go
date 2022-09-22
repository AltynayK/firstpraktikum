package app

import "flag"

type Config struct {
	ServerAddress   string
	BaseURL         string
	FileStoragePath string
	DatabaseDNS     string
}

var config *Config

func NewConfig() *Config {
	flag.StringVar(&config.ServerAddress, "a", "127.0.0.1:8080", "ServerAddress - адрес запуска HTTP-сервера")
	flag.StringVar(&config.BaseURL, "b", "http://"+config.ServerAddress, "BaseURL")
	//increment#
	flag.StringVar(&config.FileStoragePath, "f", "texts.txt", "FileStoragePath - путь до файла LongURL")
	//flag.StringVar(&DatabaseDNS, "d", "host=localhost port=5432 user=altynay password=password dbname=somedb sslmode=disable", "DatabaseDNS")
	flag.StringVar(&config.DatabaseDNS, "d", "", "DatabaseDNS")
	flag.Parse()
	return &Config{
		ServerAddress:   config.ServerAddress,
		BaseURL:         config.BaseURL,
		FileStoragePath: config.FileStoragePath,
		DatabaseDNS:     config.DatabaseDNS,
	}
}
