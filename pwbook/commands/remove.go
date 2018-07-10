package commands

import (
	"github.com/sankt-petersbug/pwbook/pwbook"
	"github.com/spf13/cobra"
)

const removeTemplate = `Entry Removed
----------------------------------------------------
Name: {{.Data.Key}}
`

// NewRemoveCommand creates a new `pwbook remove` command
func NewRemoveCommand(ctx pwbook.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove [entry name]",
		Short: "Removes an entry",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			key := args[0]

			if err := ctx.Store.Remove(key); err != nil {
				return err
			}

			f := Formatter{
				Name:     "RemoveEntry",
				Template: removeTemplate,
				Output:   ctx.Output,
				Location: ctx.Location,
			}
			return f.Write(pwbook.Entry{Key: key})
		},
	}

	return cmd
}
