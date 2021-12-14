# \SitesPermissionApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SitesCreatePermissions**](SitesPermissionApi.md#SitesCreatePermissions) | **Post** /sites/{site-id}/permissions | Create new navigation property to permissions for sites
[**SitesDeletePermissions**](SitesPermissionApi.md#SitesDeletePermissions) | **Delete** /sites/{site-id}/permissions/{permission-id} | Delete navigation property permissions for sites
[**SitesGetPermissions**](SitesPermissionApi.md#SitesGetPermissions) | **Get** /sites/{site-id}/permissions/{permission-id} | Get permissions from sites
[**SitesListPermissions**](SitesPermissionApi.md#SitesListPermissions) | **Get** /sites/{site-id}/permissions | Get permissions from sites
[**SitesUpdatePermissions**](SitesPermissionApi.md#SitesUpdatePermissions) | **Patch** /sites/{site-id}/permissions/{permission-id} | Update the navigation property permissions in sites



## SitesCreatePermissions

> MicrosoftGraphPermission SitesCreatePermissions(ctx, siteId).MicrosoftGraphPermission(microsoftGraphPermission).Execute()

Create new navigation property to permissions for sites



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
    siteId := "siteId_example" // string | key: id of site
    microsoftGraphPermission := *openapiclient.NewMicrosoftGraphPermission() // MicrosoftGraphPermission | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SitesPermissionApi.SitesCreatePermissions(context.Background(), siteId).MicrosoftGraphPermission(microsoftGraphPermission).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesPermissionApi.SitesCreatePermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SitesCreatePermissions`: MicrosoftGraphPermission
    fmt.Fprintf(os.Stdout, "Response from `SitesPermissionApi.SitesCreatePermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesCreatePermissionsRequest struct via the builder pattern


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


## SitesDeletePermissions

> SitesDeletePermissions(ctx, siteId, permissionId).IfMatch(ifMatch).Execute()

Delete navigation property permissions for sites



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
    siteId := "siteId_example" // string | key: id of site
    permissionId := "permissionId_example" // string | key: id of permission
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SitesPermissionApi.SitesDeletePermissions(context.Background(), siteId, permissionId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesPermissionApi.SitesDeletePermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 
**permissionId** | **string** | key: id of permission | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesDeletePermissionsRequest struct via the builder pattern


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


## SitesGetPermissions

> MicrosoftGraphPermission SitesGetPermissions(ctx, siteId, permissionId).Select_(select_).Expand(expand).Execute()

Get permissions from sites



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
    siteId := "siteId_example" // string | key: id of site
    permissionId := "permissionId_example" // string | key: id of permission
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SitesPermissionApi.SitesGetPermissions(context.Background(), siteId, permissionId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesPermissionApi.SitesGetPermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SitesGetPermissions`: MicrosoftGraphPermission
    fmt.Fprintf(os.Stdout, "Response from `SitesPermissionApi.SitesGetPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 
**permissionId** | **string** | key: id of permission | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesGetPermissionsRequest struct via the builder pattern


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


## SitesListPermissions

> CollectionOfPermission SitesListPermissions(ctx, siteId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get permissions from sites



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
    siteId := "siteId_example" // string | key: id of site
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
    resp, r, err := api_client.SitesPermissionApi.SitesListPermissions(context.Background(), siteId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesPermissionApi.SitesListPermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SitesListPermissions`: CollectionOfPermission
    fmt.Fprintf(os.Stdout, "Response from `SitesPermissionApi.SitesListPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesListPermissionsRequest struct via the builder pattern


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


## SitesUpdatePermissions

> SitesUpdatePermissions(ctx, siteId, permissionId).MicrosoftGraphPermission(microsoftGraphPermission).Execute()

Update the navigation property permissions in sites



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
    siteId := "siteId_example" // string | key: id of site
    permissionId := "permissionId_example" // string | key: id of permission
    microsoftGraphPermission := *openapiclient.NewMicrosoftGraphPermission() // MicrosoftGraphPermission | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SitesPermissionApi.SitesUpdatePermissions(context.Background(), siteId, permissionId).MicrosoftGraphPermission(microsoftGraphPermission).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesPermissionApi.SitesUpdatePermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 
**permissionId** | **string** | key: id of permission | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesUpdatePermissionsRequest struct via the builder pattern


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

