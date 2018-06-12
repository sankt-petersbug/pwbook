/*
 Copyright © 2018 Sankt Petersbug <sankt.petersbug@gmail.com>

 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at

     http:www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/

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
