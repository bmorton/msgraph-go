# MicrosoftGraphVerifiedPublisher

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddedDateTime** | Pointer to **NullableTime** | The timestamp when the verified publisher was first added or most recently updated. | [optional] 
**DisplayName** | Pointer to **NullableString** | The verified publisher name from the app publisher&#39;s Partner Center account. | [optional] 
**VerifiedPublisherId** | Pointer to **NullableString** | The ID of the verified publisher from the app publisher&#39;s Partner Center account. | [optional] 

## Methods

### NewMicrosoftGraphVerifiedPublisher

`func NewMicrosoftGraphVerifiedPublisher() *MicrosoftGraphVerifiedPublisher`

NewMicrosoftGraphVerifiedPublisher instantiates a new MicrosoftGraphVerifiedPublisher object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphVerifiedPublisherWithDefaults

`func NewMicrosoftGraphVerifiedPublisherWithDefaults() *MicrosoftGraphVerifiedPublisher`

NewMicrosoftGraphVerifiedPublisherWithDefaults instantiates a new MicrosoftGraphVerifiedPublisher object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddedDateTime

`func (o *MicrosoftGraphVerifiedPublisher) GetAddedDateTime() time.Time`

GetAddedDateTime returns the AddedDateTime field if non-nil, zero value otherwise.

### GetAddedDateTimeOk

`func (o *MicrosoftGraphVerifiedPublisher) GetAddedDateTimeOk() (*time.Time, bool)`

GetAddedDateTimeOk returns a tuple with the AddedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedDateTime

`func (o *MicrosoftGraphVerifiedPublisher) SetAddedDateTime(v time.Time)`

SetAddedDateTime sets AddedDateTime field to given value.

### HasAddedDateTime

`func (o *MicrosoftGraphVerifiedPublisher) HasAddedDateTime() bool`

HasAddedDateTime returns a boolean if a field has been set.

### SetAddedDateTimeNil

`func (o *MicrosoftGraphVerifiedPublisher) SetAddedDateTimeNil(b bool)`

 SetAddedDateTimeNil sets the value for AddedDateTime to be an explicit nil

### UnsetAddedDateTime
`func (o *MicrosoftGraphVerifiedPublisher) UnsetAddedDateTime()`

UnsetAddedDateTime ensures that no value is present for AddedDateTime, not even an explicit nil
### GetDisplayName

`func (o *MicrosoftGraphVerifiedPublisher) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphVerifiedPublisher) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MicrosoftGraphVerifiedPublisher) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MicrosoftGraphVerifiedPublisher) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *MicrosoftGraphVerifiedPublisher) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *MicrosoftGraphVerifiedPublisher) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetVerifiedPublisherId

`func (o *MicrosoftGraphVerifiedPublisher) GetVerifiedPublisherId() string`

GetVerifiedPublisherId returns the VerifiedPublisherId field if non-nil, zero value otherwise.

### GetVerifiedPublisherIdOk

`func (o *MicrosoftGraphVerifiedPublisher) GetVerifiedPublisherIdOk() (*string, bool)`

GetVerifiedPublisherIdOk returns a tuple with the VerifiedPublisherId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifiedPublisherId

`func (o *MicrosoftGraphVerifiedPublisher) SetVerifiedPublisherId(v string)`

SetVerifiedPublisherId sets VerifiedPublisherId field to given value.

### HasVerifiedPublisherId

`func (o *MicrosoftGraphVerifiedPublisher) HasVerifiedPublisherId() bool`

HasVerifiedPublisherId returns a boolean if a field has been set.

### SetVerifiedPublisherIdNil

`func (o *MicrosoftGraphVerifiedPublisher) SetVerifiedPublisherIdNil(b bool)`

 SetVerifiedPublisherIdNil sets the value for VerifiedPublisherId to be an explicit nil

### UnsetVerifiedPublisherId
`func (o *MicrosoftGraphVerifiedPublisher) UnsetVerifiedPublisherId()`

UnsetVerifiedPublisherId ensures that no value is present for VerifiedPublisherId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


