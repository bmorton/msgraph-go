# \AuthenticationMethodsPolicyAuthenticationMethodConfigurationApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthenticationMethodsPolicyCreateAuthenticationMethodConfigurations**](AuthenticationMethodsPolicyAuthenticationMethodConfigurationApi.md#AuthenticationMethodsPolicyCreateAuthenticationMethodConfigurations) | **Post** /authenticationMethodsPolicy/authenticationMethodConfigurations | Create new navigation property to authenticationMethodConfigurations for authenticationMethodsPolicy
[**AuthenticationMethodsPolicyDeleteAuthenticationMethodConfigurations**](AuthenticationMethodsPolicyAuthenticationMethodConfigurationApi.md#AuthenticationMethodsPolicyDeleteAuthenticationMethodConfigurations) | **Delete** /authenticationMethodsPolicy/authenticationMethodConfigurations/{authenticationMethodConfiguration-id} | Delete navigation property authenticationMethodConfigurations for authenticationMethodsPolicy
[**AuthenticationMethodsPolicyGetAuthenticationMethodConfigurations**](AuthenticationMethodsPolicyAuthenticationMethodConfigurationApi.md#AuthenticationMethodsPolicyGetAuthenticationMethodConfigurations) | **Get** /authenticationMethodsPolicy/authenticationMethodConfigurations/{authenticationMethodConfiguration-id} | Get authenticationMethodConfigurations from authenticationMethodsPolicy
[**AuthenticationMethodsPolicyListAuthenticationMethodConfigurations**](AuthenticationMethodsPolicyAuthenticationMethodConfigurationApi.md#AuthenticationMethodsPolicyListAuthenticationMethodConfigurations) | **Get** /authenticationMethodsPolicy/authenticationMethodConfigurations | Get authenticationMethodConfigurations from authenticationMethodsPolicy
[**AuthenticationMethodsPolicyUpdateAuthenticationMethodConfigurations**](AuthenticationMethodsPolicyAuthenticationMethodConfigurationApi.md#AuthenticationMethodsPolicyUpdateAuthenticationMethodConfigurations) | **Patch** /authenticationMethodsPolicy/authenticationMethodConfigurations/{authenticationMethodConfiguration-id} | Update the navigation property authenticationMethodConfigurations in authenticationMethodsPolicy



## AuthenticationMethodsPolicyCreateAuthenticationMethodConfigurations

> MicrosoftGraphAuthenticationMethodConfiguration AuthenticationMethodsPolicyCreateAuthenticationMethodConfigurations(ctx).MicrosoftGraphAuthenticationMethodConfiguration(microsoftGraphAuthenticationMethodConfiguration).Execute()

Create new navigation property to authenticationMethodConfigurations for authenticationMethodsPolicy



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
    microsoftGraphAuthenticationMethodConfiguration := *openapiclient.NewMicrosoftGraphAuthenticationMethodConfiguration() // MicrosoftGraphAuthenticationMethodConfiguration | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthenticationMethodsPolicyAuthenticationMethodConfigurationApi.AuthenticationMethodsPolicyCreateAuthenticationMethodConfigurations(context.Background()).MicrosoftGraphAuthenticationMethodConfiguration(microsoftGraphAuthenticationMethodConfiguration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationMethodsPolicyAuthenticationMethodConfigurationApi.AuthenticationMethodsPolicyCreateAuthenticationMethodConfigurations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthenticationMethodsPolicyCreateAuthenticationMethodConfigurations`: MicrosoftGraphAuthenticationMethodConfiguration
    fmt.Fprintf(os.Stdout, "Response from `AuthenticationMethodsPolicyAuthenticationMethodConfigurationApi.AuthenticationMethodsPolicyCreateAuthenticationMethodConfigurations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthenticationMethodsPolicyCreateAuthenticationMethodConfigurationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphAuthenticationMethodConfiguration** | [**MicrosoftGraphAuthenticationMethodConfiguration**](MicrosoftGraphAuthenticationMethodConfiguration.md) | New navigation property | 

### Return type

[**MicrosoftGraphAuthenticationMethodConfiguration**](MicrosoftGraphAuthenticationMethodConfiguration.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthenticationMethodsPolicyDeleteAuthenticationMethodConfigurations

> AuthenticationMethodsPolicyDeleteAuthenticationMethodConfigurations(ctx, authenticationMethodConfigurationId).IfMatch(ifMatch).Execute()

Delete navigation property authenticationMethodConfigurations for authenticationMethodsPolicy



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
    authenticationMethodConfigurationId := "authenticationMethodConfigurationId_example" // string | key: id of authenticationMethodConfiguration
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthenticationMethodsPolicyAuthenticationMethodConfigurationApi.AuthenticationMethodsPolicyDeleteAuthenticationMethodConfigurations(context.Background(), authenticationMethodConfigurationId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationMethodsPolicyAuthenticationMethodConfigurationApi.AuthenticationMethodsPolicyDeleteAuthenticationMethodConfigurations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authenticationMethodConfigurationId** | **string** | key: id of authenticationMethodConfiguration | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuthenticationMethodsPolicyDeleteAuthenticationMethodConfigurationsRequest struct via the builder pattern


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


## AuthenticationMethodsPolicyGetAuthenticationMethodConfigurations

> MicrosoftGraphAuthenticationMethodConfiguration AuthenticationMethodsPolicyGetAuthenticationMethodConfigurations(ctx, authenticationMethodConfigurationId).Select_(select_).Expand(expand).Execute()

Get authenticationMethodConfigurations from authenticationMethodsPolicy



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
    authenticationMethodConfigurationId := "authenticationMethodConfigurationId_example" // string | key: id of authenticationMethodConfiguration
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthenticationMethodsPolicyAuthenticationMethodConfigurationApi.AuthenticationMethodsPolicyGetAuthenticationMethodConfigurations(context.Background(), authenticationMethodConfigurationId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationMethodsPolicyAuthenticationMethodConfigurationApi.AuthenticationMethodsPolicyGetAuthenticationMethodConfigurations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthenticationMethodsPolicyGetAuthenticationMethodConfigurations`: MicrosoftGraphAuthenticationMethodConfiguration
    fmt.Fprintf(os.Stdout, "Response from `AuthenticationMethodsPolicyAuthenticationMethodConfigurationApi.AuthenticationMethodsPolicyGetAuthenticationMethodConfigurations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authenticationMethodConfigurationId** | **string** | key: id of authenticationMethodConfiguration | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuthenticationMethodsPolicyGetAuthenticationMethodConfigurationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphAuthenticationMethodConfiguration**](MicrosoftGraphAuthenticationMethodConfiguration.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthenticationMethodsPolicyListAuthenticationMethodConfigurations

> CollectionOfAuthenticationMethodConfiguration AuthenticationMethodsPolicyListAuthenticationMethodConfigurations(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get authenticationMethodConfigurations from authenticationMethodsPolicy



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
    resp, r, err := api_client.AuthenticationMethodsPolicyAuthenticationMethodConfigurationApi.AuthenticationMethodsPolicyListAuthenticationMethodConfigurations(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationMethodsPolicyAuthenticationMethodConfigurationApi.AuthenticationMethodsPolicyListAuthenticationMethodConfigurations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthenticationMethodsPolicyListAuthenticationMethodConfigurations`: CollectionOfAuthenticationMethodConfiguration
    fmt.Fprintf(os.Stdout, "Response from `AuthenticationMethodsPolicyAuthenticationMethodConfigurationApi.AuthenticationMethodsPolicyListAuthenticationMethodConfigurations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthenticationMethodsPolicyListAuthenticationMethodConfigurationsRequest struct via the builder pattern


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

[**CollectionOfAuthenticationMethodConfiguration**](CollectionOfAuthenticationMethodConfiguration.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthenticationMethodsPolicyUpdateAuthenticationMethodConfigurations

> AuthenticationMethodsPolicyUpdateAuthenticationMethodConfigurations(ctx, authenticationMethodConfigurationId).MicrosoftGraphAuthenticationMethodConfiguration(microsoftGraphAuthenticationMethodConfiguration).Execute()

Update the navigation property authenticationMethodConfigurations in authenticationMethodsPolicy



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
    authenticationMethodConfigurationId := "authenticationMethodConfigurationId_example" // string | key: id of authenticationMethodConfiguration
    microsoftGraphAuthenticationMethodConfiguration := *openapiclient.NewMicrosoftGraphAuthenticationMethodConfiguration() // MicrosoftGraphAuthenticationMethodConfiguration | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthenticationMethodsPolicyAuthenticationMethodConfigurationApi.AuthenticationMethodsPolicyUpdateAuthenticationMethodConfigurations(context.Background(), authenticationMethodConfigurationId).MicrosoftGraphAuthenticationMethodConfiguration(microsoftGraphAuthenticationMethodConfiguration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationMethodsPolicyAuthenticationMethodConfigurationApi.AuthenticationMethodsPolicyUpdateAuthenticationMethodConfigurations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authenticationMethodConfigurationId** | **string** | key: id of authenticationMethodConfiguration | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuthenticationMethodsPolicyUpdateAuthenticationMethodConfigurationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphAuthenticationMethodConfiguration** | [**MicrosoftGraphAuthenticationMethodConfiguration**](MicrosoftGraphAuthenticationMethodConfiguration.md) | New navigation property values | 

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

