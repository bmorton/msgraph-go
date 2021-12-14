/*
Partial Graph API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// MicrosoftGraphPlannerTaskDetails struct for MicrosoftGraphPlannerTaskDetails
type MicrosoftGraphPlannerTaskDetails struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	// The collection of checklist items on the task.
	Checklist AnyOfobject `json:"checklist,omitempty"`
	// Description of the task
	Description NullableString `json:"description,omitempty"`
	// This sets the type of preview that shows up on the task. The possible values are: automatic, noPreview, checklist, description, reference. When set to automatic the displayed preview is chosen by the app viewing the task.
	PreviewType AnyOfmicrosoftGraphPlannerPreviewType `json:"previewType,omitempty"`
	// The collection of references on the task.
	References AnyOfobject `json:"references,omitempty"`
}

// NewMicrosoftGraphPlannerTaskDetails instantiates a new MicrosoftGraphPlannerTaskDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphPlannerTaskDetails() *MicrosoftGraphPlannerTaskDetails {
	this := MicrosoftGraphPlannerTaskDetails{}
	return &this
}

// NewMicrosoftGraphPlannerTaskDetailsWithDefaults instantiates a new MicrosoftGraphPlannerTaskDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphPlannerTaskDetailsWithDefaults() *MicrosoftGraphPlannerTaskDetails {
	this := MicrosoftGraphPlannerTaskDetails{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MicrosoftGraphPlannerTaskDetails) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphPlannerTaskDetails) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphPlannerTaskDetails) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphPlannerTaskDetails) SetId(v string) {
	o.Id = &v
}

// GetChecklist returns the Checklist field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphPlannerTaskDetails) GetChecklist() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Checklist
}

// GetChecklistOk returns a tuple with the Checklist field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphPlannerTaskDetails) GetChecklistOk() (*AnyOfobject, bool) {
	if o == nil || o.Checklist == nil {
		return nil, false
	}
	return &o.Checklist, true
}

// HasChecklist returns a boolean if a field has been set.
func (o *MicrosoftGraphPlannerTaskDetails) HasChecklist() bool {
	if o != nil && o.Checklist != nil {
		return true
	}

	return false
}

// SetChecklist gets a reference to the given AnyOfobject and assigns it to the Checklist field.
func (o *MicrosoftGraphPlannerTaskDetails) SetChecklist(v AnyOfobject) {
	o.Checklist = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphPlannerTaskDetails) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphPlannerTaskDetails) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *MicrosoftGraphPlannerTaskDetails) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *MicrosoftGraphPlannerTaskDetails) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *MicrosoftGraphPlannerTaskDetails) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *MicrosoftGraphPlannerTaskDetails) UnsetDescription() {
	o.Description.Unset()
}

// GetPreviewType returns the PreviewType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphPlannerTaskDetails) GetPreviewType() AnyOfmicrosoftGraphPlannerPreviewType {
	if o == nil  {
		var ret AnyOfmicrosoftGraphPlannerPreviewType
		return ret
	}
	return o.PreviewType
}

// GetPreviewTypeOk returns a tuple with the PreviewType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphPlannerTaskDetails) GetPreviewTypeOk() (*AnyOfmicrosoftGraphPlannerPreviewType, bool) {
	if o == nil || o.PreviewType == nil {
		return nil, false
	}
	return &o.PreviewType, true
}

// HasPreviewType returns a boolean if a field has been set.
func (o *MicrosoftGraphPlannerTaskDetails) HasPreviewType() bool {
	if o != nil && o.PreviewType != nil {
		return true
	}

	return false
}

// SetPreviewType gets a reference to the given AnyOfmicrosoftGraphPlannerPreviewType and assigns it to the PreviewType field.
func (o *MicrosoftGraphPlannerTaskDetails) SetPreviewType(v AnyOfmicrosoftGraphPlannerPreviewType) {
	o.PreviewType = v
}

// GetReferences returns the References field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphPlannerTaskDetails) GetReferences() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.References
}

// GetReferencesOk returns a tuple with the References field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphPlannerTaskDetails) GetReferencesOk() (*AnyOfobject, bool) {
	if o == nil || o.References == nil {
		return nil, false
	}
	return &o.References, true
}

// HasReferences returns a boolean if a field has been set.
func (o *MicrosoftGraphPlannerTaskDetails) HasReferences() bool {
	if o != nil && o.References != nil {
		return true
	}

	return false
}

// SetReferences gets a reference to the given AnyOfobject and assigns it to the References field.
func (o *MicrosoftGraphPlannerTaskDetails) SetReferences(v AnyOfobject) {
	o.References = v
}

func (o MicrosoftGraphPlannerTaskDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Checklist != nil {
		toSerialize["checklist"] = o.Checklist
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.PreviewType != nil {
		toSerialize["previewType"] = o.PreviewType
	}
	if o.References != nil {
		toSerialize["references"] = o.References
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphPlannerTaskDetails struct {
	value *MicrosoftGraphPlannerTaskDetails
	isSet bool
}

func (v NullableMicrosoftGraphPlannerTaskDetails) Get() *MicrosoftGraphPlannerTaskDetails {
	return v.value
}

func (v *NullableMicrosoftGraphPlannerTaskDetails) Set(val *MicrosoftGraphPlannerTaskDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphPlannerTaskDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphPlannerTaskDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphPlannerTaskDetails(val *MicrosoftGraphPlannerTaskDetails) *NullableMicrosoftGraphPlannerTaskDetails {
	return &NullableMicrosoftGraphPlannerTaskDetails{value: val, isSet: true}
}

func (v NullableMicrosoftGraphPlannerTaskDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphPlannerTaskDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


