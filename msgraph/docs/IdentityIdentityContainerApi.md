# \IdentityIdentityContainerApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IdentityIdentityContainerGetIdentityContainer**](IdentityIdentityContainerApi.md#IdentityIdentityContainerGetIdentityContainer) | **Get** /identity | Get identity
[**IdentityIdentityContainerUpdateIdentityContainer**](IdentityIdentityContainerApi.md#IdentityIdentityContainerUpdateIdentityContainer) | **Patch** /identity | Update identity



## IdentityIdentityContainerGetIdentityContainer

> MicrosoftGraphIdentityContainer IdentityIdentityContainerGetIdentityContainer(ctx).Select_(select_).Expand(expand).Execute()

Get identity

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IdentityIdentityContainerApi.IdentityIdentityContainerGetIdentityContainer(context.Background()).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityIdentityContainerApi.IdentityIdentityContainerGetIdentityContainer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IdentityIdentityContainerGetIdentityContainer`: MicrosoftGraphIdentityContainer
    fmt.Fprintf(os.Stdout, "Response from `IdentityIdentityContainerApi.IdentityIdentityContainerGetIdentityContainer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIdentityIdentityContainerGetIdentityContainerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphIdentityContainer**](MicrosoftGraphIdentityContainer.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IdentityIdentityContainerUpdateIdentityContainer

> IdentityIdentityContainerUpdateIdentityContainer(ctx).MicrosoftGraphIdentityContainer(microsoftGraphIdentityContainer).Execute()

Update identity

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    microsoftGraphIdentityContainer := *openapiclient.NewMicrosoftGraphIdentityContainer() // MicrosoftGraphIdentityContainer | New property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IdentityIdentityContainerApi.IdentityIdentityContainerUpdateIdentityContainer(context.Background()).MicrosoftGraphIdentityContainer(microsoftGraphIdentityContainer).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityIdentityContainerApi.IdentityIdentityContainerUpdateIdentityContainer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIdentityIdentityContainerUpdateIdentityContainerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphIdentityContainer** | [**MicrosoftGraphIdentityContainer**](MicrosoftGraphIdentityContainer.md) | New property values | 

### Return type

 (empty response body)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

