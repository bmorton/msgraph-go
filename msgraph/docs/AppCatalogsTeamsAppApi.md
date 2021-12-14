# \AppCatalogsTeamsAppApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppCatalogsCreateTeamsApps**](AppCatalogsTeamsAppApi.md#AppCatalogsCreateTeamsApps) | **Post** /appCatalogs/teamsApps | Create new navigation property to teamsApps for appCatalogs
[**AppCatalogsDeleteTeamsApps**](AppCatalogsTeamsAppApi.md#AppCatalogsDeleteTeamsApps) | **Delete** /appCatalogs/teamsApps/{teamsApp-id} | Delete navigation property teamsApps for appCatalogs
[**AppCatalogsGetTeamsApps**](AppCatalogsTeamsAppApi.md#AppCatalogsGetTeamsApps) | **Get** /appCatalogs/teamsApps/{teamsApp-id} | Get teamsApps from appCatalogs
[**AppCatalogsListTeamsApps**](AppCatalogsTeamsAppApi.md#AppCatalogsListTeamsApps) | **Get** /appCatalogs/teamsApps | Get teamsApps from appCatalogs
[**AppCatalogsTeamsAppsAppDefinitionsDeleteBot**](AppCatalogsTeamsAppApi.md#AppCatalogsTeamsAppsAppDefinitionsDeleteBot) | **Delete** /appCatalogs/teamsApps/{teamsApp-id}/appDefinitions/{teamsAppDefinition-id}/bot | Delete navigation property bot for appCatalogs
[**AppCatalogsTeamsAppsAppDefinitionsGetBot**](AppCatalogsTeamsAppApi.md#AppCatalogsTeamsAppsAppDefinitionsGetBot) | **Get** /appCatalogs/teamsApps/{teamsApp-id}/appDefinitions/{teamsAppDefinition-id}/bot | Get bot from appCatalogs
[**AppCatalogsTeamsAppsAppDefinitionsUpdateBot**](AppCatalogsTeamsAppApi.md#AppCatalogsTeamsAppsAppDefinitionsUpdateBot) | **Patch** /appCatalogs/teamsApps/{teamsApp-id}/appDefinitions/{teamsAppDefinition-id}/bot | Update the navigation property bot in appCatalogs
[**AppCatalogsTeamsAppsCreateAppDefinitions**](AppCatalogsTeamsAppApi.md#AppCatalogsTeamsAppsCreateAppDefinitions) | **Post** /appCatalogs/teamsApps/{teamsApp-id}/appDefinitions | Create new navigation property to appDefinitions for appCatalogs
[**AppCatalogsTeamsAppsDeleteAppDefinitions**](AppCatalogsTeamsAppApi.md#AppCatalogsTeamsAppsDeleteAppDefinitions) | **Delete** /appCatalogs/teamsApps/{teamsApp-id}/appDefinitions/{teamsAppDefinition-id} | Delete navigation property appDefinitions for appCatalogs
[**AppCatalogsTeamsAppsGetAppDefinitions**](AppCatalogsTeamsAppApi.md#AppCatalogsTeamsAppsGetAppDefinitions) | **Get** /appCatalogs/teamsApps/{teamsApp-id}/appDefinitions/{teamsAppDefinition-id} | Get appDefinitions from appCatalogs
[**AppCatalogsTeamsAppsListAppDefinitions**](AppCatalogsTeamsAppApi.md#AppCatalogsTeamsAppsListAppDefinitions) | **Get** /appCatalogs/teamsApps/{teamsApp-id}/appDefinitions | Get appDefinitions from appCatalogs
[**AppCatalogsTeamsAppsUpdateAppDefinitions**](AppCatalogsTeamsAppApi.md#AppCatalogsTeamsAppsUpdateAppDefinitions) | **Patch** /appCatalogs/teamsApps/{teamsApp-id}/appDefinitions/{teamsAppDefinition-id} | Update the navigation property appDefinitions in appCatalogs
[**AppCatalogsUpdateTeamsApps**](AppCatalogsTeamsAppApi.md#AppCatalogsUpdateTeamsApps) | **Patch** /appCatalogs/teamsApps/{teamsApp-id} | Update the navigation property teamsApps in appCatalogs



## AppCatalogsCreateTeamsApps

> MicrosoftGraphTeamsApp AppCatalogsCreateTeamsApps(ctx).MicrosoftGraphTeamsApp(microsoftGraphTeamsApp).Execute()

Create new navigation property to teamsApps for appCatalogs

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
    microsoftGraphTeamsApp := *openapiclient.NewMicrosoftGraphTeamsApp() // MicrosoftGraphTeamsApp | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AppCatalogsTeamsAppApi.AppCatalogsCreateTeamsApps(context.Background()).MicrosoftGraphTeamsApp(microsoftGraphTeamsApp).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppCatalogsTeamsAppApi.AppCatalogsCreateTeamsApps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppCatalogsCreateTeamsApps`: MicrosoftGraphTeamsApp
    fmt.Fprintf(os.Stdout, "Response from `AppCatalogsTeamsAppApi.AppCatalogsCreateTeamsApps`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppCatalogsCreateTeamsAppsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphTeamsApp** | [**MicrosoftGraphTeamsApp**](MicrosoftGraphTeamsApp.md) | New navigation property | 

### Return type

[**MicrosoftGraphTeamsApp**](MicrosoftGraphTeamsApp.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppCatalogsDeleteTeamsApps

> AppCatalogsDeleteTeamsApps(ctx, teamsAppId).IfMatch(ifMatch).Execute()

Delete navigation property teamsApps for appCatalogs

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
    teamsAppId := "teamsAppId_example" // string | key: id of teamsApp
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AppCatalogsTeamsAppApi.AppCatalogsDeleteTeamsApps(context.Background(), teamsAppId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppCatalogsTeamsAppApi.AppCatalogsDeleteTeamsApps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamsAppId** | **string** | key: id of teamsApp | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppCatalogsDeleteTeamsAppsRequest struct via the builder pattern


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


## AppCatalogsGetTeamsApps

> MicrosoftGraphTeamsApp AppCatalogsGetTeamsApps(ctx, teamsAppId).Select_(select_).Expand(expand).Execute()

Get teamsApps from appCatalogs

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
    teamsAppId := "teamsAppId_example" // string | key: id of teamsApp
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AppCatalogsTeamsAppApi.AppCatalogsGetTeamsApps(context.Background(), teamsAppId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppCatalogsTeamsAppApi.AppCatalogsGetTeamsApps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppCatalogsGetTeamsApps`: MicrosoftGraphTeamsApp
    fmt.Fprintf(os.Stdout, "Response from `AppCatalogsTeamsAppApi.AppCatalogsGetTeamsApps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamsAppId** | **string** | key: id of teamsApp | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppCatalogsGetTeamsAppsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphTeamsApp**](MicrosoftGraphTeamsApp.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppCatalogsListTeamsApps

> CollectionOfTeamsApp AppCatalogsListTeamsApps(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get teamsApps from appCatalogs

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
    resp, r, err := api_client.AppCatalogsTeamsAppApi.AppCatalogsListTeamsApps(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppCatalogsTeamsAppApi.AppCatalogsListTeamsApps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppCatalogsListTeamsApps`: CollectionOfTeamsApp
    fmt.Fprintf(os.Stdout, "Response from `AppCatalogsTeamsAppApi.AppCatalogsListTeamsApps`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppCatalogsListTeamsAppsRequest struct via the builder pattern


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

[**CollectionOfTeamsApp**](CollectionOfTeamsApp.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppCatalogsTeamsAppsAppDefinitionsDeleteBot

> AppCatalogsTeamsAppsAppDefinitionsDeleteBot(ctx, teamsAppId, teamsAppDefinitionId).IfMatch(ifMatch).Execute()

Delete navigation property bot for appCatalogs



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
    teamsAppId := "teamsAppId_example" // string | key: id of teamsApp
    teamsAppDefinitionId := "teamsAppDefinitionId_example" // string | key: id of teamsAppDefinition
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AppCatalogsTeamsAppApi.AppCatalogsTeamsAppsAppDefinitionsDeleteBot(context.Background(), teamsAppId, teamsAppDefinitionId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppCatalogsTeamsAppApi.AppCatalogsTeamsAppsAppDefinitionsDeleteBot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamsAppId** | **string** | key: id of teamsApp | 
**teamsAppDefinitionId** | **string** | key: id of teamsAppDefinition | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppCatalogsTeamsAppsAppDefinitionsDeleteBotRequest struct via the builder pattern


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


## AppCatalogsTeamsAppsAppDefinitionsGetBot

> MicrosoftGraphTeamworkBot AppCatalogsTeamsAppsAppDefinitionsGetBot(ctx, teamsAppId, teamsAppDefinitionId).Select_(select_).Expand(expand).Execute()

Get bot from appCatalogs



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
    teamsAppId := "teamsAppId_example" // string | key: id of teamsApp
    teamsAppDefinitionId := "teamsAppDefinitionId_example" // string | key: id of teamsAppDefinition
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AppCatalogsTeamsAppApi.AppCatalogsTeamsAppsAppDefinitionsGetBot(context.Background(), teamsAppId, teamsAppDefinitionId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppCatalogsTeamsAppApi.AppCatalogsTeamsAppsAppDefinitionsGetBot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppCatalogsTeamsAppsAppDefinitionsGetBot`: MicrosoftGraphTeamworkBot
    fmt.Fprintf(os.Stdout, "Response from `AppCatalogsTeamsAppApi.AppCatalogsTeamsAppsAppDefinitionsGetBot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamsAppId** | **string** | key: id of teamsApp | 
**teamsAppDefinitionId** | **string** | key: id of teamsAppDefinition | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppCatalogsTeamsAppsAppDefinitionsGetBotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphTeamworkBot**](MicrosoftGraphTeamworkBot.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppCatalogsTeamsAppsAppDefinitionsUpdateBot

> AppCatalogsTeamsAppsAppDefinitionsUpdateBot(ctx, teamsAppId, teamsAppDefinitionId).MicrosoftGraphTeamworkBot(microsoftGraphTeamworkBot).Execute()

Update the navigation property bot in appCatalogs



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
    teamsAppId := "teamsAppId_example" // string | key: id of teamsApp
    teamsAppDefinitionId := "teamsAppDefinitionId_example" // string | key: id of teamsAppDefinition
    microsoftGraphTeamworkBot := *openapiclient.NewMicrosoftGraphTeamworkBot() // MicrosoftGraphTeamworkBot | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AppCatalogsTeamsAppApi.AppCatalogsTeamsAppsAppDefinitionsUpdateBot(context.Background(), teamsAppId, teamsAppDefinitionId).MicrosoftGraphTeamworkBot(microsoftGraphTeamworkBot).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppCatalogsTeamsAppApi.AppCatalogsTeamsAppsAppDefinitionsUpdateBot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamsAppId** | **string** | key: id of teamsApp | 
**teamsAppDefinitionId** | **string** | key: id of teamsAppDefinition | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppCatalogsTeamsAppsAppDefinitionsUpdateBotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **microsoftGraphTeamworkBot** | [**MicrosoftGraphTeamworkBot**](MicrosoftGraphTeamworkBot.md) | New navigation property values | 

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


## AppCatalogsTeamsAppsCreateAppDefinitions

> MicrosoftGraphTeamsAppDefinition AppCatalogsTeamsAppsCreateAppDefinitions(ctx, teamsAppId).MicrosoftGraphTeamsAppDefinition(microsoftGraphTeamsAppDefinition).Execute()

Create new navigation property to appDefinitions for appCatalogs



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
    teamsAppId := "teamsAppId_example" // string | key: id of teamsApp
    microsoftGraphTeamsAppDefinition := *openapiclient.NewMicrosoftGraphTeamsAppDefinition() // MicrosoftGraphTeamsAppDefinition | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AppCatalogsTeamsAppApi.AppCatalogsTeamsAppsCreateAppDefinitions(context.Background(), teamsAppId).MicrosoftGraphTeamsAppDefinition(microsoftGraphTeamsAppDefinition).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppCatalogsTeamsAppApi.AppCatalogsTeamsAppsCreateAppDefinitions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppCatalogsTeamsAppsCreateAppDefinitions`: MicrosoftGraphTeamsAppDefinition
    fmt.Fprintf(os.Stdout, "Response from `AppCatalogsTeamsAppApi.AppCatalogsTeamsAppsCreateAppDefinitions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamsAppId** | **string** | key: id of teamsApp | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppCatalogsTeamsAppsCreateAppDefinitionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphTeamsAppDefinition** | [**MicrosoftGraphTeamsAppDefinition**](MicrosoftGraphTeamsAppDefinition.md) | New navigation property | 

### Return type

[**MicrosoftGraphTeamsAppDefinition**](MicrosoftGraphTeamsAppDefinition.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppCatalogsTeamsAppsDeleteAppDefinitions

> AppCatalogsTeamsAppsDeleteAppDefinitions(ctx, teamsAppId, teamsAppDefinitionId).IfMatch(ifMatch).Execute()

Delete navigation property appDefinitions for appCatalogs



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
    teamsAppId := "teamsAppId_example" // string | key: id of teamsApp
    teamsAppDefinitionId := "teamsAppDefinitionId_example" // string | key: id of teamsAppDefinition
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AppCatalogsTeamsAppApi.AppCatalogsTeamsAppsDeleteAppDefinitions(context.Background(), teamsAppId, teamsAppDefinitionId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppCatalogsTeamsAppApi.AppCatalogsTeamsAppsDeleteAppDefinitions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamsAppId** | **string** | key: id of teamsApp | 
**teamsAppDefinitionId** | **string** | key: id of teamsAppDefinition | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppCatalogsTeamsAppsDeleteAppDefinitionsRequest struct via the builder pattern


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


## AppCatalogsTeamsAppsGetAppDefinitions

> MicrosoftGraphTeamsAppDefinition AppCatalogsTeamsAppsGetAppDefinitions(ctx, teamsAppId, teamsAppDefinitionId).Select_(select_).Expand(expand).Execute()

Get appDefinitions from appCatalogs



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
    teamsAppId := "teamsAppId_example" // string | key: id of teamsApp
    teamsAppDefinitionId := "teamsAppDefinitionId_example" // string | key: id of teamsAppDefinition
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AppCatalogsTeamsAppApi.AppCatalogsTeamsAppsGetAppDefinitions(context.Background(), teamsAppId, teamsAppDefinitionId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppCatalogsTeamsAppApi.AppCatalogsTeamsAppsGetAppDefinitions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppCatalogsTeamsAppsGetAppDefinitions`: MicrosoftGraphTeamsAppDefinition
    fmt.Fprintf(os.Stdout, "Response from `AppCatalogsTeamsAppApi.AppCatalogsTeamsAppsGetAppDefinitions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamsAppId** | **string** | key: id of teamsApp | 
**teamsAppDefinitionId** | **string** | key: id of teamsAppDefinition | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppCatalogsTeamsAppsGetAppDefinitionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphTeamsAppDefinition**](MicrosoftGraphTeamsAppDefinition.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppCatalogsTeamsAppsListAppDefinitions

> CollectionOfTeamsAppDefinition AppCatalogsTeamsAppsListAppDefinitions(ctx, teamsAppId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get appDefinitions from appCatalogs



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
    teamsAppId := "teamsAppId_example" // string | key: id of teamsApp
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
    resp, r, err := api_client.AppCatalogsTeamsAppApi.AppCatalogsTeamsAppsListAppDefinitions(context.Background(), teamsAppId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppCatalogsTeamsAppApi.AppCatalogsTeamsAppsListAppDefinitions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppCatalogsTeamsAppsListAppDefinitions`: CollectionOfTeamsAppDefinition
    fmt.Fprintf(os.Stdout, "Response from `AppCatalogsTeamsAppApi.AppCatalogsTeamsAppsListAppDefinitions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamsAppId** | **string** | key: id of teamsApp | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppCatalogsTeamsAppsListAppDefinitionsRequest struct via the builder pattern


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

[**CollectionOfTeamsAppDefinition**](CollectionOfTeamsAppDefinition.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppCatalogsTeamsAppsUpdateAppDefinitions

> AppCatalogsTeamsAppsUpdateAppDefinitions(ctx, teamsAppId, teamsAppDefinitionId).MicrosoftGraphTeamsAppDefinition(microsoftGraphTeamsAppDefinition).Execute()

Update the navigation property appDefinitions in appCatalogs



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
    teamsAppId := "teamsAppId_example" // string | key: id of teamsApp
    teamsAppDefinitionId := "teamsAppDefinitionId_example" // string | key: id of teamsAppDefinition
    microsoftGraphTeamsAppDefinition := *openapiclient.NewMicrosoftGraphTeamsAppDefinition() // MicrosoftGraphTeamsAppDefinition | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AppCatalogsTeamsAppApi.AppCatalogsTeamsAppsUpdateAppDefinitions(context.Background(), teamsAppId, teamsAppDefinitionId).MicrosoftGraphTeamsAppDefinition(microsoftGraphTeamsAppDefinition).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppCatalogsTeamsAppApi.AppCatalogsTeamsAppsUpdateAppDefinitions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamsAppId** | **string** | key: id of teamsApp | 
**teamsAppDefinitionId** | **string** | key: id of teamsAppDefinition | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppCatalogsTeamsAppsUpdateAppDefinitionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **microsoftGraphTeamsAppDefinition** | [**MicrosoftGraphTeamsAppDefinition**](MicrosoftGraphTeamsAppDefinition.md) | New navigation property values | 

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


## AppCatalogsUpdateTeamsApps

> AppCatalogsUpdateTeamsApps(ctx, teamsAppId).MicrosoftGraphTeamsApp(microsoftGraphTeamsApp).Execute()

Update the navigation property teamsApps in appCatalogs

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
    teamsAppId := "teamsAppId_example" // string | key: id of teamsApp
    microsoftGraphTeamsApp := *openapiclient.NewMicrosoftGraphTeamsApp() // MicrosoftGraphTeamsApp | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AppCatalogsTeamsAppApi.AppCatalogsUpdateTeamsApps(context.Background(), teamsAppId).MicrosoftGraphTeamsApp(microsoftGraphTeamsApp).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppCatalogsTeamsAppApi.AppCatalogsUpdateTeamsApps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamsAppId** | **string** | key: id of teamsApp | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppCatalogsUpdateTeamsAppsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphTeamsApp** | [**MicrosoftGraphTeamsApp**](MicrosoftGraphTeamsApp.md) | New navigation property values | 

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

