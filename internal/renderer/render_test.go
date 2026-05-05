package renderer

import "testing"

func Test_fixForOut(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		str  string
		want string
	}{
		{name: "fix", str: "hello", want: "hello\n"},
		{name: "no fix", str: "hello\n", want: "hello\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := fixForOut(tt.str)
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Fatalf("want: %s but got %s", tt.want, got)
			}
		})
	}
}