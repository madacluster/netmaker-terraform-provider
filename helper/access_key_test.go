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

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, _ := NewClient(&tt.fields.HostURL, &tt.fields.Auth.Username, &tt.fields.Auth.Password)
			got, err := c.CreateKey(tt.args.networkID, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.CreateKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.CreateKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
