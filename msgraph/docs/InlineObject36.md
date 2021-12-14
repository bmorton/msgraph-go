# InlineObject36

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Participants** | Pointer to [**[]MicrosoftGraphInvitationParticipantInfo**](MicrosoftGraphInvitationParticipantInfo.md) |  | [optional] 
**ClientContext** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewInlineObject36

`func NewInlineObject36() *InlineObject36`

NewInlineObject36 instantiates a new InlineObject36 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject36WithDefaults

`func NewInlineObject36WithDefaults() *InlineObject36`

NewInlineObject36WithDefaults instantiates a new InlineObject36 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParticipants

`func (o *InlineObject36) GetParticipants() []MicrosoftGraphInvitationParticipantInfo`

GetParticipants returns the Participants field if non-nil, zero value otherwise.

### GetParticipantsOk

`func (o *InlineObject36) GetParticipantsOk() (*[]MicrosoftGraphInvitationParticipantInfo, bool)`

GetParticipantsOk returns a tuple with the Participants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipants

`func (o *InlineObject36) SetParticipants(v []MicrosoftGraphInvitationParticipantInfo)`

SetParticipants sets Participants field to given value.

### HasParticipants

`func (o *InlineObject36) HasParticipants() bool`

HasParticipants returns a boolean if a field has been set.

### GetClientContext

`func (o *InlineObject36) GetClientContext() string`

GetClientContext returns the ClientContext field if non-nil, zero value otherwise.

### GetClientContextOk

`func (o *InlineObject36) GetClientContextOk() (*string, bool)`

GetClientContextOk returns a tuple with the ClientContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientContext

`func (o *InlineObject36) SetClientContext(v string)`

SetClientContext sets ClientContext field to given value.

### HasClientContext

`func (o *InlineObject36) HasClientContext() bool`

HasClientContext returns a boolean if a field has been set.

### SetClientContextNil

`func (o *InlineObject36) SetClientContextNil(b bool)`

 SetClientContextNil sets the value for ClientContext to be an explicit nil

### UnsetClientContext
`func (o *InlineObject36) UnsetClientContext()`

UnsetClientContext ensures that no value is present for ClientContext, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


