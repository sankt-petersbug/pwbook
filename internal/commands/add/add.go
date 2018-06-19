package add

import (
	"fmt"

	"github.com/sankt-petersbug/pwbook/internal/formatter"
	"github.com/sankt-petersbug/pwbook/internal/password"
	"github.com/sankt-petersbug/pwbook/internal/store"
	"github.com/spf13/cobra"
)

const template = `Entry Added
----------------------------------------------------
Name: {{.Key}}
Password: {{.Value}}
Created At: {{.CreatedAt.Format "02 Jan 06 15:04 MST"}}
`

// NewCommand creates a cobra.command for add command
func NewCommand(pwbookStore *store.Store) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add [entry name]",
		Short: "Add a new entry",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			key := args[0]

			value, err := password.GenerateStrong()
			if err != nil {
				return err
			}

			entry, err := pwbookStore.Create(key, value)
			if err != nil {
				return err
			}

			c := formatter.Context{"AddEntry", template}
			out, err := c.Format(entry)
			if err != nil {
				return err
			}

			fmt.Println(out)

			return nil
		},
	}

	return cmd
}
