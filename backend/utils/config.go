package utils

type ConfigInterface interface {
	GetInt(key string) int
	GetInt64(key string) int64
	GetFloat64(key string) float64
	GetBool(key string) bool
	GetString(key string) string
	GetStringSlice(key string) []string
}
