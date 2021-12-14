# \TeamsTemplatesTeamsTemplateApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TeamsTemplatesTeamsTemplateCreateTeamsTemplate**](TeamsTemplatesTeamsTemplateApi.md#TeamsTemplatesTeamsTemplateCreateTeamsTemplate) | **Post** /teamsTemplates | Add new entity to teamsTemplates
[**TeamsTemplatesTeamsTemplateDeleteTeamsTemplate**](TeamsTemplatesTeamsTemplateApi.md#TeamsTemplatesTeamsTemplateDeleteTeamsTemplate) | **Delete** /teamsTemplates/{teamsTemplate-id} | Delete entity from teamsTemplates
[**TeamsTemplatesTeamsTemplateGetTeamsTemplate**](TeamsTemplatesTeamsTemplateApi.md#TeamsTemplatesTeamsTemplateGetTeamsTemplate) | **Get** /teamsTemplates/{teamsTemplate-id} | Get entity from teamsTemplates by key
[**TeamsTemplatesTeamsTemplateListTeamsTemplate**](TeamsTemplatesTeamsTemplateApi.md#TeamsTemplatesTeamsTemplateListTeamsTemplate) | **Get** /teamsTemplates | Get entities from teamsTemplates
[**TeamsTemplatesTeamsTemplateUpdateTeamsTemplate**](TeamsTemplatesTeamsTemplateApi.md#TeamsTemplatesTeamsTemplateUpdateTeamsTemplate) | **Patch** /teamsTemplates/{teamsTemplate-id} | Update entity in teamsTemplates



## TeamsTemplatesTeamsTemplateCreateTeamsTemplate

> MicrosoftGraphTeamsTemplate TeamsTemplatesTeamsTemplateCreateTeamsTemplate(ctx).MicrosoftGraphTeamsTemplate(microsoftGraphTeamsTemplate).Execute()

Add new entity to teamsTemplates

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
    microsoftGraphTeamsTemplate := *openapiclient.NewMicrosoftGraphTeamsTemplate() // MicrosoftGraphTeamsTemplate | New entity

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsTemplatesTeamsTemplateApi.TeamsTemplatesTeamsTemplateCreateTeamsTemplate(context.Background()).MicrosoftGraphTeamsTemplate(microsoftGraphTeamsTemplate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsTemplatesTeamsTemplateApi.TeamsTemplatesTeamsTemplateCreateTeamsTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsTemplatesTeamsTemplateCreateTeamsTemplate`: MicrosoftGraphTeamsTemplate
    fmt.Fprintf(os.Stdout, "Response from `TeamsTemplatesTeamsTemplateApi.TeamsTemplatesTeamsTemplateCreateTeamsTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTeamsTemplatesTeamsTemplateCreateTeamsTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphTeamsTemplate** | [**MicrosoftGraphTeamsTemplate**](MicrosoftGraphTeamsTemplate.md) | New entity | 

### Return type

[**MicrosoftGraphTeamsTemplate**](MicrosoftGraphTeamsTemplate.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsTemplatesTeamsTemplateDeleteTeamsTemplate

> TeamsTemplatesTeamsTemplateDeleteTeamsTemplate(ctx, teamsTemplateId).IfMatch(ifMatch).Execute()

Delete entity from teamsTemplates

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
    teamsTemplateId := "teamsTemplateId_example" // string | key: id of teamsTemplate
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsTemplatesTeamsTemplateApi.TeamsTemplatesTeamsTemplateDeleteTeamsTemplate(context.Background(), teamsTemplateId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsTemplatesTeamsTemplateApi.TeamsTemplatesTeamsTemplateDeleteTeamsTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamsTemplateId** | **string** | key: id of teamsTemplate | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsTemplatesTeamsTemplateDeleteTeamsTemplateRequest struct via the builder pattern


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


## TeamsTemplatesTeamsTemplateGetTeamsTemplate

> MicrosoftGraphTeamsTemplate TeamsTemplatesTeamsTemplateGetTeamsTemplate(ctx, teamsTemplateId).Select_(select_).Expand(expand).Execute()

Get entity from teamsTemplates by key

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
    teamsTemplateId := "teamsTemplateId_example" // string | key: id of teamsTemplate
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsTemplatesTeamsTemplateApi.TeamsTemplatesTeamsTemplateGetTeamsTemplate(context.Background(), teamsTemplateId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsTemplatesTeamsTemplateApi.TeamsTemplatesTeamsTemplateGetTeamsTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsTemplatesTeamsTemplateGetTeamsTemplate`: MicrosoftGraphTeamsTemplate
    fmt.Fprintf(os.Stdout, "Response from `TeamsTemplatesTeamsTemplateApi.TeamsTemplatesTeamsTemplateGetTeamsTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamsTemplateId** | **string** | key: id of teamsTemplate | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsTemplatesTeamsTemplateGetTeamsTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphTeamsTemplate**](MicrosoftGraphTeamsTemplate.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsTemplatesTeamsTemplateListTeamsTemplate

> CollectionOfTeamsTemplate TeamsTemplatesTeamsTemplateListTeamsTemplate(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get entities from teamsTemplates

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
    resp, r, err := api_client.TeamsTemplatesTeamsTemplateApi.TeamsTemplatesTeamsTemplateListTeamsTemplate(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsTemplatesTeamsTemplateApi.TeamsTemplatesTeamsTemplateListTeamsTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsTemplatesTeamsTemplateListTeamsTemplate`: CollectionOfTeamsTemplate
    fmt.Fprintf(os.Stdout, "Response from `TeamsTemplatesTeamsTemplateApi.TeamsTemplatesTeamsTemplateListTeamsTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTeamsTemplatesTeamsTemplateListTeamsTemplateRequest struct via the builder pattern


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

[**CollectionOfTeamsTemplate**](CollectionOfTeamsTemplate.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsTemplatesTeamsTemplateUpdateTeamsTemplate

> TeamsTemplatesTeamsTemplateUpdateTeamsTemplate(ctx, teamsTemplateId).MicrosoftGraphTeamsTemplate(microsoftGraphTeamsTemplate).Execute()

Update entity in teamsTemplates

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
    teamsTemplateId := "teamsTemplateId_example" // string | key: id of teamsTemplate
    microsoftGraphTeamsTemplate := *openapiclient.NewMicrosoftGraphTeamsTemplate() // MicrosoftGraphTeamsTemplate | New property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsTemplatesTeamsTemplateApi.TeamsTemplatesTeamsTemplateUpdateTeamsTemplate(context.Background(), teamsTemplateId).MicrosoftGraphTeamsTemplate(microsoftGraphTeamsTemplate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsTemplatesTeamsTemplateApi.TeamsTemplatesTeamsTemplateUpdateTeamsTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamsTemplateId** | **string** | key: id of teamsTemplate | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsTemplatesTeamsTemplateUpdateTeamsTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphTeamsTemplate** | [**MicrosoftGraphTeamsTemplate**](MicrosoftGraphTeamsTemplate.md) | New property values | 

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

