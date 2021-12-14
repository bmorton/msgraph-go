# \ExternalExternalApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExternalExternalGetExternal**](ExternalExternalApi.md#ExternalExternalGetExternal) | **Get** /external | Get external
[**ExternalExternalUpdateExternal**](ExternalExternalApi.md#ExternalExternalUpdateExternal) | **Patch** /external | Update external



## ExternalExternalGetExternal

> MicrosoftGraphExternalConnectorsExternal ExternalExternalGetExternal(ctx).Select_(select_).Expand(expand).Execute()

Get external

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
    resp, r, err := api_client.ExternalExternalApi.ExternalExternalGetExternal(context.Background()).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExternalExternalApi.ExternalExternalGetExternal``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExternalExternalGetExternal`: MicrosoftGraphExternalConnectorsExternal
    fmt.Fprintf(os.Stdout, "Response from `ExternalExternalApi.ExternalExternalGetExternal`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExternalExternalGetExternalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphExternalConnectorsExternal**](MicrosoftGraphExternalConnectorsExternal.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExternalExternalUpdateExternal

> ExternalExternalUpdateExternal(ctx).MicrosoftGraphExternalConnectorsExternal(microsoftGraphExternalConnectorsExternal).Execute()

Update external

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
    microsoftGraphExternalConnectorsExternal := *openapiclient.NewMicrosoftGraphExternalConnectorsExternal() // MicrosoftGraphExternalConnectorsExternal | New property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ExternalExternalApi.ExternalExternalUpdateExternal(context.Background()).MicrosoftGraphExternalConnectorsExternal(microsoftGraphExternalConnectorsExternal).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExternalExternalApi.ExternalExternalUpdateExternal``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExternalExternalUpdateExternalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphExternalConnectorsExternal** | [**MicrosoftGraphExternalConnectorsExternal**](MicrosoftGraphExternalConnectorsExternal.md) | New property values | 

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

