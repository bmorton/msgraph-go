# MicrosoftGraphOutlookUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**MasterCategories** | Pointer to [**[]MicrosoftGraphOutlookCategory**](MicrosoftGraphOutlookCategory.md) | A list of categories defined for the user. | [optional] 

## Methods

### NewMicrosoftGraphOutlookUser

`func NewMicrosoftGraphOutlookUser() *MicrosoftGraphOutlookUser`

NewMicrosoftGraphOutlookUser instantiates a new MicrosoftGraphOutlookUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphOutlookUserWithDefaults

`func NewMicrosoftGraphOutlookUserWithDefaults() *MicrosoftGraphOutlookUser`

NewMicrosoftGraphOutlookUserWithDefaults instantiates a new MicrosoftGraphOutlookUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphOutlookUser) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphOutlookUser) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphOutlookUser) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphOutlookUser) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMasterCategories

`func (o *MicrosoftGraphOutlookUser) GetMasterCategories() []MicrosoftGraphOutlookCategory`

GetMasterCategories returns the MasterCategories field if non-nil, zero value otherwise.

### GetMasterCategoriesOk

`func (o *MicrosoftGraphOutlookUser) GetMasterCategoriesOk() (*[]MicrosoftGraphOutlookCategory, bool)`

GetMasterCategoriesOk returns a tuple with the MasterCategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterCategories

`func (o *MicrosoftGraphOutlookUser) SetMasterCategories(v []MicrosoftGraphOutlookCategory)`

SetMasterCategories sets MasterCategories field to given value.

### HasMasterCategories

`func (o *MicrosoftGraphOutlookUser) HasMasterCategories() bool`

HasMasterCategories returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


