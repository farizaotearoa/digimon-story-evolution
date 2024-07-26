package utils

import (
	"fmt"
	_viper "github.com/spf13/viper"
	"log"
	"os"
	"strings"
	"sync"
)

type Viper interface {
	ConfigInterface
}

type viper struct {
	viper *_viper.Viper
	*sync.Mutex
}

func (v *viper) GetInt(key string) int         { return v.viper.GetInt(ConfigRootKey + key) }
func (v *viper) GetInt64(key string) int64     { return v.viper.GetInt64(ConfigRootKey + key) }
func (v *viper) GetFloat64(key string) float64 { return v.viper.GetFloat64(ConfigRootKey + key) }
func (v *viper) GetBool(key string) bool       { return v.viper.GetBool(ConfigRootKey + key) }
func (v *viper) GetString(key string) string   { return v.viper.GetString(ConfigRootKey + key) }
func (v *viper) GetStringSlice(key string) []string {
	return v.viper.GetStringSlice(ConfigRootKey + key)
}

var Config Viper

func InitConfig() error {
	v := _viper.New()
	v.SetConfigName(".env")
	v.SetConfigType("json")
	v.AddConfigPath(".")

	if err := v.ReadInConfig(); err != nil {
		fmt.Println("Config file not found; falling back to environment variables.")
	}

	v.AutomaticEnv()
	for _, env := range os.Environ() {
		log.Println(env)
	}
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	Config = &viper{
		viper: v,
		Mutex: &sync.Mutex{},
	}

	_viper.Set("test_key", "test_value")
	testValue := _viper.GetString("test_key")
	log.Printf("Test Key Value: %s", testValue)

	return nil
}
