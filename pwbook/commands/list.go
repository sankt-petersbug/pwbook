package commands

import (
	"github.com/sankt-petersbug/pwbook/pwbook"
	"github.com/spf13/cobra"
)

const listTemplate = `Name\tPassword\tLast Updated
----------------------------------------------------
{{- range .Data}}
{{.Key}}\t{{.Value}}\t{{.ModifiedSince}} days old
{{- end}}
{{- if .Data}}
----------------------------------------------------
{{- end}}
Total {{.Data | len}} entries
`

// NewListCommand creates a new `pwbook list` command
func NewListCommand(ctx pwbook.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List available entries",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			entries, err := ctx.Store.List()
			if err != nil {
				return err
			}

			f := Formatter{
				Name:     "ListEntires",
				Template: listTemplate,
				Output:   ctx.Output,
				Location: ctx.Location,
			}
			return f.Write(entries)
		},
	}

	return cmd
}
