package config

import (
	"os"
	"strconv"

	"github.com/sirupsen/logrus"
)

const (
	// LogLevelEnvVar is the name of the environment variable that controls
	// the log level of the application logger
	LogLevelEnvVar = "LOG_LEVEL"

	// PortEnvVar is the name of the environment variable that controls the
	// value of the port the service should listen on
	PortEnvVar = "PORT"

	// BrokerAddressEnvVar is the name of the environment variable that controls the
	// value of the kafka broker address
	BrokerAddressEnvVar = "BROKER_ADDRESS"

	defaultLogLevel      = logrus.DebugLevel // used if LOG_LEVEL not set
	defaultPort          = 8080              // used if PORT not set
	defaultBrokerAddress = "localhost"       // used if BROKER_ADDRESS not set
)

// LogLevel returns the log level set in the environment, or debug if not defined
func LogLevel() logrus.Level {
	var (
		level logrus.Level
		err   error
	)

	if level, err = logrus.ParseLevel(os.Getenv(LogLevelEnvVar)); err != nil {
		return defaultLogLevel
	}

	return level
}

// Port returns the port the service should listen on, or 3000 if not defined or
// is not a valid port
func Port() int {
	var (
		rawPort string
		found   bool
		port    int
		err     error
	)

	if rawPort, found = os.LookupEnv(PortEnvVar); !found {
		return defaultPort
	}

	if port, err = strconv.Atoi(rawPort); err != nil {
		return defaultPort
	}

	return port
}

// BrokerAddress returns the address the kafka broker is listening on, or localhost if not defined
func BrokerAddress() string {
	var brokerAddress string
	var found bool

	if brokerAddress, found = os.LookupEnv(BrokerAddressEnvVar); !found {
		return defaultBrokerAddress
	}

	if len(brokerAddress) == 0 {
		return defaultBrokerAddress
	}

	return brokerAddress
}
