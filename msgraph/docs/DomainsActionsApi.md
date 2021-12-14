# \DomainsActionsApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DomainsDomainForceDelete**](DomainsActionsApi.md#DomainsDomainForceDelete) | **Post** /domains/{domain-id}/microsoft.graph.forceDelete | Invoke action forceDelete
[**DomainsDomainVerify**](DomainsActionsApi.md#DomainsDomainVerify) | **Post** /domains/{domain-id}/microsoft.graph.verify | Invoke action verify



## DomainsDomainForceDelete

> DomainsDomainForceDelete(ctx, domainId).InlineObject134(inlineObject134).Execute()

Invoke action forceDelete

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
    domainId := "domainId_example" // string | key: id of domain
    inlineObject134 := *openapiclient.NewInlineObject134() // InlineObject134 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DomainsActionsApi.DomainsDomainForceDelete(context.Background(), domainId).InlineObject134(inlineObject134).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainsActionsApi.DomainsDomainForceDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **string** | key: id of domain | 

### Other Parameters

Other parameters are passed through a pointer to a apiDomainsDomainForceDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject134** | [**InlineObject134**](InlineObject134.md) |  | 

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


## DomainsDomainVerify

> AnyOfmicrosoftGraphDomain DomainsDomainVerify(ctx, domainId).Execute()

Invoke action verify

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
    domainId := "domainId_example" // string | key: id of domain

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DomainsActionsApi.DomainsDomainVerify(context.Background(), domainId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainsActionsApi.DomainsDomainVerify``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DomainsDomainVerify`: AnyOfmicrosoftGraphDomain
    fmt.Fprintf(os.Stdout, "Response from `DomainsActionsApi.DomainsDomainVerify`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **string** | key: id of domain | 

### Other Parameters

Other parameters are passed through a pointer to a apiDomainsDomainVerifyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AnyOfmicrosoftGraphDomain**](anyOf&lt;microsoft.graph.domain&gt;.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

