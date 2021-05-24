package parsing

import "testing"

func Test_correctUrl(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "first",
			args: args{
				url: "httttps://google.com////",
			},
			want: "https://www.google.com",
		},
		{
			name: "second",
			args: args{
				url: "google.com",
			},
			want: "https://www.google.com",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := correctUrl(tt.args.url); got != tt.want {
				t.Errorf("correctUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}
