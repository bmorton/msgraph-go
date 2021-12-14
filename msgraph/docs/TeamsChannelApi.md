# \TeamsChannelApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TeamsChannelsCreateMembers**](TeamsChannelApi.md#TeamsChannelsCreateMembers) | **Post** /teams/{team-id}/channels/{channel-id}/members | Create new navigation property to members for teams
[**TeamsChannelsCreateMessages**](TeamsChannelApi.md#TeamsChannelsCreateMessages) | **Post** /teams/{team-id}/channels/{channel-id}/messages | Create new navigation property to messages for teams
[**TeamsChannelsCreateTabs**](TeamsChannelApi.md#TeamsChannelsCreateTabs) | **Post** /teams/{team-id}/channels/{channel-id}/tabs | Create new navigation property to tabs for teams
[**TeamsChannelsDeleteFilesFolder**](TeamsChannelApi.md#TeamsChannelsDeleteFilesFolder) | **Delete** /teams/{team-id}/channels/{channel-id}/filesFolder | Delete navigation property filesFolder for teams
[**TeamsChannelsDeleteMembers**](TeamsChannelApi.md#TeamsChannelsDeleteMembers) | **Delete** /teams/{team-id}/channels/{channel-id}/members/{conversationMember-id} | Delete navigation property members for teams
[**TeamsChannelsDeleteMessages**](TeamsChannelApi.md#TeamsChannelsDeleteMessages) | **Delete** /teams/{team-id}/channels/{channel-id}/messages/{chatMessage-id} | Delete navigation property messages for teams
[**TeamsChannelsDeleteTabs**](TeamsChannelApi.md#TeamsChannelsDeleteTabs) | **Delete** /teams/{team-id}/channels/{channel-id}/tabs/{teamsTab-id} | Delete navigation property tabs for teams
[**TeamsChannelsGetFilesFolder**](TeamsChannelApi.md#TeamsChannelsGetFilesFolder) | **Get** /teams/{team-id}/channels/{channel-id}/filesFolder | Get filesFolder from teams
[**TeamsChannelsGetFilesFolderContent**](TeamsChannelApi.md#TeamsChannelsGetFilesFolderContent) | **Get** /teams/{team-id}/channels/{channel-id}/filesFolder/content | Get media content for the navigation property filesFolder from teams
[**TeamsChannelsGetMembers**](TeamsChannelApi.md#TeamsChannelsGetMembers) | **Get** /teams/{team-id}/channels/{channel-id}/members/{conversationMember-id} | Get members from teams
[**TeamsChannelsGetMessages**](TeamsChannelApi.md#TeamsChannelsGetMessages) | **Get** /teams/{team-id}/channels/{channel-id}/messages/{chatMessage-id} | Get messages from teams
[**TeamsChannelsGetTabs**](TeamsChannelApi.md#TeamsChannelsGetTabs) | **Get** /teams/{team-id}/channels/{channel-id}/tabs/{teamsTab-id} | Get tabs from teams
[**TeamsChannelsListMembers**](TeamsChannelApi.md#TeamsChannelsListMembers) | **Get** /teams/{team-id}/channels/{channel-id}/members | Get members from teams
[**TeamsChannelsListMessages**](TeamsChannelApi.md#TeamsChannelsListMessages) | **Get** /teams/{team-id}/channels/{channel-id}/messages | Get messages from teams
[**TeamsChannelsListTabs**](TeamsChannelApi.md#TeamsChannelsListTabs) | **Get** /teams/{team-id}/channels/{channel-id}/tabs | Get tabs from teams
[**TeamsChannelsMessagesCreateHostedContents**](TeamsChannelApi.md#TeamsChannelsMessagesCreateHostedContents) | **Post** /teams/{team-id}/channels/{channel-id}/messages/{chatMessage-id}/hostedContents | Create new navigation property to hostedContents for teams
[**TeamsChannelsMessagesCreateReplies**](TeamsChannelApi.md#TeamsChannelsMessagesCreateReplies) | **Post** /teams/{team-id}/channels/{channel-id}/messages/{chatMessage-id}/replies | Create new navigation property to replies for teams
[**TeamsChannelsMessagesDeleteHostedContents**](TeamsChannelApi.md#TeamsChannelsMessagesDeleteHostedContents) | **Delete** /teams/{team-id}/channels/{channel-id}/messages/{chatMessage-id}/hostedContents/{chatMessageHostedContent-id} | Delete navigation property hostedContents for teams
[**TeamsChannelsMessagesDeleteReplies**](TeamsChannelApi.md#TeamsChannelsMessagesDeleteReplies) | **Delete** /teams/{team-id}/channels/{channel-id}/messages/{chatMessage-id}/replies/{chatMessage-id1} | Delete navigation property replies for teams
[**TeamsChannelsMessagesGetHostedContents**](TeamsChannelApi.md#TeamsChannelsMessagesGetHostedContents) | **Get** /teams/{team-id}/channels/{channel-id}/messages/{chatMessage-id}/hostedContents/{chatMessageHostedContent-id} | Get hostedContents from teams
[**TeamsChannelsMessagesGetReplies**](TeamsChannelApi.md#TeamsChannelsMessagesGetReplies) | **Get** /teams/{team-id}/channels/{channel-id}/messages/{chatMessage-id}/replies/{chatMessage-id1} | Get replies from teams
[**TeamsChannelsMessagesListHostedContents**](TeamsChannelApi.md#TeamsChannelsMessagesListHostedContents) | **Get** /teams/{team-id}/channels/{channel-id}/messages/{chatMessage-id}/hostedContents | Get hostedContents from teams
[**TeamsChannelsMessagesListReplies**](TeamsChannelApi.md#TeamsChannelsMessagesListReplies) | **Get** /teams/{team-id}/channels/{channel-id}/messages/{chatMessage-id}/replies | Get replies from teams
[**TeamsChannelsMessagesUpdateHostedContents**](TeamsChannelApi.md#TeamsChannelsMessagesUpdateHostedContents) | **Patch** /teams/{team-id}/channels/{channel-id}/messages/{chatMessage-id}/hostedContents/{chatMessageHostedContent-id} | Update the navigation property hostedContents in teams
[**TeamsChannelsMessagesUpdateReplies**](TeamsChannelApi.md#TeamsChannelsMessagesUpdateReplies) | **Patch** /teams/{team-id}/channels/{channel-id}/messages/{chatMessage-id}/replies/{chatMessage-id1} | Update the navigation property replies in teams
[**TeamsChannelsTabsDeleteRefTeamsApp**](TeamsChannelApi.md#TeamsChannelsTabsDeleteRefTeamsApp) | **Delete** /teams/{team-id}/channels/{channel-id}/tabs/{teamsTab-id}/teamsApp/$ref | Delete ref of navigation property teamsApp for teams
[**TeamsChannelsTabsGetRefTeamsApp**](TeamsChannelApi.md#TeamsChannelsTabsGetRefTeamsApp) | **Get** /teams/{team-id}/channels/{channel-id}/tabs/{teamsTab-id}/teamsApp/$ref | Get ref of teamsApp from teams
[**TeamsChannelsTabsGetTeamsApp**](TeamsChannelApi.md#TeamsChannelsTabsGetTeamsApp) | **Get** /teams/{team-id}/channels/{channel-id}/tabs/{teamsTab-id}/teamsApp | Get teamsApp from teams
[**TeamsChannelsTabsUpdateRefTeamsApp**](TeamsChannelApi.md#TeamsChannelsTabsUpdateRefTeamsApp) | **Put** /teams/{team-id}/channels/{channel-id}/tabs/{teamsTab-id}/teamsApp/$ref | Update the ref of navigation property teamsApp in teams
[**TeamsChannelsUpdateFilesFolder**](TeamsChannelApi.md#TeamsChannelsUpdateFilesFolder) | **Patch** /teams/{team-id}/channels/{channel-id}/filesFolder | Update the navigation property filesFolder in teams
[**TeamsChannelsUpdateFilesFolderContent**](TeamsChannelApi.md#TeamsChannelsUpdateFilesFolderContent) | **Put** /teams/{team-id}/channels/{channel-id}/filesFolder/content | Update media content for the navigation property filesFolder in teams
[**TeamsChannelsUpdateMembers**](TeamsChannelApi.md#TeamsChannelsUpdateMembers) | **Patch** /teams/{team-id}/channels/{channel-id}/members/{conversationMember-id} | Update the navigation property members in teams
[**TeamsChannelsUpdateMessages**](TeamsChannelApi.md#TeamsChannelsUpdateMessages) | **Patch** /teams/{team-id}/channels/{channel-id}/messages/{chatMessage-id} | Update the navigation property messages in teams
[**TeamsChannelsUpdateTabs**](TeamsChannelApi.md#TeamsChannelsUpdateTabs) | **Patch** /teams/{team-id}/channels/{channel-id}/tabs/{teamsTab-id} | Update the navigation property tabs in teams
[**TeamsCreateChannels**](TeamsChannelApi.md#TeamsCreateChannels) | **Post** /teams/{team-id}/channels | Create new navigation property to channels for teams
[**TeamsDeleteChannels**](TeamsChannelApi.md#TeamsDeleteChannels) | **Delete** /teams/{team-id}/channels/{channel-id} | Delete navigation property channels for teams
[**TeamsDeletePrimaryChannel**](TeamsChannelApi.md#TeamsDeletePrimaryChannel) | **Delete** /teams/{team-id}/primaryChannel | Delete navigation property primaryChannel for teams
[**TeamsGetChannels**](TeamsChannelApi.md#TeamsGetChannels) | **Get** /teams/{team-id}/channels/{channel-id} | Get channels from teams
[**TeamsGetPrimaryChannel**](TeamsChannelApi.md#TeamsGetPrimaryChannel) | **Get** /teams/{team-id}/primaryChannel | Get primaryChannel from teams
[**TeamsListChannels**](TeamsChannelApi.md#TeamsListChannels) | **Get** /teams/{team-id}/channels | Get channels from teams
[**TeamsPrimaryChannelCreateMembers**](TeamsChannelApi.md#TeamsPrimaryChannelCreateMembers) | **Post** /teams/{team-id}/primaryChannel/members | Create new navigation property to members for teams
[**TeamsPrimaryChannelCreateMessages**](TeamsChannelApi.md#TeamsPrimaryChannelCreateMessages) | **Post** /teams/{team-id}/primaryChannel/messages | Create new navigation property to messages for teams
[**TeamsPrimaryChannelCreateTabs**](TeamsChannelApi.md#TeamsPrimaryChannelCreateTabs) | **Post** /teams/{team-id}/primaryChannel/tabs | Create new navigation property to tabs for teams
[**TeamsPrimaryChannelDeleteFilesFolder**](TeamsChannelApi.md#TeamsPrimaryChannelDeleteFilesFolder) | **Delete** /teams/{team-id}/primaryChannel/filesFolder | Delete navigation property filesFolder for teams
[**TeamsPrimaryChannelDeleteMembers**](TeamsChannelApi.md#TeamsPrimaryChannelDeleteMembers) | **Delete** /teams/{team-id}/primaryChannel/members/{conversationMember-id} | Delete navigation property members for teams
[**TeamsPrimaryChannelDeleteMessages**](TeamsChannelApi.md#TeamsPrimaryChannelDeleteMessages) | **Delete** /teams/{team-id}/primaryChannel/messages/{chatMessage-id} | Delete navigation property messages for teams
[**TeamsPrimaryChannelDeleteTabs**](TeamsChannelApi.md#TeamsPrimaryChannelDeleteTabs) | **Delete** /teams/{team-id}/primaryChannel/tabs/{teamsTab-id} | Delete navigation property tabs for teams
[**TeamsPrimaryChannelGetFilesFolder**](TeamsChannelApi.md#TeamsPrimaryChannelGetFilesFolder) | **Get** /teams/{team-id}/primaryChannel/filesFolder | Get filesFolder from teams
[**TeamsPrimaryChannelGetFilesFolderContent**](TeamsChannelApi.md#TeamsPrimaryChannelGetFilesFolderContent) | **Get** /teams/{team-id}/primaryChannel/filesFolder/content | Get media content for the navigation property filesFolder from teams
[**TeamsPrimaryChannelGetMembers**](TeamsChannelApi.md#TeamsPrimaryChannelGetMembers) | **Get** /teams/{team-id}/primaryChannel/members/{conversationMember-id} | Get members from teams
[**TeamsPrimaryChannelGetMessages**](TeamsChannelApi.md#TeamsPrimaryChannelGetMessages) | **Get** /teams/{team-id}/primaryChannel/messages/{chatMessage-id} | Get messages from teams
[**TeamsPrimaryChannelGetTabs**](TeamsChannelApi.md#TeamsPrimaryChannelGetTabs) | **Get** /teams/{team-id}/primaryChannel/tabs/{teamsTab-id} | Get tabs from teams
[**TeamsPrimaryChannelListMembers**](TeamsChannelApi.md#TeamsPrimaryChannelListMembers) | **Get** /teams/{team-id}/primaryChannel/members | Get members from teams
[**TeamsPrimaryChannelListMessages**](TeamsChannelApi.md#TeamsPrimaryChannelListMessages) | **Get** /teams/{team-id}/primaryChannel/messages | Get messages from teams
[**TeamsPrimaryChannelListTabs**](TeamsChannelApi.md#TeamsPrimaryChannelListTabs) | **Get** /teams/{team-id}/primaryChannel/tabs | Get tabs from teams
[**TeamsPrimaryChannelMessagesCreateHostedContents**](TeamsChannelApi.md#TeamsPrimaryChannelMessagesCreateHostedContents) | **Post** /teams/{team-id}/primaryChannel/messages/{chatMessage-id}/hostedContents | Create new navigation property to hostedContents for teams
[**TeamsPrimaryChannelMessagesCreateReplies**](TeamsChannelApi.md#TeamsPrimaryChannelMessagesCreateReplies) | **Post** /teams/{team-id}/primaryChannel/messages/{chatMessage-id}/replies | Create new navigation property to replies for teams
[**TeamsPrimaryChannelMessagesDeleteHostedContents**](TeamsChannelApi.md#TeamsPrimaryChannelMessagesDeleteHostedContents) | **Delete** /teams/{team-id}/primaryChannel/messages/{chatMessage-id}/hostedContents/{chatMessageHostedContent-id} | Delete navigation property hostedContents for teams
[**TeamsPrimaryChannelMessagesDeleteReplies**](TeamsChannelApi.md#TeamsPrimaryChannelMessagesDeleteReplies) | **Delete** /teams/{team-id}/primaryChannel/messages/{chatMessage-id}/replies/{chatMessage-id1} | Delete navigation property replies for teams
[**TeamsPrimaryChannelMessagesGetHostedContents**](TeamsChannelApi.md#TeamsPrimaryChannelMessagesGetHostedContents) | **Get** /teams/{team-id}/primaryChannel/messages/{chatMessage-id}/hostedContents/{chatMessageHostedContent-id} | Get hostedContents from teams
[**TeamsPrimaryChannelMessagesGetReplies**](TeamsChannelApi.md#TeamsPrimaryChannelMessagesGetReplies) | **Get** /teams/{team-id}/primaryChannel/messages/{chatMessage-id}/replies/{chatMessage-id1} | Get replies from teams
[**TeamsPrimaryChannelMessagesListHostedContents**](TeamsChannelApi.md#TeamsPrimaryChannelMessagesListHostedContents) | **Get** /teams/{team-id}/primaryChannel/messages/{chatMessage-id}/hostedContents | Get hostedContents from teams
[**TeamsPrimaryChannelMessagesListReplies**](TeamsChannelApi.md#TeamsPrimaryChannelMessagesListReplies) | **Get** /teams/{team-id}/primaryChannel/messages/{chatMessage-id}/replies | Get replies from teams
[**TeamsPrimaryChannelMessagesUpdateHostedContents**](TeamsChannelApi.md#TeamsPrimaryChannelMessagesUpdateHostedContents) | **Patch** /teams/{team-id}/primaryChannel/messages/{chatMessage-id}/hostedContents/{chatMessageHostedContent-id} | Update the navigation property hostedContents in teams
[**TeamsPrimaryChannelMessagesUpdateReplies**](TeamsChannelApi.md#TeamsPrimaryChannelMessagesUpdateReplies) | **Patch** /teams/{team-id}/primaryChannel/messages/{chatMessage-id}/replies/{chatMessage-id1} | Update the navigation property replies in teams
[**TeamsPrimaryChannelTabsDeleteRefTeamsApp**](TeamsChannelApi.md#TeamsPrimaryChannelTabsDeleteRefTeamsApp) | **Delete** /teams/{team-id}/primaryChannel/tabs/{teamsTab-id}/teamsApp/$ref | Delete ref of navigation property teamsApp for teams
[**TeamsPrimaryChannelTabsGetRefTeamsApp**](TeamsChannelApi.md#TeamsPrimaryChannelTabsGetRefTeamsApp) | **Get** /teams/{team-id}/primaryChannel/tabs/{teamsTab-id}/teamsApp/$ref | Get ref of teamsApp from teams
[**TeamsPrimaryChannelTabsGetTeamsApp**](TeamsChannelApi.md#TeamsPrimaryChannelTabsGetTeamsApp) | **Get** /teams/{team-id}/primaryChannel/tabs/{teamsTab-id}/teamsApp | Get teamsApp from teams
[**TeamsPrimaryChannelTabsUpdateRefTeamsApp**](TeamsChannelApi.md#TeamsPrimaryChannelTabsUpdateRefTeamsApp) | **Put** /teams/{team-id}/primaryChannel/tabs/{teamsTab-id}/teamsApp/$ref | Update the ref of navigation property teamsApp in teams
[**TeamsPrimaryChannelUpdateFilesFolder**](TeamsChannelApi.md#TeamsPrimaryChannelUpdateFilesFolder) | **Patch** /teams/{team-id}/primaryChannel/filesFolder | Update the navigation property filesFolder in teams
[**TeamsPrimaryChannelUpdateFilesFolderContent**](TeamsChannelApi.md#TeamsPrimaryChannelUpdateFilesFolderContent) | **Put** /teams/{team-id}/primaryChannel/filesFolder/content | Update media content for the navigation property filesFolder in teams
[**TeamsPrimaryChannelUpdateMembers**](TeamsChannelApi.md#TeamsPrimaryChannelUpdateMembers) | **Patch** /teams/{team-id}/primaryChannel/members/{conversationMember-id} | Update the navigation property members in teams
[**TeamsPrimaryChannelUpdateMessages**](TeamsChannelApi.md#TeamsPrimaryChannelUpdateMessages) | **Patch** /teams/{team-id}/primaryChannel/messages/{chatMessage-id} | Update the navigation property messages in teams
[**TeamsPrimaryChannelUpdateTabs**](TeamsChannelApi.md#TeamsPrimaryChannelUpdateTabs) | **Patch** /teams/{team-id}/primaryChannel/tabs/{teamsTab-id} | Update the navigation property tabs in teams
[**TeamsUpdateChannels**](TeamsChannelApi.md#TeamsUpdateChannels) | **Patch** /teams/{team-id}/channels/{channel-id} | Update the navigation property channels in teams
[**TeamsUpdatePrimaryChannel**](TeamsChannelApi.md#TeamsUpdatePrimaryChannel) | **Patch** /teams/{team-id}/primaryChannel | Update the navigation property primaryChannel in teams



## TeamsChannelsCreateMembers

> MicrosoftGraphConversationMember TeamsChannelsCreateMembers(ctx, teamId, channelId).MicrosoftGraphConversationMember(microsoftGraphConversationMember).Execute()

Create new navigation property to members for teams



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
    microsoftGraphConversationMember := *openapiclient.NewMicrosoftGraphConversationMember() // MicrosoftGraphConversationMember | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsChannelApi.TeamsChannelsCreateMembers(context.Background(), teamId, channelId).MicrosoftGraphConversationMember(microsoftGraphConversationMember).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsChannelApi.TeamsChannelsCreateMembers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsChannelsCreateMembers`: MicrosoftGraphConversationMember
    fmt.Fprintf(os.Stdout, "Response from `TeamsChannelApi.TeamsChannelsCreateMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** | key: id of team | 
**channelId** | **string** | key: id of channel | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsChannelsCreateMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **microsoftGraphConversationMember** | [**MicrosoftGraphConversationMember**](MicrosoftGraphConversationMember.md) | New navigation property | 

### Return type

[**MicrosoftGraphConversationMember**](MicrosoftGraphConversationMember.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsChannelsCreateMessages

> MicrosoftGraphChatMessage TeamsChannelsCreateMessages(ctx, teamId, channelId).MicrosoftGraphChatMessage(microsoftGraphChatMessage).Execute()

Create new navigation property to messages for teams



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
    microsoftGraphChatMessage := *openapiclient.NewMicrosoftGraphChatMessage() // MicrosoftGraphChatMessage | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsChannelApi.TeamsChannelsCreateMessages(context.Background(), teamId, channelId).MicrosoftGraphChatMessage(microsoftGraphChatMessage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsChannelApi.TeamsChannelsCreateMessages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsChannelsCreateMessages`: MicrosoftGraphChatMessage
    fmt.Fprintf(os.Stdout, "Response from `TeamsChannelApi.TeamsChannelsCreateMessages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** | key: id of team | 
**channelId** | **string** | key: id of channel | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsChannelsCreateMessagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **microsoftGraphChatMessage** | [**MicrosoftGraphChatMessage**](MicrosoftGraphChatMessage.md) | New navigation property | 

### Return type

[**MicrosoftGraphChatMessage**](MicrosoftGraphChatMessage.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsChannelsCreateTabs

> MicrosoftGraphTeamsTab TeamsChannelsCreateTabs(ctx, teamId, channelId).MicrosoftGraphTeamsTab(microsoftGraphTeamsTab).Execute()

Create new navigation property to tabs for teams



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
    microsoftGraphTeamsTab := *openapiclient.NewMicrosoftGraphTeamsTab() // MicrosoftGraphTeamsTab | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsChannelApi.TeamsChannelsCreateTabs(context.Background(), teamId, channelId).MicrosoftGraphTeamsTab(microsoftGraphTeamsTab).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsChannelApi.TeamsChannelsCreateTabs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsChannelsCreateTabs`: MicrosoftGraphTeamsTab
    fmt.Fprintf(os.Stdout, "Response from `TeamsChannelApi.TeamsChannelsCreateTabs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** | key: id of team | 
**channelId** | **string** | key: id of channel | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsChannelsCreateTabsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **microsoftGraphTeamsTab** | [**MicrosoftGraphTeamsTab**](MicrosoftGraphTeamsTab.md) | New navigation property | 

### Return type

[**MicrosoftGraphTeamsTab**](MicrosoftGraphTeamsTab.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsChannelsDeleteFilesFolder

> TeamsChannelsDeleteFilesFolder(ctx, teamId, channelId).IfMatch(ifMatch).Execute()

Delete navigation property filesFolder for teams



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
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsChannelApi.TeamsChannelsDeleteFilesFolder(context.Background(), teamId, channelId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsChannelApi.TeamsChannelsDeleteFilesFolder``: %v\n", err)
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

Other parameters are passed through a pointer to a apiTeamsChannelsDeleteFilesFolderRequest struct via the builder pattern


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


## TeamsChannelsDeleteMembers

> TeamsChannelsDeleteMembers(ctx, teamId, channelId, conversationMemberId).IfMatch(ifMatch).Execute()

Delete navigation property members for teams



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
    conversationMemberId := "conversationMemberId_example" // string | key: id of conversationMember
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsChannelApi.TeamsChannelsDeleteMembers(context.Background(), teamId, channelId, conversationMemberId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsChannelApi.TeamsChannelsDeleteMembers``: %v\n", err)
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
**conversationMemberId** | **string** | key: id of conversationMember | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsChannelsDeleteMembersRequest struct via the builder pattern


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


## TeamsChannelsDeleteMessages

> TeamsChannelsDeleteMessages(ctx, teamId, channelId, chatMessageId).IfMatch(ifMatch).Execute()

Delete navigation property messages for teams



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
    chatMessageId := "chatMessageId_example" // string | key: id of chatMessage
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsChannelApi.TeamsChannelsDeleteMessages(context.Background(), teamId, channelId, chatMessageId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsChannelApi.TeamsChannelsDeleteMessages``: %v\n", err)
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
**chatMessageId** | **string** | key: id of chatMessage | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsChannelsDeleteMessagesRequest struct via the builder pattern


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


## TeamsChannelsDeleteTabs

> TeamsChannelsDeleteTabs(ctx, teamId, channelId, teamsTabId).IfMatch(ifMatch).Execute()

Delete navigation property tabs for teams



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
    teamsTabId := "teamsTabId_example" // string | key: id of teamsTab
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsChannelApi.TeamsChannelsDeleteTabs(context.Background(), teamId, channelId, teamsTabId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsChannelApi.TeamsChannelsDeleteTabs``: %v\n", err)
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
**teamsTabId** | **string** | key: id of teamsTab | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsChannelsDeleteTabsRequest struct via the builder pattern


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


## TeamsChannelsGetFilesFolder

> MicrosoftGraphDriveItem TeamsChannelsGetFilesFolder(ctx, teamId, channelId).Select_(select_).Expand(expand).Execute()

Get filesFolder from teams



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
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsChannelApi.TeamsChannelsGetFilesFolder(context.Background(), teamId, channelId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsChannelApi.TeamsChannelsGetFilesFolder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsChannelsGetFilesFolder`: MicrosoftGraphDriveItem
    fmt.Fprintf(os.Stdout, "Response from `TeamsChannelApi.TeamsChannelsGetFilesFolder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** | key: id of team | 
**channelId** | **string** | key: id of channel | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsChannelsGetFilesFolderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphDriveItem**](MicrosoftGraphDriveItem.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsChannelsGetFilesFolderContent

> *os.File TeamsChannelsGetFilesFolderContent(ctx, teamId, channelId).Execute()

Get media content for the navigation property filesFolder from teams

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
    resp, r, err := api_client.TeamsChannelApi.TeamsChannelsGetFilesFolderContent(context.Background(), teamId, channelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsChannelApi.TeamsChannelsGetFilesFolderContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsChannelsGetFilesFolderContent`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `TeamsChannelApi.TeamsChannelsGetFilesFolderContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** | key: id of team | 
**channelId** | **string** | key: id of channel | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsChannelsGetFilesFolderContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[***os.File**](*os.File.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsChannelsGetMembers

> MicrosoftGraphConversationMember TeamsChannelsGetMembers(ctx, teamId, channelId, conversationMemberId).Select_(select_).Expand(expand).Execute()

Get members from teams



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
    conversationMemberId := "conversationMemberId_example" // string | key: id of conversationMember
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsChannelApi.TeamsChannelsGetMembers(context.Background(), teamId, channelId, conversationMemberId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsChannelApi.TeamsChannelsGetMembers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsChannelsGetMembers`: MicrosoftGraphConversationMember
    fmt.Fprintf(os.Stdout, "Response from `TeamsChannelApi.TeamsChannelsGetMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** | key: id of team | 
**channelId** | **string** | key: id of channel | 
**conversationMemberId** | **string** | key: id of conversationMember | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsChannelsGetMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphConversationMember**](MicrosoftGraphConversationMember.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsChannelsGetMessages

> MicrosoftGraphChatMessage TeamsChannelsGetMessages(ctx, teamId, channelId, chatMessageId).Select_(select_).Expand(expand).Execute()

Get messages from teams



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
    chatMessageId := "chatMessageId_example" // string | key: id of chatMessage
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsChannelApi.TeamsChannelsGetMessages(context.Background(), teamId, channelId, chatMessageId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsChannelApi.TeamsChannelsGetMessages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsChannelsGetMessages`: MicrosoftGraphChatMessage
    fmt.Fprintf(os.Stdout, "Response from `TeamsChannelApi.TeamsChannelsGetMessages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** | key: id of team | 
**channelId** | **string** | key: id of channel | 
**chatMessageId** | **string** | key: id of chatMessage | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsChannelsGetMessagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphChatMessage**](MicrosoftGraphChatMessage.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsChannelsGetTabs

> MicrosoftGraphTeamsTab TeamsChannelsGetTabs(ctx, teamId, channelId, teamsTabId).Select_(select_).Expand(expand).Execute()

Get tabs from teams



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
    teamsTabId := "teamsTabId_example" // string | key: id of teamsTab
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsChannelApi.TeamsChannelsGetTabs(context.Background(), teamId, channelId, teamsTabId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsChannelApi.TeamsChannelsGetTabs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsChannelsGetTabs`: MicrosoftGraphTeamsTab
    fmt.Fprintf(os.Stdout, "Response from `TeamsChannelApi.TeamsChannelsGetTabs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** | key: id of team | 
**channelId** | **string** | key: id of channel | 
**teamsTabId** | **string** | key: id of teamsTab | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsChannelsGetTabsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphTeamsTab**](MicrosoftGraphTeamsTab.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsChannelsListMembers

> CollectionOfConversationMember TeamsChannelsListMembers(ctx, teamId, channelId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get members from teams



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
    resp, r, err := api_client.TeamsChannelApi.TeamsChannelsListMembers(context.Background(), teamId, channelId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsChannelApi.TeamsChannelsListMembers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsChannelsListMembers`: CollectionOfConversationMember
    fmt.Fprintf(os.Stdout, "Response from `TeamsChannelApi.TeamsChannelsListMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** | key: id of team | 
**channelId** | **string** | key: id of channel | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsChannelsListMembersRequest struct via the builder pattern


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

[**CollectionOfConversationMember**](CollectionOfConversationMember.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsChannelsListMessages

> CollectionOfChatMessage TeamsChannelsListMessages(ctx, teamId, channelId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get messages from teams



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
    resp, r, err := api_client.TeamsChannelApi.TeamsChannelsListMessages(context.Background(), teamId, channelId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsChannelApi.TeamsChannelsListMessages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsChannelsListMessages`: CollectionOfChatMessage
    fmt.Fprintf(os.Stdout, "Response from `TeamsChannelApi.TeamsChannelsListMessages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** | key: id of team | 
**channelId** | **string** | key: id of channel | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsChannelsListMessagesRequest struct via the builder pattern


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

[**CollectionOfChatMessage**](CollectionOfChatMessage.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsChannelsListTabs

> CollectionOfTeamsTab TeamsChannelsListTabs(ctx, teamId, channelId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get tabs from teams



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
    resp, r, err := api_client.TeamsChannelApi.TeamsChannelsListTabs(context.Background(), teamId, channelId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsChannelApi.TeamsChannelsListTabs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsChannelsListTabs`: CollectionOfTeamsTab
    fmt.Fprintf(os.Stdout, "Response from `TeamsChannelApi.TeamsChannelsListTabs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** | key: id of team | 
**channelId** | **string** | key: id of channel | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsChannelsListTabsRequest struct via the builder pattern


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

[**CollectionOfTeamsTab**](CollectionOfTeamsTab.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsChannelsMessagesCreateHostedContents

> MicrosoftGraphChatMessageHostedContent TeamsChannelsMessagesCreateHostedContents(ctx, teamId, channelId, chatMessageId).MicrosoftGraphChatMessageHostedContent(microsoftGraphChatMessageHostedContent).Execute()

Create new navigation property to hostedContents for teams



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
    chatMessageId := "chatMessageId_example" // string | key: id of chatMessage
    microsoftGraphChatMessageHostedContent := *openapiclient.NewMicrosoftGraphChatMessageHostedContent() // MicrosoftGraphChatMessageHostedContent | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsChannelApi.TeamsChannelsMessagesCreateHostedContents(context.Background(), teamId, channelId, chatMessageId).MicrosoftGraphChatMessageHostedContent(microsoftGraphChatMessageHostedContent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsChannelApi.TeamsChannelsMessagesCreateHostedContents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsChannelsMessagesCreateHostedContents`: MicrosoftGraphChatMessageHostedContent
    fmt.Fprintf(os.Stdout, "Response from `TeamsChannelApi.TeamsChannelsMessagesCreateHostedContents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** | key: id of team | 
**channelId** | **string** | key: id of channel | 
**chatMessageId** | **string** | key: id of chatMessage | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsChannelsMessagesCreateHostedContentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **microsoftGraphChatMessageHostedContent** | [**MicrosoftGraphChatMessageHostedContent**](MicrosoftGraphChatMessageHostedContent.md) | New navigation property | 

### Return type

[**MicrosoftGraphChatMessageHostedContent**](MicrosoftGraphChatMessageHostedContent.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsChannelsMessagesCreateReplies

> MicrosoftGraphChatMessage TeamsChannelsMessagesCreateReplies(ctx, teamId, channelId, chatMessageId).MicrosoftGraphChatMessage(microsoftGraphChatMessage).Execute()

Create new navigation property to replies for teams



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
    chatMessageId := "chatMessageId_example" // string | key: id of chatMessage
    microsoftGraphChatMessage := *openapiclient.NewMicrosoftGraphChatMessage() // MicrosoftGraphChatMessage | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsChannelApi.TeamsChannelsMessagesCreateReplies(context.Background(), teamId, channelId, chatMessageId).MicrosoftGraphChatMessage(microsoftGraphChatMessage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsChannelApi.TeamsChannelsMessagesCreateReplies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsChannelsMessagesCreateReplies`: MicrosoftGraphChatMessage
    fmt.Fprintf(os.Stdout, "Response from `TeamsChannelApi.TeamsChannelsMessagesCreateReplies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** | key: id of team | 
**channelId** | **string** | key: id of channel | 
**chatMessageId** | **string** | key: id of chatMessage | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsChannelsMessagesCreateRepliesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **microsoftGraphChatMessage** | [**MicrosoftGraphChatMessage**](MicrosoftGraphChatMessage.md) | New navigation property | 

### Return type

[**MicrosoftGraphChatMessage**](MicrosoftGraphChatMessage.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsChannelsMessagesDeleteHostedContents

> TeamsChannelsMessagesDeleteHostedContents(ctx, teamId, channelId, chatMessageId, chatMessageHostedContentId).IfMatch(ifMatch).Execute()

Delete navigation property hostedContents for teams



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
    chatMessageId := "chatMessageId_example" // string | key: id of chatMessage
    chatMessageHostedContentId := "chatMessageHostedContentId_example" // string | key: id of chatMessageHostedContent
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsChannelApi.TeamsChannelsMessagesDeleteHostedContents(context.Background(), teamId, channelId, chatMessageId, chatMessageHostedContentId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsChannelApi.TeamsChannelsMessagesDeleteHostedContents``: %v\n", err)
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
**chatMessageId** | **string** | key: id of chatMessage | 
**chatMessageHostedContentId** | **string** | key: id of chatMessageHostedContent | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsChannelsMessagesDeleteHostedContentsRequest struct via the builder pattern


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


## TeamsChannelsMessagesDeleteReplies

> TeamsChannelsMessagesDeleteReplies(ctx, teamId, channelId, chatMessageId, chatMessageId1).IfMatch(ifMatch).Execute()

Delete navigation property replies for teams



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
    chatMessageId := "chatMessageId_example" // string | key: id of chatMessage
    chatMessageId1 := "chatMessageId1_example" // string | key: id of chatMessage
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsChannelApi.TeamsChannelsMessagesDeleteReplies(context.Background(), teamId, channelId, chatMessageId, chatMessageId1).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsChannelApi.TeamsChannelsMessagesDeleteReplies``: %v\n", err)
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
**chatMessageId** | **string** | key: id of chatMessage | 
**chatMessageId1** | **string** | key: id of chatMessage | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsChannelsMessagesDeleteRepliesRequest struct via the builder pattern


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


## TeamsChannelsMessagesGetHostedContents

> MicrosoftGraphChatMessageHostedContent TeamsChannelsMessagesGetHostedContents(ctx, teamId, channelId, chatMessageId, chatMessageHostedContentId).Select_(select_).Expand(expand).Execute()

Get hostedContents from teams



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
    chatMessageId := "chatMessageId_example" // string | key: id of chatMessage
    chatMessageHostedContentId := "chatMessageHostedContentId_example" // string | key: id of chatMessageHostedContent
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsChannelApi.TeamsChannelsMessagesGetHostedContents(context.Background(), teamId, channelId, chatMessageId, chatMessageHostedContentId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsChannelApi.TeamsChannelsMessagesGetHostedContents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsChannelsMessagesGetHostedContents`: MicrosoftGraphChatMessageHostedContent
    fmt.Fprintf(os.Stdout, "Response from `TeamsChannelApi.TeamsChannelsMessagesGetHostedContents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** | key: id of team | 
**channelId** | **string** | key: id of channel | 
**chatMessageId** | **string** | key: id of chatMessage | 
**chatMessageHostedContentId** | **string** | key: id of chatMessageHostedContent | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsChannelsMessagesGetHostedContentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphChatMessageHostedContent**](MicrosoftGraphChatMessageHostedContent.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsChannelsMessagesGetReplies

> MicrosoftGraphChatMessage TeamsChannelsMessagesGetReplies(ctx, teamId, channelId, chatMessageId, chatMessageId1).Select_(select_).Expand(expand).Execute()

Get replies from teams



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
    chatMessageId := "chatMessageId_example" // string | key: id of chatMessage
    chatMessageId1 := "chatMessageId1_example" // string | key: id of chatMessage
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsChannelApi.TeamsChannelsMessagesGetReplies(context.Background(), teamId, channelId, chatMessageId, chatMessageId1).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsChannelApi.TeamsChannelsMessagesGetReplies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsChannelsMessagesGetReplies`: MicrosoftGraphChatMessage
    fmt.Fprintf(os.Stdout, "Response from `TeamsChannelApi.TeamsChannelsMessagesGetReplies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** | key: id of team | 
**channelId** | **string** | key: id of channel | 
**chatMessageId** | **string** | key: id of chatMessage | 
**chatMessageId1** | **string** | key: id of chatMessage | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsChannelsMessagesGetRepliesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphChatMessage**](MicrosoftGraphChatMessage.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsChannelsMessagesListHostedContents

> CollectionOfChatMessageHostedContent TeamsChannelsMessagesListHostedContents(ctx, teamId, channelId, chatMessageId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get hostedContents from teams



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
    chatMessageId := "chatMessageId_example" // string | key: id of chatMessage
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
    resp, r, err := api_client.TeamsChannelApi.TeamsChannelsMessagesListHostedContents(context.Background(), teamId, channelId, chatMessageId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsChannelApi.TeamsChannelsMessagesListHostedContents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsChannelsMessagesListHostedContents`: CollectionOfChatMessageHostedContent
    fmt.Fprintf(os.Stdout, "Response from `TeamsChannelApi.TeamsChannelsMessagesListHostedContents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** | key: id of team | 
**channelId** | **string** | key: id of channel | 
**chatMessageId** | **string** | key: id of chatMessage | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsChannelsMessagesListHostedContentsRequest struct via the builder pattern


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

[**CollectionOfChatMessageHostedContent**](CollectionOfChatMessageHostedContent.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsChannelsMessagesListReplies

> CollectionOfChatMessage TeamsChannelsMessagesListReplies(ctx, teamId, channelId, chatMessageId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get replies from teams



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
    chatMessageId := "chatMessageId_example" // string | key: id of chatMessage
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
    resp, r, err := api_client.TeamsChannelApi.TeamsChannelsMessagesListReplies(context.Background(), teamId, channelId, chatMessageId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsChannelApi.TeamsChannelsMessagesListReplies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsChannelsMessagesListReplies`: CollectionOfChatMessage
    fmt.Fprintf(os.Stdout, "Response from `TeamsChannelApi.TeamsChannelsMessagesListReplies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** | key: id of team | 
**channelId** | **string** | key: id of channel | 
**chatMessageId** | **string** | key: id of chatMessage | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsChannelsMessagesListRepliesRequest struct via the builder pattern


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

[**CollectionOfChatMessage**](CollectionOfChatMessage.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsChannelsMessagesUpdateHostedContents

> TeamsChannelsMessagesUpdateHostedContents(ctx, teamId, channelId, chatMessageId, chatMessageHostedContentId).MicrosoftGraphChatMessageHostedContent(microsoftGraphChatMessageHostedContent).Execute()

Update the navigation property hostedContents in teams



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
    chatMessageId := "chatMessageId_example" // string | key: id of chatMessage
    chatMessageHostedContentId := "chatMessageHostedContentId_example" // string | key: id of chatMessageHostedContent
    microsoftGraphChatMessageHostedContent := *openapiclient.NewMicrosoftGraphChatMessageHostedContent() // MicrosoftGraphChatMessageHostedContent | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsChannelApi.TeamsChannelsMessagesUpdateHostedContents(context.Background(), teamId, channelId, chatMessageId, chatMessageHostedContentId).MicrosoftGraphChatMessageHostedContent(microsoftGraphChatMessageHostedContent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsChannelApi.TeamsChannelsMessagesUpdateHostedContents``: %v\n", err)
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
**chatMessageId** | **string** | key: id of chatMessage | 
**chatMessageHostedContentId** | **string** | key: id of chatMessageHostedContent | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsChannelsMessagesUpdateHostedContentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **microsoftGraphChatMessageHostedContent** | [**MicrosoftGraphChatMessageHostedContent**](MicrosoftGraphChatMessageHostedContent.md) | New navigation property values | 

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


## TeamsChannelsMessagesUpdateReplies

> TeamsChannelsMessagesUpdateReplies(ctx, teamId, channelId, chatMessageId, chatMessageId1).MicrosoftGraphChatMessage(microsoftGraphChatMessage).Execute()

Update the navigation property replies in teams



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
    chatMessageId := "chatMessageId_example" // string | key: id of chatMessage
    chatMessageId1 := "chatMessageId1_example" // string | key: id of chatMessage
    microsoftGraphChatMessage := *openapiclient.NewMicrosoftGraphChatMessage() // MicrosoftGraphChatMessage | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsChannelApi.TeamsChannelsMessagesUpdateReplies(context.Background(), teamId, channelId, chatMessageId, chatMessageId1).MicrosoftGraphChatMessage(microsoftGraphChatMessage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsChannelApi.TeamsChannelsMessagesUpdateReplies``: %v\n", err)
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
**chatMessageId** | **string** | key: id of chatMessage | 
**chatMessageId1** | **string** | key: id of chatMessage | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsChannelsMessagesUpdateRepliesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **microsoftGraphChatMessage** | [**MicrosoftGraphChatMessage**](MicrosoftGraphChatMessage.md) | New navigation property values | 

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


## TeamsChannelsTabsDeleteRefTeamsApp

> TeamsChannelsTabsDeleteRefTeamsApp(ctx, teamId, channelId, teamsTabId).IfMatch(ifMatch).Execute()

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
    channelId := "channelId_example" // string | key: id of channel
    teamsTabId := "teamsTabId_example" // string | key: id of teamsTab
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsChannelApi.TeamsChannelsTabsDeleteRefTeamsApp(context.Background(), teamId, channelId, teamsTabId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsChannelApi.TeamsChannelsTabsDeleteRefTeamsApp``: %v\n", err)
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
**teamsTabId** | **string** | key: id of teamsTab | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsChannelsTabsDeleteRefTeamsAppRequest struct via the builder pattern


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


## TeamsChannelsTabsGetRefTeamsApp

> string TeamsChannelsTabsGetRefTeamsApp(ctx, teamId, channelId, teamsTabId).Execute()

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
    channelId := "channelId_example" // string | key: id of channel
    teamsTabId := "teamsTabId_example" // string | key: id of teamsTab

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsChannelApi.TeamsChannelsTabsGetRefTeamsApp(context.Background(), teamId, channelId, teamsTabId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsChannelApi.TeamsChannelsTabsGetRefTeamsApp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsChannelsTabsGetRefTeamsApp`: string
    fmt.Fprintf(os.Stdout, "Response from `TeamsChannelApi.TeamsChannelsTabsGetRefTeamsApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** | key: id of team | 
**channelId** | **string** | key: id of channel | 
**teamsTabId** | **string** | key: id of teamsTab | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsChannelsTabsGetRefTeamsAppRequest struct via the builder pattern


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


## TeamsChannelsTabsGetTeamsApp

> MicrosoftGraphTeamsApp TeamsChannelsTabsGetTeamsApp(ctx, teamId, channelId, teamsTabId).Select_(select_).Expand(expand).Execute()

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
    channelId := "channelId_example" // string | key: id of channel
    teamsTabId := "teamsTabId_example" // string | key: id of teamsTab
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsChannelApi.TeamsChannelsTabsGetTeamsApp(context.Background(), teamId, channelId, teamsTabId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsChannelApi.TeamsChannelsTabsGetTeamsApp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsChannelsTabsGetTeamsApp`: MicrosoftGraphTeamsApp
    fmt.Fprintf(os.Stdout, "Response from `TeamsChannelApi.TeamsChannelsTabsGetTeamsApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** | key: id of team | 
**channelId** | **string** | key: id of channel | 
**teamsTabId** | **string** | key: id of teamsTab | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsChannelsTabsGetTeamsAppRequest struct via the builder pattern


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


## TeamsChannelsTabsUpdateRefTeamsApp

> TeamsChannelsTabsUpdateRefTeamsApp(ctx, teamId, channelId, teamsTabId).RequestBody(requestBody).Execute()

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
    channelId := "channelId_example" // string | key: id of channel
    teamsTabId := "teamsTabId_example" // string | key: id of teamsTab
    requestBody := map[string]map[string]interface{}{"key": map[string]interface{}(123)} // map[string]map[string]interface{} | New navigation property ref values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsChannelApi.TeamsChannelsTabsUpdateRefTeamsApp(context.Background(), teamId, channelId, teamsTabId).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsChannelApi.TeamsChannelsTabsUpdateRefTeamsApp``: %v\n", err)
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
**teamsTabId** | **string** | key: id of teamsTab | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsChannelsTabsUpdateRefTeamsAppRequest struct via the builder pattern


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


## TeamsChannelsUpdateFilesFolder

> TeamsChannelsUpdateFilesFolder(ctx, teamId, channelId).MicrosoftGraphDriveItem(microsoftGraphDriveItem).Execute()

Update the navigation property filesFolder in teams



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
    microsoftGraphDriveItem := *openapiclient.NewMicrosoftGraphDriveItem() // MicrosoftGraphDriveItem | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsChannelApi.TeamsChannelsUpdateFilesFolder(context.Background(), teamId, channelId).MicrosoftGraphDriveItem(microsoftGraphDriveItem).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsChannelApi.TeamsChannelsUpdateFilesFolder``: %v\n", err)
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

Other parameters are passed through a pointer to a apiTeamsChannelsUpdateFilesFolderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **microsoftGraphDriveItem** | [**MicrosoftGraphDriveItem**](MicrosoftGraphDriveItem.md) | New navigation property values | 

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


## TeamsChannelsUpdateFilesFolderContent

> TeamsChannelsUpdateFilesFolderContent(ctx, teamId, channelId).Body(body).Execute()

Update media content for the navigation property filesFolder in teams

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
    body := os.NewFile(1234, "some_file") // *os.File | New media content.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsChannelApi.TeamsChannelsUpdateFilesFolderContent(context.Background(), teamId, channelId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsChannelApi.TeamsChannelsUpdateFilesFolderContent``: %v\n", err)
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

Other parameters are passed through a pointer to a apiTeamsChannelsUpdateFilesFolderContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | ***os.File** | New media content. | 

### Return type

 (empty response body)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/octet-stream
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsChannelsUpdateMembers

> TeamsChannelsUpdateMembers(ctx, teamId, channelId, conversationMemberId).MicrosoftGraphConversationMember(microsoftGraphConversationMember).Execute()

Update the navigation property members in teams



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
    conversationMemberId := "conversationMemberId_example" // string | key: id of conversationMember
    microsoftGraphConversationMember := *openapiclient.NewMicrosoftGraphConversationMember() // MicrosoftGraphConversationMember | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsChannelApi.TeamsChannelsUpdateMembers(context.Background(), teamId, channelId, conversationMemberId).MicrosoftGraphConversationMember(microsoftGraphConversationMember).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsChannelApi.TeamsChannelsUpdateMembers``: %v\n", err)
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
**conversationMemberId** | **string** | key: id of conversationMember | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsChannelsUpdateMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **microsoftGraphConversationMember** | [**MicrosoftGraphConversationMember**](MicrosoftGraphConversationMember.md) | New navigation property values | 

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


## TeamsChannelsUpdateMessages

> TeamsChannelsUpdateMessages(ctx, teamId, channelId, chatMessageId).MicrosoftGraphChatMessage(microsoftGraphChatMessage).Execute()

Update the navigation property messages in teams



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
    chatMessageId := "chatMessageId_example" // string | key: id of chatMessage
    microsoftGraphChatMessage := *openapiclient.NewMicrosoftGraphChatMessage() // MicrosoftGraphChatMessage | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsChannelApi.TeamsChannelsUpdateMessages(context.Background(), teamId, channelId, chatMessageId).MicrosoftGraphChatMessage(microsoftGraphChatMessage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsChannelApi.TeamsChannelsUpdateMessages``: %v\n", err)
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
**chatMessageId** | **string** | key: id of chatMessage | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsChannelsUpdateMessagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **microsoftGraphChatMessage** | [**MicrosoftGraphChatMessage**](MicrosoftGraphChatMessage.md) | New navigation property values | 

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


## TeamsChannelsUpdateTabs

> TeamsChannelsUpdateTabs(ctx, teamId, channelId, teamsTabId).MicrosoftGraphTeamsTab(microsoftGraphTeamsTab).Execute()

Update the navigation property tabs in teams



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
    teamsTabId := "teamsTabId_example" // string | key: id of teamsTab
    microsoftGraphTeamsTab := *openapiclient.NewMicrosoftGraphTeamsTab() // MicrosoftGraphTeamsTab | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsChannelApi.TeamsChannelsUpdateTabs(context.Background(), teamId, channelId, teamsTabId).MicrosoftGraphTeamsTab(microsoftGraphTeamsTab).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsChannelApi.TeamsChannelsUpdateTabs``: %v\n", err)
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
**teamsTabId** | **string** | key: id of teamsTab | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsChannelsUpdateTabsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **microsoftGraphTeamsTab** | [**MicrosoftGraphTeamsTab**](MicrosoftGraphTeamsTab.md) | New navigation property values | 

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


## TeamsCreateChannels

> MicrosoftGraphChannel TeamsCreateChannels(ctx, teamId).MicrosoftGraphChannel(microsoftGraphChannel).Execute()

Create new navigation property to channels for teams



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
    microsoftGraphChannel := *openapiclient.NewMicrosoftGraphChannel() // MicrosoftGraphChannel | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsChannelApi.TeamsCreateChannels(context.Background(), teamId).MicrosoftGraphChannel(microsoftGraphChannel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsChannelApi.TeamsCreateChannels``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsCreateChannels`: MicrosoftGraphChannel
    fmt.Fprintf(os.Stdout, "Response from `TeamsChannelApi.TeamsCreateChannels`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** | key: id of team | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsCreateChannelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphChannel** | [**MicrosoftGraphChannel**](MicrosoftGraphChannel.md) | New navigation property | 

### Return type

[**MicrosoftGraphChannel**](MicrosoftGraphChannel.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsDeleteChannels

> TeamsDeleteChannels(ctx, teamId, channelId).IfMatch(ifMatch).Execute()

Delete navigation property channels for teams



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
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsChannelApi.TeamsDeleteChannels(context.Background(), teamId, channelId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsChannelApi.TeamsDeleteChannels``: %v\n", err)
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

Other parameters are passed through a pointer to a apiTeamsDeleteChannelsRequest struct via the builder pattern


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


## TeamsDeletePrimaryChannel

> TeamsDeletePrimaryChannel(ctx, teamId).IfMatch(ifMatch).Execute()

Delete navigation property primaryChannel for teams



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
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsChannelApi.TeamsDeletePrimaryChannel(context.Background(), teamId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsChannelApi.TeamsDeletePrimaryChannel``: %v\n", err)
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

Other parameters are passed through a pointer to a apiTeamsDeletePrimaryChannelRequest struct via the builder pattern


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


## TeamsGetChannels

> MicrosoftGraphChannel TeamsGetChannels(ctx, teamId, channelId).Select_(select_).Expand(expand).Execute()

Get channels from teams



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
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsChannelApi.TeamsGetChannels(context.Background(), teamId, channelId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsChannelApi.TeamsGetChannels``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsGetChannels`: MicrosoftGraphChannel
    fmt.Fprintf(os.Stdout, "Response from `TeamsChannelApi.TeamsGetChannels`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** | key: id of team | 
**channelId** | **string** | key: id of channel | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsGetChannelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphChannel**](MicrosoftGraphChannel.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsGetPrimaryChannel

> MicrosoftGraphChannel TeamsGetPrimaryChannel(ctx, teamId).Select_(select_).Expand(expand).Execute()

Get primaryChannel from teams



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
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsChannelApi.TeamsGetPrimaryChannel(context.Background(), teamId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsChannelApi.TeamsGetPrimaryChannel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsGetPrimaryChannel`: MicrosoftGraphChannel
    fmt.Fprintf(os.Stdout, "Response from `TeamsChannelApi.TeamsGetPrimaryChannel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** | key: id of team | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsGetPrimaryChannelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphChannel**](MicrosoftGraphChannel.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsListChannels

> CollectionOfChannel TeamsListChannels(ctx, teamId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get channels from teams



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
    resp, r, err := api_client.TeamsChannelApi.TeamsListChannels(context.Background(), teamId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsChannelApi.TeamsListChannels``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsListChannels`: CollectionOfChannel
    fmt.Fprintf(os.Stdout, "Response from `TeamsChannelApi.TeamsListChannels`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** | key: id of team | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsListChannelsRequest struct via the builder pattern


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

[**CollectionOfChannel**](CollectionOfChannel.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsPrimaryChannelCreateMembers

> MicrosoftGraphConversationMember TeamsPrimaryChannelCreateMembers(ctx, teamId).MicrosoftGraphConversationMember(microsoftGraphConversationMember).Execute()

Create new navigation property to members for teams



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
    microsoftGraphConversationMember := *openapiclient.NewMicrosoftGraphConversationMember() // MicrosoftGraphConversationMember | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsChannelApi.TeamsPrimaryChannelCreateMembers(context.Background(), teamId).MicrosoftGraphConversationMember(microsoftGraphConversationMember).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsChannelApi.TeamsPrimaryChannelCreateMembers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsPrimaryChannelCreateMembers`: MicrosoftGraphConversationMember
    fmt.Fprintf(os.Stdout, "Response from `TeamsChannelApi.TeamsPrimaryChannelCreateMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** | key: id of team | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsPrimaryChannelCreateMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphConversationMember** | [**MicrosoftGraphConversationMember**](MicrosoftGraphConversationMember.md) | New navigation property | 

### Return type

[**MicrosoftGraphConversationMember**](MicrosoftGraphConversationMember.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsPrimaryChannelCreateMessages

> MicrosoftGraphChatMessage TeamsPrimaryChannelCreateMessages(ctx, teamId).MicrosoftGraphChatMessage(microsoftGraphChatMessage).Execute()

Create new navigation property to messages for teams



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
    microsoftGraphChatMessage := *openapiclient.NewMicrosoftGraphChatMessage() // MicrosoftGraphChatMessage | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsChannelApi.TeamsPrimaryChannelCreateMessages(context.Background(), teamId).MicrosoftGraphChatMessage(microsoftGraphChatMessage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsChannelApi.TeamsPrimaryChannelCreateMessages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsPrimaryChannelCreateMessages`: MicrosoftGraphChatMessage
    fmt.Fprintf(os.Stdout, "Response from `TeamsChannelApi.TeamsPrimaryChannelCreateMessages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** | key: id of team | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsPrimaryChannelCreateMessagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphChatMessage** | [**MicrosoftGraphChatMessage**](MicrosoftGraphChatMessage.md) | New navigation property | 

### Return type

[**MicrosoftGraphChatMessage**](MicrosoftGraphChatMessage.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsPrimaryChannelCreateTabs

> MicrosoftGraphTeamsTab TeamsPrimaryChannelCreateTabs(ctx, teamId).MicrosoftGraphTeamsTab(microsoftGraphTeamsTab).Execute()

Create new navigation property to tabs for teams



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
    microsoftGraphTeamsTab := *openapiclient.NewMicrosoftGraphTeamsTab() // MicrosoftGraphTeamsTab | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsChannelApi.TeamsPrimaryChannelCreateTabs(context.Background(), teamId).MicrosoftGraphTeamsTab(microsoftGraphTeamsTab).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsChannelApi.TeamsPrimaryChannelCreateTabs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsPrimaryChannelCreateTabs`: MicrosoftGraphTeamsTab
    fmt.Fprintf(os.Stdout, "Response from `TeamsChannelApi.TeamsPrimaryChannelCreateTabs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** | key: id of team | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsPrimaryChannelCreateTabsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphTeamsTab** | [**MicrosoftGraphTeamsTab**](MicrosoftGraphTeamsTab.md) | New navigation property | 

### Return type

[**MicrosoftGraphTeamsTab**](MicrosoftGraphTeamsTab.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsPrimaryChannelDeleteFilesFolder

> TeamsPrimaryChannelDeleteFilesFolder(ctx, teamId).IfMatch(ifMatch).Execute()

Delete navigation property filesFolder for teams



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
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsChannelApi.TeamsPrimaryChannelDeleteFilesFolder(context.Background(), teamId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsChannelApi.TeamsPrimaryChannelDeleteFilesFolder``: %v\n", err)
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

Other parameters are passed through a pointer to a apiTeamsPrimaryChannelDeleteFilesFolderRequest struct via the builder pattern


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


## TeamsPrimaryChannelDeleteMembers

> TeamsPrimaryChannelDeleteMembers(ctx, teamId, conversationMemberId).IfMatch(ifMatch).Execute()

Delete navigation property members for teams



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
    conversationMemberId := "conversationMemberId_example" // string | key: id of conversationMember
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsChannelApi.TeamsPrimaryChannelDeleteMembers(context.Background(), teamId, conversationMemberId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsChannelApi.TeamsPrimaryChannelDeleteMembers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** | key: id of team | 
**conversationMemberId** | **string** | key: id of conversationMember | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsPrimaryChannelDeleteMembersRequest struct via the builder pattern


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


## TeamsPrimaryChannelDeleteMessages

> TeamsPrimaryChannelDeleteMessages(ctx, teamId, chatMessageId).IfMatch(ifMatch).Execute()

Delete navigation property messages for teams



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
    chatMessageId := "chatMessageId_example" // string | key: id of chatMessage
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsChannelApi.TeamsPrimaryChannelDeleteMessages(context.Background(), teamId, chatMessageId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsChannelApi.TeamsPrimaryChannelDeleteMessages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** | key: id of team | 
**chatMessageId** | **string** | key: id of chatMessage | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsPrimaryChannelDeleteMessagesRequest struct via the builder pattern


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


## TeamsPrimaryChannelDeleteTabs

> TeamsPrimaryChannelDeleteTabs(ctx, teamId, teamsTabId).IfMatch(ifMatch).Execute()

Delete navigation property tabs for teams



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
    teamsTabId := "teamsTabId_example" // string | key: id of teamsTab
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsChannelApi.TeamsPrimaryChannelDeleteTabs(context.Background(), teamId, teamsTabId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsChannelApi.TeamsPrimaryChannelDeleteTabs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** | key: id of team | 
**teamsTabId** | **string** | key: id of teamsTab | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsPrimaryChannelDeleteTabsRequest struct via the builder pattern


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


## TeamsPrimaryChannelGetFilesFolder

> MicrosoftGraphDriveItem TeamsPrimaryChannelGetFilesFolder(ctx, teamId).Select_(select_).Expand(expand).Execute()

Get filesFolder from teams



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
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsChannelApi.TeamsPrimaryChannelGetFilesFolder(context.Background(), teamId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsChannelApi.TeamsPrimaryChannelGetFilesFolder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsPrimaryChannelGetFilesFolder`: MicrosoftGraphDriveItem
    fmt.Fprintf(os.Stdout, "Response from `TeamsChannelApi.TeamsPrimaryChannelGetFilesFolder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** | key: id of team | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsPrimaryChannelGetFilesFolderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphDriveItem**](MicrosoftGraphDriveItem.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsPrimaryChannelGetFilesFolderContent

> *os.File TeamsPrimaryChannelGetFilesFolderContent(ctx, teamId).Execute()

Get media content for the navigation property filesFolder from teams

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
    resp, r, err := api_client.TeamsChannelApi.TeamsPrimaryChannelGetFilesFolderContent(context.Background(), teamId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsChannelApi.TeamsPrimaryChannelGetFilesFolderContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsPrimaryChannelGetFilesFolderContent`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `TeamsChannelApi.TeamsPrimaryChannelGetFilesFolderContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** | key: id of team | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsPrimaryChannelGetFilesFolderContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsPrimaryChannelGetMembers

> MicrosoftGraphConversationMember TeamsPrimaryChannelGetMembers(ctx, teamId, conversationMemberId).Select_(select_).Expand(expand).Execute()

Get members from teams



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
    conversationMemberId := "conversationMemberId_example" // string | key: id of conversationMember
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsChannelApi.TeamsPrimaryChannelGetMembers(context.Background(), teamId, conversationMemberId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsChannelApi.TeamsPrimaryChannelGetMembers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsPrimaryChannelGetMembers`: MicrosoftGraphConversationMember
    fmt.Fprintf(os.Stdout, "Response from `TeamsChannelApi.TeamsPrimaryChannelGetMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** | key: id of team | 
**conversationMemberId** | **string** | key: id of conversationMember | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsPrimaryChannelGetMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphConversationMember**](MicrosoftGraphConversationMember.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsPrimaryChannelGetMessages

> MicrosoftGraphChatMessage TeamsPrimaryChannelGetMessages(ctx, teamId, chatMessageId).Select_(select_).Expand(expand).Execute()

Get messages from teams



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
    chatMessageId := "chatMessageId_example" // string | key: id of chatMessage
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsChannelApi.TeamsPrimaryChannelGetMessages(context.Background(), teamId, chatMessageId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsChannelApi.TeamsPrimaryChannelGetMessages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsPrimaryChannelGetMessages`: MicrosoftGraphChatMessage
    fmt.Fprintf(os.Stdout, "Response from `TeamsChannelApi.TeamsPrimaryChannelGetMessages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** | key: id of team | 
**chatMessageId** | **string** | key: id of chatMessage | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsPrimaryChannelGetMessagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphChatMessage**](MicrosoftGraphChatMessage.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsPrimaryChannelGetTabs

> MicrosoftGraphTeamsTab TeamsPrimaryChannelGetTabs(ctx, teamId, teamsTabId).Select_(select_).Expand(expand).Execute()

Get tabs from teams



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
    teamsTabId := "teamsTabId_example" // string | key: id of teamsTab
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsChannelApi.TeamsPrimaryChannelGetTabs(context.Background(), teamId, teamsTabId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsChannelApi.TeamsPrimaryChannelGetTabs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsPrimaryChannelGetTabs`: MicrosoftGraphTeamsTab
    fmt.Fprintf(os.Stdout, "Response from `TeamsChannelApi.TeamsPrimaryChannelGetTabs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** | key: id of team | 
**teamsTabId** | **string** | key: id of teamsTab | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsPrimaryChannelGetTabsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphTeamsTab**](MicrosoftGraphTeamsTab.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsPrimaryChannelListMembers

> CollectionOfConversationMember TeamsPrimaryChannelListMembers(ctx, teamId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get members from teams



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
    resp, r, err := api_client.TeamsChannelApi.TeamsPrimaryChannelListMembers(context.Background(), teamId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsChannelApi.TeamsPrimaryChannelListMembers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsPrimaryChannelListMembers`: CollectionOfConversationMember
    fmt.Fprintf(os.Stdout, "Response from `TeamsChannelApi.TeamsPrimaryChannelListMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** | key: id of team | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsPrimaryChannelListMembersRequest struct via the builder pattern


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

[**CollectionOfConversationMember**](CollectionOfConversationMember.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsPrimaryChannelListMessages

> CollectionOfChatMessage TeamsPrimaryChannelListMessages(ctx, teamId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get messages from teams



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
    resp, r, err := api_client.TeamsChannelApi.TeamsPrimaryChannelListMessages(context.Background(), teamId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsChannelApi.TeamsPrimaryChannelListMessages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsPrimaryChannelListMessages`: CollectionOfChatMessage
    fmt.Fprintf(os.Stdout, "Response from `TeamsChannelApi.TeamsPrimaryChannelListMessages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** | key: id of team | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsPrimaryChannelListMessagesRequest struct via the builder pattern


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

[**CollectionOfChatMessage**](CollectionOfChatMessage.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsPrimaryChannelListTabs

> CollectionOfTeamsTab TeamsPrimaryChannelListTabs(ctx, teamId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get tabs from teams



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
    resp, r, err := api_client.TeamsChannelApi.TeamsPrimaryChannelListTabs(context.Background(), teamId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsChannelApi.TeamsPrimaryChannelListTabs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsPrimaryChannelListTabs`: CollectionOfTeamsTab
    fmt.Fprintf(os.Stdout, "Response from `TeamsChannelApi.TeamsPrimaryChannelListTabs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** | key: id of team | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsPrimaryChannelListTabsRequest struct via the builder pattern


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

[**CollectionOfTeamsTab**](CollectionOfTeamsTab.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsPrimaryChannelMessagesCreateHostedContents

> MicrosoftGraphChatMessageHostedContent TeamsPrimaryChannelMessagesCreateHostedContents(ctx, teamId, chatMessageId).MicrosoftGraphChatMessageHostedContent(microsoftGraphChatMessageHostedContent).Execute()

Create new navigation property to hostedContents for teams



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
    chatMessageId := "chatMessageId_example" // string | key: id of chatMessage
    microsoftGraphChatMessageHostedContent := *openapiclient.NewMicrosoftGraphChatMessageHostedContent() // MicrosoftGraphChatMessageHostedContent | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsChannelApi.TeamsPrimaryChannelMessagesCreateHostedContents(context.Background(), teamId, chatMessageId).MicrosoftGraphChatMessageHostedContent(microsoftGraphChatMessageHostedContent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsChannelApi.TeamsPrimaryChannelMessagesCreateHostedContents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsPrimaryChannelMessagesCreateHostedContents`: MicrosoftGraphChatMessageHostedContent
    fmt.Fprintf(os.Stdout, "Response from `TeamsChannelApi.TeamsPrimaryChannelMessagesCreateHostedContents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** | key: id of team | 
**chatMessageId** | **string** | key: id of chatMessage | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsPrimaryChannelMessagesCreateHostedContentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **microsoftGraphChatMessageHostedContent** | [**MicrosoftGraphChatMessageHostedContent**](MicrosoftGraphChatMessageHostedContent.md) | New navigation property | 

### Return type

[**MicrosoftGraphChatMessageHostedContent**](MicrosoftGraphChatMessageHostedContent.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsPrimaryChannelMessagesCreateReplies

> MicrosoftGraphChatMessage TeamsPrimaryChannelMessagesCreateReplies(ctx, teamId, chatMessageId).MicrosoftGraphChatMessage(microsoftGraphChatMessage).Execute()

Create new navigation property to replies for teams



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
    chatMessageId := "chatMessageId_example" // string | key: id of chatMessage
    microsoftGraphChatMessage := *openapiclient.NewMicrosoftGraphChatMessage() // MicrosoftGraphChatMessage | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsChannelApi.TeamsPrimaryChannelMessagesCreateReplies(context.Background(), teamId, chatMessageId).MicrosoftGraphChatMessage(microsoftGraphChatMessage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsChannelApi.TeamsPrimaryChannelMessagesCreateReplies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsPrimaryChannelMessagesCreateReplies`: MicrosoftGraphChatMessage
    fmt.Fprintf(os.Stdout, "Response from `TeamsChannelApi.TeamsPrimaryChannelMessagesCreateReplies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** | key: id of team | 
**chatMessageId** | **string** | key: id of chatMessage | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsPrimaryChannelMessagesCreateRepliesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **microsoftGraphChatMessage** | [**MicrosoftGraphChatMessage**](MicrosoftGraphChatMessage.md) | New navigation property | 

### Return type

[**MicrosoftGraphChatMessage**](MicrosoftGraphChatMessage.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsPrimaryChannelMessagesDeleteHostedContents

> TeamsPrimaryChannelMessagesDeleteHostedContents(ctx, teamId, chatMessageId, chatMessageHostedContentId).IfMatch(ifMatch).Execute()

Delete navigation property hostedContents for teams



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
    chatMessageId := "chatMessageId_example" // string | key: id of chatMessage
    chatMessageHostedContentId := "chatMessageHostedContentId_example" // string | key: id of chatMessageHostedContent
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsChannelApi.TeamsPrimaryChannelMessagesDeleteHostedContents(context.Background(), teamId, chatMessageId, chatMessageHostedContentId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsChannelApi.TeamsPrimaryChannelMessagesDeleteHostedContents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** | key: id of team | 
**chatMessageId** | **string** | key: id of chatMessage | 
**chatMessageHostedContentId** | **string** | key: id of chatMessageHostedContent | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsPrimaryChannelMessagesDeleteHostedContentsRequest struct via the builder pattern


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


## TeamsPrimaryChannelMessagesDeleteReplies

> TeamsPrimaryChannelMessagesDeleteReplies(ctx, teamId, chatMessageId, chatMessageId1).IfMatch(ifMatch).Execute()

Delete navigation property replies for teams



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
    chatMessageId := "chatMessageId_example" // string | key: id of chatMessage
    chatMessageId1 := "chatMessageId1_example" // string | key: id of chatMessage
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsChannelApi.TeamsPrimaryChannelMessagesDeleteReplies(context.Background(), teamId, chatMessageId, chatMessageId1).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsChannelApi.TeamsPrimaryChannelMessagesDeleteReplies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** | key: id of team | 
**chatMessageId** | **string** | key: id of chatMessage | 
**chatMessageId1** | **string** | key: id of chatMessage | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsPrimaryChannelMessagesDeleteRepliesRequest struct via the builder pattern


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


## TeamsPrimaryChannelMessagesGetHostedContents

> MicrosoftGraphChatMessageHostedContent TeamsPrimaryChannelMessagesGetHostedContents(ctx, teamId, chatMessageId, chatMessageHostedContentId).Select_(select_).Expand(expand).Execute()

Get hostedContents from teams



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
    chatMessageId := "chatMessageId_example" // string | key: id of chatMessage
    chatMessageHostedContentId := "chatMessageHostedContentId_example" // string | key: id of chatMessageHostedContent
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsChannelApi.TeamsPrimaryChannelMessagesGetHostedContents(context.Background(), teamId, chatMessageId, chatMessageHostedContentId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsChannelApi.TeamsPrimaryChannelMessagesGetHostedContents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsPrimaryChannelMessagesGetHostedContents`: MicrosoftGraphChatMessageHostedContent
    fmt.Fprintf(os.Stdout, "Response from `TeamsChannelApi.TeamsPrimaryChannelMessagesGetHostedContents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** | key: id of team | 
**chatMessageId** | **string** | key: id of chatMessage | 
**chatMessageHostedContentId** | **string** | key: id of chatMessageHostedContent | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsPrimaryChannelMessagesGetHostedContentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphChatMessageHostedContent**](MicrosoftGraphChatMessageHostedContent.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsPrimaryChannelMessagesGetReplies

> MicrosoftGraphChatMessage TeamsPrimaryChannelMessagesGetReplies(ctx, teamId, chatMessageId, chatMessageId1).Select_(select_).Expand(expand).Execute()

Get replies from teams



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
    chatMessageId := "chatMessageId_example" // string | key: id of chatMessage
    chatMessageId1 := "chatMessageId1_example" // string | key: id of chatMessage
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsChannelApi.TeamsPrimaryChannelMessagesGetReplies(context.Background(), teamId, chatMessageId, chatMessageId1).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsChannelApi.TeamsPrimaryChannelMessagesGetReplies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsPrimaryChannelMessagesGetReplies`: MicrosoftGraphChatMessage
    fmt.Fprintf(os.Stdout, "Response from `TeamsChannelApi.TeamsPrimaryChannelMessagesGetReplies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** | key: id of team | 
**chatMessageId** | **string** | key: id of chatMessage | 
**chatMessageId1** | **string** | key: id of chatMessage | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsPrimaryChannelMessagesGetRepliesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphChatMessage**](MicrosoftGraphChatMessage.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsPrimaryChannelMessagesListHostedContents

> CollectionOfChatMessageHostedContent TeamsPrimaryChannelMessagesListHostedContents(ctx, teamId, chatMessageId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get hostedContents from teams



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
    chatMessageId := "chatMessageId_example" // string | key: id of chatMessage
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
    resp, r, err := api_client.TeamsChannelApi.TeamsPrimaryChannelMessagesListHostedContents(context.Background(), teamId, chatMessageId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsChannelApi.TeamsPrimaryChannelMessagesListHostedContents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsPrimaryChannelMessagesListHostedContents`: CollectionOfChatMessageHostedContent
    fmt.Fprintf(os.Stdout, "Response from `TeamsChannelApi.TeamsPrimaryChannelMessagesListHostedContents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** | key: id of team | 
**chatMessageId** | **string** | key: id of chatMessage | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsPrimaryChannelMessagesListHostedContentsRequest struct via the builder pattern


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

[**CollectionOfChatMessageHostedContent**](CollectionOfChatMessageHostedContent.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsPrimaryChannelMessagesListReplies

> CollectionOfChatMessage TeamsPrimaryChannelMessagesListReplies(ctx, teamId, chatMessageId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get replies from teams



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
    chatMessageId := "chatMessageId_example" // string | key: id of chatMessage
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
    resp, r, err := api_client.TeamsChannelApi.TeamsPrimaryChannelMessagesListReplies(context.Background(), teamId, chatMessageId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsChannelApi.TeamsPrimaryChannelMessagesListReplies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsPrimaryChannelMessagesListReplies`: CollectionOfChatMessage
    fmt.Fprintf(os.Stdout, "Response from `TeamsChannelApi.TeamsPrimaryChannelMessagesListReplies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** | key: id of team | 
**chatMessageId** | **string** | key: id of chatMessage | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsPrimaryChannelMessagesListRepliesRequest struct via the builder pattern


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

[**CollectionOfChatMessage**](CollectionOfChatMessage.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsPrimaryChannelMessagesUpdateHostedContents

> TeamsPrimaryChannelMessagesUpdateHostedContents(ctx, teamId, chatMessageId, chatMessageHostedContentId).MicrosoftGraphChatMessageHostedContent(microsoftGraphChatMessageHostedContent).Execute()

Update the navigation property hostedContents in teams



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
    chatMessageId := "chatMessageId_example" // string | key: id of chatMessage
    chatMessageHostedContentId := "chatMessageHostedContentId_example" // string | key: id of chatMessageHostedContent
    microsoftGraphChatMessageHostedContent := *openapiclient.NewMicrosoftGraphChatMessageHostedContent() // MicrosoftGraphChatMessageHostedContent | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsChannelApi.TeamsPrimaryChannelMessagesUpdateHostedContents(context.Background(), teamId, chatMessageId, chatMessageHostedContentId).MicrosoftGraphChatMessageHostedContent(microsoftGraphChatMessageHostedContent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsChannelApi.TeamsPrimaryChannelMessagesUpdateHostedContents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** | key: id of team | 
**chatMessageId** | **string** | key: id of chatMessage | 
**chatMessageHostedContentId** | **string** | key: id of chatMessageHostedContent | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsPrimaryChannelMessagesUpdateHostedContentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **microsoftGraphChatMessageHostedContent** | [**MicrosoftGraphChatMessageHostedContent**](MicrosoftGraphChatMessageHostedContent.md) | New navigation property values | 

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


## TeamsPrimaryChannelMessagesUpdateReplies

> TeamsPrimaryChannelMessagesUpdateReplies(ctx, teamId, chatMessageId, chatMessageId1).MicrosoftGraphChatMessage(microsoftGraphChatMessage).Execute()

Update the navigation property replies in teams



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
    chatMessageId := "chatMessageId_example" // string | key: id of chatMessage
    chatMessageId1 := "chatMessageId1_example" // string | key: id of chatMessage
    microsoftGraphChatMessage := *openapiclient.NewMicrosoftGraphChatMessage() // MicrosoftGraphChatMessage | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsChannelApi.TeamsPrimaryChannelMessagesUpdateReplies(context.Background(), teamId, chatMessageId, chatMessageId1).MicrosoftGraphChatMessage(microsoftGraphChatMessage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsChannelApi.TeamsPrimaryChannelMessagesUpdateReplies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** | key: id of team | 
**chatMessageId** | **string** | key: id of chatMessage | 
**chatMessageId1** | **string** | key: id of chatMessage | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsPrimaryChannelMessagesUpdateRepliesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **microsoftGraphChatMessage** | [**MicrosoftGraphChatMessage**](MicrosoftGraphChatMessage.md) | New navigation property values | 

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


## TeamsPrimaryChannelTabsDeleteRefTeamsApp

> TeamsPrimaryChannelTabsDeleteRefTeamsApp(ctx, teamId, teamsTabId).IfMatch(ifMatch).Execute()

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
    teamsTabId := "teamsTabId_example" // string | key: id of teamsTab
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsChannelApi.TeamsPrimaryChannelTabsDeleteRefTeamsApp(context.Background(), teamId, teamsTabId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsChannelApi.TeamsPrimaryChannelTabsDeleteRefTeamsApp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** | key: id of team | 
**teamsTabId** | **string** | key: id of teamsTab | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsPrimaryChannelTabsDeleteRefTeamsAppRequest struct via the builder pattern


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


## TeamsPrimaryChannelTabsGetRefTeamsApp

> string TeamsPrimaryChannelTabsGetRefTeamsApp(ctx, teamId, teamsTabId).Execute()

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
    teamsTabId := "teamsTabId_example" // string | key: id of teamsTab

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsChannelApi.TeamsPrimaryChannelTabsGetRefTeamsApp(context.Background(), teamId, teamsTabId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsChannelApi.TeamsPrimaryChannelTabsGetRefTeamsApp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsPrimaryChannelTabsGetRefTeamsApp`: string
    fmt.Fprintf(os.Stdout, "Response from `TeamsChannelApi.TeamsPrimaryChannelTabsGetRefTeamsApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** | key: id of team | 
**teamsTabId** | **string** | key: id of teamsTab | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsPrimaryChannelTabsGetRefTeamsAppRequest struct via the builder pattern


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


## TeamsPrimaryChannelTabsGetTeamsApp

> MicrosoftGraphTeamsApp TeamsPrimaryChannelTabsGetTeamsApp(ctx, teamId, teamsTabId).Select_(select_).Expand(expand).Execute()

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
    teamsTabId := "teamsTabId_example" // string | key: id of teamsTab
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsChannelApi.TeamsPrimaryChannelTabsGetTeamsApp(context.Background(), teamId, teamsTabId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsChannelApi.TeamsPrimaryChannelTabsGetTeamsApp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsPrimaryChannelTabsGetTeamsApp`: MicrosoftGraphTeamsApp
    fmt.Fprintf(os.Stdout, "Response from `TeamsChannelApi.TeamsPrimaryChannelTabsGetTeamsApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** | key: id of team | 
**teamsTabId** | **string** | key: id of teamsTab | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsPrimaryChannelTabsGetTeamsAppRequest struct via the builder pattern


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


## TeamsPrimaryChannelTabsUpdateRefTeamsApp

> TeamsPrimaryChannelTabsUpdateRefTeamsApp(ctx, teamId, teamsTabId).RequestBody(requestBody).Execute()

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
    teamsTabId := "teamsTabId_example" // string | key: id of teamsTab
    requestBody := map[string]map[string]interface{}{"key": map[string]interface{}(123)} // map[string]map[string]interface{} | New navigation property ref values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsChannelApi.TeamsPrimaryChannelTabsUpdateRefTeamsApp(context.Background(), teamId, teamsTabId).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsChannelApi.TeamsPrimaryChannelTabsUpdateRefTeamsApp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** | key: id of team | 
**teamsTabId** | **string** | key: id of teamsTab | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsPrimaryChannelTabsUpdateRefTeamsAppRequest struct via the builder pattern


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


## TeamsPrimaryChannelUpdateFilesFolder

> TeamsPrimaryChannelUpdateFilesFolder(ctx, teamId).MicrosoftGraphDriveItem(microsoftGraphDriveItem).Execute()

Update the navigation property filesFolder in teams



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
    microsoftGraphDriveItem := *openapiclient.NewMicrosoftGraphDriveItem() // MicrosoftGraphDriveItem | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsChannelApi.TeamsPrimaryChannelUpdateFilesFolder(context.Background(), teamId).MicrosoftGraphDriveItem(microsoftGraphDriveItem).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsChannelApi.TeamsPrimaryChannelUpdateFilesFolder``: %v\n", err)
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

Other parameters are passed through a pointer to a apiTeamsPrimaryChannelUpdateFilesFolderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphDriveItem** | [**MicrosoftGraphDriveItem**](MicrosoftGraphDriveItem.md) | New navigation property values | 

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


## TeamsPrimaryChannelUpdateFilesFolderContent

> TeamsPrimaryChannelUpdateFilesFolderContent(ctx, teamId).Body(body).Execute()

Update media content for the navigation property filesFolder in teams

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
    body := os.NewFile(1234, "some_file") // *os.File | New media content.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsChannelApi.TeamsPrimaryChannelUpdateFilesFolderContent(context.Background(), teamId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsChannelApi.TeamsPrimaryChannelUpdateFilesFolderContent``: %v\n", err)
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

Other parameters are passed through a pointer to a apiTeamsPrimaryChannelUpdateFilesFolderContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | ***os.File** | New media content. | 

### Return type

 (empty response body)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/octet-stream
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsPrimaryChannelUpdateMembers

> TeamsPrimaryChannelUpdateMembers(ctx, teamId, conversationMemberId).MicrosoftGraphConversationMember(microsoftGraphConversationMember).Execute()

Update the navigation property members in teams



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
    conversationMemberId := "conversationMemberId_example" // string | key: id of conversationMember
    microsoftGraphConversationMember := *openapiclient.NewMicrosoftGraphConversationMember() // MicrosoftGraphConversationMember | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsChannelApi.TeamsPrimaryChannelUpdateMembers(context.Background(), teamId, conversationMemberId).MicrosoftGraphConversationMember(microsoftGraphConversationMember).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsChannelApi.TeamsPrimaryChannelUpdateMembers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** | key: id of team | 
**conversationMemberId** | **string** | key: id of conversationMember | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsPrimaryChannelUpdateMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **microsoftGraphConversationMember** | [**MicrosoftGraphConversationMember**](MicrosoftGraphConversationMember.md) | New navigation property values | 

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


## TeamsPrimaryChannelUpdateMessages

> TeamsPrimaryChannelUpdateMessages(ctx, teamId, chatMessageId).MicrosoftGraphChatMessage(microsoftGraphChatMessage).Execute()

Update the navigation property messages in teams



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
    chatMessageId := "chatMessageId_example" // string | key: id of chatMessage
    microsoftGraphChatMessage := *openapiclient.NewMicrosoftGraphChatMessage() // MicrosoftGraphChatMessage | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsChannelApi.TeamsPrimaryChannelUpdateMessages(context.Background(), teamId, chatMessageId).MicrosoftGraphChatMessage(microsoftGraphChatMessage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsChannelApi.TeamsPrimaryChannelUpdateMessages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** | key: id of team | 
**chatMessageId** | **string** | key: id of chatMessage | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsPrimaryChannelUpdateMessagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **microsoftGraphChatMessage** | [**MicrosoftGraphChatMessage**](MicrosoftGraphChatMessage.md) | New navigation property values | 

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


## TeamsPrimaryChannelUpdateTabs

> TeamsPrimaryChannelUpdateTabs(ctx, teamId, teamsTabId).MicrosoftGraphTeamsTab(microsoftGraphTeamsTab).Execute()

Update the navigation property tabs in teams



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
    teamsTabId := "teamsTabId_example" // string | key: id of teamsTab
    microsoftGraphTeamsTab := *openapiclient.NewMicrosoftGraphTeamsTab() // MicrosoftGraphTeamsTab | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsChannelApi.TeamsPrimaryChannelUpdateTabs(context.Background(), teamId, teamsTabId).MicrosoftGraphTeamsTab(microsoftGraphTeamsTab).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsChannelApi.TeamsPrimaryChannelUpdateTabs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** | key: id of team | 
**teamsTabId** | **string** | key: id of teamsTab | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsPrimaryChannelUpdateTabsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **microsoftGraphTeamsTab** | [**MicrosoftGraphTeamsTab**](MicrosoftGraphTeamsTab.md) | New navigation property values | 

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


## TeamsUpdateChannels

> TeamsUpdateChannels(ctx, teamId, channelId).MicrosoftGraphChannel(microsoftGraphChannel).Execute()

Update the navigation property channels in teams



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
    microsoftGraphChannel := *openapiclient.NewMicrosoftGraphChannel() // MicrosoftGraphChannel | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsChannelApi.TeamsUpdateChannels(context.Background(), teamId, channelId).MicrosoftGraphChannel(microsoftGraphChannel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsChannelApi.TeamsUpdateChannels``: %v\n", err)
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

Other parameters are passed through a pointer to a apiTeamsUpdateChannelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **microsoftGraphChannel** | [**MicrosoftGraphChannel**](MicrosoftGraphChannel.md) | New navigation property values | 

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


## TeamsUpdatePrimaryChannel

> TeamsUpdatePrimaryChannel(ctx, teamId).MicrosoftGraphChannel(microsoftGraphChannel).Execute()

Update the navigation property primaryChannel in teams



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
    microsoftGraphChannel := *openapiclient.NewMicrosoftGraphChannel() // MicrosoftGraphChannel | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsChannelApi.TeamsUpdatePrimaryChannel(context.Background(), teamId).MicrosoftGraphChannel(microsoftGraphChannel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsChannelApi.TeamsUpdatePrimaryChannel``: %v\n", err)
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

Other parameters are passed through a pointer to a apiTeamsUpdatePrimaryChannelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphChannel** | [**MicrosoftGraphChannel**](MicrosoftGraphChannel.md) | New navigation property values | 

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

