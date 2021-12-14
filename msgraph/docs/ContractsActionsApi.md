# \ContractsActionsApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContractsContractCheckMemberGroups**](ContractsActionsApi.md#ContractsContractCheckMemberGroups) | **Post** /contracts/{contract-id}/microsoft.graph.checkMemberGroups | Invoke action checkMemberGroups
[**ContractsContractCheckMemberObjects**](ContractsActionsApi.md#ContractsContractCheckMemberObjects) | **Post** /contracts/{contract-id}/microsoft.graph.checkMemberObjects | Invoke action checkMemberObjects
[**ContractsContractGetMemberGroups**](ContractsActionsApi.md#ContractsContractGetMemberGroups) | **Post** /contracts/{contract-id}/microsoft.graph.getMemberGroups | Invoke action getMemberGroups
[**ContractsContractGetMemberObjects**](ContractsActionsApi.md#ContractsContractGetMemberObjects) | **Post** /contracts/{contract-id}/microsoft.graph.getMemberObjects | Invoke action getMemberObjects
[**ContractsContractRestore**](ContractsActionsApi.md#ContractsContractRestore) | **Post** /contracts/{contract-id}/microsoft.graph.restore | Invoke action restore
[**ContractsGetAvailableExtensionProperties**](ContractsActionsApi.md#ContractsGetAvailableExtensionProperties) | **Post** /contracts/microsoft.graph.getAvailableExtensionProperties | Invoke action getAvailableExtensionProperties
[**ContractsGetByIds**](ContractsActionsApi.md#ContractsGetByIds) | **Post** /contracts/microsoft.graph.getByIds | Invoke action getByIds
[**ContractsValidateProperties**](ContractsActionsApi.md#ContractsValidateProperties) | **Post** /contracts/microsoft.graph.validateProperties | Invoke action validateProperties



## ContractsContractCheckMemberGroups

> []string ContractsContractCheckMemberGroups(ctx, contractId).InlineObject49(inlineObject49).Execute()

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
    contractId := "contractId_example" // string | key: id of contract
    inlineObject49 := *openapiclient.NewInlineObject49() // InlineObject49 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContractsActionsApi.ContractsContractCheckMemberGroups(context.Background(), contractId).InlineObject49(inlineObject49).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContractsActionsApi.ContractsContractCheckMemberGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContractsContractCheckMemberGroups`: []string
    fmt.Fprintf(os.Stdout, "Response from `ContractsActionsApi.ContractsContractCheckMemberGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contractId** | **string** | key: id of contract | 

### Other Parameters

Other parameters are passed through a pointer to a apiContractsContractCheckMemberGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject49** | [**InlineObject49**](InlineObject49.md) |  | 

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


## ContractsContractCheckMemberObjects

> []string ContractsContractCheckMemberObjects(ctx, contractId).InlineObject50(inlineObject50).Execute()

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
    contractId := "contractId_example" // string | key: id of contract
    inlineObject50 := *openapiclient.NewInlineObject50() // InlineObject50 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContractsActionsApi.ContractsContractCheckMemberObjects(context.Background(), contractId).InlineObject50(inlineObject50).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContractsActionsApi.ContractsContractCheckMemberObjects``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContractsContractCheckMemberObjects`: []string
    fmt.Fprintf(os.Stdout, "Response from `ContractsActionsApi.ContractsContractCheckMemberObjects`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contractId** | **string** | key: id of contract | 

### Other Parameters

Other parameters are passed through a pointer to a apiContractsContractCheckMemberObjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject50** | [**InlineObject50**](InlineObject50.md) |  | 

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


## ContractsContractGetMemberGroups

> []string ContractsContractGetMemberGroups(ctx, contractId).InlineObject51(inlineObject51).Execute()

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
    contractId := "contractId_example" // string | key: id of contract
    inlineObject51 := *openapiclient.NewInlineObject51() // InlineObject51 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContractsActionsApi.ContractsContractGetMemberGroups(context.Background(), contractId).InlineObject51(inlineObject51).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContractsActionsApi.ContractsContractGetMemberGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContractsContractGetMemberGroups`: []string
    fmt.Fprintf(os.Stdout, "Response from `ContractsActionsApi.ContractsContractGetMemberGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contractId** | **string** | key: id of contract | 

### Other Parameters

Other parameters are passed through a pointer to a apiContractsContractGetMemberGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject51** | [**InlineObject51**](InlineObject51.md) |  | 

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


## ContractsContractGetMemberObjects

> []string ContractsContractGetMemberObjects(ctx, contractId).InlineObject52(inlineObject52).Execute()

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
    contractId := "contractId_example" // string | key: id of contract
    inlineObject52 := *openapiclient.NewInlineObject52() // InlineObject52 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContractsActionsApi.ContractsContractGetMemberObjects(context.Background(), contractId).InlineObject52(inlineObject52).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContractsActionsApi.ContractsContractGetMemberObjects``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContractsContractGetMemberObjects`: []string
    fmt.Fprintf(os.Stdout, "Response from `ContractsActionsApi.ContractsContractGetMemberObjects`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contractId** | **string** | key: id of contract | 

### Other Parameters

Other parameters are passed through a pointer to a apiContractsContractGetMemberObjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject52** | [**InlineObject52**](InlineObject52.md) |  | 

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


## ContractsContractRestore

> AnyOfmicrosoftGraphDirectoryObject ContractsContractRestore(ctx, contractId).Execute()

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
    contractId := "contractId_example" // string | key: id of contract

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContractsActionsApi.ContractsContractRestore(context.Background(), contractId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContractsActionsApi.ContractsContractRestore``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContractsContractRestore`: AnyOfmicrosoftGraphDirectoryObject
    fmt.Fprintf(os.Stdout, "Response from `ContractsActionsApi.ContractsContractRestore`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contractId** | **string** | key: id of contract | 

### Other Parameters

Other parameters are passed through a pointer to a apiContractsContractRestoreRequest struct via the builder pattern


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


## ContractsGetAvailableExtensionProperties

> []MicrosoftGraphExtensionProperty ContractsGetAvailableExtensionProperties(ctx).InlineObject53(inlineObject53).Execute()

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
    inlineObject53 := *openapiclient.NewInlineObject53() // InlineObject53 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContractsActionsApi.ContractsGetAvailableExtensionProperties(context.Background()).InlineObject53(inlineObject53).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContractsActionsApi.ContractsGetAvailableExtensionProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContractsGetAvailableExtensionProperties`: []MicrosoftGraphExtensionProperty
    fmt.Fprintf(os.Stdout, "Response from `ContractsActionsApi.ContractsGetAvailableExtensionProperties`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContractsGetAvailableExtensionPropertiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject53** | [**InlineObject53**](InlineObject53.md) |  | 

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


## ContractsGetByIds

> []MicrosoftGraphDirectoryObject ContractsGetByIds(ctx).InlineObject54(inlineObject54).Execute()

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
    inlineObject54 := *openapiclient.NewInlineObject54() // InlineObject54 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContractsActionsApi.ContractsGetByIds(context.Background()).InlineObject54(inlineObject54).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContractsActionsApi.ContractsGetByIds``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContractsGetByIds`: []MicrosoftGraphDirectoryObject
    fmt.Fprintf(os.Stdout, "Response from `ContractsActionsApi.ContractsGetByIds`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContractsGetByIdsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject54** | [**InlineObject54**](InlineObject54.md) |  | 

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


## ContractsValidateProperties

> ContractsValidateProperties(ctx).InlineObject55(inlineObject55).Execute()

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
    inlineObject55 := *openapiclient.NewInlineObject55() // InlineObject55 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContractsActionsApi.ContractsValidateProperties(context.Background()).InlineObject55(inlineObject55).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContractsActionsApi.ContractsValidateProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContractsValidatePropertiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject55** | [**InlineObject55**](InlineObject55.md) |  | 

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

