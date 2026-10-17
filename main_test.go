package main

import "testing"

func TestGreet(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want string
	}{
		{name: "no name", arg: "", want: "Hello, World!"},
		{name: "with name", arg: "Gopher", want: "Hello, Gopher!"},
		{name: "name with space", arg: "Go 语言", want: "Hello, Go 语言!"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := greet(tt.arg); got != tt.want {
				t.Errorf("greet(%q) = %q, want %q", tt.arg, got, tt.want)
			}
		})
	}
}
