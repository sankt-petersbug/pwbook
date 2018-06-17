package formatter

import (
	"testing"
)

func TestContextFormat(t *testing.T) {
	testCases := []struct {
		name          string
		template      string
		data          interface{}
		errorExpected bool
		expected      string
	}{
		{
			name:          "Simple template",
			template:      "Simple template with name {{.}}",
			data:          "YO!",
			errorExpected: false,
			expected:      "Simple template with name YO!",
		},
		{
			name:          "Template with tabs",
			template:      "Template\twith\ttabs",
			data:          nil,
			errorExpected: false,
			expected:      "Template            with                tabs",
		},
		{
			name:          "Invalid template",
			template:      "This template is not {{ valid",
			errorExpected: true,
			expected:      "",
		},
		{
			name:          "Invalid data",
			template:      "Template requires {{.Name}} Name attribute",
			data:          "This does not have Name attribute",
			errorExpected: true,
			expected:      "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			c := Context{tc.name, tc.template}
			result, err := c.Format(tc.data)

			if (err != nil) != tc.errorExpected {
				t.Errorf("expected error != nil: %v, err: %v", tc.errorExpected, err)
			}

			if result != tc.expected {
				t.Errorf("expected result: %v, saw: %v", tc.expected, result)
			}
		})
	}
}
