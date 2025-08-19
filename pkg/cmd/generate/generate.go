package generate

import (
	"github.com/meopedevts/typedvo-cli/internal/cmd/generate"
	"github.com/spf13/cobra"
)

func NewCmdGenerate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "generate",
		Short: "Geração de código automática",
		RunE: func(cmd *cobra.Command, args []string) error {
			err := generate.Run()
			if err != nil {
				return err
			}
			return nil
		},
	}

	return cmd
}
