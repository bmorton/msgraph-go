# \SearchActionsApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SearchQuery**](SearchActionsApi.md#SearchQuery) | **Post** /search/microsoft.graph.query | Invoke action query



## SearchQuery

> []MicrosoftGraphSearchResponse SearchQuery(ctx).InlineObject714(inlineObject714).Execute()

Invoke action query

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
    inlineObject714 := *openapiclient.NewInlineObject714() // InlineObject714 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SearchActionsApi.SearchQuery(context.Background()).InlineObject714(inlineObject714).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchActionsApi.SearchQuery``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchQuery`: []MicrosoftGraphSearchResponse
    fmt.Fprintf(os.Stdout, "Response from `SearchActionsApi.SearchQuery`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject714** | [**InlineObject714**](InlineObject714.md) |  | 

### Return type

[**[]MicrosoftGraphSearchResponse**](MicrosoftGraphSearchResponse.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

