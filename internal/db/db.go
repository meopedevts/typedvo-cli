package db

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/meopedevts/typedvo-cli/internal/models"
	go_ora "github.com/sijms/go-ora/v2"
)

var (
	ErrDatabaseConnection = errors.New("erro durante a conexão com o banco de dados")
	ErrConfigNotFound     = errors.New("nenhuma configuração de banco de dados encontrada")
)

func New(config *models.DatabaseConfig) (*sql.DB, error) {
	var connStr string
	if config != nil {
		connStr = go_ora.BuildUrl(config.Host, int(config.Port), config.Service, config.Username, config.Password, nil)
	} else {
		return nil, fmt.Errorf("%w, %w", ErrDatabaseConnection, ErrConfigNotFound)
	}

	sql, err := sql.Open("oracle", connStr)
	if err != nil {
		return nil, fmt.Errorf("%w, %w", ErrDatabaseConnection, err)
	}

	err = sql.Ping()
	if err != nil {
		return nil, fmt.Errorf("%w, %w", ErrDatabaseConnection, err)
	}

	return sql, nil
}
