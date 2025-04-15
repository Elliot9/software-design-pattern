package proxy

import (
	"github/elliot9/class12/infra"
	"github/elliot9/class12/internal/core"
)

const (
	PASSWORD_ENV_KEY = "PASSWORD"
	PASSWORD         = "1qaz2wsx"
)

type PasswordProtectionDatabaseProxy struct {
	envReader *infra.EnvReader
	database  core.Database
}

type PasswordProtectionError struct{}

func (e *PasswordProtectionError) Error() string {
	return "password check failed"
}

func NewPasswordProtectionDatabaseProxy(envReader *infra.EnvReader, database core.Database) *PasswordProtectionDatabaseProxy {
	return &PasswordProtectionDatabaseProxy{
		envReader: envReader,
		database:  database,
	}
}

func (p *PasswordProtectionDatabaseProxy) GetEmployeeById(id int) (core.Employee, error) {
	if !p.validatePassword() {
		return nil, &PasswordProtectionError{}
	}

	return p.database.GetEmployeeById(id)
}

func (p *PasswordProtectionDatabaseProxy) validatePassword() bool {
	return p.envReader.GetEnv(PASSWORD_ENV_KEY) == PASSWORD
}

var _ core.Database = &PasswordProtectionDatabaseProxy{}
