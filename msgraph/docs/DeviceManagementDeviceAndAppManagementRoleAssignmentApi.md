# \DeviceManagementDeviceAndAppManagementRoleAssignmentApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceManagementCreateRoleAssignments**](DeviceManagementDeviceAndAppManagementRoleAssignmentApi.md#DeviceManagementCreateRoleAssignments) | **Post** /deviceManagement/roleAssignments | Create new navigation property to roleAssignments for deviceManagement
[**DeviceManagementDeleteRoleAssignments**](DeviceManagementDeviceAndAppManagementRoleAssignmentApi.md#DeviceManagementDeleteRoleAssignments) | **Delete** /deviceManagement/roleAssignments/{deviceAndAppManagementRoleAssignment-id} | Delete navigation property roleAssignments for deviceManagement
[**DeviceManagementGetRoleAssignments**](DeviceManagementDeviceAndAppManagementRoleAssignmentApi.md#DeviceManagementGetRoleAssignments) | **Get** /deviceManagement/roleAssignments/{deviceAndAppManagementRoleAssignment-id} | Get roleAssignments from deviceManagement
[**DeviceManagementListRoleAssignments**](DeviceManagementDeviceAndAppManagementRoleAssignmentApi.md#DeviceManagementListRoleAssignments) | **Get** /deviceManagement/roleAssignments | Get roleAssignments from deviceManagement
[**DeviceManagementUpdateRoleAssignments**](DeviceManagementDeviceAndAppManagementRoleAssignmentApi.md#DeviceManagementUpdateRoleAssignments) | **Patch** /deviceManagement/roleAssignments/{deviceAndAppManagementRoleAssignment-id} | Update the navigation property roleAssignments in deviceManagement



## DeviceManagementCreateRoleAssignments

> MicrosoftGraphDeviceAndAppManagementRoleAssignment DeviceManagementCreateRoleAssignments(ctx).MicrosoftGraphDeviceAndAppManagementRoleAssignment(microsoftGraphDeviceAndAppManagementRoleAssignment).Execute()

Create new navigation property to roleAssignments for deviceManagement



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
    microsoftGraphDeviceAndAppManagementRoleAssignment := *openapiclient.NewMicrosoftGraphDeviceAndAppManagementRoleAssignment() // MicrosoftGraphDeviceAndAppManagementRoleAssignment | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementDeviceAndAppManagementRoleAssignmentApi.DeviceManagementCreateRoleAssignments(context.Background()).MicrosoftGraphDeviceAndAppManagementRoleAssignment(microsoftGraphDeviceAndAppManagementRoleAssignment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementDeviceAndAppManagementRoleAssignmentApi.DeviceManagementCreateRoleAssignments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceManagementCreateRoleAssignments`: MicrosoftGraphDeviceAndAppManagementRoleAssignment
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementDeviceAndAppManagementRoleAssignmentApi.DeviceManagementCreateRoleAssignments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementCreateRoleAssignmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphDeviceAndAppManagementRoleAssignment** | [**MicrosoftGraphDeviceAndAppManagementRoleAssignment**](MicrosoftGraphDeviceAndAppManagementRoleAssignment.md) | New navigation property | 

### Return type

[**MicrosoftGraphDeviceAndAppManagementRoleAssignment**](MicrosoftGraphDeviceAndAppManagementRoleAssignment.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementDeleteRoleAssignments

> DeviceManagementDeleteRoleAssignments(ctx, deviceAndAppManagementRoleAssignmentId).IfMatch(ifMatch).Execute()

Delete navigation property roleAssignments for deviceManagement



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
    deviceAndAppManagementRoleAssignmentId := "deviceAndAppManagementRoleAssignmentId_example" // string | key: id of deviceAndAppManagementRoleAssignment
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementDeviceAndAppManagementRoleAssignmentApi.DeviceManagementDeleteRoleAssignments(context.Background(), deviceAndAppManagementRoleAssignmentId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementDeviceAndAppManagementRoleAssignmentApi.DeviceManagementDeleteRoleAssignments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceAndAppManagementRoleAssignmentId** | **string** | key: id of deviceAndAppManagementRoleAssignment | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementDeleteRoleAssignmentsRequest struct via the builder pattern


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


## DeviceManagementGetRoleAssignments

> MicrosoftGraphDeviceAndAppManagementRoleAssignment DeviceManagementGetRoleAssignments(ctx, deviceAndAppManagementRoleAssignmentId).Select_(select_).Expand(expand).Execute()

Get roleAssignments from deviceManagement



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
    deviceAndAppManagementRoleAssignmentId := "deviceAndAppManagementRoleAssignmentId_example" // string | key: id of deviceAndAppManagementRoleAssignment
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementDeviceAndAppManagementRoleAssignmentApi.DeviceManagementGetRoleAssignments(context.Background(), deviceAndAppManagementRoleAssignmentId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementDeviceAndAppManagementRoleAssignmentApi.DeviceManagementGetRoleAssignments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceManagementGetRoleAssignments`: MicrosoftGraphDeviceAndAppManagementRoleAssignment
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementDeviceAndAppManagementRoleAssignmentApi.DeviceManagementGetRoleAssignments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceAndAppManagementRoleAssignmentId** | **string** | key: id of deviceAndAppManagementRoleAssignment | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementGetRoleAssignmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphDeviceAndAppManagementRoleAssignment**](MicrosoftGraphDeviceAndAppManagementRoleAssignment.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementListRoleAssignments

> CollectionOfDeviceAndAppManagementRoleAssignment DeviceManagementListRoleAssignments(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get roleAssignments from deviceManagement



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
    resp, r, err := api_client.DeviceManagementDeviceAndAppManagementRoleAssignmentApi.DeviceManagementListRoleAssignments(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementDeviceAndAppManagementRoleAssignmentApi.DeviceManagementListRoleAssignments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceManagementListRoleAssignments`: CollectionOfDeviceAndAppManagementRoleAssignment
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementDeviceAndAppManagementRoleAssignmentApi.DeviceManagementListRoleAssignments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementListRoleAssignmentsRequest struct via the builder pattern


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

[**CollectionOfDeviceAndAppManagementRoleAssignment**](CollectionOfDeviceAndAppManagementRoleAssignment.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementUpdateRoleAssignments

> DeviceManagementUpdateRoleAssignments(ctx, deviceAndAppManagementRoleAssignmentId).MicrosoftGraphDeviceAndAppManagementRoleAssignment(microsoftGraphDeviceAndAppManagementRoleAssignment).Execute()

Update the navigation property roleAssignments in deviceManagement



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
    deviceAndAppManagementRoleAssignmentId := "deviceAndAppManagementRoleAssignmentId_example" // string | key: id of deviceAndAppManagementRoleAssignment
    microsoftGraphDeviceAndAppManagementRoleAssignment := *openapiclient.NewMicrosoftGraphDeviceAndAppManagementRoleAssignment() // MicrosoftGraphDeviceAndAppManagementRoleAssignment | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementDeviceAndAppManagementRoleAssignmentApi.DeviceManagementUpdateRoleAssignments(context.Background(), deviceAndAppManagementRoleAssignmentId).MicrosoftGraphDeviceAndAppManagementRoleAssignment(microsoftGraphDeviceAndAppManagementRoleAssignment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementDeviceAndAppManagementRoleAssignmentApi.DeviceManagementUpdateRoleAssignments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceAndAppManagementRoleAssignmentId** | **string** | key: id of deviceAndAppManagementRoleAssignment | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementUpdateRoleAssignmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphDeviceAndAppManagementRoleAssignment** | [**MicrosoftGraphDeviceAndAppManagementRoleAssignment**](MicrosoftGraphDeviceAndAppManagementRoleAssignment.md) | New navigation property values | 

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

