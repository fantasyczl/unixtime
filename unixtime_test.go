package main

import "testing"

func Test_isDayFormat(t *testing.T) {
	type args struct {
		dateStr string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "valid day format",
			args: args{dateStr: "2023-10-01"},
			want: true,
		},
		{
			name: "valid day format with time",
			args: args{dateStr: "2023-10-01 12:34:56"},
			want: false,
		},
		{
			name: "invalid format with extra characters",
			args: args{dateStr: "2023-10-01T12:34:56"},
			want: false,
		},
		{
			name: "invalid format with missing parts",
			args: args{dateStr: "2023-10"},
			want: false,
		},
		{
			name: "invalid format with wrong separators",
			args: args{dateStr: "2023/10/01"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isDayFormat(tt.args.dateStr); got != tt.want {
				t.Errorf("isDayFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}
