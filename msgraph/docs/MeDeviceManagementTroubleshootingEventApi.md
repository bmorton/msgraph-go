# \MeDeviceManagementTroubleshootingEventApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MeCreateDeviceManagementTroubleshootingEvents**](MeDeviceManagementTroubleshootingEventApi.md#MeCreateDeviceManagementTroubleshootingEvents) | **Post** /me/deviceManagementTroubleshootingEvents | Create new navigation property to deviceManagementTroubleshootingEvents for me
[**MeDeleteDeviceManagementTroubleshootingEvents**](MeDeviceManagementTroubleshootingEventApi.md#MeDeleteDeviceManagementTroubleshootingEvents) | **Delete** /me/deviceManagementTroubleshootingEvents/{deviceManagementTroubleshootingEvent-id} | Delete navigation property deviceManagementTroubleshootingEvents for me
[**MeGetDeviceManagementTroubleshootingEvents**](MeDeviceManagementTroubleshootingEventApi.md#MeGetDeviceManagementTroubleshootingEvents) | **Get** /me/deviceManagementTroubleshootingEvents/{deviceManagementTroubleshootingEvent-id} | Get deviceManagementTroubleshootingEvents from me
[**MeListDeviceManagementTroubleshootingEvents**](MeDeviceManagementTroubleshootingEventApi.md#MeListDeviceManagementTroubleshootingEvents) | **Get** /me/deviceManagementTroubleshootingEvents | Get deviceManagementTroubleshootingEvents from me
[**MeUpdateDeviceManagementTroubleshootingEvents**](MeDeviceManagementTroubleshootingEventApi.md#MeUpdateDeviceManagementTroubleshootingEvents) | **Patch** /me/deviceManagementTroubleshootingEvents/{deviceManagementTroubleshootingEvent-id} | Update the navigation property deviceManagementTroubleshootingEvents in me



## MeCreateDeviceManagementTroubleshootingEvents

> MicrosoftGraphDeviceManagementTroubleshootingEvent MeCreateDeviceManagementTroubleshootingEvents(ctx).MicrosoftGraphDeviceManagementTroubleshootingEvent(microsoftGraphDeviceManagementTroubleshootingEvent).Execute()

Create new navigation property to deviceManagementTroubleshootingEvents for me



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
    microsoftGraphDeviceManagementTroubleshootingEvent := *openapiclient.NewMicrosoftGraphDeviceManagementTroubleshootingEvent() // MicrosoftGraphDeviceManagementTroubleshootingEvent | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeDeviceManagementTroubleshootingEventApi.MeCreateDeviceManagementTroubleshootingEvents(context.Background()).MicrosoftGraphDeviceManagementTroubleshootingEvent(microsoftGraphDeviceManagementTroubleshootingEvent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeDeviceManagementTroubleshootingEventApi.MeCreateDeviceManagementTroubleshootingEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeCreateDeviceManagementTroubleshootingEvents`: MicrosoftGraphDeviceManagementTroubleshootingEvent
    fmt.Fprintf(os.Stdout, "Response from `MeDeviceManagementTroubleshootingEventApi.MeCreateDeviceManagementTroubleshootingEvents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMeCreateDeviceManagementTroubleshootingEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphDeviceManagementTroubleshootingEvent** | [**MicrosoftGraphDeviceManagementTroubleshootingEvent**](MicrosoftGraphDeviceManagementTroubleshootingEvent.md) | New navigation property | 

### Return type

[**MicrosoftGraphDeviceManagementTroubleshootingEvent**](MicrosoftGraphDeviceManagementTroubleshootingEvent.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeDeleteDeviceManagementTroubleshootingEvents

> MeDeleteDeviceManagementTroubleshootingEvents(ctx, deviceManagementTroubleshootingEventId).IfMatch(ifMatch).Execute()

Delete navigation property deviceManagementTroubleshootingEvents for me



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
    deviceManagementTroubleshootingEventId := "deviceManagementTroubleshootingEventId_example" // string | key: id of deviceManagementTroubleshootingEvent
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeDeviceManagementTroubleshootingEventApi.MeDeleteDeviceManagementTroubleshootingEvents(context.Background(), deviceManagementTroubleshootingEventId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeDeviceManagementTroubleshootingEventApi.MeDeleteDeviceManagementTroubleshootingEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceManagementTroubleshootingEventId** | **string** | key: id of deviceManagementTroubleshootingEvent | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeDeleteDeviceManagementTroubleshootingEventsRequest struct via the builder pattern


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


## MeGetDeviceManagementTroubleshootingEvents

> MicrosoftGraphDeviceManagementTroubleshootingEvent MeGetDeviceManagementTroubleshootingEvents(ctx, deviceManagementTroubleshootingEventId).Select_(select_).Expand(expand).Execute()

Get deviceManagementTroubleshootingEvents from me



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
    deviceManagementTroubleshootingEventId := "deviceManagementTroubleshootingEventId_example" // string | key: id of deviceManagementTroubleshootingEvent
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeDeviceManagementTroubleshootingEventApi.MeGetDeviceManagementTroubleshootingEvents(context.Background(), deviceManagementTroubleshootingEventId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeDeviceManagementTroubleshootingEventApi.MeGetDeviceManagementTroubleshootingEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeGetDeviceManagementTroubleshootingEvents`: MicrosoftGraphDeviceManagementTroubleshootingEvent
    fmt.Fprintf(os.Stdout, "Response from `MeDeviceManagementTroubleshootingEventApi.MeGetDeviceManagementTroubleshootingEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceManagementTroubleshootingEventId** | **string** | key: id of deviceManagementTroubleshootingEvent | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeGetDeviceManagementTroubleshootingEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphDeviceManagementTroubleshootingEvent**](MicrosoftGraphDeviceManagementTroubleshootingEvent.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeListDeviceManagementTroubleshootingEvents

> CollectionOfDeviceManagementTroubleshootingEvent MeListDeviceManagementTroubleshootingEvents(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get deviceManagementTroubleshootingEvents from me



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
    resp, r, err := api_client.MeDeviceManagementTroubleshootingEventApi.MeListDeviceManagementTroubleshootingEvents(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeDeviceManagementTroubleshootingEventApi.MeListDeviceManagementTroubleshootingEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeListDeviceManagementTroubleshootingEvents`: CollectionOfDeviceManagementTroubleshootingEvent
    fmt.Fprintf(os.Stdout, "Response from `MeDeviceManagementTroubleshootingEventApi.MeListDeviceManagementTroubleshootingEvents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMeListDeviceManagementTroubleshootingEventsRequest struct via the builder pattern


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

[**CollectionOfDeviceManagementTroubleshootingEvent**](CollectionOfDeviceManagementTroubleshootingEvent.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeUpdateDeviceManagementTroubleshootingEvents

> MeUpdateDeviceManagementTroubleshootingEvents(ctx, deviceManagementTroubleshootingEventId).MicrosoftGraphDeviceManagementTroubleshootingEvent(microsoftGraphDeviceManagementTroubleshootingEvent).Execute()

Update the navigation property deviceManagementTroubleshootingEvents in me



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
    deviceManagementTroubleshootingEventId := "deviceManagementTroubleshootingEventId_example" // string | key: id of deviceManagementTroubleshootingEvent
    microsoftGraphDeviceManagementTroubleshootingEvent := *openapiclient.NewMicrosoftGraphDeviceManagementTroubleshootingEvent() // MicrosoftGraphDeviceManagementTroubleshootingEvent | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeDeviceManagementTroubleshootingEventApi.MeUpdateDeviceManagementTroubleshootingEvents(context.Background(), deviceManagementTroubleshootingEventId).MicrosoftGraphDeviceManagementTroubleshootingEvent(microsoftGraphDeviceManagementTroubleshootingEvent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeDeviceManagementTroubleshootingEventApi.MeUpdateDeviceManagementTroubleshootingEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceManagementTroubleshootingEventId** | **string** | key: id of deviceManagementTroubleshootingEvent | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeUpdateDeviceManagementTroubleshootingEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphDeviceManagementTroubleshootingEvent** | [**MicrosoftGraphDeviceManagementTroubleshootingEvent**](MicrosoftGraphDeviceManagementTroubleshootingEvent.md) | New navigation property values | 

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

