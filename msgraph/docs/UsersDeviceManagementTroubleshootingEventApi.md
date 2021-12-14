# \UsersDeviceManagementTroubleshootingEventApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersCreateDeviceManagementTroubleshootingEvents**](UsersDeviceManagementTroubleshootingEventApi.md#UsersCreateDeviceManagementTroubleshootingEvents) | **Post** /users/{user-id}/deviceManagementTroubleshootingEvents | Create new navigation property to deviceManagementTroubleshootingEvents for users
[**UsersDeleteDeviceManagementTroubleshootingEvents**](UsersDeviceManagementTroubleshootingEventApi.md#UsersDeleteDeviceManagementTroubleshootingEvents) | **Delete** /users/{user-id}/deviceManagementTroubleshootingEvents/{deviceManagementTroubleshootingEvent-id} | Delete navigation property deviceManagementTroubleshootingEvents for users
[**UsersGetDeviceManagementTroubleshootingEvents**](UsersDeviceManagementTroubleshootingEventApi.md#UsersGetDeviceManagementTroubleshootingEvents) | **Get** /users/{user-id}/deviceManagementTroubleshootingEvents/{deviceManagementTroubleshootingEvent-id} | Get deviceManagementTroubleshootingEvents from users
[**UsersListDeviceManagementTroubleshootingEvents**](UsersDeviceManagementTroubleshootingEventApi.md#UsersListDeviceManagementTroubleshootingEvents) | **Get** /users/{user-id}/deviceManagementTroubleshootingEvents | Get deviceManagementTroubleshootingEvents from users
[**UsersUpdateDeviceManagementTroubleshootingEvents**](UsersDeviceManagementTroubleshootingEventApi.md#UsersUpdateDeviceManagementTroubleshootingEvents) | **Patch** /users/{user-id}/deviceManagementTroubleshootingEvents/{deviceManagementTroubleshootingEvent-id} | Update the navigation property deviceManagementTroubleshootingEvents in users



## UsersCreateDeviceManagementTroubleshootingEvents

> MicrosoftGraphDeviceManagementTroubleshootingEvent UsersCreateDeviceManagementTroubleshootingEvents(ctx, userId).MicrosoftGraphDeviceManagementTroubleshootingEvent(microsoftGraphDeviceManagementTroubleshootingEvent).Execute()

Create new navigation property to deviceManagementTroubleshootingEvents for users



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
    userId := "userId_example" // string | key: id of user
    microsoftGraphDeviceManagementTroubleshootingEvent := *openapiclient.NewMicrosoftGraphDeviceManagementTroubleshootingEvent() // MicrosoftGraphDeviceManagementTroubleshootingEvent | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersDeviceManagementTroubleshootingEventApi.UsersCreateDeviceManagementTroubleshootingEvents(context.Background(), userId).MicrosoftGraphDeviceManagementTroubleshootingEvent(microsoftGraphDeviceManagementTroubleshootingEvent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersDeviceManagementTroubleshootingEventApi.UsersCreateDeviceManagementTroubleshootingEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersCreateDeviceManagementTroubleshootingEvents`: MicrosoftGraphDeviceManagementTroubleshootingEvent
    fmt.Fprintf(os.Stdout, "Response from `UsersDeviceManagementTroubleshootingEventApi.UsersCreateDeviceManagementTroubleshootingEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersCreateDeviceManagementTroubleshootingEventsRequest struct via the builder pattern


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


## UsersDeleteDeviceManagementTroubleshootingEvents

> UsersDeleteDeviceManagementTroubleshootingEvents(ctx, userId, deviceManagementTroubleshootingEventId).IfMatch(ifMatch).Execute()

Delete navigation property deviceManagementTroubleshootingEvents for users



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
    userId := "userId_example" // string | key: id of user
    deviceManagementTroubleshootingEventId := "deviceManagementTroubleshootingEventId_example" // string | key: id of deviceManagementTroubleshootingEvent
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersDeviceManagementTroubleshootingEventApi.UsersDeleteDeviceManagementTroubleshootingEvents(context.Background(), userId, deviceManagementTroubleshootingEventId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersDeviceManagementTroubleshootingEventApi.UsersDeleteDeviceManagementTroubleshootingEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**deviceManagementTroubleshootingEventId** | **string** | key: id of deviceManagementTroubleshootingEvent | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersDeleteDeviceManagementTroubleshootingEventsRequest struct via the builder pattern


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


## UsersGetDeviceManagementTroubleshootingEvents

> MicrosoftGraphDeviceManagementTroubleshootingEvent UsersGetDeviceManagementTroubleshootingEvents(ctx, userId, deviceManagementTroubleshootingEventId).Select_(select_).Expand(expand).Execute()

Get deviceManagementTroubleshootingEvents from users



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
    userId := "userId_example" // string | key: id of user
    deviceManagementTroubleshootingEventId := "deviceManagementTroubleshootingEventId_example" // string | key: id of deviceManagementTroubleshootingEvent
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersDeviceManagementTroubleshootingEventApi.UsersGetDeviceManagementTroubleshootingEvents(context.Background(), userId, deviceManagementTroubleshootingEventId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersDeviceManagementTroubleshootingEventApi.UsersGetDeviceManagementTroubleshootingEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersGetDeviceManagementTroubleshootingEvents`: MicrosoftGraphDeviceManagementTroubleshootingEvent
    fmt.Fprintf(os.Stdout, "Response from `UsersDeviceManagementTroubleshootingEventApi.UsersGetDeviceManagementTroubleshootingEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**deviceManagementTroubleshootingEventId** | **string** | key: id of deviceManagementTroubleshootingEvent | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersGetDeviceManagementTroubleshootingEventsRequest struct via the builder pattern


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


## UsersListDeviceManagementTroubleshootingEvents

> CollectionOfDeviceManagementTroubleshootingEvent UsersListDeviceManagementTroubleshootingEvents(ctx, userId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get deviceManagementTroubleshootingEvents from users



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
    userId := "userId_example" // string | key: id of user
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
    resp, r, err := api_client.UsersDeviceManagementTroubleshootingEventApi.UsersListDeviceManagementTroubleshootingEvents(context.Background(), userId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersDeviceManagementTroubleshootingEventApi.UsersListDeviceManagementTroubleshootingEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersListDeviceManagementTroubleshootingEvents`: CollectionOfDeviceManagementTroubleshootingEvent
    fmt.Fprintf(os.Stdout, "Response from `UsersDeviceManagementTroubleshootingEventApi.UsersListDeviceManagementTroubleshootingEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersListDeviceManagementTroubleshootingEventsRequest struct via the builder pattern


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


## UsersUpdateDeviceManagementTroubleshootingEvents

> UsersUpdateDeviceManagementTroubleshootingEvents(ctx, userId, deviceManagementTroubleshootingEventId).MicrosoftGraphDeviceManagementTroubleshootingEvent(microsoftGraphDeviceManagementTroubleshootingEvent).Execute()

Update the navigation property deviceManagementTroubleshootingEvents in users



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
    userId := "userId_example" // string | key: id of user
    deviceManagementTroubleshootingEventId := "deviceManagementTroubleshootingEventId_example" // string | key: id of deviceManagementTroubleshootingEvent
    microsoftGraphDeviceManagementTroubleshootingEvent := *openapiclient.NewMicrosoftGraphDeviceManagementTroubleshootingEvent() // MicrosoftGraphDeviceManagementTroubleshootingEvent | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersDeviceManagementTroubleshootingEventApi.UsersUpdateDeviceManagementTroubleshootingEvents(context.Background(), userId, deviceManagementTroubleshootingEventId).MicrosoftGraphDeviceManagementTroubleshootingEvent(microsoftGraphDeviceManagementTroubleshootingEvent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersDeviceManagementTroubleshootingEventApi.UsersUpdateDeviceManagementTroubleshootingEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**deviceManagementTroubleshootingEventId** | **string** | key: id of deviceManagementTroubleshootingEvent | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersUpdateDeviceManagementTroubleshootingEventsRequest struct via the builder pattern


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

