# \DriveDriveItemApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DriveCreateBundles**](DriveDriveItemApi.md#DriveCreateBundles) | **Post** /drive/bundles | Create new navigation property to bundles for drive
[**DriveCreateFollowing**](DriveDriveItemApi.md#DriveCreateFollowing) | **Post** /drive/following | Create new navigation property to following for drive
[**DriveCreateItems**](DriveDriveItemApi.md#DriveCreateItems) | **Post** /drive/items | Create new navigation property to items for drive
[**DriveCreateSpecial**](DriveDriveItemApi.md#DriveCreateSpecial) | **Post** /drive/special | Create new navigation property to special for drive
[**DriveDeleteBundles**](DriveDriveItemApi.md#DriveDeleteBundles) | **Delete** /drive/bundles/{driveItem-id} | Delete navigation property bundles for drive
[**DriveDeleteFollowing**](DriveDriveItemApi.md#DriveDeleteFollowing) | **Delete** /drive/following/{driveItem-id} | Delete navigation property following for drive
[**DriveDeleteItems**](DriveDriveItemApi.md#DriveDeleteItems) | **Delete** /drive/items/{driveItem-id} | Delete navigation property items for drive
[**DriveDeleteRoot**](DriveDriveItemApi.md#DriveDeleteRoot) | **Delete** /drive/root | Delete navigation property root for drive
[**DriveDeleteSpecial**](DriveDriveItemApi.md#DriveDeleteSpecial) | **Delete** /drive/special/{driveItem-id} | Delete navigation property special for drive
[**DriveGetBundles**](DriveDriveItemApi.md#DriveGetBundles) | **Get** /drive/bundles/{driveItem-id} | Get bundles from drive
[**DriveGetBundlesContent**](DriveDriveItemApi.md#DriveGetBundlesContent) | **Get** /drive/bundles/{driveItem-id}/content | Get media content for the navigation property bundles from drive
[**DriveGetFollowing**](DriveDriveItemApi.md#DriveGetFollowing) | **Get** /drive/following/{driveItem-id} | Get following from drive
[**DriveGetFollowingContent**](DriveDriveItemApi.md#DriveGetFollowingContent) | **Get** /drive/following/{driveItem-id}/content | Get media content for the navigation property following from drive
[**DriveGetItems**](DriveDriveItemApi.md#DriveGetItems) | **Get** /drive/items/{driveItem-id} | Get items from drive
[**DriveGetItemsContent**](DriveDriveItemApi.md#DriveGetItemsContent) | **Get** /drive/items/{driveItem-id}/content | Get media content for the navigation property items from drive
[**DriveGetRoot**](DriveDriveItemApi.md#DriveGetRoot) | **Get** /drive/root | Get root from drive
[**DriveGetRootContent**](DriveDriveItemApi.md#DriveGetRootContent) | **Get** /drive/root/content | Get media content for the navigation property root from drive
[**DriveGetSpecial**](DriveDriveItemApi.md#DriveGetSpecial) | **Get** /drive/special/{driveItem-id} | Get special from drive
[**DriveGetSpecialContent**](DriveDriveItemApi.md#DriveGetSpecialContent) | **Get** /drive/special/{driveItem-id}/content | Get media content for the navigation property special from drive
[**DriveListBundles**](DriveDriveItemApi.md#DriveListBundles) | **Get** /drive/bundles | Get bundles from drive
[**DriveListFollowing**](DriveDriveItemApi.md#DriveListFollowing) | **Get** /drive/following | Get following from drive
[**DriveListItems**](DriveDriveItemApi.md#DriveListItems) | **Get** /drive/items | Get items from drive
[**DriveListSpecial**](DriveDriveItemApi.md#DriveListSpecial) | **Get** /drive/special | Get special from drive
[**DriveUpdateBundles**](DriveDriveItemApi.md#DriveUpdateBundles) | **Patch** /drive/bundles/{driveItem-id} | Update the navigation property bundles in drive
[**DriveUpdateBundlesContent**](DriveDriveItemApi.md#DriveUpdateBundlesContent) | **Put** /drive/bundles/{driveItem-id}/content | Update media content for the navigation property bundles in drive
[**DriveUpdateFollowing**](DriveDriveItemApi.md#DriveUpdateFollowing) | **Patch** /drive/following/{driveItem-id} | Update the navigation property following in drive
[**DriveUpdateFollowingContent**](DriveDriveItemApi.md#DriveUpdateFollowingContent) | **Put** /drive/following/{driveItem-id}/content | Update media content for the navigation property following in drive
[**DriveUpdateItems**](DriveDriveItemApi.md#DriveUpdateItems) | **Patch** /drive/items/{driveItem-id} | Update the navigation property items in drive
[**DriveUpdateItemsContent**](DriveDriveItemApi.md#DriveUpdateItemsContent) | **Put** /drive/items/{driveItem-id}/content | Update media content for the navigation property items in drive
[**DriveUpdateRoot**](DriveDriveItemApi.md#DriveUpdateRoot) | **Patch** /drive/root | Update the navigation property root in drive
[**DriveUpdateRootContent**](DriveDriveItemApi.md#DriveUpdateRootContent) | **Put** /drive/root/content | Update media content for the navigation property root in drive
[**DriveUpdateSpecial**](DriveDriveItemApi.md#DriveUpdateSpecial) | **Patch** /drive/special/{driveItem-id} | Update the navigation property special in drive
[**DriveUpdateSpecialContent**](DriveDriveItemApi.md#DriveUpdateSpecialContent) | **Put** /drive/special/{driveItem-id}/content | Update media content for the navigation property special in drive



## DriveCreateBundles

> MicrosoftGraphDriveItem DriveCreateBundles(ctx).MicrosoftGraphDriveItem(microsoftGraphDriveItem).Execute()

Create new navigation property to bundles for drive



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
    microsoftGraphDriveItem := *openapiclient.NewMicrosoftGraphDriveItem() // MicrosoftGraphDriveItem | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DriveDriveItemApi.DriveCreateBundles(context.Background()).MicrosoftGraphDriveItem(microsoftGraphDriveItem).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DriveDriveItemApi.DriveCreateBundles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DriveCreateBundles`: MicrosoftGraphDriveItem
    fmt.Fprintf(os.Stdout, "Response from `DriveDriveItemApi.DriveCreateBundles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDriveCreateBundlesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphDriveItem** | [**MicrosoftGraphDriveItem**](MicrosoftGraphDriveItem.md) | New navigation property | 

### Return type

[**MicrosoftGraphDriveItem**](MicrosoftGraphDriveItem.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DriveCreateFollowing

> MicrosoftGraphDriveItem DriveCreateFollowing(ctx).MicrosoftGraphDriveItem(microsoftGraphDriveItem).Execute()

Create new navigation property to following for drive



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
    microsoftGraphDriveItem := *openapiclient.NewMicrosoftGraphDriveItem() // MicrosoftGraphDriveItem | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DriveDriveItemApi.DriveCreateFollowing(context.Background()).MicrosoftGraphDriveItem(microsoftGraphDriveItem).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DriveDriveItemApi.DriveCreateFollowing``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DriveCreateFollowing`: MicrosoftGraphDriveItem
    fmt.Fprintf(os.Stdout, "Response from `DriveDriveItemApi.DriveCreateFollowing`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDriveCreateFollowingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphDriveItem** | [**MicrosoftGraphDriveItem**](MicrosoftGraphDriveItem.md) | New navigation property | 

### Return type

[**MicrosoftGraphDriveItem**](MicrosoftGraphDriveItem.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DriveCreateItems

> MicrosoftGraphDriveItem DriveCreateItems(ctx).MicrosoftGraphDriveItem(microsoftGraphDriveItem).Execute()

Create new navigation property to items for drive



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
    microsoftGraphDriveItem := *openapiclient.NewMicrosoftGraphDriveItem() // MicrosoftGraphDriveItem | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DriveDriveItemApi.DriveCreateItems(context.Background()).MicrosoftGraphDriveItem(microsoftGraphDriveItem).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DriveDriveItemApi.DriveCreateItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DriveCreateItems`: MicrosoftGraphDriveItem
    fmt.Fprintf(os.Stdout, "Response from `DriveDriveItemApi.DriveCreateItems`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDriveCreateItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphDriveItem** | [**MicrosoftGraphDriveItem**](MicrosoftGraphDriveItem.md) | New navigation property | 

### Return type

[**MicrosoftGraphDriveItem**](MicrosoftGraphDriveItem.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DriveCreateSpecial

> MicrosoftGraphDriveItem DriveCreateSpecial(ctx).MicrosoftGraphDriveItem(microsoftGraphDriveItem).Execute()

Create new navigation property to special for drive



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
    microsoftGraphDriveItem := *openapiclient.NewMicrosoftGraphDriveItem() // MicrosoftGraphDriveItem | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DriveDriveItemApi.DriveCreateSpecial(context.Background()).MicrosoftGraphDriveItem(microsoftGraphDriveItem).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DriveDriveItemApi.DriveCreateSpecial``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DriveCreateSpecial`: MicrosoftGraphDriveItem
    fmt.Fprintf(os.Stdout, "Response from `DriveDriveItemApi.DriveCreateSpecial`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDriveCreateSpecialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphDriveItem** | [**MicrosoftGraphDriveItem**](MicrosoftGraphDriveItem.md) | New navigation property | 

### Return type

[**MicrosoftGraphDriveItem**](MicrosoftGraphDriveItem.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DriveDeleteBundles

> DriveDeleteBundles(ctx, driveItemId).IfMatch(ifMatch).Execute()

Delete navigation property bundles for drive



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
    resp, r, err := api_client.DriveDriveItemApi.DriveDeleteBundles(context.Background(), driveItemId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DriveDriveItemApi.DriveDeleteBundles``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDriveDeleteBundlesRequest struct via the builder pattern


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


## DriveDeleteFollowing

> DriveDeleteFollowing(ctx, driveItemId).IfMatch(ifMatch).Execute()

Delete navigation property following for drive



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
    resp, r, err := api_client.DriveDriveItemApi.DriveDeleteFollowing(context.Background(), driveItemId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DriveDriveItemApi.DriveDeleteFollowing``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDriveDeleteFollowingRequest struct via the builder pattern


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


## DriveDeleteItems

> DriveDeleteItems(ctx, driveItemId).IfMatch(ifMatch).Execute()

Delete navigation property items for drive



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
    resp, r, err := api_client.DriveDriveItemApi.DriveDeleteItems(context.Background(), driveItemId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DriveDriveItemApi.DriveDeleteItems``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDriveDeleteItemsRequest struct via the builder pattern


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


## DriveDeleteRoot

> DriveDeleteRoot(ctx).IfMatch(ifMatch).Execute()

Delete navigation property root for drive



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
    resp, r, err := api_client.DriveDriveItemApi.DriveDeleteRoot(context.Background()).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DriveDriveItemApi.DriveDeleteRoot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDriveDeleteRootRequest struct via the builder pattern


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


## DriveDeleteSpecial

> DriveDeleteSpecial(ctx, driveItemId).IfMatch(ifMatch).Execute()

Delete navigation property special for drive



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
    resp, r, err := api_client.DriveDriveItemApi.DriveDeleteSpecial(context.Background(), driveItemId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DriveDriveItemApi.DriveDeleteSpecial``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDriveDeleteSpecialRequest struct via the builder pattern


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


## DriveGetBundles

> MicrosoftGraphDriveItem DriveGetBundles(ctx, driveItemId).Select_(select_).Expand(expand).Execute()

Get bundles from drive



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
    resp, r, err := api_client.DriveDriveItemApi.DriveGetBundles(context.Background(), driveItemId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DriveDriveItemApi.DriveGetBundles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DriveGetBundles`: MicrosoftGraphDriveItem
    fmt.Fprintf(os.Stdout, "Response from `DriveDriveItemApi.DriveGetBundles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string** | key: id of driveItem | 

### Other Parameters

Other parameters are passed through a pointer to a apiDriveGetBundlesRequest struct via the builder pattern


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


## DriveGetBundlesContent

> *os.File DriveGetBundlesContent(ctx, driveItemId).Execute()

Get media content for the navigation property bundles from drive

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
    resp, r, err := api_client.DriveDriveItemApi.DriveGetBundlesContent(context.Background(), driveItemId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DriveDriveItemApi.DriveGetBundlesContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DriveGetBundlesContent`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `DriveDriveItemApi.DriveGetBundlesContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string** | key: id of driveItem | 

### Other Parameters

Other parameters are passed through a pointer to a apiDriveGetBundlesContentRequest struct via the builder pattern


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


## DriveGetFollowing

> MicrosoftGraphDriveItem DriveGetFollowing(ctx, driveItemId).Select_(select_).Expand(expand).Execute()

Get following from drive



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
    resp, r, err := api_client.DriveDriveItemApi.DriveGetFollowing(context.Background(), driveItemId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DriveDriveItemApi.DriveGetFollowing``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DriveGetFollowing`: MicrosoftGraphDriveItem
    fmt.Fprintf(os.Stdout, "Response from `DriveDriveItemApi.DriveGetFollowing`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string** | key: id of driveItem | 

### Other Parameters

Other parameters are passed through a pointer to a apiDriveGetFollowingRequest struct via the builder pattern


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


## DriveGetFollowingContent

> *os.File DriveGetFollowingContent(ctx, driveItemId).Execute()

Get media content for the navigation property following from drive

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
    resp, r, err := api_client.DriveDriveItemApi.DriveGetFollowingContent(context.Background(), driveItemId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DriveDriveItemApi.DriveGetFollowingContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DriveGetFollowingContent`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `DriveDriveItemApi.DriveGetFollowingContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string** | key: id of driveItem | 

### Other Parameters

Other parameters are passed through a pointer to a apiDriveGetFollowingContentRequest struct via the builder pattern


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


## DriveGetItems

> MicrosoftGraphDriveItem DriveGetItems(ctx, driveItemId).Select_(select_).Expand(expand).Execute()

Get items from drive



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
    resp, r, err := api_client.DriveDriveItemApi.DriveGetItems(context.Background(), driveItemId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DriveDriveItemApi.DriveGetItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DriveGetItems`: MicrosoftGraphDriveItem
    fmt.Fprintf(os.Stdout, "Response from `DriveDriveItemApi.DriveGetItems`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string** | key: id of driveItem | 

### Other Parameters

Other parameters are passed through a pointer to a apiDriveGetItemsRequest struct via the builder pattern


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


## DriveGetItemsContent

> *os.File DriveGetItemsContent(ctx, driveItemId).Execute()

Get media content for the navigation property items from drive

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
    resp, r, err := api_client.DriveDriveItemApi.DriveGetItemsContent(context.Background(), driveItemId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DriveDriveItemApi.DriveGetItemsContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DriveGetItemsContent`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `DriveDriveItemApi.DriveGetItemsContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string** | key: id of driveItem | 

### Other Parameters

Other parameters are passed through a pointer to a apiDriveGetItemsContentRequest struct via the builder pattern


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


## DriveGetRoot

> MicrosoftGraphDriveItem DriveGetRoot(ctx).Select_(select_).Expand(expand).Execute()

Get root from drive



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
    resp, r, err := api_client.DriveDriveItemApi.DriveGetRoot(context.Background()).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DriveDriveItemApi.DriveGetRoot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DriveGetRoot`: MicrosoftGraphDriveItem
    fmt.Fprintf(os.Stdout, "Response from `DriveDriveItemApi.DriveGetRoot`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDriveGetRootRequest struct via the builder pattern


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


## DriveGetRootContent

> *os.File DriveGetRootContent(ctx).Execute()

Get media content for the navigation property root from drive

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DriveDriveItemApi.DriveGetRootContent(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DriveDriveItemApi.DriveGetRootContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DriveGetRootContent`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `DriveDriveItemApi.DriveGetRootContent`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDriveGetRootContentRequest struct via the builder pattern


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


## DriveGetSpecial

> MicrosoftGraphDriveItem DriveGetSpecial(ctx, driveItemId).Select_(select_).Expand(expand).Execute()

Get special from drive



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
    resp, r, err := api_client.DriveDriveItemApi.DriveGetSpecial(context.Background(), driveItemId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DriveDriveItemApi.DriveGetSpecial``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DriveGetSpecial`: MicrosoftGraphDriveItem
    fmt.Fprintf(os.Stdout, "Response from `DriveDriveItemApi.DriveGetSpecial`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string** | key: id of driveItem | 

### Other Parameters

Other parameters are passed through a pointer to a apiDriveGetSpecialRequest struct via the builder pattern


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


## DriveGetSpecialContent

> *os.File DriveGetSpecialContent(ctx, driveItemId).Execute()

Get media content for the navigation property special from drive

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
    resp, r, err := api_client.DriveDriveItemApi.DriveGetSpecialContent(context.Background(), driveItemId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DriveDriveItemApi.DriveGetSpecialContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DriveGetSpecialContent`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `DriveDriveItemApi.DriveGetSpecialContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string** | key: id of driveItem | 

### Other Parameters

Other parameters are passed through a pointer to a apiDriveGetSpecialContentRequest struct via the builder pattern


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


## DriveListBundles

> CollectionOfDriveItem DriveListBundles(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get bundles from drive



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
    resp, r, err := api_client.DriveDriveItemApi.DriveListBundles(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DriveDriveItemApi.DriveListBundles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DriveListBundles`: CollectionOfDriveItem
    fmt.Fprintf(os.Stdout, "Response from `DriveDriveItemApi.DriveListBundles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDriveListBundlesRequest struct via the builder pattern


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

[**CollectionOfDriveItem**](CollectionOfDriveItem.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DriveListFollowing

> CollectionOfDriveItem DriveListFollowing(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get following from drive



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
    resp, r, err := api_client.DriveDriveItemApi.DriveListFollowing(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DriveDriveItemApi.DriveListFollowing``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DriveListFollowing`: CollectionOfDriveItem
    fmt.Fprintf(os.Stdout, "Response from `DriveDriveItemApi.DriveListFollowing`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDriveListFollowingRequest struct via the builder pattern


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

[**CollectionOfDriveItem**](CollectionOfDriveItem.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DriveListItems

> CollectionOfDriveItem DriveListItems(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get items from drive



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
    resp, r, err := api_client.DriveDriveItemApi.DriveListItems(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DriveDriveItemApi.DriveListItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DriveListItems`: CollectionOfDriveItem
    fmt.Fprintf(os.Stdout, "Response from `DriveDriveItemApi.DriveListItems`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDriveListItemsRequest struct via the builder pattern


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

[**CollectionOfDriveItem**](CollectionOfDriveItem.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DriveListSpecial

> CollectionOfDriveItem DriveListSpecial(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get special from drive



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
    resp, r, err := api_client.DriveDriveItemApi.DriveListSpecial(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DriveDriveItemApi.DriveListSpecial``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DriveListSpecial`: CollectionOfDriveItem
    fmt.Fprintf(os.Stdout, "Response from `DriveDriveItemApi.DriveListSpecial`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDriveListSpecialRequest struct via the builder pattern


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

[**CollectionOfDriveItem**](CollectionOfDriveItem.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DriveUpdateBundles

> DriveUpdateBundles(ctx, driveItemId).MicrosoftGraphDriveItem(microsoftGraphDriveItem).Execute()

Update the navigation property bundles in drive



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
    resp, r, err := api_client.DriveDriveItemApi.DriveUpdateBundles(context.Background(), driveItemId).MicrosoftGraphDriveItem(microsoftGraphDriveItem).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DriveDriveItemApi.DriveUpdateBundles``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDriveUpdateBundlesRequest struct via the builder pattern


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


## DriveUpdateBundlesContent

> DriveUpdateBundlesContent(ctx, driveItemId).Body(body).Execute()

Update media content for the navigation property bundles in drive

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
    resp, r, err := api_client.DriveDriveItemApi.DriveUpdateBundlesContent(context.Background(), driveItemId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DriveDriveItemApi.DriveUpdateBundlesContent``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDriveUpdateBundlesContentRequest struct via the builder pattern


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


## DriveUpdateFollowing

> DriveUpdateFollowing(ctx, driveItemId).MicrosoftGraphDriveItem(microsoftGraphDriveItem).Execute()

Update the navigation property following in drive



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
    resp, r, err := api_client.DriveDriveItemApi.DriveUpdateFollowing(context.Background(), driveItemId).MicrosoftGraphDriveItem(microsoftGraphDriveItem).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DriveDriveItemApi.DriveUpdateFollowing``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDriveUpdateFollowingRequest struct via the builder pattern


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


## DriveUpdateFollowingContent

> DriveUpdateFollowingContent(ctx, driveItemId).Body(body).Execute()

Update media content for the navigation property following in drive

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
    resp, r, err := api_client.DriveDriveItemApi.DriveUpdateFollowingContent(context.Background(), driveItemId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DriveDriveItemApi.DriveUpdateFollowingContent``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDriveUpdateFollowingContentRequest struct via the builder pattern


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


## DriveUpdateItems

> DriveUpdateItems(ctx, driveItemId).MicrosoftGraphDriveItem(microsoftGraphDriveItem).Execute()

Update the navigation property items in drive



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
    resp, r, err := api_client.DriveDriveItemApi.DriveUpdateItems(context.Background(), driveItemId).MicrosoftGraphDriveItem(microsoftGraphDriveItem).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DriveDriveItemApi.DriveUpdateItems``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDriveUpdateItemsRequest struct via the builder pattern


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


## DriveUpdateItemsContent

> DriveUpdateItemsContent(ctx, driveItemId).Body(body).Execute()

Update media content for the navigation property items in drive

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
    resp, r, err := api_client.DriveDriveItemApi.DriveUpdateItemsContent(context.Background(), driveItemId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DriveDriveItemApi.DriveUpdateItemsContent``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDriveUpdateItemsContentRequest struct via the builder pattern


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


## DriveUpdateRoot

> DriveUpdateRoot(ctx).MicrosoftGraphDriveItem(microsoftGraphDriveItem).Execute()

Update the navigation property root in drive



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
    microsoftGraphDriveItem := *openapiclient.NewMicrosoftGraphDriveItem() // MicrosoftGraphDriveItem | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DriveDriveItemApi.DriveUpdateRoot(context.Background()).MicrosoftGraphDriveItem(microsoftGraphDriveItem).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DriveDriveItemApi.DriveUpdateRoot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDriveUpdateRootRequest struct via the builder pattern


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


## DriveUpdateRootContent

> DriveUpdateRootContent(ctx).Body(body).Execute()

Update media content for the navigation property root in drive

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
    body := os.NewFile(1234, "some_file") // *os.File | New media content.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DriveDriveItemApi.DriveUpdateRootContent(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DriveDriveItemApi.DriveUpdateRootContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDriveUpdateRootContentRequest struct via the builder pattern


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


## DriveUpdateSpecial

> DriveUpdateSpecial(ctx, driveItemId).MicrosoftGraphDriveItem(microsoftGraphDriveItem).Execute()

Update the navigation property special in drive



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
    resp, r, err := api_client.DriveDriveItemApi.DriveUpdateSpecial(context.Background(), driveItemId).MicrosoftGraphDriveItem(microsoftGraphDriveItem).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DriveDriveItemApi.DriveUpdateSpecial``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDriveUpdateSpecialRequest struct via the builder pattern


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


## DriveUpdateSpecialContent

> DriveUpdateSpecialContent(ctx, driveItemId).Body(body).Execute()

Update media content for the navigation property special in drive

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
    resp, r, err := api_client.DriveDriveItemApi.DriveUpdateSpecialContent(context.Background(), driveItemId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DriveDriveItemApi.DriveUpdateSpecialContent``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDriveUpdateSpecialContentRequest struct via the builder pattern


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

