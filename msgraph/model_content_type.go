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

// ContentType struct for ContentType
type ContentType struct {
	// List of canonical URLs for hub sites with which this content type is associated to. This will contain all hub sites where this content type is queued to be enforced or is already enforced. Enforcing a content type means that the content type will be applied to the lists in the enforced sites.
	AssociatedHubsUrls *[]*string `json:"associatedHubsUrls,omitempty"`
	// The descriptive text for the item.
	Description NullableString `json:"description,omitempty"`
	// Document Set metadata.
	DocumentSet AnyOfmicrosoftGraphDocumentSet `json:"documentSet,omitempty"`
	// Document template metadata. To make sure that documents have consistent content across a site and its subsites, you can associate a Word, Excel, or PowerPoint template with a site content type.
	DocumentTemplate AnyOfmicrosoftGraphDocumentSetContent `json:"documentTemplate,omitempty"`
	// The name of the group this content type belongs to. Helps organize related content types.
	Group NullableString `json:"group,omitempty"`
	// Indicates whether the content type is hidden in the list's 'New' menu.
	Hidden NullableBool `json:"hidden,omitempty"`
	// If this content type is inherited from another scope (like a site), provides a reference to the item where the content type is defined.
	InheritedFrom AnyOfmicrosoftGraphItemReference `json:"inheritedFrom,omitempty"`
	// Specifies if a content type is a built-in content type.
	IsBuiltIn NullableBool `json:"isBuiltIn,omitempty"`
	// The name of the content type.
	Name NullableString `json:"name,omitempty"`
	// Specifies the order in which the content type appears in the selection UI.
	Order AnyOfmicrosoftGraphContentTypeOrder `json:"order,omitempty"`
	// The unique identifier of the content type.
	ParentId NullableString `json:"parentId,omitempty"`
	// If true, any changes made to the content type will be pushed to inherited content types and lists that implement the content type.
	PropagateChanges NullableBool `json:"propagateChanges,omitempty"`
	// If true, the content type can't be modified unless this value is first set to false.
	ReadOnly NullableBool `json:"readOnly,omitempty"`
	// If true, the content type can't be modified by users or through push-down operations. Only site collection administrators can seal or unseal content types.
	Sealed NullableBool `json:"sealed,omitempty"`
	// Parent contentType from which this content type is derived.
	Base AnyOfmicrosoftGraphContentType `json:"base,omitempty"`
	// The collection of content types that are ancestors of this content type.
	BaseTypes *[]MicrosoftGraphContentType `json:"baseTypes,omitempty"`
	// The collection of columns that are required by this content type.
	ColumnLinks *[]MicrosoftGraphColumnLink `json:"columnLinks,omitempty"`
	// Column order information in a content type.
	ColumnPositions *[]MicrosoftGraphColumnDefinition `json:"columnPositions,omitempty"`
	// The collection of column definitions for this contentType.
	Columns *[]MicrosoftGraphColumnDefinition `json:"columns,omitempty"`
}

// NewContentType instantiates a new ContentType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContentType() *ContentType {
	this := ContentType{}
	return &this
}

// NewContentTypeWithDefaults instantiates a new ContentType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContentTypeWithDefaults() *ContentType {
	this := ContentType{}
	return &this
}

// GetAssociatedHubsUrls returns the AssociatedHubsUrls field value if set, zero value otherwise.
func (o *ContentType) GetAssociatedHubsUrls() []*string {
	if o == nil || o.AssociatedHubsUrls == nil {
		var ret []*string
		return ret
	}
	return *o.AssociatedHubsUrls
}

// GetAssociatedHubsUrlsOk returns a tuple with the AssociatedHubsUrls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentType) GetAssociatedHubsUrlsOk() (*[]*string, bool) {
	if o == nil || o.AssociatedHubsUrls == nil {
		return nil, false
	}
	return o.AssociatedHubsUrls, true
}

