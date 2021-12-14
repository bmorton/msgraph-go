# \ServicePrincipalsActionsApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ServicePrincipalsGetAvailableExtensionProperties**](ServicePrincipalsActionsApi.md#ServicePrincipalsGetAvailableExtensionProperties) | **Post** /servicePrincipals/microsoft.graph.getAvailableExtensionProperties | Invoke action getAvailableExtensionProperties
[**ServicePrincipalsGetByIds**](ServicePrincipalsActionsApi.md#ServicePrincipalsGetByIds) | **Post** /servicePrincipals/microsoft.graph.getByIds | Invoke action getByIds
[**ServicePrincipalsServicePrincipalAddKey**](ServicePrincipalsActionsApi.md#ServicePrincipalsServicePrincipalAddKey) | **Post** /servicePrincipals/{servicePrincipal-id}/microsoft.graph.addKey | Invoke action addKey
[**ServicePrincipalsServicePrincipalAddPassword**](ServicePrincipalsActionsApi.md#ServicePrincipalsServicePrincipalAddPassword) | **Post** /servicePrincipals/{servicePrincipal-id}/microsoft.graph.addPassword | Invoke action addPassword
[**ServicePrincipalsServicePrincipalCheckMemberGroups**](ServicePrincipalsActionsApi.md#ServicePrincipalsServicePrincipalCheckMemberGroups) | **Post** /servicePrincipals/{servicePrincipal-id}/microsoft.graph.checkMemberGroups | Invoke action checkMemberGroups
[**ServicePrincipalsServicePrincipalCheckMemberObjects**](ServicePrincipalsActionsApi.md#ServicePrincipalsServicePrincipalCheckMemberObjects) | **Post** /servicePrincipals/{servicePrincipal-id}/microsoft.graph.checkMemberObjects | Invoke action checkMemberObjects
[**ServicePrincipalsServicePrincipalGetMemberGroups**](ServicePrincipalsActionsApi.md#ServicePrincipalsServicePrincipalGetMemberGroups) | **Post** /servicePrincipals/{servicePrincipal-id}/microsoft.graph.getMemberGroups | Invoke action getMemberGroups
[**ServicePrincipalsServicePrincipalGetMemberObjects**](ServicePrincipalsActionsApi.md#ServicePrincipalsServicePrincipalGetMemberObjects) | **Post** /servicePrincipals/{servicePrincipal-id}/microsoft.graph.getMemberObjects | Invoke action getMemberObjects
[**ServicePrincipalsServicePrincipalRemoveKey**](ServicePrincipalsActionsApi.md#ServicePrincipalsServicePrincipalRemoveKey) | **Post** /servicePrincipals/{servicePrincipal-id}/microsoft.graph.removeKey | Invoke action removeKey
[**ServicePrincipalsServicePrincipalRemovePassword**](ServicePrincipalsActionsApi.md#ServicePrincipalsServicePrincipalRemovePassword) | **Post** /servicePrincipals/{servicePrincipal-id}/microsoft.graph.removePassword | Invoke action removePassword
[**ServicePrincipalsServicePrincipalRestore**](ServicePrincipalsActionsApi.md#ServicePrincipalsServicePrincipalRestore) | **Post** /servicePrincipals/{servicePrincipal-id}/microsoft.graph.restore | Invoke action restore
[**ServicePrincipalsValidateProperties**](ServicePrincipalsActionsApi.md#ServicePrincipalsValidateProperties) | **Post** /servicePrincipals/microsoft.graph.validateProperties | Invoke action validateProperties



## ServicePrincipalsGetAvailableExtensionProperties

> []MicrosoftGraphExtensionProperty ServicePrincipalsGetAvailableExtensionProperties(ctx).InlineObject723(inlineObject723).Execute()

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
    inlineObject723 := *openapiclient.NewInlineObject723() // InlineObject723 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServicePrincipalsActionsApi.ServicePrincipalsGetAvailableExtensionProperties(context.Background()).InlineObject723(inlineObject723).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicePrincipalsActionsApi.ServicePrincipalsGetAvailableExtensionProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServicePrincipalsGetAvailableExtensionProperties`: []MicrosoftGraphExtensionProperty
    fmt.Fprintf(os.Stdout, "Response from `ServicePrincipalsActionsApi.ServicePrincipalsGetAvailableExtensionProperties`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiServicePrincipalsGetAvailableExtensionPropertiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject723** | [**InlineObject723**](InlineObject723.md) |  | 

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


## ServicePrincipalsGetByIds

> []MicrosoftGraphDirectoryObject ServicePrincipalsGetByIds(ctx).InlineObject724(inlineObject724).Execute()

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
    inlineObject724 := *openapiclient.NewInlineObject724() // InlineObject724 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServicePrincipalsActionsApi.ServicePrincipalsGetByIds(context.Background()).InlineObject724(inlineObject724).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicePrincipalsActionsApi.ServicePrincipalsGetByIds``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServicePrincipalsGetByIds`: []MicrosoftGraphDirectoryObject
    fmt.Fprintf(os.Stdout, "Response from `ServicePrincipalsActionsApi.ServicePrincipalsGetByIds`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiServicePrincipalsGetByIdsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject724** | [**InlineObject724**](InlineObject724.md) |  | 

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


## ServicePrincipalsServicePrincipalAddKey

> MicrosoftGraphKeyCredential ServicePrincipalsServicePrincipalAddKey(ctx, servicePrincipalId).InlineObject715(inlineObject715).Execute()

Invoke action addKey

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
    servicePrincipalId := "servicePrincipalId_example" // string | key: id of servicePrincipal
    inlineObject715 := *openapiclient.NewInlineObject715() // InlineObject715 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServicePrincipalsActionsApi.ServicePrincipalsServicePrincipalAddKey(context.Background(), servicePrincipalId).InlineObject715(inlineObject715).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicePrincipalsActionsApi.ServicePrincipalsServicePrincipalAddKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServicePrincipalsServicePrincipalAddKey`: MicrosoftGraphKeyCredential
    fmt.Fprintf(os.Stdout, "Response from `ServicePrincipalsActionsApi.ServicePrincipalsServicePrincipalAddKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**servicePrincipalId** | **string** | key: id of servicePrincipal | 

### Other Parameters

Other parameters are passed through a pointer to a apiServicePrincipalsServicePrincipalAddKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject715** | [**InlineObject715**](InlineObject715.md) |  | 

### Return type

[**MicrosoftGraphKeyCredential**](MicrosoftGraphKeyCredential.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServicePrincipalsServicePrincipalAddPassword

> MicrosoftGraphPasswordCredential ServicePrincipalsServicePrincipalAddPassword(ctx, servicePrincipalId).InlineObject716(inlineObject716).Execute()

Invoke action addPassword

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
    servicePrincipalId := "servicePrincipalId_example" // string | key: id of servicePrincipal
    inlineObject716 := *openapiclient.NewInlineObject716() // InlineObject716 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServicePrincipalsActionsApi.ServicePrincipalsServicePrincipalAddPassword(context.Background(), servicePrincipalId).InlineObject716(inlineObject716).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicePrincipalsActionsApi.ServicePrincipalsServicePrincipalAddPassword``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServicePrincipalsServicePrincipalAddPassword`: MicrosoftGraphPasswordCredential
    fmt.Fprintf(os.Stdout, "Response from `ServicePrincipalsActionsApi.ServicePrincipalsServicePrincipalAddPassword`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**servicePrincipalId** | **string** | key: id of servicePrincipal | 

### Other Parameters

Other parameters are passed through a pointer to a apiServicePrincipalsServicePrincipalAddPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject716** | [**InlineObject716**](InlineObject716.md) |  | 

### Return type

[**MicrosoftGraphPasswordCredential**](MicrosoftGraphPasswordCredential.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServicePrincipalsServicePrincipalCheckMemberGroups

> []string ServicePrincipalsServicePrincipalCheckMemberGroups(ctx, servicePrincipalId).InlineObject717(inlineObject717).Execute()

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
    servicePrincipalId := "servicePrincipalId_example" // string | key: id of servicePrincipal
    inlineObject717 := *openapiclient.NewInlineObject717() // InlineObject717 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServicePrincipalsActionsApi.ServicePrincipalsServicePrincipalCheckMemberGroups(context.Background(), servicePrincipalId).InlineObject717(inlineObject717).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicePrincipalsActionsApi.ServicePrincipalsServicePrincipalCheckMemberGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServicePrincipalsServicePrincipalCheckMemberGroups`: []string
    fmt.Fprintf(os.Stdout, "Response from `ServicePrincipalsActionsApi.ServicePrincipalsServicePrincipalCheckMemberGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**servicePrincipalId** | **string** | key: id of servicePrincipal | 

### Other Parameters

Other parameters are passed through a pointer to a apiServicePrincipalsServicePrincipalCheckMemberGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject717** | [**InlineObject717**](InlineObject717.md) |  | 

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


## ServicePrincipalsServicePrincipalCheckMemberObjects

> []string ServicePrincipalsServicePrincipalCheckMemberObjects(ctx, servicePrincipalId).InlineObject718(inlineObject718).Execute()

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
    servicePrincipalId := "servicePrincipalId_example" // string | key: id of servicePrincipal
    inlineObject718 := *openapiclient.NewInlineObject718() // InlineObject718 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServicePrincipalsActionsApi.ServicePrincipalsServicePrincipalCheckMemberObjects(context.Background(), servicePrincipalId).InlineObject718(inlineObject718).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicePrincipalsActionsApi.ServicePrincipalsServicePrincipalCheckMemberObjects``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServicePrincipalsServicePrincipalCheckMemberObjects`: []string
    fmt.Fprintf(os.Stdout, "Response from `ServicePrincipalsActionsApi.ServicePrincipalsServicePrincipalCheckMemberObjects`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**servicePrincipalId** | **string** | key: id of servicePrincipal | 

### Other Parameters

Other parameters are passed through a pointer to a apiServicePrincipalsServicePrincipalCheckMemberObjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject718** | [**InlineObject718**](InlineObject718.md) |  | 

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


## ServicePrincipalsServicePrincipalGetMemberGroups

> []string ServicePrincipalsServicePrincipalGetMemberGroups(ctx, servicePrincipalId).InlineObject719(inlineObject719).Execute()

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
    servicePrincipalId := "servicePrincipalId_example" // string | key: id of servicePrincipal
    inlineObject719 := *openapiclient.NewInlineObject719() // InlineObject719 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServicePrincipalsActionsApi.ServicePrincipalsServicePrincipalGetMemberGroups(context.Background(), servicePrincipalId).InlineObject719(inlineObject719).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicePrincipalsActionsApi.ServicePrincipalsServicePrincipalGetMemberGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServicePrincipalsServicePrincipalGetMemberGroups`: []string
    fmt.Fprintf(os.Stdout, "Response from `ServicePrincipalsActionsApi.ServicePrincipalsServicePrincipalGetMemberGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**servicePrincipalId** | **string** | key: id of servicePrincipal | 

### Other Parameters

Other parameters are passed through a pointer to a apiServicePrincipalsServicePrincipalGetMemberGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject719** | [**InlineObject719**](InlineObject719.md) |  | 

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


## ServicePrincipalsServicePrincipalGetMemberObjects

> []string ServicePrincipalsServicePrincipalGetMemberObjects(ctx, servicePrincipalId).InlineObject720(inlineObject720).Execute()

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
    servicePrincipalId := "servicePrincipalId_example" // string | key: id of servicePrincipal
    inlineObject720 := *openapiclient.NewInlineObject720() // InlineObject720 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServicePrincipalsActionsApi.ServicePrincipalsServicePrincipalGetMemberObjects(context.Background(), servicePrincipalId).InlineObject720(inlineObject720).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicePrincipalsActionsApi.ServicePrincipalsServicePrincipalGetMemberObjects``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServicePrincipalsServicePrincipalGetMemberObjects`: []string
    fmt.Fprintf(os.Stdout, "Response from `ServicePrincipalsActionsApi.ServicePrincipalsServicePrincipalGetMemberObjects`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**servicePrincipalId** | **string** | key: id of servicePrincipal | 

### Other Parameters

Other parameters are passed through a pointer to a apiServicePrincipalsServicePrincipalGetMemberObjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject720** | [**InlineObject720**](InlineObject720.md) |  | 

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


## ServicePrincipalsServicePrincipalRemoveKey

> ServicePrincipalsServicePrincipalRemoveKey(ctx, servicePrincipalId).InlineObject721(inlineObject721).Execute()

Invoke action removeKey

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
    servicePrincipalId := "servicePrincipalId_example" // string | key: id of servicePrincipal
    inlineObject721 := *openapiclient.NewInlineObject721() // InlineObject721 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServicePrincipalsActionsApi.ServicePrincipalsServicePrincipalRemoveKey(context.Background(), servicePrincipalId).InlineObject721(inlineObject721).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicePrincipalsActionsApi.ServicePrincipalsServicePrincipalRemoveKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**servicePrincipalId** | **string** | key: id of servicePrincipal | 

### Other Parameters

Other parameters are passed through a pointer to a apiServicePrincipalsServicePrincipalRemoveKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject721** | [**InlineObject721**](InlineObject721.md) |  | 

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


## ServicePrincipalsServicePrincipalRemovePassword

> ServicePrincipalsServicePrincipalRemovePassword(ctx, servicePrincipalId).InlineObject722(inlineObject722).Execute()

Invoke action removePassword

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
    servicePrincipalId := "servicePrincipalId_example" // string | key: id of servicePrincipal
    inlineObject722 := *openapiclient.NewInlineObject722() // InlineObject722 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServicePrincipalsActionsApi.ServicePrincipalsServicePrincipalRemovePassword(context.Background(), servicePrincipalId).InlineObject722(inlineObject722).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicePrincipalsActionsApi.ServicePrincipalsServicePrincipalRemovePassword``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**servicePrincipalId** | **string** | key: id of servicePrincipal | 

### Other Parameters

Other parameters are passed through a pointer to a apiServicePrincipalsServicePrincipalRemovePasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject722** | [**InlineObject722**](InlineObject722.md) |  | 

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


## ServicePrincipalsServicePrincipalRestore

> AnyOfmicrosoftGraphDirectoryObject ServicePrincipalsServicePrincipalRestore(ctx, servicePrincipalId).Execute()

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
    servicePrincipalId := "servicePrincipalId_example" // string | key: id of servicePrincipal

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServicePrincipalsActionsApi.ServicePrincipalsServicePrincipalRestore(context.Background(), servicePrincipalId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicePrincipalsActionsApi.ServicePrincipalsServicePrincipalRestore``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServicePrincipalsServicePrincipalRestore`: AnyOfmicrosoftGraphDirectoryObject
    fmt.Fprintf(os.Stdout, "Response from `ServicePrincipalsActionsApi.ServicePrincipalsServicePrincipalRestore`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**servicePrincipalId** | **string** | key: id of servicePrincipal | 

### Other Parameters

Other parameters are passed through a pointer to a apiServicePrincipalsServicePrincipalRestoreRequest struct via the builder pattern


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


## ServicePrincipalsValidateProperties

> ServicePrincipalsValidateProperties(ctx).InlineObject725(inlineObject725).Execute()

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
    inlineObject725 := *openapiclient.NewInlineObject725() // InlineObject725 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServicePrincipalsActionsApi.ServicePrincipalsValidateProperties(context.Background()).InlineObject725(inlineObject725).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicePrincipalsActionsApi.ServicePrincipalsValidateProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiServicePrincipalsValidatePropertiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject725** | [**InlineObject725**](InlineObject725.md) |  | 

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

