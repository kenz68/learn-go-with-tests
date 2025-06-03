package main

import "testing"

func Test_isPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "empty string",
			args: args{s: ""},
			want: true,
		},
		{
			name: "single character",
			args: args{s: "a"},
			want: true,
		},
		{
			name: "two characters, palindrome",
			args: args{s: "aa"},
			want: true,
		},
		{
			name: "two characters, not palindrome",
			args: args{s: "ab"},
			want: false,
		},
		{
			name: "odd length palindrome",
			args: args{s: "racecar"},
			want: true,
		},
		{
			name: "special characters, palindrome",
			args: args{s: "@#$%^&*"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.s); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