// HasAssociatedHubsUrls returns a boolean if a field has been set.
func (o *ContentType) HasAssociatedHubsUrls() bool {
	if o != nil && o.AssociatedHubsUrls != nil {
		return true
	}

	return false
}

// SetAssociatedHubsUrls gets a reference to the given []*string and assigns it to the AssociatedHubsUrls field.
func (o *ContentType) SetAssociatedHubsUrls(v []*string) {
	o.AssociatedHubsUrls = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ContentType) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContentType) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *ContentType) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *ContentType) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *ContentType) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *ContentType) UnsetDescription() {
	o.Description.Unset()
}

// GetDocumentSet returns the DocumentSet field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ContentType) GetDocumentSet() AnyOfmicrosoftGraphDocumentSet {
	if o == nil  {
		var ret AnyOfmicrosoftGraphDocumentSet
		return ret
	}
	return o.DocumentSet
}

// GetDocumentSetOk returns a tuple with the DocumentSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContentType) GetDocumentSetOk() (*AnyOfmicrosoftGraphDocumentSet, bool) {
	if o == nil || o.DocumentSet == nil {
		return nil, false
	}
	return &o.DocumentSet, true
}

// HasDocumentSet returns a boolean if a field has been set.
func (o *ContentType) HasDocumentSet() bool {
	if o != nil && o.DocumentSet != nil {
		return true
	}

	return false
}

// SetDocumentSet gets a reference to the given AnyOfmicrosoftGraphDocumentSet and assigns it to the DocumentSet field.
func (o *ContentType) SetDocumentSet(v AnyOfmicrosoftGraphDocumentSet) {
	o.DocumentSet = v
}

// GetDocumentTemplate returns the DocumentTemplate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ContentType) GetDocumentTemplate() AnyOfmicrosoftGraphDocumentSetContent {
	if o == nil  {
		var ret AnyOfmicrosoftGraphDocumentSetContent
		return ret
	}
	return o.DocumentTemplate
}

// GetDocumentTemplateOk returns a tuple with the DocumentTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContentType) GetDocumentTemplateOk() (*AnyOfmicrosoftGraphDocumentSetContent, bool) {
	if o == nil || o.DocumentTemplate == nil {
		return nil, false
	}
	return &o.DocumentTemplate, true
}

// HasDocumentTemplate returns a boolean if a field has been set.
func (o *ContentType) HasDocumentTemplate() bool {
	if o != nil && o.DocumentTemplate != nil {
		return true
	}

	return false
}

// SetDocumentTemplate gets a reference to the given AnyOfmicrosoftGraphDocumentSetContent and assigns it to the DocumentTemplate field.
func (o *ContentType) SetDocumentTemplate(v AnyOfmicrosoftGraphDocumentSetContent) {
	o.DocumentTemplate = v
}

// GetGroup returns the Group field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ContentType) GetGroup() string {
	if o == nil || o.Group.Get() == nil {
		var ret string
		return ret
	}
	return *o.Group.Get()
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContentType) GetGroupOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Group.Get(), o.Group.IsSet()
}

// HasGroup returns a boolean if a field has been set.
func (o *ContentType) HasGroup() bool {
	if o != nil && o.Group.IsSet() {
		return true
	}

	return false
}

// SetGroup gets a reference to the given NullableString and assigns it to the Group field.
func (o *ContentType) SetGroup(v string) {
	o.Group.Set(&v)
}
// SetGroupNil sets the value for Group to be an explicit nil
func (o *ContentType) SetGroupNil() {
	o.Group.Set(nil)
}

// UnsetGroup ensures that no value is present for Group, not even an explicit nil
func (o *ContentType) UnsetGroup() {
	o.Group.Unset()
}

// GetHidden returns the Hidden field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ContentType) GetHidden() bool {
	if o == nil || o.Hidden.Get() == nil {
		var ret bool
		return ret
	}
	return *o.Hidden.Get()
}

// GetHiddenOk returns a tuple with the Hidden field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContentType) GetHiddenOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Hidden.Get(), o.Hidden.IsSet()
}

