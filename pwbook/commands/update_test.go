package commands

import (
	"bytes"
	"errors"
	"io/ioutil"
	"testing"
	"time"

	"github.com/sankt-petersbug/pwbook/pwbook"
)

func TestNewUpdateCommandErrors(t *testing.T) {
	testCases := []struct {
		name          string
		args          []string
		updateFunc    func(key string, value string) (pwbook.Entry, error)
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
			updateFunc: func(key string, value string) (pwbook.Entry, error) {
				return pwbook.Entry{}, errors.New("error from store")
			},
			expectedError: "error from store",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var buf bytes.Buffer

			ctx := pwbook.Context{
				Store: &FakeStore{UpdateFunc: tc.updateFunc}, 
				Output: &buf,
				Location: nil,
			}

			cmd := NewUpdateCommand(ctx)
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

func TestNewUpdateCommandSuccess(t *testing.T) {
	testCases := []struct {
		name           string
		args           []string
		updateFunc     func(key string, value string) (pwbook.Entry, error)
		expectedOutput string
	}{
		{
			name: "update entry",
			args: []string{"Entry"},
			updateFunc: func(key string, value string) (pwbook.Entry, error) {
				entry := pwbook.Entry{
					Key:        key,
					Value:      "Generated_password",
					ModifiedAt: time.Date(2018, time.January, 1, 1, 0, 0, 0, time.UTC),
				}

				return entry, nil
			},
			expectedOutput: `Entry Updated
----------------------------------------------------
Name: Entry
Password: Generated_password
Updated At: 01 Jan 18 01:00 UTC
`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var buf bytes.Buffer

			ctx := pwbook.Context{
				Store: &FakeStore{UpdateFunc: tc.updateFunc}, 
				Output: &buf,
				Location: time.UTC,
			}

			cmd := NewUpdateCommand(ctx)
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
