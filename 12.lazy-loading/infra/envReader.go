package infra

import (
	"os"
)

type EnvReader struct{}

func NewEnvReader() *EnvReader {
	return &EnvReader{}
}

func (e *EnvReader) GetEnv(key string) string {
	return os.Getenv(key)
}
