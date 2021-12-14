# MicrosoftGraphWorkbookNamedItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**Comment** | Pointer to **NullableString** | Represents the comment associated with this name. | [optional] 
**Name** | Pointer to **NullableString** | The name of the object. Read-only. | [optional] 
**Scope** | Pointer to **string** | Indicates whether the name is scoped to the workbook or to a specific worksheet. Read-only. | [optional] 
**Type** | Pointer to **NullableString** | Indicates what type of reference is associated with the name. The possible values are: String, Integer, Double, Boolean, Range. Read-only. | [optional] 
**Value** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) | Represents the formula that the name is defined to refer to. E.g. &#x3D;Sheet14!$B$2:$H$12, &#x3D;4.75, etc. Read-only. | [optional] 
**Visible** | Pointer to **bool** | Specifies whether the object is visible or not. | [optional] 
**Worksheet** | Pointer to [**AnyOfmicrosoftGraphWorkbookWorksheet**](anyOf&lt;microsoft.graph.workbookWorksheet&gt;.md) | Returns the worksheet on which the named item is scoped to. Available only if the item is scoped to the worksheet. Read-only. | [optional] 

## Methods

### NewMicrosoftGraphWorkbookNamedItem

`func NewMicrosoftGraphWorkbookNamedItem() *MicrosoftGraphWorkbookNamedItem`

NewMicrosoftGraphWorkbookNamedItem instantiates a new MicrosoftGraphWorkbookNamedItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphWorkbookNamedItemWithDefaults

`func NewMicrosoftGraphWorkbookNamedItemWithDefaults() *MicrosoftGraphWorkbookNamedItem`

NewMicrosoftGraphWorkbookNamedItemWithDefaults instantiates a new MicrosoftGraphWorkbookNamedItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphWorkbookNamedItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphWorkbookNamedItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphWorkbookNamedItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphWorkbookNamedItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetComment

`func (o *MicrosoftGraphWorkbookNamedItem) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *MicrosoftGraphWorkbookNamedItem) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *MicrosoftGraphWorkbookNamedItem) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *MicrosoftGraphWorkbookNamedItem) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetCommentNil

`func (o *MicrosoftGraphWorkbookNamedItem) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *MicrosoftGraphWorkbookNamedItem) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil
### GetName

`func (o *MicrosoftGraphWorkbookNamedItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MicrosoftGraphWorkbookNamedItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MicrosoftGraphWorkbookNamedItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MicrosoftGraphWorkbookNamedItem) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *MicrosoftGraphWorkbookNamedItem) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *MicrosoftGraphWorkbookNamedItem) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetScope

`func (o *MicrosoftGraphWorkbookNamedItem) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *MicrosoftGraphWorkbookNamedItem) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *MicrosoftGraphWorkbookNamedItem) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *MicrosoftGraphWorkbookNamedItem) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetType

`func (o *MicrosoftGraphWorkbookNamedItem) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MicrosoftGraphWorkbookNamedItem) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MicrosoftGraphWorkbookNamedItem) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *MicrosoftGraphWorkbookNamedItem) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *MicrosoftGraphWorkbookNamedItem) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *MicrosoftGraphWorkbookNamedItem) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetValue

`func (o *MicrosoftGraphWorkbookNamedItem) GetValue() AnyOfobject`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *MicrosoftGraphWorkbookNamedItem) GetValueOk() (*AnyOfobject, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *MicrosoftGraphWorkbookNamedItem) SetValue(v AnyOfobject)`

SetValue sets Value field to given value.

### HasValue

`func (o *MicrosoftGraphWorkbookNamedItem) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *MicrosoftGraphWorkbookNamedItem) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *MicrosoftGraphWorkbookNamedItem) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetVisible

`func (o *MicrosoftGraphWorkbookNamedItem) GetVisible() bool`

GetVisible returns the Visible field if non-nil, zero value otherwise.

### GetVisibleOk

`func (o *MicrosoftGraphWorkbookNamedItem) GetVisibleOk() (*bool, bool)`

GetVisibleOk returns a tuple with the Visible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisible

`func (o *MicrosoftGraphWorkbookNamedItem) SetVisible(v bool)`

SetVisible sets Visible field to given value.

### HasVisible

`func (o *MicrosoftGraphWorkbookNamedItem) HasVisible() bool`

HasVisible returns a boolean if a field has been set.

### GetWorksheet

`func (o *MicrosoftGraphWorkbookNamedItem) GetWorksheet() AnyOfmicrosoftGraphWorkbookWorksheet`

GetWorksheet returns the Worksheet field if non-nil, zero value otherwise.

### GetWorksheetOk

`func (o *MicrosoftGraphWorkbookNamedItem) GetWorksheetOk() (*AnyOfmicrosoftGraphWorkbookWorksheet, bool)`

GetWorksheetOk returns a tuple with the Worksheet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorksheet

`func (o *MicrosoftGraphWorkbookNamedItem) SetWorksheet(v AnyOfmicrosoftGraphWorkbookWorksheet)`

SetWorksheet sets Worksheet field to given value.

### HasWorksheet

`func (o *MicrosoftGraphWorkbookNamedItem) HasWorksheet() bool`

HasWorksheet returns a boolean if a field has been set.

### SetWorksheetNil

`func (o *MicrosoftGraphWorkbookNamedItem) SetWorksheetNil(b bool)`

 SetWorksheetNil sets the value for Worksheet to be an explicit nil

### UnsetWorksheet
`func (o *MicrosoftGraphWorkbookNamedItem) UnsetWorksheet()`

UnsetWorksheet ensures that no value is present for Worksheet, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


