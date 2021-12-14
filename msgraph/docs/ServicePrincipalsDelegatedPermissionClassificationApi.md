# \ServicePrincipalsDelegatedPermissionClassificationApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ServicePrincipalsCreateDelegatedPermissionClassifications**](ServicePrincipalsDelegatedPermissionClassificationApi.md#ServicePrincipalsCreateDelegatedPermissionClassifications) | **Post** /servicePrincipals/{servicePrincipal-id}/delegatedPermissionClassifications | Create new navigation property to delegatedPermissionClassifications for servicePrincipals
[**ServicePrincipalsDeleteDelegatedPermissionClassifications**](ServicePrincipalsDelegatedPermissionClassificationApi.md#ServicePrincipalsDeleteDelegatedPermissionClassifications) | **Delete** /servicePrincipals/{servicePrincipal-id}/delegatedPermissionClassifications/{delegatedPermissionClassification-id} | Delete navigation property delegatedPermissionClassifications for servicePrincipals
[**ServicePrincipalsGetDelegatedPermissionClassifications**](ServicePrincipalsDelegatedPermissionClassificationApi.md#ServicePrincipalsGetDelegatedPermissionClassifications) | **Get** /servicePrincipals/{servicePrincipal-id}/delegatedPermissionClassifications/{delegatedPermissionClassification-id} | Get delegatedPermissionClassifications from servicePrincipals
[**ServicePrincipalsListDelegatedPermissionClassifications**](ServicePrincipalsDelegatedPermissionClassificationApi.md#ServicePrincipalsListDelegatedPermissionClassifications) | **Get** /servicePrincipals/{servicePrincipal-id}/delegatedPermissionClassifications | Get delegatedPermissionClassifications from servicePrincipals
[**ServicePrincipalsUpdateDelegatedPermissionClassifications**](ServicePrincipalsDelegatedPermissionClassificationApi.md#ServicePrincipalsUpdateDelegatedPermissionClassifications) | **Patch** /servicePrincipals/{servicePrincipal-id}/delegatedPermissionClassifications/{delegatedPermissionClassification-id} | Update the navigation property delegatedPermissionClassifications in servicePrincipals



## ServicePrincipalsCreateDelegatedPermissionClassifications

> MicrosoftGraphDelegatedPermissionClassification ServicePrincipalsCreateDelegatedPermissionClassifications(ctx, servicePrincipalId).MicrosoftGraphDelegatedPermissionClassification(microsoftGraphDelegatedPermissionClassification).Execute()

Create new navigation property to delegatedPermissionClassifications for servicePrincipals



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
    servicePrincipalId := "servicePrincipalId_example" // string | key: id of servicePrincipal
    microsoftGraphDelegatedPermissionClassification := *openapiclient.NewMicrosoftGraphDelegatedPermissionClassification() // MicrosoftGraphDelegatedPermissionClassification | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServicePrincipalsDelegatedPermissionClassificationApi.ServicePrincipalsCreateDelegatedPermissionClassifications(context.Background(), servicePrincipalId).MicrosoftGraphDelegatedPermissionClassification(microsoftGraphDelegatedPermissionClassification).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicePrincipalsDelegatedPermissionClassificationApi.ServicePrincipalsCreateDelegatedPermissionClassifications``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServicePrincipalsCreateDelegatedPermissionClassifications`: MicrosoftGraphDelegatedPermissionClassification
    fmt.Fprintf(os.Stdout, "Response from `ServicePrincipalsDelegatedPermissionClassificationApi.ServicePrincipalsCreateDelegatedPermissionClassifications`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**servicePrincipalId** | **string** | key: id of servicePrincipal | 

### Other Parameters

Other parameters are passed through a pointer to a apiServicePrincipalsCreateDelegatedPermissionClassificationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphDelegatedPermissionClassification** | [**MicrosoftGraphDelegatedPermissionClassification**](MicrosoftGraphDelegatedPermissionClassification.md) | New navigation property | 

### Return type

[**MicrosoftGraphDelegatedPermissionClassification**](MicrosoftGraphDelegatedPermissionClassification.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServicePrincipalsDeleteDelegatedPermissionClassifications

> ServicePrincipalsDeleteDelegatedPermissionClassifications(ctx, servicePrincipalId, delegatedPermissionClassificationId).IfMatch(ifMatch).Execute()

Delete navigation property delegatedPermissionClassifications for servicePrincipals



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
    servicePrincipalId := "servicePrincipalId_example" // string | key: id of servicePrincipal
    delegatedPermissionClassificationId := "delegatedPermissionClassificationId_example" // string | key: id of delegatedPermissionClassification
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServicePrincipalsDelegatedPermissionClassificationApi.ServicePrincipalsDeleteDelegatedPermissionClassifications(context.Background(), servicePrincipalId, delegatedPermissionClassificationId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicePrincipalsDelegatedPermissionClassificationApi.ServicePrincipalsDeleteDelegatedPermissionClassifications``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**servicePrincipalId** | **string** | key: id of servicePrincipal | 
**delegatedPermissionClassificationId** | **string** | key: id of delegatedPermissionClassification | 

### Other Parameters

Other parameters are passed through a pointer to a apiServicePrincipalsDeleteDelegatedPermissionClassificationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ifMatch** | **string** | ETag | 

### Return type

 (empty response body)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServicePrincipalsGetDelegatedPermissionClassifications

> MicrosoftGraphDelegatedPermissionClassification ServicePrincipalsGetDelegatedPermissionClassifications(ctx, servicePrincipalId, delegatedPermissionClassificationId).Select_(select_).Expand(expand).Execute()

Get delegatedPermissionClassifications from servicePrincipals



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
    servicePrincipalId := "servicePrincipalId_example" // string | key: id of servicePrincipal
    delegatedPermissionClassificationId := "delegatedPermissionClassificationId_example" // string | key: id of delegatedPermissionClassification
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServicePrincipalsDelegatedPermissionClassificationApi.ServicePrincipalsGetDelegatedPermissionClassifications(context.Background(), servicePrincipalId, delegatedPermissionClassificationId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicePrincipalsDelegatedPermissionClassificationApi.ServicePrincipalsGetDelegatedPermissionClassifications``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServicePrincipalsGetDelegatedPermissionClassifications`: MicrosoftGraphDelegatedPermissionClassification
    fmt.Fprintf(os.Stdout, "Response from `ServicePrincipalsDelegatedPermissionClassificationApi.ServicePrincipalsGetDelegatedPermissionClassifications`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**servicePrincipalId** | **string** | key: id of servicePrincipal | 
**delegatedPermissionClassificationId** | **string** | key: id of delegatedPermissionClassification | 

### Other Parameters

Other parameters are passed through a pointer to a apiServicePrincipalsGetDelegatedPermissionClassificationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphDelegatedPermissionClassification**](MicrosoftGraphDelegatedPermissionClassification.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServicePrincipalsListDelegatedPermissionClassifications

> CollectionOfDelegatedPermissionClassification ServicePrincipalsListDelegatedPermissionClassifications(ctx, servicePrincipalId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get delegatedPermissionClassifications from servicePrincipals



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
    servicePrincipalId := "servicePrincipalId_example" // string | key: id of servicePrincipal
    top := int32(50) // int32 | Show only the first n items (optional)
    skip := int32(56) // int32 | Skip the first n items (optional)
    search := "search_example" // string | Search items by search phrases (optional)
    filter := "filter_example" // string | Filter items by property values (optional)
    count := true // bool | Include count of items (optional)
    orderby := []string{"Orderby_example"} // []string | Order items by property values (optional)
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServicePrincipalsDelegatedPermissionClassificationApi.ServicePrincipalsListDelegatedPermissionClassifications(context.Background(), servicePrincipalId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicePrincipalsDelegatedPermissionClassificationApi.ServicePrincipalsListDelegatedPermissionClassifications``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServicePrincipalsListDelegatedPermissionClassifications`: CollectionOfDelegatedPermissionClassification
    fmt.Fprintf(os.Stdout, "Response from `ServicePrincipalsDelegatedPermissionClassificationApi.ServicePrincipalsListDelegatedPermissionClassifications`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**servicePrincipalId** | **string** | key: id of servicePrincipal | 

### Other Parameters

Other parameters are passed through a pointer to a apiServicePrincipalsListDelegatedPermissionClassificationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **top** | **int32** | Show only the first n items | 
 **skip** | **int32** | Skip the first n items | 
 **search** | **string** | Search items by search phrases | 
 **filter** | **string** | Filter items by property values | 
 **count** | **bool** | Include count of items | 
 **orderby** | **[]string** | Order items by property values | 
 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**CollectionOfDelegatedPermissionClassification**](CollectionOfDelegatedPermissionClassification.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServicePrincipalsUpdateDelegatedPermissionClassifications

> ServicePrincipalsUpdateDelegatedPermissionClassifications(ctx, servicePrincipalId, delegatedPermissionClassificationId).MicrosoftGraphDelegatedPermissionClassification(microsoftGraphDelegatedPermissionClassification).Execute()

Update the navigation property delegatedPermissionClassifications in servicePrincipals



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
    servicePrincipalId := "servicePrincipalId_example" // string | key: id of servicePrincipal
    delegatedPermissionClassificationId := "delegatedPermissionClassificationId_example" // string | key: id of delegatedPermissionClassification
    microsoftGraphDelegatedPermissionClassification := *openapiclient.NewMicrosoftGraphDelegatedPermissionClassification() // MicrosoftGraphDelegatedPermissionClassification | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServicePrincipalsDelegatedPermissionClassificationApi.ServicePrincipalsUpdateDelegatedPermissionClassifications(context.Background(), servicePrincipalId, delegatedPermissionClassificationId).MicrosoftGraphDelegatedPermissionClassification(microsoftGraphDelegatedPermissionClassification).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicePrincipalsDelegatedPermissionClassificationApi.ServicePrincipalsUpdateDelegatedPermissionClassifications``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**servicePrincipalId** | **string** | key: id of servicePrincipal | 
**delegatedPermissionClassificationId** | **string** | key: id of delegatedPermissionClassification | 

### Other Parameters

Other parameters are passed through a pointer to a apiServicePrincipalsUpdateDelegatedPermissionClassificationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **microsoftGraphDelegatedPermissionClassification** | [**MicrosoftGraphDelegatedPermissionClassification**](MicrosoftGraphDelegatedPermissionClassification.md) | New navigation property values | 

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

