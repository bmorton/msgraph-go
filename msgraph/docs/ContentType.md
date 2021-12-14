# ContentType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssociatedHubsUrls** | Pointer to **[]string** | List of canonical URLs for hub sites with which this content type is associated to. This will contain all hub sites where this content type is queued to be enforced or is already enforced. Enforcing a content type means that the content type will be applied to the lists in the enforced sites. | [optional] 
**Description** | Pointer to **NullableString** | The descriptive text for the item. | [optional] 
**DocumentSet** | Pointer to [**AnyOfmicrosoftGraphDocumentSet**](anyOf&lt;microsoft.graph.documentSet&gt;.md) | Document Set metadata. | [optional] 
**DocumentTemplate** | Pointer to [**AnyOfmicrosoftGraphDocumentSetContent**](anyOf&lt;microsoft.graph.documentSetContent&gt;.md) | Document template metadata. To make sure that documents have consistent content across a site and its subsites, you can associate a Word, Excel, or PowerPoint template with a site content type. | [optional] 
**Group** | Pointer to **NullableString** | The name of the group this content type belongs to. Helps organize related content types. | [optional] 
**Hidden** | Pointer to **NullableBool** | Indicates whether the content type is hidden in the list&#39;s &#39;New&#39; menu. | [optional] 
**InheritedFrom** | Pointer to [**AnyOfmicrosoftGraphItemReference**](anyOf&lt;microsoft.graph.itemReference&gt;.md) | If this content type is inherited from another scope (like a site), provides a reference to the item where the content type is defined. | [optional] 
**IsBuiltIn** | Pointer to **NullableBool** | Specifies if a content type is a built-in content type. | [optional] 
**Name** | Pointer to **NullableString** | The name of the content type. | [optional] 
**Order** | Pointer to [**AnyOfmicrosoftGraphContentTypeOrder**](anyOf&lt;microsoft.graph.contentTypeOrder&gt;.md) | Specifies the order in which the content type appears in the selection UI. | [optional] 
**ParentId** | Pointer to **NullableString** | The unique identifier of the content type. | [optional] 
**PropagateChanges** | Pointer to **NullableBool** | If true, any changes made to the content type will be pushed to inherited content types and lists that implement the content type. | [optional] 
**ReadOnly** | Pointer to **NullableBool** | If true, the content type can&#39;t be modified unless this value is first set to false. | [optional] 
**Sealed** | Pointer to **NullableBool** | If true, the content type can&#39;t be modified by users or through push-down operations. Only site collection administrators can seal or unseal content types. | [optional] 
**Base** | Pointer to [**AnyOfmicrosoftGraphContentType**](anyOf&lt;microsoft.graph.contentType&gt;.md) | Parent contentType from which this content type is derived. | [optional] 
**BaseTypes** | Pointer to [**[]MicrosoftGraphContentType**](MicrosoftGraphContentType.md) | The collection of content types that are ancestors of this content type. | [optional] 
**ColumnLinks** | Pointer to [**[]MicrosoftGraphColumnLink**](MicrosoftGraphColumnLink.md) | The collection of columns that are required by this content type. | [optional] 
**ColumnPositions** | Pointer to [**[]MicrosoftGraphColumnDefinition**](MicrosoftGraphColumnDefinition.md) | Column order information in a content type. | [optional] 
**Columns** | Pointer to [**[]MicrosoftGraphColumnDefinition**](MicrosoftGraphColumnDefinition.md) | The collection of column definitions for this contentType. | [optional] 

## Methods

### NewContentType

`func NewContentType() *ContentType`

NewContentType instantiates a new ContentType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentTypeWithDefaults

`func NewContentTypeWithDefaults() *ContentType`

NewContentTypeWithDefaults instantiates a new ContentType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssociatedHubsUrls

`func (o *ContentType) GetAssociatedHubsUrls() []*string`

GetAssociatedHubsUrls returns the AssociatedHubsUrls field if non-nil, zero value otherwise.

### GetAssociatedHubsUrlsOk

`func (o *ContentType) GetAssociatedHubsUrlsOk() (*[]*string, bool)`

GetAssociatedHubsUrlsOk returns a tuple with the AssociatedHubsUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedHubsUrls

