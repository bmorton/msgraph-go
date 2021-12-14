# \DeviceAppManagementManagedAppStatusApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceAppManagementCreateManagedAppStatuses**](DeviceAppManagementManagedAppStatusApi.md#DeviceAppManagementCreateManagedAppStatuses) | **Post** /deviceAppManagement/managedAppStatuses | Create new navigation property to managedAppStatuses for deviceAppManagement
[**DeviceAppManagementDeleteManagedAppStatuses**](DeviceAppManagementManagedAppStatusApi.md#DeviceAppManagementDeleteManagedAppStatuses) | **Delete** /deviceAppManagement/managedAppStatuses/{managedAppStatus-id} | Delete navigation property managedAppStatuses for deviceAppManagement
[**DeviceAppManagementGetManagedAppStatuses**](DeviceAppManagementManagedAppStatusApi.md#DeviceAppManagementGetManagedAppStatuses) | **Get** /deviceAppManagement/managedAppStatuses/{managedAppStatus-id} | Get managedAppStatuses from deviceAppManagement
[**DeviceAppManagementListManagedAppStatuses**](DeviceAppManagementManagedAppStatusApi.md#DeviceAppManagementListManagedAppStatuses) | **Get** /deviceAppManagement/managedAppStatuses | Get managedAppStatuses from deviceAppManagement
[**DeviceAppManagementUpdateManagedAppStatuses**](DeviceAppManagementManagedAppStatusApi.md#DeviceAppManagementUpdateManagedAppStatuses) | **Patch** /deviceAppManagement/managedAppStatuses/{managedAppStatus-id} | Update the navigation property managedAppStatuses in deviceAppManagement



## DeviceAppManagementCreateManagedAppStatuses

> MicrosoftGraphManagedAppStatus DeviceAppManagementCreateManagedAppStatuses(ctx).MicrosoftGraphManagedAppStatus(microsoftGraphManagedAppStatus).Execute()

Create new navigation property to managedAppStatuses for deviceAppManagement



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
    microsoftGraphManagedAppStatus := *openapiclient.NewMicrosoftGraphManagedAppStatus() // MicrosoftGraphManagedAppStatus | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceAppManagementManagedAppStatusApi.DeviceAppManagementCreateManagedAppStatuses(context.Background()).MicrosoftGraphManagedAppStatus(microsoftGraphManagedAppStatus).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAppManagementManagedAppStatusApi.DeviceAppManagementCreateManagedAppStatuses``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceAppManagementCreateManagedAppStatuses`: MicrosoftGraphManagedAppStatus
    fmt.Fprintf(os.Stdout, "Response from `DeviceAppManagementManagedAppStatusApi.DeviceAppManagementCreateManagedAppStatuses`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeviceAppManagementCreateManagedAppStatusesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphManagedAppStatus** | [**MicrosoftGraphManagedAppStatus**](MicrosoftGraphManagedAppStatus.md) | New navigation property | 

### Return type

[**MicrosoftGraphManagedAppStatus**](MicrosoftGraphManagedAppStatus.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementDeleteManagedAppStatuses

> DeviceAppManagementDeleteManagedAppStatuses(ctx, managedAppStatusId).IfMatch(ifMatch).Execute()

Delete navigation property managedAppStatuses for deviceAppManagement



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
    managedAppStatusId := "managedAppStatusId_example" // string | key: id of managedAppStatus
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceAppManagementManagedAppStatusApi.DeviceAppManagementDeleteManagedAppStatuses(context.Background(), managedAppStatusId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAppManagementManagedAppStatusApi.DeviceAppManagementDeleteManagedAppStatuses``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managedAppStatusId** | **string** | key: id of managedAppStatus | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceAppManagementDeleteManagedAppStatusesRequest struct via the builder pattern


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


## DeviceAppManagementGetManagedAppStatuses

> MicrosoftGraphManagedAppStatus DeviceAppManagementGetManagedAppStatuses(ctx, managedAppStatusId).Select_(select_).Expand(expand).Execute()

Get managedAppStatuses from deviceAppManagement



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
    managedAppStatusId := "managedAppStatusId_example" // string | key: id of managedAppStatus
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceAppManagementManagedAppStatusApi.DeviceAppManagementGetManagedAppStatuses(context.Background(), managedAppStatusId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAppManagementManagedAppStatusApi.DeviceAppManagementGetManagedAppStatuses``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceAppManagementGetManagedAppStatuses`: MicrosoftGraphManagedAppStatus
    fmt.Fprintf(os.Stdout, "Response from `DeviceAppManagementManagedAppStatusApi.DeviceAppManagementGetManagedAppStatuses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managedAppStatusId** | **string** | key: id of managedAppStatus | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceAppManagementGetManagedAppStatusesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphManagedAppStatus**](MicrosoftGraphManagedAppStatus.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementListManagedAppStatuses

> CollectionOfManagedAppStatus DeviceAppManagementListManagedAppStatuses(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get managedAppStatuses from deviceAppManagement



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
    resp, r, err := api_client.DeviceAppManagementManagedAppStatusApi.DeviceAppManagementListManagedAppStatuses(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAppManagementManagedAppStatusApi.DeviceAppManagementListManagedAppStatuses``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceAppManagementListManagedAppStatuses`: CollectionOfManagedAppStatus
    fmt.Fprintf(os.Stdout, "Response from `DeviceAppManagementManagedAppStatusApi.DeviceAppManagementListManagedAppStatuses`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeviceAppManagementListManagedAppStatusesRequest struct via the builder pattern


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

[**CollectionOfManagedAppStatus**](CollectionOfManagedAppStatus.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementUpdateManagedAppStatuses

> DeviceAppManagementUpdateManagedAppStatuses(ctx, managedAppStatusId).MicrosoftGraphManagedAppStatus(microsoftGraphManagedAppStatus).Execute()

Update the navigation property managedAppStatuses in deviceAppManagement



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
    managedAppStatusId := "managedAppStatusId_example" // string | key: id of managedAppStatus
    microsoftGraphManagedAppStatus := *openapiclient.NewMicrosoftGraphManagedAppStatus() // MicrosoftGraphManagedAppStatus | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceAppManagementManagedAppStatusApi.DeviceAppManagementUpdateManagedAppStatuses(context.Background(), managedAppStatusId).MicrosoftGraphManagedAppStatus(microsoftGraphManagedAppStatus).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAppManagementManagedAppStatusApi.DeviceAppManagementUpdateManagedAppStatuses``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managedAppStatusId** | **string** | key: id of managedAppStatus | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceAppManagementUpdateManagedAppStatusesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphManagedAppStatus** | [**MicrosoftGraphManagedAppStatus**](MicrosoftGraphManagedAppStatus.md) | New navigation property values | 

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

