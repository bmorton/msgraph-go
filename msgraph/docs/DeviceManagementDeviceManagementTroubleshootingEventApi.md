# \DeviceManagementDeviceManagementTroubleshootingEventApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceManagementCreateTroubleshootingEvents**](DeviceManagementDeviceManagementTroubleshootingEventApi.md#DeviceManagementCreateTroubleshootingEvents) | **Post** /deviceManagement/troubleshootingEvents | Create new navigation property to troubleshootingEvents for deviceManagement
[**DeviceManagementDeleteTroubleshootingEvents**](DeviceManagementDeviceManagementTroubleshootingEventApi.md#DeviceManagementDeleteTroubleshootingEvents) | **Delete** /deviceManagement/troubleshootingEvents/{deviceManagementTroubleshootingEvent-id} | Delete navigation property troubleshootingEvents for deviceManagement
[**DeviceManagementGetTroubleshootingEvents**](DeviceManagementDeviceManagementTroubleshootingEventApi.md#DeviceManagementGetTroubleshootingEvents) | **Get** /deviceManagement/troubleshootingEvents/{deviceManagementTroubleshootingEvent-id} | Get troubleshootingEvents from deviceManagement
[**DeviceManagementListTroubleshootingEvents**](DeviceManagementDeviceManagementTroubleshootingEventApi.md#DeviceManagementListTroubleshootingEvents) | **Get** /deviceManagement/troubleshootingEvents | Get troubleshootingEvents from deviceManagement
[**DeviceManagementUpdateTroubleshootingEvents**](DeviceManagementDeviceManagementTroubleshootingEventApi.md#DeviceManagementUpdateTroubleshootingEvents) | **Patch** /deviceManagement/troubleshootingEvents/{deviceManagementTroubleshootingEvent-id} | Update the navigation property troubleshootingEvents in deviceManagement



## DeviceManagementCreateTroubleshootingEvents

> MicrosoftGraphDeviceManagementTroubleshootingEvent DeviceManagementCreateTroubleshootingEvents(ctx).MicrosoftGraphDeviceManagementTroubleshootingEvent(microsoftGraphDeviceManagementTroubleshootingEvent).Execute()

Create new navigation property to troubleshootingEvents for deviceManagement



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
    resp, r, err := api_client.DeviceManagementDeviceManagementTroubleshootingEventApi.DeviceManagementCreateTroubleshootingEvents(context.Background()).MicrosoftGraphDeviceManagementTroubleshootingEvent(microsoftGraphDeviceManagementTroubleshootingEvent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementDeviceManagementTroubleshootingEventApi.DeviceManagementCreateTroubleshootingEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceManagementCreateTroubleshootingEvents`: MicrosoftGraphDeviceManagementTroubleshootingEvent
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementDeviceManagementTroubleshootingEventApi.DeviceManagementCreateTroubleshootingEvents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementCreateTroubleshootingEventsRequest struct via the builder pattern


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


## DeviceManagementDeleteTroubleshootingEvents

> DeviceManagementDeleteTroubleshootingEvents(ctx, deviceManagementTroubleshootingEventId).IfMatch(ifMatch).Execute()

Delete navigation property troubleshootingEvents for deviceManagement



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
    resp, r, err := api_client.DeviceManagementDeviceManagementTroubleshootingEventApi.DeviceManagementDeleteTroubleshootingEvents(context.Background(), deviceManagementTroubleshootingEventId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementDeviceManagementTroubleshootingEventApi.DeviceManagementDeleteTroubleshootingEvents``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeviceManagementDeleteTroubleshootingEventsRequest struct via the builder pattern


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


## DeviceManagementGetTroubleshootingEvents

> MicrosoftGraphDeviceManagementTroubleshootingEvent DeviceManagementGetTroubleshootingEvents(ctx, deviceManagementTroubleshootingEventId).Select_(select_).Expand(expand).Execute()

Get troubleshootingEvents from deviceManagement



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
    resp, r, err := api_client.DeviceManagementDeviceManagementTroubleshootingEventApi.DeviceManagementGetTroubleshootingEvents(context.Background(), deviceManagementTroubleshootingEventId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementDeviceManagementTroubleshootingEventApi.DeviceManagementGetTroubleshootingEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceManagementGetTroubleshootingEvents`: MicrosoftGraphDeviceManagementTroubleshootingEvent
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementDeviceManagementTroubleshootingEventApi.DeviceManagementGetTroubleshootingEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceManagementTroubleshootingEventId** | **string** | key: id of deviceManagementTroubleshootingEvent | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementGetTroubleshootingEventsRequest struct via the builder pattern


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


## DeviceManagementListTroubleshootingEvents

> CollectionOfDeviceManagementTroubleshootingEvent DeviceManagementListTroubleshootingEvents(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get troubleshootingEvents from deviceManagement



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
    resp, r, err := api_client.DeviceManagementDeviceManagementTroubleshootingEventApi.DeviceManagementListTroubleshootingEvents(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementDeviceManagementTroubleshootingEventApi.DeviceManagementListTroubleshootingEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceManagementListTroubleshootingEvents`: CollectionOfDeviceManagementTroubleshootingEvent
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementDeviceManagementTroubleshootingEventApi.DeviceManagementListTroubleshootingEvents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementListTroubleshootingEventsRequest struct via the builder pattern


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


## DeviceManagementUpdateTroubleshootingEvents

> DeviceManagementUpdateTroubleshootingEvents(ctx, deviceManagementTroubleshootingEventId).MicrosoftGraphDeviceManagementTroubleshootingEvent(microsoftGraphDeviceManagementTroubleshootingEvent).Execute()

Update the navigation property troubleshootingEvents in deviceManagement



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
    resp, r, err := api_client.DeviceManagementDeviceManagementTroubleshootingEventApi.DeviceManagementUpdateTroubleshootingEvents(context.Background(), deviceManagementTroubleshootingEventId).MicrosoftGraphDeviceManagementTroubleshootingEvent(microsoftGraphDeviceManagementTroubleshootingEvent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementDeviceManagementTroubleshootingEventApi.DeviceManagementUpdateTroubleshootingEvents``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeviceManagementUpdateTroubleshootingEventsRequest struct via the builder pattern


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

