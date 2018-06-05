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
    "github.com/spf13/cobra"
    "github.com/sankt-petersbug/pwbook/internal/store"
    "github.com/sankt-petersbug/pwbook/internal/commands/add"
    "github.com/sankt-petersbug/pwbook/internal/commands/list"
    "github.com/sankt-petersbug/pwbook/internal/commands/update"
    "github.com/sankt-petersbug/pwbook/internal/commands/remove"
)

// NewPWBookCommand creates a root cobra.command and add subcommands
func NewPWBookCommand(pwbookStore store.Store) *cobra.Command {
    cmd := &cobra.Command{
        Use:   "pwbook",
        Short: "a CLI for managing passwords like Sankt Petersbug does",
    }

    cmd.AddCommand(
        add.NewCommand(pwbookStore),
        list.NewCommand(pwbookStore),
        update.NewCommand(pwbookStore),
        remove.NewCommand(pwbookStore),
    )

    return cmd
}
