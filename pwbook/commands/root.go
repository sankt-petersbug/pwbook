package commands

import (
	"github.com/sankt-petersbug/pwbook/pwbook"
	"github.com/spf13/cobra"
)

// NewPWBookCommand creates a new `pwbook` command
func NewPWBookCommand(ctx pwbook.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pwbook",
		Short: "a CLI for managing passwords like Sankt Petersbug does",
	}

	cmd.AddCommand(
		NewAddCommand(ctx),
		NewListCommand(ctx),
		NewUpdateCommand(ctx),
		NewRemoveCommand(ctx),
	)

	return cmd
}
