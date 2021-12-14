# \WorkbooksListItemApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WorkbooksDeleteListItem**](WorkbooksListItemApi.md#WorkbooksDeleteListItem) | **Delete** /workbooks/{driveItem-id}/listItem | Delete navigation property listItem for workbooks
[**WorkbooksGetListItem**](WorkbooksListItemApi.md#WorkbooksGetListItem) | **Get** /workbooks/{driveItem-id}/listItem | Get listItem from workbooks
[**WorkbooksListItemCreateVersions**](WorkbooksListItemApi.md#WorkbooksListItemCreateVersions) | **Post** /workbooks/{driveItem-id}/listItem/versions | Create new navigation property to versions for workbooks
[**WorkbooksListItemDeleteDriveItem**](WorkbooksListItemApi.md#WorkbooksListItemDeleteDriveItem) | **Delete** /workbooks/{driveItem-id}/listItem/driveItem | Delete navigation property driveItem for workbooks
[**WorkbooksListItemDeleteFields**](WorkbooksListItemApi.md#WorkbooksListItemDeleteFields) | **Delete** /workbooks/{driveItem-id}/listItem/fields | Delete navigation property fields for workbooks
[**WorkbooksListItemDeleteRefAnalytics**](WorkbooksListItemApi.md#WorkbooksListItemDeleteRefAnalytics) | **Delete** /workbooks/{driveItem-id}/listItem/analytics/$ref | Delete ref of navigation property analytics for workbooks
[**WorkbooksListItemDeleteVersions**](WorkbooksListItemApi.md#WorkbooksListItemDeleteVersions) | **Delete** /workbooks/{driveItem-id}/listItem/versions/{listItemVersion-id} | Delete navigation property versions for workbooks
[**WorkbooksListItemGetAnalytics**](WorkbooksListItemApi.md#WorkbooksListItemGetAnalytics) | **Get** /workbooks/{driveItem-id}/listItem/analytics | Get analytics from workbooks
[**WorkbooksListItemGetDriveItem**](WorkbooksListItemApi.md#WorkbooksListItemGetDriveItem) | **Get** /workbooks/{driveItem-id}/listItem/driveItem | Get driveItem from workbooks
[**WorkbooksListItemGetDriveItemContent**](WorkbooksListItemApi.md#WorkbooksListItemGetDriveItemContent) | **Get** /workbooks/{driveItem-id}/listItem/driveItem/content | Get media content for the navigation property driveItem from workbooks
[**WorkbooksListItemGetFields**](WorkbooksListItemApi.md#WorkbooksListItemGetFields) | **Get** /workbooks/{driveItem-id}/listItem/fields | Get fields from workbooks
[**WorkbooksListItemGetRefAnalytics**](WorkbooksListItemApi.md#WorkbooksListItemGetRefAnalytics) | **Get** /workbooks/{driveItem-id}/listItem/analytics/$ref | Get ref of analytics from workbooks
[**WorkbooksListItemGetVersions**](WorkbooksListItemApi.md#WorkbooksListItemGetVersions) | **Get** /workbooks/{driveItem-id}/listItem/versions/{listItemVersion-id} | Get versions from workbooks
[**WorkbooksListItemListVersions**](WorkbooksListItemApi.md#WorkbooksListItemListVersions) | **Get** /workbooks/{driveItem-id}/listItem/versions | Get versions from workbooks
[**WorkbooksListItemUpdateDriveItem**](WorkbooksListItemApi.md#WorkbooksListItemUpdateDriveItem) | **Patch** /workbooks/{driveItem-id}/listItem/driveItem | Update the navigation property driveItem in workbooks
[**WorkbooksListItemUpdateDriveItemContent**](WorkbooksListItemApi.md#WorkbooksListItemUpdateDriveItemContent) | **Put** /workbooks/{driveItem-id}/listItem/driveItem/content | Update media content for the navigation property driveItem in workbooks
[**WorkbooksListItemUpdateFields**](WorkbooksListItemApi.md#WorkbooksListItemUpdateFields) | **Patch** /workbooks/{driveItem-id}/listItem/fields | Update the navigation property fields in workbooks
[**WorkbooksListItemUpdateRefAnalytics**](WorkbooksListItemApi.md#WorkbooksListItemUpdateRefAnalytics) | **Put** /workbooks/{driveItem-id}/listItem/analytics/$ref | Update the ref of navigation property analytics in workbooks
[**WorkbooksListItemUpdateVersions**](WorkbooksListItemApi.md#WorkbooksListItemUpdateVersions) | **Patch** /workbooks/{driveItem-id}/listItem/versions/{listItemVersion-id} | Update the navigation property versions in workbooks
[**WorkbooksListItemVersionsDeleteFields**](WorkbooksListItemApi.md#WorkbooksListItemVersionsDeleteFields) | **Delete** /workbooks/{driveItem-id}/listItem/versions/{listItemVersion-id}/fields | Delete navigation property fields for workbooks
[**WorkbooksListItemVersionsGetFields**](WorkbooksListItemApi.md#WorkbooksListItemVersionsGetFields) | **Get** /workbooks/{driveItem-id}/listItem/versions/{listItemVersion-id}/fields | Get fields from workbooks
[**WorkbooksListItemVersionsUpdateFields**](WorkbooksListItemApi.md#WorkbooksListItemVersionsUpdateFields) | **Patch** /workbooks/{driveItem-id}/listItem/versions/{listItemVersion-id}/fields | Update the navigation property fields in workbooks
[**WorkbooksUpdateListItem**](WorkbooksListItemApi.md#WorkbooksUpdateListItem) | **Patch** /workbooks/{driveItem-id}/listItem | Update the navigation property listItem in workbooks



## WorkbooksDeleteListItem

> WorkbooksDeleteListItem(ctx, driveItemId).IfMatch(ifMatch).Execute()

Delete navigation property listItem for workbooks



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
    driveItemId := "driveItemId_example" // string | key: id of driveItem
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkbooksListItemApi.WorkbooksDeleteListItem(context.Background(), driveItemId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkbooksListItemApi.WorkbooksDeleteListItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string** | key: id of driveItem | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkbooksDeleteListItemRequest struct via the builder pattern


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


## WorkbooksGetListItem

> MicrosoftGraphListItem WorkbooksGetListItem(ctx, driveItemId).Select_(select_).Expand(expand).Execute()

Get listItem from workbooks



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
    driveItemId := "driveItemId_example" // string | key: id of driveItem
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkbooksListItemApi.WorkbooksGetListItem(context.Background(), driveItemId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkbooksListItemApi.WorkbooksGetListItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WorkbooksGetListItem`: MicrosoftGraphListItem
    fmt.Fprintf(os.Stdout, "Response from `WorkbooksListItemApi.WorkbooksGetListItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string** | key: id of driveItem | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkbooksGetListItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphListItem**](MicrosoftGraphListItem.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkbooksListItemCreateVersions

> MicrosoftGraphListItemVersion WorkbooksListItemCreateVersions(ctx, driveItemId).MicrosoftGraphListItemVersion(microsoftGraphListItemVersion).Execute()

Create new navigation property to versions for workbooks



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
    driveItemId := "driveItemId_example" // string | key: id of driveItem
    microsoftGraphListItemVersion := *openapiclient.NewMicrosoftGraphListItemVersion() // MicrosoftGraphListItemVersion | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkbooksListItemApi.WorkbooksListItemCreateVersions(context.Background(), driveItemId).MicrosoftGraphListItemVersion(microsoftGraphListItemVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkbooksListItemApi.WorkbooksListItemCreateVersions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WorkbooksListItemCreateVersions`: MicrosoftGraphListItemVersion
    fmt.Fprintf(os.Stdout, "Response from `WorkbooksListItemApi.WorkbooksListItemCreateVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string** | key: id of driveItem | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkbooksListItemCreateVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphListItemVersion** | [**MicrosoftGraphListItemVersion**](MicrosoftGraphListItemVersion.md) | New navigation property | 

### Return type

[**MicrosoftGraphListItemVersion**](MicrosoftGraphListItemVersion.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkbooksListItemDeleteDriveItem

> WorkbooksListItemDeleteDriveItem(ctx, driveItemId).IfMatch(ifMatch).Execute()

Delete navigation property driveItem for workbooks



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
    driveItemId := "driveItemId_example" // string | key: id of driveItem
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkbooksListItemApi.WorkbooksListItemDeleteDriveItem(context.Background(), driveItemId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkbooksListItemApi.WorkbooksListItemDeleteDriveItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string** | key: id of driveItem | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkbooksListItemDeleteDriveItemRequest struct via the builder pattern


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


## WorkbooksListItemDeleteFields

> WorkbooksListItemDeleteFields(ctx, driveItemId).IfMatch(ifMatch).Execute()

Delete navigation property fields for workbooks



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
    driveItemId := "driveItemId_example" // string | key: id of driveItem
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkbooksListItemApi.WorkbooksListItemDeleteFields(context.Background(), driveItemId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkbooksListItemApi.WorkbooksListItemDeleteFields``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string** | key: id of driveItem | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkbooksListItemDeleteFieldsRequest struct via the builder pattern


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


## WorkbooksListItemDeleteRefAnalytics

> WorkbooksListItemDeleteRefAnalytics(ctx, driveItemId).IfMatch(ifMatch).Execute()

Delete ref of navigation property analytics for workbooks



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
    driveItemId := "driveItemId_example" // string | key: id of driveItem
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkbooksListItemApi.WorkbooksListItemDeleteRefAnalytics(context.Background(), driveItemId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkbooksListItemApi.WorkbooksListItemDeleteRefAnalytics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string** | key: id of driveItem | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkbooksListItemDeleteRefAnalyticsRequest struct via the builder pattern


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


## WorkbooksListItemDeleteVersions

> WorkbooksListItemDeleteVersions(ctx, driveItemId, listItemVersionId).IfMatch(ifMatch).Execute()

Delete navigation property versions for workbooks



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
    driveItemId := "driveItemId_example" // string | key: id of driveItem
    listItemVersionId := "listItemVersionId_example" // string | key: id of listItemVersion
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkbooksListItemApi.WorkbooksListItemDeleteVersions(context.Background(), driveItemId, listItemVersionId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkbooksListItemApi.WorkbooksListItemDeleteVersions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string** | key: id of driveItem | 
**listItemVersionId** | **string** | key: id of listItemVersion | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkbooksListItemDeleteVersionsRequest struct via the builder pattern


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


## WorkbooksListItemGetAnalytics

> MicrosoftGraphItemAnalytics WorkbooksListItemGetAnalytics(ctx, driveItemId).Select_(select_).Expand(expand).Execute()

Get analytics from workbooks



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
    driveItemId := "driveItemId_example" // string | key: id of driveItem
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkbooksListItemApi.WorkbooksListItemGetAnalytics(context.Background(), driveItemId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkbooksListItemApi.WorkbooksListItemGetAnalytics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WorkbooksListItemGetAnalytics`: MicrosoftGraphItemAnalytics
    fmt.Fprintf(os.Stdout, "Response from `WorkbooksListItemApi.WorkbooksListItemGetAnalytics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string** | key: id of driveItem | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkbooksListItemGetAnalyticsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphItemAnalytics**](MicrosoftGraphItemAnalytics.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkbooksListItemGetDriveItem

> MicrosoftGraphDriveItem WorkbooksListItemGetDriveItem(ctx, driveItemId).Select_(select_).Expand(expand).Execute()

Get driveItem from workbooks



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
    driveItemId := "driveItemId_example" // string | key: id of driveItem
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkbooksListItemApi.WorkbooksListItemGetDriveItem(context.Background(), driveItemId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkbooksListItemApi.WorkbooksListItemGetDriveItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WorkbooksListItemGetDriveItem`: MicrosoftGraphDriveItem
    fmt.Fprintf(os.Stdout, "Response from `WorkbooksListItemApi.WorkbooksListItemGetDriveItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string** | key: id of driveItem | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkbooksListItemGetDriveItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphDriveItem**](MicrosoftGraphDriveItem.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkbooksListItemGetDriveItemContent

> *os.File WorkbooksListItemGetDriveItemContent(ctx, driveItemId).Execute()

Get media content for the navigation property driveItem from workbooks



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
    driveItemId := "driveItemId_example" // string | key: id of driveItem

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkbooksListItemApi.WorkbooksListItemGetDriveItemContent(context.Background(), driveItemId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkbooksListItemApi.WorkbooksListItemGetDriveItemContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WorkbooksListItemGetDriveItemContent`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `WorkbooksListItemApi.WorkbooksListItemGetDriveItemContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string** | key: id of driveItem | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkbooksListItemGetDriveItemContentRequest struct via the builder pattern


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


## WorkbooksListItemGetFields

> MicrosoftGraphFieldValueSet WorkbooksListItemGetFields(ctx, driveItemId).Select_(select_).Expand(expand).Execute()

Get fields from workbooks



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
    driveItemId := "driveItemId_example" // string | key: id of driveItem
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkbooksListItemApi.WorkbooksListItemGetFields(context.Background(), driveItemId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkbooksListItemApi.WorkbooksListItemGetFields``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WorkbooksListItemGetFields`: MicrosoftGraphFieldValueSet
    fmt.Fprintf(os.Stdout, "Response from `WorkbooksListItemApi.WorkbooksListItemGetFields`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string** | key: id of driveItem | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkbooksListItemGetFieldsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphFieldValueSet**](MicrosoftGraphFieldValueSet.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkbooksListItemGetRefAnalytics

> string WorkbooksListItemGetRefAnalytics(ctx, driveItemId).Execute()

Get ref of analytics from workbooks



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
    driveItemId := "driveItemId_example" // string | key: id of driveItem

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkbooksListItemApi.WorkbooksListItemGetRefAnalytics(context.Background(), driveItemId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkbooksListItemApi.WorkbooksListItemGetRefAnalytics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WorkbooksListItemGetRefAnalytics`: string
    fmt.Fprintf(os.Stdout, "Response from `WorkbooksListItemApi.WorkbooksListItemGetRefAnalytics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string** | key: id of driveItem | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkbooksListItemGetRefAnalyticsRequest struct via the builder pattern


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


## WorkbooksListItemGetVersions

> MicrosoftGraphListItemVersion WorkbooksListItemGetVersions(ctx, driveItemId, listItemVersionId).Select_(select_).Expand(expand).Execute()

Get versions from workbooks



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
    driveItemId := "driveItemId_example" // string | key: id of driveItem
    listItemVersionId := "listItemVersionId_example" // string | key: id of listItemVersion
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkbooksListItemApi.WorkbooksListItemGetVersions(context.Background(), driveItemId, listItemVersionId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkbooksListItemApi.WorkbooksListItemGetVersions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WorkbooksListItemGetVersions`: MicrosoftGraphListItemVersion
    fmt.Fprintf(os.Stdout, "Response from `WorkbooksListItemApi.WorkbooksListItemGetVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string** | key: id of driveItem | 
**listItemVersionId** | **string** | key: id of listItemVersion | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkbooksListItemGetVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphListItemVersion**](MicrosoftGraphListItemVersion.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkbooksListItemListVersions

> CollectionOfListItemVersion WorkbooksListItemListVersions(ctx, driveItemId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get versions from workbooks



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
    driveItemId := "driveItemId_example" // string | key: id of driveItem
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
    resp, r, err := api_client.WorkbooksListItemApi.WorkbooksListItemListVersions(context.Background(), driveItemId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkbooksListItemApi.WorkbooksListItemListVersions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WorkbooksListItemListVersions`: CollectionOfListItemVersion
    fmt.Fprintf(os.Stdout, "Response from `WorkbooksListItemApi.WorkbooksListItemListVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string** | key: id of driveItem | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkbooksListItemListVersionsRequest struct via the builder pattern


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

[**CollectionOfListItemVersion**](CollectionOfListItemVersion.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkbooksListItemUpdateDriveItem

> WorkbooksListItemUpdateDriveItem(ctx, driveItemId).MicrosoftGraphDriveItem(microsoftGraphDriveItem).Execute()

Update the navigation property driveItem in workbooks



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
    driveItemId := "driveItemId_example" // string | key: id of driveItem
    microsoftGraphDriveItem := *openapiclient.NewMicrosoftGraphDriveItem() // MicrosoftGraphDriveItem | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkbooksListItemApi.WorkbooksListItemUpdateDriveItem(context.Background(), driveItemId).MicrosoftGraphDriveItem(microsoftGraphDriveItem).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkbooksListItemApi.WorkbooksListItemUpdateDriveItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string** | key: id of driveItem | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkbooksListItemUpdateDriveItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphDriveItem** | [**MicrosoftGraphDriveItem**](MicrosoftGraphDriveItem.md) | New navigation property values | 

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


## WorkbooksListItemUpdateDriveItemContent

> WorkbooksListItemUpdateDriveItemContent(ctx, driveItemId).Body(body).Execute()

Update media content for the navigation property driveItem in workbooks



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
    driveItemId := "driveItemId_example" // string | key: id of driveItem
    body := os.NewFile(1234, "some_file") // *os.File | New media content.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkbooksListItemApi.WorkbooksListItemUpdateDriveItemContent(context.Background(), driveItemId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkbooksListItemApi.WorkbooksListItemUpdateDriveItemContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string** | key: id of driveItem | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkbooksListItemUpdateDriveItemContentRequest struct via the builder pattern


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


## WorkbooksListItemUpdateFields

> WorkbooksListItemUpdateFields(ctx, driveItemId).MicrosoftGraphFieldValueSet(microsoftGraphFieldValueSet).Execute()

Update the navigation property fields in workbooks



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
    driveItemId := "driveItemId_example" // string | key: id of driveItem
    microsoftGraphFieldValueSet := *openapiclient.NewMicrosoftGraphFieldValueSet() // MicrosoftGraphFieldValueSet | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkbooksListItemApi.WorkbooksListItemUpdateFields(context.Background(), driveItemId).MicrosoftGraphFieldValueSet(microsoftGraphFieldValueSet).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkbooksListItemApi.WorkbooksListItemUpdateFields``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string** | key: id of driveItem | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkbooksListItemUpdateFieldsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphFieldValueSet** | [**MicrosoftGraphFieldValueSet**](MicrosoftGraphFieldValueSet.md) | New navigation property values | 

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


## WorkbooksListItemUpdateRefAnalytics

> WorkbooksListItemUpdateRefAnalytics(ctx, driveItemId).RequestBody(requestBody).Execute()

Update the ref of navigation property analytics in workbooks



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
    driveItemId := "driveItemId_example" // string | key: id of driveItem
    requestBody := map[string]map[string]interface{}{"key": map[string]interface{}(123)} // map[string]map[string]interface{} | New navigation property ref values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkbooksListItemApi.WorkbooksListItemUpdateRefAnalytics(context.Background(), driveItemId).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkbooksListItemApi.WorkbooksListItemUpdateRefAnalytics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string** | key: id of driveItem | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkbooksListItemUpdateRefAnalyticsRequest struct via the builder pattern


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


## WorkbooksListItemUpdateVersions

> WorkbooksListItemUpdateVersions(ctx, driveItemId, listItemVersionId).MicrosoftGraphListItemVersion(microsoftGraphListItemVersion).Execute()

Update the navigation property versions in workbooks



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
    driveItemId := "driveItemId_example" // string | key: id of driveItem
    listItemVersionId := "listItemVersionId_example" // string | key: id of listItemVersion
    microsoftGraphListItemVersion := *openapiclient.NewMicrosoftGraphListItemVersion() // MicrosoftGraphListItemVersion | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkbooksListItemApi.WorkbooksListItemUpdateVersions(context.Background(), driveItemId, listItemVersionId).MicrosoftGraphListItemVersion(microsoftGraphListItemVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkbooksListItemApi.WorkbooksListItemUpdateVersions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string** | key: id of driveItem | 
**listItemVersionId** | **string** | key: id of listItemVersion | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkbooksListItemUpdateVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **microsoftGraphListItemVersion** | [**MicrosoftGraphListItemVersion**](MicrosoftGraphListItemVersion.md) | New navigation property values | 

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


## WorkbooksListItemVersionsDeleteFields

> WorkbooksListItemVersionsDeleteFields(ctx, driveItemId, listItemVersionId).IfMatch(ifMatch).Execute()

Delete navigation property fields for workbooks



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
    driveItemId := "driveItemId_example" // string | key: id of driveItem
    listItemVersionId := "listItemVersionId_example" // string | key: id of listItemVersion
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkbooksListItemApi.WorkbooksListItemVersionsDeleteFields(context.Background(), driveItemId, listItemVersionId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkbooksListItemApi.WorkbooksListItemVersionsDeleteFields``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string** | key: id of driveItem | 
**listItemVersionId** | **string** | key: id of listItemVersion | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkbooksListItemVersionsDeleteFieldsRequest struct via the builder pattern


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


## WorkbooksListItemVersionsGetFields

> MicrosoftGraphFieldValueSet WorkbooksListItemVersionsGetFields(ctx, driveItemId, listItemVersionId).Select_(select_).Expand(expand).Execute()

Get fields from workbooks



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
    driveItemId := "driveItemId_example" // string | key: id of driveItem
    listItemVersionId := "listItemVersionId_example" // string | key: id of listItemVersion
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkbooksListItemApi.WorkbooksListItemVersionsGetFields(context.Background(), driveItemId, listItemVersionId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkbooksListItemApi.WorkbooksListItemVersionsGetFields``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WorkbooksListItemVersionsGetFields`: MicrosoftGraphFieldValueSet
    fmt.Fprintf(os.Stdout, "Response from `WorkbooksListItemApi.WorkbooksListItemVersionsGetFields`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string** | key: id of driveItem | 
**listItemVersionId** | **string** | key: id of listItemVersion | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkbooksListItemVersionsGetFieldsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphFieldValueSet**](MicrosoftGraphFieldValueSet.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkbooksListItemVersionsUpdateFields

> WorkbooksListItemVersionsUpdateFields(ctx, driveItemId, listItemVersionId).MicrosoftGraphFieldValueSet(microsoftGraphFieldValueSet).Execute()

Update the navigation property fields in workbooks



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
    driveItemId := "driveItemId_example" // string | key: id of driveItem
    listItemVersionId := "listItemVersionId_example" // string | key: id of listItemVersion
    microsoftGraphFieldValueSet := *openapiclient.NewMicrosoftGraphFieldValueSet() // MicrosoftGraphFieldValueSet | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkbooksListItemApi.WorkbooksListItemVersionsUpdateFields(context.Background(), driveItemId, listItemVersionId).MicrosoftGraphFieldValueSet(microsoftGraphFieldValueSet).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkbooksListItemApi.WorkbooksListItemVersionsUpdateFields``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string** | key: id of driveItem | 
**listItemVersionId** | **string** | key: id of listItemVersion | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkbooksListItemVersionsUpdateFieldsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **microsoftGraphFieldValueSet** | [**MicrosoftGraphFieldValueSet**](MicrosoftGraphFieldValueSet.md) | New navigation property values | 

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


## WorkbooksUpdateListItem

> WorkbooksUpdateListItem(ctx, driveItemId).MicrosoftGraphListItem(microsoftGraphListItem).Execute()

Update the navigation property listItem in workbooks



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
    driveItemId := "driveItemId_example" // string | key: id of driveItem
    microsoftGraphListItem := *openapiclient.NewMicrosoftGraphListItem() // MicrosoftGraphListItem | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkbooksListItemApi.WorkbooksUpdateListItem(context.Background(), driveItemId).MicrosoftGraphListItem(microsoftGraphListItem).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkbooksListItemApi.WorkbooksUpdateListItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string** | key: id of driveItem | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkbooksUpdateListItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphListItem** | [**MicrosoftGraphListItem**](MicrosoftGraphListItem.md) | New navigation property values | 

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

