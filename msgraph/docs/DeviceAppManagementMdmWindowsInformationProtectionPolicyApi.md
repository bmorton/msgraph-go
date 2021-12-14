# \DeviceAppManagementMdmWindowsInformationProtectionPolicyApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceAppManagementCreateMdmWindowsInformationProtectionPolicies**](DeviceAppManagementMdmWindowsInformationProtectionPolicyApi.md#DeviceAppManagementCreateMdmWindowsInformationProtectionPolicies) | **Post** /deviceAppManagement/mdmWindowsInformationProtectionPolicies | Create new navigation property to mdmWindowsInformationProtectionPolicies for deviceAppManagement
[**DeviceAppManagementDeleteMdmWindowsInformationProtectionPolicies**](DeviceAppManagementMdmWindowsInformationProtectionPolicyApi.md#DeviceAppManagementDeleteMdmWindowsInformationProtectionPolicies) | **Delete** /deviceAppManagement/mdmWindowsInformationProtectionPolicies/{mdmWindowsInformationProtectionPolicy-id} | Delete navigation property mdmWindowsInformationProtectionPolicies for deviceAppManagement
[**DeviceAppManagementGetMdmWindowsInformationProtectionPolicies**](DeviceAppManagementMdmWindowsInformationProtectionPolicyApi.md#DeviceAppManagementGetMdmWindowsInformationProtectionPolicies) | **Get** /deviceAppManagement/mdmWindowsInformationProtectionPolicies/{mdmWindowsInformationProtectionPolicy-id} | Get mdmWindowsInformationProtectionPolicies from deviceAppManagement
[**DeviceAppManagementListMdmWindowsInformationProtectionPolicies**](DeviceAppManagementMdmWindowsInformationProtectionPolicyApi.md#DeviceAppManagementListMdmWindowsInformationProtectionPolicies) | **Get** /deviceAppManagement/mdmWindowsInformationProtectionPolicies | Get mdmWindowsInformationProtectionPolicies from deviceAppManagement
[**DeviceAppManagementUpdateMdmWindowsInformationProtectionPolicies**](DeviceAppManagementMdmWindowsInformationProtectionPolicyApi.md#DeviceAppManagementUpdateMdmWindowsInformationProtectionPolicies) | **Patch** /deviceAppManagement/mdmWindowsInformationProtectionPolicies/{mdmWindowsInformationProtectionPolicy-id} | Update the navigation property mdmWindowsInformationProtectionPolicies in deviceAppManagement



## DeviceAppManagementCreateMdmWindowsInformationProtectionPolicies

> MicrosoftGraphMdmWindowsInformationProtectionPolicy DeviceAppManagementCreateMdmWindowsInformationProtectionPolicies(ctx).MicrosoftGraphMdmWindowsInformationProtectionPolicy(microsoftGraphMdmWindowsInformationProtectionPolicy).Execute()

Create new navigation property to mdmWindowsInformationProtectionPolicies for deviceAppManagement



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
    microsoftGraphMdmWindowsInformationProtectionPolicy := *openapiclient.NewMicrosoftGraphMdmWindowsInformationProtectionPolicy() // MicrosoftGraphMdmWindowsInformationProtectionPolicy | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceAppManagementMdmWindowsInformationProtectionPolicyApi.DeviceAppManagementCreateMdmWindowsInformationProtectionPolicies(context.Background()).MicrosoftGraphMdmWindowsInformationProtectionPolicy(microsoftGraphMdmWindowsInformationProtectionPolicy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAppManagementMdmWindowsInformationProtectionPolicyApi.DeviceAppManagementCreateMdmWindowsInformationProtectionPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceAppManagementCreateMdmWindowsInformationProtectionPolicies`: MicrosoftGraphMdmWindowsInformationProtectionPolicy
    fmt.Fprintf(os.Stdout, "Response from `DeviceAppManagementMdmWindowsInformationProtectionPolicyApi.DeviceAppManagementCreateMdmWindowsInformationProtectionPolicies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeviceAppManagementCreateMdmWindowsInformationProtectionPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphMdmWindowsInformationProtectionPolicy** | [**MicrosoftGraphMdmWindowsInformationProtectionPolicy**](MicrosoftGraphMdmWindowsInformationProtectionPolicy.md) | New navigation property | 

### Return type

[**MicrosoftGraphMdmWindowsInformationProtectionPolicy**](MicrosoftGraphMdmWindowsInformationProtectionPolicy.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementDeleteMdmWindowsInformationProtectionPolicies

> DeviceAppManagementDeleteMdmWindowsInformationProtectionPolicies(ctx, mdmWindowsInformationProtectionPolicyId).IfMatch(ifMatch).Execute()

Delete navigation property mdmWindowsInformationProtectionPolicies for deviceAppManagement



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
    mdmWindowsInformationProtectionPolicyId := "mdmWindowsInformationProtectionPolicyId_example" // string | key: id of mdmWindowsInformationProtectionPolicy
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceAppManagementMdmWindowsInformationProtectionPolicyApi.DeviceAppManagementDeleteMdmWindowsInformationProtectionPolicies(context.Background(), mdmWindowsInformationProtectionPolicyId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAppManagementMdmWindowsInformationProtectionPolicyApi.DeviceAppManagementDeleteMdmWindowsInformationProtectionPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mdmWindowsInformationProtectionPolicyId** | **string** | key: id of mdmWindowsInformationProtectionPolicy | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceAppManagementDeleteMdmWindowsInformationProtectionPoliciesRequest struct via the builder pattern


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


## DeviceAppManagementGetMdmWindowsInformationProtectionPolicies

> MicrosoftGraphMdmWindowsInformationProtectionPolicy DeviceAppManagementGetMdmWindowsInformationProtectionPolicies(ctx, mdmWindowsInformationProtectionPolicyId).Select_(select_).Expand(expand).Execute()

Get mdmWindowsInformationProtectionPolicies from deviceAppManagement



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
    mdmWindowsInformationProtectionPolicyId := "mdmWindowsInformationProtectionPolicyId_example" // string | key: id of mdmWindowsInformationProtectionPolicy
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceAppManagementMdmWindowsInformationProtectionPolicyApi.DeviceAppManagementGetMdmWindowsInformationProtectionPolicies(context.Background(), mdmWindowsInformationProtectionPolicyId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAppManagementMdmWindowsInformationProtectionPolicyApi.DeviceAppManagementGetMdmWindowsInformationProtectionPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceAppManagementGetMdmWindowsInformationProtectionPolicies`: MicrosoftGraphMdmWindowsInformationProtectionPolicy
    fmt.Fprintf(os.Stdout, "Response from `DeviceAppManagementMdmWindowsInformationProtectionPolicyApi.DeviceAppManagementGetMdmWindowsInformationProtectionPolicies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mdmWindowsInformationProtectionPolicyId** | **string** | key: id of mdmWindowsInformationProtectionPolicy | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceAppManagementGetMdmWindowsInformationProtectionPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphMdmWindowsInformationProtectionPolicy**](MicrosoftGraphMdmWindowsInformationProtectionPolicy.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementListMdmWindowsInformationProtectionPolicies

> CollectionOfMdmWindowsInformationProtectionPolicy DeviceAppManagementListMdmWindowsInformationProtectionPolicies(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get mdmWindowsInformationProtectionPolicies from deviceAppManagement



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
    resp, r, err := api_client.DeviceAppManagementMdmWindowsInformationProtectionPolicyApi.DeviceAppManagementListMdmWindowsInformationProtectionPolicies(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAppManagementMdmWindowsInformationProtectionPolicyApi.DeviceAppManagementListMdmWindowsInformationProtectionPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceAppManagementListMdmWindowsInformationProtectionPolicies`: CollectionOfMdmWindowsInformationProtectionPolicy
    fmt.Fprintf(os.Stdout, "Response from `DeviceAppManagementMdmWindowsInformationProtectionPolicyApi.DeviceAppManagementListMdmWindowsInformationProtectionPolicies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeviceAppManagementListMdmWindowsInformationProtectionPoliciesRequest struct via the builder pattern


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

[**CollectionOfMdmWindowsInformationProtectionPolicy**](CollectionOfMdmWindowsInformationProtectionPolicy.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementUpdateMdmWindowsInformationProtectionPolicies

> DeviceAppManagementUpdateMdmWindowsInformationProtectionPolicies(ctx, mdmWindowsInformationProtectionPolicyId).MicrosoftGraphMdmWindowsInformationProtectionPolicy(microsoftGraphMdmWindowsInformationProtectionPolicy).Execute()

Update the navigation property mdmWindowsInformationProtectionPolicies in deviceAppManagement



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
    mdmWindowsInformationProtectionPolicyId := "mdmWindowsInformationProtectionPolicyId_example" // string | key: id of mdmWindowsInformationProtectionPolicy
    microsoftGraphMdmWindowsInformationProtectionPolicy := *openapiclient.NewMicrosoftGraphMdmWindowsInformationProtectionPolicy() // MicrosoftGraphMdmWindowsInformationProtectionPolicy | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceAppManagementMdmWindowsInformationProtectionPolicyApi.DeviceAppManagementUpdateMdmWindowsInformationProtectionPolicies(context.Background(), mdmWindowsInformationProtectionPolicyId).MicrosoftGraphMdmWindowsInformationProtectionPolicy(microsoftGraphMdmWindowsInformationProtectionPolicy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAppManagementMdmWindowsInformationProtectionPolicyApi.DeviceAppManagementUpdateMdmWindowsInformationProtectionPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mdmWindowsInformationProtectionPolicyId** | **string** | key: id of mdmWindowsInformationProtectionPolicy | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceAppManagementUpdateMdmWindowsInformationProtectionPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphMdmWindowsInformationProtectionPolicy** | [**MicrosoftGraphMdmWindowsInformationProtectionPolicy**](MicrosoftGraphMdmWindowsInformationProtectionPolicy.md) | New navigation property values | 

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

