package formatter

import (
	"bytes"
	"strings"
	"text/tabwriter"
	"text/template"
)

// Context contains information required by the formatter to stringify the output as desired
type Context struct {
	Name     string
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
