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

// InlineObject26 struct for InlineObject26
type InlineObject26 struct {
	Prompts *[]*AnyOfobject `json:"prompts,omitempty"`
	BargeInAllowed NullableBool `json:"bargeInAllowed,omitempty"`
	InitialSilenceTimeoutInSeconds NullableInt32 `json:"initialSilenceTimeoutInSeconds,omitempty"`
	MaxSilenceTimeoutInSeconds NullableInt32 `json:"maxSilenceTimeoutInSeconds,omitempty"`
	MaxRecordDurationInSeconds NullableInt32 `json:"maxRecordDurationInSeconds,omitempty"`
	PlayBeep NullableBool `json:"playBeep,omitempty"`
	StopTones *[]*string `json:"stopTones,omitempty"`
	ClientContext NullableString `json:"clientContext,omitempty"`
}

// NewInlineObject26 instantiates a new InlineObject26 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject26() *InlineObject26 {
	this := InlineObject26{}
	var bargeInAllowed bool = false
	this.BargeInAllowed = *NewNullableBool(&bargeInAllowed)
	var playBeep bool = false
	this.PlayBeep = *NewNullableBool(&playBeep)
	return &this
}

// NewInlineObject26WithDefaults instantiates a new InlineObject26 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject26WithDefaults() *InlineObject26 {
	this := InlineObject26{}
	var bargeInAllowed bool = false
	this.BargeInAllowed = *NewNullableBool(&bargeInAllowed)
	var playBeep bool = false
	this.PlayBeep = *NewNullableBool(&playBeep)
	return &this
}

// GetPrompts returns the Prompts field value if set, zero value otherwise.
func (o *InlineObject26) GetPrompts() []*AnyOfobject {
	if o == nil || o.Prompts == nil {
		var ret []*AnyOfobject
		return ret
	}
	return *o.Prompts
}

// GetPromptsOk returns a tuple with the Prompts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject26) GetPromptsOk() (*[]*AnyOfobject, bool) {
	if o == nil || o.Prompts == nil {
		return nil, false
	}
	return o.Prompts, true
}

// HasPrompts returns a boolean if a field has been set.
func (o *InlineObject26) HasPrompts() bool {
	if o != nil && o.Prompts != nil {
		return true
	}

	return false
}

// SetPrompts gets a reference to the given []*AnyOfobject and assigns it to the Prompts field.
func (o *InlineObject26) SetPrompts(v []*AnyOfobject) {
	o.Prompts = &v
}

// GetBargeInAllowed returns the BargeInAllowed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject26) GetBargeInAllowed() bool {
	if o == nil || o.BargeInAllowed.Get() == nil {
		var ret bool
		return ret
	}
	return *o.BargeInAllowed.Get()
}

// GetBargeInAllowedOk returns a tuple with the BargeInAllowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject26) GetBargeInAllowedOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.BargeInAllowed.Get(), o.BargeInAllowed.IsSet()
}

// HasBargeInAllowed returns a boolean if a field has been set.
func (o *InlineObject26) HasBargeInAllowed() bool {
	if o != nil && o.BargeInAllowed.IsSet() {
		return true
	}

	return false
}

// SetBargeInAllowed gets a reference to the given NullableBool and assigns it to the BargeInAllowed field.
func (o *InlineObject26) SetBargeInAllowed(v bool) {
	o.BargeInAllowed.Set(&v)
}
// SetBargeInAllowedNil sets the value for BargeInAllowed to be an explicit nil
func (o *InlineObject26) SetBargeInAllowedNil() {
	o.BargeInAllowed.Set(nil)
}

// UnsetBargeInAllowed ensures that no value is present for BargeInAllowed, not even an explicit nil
func (o *InlineObject26) UnsetBargeInAllowed() {
	o.BargeInAllowed.Unset()
}

// GetInitialSilenceTimeoutInSeconds returns the InitialSilenceTimeoutInSeconds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject26) GetInitialSilenceTimeoutInSeconds() int32 {
	if o == nil || o.InitialSilenceTimeoutInSeconds.Get() == nil {
		var ret int32
		return ret
	}
	return *o.InitialSilenceTimeoutInSeconds.Get()
}

// GetInitialSilenceTimeoutInSecondsOk returns a tuple with the InitialSilenceTimeoutInSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject26) GetInitialSilenceTimeoutInSecondsOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.InitialSilenceTimeoutInSeconds.Get(), o.InitialSilenceTimeoutInSeconds.IsSet()
}

// HasInitialSilenceTimeoutInSeconds returns a boolean if a field has been set.
func (o *InlineObject26) HasInitialSilenceTimeoutInSeconds() bool {
	if o != nil && o.InitialSilenceTimeoutInSeconds.IsSet() {
		return true
	}

	return false
}

