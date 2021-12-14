# \DirectoryRoleTemplatesActionsApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DirectoryRoleTemplatesDirectoryRoleTemplateCheckMemberGroups**](DirectoryRoleTemplatesActionsApi.md#DirectoryRoleTemplatesDirectoryRoleTemplateCheckMemberGroups) | **Post** /directoryRoleTemplates/{directoryRoleTemplate-id}/microsoft.graph.checkMemberGroups | Invoke action checkMemberGroups
[**DirectoryRoleTemplatesDirectoryRoleTemplateCheckMemberObjects**](DirectoryRoleTemplatesActionsApi.md#DirectoryRoleTemplatesDirectoryRoleTemplateCheckMemberObjects) | **Post** /directoryRoleTemplates/{directoryRoleTemplate-id}/microsoft.graph.checkMemberObjects | Invoke action checkMemberObjects
[**DirectoryRoleTemplatesDirectoryRoleTemplateGetMemberGroups**](DirectoryRoleTemplatesActionsApi.md#DirectoryRoleTemplatesDirectoryRoleTemplateGetMemberGroups) | **Post** /directoryRoleTemplates/{directoryRoleTemplate-id}/microsoft.graph.getMemberGroups | Invoke action getMemberGroups
[**DirectoryRoleTemplatesDirectoryRoleTemplateGetMemberObjects**](DirectoryRoleTemplatesActionsApi.md#DirectoryRoleTemplatesDirectoryRoleTemplateGetMemberObjects) | **Post** /directoryRoleTemplates/{directoryRoleTemplate-id}/microsoft.graph.getMemberObjects | Invoke action getMemberObjects
[**DirectoryRoleTemplatesDirectoryRoleTemplateRestore**](DirectoryRoleTemplatesActionsApi.md#DirectoryRoleTemplatesDirectoryRoleTemplateRestore) | **Post** /directoryRoleTemplates/{directoryRoleTemplate-id}/microsoft.graph.restore | Invoke action restore
[**DirectoryRoleTemplatesGetAvailableExtensionProperties**](DirectoryRoleTemplatesActionsApi.md#DirectoryRoleTemplatesGetAvailableExtensionProperties) | **Post** /directoryRoleTemplates/microsoft.graph.getAvailableExtensionProperties | Invoke action getAvailableExtensionProperties
[**DirectoryRoleTemplatesGetByIds**](DirectoryRoleTemplatesActionsApi.md#DirectoryRoleTemplatesGetByIds) | **Post** /directoryRoleTemplates/microsoft.graph.getByIds | Invoke action getByIds
[**DirectoryRoleTemplatesValidateProperties**](DirectoryRoleTemplatesActionsApi.md#DirectoryRoleTemplatesValidateProperties) | **Post** /directoryRoleTemplates/microsoft.graph.validateProperties | Invoke action validateProperties



## DirectoryRoleTemplatesDirectoryRoleTemplateCheckMemberGroups

> []string DirectoryRoleTemplatesDirectoryRoleTemplateCheckMemberGroups(ctx, directoryRoleTemplateId).InlineObject127(inlineObject127).Execute()

Invoke action checkMemberGroups

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
    inlineObject127 := *openapiclient.NewInlineObject127() // InlineObject127 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DirectoryRoleTemplatesActionsApi.DirectoryRoleTemplatesDirectoryRoleTemplateCheckMemberGroups(context.Background(), directoryRoleTemplateId).InlineObject127(inlineObject127).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DirectoryRoleTemplatesActionsApi.DirectoryRoleTemplatesDirectoryRoleTemplateCheckMemberGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DirectoryRoleTemplatesDirectoryRoleTemplateCheckMemberGroups`: []string
    fmt.Fprintf(os.Stdout, "Response from `DirectoryRoleTemplatesActionsApi.DirectoryRoleTemplatesDirectoryRoleTemplateCheckMemberGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**directoryRoleTemplateId** | **string** | key: id of directoryRoleTemplate | 

### Other Parameters

Other parameters are passed through a pointer to a apiDirectoryRoleTemplatesDirectoryRoleTemplateCheckMemberGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject127** | [**InlineObject127**](InlineObject127.md) |  | 

### Return type

**[]string**

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DirectoryRoleTemplatesDirectoryRoleTemplateCheckMemberObjects

> []string DirectoryRoleTemplatesDirectoryRoleTemplateCheckMemberObjects(ctx, directoryRoleTemplateId).InlineObject128(inlineObject128).Execute()

Invoke action checkMemberObjects

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
    inlineObject128 := *openapiclient.NewInlineObject128() // InlineObject128 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DirectoryRoleTemplatesActionsApi.DirectoryRoleTemplatesDirectoryRoleTemplateCheckMemberObjects(context.Background(), directoryRoleTemplateId).InlineObject128(inlineObject128).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DirectoryRoleTemplatesActionsApi.DirectoryRoleTemplatesDirectoryRoleTemplateCheckMemberObjects``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DirectoryRoleTemplatesDirectoryRoleTemplateCheckMemberObjects`: []string
    fmt.Fprintf(os.Stdout, "Response from `DirectoryRoleTemplatesActionsApi.DirectoryRoleTemplatesDirectoryRoleTemplateCheckMemberObjects`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**directoryRoleTemplateId** | **string** | key: id of directoryRoleTemplate | 

### Other Parameters

Other parameters are passed through a pointer to a apiDirectoryRoleTemplatesDirectoryRoleTemplateCheckMemberObjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject128** | [**InlineObject128**](InlineObject128.md) |  | 

### Return type

**[]string**

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DirectoryRoleTemplatesDirectoryRoleTemplateGetMemberGroups

> []string DirectoryRoleTemplatesDirectoryRoleTemplateGetMemberGroups(ctx, directoryRoleTemplateId).InlineObject129(inlineObject129).Execute()

Invoke action getMemberGroups

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
    inlineObject129 := *openapiclient.NewInlineObject129() // InlineObject129 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DirectoryRoleTemplatesActionsApi.DirectoryRoleTemplatesDirectoryRoleTemplateGetMemberGroups(context.Background(), directoryRoleTemplateId).InlineObject129(inlineObject129).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DirectoryRoleTemplatesActionsApi.DirectoryRoleTemplatesDirectoryRoleTemplateGetMemberGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DirectoryRoleTemplatesDirectoryRoleTemplateGetMemberGroups`: []string
    fmt.Fprintf(os.Stdout, "Response from `DirectoryRoleTemplatesActionsApi.DirectoryRoleTemplatesDirectoryRoleTemplateGetMemberGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**directoryRoleTemplateId** | **string** | key: id of directoryRoleTemplate | 

### Other Parameters

Other parameters are passed through a pointer to a apiDirectoryRoleTemplatesDirectoryRoleTemplateGetMemberGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject129** | [**InlineObject129**](InlineObject129.md) |  | 

### Return type

**[]string**

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DirectoryRoleTemplatesDirectoryRoleTemplateGetMemberObjects

> []string DirectoryRoleTemplatesDirectoryRoleTemplateGetMemberObjects(ctx, directoryRoleTemplateId).InlineObject130(inlineObject130).Execute()

Invoke action getMemberObjects

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
    inlineObject130 := *openapiclient.NewInlineObject130() // InlineObject130 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DirectoryRoleTemplatesActionsApi.DirectoryRoleTemplatesDirectoryRoleTemplateGetMemberObjects(context.Background(), directoryRoleTemplateId).InlineObject130(inlineObject130).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DirectoryRoleTemplatesActionsApi.DirectoryRoleTemplatesDirectoryRoleTemplateGetMemberObjects``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DirectoryRoleTemplatesDirectoryRoleTemplateGetMemberObjects`: []string
    fmt.Fprintf(os.Stdout, "Response from `DirectoryRoleTemplatesActionsApi.DirectoryRoleTemplatesDirectoryRoleTemplateGetMemberObjects`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**directoryRoleTemplateId** | **string** | key: id of directoryRoleTemplate | 

### Other Parameters

Other parameters are passed through a pointer to a apiDirectoryRoleTemplatesDirectoryRoleTemplateGetMemberObjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject130** | [**InlineObject130**](InlineObject130.md) |  | 

### Return type

**[]string**

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DirectoryRoleTemplatesDirectoryRoleTemplateRestore

> AnyOfmicrosoftGraphDirectoryObject DirectoryRoleTemplatesDirectoryRoleTemplateRestore(ctx, directoryRoleTemplateId).Execute()

Invoke action restore

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DirectoryRoleTemplatesActionsApi.DirectoryRoleTemplatesDirectoryRoleTemplateRestore(context.Background(), directoryRoleTemplateId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DirectoryRoleTemplatesActionsApi.DirectoryRoleTemplatesDirectoryRoleTemplateRestore``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DirectoryRoleTemplatesDirectoryRoleTemplateRestore`: AnyOfmicrosoftGraphDirectoryObject
    fmt.Fprintf(os.Stdout, "Response from `DirectoryRoleTemplatesActionsApi.DirectoryRoleTemplatesDirectoryRoleTemplateRestore`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**directoryRoleTemplateId** | **string** | key: id of directoryRoleTemplate | 

### Other Parameters

Other parameters are passed through a pointer to a apiDirectoryRoleTemplatesDirectoryRoleTemplateRestoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AnyOfmicrosoftGraphDirectoryObject**](anyOf&lt;microsoft.graph.directoryObject&gt;.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DirectoryRoleTemplatesGetAvailableExtensionProperties

> []MicrosoftGraphExtensionProperty DirectoryRoleTemplatesGetAvailableExtensionProperties(ctx).InlineObject131(inlineObject131).Execute()

Invoke action getAvailableExtensionProperties

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
    inlineObject131 := *openapiclient.NewInlineObject131() // InlineObject131 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DirectoryRoleTemplatesActionsApi.DirectoryRoleTemplatesGetAvailableExtensionProperties(context.Background()).InlineObject131(inlineObject131).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DirectoryRoleTemplatesActionsApi.DirectoryRoleTemplatesGetAvailableExtensionProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DirectoryRoleTemplatesGetAvailableExtensionProperties`: []MicrosoftGraphExtensionProperty
    fmt.Fprintf(os.Stdout, "Response from `DirectoryRoleTemplatesActionsApi.DirectoryRoleTemplatesGetAvailableExtensionProperties`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDirectoryRoleTemplatesGetAvailableExtensionPropertiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject131** | [**InlineObject131**](InlineObject131.md) |  | 

### Return type

[**[]MicrosoftGraphExtensionProperty**](MicrosoftGraphExtensionProperty.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DirectoryRoleTemplatesGetByIds

> []MicrosoftGraphDirectoryObject DirectoryRoleTemplatesGetByIds(ctx).InlineObject132(inlineObject132).Execute()

Invoke action getByIds

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
    inlineObject132 := *openapiclient.NewInlineObject132() // InlineObject132 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DirectoryRoleTemplatesActionsApi.DirectoryRoleTemplatesGetByIds(context.Background()).InlineObject132(inlineObject132).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DirectoryRoleTemplatesActionsApi.DirectoryRoleTemplatesGetByIds``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DirectoryRoleTemplatesGetByIds`: []MicrosoftGraphDirectoryObject
    fmt.Fprintf(os.Stdout, "Response from `DirectoryRoleTemplatesActionsApi.DirectoryRoleTemplatesGetByIds`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDirectoryRoleTemplatesGetByIdsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject132** | [**InlineObject132**](InlineObject132.md) |  | 

### Return type

[**[]MicrosoftGraphDirectoryObject**](MicrosoftGraphDirectoryObject.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DirectoryRoleTemplatesValidateProperties

> DirectoryRoleTemplatesValidateProperties(ctx).InlineObject133(inlineObject133).Execute()

Invoke action validateProperties

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
    inlineObject133 := *openapiclient.NewInlineObject133() // InlineObject133 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DirectoryRoleTemplatesActionsApi.DirectoryRoleTemplatesValidateProperties(context.Background()).InlineObject133(inlineObject133).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DirectoryRoleTemplatesActionsApi.DirectoryRoleTemplatesValidateProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDirectoryRoleTemplatesValidatePropertiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject133** | [**InlineObject133**](InlineObject133.md) |  | 

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

