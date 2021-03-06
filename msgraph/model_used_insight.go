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

// UsedInsight struct for UsedInsight
type UsedInsight struct {
	// Information about when the item was last viewed or modified by the user. Read only.
	LastUsed AnyOfmicrosoftGraphUsageDetails `json:"lastUsed,omitempty"`
	// Reference properties of the used document, such as the url and type of the document. Read-only
	ResourceReference AnyOfmicrosoftGraphResourceReference `json:"resourceReference,omitempty"`
	// Properties that you can use to visualize the document in your experience. Read-only
	ResourceVisualization AnyOfmicrosoftGraphResourceVisualization `json:"resourceVisualization,omitempty"`
	// Used for navigating to the item that was used. For file attachments, the type is fileAttachment. For linked attachments, the type is driveItem.
	Resource AnyOfmicrosoftGraphEntity `json:"resource,omitempty"`
}

// NewUsedInsight instantiates a new UsedInsight object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsedInsight() *UsedInsight {
	this := UsedInsight{}
	return &this
}

// NewUsedInsightWithDefaults instantiates a new UsedInsight object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsedInsightWithDefaults() *UsedInsight {
	this := UsedInsight{}
	return &this
}

// GetLastUsed returns the LastUsed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsedInsight) GetLastUsed() AnyOfmicrosoftGraphUsageDetails {
	if o == nil  {
		var ret AnyOfmicrosoftGraphUsageDetails
		return ret
	}
	return o.LastUsed
}

// GetLastUsedOk returns a tuple with the LastUsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UsedInsight) GetLastUsedOk() (*AnyOfmicrosoftGraphUsageDetails, bool) {
	if o == nil || o.LastUsed == nil {
		return nil, false
	}
	return &o.LastUsed, true
}

// HasLastUsed returns a boolean if a field has been set.
func (o *UsedInsight) HasLastUsed() bool {
	if o != nil && o.LastUsed != nil {
		return true
	}

	return false
}

// SetLastUsed gets a reference to the given AnyOfmicrosoftGraphUsageDetails and assigns it to the LastUsed field.
func (o *UsedInsight) SetLastUsed(v AnyOfmicrosoftGraphUsageDetails) {
	o.LastUsed = v
}

// GetResourceReference returns the ResourceReference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsedInsight) GetResourceReference() AnyOfmicrosoftGraphResourceReference {
	if o == nil  {
		var ret AnyOfmicrosoftGraphResourceReference
		return ret
	}
	return o.ResourceReference
}

// GetResourceReferenceOk returns a tuple with the ResourceReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UsedInsight) GetResourceReferenceOk() (*AnyOfmicrosoftGraphResourceReference, bool) {
	if o == nil || o.ResourceReference == nil {
		return nil, false
	}
	return &o.ResourceReference, true
}

// HasResourceReference returns a boolean if a field has been set.
func (o *UsedInsight) HasResourceReference() bool {
	if o != nil && o.ResourceReference != nil {
		return true
	}

	return false
}

// SetResourceReference gets a reference to the given AnyOfmicrosoftGraphResourceReference and assigns it to the ResourceReference field.
func (o *UsedInsight) SetResourceReference(v AnyOfmicrosoftGraphResourceReference) {
	o.ResourceReference = v
}

// GetResourceVisualization returns the ResourceVisualization field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsedInsight) GetResourceVisualization() AnyOfmicrosoftGraphResourceVisualization {
	if o == nil  {
		var ret AnyOfmicrosoftGraphResourceVisualization
		return ret
	}
	return o.ResourceVisualization
}

// GetResourceVisualizationOk returns a tuple with the ResourceVisualization field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UsedInsight) GetResourceVisualizationOk() (*AnyOfmicrosoftGraphResourceVisualization, bool) {
	if o == nil || o.ResourceVisualization == nil {
		return nil, false
	}
	return &o.ResourceVisualization, true
}

// HasResourceVisualization returns a boolean if a field has been set.
func (o *UsedInsight) HasResourceVisualization() bool {
	if o != nil && o.ResourceVisualization != nil {
		return true
	}

	return false
}

// SetResourceVisualization gets a reference to the given AnyOfmicrosoftGraphResourceVisualization and assigns it to the ResourceVisualization field.
func (o *UsedInsight) SetResourceVisualization(v AnyOfmicrosoftGraphResourceVisualization) {
	o.ResourceVisualization = v
}

// GetResource returns the Resource field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsedInsight) GetResource() AnyOfmicrosoftGraphEntity {
	if o == nil  {
		var ret AnyOfmicrosoftGraphEntity
		return ret
	}
	return o.Resource
}

// GetResourceOk returns a tuple with the Resource field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UsedInsight) GetResourceOk() (*AnyOfmicrosoftGraphEntity, bool) {
	if o == nil || o.Resource == nil {
		return nil, false
	}
	return &o.Resource, true
}

// HasResource returns a boolean if a field has been set.
func (o *UsedInsight) HasResource() bool {
	if o != nil && o.Resource != nil {
		return true
	}

	return false
}

// SetResource gets a reference to the given AnyOfmicrosoftGraphEntity and assigns it to the Resource field.
func (o *UsedInsight) SetResource(v AnyOfmicrosoftGraphEntity) {
	o.Resource = v
}

func (o UsedInsight) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.LastUsed != nil {
		toSerialize["lastUsed"] = o.LastUsed
	}
	if o.ResourceReference != nil {
		toSerialize["resourceReference"] = o.ResourceReference
	}
	if o.ResourceVisualization != nil {
		toSerialize["resourceVisualization"] = o.ResourceVisualization
	}
	if o.Resource != nil {
		toSerialize["resource"] = o.Resource
	}
	return json.Marshal(toSerialize)
}

type NullableUsedInsight struct {
	value *UsedInsight
	isSet bool
}

func (v NullableUsedInsight) Get() *UsedInsight {
	return v.value
}

func (v *NullableUsedInsight) Set(val *UsedInsight) {
	v.value = val
	v.isSet = true
}

func (v NullableUsedInsight) IsSet() bool {
	return v.isSet
}

func (v *NullableUsedInsight) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsedInsight(val *UsedInsight) *NullableUsedInsight {
	return &NullableUsedInsight{value: val, isSet: true}
}

func (v NullableUsedInsight) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsedInsight) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


