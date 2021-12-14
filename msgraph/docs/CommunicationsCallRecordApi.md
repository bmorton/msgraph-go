# \CommunicationsCallRecordApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CommunicationsCallRecordsCreateSessions**](CommunicationsCallRecordApi.md#CommunicationsCallRecordsCreateSessions) | **Post** /communications/callRecords/{callRecord-id}/sessions | Create new navigation property to sessions for communications
[**CommunicationsCallRecordsDeleteSessions**](CommunicationsCallRecordApi.md#CommunicationsCallRecordsDeleteSessions) | **Delete** /communications/callRecords/{callRecord-id}/sessions/{session-id} | Delete navigation property sessions for communications
[**CommunicationsCallRecordsGetSessions**](CommunicationsCallRecordApi.md#CommunicationsCallRecordsGetSessions) | **Get** /communications/callRecords/{callRecord-id}/sessions/{session-id} | Get sessions from communications
[**CommunicationsCallRecordsListSessions**](CommunicationsCallRecordApi.md#CommunicationsCallRecordsListSessions) | **Get** /communications/callRecords/{callRecord-id}/sessions | Get sessions from communications
[**CommunicationsCallRecordsSessionsCreateSegments**](CommunicationsCallRecordApi.md#CommunicationsCallRecordsSessionsCreateSegments) | **Post** /communications/callRecords/{callRecord-id}/sessions/{session-id}/segments | Create new navigation property to segments for communications
[**CommunicationsCallRecordsSessionsDeleteSegments**](CommunicationsCallRecordApi.md#CommunicationsCallRecordsSessionsDeleteSegments) | **Delete** /communications/callRecords/{callRecord-id}/sessions/{session-id}/segments/{segment-id} | Delete navigation property segments for communications
[**CommunicationsCallRecordsSessionsGetSegments**](CommunicationsCallRecordApi.md#CommunicationsCallRecordsSessionsGetSegments) | **Get** /communications/callRecords/{callRecord-id}/sessions/{session-id}/segments/{segment-id} | Get segments from communications
[**CommunicationsCallRecordsSessionsListSegments**](CommunicationsCallRecordApi.md#CommunicationsCallRecordsSessionsListSegments) | **Get** /communications/callRecords/{callRecord-id}/sessions/{session-id}/segments | Get segments from communications
[**CommunicationsCallRecordsSessionsUpdateSegments**](CommunicationsCallRecordApi.md#CommunicationsCallRecordsSessionsUpdateSegments) | **Patch** /communications/callRecords/{callRecord-id}/sessions/{session-id}/segments/{segment-id} | Update the navigation property segments in communications
[**CommunicationsCallRecordsUpdateSessions**](CommunicationsCallRecordApi.md#CommunicationsCallRecordsUpdateSessions) | **Patch** /communications/callRecords/{callRecord-id}/sessions/{session-id} | Update the navigation property sessions in communications
[**CommunicationsCreateCallRecords**](CommunicationsCallRecordApi.md#CommunicationsCreateCallRecords) | **Post** /communications/callRecords | Create new navigation property to callRecords for communications
[**CommunicationsDeleteCallRecords**](CommunicationsCallRecordApi.md#CommunicationsDeleteCallRecords) | **Delete** /communications/callRecords/{callRecord-id} | Delete navigation property callRecords for communications
[**CommunicationsGetCallRecords**](CommunicationsCallRecordApi.md#CommunicationsGetCallRecords) | **Get** /communications/callRecords/{callRecord-id} | Get callRecords from communications
[**CommunicationsListCallRecords**](CommunicationsCallRecordApi.md#CommunicationsListCallRecords) | **Get** /communications/callRecords | Get callRecords from communications
[**CommunicationsUpdateCallRecords**](CommunicationsCallRecordApi.md#CommunicationsUpdateCallRecords) | **Patch** /communications/callRecords/{callRecord-id} | Update the navigation property callRecords in communications



## CommunicationsCallRecordsCreateSessions

> MicrosoftGraphCallRecordsSession CommunicationsCallRecordsCreateSessions(ctx, callRecordId).MicrosoftGraphCallRecordsSession(microsoftGraphCallRecordsSession).Execute()

Create new navigation property to sessions for communications



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
    callRecordId := "callRecordId_example" // string | key: id of callRecord
    microsoftGraphCallRecordsSession := *openapiclient.NewMicrosoftGraphCallRecordsSession() // MicrosoftGraphCallRecordsSession | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CommunicationsCallRecordApi.CommunicationsCallRecordsCreateSessions(context.Background(), callRecordId).MicrosoftGraphCallRecordsSession(microsoftGraphCallRecordsSession).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommunicationsCallRecordApi.CommunicationsCallRecordsCreateSessions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CommunicationsCallRecordsCreateSessions`: MicrosoftGraphCallRecordsSession
    fmt.Fprintf(os.Stdout, "Response from `CommunicationsCallRecordApi.CommunicationsCallRecordsCreateSessions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callRecordId** | **string** | key: id of callRecord | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommunicationsCallRecordsCreateSessionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphCallRecordsSession** | [**MicrosoftGraphCallRecordsSession**](MicrosoftGraphCallRecordsSession.md) | New navigation property | 

### Return type

[**MicrosoftGraphCallRecordsSession**](MicrosoftGraphCallRecordsSession.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommunicationsCallRecordsDeleteSessions

> CommunicationsCallRecordsDeleteSessions(ctx, callRecordId, sessionId).IfMatch(ifMatch).Execute()

Delete navigation property sessions for communications



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
    callRecordId := "callRecordId_example" // string | key: id of callRecord
    sessionId := "sessionId_example" // string | key: id of session
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CommunicationsCallRecordApi.CommunicationsCallRecordsDeleteSessions(context.Background(), callRecordId, sessionId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommunicationsCallRecordApi.CommunicationsCallRecordsDeleteSessions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callRecordId** | **string** | key: id of callRecord | 
**sessionId** | **string** | key: id of session | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommunicationsCallRecordsDeleteSessionsRequest struct via the builder pattern


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


## CommunicationsCallRecordsGetSessions

> MicrosoftGraphCallRecordsSession CommunicationsCallRecordsGetSessions(ctx, callRecordId, sessionId).Select_(select_).Expand(expand).Execute()

Get sessions from communications



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
    callRecordId := "callRecordId_example" // string | key: id of callRecord
    sessionId := "sessionId_example" // string | key: id of session
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CommunicationsCallRecordApi.CommunicationsCallRecordsGetSessions(context.Background(), callRecordId, sessionId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommunicationsCallRecordApi.CommunicationsCallRecordsGetSessions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CommunicationsCallRecordsGetSessions`: MicrosoftGraphCallRecordsSession
    fmt.Fprintf(os.Stdout, "Response from `CommunicationsCallRecordApi.CommunicationsCallRecordsGetSessions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callRecordId** | **string** | key: id of callRecord | 
**sessionId** | **string** | key: id of session | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommunicationsCallRecordsGetSessionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphCallRecordsSession**](MicrosoftGraphCallRecordsSession.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommunicationsCallRecordsListSessions

> CollectionOfSession CommunicationsCallRecordsListSessions(ctx, callRecordId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get sessions from communications



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
    callRecordId := "callRecordId_example" // string | key: id of callRecord
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
    resp, r, err := api_client.CommunicationsCallRecordApi.CommunicationsCallRecordsListSessions(context.Background(), callRecordId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommunicationsCallRecordApi.CommunicationsCallRecordsListSessions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CommunicationsCallRecordsListSessions`: CollectionOfSession
    fmt.Fprintf(os.Stdout, "Response from `CommunicationsCallRecordApi.CommunicationsCallRecordsListSessions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callRecordId** | **string** | key: id of callRecord | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommunicationsCallRecordsListSessionsRequest struct via the builder pattern


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

[**CollectionOfSession**](CollectionOfSession.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommunicationsCallRecordsSessionsCreateSegments

> MicrosoftGraphCallRecordsSegment CommunicationsCallRecordsSessionsCreateSegments(ctx, callRecordId, sessionId).MicrosoftGraphCallRecordsSegment(microsoftGraphCallRecordsSegment).Execute()

Create new navigation property to segments for communications



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
    callRecordId := "callRecordId_example" // string | key: id of callRecord
    sessionId := "sessionId_example" // string | key: id of session
    microsoftGraphCallRecordsSegment := *openapiclient.NewMicrosoftGraphCallRecordsSegment() // MicrosoftGraphCallRecordsSegment | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CommunicationsCallRecordApi.CommunicationsCallRecordsSessionsCreateSegments(context.Background(), callRecordId, sessionId).MicrosoftGraphCallRecordsSegment(microsoftGraphCallRecordsSegment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommunicationsCallRecordApi.CommunicationsCallRecordsSessionsCreateSegments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CommunicationsCallRecordsSessionsCreateSegments`: MicrosoftGraphCallRecordsSegment
    fmt.Fprintf(os.Stdout, "Response from `CommunicationsCallRecordApi.CommunicationsCallRecordsSessionsCreateSegments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callRecordId** | **string** | key: id of callRecord | 
**sessionId** | **string** | key: id of session | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommunicationsCallRecordsSessionsCreateSegmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **microsoftGraphCallRecordsSegment** | [**MicrosoftGraphCallRecordsSegment**](MicrosoftGraphCallRecordsSegment.md) | New navigation property | 

### Return type

[**MicrosoftGraphCallRecordsSegment**](MicrosoftGraphCallRecordsSegment.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommunicationsCallRecordsSessionsDeleteSegments

> CommunicationsCallRecordsSessionsDeleteSegments(ctx, callRecordId, sessionId, segmentId).IfMatch(ifMatch).Execute()

Delete navigation property segments for communications



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
    callRecordId := "callRecordId_example" // string | key: id of callRecord
    sessionId := "sessionId_example" // string | key: id of session
    segmentId := "segmentId_example" // string | key: id of segment
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CommunicationsCallRecordApi.CommunicationsCallRecordsSessionsDeleteSegments(context.Background(), callRecordId, sessionId, segmentId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommunicationsCallRecordApi.CommunicationsCallRecordsSessionsDeleteSegments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callRecordId** | **string** | key: id of callRecord | 
**sessionId** | **string** | key: id of session | 
**segmentId** | **string** | key: id of segment | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommunicationsCallRecordsSessionsDeleteSegmentsRequest struct via the builder pattern


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


## CommunicationsCallRecordsSessionsGetSegments

> MicrosoftGraphCallRecordsSegment CommunicationsCallRecordsSessionsGetSegments(ctx, callRecordId, sessionId, segmentId).Select_(select_).Expand(expand).Execute()

Get segments from communications



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
    callRecordId := "callRecordId_example" // string | key: id of callRecord
    sessionId := "sessionId_example" // string | key: id of session
    segmentId := "segmentId_example" // string | key: id of segment
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CommunicationsCallRecordApi.CommunicationsCallRecordsSessionsGetSegments(context.Background(), callRecordId, sessionId, segmentId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommunicationsCallRecordApi.CommunicationsCallRecordsSessionsGetSegments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CommunicationsCallRecordsSessionsGetSegments`: MicrosoftGraphCallRecordsSegment
    fmt.Fprintf(os.Stdout, "Response from `CommunicationsCallRecordApi.CommunicationsCallRecordsSessionsGetSegments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callRecordId** | **string** | key: id of callRecord | 
**sessionId** | **string** | key: id of session | 
**segmentId** | **string** | key: id of segment | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommunicationsCallRecordsSessionsGetSegmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphCallRecordsSegment**](MicrosoftGraphCallRecordsSegment.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommunicationsCallRecordsSessionsListSegments

> CollectionOfSegment CommunicationsCallRecordsSessionsListSegments(ctx, callRecordId, sessionId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get segments from communications



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
    callRecordId := "callRecordId_example" // string | key: id of callRecord
    sessionId := "sessionId_example" // string | key: id of session
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
    resp, r, err := api_client.CommunicationsCallRecordApi.CommunicationsCallRecordsSessionsListSegments(context.Background(), callRecordId, sessionId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommunicationsCallRecordApi.CommunicationsCallRecordsSessionsListSegments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CommunicationsCallRecordsSessionsListSegments`: CollectionOfSegment
    fmt.Fprintf(os.Stdout, "Response from `CommunicationsCallRecordApi.CommunicationsCallRecordsSessionsListSegments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callRecordId** | **string** | key: id of callRecord | 
**sessionId** | **string** | key: id of session | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommunicationsCallRecordsSessionsListSegmentsRequest struct via the builder pattern


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

[**CollectionOfSegment**](CollectionOfSegment.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommunicationsCallRecordsSessionsUpdateSegments

> CommunicationsCallRecordsSessionsUpdateSegments(ctx, callRecordId, sessionId, segmentId).MicrosoftGraphCallRecordsSegment(microsoftGraphCallRecordsSegment).Execute()

Update the navigation property segments in communications



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
    callRecordId := "callRecordId_example" // string | key: id of callRecord
    sessionId := "sessionId_example" // string | key: id of session
    segmentId := "segmentId_example" // string | key: id of segment
    microsoftGraphCallRecordsSegment := *openapiclient.NewMicrosoftGraphCallRecordsSegment() // MicrosoftGraphCallRecordsSegment | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CommunicationsCallRecordApi.CommunicationsCallRecordsSessionsUpdateSegments(context.Background(), callRecordId, sessionId, segmentId).MicrosoftGraphCallRecordsSegment(microsoftGraphCallRecordsSegment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommunicationsCallRecordApi.CommunicationsCallRecordsSessionsUpdateSegments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callRecordId** | **string** | key: id of callRecord | 
**sessionId** | **string** | key: id of session | 
**segmentId** | **string** | key: id of segment | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommunicationsCallRecordsSessionsUpdateSegmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **microsoftGraphCallRecordsSegment** | [**MicrosoftGraphCallRecordsSegment**](MicrosoftGraphCallRecordsSegment.md) | New navigation property values | 

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


## CommunicationsCallRecordsUpdateSessions

> CommunicationsCallRecordsUpdateSessions(ctx, callRecordId, sessionId).MicrosoftGraphCallRecordsSession(microsoftGraphCallRecordsSession).Execute()

Update the navigation property sessions in communications



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
    callRecordId := "callRecordId_example" // string | key: id of callRecord
    sessionId := "sessionId_example" // string | key: id of session
    microsoftGraphCallRecordsSession := *openapiclient.NewMicrosoftGraphCallRecordsSession() // MicrosoftGraphCallRecordsSession | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CommunicationsCallRecordApi.CommunicationsCallRecordsUpdateSessions(context.Background(), callRecordId, sessionId).MicrosoftGraphCallRecordsSession(microsoftGraphCallRecordsSession).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommunicationsCallRecordApi.CommunicationsCallRecordsUpdateSessions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callRecordId** | **string** | key: id of callRecord | 
**sessionId** | **string** | key: id of session | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommunicationsCallRecordsUpdateSessionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **microsoftGraphCallRecordsSession** | [**MicrosoftGraphCallRecordsSession**](MicrosoftGraphCallRecordsSession.md) | New navigation property values | 

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


## CommunicationsCreateCallRecords

> MicrosoftGraphCallRecordsCallRecord CommunicationsCreateCallRecords(ctx).MicrosoftGraphCallRecordsCallRecord(microsoftGraphCallRecordsCallRecord).Execute()

Create new navigation property to callRecords for communications

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
    microsoftGraphCallRecordsCallRecord := *openapiclient.NewMicrosoftGraphCallRecordsCallRecord() // MicrosoftGraphCallRecordsCallRecord | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CommunicationsCallRecordApi.CommunicationsCreateCallRecords(context.Background()).MicrosoftGraphCallRecordsCallRecord(microsoftGraphCallRecordsCallRecord).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommunicationsCallRecordApi.CommunicationsCreateCallRecords``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CommunicationsCreateCallRecords`: MicrosoftGraphCallRecordsCallRecord
    fmt.Fprintf(os.Stdout, "Response from `CommunicationsCallRecordApi.CommunicationsCreateCallRecords`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCommunicationsCreateCallRecordsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphCallRecordsCallRecord** | [**MicrosoftGraphCallRecordsCallRecord**](MicrosoftGraphCallRecordsCallRecord.md) | New navigation property | 

### Return type

[**MicrosoftGraphCallRecordsCallRecord**](MicrosoftGraphCallRecordsCallRecord.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommunicationsDeleteCallRecords

> CommunicationsDeleteCallRecords(ctx, callRecordId).IfMatch(ifMatch).Execute()

Delete navigation property callRecords for communications

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
    callRecordId := "callRecordId_example" // string | key: id of callRecord
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CommunicationsCallRecordApi.CommunicationsDeleteCallRecords(context.Background(), callRecordId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommunicationsCallRecordApi.CommunicationsDeleteCallRecords``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callRecordId** | **string** | key: id of callRecord | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommunicationsDeleteCallRecordsRequest struct via the builder pattern


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


## CommunicationsGetCallRecords

> MicrosoftGraphCallRecordsCallRecord CommunicationsGetCallRecords(ctx, callRecordId).Select_(select_).Expand(expand).Execute()

Get callRecords from communications

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
    callRecordId := "callRecordId_example" // string | key: id of callRecord
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CommunicationsCallRecordApi.CommunicationsGetCallRecords(context.Background(), callRecordId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommunicationsCallRecordApi.CommunicationsGetCallRecords``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CommunicationsGetCallRecords`: MicrosoftGraphCallRecordsCallRecord
    fmt.Fprintf(os.Stdout, "Response from `CommunicationsCallRecordApi.CommunicationsGetCallRecords`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callRecordId** | **string** | key: id of callRecord | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommunicationsGetCallRecordsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphCallRecordsCallRecord**](MicrosoftGraphCallRecordsCallRecord.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommunicationsListCallRecords

> CollectionOfCallRecord CommunicationsListCallRecords(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get callRecords from communications

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
    resp, r, err := api_client.CommunicationsCallRecordApi.CommunicationsListCallRecords(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommunicationsCallRecordApi.CommunicationsListCallRecords``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CommunicationsListCallRecords`: CollectionOfCallRecord
    fmt.Fprintf(os.Stdout, "Response from `CommunicationsCallRecordApi.CommunicationsListCallRecords`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCommunicationsListCallRecordsRequest struct via the builder pattern


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

[**CollectionOfCallRecord**](CollectionOfCallRecord.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommunicationsUpdateCallRecords

> CommunicationsUpdateCallRecords(ctx, callRecordId).MicrosoftGraphCallRecordsCallRecord(microsoftGraphCallRecordsCallRecord).Execute()

Update the navigation property callRecords in communications

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
    callRecordId := "callRecordId_example" // string | key: id of callRecord
    microsoftGraphCallRecordsCallRecord := *openapiclient.NewMicrosoftGraphCallRecordsCallRecord() // MicrosoftGraphCallRecordsCallRecord | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CommunicationsCallRecordApi.CommunicationsUpdateCallRecords(context.Background(), callRecordId).MicrosoftGraphCallRecordsCallRecord(microsoftGraphCallRecordsCallRecord).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommunicationsCallRecordApi.CommunicationsUpdateCallRecords``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callRecordId** | **string** | key: id of callRecord | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommunicationsUpdateCallRecordsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphCallRecordsCallRecord** | [**MicrosoftGraphCallRecordsCallRecord**](MicrosoftGraphCallRecordsCallRecord.md) | New navigation property values | 

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

