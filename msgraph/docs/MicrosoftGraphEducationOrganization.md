# MicrosoftGraphEducationOrganization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**Description** | Pointer to **NullableString** | Organization description. | [optional] 
**DisplayName** | Pointer to **string** | Organization display name. | [optional] 
**ExternalSource** | Pointer to [**AnyOfmicrosoftGraphEducationExternalSource**](anyOf&lt;microsoft.graph.educationExternalSource&gt;.md) | Source where this organization was created from. Possible values are: sis, manual. | [optional] 
**ExternalSourceDetail** | Pointer to **NullableString** | The name of the external source this resources was generated from. | [optional] 

## Methods

### NewMicrosoftGraphEducationOrganization

`func NewMicrosoftGraphEducationOrganization() *MicrosoftGraphEducationOrganization`

NewMicrosoftGraphEducationOrganization instantiates a new MicrosoftGraphEducationOrganization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphEducationOrganizationWithDefaults

`func NewMicrosoftGraphEducationOrganizationWithDefaults() *MicrosoftGraphEducationOrganization`

NewMicrosoftGraphEducationOrganizationWithDefaults instantiates a new MicrosoftGraphEducationOrganization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphEducationOrganization) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphEducationOrganization) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphEducationOrganization) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphEducationOrganization) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDescription

`func (o *MicrosoftGraphEducationOrganization) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MicrosoftGraphEducationOrganization) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MicrosoftGraphEducationOrganization) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MicrosoftGraphEducationOrganization) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *MicrosoftGraphEducationOrganization) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *MicrosoftGraphEducationOrganization) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDisplayName

`func (o *MicrosoftGraphEducationOrganization) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphEducationOrganization) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MicrosoftGraphEducationOrganization) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MicrosoftGraphEducationOrganization) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetExternalSource

`func (o *MicrosoftGraphEducationOrganization) GetExternalSource() AnyOfmicrosoftGraphEducationExternalSource`

GetExternalSource returns the ExternalSource field if non-nil, zero value otherwise.

### GetExternalSourceOk

`func (o *MicrosoftGraphEducationOrganization) GetExternalSourceOk() (*AnyOfmicrosoftGraphEducationExternalSource, bool)`

GetExternalSourceOk returns a tuple with the ExternalSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalSource

`func (o *MicrosoftGraphEducationOrganization) SetExternalSource(v AnyOfmicrosoftGraphEducationExternalSource)`

SetExternalSource sets ExternalSource field to given value.

### HasExternalSource

`func (o *MicrosoftGraphEducationOrganization) HasExternalSource() bool`

HasExternalSource returns a boolean if a field has been set.

### SetExternalSourceNil

`func (o *MicrosoftGraphEducationOrganization) SetExternalSourceNil(b bool)`

 SetExternalSourceNil sets the value for ExternalSource to be an explicit nil

### UnsetExternalSource
`func (o *MicrosoftGraphEducationOrganization) UnsetExternalSource()`

UnsetExternalSource ensures that no value is present for ExternalSource, not even an explicit nil
### GetExternalSourceDetail

`func (o *MicrosoftGraphEducationOrganization) GetExternalSourceDetail() string`

GetExternalSourceDetail returns the ExternalSourceDetail field if non-nil, zero value otherwise.

### GetExternalSourceDetailOk

`func (o *MicrosoftGraphEducationOrganization) GetExternalSourceDetailOk() (*string, bool)`

GetExternalSourceDetailOk returns a tuple with the ExternalSourceDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalSourceDetail

`func (o *MicrosoftGraphEducationOrganization) SetExternalSourceDetail(v string)`

SetExternalSourceDetail sets ExternalSourceDetail field to given value.

### HasExternalSourceDetail

`func (o *MicrosoftGraphEducationOrganization) HasExternalSourceDetail() bool`

HasExternalSourceDetail returns a boolean if a field has been set.

### SetExternalSourceDetailNil

`func (o *MicrosoftGraphEducationOrganization) SetExternalSourceDetailNil(b bool)`

 SetExternalSourceDetailNil sets the value for ExternalSourceDetail to be an explicit nil

### UnsetExternalSourceDetail
`func (o *MicrosoftGraphEducationOrganization) UnsetExternalSourceDetail()`

UnsetExternalSourceDetail ensures that no value is present for ExternalSourceDetail, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


