# \SharesListItemApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SharesDeleteListItem**](SharesListItemApi.md#SharesDeleteListItem) | **Delete** /shares/{sharedDriveItem-id}/listItem | Delete navigation property listItem for shares
[**SharesGetListItem**](SharesListItemApi.md#SharesGetListItem) | **Get** /shares/{sharedDriveItem-id}/listItem | Get listItem from shares
[**SharesListItemCreateVersions**](SharesListItemApi.md#SharesListItemCreateVersions) | **Post** /shares/{sharedDriveItem-id}/listItem/versions | Create new navigation property to versions for shares
[**SharesListItemDeleteDriveItem**](SharesListItemApi.md#SharesListItemDeleteDriveItem) | **Delete** /shares/{sharedDriveItem-id}/listItem/driveItem | Delete navigation property driveItem for shares
[**SharesListItemDeleteFields**](SharesListItemApi.md#SharesListItemDeleteFields) | **Delete** /shares/{sharedDriveItem-id}/listItem/fields | Delete navigation property fields for shares
[**SharesListItemDeleteRefAnalytics**](SharesListItemApi.md#SharesListItemDeleteRefAnalytics) | **Delete** /shares/{sharedDriveItem-id}/listItem/analytics/$ref | Delete ref of navigation property analytics for shares
[**SharesListItemDeleteVersions**](SharesListItemApi.md#SharesListItemDeleteVersions) | **Delete** /shares/{sharedDriveItem-id}/listItem/versions/{listItemVersion-id} | Delete navigation property versions for shares
[**SharesListItemGetAnalytics**](SharesListItemApi.md#SharesListItemGetAnalytics) | **Get** /shares/{sharedDriveItem-id}/listItem/analytics | Get analytics from shares
[**SharesListItemGetDriveItem**](SharesListItemApi.md#SharesListItemGetDriveItem) | **Get** /shares/{sharedDriveItem-id}/listItem/driveItem | Get driveItem from shares
[**SharesListItemGetDriveItemContent**](SharesListItemApi.md#SharesListItemGetDriveItemContent) | **Get** /shares/{sharedDriveItem-id}/listItem/driveItem/content | Get media content for the navigation property driveItem from shares
[**SharesListItemGetFields**](SharesListItemApi.md#SharesListItemGetFields) | **Get** /shares/{sharedDriveItem-id}/listItem/fields | Get fields from shares
[**SharesListItemGetRefAnalytics**](SharesListItemApi.md#SharesListItemGetRefAnalytics) | **Get** /shares/{sharedDriveItem-id}/listItem/analytics/$ref | Get ref of analytics from shares
[**SharesListItemGetVersions**](SharesListItemApi.md#SharesListItemGetVersions) | **Get** /shares/{sharedDriveItem-id}/listItem/versions/{listItemVersion-id} | Get versions from shares
[**SharesListItemListVersions**](SharesListItemApi.md#SharesListItemListVersions) | **Get** /shares/{sharedDriveItem-id}/listItem/versions | Get versions from shares
[**SharesListItemUpdateDriveItem**](SharesListItemApi.md#SharesListItemUpdateDriveItem) | **Patch** /shares/{sharedDriveItem-id}/listItem/driveItem | Update the navigation property driveItem in shares
[**SharesListItemUpdateDriveItemContent**](SharesListItemApi.md#SharesListItemUpdateDriveItemContent) | **Put** /shares/{sharedDriveItem-id}/listItem/driveItem/content | Update media content for the navigation property driveItem in shares
[**SharesListItemUpdateFields**](SharesListItemApi.md#SharesListItemUpdateFields) | **Patch** /shares/{sharedDriveItem-id}/listItem/fields | Update the navigation property fields in shares
[**SharesListItemUpdateRefAnalytics**](SharesListItemApi.md#SharesListItemUpdateRefAnalytics) | **Put** /shares/{sharedDriveItem-id}/listItem/analytics/$ref | Update the ref of navigation property analytics in shares
[**SharesListItemUpdateVersions**](SharesListItemApi.md#SharesListItemUpdateVersions) | **Patch** /shares/{sharedDriveItem-id}/listItem/versions/{listItemVersion-id} | Update the navigation property versions in shares
[**SharesListItemVersionsDeleteFields**](SharesListItemApi.md#SharesListItemVersionsDeleteFields) | **Delete** /shares/{sharedDriveItem-id}/listItem/versions/{listItemVersion-id}/fields | Delete navigation property fields for shares
[**SharesListItemVersionsGetFields**](SharesListItemApi.md#SharesListItemVersionsGetFields) | **Get** /shares/{sharedDriveItem-id}/listItem/versions/{listItemVersion-id}/fields | Get fields from shares
[**SharesListItemVersionsUpdateFields**](SharesListItemApi.md#SharesListItemVersionsUpdateFields) | **Patch** /shares/{sharedDriveItem-id}/listItem/versions/{listItemVersion-id}/fields | Update the navigation property fields in shares
[**SharesUpdateListItem**](SharesListItemApi.md#SharesUpdateListItem) | **Patch** /shares/{sharedDriveItem-id}/listItem | Update the navigation property listItem in shares



## SharesDeleteListItem

> SharesDeleteListItem(ctx, sharedDriveItemId).IfMatch(ifMatch).Execute()

Delete navigation property listItem for shares



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
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SharesListItemApi.SharesDeleteListItem(context.Background(), sharedDriveItemId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SharesListItemApi.SharesDeleteListItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedDriveItemId** | **string** | key: id of sharedDriveItem | 

### Other Parameters

Other parameters are passed through a pointer to a apiSharesDeleteListItemRequest struct via the builder pattern


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


## SharesGetListItem

> MicrosoftGraphListItem SharesGetListItem(ctx, sharedDriveItemId).Select_(select_).Expand(expand).Execute()

Get listItem from shares



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
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SharesListItemApi.SharesGetListItem(context.Background(), sharedDriveItemId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SharesListItemApi.SharesGetListItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SharesGetListItem`: MicrosoftGraphListItem
    fmt.Fprintf(os.Stdout, "Response from `SharesListItemApi.SharesGetListItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedDriveItemId** | **string** | key: id of sharedDriveItem | 

### Other Parameters

Other parameters are passed through a pointer to a apiSharesGetListItemRequest struct via the builder pattern


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


## SharesListItemCreateVersions

> MicrosoftGraphListItemVersion SharesListItemCreateVersions(ctx, sharedDriveItemId).MicrosoftGraphListItemVersion(microsoftGraphListItemVersion).Execute()

Create new navigation property to versions for shares



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
    microsoftGraphListItemVersion := *openapiclient.NewMicrosoftGraphListItemVersion() // MicrosoftGraphListItemVersion | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SharesListItemApi.SharesListItemCreateVersions(context.Background(), sharedDriveItemId).MicrosoftGraphListItemVersion(microsoftGraphListItemVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SharesListItemApi.SharesListItemCreateVersions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SharesListItemCreateVersions`: MicrosoftGraphListItemVersion
    fmt.Fprintf(os.Stdout, "Response from `SharesListItemApi.SharesListItemCreateVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedDriveItemId** | **string** | key: id of sharedDriveItem | 

### Other Parameters

Other parameters are passed through a pointer to a apiSharesListItemCreateVersionsRequest struct via the builder pattern


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


## SharesListItemDeleteDriveItem

> SharesListItemDeleteDriveItem(ctx, sharedDriveItemId).IfMatch(ifMatch).Execute()

Delete navigation property driveItem for shares



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
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SharesListItemApi.SharesListItemDeleteDriveItem(context.Background(), sharedDriveItemId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SharesListItemApi.SharesListItemDeleteDriveItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedDriveItemId** | **string** | key: id of sharedDriveItem | 

### Other Parameters

Other parameters are passed through a pointer to a apiSharesListItemDeleteDriveItemRequest struct via the builder pattern


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


## SharesListItemDeleteFields

> SharesListItemDeleteFields(ctx, sharedDriveItemId).IfMatch(ifMatch).Execute()

Delete navigation property fields for shares



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
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SharesListItemApi.SharesListItemDeleteFields(context.Background(), sharedDriveItemId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SharesListItemApi.SharesListItemDeleteFields``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedDriveItemId** | **string** | key: id of sharedDriveItem | 

### Other Parameters

Other parameters are passed through a pointer to a apiSharesListItemDeleteFieldsRequest struct via the builder pattern


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


## SharesListItemDeleteRefAnalytics

> SharesListItemDeleteRefAnalytics(ctx, sharedDriveItemId).IfMatch(ifMatch).Execute()

Delete ref of navigation property analytics for shares



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
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SharesListItemApi.SharesListItemDeleteRefAnalytics(context.Background(), sharedDriveItemId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SharesListItemApi.SharesListItemDeleteRefAnalytics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedDriveItemId** | **string** | key: id of sharedDriveItem | 

### Other Parameters

Other parameters are passed through a pointer to a apiSharesListItemDeleteRefAnalyticsRequest struct via the builder pattern


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


## SharesListItemDeleteVersions

> SharesListItemDeleteVersions(ctx, sharedDriveItemId, listItemVersionId).IfMatch(ifMatch).Execute()

Delete navigation property versions for shares



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
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SharesListItemApi.SharesListItemDeleteVersions(context.Background(), sharedDriveItemId, listItemVersionId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SharesListItemApi.SharesListItemDeleteVersions``: %v\n", err)
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

Other parameters are passed through a pointer to a apiSharesListItemDeleteVersionsRequest struct via the builder pattern


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


## SharesListItemGetAnalytics

> MicrosoftGraphItemAnalytics SharesListItemGetAnalytics(ctx, sharedDriveItemId).Select_(select_).Expand(expand).Execute()

Get analytics from shares



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
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SharesListItemApi.SharesListItemGetAnalytics(context.Background(), sharedDriveItemId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SharesListItemApi.SharesListItemGetAnalytics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SharesListItemGetAnalytics`: MicrosoftGraphItemAnalytics
    fmt.Fprintf(os.Stdout, "Response from `SharesListItemApi.SharesListItemGetAnalytics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedDriveItemId** | **string** | key: id of sharedDriveItem | 

### Other Parameters

Other parameters are passed through a pointer to a apiSharesListItemGetAnalyticsRequest struct via the builder pattern


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


## SharesListItemGetDriveItem

> MicrosoftGraphDriveItem SharesListItemGetDriveItem(ctx, sharedDriveItemId).Select_(select_).Expand(expand).Execute()

Get driveItem from shares



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
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SharesListItemApi.SharesListItemGetDriveItem(context.Background(), sharedDriveItemId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SharesListItemApi.SharesListItemGetDriveItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SharesListItemGetDriveItem`: MicrosoftGraphDriveItem
    fmt.Fprintf(os.Stdout, "Response from `SharesListItemApi.SharesListItemGetDriveItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedDriveItemId** | **string** | key: id of sharedDriveItem | 

### Other Parameters

Other parameters are passed through a pointer to a apiSharesListItemGetDriveItemRequest struct via the builder pattern


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


## SharesListItemGetDriveItemContent

> *os.File SharesListItemGetDriveItemContent(ctx, sharedDriveItemId).Execute()

Get media content for the navigation property driveItem from shares

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SharesListItemApi.SharesListItemGetDriveItemContent(context.Background(), sharedDriveItemId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SharesListItemApi.SharesListItemGetDriveItemContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SharesListItemGetDriveItemContent`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `SharesListItemApi.SharesListItemGetDriveItemContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedDriveItemId** | **string** | key: id of sharedDriveItem | 

### Other Parameters

Other parameters are passed through a pointer to a apiSharesListItemGetDriveItemContentRequest struct via the builder pattern


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


## SharesListItemGetFields

> MicrosoftGraphFieldValueSet SharesListItemGetFields(ctx, sharedDriveItemId).Select_(select_).Expand(expand).Execute()

Get fields from shares



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
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SharesListItemApi.SharesListItemGetFields(context.Background(), sharedDriveItemId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SharesListItemApi.SharesListItemGetFields``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SharesListItemGetFields`: MicrosoftGraphFieldValueSet
    fmt.Fprintf(os.Stdout, "Response from `SharesListItemApi.SharesListItemGetFields`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedDriveItemId** | **string** | key: id of sharedDriveItem | 

### Other Parameters

Other parameters are passed through a pointer to a apiSharesListItemGetFieldsRequest struct via the builder pattern


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


## SharesListItemGetRefAnalytics

> string SharesListItemGetRefAnalytics(ctx, sharedDriveItemId).Execute()

Get ref of analytics from shares



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SharesListItemApi.SharesListItemGetRefAnalytics(context.Background(), sharedDriveItemId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SharesListItemApi.SharesListItemGetRefAnalytics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SharesListItemGetRefAnalytics`: string
    fmt.Fprintf(os.Stdout, "Response from `SharesListItemApi.SharesListItemGetRefAnalytics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedDriveItemId** | **string** | key: id of sharedDriveItem | 

### Other Parameters

Other parameters are passed through a pointer to a apiSharesListItemGetRefAnalyticsRequest struct via the builder pattern


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


## SharesListItemGetVersions

> MicrosoftGraphListItemVersion SharesListItemGetVersions(ctx, sharedDriveItemId, listItemVersionId).Select_(select_).Expand(expand).Execute()

Get versions from shares



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
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SharesListItemApi.SharesListItemGetVersions(context.Background(), sharedDriveItemId, listItemVersionId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SharesListItemApi.SharesListItemGetVersions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SharesListItemGetVersions`: MicrosoftGraphListItemVersion
    fmt.Fprintf(os.Stdout, "Response from `SharesListItemApi.SharesListItemGetVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedDriveItemId** | **string** | key: id of sharedDriveItem | 
**listItemVersionId** | **string** | key: id of listItemVersion | 

### Other Parameters

Other parameters are passed through a pointer to a apiSharesListItemGetVersionsRequest struct via the builder pattern


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


## SharesListItemListVersions

> CollectionOfListItemVersion SharesListItemListVersions(ctx, sharedDriveItemId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get versions from shares



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
    resp, r, err := api_client.SharesListItemApi.SharesListItemListVersions(context.Background(), sharedDriveItemId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SharesListItemApi.SharesListItemListVersions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SharesListItemListVersions`: CollectionOfListItemVersion
    fmt.Fprintf(os.Stdout, "Response from `SharesListItemApi.SharesListItemListVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedDriveItemId** | **string** | key: id of sharedDriveItem | 

### Other Parameters

Other parameters are passed through a pointer to a apiSharesListItemListVersionsRequest struct via the builder pattern


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


## SharesListItemUpdateDriveItem

> SharesListItemUpdateDriveItem(ctx, sharedDriveItemId).MicrosoftGraphDriveItem(microsoftGraphDriveItem).Execute()

Update the navigation property driveItem in shares



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
    microsoftGraphDriveItem := *openapiclient.NewMicrosoftGraphDriveItem() // MicrosoftGraphDriveItem | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SharesListItemApi.SharesListItemUpdateDriveItem(context.Background(), sharedDriveItemId).MicrosoftGraphDriveItem(microsoftGraphDriveItem).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SharesListItemApi.SharesListItemUpdateDriveItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedDriveItemId** | **string** | key: id of sharedDriveItem | 

### Other Parameters

Other parameters are passed through a pointer to a apiSharesListItemUpdateDriveItemRequest struct via the builder pattern


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


## SharesListItemUpdateDriveItemContent

> SharesListItemUpdateDriveItemContent(ctx, sharedDriveItemId).Body(body).Execute()

Update media content for the navigation property driveItem in shares

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
    body := os.NewFile(1234, "some_file") // *os.File | New media content.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SharesListItemApi.SharesListItemUpdateDriveItemContent(context.Background(), sharedDriveItemId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SharesListItemApi.SharesListItemUpdateDriveItemContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedDriveItemId** | **string** | key: id of sharedDriveItem | 

### Other Parameters

Other parameters are passed through a pointer to a apiSharesListItemUpdateDriveItemContentRequest struct via the builder pattern


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


## SharesListItemUpdateFields

> SharesListItemUpdateFields(ctx, sharedDriveItemId).MicrosoftGraphFieldValueSet(microsoftGraphFieldValueSet).Execute()

Update the navigation property fields in shares



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
    microsoftGraphFieldValueSet := *openapiclient.NewMicrosoftGraphFieldValueSet() // MicrosoftGraphFieldValueSet | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SharesListItemApi.SharesListItemUpdateFields(context.Background(), sharedDriveItemId).MicrosoftGraphFieldValueSet(microsoftGraphFieldValueSet).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SharesListItemApi.SharesListItemUpdateFields``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedDriveItemId** | **string** | key: id of sharedDriveItem | 

### Other Parameters

Other parameters are passed through a pointer to a apiSharesListItemUpdateFieldsRequest struct via the builder pattern


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


## SharesListItemUpdateRefAnalytics

> SharesListItemUpdateRefAnalytics(ctx, sharedDriveItemId).RequestBody(requestBody).Execute()

Update the ref of navigation property analytics in shares



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
    requestBody := map[string]map[string]interface{}{"key": map[string]interface{}(123)} // map[string]map[string]interface{} | New navigation property ref values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SharesListItemApi.SharesListItemUpdateRefAnalytics(context.Background(), sharedDriveItemId).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SharesListItemApi.SharesListItemUpdateRefAnalytics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedDriveItemId** | **string** | key: id of sharedDriveItem | 

### Other Parameters

Other parameters are passed through a pointer to a apiSharesListItemUpdateRefAnalyticsRequest struct via the builder pattern


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


## SharesListItemUpdateVersions

> SharesListItemUpdateVersions(ctx, sharedDriveItemId, listItemVersionId).MicrosoftGraphListItemVersion(microsoftGraphListItemVersion).Execute()

Update the navigation property versions in shares



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
    microsoftGraphListItemVersion := *openapiclient.NewMicrosoftGraphListItemVersion() // MicrosoftGraphListItemVersion | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SharesListItemApi.SharesListItemUpdateVersions(context.Background(), sharedDriveItemId, listItemVersionId).MicrosoftGraphListItemVersion(microsoftGraphListItemVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SharesListItemApi.SharesListItemUpdateVersions``: %v\n", err)
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

Other parameters are passed through a pointer to a apiSharesListItemUpdateVersionsRequest struct via the builder pattern


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


## SharesListItemVersionsDeleteFields

> SharesListItemVersionsDeleteFields(ctx, sharedDriveItemId, listItemVersionId).IfMatch(ifMatch).Execute()

Delete navigation property fields for shares



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
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SharesListItemApi.SharesListItemVersionsDeleteFields(context.Background(), sharedDriveItemId, listItemVersionId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SharesListItemApi.SharesListItemVersionsDeleteFields``: %v\n", err)
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

Other parameters are passed through a pointer to a apiSharesListItemVersionsDeleteFieldsRequest struct via the builder pattern


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


## SharesListItemVersionsGetFields

> MicrosoftGraphFieldValueSet SharesListItemVersionsGetFields(ctx, sharedDriveItemId, listItemVersionId).Select_(select_).Expand(expand).Execute()

Get fields from shares



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
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SharesListItemApi.SharesListItemVersionsGetFields(context.Background(), sharedDriveItemId, listItemVersionId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SharesListItemApi.SharesListItemVersionsGetFields``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SharesListItemVersionsGetFields`: MicrosoftGraphFieldValueSet
    fmt.Fprintf(os.Stdout, "Response from `SharesListItemApi.SharesListItemVersionsGetFields`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedDriveItemId** | **string** | key: id of sharedDriveItem | 
**listItemVersionId** | **string** | key: id of listItemVersion | 

### Other Parameters

Other parameters are passed through a pointer to a apiSharesListItemVersionsGetFieldsRequest struct via the builder pattern


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


## SharesListItemVersionsUpdateFields

> SharesListItemVersionsUpdateFields(ctx, sharedDriveItemId, listItemVersionId).MicrosoftGraphFieldValueSet(microsoftGraphFieldValueSet).Execute()

Update the navigation property fields in shares



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
    microsoftGraphFieldValueSet := *openapiclient.NewMicrosoftGraphFieldValueSet() // MicrosoftGraphFieldValueSet | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SharesListItemApi.SharesListItemVersionsUpdateFields(context.Background(), sharedDriveItemId, listItemVersionId).MicrosoftGraphFieldValueSet(microsoftGraphFieldValueSet).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SharesListItemApi.SharesListItemVersionsUpdateFields``: %v\n", err)
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

Other parameters are passed through a pointer to a apiSharesListItemVersionsUpdateFieldsRequest struct via the builder pattern


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


## SharesUpdateListItem

> SharesUpdateListItem(ctx, sharedDriveItemId).MicrosoftGraphListItem(microsoftGraphListItem).Execute()

Update the navigation property listItem in shares



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
    microsoftGraphListItem := *openapiclient.NewMicrosoftGraphListItem() // MicrosoftGraphListItem | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SharesListItemApi.SharesUpdateListItem(context.Background(), sharedDriveItemId).MicrosoftGraphListItem(microsoftGraphListItem).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SharesListItemApi.SharesUpdateListItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedDriveItemId** | **string** | key: id of sharedDriveItem | 

### Other Parameters

Other parameters are passed through a pointer to a apiSharesUpdateListItemRequest struct via the builder pattern


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

