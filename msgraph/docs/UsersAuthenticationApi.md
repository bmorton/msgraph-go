# \UsersAuthenticationApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersAuthenticationCreateFido2Methods**](UsersAuthenticationApi.md#UsersAuthenticationCreateFido2Methods) | **Post** /users/{user-id}/authentication/fido2Methods | Create new navigation property to fido2Methods for users
[**UsersAuthenticationCreateMethods**](UsersAuthenticationApi.md#UsersAuthenticationCreateMethods) | **Post** /users/{user-id}/authentication/methods | Create new navigation property to methods for users
[**UsersAuthenticationCreateMicrosoftAuthenticatorMethods**](UsersAuthenticationApi.md#UsersAuthenticationCreateMicrosoftAuthenticatorMethods) | **Post** /users/{user-id}/authentication/microsoftAuthenticatorMethods | Create new navigation property to microsoftAuthenticatorMethods for users
[**UsersAuthenticationCreateWindowsHelloForBusinessMethods**](UsersAuthenticationApi.md#UsersAuthenticationCreateWindowsHelloForBusinessMethods) | **Post** /users/{user-id}/authentication/windowsHelloForBusinessMethods | Create new navigation property to windowsHelloForBusinessMethods for users
[**UsersAuthenticationDeleteFido2Methods**](UsersAuthenticationApi.md#UsersAuthenticationDeleteFido2Methods) | **Delete** /users/{user-id}/authentication/fido2Methods/{fido2AuthenticationMethod-id} | Delete navigation property fido2Methods for users
[**UsersAuthenticationDeleteMethods**](UsersAuthenticationApi.md#UsersAuthenticationDeleteMethods) | **Delete** /users/{user-id}/authentication/methods/{authenticationMethod-id} | Delete navigation property methods for users
[**UsersAuthenticationDeleteMicrosoftAuthenticatorMethods**](UsersAuthenticationApi.md#UsersAuthenticationDeleteMicrosoftAuthenticatorMethods) | **Delete** /users/{user-id}/authentication/microsoftAuthenticatorMethods/{microsoftAuthenticatorAuthenticationMethod-id} | Delete navigation property microsoftAuthenticatorMethods for users
[**UsersAuthenticationDeleteWindowsHelloForBusinessMethods**](UsersAuthenticationApi.md#UsersAuthenticationDeleteWindowsHelloForBusinessMethods) | **Delete** /users/{user-id}/authentication/windowsHelloForBusinessMethods/{windowsHelloForBusinessAuthenticationMethod-id} | Delete navigation property windowsHelloForBusinessMethods for users
[**UsersAuthenticationGetFido2Methods**](UsersAuthenticationApi.md#UsersAuthenticationGetFido2Methods) | **Get** /users/{user-id}/authentication/fido2Methods/{fido2AuthenticationMethod-id} | Get fido2Methods from users
[**UsersAuthenticationGetMethods**](UsersAuthenticationApi.md#UsersAuthenticationGetMethods) | **Get** /users/{user-id}/authentication/methods/{authenticationMethod-id} | Get methods from users
[**UsersAuthenticationGetMicrosoftAuthenticatorMethods**](UsersAuthenticationApi.md#UsersAuthenticationGetMicrosoftAuthenticatorMethods) | **Get** /users/{user-id}/authentication/microsoftAuthenticatorMethods/{microsoftAuthenticatorAuthenticationMethod-id} | Get microsoftAuthenticatorMethods from users
[**UsersAuthenticationGetWindowsHelloForBusinessMethods**](UsersAuthenticationApi.md#UsersAuthenticationGetWindowsHelloForBusinessMethods) | **Get** /users/{user-id}/authentication/windowsHelloForBusinessMethods/{windowsHelloForBusinessAuthenticationMethod-id} | Get windowsHelloForBusinessMethods from users
[**UsersAuthenticationListFido2Methods**](UsersAuthenticationApi.md#UsersAuthenticationListFido2Methods) | **Get** /users/{user-id}/authentication/fido2Methods | Get fido2Methods from users
[**UsersAuthenticationListMethods**](UsersAuthenticationApi.md#UsersAuthenticationListMethods) | **Get** /users/{user-id}/authentication/methods | Get methods from users
[**UsersAuthenticationListMicrosoftAuthenticatorMethods**](UsersAuthenticationApi.md#UsersAuthenticationListMicrosoftAuthenticatorMethods) | **Get** /users/{user-id}/authentication/microsoftAuthenticatorMethods | Get microsoftAuthenticatorMethods from users
[**UsersAuthenticationListWindowsHelloForBusinessMethods**](UsersAuthenticationApi.md#UsersAuthenticationListWindowsHelloForBusinessMethods) | **Get** /users/{user-id}/authentication/windowsHelloForBusinessMethods | Get windowsHelloForBusinessMethods from users
[**UsersAuthenticationMicrosoftAuthenticatorMethodsDeleteDevice**](UsersAuthenticationApi.md#UsersAuthenticationMicrosoftAuthenticatorMethodsDeleteDevice) | **Delete** /users/{user-id}/authentication/microsoftAuthenticatorMethods/{microsoftAuthenticatorAuthenticationMethod-id}/device | Delete navigation property device for users
[**UsersAuthenticationMicrosoftAuthenticatorMethodsGetDevice**](UsersAuthenticationApi.md#UsersAuthenticationMicrosoftAuthenticatorMethodsGetDevice) | **Get** /users/{user-id}/authentication/microsoftAuthenticatorMethods/{microsoftAuthenticatorAuthenticationMethod-id}/device | Get device from users
[**UsersAuthenticationMicrosoftAuthenticatorMethodsUpdateDevice**](UsersAuthenticationApi.md#UsersAuthenticationMicrosoftAuthenticatorMethodsUpdateDevice) | **Patch** /users/{user-id}/authentication/microsoftAuthenticatorMethods/{microsoftAuthenticatorAuthenticationMethod-id}/device | Update the navigation property device in users
[**UsersAuthenticationUpdateFido2Methods**](UsersAuthenticationApi.md#UsersAuthenticationUpdateFido2Methods) | **Patch** /users/{user-id}/authentication/fido2Methods/{fido2AuthenticationMethod-id} | Update the navigation property fido2Methods in users
[**UsersAuthenticationUpdateMethods**](UsersAuthenticationApi.md#UsersAuthenticationUpdateMethods) | **Patch** /users/{user-id}/authentication/methods/{authenticationMethod-id} | Update the navigation property methods in users
[**UsersAuthenticationUpdateMicrosoftAuthenticatorMethods**](UsersAuthenticationApi.md#UsersAuthenticationUpdateMicrosoftAuthenticatorMethods) | **Patch** /users/{user-id}/authentication/microsoftAuthenticatorMethods/{microsoftAuthenticatorAuthenticationMethod-id} | Update the navigation property microsoftAuthenticatorMethods in users
[**UsersAuthenticationUpdateWindowsHelloForBusinessMethods**](UsersAuthenticationApi.md#UsersAuthenticationUpdateWindowsHelloForBusinessMethods) | **Patch** /users/{user-id}/authentication/windowsHelloForBusinessMethods/{windowsHelloForBusinessAuthenticationMethod-id} | Update the navigation property windowsHelloForBusinessMethods in users
[**UsersAuthenticationWindowsHelloForBusinessMethodsDeleteDevice**](UsersAuthenticationApi.md#UsersAuthenticationWindowsHelloForBusinessMethodsDeleteDevice) | **Delete** /users/{user-id}/authentication/windowsHelloForBusinessMethods/{windowsHelloForBusinessAuthenticationMethod-id}/device | Delete navigation property device for users
[**UsersAuthenticationWindowsHelloForBusinessMethodsGetDevice**](UsersAuthenticationApi.md#UsersAuthenticationWindowsHelloForBusinessMethodsGetDevice) | **Get** /users/{user-id}/authentication/windowsHelloForBusinessMethods/{windowsHelloForBusinessAuthenticationMethod-id}/device | Get device from users
[**UsersAuthenticationWindowsHelloForBusinessMethodsUpdateDevice**](UsersAuthenticationApi.md#UsersAuthenticationWindowsHelloForBusinessMethodsUpdateDevice) | **Patch** /users/{user-id}/authentication/windowsHelloForBusinessMethods/{windowsHelloForBusinessAuthenticationMethod-id}/device | Update the navigation property device in users
[**UsersDeleteAuthentication**](UsersAuthenticationApi.md#UsersDeleteAuthentication) | **Delete** /users/{user-id}/authentication | Delete navigation property authentication for users
[**UsersGetAuthentication**](UsersAuthenticationApi.md#UsersGetAuthentication) | **Get** /users/{user-id}/authentication | Get authentication from users
[**UsersUpdateAuthentication**](UsersAuthenticationApi.md#UsersUpdateAuthentication) | **Patch** /users/{user-id}/authentication | Update the navigation property authentication in users



## UsersAuthenticationCreateFido2Methods

> MicrosoftGraphFido2AuthenticationMethod UsersAuthenticationCreateFido2Methods(ctx, userId).MicrosoftGraphFido2AuthenticationMethod(microsoftGraphFido2AuthenticationMethod).Execute()

Create new navigation property to fido2Methods for users

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
    userId := "userId_example" // string | key: id of user
    microsoftGraphFido2AuthenticationMethod := *openapiclient.NewMicrosoftGraphFido2AuthenticationMethod() // MicrosoftGraphFido2AuthenticationMethod | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersAuthenticationApi.UsersAuthenticationCreateFido2Methods(context.Background(), userId).MicrosoftGraphFido2AuthenticationMethod(microsoftGraphFido2AuthenticationMethod).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersAuthenticationApi.UsersAuthenticationCreateFido2Methods``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersAuthenticationCreateFido2Methods`: MicrosoftGraphFido2AuthenticationMethod
    fmt.Fprintf(os.Stdout, "Response from `UsersAuthenticationApi.UsersAuthenticationCreateFido2Methods`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersAuthenticationCreateFido2MethodsRequest struct via the builder pattern


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


## UsersAuthenticationCreateMethods

> MicrosoftGraphAuthenticationMethod UsersAuthenticationCreateMethods(ctx, userId).MicrosoftGraphAuthenticationMethod(microsoftGraphAuthenticationMethod).Execute()

Create new navigation property to methods for users

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
    userId := "userId_example" // string | key: id of user
    microsoftGraphAuthenticationMethod := *openapiclient.NewMicrosoftGraphAuthenticationMethod() // MicrosoftGraphAuthenticationMethod | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersAuthenticationApi.UsersAuthenticationCreateMethods(context.Background(), userId).MicrosoftGraphAuthenticationMethod(microsoftGraphAuthenticationMethod).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersAuthenticationApi.UsersAuthenticationCreateMethods``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersAuthenticationCreateMethods`: MicrosoftGraphAuthenticationMethod
    fmt.Fprintf(os.Stdout, "Response from `UsersAuthenticationApi.UsersAuthenticationCreateMethods`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersAuthenticationCreateMethodsRequest struct via the builder pattern


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


## UsersAuthenticationCreateMicrosoftAuthenticatorMethods

> MicrosoftGraphMicrosoftAuthenticatorAuthenticationMethod UsersAuthenticationCreateMicrosoftAuthenticatorMethods(ctx, userId).MicrosoftGraphMicrosoftAuthenticatorAuthenticationMethod(microsoftGraphMicrosoftAuthenticatorAuthenticationMethod).Execute()

Create new navigation property to microsoftAuthenticatorMethods for users

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
    userId := "userId_example" // string | key: id of user
    microsoftGraphMicrosoftAuthenticatorAuthenticationMethod := *openapiclient.NewMicrosoftGraphMicrosoftAuthenticatorAuthenticationMethod() // MicrosoftGraphMicrosoftAuthenticatorAuthenticationMethod | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersAuthenticationApi.UsersAuthenticationCreateMicrosoftAuthenticatorMethods(context.Background(), userId).MicrosoftGraphMicrosoftAuthenticatorAuthenticationMethod(microsoftGraphMicrosoftAuthenticatorAuthenticationMethod).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersAuthenticationApi.UsersAuthenticationCreateMicrosoftAuthenticatorMethods``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersAuthenticationCreateMicrosoftAuthenticatorMethods`: MicrosoftGraphMicrosoftAuthenticatorAuthenticationMethod
    fmt.Fprintf(os.Stdout, "Response from `UsersAuthenticationApi.UsersAuthenticationCreateMicrosoftAuthenticatorMethods`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersAuthenticationCreateMicrosoftAuthenticatorMethodsRequest struct via the builder pattern


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


## UsersAuthenticationCreateWindowsHelloForBusinessMethods

> MicrosoftGraphWindowsHelloForBusinessAuthenticationMethod UsersAuthenticationCreateWindowsHelloForBusinessMethods(ctx, userId).MicrosoftGraphWindowsHelloForBusinessAuthenticationMethod(microsoftGraphWindowsHelloForBusinessAuthenticationMethod).Execute()

Create new navigation property to windowsHelloForBusinessMethods for users

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
    userId := "userId_example" // string | key: id of user
    microsoftGraphWindowsHelloForBusinessAuthenticationMethod := *openapiclient.NewMicrosoftGraphWindowsHelloForBusinessAuthenticationMethod() // MicrosoftGraphWindowsHelloForBusinessAuthenticationMethod | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersAuthenticationApi.UsersAuthenticationCreateWindowsHelloForBusinessMethods(context.Background(), userId).MicrosoftGraphWindowsHelloForBusinessAuthenticationMethod(microsoftGraphWindowsHelloForBusinessAuthenticationMethod).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersAuthenticationApi.UsersAuthenticationCreateWindowsHelloForBusinessMethods``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersAuthenticationCreateWindowsHelloForBusinessMethods`: MicrosoftGraphWindowsHelloForBusinessAuthenticationMethod
    fmt.Fprintf(os.Stdout, "Response from `UsersAuthenticationApi.UsersAuthenticationCreateWindowsHelloForBusinessMethods`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersAuthenticationCreateWindowsHelloForBusinessMethodsRequest struct via the builder pattern


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


## UsersAuthenticationDeleteFido2Methods

> UsersAuthenticationDeleteFido2Methods(ctx, userId, fido2AuthenticationMethodId).IfMatch(ifMatch).Execute()

Delete navigation property fido2Methods for users

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
    userId := "userId_example" // string | key: id of user
    fido2AuthenticationMethodId := "fido2AuthenticationMethodId_example" // string | key: id of fido2AuthenticationMethod
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersAuthenticationApi.UsersAuthenticationDeleteFido2Methods(context.Background(), userId, fido2AuthenticationMethodId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersAuthenticationApi.UsersAuthenticationDeleteFido2Methods``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**fido2AuthenticationMethodId** | **string** | key: id of fido2AuthenticationMethod | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersAuthenticationDeleteFido2MethodsRequest struct via the builder pattern


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


## UsersAuthenticationDeleteMethods

> UsersAuthenticationDeleteMethods(ctx, userId, authenticationMethodId).IfMatch(ifMatch).Execute()

Delete navigation property methods for users

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
    userId := "userId_example" // string | key: id of user
    authenticationMethodId := "authenticationMethodId_example" // string | key: id of authenticationMethod
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersAuthenticationApi.UsersAuthenticationDeleteMethods(context.Background(), userId, authenticationMethodId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersAuthenticationApi.UsersAuthenticationDeleteMethods``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**authenticationMethodId** | **string** | key: id of authenticationMethod | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersAuthenticationDeleteMethodsRequest struct via the builder pattern


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


## UsersAuthenticationDeleteMicrosoftAuthenticatorMethods

> UsersAuthenticationDeleteMicrosoftAuthenticatorMethods(ctx, userId, microsoftAuthenticatorAuthenticationMethodId).IfMatch(ifMatch).Execute()

Delete navigation property microsoftAuthenticatorMethods for users

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
    userId := "userId_example" // string | key: id of user
    microsoftAuthenticatorAuthenticationMethodId := "microsoftAuthenticatorAuthenticationMethodId_example" // string | key: id of microsoftAuthenticatorAuthenticationMethod
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersAuthenticationApi.UsersAuthenticationDeleteMicrosoftAuthenticatorMethods(context.Background(), userId, microsoftAuthenticatorAuthenticationMethodId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersAuthenticationApi.UsersAuthenticationDeleteMicrosoftAuthenticatorMethods``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**microsoftAuthenticatorAuthenticationMethodId** | **string** | key: id of microsoftAuthenticatorAuthenticationMethod | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersAuthenticationDeleteMicrosoftAuthenticatorMethodsRequest struct via the builder pattern


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


## UsersAuthenticationDeleteWindowsHelloForBusinessMethods

> UsersAuthenticationDeleteWindowsHelloForBusinessMethods(ctx, userId, windowsHelloForBusinessAuthenticationMethodId).IfMatch(ifMatch).Execute()

Delete navigation property windowsHelloForBusinessMethods for users

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
    userId := "userId_example" // string | key: id of user
    windowsHelloForBusinessAuthenticationMethodId := "windowsHelloForBusinessAuthenticationMethodId_example" // string | key: id of windowsHelloForBusinessAuthenticationMethod
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersAuthenticationApi.UsersAuthenticationDeleteWindowsHelloForBusinessMethods(context.Background(), userId, windowsHelloForBusinessAuthenticationMethodId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersAuthenticationApi.UsersAuthenticationDeleteWindowsHelloForBusinessMethods``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**windowsHelloForBusinessAuthenticationMethodId** | **string** | key: id of windowsHelloForBusinessAuthenticationMethod | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersAuthenticationDeleteWindowsHelloForBusinessMethodsRequest struct via the builder pattern


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


## UsersAuthenticationGetFido2Methods

> MicrosoftGraphFido2AuthenticationMethod UsersAuthenticationGetFido2Methods(ctx, userId, fido2AuthenticationMethodId).Select_(select_).Expand(expand).Execute()

Get fido2Methods from users

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
    userId := "userId_example" // string | key: id of user
    fido2AuthenticationMethodId := "fido2AuthenticationMethodId_example" // string | key: id of fido2AuthenticationMethod
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersAuthenticationApi.UsersAuthenticationGetFido2Methods(context.Background(), userId, fido2AuthenticationMethodId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersAuthenticationApi.UsersAuthenticationGetFido2Methods``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersAuthenticationGetFido2Methods`: MicrosoftGraphFido2AuthenticationMethod
    fmt.Fprintf(os.Stdout, "Response from `UsersAuthenticationApi.UsersAuthenticationGetFido2Methods`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**fido2AuthenticationMethodId** | **string** | key: id of fido2AuthenticationMethod | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersAuthenticationGetFido2MethodsRequest struct via the builder pattern


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


## UsersAuthenticationGetMethods

> MicrosoftGraphAuthenticationMethod UsersAuthenticationGetMethods(ctx, userId, authenticationMethodId).Select_(select_).Expand(expand).Execute()

Get methods from users

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
    userId := "userId_example" // string | key: id of user
    authenticationMethodId := "authenticationMethodId_example" // string | key: id of authenticationMethod
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersAuthenticationApi.UsersAuthenticationGetMethods(context.Background(), userId, authenticationMethodId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersAuthenticationApi.UsersAuthenticationGetMethods``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersAuthenticationGetMethods`: MicrosoftGraphAuthenticationMethod
    fmt.Fprintf(os.Stdout, "Response from `UsersAuthenticationApi.UsersAuthenticationGetMethods`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**authenticationMethodId** | **string** | key: id of authenticationMethod | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersAuthenticationGetMethodsRequest struct via the builder pattern


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


## UsersAuthenticationGetMicrosoftAuthenticatorMethods

> MicrosoftGraphMicrosoftAuthenticatorAuthenticationMethod UsersAuthenticationGetMicrosoftAuthenticatorMethods(ctx, userId, microsoftAuthenticatorAuthenticationMethodId).Select_(select_).Expand(expand).Execute()

Get microsoftAuthenticatorMethods from users

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
    userId := "userId_example" // string | key: id of user
    microsoftAuthenticatorAuthenticationMethodId := "microsoftAuthenticatorAuthenticationMethodId_example" // string | key: id of microsoftAuthenticatorAuthenticationMethod
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersAuthenticationApi.UsersAuthenticationGetMicrosoftAuthenticatorMethods(context.Background(), userId, microsoftAuthenticatorAuthenticationMethodId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersAuthenticationApi.UsersAuthenticationGetMicrosoftAuthenticatorMethods``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersAuthenticationGetMicrosoftAuthenticatorMethods`: MicrosoftGraphMicrosoftAuthenticatorAuthenticationMethod
    fmt.Fprintf(os.Stdout, "Response from `UsersAuthenticationApi.UsersAuthenticationGetMicrosoftAuthenticatorMethods`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**microsoftAuthenticatorAuthenticationMethodId** | **string** | key: id of microsoftAuthenticatorAuthenticationMethod | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersAuthenticationGetMicrosoftAuthenticatorMethodsRequest struct via the builder pattern


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


## UsersAuthenticationGetWindowsHelloForBusinessMethods

> MicrosoftGraphWindowsHelloForBusinessAuthenticationMethod UsersAuthenticationGetWindowsHelloForBusinessMethods(ctx, userId, windowsHelloForBusinessAuthenticationMethodId).Select_(select_).Expand(expand).Execute()

Get windowsHelloForBusinessMethods from users

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
    userId := "userId_example" // string | key: id of user
    windowsHelloForBusinessAuthenticationMethodId := "windowsHelloForBusinessAuthenticationMethodId_example" // string | key: id of windowsHelloForBusinessAuthenticationMethod
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersAuthenticationApi.UsersAuthenticationGetWindowsHelloForBusinessMethods(context.Background(), userId, windowsHelloForBusinessAuthenticationMethodId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersAuthenticationApi.UsersAuthenticationGetWindowsHelloForBusinessMethods``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersAuthenticationGetWindowsHelloForBusinessMethods`: MicrosoftGraphWindowsHelloForBusinessAuthenticationMethod
    fmt.Fprintf(os.Stdout, "Response from `UsersAuthenticationApi.UsersAuthenticationGetWindowsHelloForBusinessMethods`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**windowsHelloForBusinessAuthenticationMethodId** | **string** | key: id of windowsHelloForBusinessAuthenticationMethod | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersAuthenticationGetWindowsHelloForBusinessMethodsRequest struct via the builder pattern


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


## UsersAuthenticationListFido2Methods

> CollectionOfFido2AuthenticationMethod UsersAuthenticationListFido2Methods(ctx, userId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get fido2Methods from users

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
    userId := "userId_example" // string | key: id of user
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
    resp, r, err := api_client.UsersAuthenticationApi.UsersAuthenticationListFido2Methods(context.Background(), userId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersAuthenticationApi.UsersAuthenticationListFido2Methods``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersAuthenticationListFido2Methods`: CollectionOfFido2AuthenticationMethod
    fmt.Fprintf(os.Stdout, "Response from `UsersAuthenticationApi.UsersAuthenticationListFido2Methods`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersAuthenticationListFido2MethodsRequest struct via the builder pattern


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


## UsersAuthenticationListMethods

> CollectionOfAuthenticationMethod UsersAuthenticationListMethods(ctx, userId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get methods from users

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
    userId := "userId_example" // string | key: id of user
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
    resp, r, err := api_client.UsersAuthenticationApi.UsersAuthenticationListMethods(context.Background(), userId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersAuthenticationApi.UsersAuthenticationListMethods``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersAuthenticationListMethods`: CollectionOfAuthenticationMethod
    fmt.Fprintf(os.Stdout, "Response from `UsersAuthenticationApi.UsersAuthenticationListMethods`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersAuthenticationListMethodsRequest struct via the builder pattern


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


## UsersAuthenticationListMicrosoftAuthenticatorMethods

> CollectionOfMicrosoftAuthenticatorAuthenticationMethod UsersAuthenticationListMicrosoftAuthenticatorMethods(ctx, userId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get microsoftAuthenticatorMethods from users

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
    userId := "userId_example" // string | key: id of user
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
    resp, r, err := api_client.UsersAuthenticationApi.UsersAuthenticationListMicrosoftAuthenticatorMethods(context.Background(), userId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersAuthenticationApi.UsersAuthenticationListMicrosoftAuthenticatorMethods``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersAuthenticationListMicrosoftAuthenticatorMethods`: CollectionOfMicrosoftAuthenticatorAuthenticationMethod
    fmt.Fprintf(os.Stdout, "Response from `UsersAuthenticationApi.UsersAuthenticationListMicrosoftAuthenticatorMethods`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersAuthenticationListMicrosoftAuthenticatorMethodsRequest struct via the builder pattern


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


## UsersAuthenticationListWindowsHelloForBusinessMethods

> CollectionOfWindowsHelloForBusinessAuthenticationMethod UsersAuthenticationListWindowsHelloForBusinessMethods(ctx, userId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get windowsHelloForBusinessMethods from users

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
    userId := "userId_example" // string | key: id of user
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
    resp, r, err := api_client.UsersAuthenticationApi.UsersAuthenticationListWindowsHelloForBusinessMethods(context.Background(), userId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersAuthenticationApi.UsersAuthenticationListWindowsHelloForBusinessMethods``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersAuthenticationListWindowsHelloForBusinessMethods`: CollectionOfWindowsHelloForBusinessAuthenticationMethod
    fmt.Fprintf(os.Stdout, "Response from `UsersAuthenticationApi.UsersAuthenticationListWindowsHelloForBusinessMethods`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersAuthenticationListWindowsHelloForBusinessMethodsRequest struct via the builder pattern


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


## UsersAuthenticationMicrosoftAuthenticatorMethodsDeleteDevice

> UsersAuthenticationMicrosoftAuthenticatorMethodsDeleteDevice(ctx, userId, microsoftAuthenticatorAuthenticationMethodId).IfMatch(ifMatch).Execute()

Delete navigation property device for users



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
    userId := "userId_example" // string | key: id of user
    microsoftAuthenticatorAuthenticationMethodId := "microsoftAuthenticatorAuthenticationMethodId_example" // string | key: id of microsoftAuthenticatorAuthenticationMethod
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersAuthenticationApi.UsersAuthenticationMicrosoftAuthenticatorMethodsDeleteDevice(context.Background(), userId, microsoftAuthenticatorAuthenticationMethodId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersAuthenticationApi.UsersAuthenticationMicrosoftAuthenticatorMethodsDeleteDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**microsoftAuthenticatorAuthenticationMethodId** | **string** | key: id of microsoftAuthenticatorAuthenticationMethod | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersAuthenticationMicrosoftAuthenticatorMethodsDeleteDeviceRequest struct via the builder pattern


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


## UsersAuthenticationMicrosoftAuthenticatorMethodsGetDevice

> MicrosoftGraphDevice UsersAuthenticationMicrosoftAuthenticatorMethodsGetDevice(ctx, userId, microsoftAuthenticatorAuthenticationMethodId).Select_(select_).Expand(expand).Execute()

Get device from users



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
    userId := "userId_example" // string | key: id of user
    microsoftAuthenticatorAuthenticationMethodId := "microsoftAuthenticatorAuthenticationMethodId_example" // string | key: id of microsoftAuthenticatorAuthenticationMethod
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersAuthenticationApi.UsersAuthenticationMicrosoftAuthenticatorMethodsGetDevice(context.Background(), userId, microsoftAuthenticatorAuthenticationMethodId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersAuthenticationApi.UsersAuthenticationMicrosoftAuthenticatorMethodsGetDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersAuthenticationMicrosoftAuthenticatorMethodsGetDevice`: MicrosoftGraphDevice
    fmt.Fprintf(os.Stdout, "Response from `UsersAuthenticationApi.UsersAuthenticationMicrosoftAuthenticatorMethodsGetDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**microsoftAuthenticatorAuthenticationMethodId** | **string** | key: id of microsoftAuthenticatorAuthenticationMethod | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersAuthenticationMicrosoftAuthenticatorMethodsGetDeviceRequest struct via the builder pattern


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


## UsersAuthenticationMicrosoftAuthenticatorMethodsUpdateDevice

> UsersAuthenticationMicrosoftAuthenticatorMethodsUpdateDevice(ctx, userId, microsoftAuthenticatorAuthenticationMethodId).MicrosoftGraphDevice(microsoftGraphDevice).Execute()

Update the navigation property device in users



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
    userId := "userId_example" // string | key: id of user
    microsoftAuthenticatorAuthenticationMethodId := "microsoftAuthenticatorAuthenticationMethodId_example" // string | key: id of microsoftAuthenticatorAuthenticationMethod
    microsoftGraphDevice := *openapiclient.NewMicrosoftGraphDevice() // MicrosoftGraphDevice | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersAuthenticationApi.UsersAuthenticationMicrosoftAuthenticatorMethodsUpdateDevice(context.Background(), userId, microsoftAuthenticatorAuthenticationMethodId).MicrosoftGraphDevice(microsoftGraphDevice).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersAuthenticationApi.UsersAuthenticationMicrosoftAuthenticatorMethodsUpdateDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**microsoftAuthenticatorAuthenticationMethodId** | **string** | key: id of microsoftAuthenticatorAuthenticationMethod | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersAuthenticationMicrosoftAuthenticatorMethodsUpdateDeviceRequest struct via the builder pattern


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


## UsersAuthenticationUpdateFido2Methods

> UsersAuthenticationUpdateFido2Methods(ctx, userId, fido2AuthenticationMethodId).MicrosoftGraphFido2AuthenticationMethod(microsoftGraphFido2AuthenticationMethod).Execute()

Update the navigation property fido2Methods in users

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
    userId := "userId_example" // string | key: id of user
    fido2AuthenticationMethodId := "fido2AuthenticationMethodId_example" // string | key: id of fido2AuthenticationMethod
    microsoftGraphFido2AuthenticationMethod := *openapiclient.NewMicrosoftGraphFido2AuthenticationMethod() // MicrosoftGraphFido2AuthenticationMethod | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersAuthenticationApi.UsersAuthenticationUpdateFido2Methods(context.Background(), userId, fido2AuthenticationMethodId).MicrosoftGraphFido2AuthenticationMethod(microsoftGraphFido2AuthenticationMethod).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersAuthenticationApi.UsersAuthenticationUpdateFido2Methods``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**fido2AuthenticationMethodId** | **string** | key: id of fido2AuthenticationMethod | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersAuthenticationUpdateFido2MethodsRequest struct via the builder pattern


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


## UsersAuthenticationUpdateMethods

> UsersAuthenticationUpdateMethods(ctx, userId, authenticationMethodId).MicrosoftGraphAuthenticationMethod(microsoftGraphAuthenticationMethod).Execute()

Update the navigation property methods in users

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
    userId := "userId_example" // string | key: id of user
    authenticationMethodId := "authenticationMethodId_example" // string | key: id of authenticationMethod
    microsoftGraphAuthenticationMethod := *openapiclient.NewMicrosoftGraphAuthenticationMethod() // MicrosoftGraphAuthenticationMethod | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersAuthenticationApi.UsersAuthenticationUpdateMethods(context.Background(), userId, authenticationMethodId).MicrosoftGraphAuthenticationMethod(microsoftGraphAuthenticationMethod).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersAuthenticationApi.UsersAuthenticationUpdateMethods``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**authenticationMethodId** | **string** | key: id of authenticationMethod | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersAuthenticationUpdateMethodsRequest struct via the builder pattern


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


## UsersAuthenticationUpdateMicrosoftAuthenticatorMethods

> UsersAuthenticationUpdateMicrosoftAuthenticatorMethods(ctx, userId, microsoftAuthenticatorAuthenticationMethodId).MicrosoftGraphMicrosoftAuthenticatorAuthenticationMethod(microsoftGraphMicrosoftAuthenticatorAuthenticationMethod).Execute()

Update the navigation property microsoftAuthenticatorMethods in users

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
    userId := "userId_example" // string | key: id of user
    microsoftAuthenticatorAuthenticationMethodId := "microsoftAuthenticatorAuthenticationMethodId_example" // string | key: id of microsoftAuthenticatorAuthenticationMethod
    microsoftGraphMicrosoftAuthenticatorAuthenticationMethod := *openapiclient.NewMicrosoftGraphMicrosoftAuthenticatorAuthenticationMethod() // MicrosoftGraphMicrosoftAuthenticatorAuthenticationMethod | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersAuthenticationApi.UsersAuthenticationUpdateMicrosoftAuthenticatorMethods(context.Background(), userId, microsoftAuthenticatorAuthenticationMethodId).MicrosoftGraphMicrosoftAuthenticatorAuthenticationMethod(microsoftGraphMicrosoftAuthenticatorAuthenticationMethod).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersAuthenticationApi.UsersAuthenticationUpdateMicrosoftAuthenticatorMethods``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**microsoftAuthenticatorAuthenticationMethodId** | **string** | key: id of microsoftAuthenticatorAuthenticationMethod | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersAuthenticationUpdateMicrosoftAuthenticatorMethodsRequest struct via the builder pattern


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


## UsersAuthenticationUpdateWindowsHelloForBusinessMethods

> UsersAuthenticationUpdateWindowsHelloForBusinessMethods(ctx, userId, windowsHelloForBusinessAuthenticationMethodId).MicrosoftGraphWindowsHelloForBusinessAuthenticationMethod(microsoftGraphWindowsHelloForBusinessAuthenticationMethod).Execute()

Update the navigation property windowsHelloForBusinessMethods in users

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
    userId := "userId_example" // string | key: id of user
    windowsHelloForBusinessAuthenticationMethodId := "windowsHelloForBusinessAuthenticationMethodId_example" // string | key: id of windowsHelloForBusinessAuthenticationMethod
    microsoftGraphWindowsHelloForBusinessAuthenticationMethod := *openapiclient.NewMicrosoftGraphWindowsHelloForBusinessAuthenticationMethod() // MicrosoftGraphWindowsHelloForBusinessAuthenticationMethod | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersAuthenticationApi.UsersAuthenticationUpdateWindowsHelloForBusinessMethods(context.Background(), userId, windowsHelloForBusinessAuthenticationMethodId).MicrosoftGraphWindowsHelloForBusinessAuthenticationMethod(microsoftGraphWindowsHelloForBusinessAuthenticationMethod).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersAuthenticationApi.UsersAuthenticationUpdateWindowsHelloForBusinessMethods``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**windowsHelloForBusinessAuthenticationMethodId** | **string** | key: id of windowsHelloForBusinessAuthenticationMethod | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersAuthenticationUpdateWindowsHelloForBusinessMethodsRequest struct via the builder pattern


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


## UsersAuthenticationWindowsHelloForBusinessMethodsDeleteDevice

> UsersAuthenticationWindowsHelloForBusinessMethodsDeleteDevice(ctx, userId, windowsHelloForBusinessAuthenticationMethodId).IfMatch(ifMatch).Execute()

Delete navigation property device for users



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
    userId := "userId_example" // string | key: id of user
    windowsHelloForBusinessAuthenticationMethodId := "windowsHelloForBusinessAuthenticationMethodId_example" // string | key: id of windowsHelloForBusinessAuthenticationMethod
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersAuthenticationApi.UsersAuthenticationWindowsHelloForBusinessMethodsDeleteDevice(context.Background(), userId, windowsHelloForBusinessAuthenticationMethodId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersAuthenticationApi.UsersAuthenticationWindowsHelloForBusinessMethodsDeleteDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**windowsHelloForBusinessAuthenticationMethodId** | **string** | key: id of windowsHelloForBusinessAuthenticationMethod | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersAuthenticationWindowsHelloForBusinessMethodsDeleteDeviceRequest struct via the builder pattern


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


## UsersAuthenticationWindowsHelloForBusinessMethodsGetDevice

> MicrosoftGraphDevice UsersAuthenticationWindowsHelloForBusinessMethodsGetDevice(ctx, userId, windowsHelloForBusinessAuthenticationMethodId).Select_(select_).Expand(expand).Execute()

Get device from users



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
    userId := "userId_example" // string | key: id of user
    windowsHelloForBusinessAuthenticationMethodId := "windowsHelloForBusinessAuthenticationMethodId_example" // string | key: id of windowsHelloForBusinessAuthenticationMethod
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersAuthenticationApi.UsersAuthenticationWindowsHelloForBusinessMethodsGetDevice(context.Background(), userId, windowsHelloForBusinessAuthenticationMethodId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersAuthenticationApi.UsersAuthenticationWindowsHelloForBusinessMethodsGetDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersAuthenticationWindowsHelloForBusinessMethodsGetDevice`: MicrosoftGraphDevice
    fmt.Fprintf(os.Stdout, "Response from `UsersAuthenticationApi.UsersAuthenticationWindowsHelloForBusinessMethodsGetDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**windowsHelloForBusinessAuthenticationMethodId** | **string** | key: id of windowsHelloForBusinessAuthenticationMethod | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersAuthenticationWindowsHelloForBusinessMethodsGetDeviceRequest struct via the builder pattern


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


## UsersAuthenticationWindowsHelloForBusinessMethodsUpdateDevice

> UsersAuthenticationWindowsHelloForBusinessMethodsUpdateDevice(ctx, userId, windowsHelloForBusinessAuthenticationMethodId).MicrosoftGraphDevice(microsoftGraphDevice).Execute()

Update the navigation property device in users



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
    userId := "userId_example" // string | key: id of user
    windowsHelloForBusinessAuthenticationMethodId := "windowsHelloForBusinessAuthenticationMethodId_example" // string | key: id of windowsHelloForBusinessAuthenticationMethod
    microsoftGraphDevice := *openapiclient.NewMicrosoftGraphDevice() // MicrosoftGraphDevice | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersAuthenticationApi.UsersAuthenticationWindowsHelloForBusinessMethodsUpdateDevice(context.Background(), userId, windowsHelloForBusinessAuthenticationMethodId).MicrosoftGraphDevice(microsoftGraphDevice).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersAuthenticationApi.UsersAuthenticationWindowsHelloForBusinessMethodsUpdateDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**windowsHelloForBusinessAuthenticationMethodId** | **string** | key: id of windowsHelloForBusinessAuthenticationMethod | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersAuthenticationWindowsHelloForBusinessMethodsUpdateDeviceRequest struct via the builder pattern


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


## UsersDeleteAuthentication

> UsersDeleteAuthentication(ctx, userId).IfMatch(ifMatch).Execute()

Delete navigation property authentication for users

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
    userId := "userId_example" // string | key: id of user
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersAuthenticationApi.UsersDeleteAuthentication(context.Background(), userId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersAuthenticationApi.UsersDeleteAuthentication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersDeleteAuthenticationRequest struct via the builder pattern


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


## UsersGetAuthentication

> MicrosoftGraphAuthentication UsersGetAuthentication(ctx, userId).Select_(select_).Expand(expand).Execute()

Get authentication from users

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
    userId := "userId_example" // string | key: id of user
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersAuthenticationApi.UsersGetAuthentication(context.Background(), userId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersAuthenticationApi.UsersGetAuthentication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersGetAuthentication`: MicrosoftGraphAuthentication
    fmt.Fprintf(os.Stdout, "Response from `UsersAuthenticationApi.UsersGetAuthentication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersGetAuthenticationRequest struct via the builder pattern


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


## UsersUpdateAuthentication

> UsersUpdateAuthentication(ctx, userId).MicrosoftGraphAuthentication(microsoftGraphAuthentication).Execute()

Update the navigation property authentication in users

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
    userId := "userId_example" // string | key: id of user
    microsoftGraphAuthentication := *openapiclient.NewMicrosoftGraphAuthentication() // MicrosoftGraphAuthentication | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersAuthenticationApi.UsersUpdateAuthentication(context.Background(), userId).MicrosoftGraphAuthentication(microsoftGraphAuthentication).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersAuthenticationApi.UsersUpdateAuthentication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersUpdateAuthenticationRequest struct via the builder pattern


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

