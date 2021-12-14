# \DeviceAppManagementMobileAppApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceAppManagementCreateMobileApps**](DeviceAppManagementMobileAppApi.md#DeviceAppManagementCreateMobileApps) | **Post** /deviceAppManagement/mobileApps | Create new navigation property to mobileApps for deviceAppManagement
[**DeviceAppManagementDeleteMobileApps**](DeviceAppManagementMobileAppApi.md#DeviceAppManagementDeleteMobileApps) | **Delete** /deviceAppManagement/mobileApps/{mobileApp-id} | Delete navigation property mobileApps for deviceAppManagement
[**DeviceAppManagementGetMobileApps**](DeviceAppManagementMobileAppApi.md#DeviceAppManagementGetMobileApps) | **Get** /deviceAppManagement/mobileApps/{mobileApp-id} | Get mobileApps from deviceAppManagement
[**DeviceAppManagementListMobileApps**](DeviceAppManagementMobileAppApi.md#DeviceAppManagementListMobileApps) | **Get** /deviceAppManagement/mobileApps | Get mobileApps from deviceAppManagement
[**DeviceAppManagementMobileAppsCreateAssignments**](DeviceAppManagementMobileAppApi.md#DeviceAppManagementMobileAppsCreateAssignments) | **Post** /deviceAppManagement/mobileApps/{mobileApp-id}/assignments | Create new navigation property to assignments for deviceAppManagement
[**DeviceAppManagementMobileAppsCreateRefCategories**](DeviceAppManagementMobileAppApi.md#DeviceAppManagementMobileAppsCreateRefCategories) | **Post** /deviceAppManagement/mobileApps/{mobileApp-id}/categories/$ref | Create new navigation property ref to categories for deviceAppManagement
[**DeviceAppManagementMobileAppsDeleteAssignments**](DeviceAppManagementMobileAppApi.md#DeviceAppManagementMobileAppsDeleteAssignments) | **Delete** /deviceAppManagement/mobileApps/{mobileApp-id}/assignments/{mobileAppAssignment-id} | Delete navigation property assignments for deviceAppManagement
[**DeviceAppManagementMobileAppsGetAssignments**](DeviceAppManagementMobileAppApi.md#DeviceAppManagementMobileAppsGetAssignments) | **Get** /deviceAppManagement/mobileApps/{mobileApp-id}/assignments/{mobileAppAssignment-id} | Get assignments from deviceAppManagement
[**DeviceAppManagementMobileAppsListAssignments**](DeviceAppManagementMobileAppApi.md#DeviceAppManagementMobileAppsListAssignments) | **Get** /deviceAppManagement/mobileApps/{mobileApp-id}/assignments | Get assignments from deviceAppManagement
[**DeviceAppManagementMobileAppsListCategories**](DeviceAppManagementMobileAppApi.md#DeviceAppManagementMobileAppsListCategories) | **Get** /deviceAppManagement/mobileApps/{mobileApp-id}/categories | Get categories from deviceAppManagement
[**DeviceAppManagementMobileAppsListRefCategories**](DeviceAppManagementMobileAppApi.md#DeviceAppManagementMobileAppsListRefCategories) | **Get** /deviceAppManagement/mobileApps/{mobileApp-id}/categories/$ref | Get ref of categories from deviceAppManagement
[**DeviceAppManagementMobileAppsUpdateAssignments**](DeviceAppManagementMobileAppApi.md#DeviceAppManagementMobileAppsUpdateAssignments) | **Patch** /deviceAppManagement/mobileApps/{mobileApp-id}/assignments/{mobileAppAssignment-id} | Update the navigation property assignments in deviceAppManagement
[**DeviceAppManagementUpdateMobileApps**](DeviceAppManagementMobileAppApi.md#DeviceAppManagementUpdateMobileApps) | **Patch** /deviceAppManagement/mobileApps/{mobileApp-id} | Update the navigation property mobileApps in deviceAppManagement



## DeviceAppManagementCreateMobileApps

> MicrosoftGraphMobileApp DeviceAppManagementCreateMobileApps(ctx).MicrosoftGraphMobileApp(microsoftGraphMobileApp).Execute()

Create new navigation property to mobileApps for deviceAppManagement



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
    microsoftGraphMobileApp := *openapiclient.NewMicrosoftGraphMobileApp() // MicrosoftGraphMobileApp | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceAppManagementMobileAppApi.DeviceAppManagementCreateMobileApps(context.Background()).MicrosoftGraphMobileApp(microsoftGraphMobileApp).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAppManagementMobileAppApi.DeviceAppManagementCreateMobileApps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceAppManagementCreateMobileApps`: MicrosoftGraphMobileApp
    fmt.Fprintf(os.Stdout, "Response from `DeviceAppManagementMobileAppApi.DeviceAppManagementCreateMobileApps`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeviceAppManagementCreateMobileAppsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphMobileApp** | [**MicrosoftGraphMobileApp**](MicrosoftGraphMobileApp.md) | New navigation property | 

### Return type

[**MicrosoftGraphMobileApp**](MicrosoftGraphMobileApp.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementDeleteMobileApps

> DeviceAppManagementDeleteMobileApps(ctx, mobileAppId).IfMatch(ifMatch).Execute()

Delete navigation property mobileApps for deviceAppManagement



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
    mobileAppId := "mobileAppId_example" // string | key: id of mobileApp
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceAppManagementMobileAppApi.DeviceAppManagementDeleteMobileApps(context.Background(), mobileAppId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAppManagementMobileAppApi.DeviceAppManagementDeleteMobileApps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mobileAppId** | **string** | key: id of mobileApp | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceAppManagementDeleteMobileAppsRequest struct via the builder pattern


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


## DeviceAppManagementGetMobileApps

> MicrosoftGraphMobileApp DeviceAppManagementGetMobileApps(ctx, mobileAppId).Select_(select_).Expand(expand).Execute()

Get mobileApps from deviceAppManagement



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
    mobileAppId := "mobileAppId_example" // string | key: id of mobileApp
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceAppManagementMobileAppApi.DeviceAppManagementGetMobileApps(context.Background(), mobileAppId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAppManagementMobileAppApi.DeviceAppManagementGetMobileApps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceAppManagementGetMobileApps`: MicrosoftGraphMobileApp
    fmt.Fprintf(os.Stdout, "Response from `DeviceAppManagementMobileAppApi.DeviceAppManagementGetMobileApps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mobileAppId** | **string** | key: id of mobileApp | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceAppManagementGetMobileAppsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphMobileApp**](MicrosoftGraphMobileApp.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementListMobileApps

> CollectionOfMobileApp DeviceAppManagementListMobileApps(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get mobileApps from deviceAppManagement



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
    resp, r, err := api_client.DeviceAppManagementMobileAppApi.DeviceAppManagementListMobileApps(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAppManagementMobileAppApi.DeviceAppManagementListMobileApps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceAppManagementListMobileApps`: CollectionOfMobileApp
    fmt.Fprintf(os.Stdout, "Response from `DeviceAppManagementMobileAppApi.DeviceAppManagementListMobileApps`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeviceAppManagementListMobileAppsRequest struct via the builder pattern


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

[**CollectionOfMobileApp**](CollectionOfMobileApp.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementMobileAppsCreateAssignments

> MicrosoftGraphMobileAppAssignment DeviceAppManagementMobileAppsCreateAssignments(ctx, mobileAppId).MicrosoftGraphMobileAppAssignment(microsoftGraphMobileAppAssignment).Execute()

Create new navigation property to assignments for deviceAppManagement



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
    mobileAppId := "mobileAppId_example" // string | key: id of mobileApp
    microsoftGraphMobileAppAssignment := *openapiclient.NewMicrosoftGraphMobileAppAssignment() // MicrosoftGraphMobileAppAssignment | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceAppManagementMobileAppApi.DeviceAppManagementMobileAppsCreateAssignments(context.Background(), mobileAppId).MicrosoftGraphMobileAppAssignment(microsoftGraphMobileAppAssignment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAppManagementMobileAppApi.DeviceAppManagementMobileAppsCreateAssignments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceAppManagementMobileAppsCreateAssignments`: MicrosoftGraphMobileAppAssignment
    fmt.Fprintf(os.Stdout, "Response from `DeviceAppManagementMobileAppApi.DeviceAppManagementMobileAppsCreateAssignments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mobileAppId** | **string** | key: id of mobileApp | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceAppManagementMobileAppsCreateAssignmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphMobileAppAssignment** | [**MicrosoftGraphMobileAppAssignment**](MicrosoftGraphMobileAppAssignment.md) | New navigation property | 

### Return type

[**MicrosoftGraphMobileAppAssignment**](MicrosoftGraphMobileAppAssignment.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementMobileAppsCreateRefCategories

> map[string]interface{} DeviceAppManagementMobileAppsCreateRefCategories(ctx, mobileAppId).RequestBody(requestBody).Execute()

Create new navigation property ref to categories for deviceAppManagement



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
    mobileAppId := "mobileAppId_example" // string | key: id of mobileApp
    requestBody := map[string]map[string]interface{}{"key": map[string]interface{}(123)} // map[string]map[string]interface{} | New navigation property ref value

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceAppManagementMobileAppApi.DeviceAppManagementMobileAppsCreateRefCategories(context.Background(), mobileAppId).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAppManagementMobileAppApi.DeviceAppManagementMobileAppsCreateRefCategories``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceAppManagementMobileAppsCreateRefCategories`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `DeviceAppManagementMobileAppApi.DeviceAppManagementMobileAppsCreateRefCategories`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mobileAppId** | **string** | key: id of mobileApp | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceAppManagementMobileAppsCreateRefCategoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBody** | **map[string]map[string]interface{}** | New navigation property ref value | 

### Return type

**map[string]interface{}**

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementMobileAppsDeleteAssignments

> DeviceAppManagementMobileAppsDeleteAssignments(ctx, mobileAppId, mobileAppAssignmentId).IfMatch(ifMatch).Execute()

Delete navigation property assignments for deviceAppManagement



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
    mobileAppId := "mobileAppId_example" // string | key: id of mobileApp
    mobileAppAssignmentId := "mobileAppAssignmentId_example" // string | key: id of mobileAppAssignment
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceAppManagementMobileAppApi.DeviceAppManagementMobileAppsDeleteAssignments(context.Background(), mobileAppId, mobileAppAssignmentId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAppManagementMobileAppApi.DeviceAppManagementMobileAppsDeleteAssignments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mobileAppId** | **string** | key: id of mobileApp | 
**mobileAppAssignmentId** | **string** | key: id of mobileAppAssignment | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceAppManagementMobileAppsDeleteAssignmentsRequest struct via the builder pattern


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


## DeviceAppManagementMobileAppsGetAssignments

> MicrosoftGraphMobileAppAssignment DeviceAppManagementMobileAppsGetAssignments(ctx, mobileAppId, mobileAppAssignmentId).Select_(select_).Expand(expand).Execute()

Get assignments from deviceAppManagement



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
    mobileAppId := "mobileAppId_example" // string | key: id of mobileApp
    mobileAppAssignmentId := "mobileAppAssignmentId_example" // string | key: id of mobileAppAssignment
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceAppManagementMobileAppApi.DeviceAppManagementMobileAppsGetAssignments(context.Background(), mobileAppId, mobileAppAssignmentId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAppManagementMobileAppApi.DeviceAppManagementMobileAppsGetAssignments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceAppManagementMobileAppsGetAssignments`: MicrosoftGraphMobileAppAssignment
    fmt.Fprintf(os.Stdout, "Response from `DeviceAppManagementMobileAppApi.DeviceAppManagementMobileAppsGetAssignments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mobileAppId** | **string** | key: id of mobileApp | 
**mobileAppAssignmentId** | **string** | key: id of mobileAppAssignment | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceAppManagementMobileAppsGetAssignmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphMobileAppAssignment**](MicrosoftGraphMobileAppAssignment.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementMobileAppsListAssignments

> CollectionOfMobileAppAssignment DeviceAppManagementMobileAppsListAssignments(ctx, mobileAppId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get assignments from deviceAppManagement



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
    mobileAppId := "mobileAppId_example" // string | key: id of mobileApp
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
    resp, r, err := api_client.DeviceAppManagementMobileAppApi.DeviceAppManagementMobileAppsListAssignments(context.Background(), mobileAppId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAppManagementMobileAppApi.DeviceAppManagementMobileAppsListAssignments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceAppManagementMobileAppsListAssignments`: CollectionOfMobileAppAssignment
    fmt.Fprintf(os.Stdout, "Response from `DeviceAppManagementMobileAppApi.DeviceAppManagementMobileAppsListAssignments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mobileAppId** | **string** | key: id of mobileApp | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceAppManagementMobileAppsListAssignmentsRequest struct via the builder pattern


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

[**CollectionOfMobileAppAssignment**](CollectionOfMobileAppAssignment.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementMobileAppsListCategories

> CollectionOfMobileAppCategory DeviceAppManagementMobileAppsListCategories(ctx, mobileAppId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get categories from deviceAppManagement



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
    mobileAppId := "mobileAppId_example" // string | key: id of mobileApp
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
    resp, r, err := api_client.DeviceAppManagementMobileAppApi.DeviceAppManagementMobileAppsListCategories(context.Background(), mobileAppId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAppManagementMobileAppApi.DeviceAppManagementMobileAppsListCategories``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceAppManagementMobileAppsListCategories`: CollectionOfMobileAppCategory
    fmt.Fprintf(os.Stdout, "Response from `DeviceAppManagementMobileAppApi.DeviceAppManagementMobileAppsListCategories`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mobileAppId** | **string** | key: id of mobileApp | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceAppManagementMobileAppsListCategoriesRequest struct via the builder pattern


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


## DeviceAppManagementMobileAppsListRefCategories

> CollectionOfLinksOfMobileAppCategory DeviceAppManagementMobileAppsListRefCategories(ctx, mobileAppId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Execute()

Get ref of categories from deviceAppManagement



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
    mobileAppId := "mobileAppId_example" // string | key: id of mobileApp
    top := int32(50) // int32 | Show only the first n items (optional)
    skip := int32(56) // int32 | Skip the first n items (optional)
    search := "search_example" // string | Search items by search phrases (optional)
    filter := "filter_example" // string | Filter items by property values (optional)
    count := true // bool | Include count of items (optional)
    orderby := []string{"Orderby_example"} // []string | Order items by property values (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceAppManagementMobileAppApi.DeviceAppManagementMobileAppsListRefCategories(context.Background(), mobileAppId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAppManagementMobileAppApi.DeviceAppManagementMobileAppsListRefCategories``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceAppManagementMobileAppsListRefCategories`: CollectionOfLinksOfMobileAppCategory
    fmt.Fprintf(os.Stdout, "Response from `DeviceAppManagementMobileAppApi.DeviceAppManagementMobileAppsListRefCategories`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mobileAppId** | **string** | key: id of mobileApp | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceAppManagementMobileAppsListRefCategoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **top** | **int32** | Show only the first n items | 
 **skip** | **int32** | Skip the first n items | 
 **search** | **string** | Search items by search phrases | 
 **filter** | **string** | Filter items by property values | 
 **count** | **bool** | Include count of items | 
 **orderby** | **[]string** | Order items by property values | 

### Return type

[**CollectionOfLinksOfMobileAppCategory**](CollectionOfLinksOfMobileAppCategory.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementMobileAppsUpdateAssignments

> DeviceAppManagementMobileAppsUpdateAssignments(ctx, mobileAppId, mobileAppAssignmentId).MicrosoftGraphMobileAppAssignment(microsoftGraphMobileAppAssignment).Execute()

Update the navigation property assignments in deviceAppManagement



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
    mobileAppId := "mobileAppId_example" // string | key: id of mobileApp
    mobileAppAssignmentId := "mobileAppAssignmentId_example" // string | key: id of mobileAppAssignment
    microsoftGraphMobileAppAssignment := *openapiclient.NewMicrosoftGraphMobileAppAssignment() // MicrosoftGraphMobileAppAssignment | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceAppManagementMobileAppApi.DeviceAppManagementMobileAppsUpdateAssignments(context.Background(), mobileAppId, mobileAppAssignmentId).MicrosoftGraphMobileAppAssignment(microsoftGraphMobileAppAssignment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAppManagementMobileAppApi.DeviceAppManagementMobileAppsUpdateAssignments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mobileAppId** | **string** | key: id of mobileApp | 
**mobileAppAssignmentId** | **string** | key: id of mobileAppAssignment | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceAppManagementMobileAppsUpdateAssignmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **microsoftGraphMobileAppAssignment** | [**MicrosoftGraphMobileAppAssignment**](MicrosoftGraphMobileAppAssignment.md) | New navigation property values | 

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


## DeviceAppManagementUpdateMobileApps

> DeviceAppManagementUpdateMobileApps(ctx, mobileAppId).MicrosoftGraphMobileApp(microsoftGraphMobileApp).Execute()

Update the navigation property mobileApps in deviceAppManagement



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
    mobileAppId := "mobileAppId_example" // string | key: id of mobileApp
    microsoftGraphMobileApp := *openapiclient.NewMicrosoftGraphMobileApp() // MicrosoftGraphMobileApp | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceAppManagementMobileAppApi.DeviceAppManagementUpdateMobileApps(context.Background(), mobileAppId).MicrosoftGraphMobileApp(microsoftGraphMobileApp).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAppManagementMobileAppApi.DeviceAppManagementUpdateMobileApps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mobileAppId** | **string** | key: id of mobileApp | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceAppManagementUpdateMobileAppsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphMobileApp** | [**MicrosoftGraphMobileApp**](MicrosoftGraphMobileApp.md) | New navigation property values | 

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

