package main

import "testing"

func Test_digitToRoman(t *testing.T) {
	type args struct {
		in int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "Test 1", args: args{in: 1}, want: "I"},
		{name: "Test 2", args: args{in: 11}, want: "XI"},
		{name: "Test 3", args: args{in: 9}, want: "IX"},
		{name: "Test 4", args: args{in: 4}, want: "IV"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := digitToRoman(tt.args.in); got != tt.want {
				t.Errorf("digitToRoman() = %v, want %v", got, tt.want)
			}
		})
	}
}
