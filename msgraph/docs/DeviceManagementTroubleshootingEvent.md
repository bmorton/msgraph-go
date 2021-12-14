# DeviceManagementTroubleshootingEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CorrelationId** | Pointer to **NullableString** | Id used for tracing the failure in the service. | [optional] 
**EventDateTime** | Pointer to **time.Time** | Time when the event occurred . | [optional] 

## Methods

### NewDeviceManagementTroubleshootingEvent

`func NewDeviceManagementTroubleshootingEvent() *DeviceManagementTroubleshootingEvent`

NewDeviceManagementTroubleshootingEvent instantiates a new DeviceManagementTroubleshootingEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceManagementTroubleshootingEventWithDefaults

`func NewDeviceManagementTroubleshootingEventWithDefaults() *DeviceManagementTroubleshootingEvent`

NewDeviceManagementTroubleshootingEventWithDefaults instantiates a new DeviceManagementTroubleshootingEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCorrelationId

`func (o *DeviceManagementTroubleshootingEvent) GetCorrelationId() string`

GetCorrelationId returns the CorrelationId field if non-nil, zero value otherwise.

### GetCorrelationIdOk

`func (o *DeviceManagementTroubleshootingEvent) GetCorrelationIdOk() (*string, bool)`

GetCorrelationIdOk returns a tuple with the CorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationId

`func (o *DeviceManagementTroubleshootingEvent) SetCorrelationId(v string)`

SetCorrelationId sets CorrelationId field to given value.

### HasCorrelationId

`func (o *DeviceManagementTroubleshootingEvent) HasCorrelationId() bool`

HasCorrelationId returns a boolean if a field has been set.

### SetCorrelationIdNil

`func (o *DeviceManagementTroubleshootingEvent) SetCorrelationIdNil(b bool)`

 SetCorrelationIdNil sets the value for CorrelationId to be an explicit nil

### UnsetCorrelationId
`func (o *DeviceManagementTroubleshootingEvent) UnsetCorrelationId()`

UnsetCorrelationId ensures that no value is present for CorrelationId, not even an explicit nil
### GetEventDateTime

`func (o *DeviceManagementTroubleshootingEvent) GetEventDateTime() time.Time`

GetEventDateTime returns the EventDateTime field if non-nil, zero value otherwise.

### GetEventDateTimeOk

`func (o *DeviceManagementTroubleshootingEvent) GetEventDateTimeOk() (*time.Time, bool)`

GetEventDateTimeOk returns a tuple with the EventDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventDateTime

`func (o *DeviceManagementTroubleshootingEvent) SetEventDateTime(v time.Time)`

SetEventDateTime sets EventDateTime field to given value.

### HasEventDateTime

`func (o *DeviceManagementTroubleshootingEvent) HasEventDateTime() bool`

HasEventDateTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