// HasHidden returns a boolean if a field has been set.
func (o *ContentType) HasHidden() bool {
	if o != nil && o.Hidden.IsSet() {
		return true
	}

	return false
}

// SetHidden gets a reference to the given NullableBool and assigns it to the Hidden field.
func (o *ContentType) SetHidden(v bool) {
	o.Hidden.Set(&v)
}
// SetHiddenNil sets the value for Hidden to be an explicit nil
func (o *ContentType) SetHiddenNil() {
	o.Hidden.Set(nil)
}

// UnsetHidden ensures that no value is present for Hidden, not even an explicit nil
func (o *ContentType) UnsetHidden() {
	o.Hidden.Unset()
}

// GetInheritedFrom returns the InheritedFrom field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ContentType) GetInheritedFrom() AnyOfmicrosoftGraphItemReference {
	if o == nil  {
		var ret AnyOfmicrosoftGraphItemReference
		return ret
	}
	return o.InheritedFrom
}

// GetInheritedFromOk returns a tuple with the InheritedFrom field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContentType) GetInheritedFromOk() (*AnyOfmicrosoftGraphItemReference, bool) {
	if o == nil || o.InheritedFrom == nil {
		return nil, false
	}
	return &o.InheritedFrom, true
}

// HasInheritedFrom returns a boolean if a field has been set.
func (o *ContentType) HasInheritedFrom() bool {
	if o != nil && o.InheritedFrom != nil {
		return true
	}

	return false
}

// SetInheritedFrom gets a reference to the given AnyOfmicrosoftGraphItemReference and assigns it to the InheritedFrom field.
func (o *ContentType) SetInheritedFrom(v AnyOfmicrosoftGraphItemReference) {
	o.InheritedFrom = v
}

// GetIsBuiltIn returns the IsBuiltIn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ContentType) GetIsBuiltIn() bool {
	if o == nil || o.IsBuiltIn.Get() == nil {
		var ret bool
		return ret
	}
	return *o.IsBuiltIn.Get()
}

// GetIsBuiltInOk returns a tuple with the IsBuiltIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContentType) GetIsBuiltInOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IsBuiltIn.Get(), o.IsBuiltIn.IsSet()
}

// HasIsBuiltIn returns a boolean if a field has been set.
func (o *ContentType) HasIsBuiltIn() bool {
	if o != nil && o.IsBuiltIn.IsSet() {
		return true
	}

	return false
}

// SetIsBuiltIn gets a reference to the given NullableBool and assigns it to the IsBuiltIn field.
func (o *ContentType) SetIsBuiltIn(v bool) {
	o.IsBuiltIn.Set(&v)
}
// SetIsBuiltInNil sets the value for IsBuiltIn to be an explicit nil
func (o *ContentType) SetIsBuiltInNil() {
	o.IsBuiltIn.Set(nil)
}

// UnsetIsBuiltIn ensures that no value is present for IsBuiltIn, not even an explicit nil
func (o *ContentType) UnsetIsBuiltIn() {
	o.IsBuiltIn.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ContentType) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContentType) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *ContentType) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *ContentType) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *ContentType) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *ContentType) UnsetName() {
	o.Name.Unset()
}

// GetOrder returns the Order field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ContentType) GetOrder() AnyOfmicrosoftGraphContentTypeOrder {
	if o == nil  {
		var ret AnyOfmicrosoftGraphContentTypeOrder
		return ret
	}
	return o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContentType) GetOrderOk() (*AnyOfmicrosoftGraphContentTypeOrder, bool) {
	if o == nil || o.Order == nil {
		return nil, false
	}
	return &o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *ContentType) HasOrder() bool {
	if o != nil && o.Order != nil {
		return true
	}

	return false
}

// SetOrder gets a reference to the given AnyOfmicrosoftGraphContentTypeOrder and assigns it to the Order field.
func (o *ContentType) SetOrder(v AnyOfmicrosoftGraphContentTypeOrder) {
	o.Order = v
}

