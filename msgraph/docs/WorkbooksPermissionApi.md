# \WorkbooksPermissionApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WorkbooksCreatePermissions**](WorkbooksPermissionApi.md#WorkbooksCreatePermissions) | **Post** /workbooks/{driveItem-id}/permissions | Create new navigation property to permissions for workbooks
[**WorkbooksDeletePermissions**](WorkbooksPermissionApi.md#WorkbooksDeletePermissions) | **Delete** /workbooks/{driveItem-id}/permissions/{permission-id} | Delete navigation property permissions for workbooks
[**WorkbooksGetPermissions**](WorkbooksPermissionApi.md#WorkbooksGetPermissions) | **Get** /workbooks/{driveItem-id}/permissions/{permission-id} | Get permissions from workbooks
[**WorkbooksListPermissions**](WorkbooksPermissionApi.md#WorkbooksListPermissions) | **Get** /workbooks/{driveItem-id}/permissions | Get permissions from workbooks
[**WorkbooksUpdatePermissions**](WorkbooksPermissionApi.md#WorkbooksUpdatePermissions) | **Patch** /workbooks/{driveItem-id}/permissions/{permission-id} | Update the navigation property permissions in workbooks



## WorkbooksCreatePermissions

> MicrosoftGraphPermission WorkbooksCreatePermissions(ctx, driveItemId).MicrosoftGraphPermission(microsoftGraphPermission).Execute()

Create new navigation property to permissions for workbooks



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
    microsoftGraphPermission := *openapiclient.NewMicrosoftGraphPermission() // MicrosoftGraphPermission | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkbooksPermissionApi.WorkbooksCreatePermissions(context.Background(), driveItemId).MicrosoftGraphPermission(microsoftGraphPermission).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkbooksPermissionApi.WorkbooksCreatePermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WorkbooksCreatePermissions`: MicrosoftGraphPermission
    fmt.Fprintf(os.Stdout, "Response from `WorkbooksPermissionApi.WorkbooksCreatePermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string** | key: id of driveItem | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkbooksCreatePermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphPermission** | [**MicrosoftGraphPermission**](MicrosoftGraphPermission.md) | New navigation property | 

### Return type

[**MicrosoftGraphPermission**](MicrosoftGraphPermission.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkbooksDeletePermissions

> WorkbooksDeletePermissions(ctx, driveItemId, permissionId).IfMatch(ifMatch).Execute()

Delete navigation property permissions for workbooks



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
    permissionId := "permissionId_example" // string | key: id of permission
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkbooksPermissionApi.WorkbooksDeletePermissions(context.Background(), driveItemId, permissionId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkbooksPermissionApi.WorkbooksDeletePermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string** | key: id of driveItem | 
**permissionId** | **string** | key: id of permission | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkbooksDeletePermissionsRequest struct via the builder pattern


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


## WorkbooksGetPermissions

> MicrosoftGraphPermission WorkbooksGetPermissions(ctx, driveItemId, permissionId).Select_(select_).Expand(expand).Execute()

Get permissions from workbooks



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
    permissionId := "permissionId_example" // string | key: id of permission
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkbooksPermissionApi.WorkbooksGetPermissions(context.Background(), driveItemId, permissionId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkbooksPermissionApi.WorkbooksGetPermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WorkbooksGetPermissions`: MicrosoftGraphPermission
    fmt.Fprintf(os.Stdout, "Response from `WorkbooksPermissionApi.WorkbooksGetPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string** | key: id of driveItem | 
**permissionId** | **string** | key: id of permission | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkbooksGetPermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphPermission**](MicrosoftGraphPermission.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkbooksListPermissions

> CollectionOfPermission WorkbooksListPermissions(ctx, driveItemId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get permissions from workbooks



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
    resp, r, err := api_client.WorkbooksPermissionApi.WorkbooksListPermissions(context.Background(), driveItemId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkbooksPermissionApi.WorkbooksListPermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WorkbooksListPermissions`: CollectionOfPermission
    fmt.Fprintf(os.Stdout, "Response from `WorkbooksPermissionApi.WorkbooksListPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string** | key: id of driveItem | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkbooksListPermissionsRequest struct via the builder pattern


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

[**CollectionOfPermission**](CollectionOfPermission.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkbooksUpdatePermissions

> WorkbooksUpdatePermissions(ctx, driveItemId, permissionId).MicrosoftGraphPermission(microsoftGraphPermission).Execute()

Update the navigation property permissions in workbooks



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
    permissionId := "permissionId_example" // string | key: id of permission
    microsoftGraphPermission := *openapiclient.NewMicrosoftGraphPermission() // MicrosoftGraphPermission | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkbooksPermissionApi.WorkbooksUpdatePermissions(context.Background(), driveItemId, permissionId).MicrosoftGraphPermission(microsoftGraphPermission).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkbooksPermissionApi.WorkbooksUpdatePermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string** | key: id of driveItem | 
**permissionId** | **string** | key: id of permission | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkbooksUpdatePermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **microsoftGraphPermission** | [**MicrosoftGraphPermission**](MicrosoftGraphPermission.md) | New navigation property values | 

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