`func (o *ContentType) SetAssociatedHubsUrls(v []*string)`

SetAssociatedHubsUrls sets AssociatedHubsUrls field to given value.

### HasAssociatedHubsUrls

`func (o *ContentType) HasAssociatedHubsUrls() bool`

HasAssociatedHubsUrls returns a boolean if a field has been set.

### GetDescription

`func (o *ContentType) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ContentType) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ContentType) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ContentType) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ContentType) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ContentType) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDocumentSet

`func (o *ContentType) GetDocumentSet() AnyOfmicrosoftGraphDocumentSet`

GetDocumentSet returns the DocumentSet field if non-nil, zero value otherwise.

### GetDocumentSetOk

`func (o *ContentType) GetDocumentSetOk() (*AnyOfmicrosoftGraphDocumentSet, bool)`

GetDocumentSetOk returns a tuple with the DocumentSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentSet

`func (o *ContentType) SetDocumentSet(v AnyOfmicrosoftGraphDocumentSet)`

SetDocumentSet sets DocumentSet field to given value.

### HasDocumentSet

`func (o *ContentType) HasDocumentSet() bool`

HasDocumentSet returns a boolean if a field has been set.

### SetDocumentSetNil

`func (o *ContentType) SetDocumentSetNil(b bool)`

 SetDocumentSetNil sets the value for DocumentSet to be an explicit nil

### UnsetDocumentSet
`func (o *ContentType) UnsetDocumentSet()`

UnsetDocumentSet ensures that no value is present for DocumentSet, not even an explicit nil
### GetDocumentTemplate

`func (o *ContentType) GetDocumentTemplate() AnyOfmicrosoftGraphDocumentSetContent`

GetDocumentTemplate returns the DocumentTemplate field if non-nil, zero value otherwise.

### GetDocumentTemplateOk

`func (o *ContentType) GetDocumentTemplateOk() (*AnyOfmicrosoftGraphDocumentSetContent, bool)`

GetDocumentTemplateOk returns a tuple with the DocumentTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentTemplate

`func (o *ContentType) SetDocumentTemplate(v AnyOfmicrosoftGraphDocumentSetContent)`

SetDocumentTemplate sets DocumentTemplate field to given value.

### HasDocumentTemplate

`func (o *ContentType) HasDocumentTemplate() bool`

HasDocumentTemplate returns a boolean if a field has been set.

### SetDocumentTemplateNil

`func (o *ContentType) SetDocumentTemplateNil(b bool)`

 SetDocumentTemplateNil sets the value for DocumentTemplate to be an explicit nil

### UnsetDocumentTemplate
`func (o *ContentType) UnsetDocumentTemplate()`

UnsetDocumentTemplate ensures that no value is present for DocumentTemplate, not even an explicit nil
### GetGroup

`func (o *ContentType) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *ContentType) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *ContentType) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *ContentType) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### SetGroupNil

`func (o *ContentType) SetGroupNil(b bool)`

 SetGroupNil sets the value for Group to be an explicit nil

### UnsetGroup
`func (o *ContentType) UnsetGroup()`

UnsetGroup ensures that no value is present for Group, not even an explicit nil
### GetHidden

`func (o *ContentType) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *ContentType) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *ContentType) SetHidden(v bool)`

SetHidden sets Hidden field to given value.

### HasHidden

`func (o *ContentType) HasHidden() bool`

HasHidden returns a boolean if a field has been set.

### SetHiddenNil

`func (o *ContentType) SetHiddenNil(b bool)`

 SetHiddenNil sets the value for Hidden to be an explicit nil

### UnsetHidden
`func (o *ContentType) UnsetHidden()`

UnsetHidden ensures that no value is present for Hidden, not even an explicit nil
### GetInheritedFrom

`func (o *ContentType) GetInheritedFrom() AnyOfmicrosoftGraphItemReference`

GetInheritedFrom returns the InheritedFrom field if non-nil, zero value otherwise.

### GetInheritedFromOk

`func (o *ContentType) GetInheritedFromOk() (*AnyOfmicrosoftGraphItemReference, bool)`

GetInheritedFromOk returns a tuple with the InheritedFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInheritedFrom

`func (o *ContentType) SetInheritedFrom(v AnyOfmicrosoftGraphItemReference)`

