/*
 Copyright © 2018 Sankt Petersbug <sankt.petersbug@gmail.com>

 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at

     http:www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/

package store

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
