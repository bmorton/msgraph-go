# \DeviceAppManagementDefaultManagedAppProtectionApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceAppManagementCreateDefaultManagedAppProtections**](DeviceAppManagementDefaultManagedAppProtectionApi.md#DeviceAppManagementCreateDefaultManagedAppProtections) | **Post** /deviceAppManagement/defaultManagedAppProtections | Create new navigation property to defaultManagedAppProtections for deviceAppManagement
[**DeviceAppManagementDefaultManagedAppProtectionsCreateApps**](DeviceAppManagementDefaultManagedAppProtectionApi.md#DeviceAppManagementDefaultManagedAppProtectionsCreateApps) | **Post** /deviceAppManagement/defaultManagedAppProtections/{defaultManagedAppProtection-id}/apps | Create new navigation property to apps for deviceAppManagement
[**DeviceAppManagementDefaultManagedAppProtectionsDeleteApps**](DeviceAppManagementDefaultManagedAppProtectionApi.md#DeviceAppManagementDefaultManagedAppProtectionsDeleteApps) | **Delete** /deviceAppManagement/defaultManagedAppProtections/{defaultManagedAppProtection-id}/apps/{managedMobileApp-id} | Delete navigation property apps for deviceAppManagement
[**DeviceAppManagementDefaultManagedAppProtectionsDeleteDeploymentSummary**](DeviceAppManagementDefaultManagedAppProtectionApi.md#DeviceAppManagementDefaultManagedAppProtectionsDeleteDeploymentSummary) | **Delete** /deviceAppManagement/defaultManagedAppProtections/{defaultManagedAppProtection-id}/deploymentSummary | Delete navigation property deploymentSummary for deviceAppManagement
[**DeviceAppManagementDefaultManagedAppProtectionsGetApps**](DeviceAppManagementDefaultManagedAppProtectionApi.md#DeviceAppManagementDefaultManagedAppProtectionsGetApps) | **Get** /deviceAppManagement/defaultManagedAppProtections/{defaultManagedAppProtection-id}/apps/{managedMobileApp-id} | Get apps from deviceAppManagement
[**DeviceAppManagementDefaultManagedAppProtectionsGetDeploymentSummary**](DeviceAppManagementDefaultManagedAppProtectionApi.md#DeviceAppManagementDefaultManagedAppProtectionsGetDeploymentSummary) | **Get** /deviceAppManagement/defaultManagedAppProtections/{defaultManagedAppProtection-id}/deploymentSummary | Get deploymentSummary from deviceAppManagement
[**DeviceAppManagementDefaultManagedAppProtectionsListApps**](DeviceAppManagementDefaultManagedAppProtectionApi.md#DeviceAppManagementDefaultManagedAppProtectionsListApps) | **Get** /deviceAppManagement/defaultManagedAppProtections/{defaultManagedAppProtection-id}/apps | Get apps from deviceAppManagement
[**DeviceAppManagementDefaultManagedAppProtectionsUpdateApps**](DeviceAppManagementDefaultManagedAppProtectionApi.md#DeviceAppManagementDefaultManagedAppProtectionsUpdateApps) | **Patch** /deviceAppManagement/defaultManagedAppProtections/{defaultManagedAppProtection-id}/apps/{managedMobileApp-id} | Update the navigation property apps in deviceAppManagement
[**DeviceAppManagementDefaultManagedAppProtectionsUpdateDeploymentSummary**](DeviceAppManagementDefaultManagedAppProtectionApi.md#DeviceAppManagementDefaultManagedAppProtectionsUpdateDeploymentSummary) | **Patch** /deviceAppManagement/defaultManagedAppProtections/{defaultManagedAppProtection-id}/deploymentSummary | Update the navigation property deploymentSummary in deviceAppManagement
[**DeviceAppManagementDeleteDefaultManagedAppProtections**](DeviceAppManagementDefaultManagedAppProtectionApi.md#DeviceAppManagementDeleteDefaultManagedAppProtections) | **Delete** /deviceAppManagement/defaultManagedAppProtections/{defaultManagedAppProtection-id} | Delete navigation property defaultManagedAppProtections for deviceAppManagement
[**DeviceAppManagementGetDefaultManagedAppProtections**](DeviceAppManagementDefaultManagedAppProtectionApi.md#DeviceAppManagementGetDefaultManagedAppProtections) | **Get** /deviceAppManagement/defaultManagedAppProtections/{defaultManagedAppProtection-id} | Get defaultManagedAppProtections from deviceAppManagement
[**DeviceAppManagementListDefaultManagedAppProtections**](DeviceAppManagementDefaultManagedAppProtectionApi.md#DeviceAppManagementListDefaultManagedAppProtections) | **Get** /deviceAppManagement/defaultManagedAppProtections | Get defaultManagedAppProtections from deviceAppManagement
[**DeviceAppManagementUpdateDefaultManagedAppProtections**](DeviceAppManagementDefaultManagedAppProtectionApi.md#DeviceAppManagementUpdateDefaultManagedAppProtections) | **Patch** /deviceAppManagement/defaultManagedAppProtections/{defaultManagedAppProtection-id} | Update the navigation property defaultManagedAppProtections in deviceAppManagement



## DeviceAppManagementCreateDefaultManagedAppProtections

> MicrosoftGraphDefaultManagedAppProtection DeviceAppManagementCreateDefaultManagedAppProtections(ctx).MicrosoftGraphDefaultManagedAppProtection(microsoftGraphDefaultManagedAppProtection).Execute()

Create new navigation property to defaultManagedAppProtections for deviceAppManagement



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
    microsoftGraphDefaultManagedAppProtection := *openapiclient.NewMicrosoftGraphDefaultManagedAppProtection() // MicrosoftGraphDefaultManagedAppProtection | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceAppManagementDefaultManagedAppProtectionApi.DeviceAppManagementCreateDefaultManagedAppProtections(context.Background()).MicrosoftGraphDefaultManagedAppProtection(microsoftGraphDefaultManagedAppProtection).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAppManagementDefaultManagedAppProtectionApi.DeviceAppManagementCreateDefaultManagedAppProtections``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceAppManagementCreateDefaultManagedAppProtections`: MicrosoftGraphDefaultManagedAppProtection
    fmt.Fprintf(os.Stdout, "Response from `DeviceAppManagementDefaultManagedAppProtectionApi.DeviceAppManagementCreateDefaultManagedAppProtections`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeviceAppManagementCreateDefaultManagedAppProtectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphDefaultManagedAppProtection** | [**MicrosoftGraphDefaultManagedAppProtection**](MicrosoftGraphDefaultManagedAppProtection.md) | New navigation property | 

### Return type

[**MicrosoftGraphDefaultManagedAppProtection**](MicrosoftGraphDefaultManagedAppProtection.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementDefaultManagedAppProtectionsCreateApps

> MicrosoftGraphManagedMobileApp DeviceAppManagementDefaultManagedAppProtectionsCreateApps(ctx, defaultManagedAppProtectionId).MicrosoftGraphManagedMobileApp(microsoftGraphManagedMobileApp).Execute()

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
    defaultManagedAppProtectionId := "defaultManagedAppProtectionId_example" // string | key: id of defaultManagedAppProtection
    microsoftGraphManagedMobileApp := *openapiclient.NewMicrosoftGraphManagedMobileApp() // MicrosoftGraphManagedMobileApp | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceAppManagementDefaultManagedAppProtectionApi.DeviceAppManagementDefaultManagedAppProtectionsCreateApps(context.Background(), defaultManagedAppProtectionId).MicrosoftGraphManagedMobileApp(microsoftGraphManagedMobileApp).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAppManagementDefaultManagedAppProtectionApi.DeviceAppManagementDefaultManagedAppProtectionsCreateApps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceAppManagementDefaultManagedAppProtectionsCreateApps`: MicrosoftGraphManagedMobileApp
    fmt.Fprintf(os.Stdout, "Response from `DeviceAppManagementDefaultManagedAppProtectionApi.DeviceAppManagementDefaultManagedAppProtectionsCreateApps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**defaultManagedAppProtectionId** | **string** | key: id of defaultManagedAppProtection | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceAppManagementDefaultManagedAppProtectionsCreateAppsRequest struct via the builder pattern


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


## DeviceAppManagementDefaultManagedAppProtectionsDeleteApps

> DeviceAppManagementDefaultManagedAppProtectionsDeleteApps(ctx, defaultManagedAppProtectionId, managedMobileAppId).IfMatch(ifMatch).Execute()

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
    defaultManagedAppProtectionId := "defaultManagedAppProtectionId_example" // string | key: id of defaultManagedAppProtection
    managedMobileAppId := "managedMobileAppId_example" // string | key: id of managedMobileApp
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceAppManagementDefaultManagedAppProtectionApi.DeviceAppManagementDefaultManagedAppProtectionsDeleteApps(context.Background(), defaultManagedAppProtectionId, managedMobileAppId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAppManagementDefaultManagedAppProtectionApi.DeviceAppManagementDefaultManagedAppProtectionsDeleteApps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**defaultManagedAppProtectionId** | **string** | key: id of defaultManagedAppProtection | 
**managedMobileAppId** | **string** | key: id of managedMobileApp | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceAppManagementDefaultManagedAppProtectionsDeleteAppsRequest struct via the builder pattern


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


## DeviceAppManagementDefaultManagedAppProtectionsDeleteDeploymentSummary

> DeviceAppManagementDefaultManagedAppProtectionsDeleteDeploymentSummary(ctx, defaultManagedAppProtectionId).IfMatch(ifMatch).Execute()

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
    defaultManagedAppProtectionId := "defaultManagedAppProtectionId_example" // string | key: id of defaultManagedAppProtection
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceAppManagementDefaultManagedAppProtectionApi.DeviceAppManagementDefaultManagedAppProtectionsDeleteDeploymentSummary(context.Background(), defaultManagedAppProtectionId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAppManagementDefaultManagedAppProtectionApi.DeviceAppManagementDefaultManagedAppProtectionsDeleteDeploymentSummary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**defaultManagedAppProtectionId** | **string** | key: id of defaultManagedAppProtection | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceAppManagementDefaultManagedAppProtectionsDeleteDeploymentSummaryRequest struct via the builder pattern


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


## DeviceAppManagementDefaultManagedAppProtectionsGetApps

> MicrosoftGraphManagedMobileApp DeviceAppManagementDefaultManagedAppProtectionsGetApps(ctx, defaultManagedAppProtectionId, managedMobileAppId).Select_(select_).Expand(expand).Execute()

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
    defaultManagedAppProtectionId := "defaultManagedAppProtectionId_example" // string | key: id of defaultManagedAppProtection
    managedMobileAppId := "managedMobileAppId_example" // string | key: id of managedMobileApp
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceAppManagementDefaultManagedAppProtectionApi.DeviceAppManagementDefaultManagedAppProtectionsGetApps(context.Background(), defaultManagedAppProtectionId, managedMobileAppId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAppManagementDefaultManagedAppProtectionApi.DeviceAppManagementDefaultManagedAppProtectionsGetApps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceAppManagementDefaultManagedAppProtectionsGetApps`: MicrosoftGraphManagedMobileApp
    fmt.Fprintf(os.Stdout, "Response from `DeviceAppManagementDefaultManagedAppProtectionApi.DeviceAppManagementDefaultManagedAppProtectionsGetApps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**defaultManagedAppProtectionId** | **string** | key: id of defaultManagedAppProtection | 
**managedMobileAppId** | **string** | key: id of managedMobileApp | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceAppManagementDefaultManagedAppProtectionsGetAppsRequest struct via the builder pattern


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


## DeviceAppManagementDefaultManagedAppProtectionsGetDeploymentSummary

> MicrosoftGraphManagedAppPolicyDeploymentSummary DeviceAppManagementDefaultManagedAppProtectionsGetDeploymentSummary(ctx, defaultManagedAppProtectionId).Select_(select_).Expand(expand).Execute()

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
    defaultManagedAppProtectionId := "defaultManagedAppProtectionId_example" // string | key: id of defaultManagedAppProtection
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceAppManagementDefaultManagedAppProtectionApi.DeviceAppManagementDefaultManagedAppProtectionsGetDeploymentSummary(context.Background(), defaultManagedAppProtectionId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAppManagementDefaultManagedAppProtectionApi.DeviceAppManagementDefaultManagedAppProtectionsGetDeploymentSummary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceAppManagementDefaultManagedAppProtectionsGetDeploymentSummary`: MicrosoftGraphManagedAppPolicyDeploymentSummary
    fmt.Fprintf(os.Stdout, "Response from `DeviceAppManagementDefaultManagedAppProtectionApi.DeviceAppManagementDefaultManagedAppProtectionsGetDeploymentSummary`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**defaultManagedAppProtectionId** | **string** | key: id of defaultManagedAppProtection | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceAppManagementDefaultManagedAppProtectionsGetDeploymentSummaryRequest struct via the builder pattern


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


## DeviceAppManagementDefaultManagedAppProtectionsListApps

> CollectionOfManagedMobileApp DeviceAppManagementDefaultManagedAppProtectionsListApps(ctx, defaultManagedAppProtectionId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

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
    defaultManagedAppProtectionId := "defaultManagedAppProtectionId_example" // string | key: id of defaultManagedAppProtection
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
    resp, r, err := api_client.DeviceAppManagementDefaultManagedAppProtectionApi.DeviceAppManagementDefaultManagedAppProtectionsListApps(context.Background(), defaultManagedAppProtectionId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAppManagementDefaultManagedAppProtectionApi.DeviceAppManagementDefaultManagedAppProtectionsListApps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceAppManagementDefaultManagedAppProtectionsListApps`: CollectionOfManagedMobileApp
    fmt.Fprintf(os.Stdout, "Response from `DeviceAppManagementDefaultManagedAppProtectionApi.DeviceAppManagementDefaultManagedAppProtectionsListApps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**defaultManagedAppProtectionId** | **string** | key: id of defaultManagedAppProtection | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceAppManagementDefaultManagedAppProtectionsListAppsRequest struct via the builder pattern


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


## DeviceAppManagementDefaultManagedAppProtectionsUpdateApps

> DeviceAppManagementDefaultManagedAppProtectionsUpdateApps(ctx, defaultManagedAppProtectionId, managedMobileAppId).MicrosoftGraphManagedMobileApp(microsoftGraphManagedMobileApp).Execute()

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
    defaultManagedAppProtectionId := "defaultManagedAppProtectionId_example" // string | key: id of defaultManagedAppProtection
    managedMobileAppId := "managedMobileAppId_example" // string | key: id of managedMobileApp
    microsoftGraphManagedMobileApp := *openapiclient.NewMicrosoftGraphManagedMobileApp() // MicrosoftGraphManagedMobileApp | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceAppManagementDefaultManagedAppProtectionApi.DeviceAppManagementDefaultManagedAppProtectionsUpdateApps(context.Background(), defaultManagedAppProtectionId, managedMobileAppId).MicrosoftGraphManagedMobileApp(microsoftGraphManagedMobileApp).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAppManagementDefaultManagedAppProtectionApi.DeviceAppManagementDefaultManagedAppProtectionsUpdateApps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**defaultManagedAppProtectionId** | **string** | key: id of defaultManagedAppProtection | 
**managedMobileAppId** | **string** | key: id of managedMobileApp | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceAppManagementDefaultManagedAppProtectionsUpdateAppsRequest struct via the builder pattern


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


## DeviceAppManagementDefaultManagedAppProtectionsUpdateDeploymentSummary

> DeviceAppManagementDefaultManagedAppProtectionsUpdateDeploymentSummary(ctx, defaultManagedAppProtectionId).MicrosoftGraphManagedAppPolicyDeploymentSummary(microsoftGraphManagedAppPolicyDeploymentSummary).Execute()

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
    defaultManagedAppProtectionId := "defaultManagedAppProtectionId_example" // string | key: id of defaultManagedAppProtection
    microsoftGraphManagedAppPolicyDeploymentSummary := *openapiclient.NewMicrosoftGraphManagedAppPolicyDeploymentSummary() // MicrosoftGraphManagedAppPolicyDeploymentSummary | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceAppManagementDefaultManagedAppProtectionApi.DeviceAppManagementDefaultManagedAppProtectionsUpdateDeploymentSummary(context.Background(), defaultManagedAppProtectionId).MicrosoftGraphManagedAppPolicyDeploymentSummary(microsoftGraphManagedAppPolicyDeploymentSummary).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAppManagementDefaultManagedAppProtectionApi.DeviceAppManagementDefaultManagedAppProtectionsUpdateDeploymentSummary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**defaultManagedAppProtectionId** | **string** | key: id of defaultManagedAppProtection | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceAppManagementDefaultManagedAppProtectionsUpdateDeploymentSummaryRequest struct via the builder pattern


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


## DeviceAppManagementDeleteDefaultManagedAppProtections

> DeviceAppManagementDeleteDefaultManagedAppProtections(ctx, defaultManagedAppProtectionId).IfMatch(ifMatch).Execute()

Delete navigation property defaultManagedAppProtections for deviceAppManagement



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
    defaultManagedAppProtectionId := "defaultManagedAppProtectionId_example" // string | key: id of defaultManagedAppProtection
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceAppManagementDefaultManagedAppProtectionApi.DeviceAppManagementDeleteDefaultManagedAppProtections(context.Background(), defaultManagedAppProtectionId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAppManagementDefaultManagedAppProtectionApi.DeviceAppManagementDeleteDefaultManagedAppProtections``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**defaultManagedAppProtectionId** | **string** | key: id of defaultManagedAppProtection | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceAppManagementDeleteDefaultManagedAppProtectionsRequest struct via the builder pattern


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


## DeviceAppManagementGetDefaultManagedAppProtections

> MicrosoftGraphDefaultManagedAppProtection DeviceAppManagementGetDefaultManagedAppProtections(ctx, defaultManagedAppProtectionId).Select_(select_).Expand(expand).Execute()

Get defaultManagedAppProtections from deviceAppManagement



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
    defaultManagedAppProtectionId := "defaultManagedAppProtectionId_example" // string | key: id of defaultManagedAppProtection
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceAppManagementDefaultManagedAppProtectionApi.DeviceAppManagementGetDefaultManagedAppProtections(context.Background(), defaultManagedAppProtectionId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAppManagementDefaultManagedAppProtectionApi.DeviceAppManagementGetDefaultManagedAppProtections``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceAppManagementGetDefaultManagedAppProtections`: MicrosoftGraphDefaultManagedAppProtection
    fmt.Fprintf(os.Stdout, "Response from `DeviceAppManagementDefaultManagedAppProtectionApi.DeviceAppManagementGetDefaultManagedAppProtections`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**defaultManagedAppProtectionId** | **string** | key: id of defaultManagedAppProtection | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceAppManagementGetDefaultManagedAppProtectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphDefaultManagedAppProtection**](MicrosoftGraphDefaultManagedAppProtection.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementListDefaultManagedAppProtections

> CollectionOfDefaultManagedAppProtection DeviceAppManagementListDefaultManagedAppProtections(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get defaultManagedAppProtections from deviceAppManagement



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
    resp, r, err := api_client.DeviceAppManagementDefaultManagedAppProtectionApi.DeviceAppManagementListDefaultManagedAppProtections(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAppManagementDefaultManagedAppProtectionApi.DeviceAppManagementListDefaultManagedAppProtections``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceAppManagementListDefaultManagedAppProtections`: CollectionOfDefaultManagedAppProtection
    fmt.Fprintf(os.Stdout, "Response from `DeviceAppManagementDefaultManagedAppProtectionApi.DeviceAppManagementListDefaultManagedAppProtections`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeviceAppManagementListDefaultManagedAppProtectionsRequest struct via the builder pattern


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

[**CollectionOfDefaultManagedAppProtection**](CollectionOfDefaultManagedAppProtection.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementUpdateDefaultManagedAppProtections

> DeviceAppManagementUpdateDefaultManagedAppProtections(ctx, defaultManagedAppProtectionId).MicrosoftGraphDefaultManagedAppProtection(microsoftGraphDefaultManagedAppProtection).Execute()

Update the navigation property defaultManagedAppProtections in deviceAppManagement



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
    defaultManagedAppProtectionId := "defaultManagedAppProtectionId_example" // string | key: id of defaultManagedAppProtection
    microsoftGraphDefaultManagedAppProtection := *openapiclient.NewMicrosoftGraphDefaultManagedAppProtection() // MicrosoftGraphDefaultManagedAppProtection | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceAppManagementDefaultManagedAppProtectionApi.DeviceAppManagementUpdateDefaultManagedAppProtections(context.Background(), defaultManagedAppProtectionId).MicrosoftGraphDefaultManagedAppProtection(microsoftGraphDefaultManagedAppProtection).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAppManagementDefaultManagedAppProtectionApi.DeviceAppManagementUpdateDefaultManagedAppProtections``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**defaultManagedAppProtectionId** | **string** | key: id of defaultManagedAppProtection | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceAppManagementUpdateDefaultManagedAppProtectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphDefaultManagedAppProtection** | [**MicrosoftGraphDefaultManagedAppProtection**](MicrosoftGraphDefaultManagedAppProtection.md) | New navigation property values | 

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

