package commands

import (
	"github.com/sankt-petersbug/pwbook/internal/commands/add"
	"github.com/sankt-petersbug/pwbook/internal/commands/list"
	"github.com/sankt-petersbug/pwbook/internal/commands/remove"
	"github.com/sankt-petersbug/pwbook/internal/commands/update"
	"github.com/sankt-petersbug/pwbook/internal/store"
	"github.com/spf13/cobra"
)

// NewPWBookCommand creates a root cobra.command and add subcommands
func NewPWBookCommand(pwbookStore *store.Store) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pwbook",
		Short: "a CLI for managing passwords like Sankt Petersbug does",
	}

	cmd.AddCommand(
		add.NewCommand(pwbookStore),
		list.NewCommand(pwbookStore),
		update.NewCommand(pwbookStore),
		remove.NewCommand(pwbookStore),
	)

	return cmd
}
