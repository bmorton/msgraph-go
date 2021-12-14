# \TeamsActionsApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TeamsTeamArchive**](TeamsActionsApi.md#TeamsTeamArchive) | **Post** /teams/{team-id}/microsoft.graph.archive | Invoke action archive
[**TeamsTeamChannelsChannelCompleteMigration**](TeamsActionsApi.md#TeamsTeamChannelsChannelCompleteMigration) | **Post** /teams/{team-id}/channels/{channel-id}/microsoft.graph.completeMigration | Invoke action completeMigration
[**TeamsTeamChannelsChannelMembersAdd**](TeamsActionsApi.md#TeamsTeamChannelsChannelMembersAdd) | **Post** /teams/{team-id}/channels/{channel-id}/members/microsoft.graph.add | Invoke action add
[**TeamsTeamChannelsChannelProvisionEmail**](TeamsActionsApi.md#TeamsTeamChannelsChannelProvisionEmail) | **Post** /teams/{team-id}/channels/{channel-id}/microsoft.graph.provisionEmail | Invoke action provisionEmail
[**TeamsTeamChannelsChannelRemoveEmail**](TeamsActionsApi.md#TeamsTeamChannelsChannelRemoveEmail) | **Post** /teams/{team-id}/channels/{channel-id}/microsoft.graph.removeEmail | Invoke action removeEmail
[**TeamsTeamClone**](TeamsActionsApi.md#TeamsTeamClone) | **Post** /teams/{team-id}/microsoft.graph.clone | Invoke action clone
[**TeamsTeamCompleteMigration**](TeamsActionsApi.md#TeamsTeamCompleteMigration) | **Post** /teams/{team-id}/microsoft.graph.completeMigration | Invoke action completeMigration
[**TeamsTeamInstalledAppsTeamsAppInstallationUpgrade**](TeamsActionsApi.md#TeamsTeamInstalledAppsTeamsAppInstallationUpgrade) | **Post** /teams/{team-id}/installedApps/{teamsAppInstallation-id}/microsoft.graph.upgrade | Invoke action upgrade
[**TeamsTeamMembersAdd**](TeamsActionsApi.md#TeamsTeamMembersAdd) | **Post** /teams/{team-id}/members/microsoft.graph.add | Invoke action add
[**TeamsTeamPrimaryChannelCompleteMigration**](TeamsActionsApi.md#TeamsTeamPrimaryChannelCompleteMigration) | **Post** /teams/{team-id}/primaryChannel/microsoft.graph.completeMigration | Invoke action completeMigration
[**TeamsTeamPrimaryChannelMembersAdd**](TeamsActionsApi.md#TeamsTeamPrimaryChannelMembersAdd) | **Post** /teams/{team-id}/primaryChannel/members/microsoft.graph.add | Invoke action add
[**TeamsTeamPrimaryChannelProvisionEmail**](TeamsActionsApi.md#TeamsTeamPrimaryChannelProvisionEmail) | **Post** /teams/{team-id}/primaryChannel/microsoft.graph.provisionEmail | Invoke action provisionEmail
[**TeamsTeamPrimaryChannelRemoveEmail**](TeamsActionsApi.md#TeamsTeamPrimaryChannelRemoveEmail) | **Post** /teams/{team-id}/primaryChannel/microsoft.graph.removeEmail | Invoke action removeEmail
[**TeamsTeamScheduleShare**](TeamsActionsApi.md#TeamsTeamScheduleShare) | **Post** /teams/{team-id}/schedule/microsoft.graph.share | Invoke action share
[**TeamsTeamSendActivityNotification**](TeamsActionsApi.md#TeamsTeamSendActivityNotification) | **Post** /teams/{team-id}/microsoft.graph.sendActivityNotification | Invoke action sendActivityNotification
[**TeamsTeamUnarchive**](TeamsActionsApi.md#TeamsTeamUnarchive) | **Post** /teams/{team-id}/microsoft.graph.unarchive | Invoke action unarchive



## TeamsTeamArchive

> TeamsTeamArchive(ctx, teamId).InlineObject848(inlineObject848).Execute()

Invoke action archive

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
    inlineObject848 := *openapiclient.NewInlineObject848() // InlineObject848 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsActionsApi.TeamsTeamArchive(context.Background(), teamId).InlineObject848(inlineObject848).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsActionsApi.TeamsTeamArchive``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** | key: id of team | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsTeamArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject848** | [**InlineObject848**](InlineObject848.md) |  | 

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


## TeamsTeamChannelsChannelCompleteMigration

> TeamsTeamChannelsChannelCompleteMigration(ctx, teamId, channelId).Execute()

Invoke action completeMigration

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
    channelId := "channelId_example" // string | key: id of channel

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsActionsApi.TeamsTeamChannelsChannelCompleteMigration(context.Background(), teamId, channelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsActionsApi.TeamsTeamChannelsChannelCompleteMigration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** | key: id of team | 
**channelId** | **string** | key: id of channel | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsTeamChannelsChannelCompleteMigrationRequest struct via the builder pattern


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


## TeamsTeamChannelsChannelMembersAdd

> []*AnyOfmicrosoftGraphActionResultPart TeamsTeamChannelsChannelMembersAdd(ctx, teamId, channelId).InlineObject846(inlineObject846).Execute()

Invoke action add

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
    channelId := "channelId_example" // string | key: id of channel
    inlineObject846 := *openapiclient.NewInlineObject846() // InlineObject846 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsActionsApi.TeamsTeamChannelsChannelMembersAdd(context.Background(), teamId, channelId).InlineObject846(inlineObject846).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsActionsApi.TeamsTeamChannelsChannelMembersAdd``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsTeamChannelsChannelMembersAdd`: []*AnyOfmicrosoftGraphActionResultPart
    fmt.Fprintf(os.Stdout, "Response from `TeamsActionsApi.TeamsTeamChannelsChannelMembersAdd`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** | key: id of team | 
**channelId** | **string** | key: id of channel | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsTeamChannelsChannelMembersAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **inlineObject846** | [**InlineObject846**](InlineObject846.md) |  | 

### Return type

[**[]*AnyOfmicrosoftGraphActionResultPart**](anyOf&lt;microsoft.graph.actionResultPart&gt;.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsTeamChannelsChannelProvisionEmail

> AnyOfmicrosoftGraphProvisionChannelEmailResult TeamsTeamChannelsChannelProvisionEmail(ctx, teamId, channelId).Execute()

Invoke action provisionEmail

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
    channelId := "channelId_example" // string | key: id of channel

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsActionsApi.TeamsTeamChannelsChannelProvisionEmail(context.Background(), teamId, channelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsActionsApi.TeamsTeamChannelsChannelProvisionEmail``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsTeamChannelsChannelProvisionEmail`: AnyOfmicrosoftGraphProvisionChannelEmailResult
    fmt.Fprintf(os.Stdout, "Response from `TeamsActionsApi.TeamsTeamChannelsChannelProvisionEmail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** | key: id of team | 
**channelId** | **string** | key: id of channel | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsTeamChannelsChannelProvisionEmailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AnyOfmicrosoftGraphProvisionChannelEmailResult**](anyOf&lt;microsoft.graph.provisionChannelEmailResult&gt;.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsTeamChannelsChannelRemoveEmail

> TeamsTeamChannelsChannelRemoveEmail(ctx, teamId, channelId).Execute()

Invoke action removeEmail

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
    channelId := "channelId_example" // string | key: id of channel

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsActionsApi.TeamsTeamChannelsChannelRemoveEmail(context.Background(), teamId, channelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsActionsApi.TeamsTeamChannelsChannelRemoveEmail``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** | key: id of team | 
**channelId** | **string** | key: id of channel | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsTeamChannelsChannelRemoveEmailRequest struct via the builder pattern


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


## TeamsTeamClone

> TeamsTeamClone(ctx, teamId).InlineObject849(inlineObject849).Execute()

Invoke action clone

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
    inlineObject849 := *openapiclient.NewInlineObject849() // InlineObject849 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsActionsApi.TeamsTeamClone(context.Background(), teamId).InlineObject849(inlineObject849).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsActionsApi.TeamsTeamClone``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** | key: id of team | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsTeamCloneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject849** | [**InlineObject849**](InlineObject849.md) |  | 

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


## TeamsTeamCompleteMigration

> TeamsTeamCompleteMigration(ctx, teamId).Execute()

Invoke action completeMigration

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsActionsApi.TeamsTeamCompleteMigration(context.Background(), teamId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsActionsApi.TeamsTeamCompleteMigration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** | key: id of team | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsTeamCompleteMigrationRequest struct via the builder pattern


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


## TeamsTeamInstalledAppsTeamsAppInstallationUpgrade

> TeamsTeamInstalledAppsTeamsAppInstallationUpgrade(ctx, teamId, teamsAppInstallationId).Execute()

Invoke action upgrade

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
    resp, r, err := api_client.TeamsActionsApi.TeamsTeamInstalledAppsTeamsAppInstallationUpgrade(context.Background(), teamId, teamsAppInstallationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsActionsApi.TeamsTeamInstalledAppsTeamsAppInstallationUpgrade``: %v\n", err)
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

Other parameters are passed through a pointer to a apiTeamsTeamInstalledAppsTeamsAppInstallationUpgradeRequest struct via the builder pattern


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


## TeamsTeamMembersAdd

> []*AnyOfmicrosoftGraphActionResultPart TeamsTeamMembersAdd(ctx, teamId).InlineObject847(inlineObject847).Execute()

Invoke action add

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
    inlineObject847 := *openapiclient.NewInlineObject847() // InlineObject847 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsActionsApi.TeamsTeamMembersAdd(context.Background(), teamId).InlineObject847(inlineObject847).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsActionsApi.TeamsTeamMembersAdd``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsTeamMembersAdd`: []*AnyOfmicrosoftGraphActionResultPart
    fmt.Fprintf(os.Stdout, "Response from `TeamsActionsApi.TeamsTeamMembersAdd`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** | key: id of team | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsTeamMembersAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject847** | [**InlineObject847**](InlineObject847.md) |  | 

### Return type

[**[]*AnyOfmicrosoftGraphActionResultPart**](anyOf&lt;microsoft.graph.actionResultPart&gt;.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsTeamPrimaryChannelCompleteMigration

> TeamsTeamPrimaryChannelCompleteMigration(ctx, teamId).Execute()

Invoke action completeMigration

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsActionsApi.TeamsTeamPrimaryChannelCompleteMigration(context.Background(), teamId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsActionsApi.TeamsTeamPrimaryChannelCompleteMigration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** | key: id of team | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsTeamPrimaryChannelCompleteMigrationRequest struct via the builder pattern


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


## TeamsTeamPrimaryChannelMembersAdd

> []*AnyOfmicrosoftGraphActionResultPart TeamsTeamPrimaryChannelMembersAdd(ctx, teamId).InlineObject851(inlineObject851).Execute()

Invoke action add

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
    inlineObject851 := *openapiclient.NewInlineObject851() // InlineObject851 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsActionsApi.TeamsTeamPrimaryChannelMembersAdd(context.Background(), teamId).InlineObject851(inlineObject851).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsActionsApi.TeamsTeamPrimaryChannelMembersAdd``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsTeamPrimaryChannelMembersAdd`: []*AnyOfmicrosoftGraphActionResultPart
    fmt.Fprintf(os.Stdout, "Response from `TeamsActionsApi.TeamsTeamPrimaryChannelMembersAdd`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** | key: id of team | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsTeamPrimaryChannelMembersAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject851** | [**InlineObject851**](InlineObject851.md) |  | 

### Return type

[**[]*AnyOfmicrosoftGraphActionResultPart**](anyOf&lt;microsoft.graph.actionResultPart&gt;.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsTeamPrimaryChannelProvisionEmail

> AnyOfmicrosoftGraphProvisionChannelEmailResult TeamsTeamPrimaryChannelProvisionEmail(ctx, teamId).Execute()

Invoke action provisionEmail

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsActionsApi.TeamsTeamPrimaryChannelProvisionEmail(context.Background(), teamId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsActionsApi.TeamsTeamPrimaryChannelProvisionEmail``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsTeamPrimaryChannelProvisionEmail`: AnyOfmicrosoftGraphProvisionChannelEmailResult
    fmt.Fprintf(os.Stdout, "Response from `TeamsActionsApi.TeamsTeamPrimaryChannelProvisionEmail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** | key: id of team | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsTeamPrimaryChannelProvisionEmailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AnyOfmicrosoftGraphProvisionChannelEmailResult**](anyOf&lt;microsoft.graph.provisionChannelEmailResult&gt;.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsTeamPrimaryChannelRemoveEmail

> TeamsTeamPrimaryChannelRemoveEmail(ctx, teamId).Execute()

Invoke action removeEmail

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsActionsApi.TeamsTeamPrimaryChannelRemoveEmail(context.Background(), teamId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsActionsApi.TeamsTeamPrimaryChannelRemoveEmail``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** | key: id of team | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsTeamPrimaryChannelRemoveEmailRequest struct via the builder pattern


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


## TeamsTeamScheduleShare

> TeamsTeamScheduleShare(ctx, teamId).InlineObject852(inlineObject852).Execute()

Invoke action share

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
    inlineObject852 := *openapiclient.NewInlineObject852() // InlineObject852 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsActionsApi.TeamsTeamScheduleShare(context.Background(), teamId).InlineObject852(inlineObject852).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsActionsApi.TeamsTeamScheduleShare``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** | key: id of team | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsTeamScheduleShareRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject852** | [**InlineObject852**](InlineObject852.md) |  | 

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


## TeamsTeamSendActivityNotification

> TeamsTeamSendActivityNotification(ctx, teamId).InlineObject850(inlineObject850).Execute()

Invoke action sendActivityNotification

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
    inlineObject850 := *openapiclient.NewInlineObject850() // InlineObject850 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsActionsApi.TeamsTeamSendActivityNotification(context.Background(), teamId).InlineObject850(inlineObject850).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsActionsApi.TeamsTeamSendActivityNotification``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** | key: id of team | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsTeamSendActivityNotificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject850** | [**InlineObject850**](InlineObject850.md) |  | 

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


## TeamsTeamUnarchive

> TeamsTeamUnarchive(ctx, teamId).Execute()

Invoke action unarchive

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsActionsApi.TeamsTeamUnarchive(context.Background(), teamId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsActionsApi.TeamsTeamUnarchive``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** | key: id of team | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsTeamUnarchiveRequest struct via the builder pattern


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

