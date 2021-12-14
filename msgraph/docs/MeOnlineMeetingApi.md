# \MeOnlineMeetingApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MeCreateOnlineMeetings**](MeOnlineMeetingApi.md#MeCreateOnlineMeetings) | **Post** /me/onlineMeetings | Create new navigation property to onlineMeetings for me
[**MeDeleteOnlineMeetings**](MeOnlineMeetingApi.md#MeDeleteOnlineMeetings) | **Delete** /me/onlineMeetings/{onlineMeeting-id} | Delete navigation property onlineMeetings for me
[**MeGetOnlineMeetings**](MeOnlineMeetingApi.md#MeGetOnlineMeetings) | **Get** /me/onlineMeetings/{onlineMeeting-id} | Get onlineMeetings from me
[**MeGetOnlineMeetingsAttendeeReport**](MeOnlineMeetingApi.md#MeGetOnlineMeetingsAttendeeReport) | **Get** /me/onlineMeetings/{onlineMeeting-id}/attendeeReport | Get media content for the navigation property onlineMeetings from me
[**MeListOnlineMeetings**](MeOnlineMeetingApi.md#MeListOnlineMeetings) | **Get** /me/onlineMeetings | Get onlineMeetings from me
[**MeUpdateOnlineMeetings**](MeOnlineMeetingApi.md#MeUpdateOnlineMeetings) | **Patch** /me/onlineMeetings/{onlineMeeting-id} | Update the navigation property onlineMeetings in me
[**MeUpdateOnlineMeetingsAttendeeReport**](MeOnlineMeetingApi.md#MeUpdateOnlineMeetingsAttendeeReport) | **Put** /me/onlineMeetings/{onlineMeeting-id}/attendeeReport | Update media content for the navigation property onlineMeetings in me



## MeCreateOnlineMeetings

> MicrosoftGraphOnlineMeeting MeCreateOnlineMeetings(ctx).MicrosoftGraphOnlineMeeting(microsoftGraphOnlineMeeting).Execute()

Create new navigation property to onlineMeetings for me

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
    microsoftGraphOnlineMeeting := *openapiclient.NewMicrosoftGraphOnlineMeeting() // MicrosoftGraphOnlineMeeting | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeOnlineMeetingApi.MeCreateOnlineMeetings(context.Background()).MicrosoftGraphOnlineMeeting(microsoftGraphOnlineMeeting).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeOnlineMeetingApi.MeCreateOnlineMeetings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeCreateOnlineMeetings`: MicrosoftGraphOnlineMeeting
    fmt.Fprintf(os.Stdout, "Response from `MeOnlineMeetingApi.MeCreateOnlineMeetings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMeCreateOnlineMeetingsRequest struct via the builder pattern


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


## MeDeleteOnlineMeetings

> MeDeleteOnlineMeetings(ctx, onlineMeetingId).IfMatch(ifMatch).Execute()

Delete navigation property onlineMeetings for me

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
    onlineMeetingId := "onlineMeetingId_example" // string | key: id of onlineMeeting
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeOnlineMeetingApi.MeDeleteOnlineMeetings(context.Background(), onlineMeetingId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeOnlineMeetingApi.MeDeleteOnlineMeetings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**onlineMeetingId** | **string** | key: id of onlineMeeting | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeDeleteOnlineMeetingsRequest struct via the builder pattern


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


## MeGetOnlineMeetings

> MicrosoftGraphOnlineMeeting MeGetOnlineMeetings(ctx, onlineMeetingId).Select_(select_).Expand(expand).Execute()

Get onlineMeetings from me

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
    onlineMeetingId := "onlineMeetingId_example" // string | key: id of onlineMeeting
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeOnlineMeetingApi.MeGetOnlineMeetings(context.Background(), onlineMeetingId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeOnlineMeetingApi.MeGetOnlineMeetings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeGetOnlineMeetings`: MicrosoftGraphOnlineMeeting
    fmt.Fprintf(os.Stdout, "Response from `MeOnlineMeetingApi.MeGetOnlineMeetings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**onlineMeetingId** | **string** | key: id of onlineMeeting | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeGetOnlineMeetingsRequest struct via the builder pattern


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


## MeGetOnlineMeetingsAttendeeReport

> *os.File MeGetOnlineMeetingsAttendeeReport(ctx, onlineMeetingId).Execute()

Get media content for the navigation property onlineMeetings from me

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
    onlineMeetingId := "onlineMeetingId_example" // string | key: id of onlineMeeting

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeOnlineMeetingApi.MeGetOnlineMeetingsAttendeeReport(context.Background(), onlineMeetingId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeOnlineMeetingApi.MeGetOnlineMeetingsAttendeeReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeGetOnlineMeetingsAttendeeReport`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `MeOnlineMeetingApi.MeGetOnlineMeetingsAttendeeReport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**onlineMeetingId** | **string** | key: id of onlineMeeting | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeGetOnlineMeetingsAttendeeReportRequest struct via the builder pattern


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


## MeListOnlineMeetings

> CollectionOfOnlineMeeting MeListOnlineMeetings(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get onlineMeetings from me

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
    resp, r, err := api_client.MeOnlineMeetingApi.MeListOnlineMeetings(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeOnlineMeetingApi.MeListOnlineMeetings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeListOnlineMeetings`: CollectionOfOnlineMeeting
    fmt.Fprintf(os.Stdout, "Response from `MeOnlineMeetingApi.MeListOnlineMeetings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMeListOnlineMeetingsRequest struct via the builder pattern


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


## MeUpdateOnlineMeetings

> MeUpdateOnlineMeetings(ctx, onlineMeetingId).MicrosoftGraphOnlineMeeting(microsoftGraphOnlineMeeting).Execute()

Update the navigation property onlineMeetings in me

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
    onlineMeetingId := "onlineMeetingId_example" // string | key: id of onlineMeeting
    microsoftGraphOnlineMeeting := *openapiclient.NewMicrosoftGraphOnlineMeeting() // MicrosoftGraphOnlineMeeting | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeOnlineMeetingApi.MeUpdateOnlineMeetings(context.Background(), onlineMeetingId).MicrosoftGraphOnlineMeeting(microsoftGraphOnlineMeeting).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeOnlineMeetingApi.MeUpdateOnlineMeetings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**onlineMeetingId** | **string** | key: id of onlineMeeting | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeUpdateOnlineMeetingsRequest struct via the builder pattern


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


## MeUpdateOnlineMeetingsAttendeeReport

> MeUpdateOnlineMeetingsAttendeeReport(ctx, onlineMeetingId).Body(body).Execute()

Update media content for the navigation property onlineMeetings in me

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
    onlineMeetingId := "onlineMeetingId_example" // string | key: id of onlineMeeting
    body := os.NewFile(1234, "some_file") // *os.File | New media content.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeOnlineMeetingApi.MeUpdateOnlineMeetingsAttendeeReport(context.Background(), onlineMeetingId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeOnlineMeetingApi.MeUpdateOnlineMeetingsAttendeeReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**onlineMeetingId** | **string** | key: id of onlineMeeting | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeUpdateOnlineMeetingsAttendeeReportRequest struct via the builder pattern


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

