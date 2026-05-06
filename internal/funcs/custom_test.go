package funcs_test

import (
	"templer/internal/funcs"
	"testing"
)

func TestCall(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		fname   string
		fn      string
		args    map[string]any
		want    any
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "simple add", fname: "add", fn: "a + b",
			args: map[string]any{"a": 2, "b": 4}, want: 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := funcs.Call(tt.fname, tt.fn, tt.args)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("Call() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("Call() succeeded unexpectedly")
			}
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Errorf("Call() = %v, want %v", got, tt.want)
			}
		})
	}
}
