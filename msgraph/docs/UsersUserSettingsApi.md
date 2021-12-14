# \UsersUserSettingsApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersDeleteSettings**](UsersUserSettingsApi.md#UsersDeleteSettings) | **Delete** /users/{user-id}/settings | Delete navigation property settings for users
[**UsersGetSettings**](UsersUserSettingsApi.md#UsersGetSettings) | **Get** /users/{user-id}/settings | Get settings from users
[**UsersSettingsDeleteShiftPreferences**](UsersUserSettingsApi.md#UsersSettingsDeleteShiftPreferences) | **Delete** /users/{user-id}/settings/shiftPreferences | Delete navigation property shiftPreferences for users
[**UsersSettingsGetShiftPreferences**](UsersUserSettingsApi.md#UsersSettingsGetShiftPreferences) | **Get** /users/{user-id}/settings/shiftPreferences | Get shiftPreferences from users
[**UsersSettingsUpdateShiftPreferences**](UsersUserSettingsApi.md#UsersSettingsUpdateShiftPreferences) | **Patch** /users/{user-id}/settings/shiftPreferences | Update the navigation property shiftPreferences in users
[**UsersUpdateSettings**](UsersUserSettingsApi.md#UsersUpdateSettings) | **Patch** /users/{user-id}/settings | Update the navigation property settings in users



## UsersDeleteSettings

> UsersDeleteSettings(ctx, userId).IfMatch(ifMatch).Execute()

Delete navigation property settings for users



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
    resp, r, err := api_client.UsersUserSettingsApi.UsersDeleteSettings(context.Background(), userId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersUserSettingsApi.UsersDeleteSettings``: %v\n", err)
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

Other parameters are passed through a pointer to a apiUsersDeleteSettingsRequest struct via the builder pattern


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


## UsersGetSettings

> MicrosoftGraphUserSettings UsersGetSettings(ctx, userId).Select_(select_).Expand(expand).Execute()

Get settings from users



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
    resp, r, err := api_client.UsersUserSettingsApi.UsersGetSettings(context.Background(), userId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersUserSettingsApi.UsersGetSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersGetSettings`: MicrosoftGraphUserSettings
    fmt.Fprintf(os.Stdout, "Response from `UsersUserSettingsApi.UsersGetSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersGetSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphUserSettings**](MicrosoftGraphUserSettings.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersSettingsDeleteShiftPreferences

> UsersSettingsDeleteShiftPreferences(ctx, userId).IfMatch(ifMatch).Execute()

Delete navigation property shiftPreferences for users



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
    resp, r, err := api_client.UsersUserSettingsApi.UsersSettingsDeleteShiftPreferences(context.Background(), userId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersUserSettingsApi.UsersSettingsDeleteShiftPreferences``: %v\n", err)
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

Other parameters are passed through a pointer to a apiUsersSettingsDeleteShiftPreferencesRequest struct via the builder pattern


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


## UsersSettingsGetShiftPreferences

> MicrosoftGraphShiftPreferences UsersSettingsGetShiftPreferences(ctx, userId).Select_(select_).Expand(expand).Execute()

Get shiftPreferences from users



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
    resp, r, err := api_client.UsersUserSettingsApi.UsersSettingsGetShiftPreferences(context.Background(), userId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersUserSettingsApi.UsersSettingsGetShiftPreferences``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersSettingsGetShiftPreferences`: MicrosoftGraphShiftPreferences
    fmt.Fprintf(os.Stdout, "Response from `UsersUserSettingsApi.UsersSettingsGetShiftPreferences`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersSettingsGetShiftPreferencesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphShiftPreferences**](MicrosoftGraphShiftPreferences.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersSettingsUpdateShiftPreferences

> UsersSettingsUpdateShiftPreferences(ctx, userId).MicrosoftGraphShiftPreferences(microsoftGraphShiftPreferences).Execute()

Update the navigation property shiftPreferences in users



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
    microsoftGraphShiftPreferences := *openapiclient.NewMicrosoftGraphShiftPreferences() // MicrosoftGraphShiftPreferences | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersUserSettingsApi.UsersSettingsUpdateShiftPreferences(context.Background(), userId).MicrosoftGraphShiftPreferences(microsoftGraphShiftPreferences).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersUserSettingsApi.UsersSettingsUpdateShiftPreferences``: %v\n", err)
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

Other parameters are passed through a pointer to a apiUsersSettingsUpdateShiftPreferencesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphShiftPreferences** | [**MicrosoftGraphShiftPreferences**](MicrosoftGraphShiftPreferences.md) | New navigation property values | 

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


## UsersUpdateSettings

> UsersUpdateSettings(ctx, userId).MicrosoftGraphUserSettings(microsoftGraphUserSettings).Execute()

Update the navigation property settings in users



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
    microsoftGraphUserSettings := *openapiclient.NewMicrosoftGraphUserSettings() // MicrosoftGraphUserSettings | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersUserSettingsApi.UsersUpdateSettings(context.Background(), userId).MicrosoftGraphUserSettings(microsoftGraphUserSettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersUserSettingsApi.UsersUpdateSettings``: %v\n", err)
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

Other parameters are passed through a pointer to a apiUsersUpdateSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphUserSettings** | [**MicrosoftGraphUserSettings**](MicrosoftGraphUserSettings.md) | New navigation property values | 

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

