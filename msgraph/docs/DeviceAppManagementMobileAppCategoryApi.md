# \DeviceAppManagementMobileAppCategoryApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceAppManagementCreateMobileAppCategories**](DeviceAppManagementMobileAppCategoryApi.md#DeviceAppManagementCreateMobileAppCategories) | **Post** /deviceAppManagement/mobileAppCategories | Create new navigation property to mobileAppCategories for deviceAppManagement
[**DeviceAppManagementDeleteMobileAppCategories**](DeviceAppManagementMobileAppCategoryApi.md#DeviceAppManagementDeleteMobileAppCategories) | **Delete** /deviceAppManagement/mobileAppCategories/{mobileAppCategory-id} | Delete navigation property mobileAppCategories for deviceAppManagement
[**DeviceAppManagementGetMobileAppCategories**](DeviceAppManagementMobileAppCategoryApi.md#DeviceAppManagementGetMobileAppCategories) | **Get** /deviceAppManagement/mobileAppCategories/{mobileAppCategory-id} | Get mobileAppCategories from deviceAppManagement
[**DeviceAppManagementListMobileAppCategories**](DeviceAppManagementMobileAppCategoryApi.md#DeviceAppManagementListMobileAppCategories) | **Get** /deviceAppManagement/mobileAppCategories | Get mobileAppCategories from deviceAppManagement
[**DeviceAppManagementUpdateMobileAppCategories**](DeviceAppManagementMobileAppCategoryApi.md#DeviceAppManagementUpdateMobileAppCategories) | **Patch** /deviceAppManagement/mobileAppCategories/{mobileAppCategory-id} | Update the navigation property mobileAppCategories in deviceAppManagement



## DeviceAppManagementCreateMobileAppCategories

> MicrosoftGraphMobileAppCategory DeviceAppManagementCreateMobileAppCategories(ctx).MicrosoftGraphMobileAppCategory(microsoftGraphMobileAppCategory).Execute()

Create new navigation property to mobileAppCategories for deviceAppManagement



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
    microsoftGraphMobileAppCategory := *openapiclient.NewMicrosoftGraphMobileAppCategory() // MicrosoftGraphMobileAppCategory | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceAppManagementMobileAppCategoryApi.DeviceAppManagementCreateMobileAppCategories(context.Background()).MicrosoftGraphMobileAppCategory(microsoftGraphMobileAppCategory).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAppManagementMobileAppCategoryApi.DeviceAppManagementCreateMobileAppCategories``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceAppManagementCreateMobileAppCategories`: MicrosoftGraphMobileAppCategory
    fmt.Fprintf(os.Stdout, "Response from `DeviceAppManagementMobileAppCategoryApi.DeviceAppManagementCreateMobileAppCategories`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeviceAppManagementCreateMobileAppCategoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphMobileAppCategory** | [**MicrosoftGraphMobileAppCategory**](MicrosoftGraphMobileAppCategory.md) | New navigation property | 

### Return type

[**MicrosoftGraphMobileAppCategory**](MicrosoftGraphMobileAppCategory.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementDeleteMobileAppCategories

> DeviceAppManagementDeleteMobileAppCategories(ctx, mobileAppCategoryId).IfMatch(ifMatch).Execute()

Delete navigation property mobileAppCategories for deviceAppManagement



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
    mobileAppCategoryId := "mobileAppCategoryId_example" // string | key: id of mobileAppCategory
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceAppManagementMobileAppCategoryApi.DeviceAppManagementDeleteMobileAppCategories(context.Background(), mobileAppCategoryId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAppManagementMobileAppCategoryApi.DeviceAppManagementDeleteMobileAppCategories``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mobileAppCategoryId** | **string** | key: id of mobileAppCategory | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceAppManagementDeleteMobileAppCategoriesRequest struct via the builder pattern


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


## DeviceAppManagementGetMobileAppCategories

> MicrosoftGraphMobileAppCategory DeviceAppManagementGetMobileAppCategories(ctx, mobileAppCategoryId).Select_(select_).Expand(expand).Execute()

Get mobileAppCategories from deviceAppManagement



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
    mobileAppCategoryId := "mobileAppCategoryId_example" // string | key: id of mobileAppCategory
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceAppManagementMobileAppCategoryApi.DeviceAppManagementGetMobileAppCategories(context.Background(), mobileAppCategoryId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAppManagementMobileAppCategoryApi.DeviceAppManagementGetMobileAppCategories``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceAppManagementGetMobileAppCategories`: MicrosoftGraphMobileAppCategory
    fmt.Fprintf(os.Stdout, "Response from `DeviceAppManagementMobileAppCategoryApi.DeviceAppManagementGetMobileAppCategories`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mobileAppCategoryId** | **string** | key: id of mobileAppCategory | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceAppManagementGetMobileAppCategoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphMobileAppCategory**](MicrosoftGraphMobileAppCategory.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementListMobileAppCategories

> CollectionOfMobileAppCategory DeviceAppManagementListMobileAppCategories(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get mobileAppCategories from deviceAppManagement



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
    resp, r, err := api_client.DeviceAppManagementMobileAppCategoryApi.DeviceAppManagementListMobileAppCategories(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAppManagementMobileAppCategoryApi.DeviceAppManagementListMobileAppCategories``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceAppManagementListMobileAppCategories`: CollectionOfMobileAppCategory
    fmt.Fprintf(os.Stdout, "Response from `DeviceAppManagementMobileAppCategoryApi.DeviceAppManagementListMobileAppCategories`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeviceAppManagementListMobileAppCategoriesRequest struct via the builder pattern


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

[**CollectionOfMobileAppCategory**](CollectionOfMobileAppCategory.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementUpdateMobileAppCategories

> DeviceAppManagementUpdateMobileAppCategories(ctx, mobileAppCategoryId).MicrosoftGraphMobileAppCategory(microsoftGraphMobileAppCategory).Execute()

Update the navigation property mobileAppCategories in deviceAppManagement



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
    mobileAppCategoryId := "mobileAppCategoryId_example" // string | key: id of mobileAppCategory
    microsoftGraphMobileAppCategory := *openapiclient.NewMicrosoftGraphMobileAppCategory() // MicrosoftGraphMobileAppCategory | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceAppManagementMobileAppCategoryApi.DeviceAppManagementUpdateMobileAppCategories(context.Background(), mobileAppCategoryId).MicrosoftGraphMobileAppCategory(microsoftGraphMobileAppCategory).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAppManagementMobileAppCategoryApi.DeviceAppManagementUpdateMobileAppCategories``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mobileAppCategoryId** | **string** | key: id of mobileAppCategory | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceAppManagementUpdateMobileAppCategoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphMobileAppCategory** | [**MicrosoftGraphMobileAppCategory**](MicrosoftGraphMobileAppCategory.md) | New navigation property values | 

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

