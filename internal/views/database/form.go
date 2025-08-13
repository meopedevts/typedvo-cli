package database

import (
	"errors"
	"strconv"

	"github.com/charmbracelet/huh"
	"github.com/meopedevts/typedvo-cli/internal/models"
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
			Validate(requiredValidation).
			Value(&host),
		huh.NewInput().
			Title("Porta").
			Placeholder("1521").
			Validate(requiredValidation).
			Value(&portStr),
		huh.NewInput().
			Title("Serviço / SID").
			Placeholder("XE").
			Validate(requiredValidation).
			Value(&serviceName),
		huh.NewInput().
			Title("Usuário").
			Placeholder("oracle").
			Validate(requiredValidation).
			Value(&username),
		huh.NewInput().
			Title("Senha").
			Placeholder("******").
			Validate(requiredValidation).
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

func requiredValidation(v string) error {
	if len(v) <= 1 {
		return errors.New("Campo obrigatório, deve ser preenchido.")
	}
	return nil
}
