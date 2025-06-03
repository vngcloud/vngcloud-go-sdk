# VngCloud Go SDK

[![Release VngCloud GoSDK project](https://github.com/vngcloud/vngcloud-go-sdk/actions/workflows/release_build.yml/badge.svg)](https://github.com/vngcloud/vngcloud-go-sdk/actions/workflows/release_build.yml)

<hr>

###### 🌈 Table of contents

- [Introduction](#introduction)
- [Usage](#usage)
- [Contributing](#contributing)

<hr>

# Introduction

- `vngcloud-go-sdk` is a Go SDK for VNG Cloud services. It helps you to interact with VNG Cloud services easily.

# Usage
- You can install the SDK by running the following command:
  ```bash
  go get github.com/vngcloud/vngcloud-go-sdk
  ```

- Now for example, imagine you want to list all available VngCloud load-balancer packages. You can implement this code in your Go Application:
  ```go
  package main

  import (
    "fmt"
    lctx "context"

    lsclient "github.com/vngcloud/vngcloud-go-sdk/v2/client"
    lslbv2 "github.com/vngcloud/vngcloud-go-sdk/v2/vngcloud/services/loadbalancer/v2"
  )

  func main() {
    vngcloud := validSdkConfig()
    opt := lslbv2.NewListLoadBalancerPackagesRequest()
    packages, sdkerr := vngcloud.VLBGateway().V2().LoadBalancerService().ListLoadBalancerPackages(opt)
    if sdkerr != nil {
      fmt.Printf("Expect nil but got %+v", sdkerr)
    }

    if packages == nil {
      fmt.Printf("Expect not nil but got nil")
    }

    for _, pkg := range packages.Items {
      fmt.Printf("Package: %+v", pkg)
    }
  }

  func validSdkConfig() lsclient.IClient {
    clientId, clientSecret := "__PUT_YOUR_CLIENT_ID__", "__PUT_YOUR_CLIENT_SECRET__"
    sdkConfig := lsclient.NewSdkConfigure().
      WithClientId(clientId).
      WithClientSecret(clientSecret).
      WithProjectId("__PUT_YOUR_VNGCLOUD_PROJECT_ID__").
      WithZoneId("65e12ffcb6d82cd39f8cf023").
      WithIamEndpoint("https://iamapis.vngcloud.vn/accounts-api").
      WithVServerEndpoint("https://hcm-3.api.vngcloud.vn/vserver/vserver-gateway").
      WithVLBEndpoint("https://hcm-3.api.vngcloud.vn/vserver/vlb-gateway").
      WithVNetworkEndpoint("https://vnetwork-hcm03.vngcloud.vn/vnetwork-gateway/vnetwork")

    return lsclient.NewClient(lctx.TODO()).WithRetryCount(1).WithSleep(10).Configure(sdkConfig)
  }
  ```

# Contributing

- To release a new version of the SDK, you need to create a new tag with the format `vX.Y.Z` and push it to the repository. The GitHub Actions will automatically build and release the new version to the GitHub Packages. For instance, to release version `v2.11.1`, you can run the following commands:

  ```bash
  git tag -am "[release] release new version" v2.11.0
  git push --tags
  ```

- To get the latest version of the SDK, you can run the following command:
  ```bash
  git tag -l --sort=-creatordate | head -n 1
  ```