# MicrosoftGraphConditionalAccessDevices

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceFilter** | Pointer to [**AnyOfmicrosoftGraphConditionalAccessFilter**](anyOf&lt;microsoft.graph.conditionalAccessFilter&gt;.md) | Filter that defines the dynamic-device-syntax rule to include/exclude devices. A filter can use device properties (such as extension attributes) to include/exclude them. | [optional] 

## Methods

### NewMicrosoftGraphConditionalAccessDevices

`func NewMicrosoftGraphConditionalAccessDevices() *MicrosoftGraphConditionalAccessDevices`

NewMicrosoftGraphConditionalAccessDevices instantiates a new MicrosoftGraphConditionalAccessDevices object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphConditionalAccessDevicesWithDefaults

`func NewMicrosoftGraphConditionalAccessDevicesWithDefaults() *MicrosoftGraphConditionalAccessDevices`

NewMicrosoftGraphConditionalAccessDevicesWithDefaults instantiates a new MicrosoftGraphConditionalAccessDevices object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceFilter

`func (o *MicrosoftGraphConditionalAccessDevices) GetDeviceFilter() AnyOfmicrosoftGraphConditionalAccessFilter`

GetDeviceFilter returns the DeviceFilter field if non-nil, zero value otherwise.

### GetDeviceFilterOk

`func (o *MicrosoftGraphConditionalAccessDevices) GetDeviceFilterOk() (*AnyOfmicrosoftGraphConditionalAccessFilter, bool)`

GetDeviceFilterOk returns a tuple with the DeviceFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceFilter

`func (o *MicrosoftGraphConditionalAccessDevices) SetDeviceFilter(v AnyOfmicrosoftGraphConditionalAccessFilter)`

SetDeviceFilter sets DeviceFilter field to given value.

### HasDeviceFilter

`func (o *MicrosoftGraphConditionalAccessDevices) HasDeviceFilter() bool`

HasDeviceFilter returns a boolean if a field has been set.

### SetDeviceFilterNil

`func (o *MicrosoftGraphConditionalAccessDevices) SetDeviceFilterNil(b bool)`

 SetDeviceFilterNil sets the value for DeviceFilter to be an explicit nil

### UnsetDeviceFilter
`func (o *MicrosoftGraphConditionalAccessDevices) UnsetDeviceFilter()`

UnsetDeviceFilter ensures that no value is present for DeviceFilter, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


