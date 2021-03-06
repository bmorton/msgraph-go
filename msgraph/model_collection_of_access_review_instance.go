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

// CollectionOfAccessReviewInstance struct for CollectionOfAccessReviewInstance
type CollectionOfAccessReviewInstance struct {
	Value *[]MicrosoftGraphAccessReviewInstance `json:"value,omitempty"`
	OdataNextLink *string `json:"@odata.nextLink,omitempty"`
}

// NewCollectionOfAccessReviewInstance instantiates a new CollectionOfAccessReviewInstance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionOfAccessReviewInstance() *CollectionOfAccessReviewInstance {
	this := CollectionOfAccessReviewInstance{}
	return &this
}

// NewCollectionOfAccessReviewInstanceWithDefaults instantiates a new CollectionOfAccessReviewInstance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionOfAccessReviewInstanceWithDefaults() *CollectionOfAccessReviewInstance {
	this := CollectionOfAccessReviewInstance{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CollectionOfAccessReviewInstance) GetValue() []MicrosoftGraphAccessReviewInstance {
	if o == nil || o.Value == nil {
		var ret []MicrosoftGraphAccessReviewInstance
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfAccessReviewInstance) GetValueOk() (*[]MicrosoftGraphAccessReviewInstance, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CollectionOfAccessReviewInstance) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given []MicrosoftGraphAccessReviewInstance and assigns it to the Value field.
func (o *CollectionOfAccessReviewInstance) SetValue(v []MicrosoftGraphAccessReviewInstance) {
	o.Value = &v
}

// GetOdataNextLink returns the OdataNextLink field value if set, zero value otherwise.
func (o *CollectionOfAccessReviewInstance) GetOdataNextLink() string {
	if o == nil || o.OdataNextLink == nil {
		var ret string
		return ret
	}
	return *o.OdataNextLink
}

// GetOdataNextLinkOk returns a tuple with the OdataNextLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfAccessReviewInstance) GetOdataNextLinkOk() (*string, bool) {
	if o == nil || o.OdataNextLink == nil {
		return nil, false
	}
	return o.OdataNextLink, true
}

// HasOdataNextLink returns a boolean if a field has been set.
func (o *CollectionOfAccessReviewInstance) HasOdataNextLink() bool {
	if o != nil && o.OdataNextLink != nil {
		return true
	}

	return false
}

// SetOdataNextLink gets a reference to the given string and assigns it to the OdataNextLink field.
func (o *CollectionOfAccessReviewInstance) SetOdataNextLink(v string) {
	o.OdataNextLink = &v
}

func (o CollectionOfAccessReviewInstance) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.OdataNextLink != nil {
		toSerialize["@odata.nextLink"] = o.OdataNextLink
	}
	return json.Marshal(toSerialize)
}

type NullableCollectionOfAccessReviewInstance struct {
	value *CollectionOfAccessReviewInstance
	isSet bool
}

func (v NullableCollectionOfAccessReviewInstance) Get() *CollectionOfAccessReviewInstance {
	return v.value
}

func (v *NullableCollectionOfAccessReviewInstance) Set(val *CollectionOfAccessReviewInstance) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionOfAccessReviewInstance) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionOfAccessReviewInstance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionOfAccessReviewInstance(val *CollectionOfAccessReviewInstance) *NullableCollectionOfAccessReviewInstance {
	return &NullableCollectionOfAccessReviewInstance{value: val, isSet: true}
}

func (v NullableCollectionOfAccessReviewInstance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionOfAccessReviewInstance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


