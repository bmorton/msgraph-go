# \DeviceManagementDeviceEnrollmentConfigurationApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceManagementCreateDeviceEnrollmentConfigurations**](DeviceManagementDeviceEnrollmentConfigurationApi.md#DeviceManagementCreateDeviceEnrollmentConfigurations) | **Post** /deviceManagement/deviceEnrollmentConfigurations | Create new navigation property to deviceEnrollmentConfigurations for deviceManagement
[**DeviceManagementDeleteDeviceEnrollmentConfigurations**](DeviceManagementDeviceEnrollmentConfigurationApi.md#DeviceManagementDeleteDeviceEnrollmentConfigurations) | **Delete** /deviceManagement/deviceEnrollmentConfigurations/{deviceEnrollmentConfiguration-id} | Delete navigation property deviceEnrollmentConfigurations for deviceManagement
[**DeviceManagementDeviceEnrollmentConfigurationsCreateAssignments**](DeviceManagementDeviceEnrollmentConfigurationApi.md#DeviceManagementDeviceEnrollmentConfigurationsCreateAssignments) | **Post** /deviceManagement/deviceEnrollmentConfigurations/{deviceEnrollmentConfiguration-id}/assignments | Create new navigation property to assignments for deviceManagement
[**DeviceManagementDeviceEnrollmentConfigurationsDeleteAssignments**](DeviceManagementDeviceEnrollmentConfigurationApi.md#DeviceManagementDeviceEnrollmentConfigurationsDeleteAssignments) | **Delete** /deviceManagement/deviceEnrollmentConfigurations/{deviceEnrollmentConfiguration-id}/assignments/{enrollmentConfigurationAssignment-id} | Delete navigation property assignments for deviceManagement
[**DeviceManagementDeviceEnrollmentConfigurationsGetAssignments**](DeviceManagementDeviceEnrollmentConfigurationApi.md#DeviceManagementDeviceEnrollmentConfigurationsGetAssignments) | **Get** /deviceManagement/deviceEnrollmentConfigurations/{deviceEnrollmentConfiguration-id}/assignments/{enrollmentConfigurationAssignment-id} | Get assignments from deviceManagement
[**DeviceManagementDeviceEnrollmentConfigurationsListAssignments**](DeviceManagementDeviceEnrollmentConfigurationApi.md#DeviceManagementDeviceEnrollmentConfigurationsListAssignments) | **Get** /deviceManagement/deviceEnrollmentConfigurations/{deviceEnrollmentConfiguration-id}/assignments | Get assignments from deviceManagement
[**DeviceManagementDeviceEnrollmentConfigurationsUpdateAssignments**](DeviceManagementDeviceEnrollmentConfigurationApi.md#DeviceManagementDeviceEnrollmentConfigurationsUpdateAssignments) | **Patch** /deviceManagement/deviceEnrollmentConfigurations/{deviceEnrollmentConfiguration-id}/assignments/{enrollmentConfigurationAssignment-id} | Update the navigation property assignments in deviceManagement
[**DeviceManagementGetDeviceEnrollmentConfigurations**](DeviceManagementDeviceEnrollmentConfigurationApi.md#DeviceManagementGetDeviceEnrollmentConfigurations) | **Get** /deviceManagement/deviceEnrollmentConfigurations/{deviceEnrollmentConfiguration-id} | Get deviceEnrollmentConfigurations from deviceManagement
[**DeviceManagementListDeviceEnrollmentConfigurations**](DeviceManagementDeviceEnrollmentConfigurationApi.md#DeviceManagementListDeviceEnrollmentConfigurations) | **Get** /deviceManagement/deviceEnrollmentConfigurations | Get deviceEnrollmentConfigurations from deviceManagement
[**DeviceManagementUpdateDeviceEnrollmentConfigurations**](DeviceManagementDeviceEnrollmentConfigurationApi.md#DeviceManagementUpdateDeviceEnrollmentConfigurations) | **Patch** /deviceManagement/deviceEnrollmentConfigurations/{deviceEnrollmentConfiguration-id} | Update the navigation property deviceEnrollmentConfigurations in deviceManagement



## DeviceManagementCreateDeviceEnrollmentConfigurations

> MicrosoftGraphDeviceEnrollmentConfiguration DeviceManagementCreateDeviceEnrollmentConfigurations(ctx).MicrosoftGraphDeviceEnrollmentConfiguration(microsoftGraphDeviceEnrollmentConfiguration).Execute()

Create new navigation property to deviceEnrollmentConfigurations for deviceManagement



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
    microsoftGraphDeviceEnrollmentConfiguration := *openapiclient.NewMicrosoftGraphDeviceEnrollmentConfiguration() // MicrosoftGraphDeviceEnrollmentConfiguration | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementDeviceEnrollmentConfigurationApi.DeviceManagementCreateDeviceEnrollmentConfigurations(context.Background()).MicrosoftGraphDeviceEnrollmentConfiguration(microsoftGraphDeviceEnrollmentConfiguration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementDeviceEnrollmentConfigurationApi.DeviceManagementCreateDeviceEnrollmentConfigurations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceManagementCreateDeviceEnrollmentConfigurations`: MicrosoftGraphDeviceEnrollmentConfiguration
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementDeviceEnrollmentConfigurationApi.DeviceManagementCreateDeviceEnrollmentConfigurations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementCreateDeviceEnrollmentConfigurationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphDeviceEnrollmentConfiguration** | [**MicrosoftGraphDeviceEnrollmentConfiguration**](MicrosoftGraphDeviceEnrollmentConfiguration.md) | New navigation property | 

### Return type

[**MicrosoftGraphDeviceEnrollmentConfiguration**](MicrosoftGraphDeviceEnrollmentConfiguration.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementDeleteDeviceEnrollmentConfigurations

> DeviceManagementDeleteDeviceEnrollmentConfigurations(ctx, deviceEnrollmentConfigurationId).IfMatch(ifMatch).Execute()

Delete navigation property deviceEnrollmentConfigurations for deviceManagement



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
    deviceEnrollmentConfigurationId := "deviceEnrollmentConfigurationId_example" // string | key: id of deviceEnrollmentConfiguration
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementDeviceEnrollmentConfigurationApi.DeviceManagementDeleteDeviceEnrollmentConfigurations(context.Background(), deviceEnrollmentConfigurationId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementDeviceEnrollmentConfigurationApi.DeviceManagementDeleteDeviceEnrollmentConfigurations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceEnrollmentConfigurationId** | **string** | key: id of deviceEnrollmentConfiguration | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementDeleteDeviceEnrollmentConfigurationsRequest struct via the builder pattern


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


## DeviceManagementDeviceEnrollmentConfigurationsCreateAssignments

> MicrosoftGraphEnrollmentConfigurationAssignment DeviceManagementDeviceEnrollmentConfigurationsCreateAssignments(ctx, deviceEnrollmentConfigurationId).MicrosoftGraphEnrollmentConfigurationAssignment(microsoftGraphEnrollmentConfigurationAssignment).Execute()

Create new navigation property to assignments for deviceManagement



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
    deviceEnrollmentConfigurationId := "deviceEnrollmentConfigurationId_example" // string | key: id of deviceEnrollmentConfiguration
    microsoftGraphEnrollmentConfigurationAssignment := *openapiclient.NewMicrosoftGraphEnrollmentConfigurationAssignment() // MicrosoftGraphEnrollmentConfigurationAssignment | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementDeviceEnrollmentConfigurationApi.DeviceManagementDeviceEnrollmentConfigurationsCreateAssignments(context.Background(), deviceEnrollmentConfigurationId).MicrosoftGraphEnrollmentConfigurationAssignment(microsoftGraphEnrollmentConfigurationAssignment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementDeviceEnrollmentConfigurationApi.DeviceManagementDeviceEnrollmentConfigurationsCreateAssignments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceManagementDeviceEnrollmentConfigurationsCreateAssignments`: MicrosoftGraphEnrollmentConfigurationAssignment
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementDeviceEnrollmentConfigurationApi.DeviceManagementDeviceEnrollmentConfigurationsCreateAssignments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceEnrollmentConfigurationId** | **string** | key: id of deviceEnrollmentConfiguration | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementDeviceEnrollmentConfigurationsCreateAssignmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphEnrollmentConfigurationAssignment** | [**MicrosoftGraphEnrollmentConfigurationAssignment**](MicrosoftGraphEnrollmentConfigurationAssignment.md) | New navigation property | 

### Return type

[**MicrosoftGraphEnrollmentConfigurationAssignment**](MicrosoftGraphEnrollmentConfigurationAssignment.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementDeviceEnrollmentConfigurationsDeleteAssignments

> DeviceManagementDeviceEnrollmentConfigurationsDeleteAssignments(ctx, deviceEnrollmentConfigurationId, enrollmentConfigurationAssignmentId).IfMatch(ifMatch).Execute()

Delete navigation property assignments for deviceManagement



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
    deviceEnrollmentConfigurationId := "deviceEnrollmentConfigurationId_example" // string | key: id of deviceEnrollmentConfiguration
    enrollmentConfigurationAssignmentId := "enrollmentConfigurationAssignmentId_example" // string | key: id of enrollmentConfigurationAssignment
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementDeviceEnrollmentConfigurationApi.DeviceManagementDeviceEnrollmentConfigurationsDeleteAssignments(context.Background(), deviceEnrollmentConfigurationId, enrollmentConfigurationAssignmentId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementDeviceEnrollmentConfigurationApi.DeviceManagementDeviceEnrollmentConfigurationsDeleteAssignments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceEnrollmentConfigurationId** | **string** | key: id of deviceEnrollmentConfiguration | 
**enrollmentConfigurationAssignmentId** | **string** | key: id of enrollmentConfigurationAssignment | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementDeviceEnrollmentConfigurationsDeleteAssignmentsRequest struct via the builder pattern


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


## DeviceManagementDeviceEnrollmentConfigurationsGetAssignments

> MicrosoftGraphEnrollmentConfigurationAssignment DeviceManagementDeviceEnrollmentConfigurationsGetAssignments(ctx, deviceEnrollmentConfigurationId, enrollmentConfigurationAssignmentId).Select_(select_).Expand(expand).Execute()

Get assignments from deviceManagement



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
    deviceEnrollmentConfigurationId := "deviceEnrollmentConfigurationId_example" // string | key: id of deviceEnrollmentConfiguration
    enrollmentConfigurationAssignmentId := "enrollmentConfigurationAssignmentId_example" // string | key: id of enrollmentConfigurationAssignment
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementDeviceEnrollmentConfigurationApi.DeviceManagementDeviceEnrollmentConfigurationsGetAssignments(context.Background(), deviceEnrollmentConfigurationId, enrollmentConfigurationAssignmentId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementDeviceEnrollmentConfigurationApi.DeviceManagementDeviceEnrollmentConfigurationsGetAssignments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceManagementDeviceEnrollmentConfigurationsGetAssignments`: MicrosoftGraphEnrollmentConfigurationAssignment
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementDeviceEnrollmentConfigurationApi.DeviceManagementDeviceEnrollmentConfigurationsGetAssignments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceEnrollmentConfigurationId** | **string** | key: id of deviceEnrollmentConfiguration | 
**enrollmentConfigurationAssignmentId** | **string** | key: id of enrollmentConfigurationAssignment | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementDeviceEnrollmentConfigurationsGetAssignmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphEnrollmentConfigurationAssignment**](MicrosoftGraphEnrollmentConfigurationAssignment.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementDeviceEnrollmentConfigurationsListAssignments

> CollectionOfEnrollmentConfigurationAssignment DeviceManagementDeviceEnrollmentConfigurationsListAssignments(ctx, deviceEnrollmentConfigurationId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get assignments from deviceManagement



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
    deviceEnrollmentConfigurationId := "deviceEnrollmentConfigurationId_example" // string | key: id of deviceEnrollmentConfiguration
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
    resp, r, err := api_client.DeviceManagementDeviceEnrollmentConfigurationApi.DeviceManagementDeviceEnrollmentConfigurationsListAssignments(context.Background(), deviceEnrollmentConfigurationId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementDeviceEnrollmentConfigurationApi.DeviceManagementDeviceEnrollmentConfigurationsListAssignments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceManagementDeviceEnrollmentConfigurationsListAssignments`: CollectionOfEnrollmentConfigurationAssignment
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementDeviceEnrollmentConfigurationApi.DeviceManagementDeviceEnrollmentConfigurationsListAssignments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceEnrollmentConfigurationId** | **string** | key: id of deviceEnrollmentConfiguration | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementDeviceEnrollmentConfigurationsListAssignmentsRequest struct via the builder pattern


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

[**CollectionOfEnrollmentConfigurationAssignment**](CollectionOfEnrollmentConfigurationAssignment.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementDeviceEnrollmentConfigurationsUpdateAssignments

> DeviceManagementDeviceEnrollmentConfigurationsUpdateAssignments(ctx, deviceEnrollmentConfigurationId, enrollmentConfigurationAssignmentId).MicrosoftGraphEnrollmentConfigurationAssignment(microsoftGraphEnrollmentConfigurationAssignment).Execute()

Update the navigation property assignments in deviceManagement



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
    deviceEnrollmentConfigurationId := "deviceEnrollmentConfigurationId_example" // string | key: id of deviceEnrollmentConfiguration
    enrollmentConfigurationAssignmentId := "enrollmentConfigurationAssignmentId_example" // string | key: id of enrollmentConfigurationAssignment
    microsoftGraphEnrollmentConfigurationAssignment := *openapiclient.NewMicrosoftGraphEnrollmentConfigurationAssignment() // MicrosoftGraphEnrollmentConfigurationAssignment | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementDeviceEnrollmentConfigurationApi.DeviceManagementDeviceEnrollmentConfigurationsUpdateAssignments(context.Background(), deviceEnrollmentConfigurationId, enrollmentConfigurationAssignmentId).MicrosoftGraphEnrollmentConfigurationAssignment(microsoftGraphEnrollmentConfigurationAssignment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementDeviceEnrollmentConfigurationApi.DeviceManagementDeviceEnrollmentConfigurationsUpdateAssignments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceEnrollmentConfigurationId** | **string** | key: id of deviceEnrollmentConfiguration | 
**enrollmentConfigurationAssignmentId** | **string** | key: id of enrollmentConfigurationAssignment | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementDeviceEnrollmentConfigurationsUpdateAssignmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **microsoftGraphEnrollmentConfigurationAssignment** | [**MicrosoftGraphEnrollmentConfigurationAssignment**](MicrosoftGraphEnrollmentConfigurationAssignment.md) | New navigation property values | 

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


## DeviceManagementGetDeviceEnrollmentConfigurations

> MicrosoftGraphDeviceEnrollmentConfiguration DeviceManagementGetDeviceEnrollmentConfigurations(ctx, deviceEnrollmentConfigurationId).Select_(select_).Expand(expand).Execute()

Get deviceEnrollmentConfigurations from deviceManagement



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
    deviceEnrollmentConfigurationId := "deviceEnrollmentConfigurationId_example" // string | key: id of deviceEnrollmentConfiguration
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementDeviceEnrollmentConfigurationApi.DeviceManagementGetDeviceEnrollmentConfigurations(context.Background(), deviceEnrollmentConfigurationId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementDeviceEnrollmentConfigurationApi.DeviceManagementGetDeviceEnrollmentConfigurations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceManagementGetDeviceEnrollmentConfigurations`: MicrosoftGraphDeviceEnrollmentConfiguration
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementDeviceEnrollmentConfigurationApi.DeviceManagementGetDeviceEnrollmentConfigurations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceEnrollmentConfigurationId** | **string** | key: id of deviceEnrollmentConfiguration | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementGetDeviceEnrollmentConfigurationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphDeviceEnrollmentConfiguration**](MicrosoftGraphDeviceEnrollmentConfiguration.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementListDeviceEnrollmentConfigurations

> CollectionOfDeviceEnrollmentConfiguration DeviceManagementListDeviceEnrollmentConfigurations(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get deviceEnrollmentConfigurations from deviceManagement



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
    resp, r, err := api_client.DeviceManagementDeviceEnrollmentConfigurationApi.DeviceManagementListDeviceEnrollmentConfigurations(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementDeviceEnrollmentConfigurationApi.DeviceManagementListDeviceEnrollmentConfigurations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceManagementListDeviceEnrollmentConfigurations`: CollectionOfDeviceEnrollmentConfiguration
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementDeviceEnrollmentConfigurationApi.DeviceManagementListDeviceEnrollmentConfigurations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementListDeviceEnrollmentConfigurationsRequest struct via the builder pattern


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

[**CollectionOfDeviceEnrollmentConfiguration**](CollectionOfDeviceEnrollmentConfiguration.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementUpdateDeviceEnrollmentConfigurations

> DeviceManagementUpdateDeviceEnrollmentConfigurations(ctx, deviceEnrollmentConfigurationId).MicrosoftGraphDeviceEnrollmentConfiguration(microsoftGraphDeviceEnrollmentConfiguration).Execute()

Update the navigation property deviceEnrollmentConfigurations in deviceManagement



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
    deviceEnrollmentConfigurationId := "deviceEnrollmentConfigurationId_example" // string | key: id of deviceEnrollmentConfiguration
    microsoftGraphDeviceEnrollmentConfiguration := *openapiclient.NewMicrosoftGraphDeviceEnrollmentConfiguration() // MicrosoftGraphDeviceEnrollmentConfiguration | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementDeviceEnrollmentConfigurationApi.DeviceManagementUpdateDeviceEnrollmentConfigurations(context.Background(), deviceEnrollmentConfigurationId).MicrosoftGraphDeviceEnrollmentConfiguration(microsoftGraphDeviceEnrollmentConfiguration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementDeviceEnrollmentConfigurationApi.DeviceManagementUpdateDeviceEnrollmentConfigurations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceEnrollmentConfigurationId** | **string** | key: id of deviceEnrollmentConfiguration | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementUpdateDeviceEnrollmentConfigurationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphDeviceEnrollmentConfiguration** | [**MicrosoftGraphDeviceEnrollmentConfiguration**](MicrosoftGraphDeviceEnrollmentConfiguration.md) | New navigation property values | 

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

