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

package commands

import (
    "fmt"
    "time"
    "strings"

    "github.com/spf13/cobra"
    "github.com/sankt-petersbug/pwbook/internal/store"
)

// NewListCommand creates a cobra.command for list command
func NewListCommand(pwbookStore store.Store) *cobra.Command {
    cmd := &cobra.Command{
        Use:   "list",
        Short: "List avilable entries",
        RunE: func(cmd *cobra.Command, args []string) error {
            entries, err := pwbookStore.List()
            if err != nil {
                return err
            }

            headers := fmt.Sprintf("%-20s%-20s%s", "Name", "Password", "Last Updated")
            divider := strings.Repeat("-", len(headers))

            fmt.Println(headers)
            fmt.Println(divider)

            for _, entry := range entries {
                d := time.Since(entry.ModifiedAt)
                daysOld := int(d.Hours() / 24)

                fmt.Printf("%-20s%-20s%d days old\n", entry.Key, entry.Value, daysOld)
            }
            fmt.Println(divider)
            fmt.Printf("Total %d entries\n", len(entries))

            return nil
        },
    }

    return cmd
}