# MicrosoftGraphCallRecordsMedia

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CalleeDevice** | Pointer to [**AnyOfmicrosoftGraphCallRecordsDeviceInfo**](anyOf&lt;microsoft.graph.callRecords.deviceInfo&gt;.md) | Device information associated with the callee endpoint of this media. | [optional] 
**CalleeNetwork** | Pointer to [**AnyOfmicrosoftGraphCallRecordsNetworkInfo**](anyOf&lt;microsoft.graph.callRecords.networkInfo&gt;.md) | Network information associated with the callee endpoint of this media. | [optional] 
**CallerDevice** | Pointer to [**AnyOfmicrosoftGraphCallRecordsDeviceInfo**](anyOf&lt;microsoft.graph.callRecords.deviceInfo&gt;.md) | Device information associated with the caller endpoint of this media. | [optional] 
**CallerNetwork** | Pointer to [**AnyOfmicrosoftGraphCallRecordsNetworkInfo**](anyOf&lt;microsoft.graph.callRecords.networkInfo&gt;.md) | Network information associated with the caller endpoint of this media. | [optional] 
**Label** | Pointer to **NullableString** | How the media was identified during media negotiation stage. | [optional] 
**Streams** | Pointer to [**[]AnyOfmicrosoftGraphCallRecordsMediaStream**](AnyOfmicrosoftGraphCallRecordsMediaStream.md) | Network streams associated with this media. | [optional] 

## Methods

### NewMicrosoftGraphCallRecordsMedia

`func NewMicrosoftGraphCallRecordsMedia() *MicrosoftGraphCallRecordsMedia`

NewMicrosoftGraphCallRecordsMedia instantiates a new MicrosoftGraphCallRecordsMedia object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphCallRecordsMediaWithDefaults

`func NewMicrosoftGraphCallRecordsMediaWithDefaults() *MicrosoftGraphCallRecordsMedia`

NewMicrosoftGraphCallRecordsMediaWithDefaults instantiates a new MicrosoftGraphCallRecordsMedia object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCalleeDevice

`func (o *MicrosoftGraphCallRecordsMedia) GetCalleeDevice() AnyOfmicrosoftGraphCallRecordsDeviceInfo`

GetCalleeDevice returns the CalleeDevice field if non-nil, zero value otherwise.

### GetCalleeDeviceOk

`func (o *MicrosoftGraphCallRecordsMedia) GetCalleeDeviceOk() (*AnyOfmicrosoftGraphCallRecordsDeviceInfo, bool)`

GetCalleeDeviceOk returns a tuple with the CalleeDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalleeDevice

`func (o *MicrosoftGraphCallRecordsMedia) SetCalleeDevice(v AnyOfmicrosoftGraphCallRecordsDeviceInfo)`

SetCalleeDevice sets CalleeDevice field to given value.

### HasCalleeDevice

`func (o *MicrosoftGraphCallRecordsMedia) HasCalleeDevice() bool`

HasCalleeDevice returns a boolean if a field has been set.

### SetCalleeDeviceNil

`func (o *MicrosoftGraphCallRecordsMedia) SetCalleeDeviceNil(b bool)`

 SetCalleeDeviceNil sets the value for CalleeDevice to be an explicit nil

### UnsetCalleeDevice
`func (o *MicrosoftGraphCallRecordsMedia) UnsetCalleeDevice()`

UnsetCalleeDevice ensures that no value is present for CalleeDevice, not even an explicit nil
### GetCalleeNetwork

`func (o *MicrosoftGraphCallRecordsMedia) GetCalleeNetwork() AnyOfmicrosoftGraphCallRecordsNetworkInfo`

GetCalleeNetwork returns the CalleeNetwork field if non-nil, zero value otherwise.

### GetCalleeNetworkOk

`func (o *MicrosoftGraphCallRecordsMedia) GetCalleeNetworkOk() (*AnyOfmicrosoftGraphCallRecordsNetworkInfo, bool)`

GetCalleeNetworkOk returns a tuple with the CalleeNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalleeNetwork

`func (o *MicrosoftGraphCallRecordsMedia) SetCalleeNetwork(v AnyOfmicrosoftGraphCallRecordsNetworkInfo)`

SetCalleeNetwork sets CalleeNetwork field to given value.

### HasCalleeNetwork

`func (o *MicrosoftGraphCallRecordsMedia) HasCalleeNetwork() bool`

HasCalleeNetwork returns a boolean if a field has been set.

### SetCalleeNetworkNil

