PWBook
======

PWBook is a command line tool for managing passwords.


### Install

To start using PWBook, install Go and run `go get`:

```sh
$ go get github.com/sankt-petersbug/pwbook/cmd/pwbook
```

This will retrieve the library and install the `pwbook` command line utility into
your `$GOBIN` path.

### Usage

```sh
$ pwbook

a CLI for managing passwords like Sankt Petersbug does

Usage:
  pwbook [command]

Available Commands:
  add         Add a new entry
  help        Help about any command
  list        List avilable entries
  remove      Removes an entry
  update      Update password of an existing entry

Flags:
  -h, --help   help for pwbook

Use "pwbook [command] --help" for more information about a command.

```

See alo pwbook --help

#### Add

Creates a new entry

```sh
$ pwbook add "Entry name"
```

#### List

List avilable entries

```sh
$ pwbook list
```

#### Update

Update password of an existing entry

```sh
$ pwbook update "Entry name"
```

#### Remove

Removes an existing entry

```sh
$ pwbook remove "Entry name"
```

### TODO

- limit entry name and password length
- filter list by name, password age
- tag or groups
- encryption with master password