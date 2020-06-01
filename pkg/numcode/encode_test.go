package numcode

import "testing"

func TestEncode(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Hello", args{s: "Hello"}, "H2ll4"},
		{"Uppercase", args{s: "Etc"}, "2tc"},
		{"Empty", args{s: ""}, ""},
		{"No vowels", args{s: "BBC"}, "BBC"},
		{"Only vowels", args{s: "aeiouAEIOU"}, "1234512345"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Encode(tt.args.s); got != tt.want {
				t.Errorf("Encode() = %q, want %q", got, tt.want)
			}
		})
	}
}
