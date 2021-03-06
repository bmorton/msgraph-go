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

// MicrosoftGraphOfficeGraphInsights struct for MicrosoftGraphOfficeGraphInsights
type MicrosoftGraphOfficeGraphInsights struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	// Calculated relationship identifying documents shared with or by the user. This includes URLs, file attachments, and reference attachments to OneDrive for Business and SharePoint files found in Outlook messages and meetings. This also includes URLs and reference attachments to Teams conversations. Ordered by recency of share.
	Shared *[]MicrosoftGraphSharedInsight `json:"shared,omitempty"`
	// Calculated relationship identifying documents trending around a user. Trending documents are calculated based on activity of the user's closest network of people and include files stored in OneDrive for Business and SharePoint. Trending insights help the user to discover potentially useful content that the user has access to, but has never viewed before.
	Trending *[]MicrosoftGraphTrending `json:"trending,omitempty"`
	// Calculated relationship identifying the latest documents viewed or modified by a user, including OneDrive for Business and SharePoint documents, ranked by recency of use.
	Used *[]MicrosoftGraphUsedInsight `json:"used,omitempty"`
}

// NewMicrosoftGraphOfficeGraphInsights instantiates a new MicrosoftGraphOfficeGraphInsights object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphOfficeGraphInsights() *MicrosoftGraphOfficeGraphInsights {
	this := MicrosoftGraphOfficeGraphInsights{}
	return &this
}

// NewMicrosoftGraphOfficeGraphInsightsWithDefaults instantiates a new MicrosoftGraphOfficeGraphInsights object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphOfficeGraphInsightsWithDefaults() *MicrosoftGraphOfficeGraphInsights {
	this := MicrosoftGraphOfficeGraphInsights{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MicrosoftGraphOfficeGraphInsights) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphOfficeGraphInsights) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphOfficeGraphInsights) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphOfficeGraphInsights) SetId(v string) {
	o.Id = &v
}

// GetShared returns the Shared field value if set, zero value otherwise.
func (o *MicrosoftGraphOfficeGraphInsights) GetShared() []MicrosoftGraphSharedInsight {
	if o == nil || o.Shared == nil {
		var ret []MicrosoftGraphSharedInsight
		return ret
	}
	return *o.Shared
}

// GetSharedOk returns a tuple with the Shared field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphOfficeGraphInsights) GetSharedOk() (*[]MicrosoftGraphSharedInsight, bool) {
	if o == nil || o.Shared == nil {
		return nil, false
	}
	return o.Shared, true
}

// HasShared returns a boolean if a field has been set.
func (o *MicrosoftGraphOfficeGraphInsights) HasShared() bool {
	if o != nil && o.Shared != nil {
		return true
	}

	return false
}

// SetShared gets a reference to the given []MicrosoftGraphSharedInsight and assigns it to the Shared field.
func (o *MicrosoftGraphOfficeGraphInsights) SetShared(v []MicrosoftGraphSharedInsight) {
	o.Shared = &v
}

// GetTrending returns the Trending field value if set, zero value otherwise.
func (o *MicrosoftGraphOfficeGraphInsights) GetTrending() []MicrosoftGraphTrending {
	if o == nil || o.Trending == nil {
		var ret []MicrosoftGraphTrending
		return ret
	}
	return *o.Trending
}

// GetTrendingOk returns a tuple with the Trending field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphOfficeGraphInsights) GetTrendingOk() (*[]MicrosoftGraphTrending, bool) {
	if o == nil || o.Trending == nil {
		return nil, false
	}
	return o.Trending, true
}

// HasTrending returns a boolean if a field has been set.
func (o *MicrosoftGraphOfficeGraphInsights) HasTrending() bool {
	if o != nil && o.Trending != nil {
		return true
	}

	return false
}

// SetTrending gets a reference to the given []MicrosoftGraphTrending and assigns it to the Trending field.
func (o *MicrosoftGraphOfficeGraphInsights) SetTrending(v []MicrosoftGraphTrending) {
	o.Trending = &v
}

// GetUsed returns the Used field value if set, zero value otherwise.
func (o *MicrosoftGraphOfficeGraphInsights) GetUsed() []MicrosoftGraphUsedInsight {
	if o == nil || o.Used == nil {
		var ret []MicrosoftGraphUsedInsight
		return ret
	}
	return *o.Used
}

// GetUsedOk returns a tuple with the Used field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphOfficeGraphInsights) GetUsedOk() (*[]MicrosoftGraphUsedInsight, bool) {
	if o == nil || o.Used == nil {
		return nil, false
	}
	return o.Used, true
}

// HasUsed returns a boolean if a field has been set.
func (o *MicrosoftGraphOfficeGraphInsights) HasUsed() bool {
	if o != nil && o.Used != nil {
		return true
	}

	return false
}

// SetUsed gets a reference to the given []MicrosoftGraphUsedInsight and assigns it to the Used field.
func (o *MicrosoftGraphOfficeGraphInsights) SetUsed(v []MicrosoftGraphUsedInsight) {
	o.Used = &v
}

func (o MicrosoftGraphOfficeGraphInsights) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Shared != nil {
		toSerialize["shared"] = o.Shared
	}
	if o.Trending != nil {
		toSerialize["trending"] = o.Trending
	}
	if o.Used != nil {
		toSerialize["used"] = o.Used
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphOfficeGraphInsights struct {
	value *MicrosoftGraphOfficeGraphInsights
	isSet bool
}

func (v NullableMicrosoftGraphOfficeGraphInsights) Get() *MicrosoftGraphOfficeGraphInsights {
	return v.value
}

func (v *NullableMicrosoftGraphOfficeGraphInsights) Set(val *MicrosoftGraphOfficeGraphInsights) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphOfficeGraphInsights) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphOfficeGraphInsights) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphOfficeGraphInsights(val *MicrosoftGraphOfficeGraphInsights) *NullableMicrosoftGraphOfficeGraphInsights {
	return &NullableMicrosoftGraphOfficeGraphInsights{value: val, isSet: true}
}

func (v NullableMicrosoftGraphOfficeGraphInsights) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphOfficeGraphInsights) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


