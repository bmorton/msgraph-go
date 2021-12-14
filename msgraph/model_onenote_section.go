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

// OnenoteSection struct for OnenoteSection
type OnenoteSection struct {
	// Indicates whether this is the user's default section. Read-only.
	IsDefault NullableBool `json:"isDefault,omitempty"`
	// Links for opening the section. The oneNoteClientURL link opens the section in the OneNote native client if it's installed. The oneNoteWebURL link opens the section in OneNote on the web.
	Links AnyOfmicrosoftGraphSectionLinks `json:"links,omitempty"`
	// The pages endpoint where you can get details for all the pages in the section. Read-only.
	PagesUrl NullableString `json:"pagesUrl,omitempty"`
	// The collection of pages in the section.  Read-only. Nullable.
	Pages *[]MicrosoftGraphOnenotePage `json:"pages,omitempty"`
	// The notebook that contains the section.  Read-only.
	ParentNotebook AnyOfmicrosoftGraphNotebook `json:"parentNotebook,omitempty"`
	// The section group that contains the section.  Read-only.
	ParentSectionGroup AnyOfmicrosoftGraphSectionGroup `json:"parentSectionGroup,omitempty"`
}

// NewOnenoteSection instantiates a new OnenoteSection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOnenoteSection() *OnenoteSection {
	this := OnenoteSection{}
	return &this
}

// NewOnenoteSectionWithDefaults instantiates a new OnenoteSection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOnenoteSectionWithDefaults() *OnenoteSection {
	this := OnenoteSection{}
	return &this
}

// GetIsDefault returns the IsDefault field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OnenoteSection) GetIsDefault() bool {
	if o == nil || o.IsDefault.Get() == nil {
		var ret bool
		return ret
	}
	return *o.IsDefault.Get()
}

// GetIsDefaultOk returns a tuple with the IsDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OnenoteSection) GetIsDefaultOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IsDefault.Get(), o.IsDefault.IsSet()
}

// HasIsDefault returns a boolean if a field has been set.
func (o *OnenoteSection) HasIsDefault() bool {
	if o != nil && o.IsDefault.IsSet() {
		return true
	}

	return false
}

// SetIsDefault gets a reference to the given NullableBool and assigns it to the IsDefault field.
func (o *OnenoteSection) SetIsDefault(v bool) {
	o.IsDefault.Set(&v)
}
// SetIsDefaultNil sets the value for IsDefault to be an explicit nil
func (o *OnenoteSection) SetIsDefaultNil() {
	o.IsDefault.Set(nil)
}

// UnsetIsDefault ensures that no value is present for IsDefault, not even an explicit nil
func (o *OnenoteSection) UnsetIsDefault() {
	o.IsDefault.Unset()
}

// GetLinks returns the Links field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OnenoteSection) GetLinks() AnyOfmicrosoftGraphSectionLinks {
	if o == nil  {
		var ret AnyOfmicrosoftGraphSectionLinks
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OnenoteSection) GetLinksOk() (*AnyOfmicrosoftGraphSectionLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return &o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *OnenoteSection) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AnyOfmicrosoftGraphSectionLinks and assigns it to the Links field.
func (o *OnenoteSection) SetLinks(v AnyOfmicrosoftGraphSectionLinks) {
	o.Links = v
}

// GetPagesUrl returns the PagesUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OnenoteSection) GetPagesUrl() string {
	if o == nil || o.PagesUrl.Get() == nil {
		var ret string
		return ret
	}
	return *o.PagesUrl.Get()
}

// GetPagesUrlOk returns a tuple with the PagesUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OnenoteSection) GetPagesUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PagesUrl.Get(), o.PagesUrl.IsSet()
}

// HasPagesUrl returns a boolean if a field has been set.
func (o *OnenoteSection) HasPagesUrl() bool {
	if o != nil && o.PagesUrl.IsSet() {
		return true
	}

	return false
}

// SetPagesUrl gets a reference to the given NullableString and assigns it to the PagesUrl field.
func (o *OnenoteSection) SetPagesUrl(v string) {
	o.PagesUrl.Set(&v)
}
// SetPagesUrlNil sets the value for PagesUrl to be an explicit nil
func (o *OnenoteSection) SetPagesUrlNil() {
	o.PagesUrl.Set(nil)
}

// UnsetPagesUrl ensures that no value is present for PagesUrl, not even an explicit nil
func (o *OnenoteSection) UnsetPagesUrl() {
	o.PagesUrl.Unset()
}

