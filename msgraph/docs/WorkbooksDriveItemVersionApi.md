# \WorkbooksDriveItemVersionApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WorkbooksCreateVersions**](WorkbooksDriveItemVersionApi.md#WorkbooksCreateVersions) | **Post** /workbooks/{driveItem-id}/versions | Create new navigation property to versions for workbooks
[**WorkbooksDeleteVersions**](WorkbooksDriveItemVersionApi.md#WorkbooksDeleteVersions) | **Delete** /workbooks/{driveItem-id}/versions/{driveItemVersion-id} | Delete navigation property versions for workbooks
[**WorkbooksGetVersions**](WorkbooksDriveItemVersionApi.md#WorkbooksGetVersions) | **Get** /workbooks/{driveItem-id}/versions/{driveItemVersion-id} | Get versions from workbooks
[**WorkbooksGetVersionsContent**](WorkbooksDriveItemVersionApi.md#WorkbooksGetVersionsContent) | **Get** /workbooks/{driveItem-id}/versions/{driveItemVersion-id}/content | Get media content for the navigation property versions from workbooks
[**WorkbooksListVersions**](WorkbooksDriveItemVersionApi.md#WorkbooksListVersions) | **Get** /workbooks/{driveItem-id}/versions | Get versions from workbooks
[**WorkbooksUpdateVersions**](WorkbooksDriveItemVersionApi.md#WorkbooksUpdateVersions) | **Patch** /workbooks/{driveItem-id}/versions/{driveItemVersion-id} | Update the navigation property versions in workbooks
[**WorkbooksUpdateVersionsContent**](WorkbooksDriveItemVersionApi.md#WorkbooksUpdateVersionsContent) | **Put** /workbooks/{driveItem-id}/versions/{driveItemVersion-id}/content | Update media content for the navigation property versions in workbooks



## WorkbooksCreateVersions

> MicrosoftGraphDriveItemVersion WorkbooksCreateVersions(ctx, driveItemId).MicrosoftGraphDriveItemVersion(microsoftGraphDriveItemVersion).Execute()

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
    microsoftGraphDriveItemVersion := *openapiclient.NewMicrosoftGraphDriveItemVersion() // MicrosoftGraphDriveItemVersion | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkbooksDriveItemVersionApi.WorkbooksCreateVersions(context.Background(), driveItemId).MicrosoftGraphDriveItemVersion(microsoftGraphDriveItemVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkbooksDriveItemVersionApi.WorkbooksCreateVersions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WorkbooksCreateVersions`: MicrosoftGraphDriveItemVersion
    fmt.Fprintf(os.Stdout, "Response from `WorkbooksDriveItemVersionApi.WorkbooksCreateVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string** | key: id of driveItem | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkbooksCreateVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphDriveItemVersion** | [**MicrosoftGraphDriveItemVersion**](MicrosoftGraphDriveItemVersion.md) | New navigation property | 

### Return type

[**MicrosoftGraphDriveItemVersion**](MicrosoftGraphDriveItemVersion.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkbooksDeleteVersions

> WorkbooksDeleteVersions(ctx, driveItemId, driveItemVersionId).IfMatch(ifMatch).Execute()

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
    driveItemVersionId := "driveItemVersionId_example" // string | key: id of driveItemVersion
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkbooksDriveItemVersionApi.WorkbooksDeleteVersions(context.Background(), driveItemId, driveItemVersionId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkbooksDriveItemVersionApi.WorkbooksDeleteVersions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string** | key: id of driveItem | 
**driveItemVersionId** | **string** | key: id of driveItemVersion | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkbooksDeleteVersionsRequest struct via the builder pattern


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


## WorkbooksGetVersions

> MicrosoftGraphDriveItemVersion WorkbooksGetVersions(ctx, driveItemId, driveItemVersionId).Select_(select_).Expand(expand).Execute()

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
    driveItemVersionId := "driveItemVersionId_example" // string | key: id of driveItemVersion
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkbooksDriveItemVersionApi.WorkbooksGetVersions(context.Background(), driveItemId, driveItemVersionId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkbooksDriveItemVersionApi.WorkbooksGetVersions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WorkbooksGetVersions`: MicrosoftGraphDriveItemVersion
    fmt.Fprintf(os.Stdout, "Response from `WorkbooksDriveItemVersionApi.WorkbooksGetVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string** | key: id of driveItem | 
**driveItemVersionId** | **string** | key: id of driveItemVersion | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkbooksGetVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphDriveItemVersion**](MicrosoftGraphDriveItemVersion.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkbooksGetVersionsContent

> *os.File WorkbooksGetVersionsContent(ctx, driveItemId, driveItemVersionId).Execute()

Get media content for the navigation property versions from workbooks



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
    driveItemVersionId := "driveItemVersionId_example" // string | key: id of driveItemVersion

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkbooksDriveItemVersionApi.WorkbooksGetVersionsContent(context.Background(), driveItemId, driveItemVersionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkbooksDriveItemVersionApi.WorkbooksGetVersionsContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WorkbooksGetVersionsContent`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `WorkbooksDriveItemVersionApi.WorkbooksGetVersionsContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string** | key: id of driveItem | 
**driveItemVersionId** | **string** | key: id of driveItemVersion | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkbooksGetVersionsContentRequest struct via the builder pattern


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


## WorkbooksListVersions

> CollectionOfDriveItemVersion WorkbooksListVersions(ctx, driveItemId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

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
    resp, r, err := api_client.WorkbooksDriveItemVersionApi.WorkbooksListVersions(context.Background(), driveItemId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkbooksDriveItemVersionApi.WorkbooksListVersions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WorkbooksListVersions`: CollectionOfDriveItemVersion
    fmt.Fprintf(os.Stdout, "Response from `WorkbooksDriveItemVersionApi.WorkbooksListVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string** | key: id of driveItem | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkbooksListVersionsRequest struct via the builder pattern


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

[**CollectionOfDriveItemVersion**](CollectionOfDriveItemVersion.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkbooksUpdateVersions

> WorkbooksUpdateVersions(ctx, driveItemId, driveItemVersionId).MicrosoftGraphDriveItemVersion(microsoftGraphDriveItemVersion).Execute()

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
    driveItemVersionId := "driveItemVersionId_example" // string | key: id of driveItemVersion
    microsoftGraphDriveItemVersion := *openapiclient.NewMicrosoftGraphDriveItemVersion() // MicrosoftGraphDriveItemVersion | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkbooksDriveItemVersionApi.WorkbooksUpdateVersions(context.Background(), driveItemId, driveItemVersionId).MicrosoftGraphDriveItemVersion(microsoftGraphDriveItemVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkbooksDriveItemVersionApi.WorkbooksUpdateVersions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string** | key: id of driveItem | 
**driveItemVersionId** | **string** | key: id of driveItemVersion | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkbooksUpdateVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **microsoftGraphDriveItemVersion** | [**MicrosoftGraphDriveItemVersion**](MicrosoftGraphDriveItemVersion.md) | New navigation property values | 

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


## WorkbooksUpdateVersionsContent

> WorkbooksUpdateVersionsContent(ctx, driveItemId, driveItemVersionId).Body(body).Execute()

Update media content for the navigation property versions in workbooks



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
    driveItemVersionId := "driveItemVersionId_example" // string | key: id of driveItemVersion
    body := os.NewFile(1234, "some_file") // *os.File | New media content.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkbooksDriveItemVersionApi.WorkbooksUpdateVersionsContent(context.Background(), driveItemId, driveItemVersionId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkbooksDriveItemVersionApi.WorkbooksUpdateVersionsContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string** | key: id of driveItem | 
**driveItemVersionId** | **string** | key: id of driveItemVersion | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkbooksUpdateVersionsContentRequest struct via the builder pattern


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

