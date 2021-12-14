# \CommunicationsCallApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CommunicationsCallsCreateAudioRoutingGroups**](CommunicationsCallApi.md#CommunicationsCallsCreateAudioRoutingGroups) | **Post** /communications/calls/{call-id}/audioRoutingGroups | Create new navigation property to audioRoutingGroups for communications
[**CommunicationsCallsCreateOperations**](CommunicationsCallApi.md#CommunicationsCallsCreateOperations) | **Post** /communications/calls/{call-id}/operations | Create new navigation property to operations for communications
[**CommunicationsCallsCreateParticipants**](CommunicationsCallApi.md#CommunicationsCallsCreateParticipants) | **Post** /communications/calls/{call-id}/participants | Create new navigation property to participants for communications
[**CommunicationsCallsDeleteAudioRoutingGroups**](CommunicationsCallApi.md#CommunicationsCallsDeleteAudioRoutingGroups) | **Delete** /communications/calls/{call-id}/audioRoutingGroups/{audioRoutingGroup-id} | Delete navigation property audioRoutingGroups for communications
[**CommunicationsCallsDeleteOperations**](CommunicationsCallApi.md#CommunicationsCallsDeleteOperations) | **Delete** /communications/calls/{call-id}/operations/{commsOperation-id} | Delete navigation property operations for communications
[**CommunicationsCallsDeleteParticipants**](CommunicationsCallApi.md#CommunicationsCallsDeleteParticipants) | **Delete** /communications/calls/{call-id}/participants/{participant-id} | Delete navigation property participants for communications
[**CommunicationsCallsGetAudioRoutingGroups**](CommunicationsCallApi.md#CommunicationsCallsGetAudioRoutingGroups) | **Get** /communications/calls/{call-id}/audioRoutingGroups/{audioRoutingGroup-id} | Get audioRoutingGroups from communications
[**CommunicationsCallsGetOperations**](CommunicationsCallApi.md#CommunicationsCallsGetOperations) | **Get** /communications/calls/{call-id}/operations/{commsOperation-id} | Get operations from communications
[**CommunicationsCallsGetParticipants**](CommunicationsCallApi.md#CommunicationsCallsGetParticipants) | **Get** /communications/calls/{call-id}/participants/{participant-id} | Get participants from communications
[**CommunicationsCallsListAudioRoutingGroups**](CommunicationsCallApi.md#CommunicationsCallsListAudioRoutingGroups) | **Get** /communications/calls/{call-id}/audioRoutingGroups | Get audioRoutingGroups from communications
[**CommunicationsCallsListOperations**](CommunicationsCallApi.md#CommunicationsCallsListOperations) | **Get** /communications/calls/{call-id}/operations | Get operations from communications
[**CommunicationsCallsListParticipants**](CommunicationsCallApi.md#CommunicationsCallsListParticipants) | **Get** /communications/calls/{call-id}/participants | Get participants from communications
[**CommunicationsCallsUpdateAudioRoutingGroups**](CommunicationsCallApi.md#CommunicationsCallsUpdateAudioRoutingGroups) | **Patch** /communications/calls/{call-id}/audioRoutingGroups/{audioRoutingGroup-id} | Update the navigation property audioRoutingGroups in communications
[**CommunicationsCallsUpdateOperations**](CommunicationsCallApi.md#CommunicationsCallsUpdateOperations) | **Patch** /communications/calls/{call-id}/operations/{commsOperation-id} | Update the navigation property operations in communications
[**CommunicationsCallsUpdateParticipants**](CommunicationsCallApi.md#CommunicationsCallsUpdateParticipants) | **Patch** /communications/calls/{call-id}/participants/{participant-id} | Update the navigation property participants in communications
[**CommunicationsCreateCalls**](CommunicationsCallApi.md#CommunicationsCreateCalls) | **Post** /communications/calls | Create new navigation property to calls for communications
[**CommunicationsDeleteCalls**](CommunicationsCallApi.md#CommunicationsDeleteCalls) | **Delete** /communications/calls/{call-id} | Delete navigation property calls for communications
[**CommunicationsGetCalls**](CommunicationsCallApi.md#CommunicationsGetCalls) | **Get** /communications/calls/{call-id} | Get calls from communications
[**CommunicationsListCalls**](CommunicationsCallApi.md#CommunicationsListCalls) | **Get** /communications/calls | Get calls from communications
[**CommunicationsUpdateCalls**](CommunicationsCallApi.md#CommunicationsUpdateCalls) | **Patch** /communications/calls/{call-id} | Update the navigation property calls in communications



## CommunicationsCallsCreateAudioRoutingGroups

> MicrosoftGraphAudioRoutingGroup CommunicationsCallsCreateAudioRoutingGroups(ctx, callId).MicrosoftGraphAudioRoutingGroup(microsoftGraphAudioRoutingGroup).Execute()

Create new navigation property to audioRoutingGroups for communications



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
    callId := "callId_example" // string | key: id of call
    microsoftGraphAudioRoutingGroup := *openapiclient.NewMicrosoftGraphAudioRoutingGroup() // MicrosoftGraphAudioRoutingGroup | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CommunicationsCallApi.CommunicationsCallsCreateAudioRoutingGroups(context.Background(), callId).MicrosoftGraphAudioRoutingGroup(microsoftGraphAudioRoutingGroup).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommunicationsCallApi.CommunicationsCallsCreateAudioRoutingGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CommunicationsCallsCreateAudioRoutingGroups`: MicrosoftGraphAudioRoutingGroup
    fmt.Fprintf(os.Stdout, "Response from `CommunicationsCallApi.CommunicationsCallsCreateAudioRoutingGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callId** | **string** | key: id of call | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommunicationsCallsCreateAudioRoutingGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphAudioRoutingGroup** | [**MicrosoftGraphAudioRoutingGroup**](MicrosoftGraphAudioRoutingGroup.md) | New navigation property | 

### Return type

[**MicrosoftGraphAudioRoutingGroup**](MicrosoftGraphAudioRoutingGroup.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommunicationsCallsCreateOperations

> MicrosoftGraphCommsOperation CommunicationsCallsCreateOperations(ctx, callId).MicrosoftGraphCommsOperation(microsoftGraphCommsOperation).Execute()

Create new navigation property to operations for communications



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
    callId := "callId_example" // string | key: id of call
    microsoftGraphCommsOperation := *openapiclient.NewMicrosoftGraphCommsOperation() // MicrosoftGraphCommsOperation | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CommunicationsCallApi.CommunicationsCallsCreateOperations(context.Background(), callId).MicrosoftGraphCommsOperation(microsoftGraphCommsOperation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommunicationsCallApi.CommunicationsCallsCreateOperations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CommunicationsCallsCreateOperations`: MicrosoftGraphCommsOperation
    fmt.Fprintf(os.Stdout, "Response from `CommunicationsCallApi.CommunicationsCallsCreateOperations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callId** | **string** | key: id of call | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommunicationsCallsCreateOperationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphCommsOperation** | [**MicrosoftGraphCommsOperation**](MicrosoftGraphCommsOperation.md) | New navigation property | 

### Return type

[**MicrosoftGraphCommsOperation**](MicrosoftGraphCommsOperation.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommunicationsCallsCreateParticipants

> MicrosoftGraphParticipant CommunicationsCallsCreateParticipants(ctx, callId).MicrosoftGraphParticipant(microsoftGraphParticipant).Execute()

Create new navigation property to participants for communications



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
    callId := "callId_example" // string | key: id of call
    microsoftGraphParticipant := *openapiclient.NewMicrosoftGraphParticipant() // MicrosoftGraphParticipant | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CommunicationsCallApi.CommunicationsCallsCreateParticipants(context.Background(), callId).MicrosoftGraphParticipant(microsoftGraphParticipant).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommunicationsCallApi.CommunicationsCallsCreateParticipants``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CommunicationsCallsCreateParticipants`: MicrosoftGraphParticipant
    fmt.Fprintf(os.Stdout, "Response from `CommunicationsCallApi.CommunicationsCallsCreateParticipants`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callId** | **string** | key: id of call | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommunicationsCallsCreateParticipantsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphParticipant** | [**MicrosoftGraphParticipant**](MicrosoftGraphParticipant.md) | New navigation property | 

### Return type

[**MicrosoftGraphParticipant**](MicrosoftGraphParticipant.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommunicationsCallsDeleteAudioRoutingGroups

> CommunicationsCallsDeleteAudioRoutingGroups(ctx, callId, audioRoutingGroupId).IfMatch(ifMatch).Execute()

Delete navigation property audioRoutingGroups for communications



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
    callId := "callId_example" // string | key: id of call
    audioRoutingGroupId := "audioRoutingGroupId_example" // string | key: id of audioRoutingGroup
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CommunicationsCallApi.CommunicationsCallsDeleteAudioRoutingGroups(context.Background(), callId, audioRoutingGroupId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommunicationsCallApi.CommunicationsCallsDeleteAudioRoutingGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callId** | **string** | key: id of call | 
**audioRoutingGroupId** | **string** | key: id of audioRoutingGroup | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommunicationsCallsDeleteAudioRoutingGroupsRequest struct via the builder pattern


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


## CommunicationsCallsDeleteOperations

> CommunicationsCallsDeleteOperations(ctx, callId, commsOperationId).IfMatch(ifMatch).Execute()

Delete navigation property operations for communications



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
    callId := "callId_example" // string | key: id of call
    commsOperationId := "commsOperationId_example" // string | key: id of commsOperation
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CommunicationsCallApi.CommunicationsCallsDeleteOperations(context.Background(), callId, commsOperationId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommunicationsCallApi.CommunicationsCallsDeleteOperations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callId** | **string** | key: id of call | 
**commsOperationId** | **string** | key: id of commsOperation | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommunicationsCallsDeleteOperationsRequest struct via the builder pattern


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


## CommunicationsCallsDeleteParticipants

> CommunicationsCallsDeleteParticipants(ctx, callId, participantId).IfMatch(ifMatch).Execute()

Delete navigation property participants for communications



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
    callId := "callId_example" // string | key: id of call
    participantId := "participantId_example" // string | key: id of participant
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CommunicationsCallApi.CommunicationsCallsDeleteParticipants(context.Background(), callId, participantId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommunicationsCallApi.CommunicationsCallsDeleteParticipants``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callId** | **string** | key: id of call | 
**participantId** | **string** | key: id of participant | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommunicationsCallsDeleteParticipantsRequest struct via the builder pattern


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


## CommunicationsCallsGetAudioRoutingGroups

> MicrosoftGraphAudioRoutingGroup CommunicationsCallsGetAudioRoutingGroups(ctx, callId, audioRoutingGroupId).Select_(select_).Expand(expand).Execute()

Get audioRoutingGroups from communications



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
    callId := "callId_example" // string | key: id of call
    audioRoutingGroupId := "audioRoutingGroupId_example" // string | key: id of audioRoutingGroup
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CommunicationsCallApi.CommunicationsCallsGetAudioRoutingGroups(context.Background(), callId, audioRoutingGroupId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommunicationsCallApi.CommunicationsCallsGetAudioRoutingGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CommunicationsCallsGetAudioRoutingGroups`: MicrosoftGraphAudioRoutingGroup
    fmt.Fprintf(os.Stdout, "Response from `CommunicationsCallApi.CommunicationsCallsGetAudioRoutingGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callId** | **string** | key: id of call | 
**audioRoutingGroupId** | **string** | key: id of audioRoutingGroup | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommunicationsCallsGetAudioRoutingGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphAudioRoutingGroup**](MicrosoftGraphAudioRoutingGroup.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommunicationsCallsGetOperations

> MicrosoftGraphCommsOperation CommunicationsCallsGetOperations(ctx, callId, commsOperationId).Select_(select_).Expand(expand).Execute()

Get operations from communications



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
    callId := "callId_example" // string | key: id of call
    commsOperationId := "commsOperationId_example" // string | key: id of commsOperation
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CommunicationsCallApi.CommunicationsCallsGetOperations(context.Background(), callId, commsOperationId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommunicationsCallApi.CommunicationsCallsGetOperations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CommunicationsCallsGetOperations`: MicrosoftGraphCommsOperation
    fmt.Fprintf(os.Stdout, "Response from `CommunicationsCallApi.CommunicationsCallsGetOperations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callId** | **string** | key: id of call | 
**commsOperationId** | **string** | key: id of commsOperation | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommunicationsCallsGetOperationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphCommsOperation**](MicrosoftGraphCommsOperation.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommunicationsCallsGetParticipants

> MicrosoftGraphParticipant CommunicationsCallsGetParticipants(ctx, callId, participantId).Select_(select_).Expand(expand).Execute()

Get participants from communications



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
    callId := "callId_example" // string | key: id of call
    participantId := "participantId_example" // string | key: id of participant
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CommunicationsCallApi.CommunicationsCallsGetParticipants(context.Background(), callId, participantId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommunicationsCallApi.CommunicationsCallsGetParticipants``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CommunicationsCallsGetParticipants`: MicrosoftGraphParticipant
    fmt.Fprintf(os.Stdout, "Response from `CommunicationsCallApi.CommunicationsCallsGetParticipants`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callId** | **string** | key: id of call | 
**participantId** | **string** | key: id of participant | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommunicationsCallsGetParticipantsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphParticipant**](MicrosoftGraphParticipant.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommunicationsCallsListAudioRoutingGroups

> CollectionOfAudioRoutingGroup CommunicationsCallsListAudioRoutingGroups(ctx, callId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get audioRoutingGroups from communications



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
    callId := "callId_example" // string | key: id of call
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
    resp, r, err := api_client.CommunicationsCallApi.CommunicationsCallsListAudioRoutingGroups(context.Background(), callId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommunicationsCallApi.CommunicationsCallsListAudioRoutingGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CommunicationsCallsListAudioRoutingGroups`: CollectionOfAudioRoutingGroup
    fmt.Fprintf(os.Stdout, "Response from `CommunicationsCallApi.CommunicationsCallsListAudioRoutingGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callId** | **string** | key: id of call | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommunicationsCallsListAudioRoutingGroupsRequest struct via the builder pattern


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

[**CollectionOfAudioRoutingGroup**](CollectionOfAudioRoutingGroup.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommunicationsCallsListOperations

> CollectionOfCommsOperation CommunicationsCallsListOperations(ctx, callId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get operations from communications



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
    callId := "callId_example" // string | key: id of call
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
    resp, r, err := api_client.CommunicationsCallApi.CommunicationsCallsListOperations(context.Background(), callId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommunicationsCallApi.CommunicationsCallsListOperations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CommunicationsCallsListOperations`: CollectionOfCommsOperation
    fmt.Fprintf(os.Stdout, "Response from `CommunicationsCallApi.CommunicationsCallsListOperations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callId** | **string** | key: id of call | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommunicationsCallsListOperationsRequest struct via the builder pattern


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

[**CollectionOfCommsOperation**](CollectionOfCommsOperation.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommunicationsCallsListParticipants

> CollectionOfParticipant CommunicationsCallsListParticipants(ctx, callId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get participants from communications



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
    callId := "callId_example" // string | key: id of call
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
    resp, r, err := api_client.CommunicationsCallApi.CommunicationsCallsListParticipants(context.Background(), callId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommunicationsCallApi.CommunicationsCallsListParticipants``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CommunicationsCallsListParticipants`: CollectionOfParticipant
    fmt.Fprintf(os.Stdout, "Response from `CommunicationsCallApi.CommunicationsCallsListParticipants`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callId** | **string** | key: id of call | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommunicationsCallsListParticipantsRequest struct via the builder pattern


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

[**CollectionOfParticipant**](CollectionOfParticipant.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommunicationsCallsUpdateAudioRoutingGroups

> CommunicationsCallsUpdateAudioRoutingGroups(ctx, callId, audioRoutingGroupId).MicrosoftGraphAudioRoutingGroup(microsoftGraphAudioRoutingGroup).Execute()

Update the navigation property audioRoutingGroups in communications



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
    callId := "callId_example" // string | key: id of call
    audioRoutingGroupId := "audioRoutingGroupId_example" // string | key: id of audioRoutingGroup
    microsoftGraphAudioRoutingGroup := *openapiclient.NewMicrosoftGraphAudioRoutingGroup() // MicrosoftGraphAudioRoutingGroup | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CommunicationsCallApi.CommunicationsCallsUpdateAudioRoutingGroups(context.Background(), callId, audioRoutingGroupId).MicrosoftGraphAudioRoutingGroup(microsoftGraphAudioRoutingGroup).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommunicationsCallApi.CommunicationsCallsUpdateAudioRoutingGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callId** | **string** | key: id of call | 
**audioRoutingGroupId** | **string** | key: id of audioRoutingGroup | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommunicationsCallsUpdateAudioRoutingGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **microsoftGraphAudioRoutingGroup** | [**MicrosoftGraphAudioRoutingGroup**](MicrosoftGraphAudioRoutingGroup.md) | New navigation property values | 

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


## CommunicationsCallsUpdateOperations

> CommunicationsCallsUpdateOperations(ctx, callId, commsOperationId).MicrosoftGraphCommsOperation(microsoftGraphCommsOperation).Execute()

Update the navigation property operations in communications



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
    callId := "callId_example" // string | key: id of call
    commsOperationId := "commsOperationId_example" // string | key: id of commsOperation
    microsoftGraphCommsOperation := *openapiclient.NewMicrosoftGraphCommsOperation() // MicrosoftGraphCommsOperation | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CommunicationsCallApi.CommunicationsCallsUpdateOperations(context.Background(), callId, commsOperationId).MicrosoftGraphCommsOperation(microsoftGraphCommsOperation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommunicationsCallApi.CommunicationsCallsUpdateOperations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callId** | **string** | key: id of call | 
**commsOperationId** | **string** | key: id of commsOperation | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommunicationsCallsUpdateOperationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **microsoftGraphCommsOperation** | [**MicrosoftGraphCommsOperation**](MicrosoftGraphCommsOperation.md) | New navigation property values | 

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


## CommunicationsCallsUpdateParticipants

> CommunicationsCallsUpdateParticipants(ctx, callId, participantId).MicrosoftGraphParticipant(microsoftGraphParticipant).Execute()

Update the navigation property participants in communications



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
    callId := "callId_example" // string | key: id of call
    participantId := "participantId_example" // string | key: id of participant
    microsoftGraphParticipant := *openapiclient.NewMicrosoftGraphParticipant() // MicrosoftGraphParticipant | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CommunicationsCallApi.CommunicationsCallsUpdateParticipants(context.Background(), callId, participantId).MicrosoftGraphParticipant(microsoftGraphParticipant).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommunicationsCallApi.CommunicationsCallsUpdateParticipants``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callId** | **string** | key: id of call | 
**participantId** | **string** | key: id of participant | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommunicationsCallsUpdateParticipantsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **microsoftGraphParticipant** | [**MicrosoftGraphParticipant**](MicrosoftGraphParticipant.md) | New navigation property values | 

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


## CommunicationsCreateCalls

> MicrosoftGraphCall CommunicationsCreateCalls(ctx).MicrosoftGraphCall(microsoftGraphCall).Execute()

Create new navigation property to calls for communications

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
    microsoftGraphCall := *openapiclient.NewMicrosoftGraphCall() // MicrosoftGraphCall | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CommunicationsCallApi.CommunicationsCreateCalls(context.Background()).MicrosoftGraphCall(microsoftGraphCall).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommunicationsCallApi.CommunicationsCreateCalls``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CommunicationsCreateCalls`: MicrosoftGraphCall
    fmt.Fprintf(os.Stdout, "Response from `CommunicationsCallApi.CommunicationsCreateCalls`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCommunicationsCreateCallsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphCall** | [**MicrosoftGraphCall**](MicrosoftGraphCall.md) | New navigation property | 

### Return type

[**MicrosoftGraphCall**](MicrosoftGraphCall.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommunicationsDeleteCalls

> CommunicationsDeleteCalls(ctx, callId).IfMatch(ifMatch).Execute()

Delete navigation property calls for communications

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
    callId := "callId_example" // string | key: id of call
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CommunicationsCallApi.CommunicationsDeleteCalls(context.Background(), callId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommunicationsCallApi.CommunicationsDeleteCalls``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callId** | **string** | key: id of call | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommunicationsDeleteCallsRequest struct via the builder pattern


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


## CommunicationsGetCalls

> MicrosoftGraphCall CommunicationsGetCalls(ctx, callId).Select_(select_).Expand(expand).Execute()

Get calls from communications

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
    callId := "callId_example" // string | key: id of call
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CommunicationsCallApi.CommunicationsGetCalls(context.Background(), callId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommunicationsCallApi.CommunicationsGetCalls``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CommunicationsGetCalls`: MicrosoftGraphCall
    fmt.Fprintf(os.Stdout, "Response from `CommunicationsCallApi.CommunicationsGetCalls`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callId** | **string** | key: id of call | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommunicationsGetCallsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphCall**](MicrosoftGraphCall.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommunicationsListCalls

> CollectionOfCall CommunicationsListCalls(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get calls from communications

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
    resp, r, err := api_client.CommunicationsCallApi.CommunicationsListCalls(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommunicationsCallApi.CommunicationsListCalls``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CommunicationsListCalls`: CollectionOfCall
    fmt.Fprintf(os.Stdout, "Response from `CommunicationsCallApi.CommunicationsListCalls`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCommunicationsListCallsRequest struct via the builder pattern


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

[**CollectionOfCall**](CollectionOfCall.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommunicationsUpdateCalls

> CommunicationsUpdateCalls(ctx, callId).MicrosoftGraphCall(microsoftGraphCall).Execute()

Update the navigation property calls in communications

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
    callId := "callId_example" // string | key: id of call
    microsoftGraphCall := *openapiclient.NewMicrosoftGraphCall() // MicrosoftGraphCall | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CommunicationsCallApi.CommunicationsUpdateCalls(context.Background(), callId).MicrosoftGraphCall(microsoftGraphCall).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommunicationsCallApi.CommunicationsUpdateCalls``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callId** | **string** | key: id of call | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommunicationsUpdateCallsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphCall** | [**MicrosoftGraphCall**](MicrosoftGraphCall.md) | New navigation property values | 

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

