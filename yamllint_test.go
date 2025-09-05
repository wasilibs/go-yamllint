package yamllint

import (
	"bytes"
	"strings"
	"testing"

	"github.com/wasilibs/go-yamllint/internal/runner"
)

func TestRuns(t *testing.T) {
	tests := []struct {
		file    string
		ret     int
		message string
	}{
		{
			file:    "test.yaml",
			ret:     0,
			message: "missing document start",
		},
		{
			file:    "error.yaml",
			ret:     1,
			message: "syntax error: mapping values are not allowed here",
		},
	}

	for _, tc := range tests {
		t.Run(tc.file, func(t *testing.T) {
			stdin := bytes.Buffer{}
			stdout := bytes.Buffer{}
			stderr := bytes.Buffer{}

			ret := runner.Run("yamllint", []string{"testdata/" + tc.file}, &stdin, &stdout, &stderr, ".")
			if ret != tc.ret {
				t.Fatalf("unexpected return code: have %d, want %d", ret, tc.ret)
			}

			if !strings.Contains(stdout.String(), "missing document start") {
				t.Fatalf("unexpected stdout: %s", stdout.String())
			}
		})
	}
}
