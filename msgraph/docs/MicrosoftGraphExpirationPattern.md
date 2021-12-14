# MicrosoftGraphExpirationPattern

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Duration** | Pointer to **NullableString** | The requestor&#39;s desired duration of access represented in ISO 8601 format for durations. For example, PT3H refers to three hours.  If specified in a request, endDateTime should not be present and the type property should be set to afterDuration. | [optional] 
**EndDateTime** | Pointer to **NullableTime** | Timestamp of date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. | [optional] 
**Type** | Pointer to [**AnyOfmicrosoftGraphExpirationPatternType**](anyOf&lt;microsoft.graph.expirationPatternType&gt;.md) | The requestor&#39;s desired expiration pattern type. The possible values are: notSpecified, noExpiration, afterDateTime, afterDuration. | [optional] 

## Methods

### NewMicrosoftGraphExpirationPattern

`func NewMicrosoftGraphExpirationPattern() *MicrosoftGraphExpirationPattern`

NewMicrosoftGraphExpirationPattern instantiates a new MicrosoftGraphExpirationPattern object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphExpirationPatternWithDefaults

`func NewMicrosoftGraphExpirationPatternWithDefaults() *MicrosoftGraphExpirationPattern`

NewMicrosoftGraphExpirationPatternWithDefaults instantiates a new MicrosoftGraphExpirationPattern object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDuration

`func (o *MicrosoftGraphExpirationPattern) GetDuration() string`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *MicrosoftGraphExpirationPattern) GetDurationOk() (*string, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *MicrosoftGraphExpirationPattern) SetDuration(v string)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *MicrosoftGraphExpirationPattern) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### SetDurationNil

`func (o *MicrosoftGraphExpirationPattern) SetDurationNil(b bool)`

 SetDurationNil sets the value for Duration to be an explicit nil

### UnsetDuration
`func (o *MicrosoftGraphExpirationPattern) UnsetDuration()`

UnsetDuration ensures that no value is present for Duration, not even an explicit nil
### GetEndDateTime

`func (o *MicrosoftGraphExpirationPattern) GetEndDateTime() time.Time`

GetEndDateTime returns the EndDateTime field if non-nil, zero value otherwise.

### GetEndDateTimeOk

`func (o *MicrosoftGraphExpirationPattern) GetEndDateTimeOk() (*time.Time, bool)`

GetEndDateTimeOk returns a tuple with the EndDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDateTime

`func (o *MicrosoftGraphExpirationPattern) SetEndDateTime(v time.Time)`

SetEndDateTime sets EndDateTime field to given value.

### HasEndDateTime

`func (o *MicrosoftGraphExpirationPattern) HasEndDateTime() bool`

HasEndDateTime returns a boolean if a field has been set.

### SetEndDateTimeNil

`func (o *MicrosoftGraphExpirationPattern) SetEndDateTimeNil(b bool)`

 SetEndDateTimeNil sets the value for EndDateTime to be an explicit nil

### UnsetEndDateTime
`func (o *MicrosoftGraphExpirationPattern) UnsetEndDateTime()`

UnsetEndDateTime ensures that no value is present for EndDateTime, not even an explicit nil
### GetType

`func (o *MicrosoftGraphExpirationPattern) GetType() AnyOfmicrosoftGraphExpirationPatternType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MicrosoftGraphExpirationPattern) GetTypeOk() (*AnyOfmicrosoftGraphExpirationPatternType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MicrosoftGraphExpirationPattern) SetType(v AnyOfmicrosoftGraphExpirationPatternType)`

SetType sets Type field to given value.

### HasType

`func (o *MicrosoftGraphExpirationPattern) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *MicrosoftGraphExpirationPattern) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *MicrosoftGraphExpirationPattern) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


