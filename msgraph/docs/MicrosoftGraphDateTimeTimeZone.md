# MicrosoftGraphDateTimeTimeZone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DateTime** | Pointer to **string** | A single point of time in a combined date and time representation ({date}T{time}; for example, 2017-08-29T04:00:00.0000000). | [optional] 
**TimeZone** | Pointer to **NullableString** | Represents a time zone, for example, &#39;Pacific Standard Time&#39;. See below for more possible values. | [optional] 

## Methods

### NewMicrosoftGraphDateTimeTimeZone

`func NewMicrosoftGraphDateTimeTimeZone() *MicrosoftGraphDateTimeTimeZone`

NewMicrosoftGraphDateTimeTimeZone instantiates a new MicrosoftGraphDateTimeTimeZone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphDateTimeTimeZoneWithDefaults

`func NewMicrosoftGraphDateTimeTimeZoneWithDefaults() *MicrosoftGraphDateTimeTimeZone`

NewMicrosoftGraphDateTimeTimeZoneWithDefaults instantiates a new MicrosoftGraphDateTimeTimeZone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDateTime

`func (o *MicrosoftGraphDateTimeTimeZone) GetDateTime() string`

GetDateTime returns the DateTime field if non-nil, zero value otherwise.

### GetDateTimeOk

`func (o *MicrosoftGraphDateTimeTimeZone) GetDateTimeOk() (*string, bool)`

GetDateTimeOk returns a tuple with the DateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTime

`func (o *MicrosoftGraphDateTimeTimeZone) SetDateTime(v string)`

SetDateTime sets DateTime field to given value.

### HasDateTime

`func (o *MicrosoftGraphDateTimeTimeZone) HasDateTime() bool`

HasDateTime returns a boolean if a field has been set.

### GetTimeZone

`func (o *MicrosoftGraphDateTimeTimeZone) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *MicrosoftGraphDateTimeTimeZone) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *MicrosoftGraphDateTimeTimeZone) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *MicrosoftGraphDateTimeTimeZone) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### SetTimeZoneNil

`func (o *MicrosoftGraphDateTimeTimeZone) SetTimeZoneNil(b bool)`

 SetTimeZoneNil sets the value for TimeZone to be an explicit nil

### UnsetTimeZone
`func (o *MicrosoftGraphDateTimeTimeZone) UnsetTimeZone()`

UnsetTimeZone ensures that no value is present for TimeZone, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


