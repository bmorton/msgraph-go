# \MeUserSettingsApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MeDeleteSettings**](MeUserSettingsApi.md#MeDeleteSettings) | **Delete** /me/settings | Delete navigation property settings for me
[**MeGetSettings**](MeUserSettingsApi.md#MeGetSettings) | **Get** /me/settings | Get settings from me
[**MeSettingsDeleteShiftPreferences**](MeUserSettingsApi.md#MeSettingsDeleteShiftPreferences) | **Delete** /me/settings/shiftPreferences | Delete navigation property shiftPreferences for me
[**MeSettingsGetShiftPreferences**](MeUserSettingsApi.md#MeSettingsGetShiftPreferences) | **Get** /me/settings/shiftPreferences | Get shiftPreferences from me
[**MeSettingsUpdateShiftPreferences**](MeUserSettingsApi.md#MeSettingsUpdateShiftPreferences) | **Patch** /me/settings/shiftPreferences | Update the navigation property shiftPreferences in me
[**MeUpdateSettings**](MeUserSettingsApi.md#MeUpdateSettings) | **Patch** /me/settings | Update the navigation property settings in me



## MeDeleteSettings

> MeDeleteSettings(ctx).IfMatch(ifMatch).Execute()

Delete navigation property settings for me



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
    resp, r, err := api_client.MeUserSettingsApi.MeDeleteSettings(context.Background()).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeUserSettingsApi.MeDeleteSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMeDeleteSettingsRequest struct via the builder pattern


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


## MeGetSettings

> MicrosoftGraphUserSettings MeGetSettings(ctx).Select_(select_).Expand(expand).Execute()

Get settings from me



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
    resp, r, err := api_client.MeUserSettingsApi.MeGetSettings(context.Background()).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeUserSettingsApi.MeGetSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeGetSettings`: MicrosoftGraphUserSettings
    fmt.Fprintf(os.Stdout, "Response from `MeUserSettingsApi.MeGetSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMeGetSettingsRequest struct via the builder pattern


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


## MeSettingsDeleteShiftPreferences

> MeSettingsDeleteShiftPreferences(ctx).IfMatch(ifMatch).Execute()

Delete navigation property shiftPreferences for me



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
    resp, r, err := api_client.MeUserSettingsApi.MeSettingsDeleteShiftPreferences(context.Background()).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeUserSettingsApi.MeSettingsDeleteShiftPreferences``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMeSettingsDeleteShiftPreferencesRequest struct via the builder pattern


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


## MeSettingsGetShiftPreferences

> MicrosoftGraphShiftPreferences MeSettingsGetShiftPreferences(ctx).Select_(select_).Expand(expand).Execute()

Get shiftPreferences from me



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
    resp, r, err := api_client.MeUserSettingsApi.MeSettingsGetShiftPreferences(context.Background()).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeUserSettingsApi.MeSettingsGetShiftPreferences``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeSettingsGetShiftPreferences`: MicrosoftGraphShiftPreferences
    fmt.Fprintf(os.Stdout, "Response from `MeUserSettingsApi.MeSettingsGetShiftPreferences`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMeSettingsGetShiftPreferencesRequest struct via the builder pattern


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


## MeSettingsUpdateShiftPreferences

> MeSettingsUpdateShiftPreferences(ctx).MicrosoftGraphShiftPreferences(microsoftGraphShiftPreferences).Execute()

Update the navigation property shiftPreferences in me



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
    microsoftGraphShiftPreferences := *openapiclient.NewMicrosoftGraphShiftPreferences() // MicrosoftGraphShiftPreferences | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeUserSettingsApi.MeSettingsUpdateShiftPreferences(context.Background()).MicrosoftGraphShiftPreferences(microsoftGraphShiftPreferences).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeUserSettingsApi.MeSettingsUpdateShiftPreferences``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMeSettingsUpdateShiftPreferencesRequest struct via the builder pattern


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


## MeUpdateSettings

> MeUpdateSettings(ctx).MicrosoftGraphUserSettings(microsoftGraphUserSettings).Execute()

Update the navigation property settings in me



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
    microsoftGraphUserSettings := *openapiclient.NewMicrosoftGraphUserSettings() // MicrosoftGraphUserSettings | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeUserSettingsApi.MeUpdateSettings(context.Background()).MicrosoftGraphUserSettings(microsoftGraphUserSettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeUserSettingsApi.MeUpdateSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMeUpdateSettingsRequest struct via the builder pattern


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

