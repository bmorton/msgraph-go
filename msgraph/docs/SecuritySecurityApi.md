# \SecuritySecurityApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SecuritySecurityGetSecurity**](SecuritySecurityApi.md#SecuritySecurityGetSecurity) | **Get** /security | Get security
[**SecuritySecurityUpdateSecurity**](SecuritySecurityApi.md#SecuritySecurityUpdateSecurity) | **Patch** /security | Update security



## SecuritySecurityGetSecurity

> MicrosoftGraphSecurity SecuritySecurityGetSecurity(ctx).Select_(select_).Expand(expand).Execute()

Get security

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
    resp, r, err := api_client.SecuritySecurityApi.SecuritySecurityGetSecurity(context.Background()).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecuritySecurityApi.SecuritySecurityGetSecurity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SecuritySecurityGetSecurity`: MicrosoftGraphSecurity
    fmt.Fprintf(os.Stdout, "Response from `SecuritySecurityApi.SecuritySecurityGetSecurity`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSecuritySecurityGetSecurityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphSecurity**](MicrosoftGraphSecurity.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecuritySecurityUpdateSecurity

> SecuritySecurityUpdateSecurity(ctx).MicrosoftGraphSecurity(microsoftGraphSecurity).Execute()

Update security

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
    microsoftGraphSecurity := *openapiclient.NewMicrosoftGraphSecurity() // MicrosoftGraphSecurity | New property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SecuritySecurityApi.SecuritySecurityUpdateSecurity(context.Background()).MicrosoftGraphSecurity(microsoftGraphSecurity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecuritySecurityApi.SecuritySecurityUpdateSecurity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSecuritySecurityUpdateSecurityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphSecurity** | [**MicrosoftGraphSecurity**](MicrosoftGraphSecurity.md) | New property values | 

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

