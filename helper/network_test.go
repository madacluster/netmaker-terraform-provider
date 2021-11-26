package helper

import (
	"net/http"
	"reflect"
	"testing"
	"time"

	"github.com/gravitl/netmaker/models"
)

var host = "http://localhost:8081"
var pass = "mx4S6JsSg7JWcZ"
var user = "admin"

const ipRange = "10.101.0.0/24"

func TestClient_CreateNetwork(t *testing.T) {
	type fields struct {
		HostURL    string
		HTTPClient *http.Client
		Token      string
		Auth       AuthStruct
	}
	type args struct {
		networkID           string
		addressrange        string
		localrange          string
		islocal             string
		isdualstack         string
		addressrange6       string
		defaultudpholepunch string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.Network
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
			args{
				networkID:           "netmakertest",
				addressrange:        "10.101.10.0/24",
				localrange:          "",
				islocal:             "no",
				isdualstack:         "no",
				addressrange6:       "",
				defaultudpholepunch: "yes",
			},
			&models.Network{
				AddressRange: "10.101.10.0/24",
				NetID:        "netmakertest",
			},

			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, _ := NewClient(&tt.fields.HostURL, &tt.fields.Auth.Username, &tt.fields.Auth.Password)
			network := &models.Network{
				AddressRange:        tt.args.addressrange,
				LocalRange:          tt.args.localrange,
				IsLocal:             tt.args.islocal,
				IsDualStack:         tt.args.isdualstack,
				AddressRange6:       tt.args.addressrange6,
				DefaultUDPHolePunch: tt.args.defaultudpholepunch,
				NetID:               tt.args.networkID,
			}
			got, err := c.CreateNetwork(*network)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.CreateNetwork() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.NetID, tt.want.NetID) {
				t.Errorf("Client.CreateNetwork() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetNetworks(t *testing.T) {
	type fields struct {
		HostURL    string
		HTTPClient *http.Client
		Token      string
		Auth       AuthStruct
	}
	tests := []struct {
		name    string
		fields  fields
		want    []models.Network
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
			[]models.Network{
				{
					AddressRange: ipRange,
					NetID:        "netmakertest",
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, _ := NewClient(&tt.fields.HostURL, &tt.fields.Auth.Username, &tt.fields.Auth.Password)
			got, err := c.GetNetworks()
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.GetNetworks() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got[0].NetID, tt.want[0].NetID) {
				t.Errorf("Client.GetNetworks() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetNetwork(t *testing.T) {
	type fields struct {
		HostURL    string
		HTTPClient *http.Client
		Token      string
		Auth       AuthStruct
	}
	type args struct {
		networkID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.Network
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
			args{networkID: "netmakertest"},
			&models.Network{
				AddressRange:        "10.101.10.0/24",
				NetID:               "netmakertest",
				DefaultUDPHolePunch: "yes",
				IsDualStack:         "no",
				IsLocal:             "no",
			},

			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, _ := NewClient(&tt.fields.HostURL, &tt.fields.Auth.Username, &tt.fields.Auth.Password)

			got, err := c.GetNetwork(tt.args.networkID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.GetNetwork() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.AddressRange, tt.want.AddressRange) {
				t.Errorf("Client.GetNetwork() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_UpdateNetwork(t *testing.T) {
	type fields struct {
		HostURL    string
		HTTPClient *http.Client
		Token      string
		Auth       AuthStruct
	}
	type args struct {
		data map[string]string
	}
	c, err := NewClient(&host, &user, &pass)
	if err != nil {
		t.Errorf("Client.UpdateNetwork() error = %v", err)
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.Network
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
			args{
				data: map[string]string{
					"netid":        "netmakertest",
					"addressrange": "10.102.0.0/24",
				},
			},
			&models.Network{
				AddressRange: "10.102.0.0/24",
				NetID:        "netmakertest",
				DisplayName:  "netmakertest",
			},

			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			network, err := c.GetNetwork(tt.args.data["netid"])
			if err != nil {
				t.Errorf("Client.UpdateNetwork() error = %v", err)
			}

			networkFieldMap := mapFielsRevert(network)
			got, err := c.UpdateNetworkMap(networkFieldMap)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.UpdateNetwork() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.DisplayName, tt.want.DisplayName) {
				t.Errorf("Client.UpdateNetwork() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_DeleteNetwork(t *testing.T) {
	type fields struct {
		HostURL    string
		HTTPClient *http.Client
		Token      string
		Auth       AuthStruct
	}
	type args struct {
		networkID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
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
			args{networkID: "netmakertest"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, _ := NewClient(&tt.fields.HostURL, &tt.fields.Auth.Username, &tt.fields.Auth.Password)

			if err := c.DeleteNetwork(tt.args.networkID); (err != nil) != tt.wantErr {
				t.Errorf("Client.DeleteNetwork() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
