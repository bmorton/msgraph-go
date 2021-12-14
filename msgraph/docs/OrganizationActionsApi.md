# \OrganizationActionsApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrganizationGetAvailableExtensionProperties**](OrganizationActionsApi.md#OrganizationGetAvailableExtensionProperties) | **Post** /organization/microsoft.graph.getAvailableExtensionProperties | Invoke action getAvailableExtensionProperties
[**OrganizationGetByIds**](OrganizationActionsApi.md#OrganizationGetByIds) | **Post** /organization/microsoft.graph.getByIds | Invoke action getByIds
[**OrganizationOrganizationCheckMemberGroups**](OrganizationActionsApi.md#OrganizationOrganizationCheckMemberGroups) | **Post** /organization/{organization-id}/microsoft.graph.checkMemberGroups | Invoke action checkMemberGroups
[**OrganizationOrganizationCheckMemberObjects**](OrganizationActionsApi.md#OrganizationOrganizationCheckMemberObjects) | **Post** /organization/{organization-id}/microsoft.graph.checkMemberObjects | Invoke action checkMemberObjects
[**OrganizationOrganizationGetMemberGroups**](OrganizationActionsApi.md#OrganizationOrganizationGetMemberGroups) | **Post** /organization/{organization-id}/microsoft.graph.getMemberGroups | Invoke action getMemberGroups
[**OrganizationOrganizationGetMemberObjects**](OrganizationActionsApi.md#OrganizationOrganizationGetMemberObjects) | **Post** /organization/{organization-id}/microsoft.graph.getMemberObjects | Invoke action getMemberObjects
[**OrganizationOrganizationRestore**](OrganizationActionsApi.md#OrganizationOrganizationRestore) | **Post** /organization/{organization-id}/microsoft.graph.restore | Invoke action restore
[**OrganizationOrganizationSetMobileDeviceManagementAuthority**](OrganizationActionsApi.md#OrganizationOrganizationSetMobileDeviceManagementAuthority) | **Post** /organization/{organization-id}/microsoft.graph.setMobileDeviceManagementAuthority | Invoke action setMobileDeviceManagementAuthority
[**OrganizationValidateProperties**](OrganizationActionsApi.md#OrganizationValidateProperties) | **Post** /organization/microsoft.graph.validateProperties | Invoke action validateProperties



## OrganizationGetAvailableExtensionProperties

> []MicrosoftGraphExtensionProperty OrganizationGetAvailableExtensionProperties(ctx).InlineObject703(inlineObject703).Execute()

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
    inlineObject703 := *openapiclient.NewInlineObject703() // InlineObject703 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganizationActionsApi.OrganizationGetAvailableExtensionProperties(context.Background()).InlineObject703(inlineObject703).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationActionsApi.OrganizationGetAvailableExtensionProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganizationGetAvailableExtensionProperties`: []MicrosoftGraphExtensionProperty
    fmt.Fprintf(os.Stdout, "Response from `OrganizationActionsApi.OrganizationGetAvailableExtensionProperties`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationGetAvailableExtensionPropertiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject703** | [**InlineObject703**](InlineObject703.md) |  | 

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


## OrganizationGetByIds

> []MicrosoftGraphDirectoryObject OrganizationGetByIds(ctx).InlineObject704(inlineObject704).Execute()

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
    inlineObject704 := *openapiclient.NewInlineObject704() // InlineObject704 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganizationActionsApi.OrganizationGetByIds(context.Background()).InlineObject704(inlineObject704).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationActionsApi.OrganizationGetByIds``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganizationGetByIds`: []MicrosoftGraphDirectoryObject
    fmt.Fprintf(os.Stdout, "Response from `OrganizationActionsApi.OrganizationGetByIds`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationGetByIdsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject704** | [**InlineObject704**](InlineObject704.md) |  | 

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


## OrganizationOrganizationCheckMemberGroups

> []string OrganizationOrganizationCheckMemberGroups(ctx, organizationId).InlineObject699(inlineObject699).Execute()

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
    organizationId := "organizationId_example" // string | key: id of organization
    inlineObject699 := *openapiclient.NewInlineObject699() // InlineObject699 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganizationActionsApi.OrganizationOrganizationCheckMemberGroups(context.Background(), organizationId).InlineObject699(inlineObject699).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationActionsApi.OrganizationOrganizationCheckMemberGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganizationOrganizationCheckMemberGroups`: []string
    fmt.Fprintf(os.Stdout, "Response from `OrganizationActionsApi.OrganizationOrganizationCheckMemberGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | key: id of organization | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationOrganizationCheckMemberGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject699** | [**InlineObject699**](InlineObject699.md) |  | 

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


## OrganizationOrganizationCheckMemberObjects

> []string OrganizationOrganizationCheckMemberObjects(ctx, organizationId).InlineObject700(inlineObject700).Execute()

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
    organizationId := "organizationId_example" // string | key: id of organization
    inlineObject700 := *openapiclient.NewInlineObject700() // InlineObject700 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganizationActionsApi.OrganizationOrganizationCheckMemberObjects(context.Background(), organizationId).InlineObject700(inlineObject700).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationActionsApi.OrganizationOrganizationCheckMemberObjects``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganizationOrganizationCheckMemberObjects`: []string
    fmt.Fprintf(os.Stdout, "Response from `OrganizationActionsApi.OrganizationOrganizationCheckMemberObjects`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | key: id of organization | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationOrganizationCheckMemberObjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject700** | [**InlineObject700**](InlineObject700.md) |  | 

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


## OrganizationOrganizationGetMemberGroups

> []string OrganizationOrganizationGetMemberGroups(ctx, organizationId).InlineObject701(inlineObject701).Execute()

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
    organizationId := "organizationId_example" // string | key: id of organization
    inlineObject701 := *openapiclient.NewInlineObject701() // InlineObject701 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganizationActionsApi.OrganizationOrganizationGetMemberGroups(context.Background(), organizationId).InlineObject701(inlineObject701).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationActionsApi.OrganizationOrganizationGetMemberGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganizationOrganizationGetMemberGroups`: []string
    fmt.Fprintf(os.Stdout, "Response from `OrganizationActionsApi.OrganizationOrganizationGetMemberGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | key: id of organization | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationOrganizationGetMemberGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject701** | [**InlineObject701**](InlineObject701.md) |  | 

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


## OrganizationOrganizationGetMemberObjects

> []string OrganizationOrganizationGetMemberObjects(ctx, organizationId).InlineObject702(inlineObject702).Execute()

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
    organizationId := "organizationId_example" // string | key: id of organization
    inlineObject702 := *openapiclient.NewInlineObject702() // InlineObject702 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganizationActionsApi.OrganizationOrganizationGetMemberObjects(context.Background(), organizationId).InlineObject702(inlineObject702).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationActionsApi.OrganizationOrganizationGetMemberObjects``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganizationOrganizationGetMemberObjects`: []string
    fmt.Fprintf(os.Stdout, "Response from `OrganizationActionsApi.OrganizationOrganizationGetMemberObjects`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | key: id of organization | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationOrganizationGetMemberObjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject702** | [**InlineObject702**](InlineObject702.md) |  | 

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


## OrganizationOrganizationRestore

> AnyOfmicrosoftGraphDirectoryObject OrganizationOrganizationRestore(ctx, organizationId).Execute()

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
    organizationId := "organizationId_example" // string | key: id of organization

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganizationActionsApi.OrganizationOrganizationRestore(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationActionsApi.OrganizationOrganizationRestore``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganizationOrganizationRestore`: AnyOfmicrosoftGraphDirectoryObject
    fmt.Fprintf(os.Stdout, "Response from `OrganizationActionsApi.OrganizationOrganizationRestore`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | key: id of organization | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationOrganizationRestoreRequest struct via the builder pattern


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


## OrganizationOrganizationSetMobileDeviceManagementAuthority

> int32 OrganizationOrganizationSetMobileDeviceManagementAuthority(ctx, organizationId).Execute()

Invoke action setMobileDeviceManagementAuthority



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
    organizationId := "organizationId_example" // string | key: id of organization

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganizationActionsApi.OrganizationOrganizationSetMobileDeviceManagementAuthority(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationActionsApi.OrganizationOrganizationSetMobileDeviceManagementAuthority``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganizationOrganizationSetMobileDeviceManagementAuthority`: int32
    fmt.Fprintf(os.Stdout, "Response from `OrganizationActionsApi.OrganizationOrganizationSetMobileDeviceManagementAuthority`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | key: id of organization | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationOrganizationSetMobileDeviceManagementAuthorityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**int32**

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationValidateProperties

> OrganizationValidateProperties(ctx).InlineObject705(inlineObject705).Execute()

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
    inlineObject705 := *openapiclient.NewInlineObject705() // InlineObject705 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganizationActionsApi.OrganizationValidateProperties(context.Background()).InlineObject705(inlineObject705).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationActionsApi.OrganizationValidateProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationValidatePropertiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject705** | [**InlineObject705**](InlineObject705.md) |  | 

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

