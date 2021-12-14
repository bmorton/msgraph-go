# \OrganizationCertificateBasedAuthConfigurationApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrganizationCreateRefCertificateBasedAuthConfiguration**](OrganizationCertificateBasedAuthConfigurationApi.md#OrganizationCreateRefCertificateBasedAuthConfiguration) | **Post** /organization/{organization-id}/certificateBasedAuthConfiguration/$ref | Create new navigation property ref to certificateBasedAuthConfiguration for organization
[**OrganizationListCertificateBasedAuthConfiguration**](OrganizationCertificateBasedAuthConfigurationApi.md#OrganizationListCertificateBasedAuthConfiguration) | **Get** /organization/{organization-id}/certificateBasedAuthConfiguration | Get certificateBasedAuthConfiguration from organization
[**OrganizationListRefCertificateBasedAuthConfiguration**](OrganizationCertificateBasedAuthConfigurationApi.md#OrganizationListRefCertificateBasedAuthConfiguration) | **Get** /organization/{organization-id}/certificateBasedAuthConfiguration/$ref | Get ref of certificateBasedAuthConfiguration from organization



## OrganizationCreateRefCertificateBasedAuthConfiguration

> map[string]interface{} OrganizationCreateRefCertificateBasedAuthConfiguration(ctx, organizationId).RequestBody(requestBody).Execute()

Create new navigation property ref to certificateBasedAuthConfiguration for organization



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
    organizationId := "organizationId_example" // string | key: id of organization
    requestBody := map[string]map[string]interface{}{"key": map[string]interface{}(123)} // map[string]map[string]interface{} | New navigation property ref value

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganizationCertificateBasedAuthConfigurationApi.OrganizationCreateRefCertificateBasedAuthConfiguration(context.Background(), organizationId).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationCertificateBasedAuthConfigurationApi.OrganizationCreateRefCertificateBasedAuthConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganizationCreateRefCertificateBasedAuthConfiguration`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `OrganizationCertificateBasedAuthConfigurationApi.OrganizationCreateRefCertificateBasedAuthConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | key: id of organization | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationCreateRefCertificateBasedAuthConfigurationRequest struct via the builder pattern


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


## OrganizationListCertificateBasedAuthConfiguration

> CollectionOfCertificateBasedAuthConfiguration OrganizationListCertificateBasedAuthConfiguration(ctx, organizationId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get certificateBasedAuthConfiguration from organization



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
    organizationId := "organizationId_example" // string | key: id of organization
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
    resp, r, err := api_client.OrganizationCertificateBasedAuthConfigurationApi.OrganizationListCertificateBasedAuthConfiguration(context.Background(), organizationId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationCertificateBasedAuthConfigurationApi.OrganizationListCertificateBasedAuthConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganizationListCertificateBasedAuthConfiguration`: CollectionOfCertificateBasedAuthConfiguration
    fmt.Fprintf(os.Stdout, "Response from `OrganizationCertificateBasedAuthConfigurationApi.OrganizationListCertificateBasedAuthConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | key: id of organization | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationListCertificateBasedAuthConfigurationRequest struct via the builder pattern


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

[**CollectionOfCertificateBasedAuthConfiguration**](CollectionOfCertificateBasedAuthConfiguration.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationListRefCertificateBasedAuthConfiguration

> CollectionOfLinksOfCertificateBasedAuthConfiguration OrganizationListRefCertificateBasedAuthConfiguration(ctx, organizationId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Execute()

Get ref of certificateBasedAuthConfiguration from organization



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
    organizationId := "organizationId_example" // string | key: id of organization
    top := int32(50) // int32 | Show only the first n items (optional)
    skip := int32(56) // int32 | Skip the first n items (optional)
    search := "search_example" // string | Search items by search phrases (optional)
    filter := "filter_example" // string | Filter items by property values (optional)
    count := true // bool | Include count of items (optional)
    orderby := []string{"Orderby_example"} // []string | Order items by property values (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganizationCertificateBasedAuthConfigurationApi.OrganizationListRefCertificateBasedAuthConfiguration(context.Background(), organizationId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationCertificateBasedAuthConfigurationApi.OrganizationListRefCertificateBasedAuthConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganizationListRefCertificateBasedAuthConfiguration`: CollectionOfLinksOfCertificateBasedAuthConfiguration
    fmt.Fprintf(os.Stdout, "Response from `OrganizationCertificateBasedAuthConfigurationApi.OrganizationListRefCertificateBasedAuthConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | key: id of organization | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationListRefCertificateBasedAuthConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **top** | **int32** | Show only the first n items | 
 **skip** | **int32** | Skip the first n items | 
 **search** | **string** | Search items by search phrases | 
 **filter** | **string** | Filter items by property values | 
 **count** | **bool** | Include count of items | 
 **orderby** | **[]string** | Order items by property values | 

### Return type

[**CollectionOfLinksOfCertificateBasedAuthConfiguration**](CollectionOfLinksOfCertificateBasedAuthConfiguration.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

