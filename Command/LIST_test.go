package Command

import (
	"testing"
)

func Test_formatFileList(t *testing.T) {
	type args struct {
		fileList string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"test2", args{fileList: `drwxr-xr-x 2 anubis anubis   4096 4月   9 19:09 tmp`}, "drwxr-xr-x 2 anubis anubis   4096 04   9 19:09 tmp\n"}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := formatFileList(tt.args.fileList); got != tt.want {
				t.Errorf("formatFileList() = %v, want %v", got, tt.want)
			}
		})
	}
}
