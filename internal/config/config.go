package config

import "os"

type ENV struct {
	Env    string
	Server string
	Port   string
	Common struct {
		Postgres struct {
			Host     string
			Port     string
			User     string
			Password string
			DBName   string
			SSLMode  string
		}

		Authentication struct {
			ApiKey    string
			JWTSecret string
		}
	}
}

var env ENV

func GetENV() *ENV {
	return &env
}

// Init ...
func Init() {
	env = ENV{
		Env:    os.Getenv("ENV"),
		Server: os.Getenv("SERVER"),
		Port:   os.Getenv("PORT"),
	}
}
