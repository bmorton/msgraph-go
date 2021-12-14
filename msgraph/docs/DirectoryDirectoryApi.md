# \DirectoryDirectoryApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DirectoryDirectoryGetDirectory**](DirectoryDirectoryApi.md#DirectoryDirectoryGetDirectory) | **Get** /directory | Get directory
[**DirectoryDirectoryUpdateDirectory**](DirectoryDirectoryApi.md#DirectoryDirectoryUpdateDirectory) | **Patch** /directory | Update directory



## DirectoryDirectoryGetDirectory

> MicrosoftGraphDirectory DirectoryDirectoryGetDirectory(ctx).Select_(select_).Expand(expand).Execute()

Get directory

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
    resp, r, err := api_client.DirectoryDirectoryApi.DirectoryDirectoryGetDirectory(context.Background()).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DirectoryDirectoryApi.DirectoryDirectoryGetDirectory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DirectoryDirectoryGetDirectory`: MicrosoftGraphDirectory
    fmt.Fprintf(os.Stdout, "Response from `DirectoryDirectoryApi.DirectoryDirectoryGetDirectory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDirectoryDirectoryGetDirectoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphDirectory**](MicrosoftGraphDirectory.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DirectoryDirectoryUpdateDirectory

> DirectoryDirectoryUpdateDirectory(ctx).MicrosoftGraphDirectory(microsoftGraphDirectory).Execute()

Update directory

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
    microsoftGraphDirectory := *openapiclient.NewMicrosoftGraphDirectory() // MicrosoftGraphDirectory | New property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DirectoryDirectoryApi.DirectoryDirectoryUpdateDirectory(context.Background()).MicrosoftGraphDirectory(microsoftGraphDirectory).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DirectoryDirectoryApi.DirectoryDirectoryUpdateDirectory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDirectoryDirectoryUpdateDirectoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphDirectory** | [**MicrosoftGraphDirectory**](MicrosoftGraphDirectory.md) | New property values | 

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

