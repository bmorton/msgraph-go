# \IdentityIdentityApiConnectorApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IdentityCreateApiConnectors**](IdentityIdentityApiConnectorApi.md#IdentityCreateApiConnectors) | **Post** /identity/apiConnectors | Create new navigation property to apiConnectors for identity
[**IdentityDeleteApiConnectors**](IdentityIdentityApiConnectorApi.md#IdentityDeleteApiConnectors) | **Delete** /identity/apiConnectors/{identityApiConnector-id} | Delete navigation property apiConnectors for identity
[**IdentityGetApiConnectors**](IdentityIdentityApiConnectorApi.md#IdentityGetApiConnectors) | **Get** /identity/apiConnectors/{identityApiConnector-id} | Get apiConnectors from identity
[**IdentityListApiConnectors**](IdentityIdentityApiConnectorApi.md#IdentityListApiConnectors) | **Get** /identity/apiConnectors | Get apiConnectors from identity
[**IdentityUpdateApiConnectors**](IdentityIdentityApiConnectorApi.md#IdentityUpdateApiConnectors) | **Patch** /identity/apiConnectors/{identityApiConnector-id} | Update the navigation property apiConnectors in identity



## IdentityCreateApiConnectors

> MicrosoftGraphIdentityApiConnector IdentityCreateApiConnectors(ctx).MicrosoftGraphIdentityApiConnector(microsoftGraphIdentityApiConnector).Execute()

Create new navigation property to apiConnectors for identity



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
    microsoftGraphIdentityApiConnector := *openapiclient.NewMicrosoftGraphIdentityApiConnector() // MicrosoftGraphIdentityApiConnector | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IdentityIdentityApiConnectorApi.IdentityCreateApiConnectors(context.Background()).MicrosoftGraphIdentityApiConnector(microsoftGraphIdentityApiConnector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityIdentityApiConnectorApi.IdentityCreateApiConnectors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IdentityCreateApiConnectors`: MicrosoftGraphIdentityApiConnector
    fmt.Fprintf(os.Stdout, "Response from `IdentityIdentityApiConnectorApi.IdentityCreateApiConnectors`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIdentityCreateApiConnectorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphIdentityApiConnector** | [**MicrosoftGraphIdentityApiConnector**](MicrosoftGraphIdentityApiConnector.md) | New navigation property | 

### Return type

[**MicrosoftGraphIdentityApiConnector**](MicrosoftGraphIdentityApiConnector.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IdentityDeleteApiConnectors

> IdentityDeleteApiConnectors(ctx, identityApiConnectorId).IfMatch(ifMatch).Execute()

Delete navigation property apiConnectors for identity



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
    identityApiConnectorId := "identityApiConnectorId_example" // string | key: id of identityApiConnector
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IdentityIdentityApiConnectorApi.IdentityDeleteApiConnectors(context.Background(), identityApiConnectorId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityIdentityApiConnectorApi.IdentityDeleteApiConnectors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identityApiConnectorId** | **string** | key: id of identityApiConnector | 

### Other Parameters

Other parameters are passed through a pointer to a apiIdentityDeleteApiConnectorsRequest struct via the builder pattern


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


## IdentityGetApiConnectors

> MicrosoftGraphIdentityApiConnector IdentityGetApiConnectors(ctx, identityApiConnectorId).Select_(select_).Expand(expand).Execute()

Get apiConnectors from identity



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
    identityApiConnectorId := "identityApiConnectorId_example" // string | key: id of identityApiConnector
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IdentityIdentityApiConnectorApi.IdentityGetApiConnectors(context.Background(), identityApiConnectorId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityIdentityApiConnectorApi.IdentityGetApiConnectors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IdentityGetApiConnectors`: MicrosoftGraphIdentityApiConnector
    fmt.Fprintf(os.Stdout, "Response from `IdentityIdentityApiConnectorApi.IdentityGetApiConnectors`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identityApiConnectorId** | **string** | key: id of identityApiConnector | 

### Other Parameters

Other parameters are passed through a pointer to a apiIdentityGetApiConnectorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphIdentityApiConnector**](MicrosoftGraphIdentityApiConnector.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IdentityListApiConnectors

> CollectionOfIdentityApiConnector IdentityListApiConnectors(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get apiConnectors from identity



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
    resp, r, err := api_client.IdentityIdentityApiConnectorApi.IdentityListApiConnectors(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityIdentityApiConnectorApi.IdentityListApiConnectors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IdentityListApiConnectors`: CollectionOfIdentityApiConnector
    fmt.Fprintf(os.Stdout, "Response from `IdentityIdentityApiConnectorApi.IdentityListApiConnectors`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIdentityListApiConnectorsRequest struct via the builder pattern


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

[**CollectionOfIdentityApiConnector**](CollectionOfIdentityApiConnector.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IdentityUpdateApiConnectors

> IdentityUpdateApiConnectors(ctx, identityApiConnectorId).MicrosoftGraphIdentityApiConnector(microsoftGraphIdentityApiConnector).Execute()

Update the navigation property apiConnectors in identity



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
    identityApiConnectorId := "identityApiConnectorId_example" // string | key: id of identityApiConnector
    microsoftGraphIdentityApiConnector := *openapiclient.NewMicrosoftGraphIdentityApiConnector() // MicrosoftGraphIdentityApiConnector | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IdentityIdentityApiConnectorApi.IdentityUpdateApiConnectors(context.Background(), identityApiConnectorId).MicrosoftGraphIdentityApiConnector(microsoftGraphIdentityApiConnector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityIdentityApiConnectorApi.IdentityUpdateApiConnectors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identityApiConnectorId** | **string** | key: id of identityApiConnector | 

### Other Parameters

Other parameters are passed through a pointer to a apiIdentityUpdateApiConnectorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphIdentityApiConnector** | [**MicrosoftGraphIdentityApiConnector**](MicrosoftGraphIdentityApiConnector.md) | New navigation property values | 

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

