# \DeviceManagementWindowsAutopilotDeviceIdentityApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceManagementCreateWindowsAutopilotDeviceIdentities**](DeviceManagementWindowsAutopilotDeviceIdentityApi.md#DeviceManagementCreateWindowsAutopilotDeviceIdentities) | **Post** /deviceManagement/windowsAutopilotDeviceIdentities | Create new navigation property to windowsAutopilotDeviceIdentities for deviceManagement
[**DeviceManagementDeleteWindowsAutopilotDeviceIdentities**](DeviceManagementWindowsAutopilotDeviceIdentityApi.md#DeviceManagementDeleteWindowsAutopilotDeviceIdentities) | **Delete** /deviceManagement/windowsAutopilotDeviceIdentities/{windowsAutopilotDeviceIdentity-id} | Delete navigation property windowsAutopilotDeviceIdentities for deviceManagement
[**DeviceManagementGetWindowsAutopilotDeviceIdentities**](DeviceManagementWindowsAutopilotDeviceIdentityApi.md#DeviceManagementGetWindowsAutopilotDeviceIdentities) | **Get** /deviceManagement/windowsAutopilotDeviceIdentities/{windowsAutopilotDeviceIdentity-id} | Get windowsAutopilotDeviceIdentities from deviceManagement
[**DeviceManagementListWindowsAutopilotDeviceIdentities**](DeviceManagementWindowsAutopilotDeviceIdentityApi.md#DeviceManagementListWindowsAutopilotDeviceIdentities) | **Get** /deviceManagement/windowsAutopilotDeviceIdentities | Get windowsAutopilotDeviceIdentities from deviceManagement
[**DeviceManagementUpdateWindowsAutopilotDeviceIdentities**](DeviceManagementWindowsAutopilotDeviceIdentityApi.md#DeviceManagementUpdateWindowsAutopilotDeviceIdentities) | **Patch** /deviceManagement/windowsAutopilotDeviceIdentities/{windowsAutopilotDeviceIdentity-id} | Update the navigation property windowsAutopilotDeviceIdentities in deviceManagement



## DeviceManagementCreateWindowsAutopilotDeviceIdentities

> MicrosoftGraphWindowsAutopilotDeviceIdentity DeviceManagementCreateWindowsAutopilotDeviceIdentities(ctx).MicrosoftGraphWindowsAutopilotDeviceIdentity(microsoftGraphWindowsAutopilotDeviceIdentity).Execute()

Create new navigation property to windowsAutopilotDeviceIdentities for deviceManagement



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
    microsoftGraphWindowsAutopilotDeviceIdentity := *openapiclient.NewMicrosoftGraphWindowsAutopilotDeviceIdentity() // MicrosoftGraphWindowsAutopilotDeviceIdentity | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementWindowsAutopilotDeviceIdentityApi.DeviceManagementCreateWindowsAutopilotDeviceIdentities(context.Background()).MicrosoftGraphWindowsAutopilotDeviceIdentity(microsoftGraphWindowsAutopilotDeviceIdentity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementWindowsAutopilotDeviceIdentityApi.DeviceManagementCreateWindowsAutopilotDeviceIdentities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceManagementCreateWindowsAutopilotDeviceIdentities`: MicrosoftGraphWindowsAutopilotDeviceIdentity
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementWindowsAutopilotDeviceIdentityApi.DeviceManagementCreateWindowsAutopilotDeviceIdentities`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementCreateWindowsAutopilotDeviceIdentitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphWindowsAutopilotDeviceIdentity** | [**MicrosoftGraphWindowsAutopilotDeviceIdentity**](MicrosoftGraphWindowsAutopilotDeviceIdentity.md) | New navigation property | 

### Return type

[**MicrosoftGraphWindowsAutopilotDeviceIdentity**](MicrosoftGraphWindowsAutopilotDeviceIdentity.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementDeleteWindowsAutopilotDeviceIdentities

> DeviceManagementDeleteWindowsAutopilotDeviceIdentities(ctx, windowsAutopilotDeviceIdentityId).IfMatch(ifMatch).Execute()

Delete navigation property windowsAutopilotDeviceIdentities for deviceManagement



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
    windowsAutopilotDeviceIdentityId := "windowsAutopilotDeviceIdentityId_example" // string | key: id of windowsAutopilotDeviceIdentity
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementWindowsAutopilotDeviceIdentityApi.DeviceManagementDeleteWindowsAutopilotDeviceIdentities(context.Background(), windowsAutopilotDeviceIdentityId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementWindowsAutopilotDeviceIdentityApi.DeviceManagementDeleteWindowsAutopilotDeviceIdentities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**windowsAutopilotDeviceIdentityId** | **string** | key: id of windowsAutopilotDeviceIdentity | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementDeleteWindowsAutopilotDeviceIdentitiesRequest struct via the builder pattern


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


## DeviceManagementGetWindowsAutopilotDeviceIdentities

> MicrosoftGraphWindowsAutopilotDeviceIdentity DeviceManagementGetWindowsAutopilotDeviceIdentities(ctx, windowsAutopilotDeviceIdentityId).Select_(select_).Expand(expand).Execute()

Get windowsAutopilotDeviceIdentities from deviceManagement



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
    windowsAutopilotDeviceIdentityId := "windowsAutopilotDeviceIdentityId_example" // string | key: id of windowsAutopilotDeviceIdentity
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementWindowsAutopilotDeviceIdentityApi.DeviceManagementGetWindowsAutopilotDeviceIdentities(context.Background(), windowsAutopilotDeviceIdentityId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementWindowsAutopilotDeviceIdentityApi.DeviceManagementGetWindowsAutopilotDeviceIdentities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceManagementGetWindowsAutopilotDeviceIdentities`: MicrosoftGraphWindowsAutopilotDeviceIdentity
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementWindowsAutopilotDeviceIdentityApi.DeviceManagementGetWindowsAutopilotDeviceIdentities`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**windowsAutopilotDeviceIdentityId** | **string** | key: id of windowsAutopilotDeviceIdentity | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementGetWindowsAutopilotDeviceIdentitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphWindowsAutopilotDeviceIdentity**](MicrosoftGraphWindowsAutopilotDeviceIdentity.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementListWindowsAutopilotDeviceIdentities

> CollectionOfWindowsAutopilotDeviceIdentity DeviceManagementListWindowsAutopilotDeviceIdentities(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get windowsAutopilotDeviceIdentities from deviceManagement



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
    resp, r, err := api_client.DeviceManagementWindowsAutopilotDeviceIdentityApi.DeviceManagementListWindowsAutopilotDeviceIdentities(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementWindowsAutopilotDeviceIdentityApi.DeviceManagementListWindowsAutopilotDeviceIdentities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceManagementListWindowsAutopilotDeviceIdentities`: CollectionOfWindowsAutopilotDeviceIdentity
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementWindowsAutopilotDeviceIdentityApi.DeviceManagementListWindowsAutopilotDeviceIdentities`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementListWindowsAutopilotDeviceIdentitiesRequest struct via the builder pattern


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

[**CollectionOfWindowsAutopilotDeviceIdentity**](CollectionOfWindowsAutopilotDeviceIdentity.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementUpdateWindowsAutopilotDeviceIdentities

> DeviceManagementUpdateWindowsAutopilotDeviceIdentities(ctx, windowsAutopilotDeviceIdentityId).MicrosoftGraphWindowsAutopilotDeviceIdentity(microsoftGraphWindowsAutopilotDeviceIdentity).Execute()

Update the navigation property windowsAutopilotDeviceIdentities in deviceManagement



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
    windowsAutopilotDeviceIdentityId := "windowsAutopilotDeviceIdentityId_example" // string | key: id of windowsAutopilotDeviceIdentity
    microsoftGraphWindowsAutopilotDeviceIdentity := *openapiclient.NewMicrosoftGraphWindowsAutopilotDeviceIdentity() // MicrosoftGraphWindowsAutopilotDeviceIdentity | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementWindowsAutopilotDeviceIdentityApi.DeviceManagementUpdateWindowsAutopilotDeviceIdentities(context.Background(), windowsAutopilotDeviceIdentityId).MicrosoftGraphWindowsAutopilotDeviceIdentity(microsoftGraphWindowsAutopilotDeviceIdentity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementWindowsAutopilotDeviceIdentityApi.DeviceManagementUpdateWindowsAutopilotDeviceIdentities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**windowsAutopilotDeviceIdentityId** | **string** | key: id of windowsAutopilotDeviceIdentity | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementUpdateWindowsAutopilotDeviceIdentitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphWindowsAutopilotDeviceIdentity** | [**MicrosoftGraphWindowsAutopilotDeviceIdentity**](MicrosoftGraphWindowsAutopilotDeviceIdentity.md) | New navigation property values | 

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

