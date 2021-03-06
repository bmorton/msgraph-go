/*
Partial Graph API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package msgraph

import (
	"encoding/json"
)

// MicrosoftGraphCallRecordsMedia struct for MicrosoftGraphCallRecordsMedia
type MicrosoftGraphCallRecordsMedia struct {
	// Device information associated with the callee endpoint of this media.
	CalleeDevice AnyOfmicrosoftGraphCallRecordsDeviceInfo `json:"calleeDevice,omitempty"`
	// Network information associated with the callee endpoint of this media.
	CalleeNetwork AnyOfmicrosoftGraphCallRecordsNetworkInfo `json:"calleeNetwork,omitempty"`
	// Device information associated with the caller endpoint of this media.
	CallerDevice AnyOfmicrosoftGraphCallRecordsDeviceInfo `json:"callerDevice,omitempty"`
	// Network information associated with the caller endpoint of this media.
	CallerNetwork AnyOfmicrosoftGraphCallRecordsNetworkInfo `json:"callerNetwork,omitempty"`
	// How the media was identified during media negotiation stage.
	Label NullableString `json:"label,omitempty"`
	// Network streams associated with this media.
	Streams *[]*AnyOfmicrosoftGraphCallRecordsMediaStream `json:"streams,omitempty"`
}

// NewMicrosoftGraphCallRecordsMedia instantiates a new MicrosoftGraphCallRecordsMedia object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphCallRecordsMedia() *MicrosoftGraphCallRecordsMedia {
	this := MicrosoftGraphCallRecordsMedia{}
	return &this
}

// NewMicrosoftGraphCallRecordsMediaWithDefaults instantiates a new MicrosoftGraphCallRecordsMedia object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphCallRecordsMediaWithDefaults() *MicrosoftGraphCallRecordsMedia {
	this := MicrosoftGraphCallRecordsMedia{}
	return &this
}

// GetCalleeDevice returns the CalleeDevice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphCallRecordsMedia) GetCalleeDevice() AnyOfmicrosoftGraphCallRecordsDeviceInfo {
	if o == nil  {
		var ret AnyOfmicrosoftGraphCallRecordsDeviceInfo
		return ret
	}
	return o.CalleeDevice
}

// GetCalleeDeviceOk returns a tuple with the CalleeDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphCallRecordsMedia) GetCalleeDeviceOk() (*AnyOfmicrosoftGraphCallRecordsDeviceInfo, bool) {
	if o == nil || o.CalleeDevice == nil {
		return nil, false
	}
	return &o.CalleeDevice, true
}

// HasCalleeDevice returns a boolean if a field has been set.
func (o *MicrosoftGraphCallRecordsMedia) HasCalleeDevice() bool {
	if o != nil && o.CalleeDevice != nil {
		return true
	}

	return false
}

// SetCalleeDevice gets a reference to the given AnyOfmicrosoftGraphCallRecordsDeviceInfo and assigns it to the CalleeDevice field.
func (o *MicrosoftGraphCallRecordsMedia) SetCalleeDevice(v AnyOfmicrosoftGraphCallRecordsDeviceInfo) {
	o.CalleeDevice = v
}

// GetCalleeNetwork returns the CalleeNetwork field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphCallRecordsMedia) GetCalleeNetwork() AnyOfmicrosoftGraphCallRecordsNetworkInfo {
	if o == nil  {
		var ret AnyOfmicrosoftGraphCallRecordsNetworkInfo
		return ret
	}
	return o.CalleeNetwork
}

// GetCalleeNetworkOk returns a tuple with the CalleeNetwork field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphCallRecordsMedia) GetCalleeNetworkOk() (*AnyOfmicrosoftGraphCallRecordsNetworkInfo, bool) {
	if o == nil || o.CalleeNetwork == nil {
		return nil, false
	}
	return &o.CalleeNetwork, true
}

// HasCalleeNetwork returns a boolean if a field has been set.
func (o *MicrosoftGraphCallRecordsMedia) HasCalleeNetwork() bool {
	if o != nil && o.CalleeNetwork != nil {
		return true
	}

	return false
}

// SetCalleeNetwork gets a reference to the given AnyOfmicrosoftGraphCallRecordsNetworkInfo and assigns it to the CalleeNetwork field.
func (o *MicrosoftGraphCallRecordsMedia) SetCalleeNetwork(v AnyOfmicrosoftGraphCallRecordsNetworkInfo) {
	o.CalleeNetwork = v
}

// GetCallerDevice returns the CallerDevice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphCallRecordsMedia) GetCallerDevice() AnyOfmicrosoftGraphCallRecordsDeviceInfo {
	if o == nil  {
		var ret AnyOfmicrosoftGraphCallRecordsDeviceInfo
		return ret
	}
	return o.CallerDevice
}

// GetCallerDeviceOk returns a tuple with the CallerDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphCallRecordsMedia) GetCallerDeviceOk() (*AnyOfmicrosoftGraphCallRecordsDeviceInfo, bool) {
	if o == nil || o.CallerDevice == nil {
		return nil, false
	}
	return &o.CallerDevice, true
}

// HasCallerDevice returns a boolean if a field has been set.
func (o *MicrosoftGraphCallRecordsMedia) HasCallerDevice() bool {
	if o != nil && o.CallerDevice != nil {
		return true
	}

	return false
}

// SetCallerDevice gets a reference to the given AnyOfmicrosoftGraphCallRecordsDeviceInfo and assigns it to the CallerDevice field.
func (o *MicrosoftGraphCallRecordsMedia) SetCallerDevice(v AnyOfmicrosoftGraphCallRecordsDeviceInfo) {
	o.CallerDevice = v
}

// GetCallerNetwork returns the CallerNetwork field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphCallRecordsMedia) GetCallerNetwork() AnyOfmicrosoftGraphCallRecordsNetworkInfo {
	if o == nil  {
		var ret AnyOfmicrosoftGraphCallRecordsNetworkInfo
		return ret
	}
	return o.CallerNetwork
}

// GetCallerNetworkOk returns a tuple with the CallerNetwork field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphCallRecordsMedia) GetCallerNetworkOk() (*AnyOfmicrosoftGraphCallRecordsNetworkInfo, bool) {
	if o == nil || o.CallerNetwork == nil {
		return nil, false
	}
	return &o.CallerNetwork, true
}

// HasCallerNetwork returns a boolean if a field has been set.
func (o *MicrosoftGraphCallRecordsMedia) HasCallerNetwork() bool {
	if o != nil && o.CallerNetwork != nil {
		return true
	}

	return false
}

// SetCallerNetwork gets a reference to the given AnyOfmicrosoftGraphCallRecordsNetworkInfo and assigns it to the CallerNetwork field.
func (o *MicrosoftGraphCallRecordsMedia) SetCallerNetwork(v AnyOfmicrosoftGraphCallRecordsNetworkInfo) {
	o.CallerNetwork = v
}

// GetLabel returns the Label field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphCallRecordsMedia) GetLabel() string {
	if o == nil || o.Label.Get() == nil {
		var ret string
		return ret
	}
	return *o.Label.Get()
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphCallRecordsMedia) GetLabelOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Label.Get(), o.Label.IsSet()
}

// HasLabel returns a boolean if a field has been set.
func (o *MicrosoftGraphCallRecordsMedia) HasLabel() bool {
	if o != nil && o.Label.IsSet() {
		return true
	}

	return false
}

// SetLabel gets a reference to the given NullableString and assigns it to the Label field.
func (o *MicrosoftGraphCallRecordsMedia) SetLabel(v string) {
	o.Label.Set(&v)
}
// SetLabelNil sets the value for Label to be an explicit nil
func (o *MicrosoftGraphCallRecordsMedia) SetLabelNil() {
	o.Label.Set(nil)
}

// UnsetLabel ensures that no value is present for Label, not even an explicit nil
func (o *MicrosoftGraphCallRecordsMedia) UnsetLabel() {
	o.Label.Unset()
}

// GetStreams returns the Streams field value if set, zero value otherwise.
func (o *MicrosoftGraphCallRecordsMedia) GetStreams() []*AnyOfmicrosoftGraphCallRecordsMediaStream {
	if o == nil || o.Streams == nil {
		var ret []*AnyOfmicrosoftGraphCallRecordsMediaStream
		return ret
	}
	return *o.Streams
}

// GetStreamsOk returns a tuple with the Streams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphCallRecordsMedia) GetStreamsOk() (*[]*AnyOfmicrosoftGraphCallRecordsMediaStream, bool) {
	if o == nil || o.Streams == nil {
		return nil, false
	}
	return o.Streams, true
}

// HasStreams returns a boolean if a field has been set.
func (o *MicrosoftGraphCallRecordsMedia) HasStreams() bool {
	if o != nil && o.Streams != nil {
		return true
	}

	return false
}

// SetStreams gets a reference to the given []*AnyOfmicrosoftGraphCallRecordsMediaStream and assigns it to the Streams field.
func (o *MicrosoftGraphCallRecordsMedia) SetStreams(v []*AnyOfmicrosoftGraphCallRecordsMediaStream) {
	o.Streams = &v
}

func (o MicrosoftGraphCallRecordsMedia) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CalleeDevice != nil {
		toSerialize["calleeDevice"] = o.CalleeDevice
	}
	if o.CalleeNetwork != nil {
		toSerialize["calleeNetwork"] = o.CalleeNetwork
	}
	if o.CallerDevice != nil {
		toSerialize["callerDevice"] = o.CallerDevice
	}
	if o.CallerNetwork != nil {
		toSerialize["callerNetwork"] = o.CallerNetwork
	}
	if o.Label.IsSet() {
		toSerialize["label"] = o.Label.Get()
	}
	if o.Streams != nil {
		toSerialize["streams"] = o.Streams
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphCallRecordsMedia struct {
	value *MicrosoftGraphCallRecordsMedia
	isSet bool
}

func (v NullableMicrosoftGraphCallRecordsMedia) Get() *MicrosoftGraphCallRecordsMedia {
	return v.value
}

func (v *NullableMicrosoftGraphCallRecordsMedia) Set(val *MicrosoftGraphCallRecordsMedia) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphCallRecordsMedia) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphCallRecordsMedia) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphCallRecordsMedia(val *MicrosoftGraphCallRecordsMedia) *NullableMicrosoftGraphCallRecordsMedia {
	return &NullableMicrosoftGraphCallRecordsMedia{value: val, isSet: true}
}

func (v NullableMicrosoftGraphCallRecordsMedia) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphCallRecordsMedia) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


