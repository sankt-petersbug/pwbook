package main

import (
	"fmt"
	"os"
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/sankt-petersbug/pwbook/internal/commands"
	"github.com/sankt-petersbug/pwbook/internal/store"
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

	cmd := commands.NewPWBookCommand(&pwbookStore)
	if err := cmd.Execute(); err != nil {
		// os.Exit do not honor deferred calls
		pwbookStore.Close()
		os.Exit(1)
	}
}
