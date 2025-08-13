package database

import (
	"context"
	"fmt"

	"github.com/charmbracelet/huh/spinner"
	"github.com/meopedevts/typedvo-cli/internal/db"
	"github.com/meopedevts/typedvo-cli/internal/models"
	view "github.com/meopedevts/typedvo-cli/internal/views/database"
)

func RenderForm() error {
	dbConfig, err := view.DatabaseConfigurationForm()
	if err != nil {
		return err
	}

	err = validate(dbConfig)
	if err != nil {
		return err
	}

	return nil
}

func validate(dbConfig *models.DatabaseConfig) error {
	err := spinner.New().
		Title("🔍 Verificando conexão com banco de dados").
		ActionWithErr(func(ctx context.Context) error {
			db, err := db.New(dbConfig)
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
