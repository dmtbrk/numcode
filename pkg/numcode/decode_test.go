package numcode

import "testing"

func TestDecode(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Hello", args{s: "H2ll4"}, "Hello"},
		{"Uppercase", args{s: "2TC"}, "eTC"},
		{"Empty", args{s: ""}, ""},
		{"No vowels", args{s: "BBC"}, "BBC"},
		{"Only vowels", args{s: "12345"}, "aeiou"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Decode(tt.args.s); got != tt.want {
				t.Errorf("Decode() = %q, want %q", got, tt.want)
			}
		})
	}
}
