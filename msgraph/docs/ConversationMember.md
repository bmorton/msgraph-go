# ConversationMember

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **NullableString** | The display name of the user. | [optional] 
**Roles** | Pointer to **[]string** | The roles for that user. | [optional] 
**VisibleHistoryStartDateTime** | Pointer to **NullableTime** | The timestamp denoting how far back a conversation&#39;s history is shared with the conversation member. This property is settable only for members of a chat. | [optional] 

## Methods

### NewConversationMember

`func NewConversationMember() *ConversationMember`

NewConversationMember instantiates a new ConversationMember object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConversationMemberWithDefaults

`func NewConversationMemberWithDefaults() *ConversationMember`

NewConversationMemberWithDefaults instantiates a new ConversationMember object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *ConversationMember) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ConversationMember) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ConversationMember) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ConversationMember) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *ConversationMember) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *ConversationMember) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetRoles

`func (o *ConversationMember) GetRoles() []*string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *ConversationMember) GetRolesOk() (*[]*string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *ConversationMember) SetRoles(v []*string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *ConversationMember) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetVisibleHistoryStartDateTime

`func (o *ConversationMember) GetVisibleHistoryStartDateTime() time.Time`

GetVisibleHistoryStartDateTime returns the VisibleHistoryStartDateTime field if non-nil, zero value otherwise.

### GetVisibleHistoryStartDateTimeOk

`func (o *ConversationMember) GetVisibleHistoryStartDateTimeOk() (*time.Time, bool)`

GetVisibleHistoryStartDateTimeOk returns a tuple with the VisibleHistoryStartDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibleHistoryStartDateTime

`func (o *ConversationMember) SetVisibleHistoryStartDateTime(v time.Time)`

SetVisibleHistoryStartDateTime sets VisibleHistoryStartDateTime field to given value.

### HasVisibleHistoryStartDateTime

`func (o *ConversationMember) HasVisibleHistoryStartDateTime() bool`

HasVisibleHistoryStartDateTime returns a boolean if a field has been set.

### SetVisibleHistoryStartDateTimeNil

`func (o *ConversationMember) SetVisibleHistoryStartDateTimeNil(b bool)`

 SetVisibleHistoryStartDateTimeNil sets the value for VisibleHistoryStartDateTime to be an explicit nil

### UnsetVisibleHistoryStartDateTime
`func (o *ConversationMember) UnsetVisibleHistoryStartDateTime()`

UnsetVisibleHistoryStartDateTime ensures that no value is present for VisibleHistoryStartDateTime, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


