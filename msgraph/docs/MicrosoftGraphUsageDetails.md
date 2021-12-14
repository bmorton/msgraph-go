# MicrosoftGraphUsageDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastAccessedDateTime** | Pointer to **NullableTime** | The date and time the resource was last accessed by the user. The timestamp represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only. | [optional] 
**LastModifiedDateTime** | Pointer to **NullableTime** | The date and time the resource was last modified by the user. The timestamp represents date and time information using ISO 8601 format and is always in UTC time.For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only. | [optional] 

## Methods

### NewMicrosoftGraphUsageDetails

`func NewMicrosoftGraphUsageDetails() *MicrosoftGraphUsageDetails`

NewMicrosoftGraphUsageDetails instantiates a new MicrosoftGraphUsageDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphUsageDetailsWithDefaults

`func NewMicrosoftGraphUsageDetailsWithDefaults() *MicrosoftGraphUsageDetails`

NewMicrosoftGraphUsageDetailsWithDefaults instantiates a new MicrosoftGraphUsageDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastAccessedDateTime

`func (o *MicrosoftGraphUsageDetails) GetLastAccessedDateTime() time.Time`

GetLastAccessedDateTime returns the LastAccessedDateTime field if non-nil, zero value otherwise.

### GetLastAccessedDateTimeOk

`func (o *MicrosoftGraphUsageDetails) GetLastAccessedDateTimeOk() (*time.Time, bool)`

GetLastAccessedDateTimeOk returns a tuple with the LastAccessedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAccessedDateTime

`func (o *MicrosoftGraphUsageDetails) SetLastAccessedDateTime(v time.Time)`

SetLastAccessedDateTime sets LastAccessedDateTime field to given value.

### HasLastAccessedDateTime

`func (o *MicrosoftGraphUsageDetails) HasLastAccessedDateTime() bool`

HasLastAccessedDateTime returns a boolean if a field has been set.

### SetLastAccessedDateTimeNil

`func (o *MicrosoftGraphUsageDetails) SetLastAccessedDateTimeNil(b bool)`

 SetLastAccessedDateTimeNil sets the value for LastAccessedDateTime to be an explicit nil

### UnsetLastAccessedDateTime
`func (o *MicrosoftGraphUsageDetails) UnsetLastAccessedDateTime()`

UnsetLastAccessedDateTime ensures that no value is present for LastAccessedDateTime, not even an explicit nil
### GetLastModifiedDateTime

`func (o *MicrosoftGraphUsageDetails) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *MicrosoftGraphUsageDetails) GetLastModifiedDateTimeOk() (*time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedDateTime

`func (o *MicrosoftGraphUsageDetails) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime sets LastModifiedDateTime field to given value.

### HasLastModifiedDateTime

`func (o *MicrosoftGraphUsageDetails) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### SetLastModifiedDateTimeNil

`func (o *MicrosoftGraphUsageDetails) SetLastModifiedDateTimeNil(b bool)`

 SetLastModifiedDateTimeNil sets the value for LastModifiedDateTime to be an explicit nil

### UnsetLastModifiedDateTime
`func (o *MicrosoftGraphUsageDetails) UnsetLastModifiedDateTime()`

UnsetLastModifiedDateTime ensures that no value is present for LastModifiedDateTime, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


