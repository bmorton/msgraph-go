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

// MicrosoftGraphOnenotePatchContentCommand struct for MicrosoftGraphOnenotePatchContentCommand
type MicrosoftGraphOnenotePatchContentCommand struct {
	// The action to perform on the target element. The possible values are: replace, append, delete, insert, or prepend.
	Action AnyOfmicrosoftGraphOnenotePatchActionType `json:"action,omitempty"`
	// A string of well-formed HTML to add to the page, and any image or file binary data. If the content contains binary data, the request must be sent using the multipart/form-data content type with a 'Commands' part.
	Content NullableString `json:"content,omitempty"`
	// The location to add the supplied content, relative to the target element. The possible values are: after (default) or before.
	Position AnyOfmicrosoftGraphOnenotePatchInsertPosition `json:"position,omitempty"`
	// The element to update. Must be the #<data-id> or the generated <id> of the element, or the body or title keyword.
	Target *string `json:"target,omitempty"`
}

// NewMicrosoftGraphOnenotePatchContentCommand instantiates a new MicrosoftGraphOnenotePatchContentCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphOnenotePatchContentCommand() *MicrosoftGraphOnenotePatchContentCommand {
	this := MicrosoftGraphOnenotePatchContentCommand{}
	return &this
}

// NewMicrosoftGraphOnenotePatchContentCommandWithDefaults instantiates a new MicrosoftGraphOnenotePatchContentCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphOnenotePatchContentCommandWithDefaults() *MicrosoftGraphOnenotePatchContentCommand {
	this := MicrosoftGraphOnenotePatchContentCommand{}
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphOnenotePatchContentCommand) GetAction() AnyOfmicrosoftGraphOnenotePatchActionType {
	if o == nil  {
		var ret AnyOfmicrosoftGraphOnenotePatchActionType
		return ret
	}
	return o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphOnenotePatchContentCommand) GetActionOk() (*AnyOfmicrosoftGraphOnenotePatchActionType, bool) {
	if o == nil || o.Action == nil {
		return nil, false
	}
	return &o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *MicrosoftGraphOnenotePatchContentCommand) HasAction() bool {
	if o != nil && o.Action != nil {
		return true
	}

	return false
}

// SetAction gets a reference to the given AnyOfmicrosoftGraphOnenotePatchActionType and assigns it to the Action field.
func (o *MicrosoftGraphOnenotePatchContentCommand) SetAction(v AnyOfmicrosoftGraphOnenotePatchActionType) {
	o.Action = v
}

// GetContent returns the Content field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphOnenotePatchContentCommand) GetContent() string {
	if o == nil || o.Content.Get() == nil {
		var ret string
		return ret
	}
	return *o.Content.Get()
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphOnenotePatchContentCommand) GetContentOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Content.Get(), o.Content.IsSet()
}

// HasContent returns a boolean if a field has been set.
func (o *MicrosoftGraphOnenotePatchContentCommand) HasContent() bool {
	if o != nil && o.Content.IsSet() {
		return true
	}

	return false
}

// SetContent gets a reference to the given NullableString and assigns it to the Content field.
func (o *MicrosoftGraphOnenotePatchContentCommand) SetContent(v string) {
	o.Content.Set(&v)
}
// SetContentNil sets the value for Content to be an explicit nil
func (o *MicrosoftGraphOnenotePatchContentCommand) SetContentNil() {
	o.Content.Set(nil)
}

// UnsetContent ensures that no value is present for Content, not even an explicit nil
func (o *MicrosoftGraphOnenotePatchContentCommand) UnsetContent() {
	o.Content.Unset()
}

// GetPosition returns the Position field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphOnenotePatchContentCommand) GetPosition() AnyOfmicrosoftGraphOnenotePatchInsertPosition {
	if o == nil  {
		var ret AnyOfmicrosoftGraphOnenotePatchInsertPosition
		return ret
	}
	return o.Position
}

// GetPositionOk returns a tuple with the Position field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphOnenotePatchContentCommand) GetPositionOk() (*AnyOfmicrosoftGraphOnenotePatchInsertPosition, bool) {
	if o == nil || o.Position == nil {
		return nil, false
	}
	return &o.Position, true
}

// HasPosition returns a boolean if a field has been set.
func (o *MicrosoftGraphOnenotePatchContentCommand) HasPosition() bool {
	if o != nil && o.Position != nil {
		return true
	}

	return false
}

// SetPosition gets a reference to the given AnyOfmicrosoftGraphOnenotePatchInsertPosition and assigns it to the Position field.
func (o *MicrosoftGraphOnenotePatchContentCommand) SetPosition(v AnyOfmicrosoftGraphOnenotePatchInsertPosition) {
	o.Position = v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *MicrosoftGraphOnenotePatchContentCommand) GetTarget() string {
	if o == nil || o.Target == nil {
		var ret string
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphOnenotePatchContentCommand) GetTargetOk() (*string, bool) {
	if o == nil || o.Target == nil {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *MicrosoftGraphOnenotePatchContentCommand) HasTarget() bool {
	if o != nil && o.Target != nil {
		return true
	}

	return false
}

// SetTarget gets a reference to the given string and assigns it to the Target field.
func (o *MicrosoftGraphOnenotePatchContentCommand) SetTarget(v string) {
	o.Target = &v
}

func (o MicrosoftGraphOnenotePatchContentCommand) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Action != nil {
		toSerialize["action"] = o.Action
	}
	if o.Content.IsSet() {
		toSerialize["content"] = o.Content.Get()
	}
	if o.Position != nil {
		toSerialize["position"] = o.Position
	}
	if o.Target != nil {
		toSerialize["target"] = o.Target
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphOnenotePatchContentCommand struct {
	value *MicrosoftGraphOnenotePatchContentCommand
	isSet bool
}

func (v NullableMicrosoftGraphOnenotePatchContentCommand) Get() *MicrosoftGraphOnenotePatchContentCommand {
	return v.value
}

func (v *NullableMicrosoftGraphOnenotePatchContentCommand) Set(val *MicrosoftGraphOnenotePatchContentCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphOnenotePatchContentCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphOnenotePatchContentCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphOnenotePatchContentCommand(val *MicrosoftGraphOnenotePatchContentCommand) *NullableMicrosoftGraphOnenotePatchContentCommand {
	return &NullableMicrosoftGraphOnenotePatchContentCommand{value: val, isSet: true}
}

func (v NullableMicrosoftGraphOnenotePatchContentCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphOnenotePatchContentCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


