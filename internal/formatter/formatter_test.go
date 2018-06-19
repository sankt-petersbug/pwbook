package formatter

import (
	"testing"
)

func TestContextFormatSuccess(t *testing.T) {
	testCases := []struct {
		name     string
		template string
		data     interface{}
		expected string
	}{
		{
			name:     "Simple template",
			template: "Simple template with name {{.}}",
			data:     "YO!",
			expected: "Simple template with name YO!",
		},
		{
			name:     "Template with tabs",
			template: "Template\twith\ttabs",
			data:     nil,
			expected: "Template            with                tabs",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			c := Context{tc.name, tc.template}
			result, err := c.Format(tc.data)

			if err != nil {
				t.Errorf("expected nil, but got an error %q", err)
			}

			if result != tc.expected {
				t.Errorf("expected result: %v, saw: %v", tc.expected, result)
			}
		})
	}
}

func TestContextFormatFail(t *testing.T) {
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
			c := Context{tc.name, tc.template}
			_, err := c.Format(tc.data)

			if err == nil {
				t.Errorf("expected an error, but got nil")
			}
		})
	}
}
