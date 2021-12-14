# \DeviceManagementManagedDeviceApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceManagementCreateManagedDevices**](DeviceManagementManagedDeviceApi.md#DeviceManagementCreateManagedDevices) | **Post** /deviceManagement/managedDevices | Create new navigation property to managedDevices for deviceManagement
[**DeviceManagementDeleteManagedDevices**](DeviceManagementManagedDeviceApi.md#DeviceManagementDeleteManagedDevices) | **Delete** /deviceManagement/managedDevices/{managedDevice-id} | Delete navigation property managedDevices for deviceManagement
[**DeviceManagementGetManagedDevices**](DeviceManagementManagedDeviceApi.md#DeviceManagementGetManagedDevices) | **Get** /deviceManagement/managedDevices/{managedDevice-id} | Get managedDevices from deviceManagement
[**DeviceManagementListManagedDevices**](DeviceManagementManagedDeviceApi.md#DeviceManagementListManagedDevices) | **Get** /deviceManagement/managedDevices | Get managedDevices from deviceManagement
[**DeviceManagementManagedDevicesCreateDeviceCompliancePolicyStates**](DeviceManagementManagedDeviceApi.md#DeviceManagementManagedDevicesCreateDeviceCompliancePolicyStates) | **Post** /deviceManagement/managedDevices/{managedDevice-id}/deviceCompliancePolicyStates | Create new navigation property to deviceCompliancePolicyStates for deviceManagement
[**DeviceManagementManagedDevicesCreateDeviceConfigurationStates**](DeviceManagementManagedDeviceApi.md#DeviceManagementManagedDevicesCreateDeviceConfigurationStates) | **Post** /deviceManagement/managedDevices/{managedDevice-id}/deviceConfigurationStates | Create new navigation property to deviceConfigurationStates for deviceManagement
[**DeviceManagementManagedDevicesDeleteDeviceCategory**](DeviceManagementManagedDeviceApi.md#DeviceManagementManagedDevicesDeleteDeviceCategory) | **Delete** /deviceManagement/managedDevices/{managedDevice-id}/deviceCategory | Delete navigation property deviceCategory for deviceManagement
[**DeviceManagementManagedDevicesDeleteDeviceCompliancePolicyStates**](DeviceManagementManagedDeviceApi.md#DeviceManagementManagedDevicesDeleteDeviceCompliancePolicyStates) | **Delete** /deviceManagement/managedDevices/{managedDevice-id}/deviceCompliancePolicyStates/{deviceCompliancePolicyState-id} | Delete navigation property deviceCompliancePolicyStates for deviceManagement
[**DeviceManagementManagedDevicesDeleteDeviceConfigurationStates**](DeviceManagementManagedDeviceApi.md#DeviceManagementManagedDevicesDeleteDeviceConfigurationStates) | **Delete** /deviceManagement/managedDevices/{managedDevice-id}/deviceConfigurationStates/{deviceConfigurationState-id} | Delete navigation property deviceConfigurationStates for deviceManagement
[**DeviceManagementManagedDevicesGetDeviceCategory**](DeviceManagementManagedDeviceApi.md#DeviceManagementManagedDevicesGetDeviceCategory) | **Get** /deviceManagement/managedDevices/{managedDevice-id}/deviceCategory | Get deviceCategory from deviceManagement
[**DeviceManagementManagedDevicesGetDeviceCompliancePolicyStates**](DeviceManagementManagedDeviceApi.md#DeviceManagementManagedDevicesGetDeviceCompliancePolicyStates) | **Get** /deviceManagement/managedDevices/{managedDevice-id}/deviceCompliancePolicyStates/{deviceCompliancePolicyState-id} | Get deviceCompliancePolicyStates from deviceManagement
[**DeviceManagementManagedDevicesGetDeviceConfigurationStates**](DeviceManagementManagedDeviceApi.md#DeviceManagementManagedDevicesGetDeviceConfigurationStates) | **Get** /deviceManagement/managedDevices/{managedDevice-id}/deviceConfigurationStates/{deviceConfigurationState-id} | Get deviceConfigurationStates from deviceManagement
[**DeviceManagementManagedDevicesListDeviceCompliancePolicyStates**](DeviceManagementManagedDeviceApi.md#DeviceManagementManagedDevicesListDeviceCompliancePolicyStates) | **Get** /deviceManagement/managedDevices/{managedDevice-id}/deviceCompliancePolicyStates | Get deviceCompliancePolicyStates from deviceManagement
[**DeviceManagementManagedDevicesListDeviceConfigurationStates**](DeviceManagementManagedDeviceApi.md#DeviceManagementManagedDevicesListDeviceConfigurationStates) | **Get** /deviceManagement/managedDevices/{managedDevice-id}/deviceConfigurationStates | Get deviceConfigurationStates from deviceManagement
[**DeviceManagementManagedDevicesUpdateDeviceCategory**](DeviceManagementManagedDeviceApi.md#DeviceManagementManagedDevicesUpdateDeviceCategory) | **Patch** /deviceManagement/managedDevices/{managedDevice-id}/deviceCategory | Update the navigation property deviceCategory in deviceManagement
[**DeviceManagementManagedDevicesUpdateDeviceCompliancePolicyStates**](DeviceManagementManagedDeviceApi.md#DeviceManagementManagedDevicesUpdateDeviceCompliancePolicyStates) | **Patch** /deviceManagement/managedDevices/{managedDevice-id}/deviceCompliancePolicyStates/{deviceCompliancePolicyState-id} | Update the navigation property deviceCompliancePolicyStates in deviceManagement
[**DeviceManagementManagedDevicesUpdateDeviceConfigurationStates**](DeviceManagementManagedDeviceApi.md#DeviceManagementManagedDevicesUpdateDeviceConfigurationStates) | **Patch** /deviceManagement/managedDevices/{managedDevice-id}/deviceConfigurationStates/{deviceConfigurationState-id} | Update the navigation property deviceConfigurationStates in deviceManagement
[**DeviceManagementUpdateManagedDevices**](DeviceManagementManagedDeviceApi.md#DeviceManagementUpdateManagedDevices) | **Patch** /deviceManagement/managedDevices/{managedDevice-id} | Update the navigation property managedDevices in deviceManagement



## DeviceManagementCreateManagedDevices

> MicrosoftGraphManagedDevice DeviceManagementCreateManagedDevices(ctx).MicrosoftGraphManagedDevice(microsoftGraphManagedDevice).Execute()

Create new navigation property to managedDevices for deviceManagement



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
    microsoftGraphManagedDevice := *openapiclient.NewMicrosoftGraphManagedDevice() // MicrosoftGraphManagedDevice | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementManagedDeviceApi.DeviceManagementCreateManagedDevices(context.Background()).MicrosoftGraphManagedDevice(microsoftGraphManagedDevice).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementManagedDeviceApi.DeviceManagementCreateManagedDevices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceManagementCreateManagedDevices`: MicrosoftGraphManagedDevice
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementManagedDeviceApi.DeviceManagementCreateManagedDevices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementCreateManagedDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphManagedDevice** | [**MicrosoftGraphManagedDevice**](MicrosoftGraphManagedDevice.md) | New navigation property | 

### Return type

[**MicrosoftGraphManagedDevice**](MicrosoftGraphManagedDevice.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementDeleteManagedDevices

> DeviceManagementDeleteManagedDevices(ctx, managedDeviceId).IfMatch(ifMatch).Execute()

Delete navigation property managedDevices for deviceManagement



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
    managedDeviceId := "managedDeviceId_example" // string | key: id of managedDevice
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementManagedDeviceApi.DeviceManagementDeleteManagedDevices(context.Background(), managedDeviceId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementManagedDeviceApi.DeviceManagementDeleteManagedDevices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managedDeviceId** | **string** | key: id of managedDevice | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementDeleteManagedDevicesRequest struct via the builder pattern


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


## DeviceManagementGetManagedDevices

> MicrosoftGraphManagedDevice DeviceManagementGetManagedDevices(ctx, managedDeviceId).Select_(select_).Expand(expand).Execute()

Get managedDevices from deviceManagement



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
    managedDeviceId := "managedDeviceId_example" // string | key: id of managedDevice
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementManagedDeviceApi.DeviceManagementGetManagedDevices(context.Background(), managedDeviceId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementManagedDeviceApi.DeviceManagementGetManagedDevices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceManagementGetManagedDevices`: MicrosoftGraphManagedDevice
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementManagedDeviceApi.DeviceManagementGetManagedDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managedDeviceId** | **string** | key: id of managedDevice | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementGetManagedDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphManagedDevice**](MicrosoftGraphManagedDevice.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementListManagedDevices

> CollectionOfManagedDevice DeviceManagementListManagedDevices(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get managedDevices from deviceManagement



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
    resp, r, err := api_client.DeviceManagementManagedDeviceApi.DeviceManagementListManagedDevices(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementManagedDeviceApi.DeviceManagementListManagedDevices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceManagementListManagedDevices`: CollectionOfManagedDevice
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementManagedDeviceApi.DeviceManagementListManagedDevices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementListManagedDevicesRequest struct via the builder pattern


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

[**CollectionOfManagedDevice**](CollectionOfManagedDevice.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementManagedDevicesCreateDeviceCompliancePolicyStates

> MicrosoftGraphDeviceCompliancePolicyState DeviceManagementManagedDevicesCreateDeviceCompliancePolicyStates(ctx, managedDeviceId).MicrosoftGraphDeviceCompliancePolicyState(microsoftGraphDeviceCompliancePolicyState).Execute()

Create new navigation property to deviceCompliancePolicyStates for deviceManagement



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
    managedDeviceId := "managedDeviceId_example" // string | key: id of managedDevice
    microsoftGraphDeviceCompliancePolicyState := *openapiclient.NewMicrosoftGraphDeviceCompliancePolicyState() // MicrosoftGraphDeviceCompliancePolicyState | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementManagedDeviceApi.DeviceManagementManagedDevicesCreateDeviceCompliancePolicyStates(context.Background(), managedDeviceId).MicrosoftGraphDeviceCompliancePolicyState(microsoftGraphDeviceCompliancePolicyState).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementManagedDeviceApi.DeviceManagementManagedDevicesCreateDeviceCompliancePolicyStates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceManagementManagedDevicesCreateDeviceCompliancePolicyStates`: MicrosoftGraphDeviceCompliancePolicyState
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementManagedDeviceApi.DeviceManagementManagedDevicesCreateDeviceCompliancePolicyStates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managedDeviceId** | **string** | key: id of managedDevice | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementManagedDevicesCreateDeviceCompliancePolicyStatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphDeviceCompliancePolicyState** | [**MicrosoftGraphDeviceCompliancePolicyState**](MicrosoftGraphDeviceCompliancePolicyState.md) | New navigation property | 

### Return type

[**MicrosoftGraphDeviceCompliancePolicyState**](MicrosoftGraphDeviceCompliancePolicyState.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementManagedDevicesCreateDeviceConfigurationStates

> MicrosoftGraphDeviceConfigurationState DeviceManagementManagedDevicesCreateDeviceConfigurationStates(ctx, managedDeviceId).MicrosoftGraphDeviceConfigurationState(microsoftGraphDeviceConfigurationState).Execute()

Create new navigation property to deviceConfigurationStates for deviceManagement



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
    managedDeviceId := "managedDeviceId_example" // string | key: id of managedDevice
    microsoftGraphDeviceConfigurationState := *openapiclient.NewMicrosoftGraphDeviceConfigurationState() // MicrosoftGraphDeviceConfigurationState | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementManagedDeviceApi.DeviceManagementManagedDevicesCreateDeviceConfigurationStates(context.Background(), managedDeviceId).MicrosoftGraphDeviceConfigurationState(microsoftGraphDeviceConfigurationState).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementManagedDeviceApi.DeviceManagementManagedDevicesCreateDeviceConfigurationStates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceManagementManagedDevicesCreateDeviceConfigurationStates`: MicrosoftGraphDeviceConfigurationState
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementManagedDeviceApi.DeviceManagementManagedDevicesCreateDeviceConfigurationStates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managedDeviceId** | **string** | key: id of managedDevice | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementManagedDevicesCreateDeviceConfigurationStatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphDeviceConfigurationState** | [**MicrosoftGraphDeviceConfigurationState**](MicrosoftGraphDeviceConfigurationState.md) | New navigation property | 

### Return type

[**MicrosoftGraphDeviceConfigurationState**](MicrosoftGraphDeviceConfigurationState.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementManagedDevicesDeleteDeviceCategory

> DeviceManagementManagedDevicesDeleteDeviceCategory(ctx, managedDeviceId).IfMatch(ifMatch).Execute()

Delete navigation property deviceCategory for deviceManagement



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
    managedDeviceId := "managedDeviceId_example" // string | key: id of managedDevice
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementManagedDeviceApi.DeviceManagementManagedDevicesDeleteDeviceCategory(context.Background(), managedDeviceId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementManagedDeviceApi.DeviceManagementManagedDevicesDeleteDeviceCategory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managedDeviceId** | **string** | key: id of managedDevice | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementManagedDevicesDeleteDeviceCategoryRequest struct via the builder pattern


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


## DeviceManagementManagedDevicesDeleteDeviceCompliancePolicyStates

> DeviceManagementManagedDevicesDeleteDeviceCompliancePolicyStates(ctx, managedDeviceId, deviceCompliancePolicyStateId).IfMatch(ifMatch).Execute()

Delete navigation property deviceCompliancePolicyStates for deviceManagement



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
    managedDeviceId := "managedDeviceId_example" // string | key: id of managedDevice
    deviceCompliancePolicyStateId := "deviceCompliancePolicyStateId_example" // string | key: id of deviceCompliancePolicyState
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementManagedDeviceApi.DeviceManagementManagedDevicesDeleteDeviceCompliancePolicyStates(context.Background(), managedDeviceId, deviceCompliancePolicyStateId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementManagedDeviceApi.DeviceManagementManagedDevicesDeleteDeviceCompliancePolicyStates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managedDeviceId** | **string** | key: id of managedDevice | 
**deviceCompliancePolicyStateId** | **string** | key: id of deviceCompliancePolicyState | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementManagedDevicesDeleteDeviceCompliancePolicyStatesRequest struct via the builder pattern


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


## DeviceManagementManagedDevicesDeleteDeviceConfigurationStates

> DeviceManagementManagedDevicesDeleteDeviceConfigurationStates(ctx, managedDeviceId, deviceConfigurationStateId).IfMatch(ifMatch).Execute()

Delete navigation property deviceConfigurationStates for deviceManagement



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
    managedDeviceId := "managedDeviceId_example" // string | key: id of managedDevice
    deviceConfigurationStateId := "deviceConfigurationStateId_example" // string | key: id of deviceConfigurationState
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementManagedDeviceApi.DeviceManagementManagedDevicesDeleteDeviceConfigurationStates(context.Background(), managedDeviceId, deviceConfigurationStateId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementManagedDeviceApi.DeviceManagementManagedDevicesDeleteDeviceConfigurationStates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managedDeviceId** | **string** | key: id of managedDevice | 
**deviceConfigurationStateId** | **string** | key: id of deviceConfigurationState | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementManagedDevicesDeleteDeviceConfigurationStatesRequest struct via the builder pattern


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


## DeviceManagementManagedDevicesGetDeviceCategory

> MicrosoftGraphDeviceCategory DeviceManagementManagedDevicesGetDeviceCategory(ctx, managedDeviceId).Select_(select_).Expand(expand).Execute()

Get deviceCategory from deviceManagement



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
    managedDeviceId := "managedDeviceId_example" // string | key: id of managedDevice
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementManagedDeviceApi.DeviceManagementManagedDevicesGetDeviceCategory(context.Background(), managedDeviceId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementManagedDeviceApi.DeviceManagementManagedDevicesGetDeviceCategory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceManagementManagedDevicesGetDeviceCategory`: MicrosoftGraphDeviceCategory
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementManagedDeviceApi.DeviceManagementManagedDevicesGetDeviceCategory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managedDeviceId** | **string** | key: id of managedDevice | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementManagedDevicesGetDeviceCategoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphDeviceCategory**](MicrosoftGraphDeviceCategory.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementManagedDevicesGetDeviceCompliancePolicyStates

> MicrosoftGraphDeviceCompliancePolicyState DeviceManagementManagedDevicesGetDeviceCompliancePolicyStates(ctx, managedDeviceId, deviceCompliancePolicyStateId).Select_(select_).Expand(expand).Execute()

Get deviceCompliancePolicyStates from deviceManagement



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
    managedDeviceId := "managedDeviceId_example" // string | key: id of managedDevice
    deviceCompliancePolicyStateId := "deviceCompliancePolicyStateId_example" // string | key: id of deviceCompliancePolicyState
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementManagedDeviceApi.DeviceManagementManagedDevicesGetDeviceCompliancePolicyStates(context.Background(), managedDeviceId, deviceCompliancePolicyStateId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementManagedDeviceApi.DeviceManagementManagedDevicesGetDeviceCompliancePolicyStates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceManagementManagedDevicesGetDeviceCompliancePolicyStates`: MicrosoftGraphDeviceCompliancePolicyState
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementManagedDeviceApi.DeviceManagementManagedDevicesGetDeviceCompliancePolicyStates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managedDeviceId** | **string** | key: id of managedDevice | 
**deviceCompliancePolicyStateId** | **string** | key: id of deviceCompliancePolicyState | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementManagedDevicesGetDeviceCompliancePolicyStatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphDeviceCompliancePolicyState**](MicrosoftGraphDeviceCompliancePolicyState.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementManagedDevicesGetDeviceConfigurationStates

> MicrosoftGraphDeviceConfigurationState DeviceManagementManagedDevicesGetDeviceConfigurationStates(ctx, managedDeviceId, deviceConfigurationStateId).Select_(select_).Expand(expand).Execute()

Get deviceConfigurationStates from deviceManagement



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
    managedDeviceId := "managedDeviceId_example" // string | key: id of managedDevice
    deviceConfigurationStateId := "deviceConfigurationStateId_example" // string | key: id of deviceConfigurationState
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementManagedDeviceApi.DeviceManagementManagedDevicesGetDeviceConfigurationStates(context.Background(), managedDeviceId, deviceConfigurationStateId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementManagedDeviceApi.DeviceManagementManagedDevicesGetDeviceConfigurationStates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceManagementManagedDevicesGetDeviceConfigurationStates`: MicrosoftGraphDeviceConfigurationState
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementManagedDeviceApi.DeviceManagementManagedDevicesGetDeviceConfigurationStates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managedDeviceId** | **string** | key: id of managedDevice | 
**deviceConfigurationStateId** | **string** | key: id of deviceConfigurationState | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementManagedDevicesGetDeviceConfigurationStatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphDeviceConfigurationState**](MicrosoftGraphDeviceConfigurationState.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementManagedDevicesListDeviceCompliancePolicyStates

> CollectionOfDeviceCompliancePolicyState DeviceManagementManagedDevicesListDeviceCompliancePolicyStates(ctx, managedDeviceId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get deviceCompliancePolicyStates from deviceManagement



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
    managedDeviceId := "managedDeviceId_example" // string | key: id of managedDevice
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
    resp, r, err := api_client.DeviceManagementManagedDeviceApi.DeviceManagementManagedDevicesListDeviceCompliancePolicyStates(context.Background(), managedDeviceId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementManagedDeviceApi.DeviceManagementManagedDevicesListDeviceCompliancePolicyStates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceManagementManagedDevicesListDeviceCompliancePolicyStates`: CollectionOfDeviceCompliancePolicyState
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementManagedDeviceApi.DeviceManagementManagedDevicesListDeviceCompliancePolicyStates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managedDeviceId** | **string** | key: id of managedDevice | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementManagedDevicesListDeviceCompliancePolicyStatesRequest struct via the builder pattern


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

[**CollectionOfDeviceCompliancePolicyState**](CollectionOfDeviceCompliancePolicyState.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementManagedDevicesListDeviceConfigurationStates

> CollectionOfDeviceConfigurationState DeviceManagementManagedDevicesListDeviceConfigurationStates(ctx, managedDeviceId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get deviceConfigurationStates from deviceManagement



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
    managedDeviceId := "managedDeviceId_example" // string | key: id of managedDevice
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
    resp, r, err := api_client.DeviceManagementManagedDeviceApi.DeviceManagementManagedDevicesListDeviceConfigurationStates(context.Background(), managedDeviceId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementManagedDeviceApi.DeviceManagementManagedDevicesListDeviceConfigurationStates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceManagementManagedDevicesListDeviceConfigurationStates`: CollectionOfDeviceConfigurationState
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementManagedDeviceApi.DeviceManagementManagedDevicesListDeviceConfigurationStates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managedDeviceId** | **string** | key: id of managedDevice | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementManagedDevicesListDeviceConfigurationStatesRequest struct via the builder pattern


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

[**CollectionOfDeviceConfigurationState**](CollectionOfDeviceConfigurationState.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementManagedDevicesUpdateDeviceCategory

> DeviceManagementManagedDevicesUpdateDeviceCategory(ctx, managedDeviceId).MicrosoftGraphDeviceCategory(microsoftGraphDeviceCategory).Execute()

Update the navigation property deviceCategory in deviceManagement



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
    managedDeviceId := "managedDeviceId_example" // string | key: id of managedDevice
    microsoftGraphDeviceCategory := *openapiclient.NewMicrosoftGraphDeviceCategory() // MicrosoftGraphDeviceCategory | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementManagedDeviceApi.DeviceManagementManagedDevicesUpdateDeviceCategory(context.Background(), managedDeviceId).MicrosoftGraphDeviceCategory(microsoftGraphDeviceCategory).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementManagedDeviceApi.DeviceManagementManagedDevicesUpdateDeviceCategory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managedDeviceId** | **string** | key: id of managedDevice | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementManagedDevicesUpdateDeviceCategoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphDeviceCategory** | [**MicrosoftGraphDeviceCategory**](MicrosoftGraphDeviceCategory.md) | New navigation property values | 

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


## DeviceManagementManagedDevicesUpdateDeviceCompliancePolicyStates

> DeviceManagementManagedDevicesUpdateDeviceCompliancePolicyStates(ctx, managedDeviceId, deviceCompliancePolicyStateId).MicrosoftGraphDeviceCompliancePolicyState(microsoftGraphDeviceCompliancePolicyState).Execute()

Update the navigation property deviceCompliancePolicyStates in deviceManagement



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
    managedDeviceId := "managedDeviceId_example" // string | key: id of managedDevice
    deviceCompliancePolicyStateId := "deviceCompliancePolicyStateId_example" // string | key: id of deviceCompliancePolicyState
    microsoftGraphDeviceCompliancePolicyState := *openapiclient.NewMicrosoftGraphDeviceCompliancePolicyState() // MicrosoftGraphDeviceCompliancePolicyState | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementManagedDeviceApi.DeviceManagementManagedDevicesUpdateDeviceCompliancePolicyStates(context.Background(), managedDeviceId, deviceCompliancePolicyStateId).MicrosoftGraphDeviceCompliancePolicyState(microsoftGraphDeviceCompliancePolicyState).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementManagedDeviceApi.DeviceManagementManagedDevicesUpdateDeviceCompliancePolicyStates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managedDeviceId** | **string** | key: id of managedDevice | 
**deviceCompliancePolicyStateId** | **string** | key: id of deviceCompliancePolicyState | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementManagedDevicesUpdateDeviceCompliancePolicyStatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **microsoftGraphDeviceCompliancePolicyState** | [**MicrosoftGraphDeviceCompliancePolicyState**](MicrosoftGraphDeviceCompliancePolicyState.md) | New navigation property values | 

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


## DeviceManagementManagedDevicesUpdateDeviceConfigurationStates

> DeviceManagementManagedDevicesUpdateDeviceConfigurationStates(ctx, managedDeviceId, deviceConfigurationStateId).MicrosoftGraphDeviceConfigurationState(microsoftGraphDeviceConfigurationState).Execute()

Update the navigation property deviceConfigurationStates in deviceManagement



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
    managedDeviceId := "managedDeviceId_example" // string | key: id of managedDevice
    deviceConfigurationStateId := "deviceConfigurationStateId_example" // string | key: id of deviceConfigurationState
    microsoftGraphDeviceConfigurationState := *openapiclient.NewMicrosoftGraphDeviceConfigurationState() // MicrosoftGraphDeviceConfigurationState | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementManagedDeviceApi.DeviceManagementManagedDevicesUpdateDeviceConfigurationStates(context.Background(), managedDeviceId, deviceConfigurationStateId).MicrosoftGraphDeviceConfigurationState(microsoftGraphDeviceConfigurationState).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementManagedDeviceApi.DeviceManagementManagedDevicesUpdateDeviceConfigurationStates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managedDeviceId** | **string** | key: id of managedDevice | 
**deviceConfigurationStateId** | **string** | key: id of deviceConfigurationState | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementManagedDevicesUpdateDeviceConfigurationStatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **microsoftGraphDeviceConfigurationState** | [**MicrosoftGraphDeviceConfigurationState**](MicrosoftGraphDeviceConfigurationState.md) | New navigation property values | 

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


## DeviceManagementUpdateManagedDevices

> DeviceManagementUpdateManagedDevices(ctx, managedDeviceId).MicrosoftGraphManagedDevice(microsoftGraphManagedDevice).Execute()

Update the navigation property managedDevices in deviceManagement



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
    managedDeviceId := "managedDeviceId_example" // string | key: id of managedDevice
    microsoftGraphManagedDevice := *openapiclient.NewMicrosoftGraphManagedDevice() // MicrosoftGraphManagedDevice | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementManagedDeviceApi.DeviceManagementUpdateManagedDevices(context.Background(), managedDeviceId).MicrosoftGraphManagedDevice(microsoftGraphManagedDevice).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementManagedDeviceApi.DeviceManagementUpdateManagedDevices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managedDeviceId** | **string** | key: id of managedDevice | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementUpdateManagedDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphManagedDevice** | [**MicrosoftGraphManagedDevice**](MicrosoftGraphManagedDevice.md) | New navigation property values | 

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

