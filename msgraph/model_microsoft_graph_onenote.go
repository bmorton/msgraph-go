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

// MicrosoftGraphOnenote struct for MicrosoftGraphOnenote
type MicrosoftGraphOnenote struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
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

// NewMicrosoftGraphOnenote instantiates a new MicrosoftGraphOnenote object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphOnenote() *MicrosoftGraphOnenote {
	this := MicrosoftGraphOnenote{}
	return &this
}

// NewMicrosoftGraphOnenoteWithDefaults instantiates a new MicrosoftGraphOnenote object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphOnenoteWithDefaults() *MicrosoftGraphOnenote {
	this := MicrosoftGraphOnenote{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MicrosoftGraphOnenote) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphOnenote) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphOnenote) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphOnenote) SetId(v string) {
	o.Id = &v
}

// GetNotebooks returns the Notebooks field value if set, zero value otherwise.
func (o *MicrosoftGraphOnenote) GetNotebooks() []MicrosoftGraphNotebook {
	if o == nil || o.Notebooks == nil {
		var ret []MicrosoftGraphNotebook
		return ret
	}
	return *o.Notebooks
}

// GetNotebooksOk returns a tuple with the Notebooks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphOnenote) GetNotebooksOk() (*[]MicrosoftGraphNotebook, bool) {
	if o == nil || o.Notebooks == nil {
		return nil, false
	}
	return o.Notebooks, true
}

// HasNotebooks returns a boolean if a field has been set.
func (o *MicrosoftGraphOnenote) HasNotebooks() bool {
	if o != nil && o.Notebooks != nil {
		return true
	}

	return false
}

// SetNotebooks gets a reference to the given []MicrosoftGraphNotebook and assigns it to the Notebooks field.
func (o *MicrosoftGraphOnenote) SetNotebooks(v []MicrosoftGraphNotebook) {
	o.Notebooks = &v
}

// GetOperations returns the Operations field value if set, zero value otherwise.
func (o *MicrosoftGraphOnenote) GetOperations() []MicrosoftGraphOnenoteOperation {
	if o == nil || o.Operations == nil {
		var ret []MicrosoftGraphOnenoteOperation
		return ret
	}
	return *o.Operations
}

// GetOperationsOk returns a tuple with the Operations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphOnenote) GetOperationsOk() (*[]MicrosoftGraphOnenoteOperation, bool) {
	if o == nil || o.Operations == nil {
		return nil, false
	}
	return o.Operations, true
}

// HasOperations returns a boolean if a field has been set.
func (o *MicrosoftGraphOnenote) HasOperations() bool {
	if o != nil && o.Operations != nil {
		return true
	}

	return false
}

// SetOperations gets a reference to the given []MicrosoftGraphOnenoteOperation and assigns it to the Operations field.
func (o *MicrosoftGraphOnenote) SetOperations(v []MicrosoftGraphOnenoteOperation) {
	o.Operations = &v
}

// GetPages returns the Pages field value if set, zero value otherwise.
func (o *MicrosoftGraphOnenote) GetPages() []MicrosoftGraphOnenotePage {
	if o == nil || o.Pages == nil {
		var ret []MicrosoftGraphOnenotePage
		return ret
	}
	return *o.Pages
}

// GetPagesOk returns a tuple with the Pages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphOnenote) GetPagesOk() (*[]MicrosoftGraphOnenotePage, bool) {
	if o == nil || o.Pages == nil {
		return nil, false
	}
	return o.Pages, true
}

// HasPages returns a boolean if a field has been set.
func (o *MicrosoftGraphOnenote) HasPages() bool {
	if o != nil && o.Pages != nil {
		return true
	}

	return false
}

// SetPages gets a reference to the given []MicrosoftGraphOnenotePage and assigns it to the Pages field.
func (o *MicrosoftGraphOnenote) SetPages(v []MicrosoftGraphOnenotePage) {
	o.Pages = &v
}

// GetResources returns the Resources field value if set, zero value otherwise.
func (o *MicrosoftGraphOnenote) GetResources() []MicrosoftGraphOnenoteResource {
	if o == nil || o.Resources == nil {
		var ret []MicrosoftGraphOnenoteResource
		return ret
	}
	return *o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphOnenote) GetResourcesOk() (*[]MicrosoftGraphOnenoteResource, bool) {
	if o == nil || o.Resources == nil {
		return nil, false
	}
	return o.Resources, true
}

// HasResources returns a boolean if a field has been set.
func (o *MicrosoftGraphOnenote) HasResources() bool {
	if o != nil && o.Resources != nil {
		return true
	}

	return false
}

// SetResources gets a reference to the given []MicrosoftGraphOnenoteResource and assigns it to the Resources field.
func (o *MicrosoftGraphOnenote) SetResources(v []MicrosoftGraphOnenoteResource) {
	o.Resources = &v
}

// GetSectionGroups returns the SectionGroups field value if set, zero value otherwise.
func (o *MicrosoftGraphOnenote) GetSectionGroups() []MicrosoftGraphSectionGroup {
	if o == nil || o.SectionGroups == nil {
		var ret []MicrosoftGraphSectionGroup
		return ret
	}
	return *o.SectionGroups
}

// GetSectionGroupsOk returns a tuple with the SectionGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphOnenote) GetSectionGroupsOk() (*[]MicrosoftGraphSectionGroup, bool) {
	if o == nil || o.SectionGroups == nil {
		return nil, false
	}
	return o.SectionGroups, true
}

// HasSectionGroups returns a boolean if a field has been set.
func (o *MicrosoftGraphOnenote) HasSectionGroups() bool {
	if o != nil && o.SectionGroups != nil {
		return true
	}

	return false
}

// SetSectionGroups gets a reference to the given []MicrosoftGraphSectionGroup and assigns it to the SectionGroups field.
func (o *MicrosoftGraphOnenote) SetSectionGroups(v []MicrosoftGraphSectionGroup) {
	o.SectionGroups = &v
}

// GetSections returns the Sections field value if set, zero value otherwise.
func (o *MicrosoftGraphOnenote) GetSections() []MicrosoftGraphOnenoteSection {
	if o == nil || o.Sections == nil {
		var ret []MicrosoftGraphOnenoteSection
		return ret
	}
	return *o.Sections
}

// GetSectionsOk returns a tuple with the Sections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphOnenote) GetSectionsOk() (*[]MicrosoftGraphOnenoteSection, bool) {
	if o == nil || o.Sections == nil {
		return nil, false
	}
	return o.Sections, true
}

// HasSections returns a boolean if a field has been set.
func (o *MicrosoftGraphOnenote) HasSections() bool {
	if o != nil && o.Sections != nil {
		return true
	}

	return false
}

// SetSections gets a reference to the given []MicrosoftGraphOnenoteSection and assigns it to the Sections field.
func (o *MicrosoftGraphOnenote) SetSections(v []MicrosoftGraphOnenoteSection) {
	o.Sections = &v
}

func (o MicrosoftGraphOnenote) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
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

type NullableMicrosoftGraphOnenote struct {
	value *MicrosoftGraphOnenote
	isSet bool
}

func (v NullableMicrosoftGraphOnenote) Get() *MicrosoftGraphOnenote {
	return v.value
}

func (v *NullableMicrosoftGraphOnenote) Set(val *MicrosoftGraphOnenote) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphOnenote) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphOnenote) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphOnenote(val *MicrosoftGraphOnenote) *NullableMicrosoftGraphOnenote {
	return &NullableMicrosoftGraphOnenote{value: val, isSet: true}
}

func (v NullableMicrosoftGraphOnenote) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphOnenote) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


