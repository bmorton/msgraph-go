# \CommunicationsActionsApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CommunicationsCallsCallAnswer**](CommunicationsActionsApi.md#CommunicationsCallsCallAnswer) | **Post** /communications/calls/{call-id}/microsoft.graph.answer | Invoke action answer
[**CommunicationsCallsCallCancelMediaProcessing**](CommunicationsActionsApi.md#CommunicationsCallsCallCancelMediaProcessing) | **Post** /communications/calls/{call-id}/microsoft.graph.cancelMediaProcessing | Invoke action cancelMediaProcessing
[**CommunicationsCallsCallChangeScreenSharingRole**](CommunicationsActionsApi.md#CommunicationsCallsCallChangeScreenSharingRole) | **Post** /communications/calls/{call-id}/microsoft.graph.changeScreenSharingRole | Invoke action changeScreenSharingRole
[**CommunicationsCallsCallKeepAlive**](CommunicationsActionsApi.md#CommunicationsCallsCallKeepAlive) | **Post** /communications/calls/{call-id}/microsoft.graph.keepAlive | Invoke action keepAlive
[**CommunicationsCallsCallMute**](CommunicationsActionsApi.md#CommunicationsCallsCallMute) | **Post** /communications/calls/{call-id}/microsoft.graph.mute | Invoke action mute
[**CommunicationsCallsCallParticipantsInvite**](CommunicationsActionsApi.md#CommunicationsCallsCallParticipantsInvite) | **Post** /communications/calls/{call-id}/participants/microsoft.graph.invite | Invoke action invite
[**CommunicationsCallsCallParticipantsParticipantMute**](CommunicationsActionsApi.md#CommunicationsCallsCallParticipantsParticipantMute) | **Post** /communications/calls/{call-id}/participants/{participant-id}/microsoft.graph.mute | Invoke action mute
[**CommunicationsCallsCallParticipantsParticipantStartHoldMusic**](CommunicationsActionsApi.md#CommunicationsCallsCallParticipantsParticipantStartHoldMusic) | **Post** /communications/calls/{call-id}/participants/{participant-id}/microsoft.graph.startHoldMusic | Invoke action startHoldMusic
[**CommunicationsCallsCallParticipantsParticipantStopHoldMusic**](CommunicationsActionsApi.md#CommunicationsCallsCallParticipantsParticipantStopHoldMusic) | **Post** /communications/calls/{call-id}/participants/{participant-id}/microsoft.graph.stopHoldMusic | Invoke action stopHoldMusic
[**CommunicationsCallsCallPlayPrompt**](CommunicationsActionsApi.md#CommunicationsCallsCallPlayPrompt) | **Post** /communications/calls/{call-id}/microsoft.graph.playPrompt | Invoke action playPrompt
[**CommunicationsCallsCallRecordResponse**](CommunicationsActionsApi.md#CommunicationsCallsCallRecordResponse) | **Post** /communications/calls/{call-id}/microsoft.graph.recordResponse | Invoke action recordResponse
[**CommunicationsCallsCallRedirect**](CommunicationsActionsApi.md#CommunicationsCallsCallRedirect) | **Post** /communications/calls/{call-id}/microsoft.graph.redirect | Invoke action redirect
[**CommunicationsCallsCallReject**](CommunicationsActionsApi.md#CommunicationsCallsCallReject) | **Post** /communications/calls/{call-id}/microsoft.graph.reject | Invoke action reject
[**CommunicationsCallsCallSubscribeToTone**](CommunicationsActionsApi.md#CommunicationsCallsCallSubscribeToTone) | **Post** /communications/calls/{call-id}/microsoft.graph.subscribeToTone | Invoke action subscribeToTone
[**CommunicationsCallsCallTransfer**](CommunicationsActionsApi.md#CommunicationsCallsCallTransfer) | **Post** /communications/calls/{call-id}/microsoft.graph.transfer | Invoke action transfer
[**CommunicationsCallsCallUnmute**](CommunicationsActionsApi.md#CommunicationsCallsCallUnmute) | **Post** /communications/calls/{call-id}/microsoft.graph.unmute | Invoke action unmute
[**CommunicationsCallsCallUpdateRecordingStatus**](CommunicationsActionsApi.md#CommunicationsCallsCallUpdateRecordingStatus) | **Post** /communications/calls/{call-id}/microsoft.graph.updateRecordingStatus | Invoke action updateRecordingStatus
[**CommunicationsCallsLogTeleconferenceDeviceQuality**](CommunicationsActionsApi.md#CommunicationsCallsLogTeleconferenceDeviceQuality) | **Post** /communications/calls/microsoft.graph.logTeleconferenceDeviceQuality | Invoke action logTeleconferenceDeviceQuality
[**CommunicationsGetPresencesByUserId**](CommunicationsActionsApi.md#CommunicationsGetPresencesByUserId) | **Post** /communications/microsoft.graph.getPresencesByUserId | Invoke action getPresencesByUserId
[**CommunicationsOnlineMeetingsCreateOrGet**](CommunicationsActionsApi.md#CommunicationsOnlineMeetingsCreateOrGet) | **Post** /communications/onlineMeetings/microsoft.graph.createOrGet | Invoke action createOrGet
[**CommunicationsPresencesPresenceClearPresence**](CommunicationsActionsApi.md#CommunicationsPresencesPresenceClearPresence) | **Post** /communications/presences/{presence-id}/microsoft.graph.clearPresence | Invoke action clearPresence
[**CommunicationsPresencesPresenceSetPresence**](CommunicationsActionsApi.md#CommunicationsPresencesPresenceSetPresence) | **Post** /communications/presences/{presence-id}/microsoft.graph.setPresence | Invoke action setPresence



## CommunicationsCallsCallAnswer

> CommunicationsCallsCallAnswer(ctx, callId).InlineObject21(inlineObject21).Execute()

Invoke action answer

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
    inlineObject21 := *openapiclient.NewInlineObject21() // InlineObject21 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CommunicationsActionsApi.CommunicationsCallsCallAnswer(context.Background(), callId).InlineObject21(inlineObject21).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommunicationsActionsApi.CommunicationsCallsCallAnswer``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCommunicationsCallsCallAnswerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject21** | [**InlineObject21**](InlineObject21.md) |  | 

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


## CommunicationsCallsCallCancelMediaProcessing

> AnyOfmicrosoftGraphCancelMediaProcessingOperation CommunicationsCallsCallCancelMediaProcessing(ctx, callId).InlineObject22(inlineObject22).Execute()

Invoke action cancelMediaProcessing

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
    inlineObject22 := *openapiclient.NewInlineObject22() // InlineObject22 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CommunicationsActionsApi.CommunicationsCallsCallCancelMediaProcessing(context.Background(), callId).InlineObject22(inlineObject22).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommunicationsActionsApi.CommunicationsCallsCallCancelMediaProcessing``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CommunicationsCallsCallCancelMediaProcessing`: AnyOfmicrosoftGraphCancelMediaProcessingOperation
    fmt.Fprintf(os.Stdout, "Response from `CommunicationsActionsApi.CommunicationsCallsCallCancelMediaProcessing`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callId** | **string** | key: id of call | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommunicationsCallsCallCancelMediaProcessingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject22** | [**InlineObject22**](InlineObject22.md) |  | 

### Return type

[**AnyOfmicrosoftGraphCancelMediaProcessingOperation**](anyOf&lt;microsoft.graph.cancelMediaProcessingOperation&gt;.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommunicationsCallsCallChangeScreenSharingRole

> CommunicationsCallsCallChangeScreenSharingRole(ctx, callId).InlineObject23(inlineObject23).Execute()

Invoke action changeScreenSharingRole

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
    inlineObject23 := *openapiclient.NewInlineObject23() // InlineObject23 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CommunicationsActionsApi.CommunicationsCallsCallChangeScreenSharingRole(context.Background(), callId).InlineObject23(inlineObject23).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommunicationsActionsApi.CommunicationsCallsCallChangeScreenSharingRole``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCommunicationsCallsCallChangeScreenSharingRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject23** | [**InlineObject23**](InlineObject23.md) |  | 

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


## CommunicationsCallsCallKeepAlive

> CommunicationsCallsCallKeepAlive(ctx, callId).Execute()

Invoke action keepAlive

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CommunicationsActionsApi.CommunicationsCallsCallKeepAlive(context.Background(), callId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommunicationsActionsApi.CommunicationsCallsCallKeepAlive``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCommunicationsCallsCallKeepAliveRequest struct via the builder pattern


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


## CommunicationsCallsCallMute

> AnyOfmicrosoftGraphMuteParticipantOperation CommunicationsCallsCallMute(ctx, callId).InlineObject24(inlineObject24).Execute()

Invoke action mute

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
    inlineObject24 := *openapiclient.NewInlineObject24() // InlineObject24 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CommunicationsActionsApi.CommunicationsCallsCallMute(context.Background(), callId).InlineObject24(inlineObject24).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommunicationsActionsApi.CommunicationsCallsCallMute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CommunicationsCallsCallMute`: AnyOfmicrosoftGraphMuteParticipantOperation
    fmt.Fprintf(os.Stdout, "Response from `CommunicationsActionsApi.CommunicationsCallsCallMute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callId** | **string** | key: id of call | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommunicationsCallsCallMuteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject24** | [**InlineObject24**](InlineObject24.md) |  | 

### Return type

[**AnyOfmicrosoftGraphMuteParticipantOperation**](anyOf&lt;microsoft.graph.muteParticipantOperation&gt;.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommunicationsCallsCallParticipantsInvite

> AnyOfmicrosoftGraphInviteParticipantsOperation CommunicationsCallsCallParticipantsInvite(ctx, callId).InlineObject36(inlineObject36).Execute()

Invoke action invite

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
    inlineObject36 := *openapiclient.NewInlineObject36() // InlineObject36 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CommunicationsActionsApi.CommunicationsCallsCallParticipantsInvite(context.Background(), callId).InlineObject36(inlineObject36).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommunicationsActionsApi.CommunicationsCallsCallParticipantsInvite``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CommunicationsCallsCallParticipantsInvite`: AnyOfmicrosoftGraphInviteParticipantsOperation
    fmt.Fprintf(os.Stdout, "Response from `CommunicationsActionsApi.CommunicationsCallsCallParticipantsInvite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callId** | **string** | key: id of call | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommunicationsCallsCallParticipantsInviteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject36** | [**InlineObject36**](InlineObject36.md) |  | 

### Return type

[**AnyOfmicrosoftGraphInviteParticipantsOperation**](anyOf&lt;microsoft.graph.inviteParticipantsOperation&gt;.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommunicationsCallsCallParticipantsParticipantMute

> AnyOfmicrosoftGraphMuteParticipantOperation CommunicationsCallsCallParticipantsParticipantMute(ctx, callId, participantId).InlineObject33(inlineObject33).Execute()

Invoke action mute

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
    inlineObject33 := *openapiclient.NewInlineObject33() // InlineObject33 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CommunicationsActionsApi.CommunicationsCallsCallParticipantsParticipantMute(context.Background(), callId, participantId).InlineObject33(inlineObject33).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommunicationsActionsApi.CommunicationsCallsCallParticipantsParticipantMute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CommunicationsCallsCallParticipantsParticipantMute`: AnyOfmicrosoftGraphMuteParticipantOperation
    fmt.Fprintf(os.Stdout, "Response from `CommunicationsActionsApi.CommunicationsCallsCallParticipantsParticipantMute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callId** | **string** | key: id of call | 
**participantId** | **string** | key: id of participant | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommunicationsCallsCallParticipantsParticipantMuteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **inlineObject33** | [**InlineObject33**](InlineObject33.md) |  | 

### Return type

[**AnyOfmicrosoftGraphMuteParticipantOperation**](anyOf&lt;microsoft.graph.muteParticipantOperation&gt;.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommunicationsCallsCallParticipantsParticipantStartHoldMusic

> AnyOfmicrosoftGraphStartHoldMusicOperation CommunicationsCallsCallParticipantsParticipantStartHoldMusic(ctx, callId, participantId).InlineObject34(inlineObject34).Execute()

Invoke action startHoldMusic

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
    inlineObject34 := *openapiclient.NewInlineObject34() // InlineObject34 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CommunicationsActionsApi.CommunicationsCallsCallParticipantsParticipantStartHoldMusic(context.Background(), callId, participantId).InlineObject34(inlineObject34).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommunicationsActionsApi.CommunicationsCallsCallParticipantsParticipantStartHoldMusic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CommunicationsCallsCallParticipantsParticipantStartHoldMusic`: AnyOfmicrosoftGraphStartHoldMusicOperation
    fmt.Fprintf(os.Stdout, "Response from `CommunicationsActionsApi.CommunicationsCallsCallParticipantsParticipantStartHoldMusic`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callId** | **string** | key: id of call | 
**participantId** | **string** | key: id of participant | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommunicationsCallsCallParticipantsParticipantStartHoldMusicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **inlineObject34** | [**InlineObject34**](InlineObject34.md) |  | 

### Return type

[**AnyOfmicrosoftGraphStartHoldMusicOperation**](anyOf&lt;microsoft.graph.startHoldMusicOperation&gt;.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommunicationsCallsCallParticipantsParticipantStopHoldMusic

> AnyOfmicrosoftGraphStopHoldMusicOperation CommunicationsCallsCallParticipantsParticipantStopHoldMusic(ctx, callId, participantId).InlineObject35(inlineObject35).Execute()

Invoke action stopHoldMusic

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
    inlineObject35 := *openapiclient.NewInlineObject35() // InlineObject35 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CommunicationsActionsApi.CommunicationsCallsCallParticipantsParticipantStopHoldMusic(context.Background(), callId, participantId).InlineObject35(inlineObject35).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommunicationsActionsApi.CommunicationsCallsCallParticipantsParticipantStopHoldMusic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CommunicationsCallsCallParticipantsParticipantStopHoldMusic`: AnyOfmicrosoftGraphStopHoldMusicOperation
    fmt.Fprintf(os.Stdout, "Response from `CommunicationsActionsApi.CommunicationsCallsCallParticipantsParticipantStopHoldMusic`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callId** | **string** | key: id of call | 
**participantId** | **string** | key: id of participant | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommunicationsCallsCallParticipantsParticipantStopHoldMusicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **inlineObject35** | [**InlineObject35**](InlineObject35.md) |  | 

### Return type

[**AnyOfmicrosoftGraphStopHoldMusicOperation**](anyOf&lt;microsoft.graph.stopHoldMusicOperation&gt;.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommunicationsCallsCallPlayPrompt

> AnyOfmicrosoftGraphPlayPromptOperation CommunicationsCallsCallPlayPrompt(ctx, callId).InlineObject25(inlineObject25).Execute()

Invoke action playPrompt

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
    inlineObject25 := *openapiclient.NewInlineObject25() // InlineObject25 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CommunicationsActionsApi.CommunicationsCallsCallPlayPrompt(context.Background(), callId).InlineObject25(inlineObject25).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommunicationsActionsApi.CommunicationsCallsCallPlayPrompt``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CommunicationsCallsCallPlayPrompt`: AnyOfmicrosoftGraphPlayPromptOperation
    fmt.Fprintf(os.Stdout, "Response from `CommunicationsActionsApi.CommunicationsCallsCallPlayPrompt`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callId** | **string** | key: id of call | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommunicationsCallsCallPlayPromptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject25** | [**InlineObject25**](InlineObject25.md) |  | 

### Return type

[**AnyOfmicrosoftGraphPlayPromptOperation**](anyOf&lt;microsoft.graph.playPromptOperation&gt;.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommunicationsCallsCallRecordResponse

> AnyOfmicrosoftGraphRecordOperation CommunicationsCallsCallRecordResponse(ctx, callId).InlineObject26(inlineObject26).Execute()

Invoke action recordResponse

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
    inlineObject26 := *openapiclient.NewInlineObject26() // InlineObject26 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CommunicationsActionsApi.CommunicationsCallsCallRecordResponse(context.Background(), callId).InlineObject26(inlineObject26).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommunicationsActionsApi.CommunicationsCallsCallRecordResponse``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CommunicationsCallsCallRecordResponse`: AnyOfmicrosoftGraphRecordOperation
    fmt.Fprintf(os.Stdout, "Response from `CommunicationsActionsApi.CommunicationsCallsCallRecordResponse`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callId** | **string** | key: id of call | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommunicationsCallsCallRecordResponseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject26** | [**InlineObject26**](InlineObject26.md) |  | 

### Return type

[**AnyOfmicrosoftGraphRecordOperation**](anyOf&lt;microsoft.graph.recordOperation&gt;.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommunicationsCallsCallRedirect

> CommunicationsCallsCallRedirect(ctx, callId).InlineObject27(inlineObject27).Execute()

Invoke action redirect

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
    inlineObject27 := *openapiclient.NewInlineObject27() // InlineObject27 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CommunicationsActionsApi.CommunicationsCallsCallRedirect(context.Background(), callId).InlineObject27(inlineObject27).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommunicationsActionsApi.CommunicationsCallsCallRedirect``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCommunicationsCallsCallRedirectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject27** | [**InlineObject27**](InlineObject27.md) |  | 

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


## CommunicationsCallsCallReject

> CommunicationsCallsCallReject(ctx, callId).InlineObject28(inlineObject28).Execute()

Invoke action reject

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
    inlineObject28 := *openapiclient.NewInlineObject28() // InlineObject28 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CommunicationsActionsApi.CommunicationsCallsCallReject(context.Background(), callId).InlineObject28(inlineObject28).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommunicationsActionsApi.CommunicationsCallsCallReject``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCommunicationsCallsCallRejectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject28** | [**InlineObject28**](InlineObject28.md) |  | 

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


## CommunicationsCallsCallSubscribeToTone

> AnyOfmicrosoftGraphSubscribeToToneOperation CommunicationsCallsCallSubscribeToTone(ctx, callId).InlineObject29(inlineObject29).Execute()

Invoke action subscribeToTone

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
    inlineObject29 := *openapiclient.NewInlineObject29() // InlineObject29 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CommunicationsActionsApi.CommunicationsCallsCallSubscribeToTone(context.Background(), callId).InlineObject29(inlineObject29).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommunicationsActionsApi.CommunicationsCallsCallSubscribeToTone``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CommunicationsCallsCallSubscribeToTone`: AnyOfmicrosoftGraphSubscribeToToneOperation
    fmt.Fprintf(os.Stdout, "Response from `CommunicationsActionsApi.CommunicationsCallsCallSubscribeToTone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callId** | **string** | key: id of call | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommunicationsCallsCallSubscribeToToneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject29** | [**InlineObject29**](InlineObject29.md) |  | 

### Return type

[**AnyOfmicrosoftGraphSubscribeToToneOperation**](anyOf&lt;microsoft.graph.subscribeToToneOperation&gt;.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommunicationsCallsCallTransfer

> CommunicationsCallsCallTransfer(ctx, callId).InlineObject30(inlineObject30).Execute()

Invoke action transfer

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
    inlineObject30 := *openapiclient.NewInlineObject30() // InlineObject30 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CommunicationsActionsApi.CommunicationsCallsCallTransfer(context.Background(), callId).InlineObject30(inlineObject30).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommunicationsActionsApi.CommunicationsCallsCallTransfer``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCommunicationsCallsCallTransferRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject30** | [**InlineObject30**](InlineObject30.md) |  | 

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


## CommunicationsCallsCallUnmute

> AnyOfmicrosoftGraphUnmuteParticipantOperation CommunicationsCallsCallUnmute(ctx, callId).InlineObject31(inlineObject31).Execute()

Invoke action unmute

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
    inlineObject31 := *openapiclient.NewInlineObject31() // InlineObject31 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CommunicationsActionsApi.CommunicationsCallsCallUnmute(context.Background(), callId).InlineObject31(inlineObject31).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommunicationsActionsApi.CommunicationsCallsCallUnmute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CommunicationsCallsCallUnmute`: AnyOfmicrosoftGraphUnmuteParticipantOperation
    fmt.Fprintf(os.Stdout, "Response from `CommunicationsActionsApi.CommunicationsCallsCallUnmute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callId** | **string** | key: id of call | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommunicationsCallsCallUnmuteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject31** | [**InlineObject31**](InlineObject31.md) |  | 

### Return type

[**AnyOfmicrosoftGraphUnmuteParticipantOperation**](anyOf&lt;microsoft.graph.unmuteParticipantOperation&gt;.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommunicationsCallsCallUpdateRecordingStatus

> AnyOfmicrosoftGraphUpdateRecordingStatusOperation CommunicationsCallsCallUpdateRecordingStatus(ctx, callId).InlineObject32(inlineObject32).Execute()

Invoke action updateRecordingStatus

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
    inlineObject32 := *openapiclient.NewInlineObject32() // InlineObject32 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CommunicationsActionsApi.CommunicationsCallsCallUpdateRecordingStatus(context.Background(), callId).InlineObject32(inlineObject32).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommunicationsActionsApi.CommunicationsCallsCallUpdateRecordingStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CommunicationsCallsCallUpdateRecordingStatus`: AnyOfmicrosoftGraphUpdateRecordingStatusOperation
    fmt.Fprintf(os.Stdout, "Response from `CommunicationsActionsApi.CommunicationsCallsCallUpdateRecordingStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callId** | **string** | key: id of call | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommunicationsCallsCallUpdateRecordingStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject32** | [**InlineObject32**](InlineObject32.md) |  | 

### Return type

[**AnyOfmicrosoftGraphUpdateRecordingStatusOperation**](anyOf&lt;microsoft.graph.updateRecordingStatusOperation&gt;.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommunicationsCallsLogTeleconferenceDeviceQuality

> CommunicationsCallsLogTeleconferenceDeviceQuality(ctx).InlineObject37(inlineObject37).Execute()

Invoke action logTeleconferenceDeviceQuality

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
    inlineObject37 := *openapiclient.NewInlineObject37() // InlineObject37 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CommunicationsActionsApi.CommunicationsCallsLogTeleconferenceDeviceQuality(context.Background()).InlineObject37(inlineObject37).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommunicationsActionsApi.CommunicationsCallsLogTeleconferenceDeviceQuality``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCommunicationsCallsLogTeleconferenceDeviceQualityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject37** | [**InlineObject37**](InlineObject37.md) |  | 

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


## CommunicationsGetPresencesByUserId

> []*AnyOfmicrosoftGraphPresence CommunicationsGetPresencesByUserId(ctx).InlineObject38(inlineObject38).Execute()

Invoke action getPresencesByUserId

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
    inlineObject38 := *openapiclient.NewInlineObject38() // InlineObject38 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CommunicationsActionsApi.CommunicationsGetPresencesByUserId(context.Background()).InlineObject38(inlineObject38).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommunicationsActionsApi.CommunicationsGetPresencesByUserId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CommunicationsGetPresencesByUserId`: []*AnyOfmicrosoftGraphPresence
    fmt.Fprintf(os.Stdout, "Response from `CommunicationsActionsApi.CommunicationsGetPresencesByUserId`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCommunicationsGetPresencesByUserIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject38** | [**InlineObject38**](InlineObject38.md) |  | 

### Return type

[**[]*AnyOfmicrosoftGraphPresence**](anyOf&lt;microsoft.graph.presence&gt;.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommunicationsOnlineMeetingsCreateOrGet

> AnyOfmicrosoftGraphOnlineMeeting CommunicationsOnlineMeetingsCreateOrGet(ctx).InlineObject39(inlineObject39).Execute()

Invoke action createOrGet

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
    inlineObject39 := *openapiclient.NewInlineObject39() // InlineObject39 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CommunicationsActionsApi.CommunicationsOnlineMeetingsCreateOrGet(context.Background()).InlineObject39(inlineObject39).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommunicationsActionsApi.CommunicationsOnlineMeetingsCreateOrGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CommunicationsOnlineMeetingsCreateOrGet`: AnyOfmicrosoftGraphOnlineMeeting
    fmt.Fprintf(os.Stdout, "Response from `CommunicationsActionsApi.CommunicationsOnlineMeetingsCreateOrGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCommunicationsOnlineMeetingsCreateOrGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject39** | [**InlineObject39**](InlineObject39.md) |  | 

### Return type

[**AnyOfmicrosoftGraphOnlineMeeting**](anyOf&lt;microsoft.graph.onlineMeeting&gt;.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommunicationsPresencesPresenceClearPresence

> CommunicationsPresencesPresenceClearPresence(ctx, presenceId).InlineObject40(inlineObject40).Execute()

Invoke action clearPresence

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
    presenceId := "presenceId_example" // string | key: id of presence
    inlineObject40 := *openapiclient.NewInlineObject40() // InlineObject40 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CommunicationsActionsApi.CommunicationsPresencesPresenceClearPresence(context.Background(), presenceId).InlineObject40(inlineObject40).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommunicationsActionsApi.CommunicationsPresencesPresenceClearPresence``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**presenceId** | **string** | key: id of presence | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommunicationsPresencesPresenceClearPresenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject40** | [**InlineObject40**](InlineObject40.md) |  | 

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


## CommunicationsPresencesPresenceSetPresence

> CommunicationsPresencesPresenceSetPresence(ctx, presenceId).InlineObject41(inlineObject41).Execute()

Invoke action setPresence

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
    presenceId := "presenceId_example" // string | key: id of presence
    inlineObject41 := *openapiclient.NewInlineObject41() // InlineObject41 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CommunicationsActionsApi.CommunicationsPresencesPresenceSetPresence(context.Background(), presenceId).InlineObject41(inlineObject41).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommunicationsActionsApi.CommunicationsPresencesPresenceSetPresence``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**presenceId** | **string** | key: id of presence | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommunicationsPresencesPresenceSetPresenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject41** | [**InlineObject41**](InlineObject41.md) |  | 

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

