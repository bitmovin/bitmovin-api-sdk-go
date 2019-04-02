# Bitmovin Go API client

***Important! This is an alpha version, there may be breaking changes.***

[![bitmovin](http://bitmovin-a.akamaihd.net/webpages/bitmovin-logo-github.png)](http://www.bitmovin.com)

Go-Client which enables you to seamlessly integrate the Bitmovin API into your projects. Using this API client requires an active account.

[Sign up for a Bitmovin Account!](https://dashboard.bitmovin.com/signup)

The full API reference can be found [here](https://bitmovin.com/docs).

## Installation:

```bash
go get github.com/bitmovin/bitmovin-api-sdk-go
```

## Initializiation

```go
package main

import (
    "github.com/bitmovin/bitmovin-api-sdk-go"
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
)

func main() {
    bitmovinApi, err := bitmovin.NewBitmovinApi(func(apiClient *common.ApiClient) {
        apiClient.ApiKey = "<YOUR_API_KEY>"
    })
    if err != nil {
        panic(err)
    }

   // ...
}
```
