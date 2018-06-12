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

package list

import (
	"testing"
	"time"

	"github.com/sankt-petersbug/pwbook/internal/formatter"
	"github.com/sankt-petersbug/pwbook/internal/store"
)

func TestTemplate(t *testing.T) {
	testCases := []struct {
		name     string
		entries  []store.Entry
		expected string
	}{
		{
			name:    "0 entries",
			entries: []store.Entry{},
			expected: `Name                Password            Last Updated
----------------------------------------------------
Total 0 entries
`,
		},
		{
			name: "2 entries",
			entries: []store.Entry{
				store.Entry{"short_key", "short_value", time.Now(), time.Now()},
				store.Entry{"this_is_long_key", "This_is_long_pw", time.Now(), time.Now()},
			},
			expected: `Name                Password            Last Updated
----------------------------------------------------
short_key           short_value         0 days old
this_is_long_key    This_is_long_pw     0 days old
----------------------------------------------------
Total 2 entries
`,
		},
		{
			name: "verify last updated",
			entries: []store.Entry{
				store.Entry{"key", "value", time.Now().AddDate(0, 0, -1), time.Now().AddDate(0, 0, -1)},
			},
			expected: `Name                Password            Last Updated
----------------------------------------------------
key                 value               1 days old
----------------------------------------------------
Total 1 entries
`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			c := formatter.Context{"ListEntires", template}
			result, err := c.Format(tc.entries)
			if err != nil {
				t.Error(err)
			}

			if result != tc.expected {
				t.Errorf("expected result: %s, saw: %s", tc.expected, result)
			}
		})
	}
}
