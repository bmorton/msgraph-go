# \ApplicationsExtensionPropertyApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApplicationsCreateExtensionProperties**](ApplicationsExtensionPropertyApi.md#ApplicationsCreateExtensionProperties) | **Post** /applications/{application-id}/extensionProperties | Create new navigation property to extensionProperties for applications
[**ApplicationsDeleteExtensionProperties**](ApplicationsExtensionPropertyApi.md#ApplicationsDeleteExtensionProperties) | **Delete** /applications/{application-id}/extensionProperties/{extensionProperty-id} | Delete navigation property extensionProperties for applications
[**ApplicationsGetExtensionProperties**](ApplicationsExtensionPropertyApi.md#ApplicationsGetExtensionProperties) | **Get** /applications/{application-id}/extensionProperties/{extensionProperty-id} | Get extensionProperties from applications
[**ApplicationsListExtensionProperties**](ApplicationsExtensionPropertyApi.md#ApplicationsListExtensionProperties) | **Get** /applications/{application-id}/extensionProperties | Get extensionProperties from applications
[**ApplicationsUpdateExtensionProperties**](ApplicationsExtensionPropertyApi.md#ApplicationsUpdateExtensionProperties) | **Patch** /applications/{application-id}/extensionProperties/{extensionProperty-id} | Update the navigation property extensionProperties in applications



## ApplicationsCreateExtensionProperties

> MicrosoftGraphExtensionProperty ApplicationsCreateExtensionProperties(ctx, applicationId).MicrosoftGraphExtensionProperty(microsoftGraphExtensionProperty).Execute()

Create new navigation property to extensionProperties for applications



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
    applicationId := "applicationId_example" // string | key: id of application
    microsoftGraphExtensionProperty := *openapiclient.NewMicrosoftGraphExtensionProperty() // MicrosoftGraphExtensionProperty | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ApplicationsExtensionPropertyApi.ApplicationsCreateExtensionProperties(context.Background(), applicationId).MicrosoftGraphExtensionProperty(microsoftGraphExtensionProperty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsExtensionPropertyApi.ApplicationsCreateExtensionProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApplicationsCreateExtensionProperties`: MicrosoftGraphExtensionProperty
    fmt.Fprintf(os.Stdout, "Response from `ApplicationsExtensionPropertyApi.ApplicationsCreateExtensionProperties`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | key: id of application | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplicationsCreateExtensionPropertiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphExtensionProperty** | [**MicrosoftGraphExtensionProperty**](MicrosoftGraphExtensionProperty.md) | New navigation property | 

### Return type

[**MicrosoftGraphExtensionProperty**](MicrosoftGraphExtensionProperty.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplicationsDeleteExtensionProperties

> ApplicationsDeleteExtensionProperties(ctx, applicationId, extensionPropertyId).IfMatch(ifMatch).Execute()

Delete navigation property extensionProperties for applications



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
    applicationId := "applicationId_example" // string | key: id of application
    extensionPropertyId := "extensionPropertyId_example" // string | key: id of extensionProperty
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ApplicationsExtensionPropertyApi.ApplicationsDeleteExtensionProperties(context.Background(), applicationId, extensionPropertyId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsExtensionPropertyApi.ApplicationsDeleteExtensionProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | key: id of application | 
**extensionPropertyId** | **string** | key: id of extensionProperty | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplicationsDeleteExtensionPropertiesRequest struct via the builder pattern


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


## ApplicationsGetExtensionProperties

> MicrosoftGraphExtensionProperty ApplicationsGetExtensionProperties(ctx, applicationId, extensionPropertyId).Select_(select_).Expand(expand).Execute()

Get extensionProperties from applications



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
    applicationId := "applicationId_example" // string | key: id of application
    extensionPropertyId := "extensionPropertyId_example" // string | key: id of extensionProperty
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ApplicationsExtensionPropertyApi.ApplicationsGetExtensionProperties(context.Background(), applicationId, extensionPropertyId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsExtensionPropertyApi.ApplicationsGetExtensionProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApplicationsGetExtensionProperties`: MicrosoftGraphExtensionProperty
    fmt.Fprintf(os.Stdout, "Response from `ApplicationsExtensionPropertyApi.ApplicationsGetExtensionProperties`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | key: id of application | 
**extensionPropertyId** | **string** | key: id of extensionProperty | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplicationsGetExtensionPropertiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphExtensionProperty**](MicrosoftGraphExtensionProperty.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplicationsListExtensionProperties

> CollectionOfExtensionProperty ApplicationsListExtensionProperties(ctx, applicationId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get extensionProperties from applications



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
    applicationId := "applicationId_example" // string | key: id of application
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
    resp, r, err := api_client.ApplicationsExtensionPropertyApi.ApplicationsListExtensionProperties(context.Background(), applicationId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsExtensionPropertyApi.ApplicationsListExtensionProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApplicationsListExtensionProperties`: CollectionOfExtensionProperty
    fmt.Fprintf(os.Stdout, "Response from `ApplicationsExtensionPropertyApi.ApplicationsListExtensionProperties`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | key: id of application | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplicationsListExtensionPropertiesRequest struct via the builder pattern


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

[**CollectionOfExtensionProperty**](CollectionOfExtensionProperty.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplicationsUpdateExtensionProperties

> ApplicationsUpdateExtensionProperties(ctx, applicationId, extensionPropertyId).MicrosoftGraphExtensionProperty(microsoftGraphExtensionProperty).Execute()

Update the navigation property extensionProperties in applications



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
    applicationId := "applicationId_example" // string | key: id of application
    extensionPropertyId := "extensionPropertyId_example" // string | key: id of extensionProperty
    microsoftGraphExtensionProperty := *openapiclient.NewMicrosoftGraphExtensionProperty() // MicrosoftGraphExtensionProperty | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ApplicationsExtensionPropertyApi.ApplicationsUpdateExtensionProperties(context.Background(), applicationId, extensionPropertyId).MicrosoftGraphExtensionProperty(microsoftGraphExtensionProperty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsExtensionPropertyApi.ApplicationsUpdateExtensionProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | key: id of application | 
**extensionPropertyId** | **string** | key: id of extensionProperty | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplicationsUpdateExtensionPropertiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **microsoftGraphExtensionProperty** | [**MicrosoftGraphExtensionProperty**](MicrosoftGraphExtensionProperty.md) | New navigation property values | 

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

