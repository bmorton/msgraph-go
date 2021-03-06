/*
Partial Graph API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package msgraph

import (
	"encoding/json"
	"time"
)

// Term struct for Term
type Term struct {
	// Date and time of term creation. Read-only.
	CreatedDateTime NullableTime `json:"createdDateTime,omitempty"`
	// Description about term that is dependent on the languageTag.
	Descriptions *[]*AnyOfmicrosoftGraphTermStoreLocalizedDescription `json:"descriptions,omitempty"`
	// Label metadata for a term.
	Labels *[]*AnyOfmicrosoftGraphTermStoreLocalizedLabel `json:"labels,omitempty"`
	// Last date and time of term modification. Read-only.
	LastModifiedDateTime NullableTime `json:"lastModifiedDateTime,omitempty"`
	// Collection of properties on the term.
	Properties *[]*AnyOfmicrosoftGraphKeyValue `json:"properties,omitempty"`
	// Children of current term.
	Children *[]MicrosoftGraphTermStoreTerm `json:"children,omitempty"`
	// To indicate which terms are related to the current term as either pinned or reused.
	Relations *[]MicrosoftGraphTermStoreRelation `json:"relations,omitempty"`
	// The [set] in which the term is created.
	Set AnyOfmicrosoftGraphTermStoreSet `json:"set,omitempty"`
}

// NewTerm instantiates a new Term object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTerm() *Term {
	this := Term{}
	return &this
}

// NewTermWithDefaults instantiates a new Term object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTermWithDefaults() *Term {
	this := Term{}
	return &this
}

// GetCreatedDateTime returns the CreatedDateTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Term) GetCreatedDateTime() time.Time {
	if o == nil || o.CreatedDateTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedDateTime.Get()
}

// GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Term) GetCreatedDateTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CreatedDateTime.Get(), o.CreatedDateTime.IsSet()
}

// HasCreatedDateTime returns a boolean if a field has been set.
func (o *Term) HasCreatedDateTime() bool {
	if o != nil && o.CreatedDateTime.IsSet() {
		return true
	}

	return false
}

// SetCreatedDateTime gets a reference to the given NullableTime and assigns it to the CreatedDateTime field.
func (o *Term) SetCreatedDateTime(v time.Time) {
	o.CreatedDateTime.Set(&v)
}
// SetCreatedDateTimeNil sets the value for CreatedDateTime to be an explicit nil
func (o *Term) SetCreatedDateTimeNil() {
	o.CreatedDateTime.Set(nil)
}

// UnsetCreatedDateTime ensures that no value is present for CreatedDateTime, not even an explicit nil
func (o *Term) UnsetCreatedDateTime() {
	o.CreatedDateTime.Unset()
}

// GetDescriptions returns the Descriptions field value if set, zero value otherwise.
func (o *Term) GetDescriptions() []*AnyOfmicrosoftGraphTermStoreLocalizedDescription {
	if o == nil || o.Descriptions == nil {
		var ret []*AnyOfmicrosoftGraphTermStoreLocalizedDescription
		return ret
	}
	return *o.Descriptions
}

// GetDescriptionsOk returns a tuple with the Descriptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Term) GetDescriptionsOk() (*[]*AnyOfmicrosoftGraphTermStoreLocalizedDescription, bool) {
	if o == nil || o.Descriptions == nil {
		return nil, false
	}
	return o.Descriptions, true
}

// HasDescriptions returns a boolean if a field has been set.
func (o *Term) HasDescriptions() bool {
	if o != nil && o.Descriptions != nil {
		return true
	}

	return false
}

// SetDescriptions gets a reference to the given []*AnyOfmicrosoftGraphTermStoreLocalizedDescription and assigns it to the Descriptions field.
func (o *Term) SetDescriptions(v []*AnyOfmicrosoftGraphTermStoreLocalizedDescription) {
	o.Descriptions = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *Term) GetLabels() []*AnyOfmicrosoftGraphTermStoreLocalizedLabel {
	if o == nil || o.Labels == nil {
		var ret []*AnyOfmicrosoftGraphTermStoreLocalizedLabel
		return ret
	}
	return *o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Term) GetLabelsOk() (*[]*AnyOfmicrosoftGraphTermStoreLocalizedLabel, bool) {
	if o == nil || o.Labels == nil {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *Term) HasLabels() bool {
	if o != nil && o.Labels != nil {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []*AnyOfmicrosoftGraphTermStoreLocalizedLabel and assigns it to the Labels field.
func (o *Term) SetLabels(v []*AnyOfmicrosoftGraphTermStoreLocalizedLabel) {
	o.Labels = &v
}

// GetLastModifiedDateTime returns the LastModifiedDateTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Term) GetLastModifiedDateTime() time.Time {
	if o == nil || o.LastModifiedDateTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.LastModifiedDateTime.Get()
}

// GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Term) GetLastModifiedDateTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LastModifiedDateTime.Get(), o.LastModifiedDateTime.IsSet()
}

// HasLastModifiedDateTime returns a boolean if a field has been set.
func (o *Term) HasLastModifiedDateTime() bool {
	if o != nil && o.LastModifiedDateTime.IsSet() {
		return true
	}

	return false
}

// SetLastModifiedDateTime gets a reference to the given NullableTime and assigns it to the LastModifiedDateTime field.
func (o *Term) SetLastModifiedDateTime(v time.Time) {
	o.LastModifiedDateTime.Set(&v)
}
// SetLastModifiedDateTimeNil sets the value for LastModifiedDateTime to be an explicit nil
func (o *Term) SetLastModifiedDateTimeNil() {
	o.LastModifiedDateTime.Set(nil)
}

// UnsetLastModifiedDateTime ensures that no value is present for LastModifiedDateTime, not even an explicit nil
func (o *Term) UnsetLastModifiedDateTime() {
	o.LastModifiedDateTime.Unset()
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *Term) GetProperties() []*AnyOfmicrosoftGraphKeyValue {
	if o == nil || o.Properties == nil {
		var ret []*AnyOfmicrosoftGraphKeyValue
		return ret
	}
	return *o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Term) GetPropertiesOk() (*[]*AnyOfmicrosoftGraphKeyValue, bool) {
	if o == nil || o.Properties == nil {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *Term) HasProperties() bool {
	if o != nil && o.Properties != nil {
		return true
	}

	return false
}

// SetProperties gets a reference to the given []*AnyOfmicrosoftGraphKeyValue and assigns it to the Properties field.
func (o *Term) SetProperties(v []*AnyOfmicrosoftGraphKeyValue) {
	o.Properties = &v
}

// GetChildren returns the Children field value if set, zero value otherwise.
func (o *Term) GetChildren() []MicrosoftGraphTermStoreTerm {
	if o == nil || o.Children == nil {
		var ret []MicrosoftGraphTermStoreTerm
		return ret
	}
	return *o.Children
}

// GetChildrenOk returns a tuple with the Children field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Term) GetChildrenOk() (*[]MicrosoftGraphTermStoreTerm, bool) {
	if o == nil || o.Children == nil {
		return nil, false
	}
	return o.Children, true
}

// HasChildren returns a boolean if a field has been set.
func (o *Term) HasChildren() bool {
	if o != nil && o.Children != nil {
		return true
	}

	return false
}

// SetChildren gets a reference to the given []MicrosoftGraphTermStoreTerm and assigns it to the Children field.
func (o *Term) SetChildren(v []MicrosoftGraphTermStoreTerm) {
	o.Children = &v
}

// GetRelations returns the Relations field value if set, zero value otherwise.
func (o *Term) GetRelations() []MicrosoftGraphTermStoreRelation {
	if o == nil || o.Relations == nil {
		var ret []MicrosoftGraphTermStoreRelation
		return ret
	}
	return *o.Relations
}

// GetRelationsOk returns a tuple with the Relations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Term) GetRelationsOk() (*[]MicrosoftGraphTermStoreRelation, bool) {
	if o == nil || o.Relations == nil {
		return nil, false
	}
	return o.Relations, true
}

// HasRelations returns a boolean if a field has been set.
func (o *Term) HasRelations() bool {
	if o != nil && o.Relations != nil {
		return true
	}

	return false
}

// SetRelations gets a reference to the given []MicrosoftGraphTermStoreRelation and assigns it to the Relations field.
func (o *Term) SetRelations(v []MicrosoftGraphTermStoreRelation) {
	o.Relations = &v
}

// GetSet returns the Set field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Term) GetSet() AnyOfmicrosoftGraphTermStoreSet {
	if o == nil  {
		var ret AnyOfmicrosoftGraphTermStoreSet
		return ret
	}
	return o.Set
}

// GetSetOk returns a tuple with the Set field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Term) GetSetOk() (*AnyOfmicrosoftGraphTermStoreSet, bool) {
	if o == nil || o.Set == nil {
		return nil, false
	}
	return &o.Set, true
}

// HasSet returns a boolean if a field has been set.
func (o *Term) HasSet() bool {
	if o != nil && o.Set != nil {
		return true
	}

	return false
}

// SetSet gets a reference to the given AnyOfmicrosoftGraphTermStoreSet and assigns it to the Set field.
func (o *Term) SetSet(v AnyOfmicrosoftGraphTermStoreSet) {
	o.Set = v
}

func (o Term) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CreatedDateTime.IsSet() {
		toSerialize["createdDateTime"] = o.CreatedDateTime.Get()
	}
	if o.Descriptions != nil {
		toSerialize["descriptions"] = o.Descriptions
	}
	if o.Labels != nil {
		toSerialize["labels"] = o.Labels
	}
	if o.LastModifiedDateTime.IsSet() {
		toSerialize["lastModifiedDateTime"] = o.LastModifiedDateTime.Get()
	}
	if o.Properties != nil {
		toSerialize["properties"] = o.Properties
	}
	if o.Children != nil {
		toSerialize["children"] = o.Children
	}
	if o.Relations != nil {
		toSerialize["relations"] = o.Relations
	}
	if o.Set != nil {
		toSerialize["set"] = o.Set
	}
	return json.Marshal(toSerialize)
}

type NullableTerm struct {
	value *Term
	isSet bool
}

func (v NullableTerm) Get() *Term {
	return v.value
}

func (v *NullableTerm) Set(val *Term) {
	v.value = val
	v.isSet = true
}

func (v NullableTerm) IsSet() bool {
	return v.isSet
}

func (v *NullableTerm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTerm(val *Term) *NullableTerm {
	return &NullableTerm{value: val, isSet: true}
}

func (v NullableTerm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTerm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


