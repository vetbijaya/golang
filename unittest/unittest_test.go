package main

import (
	"testing"
)

func Test_addition(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "success - test addiition",
			args: args{a: []int{1, 2, 3, 4}},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addition(tt.args.a...); got != tt.want {
				t.Errorf("addition() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_multiply(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{

		{
			name: "success- test multiply",
			args: args{a: []int{2, 3, 4}},
			want: 24,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := multiply(tt.args.a...); got != tt.want {
				t.Errorf("multiply() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_openFile(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		
		{
			name: "success- test open file",
			args: args{filename: "unittest.go"},
			wantErr: false,
		},
		{
			name: "fail -test open file",
			args: args{filename: "unittest1.go"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := openFile(tt.args.filename); (err != nil) != tt.wantErr {
				t.Errorf("openFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
