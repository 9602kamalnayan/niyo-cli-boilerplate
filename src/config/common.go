package config

import "os"

var (
	Port    = os.Getenv("app_port")
	AppName = os.Getenv("app_name")
)
