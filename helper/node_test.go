package helper

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/gravitl/netmaker/models"
)

func CreateTestData() (*string, error) {
	c, err := NewClient(&host, &user, &pass)
	if err != nil {
		return nil, err
	}

	network := &models.Network{
		AddressRange:        "10.102.0.0/24",
		LocalRange:          "",
		IsLocal:             "no",
		IsDualStack:         "",
		AddressRange6:       "",
		DefaultUDPHolePunch: "yes",
		NetID:               "test2",
	}
	got, err := c.CreateNetwork(*network)
	if err != nil {
		return nil, err
	}
	key := &models.AccessKey{
		Name: "test",
		Uses: 10,
	}
	accessKey, err := c.CreateKey(got.NetID, *key)
	if err != nil {
		return nil, err
	}
	return &accessKey.Value, nil
}
func TestClient_CreateNetworkNode(t *testing.T) {
	key, err := CreateTestData()
	if err != nil {
		t.Errorf("Client.CreateNetworkNode() error = %v", err)

	}
	type fields struct {
		HostURL    string
		HTTPClient *http.Client
		Token      string
		Auth       AuthStruct
	}
	type args struct {
		networkID string
		node      models.Node
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.Node
		wantErr bool
	}{
		// TODO: Add test cases.
		{

			name: "Create Network Node",
			fields: fields{
				HostURL:    host,
				HTTPClient: &http.Client{},
				Auth: AuthStruct{
					Username: user,
					Password: pass,
				},
			},
			args: args{
				networkID: "test2",
				node: models.Node{

					AccessKey: *key,
					PublicKey: "DM5qhLAE20PG9BbfBCger+Ac9D2NDOwCtY1rbYDLf34=", Name: "testnode", Endpoint: "10.0.0.1", MacAddress: "01:02:03:04:05:06", Password: "password", Network: "skynet",
				},
			},
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
			got, err := c.CreateNetworkNode(tt.args.networkID, tt.args.node)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.CreateNetworkNode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.CreateNetworkNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetNodes(t *testing.T) {
	type fields struct {
		HostURL    string
		HTTPClient *http.Client
		Token      string
		Auth       AuthStruct
	}
	tests := []struct {
		name    string
		fields  fields
		want    []models.Node
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Get Nodes",
			fields: fields{
				HostURL:    host,
				HTTPClient: &http.Client{},
				Auth: AuthStruct{
					Username: user,
					Password: pass,
				},
			},
			want:    []models.Node{},
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
			got, err := c.GetNodes()
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.GetNodes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.GetNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetNetworkNodes(t *testing.T) {
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
		want    []models.Node
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				HostURL:    tt.fields.HostURL,
				HTTPClient: tt.fields.HTTPClient,
				Token:      tt.fields.Token,
				Auth:       tt.fields.Auth,
			}
			got, err := c.GetNetworkNodes(tt.args.networkID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.GetNetworkNodes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.GetNetworkNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetNode(t *testing.T) {
	type fields struct {
		HostURL    string
		HTTPClient *http.Client
		Token      string
		Auth       AuthStruct
	}
	type args struct {
		networkID string
		mac       string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    models.Node
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				HostURL:    tt.fields.HostURL,
				HTTPClient: tt.fields.HTTPClient,
				Token:      tt.fields.Token,
				Auth:       tt.fields.Auth,
			}
			got, err := c.GetNode(tt.args.networkID, tt.args.mac)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.GetNode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.GetNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetNetworkIngress(t *testing.T) {
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
		want    []models.Node
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				HostURL:    tt.fields.HostURL,
				HTTPClient: tt.fields.HTTPClient,
				Token:      tt.fields.Token,
				Auth:       tt.fields.Auth,
			}
			got, err := c.GetNetworkIngress(tt.args.networkID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.GetNetworkIngress() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.GetNetworkIngress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_CreateIngress(t *testing.T) {
	type fields struct {
		HostURL    string
		HTTPClient *http.Client
		Token      string
		Auth       AuthStruct
	}
	type args struct {
		networkID string
		mac       string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.Node
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				HostURL:    tt.fields.HostURL,
				HTTPClient: tt.fields.HTTPClient,
				Token:      tt.fields.Token,
				Auth:       tt.fields.Auth,
			}
			got, err := c.CreateIngress(tt.args.networkID, tt.args.mac)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.CreateIngress() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.CreateIngress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_DeleteIngress(t *testing.T) {
	type fields struct {
		HostURL    string
		HTTPClient *http.Client
		Token      string
		Auth       AuthStruct
	}
	type args struct {
		networkID string
		mac       string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.Node
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				HostURL:    tt.fields.HostURL,
				HTTPClient: tt.fields.HTTPClient,
				Token:      tt.fields.Token,
				Auth:       tt.fields.Auth,
			}
			got, err := c.DeleteIngress(tt.args.networkID, tt.args.mac)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.DeleteIngress() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.DeleteIngress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetNetworkEgress(t *testing.T) {
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
		want    []models.Node
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				HostURL:    tt.fields.HostURL,
				HTTPClient: tt.fields.HTTPClient,
				Token:      tt.fields.Token,
				Auth:       tt.fields.Auth,
			}
			got, err := c.GetNetworkEgress(tt.args.networkID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.GetNetworkEgress() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.GetNetworkEgress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_CreateEgress(t *testing.T) {
	type fields struct {
		HostURL    string
		HTTPClient *http.Client
		Token      string
		Auth       AuthStruct
	}
	type args struct {
		networkID string
		mac       string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.Node
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				HostURL:    tt.fields.HostURL,
				HTTPClient: tt.fields.HTTPClient,
				Token:      tt.fields.Token,
				Auth:       tt.fields.Auth,
			}
			got, err := c.CreateEgress(tt.args.networkID, tt.args.mac)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.CreateEgress() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.CreateEgress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_DeleteEgress(t *testing.T) {
	type fields struct {
		HostURL    string
		HTTPClient *http.Client
		Token      string
		Auth       AuthStruct
	}
	type args struct {
		networkID string
		mac       string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.Node
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				HostURL:    tt.fields.HostURL,
				HTTPClient: tt.fields.HTTPClient,
				Token:      tt.fields.Token,
				Auth:       tt.fields.Auth,
			}
			got, err := c.DeleteEgress(tt.args.networkID, tt.args.mac)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.DeleteEgress() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.DeleteEgress() = %v, want %v", got, tt.want)
			}
		})
	}
}
