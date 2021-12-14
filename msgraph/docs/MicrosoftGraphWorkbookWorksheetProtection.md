# MicrosoftGraphWorkbookWorksheetProtection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**Options** | Pointer to [**AnyOfmicrosoftGraphWorkbookWorksheetProtectionOptions**](anyOf&lt;microsoft.graph.workbookWorksheetProtectionOptions&gt;.md) | Sheet protection options. Read-only. | [optional] 
**Protected** | Pointer to **bool** | Indicates if the worksheet is protected.  Read-only. | [optional] 

## Methods

### NewMicrosoftGraphWorkbookWorksheetProtection

`func NewMicrosoftGraphWorkbookWorksheetProtection() *MicrosoftGraphWorkbookWorksheetProtection`

NewMicrosoftGraphWorkbookWorksheetProtection instantiates a new MicrosoftGraphWorkbookWorksheetProtection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphWorkbookWorksheetProtectionWithDefaults

`func NewMicrosoftGraphWorkbookWorksheetProtectionWithDefaults() *MicrosoftGraphWorkbookWorksheetProtection`

NewMicrosoftGraphWorkbookWorksheetProtectionWithDefaults instantiates a new MicrosoftGraphWorkbookWorksheetProtection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphWorkbookWorksheetProtection) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphWorkbookWorksheetProtection) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphWorkbookWorksheetProtection) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphWorkbookWorksheetProtection) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOptions

`func (o *MicrosoftGraphWorkbookWorksheetProtection) GetOptions() AnyOfmicrosoftGraphWorkbookWorksheetProtectionOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *MicrosoftGraphWorkbookWorksheetProtection) GetOptionsOk() (*AnyOfmicrosoftGraphWorkbookWorksheetProtectionOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *MicrosoftGraphWorkbookWorksheetProtection) SetOptions(v AnyOfmicrosoftGraphWorkbookWorksheetProtectionOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *MicrosoftGraphWorkbookWorksheetProtection) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### SetOptionsNil

`func (o *MicrosoftGraphWorkbookWorksheetProtection) SetOptionsNil(b bool)`

 SetOptionsNil sets the value for Options to be an explicit nil

### UnsetOptions
`func (o *MicrosoftGraphWorkbookWorksheetProtection) UnsetOptions()`

UnsetOptions ensures that no value is present for Options, not even an explicit nil
### GetProtected

`func (o *MicrosoftGraphWorkbookWorksheetProtection) GetProtected() bool`

GetProtected returns the Protected field if non-nil, zero value otherwise.

### GetProtectedOk

`func (o *MicrosoftGraphWorkbookWorksheetProtection) GetProtectedOk() (*bool, bool)`

GetProtectedOk returns a tuple with the Protected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtected

`func (o *MicrosoftGraphWorkbookWorksheetProtection) SetProtected(v bool)`

SetProtected sets Protected field to given value.

### HasProtected

`func (o *MicrosoftGraphWorkbookWorksheetProtection) HasProtected() bool`

HasProtected returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


