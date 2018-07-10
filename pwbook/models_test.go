package pwbook

import (
	"testing"
	"time"
)

func TestEntryModifiedSince(t *testing.T) {
	testCases := []struct {
		name       string
		modifiedAt time.Time
		expected   int
	}{
		{
			name:       "modified today",
			modifiedAt: time.Now(),
			expected:   0,
		},
		{
			name:       "modified one day ago",
			modifiedAt: time.Now().AddDate(0, 0, -1),
			expected:   1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			entry := Entry{ModifiedAt: tc.modifiedAt}
			result := entry.ModifiedSince()

			if result != tc.expected {
				t.Errorf("expected result: %v, saw: %v", tc.expected, result)
			}
		})
	}
}
