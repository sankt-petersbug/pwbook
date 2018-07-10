package main

import (
	"fmt"
	"time"
	"os"
	"path/filepath"
	"bytes"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/sankt-petersbug/pwbook/pwbook/commands"
	"github.com/sankt-petersbug/pwbook/pwbook"
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
	fpath := filepath.Join(baseDir, "pwbook.db")
	s, err := pwbook.NewPWBookStore(fpath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer s.Close()

	var buf bytes.Buffer

	loc, err := time.LoadLocation("Local")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	ctx := pwbook.Context{
		Store: &s,
		Output: &buf,
		Location: loc,
	}
	commands.NewPWBookCommand(ctx).Execute();

	if output := buf.String(); output != "" {
		fmt.Print(output)
	}
}
