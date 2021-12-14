# \DeviceManagementIosUpdateDeviceStatusApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceManagementCreateIosUpdateStatuses**](DeviceManagementIosUpdateDeviceStatusApi.md#DeviceManagementCreateIosUpdateStatuses) | **Post** /deviceManagement/iosUpdateStatuses | Create new navigation property to iosUpdateStatuses for deviceManagement
[**DeviceManagementDeleteIosUpdateStatuses**](DeviceManagementIosUpdateDeviceStatusApi.md#DeviceManagementDeleteIosUpdateStatuses) | **Delete** /deviceManagement/iosUpdateStatuses/{iosUpdateDeviceStatus-id} | Delete navigation property iosUpdateStatuses for deviceManagement
[**DeviceManagementGetIosUpdateStatuses**](DeviceManagementIosUpdateDeviceStatusApi.md#DeviceManagementGetIosUpdateStatuses) | **Get** /deviceManagement/iosUpdateStatuses/{iosUpdateDeviceStatus-id} | Get iosUpdateStatuses from deviceManagement
[**DeviceManagementListIosUpdateStatuses**](DeviceManagementIosUpdateDeviceStatusApi.md#DeviceManagementListIosUpdateStatuses) | **Get** /deviceManagement/iosUpdateStatuses | Get iosUpdateStatuses from deviceManagement
[**DeviceManagementUpdateIosUpdateStatuses**](DeviceManagementIosUpdateDeviceStatusApi.md#DeviceManagementUpdateIosUpdateStatuses) | **Patch** /deviceManagement/iosUpdateStatuses/{iosUpdateDeviceStatus-id} | Update the navigation property iosUpdateStatuses in deviceManagement



## DeviceManagementCreateIosUpdateStatuses

> MicrosoftGraphIosUpdateDeviceStatus DeviceManagementCreateIosUpdateStatuses(ctx).MicrosoftGraphIosUpdateDeviceStatus(microsoftGraphIosUpdateDeviceStatus).Execute()

Create new navigation property to iosUpdateStatuses for deviceManagement



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
    microsoftGraphIosUpdateDeviceStatus := *openapiclient.NewMicrosoftGraphIosUpdateDeviceStatus() // MicrosoftGraphIosUpdateDeviceStatus | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementIosUpdateDeviceStatusApi.DeviceManagementCreateIosUpdateStatuses(context.Background()).MicrosoftGraphIosUpdateDeviceStatus(microsoftGraphIosUpdateDeviceStatus).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementIosUpdateDeviceStatusApi.DeviceManagementCreateIosUpdateStatuses``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceManagementCreateIosUpdateStatuses`: MicrosoftGraphIosUpdateDeviceStatus
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementIosUpdateDeviceStatusApi.DeviceManagementCreateIosUpdateStatuses`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementCreateIosUpdateStatusesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphIosUpdateDeviceStatus** | [**MicrosoftGraphIosUpdateDeviceStatus**](MicrosoftGraphIosUpdateDeviceStatus.md) | New navigation property | 

### Return type

[**MicrosoftGraphIosUpdateDeviceStatus**](MicrosoftGraphIosUpdateDeviceStatus.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementDeleteIosUpdateStatuses

> DeviceManagementDeleteIosUpdateStatuses(ctx, iosUpdateDeviceStatusId).IfMatch(ifMatch).Execute()

Delete navigation property iosUpdateStatuses for deviceManagement



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
    iosUpdateDeviceStatusId := "iosUpdateDeviceStatusId_example" // string | key: id of iosUpdateDeviceStatus
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementIosUpdateDeviceStatusApi.DeviceManagementDeleteIosUpdateStatuses(context.Background(), iosUpdateDeviceStatusId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementIosUpdateDeviceStatusApi.DeviceManagementDeleteIosUpdateStatuses``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**iosUpdateDeviceStatusId** | **string** | key: id of iosUpdateDeviceStatus | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementDeleteIosUpdateStatusesRequest struct via the builder pattern


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


## DeviceManagementGetIosUpdateStatuses

> MicrosoftGraphIosUpdateDeviceStatus DeviceManagementGetIosUpdateStatuses(ctx, iosUpdateDeviceStatusId).Select_(select_).Expand(expand).Execute()

Get iosUpdateStatuses from deviceManagement



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
    iosUpdateDeviceStatusId := "iosUpdateDeviceStatusId_example" // string | key: id of iosUpdateDeviceStatus
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementIosUpdateDeviceStatusApi.DeviceManagementGetIosUpdateStatuses(context.Background(), iosUpdateDeviceStatusId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementIosUpdateDeviceStatusApi.DeviceManagementGetIosUpdateStatuses``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceManagementGetIosUpdateStatuses`: MicrosoftGraphIosUpdateDeviceStatus
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementIosUpdateDeviceStatusApi.DeviceManagementGetIosUpdateStatuses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**iosUpdateDeviceStatusId** | **string** | key: id of iosUpdateDeviceStatus | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementGetIosUpdateStatusesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphIosUpdateDeviceStatus**](MicrosoftGraphIosUpdateDeviceStatus.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementListIosUpdateStatuses

> CollectionOfIosUpdateDeviceStatus DeviceManagementListIosUpdateStatuses(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get iosUpdateStatuses from deviceManagement



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
    resp, r, err := api_client.DeviceManagementIosUpdateDeviceStatusApi.DeviceManagementListIosUpdateStatuses(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementIosUpdateDeviceStatusApi.DeviceManagementListIosUpdateStatuses``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceManagementListIosUpdateStatuses`: CollectionOfIosUpdateDeviceStatus
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementIosUpdateDeviceStatusApi.DeviceManagementListIosUpdateStatuses`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementListIosUpdateStatusesRequest struct via the builder pattern


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

[**CollectionOfIosUpdateDeviceStatus**](CollectionOfIosUpdateDeviceStatus.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementUpdateIosUpdateStatuses

> DeviceManagementUpdateIosUpdateStatuses(ctx, iosUpdateDeviceStatusId).MicrosoftGraphIosUpdateDeviceStatus(microsoftGraphIosUpdateDeviceStatus).Execute()

Update the navigation property iosUpdateStatuses in deviceManagement



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
    iosUpdateDeviceStatusId := "iosUpdateDeviceStatusId_example" // string | key: id of iosUpdateDeviceStatus
    microsoftGraphIosUpdateDeviceStatus := *openapiclient.NewMicrosoftGraphIosUpdateDeviceStatus() // MicrosoftGraphIosUpdateDeviceStatus | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementIosUpdateDeviceStatusApi.DeviceManagementUpdateIosUpdateStatuses(context.Background(), iosUpdateDeviceStatusId).MicrosoftGraphIosUpdateDeviceStatus(microsoftGraphIosUpdateDeviceStatus).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementIosUpdateDeviceStatusApi.DeviceManagementUpdateIosUpdateStatuses``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**iosUpdateDeviceStatusId** | **string** | key: id of iosUpdateDeviceStatus | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementUpdateIosUpdateStatusesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphIosUpdateDeviceStatus** | [**MicrosoftGraphIosUpdateDeviceStatus**](MicrosoftGraphIosUpdateDeviceStatus.md) | New navigation property values | 

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

