# \SearchSearchEntityApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SearchSearchEntityGetSearchEntity**](SearchSearchEntityApi.md#SearchSearchEntityGetSearchEntity) | **Get** /search | Get search
[**SearchSearchEntityUpdateSearchEntity**](SearchSearchEntityApi.md#SearchSearchEntityUpdateSearchEntity) | **Patch** /search | Update search



## SearchSearchEntityGetSearchEntity

> MicrosoftGraphSearchEntity SearchSearchEntityGetSearchEntity(ctx).Select_(select_).Expand(expand).Execute()

Get search

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
    resp, r, err := api_client.SearchSearchEntityApi.SearchSearchEntityGetSearchEntity(context.Background()).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchSearchEntityApi.SearchSearchEntityGetSearchEntity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchSearchEntityGetSearchEntity`: MicrosoftGraphSearchEntity
    fmt.Fprintf(os.Stdout, "Response from `SearchSearchEntityApi.SearchSearchEntityGetSearchEntity`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchSearchEntityGetSearchEntityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphSearchEntity**](MicrosoftGraphSearchEntity.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchSearchEntityUpdateSearchEntity

> SearchSearchEntityUpdateSearchEntity(ctx).MicrosoftGraphSearchEntity(microsoftGraphSearchEntity).Execute()

Update search

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
    microsoftGraphSearchEntity := *openapiclient.NewMicrosoftGraphSearchEntity() // MicrosoftGraphSearchEntity | New property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SearchSearchEntityApi.SearchSearchEntityUpdateSearchEntity(context.Background()).MicrosoftGraphSearchEntity(microsoftGraphSearchEntity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchSearchEntityApi.SearchSearchEntityUpdateSearchEntity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchSearchEntityUpdateSearchEntityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphSearchEntity** | [**MicrosoftGraphSearchEntity**](MicrosoftGraphSearchEntity.md) | New property values | 

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

