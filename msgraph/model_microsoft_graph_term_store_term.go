/*
Partial Graph API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// MicrosoftGraphTermStoreTerm struct for MicrosoftGraphTermStoreTerm
type MicrosoftGraphTermStoreTerm struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
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

// NewMicrosoftGraphTermStoreTerm instantiates a new MicrosoftGraphTermStoreTerm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphTermStoreTerm() *MicrosoftGraphTermStoreTerm {
	this := MicrosoftGraphTermStoreTerm{}
	return &this
}

// NewMicrosoftGraphTermStoreTermWithDefaults instantiates a new MicrosoftGraphTermStoreTerm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphTermStoreTermWithDefaults() *MicrosoftGraphTermStoreTerm {
	this := MicrosoftGraphTermStoreTerm{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MicrosoftGraphTermStoreTerm) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphTermStoreTerm) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphTermStoreTerm) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphTermStoreTerm) SetId(v string) {
	o.Id = &v
}

// GetCreatedDateTime returns the CreatedDateTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphTermStoreTerm) GetCreatedDateTime() time.Time {
	if o == nil || o.CreatedDateTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedDateTime.Get()
}

// GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphTermStoreTerm) GetCreatedDateTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CreatedDateTime.Get(), o.CreatedDateTime.IsSet()
}

// HasCreatedDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphTermStoreTerm) HasCreatedDateTime() bool {
	if o != nil && o.CreatedDateTime.IsSet() {
		return true
	}

	return false
}

// SetCreatedDateTime gets a reference to the given NullableTime and assigns it to the CreatedDateTime field.
func (o *MicrosoftGraphTermStoreTerm) SetCreatedDateTime(v time.Time) {
	o.CreatedDateTime.Set(&v)
}
// SetCreatedDateTimeNil sets the value for CreatedDateTime to be an explicit nil
func (o *MicrosoftGraphTermStoreTerm) SetCreatedDateTimeNil() {
	o.CreatedDateTime.Set(nil)
}

// UnsetCreatedDateTime ensures that no value is present for CreatedDateTime, not even an explicit nil
func (o *MicrosoftGraphTermStoreTerm) UnsetCreatedDateTime() {
	o.CreatedDateTime.Unset()
}

// GetDescriptions returns the Descriptions field value if set, zero value otherwise.
func (o *MicrosoftGraphTermStoreTerm) GetDescriptions() []*AnyOfmicrosoftGraphTermStoreLocalizedDescription {
	if o == nil || o.Descriptions == nil {
		var ret []*AnyOfmicrosoftGraphTermStoreLocalizedDescription
		return ret
	}
	return *o.Descriptions
}

// GetDescriptionsOk returns a tuple with the Descriptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphTermStoreTerm) GetDescriptionsOk() (*[]*AnyOfmicrosoftGraphTermStoreLocalizedDescription, bool) {
	if o == nil || o.Descriptions == nil {
		return nil, false
	}
	return o.Descriptions, true
}

// HasDescriptions returns a boolean if a field has been set.
func (o *MicrosoftGraphTermStoreTerm) HasDescriptions() bool {
	if o != nil && o.Descriptions != nil {
		return true
	}

	return false
}

// SetDescriptions gets a reference to the given []*AnyOfmicrosoftGraphTermStoreLocalizedDescription and assigns it to the Descriptions field.
func (o *MicrosoftGraphTermStoreTerm) SetDescriptions(v []*AnyOfmicrosoftGraphTermStoreLocalizedDescription) {
	o.Descriptions = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *MicrosoftGraphTermStoreTerm) GetLabels() []*AnyOfmicrosoftGraphTermStoreLocalizedLabel {
	if o == nil || o.Labels == nil {
		var ret []*AnyOfmicrosoftGraphTermStoreLocalizedLabel
		return ret
	}
	return *o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphTermStoreTerm) GetLabelsOk() (*[]*AnyOfmicrosoftGraphTermStoreLocalizedLabel, bool) {
	if o == nil || o.Labels == nil {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *MicrosoftGraphTermStoreTerm) HasLabels() bool {
	if o != nil && o.Labels != nil {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []*AnyOfmicrosoftGraphTermStoreLocalizedLabel and assigns it to the Labels field.
func (o *MicrosoftGraphTermStoreTerm) SetLabels(v []*AnyOfmicrosoftGraphTermStoreLocalizedLabel) {
	o.Labels = &v
}

// GetLastModifiedDateTime returns the LastModifiedDateTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphTermStoreTerm) GetLastModifiedDateTime() time.Time {
	if o == nil || o.LastModifiedDateTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.LastModifiedDateTime.Get()
}

// GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphTermStoreTerm) GetLastModifiedDateTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LastModifiedDateTime.Get(), o.LastModifiedDateTime.IsSet()
}

// HasLastModifiedDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphTermStoreTerm) HasLastModifiedDateTime() bool {
	if o != nil && o.LastModifiedDateTime.IsSet() {
		return true
	}

	return false
}

// SetLastModifiedDateTime gets a reference to the given NullableTime and assigns it to the LastModifiedDateTime field.
func (o *MicrosoftGraphTermStoreTerm) SetLastModifiedDateTime(v time.Time) {
	o.LastModifiedDateTime.Set(&v)
}
// SetLastModifiedDateTimeNil sets the value for LastModifiedDateTime to be an explicit nil
func (o *MicrosoftGraphTermStoreTerm) SetLastModifiedDateTimeNil() {
	o.LastModifiedDateTime.Set(nil)
}

// UnsetLastModifiedDateTime ensures that no value is present for LastModifiedDateTime, not even an explicit nil
func (o *MicrosoftGraphTermStoreTerm) UnsetLastModifiedDateTime() {
	o.LastModifiedDateTime.Unset()
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *MicrosoftGraphTermStoreTerm) GetProperties() []*AnyOfmicrosoftGraphKeyValue {
	if o == nil || o.Properties == nil {
		var ret []*AnyOfmicrosoftGraphKeyValue
		return ret
	}
	return *o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphTermStoreTerm) GetPropertiesOk() (*[]*AnyOfmicrosoftGraphKeyValue, bool) {
	if o == nil || o.Properties == nil {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *MicrosoftGraphTermStoreTerm) HasProperties() bool {
	if o != nil && o.Properties != nil {
		return true
	}

	return false
}

// SetProperties gets a reference to the given []*AnyOfmicrosoftGraphKeyValue and assigns it to the Properties field.
func (o *MicrosoftGraphTermStoreTerm) SetProperties(v []*AnyOfmicrosoftGraphKeyValue) {
	o.Properties = &v
}

// GetChildren returns the Children field value if set, zero value otherwise.
func (o *MicrosoftGraphTermStoreTerm) GetChildren() []MicrosoftGraphTermStoreTerm {
	if o == nil || o.Children == nil {
		var ret []MicrosoftGraphTermStoreTerm
		return ret
	}
	return *o.Children
}

// GetChildrenOk returns a tuple with the Children field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphTermStoreTerm) GetChildrenOk() (*[]MicrosoftGraphTermStoreTerm, bool) {
	if o == nil || o.Children == nil {
		return nil, false
	}
	return o.Children, true
}

// HasChildren returns a boolean if a field has been set.
func (o *MicrosoftGraphTermStoreTerm) HasChildren() bool {
	if o != nil && o.Children != nil {
		return true
	}

	return false
}

// SetChildren gets a reference to the given []MicrosoftGraphTermStoreTerm and assigns it to the Children field.
func (o *MicrosoftGraphTermStoreTerm) SetChildren(v []MicrosoftGraphTermStoreTerm) {
	o.Children = &v
}

// GetRelations returns the Relations field value if set, zero value otherwise.
func (o *MicrosoftGraphTermStoreTerm) GetRelations() []MicrosoftGraphTermStoreRelation {
	if o == nil || o.Relations == nil {
		var ret []MicrosoftGraphTermStoreRelation
		return ret
	}
	return *o.Relations
}

// GetRelationsOk returns a tuple with the Relations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphTermStoreTerm) GetRelationsOk() (*[]MicrosoftGraphTermStoreRelation, bool) {
	if o == nil || o.Relations == nil {
		return nil, false
	}
	return o.Relations, true
}

// HasRelations returns a boolean if a field has been set.
func (o *MicrosoftGraphTermStoreTerm) HasRelations() bool {
	if o != nil && o.Relations != nil {
		return true
	}

	return false
}

// SetRelations gets a reference to the given []MicrosoftGraphTermStoreRelation and assigns it to the Relations field.
func (o *MicrosoftGraphTermStoreTerm) SetRelations(v []MicrosoftGraphTermStoreRelation) {
	o.Relations = &v
}

// GetSet returns the Set field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphTermStoreTerm) GetSet() AnyOfmicrosoftGraphTermStoreSet {
	if o == nil  {
		var ret AnyOfmicrosoftGraphTermStoreSet
		return ret
	}
	return o.Set
}

// GetSetOk returns a tuple with the Set field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphTermStoreTerm) GetSetOk() (*AnyOfmicrosoftGraphTermStoreSet, bool) {
	if o == nil || o.Set == nil {
		return nil, false
	}
	return &o.Set, true
}

// HasSet returns a boolean if a field has been set.
func (o *MicrosoftGraphTermStoreTerm) HasSet() bool {
	if o != nil && o.Set != nil {
		return true
	}

	return false
}

// SetSet gets a reference to the given AnyOfmicrosoftGraphTermStoreSet and assigns it to the Set field.
func (o *MicrosoftGraphTermStoreTerm) SetSet(v AnyOfmicrosoftGraphTermStoreSet) {
	o.Set = v
}

func (o MicrosoftGraphTermStoreTerm) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
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

type NullableMicrosoftGraphTermStoreTerm struct {
	value *MicrosoftGraphTermStoreTerm
	isSet bool
}

func (v NullableMicrosoftGraphTermStoreTerm) Get() *MicrosoftGraphTermStoreTerm {
	return v.value
}

func (v *NullableMicrosoftGraphTermStoreTerm) Set(val *MicrosoftGraphTermStoreTerm) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphTermStoreTerm) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphTermStoreTerm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphTermStoreTerm(val *MicrosoftGraphTermStoreTerm) *NullableMicrosoftGraphTermStoreTerm {
	return &NullableMicrosoftGraphTermStoreTerm{value: val, isSet: true}
}

func (v NullableMicrosoftGraphTermStoreTerm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphTermStoreTerm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


