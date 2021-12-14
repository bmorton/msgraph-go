# WorkbookWorksheetProtection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Options** | Pointer to [**AnyOfmicrosoftGraphWorkbookWorksheetProtectionOptions**](anyOf&lt;microsoft.graph.workbookWorksheetProtectionOptions&gt;.md) | Sheet protection options. Read-only. | [optional] 
**Protected** | Pointer to **bool** | Indicates if the worksheet is protected.  Read-only. | [optional] 

## Methods

### NewWorkbookWorksheetProtection

`func NewWorkbookWorksheetProtection() *WorkbookWorksheetProtection`

NewWorkbookWorksheetProtection instantiates a new WorkbookWorksheetProtection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkbookWorksheetProtectionWithDefaults

`func NewWorkbookWorksheetProtectionWithDefaults() *WorkbookWorksheetProtection`

NewWorkbookWorksheetProtectionWithDefaults instantiates a new WorkbookWorksheetProtection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOptions

`func (o *WorkbookWorksheetProtection) GetOptions() AnyOfmicrosoftGraphWorkbookWorksheetProtectionOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *WorkbookWorksheetProtection) GetOptionsOk() (*AnyOfmicrosoftGraphWorkbookWorksheetProtectionOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *WorkbookWorksheetProtection) SetOptions(v AnyOfmicrosoftGraphWorkbookWorksheetProtectionOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *WorkbookWorksheetProtection) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### SetOptionsNil

`func (o *WorkbookWorksheetProtection) SetOptionsNil(b bool)`

 SetOptionsNil sets the value for Options to be an explicit nil

### UnsetOptions
`func (o *WorkbookWorksheetProtection) UnsetOptions()`

UnsetOptions ensures that no value is present for Options, not even an explicit nil
### GetProtected

`func (o *WorkbookWorksheetProtection) GetProtected() bool`

GetProtected returns the Protected field if non-nil, zero value otherwise.

### GetProtectedOk

`func (o *WorkbookWorksheetProtection) GetProtectedOk() (*bool, bool)`

GetProtectedOk returns a tuple with the Protected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtected

`func (o *WorkbookWorksheetProtection) SetProtected(v bool)`

SetProtected sets Protected field to given value.

### HasProtected

`func (o *WorkbookWorksheetProtection) HasProtected() bool`

HasProtected returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


