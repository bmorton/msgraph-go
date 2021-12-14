# \DeviceAppManagementAndroidManagedAppProtectionApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceAppManagementAndroidManagedAppProtectionsCreateApps**](DeviceAppManagementAndroidManagedAppProtectionApi.md#DeviceAppManagementAndroidManagedAppProtectionsCreateApps) | **Post** /deviceAppManagement/androidManagedAppProtections/{androidManagedAppProtection-id}/apps | Create new navigation property to apps for deviceAppManagement
[**DeviceAppManagementAndroidManagedAppProtectionsDeleteApps**](DeviceAppManagementAndroidManagedAppProtectionApi.md#DeviceAppManagementAndroidManagedAppProtectionsDeleteApps) | **Delete** /deviceAppManagement/androidManagedAppProtections/{androidManagedAppProtection-id}/apps/{managedMobileApp-id} | Delete navigation property apps for deviceAppManagement
[**DeviceAppManagementAndroidManagedAppProtectionsDeleteDeploymentSummary**](DeviceAppManagementAndroidManagedAppProtectionApi.md#DeviceAppManagementAndroidManagedAppProtectionsDeleteDeploymentSummary) | **Delete** /deviceAppManagement/androidManagedAppProtections/{androidManagedAppProtection-id}/deploymentSummary | Delete navigation property deploymentSummary for deviceAppManagement
[**DeviceAppManagementAndroidManagedAppProtectionsGetApps**](DeviceAppManagementAndroidManagedAppProtectionApi.md#DeviceAppManagementAndroidManagedAppProtectionsGetApps) | **Get** /deviceAppManagement/androidManagedAppProtections/{androidManagedAppProtection-id}/apps/{managedMobileApp-id} | Get apps from deviceAppManagement
[**DeviceAppManagementAndroidManagedAppProtectionsGetDeploymentSummary**](DeviceAppManagementAndroidManagedAppProtectionApi.md#DeviceAppManagementAndroidManagedAppProtectionsGetDeploymentSummary) | **Get** /deviceAppManagement/androidManagedAppProtections/{androidManagedAppProtection-id}/deploymentSummary | Get deploymentSummary from deviceAppManagement
[**DeviceAppManagementAndroidManagedAppProtectionsListApps**](DeviceAppManagementAndroidManagedAppProtectionApi.md#DeviceAppManagementAndroidManagedAppProtectionsListApps) | **Get** /deviceAppManagement/androidManagedAppProtections/{androidManagedAppProtection-id}/apps | Get apps from deviceAppManagement
[**DeviceAppManagementAndroidManagedAppProtectionsUpdateApps**](DeviceAppManagementAndroidManagedAppProtectionApi.md#DeviceAppManagementAndroidManagedAppProtectionsUpdateApps) | **Patch** /deviceAppManagement/androidManagedAppProtections/{androidManagedAppProtection-id}/apps/{managedMobileApp-id} | Update the navigation property apps in deviceAppManagement
[**DeviceAppManagementAndroidManagedAppProtectionsUpdateDeploymentSummary**](DeviceAppManagementAndroidManagedAppProtectionApi.md#DeviceAppManagementAndroidManagedAppProtectionsUpdateDeploymentSummary) | **Patch** /deviceAppManagement/androidManagedAppProtections/{androidManagedAppProtection-id}/deploymentSummary | Update the navigation property deploymentSummary in deviceAppManagement
[**DeviceAppManagementCreateAndroidManagedAppProtections**](DeviceAppManagementAndroidManagedAppProtectionApi.md#DeviceAppManagementCreateAndroidManagedAppProtections) | **Post** /deviceAppManagement/androidManagedAppProtections | Create new navigation property to androidManagedAppProtections for deviceAppManagement
[**DeviceAppManagementDeleteAndroidManagedAppProtections**](DeviceAppManagementAndroidManagedAppProtectionApi.md#DeviceAppManagementDeleteAndroidManagedAppProtections) | **Delete** /deviceAppManagement/androidManagedAppProtections/{androidManagedAppProtection-id} | Delete navigation property androidManagedAppProtections for deviceAppManagement
[**DeviceAppManagementGetAndroidManagedAppProtections**](DeviceAppManagementAndroidManagedAppProtectionApi.md#DeviceAppManagementGetAndroidManagedAppProtections) | **Get** /deviceAppManagement/androidManagedAppProtections/{androidManagedAppProtection-id} | Get androidManagedAppProtections from deviceAppManagement
[**DeviceAppManagementListAndroidManagedAppProtections**](DeviceAppManagementAndroidManagedAppProtectionApi.md#DeviceAppManagementListAndroidManagedAppProtections) | **Get** /deviceAppManagement/androidManagedAppProtections | Get androidManagedAppProtections from deviceAppManagement
[**DeviceAppManagementUpdateAndroidManagedAppProtections**](DeviceAppManagementAndroidManagedAppProtectionApi.md#DeviceAppManagementUpdateAndroidManagedAppProtections) | **Patch** /deviceAppManagement/androidManagedAppProtections/{androidManagedAppProtection-id} | Update the navigation property androidManagedAppProtections in deviceAppManagement



## DeviceAppManagementAndroidManagedAppProtectionsCreateApps

> MicrosoftGraphManagedMobileApp DeviceAppManagementAndroidManagedAppProtectionsCreateApps(ctx, androidManagedAppProtectionId).MicrosoftGraphManagedMobileApp(microsoftGraphManagedMobileApp).Execute()

Create new navigation property to apps for deviceAppManagement



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
    androidManagedAppProtectionId := "androidManagedAppProtectionId_example" // string | key: id of androidManagedAppProtection
    microsoftGraphManagedMobileApp := *openapiclient.NewMicrosoftGraphManagedMobileApp() // MicrosoftGraphManagedMobileApp | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceAppManagementAndroidManagedAppProtectionApi.DeviceAppManagementAndroidManagedAppProtectionsCreateApps(context.Background(), androidManagedAppProtectionId).MicrosoftGraphManagedMobileApp(microsoftGraphManagedMobileApp).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAppManagementAndroidManagedAppProtectionApi.DeviceAppManagementAndroidManagedAppProtectionsCreateApps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceAppManagementAndroidManagedAppProtectionsCreateApps`: MicrosoftGraphManagedMobileApp
    fmt.Fprintf(os.Stdout, "Response from `DeviceAppManagementAndroidManagedAppProtectionApi.DeviceAppManagementAndroidManagedAppProtectionsCreateApps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**androidManagedAppProtectionId** | **string** | key: id of androidManagedAppProtection | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceAppManagementAndroidManagedAppProtectionsCreateAppsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphManagedMobileApp** | [**MicrosoftGraphManagedMobileApp**](MicrosoftGraphManagedMobileApp.md) | New navigation property | 

### Return type

[**MicrosoftGraphManagedMobileApp**](MicrosoftGraphManagedMobileApp.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementAndroidManagedAppProtectionsDeleteApps

> DeviceAppManagementAndroidManagedAppProtectionsDeleteApps(ctx, androidManagedAppProtectionId, managedMobileAppId).IfMatch(ifMatch).Execute()

Delete navigation property apps for deviceAppManagement



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
    androidManagedAppProtectionId := "androidManagedAppProtectionId_example" // string | key: id of androidManagedAppProtection
    managedMobileAppId := "managedMobileAppId_example" // string | key: id of managedMobileApp
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceAppManagementAndroidManagedAppProtectionApi.DeviceAppManagementAndroidManagedAppProtectionsDeleteApps(context.Background(), androidManagedAppProtectionId, managedMobileAppId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAppManagementAndroidManagedAppProtectionApi.DeviceAppManagementAndroidManagedAppProtectionsDeleteApps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**androidManagedAppProtectionId** | **string** | key: id of androidManagedAppProtection | 
**managedMobileAppId** | **string** | key: id of managedMobileApp | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceAppManagementAndroidManagedAppProtectionsDeleteAppsRequest struct via the builder pattern


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


## DeviceAppManagementAndroidManagedAppProtectionsDeleteDeploymentSummary

> DeviceAppManagementAndroidManagedAppProtectionsDeleteDeploymentSummary(ctx, androidManagedAppProtectionId).IfMatch(ifMatch).Execute()

Delete navigation property deploymentSummary for deviceAppManagement



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
    androidManagedAppProtectionId := "androidManagedAppProtectionId_example" // string | key: id of androidManagedAppProtection
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceAppManagementAndroidManagedAppProtectionApi.DeviceAppManagementAndroidManagedAppProtectionsDeleteDeploymentSummary(context.Background(), androidManagedAppProtectionId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAppManagementAndroidManagedAppProtectionApi.DeviceAppManagementAndroidManagedAppProtectionsDeleteDeploymentSummary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**androidManagedAppProtectionId** | **string** | key: id of androidManagedAppProtection | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceAppManagementAndroidManagedAppProtectionsDeleteDeploymentSummaryRequest struct via the builder pattern


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


## DeviceAppManagementAndroidManagedAppProtectionsGetApps

> MicrosoftGraphManagedMobileApp DeviceAppManagementAndroidManagedAppProtectionsGetApps(ctx, androidManagedAppProtectionId, managedMobileAppId).Select_(select_).Expand(expand).Execute()

Get apps from deviceAppManagement



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
    androidManagedAppProtectionId := "androidManagedAppProtectionId_example" // string | key: id of androidManagedAppProtection
    managedMobileAppId := "managedMobileAppId_example" // string | key: id of managedMobileApp
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceAppManagementAndroidManagedAppProtectionApi.DeviceAppManagementAndroidManagedAppProtectionsGetApps(context.Background(), androidManagedAppProtectionId, managedMobileAppId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAppManagementAndroidManagedAppProtectionApi.DeviceAppManagementAndroidManagedAppProtectionsGetApps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceAppManagementAndroidManagedAppProtectionsGetApps`: MicrosoftGraphManagedMobileApp
    fmt.Fprintf(os.Stdout, "Response from `DeviceAppManagementAndroidManagedAppProtectionApi.DeviceAppManagementAndroidManagedAppProtectionsGetApps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**androidManagedAppProtectionId** | **string** | key: id of androidManagedAppProtection | 
**managedMobileAppId** | **string** | key: id of managedMobileApp | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceAppManagementAndroidManagedAppProtectionsGetAppsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphManagedMobileApp**](MicrosoftGraphManagedMobileApp.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementAndroidManagedAppProtectionsGetDeploymentSummary

> MicrosoftGraphManagedAppPolicyDeploymentSummary DeviceAppManagementAndroidManagedAppProtectionsGetDeploymentSummary(ctx, androidManagedAppProtectionId).Select_(select_).Expand(expand).Execute()

Get deploymentSummary from deviceAppManagement



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
    androidManagedAppProtectionId := "androidManagedAppProtectionId_example" // string | key: id of androidManagedAppProtection
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceAppManagementAndroidManagedAppProtectionApi.DeviceAppManagementAndroidManagedAppProtectionsGetDeploymentSummary(context.Background(), androidManagedAppProtectionId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAppManagementAndroidManagedAppProtectionApi.DeviceAppManagementAndroidManagedAppProtectionsGetDeploymentSummary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceAppManagementAndroidManagedAppProtectionsGetDeploymentSummary`: MicrosoftGraphManagedAppPolicyDeploymentSummary
    fmt.Fprintf(os.Stdout, "Response from `DeviceAppManagementAndroidManagedAppProtectionApi.DeviceAppManagementAndroidManagedAppProtectionsGetDeploymentSummary`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**androidManagedAppProtectionId** | **string** | key: id of androidManagedAppProtection | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceAppManagementAndroidManagedAppProtectionsGetDeploymentSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphManagedAppPolicyDeploymentSummary**](MicrosoftGraphManagedAppPolicyDeploymentSummary.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementAndroidManagedAppProtectionsListApps

> CollectionOfManagedMobileApp DeviceAppManagementAndroidManagedAppProtectionsListApps(ctx, androidManagedAppProtectionId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get apps from deviceAppManagement



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
    androidManagedAppProtectionId := "androidManagedAppProtectionId_example" // string | key: id of androidManagedAppProtection
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
    resp, r, err := api_client.DeviceAppManagementAndroidManagedAppProtectionApi.DeviceAppManagementAndroidManagedAppProtectionsListApps(context.Background(), androidManagedAppProtectionId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAppManagementAndroidManagedAppProtectionApi.DeviceAppManagementAndroidManagedAppProtectionsListApps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceAppManagementAndroidManagedAppProtectionsListApps`: CollectionOfManagedMobileApp
    fmt.Fprintf(os.Stdout, "Response from `DeviceAppManagementAndroidManagedAppProtectionApi.DeviceAppManagementAndroidManagedAppProtectionsListApps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**androidManagedAppProtectionId** | **string** | key: id of androidManagedAppProtection | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceAppManagementAndroidManagedAppProtectionsListAppsRequest struct via the builder pattern


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

[**CollectionOfManagedMobileApp**](CollectionOfManagedMobileApp.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementAndroidManagedAppProtectionsUpdateApps

> DeviceAppManagementAndroidManagedAppProtectionsUpdateApps(ctx, androidManagedAppProtectionId, managedMobileAppId).MicrosoftGraphManagedMobileApp(microsoftGraphManagedMobileApp).Execute()

Update the navigation property apps in deviceAppManagement



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
    androidManagedAppProtectionId := "androidManagedAppProtectionId_example" // string | key: id of androidManagedAppProtection
    managedMobileAppId := "managedMobileAppId_example" // string | key: id of managedMobileApp
    microsoftGraphManagedMobileApp := *openapiclient.NewMicrosoftGraphManagedMobileApp() // MicrosoftGraphManagedMobileApp | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceAppManagementAndroidManagedAppProtectionApi.DeviceAppManagementAndroidManagedAppProtectionsUpdateApps(context.Background(), androidManagedAppProtectionId, managedMobileAppId).MicrosoftGraphManagedMobileApp(microsoftGraphManagedMobileApp).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAppManagementAndroidManagedAppProtectionApi.DeviceAppManagementAndroidManagedAppProtectionsUpdateApps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**androidManagedAppProtectionId** | **string** | key: id of androidManagedAppProtection | 
**managedMobileAppId** | **string** | key: id of managedMobileApp | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceAppManagementAndroidManagedAppProtectionsUpdateAppsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **microsoftGraphManagedMobileApp** | [**MicrosoftGraphManagedMobileApp**](MicrosoftGraphManagedMobileApp.md) | New navigation property values | 

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


## DeviceAppManagementAndroidManagedAppProtectionsUpdateDeploymentSummary

> DeviceAppManagementAndroidManagedAppProtectionsUpdateDeploymentSummary(ctx, androidManagedAppProtectionId).MicrosoftGraphManagedAppPolicyDeploymentSummary(microsoftGraphManagedAppPolicyDeploymentSummary).Execute()

Update the navigation property deploymentSummary in deviceAppManagement



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
    androidManagedAppProtectionId := "androidManagedAppProtectionId_example" // string | key: id of androidManagedAppProtection
    microsoftGraphManagedAppPolicyDeploymentSummary := *openapiclient.NewMicrosoftGraphManagedAppPolicyDeploymentSummary() // MicrosoftGraphManagedAppPolicyDeploymentSummary | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceAppManagementAndroidManagedAppProtectionApi.DeviceAppManagementAndroidManagedAppProtectionsUpdateDeploymentSummary(context.Background(), androidManagedAppProtectionId).MicrosoftGraphManagedAppPolicyDeploymentSummary(microsoftGraphManagedAppPolicyDeploymentSummary).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAppManagementAndroidManagedAppProtectionApi.DeviceAppManagementAndroidManagedAppProtectionsUpdateDeploymentSummary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**androidManagedAppProtectionId** | **string** | key: id of androidManagedAppProtection | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceAppManagementAndroidManagedAppProtectionsUpdateDeploymentSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphManagedAppPolicyDeploymentSummary** | [**MicrosoftGraphManagedAppPolicyDeploymentSummary**](MicrosoftGraphManagedAppPolicyDeploymentSummary.md) | New navigation property values | 

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


## DeviceAppManagementCreateAndroidManagedAppProtections

> MicrosoftGraphAndroidManagedAppProtection DeviceAppManagementCreateAndroidManagedAppProtections(ctx).MicrosoftGraphAndroidManagedAppProtection(microsoftGraphAndroidManagedAppProtection).Execute()

Create new navigation property to androidManagedAppProtections for deviceAppManagement



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
    microsoftGraphAndroidManagedAppProtection := *openapiclient.NewMicrosoftGraphAndroidManagedAppProtection() // MicrosoftGraphAndroidManagedAppProtection | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceAppManagementAndroidManagedAppProtectionApi.DeviceAppManagementCreateAndroidManagedAppProtections(context.Background()).MicrosoftGraphAndroidManagedAppProtection(microsoftGraphAndroidManagedAppProtection).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAppManagementAndroidManagedAppProtectionApi.DeviceAppManagementCreateAndroidManagedAppProtections``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceAppManagementCreateAndroidManagedAppProtections`: MicrosoftGraphAndroidManagedAppProtection
    fmt.Fprintf(os.Stdout, "Response from `DeviceAppManagementAndroidManagedAppProtectionApi.DeviceAppManagementCreateAndroidManagedAppProtections`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeviceAppManagementCreateAndroidManagedAppProtectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphAndroidManagedAppProtection** | [**MicrosoftGraphAndroidManagedAppProtection**](MicrosoftGraphAndroidManagedAppProtection.md) | New navigation property | 

### Return type

[**MicrosoftGraphAndroidManagedAppProtection**](MicrosoftGraphAndroidManagedAppProtection.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementDeleteAndroidManagedAppProtections

> DeviceAppManagementDeleteAndroidManagedAppProtections(ctx, androidManagedAppProtectionId).IfMatch(ifMatch).Execute()

Delete navigation property androidManagedAppProtections for deviceAppManagement



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
    androidManagedAppProtectionId := "androidManagedAppProtectionId_example" // string | key: id of androidManagedAppProtection
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceAppManagementAndroidManagedAppProtectionApi.DeviceAppManagementDeleteAndroidManagedAppProtections(context.Background(), androidManagedAppProtectionId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAppManagementAndroidManagedAppProtectionApi.DeviceAppManagementDeleteAndroidManagedAppProtections``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**androidManagedAppProtectionId** | **string** | key: id of androidManagedAppProtection | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceAppManagementDeleteAndroidManagedAppProtectionsRequest struct via the builder pattern


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


## DeviceAppManagementGetAndroidManagedAppProtections

> MicrosoftGraphAndroidManagedAppProtection DeviceAppManagementGetAndroidManagedAppProtections(ctx, androidManagedAppProtectionId).Select_(select_).Expand(expand).Execute()

Get androidManagedAppProtections from deviceAppManagement



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
    androidManagedAppProtectionId := "androidManagedAppProtectionId_example" // string | key: id of androidManagedAppProtection
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceAppManagementAndroidManagedAppProtectionApi.DeviceAppManagementGetAndroidManagedAppProtections(context.Background(), androidManagedAppProtectionId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAppManagementAndroidManagedAppProtectionApi.DeviceAppManagementGetAndroidManagedAppProtections``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceAppManagementGetAndroidManagedAppProtections`: MicrosoftGraphAndroidManagedAppProtection
    fmt.Fprintf(os.Stdout, "Response from `DeviceAppManagementAndroidManagedAppProtectionApi.DeviceAppManagementGetAndroidManagedAppProtections`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**androidManagedAppProtectionId** | **string** | key: id of androidManagedAppProtection | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceAppManagementGetAndroidManagedAppProtectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphAndroidManagedAppProtection**](MicrosoftGraphAndroidManagedAppProtection.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementListAndroidManagedAppProtections

> CollectionOfAndroidManagedAppProtection DeviceAppManagementListAndroidManagedAppProtections(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get androidManagedAppProtections from deviceAppManagement



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
    resp, r, err := api_client.DeviceAppManagementAndroidManagedAppProtectionApi.DeviceAppManagementListAndroidManagedAppProtections(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAppManagementAndroidManagedAppProtectionApi.DeviceAppManagementListAndroidManagedAppProtections``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceAppManagementListAndroidManagedAppProtections`: CollectionOfAndroidManagedAppProtection
    fmt.Fprintf(os.Stdout, "Response from `DeviceAppManagementAndroidManagedAppProtectionApi.DeviceAppManagementListAndroidManagedAppProtections`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeviceAppManagementListAndroidManagedAppProtectionsRequest struct via the builder pattern


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

[**CollectionOfAndroidManagedAppProtection**](CollectionOfAndroidManagedAppProtection.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementUpdateAndroidManagedAppProtections

> DeviceAppManagementUpdateAndroidManagedAppProtections(ctx, androidManagedAppProtectionId).MicrosoftGraphAndroidManagedAppProtection(microsoftGraphAndroidManagedAppProtection).Execute()

Update the navigation property androidManagedAppProtections in deviceAppManagement



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
    androidManagedAppProtectionId := "androidManagedAppProtectionId_example" // string | key: id of androidManagedAppProtection
    microsoftGraphAndroidManagedAppProtection := *openapiclient.NewMicrosoftGraphAndroidManagedAppProtection() // MicrosoftGraphAndroidManagedAppProtection | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceAppManagementAndroidManagedAppProtectionApi.DeviceAppManagementUpdateAndroidManagedAppProtections(context.Background(), androidManagedAppProtectionId).MicrosoftGraphAndroidManagedAppProtection(microsoftGraphAndroidManagedAppProtection).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAppManagementAndroidManagedAppProtectionApi.DeviceAppManagementUpdateAndroidManagedAppProtections``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**androidManagedAppProtectionId** | **string** | key: id of androidManagedAppProtection | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceAppManagementUpdateAndroidManagedAppProtectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphAndroidManagedAppProtection** | [**MicrosoftGraphAndroidManagedAppProtection**](MicrosoftGraphAndroidManagedAppProtection.md) | New navigation property values | 

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

