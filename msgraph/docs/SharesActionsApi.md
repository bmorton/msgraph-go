# \SharesActionsApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SharesSharedDriveItemListContentTypesAddCopy**](SharesActionsApi.md#SharesSharedDriveItemListContentTypesAddCopy) | **Post** /shares/{sharedDriveItem-id}/list/contentTypes/microsoft.graph.addCopy | Invoke action addCopy
[**SharesSharedDriveItemListContentTypesContentTypeAssociateWithHubSites**](SharesActionsApi.md#SharesSharedDriveItemListContentTypesContentTypeAssociateWithHubSites) | **Post** /shares/{sharedDriveItem-id}/list/contentTypes/{contentType-id}/microsoft.graph.associateWithHubSites | Invoke action associateWithHubSites
[**SharesSharedDriveItemListContentTypesContentTypeBaseAssociateWithHubSites**](SharesActionsApi.md#SharesSharedDriveItemListContentTypesContentTypeBaseAssociateWithHubSites) | **Post** /shares/{sharedDriveItem-id}/list/contentTypes/{contentType-id}/base/microsoft.graph.associateWithHubSites | Invoke action associateWithHubSites
[**SharesSharedDriveItemListContentTypesContentTypeBaseCopyToDefaultContentLocation**](SharesActionsApi.md#SharesSharedDriveItemListContentTypesContentTypeBaseCopyToDefaultContentLocation) | **Post** /shares/{sharedDriveItem-id}/list/contentTypes/{contentType-id}/base/microsoft.graph.copyToDefaultContentLocation | Invoke action copyToDefaultContentLocation
[**SharesSharedDriveItemListContentTypesContentTypeBasePublish**](SharesActionsApi.md#SharesSharedDriveItemListContentTypesContentTypeBasePublish) | **Post** /shares/{sharedDriveItem-id}/list/contentTypes/{contentType-id}/base/microsoft.graph.publish | Invoke action publish
[**SharesSharedDriveItemListContentTypesContentTypeBaseTypesAddCopy**](SharesActionsApi.md#SharesSharedDriveItemListContentTypesContentTypeBaseTypesAddCopy) | **Post** /shares/{sharedDriveItem-id}/list/contentTypes/{contentType-id}/baseTypes/microsoft.graph.addCopy | Invoke action addCopy
[**SharesSharedDriveItemListContentTypesContentTypeBaseUnpublish**](SharesActionsApi.md#SharesSharedDriveItemListContentTypesContentTypeBaseUnpublish) | **Post** /shares/{sharedDriveItem-id}/list/contentTypes/{contentType-id}/base/microsoft.graph.unpublish | Invoke action unpublish
[**SharesSharedDriveItemListContentTypesContentTypeCopyToDefaultContentLocation**](SharesActionsApi.md#SharesSharedDriveItemListContentTypesContentTypeCopyToDefaultContentLocation) | **Post** /shares/{sharedDriveItem-id}/list/contentTypes/{contentType-id}/microsoft.graph.copyToDefaultContentLocation | Invoke action copyToDefaultContentLocation
[**SharesSharedDriveItemListContentTypesContentTypePublish**](SharesActionsApi.md#SharesSharedDriveItemListContentTypesContentTypePublish) | **Post** /shares/{sharedDriveItem-id}/list/contentTypes/{contentType-id}/microsoft.graph.publish | Invoke action publish
[**SharesSharedDriveItemListContentTypesContentTypeUnpublish**](SharesActionsApi.md#SharesSharedDriveItemListContentTypesContentTypeUnpublish) | **Post** /shares/{sharedDriveItem-id}/list/contentTypes/{contentType-id}/microsoft.graph.unpublish | Invoke action unpublish
[**SharesSharedDriveItemListItemVersionsListItemVersionRestoreVersion**](SharesActionsApi.md#SharesSharedDriveItemListItemVersionsListItemVersionRestoreVersion) | **Post** /shares/{sharedDriveItem-id}/listItem/versions/{listItemVersion-id}/microsoft.graph.restoreVersion | Invoke action restoreVersion
[**SharesSharedDriveItemListItemsListItemVersionsListItemVersionRestoreVersion**](SharesActionsApi.md#SharesSharedDriveItemListItemsListItemVersionsListItemVersionRestoreVersion) | **Post** /shares/{sharedDriveItem-id}/list/items/{listItem-id}/versions/{listItemVersion-id}/microsoft.graph.restoreVersion | Invoke action restoreVersion
[**SharesSharedDriveItemPermissionGrant**](SharesActionsApi.md#SharesSharedDriveItemPermissionGrant) | **Post** /shares/{sharedDriveItem-id}/permission/microsoft.graph.grant | Invoke action grant



## SharesSharedDriveItemListContentTypesAddCopy

> AnyOfmicrosoftGraphContentType SharesSharedDriveItemListContentTypesAddCopy(ctx, sharedDriveItemId).InlineObject731(inlineObject731).Execute()

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
    sharedDriveItemId := "sharedDriveItemId_example" // string | key: id of sharedDriveItem
    inlineObject731 := *openapiclient.NewInlineObject731() // InlineObject731 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SharesActionsApi.SharesSharedDriveItemListContentTypesAddCopy(context.Background(), sharedDriveItemId).InlineObject731(inlineObject731).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SharesActionsApi.SharesSharedDriveItemListContentTypesAddCopy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SharesSharedDriveItemListContentTypesAddCopy`: AnyOfmicrosoftGraphContentType
    fmt.Fprintf(os.Stdout, "Response from `SharesActionsApi.SharesSharedDriveItemListContentTypesAddCopy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedDriveItemId** | **string** | key: id of sharedDriveItem | 

### Other Parameters

Other parameters are passed through a pointer to a apiSharesSharedDriveItemListContentTypesAddCopyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject731** | [**InlineObject731**](InlineObject731.md) |  | 

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


## SharesSharedDriveItemListContentTypesContentTypeAssociateWithHubSites

> SharesSharedDriveItemListContentTypesContentTypeAssociateWithHubSites(ctx, sharedDriveItemId, contentTypeId).InlineObject729(inlineObject729).Execute()

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
    sharedDriveItemId := "sharedDriveItemId_example" // string | key: id of sharedDriveItem
    contentTypeId := "contentTypeId_example" // string | key: id of contentType
    inlineObject729 := *openapiclient.NewInlineObject729() // InlineObject729 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SharesActionsApi.SharesSharedDriveItemListContentTypesContentTypeAssociateWithHubSites(context.Background(), sharedDriveItemId, contentTypeId).InlineObject729(inlineObject729).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SharesActionsApi.SharesSharedDriveItemListContentTypesContentTypeAssociateWithHubSites``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedDriveItemId** | **string** | key: id of sharedDriveItem | 
**contentTypeId** | **string** | key: id of contentType | 

### Other Parameters

Other parameters are passed through a pointer to a apiSharesSharedDriveItemListContentTypesContentTypeAssociateWithHubSitesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **inlineObject729** | [**InlineObject729**](InlineObject729.md) |  | 

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


## SharesSharedDriveItemListContentTypesContentTypeBaseAssociateWithHubSites

> SharesSharedDriveItemListContentTypesContentTypeBaseAssociateWithHubSites(ctx, sharedDriveItemId, contentTypeId).InlineObject726(inlineObject726).Execute()

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
    sharedDriveItemId := "sharedDriveItemId_example" // string | key: id of sharedDriveItem
    contentTypeId := "contentTypeId_example" // string | key: id of contentType
    inlineObject726 := *openapiclient.NewInlineObject726() // InlineObject726 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SharesActionsApi.SharesSharedDriveItemListContentTypesContentTypeBaseAssociateWithHubSites(context.Background(), sharedDriveItemId, contentTypeId).InlineObject726(inlineObject726).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SharesActionsApi.SharesSharedDriveItemListContentTypesContentTypeBaseAssociateWithHubSites``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedDriveItemId** | **string** | key: id of sharedDriveItem | 
**contentTypeId** | **string** | key: id of contentType | 

### Other Parameters

Other parameters are passed through a pointer to a apiSharesSharedDriveItemListContentTypesContentTypeBaseAssociateWithHubSitesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **inlineObject726** | [**InlineObject726**](InlineObject726.md) |  | 

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


## SharesSharedDriveItemListContentTypesContentTypeBaseCopyToDefaultContentLocation

> SharesSharedDriveItemListContentTypesContentTypeBaseCopyToDefaultContentLocation(ctx, sharedDriveItemId, contentTypeId).InlineObject727(inlineObject727).Execute()

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
    sharedDriveItemId := "sharedDriveItemId_example" // string | key: id of sharedDriveItem
    contentTypeId := "contentTypeId_example" // string | key: id of contentType
    inlineObject727 := *openapiclient.NewInlineObject727() // InlineObject727 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SharesActionsApi.SharesSharedDriveItemListContentTypesContentTypeBaseCopyToDefaultContentLocation(context.Background(), sharedDriveItemId, contentTypeId).InlineObject727(inlineObject727).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SharesActionsApi.SharesSharedDriveItemListContentTypesContentTypeBaseCopyToDefaultContentLocation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedDriveItemId** | **string** | key: id of sharedDriveItem | 
**contentTypeId** | **string** | key: id of contentType | 

### Other Parameters

Other parameters are passed through a pointer to a apiSharesSharedDriveItemListContentTypesContentTypeBaseCopyToDefaultContentLocationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **inlineObject727** | [**InlineObject727**](InlineObject727.md) |  | 

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


## SharesSharedDriveItemListContentTypesContentTypeBasePublish

> SharesSharedDriveItemListContentTypesContentTypeBasePublish(ctx, sharedDriveItemId, contentTypeId).Execute()

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
    sharedDriveItemId := "sharedDriveItemId_example" // string | key: id of sharedDriveItem
    contentTypeId := "contentTypeId_example" // string | key: id of contentType

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SharesActionsApi.SharesSharedDriveItemListContentTypesContentTypeBasePublish(context.Background(), sharedDriveItemId, contentTypeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SharesActionsApi.SharesSharedDriveItemListContentTypesContentTypeBasePublish``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedDriveItemId** | **string** | key: id of sharedDriveItem | 
**contentTypeId** | **string** | key: id of contentType | 

### Other Parameters

Other parameters are passed through a pointer to a apiSharesSharedDriveItemListContentTypesContentTypeBasePublishRequest struct via the builder pattern


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


## SharesSharedDriveItemListContentTypesContentTypeBaseTypesAddCopy

> AnyOfmicrosoftGraphContentType SharesSharedDriveItemListContentTypesContentTypeBaseTypesAddCopy(ctx, sharedDriveItemId, contentTypeId).InlineObject728(inlineObject728).Execute()

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
    sharedDriveItemId := "sharedDriveItemId_example" // string | key: id of sharedDriveItem
    contentTypeId := "contentTypeId_example" // string | key: id of contentType
    inlineObject728 := *openapiclient.NewInlineObject728() // InlineObject728 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SharesActionsApi.SharesSharedDriveItemListContentTypesContentTypeBaseTypesAddCopy(context.Background(), sharedDriveItemId, contentTypeId).InlineObject728(inlineObject728).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SharesActionsApi.SharesSharedDriveItemListContentTypesContentTypeBaseTypesAddCopy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SharesSharedDriveItemListContentTypesContentTypeBaseTypesAddCopy`: AnyOfmicrosoftGraphContentType
    fmt.Fprintf(os.Stdout, "Response from `SharesActionsApi.SharesSharedDriveItemListContentTypesContentTypeBaseTypesAddCopy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedDriveItemId** | **string** | key: id of sharedDriveItem | 
**contentTypeId** | **string** | key: id of contentType | 

### Other Parameters

Other parameters are passed through a pointer to a apiSharesSharedDriveItemListContentTypesContentTypeBaseTypesAddCopyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **inlineObject728** | [**InlineObject728**](InlineObject728.md) |  | 

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


## SharesSharedDriveItemListContentTypesContentTypeBaseUnpublish

> SharesSharedDriveItemListContentTypesContentTypeBaseUnpublish(ctx, sharedDriveItemId, contentTypeId).Execute()

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
    sharedDriveItemId := "sharedDriveItemId_example" // string | key: id of sharedDriveItem
    contentTypeId := "contentTypeId_example" // string | key: id of contentType

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SharesActionsApi.SharesSharedDriveItemListContentTypesContentTypeBaseUnpublish(context.Background(), sharedDriveItemId, contentTypeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SharesActionsApi.SharesSharedDriveItemListContentTypesContentTypeBaseUnpublish``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedDriveItemId** | **string** | key: id of sharedDriveItem | 
**contentTypeId** | **string** | key: id of contentType | 

### Other Parameters

Other parameters are passed through a pointer to a apiSharesSharedDriveItemListContentTypesContentTypeBaseUnpublishRequest struct via the builder pattern


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


## SharesSharedDriveItemListContentTypesContentTypeCopyToDefaultContentLocation

> SharesSharedDriveItemListContentTypesContentTypeCopyToDefaultContentLocation(ctx, sharedDriveItemId, contentTypeId).InlineObject730(inlineObject730).Execute()

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
    sharedDriveItemId := "sharedDriveItemId_example" // string | key: id of sharedDriveItem
    contentTypeId := "contentTypeId_example" // string | key: id of contentType
    inlineObject730 := *openapiclient.NewInlineObject730() // InlineObject730 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SharesActionsApi.SharesSharedDriveItemListContentTypesContentTypeCopyToDefaultContentLocation(context.Background(), sharedDriveItemId, contentTypeId).InlineObject730(inlineObject730).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SharesActionsApi.SharesSharedDriveItemListContentTypesContentTypeCopyToDefaultContentLocation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedDriveItemId** | **string** | key: id of sharedDriveItem | 
**contentTypeId** | **string** | key: id of contentType | 

### Other Parameters

Other parameters are passed through a pointer to a apiSharesSharedDriveItemListContentTypesContentTypeCopyToDefaultContentLocationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **inlineObject730** | [**InlineObject730**](InlineObject730.md) |  | 

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


## SharesSharedDriveItemListContentTypesContentTypePublish

> SharesSharedDriveItemListContentTypesContentTypePublish(ctx, sharedDriveItemId, contentTypeId).Execute()

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
    sharedDriveItemId := "sharedDriveItemId_example" // string | key: id of sharedDriveItem
    contentTypeId := "contentTypeId_example" // string | key: id of contentType

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SharesActionsApi.SharesSharedDriveItemListContentTypesContentTypePublish(context.Background(), sharedDriveItemId, contentTypeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SharesActionsApi.SharesSharedDriveItemListContentTypesContentTypePublish``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedDriveItemId** | **string** | key: id of sharedDriveItem | 
**contentTypeId** | **string** | key: id of contentType | 

### Other Parameters

Other parameters are passed through a pointer to a apiSharesSharedDriveItemListContentTypesContentTypePublishRequest struct via the builder pattern


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


## SharesSharedDriveItemListContentTypesContentTypeUnpublish

> SharesSharedDriveItemListContentTypesContentTypeUnpublish(ctx, sharedDriveItemId, contentTypeId).Execute()

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
    sharedDriveItemId := "sharedDriveItemId_example" // string | key: id of sharedDriveItem
    contentTypeId := "contentTypeId_example" // string | key: id of contentType

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SharesActionsApi.SharesSharedDriveItemListContentTypesContentTypeUnpublish(context.Background(), sharedDriveItemId, contentTypeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SharesActionsApi.SharesSharedDriveItemListContentTypesContentTypeUnpublish``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedDriveItemId** | **string** | key: id of sharedDriveItem | 
**contentTypeId** | **string** | key: id of contentType | 

### Other Parameters

Other parameters are passed through a pointer to a apiSharesSharedDriveItemListContentTypesContentTypeUnpublishRequest struct via the builder pattern


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


## SharesSharedDriveItemListItemVersionsListItemVersionRestoreVersion

> SharesSharedDriveItemListItemVersionsListItemVersionRestoreVersion(ctx, sharedDriveItemId, listItemVersionId).Execute()

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
    sharedDriveItemId := "sharedDriveItemId_example" // string | key: id of sharedDriveItem
    listItemVersionId := "listItemVersionId_example" // string | key: id of listItemVersion

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SharesActionsApi.SharesSharedDriveItemListItemVersionsListItemVersionRestoreVersion(context.Background(), sharedDriveItemId, listItemVersionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SharesActionsApi.SharesSharedDriveItemListItemVersionsListItemVersionRestoreVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedDriveItemId** | **string** | key: id of sharedDriveItem | 
**listItemVersionId** | **string** | key: id of listItemVersion | 

### Other Parameters

Other parameters are passed through a pointer to a apiSharesSharedDriveItemListItemVersionsListItemVersionRestoreVersionRequest struct via the builder pattern


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


## SharesSharedDriveItemListItemsListItemVersionsListItemVersionRestoreVersion

> SharesSharedDriveItemListItemsListItemVersionsListItemVersionRestoreVersion(ctx, sharedDriveItemId, listItemId, listItemVersionId).Execute()

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
    sharedDriveItemId := "sharedDriveItemId_example" // string | key: id of sharedDriveItem
    listItemId := "listItemId_example" // string | key: id of listItem
    listItemVersionId := "listItemVersionId_example" // string | key: id of listItemVersion

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SharesActionsApi.SharesSharedDriveItemListItemsListItemVersionsListItemVersionRestoreVersion(context.Background(), sharedDriveItemId, listItemId, listItemVersionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SharesActionsApi.SharesSharedDriveItemListItemsListItemVersionsListItemVersionRestoreVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedDriveItemId** | **string** | key: id of sharedDriveItem | 
**listItemId** | **string** | key: id of listItem | 
**listItemVersionId** | **string** | key: id of listItemVersion | 

### Other Parameters

Other parameters are passed through a pointer to a apiSharesSharedDriveItemListItemsListItemVersionsListItemVersionRestoreVersionRequest struct via the builder pattern


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


## SharesSharedDriveItemPermissionGrant

> []*AnyOfmicrosoftGraphPermission SharesSharedDriveItemPermissionGrant(ctx, sharedDriveItemId).InlineObject732(inlineObject732).Execute()

Invoke action grant

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
    sharedDriveItemId := "sharedDriveItemId_example" // string | key: id of sharedDriveItem
    inlineObject732 := *openapiclient.NewInlineObject732() // InlineObject732 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SharesActionsApi.SharesSharedDriveItemPermissionGrant(context.Background(), sharedDriveItemId).InlineObject732(inlineObject732).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SharesActionsApi.SharesSharedDriveItemPermissionGrant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SharesSharedDriveItemPermissionGrant`: []*AnyOfmicrosoftGraphPermission
    fmt.Fprintf(os.Stdout, "Response from `SharesActionsApi.SharesSharedDriveItemPermissionGrant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedDriveItemId** | **string** | key: id of sharedDriveItem | 

### Other Parameters

Other parameters are passed through a pointer to a apiSharesSharedDriveItemPermissionGrantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject732** | [**InlineObject732**](InlineObject732.md) |  | 

### Return type

[**[]*AnyOfmicrosoftGraphPermission**](anyOf&lt;microsoft.graph.permission&gt;.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

