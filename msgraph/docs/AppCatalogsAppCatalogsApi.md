# \AppCatalogsAppCatalogsApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppCatalogsAppCatalogsGetAppCatalogs**](AppCatalogsAppCatalogsApi.md#AppCatalogsAppCatalogsGetAppCatalogs) | **Get** /appCatalogs | Get appCatalogs
[**AppCatalogsAppCatalogsUpdateAppCatalogs**](AppCatalogsAppCatalogsApi.md#AppCatalogsAppCatalogsUpdateAppCatalogs) | **Patch** /appCatalogs | Update appCatalogs



## AppCatalogsAppCatalogsGetAppCatalogs

> MicrosoftGraphAppCatalogs AppCatalogsAppCatalogsGetAppCatalogs(ctx).Select_(select_).Expand(expand).Execute()

Get appCatalogs

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
    resp, r, err := api_client.AppCatalogsAppCatalogsApi.AppCatalogsAppCatalogsGetAppCatalogs(context.Background()).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppCatalogsAppCatalogsApi.AppCatalogsAppCatalogsGetAppCatalogs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppCatalogsAppCatalogsGetAppCatalogs`: MicrosoftGraphAppCatalogs
    fmt.Fprintf(os.Stdout, "Response from `AppCatalogsAppCatalogsApi.AppCatalogsAppCatalogsGetAppCatalogs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppCatalogsAppCatalogsGetAppCatalogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphAppCatalogs**](MicrosoftGraphAppCatalogs.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppCatalogsAppCatalogsUpdateAppCatalogs

> AppCatalogsAppCatalogsUpdateAppCatalogs(ctx).MicrosoftGraphAppCatalogs(microsoftGraphAppCatalogs).Execute()

Update appCatalogs

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
    microsoftGraphAppCatalogs := *openapiclient.NewMicrosoftGraphAppCatalogs() // MicrosoftGraphAppCatalogs | New property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AppCatalogsAppCatalogsApi.AppCatalogsAppCatalogsUpdateAppCatalogs(context.Background()).MicrosoftGraphAppCatalogs(microsoftGraphAppCatalogs).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppCatalogsAppCatalogsApi.AppCatalogsAppCatalogsUpdateAppCatalogs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppCatalogsAppCatalogsUpdateAppCatalogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphAppCatalogs** | [**MicrosoftGraphAppCatalogs**](MicrosoftGraphAppCatalogs.md) | New property values | 

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