SetInheritedFrom sets InheritedFrom field to given value.

### HasInheritedFrom

`func (o *ContentType) HasInheritedFrom() bool`

HasInheritedFrom returns a boolean if a field has been set.

### SetInheritedFromNil

`func (o *ContentType) SetInheritedFromNil(b bool)`

 SetInheritedFromNil sets the value for InheritedFrom to be an explicit nil

### UnsetInheritedFrom
`func (o *ContentType) UnsetInheritedFrom()`

UnsetInheritedFrom ensures that no value is present for InheritedFrom, not even an explicit nil
### GetIsBuiltIn

`func (o *ContentType) GetIsBuiltIn() bool`

GetIsBuiltIn returns the IsBuiltIn field if non-nil, zero value otherwise.

### GetIsBuiltInOk

`func (o *ContentType) GetIsBuiltInOk() (*bool, bool)`

GetIsBuiltInOk returns a tuple with the IsBuiltIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBuiltIn

`func (o *ContentType) SetIsBuiltIn(v bool)`

SetIsBuiltIn sets IsBuiltIn field to given value.

### HasIsBuiltIn

`func (o *ContentType) HasIsBuiltIn() bool`

HasIsBuiltIn returns a boolean if a field has been set.

### SetIsBuiltInNil

`func (o *ContentType) SetIsBuiltInNil(b bool)`

 SetIsBuiltInNil sets the value for IsBuiltIn to be an explicit nil

### UnsetIsBuiltIn
`func (o *ContentType) UnsetIsBuiltIn()`

UnsetIsBuiltIn ensures that no value is present for IsBuiltIn, not even an explicit nil
### GetName

`func (o *ContentType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ContentType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ContentType) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ContentType) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ContentType) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ContentType) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetOrder

`func (o *ContentType) GetOrder() AnyOfmicrosoftGraphContentTypeOrder`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *ContentType) GetOrderOk() (*AnyOfmicrosoftGraphContentTypeOrder, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *ContentType) SetOrder(v AnyOfmicrosoftGraphContentTypeOrder)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *ContentType) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### SetOrderNil

`func (o *ContentType) SetOrderNil(b bool)`

 SetOrderNil sets the value for Order to be an explicit nil

### UnsetOrder
`func (o *ContentType) UnsetOrder()`

UnsetOrder ensures that no value is present for Order, not even an explicit nil
### GetParentId

`func (o *ContentType) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *ContentType) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *ContentType) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *ContentType) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### SetParentIdNil

`func (o *ContentType) SetParentIdNil(b bool)`

 SetParentIdNil sets the value for ParentId to be an explicit nil

### UnsetParentId
`func (o *ContentType) UnsetParentId()`

UnsetParentId ensures that no value is present for ParentId, not even an explicit nil
### GetPropagateChanges

`func (o *ContentType) GetPropagateChanges() bool`

GetPropagateChanges returns the PropagateChanges field if non-nil, zero value otherwise.

### GetPropagateChangesOk

`func (o *ContentType) GetPropagateChangesOk() (*bool, bool)`

GetPropagateChangesOk returns a tuple with the PropagateChanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropagateChanges

`func (o *ContentType) SetPropagateChanges(v bool)`

SetPropagateChanges sets PropagateChanges field to given value.

### HasPropagateChanges

`func (o *ContentType) HasPropagateChanges() bool`

HasPropagateChanges returns a boolean if a field has been set.

### SetPropagateChangesNil

`func (o *ContentType) SetPropagateChangesNil(b bool)`

 SetPropagateChangesNil sets the value for PropagateChanges to be an explicit nil

### UnsetPropagateChanges
`func (o *ContentType) UnsetPropagateChanges()`

UnsetPropagateChanges ensures that no value is present for PropagateChanges, not even an explicit nil
### GetReadOnly

