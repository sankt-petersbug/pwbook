package commands

import (
	"bytes"
	"errors"
	"io/ioutil"
	"strings"
	"testing"
	"time"

	"github.com/sankt-petersbug/pwbook/pwbook"
)

func TestNewListCommandError(t *testing.T) {
	testCases := []struct {
		name          string
		args          []string
		listFunc      func() ([]pwbook.Entry, error)
		expectedError string
	}{
		{
			name:          "more number of args",
			args:          []string{"args1"},
			expectedError: "unknown command",
		},
		{
			name: "return store's error msg",
			args: []string{},
			listFunc: func() ([]pwbook.Entry, error) {
				return []pwbook.Entry{}, errors.New("error from store")
			},
			expectedError: "error from store",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var buf bytes.Buffer

			ctx := pwbook.Context{
				Store:    &FakeStore{ListFunc: tc.listFunc},
				Output:   &buf,
				Location: nil,
			}

			cmd := NewListCommand(ctx)
			cmd.SetOutput(ioutil.Discard)
			cmd.SetArgs(tc.args)

			err := cmd.Execute()
			if err == nil {
				t.Fatal("expected error but saw nil")
			}
			if msg := err.Error(); !strings.Contains(msg, tc.expectedError) {
				t.Errorf("expected error message: %q, but saw: %q", tc.expectedError, msg)
			}
		})
	}
}

func TestNewListCommandSuccess(t *testing.T) {
	now := time.Now()

	testCases := []struct {
		name           string
		args           []string
		listFunc       func() ([]pwbook.Entry, error)
		expectedOutput string
	}{
		{
			name: "0 entry",
			args: []string{},
			expectedOutput: `Name                Password            Last Updated
----------------------------------------------------
Total 0 entries
`,
		},
		{
			name: "multiple entries",
			args: []string{},
			listFunc: func() ([]pwbook.Entry, error) {
				entries := []pwbook.Entry{
					{"short_key", "short_value", now, now},
					{"this_is_long_key", "This_is_long_pw", now, now},
				}

				return entries, nil
			},
			expectedOutput: `Name                Password            Last Updated
----------------------------------------------------
short_key           short_value         0 days old
this_is_long_key    This_is_long_pw     0 days old
----------------------------------------------------
Total 2 entries
`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var buf bytes.Buffer

			ctx := pwbook.Context{
				Store:    &FakeStore{ListFunc: tc.listFunc},
				Output:   &buf,
				Location: nil,
			}

			cmd := NewListCommand(ctx)
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
