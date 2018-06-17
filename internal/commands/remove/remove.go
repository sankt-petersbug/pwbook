package remove

import (
	"errors"
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
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := validate(args); err != nil {
				return err
			}

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

func validate(args []string) error {
	if len(args) == 1 {
		return nil
	}

	return errors.New("remove needs a name for the command")
}
