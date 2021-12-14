# \AdminServiceAnnouncementApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdminDeleteServiceAnnouncement**](AdminServiceAnnouncementApi.md#AdminDeleteServiceAnnouncement) | **Delete** /admin/serviceAnnouncement | Delete navigation property serviceAnnouncement for admin
[**AdminGetServiceAnnouncement**](AdminServiceAnnouncementApi.md#AdminGetServiceAnnouncement) | **Get** /admin/serviceAnnouncement | Get serviceAnnouncement from admin
[**AdminServiceAnnouncementCreateHealthOverviews**](AdminServiceAnnouncementApi.md#AdminServiceAnnouncementCreateHealthOverviews) | **Post** /admin/serviceAnnouncement/healthOverviews | Create new navigation property to healthOverviews for admin
[**AdminServiceAnnouncementCreateIssues**](AdminServiceAnnouncementApi.md#AdminServiceAnnouncementCreateIssues) | **Post** /admin/serviceAnnouncement/issues | Create new navigation property to issues for admin
[**AdminServiceAnnouncementCreateMessages**](AdminServiceAnnouncementApi.md#AdminServiceAnnouncementCreateMessages) | **Post** /admin/serviceAnnouncement/messages | Create new navigation property to messages for admin
[**AdminServiceAnnouncementDeleteHealthOverviews**](AdminServiceAnnouncementApi.md#AdminServiceAnnouncementDeleteHealthOverviews) | **Delete** /admin/serviceAnnouncement/healthOverviews/{serviceHealth-id} | Delete navigation property healthOverviews for admin
[**AdminServiceAnnouncementDeleteIssues**](AdminServiceAnnouncementApi.md#AdminServiceAnnouncementDeleteIssues) | **Delete** /admin/serviceAnnouncement/issues/{serviceHealthIssue-id} | Delete navigation property issues for admin
[**AdminServiceAnnouncementDeleteMessages**](AdminServiceAnnouncementApi.md#AdminServiceAnnouncementDeleteMessages) | **Delete** /admin/serviceAnnouncement/messages/{serviceUpdateMessage-id} | Delete navigation property messages for admin
[**AdminServiceAnnouncementGetHealthOverviews**](AdminServiceAnnouncementApi.md#AdminServiceAnnouncementGetHealthOverviews) | **Get** /admin/serviceAnnouncement/healthOverviews/{serviceHealth-id} | Get healthOverviews from admin
[**AdminServiceAnnouncementGetIssues**](AdminServiceAnnouncementApi.md#AdminServiceAnnouncementGetIssues) | **Get** /admin/serviceAnnouncement/issues/{serviceHealthIssue-id} | Get issues from admin
[**AdminServiceAnnouncementGetMessages**](AdminServiceAnnouncementApi.md#AdminServiceAnnouncementGetMessages) | **Get** /admin/serviceAnnouncement/messages/{serviceUpdateMessage-id} | Get messages from admin
[**AdminServiceAnnouncementHealthOverviewsCreateIssues**](AdminServiceAnnouncementApi.md#AdminServiceAnnouncementHealthOverviewsCreateIssues) | **Post** /admin/serviceAnnouncement/healthOverviews/{serviceHealth-id}/issues | Create new navigation property to issues for admin
[**AdminServiceAnnouncementHealthOverviewsDeleteIssues**](AdminServiceAnnouncementApi.md#AdminServiceAnnouncementHealthOverviewsDeleteIssues) | **Delete** /admin/serviceAnnouncement/healthOverviews/{serviceHealth-id}/issues/{serviceHealthIssue-id} | Delete navigation property issues for admin
[**AdminServiceAnnouncementHealthOverviewsGetIssues**](AdminServiceAnnouncementApi.md#AdminServiceAnnouncementHealthOverviewsGetIssues) | **Get** /admin/serviceAnnouncement/healthOverviews/{serviceHealth-id}/issues/{serviceHealthIssue-id} | Get issues from admin
[**AdminServiceAnnouncementHealthOverviewsListIssues**](AdminServiceAnnouncementApi.md#AdminServiceAnnouncementHealthOverviewsListIssues) | **Get** /admin/serviceAnnouncement/healthOverviews/{serviceHealth-id}/issues | Get issues from admin
[**AdminServiceAnnouncementHealthOverviewsUpdateIssues**](AdminServiceAnnouncementApi.md#AdminServiceAnnouncementHealthOverviewsUpdateIssues) | **Patch** /admin/serviceAnnouncement/healthOverviews/{serviceHealth-id}/issues/{serviceHealthIssue-id} | Update the navigation property issues in admin
[**AdminServiceAnnouncementListHealthOverviews**](AdminServiceAnnouncementApi.md#AdminServiceAnnouncementListHealthOverviews) | **Get** /admin/serviceAnnouncement/healthOverviews | Get healthOverviews from admin
[**AdminServiceAnnouncementListIssues**](AdminServiceAnnouncementApi.md#AdminServiceAnnouncementListIssues) | **Get** /admin/serviceAnnouncement/issues | Get issues from admin
[**AdminServiceAnnouncementListMessages**](AdminServiceAnnouncementApi.md#AdminServiceAnnouncementListMessages) | **Get** /admin/serviceAnnouncement/messages | Get messages from admin
[**AdminServiceAnnouncementUpdateHealthOverviews**](AdminServiceAnnouncementApi.md#AdminServiceAnnouncementUpdateHealthOverviews) | **Patch** /admin/serviceAnnouncement/healthOverviews/{serviceHealth-id} | Update the navigation property healthOverviews in admin
[**AdminServiceAnnouncementUpdateIssues**](AdminServiceAnnouncementApi.md#AdminServiceAnnouncementUpdateIssues) | **Patch** /admin/serviceAnnouncement/issues/{serviceHealthIssue-id} | Update the navigation property issues in admin
[**AdminServiceAnnouncementUpdateMessages**](AdminServiceAnnouncementApi.md#AdminServiceAnnouncementUpdateMessages) | **Patch** /admin/serviceAnnouncement/messages/{serviceUpdateMessage-id} | Update the navigation property messages in admin
[**AdminUpdateServiceAnnouncement**](AdminServiceAnnouncementApi.md#AdminUpdateServiceAnnouncement) | **Patch** /admin/serviceAnnouncement | Update the navigation property serviceAnnouncement in admin



## AdminDeleteServiceAnnouncement

> AdminDeleteServiceAnnouncement(ctx).IfMatch(ifMatch).Execute()

Delete navigation property serviceAnnouncement for admin



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
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AdminServiceAnnouncementApi.AdminDeleteServiceAnnouncement(context.Background()).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminServiceAnnouncementApi.AdminDeleteServiceAnnouncement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminDeleteServiceAnnouncementRequest struct via the builder pattern


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


## AdminGetServiceAnnouncement

> MicrosoftGraphServiceAnnouncement AdminGetServiceAnnouncement(ctx).Select_(select_).Expand(expand).Execute()

Get serviceAnnouncement from admin



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
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AdminServiceAnnouncementApi.AdminGetServiceAnnouncement(context.Background()).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminServiceAnnouncementApi.AdminGetServiceAnnouncement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdminGetServiceAnnouncement`: MicrosoftGraphServiceAnnouncement
    fmt.Fprintf(os.Stdout, "Response from `AdminServiceAnnouncementApi.AdminGetServiceAnnouncement`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminGetServiceAnnouncementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphServiceAnnouncement**](MicrosoftGraphServiceAnnouncement.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminServiceAnnouncementCreateHealthOverviews

> MicrosoftGraphServiceHealth AdminServiceAnnouncementCreateHealthOverviews(ctx).MicrosoftGraphServiceHealth(microsoftGraphServiceHealth).Execute()

Create new navigation property to healthOverviews for admin



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
    microsoftGraphServiceHealth := *openapiclient.NewMicrosoftGraphServiceHealth() // MicrosoftGraphServiceHealth | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AdminServiceAnnouncementApi.AdminServiceAnnouncementCreateHealthOverviews(context.Background()).MicrosoftGraphServiceHealth(microsoftGraphServiceHealth).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminServiceAnnouncementApi.AdminServiceAnnouncementCreateHealthOverviews``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdminServiceAnnouncementCreateHealthOverviews`: MicrosoftGraphServiceHealth
    fmt.Fprintf(os.Stdout, "Response from `AdminServiceAnnouncementApi.AdminServiceAnnouncementCreateHealthOverviews`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminServiceAnnouncementCreateHealthOverviewsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphServiceHealth** | [**MicrosoftGraphServiceHealth**](MicrosoftGraphServiceHealth.md) | New navigation property | 

### Return type

[**MicrosoftGraphServiceHealth**](MicrosoftGraphServiceHealth.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminServiceAnnouncementCreateIssues

> MicrosoftGraphServiceHealthIssue AdminServiceAnnouncementCreateIssues(ctx).MicrosoftGraphServiceHealthIssue(microsoftGraphServiceHealthIssue).Execute()

Create new navigation property to issues for admin



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
    microsoftGraphServiceHealthIssue := *openapiclient.NewMicrosoftGraphServiceHealthIssue() // MicrosoftGraphServiceHealthIssue | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AdminServiceAnnouncementApi.AdminServiceAnnouncementCreateIssues(context.Background()).MicrosoftGraphServiceHealthIssue(microsoftGraphServiceHealthIssue).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminServiceAnnouncementApi.AdminServiceAnnouncementCreateIssues``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdminServiceAnnouncementCreateIssues`: MicrosoftGraphServiceHealthIssue
    fmt.Fprintf(os.Stdout, "Response from `AdminServiceAnnouncementApi.AdminServiceAnnouncementCreateIssues`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminServiceAnnouncementCreateIssuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphServiceHealthIssue** | [**MicrosoftGraphServiceHealthIssue**](MicrosoftGraphServiceHealthIssue.md) | New navigation property | 

### Return type

[**MicrosoftGraphServiceHealthIssue**](MicrosoftGraphServiceHealthIssue.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminServiceAnnouncementCreateMessages

> MicrosoftGraphServiceUpdateMessage AdminServiceAnnouncementCreateMessages(ctx).MicrosoftGraphServiceUpdateMessage(microsoftGraphServiceUpdateMessage).Execute()

Create new navigation property to messages for admin



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
    microsoftGraphServiceUpdateMessage := *openapiclient.NewMicrosoftGraphServiceUpdateMessage() // MicrosoftGraphServiceUpdateMessage | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AdminServiceAnnouncementApi.AdminServiceAnnouncementCreateMessages(context.Background()).MicrosoftGraphServiceUpdateMessage(microsoftGraphServiceUpdateMessage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminServiceAnnouncementApi.AdminServiceAnnouncementCreateMessages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdminServiceAnnouncementCreateMessages`: MicrosoftGraphServiceUpdateMessage
    fmt.Fprintf(os.Stdout, "Response from `AdminServiceAnnouncementApi.AdminServiceAnnouncementCreateMessages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminServiceAnnouncementCreateMessagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphServiceUpdateMessage** | [**MicrosoftGraphServiceUpdateMessage**](MicrosoftGraphServiceUpdateMessage.md) | New navigation property | 

### Return type

[**MicrosoftGraphServiceUpdateMessage**](MicrosoftGraphServiceUpdateMessage.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminServiceAnnouncementDeleteHealthOverviews

> AdminServiceAnnouncementDeleteHealthOverviews(ctx, serviceHealthId).IfMatch(ifMatch).Execute()

Delete navigation property healthOverviews for admin



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
    serviceHealthId := "serviceHealthId_example" // string | key: id of serviceHealth
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AdminServiceAnnouncementApi.AdminServiceAnnouncementDeleteHealthOverviews(context.Background(), serviceHealthId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminServiceAnnouncementApi.AdminServiceAnnouncementDeleteHealthOverviews``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceHealthId** | **string** | key: id of serviceHealth | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminServiceAnnouncementDeleteHealthOverviewsRequest struct via the builder pattern


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


## AdminServiceAnnouncementDeleteIssues

> AdminServiceAnnouncementDeleteIssues(ctx, serviceHealthIssueId).IfMatch(ifMatch).Execute()

Delete navigation property issues for admin



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
    serviceHealthIssueId := "serviceHealthIssueId_example" // string | key: id of serviceHealthIssue
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AdminServiceAnnouncementApi.AdminServiceAnnouncementDeleteIssues(context.Background(), serviceHealthIssueId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminServiceAnnouncementApi.AdminServiceAnnouncementDeleteIssues``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceHealthIssueId** | **string** | key: id of serviceHealthIssue | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminServiceAnnouncementDeleteIssuesRequest struct via the builder pattern


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


## AdminServiceAnnouncementDeleteMessages

> AdminServiceAnnouncementDeleteMessages(ctx, serviceUpdateMessageId).IfMatch(ifMatch).Execute()

Delete navigation property messages for admin



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
    serviceUpdateMessageId := "serviceUpdateMessageId_example" // string | key: id of serviceUpdateMessage
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AdminServiceAnnouncementApi.AdminServiceAnnouncementDeleteMessages(context.Background(), serviceUpdateMessageId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminServiceAnnouncementApi.AdminServiceAnnouncementDeleteMessages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceUpdateMessageId** | **string** | key: id of serviceUpdateMessage | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminServiceAnnouncementDeleteMessagesRequest struct via the builder pattern


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


## AdminServiceAnnouncementGetHealthOverviews

> MicrosoftGraphServiceHealth AdminServiceAnnouncementGetHealthOverviews(ctx, serviceHealthId).Select_(select_).Expand(expand).Execute()

Get healthOverviews from admin



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
    serviceHealthId := "serviceHealthId_example" // string | key: id of serviceHealth
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AdminServiceAnnouncementApi.AdminServiceAnnouncementGetHealthOverviews(context.Background(), serviceHealthId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminServiceAnnouncementApi.AdminServiceAnnouncementGetHealthOverviews``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdminServiceAnnouncementGetHealthOverviews`: MicrosoftGraphServiceHealth
    fmt.Fprintf(os.Stdout, "Response from `AdminServiceAnnouncementApi.AdminServiceAnnouncementGetHealthOverviews`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceHealthId** | **string** | key: id of serviceHealth | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminServiceAnnouncementGetHealthOverviewsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphServiceHealth**](MicrosoftGraphServiceHealth.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminServiceAnnouncementGetIssues

> MicrosoftGraphServiceHealthIssue AdminServiceAnnouncementGetIssues(ctx, serviceHealthIssueId).Select_(select_).Expand(expand).Execute()

Get issues from admin



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
    serviceHealthIssueId := "serviceHealthIssueId_example" // string | key: id of serviceHealthIssue
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AdminServiceAnnouncementApi.AdminServiceAnnouncementGetIssues(context.Background(), serviceHealthIssueId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminServiceAnnouncementApi.AdminServiceAnnouncementGetIssues``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdminServiceAnnouncementGetIssues`: MicrosoftGraphServiceHealthIssue
    fmt.Fprintf(os.Stdout, "Response from `AdminServiceAnnouncementApi.AdminServiceAnnouncementGetIssues`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceHealthIssueId** | **string** | key: id of serviceHealthIssue | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminServiceAnnouncementGetIssuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphServiceHealthIssue**](MicrosoftGraphServiceHealthIssue.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminServiceAnnouncementGetMessages

> MicrosoftGraphServiceUpdateMessage AdminServiceAnnouncementGetMessages(ctx, serviceUpdateMessageId).Select_(select_).Expand(expand).Execute()

Get messages from admin



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
    serviceUpdateMessageId := "serviceUpdateMessageId_example" // string | key: id of serviceUpdateMessage
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AdminServiceAnnouncementApi.AdminServiceAnnouncementGetMessages(context.Background(), serviceUpdateMessageId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminServiceAnnouncementApi.AdminServiceAnnouncementGetMessages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdminServiceAnnouncementGetMessages`: MicrosoftGraphServiceUpdateMessage
    fmt.Fprintf(os.Stdout, "Response from `AdminServiceAnnouncementApi.AdminServiceAnnouncementGetMessages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceUpdateMessageId** | **string** | key: id of serviceUpdateMessage | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminServiceAnnouncementGetMessagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphServiceUpdateMessage**](MicrosoftGraphServiceUpdateMessage.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminServiceAnnouncementHealthOverviewsCreateIssues

> MicrosoftGraphServiceHealthIssue AdminServiceAnnouncementHealthOverviewsCreateIssues(ctx, serviceHealthId).MicrosoftGraphServiceHealthIssue(microsoftGraphServiceHealthIssue).Execute()

Create new navigation property to issues for admin



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
    serviceHealthId := "serviceHealthId_example" // string | key: id of serviceHealth
    microsoftGraphServiceHealthIssue := *openapiclient.NewMicrosoftGraphServiceHealthIssue() // MicrosoftGraphServiceHealthIssue | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AdminServiceAnnouncementApi.AdminServiceAnnouncementHealthOverviewsCreateIssues(context.Background(), serviceHealthId).MicrosoftGraphServiceHealthIssue(microsoftGraphServiceHealthIssue).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminServiceAnnouncementApi.AdminServiceAnnouncementHealthOverviewsCreateIssues``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdminServiceAnnouncementHealthOverviewsCreateIssues`: MicrosoftGraphServiceHealthIssue
    fmt.Fprintf(os.Stdout, "Response from `AdminServiceAnnouncementApi.AdminServiceAnnouncementHealthOverviewsCreateIssues`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceHealthId** | **string** | key: id of serviceHealth | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminServiceAnnouncementHealthOverviewsCreateIssuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphServiceHealthIssue** | [**MicrosoftGraphServiceHealthIssue**](MicrosoftGraphServiceHealthIssue.md) | New navigation property | 

### Return type

[**MicrosoftGraphServiceHealthIssue**](MicrosoftGraphServiceHealthIssue.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminServiceAnnouncementHealthOverviewsDeleteIssues

> AdminServiceAnnouncementHealthOverviewsDeleteIssues(ctx, serviceHealthId, serviceHealthIssueId).IfMatch(ifMatch).Execute()

Delete navigation property issues for admin



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
    serviceHealthId := "serviceHealthId_example" // string | key: id of serviceHealth
    serviceHealthIssueId := "serviceHealthIssueId_example" // string | key: id of serviceHealthIssue
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AdminServiceAnnouncementApi.AdminServiceAnnouncementHealthOverviewsDeleteIssues(context.Background(), serviceHealthId, serviceHealthIssueId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminServiceAnnouncementApi.AdminServiceAnnouncementHealthOverviewsDeleteIssues``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceHealthId** | **string** | key: id of serviceHealth | 
**serviceHealthIssueId** | **string** | key: id of serviceHealthIssue | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminServiceAnnouncementHealthOverviewsDeleteIssuesRequest struct via the builder pattern


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


## AdminServiceAnnouncementHealthOverviewsGetIssues

> MicrosoftGraphServiceHealthIssue AdminServiceAnnouncementHealthOverviewsGetIssues(ctx, serviceHealthId, serviceHealthIssueId).Select_(select_).Expand(expand).Execute()

Get issues from admin



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
    serviceHealthId := "serviceHealthId_example" // string | key: id of serviceHealth
    serviceHealthIssueId := "serviceHealthIssueId_example" // string | key: id of serviceHealthIssue
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AdminServiceAnnouncementApi.AdminServiceAnnouncementHealthOverviewsGetIssues(context.Background(), serviceHealthId, serviceHealthIssueId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminServiceAnnouncementApi.AdminServiceAnnouncementHealthOverviewsGetIssues``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdminServiceAnnouncementHealthOverviewsGetIssues`: MicrosoftGraphServiceHealthIssue
    fmt.Fprintf(os.Stdout, "Response from `AdminServiceAnnouncementApi.AdminServiceAnnouncementHealthOverviewsGetIssues`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceHealthId** | **string** | key: id of serviceHealth | 
**serviceHealthIssueId** | **string** | key: id of serviceHealthIssue | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminServiceAnnouncementHealthOverviewsGetIssuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphServiceHealthIssue**](MicrosoftGraphServiceHealthIssue.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminServiceAnnouncementHealthOverviewsListIssues

> CollectionOfServiceHealthIssue AdminServiceAnnouncementHealthOverviewsListIssues(ctx, serviceHealthId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get issues from admin



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
    serviceHealthId := "serviceHealthId_example" // string | key: id of serviceHealth
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
    resp, r, err := api_client.AdminServiceAnnouncementApi.AdminServiceAnnouncementHealthOverviewsListIssues(context.Background(), serviceHealthId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminServiceAnnouncementApi.AdminServiceAnnouncementHealthOverviewsListIssues``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdminServiceAnnouncementHealthOverviewsListIssues`: CollectionOfServiceHealthIssue
    fmt.Fprintf(os.Stdout, "Response from `AdminServiceAnnouncementApi.AdminServiceAnnouncementHealthOverviewsListIssues`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceHealthId** | **string** | key: id of serviceHealth | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminServiceAnnouncementHealthOverviewsListIssuesRequest struct via the builder pattern


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

[**CollectionOfServiceHealthIssue**](CollectionOfServiceHealthIssue.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminServiceAnnouncementHealthOverviewsUpdateIssues

> AdminServiceAnnouncementHealthOverviewsUpdateIssues(ctx, serviceHealthId, serviceHealthIssueId).MicrosoftGraphServiceHealthIssue(microsoftGraphServiceHealthIssue).Execute()

Update the navigation property issues in admin



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
    serviceHealthId := "serviceHealthId_example" // string | key: id of serviceHealth
    serviceHealthIssueId := "serviceHealthIssueId_example" // string | key: id of serviceHealthIssue
    microsoftGraphServiceHealthIssue := *openapiclient.NewMicrosoftGraphServiceHealthIssue() // MicrosoftGraphServiceHealthIssue | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AdminServiceAnnouncementApi.AdminServiceAnnouncementHealthOverviewsUpdateIssues(context.Background(), serviceHealthId, serviceHealthIssueId).MicrosoftGraphServiceHealthIssue(microsoftGraphServiceHealthIssue).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminServiceAnnouncementApi.AdminServiceAnnouncementHealthOverviewsUpdateIssues``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceHealthId** | **string** | key: id of serviceHealth | 
**serviceHealthIssueId** | **string** | key: id of serviceHealthIssue | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminServiceAnnouncementHealthOverviewsUpdateIssuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **microsoftGraphServiceHealthIssue** | [**MicrosoftGraphServiceHealthIssue**](MicrosoftGraphServiceHealthIssue.md) | New navigation property values | 

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


## AdminServiceAnnouncementListHealthOverviews

> CollectionOfServiceHealth AdminServiceAnnouncementListHealthOverviews(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get healthOverviews from admin



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
    resp, r, err := api_client.AdminServiceAnnouncementApi.AdminServiceAnnouncementListHealthOverviews(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminServiceAnnouncementApi.AdminServiceAnnouncementListHealthOverviews``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdminServiceAnnouncementListHealthOverviews`: CollectionOfServiceHealth
    fmt.Fprintf(os.Stdout, "Response from `AdminServiceAnnouncementApi.AdminServiceAnnouncementListHealthOverviews`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminServiceAnnouncementListHealthOverviewsRequest struct via the builder pattern


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

[**CollectionOfServiceHealth**](CollectionOfServiceHealth.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminServiceAnnouncementListIssues

> CollectionOfServiceHealthIssue AdminServiceAnnouncementListIssues(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get issues from admin



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
    resp, r, err := api_client.AdminServiceAnnouncementApi.AdminServiceAnnouncementListIssues(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminServiceAnnouncementApi.AdminServiceAnnouncementListIssues``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdminServiceAnnouncementListIssues`: CollectionOfServiceHealthIssue
    fmt.Fprintf(os.Stdout, "Response from `AdminServiceAnnouncementApi.AdminServiceAnnouncementListIssues`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminServiceAnnouncementListIssuesRequest struct via the builder pattern


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

[**CollectionOfServiceHealthIssue**](CollectionOfServiceHealthIssue.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminServiceAnnouncementListMessages

> CollectionOfServiceUpdateMessage AdminServiceAnnouncementListMessages(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get messages from admin



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
    resp, r, err := api_client.AdminServiceAnnouncementApi.AdminServiceAnnouncementListMessages(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminServiceAnnouncementApi.AdminServiceAnnouncementListMessages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdminServiceAnnouncementListMessages`: CollectionOfServiceUpdateMessage
    fmt.Fprintf(os.Stdout, "Response from `AdminServiceAnnouncementApi.AdminServiceAnnouncementListMessages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminServiceAnnouncementListMessagesRequest struct via the builder pattern


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

[**CollectionOfServiceUpdateMessage**](CollectionOfServiceUpdateMessage.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminServiceAnnouncementUpdateHealthOverviews

> AdminServiceAnnouncementUpdateHealthOverviews(ctx, serviceHealthId).MicrosoftGraphServiceHealth(microsoftGraphServiceHealth).Execute()

Update the navigation property healthOverviews in admin



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
    serviceHealthId := "serviceHealthId_example" // string | key: id of serviceHealth
    microsoftGraphServiceHealth := *openapiclient.NewMicrosoftGraphServiceHealth() // MicrosoftGraphServiceHealth | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AdminServiceAnnouncementApi.AdminServiceAnnouncementUpdateHealthOverviews(context.Background(), serviceHealthId).MicrosoftGraphServiceHealth(microsoftGraphServiceHealth).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminServiceAnnouncementApi.AdminServiceAnnouncementUpdateHealthOverviews``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceHealthId** | **string** | key: id of serviceHealth | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminServiceAnnouncementUpdateHealthOverviewsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphServiceHealth** | [**MicrosoftGraphServiceHealth**](MicrosoftGraphServiceHealth.md) | New navigation property values | 

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


## AdminServiceAnnouncementUpdateIssues

> AdminServiceAnnouncementUpdateIssues(ctx, serviceHealthIssueId).MicrosoftGraphServiceHealthIssue(microsoftGraphServiceHealthIssue).Execute()

Update the navigation property issues in admin



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
    serviceHealthIssueId := "serviceHealthIssueId_example" // string | key: id of serviceHealthIssue
    microsoftGraphServiceHealthIssue := *openapiclient.NewMicrosoftGraphServiceHealthIssue() // MicrosoftGraphServiceHealthIssue | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AdminServiceAnnouncementApi.AdminServiceAnnouncementUpdateIssues(context.Background(), serviceHealthIssueId).MicrosoftGraphServiceHealthIssue(microsoftGraphServiceHealthIssue).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminServiceAnnouncementApi.AdminServiceAnnouncementUpdateIssues``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceHealthIssueId** | **string** | key: id of serviceHealthIssue | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminServiceAnnouncementUpdateIssuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphServiceHealthIssue** | [**MicrosoftGraphServiceHealthIssue**](MicrosoftGraphServiceHealthIssue.md) | New navigation property values | 

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


## AdminServiceAnnouncementUpdateMessages

> AdminServiceAnnouncementUpdateMessages(ctx, serviceUpdateMessageId).MicrosoftGraphServiceUpdateMessage(microsoftGraphServiceUpdateMessage).Execute()

Update the navigation property messages in admin



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
    serviceUpdateMessageId := "serviceUpdateMessageId_example" // string | key: id of serviceUpdateMessage
    microsoftGraphServiceUpdateMessage := *openapiclient.NewMicrosoftGraphServiceUpdateMessage() // MicrosoftGraphServiceUpdateMessage | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AdminServiceAnnouncementApi.AdminServiceAnnouncementUpdateMessages(context.Background(), serviceUpdateMessageId).MicrosoftGraphServiceUpdateMessage(microsoftGraphServiceUpdateMessage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminServiceAnnouncementApi.AdminServiceAnnouncementUpdateMessages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceUpdateMessageId** | **string** | key: id of serviceUpdateMessage | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminServiceAnnouncementUpdateMessagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphServiceUpdateMessage** | [**MicrosoftGraphServiceUpdateMessage**](MicrosoftGraphServiceUpdateMessage.md) | New navigation property values | 

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


## AdminUpdateServiceAnnouncement

> AdminUpdateServiceAnnouncement(ctx).MicrosoftGraphServiceAnnouncement(microsoftGraphServiceAnnouncement).Execute()

Update the navigation property serviceAnnouncement in admin



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
    microsoftGraphServiceAnnouncement := *openapiclient.NewMicrosoftGraphServiceAnnouncement() // MicrosoftGraphServiceAnnouncement | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AdminServiceAnnouncementApi.AdminUpdateServiceAnnouncement(context.Background()).MicrosoftGraphServiceAnnouncement(microsoftGraphServiceAnnouncement).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminServiceAnnouncementApi.AdminUpdateServiceAnnouncement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminUpdateServiceAnnouncementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphServiceAnnouncement** | [**MicrosoftGraphServiceAnnouncement**](MicrosoftGraphServiceAnnouncement.md) | New navigation property values | 

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

