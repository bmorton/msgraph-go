# \DeviceAppManagementManagedAppPolicyApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceAppManagementCreateManagedAppPolicies**](DeviceAppManagementManagedAppPolicyApi.md#DeviceAppManagementCreateManagedAppPolicies) | **Post** /deviceAppManagement/managedAppPolicies | Create new navigation property to managedAppPolicies for deviceAppManagement
[**DeviceAppManagementDeleteManagedAppPolicies**](DeviceAppManagementManagedAppPolicyApi.md#DeviceAppManagementDeleteManagedAppPolicies) | **Delete** /deviceAppManagement/managedAppPolicies/{managedAppPolicy-id} | Delete navigation property managedAppPolicies for deviceAppManagement
[**DeviceAppManagementGetManagedAppPolicies**](DeviceAppManagementManagedAppPolicyApi.md#DeviceAppManagementGetManagedAppPolicies) | **Get** /deviceAppManagement/managedAppPolicies/{managedAppPolicy-id} | Get managedAppPolicies from deviceAppManagement
[**DeviceAppManagementListManagedAppPolicies**](DeviceAppManagementManagedAppPolicyApi.md#DeviceAppManagementListManagedAppPolicies) | **Get** /deviceAppManagement/managedAppPolicies | Get managedAppPolicies from deviceAppManagement
[**DeviceAppManagementUpdateManagedAppPolicies**](DeviceAppManagementManagedAppPolicyApi.md#DeviceAppManagementUpdateManagedAppPolicies) | **Patch** /deviceAppManagement/managedAppPolicies/{managedAppPolicy-id} | Update the navigation property managedAppPolicies in deviceAppManagement



## DeviceAppManagementCreateManagedAppPolicies

> MicrosoftGraphManagedAppPolicy DeviceAppManagementCreateManagedAppPolicies(ctx).MicrosoftGraphManagedAppPolicy(microsoftGraphManagedAppPolicy).Execute()

Create new navigation property to managedAppPolicies for deviceAppManagement



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
    microsoftGraphManagedAppPolicy := *openapiclient.NewMicrosoftGraphManagedAppPolicy() // MicrosoftGraphManagedAppPolicy | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceAppManagementManagedAppPolicyApi.DeviceAppManagementCreateManagedAppPolicies(context.Background()).MicrosoftGraphManagedAppPolicy(microsoftGraphManagedAppPolicy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAppManagementManagedAppPolicyApi.DeviceAppManagementCreateManagedAppPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceAppManagementCreateManagedAppPolicies`: MicrosoftGraphManagedAppPolicy
    fmt.Fprintf(os.Stdout, "Response from `DeviceAppManagementManagedAppPolicyApi.DeviceAppManagementCreateManagedAppPolicies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeviceAppManagementCreateManagedAppPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphManagedAppPolicy** | [**MicrosoftGraphManagedAppPolicy**](MicrosoftGraphManagedAppPolicy.md) | New navigation property | 

### Return type

[**MicrosoftGraphManagedAppPolicy**](MicrosoftGraphManagedAppPolicy.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementDeleteManagedAppPolicies

> DeviceAppManagementDeleteManagedAppPolicies(ctx, managedAppPolicyId).IfMatch(ifMatch).Execute()

Delete navigation property managedAppPolicies for deviceAppManagement



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
    managedAppPolicyId := "managedAppPolicyId_example" // string | key: id of managedAppPolicy
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceAppManagementManagedAppPolicyApi.DeviceAppManagementDeleteManagedAppPolicies(context.Background(), managedAppPolicyId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAppManagementManagedAppPolicyApi.DeviceAppManagementDeleteManagedAppPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managedAppPolicyId** | **string** | key: id of managedAppPolicy | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceAppManagementDeleteManagedAppPoliciesRequest struct via the builder pattern


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


## DeviceAppManagementGetManagedAppPolicies

> MicrosoftGraphManagedAppPolicy DeviceAppManagementGetManagedAppPolicies(ctx, managedAppPolicyId).Select_(select_).Expand(expand).Execute()

Get managedAppPolicies from deviceAppManagement



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
    managedAppPolicyId := "managedAppPolicyId_example" // string | key: id of managedAppPolicy
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceAppManagementManagedAppPolicyApi.DeviceAppManagementGetManagedAppPolicies(context.Background(), managedAppPolicyId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAppManagementManagedAppPolicyApi.DeviceAppManagementGetManagedAppPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceAppManagementGetManagedAppPolicies`: MicrosoftGraphManagedAppPolicy
    fmt.Fprintf(os.Stdout, "Response from `DeviceAppManagementManagedAppPolicyApi.DeviceAppManagementGetManagedAppPolicies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managedAppPolicyId** | **string** | key: id of managedAppPolicy | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceAppManagementGetManagedAppPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphManagedAppPolicy**](MicrosoftGraphManagedAppPolicy.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementListManagedAppPolicies

> CollectionOfManagedAppPolicy DeviceAppManagementListManagedAppPolicies(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get managedAppPolicies from deviceAppManagement



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
    resp, r, err := api_client.DeviceAppManagementManagedAppPolicyApi.DeviceAppManagementListManagedAppPolicies(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAppManagementManagedAppPolicyApi.DeviceAppManagementListManagedAppPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceAppManagementListManagedAppPolicies`: CollectionOfManagedAppPolicy
    fmt.Fprintf(os.Stdout, "Response from `DeviceAppManagementManagedAppPolicyApi.DeviceAppManagementListManagedAppPolicies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeviceAppManagementListManagedAppPoliciesRequest struct via the builder pattern


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

[**CollectionOfManagedAppPolicy**](CollectionOfManagedAppPolicy.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementUpdateManagedAppPolicies

> DeviceAppManagementUpdateManagedAppPolicies(ctx, managedAppPolicyId).MicrosoftGraphManagedAppPolicy(microsoftGraphManagedAppPolicy).Execute()

Update the navigation property managedAppPolicies in deviceAppManagement



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
    managedAppPolicyId := "managedAppPolicyId_example" // string | key: id of managedAppPolicy
    microsoftGraphManagedAppPolicy := *openapiclient.NewMicrosoftGraphManagedAppPolicy() // MicrosoftGraphManagedAppPolicy | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceAppManagementManagedAppPolicyApi.DeviceAppManagementUpdateManagedAppPolicies(context.Background(), managedAppPolicyId).MicrosoftGraphManagedAppPolicy(microsoftGraphManagedAppPolicy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAppManagementManagedAppPolicyApi.DeviceAppManagementUpdateManagedAppPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**managedAppPolicyId** | **string** | key: id of managedAppPolicy | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceAppManagementUpdateManagedAppPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphManagedAppPolicy** | [**MicrosoftGraphManagedAppPolicy**](MicrosoftGraphManagedAppPolicy.md) | New navigation property values | 

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

