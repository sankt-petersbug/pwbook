package commands

import (
	"bytes"
	"testing"
	"time"
)

func getLocation(name string) *time.Location {
	loc, err := time.LoadLocation(name)
	if err != nil {
		panic(err)
	}

	return loc
}

func TestWrtieSuccess(t *testing.T) {
	testCases := []struct {
		name     string
		template string
		loc *time.Location
		data     interface{}
		expected string
	}{
		{
			name:     "Simple template",
			template: "Simple template with name {{.Data}}",
			data:     "YO!",
			expected: "Simple template with name YO!",
		},
		{
			name:     "Template with tabs",
			template: "Template\twith\ttabs",
			data:     nil,
			expected: "Template            with                tabs",
		},
		{
			name: "Location",
			template: `Datetime {{(.Data.In .Location).Format "02 Jan 06 15:04 MST"}}`,
			loc: getLocation("America/Chicago"),
			data: time.Date(2018, time.January, 1, 1, 0, 0, 0, time.UTC),
			expected: "Datetime 31 Dec 17 19:00 CST",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var buf bytes.Buffer

			f := Formatter{
				Name: tc.name,
				Template: tc.template,
				Output: &buf,
				Location: tc.loc,
			}
			if err := f.Write(tc.data); err != nil {
				t.Fatal(err)
			}

			if out := buf.String(); out != tc.expected {
				t.Errorf("expected result: %v, saw: %v", tc.expected, out)
			}
		})
	}
}

func TestWriteErrors(t *testing.T) {
	testCases := []struct {
		name     string
		template string
		data     interface{}
	}{
		{
			name:     "Invalid template",
			template: "This template is not {{ valid",
		},
		{
			name:     "Invalid data",
			template: "Template requires {{.Name}} Name attribute",
			data:     "This does not have Name attribute",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var buf bytes.Buffer

			f := Formatter{
				Name: tc.name,
				Template: tc.template,
				Output: &buf,
			}
			if err := f.Write(tc.data); err == nil {
				t.Error("expected an error, but got nil")
			}
		})
	}
}
