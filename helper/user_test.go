package helper

import (
	"net/http"
	"testing"
)

func TestClient_CreateAdmin(t *testing.T) {
	type fields struct {
		HostURL    string
		HTTPClient *http.Client
		Token      string
		Auth       AuthStruct
	}
	type args struct {
		user User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Create admin user",
			fields: fields{
				HostURL:    host,
				HTTPClient: &http.Client{},
				Token:      "",
				Auth: AuthStruct{
					Username: user,
					Password: pass,
				},
			},
			args: args{
				user: User{
					UserName: "admin_test",
					Password: pass,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				HostURL:    "http://localhost:8081",
				HTTPClient: tt.fields.HTTPClient,
				Token:      tt.fields.Token,
				Auth:       tt.fields.Auth,
			}
			if err := c.CreateAdmin(tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("Client.CreateAdmin() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClient_CheckAdmin(t *testing.T) {
	type fields struct {
		HostURL    string
		HTTPClient *http.Client
		Token      string
		Auth       AuthStruct
	}
	tests := []struct {
		name    string
		fields  fields
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Check admin user",
			fields: fields{
				HostURL:    host,
				HTTPClient: &http.Client{},
				Token:      "",
				Auth: AuthStruct{
					Username: user,
					Password: pass,
				},
			},
			want:    true,
			wantErr: false,
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
			got, err := c.CheckAdmin()
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.CheckAdmin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Client.CheckAdmin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_CreateUser(t *testing.T) {
	type fields struct {
		HostURL    string
		HTTPClient *http.Client
		Token      string
		Auth       AuthStruct
	}
	type args struct {
		user User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Create user",
			fields: fields{
				HostURL:    host,
				HTTPClient: &http.Client{},
				Token:      "",
				Auth: AuthStruct{
					Username: user,
					Password: pass,
				},
			},
			args: args{
				user: User{
					UserName: "user_test",
					Password: pass,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, _ := NewClient(&tt.fields.HostURL, &tt.fields.Auth.Username, &tt.fields.Auth.Password)

			if err := c.CreateUser(tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("Client.CreateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClient_DeleteUser(t *testing.T) {
	type fields struct {
		HostURL    string
		HTTPClient *http.Client
		Token      string
		Auth       AuthStruct
	}
	type args struct {
		user User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Delete user",
			fields: fields{
				HostURL:    host,
				HTTPClient: &http.Client{},
				Token:      "",
				Auth: AuthStruct{
					Username: user,
					Password: pass,
				},
			},
			args: args{
				user: User{
					UserName: "user_test",
					Password: pass,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, _ := NewClient(&tt.fields.HostURL, &tt.fields.Auth.Username, &tt.fields.Auth.Password)

			if err := c.DeleteUser(tt.args.user.UserName); (err != nil) != tt.wantErr {
				t.Errorf("Client.DeleteUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
