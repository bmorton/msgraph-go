# \SharesPermissionApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SharesDeletePermission**](SharesPermissionApi.md#SharesDeletePermission) | **Delete** /shares/{sharedDriveItem-id}/permission | Delete navigation property permission for shares
[**SharesGetPermission**](SharesPermissionApi.md#SharesGetPermission) | **Get** /shares/{sharedDriveItem-id}/permission | Get permission from shares
[**SharesUpdatePermission**](SharesPermissionApi.md#SharesUpdatePermission) | **Patch** /shares/{sharedDriveItem-id}/permission | Update the navigation property permission in shares



## SharesDeletePermission

> SharesDeletePermission(ctx, sharedDriveItemId).IfMatch(ifMatch).Execute()

Delete navigation property permission for shares



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
    resp, r, err := api_client.SharesPermissionApi.SharesDeletePermission(context.Background(), sharedDriveItemId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SharesPermissionApi.SharesDeletePermission``: %v\n", err)
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

Other parameters are passed through a pointer to a apiSharesDeletePermissionRequest struct via the builder pattern


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


## SharesGetPermission

> MicrosoftGraphPermission SharesGetPermission(ctx, sharedDriveItemId).Select_(select_).Expand(expand).Execute()

Get permission from shares



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
    resp, r, err := api_client.SharesPermissionApi.SharesGetPermission(context.Background(), sharedDriveItemId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SharesPermissionApi.SharesGetPermission``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SharesGetPermission`: MicrosoftGraphPermission
    fmt.Fprintf(os.Stdout, "Response from `SharesPermissionApi.SharesGetPermission`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedDriveItemId** | **string** | key: id of sharedDriveItem | 

### Other Parameters

Other parameters are passed through a pointer to a apiSharesGetPermissionRequest struct via the builder pattern


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


## SharesUpdatePermission

> SharesUpdatePermission(ctx, sharedDriveItemId).MicrosoftGraphPermission(microsoftGraphPermission).Execute()

Update the navigation property permission in shares



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
    microsoftGraphPermission := *openapiclient.NewMicrosoftGraphPermission() // MicrosoftGraphPermission | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SharesPermissionApi.SharesUpdatePermission(context.Background(), sharedDriveItemId).MicrosoftGraphPermission(microsoftGraphPermission).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SharesPermissionApi.SharesUpdatePermission``: %v\n", err)
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

Other parameters are passed through a pointer to a apiSharesUpdatePermissionRequest struct via the builder pattern


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