// SetInitialSilenceTimeoutInSeconds gets a reference to the given NullableInt32 and assigns it to the InitialSilenceTimeoutInSeconds field.
func (o *InlineObject26) SetInitialSilenceTimeoutInSeconds(v int32) {
	o.InitialSilenceTimeoutInSeconds.Set(&v)
}
// SetInitialSilenceTimeoutInSecondsNil sets the value for InitialSilenceTimeoutInSeconds to be an explicit nil
func (o *InlineObject26) SetInitialSilenceTimeoutInSecondsNil() {
	o.InitialSilenceTimeoutInSeconds.Set(nil)
}

// UnsetInitialSilenceTimeoutInSeconds ensures that no value is present for InitialSilenceTimeoutInSeconds, not even an explicit nil
func (o *InlineObject26) UnsetInitialSilenceTimeoutInSeconds() {
	o.InitialSilenceTimeoutInSeconds.Unset()
}

// GetMaxSilenceTimeoutInSeconds returns the MaxSilenceTimeoutInSeconds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject26) GetMaxSilenceTimeoutInSeconds() int32 {
	if o == nil || o.MaxSilenceTimeoutInSeconds.Get() == nil {
		var ret int32
		return ret
	}
	return *o.MaxSilenceTimeoutInSeconds.Get()
}

// GetMaxSilenceTimeoutInSecondsOk returns a tuple with the MaxSilenceTimeoutInSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject26) GetMaxSilenceTimeoutInSecondsOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.MaxSilenceTimeoutInSeconds.Get(), o.MaxSilenceTimeoutInSeconds.IsSet()
}

// HasMaxSilenceTimeoutInSeconds returns a boolean if a field has been set.
func (o *InlineObject26) HasMaxSilenceTimeoutInSeconds() bool {
	if o != nil && o.MaxSilenceTimeoutInSeconds.IsSet() {
		return true
	}

	return false
}

// SetMaxSilenceTimeoutInSeconds gets a reference to the given NullableInt32 and assigns it to the MaxSilenceTimeoutInSeconds field.
func (o *InlineObject26) SetMaxSilenceTimeoutInSeconds(v int32) {
	o.MaxSilenceTimeoutInSeconds.Set(&v)
}
// SetMaxSilenceTimeoutInSecondsNil sets the value for MaxSilenceTimeoutInSeconds to be an explicit nil
func (o *InlineObject26) SetMaxSilenceTimeoutInSecondsNil() {
	o.MaxSilenceTimeoutInSeconds.Set(nil)
}

// UnsetMaxSilenceTimeoutInSeconds ensures that no value is present for MaxSilenceTimeoutInSeconds, not even an explicit nil
func (o *InlineObject26) UnsetMaxSilenceTimeoutInSeconds() {
	o.MaxSilenceTimeoutInSeconds.Unset()
}

// GetMaxRecordDurationInSeconds returns the MaxRecordDurationInSeconds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject26) GetMaxRecordDurationInSeconds() int32 {
	if o == nil || o.MaxRecordDurationInSeconds.Get() == nil {
		var ret int32
		return ret
	}
	return *o.MaxRecordDurationInSeconds.Get()
}

// GetMaxRecordDurationInSecondsOk returns a tuple with the MaxRecordDurationInSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject26) GetMaxRecordDurationInSecondsOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.MaxRecordDurationInSeconds.Get(), o.MaxRecordDurationInSeconds.IsSet()
}

// HasMaxRecordDurationInSeconds returns a boolean if a field has been set.
func (o *InlineObject26) HasMaxRecordDurationInSeconds() bool {
	if o != nil && o.MaxRecordDurationInSeconds.IsSet() {
		return true
	}

	return false
}

// SetMaxRecordDurationInSeconds gets a reference to the given NullableInt32 and assigns it to the MaxRecordDurationInSeconds field.
func (o *InlineObject26) SetMaxRecordDurationInSeconds(v int32) {
	o.MaxRecordDurationInSeconds.Set(&v)
}
// SetMaxRecordDurationInSecondsNil sets the value for MaxRecordDurationInSeconds to be an explicit nil
func (o *InlineObject26) SetMaxRecordDurationInSecondsNil() {
	o.MaxRecordDurationInSeconds.Set(nil)
}

// UnsetMaxRecordDurationInSeconds ensures that no value is present for MaxRecordDurationInSeconds, not even an explicit nil
func (o *InlineObject26) UnsetMaxRecordDurationInSeconds() {
	o.MaxRecordDurationInSeconds.Unset()
}

