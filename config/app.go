package config

import (
	"strconv"
	"time"
)

type App struct {
	Host        *string
	Port        int
	ReadTimeout *time.Duration
}

var app = &App{}

func AppCfg() *App {
	return app
}

func LoadApp() {
	port, _ := GoDotEnvVariable("APP_PORT")
	timeOut, _ := GoDotEnvVariable("APP_READ_TIMEOUT")
	readTimeout, _ := strconv.Atoi(*timeOut)
	duration := time.Duration(readTimeout) * time.Second

	app.Host, _ = GoDotEnvVariable("APP_HOST")
	app.Port, _ = strconv.Atoi(*port)

	app.ReadTimeout = &duration

}
