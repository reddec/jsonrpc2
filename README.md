# JSON-RPC 2 Go supporting library

[![Documentation](https://img.shields.io/badge/documentation-latest-green)](https://godoc.org/github.com/reddec/jsonrpc2)
[![license](https://img.shields.io/github/license/reddec/jsonrpc2.svg)](https://github.com/reddec/jsonrpc2)
[![](https://godoc.org/github.com/reddec/jsonrpc2?status.svg)](http://godoc.org/github.com/reddec/jsonrpc2)
[![donate](https://img.shields.io/badge/help_by️-donate❤-ff69b4)](http://reddec.net/about/#donate)
[![Download](https://api.bintray.com/packages/reddec/debian/jsonrpc2/images/download.svg)](https://bintray.com/reddec/debian/jsonrpc2/_latestVersion)

* [Formal specification](https://www.jsonrpc.org/specification)

The library aims to bring JSON-RPC 2.0 support to Go. Goals:

* Type safe by code-generation
* Reasonable good performance
* Clean, extendable and easy-to use interface
* Protocol-agnostic solution with adapters for common-cases (HTTP, TCP, etc...)


# Installation

* (recommended) look at  [releases](https://github.com/reddec/jsonrpc2/releases) page and download
* build from source `go get -v github.com/reddec/jsonrpc2/cmd/...`
* From bintray repository for most **debian**-based distribution (`trusty`, `xenial`, `bionic`, `buster`, `wheezy`):
```bash
sudo apt-key adv --keyserver hkp://keyserver.ubuntu.com:80 --recv-keys 379CE192D401AB61
echo "deb https://dl.bintray.com/reddec/debian {distribution} main" | sudo tee -a /etc/apt/sources.list
sudo apt install jsonrpc2
```

## Build requirements

* go 1.13+


# Usage as library

Please see [package documentation](https://godoc.org/github.com/reddec/jsonrpc2)


# Usage as CLI for type-safe generation

```
Usage:
  jsonrpc2-gen [OPTIONS]

Generate tiny wrapper for JSON-RPC router
Author: Baryshnikov Aleksandr <dev@baryshnikov.net>
Version: dev

Application Options:
  -i, --file=      File to scan [$GOFILE]
  -I, --interface= Interface to wrap [$INTERFACE]
      --namespace= Custom namespace for functions. If not defined - interface name will be used [$NAMESPACE]
      --wrapper=   Wrapper function name. If not defined - Register<interface> name will be used [$WRAPPER]
  -o, --output=    Generated output destination (- means STDOUT) (default: -) [$OUTPUT]
  -p, --package=   Package name (can be override by output dir) (default: events) [$PACKAGE]
  -d, --doc=       Generate markdown documentation [$DOC]
  -c, --case=[keep|camel|pascal|snake|kebab] Method name case style (default: keep) [$CASE]

Help Options:
  -h, --help       Show this help message
```

## Example

Assume you have an interface file (`user.go`) like this:

```go
package abc

// General user profile access
type User interface {
	// Get user profile
	Profile(token string) (*Profile, error)
}

```

Just invoke `jsonrpc2-gen -i user.go -o user_gen.go -I User -p abc`

You will get `user_gen.go` file like that:


```go
// Code generated by jsonrpc2-gen. DO NOT EDIT.
//go:generate jsonrpc2-gen -i user.go -o user_gen.go -I User -p abc
package abc

import (
	"encoding/json"
	jsonrpc2 "github.com/reddec/jsonrpc2"
)

func RegisterUser(router *jsonrpc2.Router, wrap User) []string {
	router.RegisterFunc("User.Profile", func(params json.RawMessage, positional bool) (interface{}, error) {
		var args struct {
			Arg0 string `json:"token"`
		}
		var err error
		if positional {
			err = jsonrpc2.UnmarshalArray(params, &args.Arg0)
		} else {
			err = json.Unmarshal(params, &args)
		}
		if err != nil {
			return nil, err
		}
		return wrap.Profile(args.Arg0)
	})

	return []string{"User.Profile"}
}
```

### Generate documentation

Add `-doc user.md` to generate documentations as described bellow. It will be generated and saved to the provided file (`user.md`) 

```markdown
# User

General user profile access


## User.Profile

Get user profile

* Method: `User.Profile`
* Returns: `*Profile`
* Arguments:

| Position | Name | Type |
|----------|------|------|
| 0 | token | `string` |
```
