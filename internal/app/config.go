package app

import "flag"

type Config struct {
	ServerAddress   string
	BaseURL         string
	FileStoragePath string
	DatabaseDNS     string
}

func NewConfig() *Config {
	ServerAddress := flag.String("a", "127.0.0.1:8080", "ServerAddress - адрес запуска HTTP-сервера")
	BaseURL := flag.String("b", "http://"+*ServerAddress, "BaseURL")
	FileStoragePath := flag.String("f", "texts.txt", "FileStoragePath - путь до файла LongURL")
	DatabaseDNS := flag.String("d", "", "DatabaseDNS")
	flag.Parse()
	return &Config{
		ServerAddress:   *ServerAddress,
		BaseURL:         *BaseURL,
		FileStoragePath: *FileStoragePath,
		DatabaseDNS:     *DatabaseDNS,
	}
}
