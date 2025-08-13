package database

import (
	"context"
	"fmt"

	"github.com/charmbracelet/huh/spinner"
	"github.com/meopedevts/typedvo-cli/internal/db"
	"github.com/meopedevts/typedvo-cli/internal/models"
	view "github.com/meopedevts/typedvo-cli/internal/views/database"
)

func Run() error {
	dbConfig, err := render()
	if err != nil {
		return err
	}

	err = validate(dbConfig)
	if err != nil {
		return err
	}

	return nil
}

func render() (*models.DatabaseConfig, error) {
	cfg, err := view.DatabaseConfigurationForm()
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

func validate(cfg *models.DatabaseConfig) error {
	err := spinner.New().
		Title("🔍 Verificando conexão com banco de dados").
		ActionWithErr(func(ctx context.Context) error {
			db, err := db.New(cfg)
			if err != nil {
				return err
			}
			defer db.Close()

			return nil
		}).
		Run()
	if err != nil {
		return err
	}
	fmt.Println("✅ Conexão com banco de dados, realizada com sucesso!")

	return nil
}
