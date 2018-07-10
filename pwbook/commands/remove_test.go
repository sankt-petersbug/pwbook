package commands

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/sankt-petersbug/pwbook/pwbook"
)

func TestNewRemoveCommandErrors(t *testing.T) {
	testCases := []struct {
		name          string
		args          []string
		removeFunc    func(key string) error
		expectedError string
	}{
		{
			name:          "less number of args",
			args:          []string{},
			expectedError: "accepts 1 arg(s), received 0",
		},
		{
			name:          "more number of args",
			args:          []string{"args1", "args2"},
			expectedError: "accepts 1 arg(s), received 2",
		},
		{
			name: "return store's error msg",
			args: []string{"args1"},
			removeFunc: func(key string) error {
				return errors.New("error from store")
			},
			expectedError: "error from store",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var buf bytes.Buffer

			ctx := pwbook.Context{
				Store:    &FakeStore{RemoveFunc: tc.removeFunc},
				Output:   &buf,
				Location: nil,
			}

			cmd := NewRemoveCommand(ctx)
			cmd.SetOutput(ioutil.Discard)
			cmd.SetArgs(tc.args)

			err := cmd.Execute()
			if err == nil {
				t.Fatal(err)
			}
			if msg := err.Error(); msg != tc.expectedError {
				t.Errorf("expected error message: %q, but saw: %q", tc.expectedError, msg)
			}
		})
	}
}

func TestNewRemoveCommandSuccess(t *testing.T) {
	testCases := []struct {
		name           string
		args           []string
		removeFunc     func(key string) error
		expectedOutput string
	}{
		{
			name: "remove entry",
			args: []string{"Entry"},
			removeFunc: func(key string) error {
				if key != "Entry" {
					msg := fmt.Sprintf("expected key %q but saw %q", "Entry", key)
					return errors.New(msg)
				}

				return nil
			},
			expectedOutput: `Entry Removed
----------------------------------------------------
Name: Entry
`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var buf bytes.Buffer

			ctx := pwbook.Context{
				Store:    &FakeStore{RemoveFunc: tc.removeFunc},
				Output:   &buf,
				Location: nil,
			}

			cmd := NewRemoveCommand(ctx)
			cmd.SetOutput(ioutil.Discard)
			cmd.SetArgs(tc.args)

			err := cmd.Execute()
			if err != nil {
				t.Fatal(err)
			}
			if output := buf.String(); output != tc.expectedOutput {
				t.Errorf("expected result: %q, but saw: %q", tc.expectedOutput, output)
			}
		})
	}
}
