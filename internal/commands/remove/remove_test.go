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

package remove

import (
    "testing"

    "github.com/sankt-petersbug/pwbook/internal/store"
    "github.com/sankt-petersbug/pwbook/internal/formatter"
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