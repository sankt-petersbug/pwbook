package pwbook

import (
	"io"
	"time"
)

// Context stores information required for `PWBookCommand`
type Context struct {
	// Store provides api to manipulate the data
	Store Store

	// Location is used to format the output
	Location *time.Location

	// Output is used to write the output
	Output io.Writer
}
