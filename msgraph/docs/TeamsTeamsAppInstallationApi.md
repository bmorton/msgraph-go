# \TeamsTeamsAppInstallationApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TeamsCreateInstalledApps**](TeamsTeamsAppInstallationApi.md#TeamsCreateInstalledApps) | **Post** /teams/{team-id}/installedApps | Create new navigation property to installedApps for teams
[**TeamsDeleteInstalledApps**](TeamsTeamsAppInstallationApi.md#TeamsDeleteInstalledApps) | **Delete** /teams/{team-id}/installedApps/{teamsAppInstallation-id} | Delete navigation property installedApps for teams
[**TeamsGetInstalledApps**](TeamsTeamsAppInstallationApi.md#TeamsGetInstalledApps) | **Get** /teams/{team-id}/installedApps/{teamsAppInstallation-id} | Get installedApps from teams
[**TeamsInstalledAppsDeleteRefTeamsApp**](TeamsTeamsAppInstallationApi.md#TeamsInstalledAppsDeleteRefTeamsApp) | **Delete** /teams/{team-id}/installedApps/{teamsAppInstallation-id}/teamsApp/$ref | Delete ref of navigation property teamsApp for teams
[**TeamsInstalledAppsDeleteRefTeamsAppDefinition**](TeamsTeamsAppInstallationApi.md#TeamsInstalledAppsDeleteRefTeamsAppDefinition) | **Delete** /teams/{team-id}/installedApps/{teamsAppInstallation-id}/teamsAppDefinition/$ref | Delete ref of navigation property teamsAppDefinition for teams
[**TeamsInstalledAppsGetRefTeamsApp**](TeamsTeamsAppInstallationApi.md#TeamsInstalledAppsGetRefTeamsApp) | **Get** /teams/{team-id}/installedApps/{teamsAppInstallation-id}/teamsApp/$ref | Get ref of teamsApp from teams
[**TeamsInstalledAppsGetRefTeamsAppDefinition**](TeamsTeamsAppInstallationApi.md#TeamsInstalledAppsGetRefTeamsAppDefinition) | **Get** /teams/{team-id}/installedApps/{teamsAppInstallation-id}/teamsAppDefinition/$ref | Get ref of teamsAppDefinition from teams
[**TeamsInstalledAppsGetTeamsApp**](TeamsTeamsAppInstallationApi.md#TeamsInstalledAppsGetTeamsApp) | **Get** /teams/{team-id}/installedApps/{teamsAppInstallation-id}/teamsApp | Get teamsApp from teams
[**TeamsInstalledAppsGetTeamsAppDefinition**](TeamsTeamsAppInstallationApi.md#TeamsInstalledAppsGetTeamsAppDefinition) | **Get** /teams/{team-id}/installedApps/{teamsAppInstallation-id}/teamsAppDefinition | Get teamsAppDefinition from teams
[**TeamsInstalledAppsUpdateRefTeamsApp**](TeamsTeamsAppInstallationApi.md#TeamsInstalledAppsUpdateRefTeamsApp) | **Put** /teams/{team-id}/installedApps/{teamsAppInstallation-id}/teamsApp/$ref | Update the ref of navigation property teamsApp in teams
[**TeamsInstalledAppsUpdateRefTeamsAppDefinition**](TeamsTeamsAppInstallationApi.md#TeamsInstalledAppsUpdateRefTeamsAppDefinition) | **Put** /teams/{team-id}/installedApps/{teamsAppInstallation-id}/teamsAppDefinition/$ref | Update the ref of navigation property teamsAppDefinition in teams
[**TeamsListInstalledApps**](TeamsTeamsAppInstallationApi.md#TeamsListInstalledApps) | **Get** /teams/{team-id}/installedApps | Get installedApps from teams
[**TeamsUpdateInstalledApps**](TeamsTeamsAppInstallationApi.md#TeamsUpdateInstalledApps) | **Patch** /teams/{team-id}/installedApps/{teamsAppInstallation-id} | Update the navigation property installedApps in teams



## TeamsCreateInstalledApps

> MicrosoftGraphTeamsAppInstallation TeamsCreateInstalledApps(ctx, teamId).MicrosoftGraphTeamsAppInstallation(microsoftGraphTeamsAppInstallation).Execute()

Create new navigation property to installedApps for teams



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
    teamId := "teamId_example" // string | key: id of team
    microsoftGraphTeamsAppInstallation := *openapiclient.NewMicrosoftGraphTeamsAppInstallation() // MicrosoftGraphTeamsAppInstallation | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsTeamsAppInstallationApi.TeamsCreateInstalledApps(context.Background(), teamId).MicrosoftGraphTeamsAppInstallation(microsoftGraphTeamsAppInstallation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsTeamsAppInstallationApi.TeamsCreateInstalledApps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsCreateInstalledApps`: MicrosoftGraphTeamsAppInstallation
    fmt.Fprintf(os.Stdout, "Response from `TeamsTeamsAppInstallationApi.TeamsCreateInstalledApps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** | key: id of team | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsCreateInstalledAppsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphTeamsAppInstallation** | [**MicrosoftGraphTeamsAppInstallation**](MicrosoftGraphTeamsAppInstallation.md) | New navigation property | 

### Return type

[**MicrosoftGraphTeamsAppInstallation**](MicrosoftGraphTeamsAppInstallation.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsDeleteInstalledApps

> TeamsDeleteInstalledApps(ctx, teamId, teamsAppInstallationId).IfMatch(ifMatch).Execute()

Delete navigation property installedApps for teams



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
    teamId := "teamId_example" // string | key: id of team
    teamsAppInstallationId := "teamsAppInstallationId_example" // string | key: id of teamsAppInstallation
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsTeamsAppInstallationApi.TeamsDeleteInstalledApps(context.Background(), teamId, teamsAppInstallationId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsTeamsAppInstallationApi.TeamsDeleteInstalledApps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** | key: id of team | 
**teamsAppInstallationId** | **string** | key: id of teamsAppInstallation | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsDeleteInstalledAppsRequest struct via the builder pattern


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


## TeamsGetInstalledApps

> MicrosoftGraphTeamsAppInstallation TeamsGetInstalledApps(ctx, teamId, teamsAppInstallationId).Select_(select_).Expand(expand).Execute()

Get installedApps from teams



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
    teamId := "teamId_example" // string | key: id of team
    teamsAppInstallationId := "teamsAppInstallationId_example" // string | key: id of teamsAppInstallation
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsTeamsAppInstallationApi.TeamsGetInstalledApps(context.Background(), teamId, teamsAppInstallationId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsTeamsAppInstallationApi.TeamsGetInstalledApps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsGetInstalledApps`: MicrosoftGraphTeamsAppInstallation
    fmt.Fprintf(os.Stdout, "Response from `TeamsTeamsAppInstallationApi.TeamsGetInstalledApps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** | key: id of team | 
**teamsAppInstallationId** | **string** | key: id of teamsAppInstallation | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsGetInstalledAppsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphTeamsAppInstallation**](MicrosoftGraphTeamsAppInstallation.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsInstalledAppsDeleteRefTeamsApp

> TeamsInstalledAppsDeleteRefTeamsApp(ctx, teamId, teamsAppInstallationId).IfMatch(ifMatch).Execute()

Delete ref of navigation property teamsApp for teams



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
    teamId := "teamId_example" // string | key: id of team
    teamsAppInstallationId := "teamsAppInstallationId_example" // string | key: id of teamsAppInstallation
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsTeamsAppInstallationApi.TeamsInstalledAppsDeleteRefTeamsApp(context.Background(), teamId, teamsAppInstallationId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsTeamsAppInstallationApi.TeamsInstalledAppsDeleteRefTeamsApp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** | key: id of team | 
**teamsAppInstallationId** | **string** | key: id of teamsAppInstallation | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsInstalledAppsDeleteRefTeamsAppRequest struct via the builder pattern


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


## TeamsInstalledAppsDeleteRefTeamsAppDefinition

> TeamsInstalledAppsDeleteRefTeamsAppDefinition(ctx, teamId, teamsAppInstallationId).IfMatch(ifMatch).Execute()

Delete ref of navigation property teamsAppDefinition for teams



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
    teamId := "teamId_example" // string | key: id of team
    teamsAppInstallationId := "teamsAppInstallationId_example" // string | key: id of teamsAppInstallation
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsTeamsAppInstallationApi.TeamsInstalledAppsDeleteRefTeamsAppDefinition(context.Background(), teamId, teamsAppInstallationId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsTeamsAppInstallationApi.TeamsInstalledAppsDeleteRefTeamsAppDefinition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** | key: id of team | 
**teamsAppInstallationId** | **string** | key: id of teamsAppInstallation | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsInstalledAppsDeleteRefTeamsAppDefinitionRequest struct via the builder pattern


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


## TeamsInstalledAppsGetRefTeamsApp

> string TeamsInstalledAppsGetRefTeamsApp(ctx, teamId, teamsAppInstallationId).Execute()

Get ref of teamsApp from teams



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
    teamId := "teamId_example" // string | key: id of team
    teamsAppInstallationId := "teamsAppInstallationId_example" // string | key: id of teamsAppInstallation

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsTeamsAppInstallationApi.TeamsInstalledAppsGetRefTeamsApp(context.Background(), teamId, teamsAppInstallationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsTeamsAppInstallationApi.TeamsInstalledAppsGetRefTeamsApp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsInstalledAppsGetRefTeamsApp`: string
    fmt.Fprintf(os.Stdout, "Response from `TeamsTeamsAppInstallationApi.TeamsInstalledAppsGetRefTeamsApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** | key: id of team | 
**teamsAppInstallationId** | **string** | key: id of teamsAppInstallation | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsInstalledAppsGetRefTeamsAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**string**

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsInstalledAppsGetRefTeamsAppDefinition

> string TeamsInstalledAppsGetRefTeamsAppDefinition(ctx, teamId, teamsAppInstallationId).Execute()

Get ref of teamsAppDefinition from teams



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
    teamId := "teamId_example" // string | key: id of team
    teamsAppInstallationId := "teamsAppInstallationId_example" // string | key: id of teamsAppInstallation

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsTeamsAppInstallationApi.TeamsInstalledAppsGetRefTeamsAppDefinition(context.Background(), teamId, teamsAppInstallationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsTeamsAppInstallationApi.TeamsInstalledAppsGetRefTeamsAppDefinition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsInstalledAppsGetRefTeamsAppDefinition`: string
    fmt.Fprintf(os.Stdout, "Response from `TeamsTeamsAppInstallationApi.TeamsInstalledAppsGetRefTeamsAppDefinition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** | key: id of team | 
**teamsAppInstallationId** | **string** | key: id of teamsAppInstallation | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsInstalledAppsGetRefTeamsAppDefinitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**string**

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsInstalledAppsGetTeamsApp

> MicrosoftGraphTeamsApp TeamsInstalledAppsGetTeamsApp(ctx, teamId, teamsAppInstallationId).Select_(select_).Expand(expand).Execute()

Get teamsApp from teams



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
    teamId := "teamId_example" // string | key: id of team
    teamsAppInstallationId := "teamsAppInstallationId_example" // string | key: id of teamsAppInstallation
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsTeamsAppInstallationApi.TeamsInstalledAppsGetTeamsApp(context.Background(), teamId, teamsAppInstallationId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsTeamsAppInstallationApi.TeamsInstalledAppsGetTeamsApp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsInstalledAppsGetTeamsApp`: MicrosoftGraphTeamsApp
    fmt.Fprintf(os.Stdout, "Response from `TeamsTeamsAppInstallationApi.TeamsInstalledAppsGetTeamsApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** | key: id of team | 
**teamsAppInstallationId** | **string** | key: id of teamsAppInstallation | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsInstalledAppsGetTeamsAppRequest struct via the builder pattern


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


## TeamsInstalledAppsGetTeamsAppDefinition

> MicrosoftGraphTeamsAppDefinition TeamsInstalledAppsGetTeamsAppDefinition(ctx, teamId, teamsAppInstallationId).Select_(select_).Expand(expand).Execute()

Get teamsAppDefinition from teams



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
    teamId := "teamId_example" // string | key: id of team
    teamsAppInstallationId := "teamsAppInstallationId_example" // string | key: id of teamsAppInstallation
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsTeamsAppInstallationApi.TeamsInstalledAppsGetTeamsAppDefinition(context.Background(), teamId, teamsAppInstallationId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsTeamsAppInstallationApi.TeamsInstalledAppsGetTeamsAppDefinition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsInstalledAppsGetTeamsAppDefinition`: MicrosoftGraphTeamsAppDefinition
    fmt.Fprintf(os.Stdout, "Response from `TeamsTeamsAppInstallationApi.TeamsInstalledAppsGetTeamsAppDefinition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** | key: id of team | 
**teamsAppInstallationId** | **string** | key: id of teamsAppInstallation | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsInstalledAppsGetTeamsAppDefinitionRequest struct via the builder pattern


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


## TeamsInstalledAppsUpdateRefTeamsApp

> TeamsInstalledAppsUpdateRefTeamsApp(ctx, teamId, teamsAppInstallationId).RequestBody(requestBody).Execute()

Update the ref of navigation property teamsApp in teams



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
    teamId := "teamId_example" // string | key: id of team
    teamsAppInstallationId := "teamsAppInstallationId_example" // string | key: id of teamsAppInstallation
    requestBody := map[string]map[string]interface{}{"key": map[string]interface{}(123)} // map[string]map[string]interface{} | New navigation property ref values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsTeamsAppInstallationApi.TeamsInstalledAppsUpdateRefTeamsApp(context.Background(), teamId, teamsAppInstallationId).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsTeamsAppInstallationApi.TeamsInstalledAppsUpdateRefTeamsApp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** | key: id of team | 
**teamsAppInstallationId** | **string** | key: id of teamsAppInstallation | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsInstalledAppsUpdateRefTeamsAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **requestBody** | **map[string]map[string]interface{}** | New navigation property ref values | 

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


## TeamsInstalledAppsUpdateRefTeamsAppDefinition

> TeamsInstalledAppsUpdateRefTeamsAppDefinition(ctx, teamId, teamsAppInstallationId).RequestBody(requestBody).Execute()

Update the ref of navigation property teamsAppDefinition in teams



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
    teamId := "teamId_example" // string | key: id of team
    teamsAppInstallationId := "teamsAppInstallationId_example" // string | key: id of teamsAppInstallation
    requestBody := map[string]map[string]interface{}{"key": map[string]interface{}(123)} // map[string]map[string]interface{} | New navigation property ref values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsTeamsAppInstallationApi.TeamsInstalledAppsUpdateRefTeamsAppDefinition(context.Background(), teamId, teamsAppInstallationId).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsTeamsAppInstallationApi.TeamsInstalledAppsUpdateRefTeamsAppDefinition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** | key: id of team | 
**teamsAppInstallationId** | **string** | key: id of teamsAppInstallation | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsInstalledAppsUpdateRefTeamsAppDefinitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **requestBody** | **map[string]map[string]interface{}** | New navigation property ref values | 

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


## TeamsListInstalledApps

> CollectionOfTeamsAppInstallation TeamsListInstalledApps(ctx, teamId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get installedApps from teams



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
    teamId := "teamId_example" // string | key: id of team
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
    resp, r, err := api_client.TeamsTeamsAppInstallationApi.TeamsListInstalledApps(context.Background(), teamId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsTeamsAppInstallationApi.TeamsListInstalledApps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsListInstalledApps`: CollectionOfTeamsAppInstallation
    fmt.Fprintf(os.Stdout, "Response from `TeamsTeamsAppInstallationApi.TeamsListInstalledApps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** | key: id of team | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsListInstalledAppsRequest struct via the builder pattern


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

[**CollectionOfTeamsAppInstallation**](CollectionOfTeamsAppInstallation.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsUpdateInstalledApps

> TeamsUpdateInstalledApps(ctx, teamId, teamsAppInstallationId).MicrosoftGraphTeamsAppInstallation(microsoftGraphTeamsAppInstallation).Execute()

Update the navigation property installedApps in teams



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
    teamId := "teamId_example" // string | key: id of team
    teamsAppInstallationId := "teamsAppInstallationId_example" // string | key: id of teamsAppInstallation
    microsoftGraphTeamsAppInstallation := *openapiclient.NewMicrosoftGraphTeamsAppInstallation() // MicrosoftGraphTeamsAppInstallation | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsTeamsAppInstallationApi.TeamsUpdateInstalledApps(context.Background(), teamId, teamsAppInstallationId).MicrosoftGraphTeamsAppInstallation(microsoftGraphTeamsAppInstallation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsTeamsAppInstallationApi.TeamsUpdateInstalledApps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** | key: id of team | 
**teamsAppInstallationId** | **string** | key: id of teamsAppInstallation | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsUpdateInstalledAppsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **microsoftGraphTeamsAppInstallation** | [**MicrosoftGraphTeamsAppInstallation**](MicrosoftGraphTeamsAppInstallation.md) | New navigation property values | 

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

