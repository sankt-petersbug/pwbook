package pwbook

import (
    "testing"
    "io/ioutil"
    "os"
    "reflect"
)

func tempfile() string {
    f, err := ioutil.TempFile("", "pwbook-")
    if err != nil {
        panic(err)
    }
    if err := f.Close(); err != nil {
        panic(err)
    }
    if err := os.Remove(f.Name()); err != nil {
        panic(err)
    }
    return f.Name()
}

func TestPWBookStoreCreateSuccess(t *testing.T) {
    f := tempfile()
    defer os.Remove(f)

    s, err := NewPWBookStore(f)
    defer s.Close()
    if err != nil {
        t.Fatal(err)
    }

    entry, err := s.Create("Name", "Password")
    if err != nil {
        t.Fatal(err)
    }

    entires, err := s.List()
    if err != nil {
        t.Fatal(err)
    }

    if l := len(entires); l != 1 {
        t.Fatalf("expected len(entries) to be 1, but saw: %d", l)
    }

    if !reflect.DeepEqual(entires[0], entry) {
        t.Fatalf("expected result: %v, saw: %v", entires[0], entry)
    } 
}

// TODO: add more tests..
