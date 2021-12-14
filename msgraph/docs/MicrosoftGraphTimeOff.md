# MicrosoftGraphTimeOff

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**CreatedDateTime** | Pointer to **NullableTime** | The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z | [optional] 
**LastModifiedBy** | Pointer to [**AnyOfmicrosoftGraphIdentitySet**](anyOf&lt;microsoft.graph.identitySet&gt;.md) | Identity of the person who last modified the entity. | [optional] 
**LastModifiedDateTime** | Pointer to **NullableTime** | The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z | [optional] 
**DraftTimeOff** | Pointer to [**AnyOfmicrosoftGraphTimeOffItem**](anyOf&lt;microsoft.graph.timeOffItem&gt;.md) | The draft version of this timeOff that is viewable by managers. Required. | [optional] 
**SharedTimeOff** | Pointer to [**AnyOfmicrosoftGraphTimeOffItem**](anyOf&lt;microsoft.graph.timeOffItem&gt;.md) | The shared version of this timeOff that is viewable by both employees and managers. Required. | [optional] 
**UserId** | Pointer to **NullableString** | ID of the user assigned to the timeOff. Required. | [optional] 

## Methods

### NewMicrosoftGraphTimeOff

`func NewMicrosoftGraphTimeOff() *MicrosoftGraphTimeOff`

NewMicrosoftGraphTimeOff instantiates a new MicrosoftGraphTimeOff object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphTimeOffWithDefaults

`func NewMicrosoftGraphTimeOffWithDefaults() *MicrosoftGraphTimeOff`

NewMicrosoftGraphTimeOffWithDefaults instantiates a new MicrosoftGraphTimeOff object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphTimeOff) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphTimeOff) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphTimeOff) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphTimeOff) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedDateTime

`func (o *MicrosoftGraphTimeOff) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphTimeOff) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphTimeOff) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *MicrosoftGraphTimeOff) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTimeNil

`func (o *MicrosoftGraphTimeOff) SetCreatedDateTimeNil(b bool)`

 SetCreatedDateTimeNil sets the value for CreatedDateTime to be an explicit nil

### UnsetCreatedDateTime
`func (o *MicrosoftGraphTimeOff) UnsetCreatedDateTime()`

UnsetCreatedDateTime ensures that no value is present for CreatedDateTime, not even an explicit nil
### GetLastModifiedBy

`func (o *MicrosoftGraphTimeOff) GetLastModifiedBy() AnyOfmicrosoftGraphIdentitySet`

GetLastModifiedBy returns the LastModifiedBy field if non-nil, zero value otherwise.

### GetLastModifiedByOk

`func (o *MicrosoftGraphTimeOff) GetLastModifiedByOk() (*AnyOfmicrosoftGraphIdentitySet, bool)`

GetLastModifiedByOk returns a tuple with the LastModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedBy

`func (o *MicrosoftGraphTimeOff) SetLastModifiedBy(v AnyOfmicrosoftGraphIdentitySet)`

SetLastModifiedBy sets LastModifiedBy field to given value.

### HasLastModifiedBy

`func (o *MicrosoftGraphTimeOff) HasLastModifiedBy() bool`

HasLastModifiedBy returns a boolean if a field has been set.

### SetLastModifiedByNil

`func (o *MicrosoftGraphTimeOff) SetLastModifiedByNil(b bool)`

 SetLastModifiedByNil sets the value for LastModifiedBy to be an explicit nil

### UnsetLastModifiedBy
`func (o *MicrosoftGraphTimeOff) UnsetLastModifiedBy()`

UnsetLastModifiedBy ensures that no value is present for LastModifiedBy, not even an explicit nil
### GetLastModifiedDateTime

`func (o *MicrosoftGraphTimeOff) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *MicrosoftGraphTimeOff) GetLastModifiedDateTimeOk() (*time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedDateTime

`func (o *MicrosoftGraphTimeOff) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime sets LastModifiedDateTime field to given value.

### HasLastModifiedDateTime

`func (o *MicrosoftGraphTimeOff) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### SetLastModifiedDateTimeNil

`func (o *MicrosoftGraphTimeOff) SetLastModifiedDateTimeNil(b bool)`

 SetLastModifiedDateTimeNil sets the value for LastModifiedDateTime to be an explicit nil

### UnsetLastModifiedDateTime
`func (o *MicrosoftGraphTimeOff) UnsetLastModifiedDateTime()`

UnsetLastModifiedDateTime ensures that no value is present for LastModifiedDateTime, not even an explicit nil
### GetDraftTimeOff

`func (o *MicrosoftGraphTimeOff) GetDraftTimeOff() AnyOfmicrosoftGraphTimeOffItem`

GetDraftTimeOff returns the DraftTimeOff field if non-nil, zero value otherwise.

### GetDraftTimeOffOk

`func (o *MicrosoftGraphTimeOff) GetDraftTimeOffOk() (*AnyOfmicrosoftGraphTimeOffItem, bool)`

GetDraftTimeOffOk returns a tuple with the DraftTimeOff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraftTimeOff

`func (o *MicrosoftGraphTimeOff) SetDraftTimeOff(v AnyOfmicrosoftGraphTimeOffItem)`

SetDraftTimeOff sets DraftTimeOff field to given value.

### HasDraftTimeOff

`func (o *MicrosoftGraphTimeOff) HasDraftTimeOff() bool`

HasDraftTimeOff returns a boolean if a field has been set.

### SetDraftTimeOffNil

`func (o *MicrosoftGraphTimeOff) SetDraftTimeOffNil(b bool)`

 SetDraftTimeOffNil sets the value for DraftTimeOff to be an explicit nil

### UnsetDraftTimeOff
`func (o *MicrosoftGraphTimeOff) UnsetDraftTimeOff()`

UnsetDraftTimeOff ensures that no value is present for DraftTimeOff, not even an explicit nil
### GetSharedTimeOff

`func (o *MicrosoftGraphTimeOff) GetSharedTimeOff() AnyOfmicrosoftGraphTimeOffItem`

GetSharedTimeOff returns the SharedTimeOff field if non-nil, zero value otherwise.

### GetSharedTimeOffOk

`func (o *MicrosoftGraphTimeOff) GetSharedTimeOffOk() (*AnyOfmicrosoftGraphTimeOffItem, bool)`

GetSharedTimeOffOk returns a tuple with the SharedTimeOff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedTimeOff

`func (o *MicrosoftGraphTimeOff) SetSharedTimeOff(v AnyOfmicrosoftGraphTimeOffItem)`

SetSharedTimeOff sets SharedTimeOff field to given value.

### HasSharedTimeOff

`func (o *MicrosoftGraphTimeOff) HasSharedTimeOff() bool`

HasSharedTimeOff returns a boolean if a field has been set.

### SetSharedTimeOffNil

`func (o *MicrosoftGraphTimeOff) SetSharedTimeOffNil(b bool)`

 SetSharedTimeOffNil sets the value for SharedTimeOff to be an explicit nil

### UnsetSharedTimeOff
`func (o *MicrosoftGraphTimeOff) UnsetSharedTimeOff()`

UnsetSharedTimeOff ensures that no value is present for SharedTimeOff, not even an explicit nil
### GetUserId

`func (o *MicrosoftGraphTimeOff) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *MicrosoftGraphTimeOff) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *MicrosoftGraphTimeOff) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *MicrosoftGraphTimeOff) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *MicrosoftGraphTimeOff) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *MicrosoftGraphTimeOff) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


