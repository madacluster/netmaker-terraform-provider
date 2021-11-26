package helper

import (
	"net/http"
	"reflect"
	"testing"
	"time"
)

func TestClient_SignIn(t *testing.T) {
	type fields struct {
		HostURL    string
		HTTPClient *http.Client
		Token      string
		Auth       AuthStruct
	}

	tests := []struct {
		name    string
		fields  fields
		want    *AuthResponse
		wantErr bool
	}{
		// TODO: Add test cases.
		{"",
			fields{
				"http://localhost:8081",
				&http.Client{Timeout: 10 * time.Second},
				"",
				AuthStruct{
					"admin",
					"mx4S6JsSg7JWcZ",
				},
			},
			&AuthResponse{
				Code: 200,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				HostURL:    tt.fields.HostURL,
				HTTPClient: tt.fields.HTTPClient,
				Token:      tt.fields.Token,
				Auth:       tt.fields.Auth,
			}
			got, err := c.SignIn()
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.SignIn() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.Code, tt.want.Code) {
				t.Errorf("Client.SignIn() = %v, want %v", got, tt.want)
			}
		})
	}
}
