# MicrosoftGraphContentType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
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

### NewMicrosoftGraphContentType

`func NewMicrosoftGraphContentType() *MicrosoftGraphContentType`

NewMicrosoftGraphContentType instantiates a new MicrosoftGraphContentType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphContentTypeWithDefaults

`func NewMicrosoftGraphContentTypeWithDefaults() *MicrosoftGraphContentType`

NewMicrosoftGraphContentTypeWithDefaults instantiates a new MicrosoftGraphContentType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphContentType) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphContentType) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphContentType) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphContentType) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAssociatedHubsUrls

`func (o *MicrosoftGraphContentType) GetAssociatedHubsUrls() []*string`

GetAssociatedHubsUrls returns the AssociatedHubsUrls field if non-nil, zero value otherwise.

### GetAssociatedHubsUrlsOk

`func (o *MicrosoftGraphContentType) GetAssociatedHubsUrlsOk() (*[]*string, bool)`

GetAssociatedHubsUrlsOk returns a tuple with the AssociatedHubsUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedHubsUrls

`func (o *MicrosoftGraphContentType) SetAssociatedHubsUrls(v []*string)`

SetAssociatedHubsUrls sets AssociatedHubsUrls field to given value.

### HasAssociatedHubsUrls

`func (o *MicrosoftGraphContentType) HasAssociatedHubsUrls() bool`

HasAssociatedHubsUrls returns a boolean if a field has been set.

### GetDescription

`func (o *MicrosoftGraphContentType) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MicrosoftGraphContentType) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MicrosoftGraphContentType) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MicrosoftGraphContentType) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *MicrosoftGraphContentType) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *MicrosoftGraphContentType) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDocumentSet

`func (o *MicrosoftGraphContentType) GetDocumentSet() AnyOfmicrosoftGraphDocumentSet`

GetDocumentSet returns the DocumentSet field if non-nil, zero value otherwise.

### GetDocumentSetOk

`func (o *MicrosoftGraphContentType) GetDocumentSetOk() (*AnyOfmicrosoftGraphDocumentSet, bool)`

GetDocumentSetOk returns a tuple with the DocumentSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentSet

`func (o *MicrosoftGraphContentType) SetDocumentSet(v AnyOfmicrosoftGraphDocumentSet)`

SetDocumentSet sets DocumentSet field to given value.

### HasDocumentSet

`func (o *MicrosoftGraphContentType) HasDocumentSet() bool`

HasDocumentSet returns a boolean if a field has been set.

### SetDocumentSetNil

`func (o *MicrosoftGraphContentType) SetDocumentSetNil(b bool)`

 SetDocumentSetNil sets the value for DocumentSet to be an explicit nil

### UnsetDocumentSet
`func (o *MicrosoftGraphContentType) UnsetDocumentSet()`

UnsetDocumentSet ensures that no value is present for DocumentSet, not even an explicit nil
### GetDocumentTemplate

`func (o *MicrosoftGraphContentType) GetDocumentTemplate() AnyOfmicrosoftGraphDocumentSetContent`

GetDocumentTemplate returns the DocumentTemplate field if non-nil, zero value otherwise.

### GetDocumentTemplateOk

`func (o *MicrosoftGraphContentType) GetDocumentTemplateOk() (*AnyOfmicrosoftGraphDocumentSetContent, bool)`

GetDocumentTemplateOk returns a tuple with the DocumentTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentTemplate

`func (o *MicrosoftGraphContentType) SetDocumentTemplate(v AnyOfmicrosoftGraphDocumentSetContent)`

SetDocumentTemplate sets DocumentTemplate field to given value.

### HasDocumentTemplate

`func (o *MicrosoftGraphContentType) HasDocumentTemplate() bool`

HasDocumentTemplate returns a boolean if a field has been set.

### SetDocumentTemplateNil

`func (o *MicrosoftGraphContentType) SetDocumentTemplateNil(b bool)`

 SetDocumentTemplateNil sets the value for DocumentTemplate to be an explicit nil

### UnsetDocumentTemplate
`func (o *MicrosoftGraphContentType) UnsetDocumentTemplate()`

UnsetDocumentTemplate ensures that no value is present for DocumentTemplate, not even an explicit nil
### GetGroup

`func (o *MicrosoftGraphContentType) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *MicrosoftGraphContentType) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *MicrosoftGraphContentType) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *MicrosoftGraphContentType) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### SetGroupNil

