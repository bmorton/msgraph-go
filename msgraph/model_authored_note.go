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

// AuthoredNote struct for AuthoredNote
type AuthoredNote struct {
	// Identity information about the note's author.
	Author AnyOfmicrosoftGraphIdentity `json:"author,omitempty"`
	// The content of the note.
	Content AnyOfmicrosoftGraphItemBody `json:"content,omitempty"`
	// The date and time when the entity was created. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	CreatedDateTime NullableTime `json:"createdDateTime,omitempty"`
}

// NewAuthoredNote instantiates a new AuthoredNote object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthoredNote() *AuthoredNote {
	this := AuthoredNote{}
	return &this
}

// NewAuthoredNoteWithDefaults instantiates a new AuthoredNote object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthoredNoteWithDefaults() *AuthoredNote {
	this := AuthoredNote{}
	return &this
}

// GetAuthor returns the Author field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuthoredNote) GetAuthor() AnyOfmicrosoftGraphIdentity {
	if o == nil  {
		var ret AnyOfmicrosoftGraphIdentity
		return ret
	}
	return o.Author
}

// GetAuthorOk returns a tuple with the Author field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuthoredNote) GetAuthorOk() (*AnyOfmicrosoftGraphIdentity, bool) {
	if o == nil || o.Author == nil {
		return nil, false
	}
	return &o.Author, true
}

// HasAuthor returns a boolean if a field has been set.
func (o *AuthoredNote) HasAuthor() bool {
	if o != nil && o.Author != nil {
		return true
	}

	return false
}

// SetAuthor gets a reference to the given AnyOfmicrosoftGraphIdentity and assigns it to the Author field.
func (o *AuthoredNote) SetAuthor(v AnyOfmicrosoftGraphIdentity) {
	o.Author = v
}

// GetContent returns the Content field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuthoredNote) GetContent() AnyOfmicrosoftGraphItemBody {
	if o == nil  {
		var ret AnyOfmicrosoftGraphItemBody
		return ret
	}
	return o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuthoredNote) GetContentOk() (*AnyOfmicrosoftGraphItemBody, bool) {
	if o == nil || o.Content == nil {
		return nil, false
	}
	return &o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *AuthoredNote) HasContent() bool {
	if o != nil && o.Content != nil {
		return true
	}

	return false
}

// SetContent gets a reference to the given AnyOfmicrosoftGraphItemBody and assigns it to the Content field.
func (o *AuthoredNote) SetContent(v AnyOfmicrosoftGraphItemBody) {
	o.Content = v
}

// GetCreatedDateTime returns the CreatedDateTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuthoredNote) GetCreatedDateTime() time.Time {
	if o == nil || o.CreatedDateTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedDateTime.Get()
}

// GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuthoredNote) GetCreatedDateTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CreatedDateTime.Get(), o.CreatedDateTime.IsSet()
}

// HasCreatedDateTime returns a boolean if a field has been set.
func (o *AuthoredNote) HasCreatedDateTime() bool {
	if o != nil && o.CreatedDateTime.IsSet() {
		return true
	}

	return false
}

// SetCreatedDateTime gets a reference to the given NullableTime and assigns it to the CreatedDateTime field.
func (o *AuthoredNote) SetCreatedDateTime(v time.Time) {
	o.CreatedDateTime.Set(&v)
}
// SetCreatedDateTimeNil sets the value for CreatedDateTime to be an explicit nil
func (o *AuthoredNote) SetCreatedDateTimeNil() {
	o.CreatedDateTime.Set(nil)
}

// UnsetCreatedDateTime ensures that no value is present for CreatedDateTime, not even an explicit nil
func (o *AuthoredNote) UnsetCreatedDateTime() {
	o.CreatedDateTime.Unset()
}

func (o AuthoredNote) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Author != nil {
		toSerialize["author"] = o.Author
	}
	if o.Content != nil {
		toSerialize["content"] = o.Content
	}
	if o.CreatedDateTime.IsSet() {
		toSerialize["createdDateTime"] = o.CreatedDateTime.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableAuthoredNote struct {
	value *AuthoredNote
	isSet bool
}

func (v NullableAuthoredNote) Get() *AuthoredNote {
	return v.value
}

func (v *NullableAuthoredNote) Set(val *AuthoredNote) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthoredNote) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthoredNote) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthoredNote(val *AuthoredNote) *NullableAuthoredNote {
	return &NullableAuthoredNote{value: val, isSet: true}
}

func (v NullableAuthoredNote) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthoredNote) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


