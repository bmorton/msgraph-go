# \InformationProtectionInformationProtectionApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InformationProtectionInformationProtectionGetInformationProtection**](InformationProtectionInformationProtectionApi.md#InformationProtectionInformationProtectionGetInformationProtection) | **Get** /informationProtection | Get informationProtection
[**InformationProtectionInformationProtectionUpdateInformationProtection**](InformationProtectionInformationProtectionApi.md#InformationProtectionInformationProtectionUpdateInformationProtection) | **Patch** /informationProtection | Update informationProtection



## InformationProtectionInformationProtectionGetInformationProtection

> MicrosoftGraphInformationProtection InformationProtectionInformationProtectionGetInformationProtection(ctx).Select_(select_).Expand(expand).Execute()

Get informationProtection

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
    resp, r, err := api_client.InformationProtectionInformationProtectionApi.InformationProtectionInformationProtectionGetInformationProtection(context.Background()).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InformationProtectionInformationProtectionApi.InformationProtectionInformationProtectionGetInformationProtection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InformationProtectionInformationProtectionGetInformationProtection`: MicrosoftGraphInformationProtection
    fmt.Fprintf(os.Stdout, "Response from `InformationProtectionInformationProtectionApi.InformationProtectionInformationProtectionGetInformationProtection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInformationProtectionInformationProtectionGetInformationProtectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphInformationProtection**](MicrosoftGraphInformationProtection.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InformationProtectionInformationProtectionUpdateInformationProtection

> InformationProtectionInformationProtectionUpdateInformationProtection(ctx).MicrosoftGraphInformationProtection(microsoftGraphInformationProtection).Execute()

Update informationProtection

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
    microsoftGraphInformationProtection := *openapiclient.NewMicrosoftGraphInformationProtection() // MicrosoftGraphInformationProtection | New property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InformationProtectionInformationProtectionApi.InformationProtectionInformationProtectionUpdateInformationProtection(context.Background()).MicrosoftGraphInformationProtection(microsoftGraphInformationProtection).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InformationProtectionInformationProtectionApi.InformationProtectionInformationProtectionUpdateInformationProtection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInformationProtectionInformationProtectionUpdateInformationProtectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphInformationProtection** | [**MicrosoftGraphInformationProtection**](MicrosoftGraphInformationProtection.md) | New property values | 

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

