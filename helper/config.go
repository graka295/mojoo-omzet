package helper

import (
	"flag"
	"fmt"
	"os"

	"github.com/jinzhu/configor"
)

// Config contract for ConfigImpl
type Config interface {
	GetConfig() *ConfigImpl
}

// ConfigImpl struct for config
type ConfigImpl struct {
	Mode     string
	Database struct {
		Dialect   string `default:"mysql"`
		Host      string
		Port      string
		Dbname    string
		Username  string
		Password  string
		Migration bool `default:"false"`
	}
}

// Load load init config
func Load() Config {
	var env *string
	if value := os.Getenv("MODE"); value != "" {
		env = &value
	} else {
		env = flag.String("env", "develop", "To switch configurations.")
		flag.Parse()
	}
	config := &ConfigImpl{}
	if err := configor.Load(config, "application."+*env+".yaml"); err != nil {
		panic(fmt.Sprintf("Failed to read application.%s.yaml: %s", *env, err))
	}
	config.Mode = *env
	return config
}

// GetConfig get config
func (x *ConfigImpl) GetConfig() *ConfigImpl {
	return x
}
