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

// Trending struct for Trending
type Trending struct {
	LastModifiedDateTime NullableTime `json:"lastModifiedDateTime,omitempty"`
	// Reference properties of the trending document, such as the url and type of the document.
	ResourceReference AnyOfmicrosoftGraphResourceReference `json:"resourceReference,omitempty"`
	// Properties that you can use to visualize the document in your experience.
	ResourceVisualization AnyOfmicrosoftGraphResourceVisualization `json:"resourceVisualization,omitempty"`
	// Value indicating how much the document is currently trending. The larger the number, the more the document is currently trending around the user (the more relevant it is). Returned documents are sorted by this value.
	Weight AnyOfnumberstringstring `json:"weight,omitempty"`
	// Used for navigating to the trending document.
	Resource AnyOfmicrosoftGraphEntity `json:"resource,omitempty"`
}

// NewTrending instantiates a new Trending object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTrending() *Trending {
	this := Trending{}
	return &this
}

// NewTrendingWithDefaults instantiates a new Trending object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTrendingWithDefaults() *Trending {
	this := Trending{}
	return &this
}

// GetLastModifiedDateTime returns the LastModifiedDateTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Trending) GetLastModifiedDateTime() time.Time {
	if o == nil || o.LastModifiedDateTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.LastModifiedDateTime.Get()
}

// GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Trending) GetLastModifiedDateTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LastModifiedDateTime.Get(), o.LastModifiedDateTime.IsSet()
}

// HasLastModifiedDateTime returns a boolean if a field has been set.
func (o *Trending) HasLastModifiedDateTime() bool {
	if o != nil && o.LastModifiedDateTime.IsSet() {
		return true
	}

	return false
}

// SetLastModifiedDateTime gets a reference to the given NullableTime and assigns it to the LastModifiedDateTime field.
func (o *Trending) SetLastModifiedDateTime(v time.Time) {
	o.LastModifiedDateTime.Set(&v)
}
// SetLastModifiedDateTimeNil sets the value for LastModifiedDateTime to be an explicit nil
func (o *Trending) SetLastModifiedDateTimeNil() {
	o.LastModifiedDateTime.Set(nil)
}

// UnsetLastModifiedDateTime ensures that no value is present for LastModifiedDateTime, not even an explicit nil
func (o *Trending) UnsetLastModifiedDateTime() {
	o.LastModifiedDateTime.Unset()
}

// GetResourceReference returns the ResourceReference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Trending) GetResourceReference() AnyOfmicrosoftGraphResourceReference {
	if o == nil  {
		var ret AnyOfmicrosoftGraphResourceReference
		return ret
	}
	return o.ResourceReference
}

// GetResourceReferenceOk returns a tuple with the ResourceReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Trending) GetResourceReferenceOk() (*AnyOfmicrosoftGraphResourceReference, bool) {
	if o == nil || o.ResourceReference == nil {
		return nil, false
	}
	return &o.ResourceReference, true
}

// HasResourceReference returns a boolean if a field has been set.
func (o *Trending) HasResourceReference() bool {
	if o != nil && o.ResourceReference != nil {
		return true
	}

	return false
}

// SetResourceReference gets a reference to the given AnyOfmicrosoftGraphResourceReference and assigns it to the ResourceReference field.
func (o *Trending) SetResourceReference(v AnyOfmicrosoftGraphResourceReference) {
	o.ResourceReference = v
}

// GetResourceVisualization returns the ResourceVisualization field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Trending) GetResourceVisualization() AnyOfmicrosoftGraphResourceVisualization {
	if o == nil  {
		var ret AnyOfmicrosoftGraphResourceVisualization
		return ret
	}
	return o.ResourceVisualization
}

// GetResourceVisualizationOk returns a tuple with the ResourceVisualization field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Trending) GetResourceVisualizationOk() (*AnyOfmicrosoftGraphResourceVisualization, bool) {
	if o == nil || o.ResourceVisualization == nil {
		return nil, false
	}
	return &o.ResourceVisualization, true
}

// HasResourceVisualization returns a boolean if a field has been set.
func (o *Trending) HasResourceVisualization() bool {
	if o != nil && o.ResourceVisualization != nil {
		return true
	}

	return false
}

// SetResourceVisualization gets a reference to the given AnyOfmicrosoftGraphResourceVisualization and assigns it to the ResourceVisualization field.
func (o *Trending) SetResourceVisualization(v AnyOfmicrosoftGraphResourceVisualization) {
	o.ResourceVisualization = v
}

// GetWeight returns the Weight field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Trending) GetWeight() AnyOfnumberstringstring {
	if o == nil  {
		var ret AnyOfnumberstringstring
		return ret
	}
	return o.Weight
}

// GetWeightOk returns a tuple with the Weight field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Trending) GetWeightOk() (*AnyOfnumberstringstring, bool) {
	if o == nil || o.Weight == nil {
		return nil, false
	}
	return &o.Weight, true
}

// HasWeight returns a boolean if a field has been set.
func (o *Trending) HasWeight() bool {
	if o != nil && o.Weight != nil {
		return true
	}

	return false
}

// SetWeight gets a reference to the given AnyOfnumberstringstring and assigns it to the Weight field.
func (o *Trending) SetWeight(v AnyOfnumberstringstring) {
	o.Weight = v
}

// GetResource returns the Resource field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Trending) GetResource() AnyOfmicrosoftGraphEntity {
	if o == nil  {
		var ret AnyOfmicrosoftGraphEntity
		return ret
	}
	return o.Resource
}

// GetResourceOk returns a tuple with the Resource field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Trending) GetResourceOk() (*AnyOfmicrosoftGraphEntity, bool) {
	if o == nil || o.Resource == nil {
		return nil, false
	}
	return &o.Resource, true
}

// HasResource returns a boolean if a field has been set.
func (o *Trending) HasResource() bool {
	if o != nil && o.Resource != nil {
		return true
	}

	return false
}

// SetResource gets a reference to the given AnyOfmicrosoftGraphEntity and assigns it to the Resource field.
func (o *Trending) SetResource(v AnyOfmicrosoftGraphEntity) {
	o.Resource = v
}

func (o Trending) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.LastModifiedDateTime.IsSet() {
		toSerialize["lastModifiedDateTime"] = o.LastModifiedDateTime.Get()
	}
	if o.ResourceReference != nil {
		toSerialize["resourceReference"] = o.ResourceReference
	}
	if o.ResourceVisualization != nil {
		toSerialize["resourceVisualization"] = o.ResourceVisualization
	}
	if o.Weight != nil {
		toSerialize["weight"] = o.Weight
	}
	if o.Resource != nil {
		toSerialize["resource"] = o.Resource
	}
	return json.Marshal(toSerialize)
}

type NullableTrending struct {
	value *Trending
	isSet bool
}

func (v NullableTrending) Get() *Trending {
	return v.value
}

func (v *NullableTrending) Set(val *Trending) {
	v.value = val
	v.isSet = true
}

func (v NullableTrending) IsSet() bool {
	return v.isSet
}

func (v *NullableTrending) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrending(val *Trending) *NullableTrending {
	return &NullableTrending{value: val, isSet: true}
}

func (v NullableTrending) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTrending) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

