# \CommunicationsCloudCommunicationsApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CommunicationsCloudCommunicationsGetCloudCommunications**](CommunicationsCloudCommunicationsApi.md#CommunicationsCloudCommunicationsGetCloudCommunications) | **Get** /communications | Get communications
[**CommunicationsCloudCommunicationsUpdateCloudCommunications**](CommunicationsCloudCommunicationsApi.md#CommunicationsCloudCommunicationsUpdateCloudCommunications) | **Patch** /communications | Update communications



## CommunicationsCloudCommunicationsGetCloudCommunications

> MicrosoftGraphCloudCommunications CommunicationsCloudCommunicationsGetCloudCommunications(ctx).Select_(select_).Expand(expand).Execute()

Get communications

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
    resp, r, err := api_client.CommunicationsCloudCommunicationsApi.CommunicationsCloudCommunicationsGetCloudCommunications(context.Background()).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommunicationsCloudCommunicationsApi.CommunicationsCloudCommunicationsGetCloudCommunications``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CommunicationsCloudCommunicationsGetCloudCommunications`: MicrosoftGraphCloudCommunications
    fmt.Fprintf(os.Stdout, "Response from `CommunicationsCloudCommunicationsApi.CommunicationsCloudCommunicationsGetCloudCommunications`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCommunicationsCloudCommunicationsGetCloudCommunicationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphCloudCommunications**](MicrosoftGraphCloudCommunications.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommunicationsCloudCommunicationsUpdateCloudCommunications

> CommunicationsCloudCommunicationsUpdateCloudCommunications(ctx).MicrosoftGraphCloudCommunications(microsoftGraphCloudCommunications).Execute()

Update communications

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
    microsoftGraphCloudCommunications := *openapiclient.NewMicrosoftGraphCloudCommunications() // MicrosoftGraphCloudCommunications | New property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CommunicationsCloudCommunicationsApi.CommunicationsCloudCommunicationsUpdateCloudCommunications(context.Background()).MicrosoftGraphCloudCommunications(microsoftGraphCloudCommunications).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommunicationsCloudCommunicationsApi.CommunicationsCloudCommunicationsUpdateCloudCommunications``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCommunicationsCloudCommunicationsUpdateCloudCommunicationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphCloudCommunications** | [**MicrosoftGraphCloudCommunications**](MicrosoftGraphCloudCommunications.md) | New property values | 

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

