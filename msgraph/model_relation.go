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

// Relation struct for Relation
type Relation struct {
	// The type of relation. Possible values are: pin, reuse.
	Relationship AnyOfmicrosoftGraphTermStoreRelationType `json:"relationship,omitempty"`
	// The from [term] of the relation. The term from which the relationship is defined. A null value would indicate the relation is directly with the [set].
	FromTerm AnyOfmicrosoftGraphTermStoreTerm `json:"fromTerm,omitempty"`
	// The [set] in which the relation is relevant.
	Set AnyOfmicrosoftGraphTermStoreSet `json:"set,omitempty"`
	// The to [term] of the relation. The term to which the relationship is defined.
	ToTerm AnyOfmicrosoftGraphTermStoreTerm `json:"toTerm,omitempty"`
}

// NewRelation instantiates a new Relation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelation() *Relation {
	this := Relation{}
	return &this
}

// NewRelationWithDefaults instantiates a new Relation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelationWithDefaults() *Relation {
	this := Relation{}
	return &this
}

// GetRelationship returns the Relationship field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Relation) GetRelationship() AnyOfmicrosoftGraphTermStoreRelationType {
	if o == nil  {
		var ret AnyOfmicrosoftGraphTermStoreRelationType
		return ret
	}
	return o.Relationship
}

// GetRelationshipOk returns a tuple with the Relationship field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Relation) GetRelationshipOk() (*AnyOfmicrosoftGraphTermStoreRelationType, bool) {
	if o == nil || o.Relationship == nil {
		return nil, false
	}
	return &o.Relationship, true
}

// HasRelationship returns a boolean if a field has been set.
func (o *Relation) HasRelationship() bool {
	if o != nil && o.Relationship != nil {
		return true
	}

	return false
}

// SetRelationship gets a reference to the given AnyOfmicrosoftGraphTermStoreRelationType and assigns it to the Relationship field.
func (o *Relation) SetRelationship(v AnyOfmicrosoftGraphTermStoreRelationType) {
	o.Relationship = v
}

// GetFromTerm returns the FromTerm field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Relation) GetFromTerm() AnyOfmicrosoftGraphTermStoreTerm {
	if o == nil  {
		var ret AnyOfmicrosoftGraphTermStoreTerm
		return ret
	}
	return o.FromTerm
}

// GetFromTermOk returns a tuple with the FromTerm field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Relation) GetFromTermOk() (*AnyOfmicrosoftGraphTermStoreTerm, bool) {
	if o == nil || o.FromTerm == nil {
		return nil, false
	}
	return &o.FromTerm, true
}

// HasFromTerm returns a boolean if a field has been set.
func (o *Relation) HasFromTerm() bool {
	if o != nil && o.FromTerm != nil {
		return true
	}

	return false
}

// SetFromTerm gets a reference to the given AnyOfmicrosoftGraphTermStoreTerm and assigns it to the FromTerm field.
func (o *Relation) SetFromTerm(v AnyOfmicrosoftGraphTermStoreTerm) {
	o.FromTerm = v
}

// GetSet returns the Set field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Relation) GetSet() AnyOfmicrosoftGraphTermStoreSet {
	if o == nil  {
		var ret AnyOfmicrosoftGraphTermStoreSet
		return ret
	}
	return o.Set
}

// GetSetOk returns a tuple with the Set field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Relation) GetSetOk() (*AnyOfmicrosoftGraphTermStoreSet, bool) {
	if o == nil || o.Set == nil {
		return nil, false
	}
	return &o.Set, true
}

// HasSet returns a boolean if a field has been set.
func (o *Relation) HasSet() bool {
	if o != nil && o.Set != nil {
		return true
	}

	return false
}

// SetSet gets a reference to the given AnyOfmicrosoftGraphTermStoreSet and assigns it to the Set field.
func (o *Relation) SetSet(v AnyOfmicrosoftGraphTermStoreSet) {
	o.Set = v
}

// GetToTerm returns the ToTerm field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Relation) GetToTerm() AnyOfmicrosoftGraphTermStoreTerm {
	if o == nil  {
		var ret AnyOfmicrosoftGraphTermStoreTerm
		return ret
	}
	return o.ToTerm
}

// GetToTermOk returns a tuple with the ToTerm field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Relation) GetToTermOk() (*AnyOfmicrosoftGraphTermStoreTerm, bool) {
	if o == nil || o.ToTerm == nil {
		return nil, false
	}
	return &o.ToTerm, true
}

// HasToTerm returns a boolean if a field has been set.
func (o *Relation) HasToTerm() bool {
	if o != nil && o.ToTerm != nil {
		return true
	}

	return false
}

// SetToTerm gets a reference to the given AnyOfmicrosoftGraphTermStoreTerm and assigns it to the ToTerm field.
func (o *Relation) SetToTerm(v AnyOfmicrosoftGraphTermStoreTerm) {
	o.ToTerm = v
}

func (o Relation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Relationship != nil {
		toSerialize["relationship"] = o.Relationship
	}
	if o.FromTerm != nil {
		toSerialize["fromTerm"] = o.FromTerm
	}
	if o.Set != nil {
		toSerialize["set"] = o.Set
	}
	if o.ToTerm != nil {
		toSerialize["toTerm"] = o.ToTerm
	}
	return json.Marshal(toSerialize)
}

type NullableRelation struct {
	value *Relation
	isSet bool
}

func (v NullableRelation) Get() *Relation {
	return v.value
}

func (v *NullableRelation) Set(val *Relation) {
	v.value = val
	v.isSet = true
}

func (v NullableRelation) IsSet() bool {
	return v.isSet
}

func (v *NullableRelation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRelation(val *Relation) *NullableRelation {
	return &NullableRelation{value: val, isSet: true}
}

func (v NullableRelation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRelation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

