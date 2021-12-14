# \DriveActionsApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DriveListContentTypesAddCopy**](DriveActionsApi.md#DriveListContentTypesAddCopy) | **Post** /drive/list/contentTypes/microsoft.graph.addCopy | Invoke action addCopy
[**DriveListContentTypesContentTypeAssociateWithHubSites**](DriveActionsApi.md#DriveListContentTypesContentTypeAssociateWithHubSites) | **Post** /drive/list/contentTypes/{contentType-id}/microsoft.graph.associateWithHubSites | Invoke action associateWithHubSites
[**DriveListContentTypesContentTypeBaseAssociateWithHubSites**](DriveActionsApi.md#DriveListContentTypesContentTypeBaseAssociateWithHubSites) | **Post** /drive/list/contentTypes/{contentType-id}/base/microsoft.graph.associateWithHubSites | Invoke action associateWithHubSites
[**DriveListContentTypesContentTypeBaseCopyToDefaultContentLocation**](DriveActionsApi.md#DriveListContentTypesContentTypeBaseCopyToDefaultContentLocation) | **Post** /drive/list/contentTypes/{contentType-id}/base/microsoft.graph.copyToDefaultContentLocation | Invoke action copyToDefaultContentLocation
[**DriveListContentTypesContentTypeBasePublish**](DriveActionsApi.md#DriveListContentTypesContentTypeBasePublish) | **Post** /drive/list/contentTypes/{contentType-id}/base/microsoft.graph.publish | Invoke action publish
[**DriveListContentTypesContentTypeBaseTypesAddCopy**](DriveActionsApi.md#DriveListContentTypesContentTypeBaseTypesAddCopy) | **Post** /drive/list/contentTypes/{contentType-id}/baseTypes/microsoft.graph.addCopy | Invoke action addCopy
[**DriveListContentTypesContentTypeBaseUnpublish**](DriveActionsApi.md#DriveListContentTypesContentTypeBaseUnpublish) | **Post** /drive/list/contentTypes/{contentType-id}/base/microsoft.graph.unpublish | Invoke action unpublish
[**DriveListContentTypesContentTypeCopyToDefaultContentLocation**](DriveActionsApi.md#DriveListContentTypesContentTypeCopyToDefaultContentLocation) | **Post** /drive/list/contentTypes/{contentType-id}/microsoft.graph.copyToDefaultContentLocation | Invoke action copyToDefaultContentLocation
[**DriveListContentTypesContentTypePublish**](DriveActionsApi.md#DriveListContentTypesContentTypePublish) | **Post** /drive/list/contentTypes/{contentType-id}/microsoft.graph.publish | Invoke action publish
[**DriveListContentTypesContentTypeUnpublish**](DriveActionsApi.md#DriveListContentTypesContentTypeUnpublish) | **Post** /drive/list/contentTypes/{contentType-id}/microsoft.graph.unpublish | Invoke action unpublish
[**DriveListItemsListItemVersionsListItemVersionRestoreVersion**](DriveActionsApi.md#DriveListItemsListItemVersionsListItemVersionRestoreVersion) | **Post** /drive/list/items/{listItem-id}/versions/{listItemVersion-id}/microsoft.graph.restoreVersion | Invoke action restoreVersion



## DriveListContentTypesAddCopy

> AnyOfmicrosoftGraphContentType DriveListContentTypesAddCopy(ctx).InlineObject140(inlineObject140).Execute()

Invoke action addCopy

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
    inlineObject140 := *openapiclient.NewInlineObject140() // InlineObject140 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DriveActionsApi.DriveListContentTypesAddCopy(context.Background()).InlineObject140(inlineObject140).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DriveActionsApi.DriveListContentTypesAddCopy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DriveListContentTypesAddCopy`: AnyOfmicrosoftGraphContentType
    fmt.Fprintf(os.Stdout, "Response from `DriveActionsApi.DriveListContentTypesAddCopy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDriveListContentTypesAddCopyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject140** | [**InlineObject140**](InlineObject140.md) |  | 

### Return type

[**AnyOfmicrosoftGraphContentType**](anyOf&lt;microsoft.graph.contentType&gt;.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DriveListContentTypesContentTypeAssociateWithHubSites

> DriveListContentTypesContentTypeAssociateWithHubSites(ctx, contentTypeId).InlineObject138(inlineObject138).Execute()

Invoke action associateWithHubSites

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
    contentTypeId := "contentTypeId_example" // string | key: id of contentType
    inlineObject138 := *openapiclient.NewInlineObject138() // InlineObject138 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DriveActionsApi.DriveListContentTypesContentTypeAssociateWithHubSites(context.Background(), contentTypeId).InlineObject138(inlineObject138).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DriveActionsApi.DriveListContentTypesContentTypeAssociateWithHubSites``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contentTypeId** | **string** | key: id of contentType | 

### Other Parameters

Other parameters are passed through a pointer to a apiDriveListContentTypesContentTypeAssociateWithHubSitesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject138** | [**InlineObject138**](InlineObject138.md) |  | 

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


## DriveListContentTypesContentTypeBaseAssociateWithHubSites

> DriveListContentTypesContentTypeBaseAssociateWithHubSites(ctx, contentTypeId).InlineObject135(inlineObject135).Execute()

Invoke action associateWithHubSites

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
    contentTypeId := "contentTypeId_example" // string | key: id of contentType
    inlineObject135 := *openapiclient.NewInlineObject135() // InlineObject135 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DriveActionsApi.DriveListContentTypesContentTypeBaseAssociateWithHubSites(context.Background(), contentTypeId).InlineObject135(inlineObject135).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DriveActionsApi.DriveListContentTypesContentTypeBaseAssociateWithHubSites``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contentTypeId** | **string** | key: id of contentType | 

### Other Parameters

Other parameters are passed through a pointer to a apiDriveListContentTypesContentTypeBaseAssociateWithHubSitesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject135** | [**InlineObject135**](InlineObject135.md) |  | 

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


## DriveListContentTypesContentTypeBaseCopyToDefaultContentLocation

> DriveListContentTypesContentTypeBaseCopyToDefaultContentLocation(ctx, contentTypeId).InlineObject136(inlineObject136).Execute()

Invoke action copyToDefaultContentLocation

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
    contentTypeId := "contentTypeId_example" // string | key: id of contentType
    inlineObject136 := *openapiclient.NewInlineObject136() // InlineObject136 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DriveActionsApi.DriveListContentTypesContentTypeBaseCopyToDefaultContentLocation(context.Background(), contentTypeId).InlineObject136(inlineObject136).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DriveActionsApi.DriveListContentTypesContentTypeBaseCopyToDefaultContentLocation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contentTypeId** | **string** | key: id of contentType | 

### Other Parameters

Other parameters are passed through a pointer to a apiDriveListContentTypesContentTypeBaseCopyToDefaultContentLocationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject136** | [**InlineObject136**](InlineObject136.md) |  | 

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


## DriveListContentTypesContentTypeBasePublish

> DriveListContentTypesContentTypeBasePublish(ctx, contentTypeId).Execute()

Invoke action publish

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
    contentTypeId := "contentTypeId_example" // string | key: id of contentType

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DriveActionsApi.DriveListContentTypesContentTypeBasePublish(context.Background(), contentTypeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DriveActionsApi.DriveListContentTypesContentTypeBasePublish``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contentTypeId** | **string** | key: id of contentType | 

### Other Parameters

Other parameters are passed through a pointer to a apiDriveListContentTypesContentTypeBasePublishRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## DriveListContentTypesContentTypeBaseTypesAddCopy

> AnyOfmicrosoftGraphContentType DriveListContentTypesContentTypeBaseTypesAddCopy(ctx, contentTypeId).InlineObject137(inlineObject137).Execute()

Invoke action addCopy

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
    contentTypeId := "contentTypeId_example" // string | key: id of contentType
    inlineObject137 := *openapiclient.NewInlineObject137() // InlineObject137 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DriveActionsApi.DriveListContentTypesContentTypeBaseTypesAddCopy(context.Background(), contentTypeId).InlineObject137(inlineObject137).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DriveActionsApi.DriveListContentTypesContentTypeBaseTypesAddCopy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DriveListContentTypesContentTypeBaseTypesAddCopy`: AnyOfmicrosoftGraphContentType
    fmt.Fprintf(os.Stdout, "Response from `DriveActionsApi.DriveListContentTypesContentTypeBaseTypesAddCopy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contentTypeId** | **string** | key: id of contentType | 

### Other Parameters

Other parameters are passed through a pointer to a apiDriveListContentTypesContentTypeBaseTypesAddCopyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject137** | [**InlineObject137**](InlineObject137.md) |  | 

### Return type

[**AnyOfmicrosoftGraphContentType**](anyOf&lt;microsoft.graph.contentType&gt;.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DriveListContentTypesContentTypeBaseUnpublish

> DriveListContentTypesContentTypeBaseUnpublish(ctx, contentTypeId).Execute()

Invoke action unpublish

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
    contentTypeId := "contentTypeId_example" // string | key: id of contentType

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DriveActionsApi.DriveListContentTypesContentTypeBaseUnpublish(context.Background(), contentTypeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DriveActionsApi.DriveListContentTypesContentTypeBaseUnpublish``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contentTypeId** | **string** | key: id of contentType | 

### Other Parameters

Other parameters are passed through a pointer to a apiDriveListContentTypesContentTypeBaseUnpublishRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## DriveListContentTypesContentTypeCopyToDefaultContentLocation

> DriveListContentTypesContentTypeCopyToDefaultContentLocation(ctx, contentTypeId).InlineObject139(inlineObject139).Execute()

Invoke action copyToDefaultContentLocation

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
    contentTypeId := "contentTypeId_example" // string | key: id of contentType
    inlineObject139 := *openapiclient.NewInlineObject139() // InlineObject139 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DriveActionsApi.DriveListContentTypesContentTypeCopyToDefaultContentLocation(context.Background(), contentTypeId).InlineObject139(inlineObject139).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DriveActionsApi.DriveListContentTypesContentTypeCopyToDefaultContentLocation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contentTypeId** | **string** | key: id of contentType | 

### Other Parameters

Other parameters are passed through a pointer to a apiDriveListContentTypesContentTypeCopyToDefaultContentLocationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject139** | [**InlineObject139**](InlineObject139.md) |  | 

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


## DriveListContentTypesContentTypePublish

> DriveListContentTypesContentTypePublish(ctx, contentTypeId).Execute()

Invoke action publish

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
    contentTypeId := "contentTypeId_example" // string | key: id of contentType

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DriveActionsApi.DriveListContentTypesContentTypePublish(context.Background(), contentTypeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DriveActionsApi.DriveListContentTypesContentTypePublish``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contentTypeId** | **string** | key: id of contentType | 

### Other Parameters

Other parameters are passed through a pointer to a apiDriveListContentTypesContentTypePublishRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## DriveListContentTypesContentTypeUnpublish

> DriveListContentTypesContentTypeUnpublish(ctx, contentTypeId).Execute()

Invoke action unpublish

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
    contentTypeId := "contentTypeId_example" // string | key: id of contentType

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DriveActionsApi.DriveListContentTypesContentTypeUnpublish(context.Background(), contentTypeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DriveActionsApi.DriveListContentTypesContentTypeUnpublish``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contentTypeId** | **string** | key: id of contentType | 

### Other Parameters

Other parameters are passed through a pointer to a apiDriveListContentTypesContentTypeUnpublishRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## DriveListItemsListItemVersionsListItemVersionRestoreVersion

> DriveListItemsListItemVersionsListItemVersionRestoreVersion(ctx, listItemId, listItemVersionId).Execute()

Invoke action restoreVersion

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
    listItemId := "listItemId_example" // string | key: id of listItem
    listItemVersionId := "listItemVersionId_example" // string | key: id of listItemVersion

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DriveActionsApi.DriveListItemsListItemVersionsListItemVersionRestoreVersion(context.Background(), listItemId, listItemVersionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DriveActionsApi.DriveListItemsListItemVersionsListItemVersionRestoreVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listItemId** | **string** | key: id of listItem | 
**listItemVersionId** | **string** | key: id of listItemVersion | 

### Other Parameters

Other parameters are passed through a pointer to a apiDriveListItemsListItemVersionsListItemVersionRestoreVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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

