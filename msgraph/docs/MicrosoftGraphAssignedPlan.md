# MicrosoftGraphAssignedPlan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssignedDateTime** | Pointer to **NullableTime** | The date and time at which the plan was assigned. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. | [optional] 
**CapabilityStatus** | Pointer to **NullableString** | Condition of the capability assignment. The possible values are Enabled, Warning, Suspended, Deleted, LockedOut. See a detailed description of each value. | [optional] 
**Service** | Pointer to **NullableString** | The name of the service; for example, &#39;Exchange&#39;. | [optional] 
**ServicePlanId** | Pointer to **NullableString** | A GUID that identifies the service plan. | [optional] 

## Methods

### NewMicrosoftGraphAssignedPlan

`func NewMicrosoftGraphAssignedPlan() *MicrosoftGraphAssignedPlan`

NewMicrosoftGraphAssignedPlan instantiates a new MicrosoftGraphAssignedPlan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphAssignedPlanWithDefaults

`func NewMicrosoftGraphAssignedPlanWithDefaults() *MicrosoftGraphAssignedPlan`

NewMicrosoftGraphAssignedPlanWithDefaults instantiates a new MicrosoftGraphAssignedPlan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignedDateTime

`func (o *MicrosoftGraphAssignedPlan) GetAssignedDateTime() time.Time`

GetAssignedDateTime returns the AssignedDateTime field if non-nil, zero value otherwise.

### GetAssignedDateTimeOk

`func (o *MicrosoftGraphAssignedPlan) GetAssignedDateTimeOk() (*time.Time, bool)`

GetAssignedDateTimeOk returns a tuple with the AssignedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedDateTime

`func (o *MicrosoftGraphAssignedPlan) SetAssignedDateTime(v time.Time)`

SetAssignedDateTime sets AssignedDateTime field to given value.

### HasAssignedDateTime

`func (o *MicrosoftGraphAssignedPlan) HasAssignedDateTime() bool`

HasAssignedDateTime returns a boolean if a field has been set.

### SetAssignedDateTimeNil

`func (o *MicrosoftGraphAssignedPlan) SetAssignedDateTimeNil(b bool)`

 SetAssignedDateTimeNil sets the value for AssignedDateTime to be an explicit nil

### UnsetAssignedDateTime
`func (o *MicrosoftGraphAssignedPlan) UnsetAssignedDateTime()`

UnsetAssignedDateTime ensures that no value is present for AssignedDateTime, not even an explicit nil
### GetCapabilityStatus

`func (o *MicrosoftGraphAssignedPlan) GetCapabilityStatus() string`

GetCapabilityStatus returns the CapabilityStatus field if non-nil, zero value otherwise.

### GetCapabilityStatusOk

`func (o *MicrosoftGraphAssignedPlan) GetCapabilityStatusOk() (*string, bool)`

GetCapabilityStatusOk returns a tuple with the CapabilityStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilityStatus

`func (o *MicrosoftGraphAssignedPlan) SetCapabilityStatus(v string)`

SetCapabilityStatus sets CapabilityStatus field to given value.

### HasCapabilityStatus

`func (o *MicrosoftGraphAssignedPlan) HasCapabilityStatus() bool`

HasCapabilityStatus returns a boolean if a field has been set.

### SetCapabilityStatusNil

`func (o *MicrosoftGraphAssignedPlan) SetCapabilityStatusNil(b bool)`

 SetCapabilityStatusNil sets the value for CapabilityStatus to be an explicit nil

### UnsetCapabilityStatus
`func (o *MicrosoftGraphAssignedPlan) UnsetCapabilityStatus()`

UnsetCapabilityStatus ensures that no value is present for CapabilityStatus, not even an explicit nil
### GetService

`func (o *MicrosoftGraphAssignedPlan) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *MicrosoftGraphAssignedPlan) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *MicrosoftGraphAssignedPlan) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *MicrosoftGraphAssignedPlan) HasService() bool`

HasService returns a boolean if a field has been set.

### SetServiceNil

`func (o *MicrosoftGraphAssignedPlan) SetServiceNil(b bool)`

 SetServiceNil sets the value for Service to be an explicit nil

### UnsetService
`func (o *MicrosoftGraphAssignedPlan) UnsetService()`

UnsetService ensures that no value is present for Service, not even an explicit nil
### GetServicePlanId

`func (o *MicrosoftGraphAssignedPlan) GetServicePlanId() string`

GetServicePlanId returns the ServicePlanId field if non-nil, zero value otherwise.

### GetServicePlanIdOk

`func (o *MicrosoftGraphAssignedPlan) GetServicePlanIdOk() (*string, bool)`

GetServicePlanIdOk returns a tuple with the ServicePlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePlanId

`func (o *MicrosoftGraphAssignedPlan) SetServicePlanId(v string)`

SetServicePlanId sets ServicePlanId field to given value.

### HasServicePlanId

`func (o *MicrosoftGraphAssignedPlan) HasServicePlanId() bool`

HasServicePlanId returns a boolean if a field has been set.

### SetServicePlanIdNil

`func (o *MicrosoftGraphAssignedPlan) SetServicePlanIdNil(b bool)`

 SetServicePlanIdNil sets the value for ServicePlanId to be an explicit nil

### UnsetServicePlanId
`func (o *MicrosoftGraphAssignedPlan) UnsetServicePlanId()`

UnsetServicePlanId ensures that no value is present for ServicePlanId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


