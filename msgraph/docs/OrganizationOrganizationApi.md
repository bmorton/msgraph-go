# \OrganizationOrganizationApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrganizationOrganizationCreateOrganization**](OrganizationOrganizationApi.md#OrganizationOrganizationCreateOrganization) | **Post** /organization | Add new entity to organization
[**OrganizationOrganizationDeleteOrganization**](OrganizationOrganizationApi.md#OrganizationOrganizationDeleteOrganization) | **Delete** /organization/{organization-id} | Delete entity from organization
[**OrganizationOrganizationGetOrganization**](OrganizationOrganizationApi.md#OrganizationOrganizationGetOrganization) | **Get** /organization/{organization-id} | Get entity from organization by key
[**OrganizationOrganizationListOrganization**](OrganizationOrganizationApi.md#OrganizationOrganizationListOrganization) | **Get** /organization | Get entities from organization
[**OrganizationOrganizationUpdateOrganization**](OrganizationOrganizationApi.md#OrganizationOrganizationUpdateOrganization) | **Patch** /organization/{organization-id} | Update entity in organization



## OrganizationOrganizationCreateOrganization

> MicrosoftGraphOrganization OrganizationOrganizationCreateOrganization(ctx).MicrosoftGraphOrganization(microsoftGraphOrganization).Execute()

Add new entity to organization

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
    microsoftGraphOrganization := *openapiclient.NewMicrosoftGraphOrganization() // MicrosoftGraphOrganization | New entity

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganizationOrganizationApi.OrganizationOrganizationCreateOrganization(context.Background()).MicrosoftGraphOrganization(microsoftGraphOrganization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationOrganizationApi.OrganizationOrganizationCreateOrganization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganizationOrganizationCreateOrganization`: MicrosoftGraphOrganization
    fmt.Fprintf(os.Stdout, "Response from `OrganizationOrganizationApi.OrganizationOrganizationCreateOrganization`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationOrganizationCreateOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphOrganization** | [**MicrosoftGraphOrganization**](MicrosoftGraphOrganization.md) | New entity | 

### Return type

[**MicrosoftGraphOrganization**](MicrosoftGraphOrganization.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationOrganizationDeleteOrganization

> OrganizationOrganizationDeleteOrganization(ctx, organizationId).IfMatch(ifMatch).Execute()

Delete entity from organization

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
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganizationOrganizationApi.OrganizationOrganizationDeleteOrganization(context.Background(), organizationId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationOrganizationApi.OrganizationOrganizationDeleteOrganization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | key: id of organization | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationOrganizationDeleteOrganizationRequest struct via the builder pattern


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


## OrganizationOrganizationGetOrganization

> MicrosoftGraphOrganization OrganizationOrganizationGetOrganization(ctx, organizationId).Select_(select_).Expand(expand).Execute()

Get entity from organization by key

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
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganizationOrganizationApi.OrganizationOrganizationGetOrganization(context.Background(), organizationId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationOrganizationApi.OrganizationOrganizationGetOrganization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganizationOrganizationGetOrganization`: MicrosoftGraphOrganization
    fmt.Fprintf(os.Stdout, "Response from `OrganizationOrganizationApi.OrganizationOrganizationGetOrganization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | key: id of organization | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationOrganizationGetOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphOrganization**](MicrosoftGraphOrganization.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationOrganizationListOrganization

> CollectionOfOrganization OrganizationOrganizationListOrganization(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get entities from organization

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
    resp, r, err := api_client.OrganizationOrganizationApi.OrganizationOrganizationListOrganization(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationOrganizationApi.OrganizationOrganizationListOrganization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganizationOrganizationListOrganization`: CollectionOfOrganization
    fmt.Fprintf(os.Stdout, "Response from `OrganizationOrganizationApi.OrganizationOrganizationListOrganization`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationOrganizationListOrganizationRequest struct via the builder pattern


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

[**CollectionOfOrganization**](CollectionOfOrganization.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationOrganizationUpdateOrganization

> OrganizationOrganizationUpdateOrganization(ctx, organizationId).MicrosoftGraphOrganization(microsoftGraphOrganization).Execute()

Update entity in organization

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
    microsoftGraphOrganization := *openapiclient.NewMicrosoftGraphOrganization() // MicrosoftGraphOrganization | New property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganizationOrganizationApi.OrganizationOrganizationUpdateOrganization(context.Background(), organizationId).MicrosoftGraphOrganization(microsoftGraphOrganization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationOrganizationApi.OrganizationOrganizationUpdateOrganization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | key: id of organization | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationOrganizationUpdateOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphOrganization** | [**MicrosoftGraphOrganization**](MicrosoftGraphOrganization.md) | New property values | 

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

