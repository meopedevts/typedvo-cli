package generate

import (
	"github.com/charmbracelet/huh"
	"github.com/meopedevts/typedvo-cli/internal/models"
	"github.com/meopedevts/typedvo-cli/internal/views/utils"
)

func GenerateForm() (*models.GenerateFormFields, error) {
	var entityName string

	form := huh.NewForm(huh.NewGroup(
		huh.NewInput().
			Title("Entidade").
			Placeholder("CabecalhoNota").
			Validate(utils.RequiredValidation).
			Value(&entityName),
	).Title("Informe a entidade para geração dos DAO e VO"))

	err := form.Run()
	if err != nil {
		return nil, err
	}

	return &models.GenerateFormFields{EntityName: entityName}, nil
}
