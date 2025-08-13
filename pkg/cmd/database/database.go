package database

import (
	"github.com/meopedevts/typedvo-cli/internal/cmd/database"
	"github.com/spf13/cobra"
)

func NewCmdDatabase() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "database",
		Short: "Configurar banco de dados",
		RunE: func(cmd *cobra.Command, args []string) error {
			err := database.Run()
			if err != nil {
				return err
			}
			return nil
		},
	}

	return cmd
}
