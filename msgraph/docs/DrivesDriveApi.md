# \DrivesDriveApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DrivesDriveCreateDrive**](DrivesDriveApi.md#DrivesDriveCreateDrive) | **Post** /drives | Add new entity to drives
[**DrivesDriveDeleteDrive**](DrivesDriveApi.md#DrivesDriveDeleteDrive) | **Delete** /drives/{drive-id} | Delete entity from drives
[**DrivesDriveGetDrive**](DrivesDriveApi.md#DrivesDriveGetDrive) | **Get** /drives/{drive-id} | Get entity from drives by key
[**DrivesDriveListDrive**](DrivesDriveApi.md#DrivesDriveListDrive) | **Get** /drives | Get entities from drives
[**DrivesDriveUpdateDrive**](DrivesDriveApi.md#DrivesDriveUpdateDrive) | **Patch** /drives/{drive-id} | Update entity in drives



## DrivesDriveCreateDrive

> MicrosoftGraphDrive DrivesDriveCreateDrive(ctx).MicrosoftGraphDrive(microsoftGraphDrive).Execute()

Add new entity to drives

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
    microsoftGraphDrive := *openapiclient.NewMicrosoftGraphDrive() // MicrosoftGraphDrive | New entity

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DrivesDriveApi.DrivesDriveCreateDrive(context.Background()).MicrosoftGraphDrive(microsoftGraphDrive).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DrivesDriveApi.DrivesDriveCreateDrive``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DrivesDriveCreateDrive`: MicrosoftGraphDrive
    fmt.Fprintf(os.Stdout, "Response from `DrivesDriveApi.DrivesDriveCreateDrive`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDrivesDriveCreateDriveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphDrive** | [**MicrosoftGraphDrive**](MicrosoftGraphDrive.md) | New entity | 

### Return type

[**MicrosoftGraphDrive**](MicrosoftGraphDrive.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DrivesDriveDeleteDrive

> DrivesDriveDeleteDrive(ctx, driveId).IfMatch(ifMatch).Execute()

Delete entity from drives

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
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DrivesDriveApi.DrivesDriveDeleteDrive(context.Background(), driveId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DrivesDriveApi.DrivesDriveDeleteDrive``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveId** | **string** | key: id of drive | 

### Other Parameters

Other parameters are passed through a pointer to a apiDrivesDriveDeleteDriveRequest struct via the builder pattern


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


## DrivesDriveGetDrive

> MicrosoftGraphDrive DrivesDriveGetDrive(ctx, driveId).Select_(select_).Expand(expand).Execute()

Get entity from drives by key

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
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DrivesDriveApi.DrivesDriveGetDrive(context.Background(), driveId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DrivesDriveApi.DrivesDriveGetDrive``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DrivesDriveGetDrive`: MicrosoftGraphDrive
    fmt.Fprintf(os.Stdout, "Response from `DrivesDriveApi.DrivesDriveGetDrive`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveId** | **string** | key: id of drive | 

### Other Parameters

Other parameters are passed through a pointer to a apiDrivesDriveGetDriveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphDrive**](MicrosoftGraphDrive.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DrivesDriveListDrive

> CollectionOfDrive DrivesDriveListDrive(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get entities from drives

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
    resp, r, err := api_client.DrivesDriveApi.DrivesDriveListDrive(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DrivesDriveApi.DrivesDriveListDrive``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DrivesDriveListDrive`: CollectionOfDrive
    fmt.Fprintf(os.Stdout, "Response from `DrivesDriveApi.DrivesDriveListDrive`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDrivesDriveListDriveRequest struct via the builder pattern


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

[**CollectionOfDrive**](CollectionOfDrive.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DrivesDriveUpdateDrive

> DrivesDriveUpdateDrive(ctx, driveId).MicrosoftGraphDrive(microsoftGraphDrive).Execute()

Update entity in drives

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
    microsoftGraphDrive := *openapiclient.NewMicrosoftGraphDrive() // MicrosoftGraphDrive | New property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DrivesDriveApi.DrivesDriveUpdateDrive(context.Background(), driveId).MicrosoftGraphDrive(microsoftGraphDrive).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DrivesDriveApi.DrivesDriveUpdateDrive``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveId** | **string** | key: id of drive | 

### Other Parameters

Other parameters are passed through a pointer to a apiDrivesDriveUpdateDriveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphDrive** | [**MicrosoftGraphDrive**](MicrosoftGraphDrive.md) | New property values | 

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

