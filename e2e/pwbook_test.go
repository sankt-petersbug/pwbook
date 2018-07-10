package pwbook

import (
	"os/exec"
	"strings"
	"testing"
)

func assertNoError(t *testing.T, err error) {
	if err != nil {
		t.Errorf("expected nil, but got an error %q", err)
	}
}

func assertOutputContains(t *testing.T, output []byte, substr string) {
	s := string(output)

	if !strings.Contains(s, substr) {
		t.Errorf("%q does not contain %q", s, substr)
	}
}

func TestE2E(t *testing.T) {
	var out []byte
	var err error
	const entryName = "entry1"

	out, err = exec.Command("pwbook", "add", entryName).Output()
	assertNoError(t, err)
	assertOutputContains(t, out, entryName)
	t.Logf("Added entry %q\n", entryName)

	out, err = exec.Command("pwbook", "list").Output()
	assertNoError(t, err)
	assertOutputContains(t, out, "entry1")
	assertOutputContains(t, out, "Total 1 entries")
	t.Logf("Verified list containing the entry %q\n", entryName)

	out, err = exec.Command("pwbook", "update", "entry1").Output()
	assertNoError(t, err)
	assertOutputContains(t, out, "entry1")
	t.Logf("Updated the entry %q\n", entryName)

	out, err = exec.Command("pwbook", "remove", "entry1").Output()
	assertNoError(t, err)
	assertOutputContains(t, out, "entry1")
	t.Logf("Removed the entry %q\n", entryName)

	out, err = exec.Command("pwbook", "list").Output()
	assertNoError(t, err)
	assertOutputContains(t, out, "Total 0 entries")
	t.Logf("Verified list not containing the entry %q\n", entryName)
}