`func (o *MicrosoftGraphCallRecordsMedia) SetCalleeNetworkNil(b bool)`

 SetCalleeNetworkNil sets the value for CalleeNetwork to be an explicit nil

### UnsetCalleeNetwork
`func (o *MicrosoftGraphCallRecordsMedia) UnsetCalleeNetwork()`

UnsetCalleeNetwork ensures that no value is present for CalleeNetwork, not even an explicit nil
### GetCallerDevice

`func (o *MicrosoftGraphCallRecordsMedia) GetCallerDevice() AnyOfmicrosoftGraphCallRecordsDeviceInfo`

GetCallerDevice returns the CallerDevice field if non-nil, zero value otherwise.

### GetCallerDeviceOk

`func (o *MicrosoftGraphCallRecordsMedia) GetCallerDeviceOk() (*AnyOfmicrosoftGraphCallRecordsDeviceInfo, bool)`

GetCallerDeviceOk returns a tuple with the CallerDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallerDevice

`func (o *MicrosoftGraphCallRecordsMedia) SetCallerDevice(v AnyOfmicrosoftGraphCallRecordsDeviceInfo)`

SetCallerDevice sets CallerDevice field to given value.

### HasCallerDevice

`func (o *MicrosoftGraphCallRecordsMedia) HasCallerDevice() bool`

HasCallerDevice returns a boolean if a field has been set.

### SetCallerDeviceNil

`func (o *MicrosoftGraphCallRecordsMedia) SetCallerDeviceNil(b bool)`

 SetCallerDeviceNil sets the value for CallerDevice to be an explicit nil

### UnsetCallerDevice
`func (o *MicrosoftGraphCallRecordsMedia) UnsetCallerDevice()`

UnsetCallerDevice ensures that no value is present for CallerDevice, not even an explicit nil
### GetCallerNetwork

`func (o *MicrosoftGraphCallRecordsMedia) GetCallerNetwork() AnyOfmicrosoftGraphCallRecordsNetworkInfo`

GetCallerNetwork returns the CallerNetwork field if non-nil, zero value otherwise.

### GetCallerNetworkOk

`func (o *MicrosoftGraphCallRecordsMedia) GetCallerNetworkOk() (*AnyOfmicrosoftGraphCallRecordsNetworkInfo, bool)`

GetCallerNetworkOk returns a tuple with the CallerNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallerNetwork

`func (o *MicrosoftGraphCallRecordsMedia) SetCallerNetwork(v AnyOfmicrosoftGraphCallRecordsNetworkInfo)`

SetCallerNetwork sets CallerNetwork field to given value.

### HasCallerNetwork

`func (o *MicrosoftGraphCallRecordsMedia) HasCallerNetwork() bool`

HasCallerNetwork returns a boolean if a field has been set.

### SetCallerNetworkNil

`func (o *MicrosoftGraphCallRecordsMedia) SetCallerNetworkNil(b bool)`

 SetCallerNetworkNil sets the value for CallerNetwork to be an explicit nil

### UnsetCallerNetwork
`func (o *MicrosoftGraphCallRecordsMedia) UnsetCallerNetwork()`

UnsetCallerNetwork ensures that no value is present for CallerNetwork, not even an explicit nil
### GetLabel

`func (o *MicrosoftGraphCallRecordsMedia) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *MicrosoftGraphCallRecordsMedia) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *MicrosoftGraphCallRecordsMedia) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *MicrosoftGraphCallRecordsMedia) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### SetLabelNil

`func (o *MicrosoftGraphCallRecordsMedia) SetLabelNil(b bool)`

 SetLabelNil sets the value for Label to be an explicit nil

### UnsetLabel
`func (o *MicrosoftGraphCallRecordsMedia) UnsetLabel()`

UnsetLabel ensures that no value is present for Label, not even an explicit nil
### GetStreams

`func (o *MicrosoftGraphCallRecordsMedia) GetStreams() []*AnyOfmicrosoftGraphCallRecordsMediaStream`

GetStreams returns the Streams field if non-nil, zero value otherwise.

### GetStreamsOk

`func (o *MicrosoftGraphCallRecordsMedia) GetStreamsOk() (*[]*AnyOfmicrosoftGraphCallRecordsMediaStream, bool)`

GetStreamsOk returns a tuple with the Streams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreams

`func (o *MicrosoftGraphCallRecordsMedia) SetStreams(v []*AnyOfmicrosoftGraphCallRecordsMediaStream)`

SetStreams sets Streams field to given value.

### HasStreams

`func (o *MicrosoftGraphCallRecordsMedia) HasStreams() bool`

HasStreams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


