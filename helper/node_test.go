package helper

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/gravitl/netmaker/models"
)

var token string

const node_id = "testnode"
const net_id = "test"
const key_name = "test"
const node_mac = "01:02:03:04:05:06"

func CreateTestData(t *testing.T, createNode bool) {
	c, err := NewClient(&host, &user, &pass)
	if err != nil {
		t.Fatal(err)
	}

	network := &models.Network{
		AddressRange:        "10.102.0.0/24",
		LocalRange:          "",
		IsLocal:             "no",
		IsDualStack:         "",
		AddressRange6:       "",
		DefaultUDPHolePunch: "yes",
		NetID:               net_id,
	}
	got, err := c.CreateNetwork(*network)
	if err != nil {
		t.Fatal(err)
	}
	key := &models.AccessKey{
		Name: key_name,
		Uses: 10,
	}
	accessKey, err := c.CreateKey(got.NetID, *key)
	token = accessKey.Value
	// token = accessKey.
	if err != nil {
		t.Fatal(err)
	}
	if createNode {
		node := models.Node{
			AccessKey: token,
			PublicKey: "DM5qhLAE20PG9BbfBCger+Ac9D2NDOwCtY1rbYDLf34=", Name: node_id, Endpoint: "10.0.0.1", MacAddress: node_mac, Password: "password", Network: net_id,
		}
		_, err := c.CreateNetworkNode(network.NetID, node)
		if err != nil {
			t.Fatal(err)
		}
	}
	t.Cleanup(func() {
		err := CleanTestData()
		if err != nil {
			t.Fatal(err)
		}
	})
}

func CleanTestData() error {
	c, err := NewClient(&host, &user, &pass)
	if err != nil {
		return err
	}
	err = c.DeleteNetworkNode(net_id, node_mac)
	if err != nil {
		return err
	}
	err = c.DeleteKey(net_id, key_name)
	if err != nil {
		return err
	}
	return c.DeleteNetwork(net_id)
}
func TestClient_CreateNetworkNode(t *testing.T) {
	CreateTestData(t, false)
	// if err != nil {
	// 	t.Errorf("Client.CreateNetworkNode() error = %v", err)

	node := models.Node{
		AccessKey: token,
		PublicKey: "DM5qhLAE20PG9BbfBCger+Ac9D2NDOwCtY1rbYDLf34=", Name: node_id, Endpoint: "10.0.0.1", MacAddress: node_mac, Password: "password", Network: net_id,
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
				networkID: net_id,
				node:      node,
			},
			want:    &node,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, err := NewClient(&host, &user, &pass)
			if err != nil {
				t.Errorf("Client.CreateNetworkNode() error = %v", err)
			}
			got, err := c.CreateNetworkNode(tt.args.networkID, tt.args.node)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.CreateNetworkNode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.Network, net_id) {
				t.Errorf("Client.CreateNetworkNode() = %v, want %v", got.Network, net_id)
			}
		})
	}
}

func TestClient_GetNodes(t *testing.T) {
	CreateTestData(t, true)
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
			c, err := NewClient(&host, &user, &pass)
			if err != nil {
				t.Errorf("Client.NetClient() error = %v", err)
			}
			got, err := c.GetNodes()
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.GetNodes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got[0].Name, "testnode") {
				t.Errorf("Client.GetNodes() = %v, want %v", got[0].Name, "testnode")
			}
		})
	}
}

func TestClient_GetNetworkNodes(t *testing.T) {
	CreateTestData(t, true)

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
		{
			name: "Get Nodes Network",
			fields: fields{
				HostURL:    host,
				HTTPClient: &http.Client{},
				Auth: AuthStruct{
					Username: user,
					Password: pass,
				},
			},
			args: args{
				networkID: net_id,
			},
			want:    []models.Node{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, err := NewClient(&host, &user, &pass)
			if err != nil {
				t.Errorf("Client.NetClient() error = %v", err)
			}
			got, err := c.GetNetworkNodes(tt.args.networkID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.GetNodes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got[0].Name, "testnode") {
				t.Errorf("Client.GetNodes() = %v, want %v", got[0].Name, "testnode")
			}
		})
	}
}

func TestClient_GetNode(t *testing.T) {
	CreateTestData(t, true)
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
		{
			name:   "Get Node",
			fields: fields{},
			args: args{
				networkID: net_id,
				mac:       node_mac,
			},
			want: models.Node{
				Name: node_id,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, err := NewClient(&host, &user, &pass)
			if err != nil {
				t.Errorf("Client.NetClient() error = %v", err)
			}
			got, err := c.GetNode(tt.args.networkID, tt.args.mac)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.GetNode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.Name, tt.want.Name) {
				t.Errorf("Client.GetNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetNetworkIngress(t *testing.T) {
	CreateTestData(t, true)
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
		{
			name:   "Get Network Ingress",
			fields: fields{},
			args: args{
				networkID: net_id,
			},
			want:    []models.Node{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, err := NewClient(&host, &user, &pass)
			if err != nil {
				t.Errorf("Client.NetClient() error = %v", err)
				return
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
	CreateTestData(t, true)
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
		{
			name:   "Create Ingress",
			fields: fields{},
			args: args{
				networkID: net_id,
				mac:       node_mac,
			},
			want: &models.Node{
				Name: node_id,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, err := NewClient(&host, &user, &pass)
			if err != nil {
				t.Errorf("Client.NetClient() error = %v", err)
			}
			got, err := c.CreateIngress(tt.args.networkID, tt.args.mac)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.CreateIngress() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.Name, tt.want.Name) {
				t.Errorf("Client.CreateIngress() = %v, want %v", got, tt.want.Name)
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
