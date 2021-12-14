# MicrosoftGraphServiceAnnouncementBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**Details** | Pointer to [**[]AnyOfmicrosoftGraphKeyValuePair**](AnyOfmicrosoftGraphKeyValuePair.md) | Additional details about service event. This property doesn&#39;t support filters. | [optional] 
**EndDateTime** | Pointer to **NullableTime** | The end time of the service event. | [optional] 
**LastModifiedDateTime** | Pointer to **time.Time** | The last modified time of the service event. | [optional] 
**StartDateTime** | Pointer to **time.Time** | The start time of the service event. | [optional] 
**Title** | Pointer to **string** | The title of the service event. | [optional] 

## Methods

### NewMicrosoftGraphServiceAnnouncementBase

`func NewMicrosoftGraphServiceAnnouncementBase() *MicrosoftGraphServiceAnnouncementBase`

NewMicrosoftGraphServiceAnnouncementBase instantiates a new MicrosoftGraphServiceAnnouncementBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphServiceAnnouncementBaseWithDefaults

`func NewMicrosoftGraphServiceAnnouncementBaseWithDefaults() *MicrosoftGraphServiceAnnouncementBase`

NewMicrosoftGraphServiceAnnouncementBaseWithDefaults instantiates a new MicrosoftGraphServiceAnnouncementBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphServiceAnnouncementBase) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphServiceAnnouncementBase) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphServiceAnnouncementBase) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphServiceAnnouncementBase) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDetails

`func (o *MicrosoftGraphServiceAnnouncementBase) GetDetails() []*AnyOfmicrosoftGraphKeyValuePair`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *MicrosoftGraphServiceAnnouncementBase) GetDetailsOk() (*[]*AnyOfmicrosoftGraphKeyValuePair, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *MicrosoftGraphServiceAnnouncementBase) SetDetails(v []*AnyOfmicrosoftGraphKeyValuePair)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *MicrosoftGraphServiceAnnouncementBase) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetEndDateTime

`func (o *MicrosoftGraphServiceAnnouncementBase) GetEndDateTime() time.Time`

GetEndDateTime returns the EndDateTime field if non-nil, zero value otherwise.

### GetEndDateTimeOk

`func (o *MicrosoftGraphServiceAnnouncementBase) GetEndDateTimeOk() (*time.Time, bool)`

GetEndDateTimeOk returns a tuple with the EndDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDateTime

`func (o *MicrosoftGraphServiceAnnouncementBase) SetEndDateTime(v time.Time)`

SetEndDateTime sets EndDateTime field to given value.

### HasEndDateTime

`func (o *MicrosoftGraphServiceAnnouncementBase) HasEndDateTime() bool`

HasEndDateTime returns a boolean if a field has been set.

### SetEndDateTimeNil

`func (o *MicrosoftGraphServiceAnnouncementBase) SetEndDateTimeNil(b bool)`

 SetEndDateTimeNil sets the value for EndDateTime to be an explicit nil

### UnsetEndDateTime
`func (o *MicrosoftGraphServiceAnnouncementBase) UnsetEndDateTime()`

UnsetEndDateTime ensures that no value is present for EndDateTime, not even an explicit nil
### GetLastModifiedDateTime

`func (o *MicrosoftGraphServiceAnnouncementBase) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *MicrosoftGraphServiceAnnouncementBase) GetLastModifiedDateTimeOk() (*time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedDateTime

`func (o *MicrosoftGraphServiceAnnouncementBase) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime sets LastModifiedDateTime field to given value.

### HasLastModifiedDateTime

`func (o *MicrosoftGraphServiceAnnouncementBase) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### GetStartDateTime

`func (o *MicrosoftGraphServiceAnnouncementBase) GetStartDateTime() time.Time`

GetStartDateTime returns the StartDateTime field if non-nil, zero value otherwise.

### GetStartDateTimeOk

`func (o *MicrosoftGraphServiceAnnouncementBase) GetStartDateTimeOk() (*time.Time, bool)`

GetStartDateTimeOk returns a tuple with the StartDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDateTime

`func (o *MicrosoftGraphServiceAnnouncementBase) SetStartDateTime(v time.Time)`

SetStartDateTime sets StartDateTime field to given value.

### HasStartDateTime

`func (o *MicrosoftGraphServiceAnnouncementBase) HasStartDateTime() bool`

HasStartDateTime returns a boolean if a field has been set.

### GetTitle

`func (o *MicrosoftGraphServiceAnnouncementBase) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *MicrosoftGraphServiceAnnouncementBase) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *MicrosoftGraphServiceAnnouncementBase) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *MicrosoftGraphServiceAnnouncementBase) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


