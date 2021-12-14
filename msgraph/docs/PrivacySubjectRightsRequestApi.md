# \PrivacySubjectRightsRequestApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PrivacyCreateSubjectRightsRequests**](PrivacySubjectRightsRequestApi.md#PrivacyCreateSubjectRightsRequests) | **Post** /privacy/subjectRightsRequests | Create new navigation property to subjectRightsRequests for privacy
[**PrivacyDeleteSubjectRightsRequests**](PrivacySubjectRightsRequestApi.md#PrivacyDeleteSubjectRightsRequests) | **Delete** /privacy/subjectRightsRequests/{subjectRightsRequest-id} | Delete navigation property subjectRightsRequests for privacy
[**PrivacyGetSubjectRightsRequests**](PrivacySubjectRightsRequestApi.md#PrivacyGetSubjectRightsRequests) | **Get** /privacy/subjectRightsRequests/{subjectRightsRequest-id} | Get subjectRightsRequests from privacy
[**PrivacyListSubjectRightsRequests**](PrivacySubjectRightsRequestApi.md#PrivacyListSubjectRightsRequests) | **Get** /privacy/subjectRightsRequests | Get subjectRightsRequests from privacy
[**PrivacySubjectRightsRequestsCreateNotes**](PrivacySubjectRightsRequestApi.md#PrivacySubjectRightsRequestsCreateNotes) | **Post** /privacy/subjectRightsRequests/{subjectRightsRequest-id}/notes | Create new navigation property to notes for privacy
[**PrivacySubjectRightsRequestsDeleteNotes**](PrivacySubjectRightsRequestApi.md#PrivacySubjectRightsRequestsDeleteNotes) | **Delete** /privacy/subjectRightsRequests/{subjectRightsRequest-id}/notes/{authoredNote-id} | Delete navigation property notes for privacy
[**PrivacySubjectRightsRequestsDeleteRefTeam**](PrivacySubjectRightsRequestApi.md#PrivacySubjectRightsRequestsDeleteRefTeam) | **Delete** /privacy/subjectRightsRequests/{subjectRightsRequest-id}/team/$ref | Delete ref of navigation property team for privacy
[**PrivacySubjectRightsRequestsGetNotes**](PrivacySubjectRightsRequestApi.md#PrivacySubjectRightsRequestsGetNotes) | **Get** /privacy/subjectRightsRequests/{subjectRightsRequest-id}/notes/{authoredNote-id} | Get notes from privacy
[**PrivacySubjectRightsRequestsGetRefTeam**](PrivacySubjectRightsRequestApi.md#PrivacySubjectRightsRequestsGetRefTeam) | **Get** /privacy/subjectRightsRequests/{subjectRightsRequest-id}/team/$ref | Get ref of team from privacy
[**PrivacySubjectRightsRequestsGetTeam**](PrivacySubjectRightsRequestApi.md#PrivacySubjectRightsRequestsGetTeam) | **Get** /privacy/subjectRightsRequests/{subjectRightsRequest-id}/team | Get team from privacy
[**PrivacySubjectRightsRequestsListNotes**](PrivacySubjectRightsRequestApi.md#PrivacySubjectRightsRequestsListNotes) | **Get** /privacy/subjectRightsRequests/{subjectRightsRequest-id}/notes | Get notes from privacy
[**PrivacySubjectRightsRequestsUpdateNotes**](PrivacySubjectRightsRequestApi.md#PrivacySubjectRightsRequestsUpdateNotes) | **Patch** /privacy/subjectRightsRequests/{subjectRightsRequest-id}/notes/{authoredNote-id} | Update the navigation property notes in privacy
[**PrivacySubjectRightsRequestsUpdateRefTeam**](PrivacySubjectRightsRequestApi.md#PrivacySubjectRightsRequestsUpdateRefTeam) | **Put** /privacy/subjectRightsRequests/{subjectRightsRequest-id}/team/$ref | Update the ref of navigation property team in privacy
[**PrivacyUpdateSubjectRightsRequests**](PrivacySubjectRightsRequestApi.md#PrivacyUpdateSubjectRightsRequests) | **Patch** /privacy/subjectRightsRequests/{subjectRightsRequest-id} | Update the navigation property subjectRightsRequests in privacy



## PrivacyCreateSubjectRightsRequests

> MicrosoftGraphSubjectRightsRequest PrivacyCreateSubjectRightsRequests(ctx).MicrosoftGraphSubjectRightsRequest(microsoftGraphSubjectRightsRequest).Execute()

Create new navigation property to subjectRightsRequests for privacy

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
    microsoftGraphSubjectRightsRequest := *openapiclient.NewMicrosoftGraphSubjectRightsRequest() // MicrosoftGraphSubjectRightsRequest | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PrivacySubjectRightsRequestApi.PrivacyCreateSubjectRightsRequests(context.Background()).MicrosoftGraphSubjectRightsRequest(microsoftGraphSubjectRightsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrivacySubjectRightsRequestApi.PrivacyCreateSubjectRightsRequests``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PrivacyCreateSubjectRightsRequests`: MicrosoftGraphSubjectRightsRequest
    fmt.Fprintf(os.Stdout, "Response from `PrivacySubjectRightsRequestApi.PrivacyCreateSubjectRightsRequests`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPrivacyCreateSubjectRightsRequestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphSubjectRightsRequest** | [**MicrosoftGraphSubjectRightsRequest**](MicrosoftGraphSubjectRightsRequest.md) | New navigation property | 

### Return type

[**MicrosoftGraphSubjectRightsRequest**](MicrosoftGraphSubjectRightsRequest.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrivacyDeleteSubjectRightsRequests

> PrivacyDeleteSubjectRightsRequests(ctx, subjectRightsRequestId).IfMatch(ifMatch).Execute()

Delete navigation property subjectRightsRequests for privacy

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
    subjectRightsRequestId := "subjectRightsRequestId_example" // string | key: id of subjectRightsRequest
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PrivacySubjectRightsRequestApi.PrivacyDeleteSubjectRightsRequests(context.Background(), subjectRightsRequestId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrivacySubjectRightsRequestApi.PrivacyDeleteSubjectRightsRequests``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subjectRightsRequestId** | **string** | key: id of subjectRightsRequest | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrivacyDeleteSubjectRightsRequestsRequest struct via the builder pattern


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


## PrivacyGetSubjectRightsRequests

> MicrosoftGraphSubjectRightsRequest PrivacyGetSubjectRightsRequests(ctx, subjectRightsRequestId).Select_(select_).Expand(expand).Execute()

Get subjectRightsRequests from privacy

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
    subjectRightsRequestId := "subjectRightsRequestId_example" // string | key: id of subjectRightsRequest
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PrivacySubjectRightsRequestApi.PrivacyGetSubjectRightsRequests(context.Background(), subjectRightsRequestId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrivacySubjectRightsRequestApi.PrivacyGetSubjectRightsRequests``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PrivacyGetSubjectRightsRequests`: MicrosoftGraphSubjectRightsRequest
    fmt.Fprintf(os.Stdout, "Response from `PrivacySubjectRightsRequestApi.PrivacyGetSubjectRightsRequests`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subjectRightsRequestId** | **string** | key: id of subjectRightsRequest | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrivacyGetSubjectRightsRequestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphSubjectRightsRequest**](MicrosoftGraphSubjectRightsRequest.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrivacyListSubjectRightsRequests

> CollectionOfSubjectRightsRequest PrivacyListSubjectRightsRequests(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get subjectRightsRequests from privacy

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
    resp, r, err := api_client.PrivacySubjectRightsRequestApi.PrivacyListSubjectRightsRequests(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrivacySubjectRightsRequestApi.PrivacyListSubjectRightsRequests``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PrivacyListSubjectRightsRequests`: CollectionOfSubjectRightsRequest
    fmt.Fprintf(os.Stdout, "Response from `PrivacySubjectRightsRequestApi.PrivacyListSubjectRightsRequests`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPrivacyListSubjectRightsRequestsRequest struct via the builder pattern


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

[**CollectionOfSubjectRightsRequest**](CollectionOfSubjectRightsRequest.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrivacySubjectRightsRequestsCreateNotes

> MicrosoftGraphAuthoredNote PrivacySubjectRightsRequestsCreateNotes(ctx, subjectRightsRequestId).MicrosoftGraphAuthoredNote(microsoftGraphAuthoredNote).Execute()

Create new navigation property to notes for privacy



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
    subjectRightsRequestId := "subjectRightsRequestId_example" // string | key: id of subjectRightsRequest
    microsoftGraphAuthoredNote := *openapiclient.NewMicrosoftGraphAuthoredNote() // MicrosoftGraphAuthoredNote | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PrivacySubjectRightsRequestApi.PrivacySubjectRightsRequestsCreateNotes(context.Background(), subjectRightsRequestId).MicrosoftGraphAuthoredNote(microsoftGraphAuthoredNote).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrivacySubjectRightsRequestApi.PrivacySubjectRightsRequestsCreateNotes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PrivacySubjectRightsRequestsCreateNotes`: MicrosoftGraphAuthoredNote
    fmt.Fprintf(os.Stdout, "Response from `PrivacySubjectRightsRequestApi.PrivacySubjectRightsRequestsCreateNotes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subjectRightsRequestId** | **string** | key: id of subjectRightsRequest | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrivacySubjectRightsRequestsCreateNotesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphAuthoredNote** | [**MicrosoftGraphAuthoredNote**](MicrosoftGraphAuthoredNote.md) | New navigation property | 

### Return type

[**MicrosoftGraphAuthoredNote**](MicrosoftGraphAuthoredNote.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrivacySubjectRightsRequestsDeleteNotes

> PrivacySubjectRightsRequestsDeleteNotes(ctx, subjectRightsRequestId, authoredNoteId).IfMatch(ifMatch).Execute()

Delete navigation property notes for privacy



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
    subjectRightsRequestId := "subjectRightsRequestId_example" // string | key: id of subjectRightsRequest
    authoredNoteId := "authoredNoteId_example" // string | key: id of authoredNote
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PrivacySubjectRightsRequestApi.PrivacySubjectRightsRequestsDeleteNotes(context.Background(), subjectRightsRequestId, authoredNoteId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrivacySubjectRightsRequestApi.PrivacySubjectRightsRequestsDeleteNotes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subjectRightsRequestId** | **string** | key: id of subjectRightsRequest | 
**authoredNoteId** | **string** | key: id of authoredNote | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrivacySubjectRightsRequestsDeleteNotesRequest struct via the builder pattern


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


## PrivacySubjectRightsRequestsDeleteRefTeam

> PrivacySubjectRightsRequestsDeleteRefTeam(ctx, subjectRightsRequestId).IfMatch(ifMatch).Execute()

Delete ref of navigation property team for privacy



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
    subjectRightsRequestId := "subjectRightsRequestId_example" // string | key: id of subjectRightsRequest
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PrivacySubjectRightsRequestApi.PrivacySubjectRightsRequestsDeleteRefTeam(context.Background(), subjectRightsRequestId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrivacySubjectRightsRequestApi.PrivacySubjectRightsRequestsDeleteRefTeam``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subjectRightsRequestId** | **string** | key: id of subjectRightsRequest | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrivacySubjectRightsRequestsDeleteRefTeamRequest struct via the builder pattern


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


## PrivacySubjectRightsRequestsGetNotes

> MicrosoftGraphAuthoredNote PrivacySubjectRightsRequestsGetNotes(ctx, subjectRightsRequestId, authoredNoteId).Select_(select_).Expand(expand).Execute()

Get notes from privacy



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
    subjectRightsRequestId := "subjectRightsRequestId_example" // string | key: id of subjectRightsRequest
    authoredNoteId := "authoredNoteId_example" // string | key: id of authoredNote
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PrivacySubjectRightsRequestApi.PrivacySubjectRightsRequestsGetNotes(context.Background(), subjectRightsRequestId, authoredNoteId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrivacySubjectRightsRequestApi.PrivacySubjectRightsRequestsGetNotes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PrivacySubjectRightsRequestsGetNotes`: MicrosoftGraphAuthoredNote
    fmt.Fprintf(os.Stdout, "Response from `PrivacySubjectRightsRequestApi.PrivacySubjectRightsRequestsGetNotes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subjectRightsRequestId** | **string** | key: id of subjectRightsRequest | 
**authoredNoteId** | **string** | key: id of authoredNote | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrivacySubjectRightsRequestsGetNotesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphAuthoredNote**](MicrosoftGraphAuthoredNote.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrivacySubjectRightsRequestsGetRefTeam

> string PrivacySubjectRightsRequestsGetRefTeam(ctx, subjectRightsRequestId).Execute()

Get ref of team from privacy



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
    subjectRightsRequestId := "subjectRightsRequestId_example" // string | key: id of subjectRightsRequest

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PrivacySubjectRightsRequestApi.PrivacySubjectRightsRequestsGetRefTeam(context.Background(), subjectRightsRequestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrivacySubjectRightsRequestApi.PrivacySubjectRightsRequestsGetRefTeam``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PrivacySubjectRightsRequestsGetRefTeam`: string
    fmt.Fprintf(os.Stdout, "Response from `PrivacySubjectRightsRequestApi.PrivacySubjectRightsRequestsGetRefTeam`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subjectRightsRequestId** | **string** | key: id of subjectRightsRequest | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrivacySubjectRightsRequestsGetRefTeamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrivacySubjectRightsRequestsGetTeam

> MicrosoftGraphTeam PrivacySubjectRightsRequestsGetTeam(ctx, subjectRightsRequestId).Select_(select_).Expand(expand).Execute()

Get team from privacy



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
    subjectRightsRequestId := "subjectRightsRequestId_example" // string | key: id of subjectRightsRequest
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PrivacySubjectRightsRequestApi.PrivacySubjectRightsRequestsGetTeam(context.Background(), subjectRightsRequestId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrivacySubjectRightsRequestApi.PrivacySubjectRightsRequestsGetTeam``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PrivacySubjectRightsRequestsGetTeam`: MicrosoftGraphTeam
    fmt.Fprintf(os.Stdout, "Response from `PrivacySubjectRightsRequestApi.PrivacySubjectRightsRequestsGetTeam`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subjectRightsRequestId** | **string** | key: id of subjectRightsRequest | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrivacySubjectRightsRequestsGetTeamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphTeam**](MicrosoftGraphTeam.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrivacySubjectRightsRequestsListNotes

> CollectionOfAuthoredNote PrivacySubjectRightsRequestsListNotes(ctx, subjectRightsRequestId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get notes from privacy



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
    subjectRightsRequestId := "subjectRightsRequestId_example" // string | key: id of subjectRightsRequest
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
    resp, r, err := api_client.PrivacySubjectRightsRequestApi.PrivacySubjectRightsRequestsListNotes(context.Background(), subjectRightsRequestId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrivacySubjectRightsRequestApi.PrivacySubjectRightsRequestsListNotes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PrivacySubjectRightsRequestsListNotes`: CollectionOfAuthoredNote
    fmt.Fprintf(os.Stdout, "Response from `PrivacySubjectRightsRequestApi.PrivacySubjectRightsRequestsListNotes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subjectRightsRequestId** | **string** | key: id of subjectRightsRequest | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrivacySubjectRightsRequestsListNotesRequest struct via the builder pattern


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

[**CollectionOfAuthoredNote**](CollectionOfAuthoredNote.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrivacySubjectRightsRequestsUpdateNotes

> PrivacySubjectRightsRequestsUpdateNotes(ctx, subjectRightsRequestId, authoredNoteId).MicrosoftGraphAuthoredNote(microsoftGraphAuthoredNote).Execute()

Update the navigation property notes in privacy



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
    subjectRightsRequestId := "subjectRightsRequestId_example" // string | key: id of subjectRightsRequest
    authoredNoteId := "authoredNoteId_example" // string | key: id of authoredNote
    microsoftGraphAuthoredNote := *openapiclient.NewMicrosoftGraphAuthoredNote() // MicrosoftGraphAuthoredNote | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PrivacySubjectRightsRequestApi.PrivacySubjectRightsRequestsUpdateNotes(context.Background(), subjectRightsRequestId, authoredNoteId).MicrosoftGraphAuthoredNote(microsoftGraphAuthoredNote).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrivacySubjectRightsRequestApi.PrivacySubjectRightsRequestsUpdateNotes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subjectRightsRequestId** | **string** | key: id of subjectRightsRequest | 
**authoredNoteId** | **string** | key: id of authoredNote | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrivacySubjectRightsRequestsUpdateNotesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **microsoftGraphAuthoredNote** | [**MicrosoftGraphAuthoredNote**](MicrosoftGraphAuthoredNote.md) | New navigation property values | 

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


## PrivacySubjectRightsRequestsUpdateRefTeam

> PrivacySubjectRightsRequestsUpdateRefTeam(ctx, subjectRightsRequestId).RequestBody(requestBody).Execute()

Update the ref of navigation property team in privacy



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
    subjectRightsRequestId := "subjectRightsRequestId_example" // string | key: id of subjectRightsRequest
    requestBody := map[string]map[string]interface{}{"key": map[string]interface{}(123)} // map[string]map[string]interface{} | New navigation property ref values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PrivacySubjectRightsRequestApi.PrivacySubjectRightsRequestsUpdateRefTeam(context.Background(), subjectRightsRequestId).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrivacySubjectRightsRequestApi.PrivacySubjectRightsRequestsUpdateRefTeam``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subjectRightsRequestId** | **string** | key: id of subjectRightsRequest | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrivacySubjectRightsRequestsUpdateRefTeamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBody** | **map[string]map[string]interface{}** | New navigation property ref values | 

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


## PrivacyUpdateSubjectRightsRequests

> PrivacyUpdateSubjectRightsRequests(ctx, subjectRightsRequestId).MicrosoftGraphSubjectRightsRequest(microsoftGraphSubjectRightsRequest).Execute()

Update the navigation property subjectRightsRequests in privacy

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
    subjectRightsRequestId := "subjectRightsRequestId_example" // string | key: id of subjectRightsRequest
    microsoftGraphSubjectRightsRequest := *openapiclient.NewMicrosoftGraphSubjectRightsRequest() // MicrosoftGraphSubjectRightsRequest | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PrivacySubjectRightsRequestApi.PrivacyUpdateSubjectRightsRequests(context.Background(), subjectRightsRequestId).MicrosoftGraphSubjectRightsRequest(microsoftGraphSubjectRightsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrivacySubjectRightsRequestApi.PrivacyUpdateSubjectRightsRequests``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subjectRightsRequestId** | **string** | key: id of subjectRightsRequest | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrivacyUpdateSubjectRightsRequestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphSubjectRightsRequest** | [**MicrosoftGraphSubjectRightsRequest**](MicrosoftGraphSubjectRightsRequest.md) | New navigation property values | 

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

