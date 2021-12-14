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

// MicrosoftGraphRecentNotebook struct for MicrosoftGraphRecentNotebook
type MicrosoftGraphRecentNotebook struct {
	// The name of the notebook.
	DisplayName NullableString `json:"displayName,omitempty"`
	// The date and time when the notebook was last modified. The timestamp represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
	LastAccessedTime NullableTime `json:"lastAccessedTime,omitempty"`
	// Links for opening the notebook. The oneNoteClientURL link opens the notebook in the OneNote client, if it's installed. The oneNoteWebURL link opens the notebook in OneNote on the web.
	Links AnyOfmicrosoftGraphRecentNotebookLinks `json:"links,omitempty"`
	// The backend store where the Notebook resides, either OneDriveForBusiness or OneDrive.
	SourceService AnyOfmicrosoftGraphOnenoteSourceService `json:"sourceService,omitempty"`
}

// NewMicrosoftGraphRecentNotebook instantiates a new MicrosoftGraphRecentNotebook object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphRecentNotebook() *MicrosoftGraphRecentNotebook {
	this := MicrosoftGraphRecentNotebook{}
	return &this
}

// NewMicrosoftGraphRecentNotebookWithDefaults instantiates a new MicrosoftGraphRecentNotebook object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphRecentNotebookWithDefaults() *MicrosoftGraphRecentNotebook {
	this := MicrosoftGraphRecentNotebook{}
	return &this
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphRecentNotebook) GetDisplayName() string {
	if o == nil || o.DisplayName.Get() == nil {
		var ret string
		return ret
	}
	return *o.DisplayName.Get()
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphRecentNotebook) GetDisplayNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DisplayName.Get(), o.DisplayName.IsSet()
}

// HasDisplayName returns a boolean if a field has been set.
func (o *MicrosoftGraphRecentNotebook) HasDisplayName() bool {
	if o != nil && o.DisplayName.IsSet() {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given NullableString and assigns it to the DisplayName field.
func (o *MicrosoftGraphRecentNotebook) SetDisplayName(v string) {
	o.DisplayName.Set(&v)
}
// SetDisplayNameNil sets the value for DisplayName to be an explicit nil
func (o *MicrosoftGraphRecentNotebook) SetDisplayNameNil() {
	o.DisplayName.Set(nil)
}

// UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
func (o *MicrosoftGraphRecentNotebook) UnsetDisplayName() {
	o.DisplayName.Unset()
}

// GetLastAccessedTime returns the LastAccessedTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphRecentNotebook) GetLastAccessedTime() time.Time {
	if o == nil || o.LastAccessedTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.LastAccessedTime.Get()
}

// GetLastAccessedTimeOk returns a tuple with the LastAccessedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphRecentNotebook) GetLastAccessedTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LastAccessedTime.Get(), o.LastAccessedTime.IsSet()
}

// HasLastAccessedTime returns a boolean if a field has been set.
func (o *MicrosoftGraphRecentNotebook) HasLastAccessedTime() bool {
	if o != nil && o.LastAccessedTime.IsSet() {
		return true
	}

	return false
}

// SetLastAccessedTime gets a reference to the given NullableTime and assigns it to the LastAccessedTime field.
func (o *MicrosoftGraphRecentNotebook) SetLastAccessedTime(v time.Time) {
	o.LastAccessedTime.Set(&v)
}
// SetLastAccessedTimeNil sets the value for LastAccessedTime to be an explicit nil
func (o *MicrosoftGraphRecentNotebook) SetLastAccessedTimeNil() {
	o.LastAccessedTime.Set(nil)
}

// UnsetLastAccessedTime ensures that no value is present for LastAccessedTime, not even an explicit nil
func (o *MicrosoftGraphRecentNotebook) UnsetLastAccessedTime() {
	o.LastAccessedTime.Unset()
}

// GetLinks returns the Links field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphRecentNotebook) GetLinks() AnyOfmicrosoftGraphRecentNotebookLinks {
	if o == nil  {
		var ret AnyOfmicrosoftGraphRecentNotebookLinks
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphRecentNotebook) GetLinksOk() (*AnyOfmicrosoftGraphRecentNotebookLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return &o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *MicrosoftGraphRecentNotebook) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AnyOfmicrosoftGraphRecentNotebookLinks and assigns it to the Links field.
func (o *MicrosoftGraphRecentNotebook) SetLinks(v AnyOfmicrosoftGraphRecentNotebookLinks) {
	o.Links = v
}

// GetSourceService returns the SourceService field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphRecentNotebook) GetSourceService() AnyOfmicrosoftGraphOnenoteSourceService {
	if o == nil  {
		var ret AnyOfmicrosoftGraphOnenoteSourceService
		return ret
	}
	return o.SourceService
}

// GetSourceServiceOk returns a tuple with the SourceService field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphRecentNotebook) GetSourceServiceOk() (*AnyOfmicrosoftGraphOnenoteSourceService, bool) {
	if o == nil || o.SourceService == nil {
		return nil, false
	}
	return &o.SourceService, true
}

// HasSourceService returns a boolean if a field has been set.
func (o *MicrosoftGraphRecentNotebook) HasSourceService() bool {
	if o != nil && o.SourceService != nil {
		return true
	}

	return false
}

// SetSourceService gets a reference to the given AnyOfmicrosoftGraphOnenoteSourceService and assigns it to the SourceService field.
func (o *MicrosoftGraphRecentNotebook) SetSourceService(v AnyOfmicrosoftGraphOnenoteSourceService) {
	o.SourceService = v
}

func (o MicrosoftGraphRecentNotebook) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DisplayName.IsSet() {
		toSerialize["displayName"] = o.DisplayName.Get()
	}
	if o.LastAccessedTime.IsSet() {
		toSerialize["lastAccessedTime"] = o.LastAccessedTime.Get()
	}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.SourceService != nil {
		toSerialize["sourceService"] = o.SourceService
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphRecentNotebook struct {
	value *MicrosoftGraphRecentNotebook
	isSet bool
}

func (v NullableMicrosoftGraphRecentNotebook) Get() *MicrosoftGraphRecentNotebook {
	return v.value
}

func (v *NullableMicrosoftGraphRecentNotebook) Set(val *MicrosoftGraphRecentNotebook) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphRecentNotebook) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphRecentNotebook) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphRecentNotebook(val *MicrosoftGraphRecentNotebook) *NullableMicrosoftGraphRecentNotebook {
	return &NullableMicrosoftGraphRecentNotebook{value: val, isSet: true}
}

func (v NullableMicrosoftGraphRecentNotebook) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphRecentNotebook) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


