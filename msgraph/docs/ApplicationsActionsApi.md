# \ApplicationsActionsApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApplicationsApplicationAddKey**](ApplicationsActionsApi.md#ApplicationsApplicationAddKey) | **Post** /applications/{application-id}/microsoft.graph.addKey | Invoke action addKey
[**ApplicationsApplicationAddPassword**](ApplicationsActionsApi.md#ApplicationsApplicationAddPassword) | **Post** /applications/{application-id}/microsoft.graph.addPassword | Invoke action addPassword
[**ApplicationsApplicationCheckMemberGroups**](ApplicationsActionsApi.md#ApplicationsApplicationCheckMemberGroups) | **Post** /applications/{application-id}/microsoft.graph.checkMemberGroups | Invoke action checkMemberGroups
[**ApplicationsApplicationCheckMemberObjects**](ApplicationsActionsApi.md#ApplicationsApplicationCheckMemberObjects) | **Post** /applications/{application-id}/microsoft.graph.checkMemberObjects | Invoke action checkMemberObjects
[**ApplicationsApplicationGetMemberGroups**](ApplicationsActionsApi.md#ApplicationsApplicationGetMemberGroups) | **Post** /applications/{application-id}/microsoft.graph.getMemberGroups | Invoke action getMemberGroups
[**ApplicationsApplicationGetMemberObjects**](ApplicationsActionsApi.md#ApplicationsApplicationGetMemberObjects) | **Post** /applications/{application-id}/microsoft.graph.getMemberObjects | Invoke action getMemberObjects
[**ApplicationsApplicationRemoveKey**](ApplicationsActionsApi.md#ApplicationsApplicationRemoveKey) | **Post** /applications/{application-id}/microsoft.graph.removeKey | Invoke action removeKey
[**ApplicationsApplicationRemovePassword**](ApplicationsActionsApi.md#ApplicationsApplicationRemovePassword) | **Post** /applications/{application-id}/microsoft.graph.removePassword | Invoke action removePassword
[**ApplicationsApplicationRestore**](ApplicationsActionsApi.md#ApplicationsApplicationRestore) | **Post** /applications/{application-id}/microsoft.graph.restore | Invoke action restore
[**ApplicationsApplicationSetVerifiedPublisher**](ApplicationsActionsApi.md#ApplicationsApplicationSetVerifiedPublisher) | **Post** /applications/{application-id}/microsoft.graph.setVerifiedPublisher | Invoke action setVerifiedPublisher
[**ApplicationsApplicationUnsetVerifiedPublisher**](ApplicationsActionsApi.md#ApplicationsApplicationUnsetVerifiedPublisher) | **Post** /applications/{application-id}/microsoft.graph.unsetVerifiedPublisher | Invoke action unsetVerifiedPublisher
[**ApplicationsGetAvailableExtensionProperties**](ApplicationsActionsApi.md#ApplicationsGetAvailableExtensionProperties) | **Post** /applications/microsoft.graph.getAvailableExtensionProperties | Invoke action getAvailableExtensionProperties
[**ApplicationsGetByIds**](ApplicationsActionsApi.md#ApplicationsGetByIds) | **Post** /applications/microsoft.graph.getByIds | Invoke action getByIds
[**ApplicationsValidateProperties**](ApplicationsActionsApi.md#ApplicationsValidateProperties) | **Post** /applications/microsoft.graph.validateProperties | Invoke action validateProperties



## ApplicationsApplicationAddKey

> MicrosoftGraphKeyCredential ApplicationsApplicationAddKey(ctx, applicationId).InlineObject6(inlineObject6).Execute()

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
    applicationId := "applicationId_example" // string | key: id of application
    inlineObject6 := *openapiclient.NewInlineObject6() // InlineObject6 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ApplicationsActionsApi.ApplicationsApplicationAddKey(context.Background(), applicationId).InlineObject6(inlineObject6).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsActionsApi.ApplicationsApplicationAddKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApplicationsApplicationAddKey`: MicrosoftGraphKeyCredential
    fmt.Fprintf(os.Stdout, "Response from `ApplicationsActionsApi.ApplicationsApplicationAddKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | key: id of application | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplicationsApplicationAddKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject6** | [**InlineObject6**](InlineObject6.md) |  | 

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


## ApplicationsApplicationAddPassword

> MicrosoftGraphPasswordCredential ApplicationsApplicationAddPassword(ctx, applicationId).InlineObject7(inlineObject7).Execute()

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
    applicationId := "applicationId_example" // string | key: id of application
    inlineObject7 := *openapiclient.NewInlineObject7() // InlineObject7 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ApplicationsActionsApi.ApplicationsApplicationAddPassword(context.Background(), applicationId).InlineObject7(inlineObject7).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsActionsApi.ApplicationsApplicationAddPassword``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApplicationsApplicationAddPassword`: MicrosoftGraphPasswordCredential
    fmt.Fprintf(os.Stdout, "Response from `ApplicationsActionsApi.ApplicationsApplicationAddPassword`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | key: id of application | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplicationsApplicationAddPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject7** | [**InlineObject7**](InlineObject7.md) |  | 

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


## ApplicationsApplicationCheckMemberGroups

> []string ApplicationsApplicationCheckMemberGroups(ctx, applicationId).InlineObject8(inlineObject8).Execute()

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
    applicationId := "applicationId_example" // string | key: id of application
    inlineObject8 := *openapiclient.NewInlineObject8() // InlineObject8 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ApplicationsActionsApi.ApplicationsApplicationCheckMemberGroups(context.Background(), applicationId).InlineObject8(inlineObject8).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsActionsApi.ApplicationsApplicationCheckMemberGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApplicationsApplicationCheckMemberGroups`: []string
    fmt.Fprintf(os.Stdout, "Response from `ApplicationsActionsApi.ApplicationsApplicationCheckMemberGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | key: id of application | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplicationsApplicationCheckMemberGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject8** | [**InlineObject8**](InlineObject8.md) |  | 

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


## ApplicationsApplicationCheckMemberObjects

> []string ApplicationsApplicationCheckMemberObjects(ctx, applicationId).InlineObject9(inlineObject9).Execute()

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
    applicationId := "applicationId_example" // string | key: id of application
    inlineObject9 := *openapiclient.NewInlineObject9() // InlineObject9 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ApplicationsActionsApi.ApplicationsApplicationCheckMemberObjects(context.Background(), applicationId).InlineObject9(inlineObject9).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsActionsApi.ApplicationsApplicationCheckMemberObjects``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApplicationsApplicationCheckMemberObjects`: []string
    fmt.Fprintf(os.Stdout, "Response from `ApplicationsActionsApi.ApplicationsApplicationCheckMemberObjects`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | key: id of application | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplicationsApplicationCheckMemberObjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject9** | [**InlineObject9**](InlineObject9.md) |  | 

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


## ApplicationsApplicationGetMemberGroups

> []string ApplicationsApplicationGetMemberGroups(ctx, applicationId).InlineObject10(inlineObject10).Execute()

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
    applicationId := "applicationId_example" // string | key: id of application
    inlineObject10 := *openapiclient.NewInlineObject10() // InlineObject10 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ApplicationsActionsApi.ApplicationsApplicationGetMemberGroups(context.Background(), applicationId).InlineObject10(inlineObject10).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsActionsApi.ApplicationsApplicationGetMemberGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApplicationsApplicationGetMemberGroups`: []string
    fmt.Fprintf(os.Stdout, "Response from `ApplicationsActionsApi.ApplicationsApplicationGetMemberGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | key: id of application | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplicationsApplicationGetMemberGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject10** | [**InlineObject10**](InlineObject10.md) |  | 

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


## ApplicationsApplicationGetMemberObjects

> []string ApplicationsApplicationGetMemberObjects(ctx, applicationId).InlineObject11(inlineObject11).Execute()

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
    applicationId := "applicationId_example" // string | key: id of application
    inlineObject11 := *openapiclient.NewInlineObject11() // InlineObject11 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ApplicationsActionsApi.ApplicationsApplicationGetMemberObjects(context.Background(), applicationId).InlineObject11(inlineObject11).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsActionsApi.ApplicationsApplicationGetMemberObjects``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApplicationsApplicationGetMemberObjects`: []string
    fmt.Fprintf(os.Stdout, "Response from `ApplicationsActionsApi.ApplicationsApplicationGetMemberObjects`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | key: id of application | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplicationsApplicationGetMemberObjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject11** | [**InlineObject11**](InlineObject11.md) |  | 

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


## ApplicationsApplicationRemoveKey

> ApplicationsApplicationRemoveKey(ctx, applicationId).InlineObject12(inlineObject12).Execute()

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
    applicationId := "applicationId_example" // string | key: id of application
    inlineObject12 := *openapiclient.NewInlineObject12() // InlineObject12 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ApplicationsActionsApi.ApplicationsApplicationRemoveKey(context.Background(), applicationId).InlineObject12(inlineObject12).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsActionsApi.ApplicationsApplicationRemoveKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | key: id of application | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplicationsApplicationRemoveKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject12** | [**InlineObject12**](InlineObject12.md) |  | 

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


## ApplicationsApplicationRemovePassword

> ApplicationsApplicationRemovePassword(ctx, applicationId).InlineObject13(inlineObject13).Execute()

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
    applicationId := "applicationId_example" // string | key: id of application
    inlineObject13 := *openapiclient.NewInlineObject13() // InlineObject13 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ApplicationsActionsApi.ApplicationsApplicationRemovePassword(context.Background(), applicationId).InlineObject13(inlineObject13).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsActionsApi.ApplicationsApplicationRemovePassword``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | key: id of application | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplicationsApplicationRemovePasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject13** | [**InlineObject13**](InlineObject13.md) |  | 

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


## ApplicationsApplicationRestore

> AnyOfmicrosoftGraphDirectoryObject ApplicationsApplicationRestore(ctx, applicationId).Execute()

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
    applicationId := "applicationId_example" // string | key: id of application

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ApplicationsActionsApi.ApplicationsApplicationRestore(context.Background(), applicationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsActionsApi.ApplicationsApplicationRestore``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApplicationsApplicationRestore`: AnyOfmicrosoftGraphDirectoryObject
    fmt.Fprintf(os.Stdout, "Response from `ApplicationsActionsApi.ApplicationsApplicationRestore`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | key: id of application | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplicationsApplicationRestoreRequest struct via the builder pattern


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


## ApplicationsApplicationSetVerifiedPublisher

> ApplicationsApplicationSetVerifiedPublisher(ctx, applicationId).InlineObject14(inlineObject14).Execute()

Invoke action setVerifiedPublisher

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
    inlineObject14 := *openapiclient.NewInlineObject14() // InlineObject14 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ApplicationsActionsApi.ApplicationsApplicationSetVerifiedPublisher(context.Background(), applicationId).InlineObject14(inlineObject14).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsActionsApi.ApplicationsApplicationSetVerifiedPublisher``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | key: id of application | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplicationsApplicationSetVerifiedPublisherRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject14** | [**InlineObject14**](InlineObject14.md) |  | 

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


## ApplicationsApplicationUnsetVerifiedPublisher

> ApplicationsApplicationUnsetVerifiedPublisher(ctx, applicationId).Execute()

Invoke action unsetVerifiedPublisher

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ApplicationsActionsApi.ApplicationsApplicationUnsetVerifiedPublisher(context.Background(), applicationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsActionsApi.ApplicationsApplicationUnsetVerifiedPublisher``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | key: id of application | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplicationsApplicationUnsetVerifiedPublisherRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## ApplicationsGetAvailableExtensionProperties

> []MicrosoftGraphExtensionProperty ApplicationsGetAvailableExtensionProperties(ctx).InlineObject15(inlineObject15).Execute()

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
    inlineObject15 := *openapiclient.NewInlineObject15() // InlineObject15 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ApplicationsActionsApi.ApplicationsGetAvailableExtensionProperties(context.Background()).InlineObject15(inlineObject15).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsActionsApi.ApplicationsGetAvailableExtensionProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApplicationsGetAvailableExtensionProperties`: []MicrosoftGraphExtensionProperty
    fmt.Fprintf(os.Stdout, "Response from `ApplicationsActionsApi.ApplicationsGetAvailableExtensionProperties`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApplicationsGetAvailableExtensionPropertiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject15** | [**InlineObject15**](InlineObject15.md) |  | 

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


## ApplicationsGetByIds

> []MicrosoftGraphDirectoryObject ApplicationsGetByIds(ctx).InlineObject16(inlineObject16).Execute()

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
    inlineObject16 := *openapiclient.NewInlineObject16() // InlineObject16 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ApplicationsActionsApi.ApplicationsGetByIds(context.Background()).InlineObject16(inlineObject16).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsActionsApi.ApplicationsGetByIds``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApplicationsGetByIds`: []MicrosoftGraphDirectoryObject
    fmt.Fprintf(os.Stdout, "Response from `ApplicationsActionsApi.ApplicationsGetByIds`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApplicationsGetByIdsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject16** | [**InlineObject16**](InlineObject16.md) |  | 

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


## ApplicationsValidateProperties

> ApplicationsValidateProperties(ctx).InlineObject17(inlineObject17).Execute()

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
    inlineObject17 := *openapiclient.NewInlineObject17() // InlineObject17 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ApplicationsActionsApi.ApplicationsValidateProperties(context.Background()).InlineObject17(inlineObject17).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsActionsApi.ApplicationsValidateProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApplicationsValidatePropertiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject17** | [**InlineObject17**](InlineObject17.md) |  | 

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

