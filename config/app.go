package config

import (
	"os"
	"strconv"
	"time"
)

type App struct {
	Host        string
	Port        int
	ReadTimeout *time.Duration
}

var app = &App{}

func AppCfg() *App {
	return app
}

func LoadApp() {
	port := os.Getenv("APP_PORT")
	timeOut := os.Getenv("APP_READ_TIMEOUT")
	readTimeout, _ := strconv.Atoi(timeOut)
	duration := time.Duration(readTimeout) * time.Second

	app.Host = os.Getenv("APP_HOST")
	app.Port, _ = strconv.Atoi(port)

	app.ReadTimeout = &duration

}
