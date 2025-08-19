package database

import (
	"strconv"

	"github.com/charmbracelet/huh"
	"github.com/meopedevts/typedvo-cli/internal/models"
	"github.com/meopedevts/typedvo-cli/internal/views/utils"
)

func DatabaseConfigurationForm() (*models.DatabaseConfig, error) {
	var (
		host        string
		portStr     string
		serviceName string
		username    string
		password    string
	)

	form := huh.NewForm(huh.NewGroup(
		huh.NewInput().
			Title("Endereço / IP").
			Placeholder("127.0.0.1").
			Validate(utils.RequiredValidation).
			Value(&host),
		huh.NewInput().
			Title("Porta").
			Placeholder("1521").
			Validate(utils.RequiredValidation).
			Value(&portStr),
		huh.NewInput().
			Title("Serviço / SID").
			Placeholder("XE").
			Validate(utils.RequiredValidation).
			Value(&serviceName),
		huh.NewInput().
			Title("Usuário").
			Placeholder("oracle").
			Validate(utils.RequiredValidation).
			Value(&username),
		huh.NewInput().
			Title("Senha").
			Placeholder("******").
			Validate(utils.RequiredValidation).
			EchoMode(huh.EchoModePassword).
			Value(&password),
	).Title("⚙️ Configurando Banco de Dados"))

	err := form.Run()
	if err != nil {
		return nil, err
	}

	port, err := strconv.Atoi(portStr)
	if err != nil {
		return nil, err
	}

	return &models.DatabaseConfig{
		Host:     host,
		Port:     uint16(port),
		Service:  serviceName,
		Username: username,
		Password: password,
	}, nil
}
