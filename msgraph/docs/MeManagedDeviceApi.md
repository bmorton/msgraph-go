# \MeManagedDeviceApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MeCreateManagedDevices**](MeManagedDeviceApi.md#MeCreateManagedDevices) | **Post** /me/managedDevices | Create new navigation property to managedDevices for me
[**MeDeleteManagedDevices**](MeManagedDeviceApi.md#MeDeleteManagedDevices) | **Delete** /me/managedDevices/{managedDevice-id} | Delete navigation property managedDevices for me
[**MeGetManagedDevices**](MeManagedDeviceApi.md#MeGetManagedDevices) | **Get** /me/managedDevices/{managedDevice-id} | Get managedDevices from me
[**MeListManagedDevices**](MeManagedDeviceApi.md#MeListManagedDevices) | **Get** /me/managedDevices | Get managedDevices from me
[**MeManagedDevicesCreateDeviceCompliancePolicyStates**](MeManagedDeviceApi.md#MeManagedDevicesCreateDeviceCompliancePolicyStates) | **Post** /me/managedDevices/{managedDevice-id}/deviceCompliancePolicyStates | Create new navigation property to deviceCompliancePolicyStates for me
[**MeManagedDevicesCreateDeviceConfigurationStates**](MeManagedDeviceApi.md#MeManagedDevicesCreateDeviceConfigurationStates) | **Post** /me/managedDevices/{managedDevice-id}/deviceConfigurationStates | Create new navigation property to deviceConfigurationStates for me
[**MeManagedDevicesDeleteDeviceCategory**](MeManagedDeviceApi.md#MeManagedDevicesDeleteDeviceCategory) | **Delete** /me/managedDevices/{managedDevice-id}/deviceCategory | Delete navigation property deviceCategory for me
[**MeManagedDevicesDeleteDeviceCompliancePolicyStates**](MeManagedDeviceApi.md#MeManagedDevicesDeleteDeviceCompliancePolicyStates) | **Delete** /me/managedDevices/{managedDevice-id}/deviceCompliancePolicyStates/{deviceCompliancePolicyState-id} | Delete navigation property deviceCompliancePolicyStates for me
[**MeManagedDevicesDeleteDeviceConfigurationStates**](MeManagedDeviceApi.md#MeManagedDevicesDeleteDeviceConfigurationStates) | **Delete** /me/managedDevices/{managedDevice-id}/deviceConfigurationStates/{deviceConfigurationState-id} | Delete navigation property deviceConfigurationStates for me
[**MeManagedDevicesGetDeviceCategory**](MeManagedDeviceApi.md#MeManagedDevicesGetDeviceCategory) | **Get** /me/managedDevices/{managedDevice-id}/deviceCategory | Get deviceCategory from me
[**MeManagedDevicesGetDeviceCompliancePolicyStates**](MeManagedDeviceApi.md#MeManagedDevicesGetDeviceCompliancePolicyStates) | **Get** /me/managedDevices/{managedDevice-id}/deviceCompliancePolicyStates/{deviceCompliancePolicyState-id} | Get deviceCompliancePolicyStates from me
[**MeManagedDevicesGetDeviceConfigurationStates**](MeManagedDeviceApi.md#MeManagedDevicesGetDeviceConfigurationStates) | **Get** /me/managedDevices/{managedDevice-id}/deviceConfigurationStates/{deviceConfigurationState-id} | Get deviceConfigurationStates from me
[**MeManagedDevicesListDeviceCompliancePolicyStates**](MeManagedDeviceApi.md#MeManagedDevicesListDeviceCompliancePolicyStates) | **Get** /me/managedDevices/{managedDevice-id}/deviceCompliancePolicyStates | Get deviceCompliancePolicyStates from me
[**MeManagedDevicesListDeviceConfigurationStates**](MeManagedDeviceApi.md#MeManagedDevicesListDeviceConfigurationStates) | **Get** /me/managedDevices/{managedDevice-id}/deviceConfigurationStates | Get deviceConfigurationStates from me
[**MeManagedDevicesUpdateDeviceCategory**](MeManagedDeviceApi.md#MeManagedDevicesUpdateDeviceCategory) | **Patch** /me/managedDevices/{managedDevice-id}/deviceCategory | Update the navigation property deviceCategory in me
[**MeManagedDevicesUpdateDeviceCompliancePolicyStates**](MeManagedDeviceApi.md#MeManagedDevicesUpdateDeviceCompliancePolicyStates) | **Patch** /me/managedDevices/{managedDevice-id}/deviceCompliancePolicyStates/{deviceCompliancePolicyState-id} | Update the navigation property deviceCompliancePolicyStates in me
[**MeManagedDevicesUpdateDeviceConfigurationStates**](MeManagedDeviceApi.md#MeManagedDevicesUpdateDeviceConfigurationStates) | **Patch** /me/managedDevices/{managedDevice-id}/deviceConfigurationStates/{deviceConfigurationState-id} | Update the navigation property deviceConfigurationStates in me
[**MeUpdateManagedDevices**](MeManagedDeviceApi.md#MeUpdateManagedDevices) | **Patch** /me/managedDevices/{managedDevice-id} | Update the navigation property managedDevices in me



## MeCreateManagedDevices

> MicrosoftGraphManagedDevice MeCreateManagedDevices(ctx).MicrosoftGraphManagedDevice(microsoftGraphManagedDevice).Execute()

Create new navigation property to managedDevices for me



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
    resp, r, err := api_client.MeManagedDeviceApi.MeCreateManagedDevices(context.Background()).MicrosoftGraphManagedDevice(microsoftGraphManagedDevice).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeManagedDeviceApi.MeCreateManagedDevices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeCreateManagedDevices`: MicrosoftGraphManagedDevice
    fmt.Fprintf(os.Stdout, "Response from `MeManagedDeviceApi.MeCreateManagedDevices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMeCreateManagedDevicesRequest struct via the builder pattern


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


## MeDeleteManagedDevices

> MeDeleteManagedDevices(ctx, managedDeviceId).IfMatch(ifMatch).Execute()

Delete navigation property managedDevices for me



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
    resp, r, err := api_client.MeManagedDeviceApi.MeDeleteManagedDevices(context.Background(), managedDeviceId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeManagedDeviceApi.MeDeleteManagedDevices``: %v\n", err)
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

Other parameters are passed through a pointer to a apiMeDeleteManagedDevicesRequest struct via the builder pattern


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


## MeGetManagedDevices

> MicrosoftGraphManagedDevice MeGetManagedDevices(ctx, managedDeviceId).Select_(select_).Expand(expand).Execute()

Get managedDevices from me



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
    resp, r, err := api_client.MeManagedDeviceApi.MeGetManagedDevices(context.Background(), managedDeviceId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeManagedDeviceApi.MeGetManagedDevices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeGetManagedDevices`: MicrosoftGraphManagedDevice
    fmt.Fprintf(os.Stdout, "Response from `MeManagedDeviceApi.MeGetManagedDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managedDeviceId** | **string** | key: id of managedDevice | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeGetManagedDevicesRequest struct via the builder pattern


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


## MeListManagedDevices

> CollectionOfManagedDevice MeListManagedDevices(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get managedDevices from me



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
    resp, r, err := api_client.MeManagedDeviceApi.MeListManagedDevices(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeManagedDeviceApi.MeListManagedDevices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeListManagedDevices`: CollectionOfManagedDevice
    fmt.Fprintf(os.Stdout, "Response from `MeManagedDeviceApi.MeListManagedDevices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMeListManagedDevicesRequest struct via the builder pattern


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


## MeManagedDevicesCreateDeviceCompliancePolicyStates

> MicrosoftGraphDeviceCompliancePolicyState MeManagedDevicesCreateDeviceCompliancePolicyStates(ctx, managedDeviceId).MicrosoftGraphDeviceCompliancePolicyState(microsoftGraphDeviceCompliancePolicyState).Execute()

Create new navigation property to deviceCompliancePolicyStates for me



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
    resp, r, err := api_client.MeManagedDeviceApi.MeManagedDevicesCreateDeviceCompliancePolicyStates(context.Background(), managedDeviceId).MicrosoftGraphDeviceCompliancePolicyState(microsoftGraphDeviceCompliancePolicyState).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeManagedDeviceApi.MeManagedDevicesCreateDeviceCompliancePolicyStates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeManagedDevicesCreateDeviceCompliancePolicyStates`: MicrosoftGraphDeviceCompliancePolicyState
    fmt.Fprintf(os.Stdout, "Response from `MeManagedDeviceApi.MeManagedDevicesCreateDeviceCompliancePolicyStates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managedDeviceId** | **string** | key: id of managedDevice | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeManagedDevicesCreateDeviceCompliancePolicyStatesRequest struct via the builder pattern


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


## MeManagedDevicesCreateDeviceConfigurationStates

> MicrosoftGraphDeviceConfigurationState MeManagedDevicesCreateDeviceConfigurationStates(ctx, managedDeviceId).MicrosoftGraphDeviceConfigurationState(microsoftGraphDeviceConfigurationState).Execute()

Create new navigation property to deviceConfigurationStates for me



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
    resp, r, err := api_client.MeManagedDeviceApi.MeManagedDevicesCreateDeviceConfigurationStates(context.Background(), managedDeviceId).MicrosoftGraphDeviceConfigurationState(microsoftGraphDeviceConfigurationState).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeManagedDeviceApi.MeManagedDevicesCreateDeviceConfigurationStates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeManagedDevicesCreateDeviceConfigurationStates`: MicrosoftGraphDeviceConfigurationState
    fmt.Fprintf(os.Stdout, "Response from `MeManagedDeviceApi.MeManagedDevicesCreateDeviceConfigurationStates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managedDeviceId** | **string** | key: id of managedDevice | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeManagedDevicesCreateDeviceConfigurationStatesRequest struct via the builder pattern


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


## MeManagedDevicesDeleteDeviceCategory

> MeManagedDevicesDeleteDeviceCategory(ctx, managedDeviceId).IfMatch(ifMatch).Execute()

Delete navigation property deviceCategory for me



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
    resp, r, err := api_client.MeManagedDeviceApi.MeManagedDevicesDeleteDeviceCategory(context.Background(), managedDeviceId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeManagedDeviceApi.MeManagedDevicesDeleteDeviceCategory``: %v\n", err)
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

Other parameters are passed through a pointer to a apiMeManagedDevicesDeleteDeviceCategoryRequest struct via the builder pattern


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


## MeManagedDevicesDeleteDeviceCompliancePolicyStates

> MeManagedDevicesDeleteDeviceCompliancePolicyStates(ctx, managedDeviceId, deviceCompliancePolicyStateId).IfMatch(ifMatch).Execute()

Delete navigation property deviceCompliancePolicyStates for me



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
    resp, r, err := api_client.MeManagedDeviceApi.MeManagedDevicesDeleteDeviceCompliancePolicyStates(context.Background(), managedDeviceId, deviceCompliancePolicyStateId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeManagedDeviceApi.MeManagedDevicesDeleteDeviceCompliancePolicyStates``: %v\n", err)
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

Other parameters are passed through a pointer to a apiMeManagedDevicesDeleteDeviceCompliancePolicyStatesRequest struct via the builder pattern


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


## MeManagedDevicesDeleteDeviceConfigurationStates

> MeManagedDevicesDeleteDeviceConfigurationStates(ctx, managedDeviceId, deviceConfigurationStateId).IfMatch(ifMatch).Execute()

Delete navigation property deviceConfigurationStates for me



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
    resp, r, err := api_client.MeManagedDeviceApi.MeManagedDevicesDeleteDeviceConfigurationStates(context.Background(), managedDeviceId, deviceConfigurationStateId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeManagedDeviceApi.MeManagedDevicesDeleteDeviceConfigurationStates``: %v\n", err)
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

Other parameters are passed through a pointer to a apiMeManagedDevicesDeleteDeviceConfigurationStatesRequest struct via the builder pattern


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


## MeManagedDevicesGetDeviceCategory

> MicrosoftGraphDeviceCategory MeManagedDevicesGetDeviceCategory(ctx, managedDeviceId).Select_(select_).Expand(expand).Execute()

Get deviceCategory from me



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
    resp, r, err := api_client.MeManagedDeviceApi.MeManagedDevicesGetDeviceCategory(context.Background(), managedDeviceId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeManagedDeviceApi.MeManagedDevicesGetDeviceCategory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeManagedDevicesGetDeviceCategory`: MicrosoftGraphDeviceCategory
    fmt.Fprintf(os.Stdout, "Response from `MeManagedDeviceApi.MeManagedDevicesGetDeviceCategory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managedDeviceId** | **string** | key: id of managedDevice | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeManagedDevicesGetDeviceCategoryRequest struct via the builder pattern


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


## MeManagedDevicesGetDeviceCompliancePolicyStates

> MicrosoftGraphDeviceCompliancePolicyState MeManagedDevicesGetDeviceCompliancePolicyStates(ctx, managedDeviceId, deviceCompliancePolicyStateId).Select_(select_).Expand(expand).Execute()

Get deviceCompliancePolicyStates from me



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
    resp, r, err := api_client.MeManagedDeviceApi.MeManagedDevicesGetDeviceCompliancePolicyStates(context.Background(), managedDeviceId, deviceCompliancePolicyStateId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeManagedDeviceApi.MeManagedDevicesGetDeviceCompliancePolicyStates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeManagedDevicesGetDeviceCompliancePolicyStates`: MicrosoftGraphDeviceCompliancePolicyState
    fmt.Fprintf(os.Stdout, "Response from `MeManagedDeviceApi.MeManagedDevicesGetDeviceCompliancePolicyStates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managedDeviceId** | **string** | key: id of managedDevice | 
**deviceCompliancePolicyStateId** | **string** | key: id of deviceCompliancePolicyState | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeManagedDevicesGetDeviceCompliancePolicyStatesRequest struct via the builder pattern


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


## MeManagedDevicesGetDeviceConfigurationStates

> MicrosoftGraphDeviceConfigurationState MeManagedDevicesGetDeviceConfigurationStates(ctx, managedDeviceId, deviceConfigurationStateId).Select_(select_).Expand(expand).Execute()

Get deviceConfigurationStates from me



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
    resp, r, err := api_client.MeManagedDeviceApi.MeManagedDevicesGetDeviceConfigurationStates(context.Background(), managedDeviceId, deviceConfigurationStateId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeManagedDeviceApi.MeManagedDevicesGetDeviceConfigurationStates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeManagedDevicesGetDeviceConfigurationStates`: MicrosoftGraphDeviceConfigurationState
    fmt.Fprintf(os.Stdout, "Response from `MeManagedDeviceApi.MeManagedDevicesGetDeviceConfigurationStates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managedDeviceId** | **string** | key: id of managedDevice | 
**deviceConfigurationStateId** | **string** | key: id of deviceConfigurationState | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeManagedDevicesGetDeviceConfigurationStatesRequest struct via the builder pattern


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


## MeManagedDevicesListDeviceCompliancePolicyStates

> CollectionOfDeviceCompliancePolicyState MeManagedDevicesListDeviceCompliancePolicyStates(ctx, managedDeviceId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get deviceCompliancePolicyStates from me



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
    resp, r, err := api_client.MeManagedDeviceApi.MeManagedDevicesListDeviceCompliancePolicyStates(context.Background(), managedDeviceId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeManagedDeviceApi.MeManagedDevicesListDeviceCompliancePolicyStates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeManagedDevicesListDeviceCompliancePolicyStates`: CollectionOfDeviceCompliancePolicyState
    fmt.Fprintf(os.Stdout, "Response from `MeManagedDeviceApi.MeManagedDevicesListDeviceCompliancePolicyStates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managedDeviceId** | **string** | key: id of managedDevice | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeManagedDevicesListDeviceCompliancePolicyStatesRequest struct via the builder pattern


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


## MeManagedDevicesListDeviceConfigurationStates

> CollectionOfDeviceConfigurationState MeManagedDevicesListDeviceConfigurationStates(ctx, managedDeviceId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get deviceConfigurationStates from me



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
    resp, r, err := api_client.MeManagedDeviceApi.MeManagedDevicesListDeviceConfigurationStates(context.Background(), managedDeviceId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeManagedDeviceApi.MeManagedDevicesListDeviceConfigurationStates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeManagedDevicesListDeviceConfigurationStates`: CollectionOfDeviceConfigurationState
    fmt.Fprintf(os.Stdout, "Response from `MeManagedDeviceApi.MeManagedDevicesListDeviceConfigurationStates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managedDeviceId** | **string** | key: id of managedDevice | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeManagedDevicesListDeviceConfigurationStatesRequest struct via the builder pattern


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


## MeManagedDevicesUpdateDeviceCategory

> MeManagedDevicesUpdateDeviceCategory(ctx, managedDeviceId).MicrosoftGraphDeviceCategory(microsoftGraphDeviceCategory).Execute()

Update the navigation property deviceCategory in me



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
    resp, r, err := api_client.MeManagedDeviceApi.MeManagedDevicesUpdateDeviceCategory(context.Background(), managedDeviceId).MicrosoftGraphDeviceCategory(microsoftGraphDeviceCategory).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeManagedDeviceApi.MeManagedDevicesUpdateDeviceCategory``: %v\n", err)
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

Other parameters are passed through a pointer to a apiMeManagedDevicesUpdateDeviceCategoryRequest struct via the builder pattern


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


## MeManagedDevicesUpdateDeviceCompliancePolicyStates

> MeManagedDevicesUpdateDeviceCompliancePolicyStates(ctx, managedDeviceId, deviceCompliancePolicyStateId).MicrosoftGraphDeviceCompliancePolicyState(microsoftGraphDeviceCompliancePolicyState).Execute()

Update the navigation property deviceCompliancePolicyStates in me



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
    resp, r, err := api_client.MeManagedDeviceApi.MeManagedDevicesUpdateDeviceCompliancePolicyStates(context.Background(), managedDeviceId, deviceCompliancePolicyStateId).MicrosoftGraphDeviceCompliancePolicyState(microsoftGraphDeviceCompliancePolicyState).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeManagedDeviceApi.MeManagedDevicesUpdateDeviceCompliancePolicyStates``: %v\n", err)
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

Other parameters are passed through a pointer to a apiMeManagedDevicesUpdateDeviceCompliancePolicyStatesRequest struct via the builder pattern


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


## MeManagedDevicesUpdateDeviceConfigurationStates

> MeManagedDevicesUpdateDeviceConfigurationStates(ctx, managedDeviceId, deviceConfigurationStateId).MicrosoftGraphDeviceConfigurationState(microsoftGraphDeviceConfigurationState).Execute()

Update the navigation property deviceConfigurationStates in me



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
    resp, r, err := api_client.MeManagedDeviceApi.MeManagedDevicesUpdateDeviceConfigurationStates(context.Background(), managedDeviceId, deviceConfigurationStateId).MicrosoftGraphDeviceConfigurationState(microsoftGraphDeviceConfigurationState).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeManagedDeviceApi.MeManagedDevicesUpdateDeviceConfigurationStates``: %v\n", err)
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

Other parameters are passed through a pointer to a apiMeManagedDevicesUpdateDeviceConfigurationStatesRequest struct via the builder pattern


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


## MeUpdateManagedDevices

> MeUpdateManagedDevices(ctx, managedDeviceId).MicrosoftGraphManagedDevice(microsoftGraphManagedDevice).Execute()

Update the navigation property managedDevices in me



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
    resp, r, err := api_client.MeManagedDeviceApi.MeUpdateManagedDevices(context.Background(), managedDeviceId).MicrosoftGraphManagedDevice(microsoftGraphManagedDevice).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeManagedDeviceApi.MeUpdateManagedDevices``: %v\n", err)
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

Other parameters are passed through a pointer to a apiMeUpdateManagedDevicesRequest struct via the builder pattern


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

