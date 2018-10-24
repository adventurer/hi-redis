package models

var AppConfig Config

type Config struct {
	Listen string
	DBIp   string
	DBPort string
	DBAuth string
}
