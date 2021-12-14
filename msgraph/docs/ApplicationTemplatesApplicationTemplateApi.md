# \ApplicationTemplatesApplicationTemplateApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApplicationTemplatesApplicationTemplateCreateApplicationTemplate**](ApplicationTemplatesApplicationTemplateApi.md#ApplicationTemplatesApplicationTemplateCreateApplicationTemplate) | **Post** /applicationTemplates | Add new entity to applicationTemplates
[**ApplicationTemplatesApplicationTemplateDeleteApplicationTemplate**](ApplicationTemplatesApplicationTemplateApi.md#ApplicationTemplatesApplicationTemplateDeleteApplicationTemplate) | **Delete** /applicationTemplates/{applicationTemplate-id} | Delete entity from applicationTemplates
[**ApplicationTemplatesApplicationTemplateGetApplicationTemplate**](ApplicationTemplatesApplicationTemplateApi.md#ApplicationTemplatesApplicationTemplateGetApplicationTemplate) | **Get** /applicationTemplates/{applicationTemplate-id} | Get entity from applicationTemplates by key
[**ApplicationTemplatesApplicationTemplateListApplicationTemplate**](ApplicationTemplatesApplicationTemplateApi.md#ApplicationTemplatesApplicationTemplateListApplicationTemplate) | **Get** /applicationTemplates | Get entities from applicationTemplates
[**ApplicationTemplatesApplicationTemplateUpdateApplicationTemplate**](ApplicationTemplatesApplicationTemplateApi.md#ApplicationTemplatesApplicationTemplateUpdateApplicationTemplate) | **Patch** /applicationTemplates/{applicationTemplate-id} | Update entity in applicationTemplates



## ApplicationTemplatesApplicationTemplateCreateApplicationTemplate

> MicrosoftGraphApplicationTemplate ApplicationTemplatesApplicationTemplateCreateApplicationTemplate(ctx).MicrosoftGraphApplicationTemplate(microsoftGraphApplicationTemplate).Execute()

Add new entity to applicationTemplates

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
    microsoftGraphApplicationTemplate := *openapiclient.NewMicrosoftGraphApplicationTemplate() // MicrosoftGraphApplicationTemplate | New entity

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ApplicationTemplatesApplicationTemplateApi.ApplicationTemplatesApplicationTemplateCreateApplicationTemplate(context.Background()).MicrosoftGraphApplicationTemplate(microsoftGraphApplicationTemplate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationTemplatesApplicationTemplateApi.ApplicationTemplatesApplicationTemplateCreateApplicationTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApplicationTemplatesApplicationTemplateCreateApplicationTemplate`: MicrosoftGraphApplicationTemplate
    fmt.Fprintf(os.Stdout, "Response from `ApplicationTemplatesApplicationTemplateApi.ApplicationTemplatesApplicationTemplateCreateApplicationTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApplicationTemplatesApplicationTemplateCreateApplicationTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphApplicationTemplate** | [**MicrosoftGraphApplicationTemplate**](MicrosoftGraphApplicationTemplate.md) | New entity | 

### Return type

[**MicrosoftGraphApplicationTemplate**](MicrosoftGraphApplicationTemplate.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplicationTemplatesApplicationTemplateDeleteApplicationTemplate

> ApplicationTemplatesApplicationTemplateDeleteApplicationTemplate(ctx, applicationTemplateId).IfMatch(ifMatch).Execute()

Delete entity from applicationTemplates

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
    applicationTemplateId := "applicationTemplateId_example" // string | key: id of applicationTemplate
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ApplicationTemplatesApplicationTemplateApi.ApplicationTemplatesApplicationTemplateDeleteApplicationTemplate(context.Background(), applicationTemplateId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationTemplatesApplicationTemplateApi.ApplicationTemplatesApplicationTemplateDeleteApplicationTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationTemplateId** | **string** | key: id of applicationTemplate | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplicationTemplatesApplicationTemplateDeleteApplicationTemplateRequest struct via the builder pattern


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


## ApplicationTemplatesApplicationTemplateGetApplicationTemplate

> MicrosoftGraphApplicationTemplate ApplicationTemplatesApplicationTemplateGetApplicationTemplate(ctx, applicationTemplateId).Select_(select_).Expand(expand).Execute()

Get entity from applicationTemplates by key

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
    applicationTemplateId := "applicationTemplateId_example" // string | key: id of applicationTemplate
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ApplicationTemplatesApplicationTemplateApi.ApplicationTemplatesApplicationTemplateGetApplicationTemplate(context.Background(), applicationTemplateId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationTemplatesApplicationTemplateApi.ApplicationTemplatesApplicationTemplateGetApplicationTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApplicationTemplatesApplicationTemplateGetApplicationTemplate`: MicrosoftGraphApplicationTemplate
    fmt.Fprintf(os.Stdout, "Response from `ApplicationTemplatesApplicationTemplateApi.ApplicationTemplatesApplicationTemplateGetApplicationTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationTemplateId** | **string** | key: id of applicationTemplate | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplicationTemplatesApplicationTemplateGetApplicationTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphApplicationTemplate**](MicrosoftGraphApplicationTemplate.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplicationTemplatesApplicationTemplateListApplicationTemplate

> CollectionOfApplicationTemplate ApplicationTemplatesApplicationTemplateListApplicationTemplate(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get entities from applicationTemplates

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
    resp, r, err := api_client.ApplicationTemplatesApplicationTemplateApi.ApplicationTemplatesApplicationTemplateListApplicationTemplate(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationTemplatesApplicationTemplateApi.ApplicationTemplatesApplicationTemplateListApplicationTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApplicationTemplatesApplicationTemplateListApplicationTemplate`: CollectionOfApplicationTemplate
    fmt.Fprintf(os.Stdout, "Response from `ApplicationTemplatesApplicationTemplateApi.ApplicationTemplatesApplicationTemplateListApplicationTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApplicationTemplatesApplicationTemplateListApplicationTemplateRequest struct via the builder pattern


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

[**CollectionOfApplicationTemplate**](CollectionOfApplicationTemplate.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplicationTemplatesApplicationTemplateUpdateApplicationTemplate

> ApplicationTemplatesApplicationTemplateUpdateApplicationTemplate(ctx, applicationTemplateId).MicrosoftGraphApplicationTemplate(microsoftGraphApplicationTemplate).Execute()

Update entity in applicationTemplates

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
    applicationTemplateId := "applicationTemplateId_example" // string | key: id of applicationTemplate
    microsoftGraphApplicationTemplate := *openapiclient.NewMicrosoftGraphApplicationTemplate() // MicrosoftGraphApplicationTemplate | New property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ApplicationTemplatesApplicationTemplateApi.ApplicationTemplatesApplicationTemplateUpdateApplicationTemplate(context.Background(), applicationTemplateId).MicrosoftGraphApplicationTemplate(microsoftGraphApplicationTemplate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationTemplatesApplicationTemplateApi.ApplicationTemplatesApplicationTemplateUpdateApplicationTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationTemplateId** | **string** | key: id of applicationTemplate | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplicationTemplatesApplicationTemplateUpdateApplicationTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphApplicationTemplate** | [**MicrosoftGraphApplicationTemplate**](MicrosoftGraphApplicationTemplate.md) | New property values | 

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

