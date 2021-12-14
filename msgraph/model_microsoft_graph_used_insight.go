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

// MicrosoftGraphUsedInsight struct for MicrosoftGraphUsedInsight
type MicrosoftGraphUsedInsight struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	// Information about when the item was last viewed or modified by the user. Read only.
	LastUsed AnyOfmicrosoftGraphUsageDetails `json:"lastUsed,omitempty"`
	// Reference properties of the used document, such as the url and type of the document. Read-only
	ResourceReference AnyOfmicrosoftGraphResourceReference `json:"resourceReference,omitempty"`
	// Properties that you can use to visualize the document in your experience. Read-only
	ResourceVisualization AnyOfmicrosoftGraphResourceVisualization `json:"resourceVisualization,omitempty"`
	// Used for navigating to the item that was used. For file attachments, the type is fileAttachment. For linked attachments, the type is driveItem.
	Resource AnyOfmicrosoftGraphEntity `json:"resource,omitempty"`
}

// NewMicrosoftGraphUsedInsight instantiates a new MicrosoftGraphUsedInsight object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphUsedInsight() *MicrosoftGraphUsedInsight {
	this := MicrosoftGraphUsedInsight{}
	return &this
}

// NewMicrosoftGraphUsedInsightWithDefaults instantiates a new MicrosoftGraphUsedInsight object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphUsedInsightWithDefaults() *MicrosoftGraphUsedInsight {
	this := MicrosoftGraphUsedInsight{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MicrosoftGraphUsedInsight) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphUsedInsight) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphUsedInsight) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphUsedInsight) SetId(v string) {
	o.Id = &v
}

// GetLastUsed returns the LastUsed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphUsedInsight) GetLastUsed() AnyOfmicrosoftGraphUsageDetails {
	if o == nil  {
		var ret AnyOfmicrosoftGraphUsageDetails
		return ret
	}
	return o.LastUsed
}

// GetLastUsedOk returns a tuple with the LastUsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphUsedInsight) GetLastUsedOk() (*AnyOfmicrosoftGraphUsageDetails, bool) {
	if o == nil || o.LastUsed == nil {
		return nil, false
	}
	return &o.LastUsed, true
}

// HasLastUsed returns a boolean if a field has been set.
func (o *MicrosoftGraphUsedInsight) HasLastUsed() bool {
	if o != nil && o.LastUsed != nil {
		return true
	}

	return false
}

// SetLastUsed gets a reference to the given AnyOfmicrosoftGraphUsageDetails and assigns it to the LastUsed field.
func (o *MicrosoftGraphUsedInsight) SetLastUsed(v AnyOfmicrosoftGraphUsageDetails) {
	o.LastUsed = v
}

// GetResourceReference returns the ResourceReference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphUsedInsight) GetResourceReference() AnyOfmicrosoftGraphResourceReference {
	if o == nil  {
		var ret AnyOfmicrosoftGraphResourceReference
		return ret
	}
	return o.ResourceReference
}

// GetResourceReferenceOk returns a tuple with the ResourceReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphUsedInsight) GetResourceReferenceOk() (*AnyOfmicrosoftGraphResourceReference, bool) {
	if o == nil || o.ResourceReference == nil {
		return nil, false
	}
	return &o.ResourceReference, true
}

// HasResourceReference returns a boolean if a field has been set.
func (o *MicrosoftGraphUsedInsight) HasResourceReference() bool {
	if o != nil && o.ResourceReference != nil {
		return true
	}

	return false
}

// SetResourceReference gets a reference to the given AnyOfmicrosoftGraphResourceReference and assigns it to the ResourceReference field.
func (o *MicrosoftGraphUsedInsight) SetResourceReference(v AnyOfmicrosoftGraphResourceReference) {
	o.ResourceReference = v
}

// GetResourceVisualization returns the ResourceVisualization field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphUsedInsight) GetResourceVisualization() AnyOfmicrosoftGraphResourceVisualization {
	if o == nil  {
		var ret AnyOfmicrosoftGraphResourceVisualization
		return ret
	}
	return o.ResourceVisualization
}

// GetResourceVisualizationOk returns a tuple with the ResourceVisualization field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphUsedInsight) GetResourceVisualizationOk() (*AnyOfmicrosoftGraphResourceVisualization, bool) {
	if o == nil || o.ResourceVisualization == nil {
		return nil, false
	}
	return &o.ResourceVisualization, true
}

// HasResourceVisualization returns a boolean if a field has been set.
func (o *MicrosoftGraphUsedInsight) HasResourceVisualization() bool {
	if o != nil && o.ResourceVisualization != nil {
		return true
	}

	return false
}

// SetResourceVisualization gets a reference to the given AnyOfmicrosoftGraphResourceVisualization and assigns it to the ResourceVisualization field.
func (o *MicrosoftGraphUsedInsight) SetResourceVisualization(v AnyOfmicrosoftGraphResourceVisualization) {
	o.ResourceVisualization = v
}

// GetResource returns the Resource field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphUsedInsight) GetResource() AnyOfmicrosoftGraphEntity {
	if o == nil  {
		var ret AnyOfmicrosoftGraphEntity
		return ret
	}
	return o.Resource
}

// GetResourceOk returns a tuple with the Resource field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphUsedInsight) GetResourceOk() (*AnyOfmicrosoftGraphEntity, bool) {
	if o == nil || o.Resource == nil {
		return nil, false
	}
	return &o.Resource, true
}

// HasResource returns a boolean if a field has been set.
func (o *MicrosoftGraphUsedInsight) HasResource() bool {
	if o != nil && o.Resource != nil {
		return true
	}

	return false
}

// SetResource gets a reference to the given AnyOfmicrosoftGraphEntity and assigns it to the Resource field.
func (o *MicrosoftGraphUsedInsight) SetResource(v AnyOfmicrosoftGraphEntity) {
	o.Resource = v
}

func (o MicrosoftGraphUsedInsight) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
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

type NullableMicrosoftGraphUsedInsight struct {
	value *MicrosoftGraphUsedInsight
	isSet bool
}

func (v NullableMicrosoftGraphUsedInsight) Get() *MicrosoftGraphUsedInsight {
	return v.value
}

func (v *NullableMicrosoftGraphUsedInsight) Set(val *MicrosoftGraphUsedInsight) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphUsedInsight) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphUsedInsight) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphUsedInsight(val *MicrosoftGraphUsedInsight) *NullableMicrosoftGraphUsedInsight {
	return &NullableMicrosoftGraphUsedInsight{value: val, isSet: true}
}

func (v NullableMicrosoftGraphUsedInsight) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphUsedInsight) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