// GetParentId returns the ParentId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ContentType) GetParentId() string {
	if o == nil || o.ParentId.Get() == nil {
		var ret string
		return ret
	}
	return *o.ParentId.Get()
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContentType) GetParentIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ParentId.Get(), o.ParentId.IsSet()
}

// HasParentId returns a boolean if a field has been set.
func (o *ContentType) HasParentId() bool {
	if o != nil && o.ParentId.IsSet() {
		return true
	}

	return false
}

// SetParentId gets a reference to the given NullableString and assigns it to the ParentId field.
func (o *ContentType) SetParentId(v string) {
	o.ParentId.Set(&v)
}
// SetParentIdNil sets the value for ParentId to be an explicit nil
func (o *ContentType) SetParentIdNil() {
	o.ParentId.Set(nil)
}

// UnsetParentId ensures that no value is present for ParentId, not even an explicit nil
func (o *ContentType) UnsetParentId() {
	o.ParentId.Unset()
}

// GetPropagateChanges returns the PropagateChanges field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ContentType) GetPropagateChanges() bool {
	if o == nil || o.PropagateChanges.Get() == nil {
		var ret bool
		return ret
	}
	return *o.PropagateChanges.Get()
}

// GetPropagateChangesOk returns a tuple with the PropagateChanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContentType) GetPropagateChangesOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PropagateChanges.Get(), o.PropagateChanges.IsSet()
}

// HasPropagateChanges returns a boolean if a field has been set.
func (o *ContentType) HasPropagateChanges() bool {
	if o != nil && o.PropagateChanges.IsSet() {
		return true
	}

	return false
}

// SetPropagateChanges gets a reference to the given NullableBool and assigns it to the PropagateChanges field.
func (o *ContentType) SetPropagateChanges(v bool) {
	o.PropagateChanges.Set(&v)
}
// SetPropagateChangesNil sets the value for PropagateChanges to be an explicit nil
func (o *ContentType) SetPropagateChangesNil() {
	o.PropagateChanges.Set(nil)
}

// UnsetPropagateChanges ensures that no value is present for PropagateChanges, not even an explicit nil
func (o *ContentType) UnsetPropagateChanges() {
	o.PropagateChanges.Unset()
}

// GetReadOnly returns the ReadOnly field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ContentType) GetReadOnly() bool {
	if o == nil || o.ReadOnly.Get() == nil {
		var ret bool
		return ret
	}
	return *o.ReadOnly.Get()
}

// GetReadOnlyOk returns a tuple with the ReadOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContentType) GetReadOnlyOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ReadOnly.Get(), o.ReadOnly.IsSet()
}

// HasReadOnly returns a boolean if a field has been set.
func (o *ContentType) HasReadOnly() bool {
	if o != nil && o.ReadOnly.IsSet() {
		return true
	}

	return false
}

// SetReadOnly gets a reference to the given NullableBool and assigns it to the ReadOnly field.
func (o *ContentType) SetReadOnly(v bool) {
	o.ReadOnly.Set(&v)
}
// SetReadOnlyNil sets the value for ReadOnly to be an explicit nil
func (o *ContentType) SetReadOnlyNil() {
	o.ReadOnly.Set(nil)
}

// UnsetReadOnly ensures that no value is present for ReadOnly, not even an explicit nil
func (o *ContentType) UnsetReadOnly() {
	o.ReadOnly.Unset()
}

// GetSealed returns the Sealed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ContentType) GetSealed() bool {
	if o == nil || o.Sealed.Get() == nil {
		var ret bool
		return ret
	}
	return *o.Sealed.Get()
}

// GetSealedOk returns a tuple with the Sealed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContentType) GetSealedOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Sealed.Get(), o.Sealed.IsSet()
}

// HasSealed returns a boolean if a field has been set.
func (o *ContentType) HasSealed() bool {
	if o != nil && o.Sealed.IsSet() {
		return true
	}

	return false
}

