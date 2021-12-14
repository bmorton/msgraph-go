# \CommunicationsOnlineMeetingApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CommunicationsCreateOnlineMeetings**](CommunicationsOnlineMeetingApi.md#CommunicationsCreateOnlineMeetings) | **Post** /communications/onlineMeetings | Create new navigation property to onlineMeetings for communications
[**CommunicationsDeleteOnlineMeetings**](CommunicationsOnlineMeetingApi.md#CommunicationsDeleteOnlineMeetings) | **Delete** /communications/onlineMeetings/{onlineMeeting-id} | Delete navigation property onlineMeetings for communications
[**CommunicationsGetOnlineMeetings**](CommunicationsOnlineMeetingApi.md#CommunicationsGetOnlineMeetings) | **Get** /communications/onlineMeetings/{onlineMeeting-id} | Get onlineMeetings from communications
[**CommunicationsGetOnlineMeetingsAttendeeReport**](CommunicationsOnlineMeetingApi.md#CommunicationsGetOnlineMeetingsAttendeeReport) | **Get** /communications/onlineMeetings/{onlineMeeting-id}/attendeeReport | Get media content for the navigation property onlineMeetings from communications
[**CommunicationsListOnlineMeetings**](CommunicationsOnlineMeetingApi.md#CommunicationsListOnlineMeetings) | **Get** /communications/onlineMeetings | Get onlineMeetings from communications
[**CommunicationsUpdateOnlineMeetings**](CommunicationsOnlineMeetingApi.md#CommunicationsUpdateOnlineMeetings) | **Patch** /communications/onlineMeetings/{onlineMeeting-id} | Update the navigation property onlineMeetings in communications
[**CommunicationsUpdateOnlineMeetingsAttendeeReport**](CommunicationsOnlineMeetingApi.md#CommunicationsUpdateOnlineMeetingsAttendeeReport) | **Put** /communications/onlineMeetings/{onlineMeeting-id}/attendeeReport | Update media content for the navigation property onlineMeetings in communications



## CommunicationsCreateOnlineMeetings

> MicrosoftGraphOnlineMeeting CommunicationsCreateOnlineMeetings(ctx).MicrosoftGraphOnlineMeeting(microsoftGraphOnlineMeeting).Execute()

Create new navigation property to onlineMeetings for communications

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
    resp, r, err := api_client.CommunicationsOnlineMeetingApi.CommunicationsCreateOnlineMeetings(context.Background()).MicrosoftGraphOnlineMeeting(microsoftGraphOnlineMeeting).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommunicationsOnlineMeetingApi.CommunicationsCreateOnlineMeetings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CommunicationsCreateOnlineMeetings`: MicrosoftGraphOnlineMeeting
    fmt.Fprintf(os.Stdout, "Response from `CommunicationsOnlineMeetingApi.CommunicationsCreateOnlineMeetings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCommunicationsCreateOnlineMeetingsRequest struct via the builder pattern


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


## CommunicationsDeleteOnlineMeetings

> CommunicationsDeleteOnlineMeetings(ctx, onlineMeetingId).IfMatch(ifMatch).Execute()

Delete navigation property onlineMeetings for communications

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
    resp, r, err := api_client.CommunicationsOnlineMeetingApi.CommunicationsDeleteOnlineMeetings(context.Background(), onlineMeetingId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommunicationsOnlineMeetingApi.CommunicationsDeleteOnlineMeetings``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCommunicationsDeleteOnlineMeetingsRequest struct via the builder pattern


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


## CommunicationsGetOnlineMeetings

> MicrosoftGraphOnlineMeeting CommunicationsGetOnlineMeetings(ctx, onlineMeetingId).Select_(select_).Expand(expand).Execute()

Get onlineMeetings from communications

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
    resp, r, err := api_client.CommunicationsOnlineMeetingApi.CommunicationsGetOnlineMeetings(context.Background(), onlineMeetingId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommunicationsOnlineMeetingApi.CommunicationsGetOnlineMeetings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CommunicationsGetOnlineMeetings`: MicrosoftGraphOnlineMeeting
    fmt.Fprintf(os.Stdout, "Response from `CommunicationsOnlineMeetingApi.CommunicationsGetOnlineMeetings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**onlineMeetingId** | **string** | key: id of onlineMeeting | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommunicationsGetOnlineMeetingsRequest struct via the builder pattern


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


## CommunicationsGetOnlineMeetingsAttendeeReport

> *os.File CommunicationsGetOnlineMeetingsAttendeeReport(ctx, onlineMeetingId).Execute()

Get media content for the navigation property onlineMeetings from communications

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
    resp, r, err := api_client.CommunicationsOnlineMeetingApi.CommunicationsGetOnlineMeetingsAttendeeReport(context.Background(), onlineMeetingId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommunicationsOnlineMeetingApi.CommunicationsGetOnlineMeetingsAttendeeReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CommunicationsGetOnlineMeetingsAttendeeReport`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `CommunicationsOnlineMeetingApi.CommunicationsGetOnlineMeetingsAttendeeReport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**onlineMeetingId** | **string** | key: id of onlineMeeting | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommunicationsGetOnlineMeetingsAttendeeReportRequest struct via the builder pattern


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


## CommunicationsListOnlineMeetings

> CollectionOfOnlineMeeting CommunicationsListOnlineMeetings(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get onlineMeetings from communications

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
    resp, r, err := api_client.CommunicationsOnlineMeetingApi.CommunicationsListOnlineMeetings(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommunicationsOnlineMeetingApi.CommunicationsListOnlineMeetings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CommunicationsListOnlineMeetings`: CollectionOfOnlineMeeting
    fmt.Fprintf(os.Stdout, "Response from `CommunicationsOnlineMeetingApi.CommunicationsListOnlineMeetings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCommunicationsListOnlineMeetingsRequest struct via the builder pattern


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


## CommunicationsUpdateOnlineMeetings

> CommunicationsUpdateOnlineMeetings(ctx, onlineMeetingId).MicrosoftGraphOnlineMeeting(microsoftGraphOnlineMeeting).Execute()

Update the navigation property onlineMeetings in communications

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
    resp, r, err := api_client.CommunicationsOnlineMeetingApi.CommunicationsUpdateOnlineMeetings(context.Background(), onlineMeetingId).MicrosoftGraphOnlineMeeting(microsoftGraphOnlineMeeting).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommunicationsOnlineMeetingApi.CommunicationsUpdateOnlineMeetings``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCommunicationsUpdateOnlineMeetingsRequest struct via the builder pattern


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


## CommunicationsUpdateOnlineMeetingsAttendeeReport

> CommunicationsUpdateOnlineMeetingsAttendeeReport(ctx, onlineMeetingId).Body(body).Execute()

Update media content for the navigation property onlineMeetings in communications

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
    resp, r, err := api_client.CommunicationsOnlineMeetingApi.CommunicationsUpdateOnlineMeetingsAttendeeReport(context.Background(), onlineMeetingId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommunicationsOnlineMeetingApi.CommunicationsUpdateOnlineMeetingsAttendeeReport``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCommunicationsUpdateOnlineMeetingsAttendeeReportRequest struct via the builder pattern


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

