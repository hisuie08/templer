package output

import "testing"

func Test_stdoutOutput_WriteFile(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		data    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "test", data: "hello", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: construct the receiver type.
			var o stdoutOutput = stdoutOutput{w: t.Output()}
			gotErr := o.WriteFile("", []byte(tt.data))
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("WriteFile() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("WriteFile() succeeded unexpectedly")
			}
		})
	}
}


