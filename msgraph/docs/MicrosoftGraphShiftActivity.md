# MicrosoftGraphShiftActivity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **NullableString** | Customer defined code for the shiftActivity. Required. | [optional] 
**DisplayName** | Pointer to **NullableString** | The name of the shiftActivity. Required. | [optional] 
**EndDateTime** | Pointer to **NullableTime** | The end date and time for the shiftActivity. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Required. | [optional] 
**IsPaid** | Pointer to **NullableBool** | Indicates whether the microsoft.graph.user should be paid for the activity during their shift. Required. | [optional] 
**StartDateTime** | Pointer to **NullableTime** | The start date and time for the shiftActivity. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Required. | [optional] 
**Theme** | Pointer to [**AnyOfmicrosoftGraphScheduleEntityTheme**](anyOf&lt;microsoft.graph.scheduleEntityTheme&gt;.md) |  | [optional] 

## Methods

### NewMicrosoftGraphShiftActivity

`func NewMicrosoftGraphShiftActivity() *MicrosoftGraphShiftActivity`

NewMicrosoftGraphShiftActivity instantiates a new MicrosoftGraphShiftActivity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphShiftActivityWithDefaults

`func NewMicrosoftGraphShiftActivityWithDefaults() *MicrosoftGraphShiftActivity`

NewMicrosoftGraphShiftActivityWithDefaults instantiates a new MicrosoftGraphShiftActivity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *MicrosoftGraphShiftActivity) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *MicrosoftGraphShiftActivity) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *MicrosoftGraphShiftActivity) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *MicrosoftGraphShiftActivity) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *MicrosoftGraphShiftActivity) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *MicrosoftGraphShiftActivity) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetDisplayName

`func (o *MicrosoftGraphShiftActivity) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphShiftActivity) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MicrosoftGraphShiftActivity) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MicrosoftGraphShiftActivity) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *MicrosoftGraphShiftActivity) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *MicrosoftGraphShiftActivity) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetEndDateTime

`func (o *MicrosoftGraphShiftActivity) GetEndDateTime() time.Time`

GetEndDateTime returns the EndDateTime field if non-nil, zero value otherwise.

### GetEndDateTimeOk

`func (o *MicrosoftGraphShiftActivity) GetEndDateTimeOk() (*time.Time, bool)`

GetEndDateTimeOk returns a tuple with the EndDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDateTime

`func (o *MicrosoftGraphShiftActivity) SetEndDateTime(v time.Time)`

SetEndDateTime sets EndDateTime field to given value.

### HasEndDateTime

`func (o *MicrosoftGraphShiftActivity) HasEndDateTime() bool`

HasEndDateTime returns a boolean if a field has been set.

### SetEndDateTimeNil

`func (o *MicrosoftGraphShiftActivity) SetEndDateTimeNil(b bool)`

 SetEndDateTimeNil sets the value for EndDateTime to be an explicit nil

### UnsetEndDateTime
`func (o *MicrosoftGraphShiftActivity) UnsetEndDateTime()`

UnsetEndDateTime ensures that no value is present for EndDateTime, not even an explicit nil
### GetIsPaid

`func (o *MicrosoftGraphShiftActivity) GetIsPaid() bool`

GetIsPaid returns the IsPaid field if non-nil, zero value otherwise.

### GetIsPaidOk

`func (o *MicrosoftGraphShiftActivity) GetIsPaidOk() (*bool, bool)`

GetIsPaidOk returns a tuple with the IsPaid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPaid

`func (o *MicrosoftGraphShiftActivity) SetIsPaid(v bool)`

SetIsPaid sets IsPaid field to given value.

### HasIsPaid

`func (o *MicrosoftGraphShiftActivity) HasIsPaid() bool`

HasIsPaid returns a boolean if a field has been set.

### SetIsPaidNil

`func (o *MicrosoftGraphShiftActivity) SetIsPaidNil(b bool)`

 SetIsPaidNil sets the value for IsPaid to be an explicit nil

### UnsetIsPaid
`func (o *MicrosoftGraphShiftActivity) UnsetIsPaid()`

UnsetIsPaid ensures that no value is present for IsPaid, not even an explicit nil
### GetStartDateTime

`func (o *MicrosoftGraphShiftActivity) GetStartDateTime() time.Time`

GetStartDateTime returns the StartDateTime field if non-nil, zero value otherwise.

### GetStartDateTimeOk

`func (o *MicrosoftGraphShiftActivity) GetStartDateTimeOk() (*time.Time, bool)`

GetStartDateTimeOk returns a tuple with the StartDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDateTime

`func (o *MicrosoftGraphShiftActivity) SetStartDateTime(v time.Time)`

SetStartDateTime sets StartDateTime field to given value.

### HasStartDateTime

`func (o *MicrosoftGraphShiftActivity) HasStartDateTime() bool`

HasStartDateTime returns a boolean if a field has been set.

### SetStartDateTimeNil

`func (o *MicrosoftGraphShiftActivity) SetStartDateTimeNil(b bool)`

 SetStartDateTimeNil sets the value for StartDateTime to be an explicit nil

### UnsetStartDateTime
`func (o *MicrosoftGraphShiftActivity) UnsetStartDateTime()`

UnsetStartDateTime ensures that no value is present for StartDateTime, not even an explicit nil
### GetTheme

`func (o *MicrosoftGraphShiftActivity) GetTheme() AnyOfmicrosoftGraphScheduleEntityTheme`

GetTheme returns the Theme field if non-nil, zero value otherwise.

### GetThemeOk

`func (o *MicrosoftGraphShiftActivity) GetThemeOk() (*AnyOfmicrosoftGraphScheduleEntityTheme, bool)`

GetThemeOk returns a tuple with the Theme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTheme

`func (o *MicrosoftGraphShiftActivity) SetTheme(v AnyOfmicrosoftGraphScheduleEntityTheme)`

SetTheme sets Theme field to given value.

### HasTheme

`func (o *MicrosoftGraphShiftActivity) HasTheme() bool`

HasTheme returns a boolean if a field has been set.

### SetThemeNil

`func (o *MicrosoftGraphShiftActivity) SetThemeNil(b bool)`

 SetThemeNil sets the value for Theme to be an explicit nil

### UnsetTheme
`func (o *MicrosoftGraphShiftActivity) UnsetTheme()`

UnsetTheme ensures that no value is present for Theme, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


