# \CommunicationsPresenceApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CommunicationsCreatePresences**](CommunicationsPresenceApi.md#CommunicationsCreatePresences) | **Post** /communications/presences | Create new navigation property to presences for communications
[**CommunicationsDeletePresences**](CommunicationsPresenceApi.md#CommunicationsDeletePresences) | **Delete** /communications/presences/{presence-id} | Delete navigation property presences for communications
[**CommunicationsGetPresences**](CommunicationsPresenceApi.md#CommunicationsGetPresences) | **Get** /communications/presences/{presence-id} | Get presences from communications
[**CommunicationsListPresences**](CommunicationsPresenceApi.md#CommunicationsListPresences) | **Get** /communications/presences | Get presences from communications
[**CommunicationsUpdatePresences**](CommunicationsPresenceApi.md#CommunicationsUpdatePresences) | **Patch** /communications/presences/{presence-id} | Update the navigation property presences in communications



## CommunicationsCreatePresences

> MicrosoftGraphPresence CommunicationsCreatePresences(ctx).MicrosoftGraphPresence(microsoftGraphPresence).Execute()

Create new navigation property to presences for communications

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
    microsoftGraphPresence := *openapiclient.NewMicrosoftGraphPresence() // MicrosoftGraphPresence | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CommunicationsPresenceApi.CommunicationsCreatePresences(context.Background()).MicrosoftGraphPresence(microsoftGraphPresence).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommunicationsPresenceApi.CommunicationsCreatePresences``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CommunicationsCreatePresences`: MicrosoftGraphPresence
    fmt.Fprintf(os.Stdout, "Response from `CommunicationsPresenceApi.CommunicationsCreatePresences`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCommunicationsCreatePresencesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphPresence** | [**MicrosoftGraphPresence**](MicrosoftGraphPresence.md) | New navigation property | 

### Return type

[**MicrosoftGraphPresence**](MicrosoftGraphPresence.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommunicationsDeletePresences

> CommunicationsDeletePresences(ctx, presenceId).IfMatch(ifMatch).Execute()

Delete navigation property presences for communications

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
    presenceId := "presenceId_example" // string | key: id of presence
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CommunicationsPresenceApi.CommunicationsDeletePresences(context.Background(), presenceId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommunicationsPresenceApi.CommunicationsDeletePresences``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**presenceId** | **string** | key: id of presence | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommunicationsDeletePresencesRequest struct via the builder pattern


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


## CommunicationsGetPresences

> MicrosoftGraphPresence CommunicationsGetPresences(ctx, presenceId).Select_(select_).Expand(expand).Execute()

Get presences from communications

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
    presenceId := "presenceId_example" // string | key: id of presence
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CommunicationsPresenceApi.CommunicationsGetPresences(context.Background(), presenceId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommunicationsPresenceApi.CommunicationsGetPresences``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CommunicationsGetPresences`: MicrosoftGraphPresence
    fmt.Fprintf(os.Stdout, "Response from `CommunicationsPresenceApi.CommunicationsGetPresences`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**presenceId** | **string** | key: id of presence | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommunicationsGetPresencesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphPresence**](MicrosoftGraphPresence.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommunicationsListPresences

> CollectionOfPresence CommunicationsListPresences(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get presences from communications

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
    resp, r, err := api_client.CommunicationsPresenceApi.CommunicationsListPresences(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommunicationsPresenceApi.CommunicationsListPresences``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CommunicationsListPresences`: CollectionOfPresence
    fmt.Fprintf(os.Stdout, "Response from `CommunicationsPresenceApi.CommunicationsListPresences`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCommunicationsListPresencesRequest struct via the builder pattern


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

[**CollectionOfPresence**](CollectionOfPresence.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommunicationsUpdatePresences

> CommunicationsUpdatePresences(ctx, presenceId).MicrosoftGraphPresence(microsoftGraphPresence).Execute()

Update the navigation property presences in communications

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
    presenceId := "presenceId_example" // string | key: id of presence
    microsoftGraphPresence := *openapiclient.NewMicrosoftGraphPresence() // MicrosoftGraphPresence | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CommunicationsPresenceApi.CommunicationsUpdatePresences(context.Background(), presenceId).MicrosoftGraphPresence(microsoftGraphPresence).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommunicationsPresenceApi.CommunicationsUpdatePresences``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**presenceId** | **string** | key: id of presence | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommunicationsUpdatePresencesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphPresence** | [**MicrosoftGraphPresence**](MicrosoftGraphPresence.md) | New navigation property values | 

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

