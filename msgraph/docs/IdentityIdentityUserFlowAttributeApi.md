# \IdentityIdentityUserFlowAttributeApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IdentityCreateUserFlowAttributes**](IdentityIdentityUserFlowAttributeApi.md#IdentityCreateUserFlowAttributes) | **Post** /identity/userFlowAttributes | Create new navigation property to userFlowAttributes for identity
[**IdentityDeleteUserFlowAttributes**](IdentityIdentityUserFlowAttributeApi.md#IdentityDeleteUserFlowAttributes) | **Delete** /identity/userFlowAttributes/{identityUserFlowAttribute-id} | Delete navigation property userFlowAttributes for identity
[**IdentityGetUserFlowAttributes**](IdentityIdentityUserFlowAttributeApi.md#IdentityGetUserFlowAttributes) | **Get** /identity/userFlowAttributes/{identityUserFlowAttribute-id} | Get userFlowAttributes from identity
[**IdentityListUserFlowAttributes**](IdentityIdentityUserFlowAttributeApi.md#IdentityListUserFlowAttributes) | **Get** /identity/userFlowAttributes | Get userFlowAttributes from identity
[**IdentityUpdateUserFlowAttributes**](IdentityIdentityUserFlowAttributeApi.md#IdentityUpdateUserFlowAttributes) | **Patch** /identity/userFlowAttributes/{identityUserFlowAttribute-id} | Update the navigation property userFlowAttributes in identity



## IdentityCreateUserFlowAttributes

> MicrosoftGraphIdentityUserFlowAttribute IdentityCreateUserFlowAttributes(ctx).MicrosoftGraphIdentityUserFlowAttribute(microsoftGraphIdentityUserFlowAttribute).Execute()

Create new navigation property to userFlowAttributes for identity



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
    microsoftGraphIdentityUserFlowAttribute := *openapiclient.NewMicrosoftGraphIdentityUserFlowAttribute() // MicrosoftGraphIdentityUserFlowAttribute | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IdentityIdentityUserFlowAttributeApi.IdentityCreateUserFlowAttributes(context.Background()).MicrosoftGraphIdentityUserFlowAttribute(microsoftGraphIdentityUserFlowAttribute).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityIdentityUserFlowAttributeApi.IdentityCreateUserFlowAttributes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IdentityCreateUserFlowAttributes`: MicrosoftGraphIdentityUserFlowAttribute
    fmt.Fprintf(os.Stdout, "Response from `IdentityIdentityUserFlowAttributeApi.IdentityCreateUserFlowAttributes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIdentityCreateUserFlowAttributesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphIdentityUserFlowAttribute** | [**MicrosoftGraphIdentityUserFlowAttribute**](MicrosoftGraphIdentityUserFlowAttribute.md) | New navigation property | 

### Return type

[**MicrosoftGraphIdentityUserFlowAttribute**](MicrosoftGraphIdentityUserFlowAttribute.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IdentityDeleteUserFlowAttributes

> IdentityDeleteUserFlowAttributes(ctx, identityUserFlowAttributeId).IfMatch(ifMatch).Execute()

Delete navigation property userFlowAttributes for identity



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
    identityUserFlowAttributeId := "identityUserFlowAttributeId_example" // string | key: id of identityUserFlowAttribute
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IdentityIdentityUserFlowAttributeApi.IdentityDeleteUserFlowAttributes(context.Background(), identityUserFlowAttributeId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityIdentityUserFlowAttributeApi.IdentityDeleteUserFlowAttributes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identityUserFlowAttributeId** | **string** | key: id of identityUserFlowAttribute | 

### Other Parameters

Other parameters are passed through a pointer to a apiIdentityDeleteUserFlowAttributesRequest struct via the builder pattern


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


## IdentityGetUserFlowAttributes

> MicrosoftGraphIdentityUserFlowAttribute IdentityGetUserFlowAttributes(ctx, identityUserFlowAttributeId).Select_(select_).Expand(expand).Execute()

Get userFlowAttributes from identity



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
    identityUserFlowAttributeId := "identityUserFlowAttributeId_example" // string | key: id of identityUserFlowAttribute
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IdentityIdentityUserFlowAttributeApi.IdentityGetUserFlowAttributes(context.Background(), identityUserFlowAttributeId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityIdentityUserFlowAttributeApi.IdentityGetUserFlowAttributes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IdentityGetUserFlowAttributes`: MicrosoftGraphIdentityUserFlowAttribute
    fmt.Fprintf(os.Stdout, "Response from `IdentityIdentityUserFlowAttributeApi.IdentityGetUserFlowAttributes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identityUserFlowAttributeId** | **string** | key: id of identityUserFlowAttribute | 

### Other Parameters

Other parameters are passed through a pointer to a apiIdentityGetUserFlowAttributesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphIdentityUserFlowAttribute**](MicrosoftGraphIdentityUserFlowAttribute.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IdentityListUserFlowAttributes

> CollectionOfIdentityUserFlowAttribute IdentityListUserFlowAttributes(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get userFlowAttributes from identity



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
    resp, r, err := api_client.IdentityIdentityUserFlowAttributeApi.IdentityListUserFlowAttributes(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityIdentityUserFlowAttributeApi.IdentityListUserFlowAttributes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IdentityListUserFlowAttributes`: CollectionOfIdentityUserFlowAttribute
    fmt.Fprintf(os.Stdout, "Response from `IdentityIdentityUserFlowAttributeApi.IdentityListUserFlowAttributes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIdentityListUserFlowAttributesRequest struct via the builder pattern


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

[**CollectionOfIdentityUserFlowAttribute**](CollectionOfIdentityUserFlowAttribute.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IdentityUpdateUserFlowAttributes

> IdentityUpdateUserFlowAttributes(ctx, identityUserFlowAttributeId).MicrosoftGraphIdentityUserFlowAttribute(microsoftGraphIdentityUserFlowAttribute).Execute()

Update the navigation property userFlowAttributes in identity



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
    identityUserFlowAttributeId := "identityUserFlowAttributeId_example" // string | key: id of identityUserFlowAttribute
    microsoftGraphIdentityUserFlowAttribute := *openapiclient.NewMicrosoftGraphIdentityUserFlowAttribute() // MicrosoftGraphIdentityUserFlowAttribute | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IdentityIdentityUserFlowAttributeApi.IdentityUpdateUserFlowAttributes(context.Background(), identityUserFlowAttributeId).MicrosoftGraphIdentityUserFlowAttribute(microsoftGraphIdentityUserFlowAttribute).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityIdentityUserFlowAttributeApi.IdentityUpdateUserFlowAttributes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identityUserFlowAttributeId** | **string** | key: id of identityUserFlowAttribute | 

### Other Parameters

Other parameters are passed through a pointer to a apiIdentityUpdateUserFlowAttributesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphIdentityUserFlowAttribute** | [**MicrosoftGraphIdentityUserFlowAttribute**](MicrosoftGraphIdentityUserFlowAttribute.md) | New navigation property values | 

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