// SetSealed gets a reference to the given NullableBool and assigns it to the Sealed field.
func (o *ContentType) SetSealed(v bool) {
	o.Sealed.Set(&v)
}
// SetSealedNil sets the value for Sealed to be an explicit nil
func (o *ContentType) SetSealedNil() {
	o.Sealed.Set(nil)
}

// UnsetSealed ensures that no value is present for Sealed, not even an explicit nil
func (o *ContentType) UnsetSealed() {
	o.Sealed.Unset()
}

// GetBase returns the Base field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ContentType) GetBase() AnyOfmicrosoftGraphContentType {
	if o == nil  {
		var ret AnyOfmicrosoftGraphContentType
		return ret
	}
	return o.Base
}

// GetBaseOk returns a tuple with the Base field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContentType) GetBaseOk() (*AnyOfmicrosoftGraphContentType, bool) {
	if o == nil || o.Base == nil {
		return nil, false
	}
	return &o.Base, true
}

// HasBase returns a boolean if a field has been set.
func (o *ContentType) HasBase() bool {
	if o != nil && o.Base != nil {
		return true
	}

	return false
}

// SetBase gets a reference to the given AnyOfmicrosoftGraphContentType and assigns it to the Base field.
func (o *ContentType) SetBase(v AnyOfmicrosoftGraphContentType) {
	o.Base = v
}

// GetBaseTypes returns the BaseTypes field value if set, zero value otherwise.
func (o *ContentType) GetBaseTypes() []MicrosoftGraphContentType {
	if o == nil || o.BaseTypes == nil {
		var ret []MicrosoftGraphContentType
		return ret
	}
	return *o.BaseTypes
}

// GetBaseTypesOk returns a tuple with the BaseTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentType) GetBaseTypesOk() (*[]MicrosoftGraphContentType, bool) {
	if o == nil || o.BaseTypes == nil {
		return nil, false
	}
	return o.BaseTypes, true
}

// HasBaseTypes returns a boolean if a field has been set.
func (o *ContentType) HasBaseTypes() bool {
	if o != nil && o.BaseTypes != nil {
		return true
	}

	return false
}

// SetBaseTypes gets a reference to the given []MicrosoftGraphContentType and assigns it to the BaseTypes field.
func (o *ContentType) SetBaseTypes(v []MicrosoftGraphContentType) {
	o.BaseTypes = &v
}

// GetColumnLinks returns the ColumnLinks field value if set, zero value otherwise.
func (o *ContentType) GetColumnLinks() []MicrosoftGraphColumnLink {
	if o == nil || o.ColumnLinks == nil {
		var ret []MicrosoftGraphColumnLink
		return ret
	}
	return *o.ColumnLinks
}

// GetColumnLinksOk returns a tuple with the ColumnLinks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentType) GetColumnLinksOk() (*[]MicrosoftGraphColumnLink, bool) {
	if o == nil || o.ColumnLinks == nil {
		return nil, false
	}
	return o.ColumnLinks, true
}

// HasColumnLinks returns a boolean if a field has been set.
func (o *ContentType) HasColumnLinks() bool {
	if o != nil && o.ColumnLinks != nil {
		return true
	}

	return false
}

// SetColumnLinks gets a reference to the given []MicrosoftGraphColumnLink and assigns it to the ColumnLinks field.
func (o *ContentType) SetColumnLinks(v []MicrosoftGraphColumnLink) {
	o.ColumnLinks = &v
}

// GetColumnPositions returns the ColumnPositions field value if set, zero value otherwise.
func (o *ContentType) GetColumnPositions() []MicrosoftGraphColumnDefinition {
	if o == nil || o.ColumnPositions == nil {
		var ret []MicrosoftGraphColumnDefinition
		return ret
	}
	return *o.ColumnPositions
}

