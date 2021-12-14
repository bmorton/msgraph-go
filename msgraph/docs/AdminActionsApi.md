# \AdminActionsApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdminServiceAnnouncementMessagesArchive**](AdminActionsApi.md#AdminServiceAnnouncementMessagesArchive) | **Post** /admin/serviceAnnouncement/messages/microsoft.graph.archive | Invoke action archive
[**AdminServiceAnnouncementMessagesFavorite**](AdminActionsApi.md#AdminServiceAnnouncementMessagesFavorite) | **Post** /admin/serviceAnnouncement/messages/microsoft.graph.favorite | Invoke action favorite
[**AdminServiceAnnouncementMessagesMarkRead**](AdminActionsApi.md#AdminServiceAnnouncementMessagesMarkRead) | **Post** /admin/serviceAnnouncement/messages/microsoft.graph.markRead | Invoke action markRead
[**AdminServiceAnnouncementMessagesMarkUnread**](AdminActionsApi.md#AdminServiceAnnouncementMessagesMarkUnread) | **Post** /admin/serviceAnnouncement/messages/microsoft.graph.markUnread | Invoke action markUnread
[**AdminServiceAnnouncementMessagesUnarchive**](AdminActionsApi.md#AdminServiceAnnouncementMessagesUnarchive) | **Post** /admin/serviceAnnouncement/messages/microsoft.graph.unarchive | Invoke action unarchive
[**AdminServiceAnnouncementMessagesUnfavorite**](AdminActionsApi.md#AdminServiceAnnouncementMessagesUnfavorite) | **Post** /admin/serviceAnnouncement/messages/microsoft.graph.unfavorite | Invoke action unfavorite



## AdminServiceAnnouncementMessagesArchive

> bool AdminServiceAnnouncementMessagesArchive(ctx).InlineObject(inlineObject).Execute()

Invoke action archive

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
    inlineObject := *openapiclient.NewInlineObject() // InlineObject | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AdminActionsApi.AdminServiceAnnouncementMessagesArchive(context.Background()).InlineObject(inlineObject).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminActionsApi.AdminServiceAnnouncementMessagesArchive``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdminServiceAnnouncementMessagesArchive`: bool
    fmt.Fprintf(os.Stdout, "Response from `AdminActionsApi.AdminServiceAnnouncementMessagesArchive`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminServiceAnnouncementMessagesArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject** | [**InlineObject**](InlineObject.md) |  | 

### Return type

**bool**

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminServiceAnnouncementMessagesFavorite

> bool AdminServiceAnnouncementMessagesFavorite(ctx).InlineObject1(inlineObject1).Execute()

Invoke action favorite

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
    inlineObject1 := *openapiclient.NewInlineObject1() // InlineObject1 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AdminActionsApi.AdminServiceAnnouncementMessagesFavorite(context.Background()).InlineObject1(inlineObject1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminActionsApi.AdminServiceAnnouncementMessagesFavorite``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdminServiceAnnouncementMessagesFavorite`: bool
    fmt.Fprintf(os.Stdout, "Response from `AdminActionsApi.AdminServiceAnnouncementMessagesFavorite`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminServiceAnnouncementMessagesFavoriteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject1** | [**InlineObject1**](InlineObject1.md) |  | 

### Return type

**bool**

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminServiceAnnouncementMessagesMarkRead

> bool AdminServiceAnnouncementMessagesMarkRead(ctx).InlineObject2(inlineObject2).Execute()

Invoke action markRead

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
    inlineObject2 := *openapiclient.NewInlineObject2() // InlineObject2 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AdminActionsApi.AdminServiceAnnouncementMessagesMarkRead(context.Background()).InlineObject2(inlineObject2).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminActionsApi.AdminServiceAnnouncementMessagesMarkRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdminServiceAnnouncementMessagesMarkRead`: bool
    fmt.Fprintf(os.Stdout, "Response from `AdminActionsApi.AdminServiceAnnouncementMessagesMarkRead`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminServiceAnnouncementMessagesMarkReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject2** | [**InlineObject2**](InlineObject2.md) |  | 

### Return type

**bool**

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminServiceAnnouncementMessagesMarkUnread

> bool AdminServiceAnnouncementMessagesMarkUnread(ctx).InlineObject3(inlineObject3).Execute()

Invoke action markUnread

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
    inlineObject3 := *openapiclient.NewInlineObject3() // InlineObject3 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AdminActionsApi.AdminServiceAnnouncementMessagesMarkUnread(context.Background()).InlineObject3(inlineObject3).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminActionsApi.AdminServiceAnnouncementMessagesMarkUnread``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdminServiceAnnouncementMessagesMarkUnread`: bool
    fmt.Fprintf(os.Stdout, "Response from `AdminActionsApi.AdminServiceAnnouncementMessagesMarkUnread`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminServiceAnnouncementMessagesMarkUnreadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject3** | [**InlineObject3**](InlineObject3.md) |  | 

### Return type

**bool**

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminServiceAnnouncementMessagesUnarchive

> bool AdminServiceAnnouncementMessagesUnarchive(ctx).InlineObject4(inlineObject4).Execute()

Invoke action unarchive

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
    inlineObject4 := *openapiclient.NewInlineObject4() // InlineObject4 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AdminActionsApi.AdminServiceAnnouncementMessagesUnarchive(context.Background()).InlineObject4(inlineObject4).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminActionsApi.AdminServiceAnnouncementMessagesUnarchive``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdminServiceAnnouncementMessagesUnarchive`: bool
    fmt.Fprintf(os.Stdout, "Response from `AdminActionsApi.AdminServiceAnnouncementMessagesUnarchive`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminServiceAnnouncementMessagesUnarchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject4** | [**InlineObject4**](InlineObject4.md) |  | 

### Return type

**bool**

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminServiceAnnouncementMessagesUnfavorite

> bool AdminServiceAnnouncementMessagesUnfavorite(ctx).InlineObject5(inlineObject5).Execute()

Invoke action unfavorite

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
    inlineObject5 := *openapiclient.NewInlineObject5() // InlineObject5 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AdminActionsApi.AdminServiceAnnouncementMessagesUnfavorite(context.Background()).InlineObject5(inlineObject5).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminActionsApi.AdminServiceAnnouncementMessagesUnfavorite``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdminServiceAnnouncementMessagesUnfavorite`: bool
    fmt.Fprintf(os.Stdout, "Response from `AdminActionsApi.AdminServiceAnnouncementMessagesUnfavorite`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminServiceAnnouncementMessagesUnfavoriteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject5** | [**InlineObject5**](InlineObject5.md) |  | 

### Return type

**bool**

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

