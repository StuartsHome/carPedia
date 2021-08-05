package settings

import (
	"flag"
	"sync"

	"github.com/caarlos0/env"
	"github.com/stuartshome/carpedia/logging"
)

type HTTPEnvSettings struct {
	Port    string `env:"HTTP_PORT" envDefault:":8100"`
	TlsCert string `env:"TLS_CERT" envDefault:""`
	TlsKey  string `env:"TLS_KEY" envDefault:""`
}

type DatabaseCreds struct {
	DBPort string `env:"DB_PORT" envDefault:":3006"`
}

type HttpSettings struct {
	ListenAddress  *string
	TlsCertificate *string
	TlsKey         *string
	TlsEnabled     bool
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

	logging.Logf("Database port is: %v", s.DatabaseCreds)
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
