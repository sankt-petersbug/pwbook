package remove

import (
	"fmt"

	"github.com/sankt-petersbug/pwbook/internal/formatter"
	"github.com/sankt-petersbug/pwbook/internal/store"
	"github.com/spf13/cobra"
)

const template = `Entry Removed
----------------------------------------------------
Name: {{.Key}}
`

// NewCommand creates a cobra.command for remove command
func NewCommand(pwbookStore *store.Store) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove [entry name]",
		Short: "Removes an entry",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			key := args[0]

			if err := pwbookStore.Delete(key); err != nil {
				return err
			}

			c := formatter.Context{"RemoveEntry", template}
			out, err := c.Format(store.Entry{Key: key})
			if err != nil {
				return err
			}

			fmt.Println(out)

			return nil
		},
	}

	return cmd
}
