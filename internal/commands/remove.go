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
    "errors"
    "strings"

    "github.com/spf13/cobra"
    "github.com/sankt-petersbug/pwbook/internal/store"
)

func NewRemoveCommand(pwbookStore store.Store) *cobra.Command {
    cmd := &cobra.Command{
        Use:   "remove [entry name]",
        Short: "Removes an entry",
        RunE: func(cmd *cobra.Command, args []string) error {
            if len(args) == 0 {
                return errors.New("remove needs a name for the command")
            }

            key := args[0]

            if err := pwbookStore.Delete(key); err != nil {
                return err
            }

            divider := strings.Repeat("-", 31)

            fmt.Println("Entry Removed")
            fmt.Println(divider)
            fmt.Printf("Name: %s\n", key)

            return nil
        },
    }

    return cmd
}