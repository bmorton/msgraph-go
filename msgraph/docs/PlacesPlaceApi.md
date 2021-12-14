# \PlacesPlaceApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PlacesPlaceCreatePlace**](PlacesPlaceApi.md#PlacesPlaceCreatePlace) | **Post** /places | Add new entity to places
[**PlacesPlaceDeletePlace**](PlacesPlaceApi.md#PlacesPlaceDeletePlace) | **Delete** /places/{place-id} | Delete entity from places
[**PlacesPlaceGetPlace**](PlacesPlaceApi.md#PlacesPlaceGetPlace) | **Get** /places/{place-id} | Get entity from places by key
[**PlacesPlaceListPlace**](PlacesPlaceApi.md#PlacesPlaceListPlace) | **Get** /places | Get entities from places
[**PlacesPlaceUpdatePlace**](PlacesPlaceApi.md#PlacesPlaceUpdatePlace) | **Patch** /places/{place-id} | Update entity in places



## PlacesPlaceCreatePlace

> MicrosoftGraphPlace PlacesPlaceCreatePlace(ctx).MicrosoftGraphPlace(microsoftGraphPlace).Execute()

Add new entity to places

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
    microsoftGraphPlace := *openapiclient.NewMicrosoftGraphPlace() // MicrosoftGraphPlace | New entity

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlacesPlaceApi.PlacesPlaceCreatePlace(context.Background()).MicrosoftGraphPlace(microsoftGraphPlace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlacesPlaceApi.PlacesPlaceCreatePlace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PlacesPlaceCreatePlace`: MicrosoftGraphPlace
    fmt.Fprintf(os.Stdout, "Response from `PlacesPlaceApi.PlacesPlaceCreatePlace`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlacesPlaceCreatePlaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphPlace** | [**MicrosoftGraphPlace**](MicrosoftGraphPlace.md) | New entity | 

### Return type

[**MicrosoftGraphPlace**](MicrosoftGraphPlace.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlacesPlaceDeletePlace

> PlacesPlaceDeletePlace(ctx, placeId).IfMatch(ifMatch).Execute()

Delete entity from places

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
    placeId := "placeId_example" // string | key: id of place
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlacesPlaceApi.PlacesPlaceDeletePlace(context.Background(), placeId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlacesPlaceApi.PlacesPlaceDeletePlace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**placeId** | **string** | key: id of place | 

### Other Parameters

Other parameters are passed through a pointer to a apiPlacesPlaceDeletePlaceRequest struct via the builder pattern


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


## PlacesPlaceGetPlace

> MicrosoftGraphPlace PlacesPlaceGetPlace(ctx, placeId).Select_(select_).Expand(expand).Execute()

Get entity from places by key

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
    placeId := "placeId_example" // string | key: id of place
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlacesPlaceApi.PlacesPlaceGetPlace(context.Background(), placeId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlacesPlaceApi.PlacesPlaceGetPlace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PlacesPlaceGetPlace`: MicrosoftGraphPlace
    fmt.Fprintf(os.Stdout, "Response from `PlacesPlaceApi.PlacesPlaceGetPlace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**placeId** | **string** | key: id of place | 

### Other Parameters

Other parameters are passed through a pointer to a apiPlacesPlaceGetPlaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphPlace**](MicrosoftGraphPlace.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlacesPlaceListPlace

> CollectionOfPlace PlacesPlaceListPlace(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get entities from places

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
    resp, r, err := api_client.PlacesPlaceApi.PlacesPlaceListPlace(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlacesPlaceApi.PlacesPlaceListPlace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PlacesPlaceListPlace`: CollectionOfPlace
    fmt.Fprintf(os.Stdout, "Response from `PlacesPlaceApi.PlacesPlaceListPlace`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlacesPlaceListPlaceRequest struct via the builder pattern


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

[**CollectionOfPlace**](CollectionOfPlace.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlacesPlaceUpdatePlace

> PlacesPlaceUpdatePlace(ctx, placeId).MicrosoftGraphPlace(microsoftGraphPlace).Execute()

Update entity in places

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
    placeId := "placeId_example" // string | key: id of place
    microsoftGraphPlace := *openapiclient.NewMicrosoftGraphPlace() // MicrosoftGraphPlace | New property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlacesPlaceApi.PlacesPlaceUpdatePlace(context.Background(), placeId).MicrosoftGraphPlace(microsoftGraphPlace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlacesPlaceApi.PlacesPlaceUpdatePlace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**placeId** | **string** | key: id of place | 

### Other Parameters

Other parameters are passed through a pointer to a apiPlacesPlaceUpdatePlaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphPlace** | [**MicrosoftGraphPlace**](MicrosoftGraphPlace.md) | New property values | 

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

