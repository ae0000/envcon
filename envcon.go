package envcon

import (
	"fmt"
	"os"
)

// Environment maps to dev, prod or testing
type Environment string

// Environment statuses
var (
	Development Environment = "development"
	Production  Environment = "production"
	Testing     Environment = "testing"
)

// Config house the environment variables used throughout the application
type Config struct {
	Environment      Environment
	DatabasePort     string
	DatabaseUser     string
	DatabasePassword string
	DatabaseName     string
	DatabaseHost     string
	Port             string
	Key              string
	AdminEmail       string
	AppName          string
	AppURL           string
}

// Setup grabs the env variables and puts them in the config
func Setup(prefix string) Config {
	c := Config{}

	c.Environment = Environment(os.Getenv(fmt.Sprintf("%s_ENVIRONMENT", prefix)))
	c.DatabasePort = os.Getenv(fmt.Sprintf("%s_DATABASE_PORT", prefix))
	c.DatabaseUser = os.Getenv(fmt.Sprintf("%s_DATABASE_USER", prefix))
	c.DatabasePassword = os.Getenv(fmt.Sprintf("%s_DATABASE_PASSWORD", prefix))
	c.DatabaseName = os.Getenv(fmt.Sprintf("%s_DATABASE_NAME", prefix))
	c.DatabaseHost = os.Getenv(fmt.Sprintf("%s_DATABASE_HOST", prefix))
	c.Port = os.Getenv(fmt.Sprintf("%s_PORT", prefix))
	c.Key = os.Getenv(fmt.Sprintf("%s_KEY", prefix))
	c.AdminEmail = os.Getenv(fmt.Sprintf("%s_ADMIN_EMAIL", prefix))
	c.AppName = os.Getenv(fmt.Sprintf("%s_APP_NAME", prefix))
	c.AppURL = os.Getenv(fmt.Sprintf("%s_APP_URL", prefix))

	return c
}
