# \UsersManagedDeviceApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersCreateManagedDevices**](UsersManagedDeviceApi.md#UsersCreateManagedDevices) | **Post** /users/{user-id}/managedDevices | Create new navigation property to managedDevices for users
[**UsersDeleteManagedDevices**](UsersManagedDeviceApi.md#UsersDeleteManagedDevices) | **Delete** /users/{user-id}/managedDevices/{managedDevice-id} | Delete navigation property managedDevices for users
[**UsersGetManagedDevices**](UsersManagedDeviceApi.md#UsersGetManagedDevices) | **Get** /users/{user-id}/managedDevices/{managedDevice-id} | Get managedDevices from users
[**UsersListManagedDevices**](UsersManagedDeviceApi.md#UsersListManagedDevices) | **Get** /users/{user-id}/managedDevices | Get managedDevices from users
[**UsersManagedDevicesCreateDeviceCompliancePolicyStates**](UsersManagedDeviceApi.md#UsersManagedDevicesCreateDeviceCompliancePolicyStates) | **Post** /users/{user-id}/managedDevices/{managedDevice-id}/deviceCompliancePolicyStates | Create new navigation property to deviceCompliancePolicyStates for users
[**UsersManagedDevicesCreateDeviceConfigurationStates**](UsersManagedDeviceApi.md#UsersManagedDevicesCreateDeviceConfigurationStates) | **Post** /users/{user-id}/managedDevices/{managedDevice-id}/deviceConfigurationStates | Create new navigation property to deviceConfigurationStates for users
[**UsersManagedDevicesDeleteDeviceCategory**](UsersManagedDeviceApi.md#UsersManagedDevicesDeleteDeviceCategory) | **Delete** /users/{user-id}/managedDevices/{managedDevice-id}/deviceCategory | Delete navigation property deviceCategory for users
[**UsersManagedDevicesDeleteDeviceCompliancePolicyStates**](UsersManagedDeviceApi.md#UsersManagedDevicesDeleteDeviceCompliancePolicyStates) | **Delete** /users/{user-id}/managedDevices/{managedDevice-id}/deviceCompliancePolicyStates/{deviceCompliancePolicyState-id} | Delete navigation property deviceCompliancePolicyStates for users
[**UsersManagedDevicesDeleteDeviceConfigurationStates**](UsersManagedDeviceApi.md#UsersManagedDevicesDeleteDeviceConfigurationStates) | **Delete** /users/{user-id}/managedDevices/{managedDevice-id}/deviceConfigurationStates/{deviceConfigurationState-id} | Delete navigation property deviceConfigurationStates for users
[**UsersManagedDevicesGetDeviceCategory**](UsersManagedDeviceApi.md#UsersManagedDevicesGetDeviceCategory) | **Get** /users/{user-id}/managedDevices/{managedDevice-id}/deviceCategory | Get deviceCategory from users
[**UsersManagedDevicesGetDeviceCompliancePolicyStates**](UsersManagedDeviceApi.md#UsersManagedDevicesGetDeviceCompliancePolicyStates) | **Get** /users/{user-id}/managedDevices/{managedDevice-id}/deviceCompliancePolicyStates/{deviceCompliancePolicyState-id} | Get deviceCompliancePolicyStates from users
[**UsersManagedDevicesGetDeviceConfigurationStates**](UsersManagedDeviceApi.md#UsersManagedDevicesGetDeviceConfigurationStates) | **Get** /users/{user-id}/managedDevices/{managedDevice-id}/deviceConfigurationStates/{deviceConfigurationState-id} | Get deviceConfigurationStates from users
[**UsersManagedDevicesListDeviceCompliancePolicyStates**](UsersManagedDeviceApi.md#UsersManagedDevicesListDeviceCompliancePolicyStates) | **Get** /users/{user-id}/managedDevices/{managedDevice-id}/deviceCompliancePolicyStates | Get deviceCompliancePolicyStates from users
[**UsersManagedDevicesListDeviceConfigurationStates**](UsersManagedDeviceApi.md#UsersManagedDevicesListDeviceConfigurationStates) | **Get** /users/{user-id}/managedDevices/{managedDevice-id}/deviceConfigurationStates | Get deviceConfigurationStates from users
[**UsersManagedDevicesUpdateDeviceCategory**](UsersManagedDeviceApi.md#UsersManagedDevicesUpdateDeviceCategory) | **Patch** /users/{user-id}/managedDevices/{managedDevice-id}/deviceCategory | Update the navigation property deviceCategory in users
[**UsersManagedDevicesUpdateDeviceCompliancePolicyStates**](UsersManagedDeviceApi.md#UsersManagedDevicesUpdateDeviceCompliancePolicyStates) | **Patch** /users/{user-id}/managedDevices/{managedDevice-id}/deviceCompliancePolicyStates/{deviceCompliancePolicyState-id} | Update the navigation property deviceCompliancePolicyStates in users
[**UsersManagedDevicesUpdateDeviceConfigurationStates**](UsersManagedDeviceApi.md#UsersManagedDevicesUpdateDeviceConfigurationStates) | **Patch** /users/{user-id}/managedDevices/{managedDevice-id}/deviceConfigurationStates/{deviceConfigurationState-id} | Update the navigation property deviceConfigurationStates in users
[**UsersUpdateManagedDevices**](UsersManagedDeviceApi.md#UsersUpdateManagedDevices) | **Patch** /users/{user-id}/managedDevices/{managedDevice-id} | Update the navigation property managedDevices in users



## UsersCreateManagedDevices

> MicrosoftGraphManagedDevice UsersCreateManagedDevices(ctx, userId).MicrosoftGraphManagedDevice(microsoftGraphManagedDevice).Execute()

Create new navigation property to managedDevices for users



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
    userId := "userId_example" // string | key: id of user
    microsoftGraphManagedDevice := *openapiclient.NewMicrosoftGraphManagedDevice() // MicrosoftGraphManagedDevice | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersManagedDeviceApi.UsersCreateManagedDevices(context.Background(), userId).MicrosoftGraphManagedDevice(microsoftGraphManagedDevice).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersManagedDeviceApi.UsersCreateManagedDevices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersCreateManagedDevices`: MicrosoftGraphManagedDevice
    fmt.Fprintf(os.Stdout, "Response from `UsersManagedDeviceApi.UsersCreateManagedDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersCreateManagedDevicesRequest struct via the builder pattern


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


## UsersDeleteManagedDevices

> UsersDeleteManagedDevices(ctx, userId, managedDeviceId).IfMatch(ifMatch).Execute()

Delete navigation property managedDevices for users



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
    userId := "userId_example" // string | key: id of user
    managedDeviceId := "managedDeviceId_example" // string | key: id of managedDevice
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersManagedDeviceApi.UsersDeleteManagedDevices(context.Background(), userId, managedDeviceId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersManagedDeviceApi.UsersDeleteManagedDevices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**managedDeviceId** | **string** | key: id of managedDevice | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersDeleteManagedDevicesRequest struct via the builder pattern


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


## UsersGetManagedDevices

> MicrosoftGraphManagedDevice UsersGetManagedDevices(ctx, userId, managedDeviceId).Select_(select_).Expand(expand).Execute()

Get managedDevices from users



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
    userId := "userId_example" // string | key: id of user
    managedDeviceId := "managedDeviceId_example" // string | key: id of managedDevice
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersManagedDeviceApi.UsersGetManagedDevices(context.Background(), userId, managedDeviceId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersManagedDeviceApi.UsersGetManagedDevices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersGetManagedDevices`: MicrosoftGraphManagedDevice
    fmt.Fprintf(os.Stdout, "Response from `UsersManagedDeviceApi.UsersGetManagedDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**managedDeviceId** | **string** | key: id of managedDevice | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersGetManagedDevicesRequest struct via the builder pattern


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


## UsersListManagedDevices

> CollectionOfManagedDevice UsersListManagedDevices(ctx, userId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get managedDevices from users



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
    userId := "userId_example" // string | key: id of user
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
    resp, r, err := api_client.UsersManagedDeviceApi.UsersListManagedDevices(context.Background(), userId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersManagedDeviceApi.UsersListManagedDevices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersListManagedDevices`: CollectionOfManagedDevice
    fmt.Fprintf(os.Stdout, "Response from `UsersManagedDeviceApi.UsersListManagedDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersListManagedDevicesRequest struct via the builder pattern


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


## UsersManagedDevicesCreateDeviceCompliancePolicyStates

> MicrosoftGraphDeviceCompliancePolicyState UsersManagedDevicesCreateDeviceCompliancePolicyStates(ctx, userId, managedDeviceId).MicrosoftGraphDeviceCompliancePolicyState(microsoftGraphDeviceCompliancePolicyState).Execute()

Create new navigation property to deviceCompliancePolicyStates for users



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
    userId := "userId_example" // string | key: id of user
    managedDeviceId := "managedDeviceId_example" // string | key: id of managedDevice
    microsoftGraphDeviceCompliancePolicyState := *openapiclient.NewMicrosoftGraphDeviceCompliancePolicyState() // MicrosoftGraphDeviceCompliancePolicyState | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersManagedDeviceApi.UsersManagedDevicesCreateDeviceCompliancePolicyStates(context.Background(), userId, managedDeviceId).MicrosoftGraphDeviceCompliancePolicyState(microsoftGraphDeviceCompliancePolicyState).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersManagedDeviceApi.UsersManagedDevicesCreateDeviceCompliancePolicyStates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersManagedDevicesCreateDeviceCompliancePolicyStates`: MicrosoftGraphDeviceCompliancePolicyState
    fmt.Fprintf(os.Stdout, "Response from `UsersManagedDeviceApi.UsersManagedDevicesCreateDeviceCompliancePolicyStates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**managedDeviceId** | **string** | key: id of managedDevice | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersManagedDevicesCreateDeviceCompliancePolicyStatesRequest struct via the builder pattern


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


## UsersManagedDevicesCreateDeviceConfigurationStates

> MicrosoftGraphDeviceConfigurationState UsersManagedDevicesCreateDeviceConfigurationStates(ctx, userId, managedDeviceId).MicrosoftGraphDeviceConfigurationState(microsoftGraphDeviceConfigurationState).Execute()

Create new navigation property to deviceConfigurationStates for users



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
    userId := "userId_example" // string | key: id of user
    managedDeviceId := "managedDeviceId_example" // string | key: id of managedDevice
    microsoftGraphDeviceConfigurationState := *openapiclient.NewMicrosoftGraphDeviceConfigurationState() // MicrosoftGraphDeviceConfigurationState | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersManagedDeviceApi.UsersManagedDevicesCreateDeviceConfigurationStates(context.Background(), userId, managedDeviceId).MicrosoftGraphDeviceConfigurationState(microsoftGraphDeviceConfigurationState).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersManagedDeviceApi.UsersManagedDevicesCreateDeviceConfigurationStates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersManagedDevicesCreateDeviceConfigurationStates`: MicrosoftGraphDeviceConfigurationState
    fmt.Fprintf(os.Stdout, "Response from `UsersManagedDeviceApi.UsersManagedDevicesCreateDeviceConfigurationStates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**managedDeviceId** | **string** | key: id of managedDevice | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersManagedDevicesCreateDeviceConfigurationStatesRequest struct via the builder pattern


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


## UsersManagedDevicesDeleteDeviceCategory

> UsersManagedDevicesDeleteDeviceCategory(ctx, userId, managedDeviceId).IfMatch(ifMatch).Execute()

Delete navigation property deviceCategory for users



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
    userId := "userId_example" // string | key: id of user
    managedDeviceId := "managedDeviceId_example" // string | key: id of managedDevice
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersManagedDeviceApi.UsersManagedDevicesDeleteDeviceCategory(context.Background(), userId, managedDeviceId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersManagedDeviceApi.UsersManagedDevicesDeleteDeviceCategory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**managedDeviceId** | **string** | key: id of managedDevice | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersManagedDevicesDeleteDeviceCategoryRequest struct via the builder pattern


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


## UsersManagedDevicesDeleteDeviceCompliancePolicyStates

> UsersManagedDevicesDeleteDeviceCompliancePolicyStates(ctx, userId, managedDeviceId, deviceCompliancePolicyStateId).IfMatch(ifMatch).Execute()

Delete navigation property deviceCompliancePolicyStates for users



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
    userId := "userId_example" // string | key: id of user
    managedDeviceId := "managedDeviceId_example" // string | key: id of managedDevice
    deviceCompliancePolicyStateId := "deviceCompliancePolicyStateId_example" // string | key: id of deviceCompliancePolicyState
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersManagedDeviceApi.UsersManagedDevicesDeleteDeviceCompliancePolicyStates(context.Background(), userId, managedDeviceId, deviceCompliancePolicyStateId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersManagedDeviceApi.UsersManagedDevicesDeleteDeviceCompliancePolicyStates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**managedDeviceId** | **string** | key: id of managedDevice | 
**deviceCompliancePolicyStateId** | **string** | key: id of deviceCompliancePolicyState | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersManagedDevicesDeleteDeviceCompliancePolicyStatesRequest struct via the builder pattern


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


## UsersManagedDevicesDeleteDeviceConfigurationStates

> UsersManagedDevicesDeleteDeviceConfigurationStates(ctx, userId, managedDeviceId, deviceConfigurationStateId).IfMatch(ifMatch).Execute()

Delete navigation property deviceConfigurationStates for users



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
    userId := "userId_example" // string | key: id of user
    managedDeviceId := "managedDeviceId_example" // string | key: id of managedDevice
    deviceConfigurationStateId := "deviceConfigurationStateId_example" // string | key: id of deviceConfigurationState
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersManagedDeviceApi.UsersManagedDevicesDeleteDeviceConfigurationStates(context.Background(), userId, managedDeviceId, deviceConfigurationStateId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersManagedDeviceApi.UsersManagedDevicesDeleteDeviceConfigurationStates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**managedDeviceId** | **string** | key: id of managedDevice | 
**deviceConfigurationStateId** | **string** | key: id of deviceConfigurationState | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersManagedDevicesDeleteDeviceConfigurationStatesRequest struct via the builder pattern


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


## UsersManagedDevicesGetDeviceCategory

> MicrosoftGraphDeviceCategory UsersManagedDevicesGetDeviceCategory(ctx, userId, managedDeviceId).Select_(select_).Expand(expand).Execute()

Get deviceCategory from users



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
    userId := "userId_example" // string | key: id of user
    managedDeviceId := "managedDeviceId_example" // string | key: id of managedDevice
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersManagedDeviceApi.UsersManagedDevicesGetDeviceCategory(context.Background(), userId, managedDeviceId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersManagedDeviceApi.UsersManagedDevicesGetDeviceCategory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersManagedDevicesGetDeviceCategory`: MicrosoftGraphDeviceCategory
    fmt.Fprintf(os.Stdout, "Response from `UsersManagedDeviceApi.UsersManagedDevicesGetDeviceCategory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**managedDeviceId** | **string** | key: id of managedDevice | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersManagedDevicesGetDeviceCategoryRequest struct via the builder pattern


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


## UsersManagedDevicesGetDeviceCompliancePolicyStates

> MicrosoftGraphDeviceCompliancePolicyState UsersManagedDevicesGetDeviceCompliancePolicyStates(ctx, userId, managedDeviceId, deviceCompliancePolicyStateId).Select_(select_).Expand(expand).Execute()

Get deviceCompliancePolicyStates from users



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
    userId := "userId_example" // string | key: id of user
    managedDeviceId := "managedDeviceId_example" // string | key: id of managedDevice
    deviceCompliancePolicyStateId := "deviceCompliancePolicyStateId_example" // string | key: id of deviceCompliancePolicyState
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersManagedDeviceApi.UsersManagedDevicesGetDeviceCompliancePolicyStates(context.Background(), userId, managedDeviceId, deviceCompliancePolicyStateId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersManagedDeviceApi.UsersManagedDevicesGetDeviceCompliancePolicyStates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersManagedDevicesGetDeviceCompliancePolicyStates`: MicrosoftGraphDeviceCompliancePolicyState
    fmt.Fprintf(os.Stdout, "Response from `UsersManagedDeviceApi.UsersManagedDevicesGetDeviceCompliancePolicyStates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**managedDeviceId** | **string** | key: id of managedDevice | 
**deviceCompliancePolicyStateId** | **string** | key: id of deviceCompliancePolicyState | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersManagedDevicesGetDeviceCompliancePolicyStatesRequest struct via the builder pattern


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


## UsersManagedDevicesGetDeviceConfigurationStates

> MicrosoftGraphDeviceConfigurationState UsersManagedDevicesGetDeviceConfigurationStates(ctx, userId, managedDeviceId, deviceConfigurationStateId).Select_(select_).Expand(expand).Execute()

Get deviceConfigurationStates from users



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
    userId := "userId_example" // string | key: id of user
    managedDeviceId := "managedDeviceId_example" // string | key: id of managedDevice
    deviceConfigurationStateId := "deviceConfigurationStateId_example" // string | key: id of deviceConfigurationState
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersManagedDeviceApi.UsersManagedDevicesGetDeviceConfigurationStates(context.Background(), userId, managedDeviceId, deviceConfigurationStateId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersManagedDeviceApi.UsersManagedDevicesGetDeviceConfigurationStates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersManagedDevicesGetDeviceConfigurationStates`: MicrosoftGraphDeviceConfigurationState
    fmt.Fprintf(os.Stdout, "Response from `UsersManagedDeviceApi.UsersManagedDevicesGetDeviceConfigurationStates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**managedDeviceId** | **string** | key: id of managedDevice | 
**deviceConfigurationStateId** | **string** | key: id of deviceConfigurationState | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersManagedDevicesGetDeviceConfigurationStatesRequest struct via the builder pattern


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


## UsersManagedDevicesListDeviceCompliancePolicyStates

> CollectionOfDeviceCompliancePolicyState UsersManagedDevicesListDeviceCompliancePolicyStates(ctx, userId, managedDeviceId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get deviceCompliancePolicyStates from users



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
    userId := "userId_example" // string | key: id of user
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
    resp, r, err := api_client.UsersManagedDeviceApi.UsersManagedDevicesListDeviceCompliancePolicyStates(context.Background(), userId, managedDeviceId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersManagedDeviceApi.UsersManagedDevicesListDeviceCompliancePolicyStates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersManagedDevicesListDeviceCompliancePolicyStates`: CollectionOfDeviceCompliancePolicyState
    fmt.Fprintf(os.Stdout, "Response from `UsersManagedDeviceApi.UsersManagedDevicesListDeviceCompliancePolicyStates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**managedDeviceId** | **string** | key: id of managedDevice | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersManagedDevicesListDeviceCompliancePolicyStatesRequest struct via the builder pattern


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


## UsersManagedDevicesListDeviceConfigurationStates

> CollectionOfDeviceConfigurationState UsersManagedDevicesListDeviceConfigurationStates(ctx, userId, managedDeviceId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get deviceConfigurationStates from users



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
    userId := "userId_example" // string | key: id of user
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
    resp, r, err := api_client.UsersManagedDeviceApi.UsersManagedDevicesListDeviceConfigurationStates(context.Background(), userId, managedDeviceId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersManagedDeviceApi.UsersManagedDevicesListDeviceConfigurationStates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersManagedDevicesListDeviceConfigurationStates`: CollectionOfDeviceConfigurationState
    fmt.Fprintf(os.Stdout, "Response from `UsersManagedDeviceApi.UsersManagedDevicesListDeviceConfigurationStates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**managedDeviceId** | **string** | key: id of managedDevice | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersManagedDevicesListDeviceConfigurationStatesRequest struct via the builder pattern


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


## UsersManagedDevicesUpdateDeviceCategory

> UsersManagedDevicesUpdateDeviceCategory(ctx, userId, managedDeviceId).MicrosoftGraphDeviceCategory(microsoftGraphDeviceCategory).Execute()

Update the navigation property deviceCategory in users



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
    userId := "userId_example" // string | key: id of user
    managedDeviceId := "managedDeviceId_example" // string | key: id of managedDevice
    microsoftGraphDeviceCategory := *openapiclient.NewMicrosoftGraphDeviceCategory() // MicrosoftGraphDeviceCategory | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersManagedDeviceApi.UsersManagedDevicesUpdateDeviceCategory(context.Background(), userId, managedDeviceId).MicrosoftGraphDeviceCategory(microsoftGraphDeviceCategory).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersManagedDeviceApi.UsersManagedDevicesUpdateDeviceCategory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**managedDeviceId** | **string** | key: id of managedDevice | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersManagedDevicesUpdateDeviceCategoryRequest struct via the builder pattern


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


## UsersManagedDevicesUpdateDeviceCompliancePolicyStates

> UsersManagedDevicesUpdateDeviceCompliancePolicyStates(ctx, userId, managedDeviceId, deviceCompliancePolicyStateId).MicrosoftGraphDeviceCompliancePolicyState(microsoftGraphDeviceCompliancePolicyState).Execute()

Update the navigation property deviceCompliancePolicyStates in users



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
    userId := "userId_example" // string | key: id of user
    managedDeviceId := "managedDeviceId_example" // string | key: id of managedDevice
    deviceCompliancePolicyStateId := "deviceCompliancePolicyStateId_example" // string | key: id of deviceCompliancePolicyState
    microsoftGraphDeviceCompliancePolicyState := *openapiclient.NewMicrosoftGraphDeviceCompliancePolicyState() // MicrosoftGraphDeviceCompliancePolicyState | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersManagedDeviceApi.UsersManagedDevicesUpdateDeviceCompliancePolicyStates(context.Background(), userId, managedDeviceId, deviceCompliancePolicyStateId).MicrosoftGraphDeviceCompliancePolicyState(microsoftGraphDeviceCompliancePolicyState).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersManagedDeviceApi.UsersManagedDevicesUpdateDeviceCompliancePolicyStates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**managedDeviceId** | **string** | key: id of managedDevice | 
**deviceCompliancePolicyStateId** | **string** | key: id of deviceCompliancePolicyState | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersManagedDevicesUpdateDeviceCompliancePolicyStatesRequest struct via the builder pattern


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


## UsersManagedDevicesUpdateDeviceConfigurationStates

> UsersManagedDevicesUpdateDeviceConfigurationStates(ctx, userId, managedDeviceId, deviceConfigurationStateId).MicrosoftGraphDeviceConfigurationState(microsoftGraphDeviceConfigurationState).Execute()

Update the navigation property deviceConfigurationStates in users



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
    userId := "userId_example" // string | key: id of user
    managedDeviceId := "managedDeviceId_example" // string | key: id of managedDevice
    deviceConfigurationStateId := "deviceConfigurationStateId_example" // string | key: id of deviceConfigurationState
    microsoftGraphDeviceConfigurationState := *openapiclient.NewMicrosoftGraphDeviceConfigurationState() // MicrosoftGraphDeviceConfigurationState | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersManagedDeviceApi.UsersManagedDevicesUpdateDeviceConfigurationStates(context.Background(), userId, managedDeviceId, deviceConfigurationStateId).MicrosoftGraphDeviceConfigurationState(microsoftGraphDeviceConfigurationState).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersManagedDeviceApi.UsersManagedDevicesUpdateDeviceConfigurationStates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**managedDeviceId** | **string** | key: id of managedDevice | 
**deviceConfigurationStateId** | **string** | key: id of deviceConfigurationState | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersManagedDevicesUpdateDeviceConfigurationStatesRequest struct via the builder pattern


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


## UsersUpdateManagedDevices

> UsersUpdateManagedDevices(ctx, userId, managedDeviceId).MicrosoftGraphManagedDevice(microsoftGraphManagedDevice).Execute()

Update the navigation property managedDevices in users



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
    userId := "userId_example" // string | key: id of user
    managedDeviceId := "managedDeviceId_example" // string | key: id of managedDevice
    microsoftGraphManagedDevice := *openapiclient.NewMicrosoftGraphManagedDevice() // MicrosoftGraphManagedDevice | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersManagedDeviceApi.UsersUpdateManagedDevices(context.Background(), userId, managedDeviceId).MicrosoftGraphManagedDevice(microsoftGraphManagedDevice).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersManagedDeviceApi.UsersUpdateManagedDevices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**managedDeviceId** | **string** | key: id of managedDevice | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersUpdateManagedDevicesRequest struct via the builder pattern


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

