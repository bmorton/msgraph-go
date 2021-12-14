# \UsersOnlineMeetingApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersCreateOnlineMeetings**](UsersOnlineMeetingApi.md#UsersCreateOnlineMeetings) | **Post** /users/{user-id}/onlineMeetings | Create new navigation property to onlineMeetings for users
[**UsersDeleteOnlineMeetings**](UsersOnlineMeetingApi.md#UsersDeleteOnlineMeetings) | **Delete** /users/{user-id}/onlineMeetings/{onlineMeeting-id} | Delete navigation property onlineMeetings for users
[**UsersGetOnlineMeetings**](UsersOnlineMeetingApi.md#UsersGetOnlineMeetings) | **Get** /users/{user-id}/onlineMeetings/{onlineMeeting-id} | Get onlineMeetings from users
[**UsersGetOnlineMeetingsAttendeeReport**](UsersOnlineMeetingApi.md#UsersGetOnlineMeetingsAttendeeReport) | **Get** /users/{user-id}/onlineMeetings/{onlineMeeting-id}/attendeeReport | Get media content for the navigation property onlineMeetings from users
[**UsersListOnlineMeetings**](UsersOnlineMeetingApi.md#UsersListOnlineMeetings) | **Get** /users/{user-id}/onlineMeetings | Get onlineMeetings from users
[**UsersUpdateOnlineMeetings**](UsersOnlineMeetingApi.md#UsersUpdateOnlineMeetings) | **Patch** /users/{user-id}/onlineMeetings/{onlineMeeting-id} | Update the navigation property onlineMeetings in users
[**UsersUpdateOnlineMeetingsAttendeeReport**](UsersOnlineMeetingApi.md#UsersUpdateOnlineMeetingsAttendeeReport) | **Put** /users/{user-id}/onlineMeetings/{onlineMeeting-id}/attendeeReport | Update media content for the navigation property onlineMeetings in users



## UsersCreateOnlineMeetings

> MicrosoftGraphOnlineMeeting UsersCreateOnlineMeetings(ctx, userId).MicrosoftGraphOnlineMeeting(microsoftGraphOnlineMeeting).Execute()

Create new navigation property to onlineMeetings for users

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
    microsoftGraphOnlineMeeting := *openapiclient.NewMicrosoftGraphOnlineMeeting() // MicrosoftGraphOnlineMeeting | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersOnlineMeetingApi.UsersCreateOnlineMeetings(context.Background(), userId).MicrosoftGraphOnlineMeeting(microsoftGraphOnlineMeeting).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersOnlineMeetingApi.UsersCreateOnlineMeetings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersCreateOnlineMeetings`: MicrosoftGraphOnlineMeeting
    fmt.Fprintf(os.Stdout, "Response from `UsersOnlineMeetingApi.UsersCreateOnlineMeetings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersCreateOnlineMeetingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphOnlineMeeting** | [**MicrosoftGraphOnlineMeeting**](MicrosoftGraphOnlineMeeting.md) | New navigation property | 

### Return type

[**MicrosoftGraphOnlineMeeting**](MicrosoftGraphOnlineMeeting.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersDeleteOnlineMeetings

> UsersDeleteOnlineMeetings(ctx, userId, onlineMeetingId).IfMatch(ifMatch).Execute()

Delete navigation property onlineMeetings for users

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
    onlineMeetingId := "onlineMeetingId_example" // string | key: id of onlineMeeting
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersOnlineMeetingApi.UsersDeleteOnlineMeetings(context.Background(), userId, onlineMeetingId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersOnlineMeetingApi.UsersDeleteOnlineMeetings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**onlineMeetingId** | **string** | key: id of onlineMeeting | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersDeleteOnlineMeetingsRequest struct via the builder pattern


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


## UsersGetOnlineMeetings

> MicrosoftGraphOnlineMeeting UsersGetOnlineMeetings(ctx, userId, onlineMeetingId).Select_(select_).Expand(expand).Execute()

Get onlineMeetings from users

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
    onlineMeetingId := "onlineMeetingId_example" // string | key: id of onlineMeeting
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersOnlineMeetingApi.UsersGetOnlineMeetings(context.Background(), userId, onlineMeetingId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersOnlineMeetingApi.UsersGetOnlineMeetings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersGetOnlineMeetings`: MicrosoftGraphOnlineMeeting
    fmt.Fprintf(os.Stdout, "Response from `UsersOnlineMeetingApi.UsersGetOnlineMeetings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**onlineMeetingId** | **string** | key: id of onlineMeeting | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersGetOnlineMeetingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphOnlineMeeting**](MicrosoftGraphOnlineMeeting.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersGetOnlineMeetingsAttendeeReport

> *os.File UsersGetOnlineMeetingsAttendeeReport(ctx, userId, onlineMeetingId).Execute()

Get media content for the navigation property onlineMeetings from users

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
    onlineMeetingId := "onlineMeetingId_example" // string | key: id of onlineMeeting

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersOnlineMeetingApi.UsersGetOnlineMeetingsAttendeeReport(context.Background(), userId, onlineMeetingId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersOnlineMeetingApi.UsersGetOnlineMeetingsAttendeeReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersGetOnlineMeetingsAttendeeReport`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `UsersOnlineMeetingApi.UsersGetOnlineMeetingsAttendeeReport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**onlineMeetingId** | **string** | key: id of onlineMeeting | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersGetOnlineMeetingsAttendeeReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[***os.File**](*os.File.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersListOnlineMeetings

> CollectionOfOnlineMeeting UsersListOnlineMeetings(ctx, userId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get onlineMeetings from users

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
    resp, r, err := api_client.UsersOnlineMeetingApi.UsersListOnlineMeetings(context.Background(), userId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersOnlineMeetingApi.UsersListOnlineMeetings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersListOnlineMeetings`: CollectionOfOnlineMeeting
    fmt.Fprintf(os.Stdout, "Response from `UsersOnlineMeetingApi.UsersListOnlineMeetings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersListOnlineMeetingsRequest struct via the builder pattern


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

[**CollectionOfOnlineMeeting**](CollectionOfOnlineMeeting.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersUpdateOnlineMeetings

> UsersUpdateOnlineMeetings(ctx, userId, onlineMeetingId).MicrosoftGraphOnlineMeeting(microsoftGraphOnlineMeeting).Execute()

Update the navigation property onlineMeetings in users

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
    onlineMeetingId := "onlineMeetingId_example" // string | key: id of onlineMeeting
    microsoftGraphOnlineMeeting := *openapiclient.NewMicrosoftGraphOnlineMeeting() // MicrosoftGraphOnlineMeeting | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersOnlineMeetingApi.UsersUpdateOnlineMeetings(context.Background(), userId, onlineMeetingId).MicrosoftGraphOnlineMeeting(microsoftGraphOnlineMeeting).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersOnlineMeetingApi.UsersUpdateOnlineMeetings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**onlineMeetingId** | **string** | key: id of onlineMeeting | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersUpdateOnlineMeetingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **microsoftGraphOnlineMeeting** | [**MicrosoftGraphOnlineMeeting**](MicrosoftGraphOnlineMeeting.md) | New navigation property values | 

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


## UsersUpdateOnlineMeetingsAttendeeReport

> UsersUpdateOnlineMeetingsAttendeeReport(ctx, userId, onlineMeetingId).Body(body).Execute()

Update media content for the navigation property onlineMeetings in users

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
    onlineMeetingId := "onlineMeetingId_example" // string | key: id of onlineMeeting
    body := os.NewFile(1234, "some_file") // *os.File | New media content.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersOnlineMeetingApi.UsersUpdateOnlineMeetingsAttendeeReport(context.Background(), userId, onlineMeetingId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersOnlineMeetingApi.UsersUpdateOnlineMeetingsAttendeeReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**onlineMeetingId** | **string** | key: id of onlineMeeting | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersUpdateOnlineMeetingsAttendeeReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | ***os.File** | New media content. | 

### Return type

 (empty response body)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/octet-stream
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

