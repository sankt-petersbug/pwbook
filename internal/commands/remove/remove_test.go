package remove

import (
	"testing"

	"github.com/sankt-petersbug/pwbook/internal/formatter"
	"github.com/sankt-petersbug/pwbook/internal/store"
)

func TestValidate(t *testing.T) {
	testCases := []struct {
		name        string
		args        []string
		expectedErr bool
	}{
		{
			name:        "success",
			args:        []string{"1"},
			expectedErr: false,
		},
		{
			name:        "fail (empty)",
			args:        []string{},
			expectedErr: true,
		},
		{
			name:        "fail (more than one args)",
			args:        []string{"1", "2"},
			expectedErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := validate(tc.args)
			if (err != nil) != tc.expectedErr {
				t.Errorf("expected error: %v, saw: %v, error: %v", tc.expectedErr, err != nil, err)
			}
		})
	}
}

func TestTemplate(t *testing.T) {
	testCases := []struct {
		name     string
		entry    store.Entry
		expected string
	}{
		{
			name: "entry",
			entry: store.Entry{
				Key: "Name",
			},
			expected: `Entry Removed
----------------------------------------------------
Name: Name
`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			c := formatter.Context{"RemoveEntry", template}
			result, err := c.Format(tc.entry)
			if err != nil {
				t.Error(err)
			}

			if result != tc.expected {
				t.Errorf("expected result: %s, saw: %s", tc.expected, result)
			}
		})
	}
}
