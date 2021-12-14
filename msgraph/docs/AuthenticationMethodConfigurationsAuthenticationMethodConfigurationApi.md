# \AuthenticationMethodConfigurationsAuthenticationMethodConfigurationApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthenticationMethodConfigurationsAuthenticationMethodConfigurationCreateAuthenticationMethodConfiguration**](AuthenticationMethodConfigurationsAuthenticationMethodConfigurationApi.md#AuthenticationMethodConfigurationsAuthenticationMethodConfigurationCreateAuthenticationMethodConfiguration) | **Post** /authenticationMethodConfigurations | Add new entity to authenticationMethodConfigurations
[**AuthenticationMethodConfigurationsAuthenticationMethodConfigurationDeleteAuthenticationMethodConfiguration**](AuthenticationMethodConfigurationsAuthenticationMethodConfigurationApi.md#AuthenticationMethodConfigurationsAuthenticationMethodConfigurationDeleteAuthenticationMethodConfiguration) | **Delete** /authenticationMethodConfigurations/{authenticationMethodConfiguration-id} | Delete entity from authenticationMethodConfigurations
[**AuthenticationMethodConfigurationsAuthenticationMethodConfigurationGetAuthenticationMethodConfiguration**](AuthenticationMethodConfigurationsAuthenticationMethodConfigurationApi.md#AuthenticationMethodConfigurationsAuthenticationMethodConfigurationGetAuthenticationMethodConfiguration) | **Get** /authenticationMethodConfigurations/{authenticationMethodConfiguration-id} | Get entity from authenticationMethodConfigurations by key
[**AuthenticationMethodConfigurationsAuthenticationMethodConfigurationListAuthenticationMethodConfiguration**](AuthenticationMethodConfigurationsAuthenticationMethodConfigurationApi.md#AuthenticationMethodConfigurationsAuthenticationMethodConfigurationListAuthenticationMethodConfiguration) | **Get** /authenticationMethodConfigurations | Get entities from authenticationMethodConfigurations
[**AuthenticationMethodConfigurationsAuthenticationMethodConfigurationUpdateAuthenticationMethodConfiguration**](AuthenticationMethodConfigurationsAuthenticationMethodConfigurationApi.md#AuthenticationMethodConfigurationsAuthenticationMethodConfigurationUpdateAuthenticationMethodConfiguration) | **Patch** /authenticationMethodConfigurations/{authenticationMethodConfiguration-id} | Update entity in authenticationMethodConfigurations



## AuthenticationMethodConfigurationsAuthenticationMethodConfigurationCreateAuthenticationMethodConfiguration

> MicrosoftGraphAuthenticationMethodConfiguration AuthenticationMethodConfigurationsAuthenticationMethodConfigurationCreateAuthenticationMethodConfiguration(ctx).MicrosoftGraphAuthenticationMethodConfiguration(microsoftGraphAuthenticationMethodConfiguration).Execute()

Add new entity to authenticationMethodConfigurations

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
    microsoftGraphAuthenticationMethodConfiguration := *openapiclient.NewMicrosoftGraphAuthenticationMethodConfiguration() // MicrosoftGraphAuthenticationMethodConfiguration | New entity

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthenticationMethodConfigurationsAuthenticationMethodConfigurationApi.AuthenticationMethodConfigurationsAuthenticationMethodConfigurationCreateAuthenticationMethodConfiguration(context.Background()).MicrosoftGraphAuthenticationMethodConfiguration(microsoftGraphAuthenticationMethodConfiguration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationMethodConfigurationsAuthenticationMethodConfigurationApi.AuthenticationMethodConfigurationsAuthenticationMethodConfigurationCreateAuthenticationMethodConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthenticationMethodConfigurationsAuthenticationMethodConfigurationCreateAuthenticationMethodConfiguration`: MicrosoftGraphAuthenticationMethodConfiguration
    fmt.Fprintf(os.Stdout, "Response from `AuthenticationMethodConfigurationsAuthenticationMethodConfigurationApi.AuthenticationMethodConfigurationsAuthenticationMethodConfigurationCreateAuthenticationMethodConfiguration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthenticationMethodConfigurationsAuthenticationMethodConfigurationCreateAuthenticationMethodConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphAuthenticationMethodConfiguration** | [**MicrosoftGraphAuthenticationMethodConfiguration**](MicrosoftGraphAuthenticationMethodConfiguration.md) | New entity | 

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


## AuthenticationMethodConfigurationsAuthenticationMethodConfigurationDeleteAuthenticationMethodConfiguration

> AuthenticationMethodConfigurationsAuthenticationMethodConfigurationDeleteAuthenticationMethodConfiguration(ctx, authenticationMethodConfigurationId).IfMatch(ifMatch).Execute()

Delete entity from authenticationMethodConfigurations

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
    resp, r, err := api_client.AuthenticationMethodConfigurationsAuthenticationMethodConfigurationApi.AuthenticationMethodConfigurationsAuthenticationMethodConfigurationDeleteAuthenticationMethodConfiguration(context.Background(), authenticationMethodConfigurationId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationMethodConfigurationsAuthenticationMethodConfigurationApi.AuthenticationMethodConfigurationsAuthenticationMethodConfigurationDeleteAuthenticationMethodConfiguration``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAuthenticationMethodConfigurationsAuthenticationMethodConfigurationDeleteAuthenticationMethodConfigurationRequest struct via the builder pattern


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


## AuthenticationMethodConfigurationsAuthenticationMethodConfigurationGetAuthenticationMethodConfiguration

> MicrosoftGraphAuthenticationMethodConfiguration AuthenticationMethodConfigurationsAuthenticationMethodConfigurationGetAuthenticationMethodConfiguration(ctx, authenticationMethodConfigurationId).Select_(select_).Expand(expand).Execute()

Get entity from authenticationMethodConfigurations by key

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
    resp, r, err := api_client.AuthenticationMethodConfigurationsAuthenticationMethodConfigurationApi.AuthenticationMethodConfigurationsAuthenticationMethodConfigurationGetAuthenticationMethodConfiguration(context.Background(), authenticationMethodConfigurationId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationMethodConfigurationsAuthenticationMethodConfigurationApi.AuthenticationMethodConfigurationsAuthenticationMethodConfigurationGetAuthenticationMethodConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthenticationMethodConfigurationsAuthenticationMethodConfigurationGetAuthenticationMethodConfiguration`: MicrosoftGraphAuthenticationMethodConfiguration
    fmt.Fprintf(os.Stdout, "Response from `AuthenticationMethodConfigurationsAuthenticationMethodConfigurationApi.AuthenticationMethodConfigurationsAuthenticationMethodConfigurationGetAuthenticationMethodConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authenticationMethodConfigurationId** | **string** | key: id of authenticationMethodConfiguration | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuthenticationMethodConfigurationsAuthenticationMethodConfigurationGetAuthenticationMethodConfigurationRequest struct via the builder pattern


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


## AuthenticationMethodConfigurationsAuthenticationMethodConfigurationListAuthenticationMethodConfiguration

> CollectionOfAuthenticationMethodConfiguration AuthenticationMethodConfigurationsAuthenticationMethodConfigurationListAuthenticationMethodConfiguration(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get entities from authenticationMethodConfigurations

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
    resp, r, err := api_client.AuthenticationMethodConfigurationsAuthenticationMethodConfigurationApi.AuthenticationMethodConfigurationsAuthenticationMethodConfigurationListAuthenticationMethodConfiguration(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationMethodConfigurationsAuthenticationMethodConfigurationApi.AuthenticationMethodConfigurationsAuthenticationMethodConfigurationListAuthenticationMethodConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthenticationMethodConfigurationsAuthenticationMethodConfigurationListAuthenticationMethodConfiguration`: CollectionOfAuthenticationMethodConfiguration
    fmt.Fprintf(os.Stdout, "Response from `AuthenticationMethodConfigurationsAuthenticationMethodConfigurationApi.AuthenticationMethodConfigurationsAuthenticationMethodConfigurationListAuthenticationMethodConfiguration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthenticationMethodConfigurationsAuthenticationMethodConfigurationListAuthenticationMethodConfigurationRequest struct via the builder pattern


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


## AuthenticationMethodConfigurationsAuthenticationMethodConfigurationUpdateAuthenticationMethodConfiguration

> AuthenticationMethodConfigurationsAuthenticationMethodConfigurationUpdateAuthenticationMethodConfiguration(ctx, authenticationMethodConfigurationId).MicrosoftGraphAuthenticationMethodConfiguration(microsoftGraphAuthenticationMethodConfiguration).Execute()

Update entity in authenticationMethodConfigurations

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
    microsoftGraphAuthenticationMethodConfiguration := *openapiclient.NewMicrosoftGraphAuthenticationMethodConfiguration() // MicrosoftGraphAuthenticationMethodConfiguration | New property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthenticationMethodConfigurationsAuthenticationMethodConfigurationApi.AuthenticationMethodConfigurationsAuthenticationMethodConfigurationUpdateAuthenticationMethodConfiguration(context.Background(), authenticationMethodConfigurationId).MicrosoftGraphAuthenticationMethodConfiguration(microsoftGraphAuthenticationMethodConfiguration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationMethodConfigurationsAuthenticationMethodConfigurationApi.AuthenticationMethodConfigurationsAuthenticationMethodConfigurationUpdateAuthenticationMethodConfiguration``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAuthenticationMethodConfigurationsAuthenticationMethodConfigurationUpdateAuthenticationMethodConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphAuthenticationMethodConfiguration** | [**MicrosoftGraphAuthenticationMethodConfiguration**](MicrosoftGraphAuthenticationMethodConfiguration.md) | New property values | 

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

