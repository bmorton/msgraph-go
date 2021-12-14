# \DeviceManagementDeviceCategoryApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceManagementCreateDeviceCategories**](DeviceManagementDeviceCategoryApi.md#DeviceManagementCreateDeviceCategories) | **Post** /deviceManagement/deviceCategories | Create new navigation property to deviceCategories for deviceManagement
[**DeviceManagementDeleteDeviceCategories**](DeviceManagementDeviceCategoryApi.md#DeviceManagementDeleteDeviceCategories) | **Delete** /deviceManagement/deviceCategories/{deviceCategory-id} | Delete navigation property deviceCategories for deviceManagement
[**DeviceManagementGetDeviceCategories**](DeviceManagementDeviceCategoryApi.md#DeviceManagementGetDeviceCategories) | **Get** /deviceManagement/deviceCategories/{deviceCategory-id} | Get deviceCategories from deviceManagement
[**DeviceManagementListDeviceCategories**](DeviceManagementDeviceCategoryApi.md#DeviceManagementListDeviceCategories) | **Get** /deviceManagement/deviceCategories | Get deviceCategories from deviceManagement
[**DeviceManagementUpdateDeviceCategories**](DeviceManagementDeviceCategoryApi.md#DeviceManagementUpdateDeviceCategories) | **Patch** /deviceManagement/deviceCategories/{deviceCategory-id} | Update the navigation property deviceCategories in deviceManagement



## DeviceManagementCreateDeviceCategories

> MicrosoftGraphDeviceCategory DeviceManagementCreateDeviceCategories(ctx).MicrosoftGraphDeviceCategory(microsoftGraphDeviceCategory).Execute()

Create new navigation property to deviceCategories for deviceManagement



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
    microsoftGraphDeviceCategory := *openapiclient.NewMicrosoftGraphDeviceCategory() // MicrosoftGraphDeviceCategory | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementDeviceCategoryApi.DeviceManagementCreateDeviceCategories(context.Background()).MicrosoftGraphDeviceCategory(microsoftGraphDeviceCategory).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementDeviceCategoryApi.DeviceManagementCreateDeviceCategories``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceManagementCreateDeviceCategories`: MicrosoftGraphDeviceCategory
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementDeviceCategoryApi.DeviceManagementCreateDeviceCategories`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementCreateDeviceCategoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphDeviceCategory** | [**MicrosoftGraphDeviceCategory**](MicrosoftGraphDeviceCategory.md) | New navigation property | 

### Return type

[**MicrosoftGraphDeviceCategory**](MicrosoftGraphDeviceCategory.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementDeleteDeviceCategories

> DeviceManagementDeleteDeviceCategories(ctx, deviceCategoryId).IfMatch(ifMatch).Execute()

Delete navigation property deviceCategories for deviceManagement



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
    deviceCategoryId := "deviceCategoryId_example" // string | key: id of deviceCategory
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementDeviceCategoryApi.DeviceManagementDeleteDeviceCategories(context.Background(), deviceCategoryId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementDeviceCategoryApi.DeviceManagementDeleteDeviceCategories``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceCategoryId** | **string** | key: id of deviceCategory | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementDeleteDeviceCategoriesRequest struct via the builder pattern


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


## DeviceManagementGetDeviceCategories

> MicrosoftGraphDeviceCategory DeviceManagementGetDeviceCategories(ctx, deviceCategoryId).Select_(select_).Expand(expand).Execute()

Get deviceCategories from deviceManagement



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
    deviceCategoryId := "deviceCategoryId_example" // string | key: id of deviceCategory
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementDeviceCategoryApi.DeviceManagementGetDeviceCategories(context.Background(), deviceCategoryId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementDeviceCategoryApi.DeviceManagementGetDeviceCategories``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceManagementGetDeviceCategories`: MicrosoftGraphDeviceCategory
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementDeviceCategoryApi.DeviceManagementGetDeviceCategories`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceCategoryId** | **string** | key: id of deviceCategory | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementGetDeviceCategoriesRequest struct via the builder pattern


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


## DeviceManagementListDeviceCategories

> CollectionOfDeviceCategory DeviceManagementListDeviceCategories(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get deviceCategories from deviceManagement



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
    resp, r, err := api_client.DeviceManagementDeviceCategoryApi.DeviceManagementListDeviceCategories(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementDeviceCategoryApi.DeviceManagementListDeviceCategories``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceManagementListDeviceCategories`: CollectionOfDeviceCategory
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementDeviceCategoryApi.DeviceManagementListDeviceCategories`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementListDeviceCategoriesRequest struct via the builder pattern


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

[**CollectionOfDeviceCategory**](CollectionOfDeviceCategory.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementUpdateDeviceCategories

> DeviceManagementUpdateDeviceCategories(ctx, deviceCategoryId).MicrosoftGraphDeviceCategory(microsoftGraphDeviceCategory).Execute()

Update the navigation property deviceCategories in deviceManagement



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
    deviceCategoryId := "deviceCategoryId_example" // string | key: id of deviceCategory
    microsoftGraphDeviceCategory := *openapiclient.NewMicrosoftGraphDeviceCategory() // MicrosoftGraphDeviceCategory | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementDeviceCategoryApi.DeviceManagementUpdateDeviceCategories(context.Background(), deviceCategoryId).MicrosoftGraphDeviceCategory(microsoftGraphDeviceCategory).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementDeviceCategoryApi.DeviceManagementUpdateDeviceCategories``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceCategoryId** | **string** | key: id of deviceCategory | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementUpdateDeviceCategoriesRequest struct via the builder pattern


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

