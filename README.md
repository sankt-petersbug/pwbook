PWBook
======

PWBook is a command line tool for managing passwords.


### Install

To start using PWBook, install Go and run `go get`:

```sh
$ go get gihub.com/sankt-petersbug/pwbook/cmd/pwbook
```

This will retrieve the library and install the `pwbook` command line utility into
your `$GOBIN` path.

### Usage

Add a new entry with randomly generated password

```sh
$ pwbook add www.github.com
```

List available entries

```sh
$ pwbook list
```

Update existing password

```sh
$ pwbook update www.github.com
```

See alo pwbook --help

### TODO

- limit entry name and password length
- filter list by name, password age
- insert, update flags: ensure strong password, password requirements (letters, digits, symbols)
- encryption with master password