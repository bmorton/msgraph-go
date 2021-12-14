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

// MicrosoftGraphAccessReviewReviewerScope struct for MicrosoftGraphAccessReviewReviewerScope
type MicrosoftGraphAccessReviewReviewerScope struct {
	// The query specifying who will be the reviewer. See table for examples.
	Query NullableString `json:"query,omitempty"`
	// In the scenario where reviewers need to be specified dynamically, this property is used to indicate the relative source of the query. This property is only required if a relative query, for example, ./manager, is specified. Possible value: decisions.
	QueryRoot NullableString `json:"queryRoot,omitempty"`
	// The type of query. Examples include MicrosoftGraph and ARM.
	QueryType NullableString `json:"queryType,omitempty"`
}

// NewMicrosoftGraphAccessReviewReviewerScope instantiates a new MicrosoftGraphAccessReviewReviewerScope object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphAccessReviewReviewerScope() *MicrosoftGraphAccessReviewReviewerScope {
	this := MicrosoftGraphAccessReviewReviewerScope{}
	return &this
}

// NewMicrosoftGraphAccessReviewReviewerScopeWithDefaults instantiates a new MicrosoftGraphAccessReviewReviewerScope object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphAccessReviewReviewerScopeWithDefaults() *MicrosoftGraphAccessReviewReviewerScope {
	this := MicrosoftGraphAccessReviewReviewerScope{}
	return &this
}

// GetQuery returns the Query field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphAccessReviewReviewerScope) GetQuery() string {
	if o == nil || o.Query.Get() == nil {
		var ret string
		return ret
	}
	return *o.Query.Get()
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphAccessReviewReviewerScope) GetQueryOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Query.Get(), o.Query.IsSet()
}

// HasQuery returns a boolean if a field has been set.
func (o *MicrosoftGraphAccessReviewReviewerScope) HasQuery() bool {
	if o != nil && o.Query.IsSet() {
		return true
	}

	return false
}

// SetQuery gets a reference to the given NullableString and assigns it to the Query field.
func (o *MicrosoftGraphAccessReviewReviewerScope) SetQuery(v string) {
	o.Query.Set(&v)
}
// SetQueryNil sets the value for Query to be an explicit nil
func (o *MicrosoftGraphAccessReviewReviewerScope) SetQueryNil() {
	o.Query.Set(nil)
}

// UnsetQuery ensures that no value is present for Query, not even an explicit nil
func (o *MicrosoftGraphAccessReviewReviewerScope) UnsetQuery() {
	o.Query.Unset()
}

// GetQueryRoot returns the QueryRoot field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphAccessReviewReviewerScope) GetQueryRoot() string {
	if o == nil || o.QueryRoot.Get() == nil {
		var ret string
		return ret
	}
	return *o.QueryRoot.Get()
}

// GetQueryRootOk returns a tuple with the QueryRoot field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphAccessReviewReviewerScope) GetQueryRootOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.QueryRoot.Get(), o.QueryRoot.IsSet()
}

// HasQueryRoot returns a boolean if a field has been set.
func (o *MicrosoftGraphAccessReviewReviewerScope) HasQueryRoot() bool {
	if o != nil && o.QueryRoot.IsSet() {
		return true
	}

	return false
}

// SetQueryRoot gets a reference to the given NullableString and assigns it to the QueryRoot field.
func (o *MicrosoftGraphAccessReviewReviewerScope) SetQueryRoot(v string) {
	o.QueryRoot.Set(&v)
}
// SetQueryRootNil sets the value for QueryRoot to be an explicit nil
func (o *MicrosoftGraphAccessReviewReviewerScope) SetQueryRootNil() {
	o.QueryRoot.Set(nil)
}

// UnsetQueryRoot ensures that no value is present for QueryRoot, not even an explicit nil
func (o *MicrosoftGraphAccessReviewReviewerScope) UnsetQueryRoot() {
	o.QueryRoot.Unset()
}

// GetQueryType returns the QueryType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphAccessReviewReviewerScope) GetQueryType() string {
	if o == nil || o.QueryType.Get() == nil {
		var ret string
		return ret
	}
	return *o.QueryType.Get()
}

// GetQueryTypeOk returns a tuple with the QueryType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphAccessReviewReviewerScope) GetQueryTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.QueryType.Get(), o.QueryType.IsSet()
}

// HasQueryType returns a boolean if a field has been set.
func (o *MicrosoftGraphAccessReviewReviewerScope) HasQueryType() bool {
	if o != nil && o.QueryType.IsSet() {
		return true
	}

	return false
}

// SetQueryType gets a reference to the given NullableString and assigns it to the QueryType field.
func (o *MicrosoftGraphAccessReviewReviewerScope) SetQueryType(v string) {
	o.QueryType.Set(&v)
}
// SetQueryTypeNil sets the value for QueryType to be an explicit nil
func (o *MicrosoftGraphAccessReviewReviewerScope) SetQueryTypeNil() {
	o.QueryType.Set(nil)
}

// UnsetQueryType ensures that no value is present for QueryType, not even an explicit nil
func (o *MicrosoftGraphAccessReviewReviewerScope) UnsetQueryType() {
	o.QueryType.Unset()
}

func (o MicrosoftGraphAccessReviewReviewerScope) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Query.IsSet() {
		toSerialize["query"] = o.Query.Get()
	}
	if o.QueryRoot.IsSet() {
		toSerialize["queryRoot"] = o.QueryRoot.Get()
	}
	if o.QueryType.IsSet() {
		toSerialize["queryType"] = o.QueryType.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphAccessReviewReviewerScope struct {
	value *MicrosoftGraphAccessReviewReviewerScope
	isSet bool
}

func (v NullableMicrosoftGraphAccessReviewReviewerScope) Get() *MicrosoftGraphAccessReviewReviewerScope {
	return v.value
}

func (v *NullableMicrosoftGraphAccessReviewReviewerScope) Set(val *MicrosoftGraphAccessReviewReviewerScope) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphAccessReviewReviewerScope) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphAccessReviewReviewerScope) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphAccessReviewReviewerScope(val *MicrosoftGraphAccessReviewReviewerScope) *NullableMicrosoftGraphAccessReviewReviewerScope {
	return &NullableMicrosoftGraphAccessReviewReviewerScope{value: val, isSet: true}
}

func (v NullableMicrosoftGraphAccessReviewReviewerScope) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphAccessReviewReviewerScope) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


