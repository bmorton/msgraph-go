# \MeAuthenticationApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MeAuthenticationCreateFido2Methods**](MeAuthenticationApi.md#MeAuthenticationCreateFido2Methods) | **Post** /me/authentication/fido2Methods | Create new navigation property to fido2Methods for me
[**MeAuthenticationCreateMethods**](MeAuthenticationApi.md#MeAuthenticationCreateMethods) | **Post** /me/authentication/methods | Create new navigation property to methods for me
[**MeAuthenticationCreateMicrosoftAuthenticatorMethods**](MeAuthenticationApi.md#MeAuthenticationCreateMicrosoftAuthenticatorMethods) | **Post** /me/authentication/microsoftAuthenticatorMethods | Create new navigation property to microsoftAuthenticatorMethods for me
[**MeAuthenticationCreateWindowsHelloForBusinessMethods**](MeAuthenticationApi.md#MeAuthenticationCreateWindowsHelloForBusinessMethods) | **Post** /me/authentication/windowsHelloForBusinessMethods | Create new navigation property to windowsHelloForBusinessMethods for me
[**MeAuthenticationDeleteFido2Methods**](MeAuthenticationApi.md#MeAuthenticationDeleteFido2Methods) | **Delete** /me/authentication/fido2Methods/{fido2AuthenticationMethod-id} | Delete navigation property fido2Methods for me
[**MeAuthenticationDeleteMethods**](MeAuthenticationApi.md#MeAuthenticationDeleteMethods) | **Delete** /me/authentication/methods/{authenticationMethod-id} | Delete navigation property methods for me
[**MeAuthenticationDeleteMicrosoftAuthenticatorMethods**](MeAuthenticationApi.md#MeAuthenticationDeleteMicrosoftAuthenticatorMethods) | **Delete** /me/authentication/microsoftAuthenticatorMethods/{microsoftAuthenticatorAuthenticationMethod-id} | Delete navigation property microsoftAuthenticatorMethods for me
[**MeAuthenticationDeleteWindowsHelloForBusinessMethods**](MeAuthenticationApi.md#MeAuthenticationDeleteWindowsHelloForBusinessMethods) | **Delete** /me/authentication/windowsHelloForBusinessMethods/{windowsHelloForBusinessAuthenticationMethod-id} | Delete navigation property windowsHelloForBusinessMethods for me
[**MeAuthenticationGetFido2Methods**](MeAuthenticationApi.md#MeAuthenticationGetFido2Methods) | **Get** /me/authentication/fido2Methods/{fido2AuthenticationMethod-id} | Get fido2Methods from me
[**MeAuthenticationGetMethods**](MeAuthenticationApi.md#MeAuthenticationGetMethods) | **Get** /me/authentication/methods/{authenticationMethod-id} | Get methods from me
[**MeAuthenticationGetMicrosoftAuthenticatorMethods**](MeAuthenticationApi.md#MeAuthenticationGetMicrosoftAuthenticatorMethods) | **Get** /me/authentication/microsoftAuthenticatorMethods/{microsoftAuthenticatorAuthenticationMethod-id} | Get microsoftAuthenticatorMethods from me
[**MeAuthenticationGetWindowsHelloForBusinessMethods**](MeAuthenticationApi.md#MeAuthenticationGetWindowsHelloForBusinessMethods) | **Get** /me/authentication/windowsHelloForBusinessMethods/{windowsHelloForBusinessAuthenticationMethod-id} | Get windowsHelloForBusinessMethods from me
[**MeAuthenticationListFido2Methods**](MeAuthenticationApi.md#MeAuthenticationListFido2Methods) | **Get** /me/authentication/fido2Methods | Get fido2Methods from me
[**MeAuthenticationListMethods**](MeAuthenticationApi.md#MeAuthenticationListMethods) | **Get** /me/authentication/methods | Get methods from me
[**MeAuthenticationListMicrosoftAuthenticatorMethods**](MeAuthenticationApi.md#MeAuthenticationListMicrosoftAuthenticatorMethods) | **Get** /me/authentication/microsoftAuthenticatorMethods | Get microsoftAuthenticatorMethods from me
[**MeAuthenticationListWindowsHelloForBusinessMethods**](MeAuthenticationApi.md#MeAuthenticationListWindowsHelloForBusinessMethods) | **Get** /me/authentication/windowsHelloForBusinessMethods | Get windowsHelloForBusinessMethods from me
[**MeAuthenticationMicrosoftAuthenticatorMethodsDeleteDevice**](MeAuthenticationApi.md#MeAuthenticationMicrosoftAuthenticatorMethodsDeleteDevice) | **Delete** /me/authentication/microsoftAuthenticatorMethods/{microsoftAuthenticatorAuthenticationMethod-id}/device | Delete navigation property device for me
[**MeAuthenticationMicrosoftAuthenticatorMethodsGetDevice**](MeAuthenticationApi.md#MeAuthenticationMicrosoftAuthenticatorMethodsGetDevice) | **Get** /me/authentication/microsoftAuthenticatorMethods/{microsoftAuthenticatorAuthenticationMethod-id}/device | Get device from me
[**MeAuthenticationMicrosoftAuthenticatorMethodsUpdateDevice**](MeAuthenticationApi.md#MeAuthenticationMicrosoftAuthenticatorMethodsUpdateDevice) | **Patch** /me/authentication/microsoftAuthenticatorMethods/{microsoftAuthenticatorAuthenticationMethod-id}/device | Update the navigation property device in me
[**MeAuthenticationUpdateFido2Methods**](MeAuthenticationApi.md#MeAuthenticationUpdateFido2Methods) | **Patch** /me/authentication/fido2Methods/{fido2AuthenticationMethod-id} | Update the navigation property fido2Methods in me
[**MeAuthenticationUpdateMethods**](MeAuthenticationApi.md#MeAuthenticationUpdateMethods) | **Patch** /me/authentication/methods/{authenticationMethod-id} | Update the navigation property methods in me
[**MeAuthenticationUpdateMicrosoftAuthenticatorMethods**](MeAuthenticationApi.md#MeAuthenticationUpdateMicrosoftAuthenticatorMethods) | **Patch** /me/authentication/microsoftAuthenticatorMethods/{microsoftAuthenticatorAuthenticationMethod-id} | Update the navigation property microsoftAuthenticatorMethods in me
[**MeAuthenticationUpdateWindowsHelloForBusinessMethods**](MeAuthenticationApi.md#MeAuthenticationUpdateWindowsHelloForBusinessMethods) | **Patch** /me/authentication/windowsHelloForBusinessMethods/{windowsHelloForBusinessAuthenticationMethod-id} | Update the navigation property windowsHelloForBusinessMethods in me
[**MeAuthenticationWindowsHelloForBusinessMethodsDeleteDevice**](MeAuthenticationApi.md#MeAuthenticationWindowsHelloForBusinessMethodsDeleteDevice) | **Delete** /me/authentication/windowsHelloForBusinessMethods/{windowsHelloForBusinessAuthenticationMethod-id}/device | Delete navigation property device for me
[**MeAuthenticationWindowsHelloForBusinessMethodsGetDevice**](MeAuthenticationApi.md#MeAuthenticationWindowsHelloForBusinessMethodsGetDevice) | **Get** /me/authentication/windowsHelloForBusinessMethods/{windowsHelloForBusinessAuthenticationMethod-id}/device | Get device from me
[**MeAuthenticationWindowsHelloForBusinessMethodsUpdateDevice**](MeAuthenticationApi.md#MeAuthenticationWindowsHelloForBusinessMethodsUpdateDevice) | **Patch** /me/authentication/windowsHelloForBusinessMethods/{windowsHelloForBusinessAuthenticationMethod-id}/device | Update the navigation property device in me
[**MeDeleteAuthentication**](MeAuthenticationApi.md#MeDeleteAuthentication) | **Delete** /me/authentication | Delete navigation property authentication for me
[**MeGetAuthentication**](MeAuthenticationApi.md#MeGetAuthentication) | **Get** /me/authentication | Get authentication from me
[**MeUpdateAuthentication**](MeAuthenticationApi.md#MeUpdateAuthentication) | **Patch** /me/authentication | Update the navigation property authentication in me



## MeAuthenticationCreateFido2Methods

> MicrosoftGraphFido2AuthenticationMethod MeAuthenticationCreateFido2Methods(ctx).MicrosoftGraphFido2AuthenticationMethod(microsoftGraphFido2AuthenticationMethod).Execute()

Create new navigation property to fido2Methods for me

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
    microsoftGraphFido2AuthenticationMethod := *openapiclient.NewMicrosoftGraphFido2AuthenticationMethod() // MicrosoftGraphFido2AuthenticationMethod | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeAuthenticationApi.MeAuthenticationCreateFido2Methods(context.Background()).MicrosoftGraphFido2AuthenticationMethod(microsoftGraphFido2AuthenticationMethod).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeAuthenticationApi.MeAuthenticationCreateFido2Methods``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeAuthenticationCreateFido2Methods`: MicrosoftGraphFido2AuthenticationMethod
    fmt.Fprintf(os.Stdout, "Response from `MeAuthenticationApi.MeAuthenticationCreateFido2Methods`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMeAuthenticationCreateFido2MethodsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphFido2AuthenticationMethod** | [**MicrosoftGraphFido2AuthenticationMethod**](MicrosoftGraphFido2AuthenticationMethod.md) | New navigation property | 

### Return type

[**MicrosoftGraphFido2AuthenticationMethod**](MicrosoftGraphFido2AuthenticationMethod.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeAuthenticationCreateMethods

> MicrosoftGraphAuthenticationMethod MeAuthenticationCreateMethods(ctx).MicrosoftGraphAuthenticationMethod(microsoftGraphAuthenticationMethod).Execute()

Create new navigation property to methods for me

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
    microsoftGraphAuthenticationMethod := *openapiclient.NewMicrosoftGraphAuthenticationMethod() // MicrosoftGraphAuthenticationMethod | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeAuthenticationApi.MeAuthenticationCreateMethods(context.Background()).MicrosoftGraphAuthenticationMethod(microsoftGraphAuthenticationMethod).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeAuthenticationApi.MeAuthenticationCreateMethods``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeAuthenticationCreateMethods`: MicrosoftGraphAuthenticationMethod
    fmt.Fprintf(os.Stdout, "Response from `MeAuthenticationApi.MeAuthenticationCreateMethods`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMeAuthenticationCreateMethodsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphAuthenticationMethod** | [**MicrosoftGraphAuthenticationMethod**](MicrosoftGraphAuthenticationMethod.md) | New navigation property | 

### Return type

[**MicrosoftGraphAuthenticationMethod**](MicrosoftGraphAuthenticationMethod.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeAuthenticationCreateMicrosoftAuthenticatorMethods

> MicrosoftGraphMicrosoftAuthenticatorAuthenticationMethod MeAuthenticationCreateMicrosoftAuthenticatorMethods(ctx).MicrosoftGraphMicrosoftAuthenticatorAuthenticationMethod(microsoftGraphMicrosoftAuthenticatorAuthenticationMethod).Execute()

Create new navigation property to microsoftAuthenticatorMethods for me

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
    microsoftGraphMicrosoftAuthenticatorAuthenticationMethod := *openapiclient.NewMicrosoftGraphMicrosoftAuthenticatorAuthenticationMethod() // MicrosoftGraphMicrosoftAuthenticatorAuthenticationMethod | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeAuthenticationApi.MeAuthenticationCreateMicrosoftAuthenticatorMethods(context.Background()).MicrosoftGraphMicrosoftAuthenticatorAuthenticationMethod(microsoftGraphMicrosoftAuthenticatorAuthenticationMethod).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeAuthenticationApi.MeAuthenticationCreateMicrosoftAuthenticatorMethods``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeAuthenticationCreateMicrosoftAuthenticatorMethods`: MicrosoftGraphMicrosoftAuthenticatorAuthenticationMethod
    fmt.Fprintf(os.Stdout, "Response from `MeAuthenticationApi.MeAuthenticationCreateMicrosoftAuthenticatorMethods`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMeAuthenticationCreateMicrosoftAuthenticatorMethodsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphMicrosoftAuthenticatorAuthenticationMethod** | [**MicrosoftGraphMicrosoftAuthenticatorAuthenticationMethod**](MicrosoftGraphMicrosoftAuthenticatorAuthenticationMethod.md) | New navigation property | 

### Return type

[**MicrosoftGraphMicrosoftAuthenticatorAuthenticationMethod**](MicrosoftGraphMicrosoftAuthenticatorAuthenticationMethod.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeAuthenticationCreateWindowsHelloForBusinessMethods

> MicrosoftGraphWindowsHelloForBusinessAuthenticationMethod MeAuthenticationCreateWindowsHelloForBusinessMethods(ctx).MicrosoftGraphWindowsHelloForBusinessAuthenticationMethod(microsoftGraphWindowsHelloForBusinessAuthenticationMethod).Execute()

Create new navigation property to windowsHelloForBusinessMethods for me

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
    microsoftGraphWindowsHelloForBusinessAuthenticationMethod := *openapiclient.NewMicrosoftGraphWindowsHelloForBusinessAuthenticationMethod() // MicrosoftGraphWindowsHelloForBusinessAuthenticationMethod | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeAuthenticationApi.MeAuthenticationCreateWindowsHelloForBusinessMethods(context.Background()).MicrosoftGraphWindowsHelloForBusinessAuthenticationMethod(microsoftGraphWindowsHelloForBusinessAuthenticationMethod).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeAuthenticationApi.MeAuthenticationCreateWindowsHelloForBusinessMethods``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeAuthenticationCreateWindowsHelloForBusinessMethods`: MicrosoftGraphWindowsHelloForBusinessAuthenticationMethod
    fmt.Fprintf(os.Stdout, "Response from `MeAuthenticationApi.MeAuthenticationCreateWindowsHelloForBusinessMethods`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMeAuthenticationCreateWindowsHelloForBusinessMethodsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphWindowsHelloForBusinessAuthenticationMethod** | [**MicrosoftGraphWindowsHelloForBusinessAuthenticationMethod**](MicrosoftGraphWindowsHelloForBusinessAuthenticationMethod.md) | New navigation property | 

### Return type

[**MicrosoftGraphWindowsHelloForBusinessAuthenticationMethod**](MicrosoftGraphWindowsHelloForBusinessAuthenticationMethod.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeAuthenticationDeleteFido2Methods

> MeAuthenticationDeleteFido2Methods(ctx, fido2AuthenticationMethodId).IfMatch(ifMatch).Execute()

Delete navigation property fido2Methods for me

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
    fido2AuthenticationMethodId := "fido2AuthenticationMethodId_example" // string | key: id of fido2AuthenticationMethod
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeAuthenticationApi.MeAuthenticationDeleteFido2Methods(context.Background(), fido2AuthenticationMethodId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeAuthenticationApi.MeAuthenticationDeleteFido2Methods``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fido2AuthenticationMethodId** | **string** | key: id of fido2AuthenticationMethod | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeAuthenticationDeleteFido2MethodsRequest struct via the builder pattern


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


## MeAuthenticationDeleteMethods

> MeAuthenticationDeleteMethods(ctx, authenticationMethodId).IfMatch(ifMatch).Execute()

Delete navigation property methods for me

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
    authenticationMethodId := "authenticationMethodId_example" // string | key: id of authenticationMethod
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeAuthenticationApi.MeAuthenticationDeleteMethods(context.Background(), authenticationMethodId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeAuthenticationApi.MeAuthenticationDeleteMethods``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authenticationMethodId** | **string** | key: id of authenticationMethod | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeAuthenticationDeleteMethodsRequest struct via the builder pattern


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


## MeAuthenticationDeleteMicrosoftAuthenticatorMethods

> MeAuthenticationDeleteMicrosoftAuthenticatorMethods(ctx, microsoftAuthenticatorAuthenticationMethodId).IfMatch(ifMatch).Execute()

Delete navigation property microsoftAuthenticatorMethods for me

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
    microsoftAuthenticatorAuthenticationMethodId := "microsoftAuthenticatorAuthenticationMethodId_example" // string | key: id of microsoftAuthenticatorAuthenticationMethod
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeAuthenticationApi.MeAuthenticationDeleteMicrosoftAuthenticatorMethods(context.Background(), microsoftAuthenticatorAuthenticationMethodId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeAuthenticationApi.MeAuthenticationDeleteMicrosoftAuthenticatorMethods``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftAuthenticatorAuthenticationMethodId** | **string** | key: id of microsoftAuthenticatorAuthenticationMethod | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeAuthenticationDeleteMicrosoftAuthenticatorMethodsRequest struct via the builder pattern


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


## MeAuthenticationDeleteWindowsHelloForBusinessMethods

> MeAuthenticationDeleteWindowsHelloForBusinessMethods(ctx, windowsHelloForBusinessAuthenticationMethodId).IfMatch(ifMatch).Execute()

Delete navigation property windowsHelloForBusinessMethods for me

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
    windowsHelloForBusinessAuthenticationMethodId := "windowsHelloForBusinessAuthenticationMethodId_example" // string | key: id of windowsHelloForBusinessAuthenticationMethod
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeAuthenticationApi.MeAuthenticationDeleteWindowsHelloForBusinessMethods(context.Background(), windowsHelloForBusinessAuthenticationMethodId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeAuthenticationApi.MeAuthenticationDeleteWindowsHelloForBusinessMethods``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**windowsHelloForBusinessAuthenticationMethodId** | **string** | key: id of windowsHelloForBusinessAuthenticationMethod | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeAuthenticationDeleteWindowsHelloForBusinessMethodsRequest struct via the builder pattern


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


## MeAuthenticationGetFido2Methods

> MicrosoftGraphFido2AuthenticationMethod MeAuthenticationGetFido2Methods(ctx, fido2AuthenticationMethodId).Select_(select_).Expand(expand).Execute()

Get fido2Methods from me

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
    fido2AuthenticationMethodId := "fido2AuthenticationMethodId_example" // string | key: id of fido2AuthenticationMethod
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeAuthenticationApi.MeAuthenticationGetFido2Methods(context.Background(), fido2AuthenticationMethodId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeAuthenticationApi.MeAuthenticationGetFido2Methods``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeAuthenticationGetFido2Methods`: MicrosoftGraphFido2AuthenticationMethod
    fmt.Fprintf(os.Stdout, "Response from `MeAuthenticationApi.MeAuthenticationGetFido2Methods`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fido2AuthenticationMethodId** | **string** | key: id of fido2AuthenticationMethod | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeAuthenticationGetFido2MethodsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphFido2AuthenticationMethod**](MicrosoftGraphFido2AuthenticationMethod.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeAuthenticationGetMethods

> MicrosoftGraphAuthenticationMethod MeAuthenticationGetMethods(ctx, authenticationMethodId).Select_(select_).Expand(expand).Execute()

Get methods from me

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
    authenticationMethodId := "authenticationMethodId_example" // string | key: id of authenticationMethod
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeAuthenticationApi.MeAuthenticationGetMethods(context.Background(), authenticationMethodId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeAuthenticationApi.MeAuthenticationGetMethods``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeAuthenticationGetMethods`: MicrosoftGraphAuthenticationMethod
    fmt.Fprintf(os.Stdout, "Response from `MeAuthenticationApi.MeAuthenticationGetMethods`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authenticationMethodId** | **string** | key: id of authenticationMethod | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeAuthenticationGetMethodsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphAuthenticationMethod**](MicrosoftGraphAuthenticationMethod.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeAuthenticationGetMicrosoftAuthenticatorMethods

> MicrosoftGraphMicrosoftAuthenticatorAuthenticationMethod MeAuthenticationGetMicrosoftAuthenticatorMethods(ctx, microsoftAuthenticatorAuthenticationMethodId).Select_(select_).Expand(expand).Execute()

Get microsoftAuthenticatorMethods from me

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
    microsoftAuthenticatorAuthenticationMethodId := "microsoftAuthenticatorAuthenticationMethodId_example" // string | key: id of microsoftAuthenticatorAuthenticationMethod
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeAuthenticationApi.MeAuthenticationGetMicrosoftAuthenticatorMethods(context.Background(), microsoftAuthenticatorAuthenticationMethodId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeAuthenticationApi.MeAuthenticationGetMicrosoftAuthenticatorMethods``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeAuthenticationGetMicrosoftAuthenticatorMethods`: MicrosoftGraphMicrosoftAuthenticatorAuthenticationMethod
    fmt.Fprintf(os.Stdout, "Response from `MeAuthenticationApi.MeAuthenticationGetMicrosoftAuthenticatorMethods`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftAuthenticatorAuthenticationMethodId** | **string** | key: id of microsoftAuthenticatorAuthenticationMethod | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeAuthenticationGetMicrosoftAuthenticatorMethodsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphMicrosoftAuthenticatorAuthenticationMethod**](MicrosoftGraphMicrosoftAuthenticatorAuthenticationMethod.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeAuthenticationGetWindowsHelloForBusinessMethods

> MicrosoftGraphWindowsHelloForBusinessAuthenticationMethod MeAuthenticationGetWindowsHelloForBusinessMethods(ctx, windowsHelloForBusinessAuthenticationMethodId).Select_(select_).Expand(expand).Execute()

Get windowsHelloForBusinessMethods from me

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
    windowsHelloForBusinessAuthenticationMethodId := "windowsHelloForBusinessAuthenticationMethodId_example" // string | key: id of windowsHelloForBusinessAuthenticationMethod
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeAuthenticationApi.MeAuthenticationGetWindowsHelloForBusinessMethods(context.Background(), windowsHelloForBusinessAuthenticationMethodId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeAuthenticationApi.MeAuthenticationGetWindowsHelloForBusinessMethods``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeAuthenticationGetWindowsHelloForBusinessMethods`: MicrosoftGraphWindowsHelloForBusinessAuthenticationMethod
    fmt.Fprintf(os.Stdout, "Response from `MeAuthenticationApi.MeAuthenticationGetWindowsHelloForBusinessMethods`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**windowsHelloForBusinessAuthenticationMethodId** | **string** | key: id of windowsHelloForBusinessAuthenticationMethod | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeAuthenticationGetWindowsHelloForBusinessMethodsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphWindowsHelloForBusinessAuthenticationMethod**](MicrosoftGraphWindowsHelloForBusinessAuthenticationMethod.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeAuthenticationListFido2Methods

> CollectionOfFido2AuthenticationMethod MeAuthenticationListFido2Methods(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get fido2Methods from me

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
    resp, r, err := api_client.MeAuthenticationApi.MeAuthenticationListFido2Methods(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeAuthenticationApi.MeAuthenticationListFido2Methods``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeAuthenticationListFido2Methods`: CollectionOfFido2AuthenticationMethod
    fmt.Fprintf(os.Stdout, "Response from `MeAuthenticationApi.MeAuthenticationListFido2Methods`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMeAuthenticationListFido2MethodsRequest struct via the builder pattern


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

[**CollectionOfFido2AuthenticationMethod**](CollectionOfFido2AuthenticationMethod.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeAuthenticationListMethods

> CollectionOfAuthenticationMethod MeAuthenticationListMethods(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get methods from me

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
    resp, r, err := api_client.MeAuthenticationApi.MeAuthenticationListMethods(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeAuthenticationApi.MeAuthenticationListMethods``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeAuthenticationListMethods`: CollectionOfAuthenticationMethod
    fmt.Fprintf(os.Stdout, "Response from `MeAuthenticationApi.MeAuthenticationListMethods`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMeAuthenticationListMethodsRequest struct via the builder pattern


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

[**CollectionOfAuthenticationMethod**](CollectionOfAuthenticationMethod.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeAuthenticationListMicrosoftAuthenticatorMethods

> CollectionOfMicrosoftAuthenticatorAuthenticationMethod MeAuthenticationListMicrosoftAuthenticatorMethods(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get microsoftAuthenticatorMethods from me

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
    resp, r, err := api_client.MeAuthenticationApi.MeAuthenticationListMicrosoftAuthenticatorMethods(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeAuthenticationApi.MeAuthenticationListMicrosoftAuthenticatorMethods``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeAuthenticationListMicrosoftAuthenticatorMethods`: CollectionOfMicrosoftAuthenticatorAuthenticationMethod
    fmt.Fprintf(os.Stdout, "Response from `MeAuthenticationApi.MeAuthenticationListMicrosoftAuthenticatorMethods`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMeAuthenticationListMicrosoftAuthenticatorMethodsRequest struct via the builder pattern


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

[**CollectionOfMicrosoftAuthenticatorAuthenticationMethod**](CollectionOfMicrosoftAuthenticatorAuthenticationMethod.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeAuthenticationListWindowsHelloForBusinessMethods

> CollectionOfWindowsHelloForBusinessAuthenticationMethod MeAuthenticationListWindowsHelloForBusinessMethods(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get windowsHelloForBusinessMethods from me

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
    resp, r, err := api_client.MeAuthenticationApi.MeAuthenticationListWindowsHelloForBusinessMethods(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeAuthenticationApi.MeAuthenticationListWindowsHelloForBusinessMethods``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeAuthenticationListWindowsHelloForBusinessMethods`: CollectionOfWindowsHelloForBusinessAuthenticationMethod
    fmt.Fprintf(os.Stdout, "Response from `MeAuthenticationApi.MeAuthenticationListWindowsHelloForBusinessMethods`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMeAuthenticationListWindowsHelloForBusinessMethodsRequest struct via the builder pattern


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

[**CollectionOfWindowsHelloForBusinessAuthenticationMethod**](CollectionOfWindowsHelloForBusinessAuthenticationMethod.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeAuthenticationMicrosoftAuthenticatorMethodsDeleteDevice

> MeAuthenticationMicrosoftAuthenticatorMethodsDeleteDevice(ctx, microsoftAuthenticatorAuthenticationMethodId).IfMatch(ifMatch).Execute()

Delete navigation property device for me



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
    microsoftAuthenticatorAuthenticationMethodId := "microsoftAuthenticatorAuthenticationMethodId_example" // string | key: id of microsoftAuthenticatorAuthenticationMethod
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeAuthenticationApi.MeAuthenticationMicrosoftAuthenticatorMethodsDeleteDevice(context.Background(), microsoftAuthenticatorAuthenticationMethodId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeAuthenticationApi.MeAuthenticationMicrosoftAuthenticatorMethodsDeleteDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftAuthenticatorAuthenticationMethodId** | **string** | key: id of microsoftAuthenticatorAuthenticationMethod | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeAuthenticationMicrosoftAuthenticatorMethodsDeleteDeviceRequest struct via the builder pattern


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


## MeAuthenticationMicrosoftAuthenticatorMethodsGetDevice

> MicrosoftGraphDevice MeAuthenticationMicrosoftAuthenticatorMethodsGetDevice(ctx, microsoftAuthenticatorAuthenticationMethodId).Select_(select_).Expand(expand).Execute()

Get device from me



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
    microsoftAuthenticatorAuthenticationMethodId := "microsoftAuthenticatorAuthenticationMethodId_example" // string | key: id of microsoftAuthenticatorAuthenticationMethod
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeAuthenticationApi.MeAuthenticationMicrosoftAuthenticatorMethodsGetDevice(context.Background(), microsoftAuthenticatorAuthenticationMethodId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeAuthenticationApi.MeAuthenticationMicrosoftAuthenticatorMethodsGetDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeAuthenticationMicrosoftAuthenticatorMethodsGetDevice`: MicrosoftGraphDevice
    fmt.Fprintf(os.Stdout, "Response from `MeAuthenticationApi.MeAuthenticationMicrosoftAuthenticatorMethodsGetDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftAuthenticatorAuthenticationMethodId** | **string** | key: id of microsoftAuthenticatorAuthenticationMethod | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeAuthenticationMicrosoftAuthenticatorMethodsGetDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphDevice**](MicrosoftGraphDevice.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeAuthenticationMicrosoftAuthenticatorMethodsUpdateDevice

> MeAuthenticationMicrosoftAuthenticatorMethodsUpdateDevice(ctx, microsoftAuthenticatorAuthenticationMethodId).MicrosoftGraphDevice(microsoftGraphDevice).Execute()

Update the navigation property device in me



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
    microsoftAuthenticatorAuthenticationMethodId := "microsoftAuthenticatorAuthenticationMethodId_example" // string | key: id of microsoftAuthenticatorAuthenticationMethod
    microsoftGraphDevice := *openapiclient.NewMicrosoftGraphDevice() // MicrosoftGraphDevice | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeAuthenticationApi.MeAuthenticationMicrosoftAuthenticatorMethodsUpdateDevice(context.Background(), microsoftAuthenticatorAuthenticationMethodId).MicrosoftGraphDevice(microsoftGraphDevice).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeAuthenticationApi.MeAuthenticationMicrosoftAuthenticatorMethodsUpdateDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftAuthenticatorAuthenticationMethodId** | **string** | key: id of microsoftAuthenticatorAuthenticationMethod | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeAuthenticationMicrosoftAuthenticatorMethodsUpdateDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphDevice** | [**MicrosoftGraphDevice**](MicrosoftGraphDevice.md) | New navigation property values | 

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


## MeAuthenticationUpdateFido2Methods

> MeAuthenticationUpdateFido2Methods(ctx, fido2AuthenticationMethodId).MicrosoftGraphFido2AuthenticationMethod(microsoftGraphFido2AuthenticationMethod).Execute()

Update the navigation property fido2Methods in me

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
    fido2AuthenticationMethodId := "fido2AuthenticationMethodId_example" // string | key: id of fido2AuthenticationMethod
    microsoftGraphFido2AuthenticationMethod := *openapiclient.NewMicrosoftGraphFido2AuthenticationMethod() // MicrosoftGraphFido2AuthenticationMethod | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeAuthenticationApi.MeAuthenticationUpdateFido2Methods(context.Background(), fido2AuthenticationMethodId).MicrosoftGraphFido2AuthenticationMethod(microsoftGraphFido2AuthenticationMethod).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeAuthenticationApi.MeAuthenticationUpdateFido2Methods``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fido2AuthenticationMethodId** | **string** | key: id of fido2AuthenticationMethod | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeAuthenticationUpdateFido2MethodsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphFido2AuthenticationMethod** | [**MicrosoftGraphFido2AuthenticationMethod**](MicrosoftGraphFido2AuthenticationMethod.md) | New navigation property values | 

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


## MeAuthenticationUpdateMethods

> MeAuthenticationUpdateMethods(ctx, authenticationMethodId).MicrosoftGraphAuthenticationMethod(microsoftGraphAuthenticationMethod).Execute()

Update the navigation property methods in me

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
    authenticationMethodId := "authenticationMethodId_example" // string | key: id of authenticationMethod
    microsoftGraphAuthenticationMethod := *openapiclient.NewMicrosoftGraphAuthenticationMethod() // MicrosoftGraphAuthenticationMethod | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeAuthenticationApi.MeAuthenticationUpdateMethods(context.Background(), authenticationMethodId).MicrosoftGraphAuthenticationMethod(microsoftGraphAuthenticationMethod).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeAuthenticationApi.MeAuthenticationUpdateMethods``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authenticationMethodId** | **string** | key: id of authenticationMethod | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeAuthenticationUpdateMethodsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphAuthenticationMethod** | [**MicrosoftGraphAuthenticationMethod**](MicrosoftGraphAuthenticationMethod.md) | New navigation property values | 

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


## MeAuthenticationUpdateMicrosoftAuthenticatorMethods

> MeAuthenticationUpdateMicrosoftAuthenticatorMethods(ctx, microsoftAuthenticatorAuthenticationMethodId).MicrosoftGraphMicrosoftAuthenticatorAuthenticationMethod(microsoftGraphMicrosoftAuthenticatorAuthenticationMethod).Execute()

Update the navigation property microsoftAuthenticatorMethods in me

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
    microsoftAuthenticatorAuthenticationMethodId := "microsoftAuthenticatorAuthenticationMethodId_example" // string | key: id of microsoftAuthenticatorAuthenticationMethod
    microsoftGraphMicrosoftAuthenticatorAuthenticationMethod := *openapiclient.NewMicrosoftGraphMicrosoftAuthenticatorAuthenticationMethod() // MicrosoftGraphMicrosoftAuthenticatorAuthenticationMethod | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeAuthenticationApi.MeAuthenticationUpdateMicrosoftAuthenticatorMethods(context.Background(), microsoftAuthenticatorAuthenticationMethodId).MicrosoftGraphMicrosoftAuthenticatorAuthenticationMethod(microsoftGraphMicrosoftAuthenticatorAuthenticationMethod).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeAuthenticationApi.MeAuthenticationUpdateMicrosoftAuthenticatorMethods``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**microsoftAuthenticatorAuthenticationMethodId** | **string** | key: id of microsoftAuthenticatorAuthenticationMethod | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeAuthenticationUpdateMicrosoftAuthenticatorMethodsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphMicrosoftAuthenticatorAuthenticationMethod** | [**MicrosoftGraphMicrosoftAuthenticatorAuthenticationMethod**](MicrosoftGraphMicrosoftAuthenticatorAuthenticationMethod.md) | New navigation property values | 

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


## MeAuthenticationUpdateWindowsHelloForBusinessMethods

> MeAuthenticationUpdateWindowsHelloForBusinessMethods(ctx, windowsHelloForBusinessAuthenticationMethodId).MicrosoftGraphWindowsHelloForBusinessAuthenticationMethod(microsoftGraphWindowsHelloForBusinessAuthenticationMethod).Execute()

Update the navigation property windowsHelloForBusinessMethods in me

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
    windowsHelloForBusinessAuthenticationMethodId := "windowsHelloForBusinessAuthenticationMethodId_example" // string | key: id of windowsHelloForBusinessAuthenticationMethod
    microsoftGraphWindowsHelloForBusinessAuthenticationMethod := *openapiclient.NewMicrosoftGraphWindowsHelloForBusinessAuthenticationMethod() // MicrosoftGraphWindowsHelloForBusinessAuthenticationMethod | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeAuthenticationApi.MeAuthenticationUpdateWindowsHelloForBusinessMethods(context.Background(), windowsHelloForBusinessAuthenticationMethodId).MicrosoftGraphWindowsHelloForBusinessAuthenticationMethod(microsoftGraphWindowsHelloForBusinessAuthenticationMethod).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeAuthenticationApi.MeAuthenticationUpdateWindowsHelloForBusinessMethods``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**windowsHelloForBusinessAuthenticationMethodId** | **string** | key: id of windowsHelloForBusinessAuthenticationMethod | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeAuthenticationUpdateWindowsHelloForBusinessMethodsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphWindowsHelloForBusinessAuthenticationMethod** | [**MicrosoftGraphWindowsHelloForBusinessAuthenticationMethod**](MicrosoftGraphWindowsHelloForBusinessAuthenticationMethod.md) | New navigation property values | 

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


## MeAuthenticationWindowsHelloForBusinessMethodsDeleteDevice

> MeAuthenticationWindowsHelloForBusinessMethodsDeleteDevice(ctx, windowsHelloForBusinessAuthenticationMethodId).IfMatch(ifMatch).Execute()

Delete navigation property device for me



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
    windowsHelloForBusinessAuthenticationMethodId := "windowsHelloForBusinessAuthenticationMethodId_example" // string | key: id of windowsHelloForBusinessAuthenticationMethod
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeAuthenticationApi.MeAuthenticationWindowsHelloForBusinessMethodsDeleteDevice(context.Background(), windowsHelloForBusinessAuthenticationMethodId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeAuthenticationApi.MeAuthenticationWindowsHelloForBusinessMethodsDeleteDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**windowsHelloForBusinessAuthenticationMethodId** | **string** | key: id of windowsHelloForBusinessAuthenticationMethod | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeAuthenticationWindowsHelloForBusinessMethodsDeleteDeviceRequest struct via the builder pattern


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


## MeAuthenticationWindowsHelloForBusinessMethodsGetDevice

> MicrosoftGraphDevice MeAuthenticationWindowsHelloForBusinessMethodsGetDevice(ctx, windowsHelloForBusinessAuthenticationMethodId).Select_(select_).Expand(expand).Execute()

Get device from me



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
    windowsHelloForBusinessAuthenticationMethodId := "windowsHelloForBusinessAuthenticationMethodId_example" // string | key: id of windowsHelloForBusinessAuthenticationMethod
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeAuthenticationApi.MeAuthenticationWindowsHelloForBusinessMethodsGetDevice(context.Background(), windowsHelloForBusinessAuthenticationMethodId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeAuthenticationApi.MeAuthenticationWindowsHelloForBusinessMethodsGetDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeAuthenticationWindowsHelloForBusinessMethodsGetDevice`: MicrosoftGraphDevice
    fmt.Fprintf(os.Stdout, "Response from `MeAuthenticationApi.MeAuthenticationWindowsHelloForBusinessMethodsGetDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**windowsHelloForBusinessAuthenticationMethodId** | **string** | key: id of windowsHelloForBusinessAuthenticationMethod | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeAuthenticationWindowsHelloForBusinessMethodsGetDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphDevice**](MicrosoftGraphDevice.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeAuthenticationWindowsHelloForBusinessMethodsUpdateDevice

> MeAuthenticationWindowsHelloForBusinessMethodsUpdateDevice(ctx, windowsHelloForBusinessAuthenticationMethodId).MicrosoftGraphDevice(microsoftGraphDevice).Execute()

Update the navigation property device in me



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
    windowsHelloForBusinessAuthenticationMethodId := "windowsHelloForBusinessAuthenticationMethodId_example" // string | key: id of windowsHelloForBusinessAuthenticationMethod
    microsoftGraphDevice := *openapiclient.NewMicrosoftGraphDevice() // MicrosoftGraphDevice | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeAuthenticationApi.MeAuthenticationWindowsHelloForBusinessMethodsUpdateDevice(context.Background(), windowsHelloForBusinessAuthenticationMethodId).MicrosoftGraphDevice(microsoftGraphDevice).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeAuthenticationApi.MeAuthenticationWindowsHelloForBusinessMethodsUpdateDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**windowsHelloForBusinessAuthenticationMethodId** | **string** | key: id of windowsHelloForBusinessAuthenticationMethod | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeAuthenticationWindowsHelloForBusinessMethodsUpdateDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphDevice** | [**MicrosoftGraphDevice**](MicrosoftGraphDevice.md) | New navigation property values | 

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


## MeDeleteAuthentication

> MeDeleteAuthentication(ctx).IfMatch(ifMatch).Execute()

Delete navigation property authentication for me

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
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeAuthenticationApi.MeDeleteAuthentication(context.Background()).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeAuthenticationApi.MeDeleteAuthentication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMeDeleteAuthenticationRequest struct via the builder pattern


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


## MeGetAuthentication

> MicrosoftGraphAuthentication MeGetAuthentication(ctx).Select_(select_).Expand(expand).Execute()

Get authentication from me

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
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeAuthenticationApi.MeGetAuthentication(context.Background()).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeAuthenticationApi.MeGetAuthentication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeGetAuthentication`: MicrosoftGraphAuthentication
    fmt.Fprintf(os.Stdout, "Response from `MeAuthenticationApi.MeGetAuthentication`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMeGetAuthenticationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphAuthentication**](MicrosoftGraphAuthentication.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeUpdateAuthentication

> MeUpdateAuthentication(ctx).MicrosoftGraphAuthentication(microsoftGraphAuthentication).Execute()

Update the navigation property authentication in me

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
    microsoftGraphAuthentication := *openapiclient.NewMicrosoftGraphAuthentication() // MicrosoftGraphAuthentication | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeAuthenticationApi.MeUpdateAuthentication(context.Background()).MicrosoftGraphAuthentication(microsoftGraphAuthentication).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeAuthenticationApi.MeUpdateAuthentication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMeUpdateAuthenticationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphAuthentication** | [**MicrosoftGraphAuthentication**](MicrosoftGraphAuthentication.md) | New navigation property values | 

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

