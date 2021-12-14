# \DeviceAppManagementIosManagedAppProtectionApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceAppManagementCreateIosManagedAppProtections**](DeviceAppManagementIosManagedAppProtectionApi.md#DeviceAppManagementCreateIosManagedAppProtections) | **Post** /deviceAppManagement/iosManagedAppProtections | Create new navigation property to iosManagedAppProtections for deviceAppManagement
[**DeviceAppManagementDeleteIosManagedAppProtections**](DeviceAppManagementIosManagedAppProtectionApi.md#DeviceAppManagementDeleteIosManagedAppProtections) | **Delete** /deviceAppManagement/iosManagedAppProtections/{iosManagedAppProtection-id} | Delete navigation property iosManagedAppProtections for deviceAppManagement
[**DeviceAppManagementGetIosManagedAppProtections**](DeviceAppManagementIosManagedAppProtectionApi.md#DeviceAppManagementGetIosManagedAppProtections) | **Get** /deviceAppManagement/iosManagedAppProtections/{iosManagedAppProtection-id} | Get iosManagedAppProtections from deviceAppManagement
[**DeviceAppManagementIosManagedAppProtectionsCreateApps**](DeviceAppManagementIosManagedAppProtectionApi.md#DeviceAppManagementIosManagedAppProtectionsCreateApps) | **Post** /deviceAppManagement/iosManagedAppProtections/{iosManagedAppProtection-id}/apps | Create new navigation property to apps for deviceAppManagement
[**DeviceAppManagementIosManagedAppProtectionsDeleteApps**](DeviceAppManagementIosManagedAppProtectionApi.md#DeviceAppManagementIosManagedAppProtectionsDeleteApps) | **Delete** /deviceAppManagement/iosManagedAppProtections/{iosManagedAppProtection-id}/apps/{managedMobileApp-id} | Delete navigation property apps for deviceAppManagement
[**DeviceAppManagementIosManagedAppProtectionsDeleteDeploymentSummary**](DeviceAppManagementIosManagedAppProtectionApi.md#DeviceAppManagementIosManagedAppProtectionsDeleteDeploymentSummary) | **Delete** /deviceAppManagement/iosManagedAppProtections/{iosManagedAppProtection-id}/deploymentSummary | Delete navigation property deploymentSummary for deviceAppManagement
[**DeviceAppManagementIosManagedAppProtectionsGetApps**](DeviceAppManagementIosManagedAppProtectionApi.md#DeviceAppManagementIosManagedAppProtectionsGetApps) | **Get** /deviceAppManagement/iosManagedAppProtections/{iosManagedAppProtection-id}/apps/{managedMobileApp-id} | Get apps from deviceAppManagement
[**DeviceAppManagementIosManagedAppProtectionsGetDeploymentSummary**](DeviceAppManagementIosManagedAppProtectionApi.md#DeviceAppManagementIosManagedAppProtectionsGetDeploymentSummary) | **Get** /deviceAppManagement/iosManagedAppProtections/{iosManagedAppProtection-id}/deploymentSummary | Get deploymentSummary from deviceAppManagement
[**DeviceAppManagementIosManagedAppProtectionsListApps**](DeviceAppManagementIosManagedAppProtectionApi.md#DeviceAppManagementIosManagedAppProtectionsListApps) | **Get** /deviceAppManagement/iosManagedAppProtections/{iosManagedAppProtection-id}/apps | Get apps from deviceAppManagement
[**DeviceAppManagementIosManagedAppProtectionsUpdateApps**](DeviceAppManagementIosManagedAppProtectionApi.md#DeviceAppManagementIosManagedAppProtectionsUpdateApps) | **Patch** /deviceAppManagement/iosManagedAppProtections/{iosManagedAppProtection-id}/apps/{managedMobileApp-id} | Update the navigation property apps in deviceAppManagement
[**DeviceAppManagementIosManagedAppProtectionsUpdateDeploymentSummary**](DeviceAppManagementIosManagedAppProtectionApi.md#DeviceAppManagementIosManagedAppProtectionsUpdateDeploymentSummary) | **Patch** /deviceAppManagement/iosManagedAppProtections/{iosManagedAppProtection-id}/deploymentSummary | Update the navigation property deploymentSummary in deviceAppManagement
[**DeviceAppManagementListIosManagedAppProtections**](DeviceAppManagementIosManagedAppProtectionApi.md#DeviceAppManagementListIosManagedAppProtections) | **Get** /deviceAppManagement/iosManagedAppProtections | Get iosManagedAppProtections from deviceAppManagement
[**DeviceAppManagementUpdateIosManagedAppProtections**](DeviceAppManagementIosManagedAppProtectionApi.md#DeviceAppManagementUpdateIosManagedAppProtections) | **Patch** /deviceAppManagement/iosManagedAppProtections/{iosManagedAppProtection-id} | Update the navigation property iosManagedAppProtections in deviceAppManagement



## DeviceAppManagementCreateIosManagedAppProtections

> MicrosoftGraphIosManagedAppProtection DeviceAppManagementCreateIosManagedAppProtections(ctx).MicrosoftGraphIosManagedAppProtection(microsoftGraphIosManagedAppProtection).Execute()

Create new navigation property to iosManagedAppProtections for deviceAppManagement



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
    microsoftGraphIosManagedAppProtection := *openapiclient.NewMicrosoftGraphIosManagedAppProtection() // MicrosoftGraphIosManagedAppProtection | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceAppManagementIosManagedAppProtectionApi.DeviceAppManagementCreateIosManagedAppProtections(context.Background()).MicrosoftGraphIosManagedAppProtection(microsoftGraphIosManagedAppProtection).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAppManagementIosManagedAppProtectionApi.DeviceAppManagementCreateIosManagedAppProtections``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceAppManagementCreateIosManagedAppProtections`: MicrosoftGraphIosManagedAppProtection
    fmt.Fprintf(os.Stdout, "Response from `DeviceAppManagementIosManagedAppProtectionApi.DeviceAppManagementCreateIosManagedAppProtections`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeviceAppManagementCreateIosManagedAppProtectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphIosManagedAppProtection** | [**MicrosoftGraphIosManagedAppProtection**](MicrosoftGraphIosManagedAppProtection.md) | New navigation property | 

### Return type

[**MicrosoftGraphIosManagedAppProtection**](MicrosoftGraphIosManagedAppProtection.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementDeleteIosManagedAppProtections

> DeviceAppManagementDeleteIosManagedAppProtections(ctx, iosManagedAppProtectionId).IfMatch(ifMatch).Execute()

Delete navigation property iosManagedAppProtections for deviceAppManagement



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
    iosManagedAppProtectionId := "iosManagedAppProtectionId_example" // string | key: id of iosManagedAppProtection
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceAppManagementIosManagedAppProtectionApi.DeviceAppManagementDeleteIosManagedAppProtections(context.Background(), iosManagedAppProtectionId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAppManagementIosManagedAppProtectionApi.DeviceAppManagementDeleteIosManagedAppProtections``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**iosManagedAppProtectionId** | **string** | key: id of iosManagedAppProtection | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceAppManagementDeleteIosManagedAppProtectionsRequest struct via the builder pattern


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


## DeviceAppManagementGetIosManagedAppProtections

> MicrosoftGraphIosManagedAppProtection DeviceAppManagementGetIosManagedAppProtections(ctx, iosManagedAppProtectionId).Select_(select_).Expand(expand).Execute()

Get iosManagedAppProtections from deviceAppManagement



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
    iosManagedAppProtectionId := "iosManagedAppProtectionId_example" // string | key: id of iosManagedAppProtection
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceAppManagementIosManagedAppProtectionApi.DeviceAppManagementGetIosManagedAppProtections(context.Background(), iosManagedAppProtectionId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAppManagementIosManagedAppProtectionApi.DeviceAppManagementGetIosManagedAppProtections``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceAppManagementGetIosManagedAppProtections`: MicrosoftGraphIosManagedAppProtection
    fmt.Fprintf(os.Stdout, "Response from `DeviceAppManagementIosManagedAppProtectionApi.DeviceAppManagementGetIosManagedAppProtections`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**iosManagedAppProtectionId** | **string** | key: id of iosManagedAppProtection | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceAppManagementGetIosManagedAppProtectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphIosManagedAppProtection**](MicrosoftGraphIosManagedAppProtection.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementIosManagedAppProtectionsCreateApps

> MicrosoftGraphManagedMobileApp DeviceAppManagementIosManagedAppProtectionsCreateApps(ctx, iosManagedAppProtectionId).MicrosoftGraphManagedMobileApp(microsoftGraphManagedMobileApp).Execute()

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
    iosManagedAppProtectionId := "iosManagedAppProtectionId_example" // string | key: id of iosManagedAppProtection
    microsoftGraphManagedMobileApp := *openapiclient.NewMicrosoftGraphManagedMobileApp() // MicrosoftGraphManagedMobileApp | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceAppManagementIosManagedAppProtectionApi.DeviceAppManagementIosManagedAppProtectionsCreateApps(context.Background(), iosManagedAppProtectionId).MicrosoftGraphManagedMobileApp(microsoftGraphManagedMobileApp).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAppManagementIosManagedAppProtectionApi.DeviceAppManagementIosManagedAppProtectionsCreateApps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceAppManagementIosManagedAppProtectionsCreateApps`: MicrosoftGraphManagedMobileApp
    fmt.Fprintf(os.Stdout, "Response from `DeviceAppManagementIosManagedAppProtectionApi.DeviceAppManagementIosManagedAppProtectionsCreateApps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**iosManagedAppProtectionId** | **string** | key: id of iosManagedAppProtection | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceAppManagementIosManagedAppProtectionsCreateAppsRequest struct via the builder pattern


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


## DeviceAppManagementIosManagedAppProtectionsDeleteApps

> DeviceAppManagementIosManagedAppProtectionsDeleteApps(ctx, iosManagedAppProtectionId, managedMobileAppId).IfMatch(ifMatch).Execute()

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
    iosManagedAppProtectionId := "iosManagedAppProtectionId_example" // string | key: id of iosManagedAppProtection
    managedMobileAppId := "managedMobileAppId_example" // string | key: id of managedMobileApp
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceAppManagementIosManagedAppProtectionApi.DeviceAppManagementIosManagedAppProtectionsDeleteApps(context.Background(), iosManagedAppProtectionId, managedMobileAppId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAppManagementIosManagedAppProtectionApi.DeviceAppManagementIosManagedAppProtectionsDeleteApps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**iosManagedAppProtectionId** | **string** | key: id of iosManagedAppProtection | 
**managedMobileAppId** | **string** | key: id of managedMobileApp | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceAppManagementIosManagedAppProtectionsDeleteAppsRequest struct via the builder pattern


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


## DeviceAppManagementIosManagedAppProtectionsDeleteDeploymentSummary

> DeviceAppManagementIosManagedAppProtectionsDeleteDeploymentSummary(ctx, iosManagedAppProtectionId).IfMatch(ifMatch).Execute()

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
    iosManagedAppProtectionId := "iosManagedAppProtectionId_example" // string | key: id of iosManagedAppProtection
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceAppManagementIosManagedAppProtectionApi.DeviceAppManagementIosManagedAppProtectionsDeleteDeploymentSummary(context.Background(), iosManagedAppProtectionId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAppManagementIosManagedAppProtectionApi.DeviceAppManagementIosManagedAppProtectionsDeleteDeploymentSummary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**iosManagedAppProtectionId** | **string** | key: id of iosManagedAppProtection | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceAppManagementIosManagedAppProtectionsDeleteDeploymentSummaryRequest struct via the builder pattern


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


## DeviceAppManagementIosManagedAppProtectionsGetApps

> MicrosoftGraphManagedMobileApp DeviceAppManagementIosManagedAppProtectionsGetApps(ctx, iosManagedAppProtectionId, managedMobileAppId).Select_(select_).Expand(expand).Execute()

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
    iosManagedAppProtectionId := "iosManagedAppProtectionId_example" // string | key: id of iosManagedAppProtection
    managedMobileAppId := "managedMobileAppId_example" // string | key: id of managedMobileApp
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceAppManagementIosManagedAppProtectionApi.DeviceAppManagementIosManagedAppProtectionsGetApps(context.Background(), iosManagedAppProtectionId, managedMobileAppId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAppManagementIosManagedAppProtectionApi.DeviceAppManagementIosManagedAppProtectionsGetApps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceAppManagementIosManagedAppProtectionsGetApps`: MicrosoftGraphManagedMobileApp
    fmt.Fprintf(os.Stdout, "Response from `DeviceAppManagementIosManagedAppProtectionApi.DeviceAppManagementIosManagedAppProtectionsGetApps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**iosManagedAppProtectionId** | **string** | key: id of iosManagedAppProtection | 
**managedMobileAppId** | **string** | key: id of managedMobileApp | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceAppManagementIosManagedAppProtectionsGetAppsRequest struct via the builder pattern


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


## DeviceAppManagementIosManagedAppProtectionsGetDeploymentSummary

> MicrosoftGraphManagedAppPolicyDeploymentSummary DeviceAppManagementIosManagedAppProtectionsGetDeploymentSummary(ctx, iosManagedAppProtectionId).Select_(select_).Expand(expand).Execute()

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
    iosManagedAppProtectionId := "iosManagedAppProtectionId_example" // string | key: id of iosManagedAppProtection
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceAppManagementIosManagedAppProtectionApi.DeviceAppManagementIosManagedAppProtectionsGetDeploymentSummary(context.Background(), iosManagedAppProtectionId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAppManagementIosManagedAppProtectionApi.DeviceAppManagementIosManagedAppProtectionsGetDeploymentSummary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceAppManagementIosManagedAppProtectionsGetDeploymentSummary`: MicrosoftGraphManagedAppPolicyDeploymentSummary
    fmt.Fprintf(os.Stdout, "Response from `DeviceAppManagementIosManagedAppProtectionApi.DeviceAppManagementIosManagedAppProtectionsGetDeploymentSummary`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**iosManagedAppProtectionId** | **string** | key: id of iosManagedAppProtection | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceAppManagementIosManagedAppProtectionsGetDeploymentSummaryRequest struct via the builder pattern


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


## DeviceAppManagementIosManagedAppProtectionsListApps

> CollectionOfManagedMobileApp DeviceAppManagementIosManagedAppProtectionsListApps(ctx, iosManagedAppProtectionId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

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
    iosManagedAppProtectionId := "iosManagedAppProtectionId_example" // string | key: id of iosManagedAppProtection
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
    resp, r, err := api_client.DeviceAppManagementIosManagedAppProtectionApi.DeviceAppManagementIosManagedAppProtectionsListApps(context.Background(), iosManagedAppProtectionId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAppManagementIosManagedAppProtectionApi.DeviceAppManagementIosManagedAppProtectionsListApps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceAppManagementIosManagedAppProtectionsListApps`: CollectionOfManagedMobileApp
    fmt.Fprintf(os.Stdout, "Response from `DeviceAppManagementIosManagedAppProtectionApi.DeviceAppManagementIosManagedAppProtectionsListApps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**iosManagedAppProtectionId** | **string** | key: id of iosManagedAppProtection | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceAppManagementIosManagedAppProtectionsListAppsRequest struct via the builder pattern


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


## DeviceAppManagementIosManagedAppProtectionsUpdateApps

> DeviceAppManagementIosManagedAppProtectionsUpdateApps(ctx, iosManagedAppProtectionId, managedMobileAppId).MicrosoftGraphManagedMobileApp(microsoftGraphManagedMobileApp).Execute()

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
    iosManagedAppProtectionId := "iosManagedAppProtectionId_example" // string | key: id of iosManagedAppProtection
    managedMobileAppId := "managedMobileAppId_example" // string | key: id of managedMobileApp
    microsoftGraphManagedMobileApp := *openapiclient.NewMicrosoftGraphManagedMobileApp() // MicrosoftGraphManagedMobileApp | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceAppManagementIosManagedAppProtectionApi.DeviceAppManagementIosManagedAppProtectionsUpdateApps(context.Background(), iosManagedAppProtectionId, managedMobileAppId).MicrosoftGraphManagedMobileApp(microsoftGraphManagedMobileApp).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAppManagementIosManagedAppProtectionApi.DeviceAppManagementIosManagedAppProtectionsUpdateApps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**iosManagedAppProtectionId** | **string** | key: id of iosManagedAppProtection | 
**managedMobileAppId** | **string** | key: id of managedMobileApp | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceAppManagementIosManagedAppProtectionsUpdateAppsRequest struct via the builder pattern


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


## DeviceAppManagementIosManagedAppProtectionsUpdateDeploymentSummary

> DeviceAppManagementIosManagedAppProtectionsUpdateDeploymentSummary(ctx, iosManagedAppProtectionId).MicrosoftGraphManagedAppPolicyDeploymentSummary(microsoftGraphManagedAppPolicyDeploymentSummary).Execute()

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
    iosManagedAppProtectionId := "iosManagedAppProtectionId_example" // string | key: id of iosManagedAppProtection
    microsoftGraphManagedAppPolicyDeploymentSummary := *openapiclient.NewMicrosoftGraphManagedAppPolicyDeploymentSummary() // MicrosoftGraphManagedAppPolicyDeploymentSummary | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceAppManagementIosManagedAppProtectionApi.DeviceAppManagementIosManagedAppProtectionsUpdateDeploymentSummary(context.Background(), iosManagedAppProtectionId).MicrosoftGraphManagedAppPolicyDeploymentSummary(microsoftGraphManagedAppPolicyDeploymentSummary).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAppManagementIosManagedAppProtectionApi.DeviceAppManagementIosManagedAppProtectionsUpdateDeploymentSummary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**iosManagedAppProtectionId** | **string** | key: id of iosManagedAppProtection | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceAppManagementIosManagedAppProtectionsUpdateDeploymentSummaryRequest struct via the builder pattern


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


## DeviceAppManagementListIosManagedAppProtections

> CollectionOfIosManagedAppProtection DeviceAppManagementListIosManagedAppProtections(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get iosManagedAppProtections from deviceAppManagement



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
    resp, r, err := api_client.DeviceAppManagementIosManagedAppProtectionApi.DeviceAppManagementListIosManagedAppProtections(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAppManagementIosManagedAppProtectionApi.DeviceAppManagementListIosManagedAppProtections``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceAppManagementListIosManagedAppProtections`: CollectionOfIosManagedAppProtection
    fmt.Fprintf(os.Stdout, "Response from `DeviceAppManagementIosManagedAppProtectionApi.DeviceAppManagementListIosManagedAppProtections`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeviceAppManagementListIosManagedAppProtectionsRequest struct via the builder pattern


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

[**CollectionOfIosManagedAppProtection**](CollectionOfIosManagedAppProtection.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementUpdateIosManagedAppProtections

> DeviceAppManagementUpdateIosManagedAppProtections(ctx, iosManagedAppProtectionId).MicrosoftGraphIosManagedAppProtection(microsoftGraphIosManagedAppProtection).Execute()

Update the navigation property iosManagedAppProtections in deviceAppManagement



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
    iosManagedAppProtectionId := "iosManagedAppProtectionId_example" // string | key: id of iosManagedAppProtection
    microsoftGraphIosManagedAppProtection := *openapiclient.NewMicrosoftGraphIosManagedAppProtection() // MicrosoftGraphIosManagedAppProtection | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceAppManagementIosManagedAppProtectionApi.DeviceAppManagementUpdateIosManagedAppProtections(context.Background(), iosManagedAppProtectionId).MicrosoftGraphIosManagedAppProtection(microsoftGraphIosManagedAppProtection).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAppManagementIosManagedAppProtectionApi.DeviceAppManagementUpdateIosManagedAppProtections``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**iosManagedAppProtectionId** | **string** | key: id of iosManagedAppProtection | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceAppManagementUpdateIosManagedAppProtectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphIosManagedAppProtection** | [**MicrosoftGraphIosManagedAppProtection**](MicrosoftGraphIosManagedAppProtection.md) | New navigation property values | 

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

