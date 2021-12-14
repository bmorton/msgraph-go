# \DrivesActionsApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DrivesDriveListContentTypesAddCopy**](DrivesActionsApi.md#DrivesDriveListContentTypesAddCopy) | **Post** /drives/{drive-id}/list/contentTypes/microsoft.graph.addCopy | Invoke action addCopy
[**DrivesDriveListContentTypesContentTypeAssociateWithHubSites**](DrivesActionsApi.md#DrivesDriveListContentTypesContentTypeAssociateWithHubSites) | **Post** /drives/{drive-id}/list/contentTypes/{contentType-id}/microsoft.graph.associateWithHubSites | Invoke action associateWithHubSites
[**DrivesDriveListContentTypesContentTypeBaseAssociateWithHubSites**](DrivesActionsApi.md#DrivesDriveListContentTypesContentTypeBaseAssociateWithHubSites) | **Post** /drives/{drive-id}/list/contentTypes/{contentType-id}/base/microsoft.graph.associateWithHubSites | Invoke action associateWithHubSites
[**DrivesDriveListContentTypesContentTypeBaseCopyToDefaultContentLocation**](DrivesActionsApi.md#DrivesDriveListContentTypesContentTypeBaseCopyToDefaultContentLocation) | **Post** /drives/{drive-id}/list/contentTypes/{contentType-id}/base/microsoft.graph.copyToDefaultContentLocation | Invoke action copyToDefaultContentLocation
[**DrivesDriveListContentTypesContentTypeBasePublish**](DrivesActionsApi.md#DrivesDriveListContentTypesContentTypeBasePublish) | **Post** /drives/{drive-id}/list/contentTypes/{contentType-id}/base/microsoft.graph.publish | Invoke action publish
[**DrivesDriveListContentTypesContentTypeBaseTypesAddCopy**](DrivesActionsApi.md#DrivesDriveListContentTypesContentTypeBaseTypesAddCopy) | **Post** /drives/{drive-id}/list/contentTypes/{contentType-id}/baseTypes/microsoft.graph.addCopy | Invoke action addCopy
[**DrivesDriveListContentTypesContentTypeBaseUnpublish**](DrivesActionsApi.md#DrivesDriveListContentTypesContentTypeBaseUnpublish) | **Post** /drives/{drive-id}/list/contentTypes/{contentType-id}/base/microsoft.graph.unpublish | Invoke action unpublish
[**DrivesDriveListContentTypesContentTypeCopyToDefaultContentLocation**](DrivesActionsApi.md#DrivesDriveListContentTypesContentTypeCopyToDefaultContentLocation) | **Post** /drives/{drive-id}/list/contentTypes/{contentType-id}/microsoft.graph.copyToDefaultContentLocation | Invoke action copyToDefaultContentLocation
[**DrivesDriveListContentTypesContentTypePublish**](DrivesActionsApi.md#DrivesDriveListContentTypesContentTypePublish) | **Post** /drives/{drive-id}/list/contentTypes/{contentType-id}/microsoft.graph.publish | Invoke action publish
[**DrivesDriveListContentTypesContentTypeUnpublish**](DrivesActionsApi.md#DrivesDriveListContentTypesContentTypeUnpublish) | **Post** /drives/{drive-id}/list/contentTypes/{contentType-id}/microsoft.graph.unpublish | Invoke action unpublish
[**DrivesDriveListItemsListItemVersionsListItemVersionRestoreVersion**](DrivesActionsApi.md#DrivesDriveListItemsListItemVersionsListItemVersionRestoreVersion) | **Post** /drives/{drive-id}/list/items/{listItem-id}/versions/{listItemVersion-id}/microsoft.graph.restoreVersion | Invoke action restoreVersion



## DrivesDriveListContentTypesAddCopy

> AnyOfmicrosoftGraphContentType DrivesDriveListContentTypesAddCopy(ctx, driveId).InlineObject146(inlineObject146).Execute()

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
    driveId := "driveId_example" // string | key: id of drive
    inlineObject146 := *openapiclient.NewInlineObject146() // InlineObject146 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DrivesActionsApi.DrivesDriveListContentTypesAddCopy(context.Background(), driveId).InlineObject146(inlineObject146).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DrivesActionsApi.DrivesDriveListContentTypesAddCopy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DrivesDriveListContentTypesAddCopy`: AnyOfmicrosoftGraphContentType
    fmt.Fprintf(os.Stdout, "Response from `DrivesActionsApi.DrivesDriveListContentTypesAddCopy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveId** | **string** | key: id of drive | 

### Other Parameters

Other parameters are passed through a pointer to a apiDrivesDriveListContentTypesAddCopyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject146** | [**InlineObject146**](InlineObject146.md) |  | 

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


## DrivesDriveListContentTypesContentTypeAssociateWithHubSites

> DrivesDriveListContentTypesContentTypeAssociateWithHubSites(ctx, driveId, contentTypeId).InlineObject144(inlineObject144).Execute()

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
    driveId := "driveId_example" // string | key: id of drive
    contentTypeId := "contentTypeId_example" // string | key: id of contentType
    inlineObject144 := *openapiclient.NewInlineObject144() // InlineObject144 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DrivesActionsApi.DrivesDriveListContentTypesContentTypeAssociateWithHubSites(context.Background(), driveId, contentTypeId).InlineObject144(inlineObject144).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DrivesActionsApi.DrivesDriveListContentTypesContentTypeAssociateWithHubSites``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveId** | **string** | key: id of drive | 
**contentTypeId** | **string** | key: id of contentType | 

### Other Parameters

Other parameters are passed through a pointer to a apiDrivesDriveListContentTypesContentTypeAssociateWithHubSitesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **inlineObject144** | [**InlineObject144**](InlineObject144.md) |  | 

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


## DrivesDriveListContentTypesContentTypeBaseAssociateWithHubSites

> DrivesDriveListContentTypesContentTypeBaseAssociateWithHubSites(ctx, driveId, contentTypeId).InlineObject141(inlineObject141).Execute()

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
    driveId := "driveId_example" // string | key: id of drive
    contentTypeId := "contentTypeId_example" // string | key: id of contentType
    inlineObject141 := *openapiclient.NewInlineObject141() // InlineObject141 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DrivesActionsApi.DrivesDriveListContentTypesContentTypeBaseAssociateWithHubSites(context.Background(), driveId, contentTypeId).InlineObject141(inlineObject141).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DrivesActionsApi.DrivesDriveListContentTypesContentTypeBaseAssociateWithHubSites``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveId** | **string** | key: id of drive | 
**contentTypeId** | **string** | key: id of contentType | 

### Other Parameters

Other parameters are passed through a pointer to a apiDrivesDriveListContentTypesContentTypeBaseAssociateWithHubSitesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **inlineObject141** | [**InlineObject141**](InlineObject141.md) |  | 

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


## DrivesDriveListContentTypesContentTypeBaseCopyToDefaultContentLocation

> DrivesDriveListContentTypesContentTypeBaseCopyToDefaultContentLocation(ctx, driveId, contentTypeId).InlineObject142(inlineObject142).Execute()

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
    driveId := "driveId_example" // string | key: id of drive
    contentTypeId := "contentTypeId_example" // string | key: id of contentType
    inlineObject142 := *openapiclient.NewInlineObject142() // InlineObject142 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DrivesActionsApi.DrivesDriveListContentTypesContentTypeBaseCopyToDefaultContentLocation(context.Background(), driveId, contentTypeId).InlineObject142(inlineObject142).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DrivesActionsApi.DrivesDriveListContentTypesContentTypeBaseCopyToDefaultContentLocation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveId** | **string** | key: id of drive | 
**contentTypeId** | **string** | key: id of contentType | 

### Other Parameters

Other parameters are passed through a pointer to a apiDrivesDriveListContentTypesContentTypeBaseCopyToDefaultContentLocationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **inlineObject142** | [**InlineObject142**](InlineObject142.md) |  | 

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


## DrivesDriveListContentTypesContentTypeBasePublish

> DrivesDriveListContentTypesContentTypeBasePublish(ctx, driveId, contentTypeId).Execute()

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
    driveId := "driveId_example" // string | key: id of drive
    contentTypeId := "contentTypeId_example" // string | key: id of contentType

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DrivesActionsApi.DrivesDriveListContentTypesContentTypeBasePublish(context.Background(), driveId, contentTypeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DrivesActionsApi.DrivesDriveListContentTypesContentTypeBasePublish``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveId** | **string** | key: id of drive | 
**contentTypeId** | **string** | key: id of contentType | 

### Other Parameters

Other parameters are passed through a pointer to a apiDrivesDriveListContentTypesContentTypeBasePublishRequest struct via the builder pattern


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


## DrivesDriveListContentTypesContentTypeBaseTypesAddCopy

> AnyOfmicrosoftGraphContentType DrivesDriveListContentTypesContentTypeBaseTypesAddCopy(ctx, driveId, contentTypeId).InlineObject143(inlineObject143).Execute()

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
    driveId := "driveId_example" // string | key: id of drive
    contentTypeId := "contentTypeId_example" // string | key: id of contentType
    inlineObject143 := *openapiclient.NewInlineObject143() // InlineObject143 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DrivesActionsApi.DrivesDriveListContentTypesContentTypeBaseTypesAddCopy(context.Background(), driveId, contentTypeId).InlineObject143(inlineObject143).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DrivesActionsApi.DrivesDriveListContentTypesContentTypeBaseTypesAddCopy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DrivesDriveListContentTypesContentTypeBaseTypesAddCopy`: AnyOfmicrosoftGraphContentType
    fmt.Fprintf(os.Stdout, "Response from `DrivesActionsApi.DrivesDriveListContentTypesContentTypeBaseTypesAddCopy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveId** | **string** | key: id of drive | 
**contentTypeId** | **string** | key: id of contentType | 

### Other Parameters

Other parameters are passed through a pointer to a apiDrivesDriveListContentTypesContentTypeBaseTypesAddCopyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **inlineObject143** | [**InlineObject143**](InlineObject143.md) |  | 

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


## DrivesDriveListContentTypesContentTypeBaseUnpublish

> DrivesDriveListContentTypesContentTypeBaseUnpublish(ctx, driveId, contentTypeId).Execute()

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
    driveId := "driveId_example" // string | key: id of drive
    contentTypeId := "contentTypeId_example" // string | key: id of contentType

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DrivesActionsApi.DrivesDriveListContentTypesContentTypeBaseUnpublish(context.Background(), driveId, contentTypeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DrivesActionsApi.DrivesDriveListContentTypesContentTypeBaseUnpublish``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveId** | **string** | key: id of drive | 
**contentTypeId** | **string** | key: id of contentType | 

### Other Parameters

Other parameters are passed through a pointer to a apiDrivesDriveListContentTypesContentTypeBaseUnpublishRequest struct via the builder pattern


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


## DrivesDriveListContentTypesContentTypeCopyToDefaultContentLocation

> DrivesDriveListContentTypesContentTypeCopyToDefaultContentLocation(ctx, driveId, contentTypeId).InlineObject145(inlineObject145).Execute()

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
    driveId := "driveId_example" // string | key: id of drive
    contentTypeId := "contentTypeId_example" // string | key: id of contentType
    inlineObject145 := *openapiclient.NewInlineObject145() // InlineObject145 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DrivesActionsApi.DrivesDriveListContentTypesContentTypeCopyToDefaultContentLocation(context.Background(), driveId, contentTypeId).InlineObject145(inlineObject145).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DrivesActionsApi.DrivesDriveListContentTypesContentTypeCopyToDefaultContentLocation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveId** | **string** | key: id of drive | 
**contentTypeId** | **string** | key: id of contentType | 

### Other Parameters

Other parameters are passed through a pointer to a apiDrivesDriveListContentTypesContentTypeCopyToDefaultContentLocationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **inlineObject145** | [**InlineObject145**](InlineObject145.md) |  | 

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


## DrivesDriveListContentTypesContentTypePublish

> DrivesDriveListContentTypesContentTypePublish(ctx, driveId, contentTypeId).Execute()

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
    driveId := "driveId_example" // string | key: id of drive
    contentTypeId := "contentTypeId_example" // string | key: id of contentType

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DrivesActionsApi.DrivesDriveListContentTypesContentTypePublish(context.Background(), driveId, contentTypeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DrivesActionsApi.DrivesDriveListContentTypesContentTypePublish``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveId** | **string** | key: id of drive | 
**contentTypeId** | **string** | key: id of contentType | 

### Other Parameters

Other parameters are passed through a pointer to a apiDrivesDriveListContentTypesContentTypePublishRequest struct via the builder pattern


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


## DrivesDriveListContentTypesContentTypeUnpublish

> DrivesDriveListContentTypesContentTypeUnpublish(ctx, driveId, contentTypeId).Execute()

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
    driveId := "driveId_example" // string | key: id of drive
    contentTypeId := "contentTypeId_example" // string | key: id of contentType

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DrivesActionsApi.DrivesDriveListContentTypesContentTypeUnpublish(context.Background(), driveId, contentTypeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DrivesActionsApi.DrivesDriveListContentTypesContentTypeUnpublish``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveId** | **string** | key: id of drive | 
**contentTypeId** | **string** | key: id of contentType | 

### Other Parameters

Other parameters are passed through a pointer to a apiDrivesDriveListContentTypesContentTypeUnpublishRequest struct via the builder pattern


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


## DrivesDriveListItemsListItemVersionsListItemVersionRestoreVersion

> DrivesDriveListItemsListItemVersionsListItemVersionRestoreVersion(ctx, driveId, listItemId, listItemVersionId).Execute()

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
    driveId := "driveId_example" // string | key: id of drive
    listItemId := "listItemId_example" // string | key: id of listItem
    listItemVersionId := "listItemVersionId_example" // string | key: id of listItemVersion

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DrivesActionsApi.DrivesDriveListItemsListItemVersionsListItemVersionRestoreVersion(context.Background(), driveId, listItemId, listItemVersionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DrivesActionsApi.DrivesDriveListItemsListItemVersionsListItemVersionRestoreVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveId** | **string** | key: id of drive | 
**listItemId** | **string** | key: id of listItem | 
**listItemVersionId** | **string** | key: id of listItemVersion | 

### Other Parameters

Other parameters are passed through a pointer to a apiDrivesDriveListItemsListItemVersionsListItemVersionRestoreVersionRequest struct via the builder pattern


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

