package cmd

import (
	databaseCmd "github.com/meopedevts/typedvo-cli/pkg/cmd/database"
	generateCmd "github.com/meopedevts/typedvo-cli/pkg/cmd/generate"
	"github.com/spf13/cobra"
)

type exitCode int

const (
	exitOk  exitCode = 0
	exitErr exitCode = 1
)

func Execute() exitCode {
	cmd := &cobra.Command{
		Use:   "typedvo",
		Short: "Gera classes TypedVO Kotlin (VO e DAO) automaticamente a partir do banco de dados.",
		Long: `Gera automaticamente as classes TypedVO (VO e DAO) Kotlin com base na estrutura do banco de dados.

A CLI extrai metadados do banco e cria os arquivos de código correspondentes, simplificando o processo de criação das classes tipadas.`,
	}

	cmd.AddCommand(databaseCmd.NewCmdDatabase())
	cmd.AddCommand(generateCmd.NewCmdGenerate())

	err := cmd.Execute()
	if err != nil {
		return exitErr
	}

	return exitOk
}
