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

// Onenote struct for Onenote
type Onenote struct {
	// The collection of OneNote notebooks that are owned by the user or group. Read-only. Nullable.
	Notebooks *[]MicrosoftGraphNotebook `json:"notebooks,omitempty"`
	// The status of OneNote operations. Getting an operations collection is not supported, but you can get the status of long-running operations if the Operation-Location header is returned in the response. Read-only. Nullable.
	Operations *[]MicrosoftGraphOnenoteOperation `json:"operations,omitempty"`
	// The pages in all OneNote notebooks that are owned by the user or group.  Read-only. Nullable.
	Pages *[]MicrosoftGraphOnenotePage `json:"pages,omitempty"`
	// The image and other file resources in OneNote pages. Getting a resources collection is not supported, but you can get the binary content of a specific resource. Read-only. Nullable.
	Resources *[]MicrosoftGraphOnenoteResource `json:"resources,omitempty"`
	// The section groups in all OneNote notebooks that are owned by the user or group.  Read-only. Nullable.
	SectionGroups *[]MicrosoftGraphSectionGroup `json:"sectionGroups,omitempty"`
	// The sections in all OneNote notebooks that are owned by the user or group.  Read-only. Nullable.
	Sections *[]MicrosoftGraphOnenoteSection `json:"sections,omitempty"`
}

// NewOnenote instantiates a new Onenote object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOnenote() *Onenote {
	this := Onenote{}
	return &this
}

// NewOnenoteWithDefaults instantiates a new Onenote object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOnenoteWithDefaults() *Onenote {
	this := Onenote{}
	return &this
}

// GetNotebooks returns the Notebooks field value if set, zero value otherwise.
func (o *Onenote) GetNotebooks() []MicrosoftGraphNotebook {
	if o == nil || o.Notebooks == nil {
		var ret []MicrosoftGraphNotebook
		return ret
	}
	return *o.Notebooks
}

// GetNotebooksOk returns a tuple with the Notebooks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Onenote) GetNotebooksOk() (*[]MicrosoftGraphNotebook, bool) {
	if o == nil || o.Notebooks == nil {
		return nil, false
	}
	return o.Notebooks, true
}

// HasNotebooks returns a boolean if a field has been set.
func (o *Onenote) HasNotebooks() bool {
	if o != nil && o.Notebooks != nil {
		return true
	}

	return false
}

// SetNotebooks gets a reference to the given []MicrosoftGraphNotebook and assigns it to the Notebooks field.
func (o *Onenote) SetNotebooks(v []MicrosoftGraphNotebook) {
	o.Notebooks = &v
}

// GetOperations returns the Operations field value if set, zero value otherwise.
func (o *Onenote) GetOperations() []MicrosoftGraphOnenoteOperation {
	if o == nil || o.Operations == nil {
		var ret []MicrosoftGraphOnenoteOperation
		return ret
	}
	return *o.Operations
}

// GetOperationsOk returns a tuple with the Operations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Onenote) GetOperationsOk() (*[]MicrosoftGraphOnenoteOperation, bool) {
	if o == nil || o.Operations == nil {
		return nil, false
	}
	return o.Operations, true
}

// HasOperations returns a boolean if a field has been set.
func (o *Onenote) HasOperations() bool {
	if o != nil && o.Operations != nil {
		return true
	}

	return false
}

// SetOperations gets a reference to the given []MicrosoftGraphOnenoteOperation and assigns it to the Operations field.
func (o *Onenote) SetOperations(v []MicrosoftGraphOnenoteOperation) {
	o.Operations = &v
}

// GetPages returns the Pages field value if set, zero value otherwise.
func (o *Onenote) GetPages() []MicrosoftGraphOnenotePage {
	if o == nil || o.Pages == nil {
		var ret []MicrosoftGraphOnenotePage
		return ret
	}
	return *o.Pages
}

// GetPagesOk returns a tuple with the Pages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Onenote) GetPagesOk() (*[]MicrosoftGraphOnenotePage, bool) {
	if o == nil || o.Pages == nil {
		return nil, false
	}
	return o.Pages, true
}

// HasPages returns a boolean if a field has been set.
func (o *Onenote) HasPages() bool {
	if o != nil && o.Pages != nil {
		return true
	}

	return false
}

