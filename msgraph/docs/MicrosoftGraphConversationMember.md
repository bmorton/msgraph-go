# MicrosoftGraphConversationMember

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**DisplayName** | Pointer to **NullableString** | The display name of the user. | [optional] 
**Roles** | Pointer to **[]string** | The roles for that user. | [optional] 
**VisibleHistoryStartDateTime** | Pointer to **NullableTime** | The timestamp denoting how far back a conversation&#39;s history is shared with the conversation member. This property is settable only for members of a chat. | [optional] 

## Methods

### NewMicrosoftGraphConversationMember

`func NewMicrosoftGraphConversationMember() *MicrosoftGraphConversationMember`

NewMicrosoftGraphConversationMember instantiates a new MicrosoftGraphConversationMember object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphConversationMemberWithDefaults

`func NewMicrosoftGraphConversationMemberWithDefaults() *MicrosoftGraphConversationMember`

NewMicrosoftGraphConversationMemberWithDefaults instantiates a new MicrosoftGraphConversationMember object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphConversationMember) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphConversationMember) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphConversationMember) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphConversationMember) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDisplayName

`func (o *MicrosoftGraphConversationMember) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphConversationMember) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MicrosoftGraphConversationMember) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MicrosoftGraphConversationMember) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *MicrosoftGraphConversationMember) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *MicrosoftGraphConversationMember) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetRoles

`func (o *MicrosoftGraphConversationMember) GetRoles() []*string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *MicrosoftGraphConversationMember) GetRolesOk() (*[]*string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *MicrosoftGraphConversationMember) SetRoles(v []*string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *MicrosoftGraphConversationMember) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetVisibleHistoryStartDateTime

`func (o *MicrosoftGraphConversationMember) GetVisibleHistoryStartDateTime() time.Time`

GetVisibleHistoryStartDateTime returns the VisibleHistoryStartDateTime field if non-nil, zero value otherwise.

### GetVisibleHistoryStartDateTimeOk

`func (o *MicrosoftGraphConversationMember) GetVisibleHistoryStartDateTimeOk() (*time.Time, bool)`

GetVisibleHistoryStartDateTimeOk returns a tuple with the VisibleHistoryStartDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibleHistoryStartDateTime

`func (o *MicrosoftGraphConversationMember) SetVisibleHistoryStartDateTime(v time.Time)`

SetVisibleHistoryStartDateTime sets VisibleHistoryStartDateTime field to given value.

### HasVisibleHistoryStartDateTime

`func (o *MicrosoftGraphConversationMember) HasVisibleHistoryStartDateTime() bool`

HasVisibleHistoryStartDateTime returns a boolean if a field has been set.

### SetVisibleHistoryStartDateTimeNil

`func (o *MicrosoftGraphConversationMember) SetVisibleHistoryStartDateTimeNil(b bool)`

 SetVisibleHistoryStartDateTimeNil sets the value for VisibleHistoryStartDateTime to be an explicit nil

### UnsetVisibleHistoryStartDateTime
`func (o *MicrosoftGraphConversationMember) UnsetVisibleHistoryStartDateTime()`

UnsetVisibleHistoryStartDateTime ensures that no value is present for VisibleHistoryStartDateTime, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


