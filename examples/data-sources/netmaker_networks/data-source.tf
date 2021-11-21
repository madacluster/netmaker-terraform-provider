terraform {
  required_providers {
    netmaker = {
      version = "0.2"
      source  = "github.com/madacluster/netmaker"
    }
  }
}

provider "netmaker" {
  username = "admin"
  password = "mx4S6JsSg7JWcZ"
}

data "netmaker_networks" "example" {
  # sample_attribute = "foo"
}