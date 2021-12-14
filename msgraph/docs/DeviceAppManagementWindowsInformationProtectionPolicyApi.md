# \DeviceAppManagementWindowsInformationProtectionPolicyApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceAppManagementCreateWindowsInformationProtectionPolicies**](DeviceAppManagementWindowsInformationProtectionPolicyApi.md#DeviceAppManagementCreateWindowsInformationProtectionPolicies) | **Post** /deviceAppManagement/windowsInformationProtectionPolicies | Create new navigation property to windowsInformationProtectionPolicies for deviceAppManagement
[**DeviceAppManagementDeleteWindowsInformationProtectionPolicies**](DeviceAppManagementWindowsInformationProtectionPolicyApi.md#DeviceAppManagementDeleteWindowsInformationProtectionPolicies) | **Delete** /deviceAppManagement/windowsInformationProtectionPolicies/{windowsInformationProtectionPolicy-id} | Delete navigation property windowsInformationProtectionPolicies for deviceAppManagement
[**DeviceAppManagementGetWindowsInformationProtectionPolicies**](DeviceAppManagementWindowsInformationProtectionPolicyApi.md#DeviceAppManagementGetWindowsInformationProtectionPolicies) | **Get** /deviceAppManagement/windowsInformationProtectionPolicies/{windowsInformationProtectionPolicy-id} | Get windowsInformationProtectionPolicies from deviceAppManagement
[**DeviceAppManagementListWindowsInformationProtectionPolicies**](DeviceAppManagementWindowsInformationProtectionPolicyApi.md#DeviceAppManagementListWindowsInformationProtectionPolicies) | **Get** /deviceAppManagement/windowsInformationProtectionPolicies | Get windowsInformationProtectionPolicies from deviceAppManagement
[**DeviceAppManagementUpdateWindowsInformationProtectionPolicies**](DeviceAppManagementWindowsInformationProtectionPolicyApi.md#DeviceAppManagementUpdateWindowsInformationProtectionPolicies) | **Patch** /deviceAppManagement/windowsInformationProtectionPolicies/{windowsInformationProtectionPolicy-id} | Update the navigation property windowsInformationProtectionPolicies in deviceAppManagement



## DeviceAppManagementCreateWindowsInformationProtectionPolicies

> MicrosoftGraphWindowsInformationProtectionPolicy DeviceAppManagementCreateWindowsInformationProtectionPolicies(ctx).MicrosoftGraphWindowsInformationProtectionPolicy(microsoftGraphWindowsInformationProtectionPolicy).Execute()

Create new navigation property to windowsInformationProtectionPolicies for deviceAppManagement



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
    microsoftGraphWindowsInformationProtectionPolicy := *openapiclient.NewMicrosoftGraphWindowsInformationProtectionPolicy() // MicrosoftGraphWindowsInformationProtectionPolicy | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceAppManagementWindowsInformationProtectionPolicyApi.DeviceAppManagementCreateWindowsInformationProtectionPolicies(context.Background()).MicrosoftGraphWindowsInformationProtectionPolicy(microsoftGraphWindowsInformationProtectionPolicy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAppManagementWindowsInformationProtectionPolicyApi.DeviceAppManagementCreateWindowsInformationProtectionPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceAppManagementCreateWindowsInformationProtectionPolicies`: MicrosoftGraphWindowsInformationProtectionPolicy
    fmt.Fprintf(os.Stdout, "Response from `DeviceAppManagementWindowsInformationProtectionPolicyApi.DeviceAppManagementCreateWindowsInformationProtectionPolicies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeviceAppManagementCreateWindowsInformationProtectionPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphWindowsInformationProtectionPolicy** | [**MicrosoftGraphWindowsInformationProtectionPolicy**](MicrosoftGraphWindowsInformationProtectionPolicy.md) | New navigation property | 

### Return type

[**MicrosoftGraphWindowsInformationProtectionPolicy**](MicrosoftGraphWindowsInformationProtectionPolicy.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementDeleteWindowsInformationProtectionPolicies

> DeviceAppManagementDeleteWindowsInformationProtectionPolicies(ctx, windowsInformationProtectionPolicyId).IfMatch(ifMatch).Execute()

Delete navigation property windowsInformationProtectionPolicies for deviceAppManagement



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
    windowsInformationProtectionPolicyId := "windowsInformationProtectionPolicyId_example" // string | key: id of windowsInformationProtectionPolicy
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceAppManagementWindowsInformationProtectionPolicyApi.DeviceAppManagementDeleteWindowsInformationProtectionPolicies(context.Background(), windowsInformationProtectionPolicyId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAppManagementWindowsInformationProtectionPolicyApi.DeviceAppManagementDeleteWindowsInformationProtectionPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**windowsInformationProtectionPolicyId** | **string** | key: id of windowsInformationProtectionPolicy | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceAppManagementDeleteWindowsInformationProtectionPoliciesRequest struct via the builder pattern


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


## DeviceAppManagementGetWindowsInformationProtectionPolicies

> MicrosoftGraphWindowsInformationProtectionPolicy DeviceAppManagementGetWindowsInformationProtectionPolicies(ctx, windowsInformationProtectionPolicyId).Select_(select_).Expand(expand).Execute()

Get windowsInformationProtectionPolicies from deviceAppManagement



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
    windowsInformationProtectionPolicyId := "windowsInformationProtectionPolicyId_example" // string | key: id of windowsInformationProtectionPolicy
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceAppManagementWindowsInformationProtectionPolicyApi.DeviceAppManagementGetWindowsInformationProtectionPolicies(context.Background(), windowsInformationProtectionPolicyId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAppManagementWindowsInformationProtectionPolicyApi.DeviceAppManagementGetWindowsInformationProtectionPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceAppManagementGetWindowsInformationProtectionPolicies`: MicrosoftGraphWindowsInformationProtectionPolicy
    fmt.Fprintf(os.Stdout, "Response from `DeviceAppManagementWindowsInformationProtectionPolicyApi.DeviceAppManagementGetWindowsInformationProtectionPolicies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**windowsInformationProtectionPolicyId** | **string** | key: id of windowsInformationProtectionPolicy | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceAppManagementGetWindowsInformationProtectionPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphWindowsInformationProtectionPolicy**](MicrosoftGraphWindowsInformationProtectionPolicy.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementListWindowsInformationProtectionPolicies

> CollectionOfWindowsInformationProtectionPolicy DeviceAppManagementListWindowsInformationProtectionPolicies(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get windowsInformationProtectionPolicies from deviceAppManagement



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
    resp, r, err := api_client.DeviceAppManagementWindowsInformationProtectionPolicyApi.DeviceAppManagementListWindowsInformationProtectionPolicies(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAppManagementWindowsInformationProtectionPolicyApi.DeviceAppManagementListWindowsInformationProtectionPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceAppManagementListWindowsInformationProtectionPolicies`: CollectionOfWindowsInformationProtectionPolicy
    fmt.Fprintf(os.Stdout, "Response from `DeviceAppManagementWindowsInformationProtectionPolicyApi.DeviceAppManagementListWindowsInformationProtectionPolicies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeviceAppManagementListWindowsInformationProtectionPoliciesRequest struct via the builder pattern


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

[**CollectionOfWindowsInformationProtectionPolicy**](CollectionOfWindowsInformationProtectionPolicy.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementUpdateWindowsInformationProtectionPolicies

> DeviceAppManagementUpdateWindowsInformationProtectionPolicies(ctx, windowsInformationProtectionPolicyId).MicrosoftGraphWindowsInformationProtectionPolicy(microsoftGraphWindowsInformationProtectionPolicy).Execute()

Update the navigation property windowsInformationProtectionPolicies in deviceAppManagement



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
    windowsInformationProtectionPolicyId := "windowsInformationProtectionPolicyId_example" // string | key: id of windowsInformationProtectionPolicy
    microsoftGraphWindowsInformationProtectionPolicy := *openapiclient.NewMicrosoftGraphWindowsInformationProtectionPolicy() // MicrosoftGraphWindowsInformationProtectionPolicy | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceAppManagementWindowsInformationProtectionPolicyApi.DeviceAppManagementUpdateWindowsInformationProtectionPolicies(context.Background(), windowsInformationProtectionPolicyId).MicrosoftGraphWindowsInformationProtectionPolicy(microsoftGraphWindowsInformationProtectionPolicy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAppManagementWindowsInformationProtectionPolicyApi.DeviceAppManagementUpdateWindowsInformationProtectionPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**windowsInformationProtectionPolicyId** | **string** | key: id of windowsInformationProtectionPolicy | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceAppManagementUpdateWindowsInformationProtectionPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphWindowsInformationProtectionPolicy** | [**MicrosoftGraphWindowsInformationProtectionPolicy**](MicrosoftGraphWindowsInformationProtectionPolicy.md) | New navigation property values | 

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

