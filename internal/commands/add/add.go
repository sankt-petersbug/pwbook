package add

import (
	"errors"
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
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := validate(args); err != nil {
				return err
			}

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

func validate(args []string) error {
	if len(args) == 1 {
		return nil
	}

	return errors.New("add needs a name for the command")
}
