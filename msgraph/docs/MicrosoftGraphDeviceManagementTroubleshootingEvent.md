# MicrosoftGraphDeviceManagementTroubleshootingEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**CorrelationId** | Pointer to **NullableString** | Id used for tracing the failure in the service. | [optional] 
**EventDateTime** | Pointer to **time.Time** | Time when the event occurred . | [optional] 

## Methods

### NewMicrosoftGraphDeviceManagementTroubleshootingEvent

`func NewMicrosoftGraphDeviceManagementTroubleshootingEvent() *MicrosoftGraphDeviceManagementTroubleshootingEvent`

NewMicrosoftGraphDeviceManagementTroubleshootingEvent instantiates a new MicrosoftGraphDeviceManagementTroubleshootingEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphDeviceManagementTroubleshootingEventWithDefaults

`func NewMicrosoftGraphDeviceManagementTroubleshootingEventWithDefaults() *MicrosoftGraphDeviceManagementTroubleshootingEvent`

NewMicrosoftGraphDeviceManagementTroubleshootingEventWithDefaults instantiates a new MicrosoftGraphDeviceManagementTroubleshootingEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphDeviceManagementTroubleshootingEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphDeviceManagementTroubleshootingEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphDeviceManagementTroubleshootingEvent) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphDeviceManagementTroubleshootingEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCorrelationId

`func (o *MicrosoftGraphDeviceManagementTroubleshootingEvent) GetCorrelationId() string`

GetCorrelationId returns the CorrelationId field if non-nil, zero value otherwise.

### GetCorrelationIdOk

`func (o *MicrosoftGraphDeviceManagementTroubleshootingEvent) GetCorrelationIdOk() (*string, bool)`

GetCorrelationIdOk returns a tuple with the CorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationId

`func (o *MicrosoftGraphDeviceManagementTroubleshootingEvent) SetCorrelationId(v string)`

SetCorrelationId sets CorrelationId field to given value.

### HasCorrelationId

`func (o *MicrosoftGraphDeviceManagementTroubleshootingEvent) HasCorrelationId() bool`

HasCorrelationId returns a boolean if a field has been set.

### SetCorrelationIdNil

`func (o *MicrosoftGraphDeviceManagementTroubleshootingEvent) SetCorrelationIdNil(b bool)`

 SetCorrelationIdNil sets the value for CorrelationId to be an explicit nil

### UnsetCorrelationId
`func (o *MicrosoftGraphDeviceManagementTroubleshootingEvent) UnsetCorrelationId()`

UnsetCorrelationId ensures that no value is present for CorrelationId, not even an explicit nil
### GetEventDateTime

`func (o *MicrosoftGraphDeviceManagementTroubleshootingEvent) GetEventDateTime() time.Time`

GetEventDateTime returns the EventDateTime field if non-nil, zero value otherwise.

### GetEventDateTimeOk

`func (o *MicrosoftGraphDeviceManagementTroubleshootingEvent) GetEventDateTimeOk() (*time.Time, bool)`

GetEventDateTimeOk returns a tuple with the EventDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventDateTime

`func (o *MicrosoftGraphDeviceManagementTroubleshootingEvent) SetEventDateTime(v time.Time)`

SetEventDateTime sets EventDateTime field to given value.

### HasEventDateTime

`func (o *MicrosoftGraphDeviceManagementTroubleshootingEvent) HasEventDateTime() bool`

HasEventDateTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