// SetPages gets a reference to the given []MicrosoftGraphOnenotePage and assigns it to the Pages field.
func (o *Onenote) SetPages(v []MicrosoftGraphOnenotePage) {
	o.Pages = &v
}

// GetResources returns the Resources field value if set, zero value otherwise.
func (o *Onenote) GetResources() []MicrosoftGraphOnenoteResource {
	if o == nil || o.Resources == nil {
		var ret []MicrosoftGraphOnenoteResource
		return ret
	}
	return *o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Onenote) GetResourcesOk() (*[]MicrosoftGraphOnenoteResource, bool) {
	if o == nil || o.Resources == nil {
		return nil, false
	}
	return o.Resources, true
}

// HasResources returns a boolean if a field has been set.
func (o *Onenote) HasResources() bool {
	if o != nil && o.Resources != nil {
		return true
	}

	return false
}

// SetResources gets a reference to the given []MicrosoftGraphOnenoteResource and assigns it to the Resources field.
func (o *Onenote) SetResources(v []MicrosoftGraphOnenoteResource) {
	o.Resources = &v
}

// GetSectionGroups returns the SectionGroups field value if set, zero value otherwise.
func (o *Onenote) GetSectionGroups() []MicrosoftGraphSectionGroup {
	if o == nil || o.SectionGroups == nil {
		var ret []MicrosoftGraphSectionGroup
		return ret
	}
	return *o.SectionGroups
}

// GetSectionGroupsOk returns a tuple with the SectionGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Onenote) GetSectionGroupsOk() (*[]MicrosoftGraphSectionGroup, bool) {
	if o == nil || o.SectionGroups == nil {
		return nil, false
	}
	return o.SectionGroups, true
}

// HasSectionGroups returns a boolean if a field has been set.
func (o *Onenote) HasSectionGroups() bool {
	if o != nil && o.SectionGroups != nil {
		return true
	}

	return false
}

// SetSectionGroups gets a reference to the given []MicrosoftGraphSectionGroup and assigns it to the SectionGroups field.
func (o *Onenote) SetSectionGroups(v []MicrosoftGraphSectionGroup) {
	o.SectionGroups = &v
}

// GetSections returns the Sections field value if set, zero value otherwise.
func (o *Onenote) GetSections() []MicrosoftGraphOnenoteSection {
	if o == nil || o.Sections == nil {
		var ret []MicrosoftGraphOnenoteSection
		return ret
	}
	return *o.Sections
}

// GetSectionsOk returns a tuple with the Sections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Onenote) GetSectionsOk() (*[]MicrosoftGraphOnenoteSection, bool) {
	if o == nil || o.Sections == nil {
		return nil, false
	}
	return o.Sections, true
}

// HasSections returns a boolean if a field has been set.
func (o *Onenote) HasSections() bool {
	if o != nil && o.Sections != nil {
		return true
	}

	return false
}

// SetSections gets a reference to the given []MicrosoftGraphOnenoteSection and assigns it to the Sections field.
func (o *Onenote) SetSections(v []MicrosoftGraphOnenoteSection) {
	o.Sections = &v
}

func (o Onenote) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Notebooks != nil {
		toSerialize["notebooks"] = o.Notebooks
	}
	if o.Operations != nil {
		toSerialize["operations"] = o.Operations
	}
	if o.Pages != nil {
		toSerialize["pages"] = o.Pages
	}
	if o.Resources != nil {
		toSerialize["resources"] = o.Resources
	}
	if o.SectionGroups != nil {
		toSerialize["sectionGroups"] = o.SectionGroups
	}
	if o.Sections != nil {
		toSerialize["sections"] = o.Sections
	}
	return json.Marshal(toSerialize)
}

type NullableOnenote struct {
	value *Onenote
	isSet bool
}

func (v NullableOnenote) Get() *Onenote {
	return v.value
}

func (v *NullableOnenote) Set(val *Onenote) {
	v.value = val
	v.isSet = true
}

func (v NullableOnenote) IsSet() bool {
	return v.isSet
}

func (v *NullableOnenote) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOnenote(val *Onenote) *NullableOnenote {
	return &NullableOnenote{value: val, isSet: true}
}

func (v NullableOnenote) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOnenote) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

