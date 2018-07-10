package commands

import (
	"github.com/sankt-petersbug/pwbook/pwbook"
	"github.com/sankt-petersbug/pwbook/pwbook/password"
	"github.com/spf13/cobra"
)

const updateTemplate = `Entry Updated
----------------------------------------------------
Name: {{.Data.Key}}
Password: {{.Data.Value}}
Updated At: {{(.Data.ModifiedAt.In .Location).Format "02 Jan 06 15:04 MST"}}
`

// NewUpdateCommand creates a new `pwbook update` command
func NewUpdateCommand(ctx pwbook.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update [entry name]",
		Short: "Update password of an existing entry",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			key := args[0]

			value, err := password.GenerateStrong()
			if err != nil {
				return err
			}

			entry, err := ctx.Store.Update(key, value)
			if err != nil {
				return err
			}

			f := Formatter{
				Name:     "updateTemplate",
				Template: updateTemplate,
				Output:   ctx.Output,
				Location: ctx.Location,
			}
			return f.Write(entry)
		},
	}

	return cmd
}
