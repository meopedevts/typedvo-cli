package generate

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/meopedevts/typedvo-cli/internal/db"
	"github.com/meopedevts/typedvo-cli/internal/models"
	view "github.com/meopedevts/typedvo-cli/internal/views/generate"
)

var (
	errCodeGeneration   = errors.New("erro durante a geração de código")
	errInstanceNotFound = errors.New("nenhuma instancia encontrada")
)

type entityField struct {
	nuCampo    int64
	nomeCampo  string
	descrCampo string
	tipCampo   string
}

func Run() error {
	form, err := render()
	if err != nil {
		return err
	}

	// TODO - Remover saporra dps
	cfg := &models.DatabaseConfig{
		Host:     "127.0.0.1",
		Port:     1521,
		Service:  "XE",
		Username: "SANKHYA",
		Password: "developer",
	}

	db, err := db.New(cfg)
	if err != nil {
		return err
	}
	defer db.Close()

	tableName, err := getTableName(db, form.EntityName)
	if err != nil {
		return err
	}

	entityFields, err := getFields(db, tableName)
	if err != nil {
		return err
	}

	err = createDir(form.EntityName)
	if err != nil {
		return err
	}

	err = generateVO(form.EntityName, entityFields)
	if err != nil {
		return err
	}

	fmt.Println("DAO e VO gerados com sucesso!")

	return nil
}

func render() (*models.GenerateFormFields, error) {
	form, err := view.GenerateForm()
	if err != nil {
		return nil, err
	}
	return form, nil
}

func getTableName(db *sql.DB, instanceName string) (string, error) {
	var tableName string

	rows, err := db.Query("SELECT NOMETAB FROM TDDINS WHERE NOMEINSTANCIA = :1",
		instanceName,
	)
	if err != nil {
		return "", err
	}

	if !rows.Next() {
		return "", fmt.Errorf("%w: %w", errCodeGeneration, errInstanceNotFound)
	}

	err = rows.Scan(&tableName)
	if err != nil {
		return "", fmt.Errorf("%w: %w", errCodeGeneration, err)
	}

	return tableName, nil
}

func getFields(db *sql.DB, tableName string) ([]entityField, error) {
	fields := make([]entityField, 0)

	rows, err := db.Query("SELECT NUCAMPO, NOMECAMPO, DESCRCAMPO, TIPCAMPO FROM TDDCAM WHERE NOMETAB = :1 AND CALCULADO = 'N' ORDER BY 2",
		tableName,
	)
	if err != nil {
		return fields, err
	}

	for rows.Next() {
		var (
			nuCampo    int
			nomeCampo  string
			descrCampo string
			tipCampo   string
		)

		err := rows.Scan(
			&nuCampo,
			&nomeCampo,
			&descrCampo,
			&tipCampo,
		)
		if err != nil {
			return fields, err
		}

		fields = append(fields, entityField{
			nuCampo:    int64(nuCampo),
			nomeCampo:  nomeCampo,
			descrCampo: descrCampo,
			tipCampo:   tipCampo,
		})
	}

	return fields, nil
}
