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

// MicrosoftGraphAuthoredNote struct for MicrosoftGraphAuthoredNote
type MicrosoftGraphAuthoredNote struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	// Identity information about the note's author.
	Author AnyOfmicrosoftGraphIdentity `json:"author,omitempty"`
	// The content of the note.
	Content AnyOfmicrosoftGraphItemBody `json:"content,omitempty"`
	// The date and time when the entity was created. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	CreatedDateTime NullableTime `json:"createdDateTime,omitempty"`
}

// NewMicrosoftGraphAuthoredNote instantiates a new MicrosoftGraphAuthoredNote object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphAuthoredNote() *MicrosoftGraphAuthoredNote {
	this := MicrosoftGraphAuthoredNote{}
	return &this
}

// NewMicrosoftGraphAuthoredNoteWithDefaults instantiates a new MicrosoftGraphAuthoredNote object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphAuthoredNoteWithDefaults() *MicrosoftGraphAuthoredNote {
	this := MicrosoftGraphAuthoredNote{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MicrosoftGraphAuthoredNote) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphAuthoredNote) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphAuthoredNote) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphAuthoredNote) SetId(v string) {
	o.Id = &v
}

// GetAuthor returns the Author field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphAuthoredNote) GetAuthor() AnyOfmicrosoftGraphIdentity {
	if o == nil  {
		var ret AnyOfmicrosoftGraphIdentity
		return ret
	}
	return o.Author
}

// GetAuthorOk returns a tuple with the Author field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphAuthoredNote) GetAuthorOk() (*AnyOfmicrosoftGraphIdentity, bool) {
	if o == nil || o.Author == nil {
		return nil, false
	}
	return &o.Author, true
}

// HasAuthor returns a boolean if a field has been set.
func (o *MicrosoftGraphAuthoredNote) HasAuthor() bool {
	if o != nil && o.Author != nil {
		return true
	}

	return false
}

// SetAuthor gets a reference to the given AnyOfmicrosoftGraphIdentity and assigns it to the Author field.
func (o *MicrosoftGraphAuthoredNote) SetAuthor(v AnyOfmicrosoftGraphIdentity) {
	o.Author = v
}

// GetContent returns the Content field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphAuthoredNote) GetContent() AnyOfmicrosoftGraphItemBody {
	if o == nil  {
		var ret AnyOfmicrosoftGraphItemBody
		return ret
	}
	return o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphAuthoredNote) GetContentOk() (*AnyOfmicrosoftGraphItemBody, bool) {
	if o == nil || o.Content == nil {
		return nil, false
	}
	return &o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *MicrosoftGraphAuthoredNote) HasContent() bool {
	if o != nil && o.Content != nil {
		return true
	}

	return false
}

// SetContent gets a reference to the given AnyOfmicrosoftGraphItemBody and assigns it to the Content field.
func (o *MicrosoftGraphAuthoredNote) SetContent(v AnyOfmicrosoftGraphItemBody) {
	o.Content = v
}

// GetCreatedDateTime returns the CreatedDateTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphAuthoredNote) GetCreatedDateTime() time.Time {
	if o == nil || o.CreatedDateTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedDateTime.Get()
}

// GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphAuthoredNote) GetCreatedDateTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CreatedDateTime.Get(), o.CreatedDateTime.IsSet()
}

// HasCreatedDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphAuthoredNote) HasCreatedDateTime() bool {
	if o != nil && o.CreatedDateTime.IsSet() {
		return true
	}

	return false
}

// SetCreatedDateTime gets a reference to the given NullableTime and assigns it to the CreatedDateTime field.
func (o *MicrosoftGraphAuthoredNote) SetCreatedDateTime(v time.Time) {
	o.CreatedDateTime.Set(&v)
}
// SetCreatedDateTimeNil sets the value for CreatedDateTime to be an explicit nil
func (o *MicrosoftGraphAuthoredNote) SetCreatedDateTimeNil() {
	o.CreatedDateTime.Set(nil)
}

// UnsetCreatedDateTime ensures that no value is present for CreatedDateTime, not even an explicit nil
func (o *MicrosoftGraphAuthoredNote) UnsetCreatedDateTime() {
	o.CreatedDateTime.Unset()
}

func (o MicrosoftGraphAuthoredNote) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
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

type NullableMicrosoftGraphAuthoredNote struct {
	value *MicrosoftGraphAuthoredNote
	isSet bool
}

func (v NullableMicrosoftGraphAuthoredNote) Get() *MicrosoftGraphAuthoredNote {
	return v.value
}

func (v *NullableMicrosoftGraphAuthoredNote) Set(val *MicrosoftGraphAuthoredNote) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphAuthoredNote) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphAuthoredNote) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphAuthoredNote(val *MicrosoftGraphAuthoredNote) *NullableMicrosoftGraphAuthoredNote {
	return &NullableMicrosoftGraphAuthoredNote{value: val, isSet: true}
}

func (v NullableMicrosoftGraphAuthoredNote) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphAuthoredNote) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


