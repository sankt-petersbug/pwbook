package pwbook

import (
	"os/exec"
	"testing"
    "strings"
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

    out, err = exec.Command("pwbook", "add", "entry1").Output()
    assertNoError(t, err)
    assertOutputContains(t, out, "entry1")

    out, err = exec.Command("pwbook", "list").Output()
    assertNoError(t, err)
    assertOutputContains(t, out, "entry1")
    assertOutputContains(t, out, "Total 1 entries")

    out, err = exec.Command("pwbook", "update", "entry1").Output()
    assertNoError(t, err)
    assertOutputContains(t, out, "entry1")

    out, err = exec.Command("pwbook", "remove", "entry1").Output()
    assertNoError(t, err)
    assertOutputContains(t, out, "entry1")

    out, err = exec.Command("pwbook", "list").Output()
    assertNoError(t, err)
    assertOutputContains(t, out, "Total 0 entries")
}