`func (o *ContentType) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *ContentType) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *ContentType) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *ContentType) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.

### SetReadOnlyNil

`func (o *ContentType) SetReadOnlyNil(b bool)`

 SetReadOnlyNil sets the value for ReadOnly to be an explicit nil

### UnsetReadOnly
`func (o *ContentType) UnsetReadOnly()`

UnsetReadOnly ensures that no value is present for ReadOnly, not even an explicit nil
### GetSealed

`func (o *ContentType) GetSealed() bool`

GetSealed returns the Sealed field if non-nil, zero value otherwise.

### GetSealedOk

`func (o *ContentType) GetSealedOk() (*bool, bool)`

GetSealedOk returns a tuple with the Sealed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSealed

`func (o *ContentType) SetSealed(v bool)`

SetSealed sets Sealed field to given value.

### HasSealed

`func (o *ContentType) HasSealed() bool`

HasSealed returns a boolean if a field has been set.

### SetSealedNil

`func (o *ContentType) SetSealedNil(b bool)`

 SetSealedNil sets the value for Sealed to be an explicit nil

### UnsetSealed
`func (o *ContentType) UnsetSealed()`

UnsetSealed ensures that no value is present for Sealed, not even an explicit nil
### GetBase

`func (o *ContentType) GetBase() AnyOfmicrosoftGraphContentType`

GetBase returns the Base field if non-nil, zero value otherwise.

### GetBaseOk

`func (o *ContentType) GetBaseOk() (*AnyOfmicrosoftGraphContentType, bool)`

GetBaseOk returns a tuple with the Base field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBase

`func (o *ContentType) SetBase(v AnyOfmicrosoftGraphContentType)`

SetBase sets Base field to given value.

### HasBase

`func (o *ContentType) HasBase() bool`

HasBase returns a boolean if a field has been set.

### SetBaseNil

`func (o *ContentType) SetBaseNil(b bool)`

 SetBaseNil sets the value for Base to be an explicit nil

### UnsetBase
`func (o *ContentType) UnsetBase()`

UnsetBase ensures that no value is present for Base, not even an explicit nil
### GetBaseTypes

`func (o *ContentType) GetBaseTypes() []MicrosoftGraphContentType`

GetBaseTypes returns the BaseTypes field if non-nil, zero value otherwise.

### GetBaseTypesOk

`func (o *ContentType) GetBaseTypesOk() (*[]MicrosoftGraphContentType, bool)`

GetBaseTypesOk returns a tuple with the BaseTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseTypes

`func (o *ContentType) SetBaseTypes(v []MicrosoftGraphContentType)`

SetBaseTypes sets BaseTypes field to given value.

### HasBaseTypes

`func (o *ContentType) HasBaseTypes() bool`

HasBaseTypes returns a boolean if a field has been set.

### GetColumnLinks

`func (o *ContentType) GetColumnLinks() []MicrosoftGraphColumnLink`

GetColumnLinks returns the ColumnLinks field if non-nil, zero value otherwise.

### GetColumnLinksOk

`func (o *ContentType) GetColumnLinksOk() (*[]MicrosoftGraphColumnLink, bool)`

GetColumnLinksOk returns a tuple with the ColumnLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnLinks

`func (o *ContentType) SetColumnLinks(v []MicrosoftGraphColumnLink)`

SetColumnLinks sets ColumnLinks field to given value.

### HasColumnLinks

`func (o *ContentType) HasColumnLinks() bool`

HasColumnLinks returns a boolean if a field has been set.

### GetColumnPositions

`func (o *ContentType) GetColumnPositions() []MicrosoftGraphColumnDefinition`

GetColumnPositions returns the ColumnPositions field if non-nil, zero value otherwise.

### GetColumnPositionsOk

`func (o *ContentType) GetColumnPositionsOk() (*[]MicrosoftGraphColumnDefinition, bool)`

GetColumnPositionsOk returns a tuple with the ColumnPositions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnPositions

`func (o *ContentType) SetColumnPositions(v []MicrosoftGraphColumnDefinition)`

SetColumnPositions sets ColumnPositions field to given value.

### HasColumnPositions

`func (o *ContentType) HasColumnPositions() bool`

HasColumnPositions returns a boolean if a field has been set.

### GetColumns

`func (o *ContentType) GetColumns() []MicrosoftGraphColumnDefinition`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *ContentType) GetColumnsOk() (*[]MicrosoftGraphColumnDefinition, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *ContentType) SetColumns(v []MicrosoftGraphColumnDefinition)`

SetColumns sets Columns field to given value.

### HasColumns

`func (o *ContentType) HasColumns() bool`

HasColumns returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


