# \GroupSettingTemplatesActionsApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GroupSettingTemplatesGetAvailableExtensionProperties**](GroupSettingTemplatesActionsApi.md#GroupSettingTemplatesGetAvailableExtensionProperties) | **Post** /groupSettingTemplates/microsoft.graph.getAvailableExtensionProperties | Invoke action getAvailableExtensionProperties
[**GroupSettingTemplatesGetByIds**](GroupSettingTemplatesActionsApi.md#GroupSettingTemplatesGetByIds) | **Post** /groupSettingTemplates/microsoft.graph.getByIds | Invoke action getByIds
[**GroupSettingTemplatesGroupSettingTemplateCheckMemberGroups**](GroupSettingTemplatesActionsApi.md#GroupSettingTemplatesGroupSettingTemplateCheckMemberGroups) | **Post** /groupSettingTemplates/{groupSettingTemplate-id}/microsoft.graph.checkMemberGroups | Invoke action checkMemberGroups
[**GroupSettingTemplatesGroupSettingTemplateCheckMemberObjects**](GroupSettingTemplatesActionsApi.md#GroupSettingTemplatesGroupSettingTemplateCheckMemberObjects) | **Post** /groupSettingTemplates/{groupSettingTemplate-id}/microsoft.graph.checkMemberObjects | Invoke action checkMemberObjects
[**GroupSettingTemplatesGroupSettingTemplateGetMemberGroups**](GroupSettingTemplatesActionsApi.md#GroupSettingTemplatesGroupSettingTemplateGetMemberGroups) | **Post** /groupSettingTemplates/{groupSettingTemplate-id}/microsoft.graph.getMemberGroups | Invoke action getMemberGroups
[**GroupSettingTemplatesGroupSettingTemplateGetMemberObjects**](GroupSettingTemplatesActionsApi.md#GroupSettingTemplatesGroupSettingTemplateGetMemberObjects) | **Post** /groupSettingTemplates/{groupSettingTemplate-id}/microsoft.graph.getMemberObjects | Invoke action getMemberObjects
[**GroupSettingTemplatesGroupSettingTemplateRestore**](GroupSettingTemplatesActionsApi.md#GroupSettingTemplatesGroupSettingTemplateRestore) | **Post** /groupSettingTemplates/{groupSettingTemplate-id}/microsoft.graph.restore | Invoke action restore
[**GroupSettingTemplatesValidateProperties**](GroupSettingTemplatesActionsApi.md#GroupSettingTemplatesValidateProperties) | **Post** /groupSettingTemplates/microsoft.graph.validateProperties | Invoke action validateProperties



## GroupSettingTemplatesGetAvailableExtensionProperties

> []MicrosoftGraphExtensionProperty GroupSettingTemplatesGetAvailableExtensionProperties(ctx).InlineObject353(inlineObject353).Execute()

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
    inlineObject353 := *openapiclient.NewInlineObject353() // InlineObject353 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupSettingTemplatesActionsApi.GroupSettingTemplatesGetAvailableExtensionProperties(context.Background()).InlineObject353(inlineObject353).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupSettingTemplatesActionsApi.GroupSettingTemplatesGetAvailableExtensionProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupSettingTemplatesGetAvailableExtensionProperties`: []MicrosoftGraphExtensionProperty
    fmt.Fprintf(os.Stdout, "Response from `GroupSettingTemplatesActionsApi.GroupSettingTemplatesGetAvailableExtensionProperties`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGroupSettingTemplatesGetAvailableExtensionPropertiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject353** | [**InlineObject353**](InlineObject353.md) |  | 

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


## GroupSettingTemplatesGetByIds

> []MicrosoftGraphDirectoryObject GroupSettingTemplatesGetByIds(ctx).InlineObject354(inlineObject354).Execute()

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
    inlineObject354 := *openapiclient.NewInlineObject354() // InlineObject354 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupSettingTemplatesActionsApi.GroupSettingTemplatesGetByIds(context.Background()).InlineObject354(inlineObject354).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupSettingTemplatesActionsApi.GroupSettingTemplatesGetByIds``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupSettingTemplatesGetByIds`: []MicrosoftGraphDirectoryObject
    fmt.Fprintf(os.Stdout, "Response from `GroupSettingTemplatesActionsApi.GroupSettingTemplatesGetByIds`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGroupSettingTemplatesGetByIdsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject354** | [**InlineObject354**](InlineObject354.md) |  | 

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


## GroupSettingTemplatesGroupSettingTemplateCheckMemberGroups

> []string GroupSettingTemplatesGroupSettingTemplateCheckMemberGroups(ctx, groupSettingTemplateId).InlineObject349(inlineObject349).Execute()

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
    groupSettingTemplateId := "groupSettingTemplateId_example" // string | key: id of groupSettingTemplate
    inlineObject349 := *openapiclient.NewInlineObject349() // InlineObject349 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupSettingTemplatesActionsApi.GroupSettingTemplatesGroupSettingTemplateCheckMemberGroups(context.Background(), groupSettingTemplateId).InlineObject349(inlineObject349).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupSettingTemplatesActionsApi.GroupSettingTemplatesGroupSettingTemplateCheckMemberGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupSettingTemplatesGroupSettingTemplateCheckMemberGroups`: []string
    fmt.Fprintf(os.Stdout, "Response from `GroupSettingTemplatesActionsApi.GroupSettingTemplatesGroupSettingTemplateCheckMemberGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupSettingTemplateId** | **string** | key: id of groupSettingTemplate | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupSettingTemplatesGroupSettingTemplateCheckMemberGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject349** | [**InlineObject349**](InlineObject349.md) |  | 

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


## GroupSettingTemplatesGroupSettingTemplateCheckMemberObjects

> []string GroupSettingTemplatesGroupSettingTemplateCheckMemberObjects(ctx, groupSettingTemplateId).InlineObject350(inlineObject350).Execute()

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
    groupSettingTemplateId := "groupSettingTemplateId_example" // string | key: id of groupSettingTemplate
    inlineObject350 := *openapiclient.NewInlineObject350() // InlineObject350 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupSettingTemplatesActionsApi.GroupSettingTemplatesGroupSettingTemplateCheckMemberObjects(context.Background(), groupSettingTemplateId).InlineObject350(inlineObject350).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupSettingTemplatesActionsApi.GroupSettingTemplatesGroupSettingTemplateCheckMemberObjects``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupSettingTemplatesGroupSettingTemplateCheckMemberObjects`: []string
    fmt.Fprintf(os.Stdout, "Response from `GroupSettingTemplatesActionsApi.GroupSettingTemplatesGroupSettingTemplateCheckMemberObjects`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupSettingTemplateId** | **string** | key: id of groupSettingTemplate | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupSettingTemplatesGroupSettingTemplateCheckMemberObjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject350** | [**InlineObject350**](InlineObject350.md) |  | 

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


## GroupSettingTemplatesGroupSettingTemplateGetMemberGroups

> []string GroupSettingTemplatesGroupSettingTemplateGetMemberGroups(ctx, groupSettingTemplateId).InlineObject351(inlineObject351).Execute()

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
    groupSettingTemplateId := "groupSettingTemplateId_example" // string | key: id of groupSettingTemplate
    inlineObject351 := *openapiclient.NewInlineObject351() // InlineObject351 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupSettingTemplatesActionsApi.GroupSettingTemplatesGroupSettingTemplateGetMemberGroups(context.Background(), groupSettingTemplateId).InlineObject351(inlineObject351).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupSettingTemplatesActionsApi.GroupSettingTemplatesGroupSettingTemplateGetMemberGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupSettingTemplatesGroupSettingTemplateGetMemberGroups`: []string
    fmt.Fprintf(os.Stdout, "Response from `GroupSettingTemplatesActionsApi.GroupSettingTemplatesGroupSettingTemplateGetMemberGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupSettingTemplateId** | **string** | key: id of groupSettingTemplate | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupSettingTemplatesGroupSettingTemplateGetMemberGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject351** | [**InlineObject351**](InlineObject351.md) |  | 

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


## GroupSettingTemplatesGroupSettingTemplateGetMemberObjects

> []string GroupSettingTemplatesGroupSettingTemplateGetMemberObjects(ctx, groupSettingTemplateId).InlineObject352(inlineObject352).Execute()

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
    groupSettingTemplateId := "groupSettingTemplateId_example" // string | key: id of groupSettingTemplate
    inlineObject352 := *openapiclient.NewInlineObject352() // InlineObject352 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupSettingTemplatesActionsApi.GroupSettingTemplatesGroupSettingTemplateGetMemberObjects(context.Background(), groupSettingTemplateId).InlineObject352(inlineObject352).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupSettingTemplatesActionsApi.GroupSettingTemplatesGroupSettingTemplateGetMemberObjects``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupSettingTemplatesGroupSettingTemplateGetMemberObjects`: []string
    fmt.Fprintf(os.Stdout, "Response from `GroupSettingTemplatesActionsApi.GroupSettingTemplatesGroupSettingTemplateGetMemberObjects`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupSettingTemplateId** | **string** | key: id of groupSettingTemplate | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupSettingTemplatesGroupSettingTemplateGetMemberObjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject352** | [**InlineObject352**](InlineObject352.md) |  | 

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


## GroupSettingTemplatesGroupSettingTemplateRestore

> AnyOfmicrosoftGraphDirectoryObject GroupSettingTemplatesGroupSettingTemplateRestore(ctx, groupSettingTemplateId).Execute()

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
    groupSettingTemplateId := "groupSettingTemplateId_example" // string | key: id of groupSettingTemplate

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupSettingTemplatesActionsApi.GroupSettingTemplatesGroupSettingTemplateRestore(context.Background(), groupSettingTemplateId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupSettingTemplatesActionsApi.GroupSettingTemplatesGroupSettingTemplateRestore``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupSettingTemplatesGroupSettingTemplateRestore`: AnyOfmicrosoftGraphDirectoryObject
    fmt.Fprintf(os.Stdout, "Response from `GroupSettingTemplatesActionsApi.GroupSettingTemplatesGroupSettingTemplateRestore`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupSettingTemplateId** | **string** | key: id of groupSettingTemplate | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupSettingTemplatesGroupSettingTemplateRestoreRequest struct via the builder pattern


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


## GroupSettingTemplatesValidateProperties

> GroupSettingTemplatesValidateProperties(ctx).InlineObject355(inlineObject355).Execute()

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
    inlineObject355 := *openapiclient.NewInlineObject355() // InlineObject355 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupSettingTemplatesActionsApi.GroupSettingTemplatesValidateProperties(context.Background()).InlineObject355(inlineObject355).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupSettingTemplatesActionsApi.GroupSettingTemplatesValidateProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGroupSettingTemplatesValidatePropertiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject355** | [**InlineObject355**](InlineObject355.md) |  | 

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

