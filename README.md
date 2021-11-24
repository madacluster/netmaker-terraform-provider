# Terraform Provider Netmaker
[![CodeFactor](https://www.codefactor.io/repository/github/madacluster/netmaker-terraform-provider/badge)](https://www.codefactor.io/repository/github/madacluster/netmaker-terraform-provider)

This repository is a *template* for a [Terraform](https://www.terraform.io) provider. It is intended as a starting point for creating Terraform providers, containing:

 - A resource, and a data source (`internal/provider/`),
 - Examples (`examples/`) and generated documentation (`docs/`),
 - Miscellaneous meta files.
 
These files contain boilerplate code that you will need to edit to create your own Terraform provider. A full guide to creating Terraform providers can be found at [Writing Custom Providers](https://www.terraform.io/docs/extend/writing-custom-providers.html).

Please see the [GitHub template repository documentation](https://help.github.com/en/github/creating-cloning-and-archiving-repositories/creating-a-repository-from-a-template) for how to create a new repository from this template on GitHub.

Once you've written your provider, you'll want to [publish it on the Terraform Registry](https://www.terraform.io/docs/registry/providers/publishing.html) so that others can use it.

## Features
**Networks API**
- [x] Get All Networks: /api/networks, GET

- [x] Create Network: /api/network, POST

- [x] Get Network: /api/networks/{network id}, GET

- [x] Update Network: /api/networks/{network id}, PUT

- [x] Delete Network: /api/networks/{network id}, DELETE

**Access Keys API**

- [ ] Get All Keys: /api/networks/{network id}/keys, GET

- [ ] Create Key: /api/networks/{network id}/keys, GET

- [ ] Delete Key: /api/networks/{network id}/keys/{keyname}, DELETE

**Nodes API**

- [ ] Create a Gateway: /api/nodes/{network id}/{macaddress}/creategateway, POST

- [ ] Delete a Gateway: /api/nodes/{network id}/{macaddress}/deletegateway, DELETE

- [ ] Get Network Nodes: /api/nodes/{network id}, GET

- [ ] Get Node: /api/nodes/{network id}/{macaddress}, GET

- [ ] Update Node: /api/nodes/{network id}/{macaddress}, PUT

- [ ] Delete Node: /api/nodes/{network id}/{macaddress}, DELETE

**Users API**
Note: Only able to create Admin user at this time. The "user" is only used by the user interface to authenticate the single admin user.

- [ ] Get User: /api/users/{username}, GET

- [ ] Update User: /api/users/{username}, PUT

- [ ] Delete User: /api/users/{username}, DELETE

- [ ] Check for Admin User: /api/users/adm/hasadmin, GET

- [ ] Create Admin User: /api/users/adm/createadmin, POST

- [x] Authenticate: /api/users/adm/authenticate, POST


## Requirements

-	[Terraform](https://www.terraform.io/downloads.html) >= 0.13.x
-	[Go](https://golang.org/doc/install) >= 1.15

## Building The Provider

1. Clone the repository
1. Enter the repository directory
1. Build the provider using the Go `install` command: 
```sh
$ go install
```

## Adding Dependencies

This provider uses [Go modules](https://github.com/golang/go/wiki/Modules).
Please see the Go documentation for the most up to date information about using Go modules.

To add a new dependency `github.com/author/dependency` to your Terraform provider:

```
go get github.com/author/dependency
go mod tidy
```

Then commit the changes to `go.mod` and `go.sum`.

## Using the provider

Fill this in for each provider

## Developing the Provider

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (see [Requirements](#requirements) above).

To compile the provider, run `go install`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

To generate or update documentation, run `go generate`.

In order to run the full suite of Acceptance tests, run `make testacc`.

*Note:* Acceptance tests create real resources, and often cost money to run.

```sh
$ make testacc
```

