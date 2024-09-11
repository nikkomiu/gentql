package config

import (
	"fmt"
	"os"
)

type App struct {
	Server   HTTPServer
	Database Database
}

type HTTPServer struct {
	Host string
	Port int
}

type Database struct {
	Driver string
	URL    string
}

func (hs HTTPServer) DisplayAddr() string {
	host := hs.Host
	if host == "" {
		host = "localhost"
	}
	return fmt.Sprintf("http://%s:%d/", host, hs.Port)
}

func (hs HTTPServer) Addr() string {
	return fmt.Sprintf("%s:%d", hs.Host, hs.Port)
}

func GetApp() App {
	return App{
		Server: HTTPServer{
			Host: "",
			Port: 8080,
		},
		Database: Database{
			Driver: "postgres",
			URL:    os.Getenv("DATABASE_URL"),
		},
	}
}