// GetPages returns the Pages field value if set, zero value otherwise.
func (o *OnenoteSection) GetPages() []MicrosoftGraphOnenotePage {
	if o == nil || o.Pages == nil {
		var ret []MicrosoftGraphOnenotePage
		return ret
	}
	return *o.Pages
}

// GetPagesOk returns a tuple with the Pages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnenoteSection) GetPagesOk() (*[]MicrosoftGraphOnenotePage, bool) {
	if o == nil || o.Pages == nil {
		return nil, false
	}
	return o.Pages, true
}

// HasPages returns a boolean if a field has been set.
func (o *OnenoteSection) HasPages() bool {
	if o != nil && o.Pages != nil {
		return true
	}

	return false
}

// SetPages gets a reference to the given []MicrosoftGraphOnenotePage and assigns it to the Pages field.
func (o *OnenoteSection) SetPages(v []MicrosoftGraphOnenotePage) {
	o.Pages = &v
}

// GetParentNotebook returns the ParentNotebook field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OnenoteSection) GetParentNotebook() AnyOfmicrosoftGraphNotebook {
	if o == nil  {
		var ret AnyOfmicrosoftGraphNotebook
		return ret
	}
	return o.ParentNotebook
}

// GetParentNotebookOk returns a tuple with the ParentNotebook field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OnenoteSection) GetParentNotebookOk() (*AnyOfmicrosoftGraphNotebook, bool) {
	if o == nil || o.ParentNotebook == nil {
		return nil, false
	}
	return &o.ParentNotebook, true
}

// HasParentNotebook returns a boolean if a field has been set.
func (o *OnenoteSection) HasParentNotebook() bool {
	if o != nil && o.ParentNotebook != nil {
		return true
	}

	return false
}

// SetParentNotebook gets a reference to the given AnyOfmicrosoftGraphNotebook and assigns it to the ParentNotebook field.
func (o *OnenoteSection) SetParentNotebook(v AnyOfmicrosoftGraphNotebook) {
	o.ParentNotebook = v
}

// GetParentSectionGroup returns the ParentSectionGroup field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OnenoteSection) GetParentSectionGroup() AnyOfmicrosoftGraphSectionGroup {
	if o == nil  {
		var ret AnyOfmicrosoftGraphSectionGroup
		return ret
	}
	return o.ParentSectionGroup
}

// GetParentSectionGroupOk returns a tuple with the ParentSectionGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OnenoteSection) GetParentSectionGroupOk() (*AnyOfmicrosoftGraphSectionGroup, bool) {
	if o == nil || o.ParentSectionGroup == nil {
		return nil, false
	}
	return &o.ParentSectionGroup, true
}

// HasParentSectionGroup returns a boolean if a field has been set.
func (o *OnenoteSection) HasParentSectionGroup() bool {
	if o != nil && o.ParentSectionGroup != nil {
		return true
	}

	return false
}

// SetParentSectionGroup gets a reference to the given AnyOfmicrosoftGraphSectionGroup and assigns it to the ParentSectionGroup field.
func (o *OnenoteSection) SetParentSectionGroup(v AnyOfmicrosoftGraphSectionGroup) {
	o.ParentSectionGroup = v
}

func (o OnenoteSection) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IsDefault.IsSet() {
		toSerialize["isDefault"] = o.IsDefault.Get()
	}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.PagesUrl.IsSet() {
		toSerialize["pagesUrl"] = o.PagesUrl.Get()
	}
	if o.Pages != nil {
		toSerialize["pages"] = o.Pages
	}
	if o.ParentNotebook != nil {
		toSerialize["parentNotebook"] = o.ParentNotebook
	}
	if o.ParentSectionGroup != nil {
		toSerialize["parentSectionGroup"] = o.ParentSectionGroup
	}
	return json.Marshal(toSerialize)
}

type NullableOnenoteSection struct {
	value *OnenoteSection
	isSet bool
}

func (v NullableOnenoteSection) Get() *OnenoteSection {
	return v.value
}

func (v *NullableOnenoteSection) Set(val *OnenoteSection) {
	v.value = val
	v.isSet = true
}

func (v NullableOnenoteSection) IsSet() bool {
	return v.isSet
}

func (v *NullableOnenoteSection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOnenoteSection(val *OnenoteSection) *NullableOnenoteSection {
	return &NullableOnenoteSection{value: val, isSet: true}
}

func (v NullableOnenoteSection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOnenoteSection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


