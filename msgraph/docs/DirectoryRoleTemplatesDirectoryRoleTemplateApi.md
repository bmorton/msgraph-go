# \DirectoryRoleTemplatesDirectoryRoleTemplateApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DirectoryRoleTemplatesDirectoryRoleTemplateCreateDirectoryRoleTemplate**](DirectoryRoleTemplatesDirectoryRoleTemplateApi.md#DirectoryRoleTemplatesDirectoryRoleTemplateCreateDirectoryRoleTemplate) | **Post** /directoryRoleTemplates | Add new entity to directoryRoleTemplates
[**DirectoryRoleTemplatesDirectoryRoleTemplateDeleteDirectoryRoleTemplate**](DirectoryRoleTemplatesDirectoryRoleTemplateApi.md#DirectoryRoleTemplatesDirectoryRoleTemplateDeleteDirectoryRoleTemplate) | **Delete** /directoryRoleTemplates/{directoryRoleTemplate-id} | Delete entity from directoryRoleTemplates
[**DirectoryRoleTemplatesDirectoryRoleTemplateGetDirectoryRoleTemplate**](DirectoryRoleTemplatesDirectoryRoleTemplateApi.md#DirectoryRoleTemplatesDirectoryRoleTemplateGetDirectoryRoleTemplate) | **Get** /directoryRoleTemplates/{directoryRoleTemplate-id} | Get entity from directoryRoleTemplates by key
[**DirectoryRoleTemplatesDirectoryRoleTemplateListDirectoryRoleTemplate**](DirectoryRoleTemplatesDirectoryRoleTemplateApi.md#DirectoryRoleTemplatesDirectoryRoleTemplateListDirectoryRoleTemplate) | **Get** /directoryRoleTemplates | Get entities from directoryRoleTemplates
[**DirectoryRoleTemplatesDirectoryRoleTemplateUpdateDirectoryRoleTemplate**](DirectoryRoleTemplatesDirectoryRoleTemplateApi.md#DirectoryRoleTemplatesDirectoryRoleTemplateUpdateDirectoryRoleTemplate) | **Patch** /directoryRoleTemplates/{directoryRoleTemplate-id} | Update entity in directoryRoleTemplates



## DirectoryRoleTemplatesDirectoryRoleTemplateCreateDirectoryRoleTemplate

> MicrosoftGraphDirectoryRoleTemplate DirectoryRoleTemplatesDirectoryRoleTemplateCreateDirectoryRoleTemplate(ctx).MicrosoftGraphDirectoryRoleTemplate(microsoftGraphDirectoryRoleTemplate).Execute()

Add new entity to directoryRoleTemplates

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
    microsoftGraphDirectoryRoleTemplate := *openapiclient.NewMicrosoftGraphDirectoryRoleTemplate() // MicrosoftGraphDirectoryRoleTemplate | New entity

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DirectoryRoleTemplatesDirectoryRoleTemplateApi.DirectoryRoleTemplatesDirectoryRoleTemplateCreateDirectoryRoleTemplate(context.Background()).MicrosoftGraphDirectoryRoleTemplate(microsoftGraphDirectoryRoleTemplate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DirectoryRoleTemplatesDirectoryRoleTemplateApi.DirectoryRoleTemplatesDirectoryRoleTemplateCreateDirectoryRoleTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DirectoryRoleTemplatesDirectoryRoleTemplateCreateDirectoryRoleTemplate`: MicrosoftGraphDirectoryRoleTemplate
    fmt.Fprintf(os.Stdout, "Response from `DirectoryRoleTemplatesDirectoryRoleTemplateApi.DirectoryRoleTemplatesDirectoryRoleTemplateCreateDirectoryRoleTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDirectoryRoleTemplatesDirectoryRoleTemplateCreateDirectoryRoleTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphDirectoryRoleTemplate** | [**MicrosoftGraphDirectoryRoleTemplate**](MicrosoftGraphDirectoryRoleTemplate.md) | New entity | 

### Return type

[**MicrosoftGraphDirectoryRoleTemplate**](MicrosoftGraphDirectoryRoleTemplate.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DirectoryRoleTemplatesDirectoryRoleTemplateDeleteDirectoryRoleTemplate

> DirectoryRoleTemplatesDirectoryRoleTemplateDeleteDirectoryRoleTemplate(ctx, directoryRoleTemplateId).IfMatch(ifMatch).Execute()

Delete entity from directoryRoleTemplates

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
    directoryRoleTemplateId := "directoryRoleTemplateId_example" // string | key: id of directoryRoleTemplate
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DirectoryRoleTemplatesDirectoryRoleTemplateApi.DirectoryRoleTemplatesDirectoryRoleTemplateDeleteDirectoryRoleTemplate(context.Background(), directoryRoleTemplateId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DirectoryRoleTemplatesDirectoryRoleTemplateApi.DirectoryRoleTemplatesDirectoryRoleTemplateDeleteDirectoryRoleTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**directoryRoleTemplateId** | **string** | key: id of directoryRoleTemplate | 

### Other Parameters

Other parameters are passed through a pointer to a apiDirectoryRoleTemplatesDirectoryRoleTemplateDeleteDirectoryRoleTemplateRequest struct via the builder pattern


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


## DirectoryRoleTemplatesDirectoryRoleTemplateGetDirectoryRoleTemplate

> MicrosoftGraphDirectoryRoleTemplate DirectoryRoleTemplatesDirectoryRoleTemplateGetDirectoryRoleTemplate(ctx, directoryRoleTemplateId).Select_(select_).Expand(expand).Execute()

Get entity from directoryRoleTemplates by key

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
    directoryRoleTemplateId := "directoryRoleTemplateId_example" // string | key: id of directoryRoleTemplate
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DirectoryRoleTemplatesDirectoryRoleTemplateApi.DirectoryRoleTemplatesDirectoryRoleTemplateGetDirectoryRoleTemplate(context.Background(), directoryRoleTemplateId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DirectoryRoleTemplatesDirectoryRoleTemplateApi.DirectoryRoleTemplatesDirectoryRoleTemplateGetDirectoryRoleTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DirectoryRoleTemplatesDirectoryRoleTemplateGetDirectoryRoleTemplate`: MicrosoftGraphDirectoryRoleTemplate
    fmt.Fprintf(os.Stdout, "Response from `DirectoryRoleTemplatesDirectoryRoleTemplateApi.DirectoryRoleTemplatesDirectoryRoleTemplateGetDirectoryRoleTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**directoryRoleTemplateId** | **string** | key: id of directoryRoleTemplate | 

### Other Parameters

Other parameters are passed through a pointer to a apiDirectoryRoleTemplatesDirectoryRoleTemplateGetDirectoryRoleTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphDirectoryRoleTemplate**](MicrosoftGraphDirectoryRoleTemplate.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DirectoryRoleTemplatesDirectoryRoleTemplateListDirectoryRoleTemplate

> CollectionOfDirectoryRoleTemplate DirectoryRoleTemplatesDirectoryRoleTemplateListDirectoryRoleTemplate(ctx).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get entities from directoryRoleTemplates

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
    skip := int32(56) // int32 | Skip the first n items (optional)
    search := "search_example" // string | Search items by search phrases (optional)
    filter := "filter_example" // string | Filter items by property values (optional)
    count := true // bool | Include count of items (optional)
    orderby := []string{"Orderby_example"} // []string | Order items by property values (optional)
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DirectoryRoleTemplatesDirectoryRoleTemplateApi.DirectoryRoleTemplatesDirectoryRoleTemplateListDirectoryRoleTemplate(context.Background()).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DirectoryRoleTemplatesDirectoryRoleTemplateApi.DirectoryRoleTemplatesDirectoryRoleTemplateListDirectoryRoleTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DirectoryRoleTemplatesDirectoryRoleTemplateListDirectoryRoleTemplate`: CollectionOfDirectoryRoleTemplate
    fmt.Fprintf(os.Stdout, "Response from `DirectoryRoleTemplatesDirectoryRoleTemplateApi.DirectoryRoleTemplatesDirectoryRoleTemplateListDirectoryRoleTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDirectoryRoleTemplatesDirectoryRoleTemplateListDirectoryRoleTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **int32** | Skip the first n items | 
 **search** | **string** | Search items by search phrases | 
 **filter** | **string** | Filter items by property values | 
 **count** | **bool** | Include count of items | 
 **orderby** | **[]string** | Order items by property values | 
 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**CollectionOfDirectoryRoleTemplate**](CollectionOfDirectoryRoleTemplate.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DirectoryRoleTemplatesDirectoryRoleTemplateUpdateDirectoryRoleTemplate

> DirectoryRoleTemplatesDirectoryRoleTemplateUpdateDirectoryRoleTemplate(ctx, directoryRoleTemplateId).MicrosoftGraphDirectoryRoleTemplate(microsoftGraphDirectoryRoleTemplate).Execute()

Update entity in directoryRoleTemplates

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
    directoryRoleTemplateId := "directoryRoleTemplateId_example" // string | key: id of directoryRoleTemplate
    microsoftGraphDirectoryRoleTemplate := *openapiclient.NewMicrosoftGraphDirectoryRoleTemplate() // MicrosoftGraphDirectoryRoleTemplate | New property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DirectoryRoleTemplatesDirectoryRoleTemplateApi.DirectoryRoleTemplatesDirectoryRoleTemplateUpdateDirectoryRoleTemplate(context.Background(), directoryRoleTemplateId).MicrosoftGraphDirectoryRoleTemplate(microsoftGraphDirectoryRoleTemplate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DirectoryRoleTemplatesDirectoryRoleTemplateApi.DirectoryRoleTemplatesDirectoryRoleTemplateUpdateDirectoryRoleTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**directoryRoleTemplateId** | **string** | key: id of directoryRoleTemplate | 

### Other Parameters

Other parameters are passed through a pointer to a apiDirectoryRoleTemplatesDirectoryRoleTemplateUpdateDirectoryRoleTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphDirectoryRoleTemplate** | [**MicrosoftGraphDirectoryRoleTemplate**](MicrosoftGraphDirectoryRoleTemplate.md) | New property values | 

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

