package jsonparse

import (
	"reflect"
	"testing"
)

func TestParseJSON(t *testing.T) {
	type args struct {
		body []byte
		urls *IncomingJSON
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
		want    *IncomingJSON
	}{
		// TODO: Add test cases.
		{
			name: "first",
			args: args{
				body: []byte(`{"url":["abc.com", "safasdf.com", "facebook.com","http://www.google.com/", "https://mail.ru/"]}`),
				urls: &IncomingJSON{},
			},
			wantErr: nil,
			want: &IncomingJSON{
				Url: []string{"abc.com", "safasdf.com", "facebook.com", "http://www.google.com/", "https://mail.ru/"},
			},
		},
		{
			name: "second",
			args: args{
				body: []byte(`{"map":52}`),
				urls: &IncomingJSON{},
			},
			wantErr: ErrParseJson,
			want: &IncomingJSON{
				Url: []string{},
			},
		},
		{
			name: "third",
			args: args{
				body: []byte(`{"urls":["abc.com", "safasdf.com"]}`),
				urls: &IncomingJSON{},
			},
			wantErr: ErrParseJson,
			want: &IncomingJSON{
				Url: []string{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// test errors

			if err := ParseJSON(tt.args.body, tt.args.urls); err != nil {

				if err != tt.wantErr {
					t.Errorf("ParseJSON() error = %v, wantErr %v", err, tt.wantErr)
				}
			} else {
				// test result
				got := tt.args.urls
				if ParseJSON(tt.args.body, got); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("ParseJSON() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}
