package helper

import (
	"net/http"
	"reflect"
	"testing"
	"time"
)

var host = "https://api.netmaker.madacluster.tech"
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
		want    *Network
		wantErr bool
	}{
		// TODO: Add test cases.
		{"",
			fields{
				"https://api.netmaker.madacluster.tech",
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
			&Network{
				Addressrange: "10.101.10.0/24",
				Netid:        "netmakertest",
			},

			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, _ := NewClient(&tt.fields.HostURL, &tt.fields.Auth.Username, &tt.fields.Auth.Password)
			network := &Network{
				Addressrange:        tt.args.addressrange,
				Localrange:          tt.args.localrange,
				Islocal:             tt.args.islocal,
				Isdualstack:         tt.args.isdualstack,
				Addressrange6:       tt.args.addressrange6,
				Defaultudpholepunch: tt.args.defaultudpholepunch,
				Netid:               tt.args.networkID,
			}
			got, err := c.CreateNetwork(*network)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.CreateNetwork() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.Netid, tt.want.Netid) {
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
		want    []Network
		wantErr bool
	}{
		// TODO: Add test cases.
		{"",
			fields{
				"https://api.netmaker.madacluster.tech",
				&http.Client{Timeout: 10 * time.Second},
				"",
				AuthStruct{
					"admin",
					"mx4S6JsSg7JWcZ",
				},
			},
			[]Network{
				{
					Addressrange: ipRange,
					Netid:        "netmakertest",
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
			if !reflect.DeepEqual(got[0].Netid, tt.want[0].Netid) {
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
		want    *Network
		wantErr bool
	}{
		// TODO: Add test cases.
		{"",
			fields{
				"https://api.netmaker.madacluster.tech",
				&http.Client{Timeout: 10 * time.Second},
				"",
				AuthStruct{
					"admin",
					"mx4S6JsSg7JWcZ",
				},
			},
			args{networkID: "netmakertest"},
			&Network{
				Addressrange:        "10.101.10.0/24",
				Netid:               "netmakertest",
				Defaultudpholepunch: "yes",
				Isdualstack:         "no",
				Islocal:             "no",
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
			if !reflect.DeepEqual(got.Addressrange, tt.want.Addressrange) {
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
		want    *Network
		wantErr bool
	}{
		// TODO: Add test cases.
		{"",
			fields{
				"https://api.netmaker.madacluster.tech",
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
			&Network{
				Addressrange: "10.102.0.0/24",
				Netid:        "netmakertest",
				Displayname:  "netmakertest",
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
			if !reflect.DeepEqual(got.Displayname, tt.want.Displayname) {
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
				"https://api.netmaker.madacluster.tech",
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
