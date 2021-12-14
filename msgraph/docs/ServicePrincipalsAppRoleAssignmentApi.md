# \ServicePrincipalsAppRoleAssignmentApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ServicePrincipalsCreateAppRoleAssignedTo**](ServicePrincipalsAppRoleAssignmentApi.md#ServicePrincipalsCreateAppRoleAssignedTo) | **Post** /servicePrincipals/{servicePrincipal-id}/appRoleAssignedTo | Create new navigation property to appRoleAssignedTo for servicePrincipals
[**ServicePrincipalsCreateAppRoleAssignments**](ServicePrincipalsAppRoleAssignmentApi.md#ServicePrincipalsCreateAppRoleAssignments) | **Post** /servicePrincipals/{servicePrincipal-id}/appRoleAssignments | Create new navigation property to appRoleAssignments for servicePrincipals
[**ServicePrincipalsDeleteAppRoleAssignedTo**](ServicePrincipalsAppRoleAssignmentApi.md#ServicePrincipalsDeleteAppRoleAssignedTo) | **Delete** /servicePrincipals/{servicePrincipal-id}/appRoleAssignedTo/{appRoleAssignment-id} | Delete navigation property appRoleAssignedTo for servicePrincipals
[**ServicePrincipalsDeleteAppRoleAssignments**](ServicePrincipalsAppRoleAssignmentApi.md#ServicePrincipalsDeleteAppRoleAssignments) | **Delete** /servicePrincipals/{servicePrincipal-id}/appRoleAssignments/{appRoleAssignment-id} | Delete navigation property appRoleAssignments for servicePrincipals
[**ServicePrincipalsGetAppRoleAssignedTo**](ServicePrincipalsAppRoleAssignmentApi.md#ServicePrincipalsGetAppRoleAssignedTo) | **Get** /servicePrincipals/{servicePrincipal-id}/appRoleAssignedTo/{appRoleAssignment-id} | Get appRoleAssignedTo from servicePrincipals
[**ServicePrincipalsGetAppRoleAssignments**](ServicePrincipalsAppRoleAssignmentApi.md#ServicePrincipalsGetAppRoleAssignments) | **Get** /servicePrincipals/{servicePrincipal-id}/appRoleAssignments/{appRoleAssignment-id} | Get appRoleAssignments from servicePrincipals
[**ServicePrincipalsListAppRoleAssignedTo**](ServicePrincipalsAppRoleAssignmentApi.md#ServicePrincipalsListAppRoleAssignedTo) | **Get** /servicePrincipals/{servicePrincipal-id}/appRoleAssignedTo | Get appRoleAssignedTo from servicePrincipals
[**ServicePrincipalsListAppRoleAssignments**](ServicePrincipalsAppRoleAssignmentApi.md#ServicePrincipalsListAppRoleAssignments) | **Get** /servicePrincipals/{servicePrincipal-id}/appRoleAssignments | Get appRoleAssignments from servicePrincipals
[**ServicePrincipalsUpdateAppRoleAssignedTo**](ServicePrincipalsAppRoleAssignmentApi.md#ServicePrincipalsUpdateAppRoleAssignedTo) | **Patch** /servicePrincipals/{servicePrincipal-id}/appRoleAssignedTo/{appRoleAssignment-id} | Update the navigation property appRoleAssignedTo in servicePrincipals
[**ServicePrincipalsUpdateAppRoleAssignments**](ServicePrincipalsAppRoleAssignmentApi.md#ServicePrincipalsUpdateAppRoleAssignments) | **Patch** /servicePrincipals/{servicePrincipal-id}/appRoleAssignments/{appRoleAssignment-id} | Update the navigation property appRoleAssignments in servicePrincipals



## ServicePrincipalsCreateAppRoleAssignedTo

> MicrosoftGraphAppRoleAssignment ServicePrincipalsCreateAppRoleAssignedTo(ctx, servicePrincipalId).MicrosoftGraphAppRoleAssignment(microsoftGraphAppRoleAssignment).Execute()

Create new navigation property to appRoleAssignedTo for servicePrincipals



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
    servicePrincipalId := "servicePrincipalId_example" // string | key: id of servicePrincipal
    microsoftGraphAppRoleAssignment := *openapiclient.NewMicrosoftGraphAppRoleAssignment() // MicrosoftGraphAppRoleAssignment | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServicePrincipalsAppRoleAssignmentApi.ServicePrincipalsCreateAppRoleAssignedTo(context.Background(), servicePrincipalId).MicrosoftGraphAppRoleAssignment(microsoftGraphAppRoleAssignment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicePrincipalsAppRoleAssignmentApi.ServicePrincipalsCreateAppRoleAssignedTo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServicePrincipalsCreateAppRoleAssignedTo`: MicrosoftGraphAppRoleAssignment
    fmt.Fprintf(os.Stdout, "Response from `ServicePrincipalsAppRoleAssignmentApi.ServicePrincipalsCreateAppRoleAssignedTo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**servicePrincipalId** | **string** | key: id of servicePrincipal | 

### Other Parameters

Other parameters are passed through a pointer to a apiServicePrincipalsCreateAppRoleAssignedToRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphAppRoleAssignment** | [**MicrosoftGraphAppRoleAssignment**](MicrosoftGraphAppRoleAssignment.md) | New navigation property | 

### Return type

[**MicrosoftGraphAppRoleAssignment**](MicrosoftGraphAppRoleAssignment.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServicePrincipalsCreateAppRoleAssignments

> MicrosoftGraphAppRoleAssignment ServicePrincipalsCreateAppRoleAssignments(ctx, servicePrincipalId).MicrosoftGraphAppRoleAssignment(microsoftGraphAppRoleAssignment).Execute()

Create new navigation property to appRoleAssignments for servicePrincipals



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
    servicePrincipalId := "servicePrincipalId_example" // string | key: id of servicePrincipal
    microsoftGraphAppRoleAssignment := *openapiclient.NewMicrosoftGraphAppRoleAssignment() // MicrosoftGraphAppRoleAssignment | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServicePrincipalsAppRoleAssignmentApi.ServicePrincipalsCreateAppRoleAssignments(context.Background(), servicePrincipalId).MicrosoftGraphAppRoleAssignment(microsoftGraphAppRoleAssignment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicePrincipalsAppRoleAssignmentApi.ServicePrincipalsCreateAppRoleAssignments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServicePrincipalsCreateAppRoleAssignments`: MicrosoftGraphAppRoleAssignment
    fmt.Fprintf(os.Stdout, "Response from `ServicePrincipalsAppRoleAssignmentApi.ServicePrincipalsCreateAppRoleAssignments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**servicePrincipalId** | **string** | key: id of servicePrincipal | 

### Other Parameters

Other parameters are passed through a pointer to a apiServicePrincipalsCreateAppRoleAssignmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphAppRoleAssignment** | [**MicrosoftGraphAppRoleAssignment**](MicrosoftGraphAppRoleAssignment.md) | New navigation property | 

### Return type

[**MicrosoftGraphAppRoleAssignment**](MicrosoftGraphAppRoleAssignment.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServicePrincipalsDeleteAppRoleAssignedTo

> ServicePrincipalsDeleteAppRoleAssignedTo(ctx, servicePrincipalId, appRoleAssignmentId).IfMatch(ifMatch).Execute()

Delete navigation property appRoleAssignedTo for servicePrincipals



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
    servicePrincipalId := "servicePrincipalId_example" // string | key: id of servicePrincipal
    appRoleAssignmentId := "appRoleAssignmentId_example" // string | key: id of appRoleAssignment
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServicePrincipalsAppRoleAssignmentApi.ServicePrincipalsDeleteAppRoleAssignedTo(context.Background(), servicePrincipalId, appRoleAssignmentId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicePrincipalsAppRoleAssignmentApi.ServicePrincipalsDeleteAppRoleAssignedTo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**servicePrincipalId** | **string** | key: id of servicePrincipal | 
**appRoleAssignmentId** | **string** | key: id of appRoleAssignment | 

### Other Parameters

Other parameters are passed through a pointer to a apiServicePrincipalsDeleteAppRoleAssignedToRequest struct via the builder pattern


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


## ServicePrincipalsDeleteAppRoleAssignments

> ServicePrincipalsDeleteAppRoleAssignments(ctx, servicePrincipalId, appRoleAssignmentId).IfMatch(ifMatch).Execute()

Delete navigation property appRoleAssignments for servicePrincipals



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
    servicePrincipalId := "servicePrincipalId_example" // string | key: id of servicePrincipal
    appRoleAssignmentId := "appRoleAssignmentId_example" // string | key: id of appRoleAssignment
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServicePrincipalsAppRoleAssignmentApi.ServicePrincipalsDeleteAppRoleAssignments(context.Background(), servicePrincipalId, appRoleAssignmentId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicePrincipalsAppRoleAssignmentApi.ServicePrincipalsDeleteAppRoleAssignments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**servicePrincipalId** | **string** | key: id of servicePrincipal | 
**appRoleAssignmentId** | **string** | key: id of appRoleAssignment | 

### Other Parameters

Other parameters are passed through a pointer to a apiServicePrincipalsDeleteAppRoleAssignmentsRequest struct via the builder pattern


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


## ServicePrincipalsGetAppRoleAssignedTo

> MicrosoftGraphAppRoleAssignment ServicePrincipalsGetAppRoleAssignedTo(ctx, servicePrincipalId, appRoleAssignmentId).Select_(select_).Expand(expand).Execute()

Get appRoleAssignedTo from servicePrincipals



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
    servicePrincipalId := "servicePrincipalId_example" // string | key: id of servicePrincipal
    appRoleAssignmentId := "appRoleAssignmentId_example" // string | key: id of appRoleAssignment
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServicePrincipalsAppRoleAssignmentApi.ServicePrincipalsGetAppRoleAssignedTo(context.Background(), servicePrincipalId, appRoleAssignmentId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicePrincipalsAppRoleAssignmentApi.ServicePrincipalsGetAppRoleAssignedTo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServicePrincipalsGetAppRoleAssignedTo`: MicrosoftGraphAppRoleAssignment
    fmt.Fprintf(os.Stdout, "Response from `ServicePrincipalsAppRoleAssignmentApi.ServicePrincipalsGetAppRoleAssignedTo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**servicePrincipalId** | **string** | key: id of servicePrincipal | 
**appRoleAssignmentId** | **string** | key: id of appRoleAssignment | 

### Other Parameters

Other parameters are passed through a pointer to a apiServicePrincipalsGetAppRoleAssignedToRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphAppRoleAssignment**](MicrosoftGraphAppRoleAssignment.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServicePrincipalsGetAppRoleAssignments

> MicrosoftGraphAppRoleAssignment ServicePrincipalsGetAppRoleAssignments(ctx, servicePrincipalId, appRoleAssignmentId).Select_(select_).Expand(expand).Execute()

Get appRoleAssignments from servicePrincipals



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
    servicePrincipalId := "servicePrincipalId_example" // string | key: id of servicePrincipal
    appRoleAssignmentId := "appRoleAssignmentId_example" // string | key: id of appRoleAssignment
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServicePrincipalsAppRoleAssignmentApi.ServicePrincipalsGetAppRoleAssignments(context.Background(), servicePrincipalId, appRoleAssignmentId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicePrincipalsAppRoleAssignmentApi.ServicePrincipalsGetAppRoleAssignments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServicePrincipalsGetAppRoleAssignments`: MicrosoftGraphAppRoleAssignment
    fmt.Fprintf(os.Stdout, "Response from `ServicePrincipalsAppRoleAssignmentApi.ServicePrincipalsGetAppRoleAssignments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**servicePrincipalId** | **string** | key: id of servicePrincipal | 
**appRoleAssignmentId** | **string** | key: id of appRoleAssignment | 

### Other Parameters

Other parameters are passed through a pointer to a apiServicePrincipalsGetAppRoleAssignmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphAppRoleAssignment**](MicrosoftGraphAppRoleAssignment.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServicePrincipalsListAppRoleAssignedTo

> CollectionOfAppRoleAssignment ServicePrincipalsListAppRoleAssignedTo(ctx, servicePrincipalId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get appRoleAssignedTo from servicePrincipals



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
    servicePrincipalId := "servicePrincipalId_example" // string | key: id of servicePrincipal
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
    resp, r, err := api_client.ServicePrincipalsAppRoleAssignmentApi.ServicePrincipalsListAppRoleAssignedTo(context.Background(), servicePrincipalId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicePrincipalsAppRoleAssignmentApi.ServicePrincipalsListAppRoleAssignedTo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServicePrincipalsListAppRoleAssignedTo`: CollectionOfAppRoleAssignment
    fmt.Fprintf(os.Stdout, "Response from `ServicePrincipalsAppRoleAssignmentApi.ServicePrincipalsListAppRoleAssignedTo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**servicePrincipalId** | **string** | key: id of servicePrincipal | 

### Other Parameters

Other parameters are passed through a pointer to a apiServicePrincipalsListAppRoleAssignedToRequest struct via the builder pattern


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

[**CollectionOfAppRoleAssignment**](CollectionOfAppRoleAssignment.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServicePrincipalsListAppRoleAssignments

> CollectionOfAppRoleAssignment ServicePrincipalsListAppRoleAssignments(ctx, servicePrincipalId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get appRoleAssignments from servicePrincipals



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
    servicePrincipalId := "servicePrincipalId_example" // string | key: id of servicePrincipal
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
    resp, r, err := api_client.ServicePrincipalsAppRoleAssignmentApi.ServicePrincipalsListAppRoleAssignments(context.Background(), servicePrincipalId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicePrincipalsAppRoleAssignmentApi.ServicePrincipalsListAppRoleAssignments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServicePrincipalsListAppRoleAssignments`: CollectionOfAppRoleAssignment
    fmt.Fprintf(os.Stdout, "Response from `ServicePrincipalsAppRoleAssignmentApi.ServicePrincipalsListAppRoleAssignments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**servicePrincipalId** | **string** | key: id of servicePrincipal | 

### Other Parameters

Other parameters are passed through a pointer to a apiServicePrincipalsListAppRoleAssignmentsRequest struct via the builder pattern


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

[**CollectionOfAppRoleAssignment**](CollectionOfAppRoleAssignment.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServicePrincipalsUpdateAppRoleAssignedTo

> ServicePrincipalsUpdateAppRoleAssignedTo(ctx, servicePrincipalId, appRoleAssignmentId).MicrosoftGraphAppRoleAssignment(microsoftGraphAppRoleAssignment).Execute()

Update the navigation property appRoleAssignedTo in servicePrincipals



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
    servicePrincipalId := "servicePrincipalId_example" // string | key: id of servicePrincipal
    appRoleAssignmentId := "appRoleAssignmentId_example" // string | key: id of appRoleAssignment
    microsoftGraphAppRoleAssignment := *openapiclient.NewMicrosoftGraphAppRoleAssignment() // MicrosoftGraphAppRoleAssignment | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServicePrincipalsAppRoleAssignmentApi.ServicePrincipalsUpdateAppRoleAssignedTo(context.Background(), servicePrincipalId, appRoleAssignmentId).MicrosoftGraphAppRoleAssignment(microsoftGraphAppRoleAssignment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicePrincipalsAppRoleAssignmentApi.ServicePrincipalsUpdateAppRoleAssignedTo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**servicePrincipalId** | **string** | key: id of servicePrincipal | 
**appRoleAssignmentId** | **string** | key: id of appRoleAssignment | 

### Other Parameters

Other parameters are passed through a pointer to a apiServicePrincipalsUpdateAppRoleAssignedToRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **microsoftGraphAppRoleAssignment** | [**MicrosoftGraphAppRoleAssignment**](MicrosoftGraphAppRoleAssignment.md) | New navigation property values | 

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


## ServicePrincipalsUpdateAppRoleAssignments

> ServicePrincipalsUpdateAppRoleAssignments(ctx, servicePrincipalId, appRoleAssignmentId).MicrosoftGraphAppRoleAssignment(microsoftGraphAppRoleAssignment).Execute()

Update the navigation property appRoleAssignments in servicePrincipals



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
    servicePrincipalId := "servicePrincipalId_example" // string | key: id of servicePrincipal
    appRoleAssignmentId := "appRoleAssignmentId_example" // string | key: id of appRoleAssignment
    microsoftGraphAppRoleAssignment := *openapiclient.NewMicrosoftGraphAppRoleAssignment() // MicrosoftGraphAppRoleAssignment | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServicePrincipalsAppRoleAssignmentApi.ServicePrincipalsUpdateAppRoleAssignments(context.Background(), servicePrincipalId, appRoleAssignmentId).MicrosoftGraphAppRoleAssignment(microsoftGraphAppRoleAssignment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicePrincipalsAppRoleAssignmentApi.ServicePrincipalsUpdateAppRoleAssignments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**servicePrincipalId** | **string** | key: id of servicePrincipal | 
**appRoleAssignmentId** | **string** | key: id of appRoleAssignment | 

### Other Parameters

Other parameters are passed through a pointer to a apiServicePrincipalsUpdateAppRoleAssignmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **microsoftGraphAppRoleAssignment** | [**MicrosoftGraphAppRoleAssignment**](MicrosoftGraphAppRoleAssignment.md) | New navigation property values | 

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