`func (o *MicrosoftGraphContentType) SetGroupNil(b bool)`

 SetGroupNil sets the value for Group to be an explicit nil

### UnsetGroup
`func (o *MicrosoftGraphContentType) UnsetGroup()`

UnsetGroup ensures that no value is present for Group, not even an explicit nil
### GetHidden

`func (o *MicrosoftGraphContentType) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *MicrosoftGraphContentType) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *MicrosoftGraphContentType) SetHidden(v bool)`

SetHidden sets Hidden field to given value.

### HasHidden

`func (o *MicrosoftGraphContentType) HasHidden() bool`

HasHidden returns a boolean if a field has been set.

### SetHiddenNil

`func (o *MicrosoftGraphContentType) SetHiddenNil(b bool)`

 SetHiddenNil sets the value for Hidden to be an explicit nil

### UnsetHidden
`func (o *MicrosoftGraphContentType) UnsetHidden()`

UnsetHidden ensures that no value is present for Hidden, not even an explicit nil
### GetInheritedFrom

`func (o *MicrosoftGraphContentType) GetInheritedFrom() AnyOfmicrosoftGraphItemReference`

GetInheritedFrom returns the InheritedFrom field if non-nil, zero value otherwise.

### GetInheritedFromOk

`func (o *MicrosoftGraphContentType) GetInheritedFromOk() (*AnyOfmicrosoftGraphItemReference, bool)`

GetInheritedFromOk returns a tuple with the InheritedFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInheritedFrom

`func (o *MicrosoftGraphContentType) SetInheritedFrom(v AnyOfmicrosoftGraphItemReference)`

SetInheritedFrom sets InheritedFrom field to given value.

### HasInheritedFrom

`func (o *MicrosoftGraphContentType) HasInheritedFrom() bool`

HasInheritedFrom returns a boolean if a field has been set.

### SetInheritedFromNil

`func (o *MicrosoftGraphContentType) SetInheritedFromNil(b bool)`

 SetInheritedFromNil sets the value for InheritedFrom to be an explicit nil

### UnsetInheritedFrom
`func (o *MicrosoftGraphContentType) UnsetInheritedFrom()`

UnsetInheritedFrom ensures that no value is present for InheritedFrom, not even an explicit nil
### GetIsBuiltIn

`func (o *MicrosoftGraphContentType) GetIsBuiltIn() bool`

GetIsBuiltIn returns the IsBuiltIn field if non-nil, zero value otherwise.

### GetIsBuiltInOk

`func (o *MicrosoftGraphContentType) GetIsBuiltInOk() (*bool, bool)`

GetIsBuiltInOk returns a tuple with the IsBuiltIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBuiltIn

`func (o *MicrosoftGraphContentType) SetIsBuiltIn(v bool)`

SetIsBuiltIn sets IsBuiltIn field to given value.

### HasIsBuiltIn

`func (o *MicrosoftGraphContentType) HasIsBuiltIn() bool`

HasIsBuiltIn returns a boolean if a field has been set.

### SetIsBuiltInNil

`func (o *MicrosoftGraphContentType) SetIsBuiltInNil(b bool)`

 SetIsBuiltInNil sets the value for IsBuiltIn to be an explicit nil

### UnsetIsBuiltIn
`func (o *MicrosoftGraphContentType) UnsetIsBuiltIn()`

UnsetIsBuiltIn ensures that no value is present for IsBuiltIn, not even an explicit nil
### GetName

