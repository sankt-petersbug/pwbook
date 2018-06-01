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

package commands

import (
    "testing"
    "io/ioutil"

    "github.com/sankt-petersbug/pwbook/internal/store"
)

func TestNewUpdateCommandWithNoArgs(t *testing.T) {
    cmd := NewUpdateCommand(store.Store{})
    cmd.SetOutput(ioutil.Discard)
    cmd.SetArgs([]string{})

    if err := cmd.Execute(); err == nil {
        t.Fatal("Must throw error")
    }
}

func TestValidateUpdateCommandArgsWithNoArgs(t *testing.T) {
    err := validateUpdateCommandArgs([]string{})
    if err == nil {
        t.Fatal("Should return an error")
    }
}

func TestValidateUpdateCommandArgs(t * testing.T) {
    testCases := map[string]struct {
        a        []string
        expected error
    }{
        "length = 1": {
            a:        []string{"args1"},
            expected: nil,
        },
        "length = 2": {
            a:        []string{"args1", "args2"},
            expected: nil,
        },
    }
    for testName, testCase := range testCases {
        actual := validateUpdateCommandArgs(testCase.a)

        if actual != testCase.expected {
            t.Errorf("[%s] expected: %v, actual: %v", testName, testCase.expected, actual)
        }
    }
}