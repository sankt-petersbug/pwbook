package commands

import (
	"io"
	"strings"
	"text/tabwriter"
	"text/template"
	"time"
)

type payload struct {
	Data     interface{}
	Location *time.Location
}

// Formatter formats data and output the result to writer
type Formatter struct {
	Name     string
	Template string
	Output   io.Writer
	Location *time.Location
}

// Write the formatted output using the context
func (f *Formatter) Write(data interface{}) error {
	t, err := template.New(f.Name).Parse(strings.Replace(f.Template, `\t`, "\t", -1))
	if err != nil {
		return err
	}

	p := payload{
		Data:     data,
		Location: f.Location,
	}
	tw := tabwriter.NewWriter(f.Output, 20, 1, 3, ' ', 0)
	if err := t.Execute(tw, p); err != nil {
		return err
	}

	tw.Flush()

	return nil
}
