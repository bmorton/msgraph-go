# \GroupSettingTemplatesGroupSettingTemplateApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GroupSettingTemplatesGroupSettingTemplateCreateGroupSettingTemplate**](GroupSettingTemplatesGroupSettingTemplateApi.md#GroupSettingTemplatesGroupSettingTemplateCreateGroupSettingTemplate) | **Post** /groupSettingTemplates | Add new entity to groupSettingTemplates
[**GroupSettingTemplatesGroupSettingTemplateDeleteGroupSettingTemplate**](GroupSettingTemplatesGroupSettingTemplateApi.md#GroupSettingTemplatesGroupSettingTemplateDeleteGroupSettingTemplate) | **Delete** /groupSettingTemplates/{groupSettingTemplate-id} | Delete entity from groupSettingTemplates
[**GroupSettingTemplatesGroupSettingTemplateGetGroupSettingTemplate**](GroupSettingTemplatesGroupSettingTemplateApi.md#GroupSettingTemplatesGroupSettingTemplateGetGroupSettingTemplate) | **Get** /groupSettingTemplates/{groupSettingTemplate-id} | Get entity from groupSettingTemplates by key
[**GroupSettingTemplatesGroupSettingTemplateListGroupSettingTemplate**](GroupSettingTemplatesGroupSettingTemplateApi.md#GroupSettingTemplatesGroupSettingTemplateListGroupSettingTemplate) | **Get** /groupSettingTemplates | Get entities from groupSettingTemplates
[**GroupSettingTemplatesGroupSettingTemplateUpdateGroupSettingTemplate**](GroupSettingTemplatesGroupSettingTemplateApi.md#GroupSettingTemplatesGroupSettingTemplateUpdateGroupSettingTemplate) | **Patch** /groupSettingTemplates/{groupSettingTemplate-id} | Update entity in groupSettingTemplates



## GroupSettingTemplatesGroupSettingTemplateCreateGroupSettingTemplate

> MicrosoftGraphGroupSettingTemplate GroupSettingTemplatesGroupSettingTemplateCreateGroupSettingTemplate(ctx).MicrosoftGraphGroupSettingTemplate(microsoftGraphGroupSettingTemplate).Execute()

Add new entity to groupSettingTemplates

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
    microsoftGraphGroupSettingTemplate := *openapiclient.NewMicrosoftGraphGroupSettingTemplate() // MicrosoftGraphGroupSettingTemplate | New entity

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupSettingTemplatesGroupSettingTemplateApi.GroupSettingTemplatesGroupSettingTemplateCreateGroupSettingTemplate(context.Background()).MicrosoftGraphGroupSettingTemplate(microsoftGraphGroupSettingTemplate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupSettingTemplatesGroupSettingTemplateApi.GroupSettingTemplatesGroupSettingTemplateCreateGroupSettingTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupSettingTemplatesGroupSettingTemplateCreateGroupSettingTemplate`: MicrosoftGraphGroupSettingTemplate
    fmt.Fprintf(os.Stdout, "Response from `GroupSettingTemplatesGroupSettingTemplateApi.GroupSettingTemplatesGroupSettingTemplateCreateGroupSettingTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGroupSettingTemplatesGroupSettingTemplateCreateGroupSettingTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphGroupSettingTemplate** | [**MicrosoftGraphGroupSettingTemplate**](MicrosoftGraphGroupSettingTemplate.md) | New entity | 

### Return type

[**MicrosoftGraphGroupSettingTemplate**](MicrosoftGraphGroupSettingTemplate.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupSettingTemplatesGroupSettingTemplateDeleteGroupSettingTemplate

> GroupSettingTemplatesGroupSettingTemplateDeleteGroupSettingTemplate(ctx, groupSettingTemplateId).IfMatch(ifMatch).Execute()

Delete entity from groupSettingTemplates

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
    groupSettingTemplateId := "groupSettingTemplateId_example" // string | key: id of groupSettingTemplate
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupSettingTemplatesGroupSettingTemplateApi.GroupSettingTemplatesGroupSettingTemplateDeleteGroupSettingTemplate(context.Background(), groupSettingTemplateId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupSettingTemplatesGroupSettingTemplateApi.GroupSettingTemplatesGroupSettingTemplateDeleteGroupSettingTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupSettingTemplateId** | **string** | key: id of groupSettingTemplate | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupSettingTemplatesGroupSettingTemplateDeleteGroupSettingTemplateRequest struct via the builder pattern


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


## GroupSettingTemplatesGroupSettingTemplateGetGroupSettingTemplate

> MicrosoftGraphGroupSettingTemplate GroupSettingTemplatesGroupSettingTemplateGetGroupSettingTemplate(ctx, groupSettingTemplateId).Select_(select_).Expand(expand).Execute()

Get entity from groupSettingTemplates by key

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
    groupSettingTemplateId := "groupSettingTemplateId_example" // string | key: id of groupSettingTemplate
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupSettingTemplatesGroupSettingTemplateApi.GroupSettingTemplatesGroupSettingTemplateGetGroupSettingTemplate(context.Background(), groupSettingTemplateId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupSettingTemplatesGroupSettingTemplateApi.GroupSettingTemplatesGroupSettingTemplateGetGroupSettingTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupSettingTemplatesGroupSettingTemplateGetGroupSettingTemplate`: MicrosoftGraphGroupSettingTemplate
    fmt.Fprintf(os.Stdout, "Response from `GroupSettingTemplatesGroupSettingTemplateApi.GroupSettingTemplatesGroupSettingTemplateGetGroupSettingTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupSettingTemplateId** | **string** | key: id of groupSettingTemplate | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupSettingTemplatesGroupSettingTemplateGetGroupSettingTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphGroupSettingTemplate**](MicrosoftGraphGroupSettingTemplate.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupSettingTemplatesGroupSettingTemplateListGroupSettingTemplate

> CollectionOfGroupSettingTemplate GroupSettingTemplatesGroupSettingTemplateListGroupSettingTemplate(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get entities from groupSettingTemplates

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
    resp, r, err := api_client.GroupSettingTemplatesGroupSettingTemplateApi.GroupSettingTemplatesGroupSettingTemplateListGroupSettingTemplate(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupSettingTemplatesGroupSettingTemplateApi.GroupSettingTemplatesGroupSettingTemplateListGroupSettingTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupSettingTemplatesGroupSettingTemplateListGroupSettingTemplate`: CollectionOfGroupSettingTemplate
    fmt.Fprintf(os.Stdout, "Response from `GroupSettingTemplatesGroupSettingTemplateApi.GroupSettingTemplatesGroupSettingTemplateListGroupSettingTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGroupSettingTemplatesGroupSettingTemplateListGroupSettingTemplateRequest struct via the builder pattern


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

[**CollectionOfGroupSettingTemplate**](CollectionOfGroupSettingTemplate.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupSettingTemplatesGroupSettingTemplateUpdateGroupSettingTemplate

> GroupSettingTemplatesGroupSettingTemplateUpdateGroupSettingTemplate(ctx, groupSettingTemplateId).MicrosoftGraphGroupSettingTemplate(microsoftGraphGroupSettingTemplate).Execute()

Update entity in groupSettingTemplates

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
    groupSettingTemplateId := "groupSettingTemplateId_example" // string | key: id of groupSettingTemplate
    microsoftGraphGroupSettingTemplate := *openapiclient.NewMicrosoftGraphGroupSettingTemplate() // MicrosoftGraphGroupSettingTemplate | New property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupSettingTemplatesGroupSettingTemplateApi.GroupSettingTemplatesGroupSettingTemplateUpdateGroupSettingTemplate(context.Background(), groupSettingTemplateId).MicrosoftGraphGroupSettingTemplate(microsoftGraphGroupSettingTemplate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupSettingTemplatesGroupSettingTemplateApi.GroupSettingTemplatesGroupSettingTemplateUpdateGroupSettingTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupSettingTemplateId** | **string** | key: id of groupSettingTemplate | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupSettingTemplatesGroupSettingTemplateUpdateGroupSettingTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphGroupSettingTemplate** | [**MicrosoftGraphGroupSettingTemplate**](MicrosoftGraphGroupSettingTemplate.md) | New property values | 

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

