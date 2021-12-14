# \PermissionGrantsResourceSpecificPermissionGrantApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PermissionGrantsResourceSpecificPermissionGrantCreateResourceSpecificPermissionGrant**](PermissionGrantsResourceSpecificPermissionGrantApi.md#PermissionGrantsResourceSpecificPermissionGrantCreateResourceSpecificPermissionGrant) | **Post** /permissionGrants | Add new entity to permissionGrants
[**PermissionGrantsResourceSpecificPermissionGrantDeleteResourceSpecificPermissionGrant**](PermissionGrantsResourceSpecificPermissionGrantApi.md#PermissionGrantsResourceSpecificPermissionGrantDeleteResourceSpecificPermissionGrant) | **Delete** /permissionGrants/{resourceSpecificPermissionGrant-id} | Delete entity from permissionGrants
[**PermissionGrantsResourceSpecificPermissionGrantGetResourceSpecificPermissionGrant**](PermissionGrantsResourceSpecificPermissionGrantApi.md#PermissionGrantsResourceSpecificPermissionGrantGetResourceSpecificPermissionGrant) | **Get** /permissionGrants/{resourceSpecificPermissionGrant-id} | Get entity from permissionGrants by key
[**PermissionGrantsResourceSpecificPermissionGrantListResourceSpecificPermissionGrant**](PermissionGrantsResourceSpecificPermissionGrantApi.md#PermissionGrantsResourceSpecificPermissionGrantListResourceSpecificPermissionGrant) | **Get** /permissionGrants | Get entities from permissionGrants
[**PermissionGrantsResourceSpecificPermissionGrantUpdateResourceSpecificPermissionGrant**](PermissionGrantsResourceSpecificPermissionGrantApi.md#PermissionGrantsResourceSpecificPermissionGrantUpdateResourceSpecificPermissionGrant) | **Patch** /permissionGrants/{resourceSpecificPermissionGrant-id} | Update entity in permissionGrants



## PermissionGrantsResourceSpecificPermissionGrantCreateResourceSpecificPermissionGrant

> MicrosoftGraphResourceSpecificPermissionGrant PermissionGrantsResourceSpecificPermissionGrantCreateResourceSpecificPermissionGrant(ctx).MicrosoftGraphResourceSpecificPermissionGrant(microsoftGraphResourceSpecificPermissionGrant).Execute()

Add new entity to permissionGrants

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
    microsoftGraphResourceSpecificPermissionGrant := *openapiclient.NewMicrosoftGraphResourceSpecificPermissionGrant() // MicrosoftGraphResourceSpecificPermissionGrant | New entity

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PermissionGrantsResourceSpecificPermissionGrantApi.PermissionGrantsResourceSpecificPermissionGrantCreateResourceSpecificPermissionGrant(context.Background()).MicrosoftGraphResourceSpecificPermissionGrant(microsoftGraphResourceSpecificPermissionGrant).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PermissionGrantsResourceSpecificPermissionGrantApi.PermissionGrantsResourceSpecificPermissionGrantCreateResourceSpecificPermissionGrant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PermissionGrantsResourceSpecificPermissionGrantCreateResourceSpecificPermissionGrant`: MicrosoftGraphResourceSpecificPermissionGrant
    fmt.Fprintf(os.Stdout, "Response from `PermissionGrantsResourceSpecificPermissionGrantApi.PermissionGrantsResourceSpecificPermissionGrantCreateResourceSpecificPermissionGrant`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPermissionGrantsResourceSpecificPermissionGrantCreateResourceSpecificPermissionGrantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphResourceSpecificPermissionGrant** | [**MicrosoftGraphResourceSpecificPermissionGrant**](MicrosoftGraphResourceSpecificPermissionGrant.md) | New entity | 

### Return type

[**MicrosoftGraphResourceSpecificPermissionGrant**](MicrosoftGraphResourceSpecificPermissionGrant.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PermissionGrantsResourceSpecificPermissionGrantDeleteResourceSpecificPermissionGrant

> PermissionGrantsResourceSpecificPermissionGrantDeleteResourceSpecificPermissionGrant(ctx, resourceSpecificPermissionGrantId).IfMatch(ifMatch).Execute()

Delete entity from permissionGrants

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
    resourceSpecificPermissionGrantId := "resourceSpecificPermissionGrantId_example" // string | key: id of resourceSpecificPermissionGrant
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PermissionGrantsResourceSpecificPermissionGrantApi.PermissionGrantsResourceSpecificPermissionGrantDeleteResourceSpecificPermissionGrant(context.Background(), resourceSpecificPermissionGrantId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PermissionGrantsResourceSpecificPermissionGrantApi.PermissionGrantsResourceSpecificPermissionGrantDeleteResourceSpecificPermissionGrant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceSpecificPermissionGrantId** | **string** | key: id of resourceSpecificPermissionGrant | 

### Other Parameters

Other parameters are passed through a pointer to a apiPermissionGrantsResourceSpecificPermissionGrantDeleteResourceSpecificPermissionGrantRequest struct via the builder pattern


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


## PermissionGrantsResourceSpecificPermissionGrantGetResourceSpecificPermissionGrant

> MicrosoftGraphResourceSpecificPermissionGrant PermissionGrantsResourceSpecificPermissionGrantGetResourceSpecificPermissionGrant(ctx, resourceSpecificPermissionGrantId).Select_(select_).Expand(expand).Execute()

Get entity from permissionGrants by key

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
    resourceSpecificPermissionGrantId := "resourceSpecificPermissionGrantId_example" // string | key: id of resourceSpecificPermissionGrant
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PermissionGrantsResourceSpecificPermissionGrantApi.PermissionGrantsResourceSpecificPermissionGrantGetResourceSpecificPermissionGrant(context.Background(), resourceSpecificPermissionGrantId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PermissionGrantsResourceSpecificPermissionGrantApi.PermissionGrantsResourceSpecificPermissionGrantGetResourceSpecificPermissionGrant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PermissionGrantsResourceSpecificPermissionGrantGetResourceSpecificPermissionGrant`: MicrosoftGraphResourceSpecificPermissionGrant
    fmt.Fprintf(os.Stdout, "Response from `PermissionGrantsResourceSpecificPermissionGrantApi.PermissionGrantsResourceSpecificPermissionGrantGetResourceSpecificPermissionGrant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceSpecificPermissionGrantId** | **string** | key: id of resourceSpecificPermissionGrant | 

### Other Parameters

Other parameters are passed through a pointer to a apiPermissionGrantsResourceSpecificPermissionGrantGetResourceSpecificPermissionGrantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphResourceSpecificPermissionGrant**](MicrosoftGraphResourceSpecificPermissionGrant.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PermissionGrantsResourceSpecificPermissionGrantListResourceSpecificPermissionGrant

> CollectionOfResourceSpecificPermissionGrant PermissionGrantsResourceSpecificPermissionGrantListResourceSpecificPermissionGrant(ctx).Search(search).Filter(filter).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get entities from permissionGrants

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
    search := "search_example" // string | Search items by search phrases (optional)
    filter := "filter_example" // string | Filter items by property values (optional)
    orderby := []string{"Orderby_example"} // []string | Order items by property values (optional)
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PermissionGrantsResourceSpecificPermissionGrantApi.PermissionGrantsResourceSpecificPermissionGrantListResourceSpecificPermissionGrant(context.Background()).Search(search).Filter(filter).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PermissionGrantsResourceSpecificPermissionGrantApi.PermissionGrantsResourceSpecificPermissionGrantListResourceSpecificPermissionGrant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PermissionGrantsResourceSpecificPermissionGrantListResourceSpecificPermissionGrant`: CollectionOfResourceSpecificPermissionGrant
    fmt.Fprintf(os.Stdout, "Response from `PermissionGrantsResourceSpecificPermissionGrantApi.PermissionGrantsResourceSpecificPermissionGrantListResourceSpecificPermissionGrant`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPermissionGrantsResourceSpecificPermissionGrantListResourceSpecificPermissionGrantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **string** | Search items by search phrases | 
 **filter** | **string** | Filter items by property values | 
 **orderby** | **[]string** | Order items by property values | 
 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**CollectionOfResourceSpecificPermissionGrant**](CollectionOfResourceSpecificPermissionGrant.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PermissionGrantsResourceSpecificPermissionGrantUpdateResourceSpecificPermissionGrant

> PermissionGrantsResourceSpecificPermissionGrantUpdateResourceSpecificPermissionGrant(ctx, resourceSpecificPermissionGrantId).MicrosoftGraphResourceSpecificPermissionGrant(microsoftGraphResourceSpecificPermissionGrant).Execute()

Update entity in permissionGrants

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
    resourceSpecificPermissionGrantId := "resourceSpecificPermissionGrantId_example" // string | key: id of resourceSpecificPermissionGrant
    microsoftGraphResourceSpecificPermissionGrant := *openapiclient.NewMicrosoftGraphResourceSpecificPermissionGrant() // MicrosoftGraphResourceSpecificPermissionGrant | New property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PermissionGrantsResourceSpecificPermissionGrantApi.PermissionGrantsResourceSpecificPermissionGrantUpdateResourceSpecificPermissionGrant(context.Background(), resourceSpecificPermissionGrantId).MicrosoftGraphResourceSpecificPermissionGrant(microsoftGraphResourceSpecificPermissionGrant).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PermissionGrantsResourceSpecificPermissionGrantApi.PermissionGrantsResourceSpecificPermissionGrantUpdateResourceSpecificPermissionGrant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceSpecificPermissionGrantId** | **string** | key: id of resourceSpecificPermissionGrant | 

### Other Parameters

Other parameters are passed through a pointer to a apiPermissionGrantsResourceSpecificPermissionGrantUpdateResourceSpecificPermissionGrantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphResourceSpecificPermissionGrant** | [**MicrosoftGraphResourceSpecificPermissionGrant**](MicrosoftGraphResourceSpecificPermissionGrant.md) | New property values | 

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

