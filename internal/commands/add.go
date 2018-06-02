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
    "errors"
    "strings"

    "github.com/spf13/cobra"
    "github.com/sankt-petersbug/pwbook/internal/store"
    "github.com/sankt-petersbug/pwbook/internal/password"
)

func NewAddCommand(pwbookStore store.Store) *cobra.Command {
    cmd := &cobra.Command{
        Use:   "add [entry name]",
        Short: "Add a new entry",
        RunE: func(cmd *cobra.Command, args []string) error {
            if len(args) == 0 {
                return errors.New("add needs a name for the command")
            }

            key := args[0]
            value := password.Generate(10, nil)

            entry, err := pwbookStore.Create(key, value)
            if err != nil {
                return err
            }

            datestr := entry.CreatedAt.Format(time.RFC822)
            divider := strings.Repeat("-", 31)

            fmt.Println("Entry Added")
            fmt.Println(divider)
            fmt.Printf("Name:       %s\n", entry.Key)
            fmt.Printf("Password:   %s\n", entry.Value)
            fmt.Printf("Updated At: %s\n", datestr)

            return nil
        },
    }

    return cmd
}