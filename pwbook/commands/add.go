package commands

import (
	"github.com/sankt-petersbug/pwbook/pwbook"
	"github.com/sankt-petersbug/pwbook/pwbook/password"
	"github.com/spf13/cobra"
)

const addTemplate = `Entry Added
----------------------------------------------------
Name: {{.Data.Key}}
Password: {{.Data.Value}}
Created At: {{(.Data.CreatedAt.In .Location).Format "02 Jan 06 15:04 MST"}}
`

// NewAddCommand creates a new `pwbook add` command
func NewAddCommand(ctx pwbook.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add [entry name]",
		Short: "Add a new entry",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			key := args[0]

			value, err := password.GenerateStrong()
			if err != nil {
				return err
			}

			entry, err := ctx.Store.Create(key, value)
			if err != nil {
				return err
			}

			f := Formatter{
				Name:     "AddEntry",
				Template: addTemplate,
				Output:   ctx.Output,
				Location: ctx.Location,
			}
			return f.Write(entry)
		},
	}

	return cmd
}
