# \CertificateBasedAuthConfigurationCertificateBasedAuthConfigurationApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CertificateBasedAuthConfigurationCertificateBasedAuthConfigurationCreateCertificateBasedAuthConfiguration**](CertificateBasedAuthConfigurationCertificateBasedAuthConfigurationApi.md#CertificateBasedAuthConfigurationCertificateBasedAuthConfigurationCreateCertificateBasedAuthConfiguration) | **Post** /certificateBasedAuthConfiguration | Add new entity to certificateBasedAuthConfiguration
[**CertificateBasedAuthConfigurationCertificateBasedAuthConfigurationDeleteCertificateBasedAuthConfiguration**](CertificateBasedAuthConfigurationCertificateBasedAuthConfigurationApi.md#CertificateBasedAuthConfigurationCertificateBasedAuthConfigurationDeleteCertificateBasedAuthConfiguration) | **Delete** /certificateBasedAuthConfiguration/{certificateBasedAuthConfiguration-id} | Delete entity from certificateBasedAuthConfiguration
[**CertificateBasedAuthConfigurationCertificateBasedAuthConfigurationGetCertificateBasedAuthConfiguration**](CertificateBasedAuthConfigurationCertificateBasedAuthConfigurationApi.md#CertificateBasedAuthConfigurationCertificateBasedAuthConfigurationGetCertificateBasedAuthConfiguration) | **Get** /certificateBasedAuthConfiguration/{certificateBasedAuthConfiguration-id} | Get entity from certificateBasedAuthConfiguration by key
[**CertificateBasedAuthConfigurationCertificateBasedAuthConfigurationListCertificateBasedAuthConfiguration**](CertificateBasedAuthConfigurationCertificateBasedAuthConfigurationApi.md#CertificateBasedAuthConfigurationCertificateBasedAuthConfigurationListCertificateBasedAuthConfiguration) | **Get** /certificateBasedAuthConfiguration | Get entities from certificateBasedAuthConfiguration
[**CertificateBasedAuthConfigurationCertificateBasedAuthConfigurationUpdateCertificateBasedAuthConfiguration**](CertificateBasedAuthConfigurationCertificateBasedAuthConfigurationApi.md#CertificateBasedAuthConfigurationCertificateBasedAuthConfigurationUpdateCertificateBasedAuthConfiguration) | **Patch** /certificateBasedAuthConfiguration/{certificateBasedAuthConfiguration-id} | Update entity in certificateBasedAuthConfiguration



## CertificateBasedAuthConfigurationCertificateBasedAuthConfigurationCreateCertificateBasedAuthConfiguration

> MicrosoftGraphCertificateBasedAuthConfiguration CertificateBasedAuthConfigurationCertificateBasedAuthConfigurationCreateCertificateBasedAuthConfiguration(ctx).MicrosoftGraphCertificateBasedAuthConfiguration(microsoftGraphCertificateBasedAuthConfiguration).Execute()

Add new entity to certificateBasedAuthConfiguration

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
    microsoftGraphCertificateBasedAuthConfiguration := *openapiclient.NewMicrosoftGraphCertificateBasedAuthConfiguration() // MicrosoftGraphCertificateBasedAuthConfiguration | New entity

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CertificateBasedAuthConfigurationCertificateBasedAuthConfigurationApi.CertificateBasedAuthConfigurationCertificateBasedAuthConfigurationCreateCertificateBasedAuthConfiguration(context.Background()).MicrosoftGraphCertificateBasedAuthConfiguration(microsoftGraphCertificateBasedAuthConfiguration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateBasedAuthConfigurationCertificateBasedAuthConfigurationApi.CertificateBasedAuthConfigurationCertificateBasedAuthConfigurationCreateCertificateBasedAuthConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CertificateBasedAuthConfigurationCertificateBasedAuthConfigurationCreateCertificateBasedAuthConfiguration`: MicrosoftGraphCertificateBasedAuthConfiguration
    fmt.Fprintf(os.Stdout, "Response from `CertificateBasedAuthConfigurationCertificateBasedAuthConfigurationApi.CertificateBasedAuthConfigurationCertificateBasedAuthConfigurationCreateCertificateBasedAuthConfiguration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCertificateBasedAuthConfigurationCertificateBasedAuthConfigurationCreateCertificateBasedAuthConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphCertificateBasedAuthConfiguration** | [**MicrosoftGraphCertificateBasedAuthConfiguration**](MicrosoftGraphCertificateBasedAuthConfiguration.md) | New entity | 

### Return type

[**MicrosoftGraphCertificateBasedAuthConfiguration**](MicrosoftGraphCertificateBasedAuthConfiguration.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CertificateBasedAuthConfigurationCertificateBasedAuthConfigurationDeleteCertificateBasedAuthConfiguration

> CertificateBasedAuthConfigurationCertificateBasedAuthConfigurationDeleteCertificateBasedAuthConfiguration(ctx, certificateBasedAuthConfigurationId).IfMatch(ifMatch).Execute()

Delete entity from certificateBasedAuthConfiguration

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
    certificateBasedAuthConfigurationId := "certificateBasedAuthConfigurationId_example" // string | key: id of certificateBasedAuthConfiguration
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CertificateBasedAuthConfigurationCertificateBasedAuthConfigurationApi.CertificateBasedAuthConfigurationCertificateBasedAuthConfigurationDeleteCertificateBasedAuthConfiguration(context.Background(), certificateBasedAuthConfigurationId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateBasedAuthConfigurationCertificateBasedAuthConfigurationApi.CertificateBasedAuthConfigurationCertificateBasedAuthConfigurationDeleteCertificateBasedAuthConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**certificateBasedAuthConfigurationId** | **string** | key: id of certificateBasedAuthConfiguration | 

### Other Parameters

Other parameters are passed through a pointer to a apiCertificateBasedAuthConfigurationCertificateBasedAuthConfigurationDeleteCertificateBasedAuthConfigurationRequest struct via the builder pattern


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


## CertificateBasedAuthConfigurationCertificateBasedAuthConfigurationGetCertificateBasedAuthConfiguration

> MicrosoftGraphCertificateBasedAuthConfiguration CertificateBasedAuthConfigurationCertificateBasedAuthConfigurationGetCertificateBasedAuthConfiguration(ctx, certificateBasedAuthConfigurationId).Select_(select_).Expand(expand).Execute()

Get entity from certificateBasedAuthConfiguration by key

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
    certificateBasedAuthConfigurationId := "certificateBasedAuthConfigurationId_example" // string | key: id of certificateBasedAuthConfiguration
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CertificateBasedAuthConfigurationCertificateBasedAuthConfigurationApi.CertificateBasedAuthConfigurationCertificateBasedAuthConfigurationGetCertificateBasedAuthConfiguration(context.Background(), certificateBasedAuthConfigurationId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateBasedAuthConfigurationCertificateBasedAuthConfigurationApi.CertificateBasedAuthConfigurationCertificateBasedAuthConfigurationGetCertificateBasedAuthConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CertificateBasedAuthConfigurationCertificateBasedAuthConfigurationGetCertificateBasedAuthConfiguration`: MicrosoftGraphCertificateBasedAuthConfiguration
    fmt.Fprintf(os.Stdout, "Response from `CertificateBasedAuthConfigurationCertificateBasedAuthConfigurationApi.CertificateBasedAuthConfigurationCertificateBasedAuthConfigurationGetCertificateBasedAuthConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**certificateBasedAuthConfigurationId** | **string** | key: id of certificateBasedAuthConfiguration | 

### Other Parameters

Other parameters are passed through a pointer to a apiCertificateBasedAuthConfigurationCertificateBasedAuthConfigurationGetCertificateBasedAuthConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphCertificateBasedAuthConfiguration**](MicrosoftGraphCertificateBasedAuthConfiguration.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CertificateBasedAuthConfigurationCertificateBasedAuthConfigurationListCertificateBasedAuthConfiguration

> CollectionOfCertificateBasedAuthConfiguration CertificateBasedAuthConfigurationCertificateBasedAuthConfigurationListCertificateBasedAuthConfiguration(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get entities from certificateBasedAuthConfiguration

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
    resp, r, err := api_client.CertificateBasedAuthConfigurationCertificateBasedAuthConfigurationApi.CertificateBasedAuthConfigurationCertificateBasedAuthConfigurationListCertificateBasedAuthConfiguration(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateBasedAuthConfigurationCertificateBasedAuthConfigurationApi.CertificateBasedAuthConfigurationCertificateBasedAuthConfigurationListCertificateBasedAuthConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CertificateBasedAuthConfigurationCertificateBasedAuthConfigurationListCertificateBasedAuthConfiguration`: CollectionOfCertificateBasedAuthConfiguration
    fmt.Fprintf(os.Stdout, "Response from `CertificateBasedAuthConfigurationCertificateBasedAuthConfigurationApi.CertificateBasedAuthConfigurationCertificateBasedAuthConfigurationListCertificateBasedAuthConfiguration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCertificateBasedAuthConfigurationCertificateBasedAuthConfigurationListCertificateBasedAuthConfigurationRequest struct via the builder pattern


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


## CertificateBasedAuthConfigurationCertificateBasedAuthConfigurationUpdateCertificateBasedAuthConfiguration

> CertificateBasedAuthConfigurationCertificateBasedAuthConfigurationUpdateCertificateBasedAuthConfiguration(ctx, certificateBasedAuthConfigurationId).MicrosoftGraphCertificateBasedAuthConfiguration(microsoftGraphCertificateBasedAuthConfiguration).Execute()

Update entity in certificateBasedAuthConfiguration

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
    certificateBasedAuthConfigurationId := "certificateBasedAuthConfigurationId_example" // string | key: id of certificateBasedAuthConfiguration
    microsoftGraphCertificateBasedAuthConfiguration := *openapiclient.NewMicrosoftGraphCertificateBasedAuthConfiguration() // MicrosoftGraphCertificateBasedAuthConfiguration | New property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CertificateBasedAuthConfigurationCertificateBasedAuthConfigurationApi.CertificateBasedAuthConfigurationCertificateBasedAuthConfigurationUpdateCertificateBasedAuthConfiguration(context.Background(), certificateBasedAuthConfigurationId).MicrosoftGraphCertificateBasedAuthConfiguration(microsoftGraphCertificateBasedAuthConfiguration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateBasedAuthConfigurationCertificateBasedAuthConfigurationApi.CertificateBasedAuthConfigurationCertificateBasedAuthConfigurationUpdateCertificateBasedAuthConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**certificateBasedAuthConfigurationId** | **string** | key: id of certificateBasedAuthConfiguration | 

### Other Parameters

Other parameters are passed through a pointer to a apiCertificateBasedAuthConfigurationCertificateBasedAuthConfigurationUpdateCertificateBasedAuthConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphCertificateBasedAuthConfiguration** | [**MicrosoftGraphCertificateBasedAuthConfiguration**](MicrosoftGraphCertificateBasedAuthConfiguration.md) | New property values | 

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

