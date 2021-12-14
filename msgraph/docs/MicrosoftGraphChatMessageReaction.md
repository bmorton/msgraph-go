# MicrosoftGraphChatMessageReaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedDateTime** | Pointer to **time.Time** | The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z | [optional] 
**ReactionType** | Pointer to **string** | Supported values are like, angry, sad, laugh, heart, surprised. | [optional] 
**User** | Pointer to [**MicrosoftGraphChatMessageReactionIdentitySet**](MicrosoftGraphChatMessageReactionIdentitySet.md) |  | [optional] 

## Methods

### NewMicrosoftGraphChatMessageReaction

`func NewMicrosoftGraphChatMessageReaction() *MicrosoftGraphChatMessageReaction`

NewMicrosoftGraphChatMessageReaction instantiates a new MicrosoftGraphChatMessageReaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphChatMessageReactionWithDefaults

`func NewMicrosoftGraphChatMessageReactionWithDefaults() *MicrosoftGraphChatMessageReaction`

NewMicrosoftGraphChatMessageReactionWithDefaults instantiates a new MicrosoftGraphChatMessageReaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedDateTime

`func (o *MicrosoftGraphChatMessageReaction) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphChatMessageReaction) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphChatMessageReaction) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *MicrosoftGraphChatMessageReaction) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### GetReactionType

`func (o *MicrosoftGraphChatMessageReaction) GetReactionType() string`

GetReactionType returns the ReactionType field if non-nil, zero value otherwise.

### GetReactionTypeOk

`func (o *MicrosoftGraphChatMessageReaction) GetReactionTypeOk() (*string, bool)`

GetReactionTypeOk returns a tuple with the ReactionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReactionType

`func (o *MicrosoftGraphChatMessageReaction) SetReactionType(v string)`

SetReactionType sets ReactionType field to given value.

### HasReactionType

`func (o *MicrosoftGraphChatMessageReaction) HasReactionType() bool`

HasReactionType returns a boolean if a field has been set.

### GetUser

`func (o *MicrosoftGraphChatMessageReaction) GetUser() MicrosoftGraphChatMessageReactionIdentitySet`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *MicrosoftGraphChatMessageReaction) GetUserOk() (*MicrosoftGraphChatMessageReactionIdentitySet, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *MicrosoftGraphChatMessageReaction) SetUser(v MicrosoftGraphChatMessageReactionIdentitySet)`

SetUser sets User field to given value.

### HasUser

`func (o *MicrosoftGraphChatMessageReaction) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