`func (o *MicrosoftGraphContentType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MicrosoftGraphContentType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MicrosoftGraphContentType) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MicrosoftGraphContentType) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *MicrosoftGraphContentType) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *MicrosoftGraphContentType) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetOrder

`func (o *MicrosoftGraphContentType) GetOrder() AnyOfmicrosoftGraphContentTypeOrder`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *MicrosoftGraphContentType) GetOrderOk() (*AnyOfmicrosoftGraphContentTypeOrder, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *MicrosoftGraphContentType) SetOrder(v AnyOfmicrosoftGraphContentTypeOrder)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *MicrosoftGraphContentType) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### SetOrderNil

`func (o *MicrosoftGraphContentType) SetOrderNil(b bool)`

 SetOrderNil sets the value for Order to be an explicit nil

### UnsetOrder
`func (o *MicrosoftGraphContentType) UnsetOrder()`

UnsetOrder ensures that no value is present for Order, not even an explicit nil
### GetParentId

`func (o *MicrosoftGraphContentType) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *MicrosoftGraphContentType) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *MicrosoftGraphContentType) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *MicrosoftGraphContentType) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### SetParentIdNil

`func (o *MicrosoftGraphContentType) SetParentIdNil(b bool)`

 SetParentIdNil sets the value for ParentId to be an explicit nil

### UnsetParentId
`func (o *MicrosoftGraphContentType) UnsetParentId()`

UnsetParentId ensures that no value is present for ParentId, not even an explicit nil
### GetPropagateChanges

`func (o *MicrosoftGraphContentType) GetPropagateChanges() bool`

GetPropagateChanges returns the PropagateChanges field if non-nil, zero value otherwise.

### GetPropagateChangesOk

`func (o *MicrosoftGraphContentType) GetPropagateChangesOk() (*bool, bool)`

GetPropagateChangesOk returns a tuple with the PropagateChanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropagateChanges

`func (o *MicrosoftGraphContentType) SetPropagateChanges(v bool)`

SetPropagateChanges sets PropagateChanges field to given value.

### HasPropagateChanges

`func (o *MicrosoftGraphContentType) HasPropagateChanges() bool`

HasPropagateChanges returns a boolean if a field has been set.

### SetPropagateChangesNil

`func (o *MicrosoftGraphContentType) SetPropagateChangesNil(b bool)`

 SetPropagateChangesNil sets the value for PropagateChanges to be an explicit nil

### UnsetPropagateChanges
`func (o *MicrosoftGraphContentType) UnsetPropagateChanges()`

UnsetPropagateChanges ensures that no value is present for PropagateChanges, not even an explicit nil
### GetReadOnly

