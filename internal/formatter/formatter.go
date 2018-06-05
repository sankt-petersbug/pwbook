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
    "bytes"
    "strings"
    "text/tabwriter"
    "text/template"
)

// Context contains information required by the formatter to stringify the output as desired
type Context struct {
    Name string
    Template string
}

// Format the data provided using this context
func (c *Context) Format(data interface{}) (string, error) {
    var tpl bytes.Buffer

    t := template.New(c.Name)
    t, err := t.Parse(strings.Replace(c.Template, `\t`, "\t", -1))
    if err != nil {
        return "", err
    }

    tw := tabwriter.NewWriter(&tpl, 20, 1, 3, ' ', 0)
    if err := t.Execute(tw, data); err != nil {
        return "", err
    }

    tw.Flush()

    return tpl.String(), nil
}