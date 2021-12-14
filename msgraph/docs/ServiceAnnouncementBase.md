# ServiceAnnouncementBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Details** | Pointer to [**[]AnyOfmicrosoftGraphKeyValuePair**](AnyOfmicrosoftGraphKeyValuePair.md) | Additional details about service event. This property doesn&#39;t support filters. | [optional] 
**EndDateTime** | Pointer to **NullableTime** | The end time of the service event. | [optional] 
**LastModifiedDateTime** | Pointer to **time.Time** | The last modified time of the service event. | [optional] 
**StartDateTime** | Pointer to **time.Time** | The start time of the service event. | [optional] 
**Title** | Pointer to **string** | The title of the service event. | [optional] 

## Methods

### NewServiceAnnouncementBase

`func NewServiceAnnouncementBase() *ServiceAnnouncementBase`

NewServiceAnnouncementBase instantiates a new ServiceAnnouncementBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceAnnouncementBaseWithDefaults

`func NewServiceAnnouncementBaseWithDefaults() *ServiceAnnouncementBase`

NewServiceAnnouncementBaseWithDefaults instantiates a new ServiceAnnouncementBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetails

`func (o *ServiceAnnouncementBase) GetDetails() []*AnyOfmicrosoftGraphKeyValuePair`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *ServiceAnnouncementBase) GetDetailsOk() (*[]*AnyOfmicrosoftGraphKeyValuePair, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *ServiceAnnouncementBase) SetDetails(v []*AnyOfmicrosoftGraphKeyValuePair)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *ServiceAnnouncementBase) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetEndDateTime

`func (o *ServiceAnnouncementBase) GetEndDateTime() time.Time`

GetEndDateTime returns the EndDateTime field if non-nil, zero value otherwise.

### GetEndDateTimeOk

`func (o *ServiceAnnouncementBase) GetEndDateTimeOk() (*time.Time, bool)`

GetEndDateTimeOk returns a tuple with the EndDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDateTime

`func (o *ServiceAnnouncementBase) SetEndDateTime(v time.Time)`

SetEndDateTime sets EndDateTime field to given value.

### HasEndDateTime

`func (o *ServiceAnnouncementBase) HasEndDateTime() bool`

HasEndDateTime returns a boolean if a field has been set.

### SetEndDateTimeNil

`func (o *ServiceAnnouncementBase) SetEndDateTimeNil(b bool)`

 SetEndDateTimeNil sets the value for EndDateTime to be an explicit nil

### UnsetEndDateTime
`func (o *ServiceAnnouncementBase) UnsetEndDateTime()`

UnsetEndDateTime ensures that no value is present for EndDateTime, not even an explicit nil
### GetLastModifiedDateTime

`func (o *ServiceAnnouncementBase) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *ServiceAnnouncementBase) GetLastModifiedDateTimeOk() (*time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedDateTime

`func (o *ServiceAnnouncementBase) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime sets LastModifiedDateTime field to given value.

### HasLastModifiedDateTime

`func (o *ServiceAnnouncementBase) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### GetStartDateTime

`func (o *ServiceAnnouncementBase) GetStartDateTime() time.Time`

GetStartDateTime returns the StartDateTime field if non-nil, zero value otherwise.

### GetStartDateTimeOk

`func (o *ServiceAnnouncementBase) GetStartDateTimeOk() (*time.Time, bool)`

GetStartDateTimeOk returns a tuple with the StartDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDateTime

`func (o *ServiceAnnouncementBase) SetStartDateTime(v time.Time)`

SetStartDateTime sets StartDateTime field to given value.

### HasStartDateTime

`func (o *ServiceAnnouncementBase) HasStartDateTime() bool`

HasStartDateTime returns a boolean if a field has been set.

### GetTitle

`func (o *ServiceAnnouncementBase) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ServiceAnnouncementBase) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ServiceAnnouncementBase) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ServiceAnnouncementBase) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