`func (o *MicrosoftGraphContentType) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *MicrosoftGraphContentType) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *MicrosoftGraphContentType) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *MicrosoftGraphContentType) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.

### SetReadOnlyNil

`func (o *MicrosoftGraphContentType) SetReadOnlyNil(b bool)`

 SetReadOnlyNil sets the value for ReadOnly to be an explicit nil

### UnsetReadOnly
`func (o *MicrosoftGraphContentType) UnsetReadOnly()`

UnsetReadOnly ensures that no value is present for ReadOnly, not even an explicit nil
### GetSealed

`func (o *MicrosoftGraphContentType) GetSealed() bool`

GetSealed returns the Sealed field if non-nil, zero value otherwise.

### GetSealedOk

`func (o *MicrosoftGraphContentType) GetSealedOk() (*bool, bool)`

GetSealedOk returns a tuple with the Sealed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSealed

`func (o *MicrosoftGraphContentType) SetSealed(v bool)`

SetSealed sets Sealed field to given value.

### HasSealed

`func (o *MicrosoftGraphContentType) HasSealed() bool`

HasSealed returns a boolean if a field has been set.

### SetSealedNil

`func (o *MicrosoftGraphContentType) SetSealedNil(b bool)`

 SetSealedNil sets the value for Sealed to be an explicit nil

### UnsetSealed
`func (o *MicrosoftGraphContentType) UnsetSealed()`

UnsetSealed ensures that no value is present for Sealed, not even an explicit nil
### GetBase

`func (o *MicrosoftGraphContentType) GetBase() AnyOfmicrosoftGraphContentType`

GetBase returns the Base field if non-nil, zero value otherwise.

### GetBaseOk

`func (o *MicrosoftGraphContentType) GetBaseOk() (*AnyOfmicrosoftGraphContentType, bool)`

GetBaseOk returns a tuple with the Base field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBase

`func (o *MicrosoftGraphContentType) SetBase(v AnyOfmicrosoftGraphContentType)`

SetBase sets Base field to given value.

### HasBase

`func (o *MicrosoftGraphContentType) HasBase() bool`

HasBase returns a boolean if a field has been set.

### SetBaseNil

`func (o *MicrosoftGraphContentType) SetBaseNil(b bool)`

 SetBaseNil sets the value for Base to be an explicit nil

### UnsetBase
`func (o *MicrosoftGraphContentType) UnsetBase()`

UnsetBase ensures that no value is present for Base, not even an explicit nil
### GetBaseTypes

`func (o *MicrosoftGraphContentType) GetBaseTypes() []MicrosoftGraphContentType`

GetBaseTypes returns the BaseTypes field if non-nil, zero value otherwise.

### GetBaseTypesOk

`func (o *MicrosoftGraphContentType) GetBaseTypesOk() (*[]MicrosoftGraphContentType, bool)`

GetBaseTypesOk returns a tuple with the BaseTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseTypes

`func (o *MicrosoftGraphContentType) SetBaseTypes(v []MicrosoftGraphContentType)`

SetBaseTypes sets BaseTypes field to given value.

### HasBaseTypes

`func (o *MicrosoftGraphContentType) HasBaseTypes() bool`

HasBaseTypes returns a boolean if a field has been set.

### GetColumnLinks

`func (o *MicrosoftGraphContentType) GetColumnLinks() []MicrosoftGraphColumnLink`

GetColumnLinks returns the ColumnLinks field if non-nil, zero value otherwise.

### GetColumnLinksOk

`func (o *MicrosoftGraphContentType) GetColumnLinksOk() (*[]MicrosoftGraphColumnLink, bool)`

GetColumnLinksOk returns a tuple with the ColumnLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnLinks

`func (o *MicrosoftGraphContentType) SetColumnLinks(v []MicrosoftGraphColumnLink)`

SetColumnLinks sets ColumnLinks field to given value.

### HasColumnLinks

`func (o *MicrosoftGraphContentType) HasColumnLinks() bool`

HasColumnLinks returns a boolean if a field has been set.

### GetColumnPositions

`func (o *MicrosoftGraphContentType) GetColumnPositions() []MicrosoftGraphColumnDefinition`

GetColumnPositions returns the ColumnPositions field if non-nil, zero value otherwise.

### GetColumnPositionsOk

`func (o *MicrosoftGraphContentType) GetColumnPositionsOk() (*[]MicrosoftGraphColumnDefinition, bool)`

GetColumnPositionsOk returns a tuple with the ColumnPositions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnPositions

`func (o *MicrosoftGraphContentType) SetColumnPositions(v []MicrosoftGraphColumnDefinition)`

SetColumnPositions sets ColumnPositions field to given value.

### HasColumnPositions

`func (o *MicrosoftGraphContentType) HasColumnPositions() bool`

HasColumnPositions returns a boolean if a field has been set.

### GetColumns

`func (o *MicrosoftGraphContentType) GetColumns() []MicrosoftGraphColumnDefinition`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *MicrosoftGraphContentType) GetColumnsOk() (*[]MicrosoftGraphColumnDefinition, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *MicrosoftGraphContentType) SetColumns(v []MicrosoftGraphColumnDefinition)`

SetColumns sets Columns field to given value.

### HasColumns

`func (o *MicrosoftGraphContentType) HasColumns() bool`

HasColumns returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


