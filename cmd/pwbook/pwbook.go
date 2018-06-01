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

package main

import (
    "os"
    "fmt"
    "path/filepath"

    homedir "github.com/mitchellh/go-homedir"
    "github.com/sankt-petersbug/pwbook/internal/store"
    "github.com/sankt-petersbug/pwbook/internal/commands"
)

func makeBaseDir() string {
    userhome, err := homedir.Dir()
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    baseDir := filepath.Join(userhome, ".pwbook")
    if err := os.Mkdir(baseDir, 0700); err != nil && !os.IsExist(err) {
        fmt.Println(err)
        os.Exit(1)
    }

    return baseDir
}

func main() {
    baseDir := makeBaseDir()
    dbPath := filepath.Join(baseDir, "pwbook.db")
    pwbookStore, err := store.NewStore(dbPath)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    defer pwbookStore.Close()

    cmd := commands.NewPWBookCommand(pwbookStore)
    if err := cmd.Execute(); err != nil {
        // os.Exit do not honor deferred calls
        pwbookStore.Close()
        os.Exit(1)
    }
}