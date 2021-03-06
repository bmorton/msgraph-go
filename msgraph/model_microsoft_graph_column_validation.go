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

// MicrosoftGraphColumnValidation struct for MicrosoftGraphColumnValidation
type MicrosoftGraphColumnValidation struct {
	// Default BCP 47 language tag for the description.
	DefaultLanguage NullableString `json:"defaultLanguage,omitempty"`
	// Localized messages that explain what is needed for this column's value to be considered valid. User will be prompted with this message if validation fails.
	Descriptions *[]*AnyOfmicrosoftGraphDisplayNameLocalization `json:"descriptions,omitempty"`
	// The formula to validate column value. For examples, see Examples of common formulas in lists.
	Formula NullableString `json:"formula,omitempty"`
}

// NewMicrosoftGraphColumnValidation instantiates a new MicrosoftGraphColumnValidation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphColumnValidation() *MicrosoftGraphColumnValidation {
	this := MicrosoftGraphColumnValidation{}
	return &this
}

// NewMicrosoftGraphColumnValidationWithDefaults instantiates a new MicrosoftGraphColumnValidation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphColumnValidationWithDefaults() *MicrosoftGraphColumnValidation {
	this := MicrosoftGraphColumnValidation{}
	return &this
}

// GetDefaultLanguage returns the DefaultLanguage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphColumnValidation) GetDefaultLanguage() string {
	if o == nil || o.DefaultLanguage.Get() == nil {
		var ret string
		return ret
	}
	return *o.DefaultLanguage.Get()
}

// GetDefaultLanguageOk returns a tuple with the DefaultLanguage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphColumnValidation) GetDefaultLanguageOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DefaultLanguage.Get(), o.DefaultLanguage.IsSet()
}

// HasDefaultLanguage returns a boolean if a field has been set.
func (o *MicrosoftGraphColumnValidation) HasDefaultLanguage() bool {
	if o != nil && o.DefaultLanguage.IsSet() {
		return true
	}

	return false
}

// SetDefaultLanguage gets a reference to the given NullableString and assigns it to the DefaultLanguage field.
func (o *MicrosoftGraphColumnValidation) SetDefaultLanguage(v string) {
	o.DefaultLanguage.Set(&v)
}
// SetDefaultLanguageNil sets the value for DefaultLanguage to be an explicit nil
func (o *MicrosoftGraphColumnValidation) SetDefaultLanguageNil() {
	o.DefaultLanguage.Set(nil)
}

// UnsetDefaultLanguage ensures that no value is present for DefaultLanguage, not even an explicit nil
func (o *MicrosoftGraphColumnValidation) UnsetDefaultLanguage() {
	o.DefaultLanguage.Unset()
}

// GetDescriptions returns the Descriptions field value if set, zero value otherwise.
func (o *MicrosoftGraphColumnValidation) GetDescriptions() []*AnyOfmicrosoftGraphDisplayNameLocalization {
	if o == nil || o.Descriptions == nil {
		var ret []*AnyOfmicrosoftGraphDisplayNameLocalization
		return ret
	}
	return *o.Descriptions
}

// GetDescriptionsOk returns a tuple with the Descriptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphColumnValidation) GetDescriptionsOk() (*[]*AnyOfmicrosoftGraphDisplayNameLocalization, bool) {
	if o == nil || o.Descriptions == nil {
		return nil, false
	}
	return o.Descriptions, true
}

// HasDescriptions returns a boolean if a field has been set.
func (o *MicrosoftGraphColumnValidation) HasDescriptions() bool {
	if o != nil && o.Descriptions != nil {
		return true
	}

	return false
}

// SetDescriptions gets a reference to the given []*AnyOfmicrosoftGraphDisplayNameLocalization and assigns it to the Descriptions field.
func (o *MicrosoftGraphColumnValidation) SetDescriptions(v []*AnyOfmicrosoftGraphDisplayNameLocalization) {
	o.Descriptions = &v
}

// GetFormula returns the Formula field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphColumnValidation) GetFormula() string {
	if o == nil || o.Formula.Get() == nil {
		var ret string
		return ret
	}
	return *o.Formula.Get()
}

// GetFormulaOk returns a tuple with the Formula field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphColumnValidation) GetFormulaOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Formula.Get(), o.Formula.IsSet()
}

// HasFormula returns a boolean if a field has been set.
func (o *MicrosoftGraphColumnValidation) HasFormula() bool {
	if o != nil && o.Formula.IsSet() {
		return true
	}

	return false
}

// SetFormula gets a reference to the given NullableString and assigns it to the Formula field.
func (o *MicrosoftGraphColumnValidation) SetFormula(v string) {
	o.Formula.Set(&v)
}
// SetFormulaNil sets the value for Formula to be an explicit nil
func (o *MicrosoftGraphColumnValidation) SetFormulaNil() {
	o.Formula.Set(nil)
}

// UnsetFormula ensures that no value is present for Formula, not even an explicit nil
func (o *MicrosoftGraphColumnValidation) UnsetFormula() {
	o.Formula.Unset()
}

func (o MicrosoftGraphColumnValidation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DefaultLanguage.IsSet() {
		toSerialize["defaultLanguage"] = o.DefaultLanguage.Get()
	}
	if o.Descriptions != nil {
		toSerialize["descriptions"] = o.Descriptions
	}
	if o.Formula.IsSet() {
		toSerialize["formula"] = o.Formula.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphColumnValidation struct {
	value *MicrosoftGraphColumnValidation
	isSet bool
}

func (v NullableMicrosoftGraphColumnValidation) Get() *MicrosoftGraphColumnValidation {
	return v.value
}

func (v *NullableMicrosoftGraphColumnValidation) Set(val *MicrosoftGraphColumnValidation) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphColumnValidation) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphColumnValidation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphColumnValidation(val *MicrosoftGraphColumnValidation) *NullableMicrosoftGraphColumnValidation {
	return &NullableMicrosoftGraphColumnValidation{value: val, isSet: true}
}

func (v NullableMicrosoftGraphColumnValidation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphColumnValidation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


