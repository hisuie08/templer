package output

import (
	"os"
	"path/filepath"
	"testing"
)

func Test_fileOutput_WriteFile(t *testing.T) {
	td := t.TempDir()
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		path    string
		data    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "current", path: filepath.Join(td, "./file.txt"), data: "", wantErr: false},
		{name: "nested", path: filepath.Join(td, "sub", "file.txt"), data: "", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: construct the receiver type.
			var o fileOutput
			gotErr := o.WriteFile(tt.path, []byte(tt.data))
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("WriteFile() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("WriteFile() succeeded unexpectedly")
			}
			if _, e := os.Stat(tt.path); e != nil {
				t.Fatalf("file not created: %s", tt.path)
			}
		})
	}
}
