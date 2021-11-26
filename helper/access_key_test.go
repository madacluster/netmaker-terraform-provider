package helper

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/gravitl/netmaker/models"
)

func TestClient_CreateKey(t *testing.T) {
	type fields struct {
		HostURL    string
		HTTPClient *http.Client
		Token      string
		Auth       AuthStruct
	}
	type args struct {
		networkID string
		key       models.AccessKey
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.AccessKey
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
				networkID: "netmakertest",
				key: models.AccessKey{
					Name: "test",
					Uses: 10,
				},
			},
			want: &models.AccessKey{
				Name: "test",
				Uses: 10,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, _ := NewClient(&tt.fields.HostURL, &tt.fields.Auth.Username, &tt.fields.Auth.Password)
			got, err := c.CreateKey(tt.args.networkID, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.CreateKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.Name, tt.want.Name) {
				t.Errorf("Client.CreateKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
