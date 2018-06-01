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

package password

import (
    "testing"
    "strings"
)

func containedIn(s1 string, s2 string) bool {
    m := make(map[rune]bool)

    for _, c := range s2 {
        m[c] = true
    }
    for _, c := range s1 {
        if _, found := m[c]; !found {
            return false
        }
    }

    return true
}

func TestGenerateWithLength(t *testing.T) {
    testCases := []struct {
        name           string
        length         int
        expectedLength int
    }{
        {
            name:           "length = 0",
            length:         0,
            expectedLength: 0,
        },
        {
            name:           "length = 1",
            length:         1,
            expectedLength: 1,
        },
        {
            name:           "length = 2",
            length:         2,
            expectedLength: 2,
        },
    }

    for _, tc := range testCases {
        s := Generate(tc.length, nil)
        length := len(s)

        if length != tc.expectedLength {
            t.Errorf("[%s] expected: %d, actual: %d", tc.name, tc.expectedLength, length)
        }
    }
}

func TestGenerateWithOptions(t *testing.T) {
    testCases := []struct {
        name      string
        options   Options
        allowed string
    }{
        {
            name: "Lower letters only",
            options: Options{lowerLetters: true},
            allowed: LowerLetters,
        },
        {
            name: "Upper letters only",
            options: Options{upperLetters: true},
            allowed: UpperLetters,
        },
        {
            name: "Digits only",
            options: Options{digits: true},
            allowed: Digits,
        },
        {
            name: "Symbols only",
            options: Options{symbols: true},
            allowed: Symbols,
        },
        {
            name: "Nothing allwoed",
            options: Options{},
            allowed: "",
        },
    }

    for _, tc := range testCases {
        s := Generate(10, &tc.options)

        if !containedIn(s, tc.allowed) {
            t.Errorf("[%s] %s is not contained in %s", tc.name, s, tc.allowed)
        }
    }
}

func TestGenerate(t *testing.T) {
    categories := []string{LowerLetters, UpperLetters, Digits, Symbols}
    allowed := strings.Join(categories, "")
    n := int(10e3)

    for i := 0; i < n; i++ {
        s := Generate(10, nil)

        if !containedIn(s, allowed) {
            t.Errorf("Generated string %s contains non-allowed chars", s)
        }
    }
}

