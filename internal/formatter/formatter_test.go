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
            name: "Simple template",
            template: "Simple template with name {{.}}",
            data: "YO!",
            errorExpected: false,
            expected: "Simple template with name YO!",
        },
        {
            name: "Template with tabs",
            template: "Template\twith\ttabs",
            data: nil,
            errorExpected: false,
            expected: "Template            with                tabs",
        },
        {
            name: "Invalid template",
            template: "This template is not {{ valid",
            errorExpected: true,
            expected: "",
        },
        {
            name: "Invalid data",
            template: "Template requires {{.Name}} Name attribute",
            data: "This does not have Name attribute",
            errorExpected: true,
            expected: "",
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