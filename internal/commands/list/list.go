package list

import (
	"fmt"

	"github.com/sankt-petersbug/pwbook/internal/formatter"
	"github.com/sankt-petersbug/pwbook/internal/store"
	"github.com/spf13/cobra"
)

const template = `Name\tPassword\tLast Updated
----------------------------------------------------
{{- range .}}
{{.Key}}\t{{.Value}}\t{{.ModifiedSince}} days old
{{- end}}
{{- if .}}
----------------------------------------------------
{{- end}}
Total {{. | len}} entries
`

// NewCommand creates a cobra.command for list command
func NewCommand(pwbookStore *store.Store) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List avilable entries",
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			entries, err := pwbookStore.List()
			if err != nil {
				return err
			}

			c := formatter.Context{"ListEntires", template}
			out, err := c.Format(entries)
			if err != nil {
				return err
			}

			fmt.Println(out)

			return nil
		},
	}

	return cmd
}