// GetColumnPositionsOk returns a tuple with the ColumnPositions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentType) GetColumnPositionsOk() (*[]MicrosoftGraphColumnDefinition, bool) {
	if o == nil || o.ColumnPositions == nil {
		return nil, false
	}
	return o.ColumnPositions, true
}

// HasColumnPositions returns a boolean if a field has been set.
func (o *ContentType) HasColumnPositions() bool {
	if o != nil && o.ColumnPositions != nil {
		return true
	}

	return false
}

// SetColumnPositions gets a reference to the given []MicrosoftGraphColumnDefinition and assigns it to the ColumnPositions field.
func (o *ContentType) SetColumnPositions(v []MicrosoftGraphColumnDefinition) {
	o.ColumnPositions = &v
}

// GetColumns returns the Columns field value if set, zero value otherwise.
func (o *ContentType) GetColumns() []MicrosoftGraphColumnDefinition {
	if o == nil || o.Columns == nil {
		var ret []MicrosoftGraphColumnDefinition
		return ret
	}
	return *o.Columns
}

// GetColumnsOk returns a tuple with the Columns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentType) GetColumnsOk() (*[]MicrosoftGraphColumnDefinition, bool) {
	if o == nil || o.Columns == nil {
		return nil, false
	}
	return o.Columns, true
}

// HasColumns returns a boolean if a field has been set.
func (o *ContentType) HasColumns() bool {
	if o != nil && o.Columns != nil {
		return true
	}

	return false
}

// SetColumns gets a reference to the given []MicrosoftGraphColumnDefinition and assigns it to the Columns field.
func (o *ContentType) SetColumns(v []MicrosoftGraphColumnDefinition) {
	o.Columns = &v
}

func (o ContentType) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AssociatedHubsUrls != nil {
		toSerialize["associatedHubsUrls"] = o.AssociatedHubsUrls
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.DocumentSet != nil {
		toSerialize["documentSet"] = o.DocumentSet
	}
	if o.DocumentTemplate != nil {
		toSerialize["documentTemplate"] = o.DocumentTemplate
	}
	if o.Group.IsSet() {
		toSerialize["group"] = o.Group.Get()
	}
	if o.Hidden.IsSet() {
		toSerialize["hidden"] = o.Hidden.Get()
	}
	if o.InheritedFrom != nil {
		toSerialize["inheritedFrom"] = o.InheritedFrom
	}
	if o.IsBuiltIn.IsSet() {
		toSerialize["isBuiltIn"] = o.IsBuiltIn.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Order != nil {
		toSerialize["order"] = o.Order
	}
	if o.ParentId.IsSet() {
		toSerialize["parentId"] = o.ParentId.Get()
	}
	if o.PropagateChanges.IsSet() {
		toSerialize["propagateChanges"] = o.PropagateChanges.Get()
	}
	if o.ReadOnly.IsSet() {
		toSerialize["readOnly"] = o.ReadOnly.Get()
	}
	if o.Sealed.IsSet() {
		toSerialize["sealed"] = o.Sealed.Get()
	}
	if o.Base != nil {
		toSerialize["base"] = o.Base
	}
	if o.BaseTypes != nil {
		toSerialize["baseTypes"] = o.BaseTypes
	}
	if o.ColumnLinks != nil {
		toSerialize["columnLinks"] = o.ColumnLinks
	}
	if o.ColumnPositions != nil {
		toSerialize["columnPositions"] = o.ColumnPositions
	}
	if o.Columns != nil {
		toSerialize["columns"] = o.Columns
	}
	return json.Marshal(toSerialize)
}

type NullableContentType struct {
	value *ContentType
	isSet bool
}

func (v NullableContentType) Get() *ContentType {
	return v.value
}

func (v *NullableContentType) Set(val *ContentType) {
	v.value = val
	v.isSet = true
}

func (v NullableContentType) IsSet() bool {
	return v.isSet
}

func (v *NullableContentType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContentType(val *ContentType) *NullableContentType {
	return &NullableContentType{value: val, isSet: true}
}

func (v NullableContentType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContentType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