// GetPlayBeep returns the PlayBeep field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject26) GetPlayBeep() bool {
	if o == nil || o.PlayBeep.Get() == nil {
		var ret bool
		return ret
	}
	return *o.PlayBeep.Get()
}

// GetPlayBeepOk returns a tuple with the PlayBeep field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject26) GetPlayBeepOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PlayBeep.Get(), o.PlayBeep.IsSet()
}

// HasPlayBeep returns a boolean if a field has been set.
func (o *InlineObject26) HasPlayBeep() bool {
	if o != nil && o.PlayBeep.IsSet() {
		return true
	}

	return false
}

// SetPlayBeep gets a reference to the given NullableBool and assigns it to the PlayBeep field.
func (o *InlineObject26) SetPlayBeep(v bool) {
	o.PlayBeep.Set(&v)
}
// SetPlayBeepNil sets the value for PlayBeep to be an explicit nil
func (o *InlineObject26) SetPlayBeepNil() {
	o.PlayBeep.Set(nil)
}

// UnsetPlayBeep ensures that no value is present for PlayBeep, not even an explicit nil
func (o *InlineObject26) UnsetPlayBeep() {
	o.PlayBeep.Unset()
}

// GetStopTones returns the StopTones field value if set, zero value otherwise.
func (o *InlineObject26) GetStopTones() []*string {
	if o == nil || o.StopTones == nil {
		var ret []*string
		return ret
	}
	return *o.StopTones
}

// GetStopTonesOk returns a tuple with the StopTones field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject26) GetStopTonesOk() (*[]*string, bool) {
	if o == nil || o.StopTones == nil {
		return nil, false
	}
	return o.StopTones, true
}

// HasStopTones returns a boolean if a field has been set.
func (o *InlineObject26) HasStopTones() bool {
	if o != nil && o.StopTones != nil {
		return true
	}

	return false
}

// SetStopTones gets a reference to the given []*string and assigns it to the StopTones field.
func (o *InlineObject26) SetStopTones(v []*string) {
	o.StopTones = &v
}

// GetClientContext returns the ClientContext field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject26) GetClientContext() string {
	if o == nil || o.ClientContext.Get() == nil {
		var ret string
		return ret
	}
	return *o.ClientContext.Get()
}

// GetClientContextOk returns a tuple with the ClientContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject26) GetClientContextOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ClientContext.Get(), o.ClientContext.IsSet()
}

// HasClientContext returns a boolean if a field has been set.
func (o *InlineObject26) HasClientContext() bool {
	if o != nil && o.ClientContext.IsSet() {
		return true
	}

	return false
}

// SetClientContext gets a reference to the given NullableString and assigns it to the ClientContext field.
func (o *InlineObject26) SetClientContext(v string) {
	o.ClientContext.Set(&v)
}
// SetClientContextNil sets the value for ClientContext to be an explicit nil
func (o *InlineObject26) SetClientContextNil() {
	o.ClientContext.Set(nil)
}

// UnsetClientContext ensures that no value is present for ClientContext, not even an explicit nil
func (o *InlineObject26) UnsetClientContext() {
	o.ClientContext.Unset()
}

func (o InlineObject26) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Prompts != nil {
		toSerialize["prompts"] = o.Prompts
	}
	if o.BargeInAllowed.IsSet() {
		toSerialize["bargeInAllowed"] = o.BargeInAllowed.Get()
	}
	if o.InitialSilenceTimeoutInSeconds.IsSet() {
		toSerialize["initialSilenceTimeoutInSeconds"] = o.InitialSilenceTimeoutInSeconds.Get()
	}
	if o.MaxSilenceTimeoutInSeconds.IsSet() {
		toSerialize["maxSilenceTimeoutInSeconds"] = o.MaxSilenceTimeoutInSeconds.Get()
	}
	if o.MaxRecordDurationInSeconds.IsSet() {
		toSerialize["maxRecordDurationInSeconds"] = o.MaxRecordDurationInSeconds.Get()
	}
	if o.PlayBeep.IsSet() {
		toSerialize["playBeep"] = o.PlayBeep.Get()
	}
	if o.StopTones != nil {
		toSerialize["stopTones"] = o.StopTones
	}
	if o.ClientContext.IsSet() {
		toSerialize["clientContext"] = o.ClientContext.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject26 struct {
	value *InlineObject26
	isSet bool
}

func (v NullableInlineObject26) Get() *InlineObject26 {
	return v.value
}

func (v *NullableInlineObject26) Set(val *InlineObject26) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject26) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject26) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject26(val *InlineObject26) *NullableInlineObject26 {
	return &NullableInlineObject26{value: val, isSet: true}
}

func (v NullableInlineObject26) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject26) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


