# \ServicePrincipalsEndpointApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ServicePrincipalsCreateEndpoints**](ServicePrincipalsEndpointApi.md#ServicePrincipalsCreateEndpoints) | **Post** /servicePrincipals/{servicePrincipal-id}/endpoints | Create new navigation property to endpoints for servicePrincipals
[**ServicePrincipalsDeleteEndpoints**](ServicePrincipalsEndpointApi.md#ServicePrincipalsDeleteEndpoints) | **Delete** /servicePrincipals/{servicePrincipal-id}/endpoints/{endpoint-id} | Delete navigation property endpoints for servicePrincipals
[**ServicePrincipalsGetEndpoints**](ServicePrincipalsEndpointApi.md#ServicePrincipalsGetEndpoints) | **Get** /servicePrincipals/{servicePrincipal-id}/endpoints/{endpoint-id} | Get endpoints from servicePrincipals
[**ServicePrincipalsListEndpoints**](ServicePrincipalsEndpointApi.md#ServicePrincipalsListEndpoints) | **Get** /servicePrincipals/{servicePrincipal-id}/endpoints | Get endpoints from servicePrincipals
[**ServicePrincipalsUpdateEndpoints**](ServicePrincipalsEndpointApi.md#ServicePrincipalsUpdateEndpoints) | **Patch** /servicePrincipals/{servicePrincipal-id}/endpoints/{endpoint-id} | Update the navigation property endpoints in servicePrincipals



## ServicePrincipalsCreateEndpoints

> MicrosoftGraphEndpoint ServicePrincipalsCreateEndpoints(ctx, servicePrincipalId).MicrosoftGraphEndpoint(microsoftGraphEndpoint).Execute()

Create new navigation property to endpoints for servicePrincipals



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
    microsoftGraphEndpoint := *openapiclient.NewMicrosoftGraphEndpoint() // MicrosoftGraphEndpoint | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServicePrincipalsEndpointApi.ServicePrincipalsCreateEndpoints(context.Background(), servicePrincipalId).MicrosoftGraphEndpoint(microsoftGraphEndpoint).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicePrincipalsEndpointApi.ServicePrincipalsCreateEndpoints``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServicePrincipalsCreateEndpoints`: MicrosoftGraphEndpoint
    fmt.Fprintf(os.Stdout, "Response from `ServicePrincipalsEndpointApi.ServicePrincipalsCreateEndpoints`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**servicePrincipalId** | **string** | key: id of servicePrincipal | 

### Other Parameters

Other parameters are passed through a pointer to a apiServicePrincipalsCreateEndpointsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphEndpoint** | [**MicrosoftGraphEndpoint**](MicrosoftGraphEndpoint.md) | New navigation property | 

### Return type

[**MicrosoftGraphEndpoint**](MicrosoftGraphEndpoint.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServicePrincipalsDeleteEndpoints

> ServicePrincipalsDeleteEndpoints(ctx, servicePrincipalId, endpointId).IfMatch(ifMatch).Execute()

Delete navigation property endpoints for servicePrincipals



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
    endpointId := "endpointId_example" // string | key: id of endpoint
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServicePrincipalsEndpointApi.ServicePrincipalsDeleteEndpoints(context.Background(), servicePrincipalId, endpointId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicePrincipalsEndpointApi.ServicePrincipalsDeleteEndpoints``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**servicePrincipalId** | **string** | key: id of servicePrincipal | 
**endpointId** | **string** | key: id of endpoint | 

### Other Parameters

Other parameters are passed through a pointer to a apiServicePrincipalsDeleteEndpointsRequest struct via the builder pattern


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


## ServicePrincipalsGetEndpoints

> MicrosoftGraphEndpoint ServicePrincipalsGetEndpoints(ctx, servicePrincipalId, endpointId).Select_(select_).Expand(expand).Execute()

Get endpoints from servicePrincipals



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
    endpointId := "endpointId_example" // string | key: id of endpoint
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServicePrincipalsEndpointApi.ServicePrincipalsGetEndpoints(context.Background(), servicePrincipalId, endpointId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicePrincipalsEndpointApi.ServicePrincipalsGetEndpoints``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServicePrincipalsGetEndpoints`: MicrosoftGraphEndpoint
    fmt.Fprintf(os.Stdout, "Response from `ServicePrincipalsEndpointApi.ServicePrincipalsGetEndpoints`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**servicePrincipalId** | **string** | key: id of servicePrincipal | 
**endpointId** | **string** | key: id of endpoint | 

### Other Parameters

Other parameters are passed through a pointer to a apiServicePrincipalsGetEndpointsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphEndpoint**](MicrosoftGraphEndpoint.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServicePrincipalsListEndpoints

> CollectionOfEndpoint ServicePrincipalsListEndpoints(ctx, servicePrincipalId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get endpoints from servicePrincipals



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
    resp, r, err := api_client.ServicePrincipalsEndpointApi.ServicePrincipalsListEndpoints(context.Background(), servicePrincipalId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicePrincipalsEndpointApi.ServicePrincipalsListEndpoints``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServicePrincipalsListEndpoints`: CollectionOfEndpoint
    fmt.Fprintf(os.Stdout, "Response from `ServicePrincipalsEndpointApi.ServicePrincipalsListEndpoints`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**servicePrincipalId** | **string** | key: id of servicePrincipal | 

### Other Parameters

Other parameters are passed through a pointer to a apiServicePrincipalsListEndpointsRequest struct via the builder pattern


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

[**CollectionOfEndpoint**](CollectionOfEndpoint.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServicePrincipalsUpdateEndpoints

> ServicePrincipalsUpdateEndpoints(ctx, servicePrincipalId, endpointId).MicrosoftGraphEndpoint(microsoftGraphEndpoint).Execute()

Update the navigation property endpoints in servicePrincipals



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
    endpointId := "endpointId_example" // string | key: id of endpoint
    microsoftGraphEndpoint := *openapiclient.NewMicrosoftGraphEndpoint() // MicrosoftGraphEndpoint | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServicePrincipalsEndpointApi.ServicePrincipalsUpdateEndpoints(context.Background(), servicePrincipalId, endpointId).MicrosoftGraphEndpoint(microsoftGraphEndpoint).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicePrincipalsEndpointApi.ServicePrincipalsUpdateEndpoints``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**servicePrincipalId** | **string** | key: id of servicePrincipal | 
**endpointId** | **string** | key: id of endpoint | 

### Other Parameters

Other parameters are passed through a pointer to a apiServicePrincipalsUpdateEndpointsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **microsoftGraphEndpoint** | [**MicrosoftGraphEndpoint**](MicrosoftGraphEndpoint.md) | New navigation property values | 

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

