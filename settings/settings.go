package settings

import (
	"flag"
	"log"
	"os"
	"sync"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
	"github.com/stuartshome/carpedia/logging"
)

type HTTPEnvSettings struct {
	Port    string `env:"HTTP_PORT" envDefault:":8100"`
	TlsCert string `env:"TLS_CERT" envDefault:""`
	TlsKey  string `env:"TLS_KEY" envDefault:""`
	IsDev   string `env:"IS_DEV" envDefault:""`
}

type DatabaseCreds struct {
	DBPort string `env:"DB_PORT" envDefault:":3006"`
	DBName string `env:"MYSQL_DATABASE" envDefault:"car_pedia"`
	DBPass string `env:"PASS" envDefault:"123456"`
	DBUser string `env:"MYSQL_USER" envDefault:"test1"`
}

type HttpSettings struct {
	ListenAddress  *string
	TlsCertificate *string
	TlsKey         *string
	TlsEnabled     bool
	IsDev          *string
}

type Settings struct {
	HttpSettings  HttpSettings
	DatabaseCreds DatabaseCreds
}

var globalSettings *Settings
var initialisationMutex sync.Mutex

func (s *Settings) init() {
	httpConfig := HTTPEnvSettings{}
	err := env.Parse(&httpConfig)
	if err != nil {
		logging.Logf("using default port after problem parsing http config ENV variables: %v", err.Error())
		httpConfig.Port = ":8080"
	}
	s.HttpSettings.IsDev = flag.String("development", httpConfig.IsDev, "If developing set this true")
	s.HttpSettings.ListenAddress = flag.String("listen", httpConfig.Port, "Address and port to bind HTTP server to")
	s.HttpSettings.TlsCertificate = flag.String("cert", httpConfig.TlsCert, "TLS certificate to use for service")
	s.HttpSettings.TlsKey = flag.String("key", httpConfig.TlsKey, "TLS key to use for service")

	flag.Parse()

	if len(*s.HttpSettings.TlsCertificate) > 0 || len(*s.HttpSettings.TlsKey) > 0 {
		s.HttpSettings.TlsEnabled = true
	}

	if len(*s.HttpSettings.ListenAddress) == 0 {
		if s.HttpSettings.TlsEnabled {
			*s.HttpSettings.ListenAddress = ":443"
		}
	}

	// Database creds
	if err := env.Parse(&s.DatabaseCreds); err != nil {
		logging.Logf("problem with parsing DatabaseCreds: %v", err)
	}
	// t := "true"
	err = godotenv.Load("script_config.env")
	if err != nil {
		log.Fatalf("error loading .env file")
	}
	if *s.HttpSettings.IsDev == "true" {
		logging.Log("hit")
		s.DatabaseCreds.DBName = os.Getenv("MYSQL_DATABASE")
		s.DatabaseCreds.DBPass = os.Getenv("PASS")
		s.DatabaseCreds.DBUser = os.Getenv("MYSQL_USER")
	}
	// if !s.HttpSettings.IsDev {
	// 	err := godotenv.Load("script_config.env")

	// 	s.DatabaseCreds.DBName = os.Getenv("MYSQL_DATABASE")
	// 	s.DatabaseCreds.DBPass = os.Getenv("PASS")
	// 	s.DatabaseCreds.DBUser = os.Getenv("MYSQL_USER")
	// }
}

func (s *Settings) logSettings() {
	if s.HttpSettings.ListenAddress != nil {
		logging.Logf("Listening for requests on port: %v", *s.HttpSettings.ListenAddress)
	}
	if s.HttpSettings.TlsEnabled {
		if s.HttpSettings.TlsCertificate != nil && s.HttpSettings.TlsKey != nil {
			logging.Logf("TLS enabled; TlsCertificate and TlsKey both set")
		} else {
			logging.Logf("TLS enabled but missing one of the TlsCertificte and TlsKey")
		}
	} else {
		logging.Log("TLS enabled")
	}

	logging.Logf("Database details are: %v", s.DatabaseCreds)
}

func Get() Settings {
	initialisationMutex.Lock()
	defer initialisationMutex.Unlock()
	if globalSettings == nil {
		s := Settings{}
		s.init()
		globalSettings = &s
		s.logSettings()
	}
	return *globalSettings
}
