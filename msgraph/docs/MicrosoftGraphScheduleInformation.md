# MicrosoftGraphScheduleInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvailabilityView** | Pointer to **NullableString** | Represents a merged view of availability of all the items in scheduleItems. The view consists of time slots. Availability during each time slot is indicated with: 0&#x3D; free, 1&#x3D; tentative, 2&#x3D; busy, 3&#x3D; out of office, 4&#x3D; working elsewhere. | [optional] 
**Error** | Pointer to [**AnyOfmicrosoftGraphFreeBusyError**](anyOf&lt;microsoft.graph.freeBusyError&gt;.md) | Error information from attempting to get the availability of the user, distribution list, or resource. | [optional] 
**ScheduleId** | Pointer to **NullableString** | An SMTP address of the user, distribution list, or resource, identifying an instance of scheduleInformation. | [optional] 
**ScheduleItems** | Pointer to [**[]AnyOfmicrosoftGraphScheduleItem**](AnyOfmicrosoftGraphScheduleItem.md) | Contains the items that describe the availability of the user or resource. | [optional] 
**WorkingHours** | Pointer to [**AnyOfmicrosoftGraphWorkingHours**](anyOf&lt;microsoft.graph.workingHours&gt;.md) | The days of the week and hours in a specific time zone that the user works. These are set as part of the user&#39;s mailboxSettings. | [optional] 

## Methods

### NewMicrosoftGraphScheduleInformation

`func NewMicrosoftGraphScheduleInformation() *MicrosoftGraphScheduleInformation`

NewMicrosoftGraphScheduleInformation instantiates a new MicrosoftGraphScheduleInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphScheduleInformationWithDefaults

`func NewMicrosoftGraphScheduleInformationWithDefaults() *MicrosoftGraphScheduleInformation`

NewMicrosoftGraphScheduleInformationWithDefaults instantiates a new MicrosoftGraphScheduleInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailabilityView

`func (o *MicrosoftGraphScheduleInformation) GetAvailabilityView() string`

GetAvailabilityView returns the AvailabilityView field if non-nil, zero value otherwise.

### GetAvailabilityViewOk

`func (o *MicrosoftGraphScheduleInformation) GetAvailabilityViewOk() (*string, bool)`

GetAvailabilityViewOk returns a tuple with the AvailabilityView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityView

`func (o *MicrosoftGraphScheduleInformation) SetAvailabilityView(v string)`

SetAvailabilityView sets AvailabilityView field to given value.

### HasAvailabilityView

`func (o *MicrosoftGraphScheduleInformation) HasAvailabilityView() bool`

HasAvailabilityView returns a boolean if a field has been set.

### SetAvailabilityViewNil

`func (o *MicrosoftGraphScheduleInformation) SetAvailabilityViewNil(b bool)`

 SetAvailabilityViewNil sets the value for AvailabilityView to be an explicit nil

### UnsetAvailabilityView
`func (o *MicrosoftGraphScheduleInformation) UnsetAvailabilityView()`

UnsetAvailabilityView ensures that no value is present for AvailabilityView, not even an explicit nil
### GetError

`func (o *MicrosoftGraphScheduleInformation) GetError() AnyOfmicrosoftGraphFreeBusyError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *MicrosoftGraphScheduleInformation) GetErrorOk() (*AnyOfmicrosoftGraphFreeBusyError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *MicrosoftGraphScheduleInformation) SetError(v AnyOfmicrosoftGraphFreeBusyError)`

SetError sets Error field to given value.

### HasError

`func (o *MicrosoftGraphScheduleInformation) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *MicrosoftGraphScheduleInformation) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *MicrosoftGraphScheduleInformation) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetScheduleId

`func (o *MicrosoftGraphScheduleInformation) GetScheduleId() string`

GetScheduleId returns the ScheduleId field if non-nil, zero value otherwise.

### GetScheduleIdOk

`func (o *MicrosoftGraphScheduleInformation) GetScheduleIdOk() (*string, bool)`

GetScheduleIdOk returns a tuple with the ScheduleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleId

`func (o *MicrosoftGraphScheduleInformation) SetScheduleId(v string)`

SetScheduleId sets ScheduleId field to given value.

### HasScheduleId

`func (o *MicrosoftGraphScheduleInformation) HasScheduleId() bool`

HasScheduleId returns a boolean if a field has been set.

### SetScheduleIdNil

`func (o *MicrosoftGraphScheduleInformation) SetScheduleIdNil(b bool)`

 SetScheduleIdNil sets the value for ScheduleId to be an explicit nil

### UnsetScheduleId
`func (o *MicrosoftGraphScheduleInformation) UnsetScheduleId()`

UnsetScheduleId ensures that no value is present for ScheduleId, not even an explicit nil
### GetScheduleItems

`func (o *MicrosoftGraphScheduleInformation) GetScheduleItems() []*AnyOfmicrosoftGraphScheduleItem`

GetScheduleItems returns the ScheduleItems field if non-nil, zero value otherwise.

### GetScheduleItemsOk

`func (o *MicrosoftGraphScheduleInformation) GetScheduleItemsOk() (*[]*AnyOfmicrosoftGraphScheduleItem, bool)`

GetScheduleItemsOk returns a tuple with the ScheduleItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleItems

`func (o *MicrosoftGraphScheduleInformation) SetScheduleItems(v []*AnyOfmicrosoftGraphScheduleItem)`

SetScheduleItems sets ScheduleItems field to given value.

### HasScheduleItems

`func (o *MicrosoftGraphScheduleInformation) HasScheduleItems() bool`

HasScheduleItems returns a boolean if a field has been set.

### GetWorkingHours

`func (o *MicrosoftGraphScheduleInformation) GetWorkingHours() AnyOfmicrosoftGraphWorkingHours`

GetWorkingHours returns the WorkingHours field if non-nil, zero value otherwise.

### GetWorkingHoursOk

`func (o *MicrosoftGraphScheduleInformation) GetWorkingHoursOk() (*AnyOfmicrosoftGraphWorkingHours, bool)`

GetWorkingHoursOk returns a tuple with the WorkingHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingHours

`func (o *MicrosoftGraphScheduleInformation) SetWorkingHours(v AnyOfmicrosoftGraphWorkingHours)`

SetWorkingHours sets WorkingHours field to given value.

### HasWorkingHours

`func (o *MicrosoftGraphScheduleInformation) HasWorkingHours() bool`

HasWorkingHours returns a boolean if a field has been set.

### SetWorkingHoursNil

`func (o *MicrosoftGraphScheduleInformation) SetWorkingHoursNil(b bool)`

 SetWorkingHoursNil sets the value for WorkingHours to be an explicit nil

### UnsetWorkingHours
`func (o *MicrosoftGraphScheduleInformation) UnsetWorkingHours()`

UnsetWorkingHours ensures that no value is present for WorkingHours, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


