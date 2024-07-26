package utils

import (
	"fmt"
	_viper "github.com/spf13/viper"
	"os"
	"strconv"
	"strings"
	"sync"
)

type Viper interface {
	ConfigInterface
}

type viper struct {
	viper  *_viper.Viper
	loaded bool
	*sync.Mutex
}

func (v *viper) GetInt(key string) int {
	fullKey := ConfigRootKey + key
	if v.loaded {
		return v.viper.GetInt(fullKey)
	} else {
		val, _ := strconv.Atoi(os.Getenv(fullKey))
		return val
	}
}
func (v *viper) GetInt64(key string) int64 {
	fullKey := ConfigRootKey + key
	if v.loaded {
		return v.viper.GetInt64(fullKey)
	} else {
		val, _ := strconv.ParseInt(os.Getenv(fullKey), 10, 64)
		return val
	}
}
func (v *viper) GetFloat64(key string) float64 {
	fullKey := ConfigRootKey + key
	if v.loaded {
		return v.viper.GetFloat64(fullKey)
	} else {
		val, _ := strconv.ParseFloat(os.Getenv(fullKey), 64)
		return val
	}
}
func (v *viper) GetBool(key string) bool {
	fullKey := ConfigRootKey + key
	if v.loaded {
		return v.viper.GetBool(fullKey)
	} else {
		val, _ := strconv.ParseBool(os.Getenv(fullKey))
		return val
	}
}
func (v *viper) GetString(key string) string {
	fullKey := ConfigRootKey + key
	if v.loaded {
		return v.viper.GetString(fullKey)
	} else {
		return os.Getenv(fullKey)
	}
}
func (v *viper) GetStringSlice(key string) []string {
	fullKey := ConfigRootKey + key
	if v.loaded {
		return v.viper.GetStringSlice(fullKey)
	} else {
		return strings.Split(os.Getenv(fullKey), ",")
	}
}

var Config Viper

func InitConfig() error {
	v := _viper.New()
	v.SetConfigName(".env")
	v.SetConfigType("json")
	v.AddConfigPath(".")
	loaded := true

	if err := v.ReadInConfig(); err != nil {
		loaded = false
		fmt.Println("Config file not found; falling back to environment variables.")
	}

	Config = &viper{
		viper:  v,
		loaded: loaded,
		Mutex:  &sync.Mutex{},
	}

	return nil
}
