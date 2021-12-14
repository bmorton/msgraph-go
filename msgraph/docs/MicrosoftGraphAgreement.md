# MicrosoftGraphAgreement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**DisplayName** | Pointer to **NullableString** | Display name of the agreement. The display name is used for internal tracking of the agreement but is not shown to end users who view the agreement. | [optional] 
**IsPerDeviceAcceptanceRequired** | Pointer to **NullableBool** | Indicates whether end users are required to accept this agreement on every device that they access it from. The end user is required to register their device in Azure AD, if they haven&#39;t already done so. | [optional] 
**IsViewingBeforeAcceptanceRequired** | Pointer to **NullableBool** | Indicates whether the user has to expand the agreement before accepting. | [optional] 
**TermsExpiration** | Pointer to [**AnyOfmicrosoftGraphTermsExpiration**](anyOf&lt;microsoft.graph.termsExpiration&gt;.md) | Expiration schedule and frequency of agreement for all users. | [optional] 
**UserReacceptRequiredFrequency** | Pointer to **NullableString** | The duration after which the user must re-accept the terms of use. The value is represented in ISO 8601 format for durations. | [optional] 
**Acceptances** | Pointer to [**[]MicrosoftGraphAgreementAcceptance**](MicrosoftGraphAgreementAcceptance.md) | Read-only. Information about acceptances of this agreement. | [optional] 
**File** | Pointer to [**AnyOfmicrosoftGraphAgreementFile**](anyOf&lt;microsoft.graph.agreementFile&gt;.md) | Default PDF linked to this agreement. | [optional] 
**Files** | Pointer to [**[]MicrosoftGraphAgreementFileLocalization**](MicrosoftGraphAgreementFileLocalization.md) | PDFs linked to this agreement. This property is in the process of being deprecated. Use the  file property instead. | [optional] 

## Methods

### NewMicrosoftGraphAgreement

`func NewMicrosoftGraphAgreement() *MicrosoftGraphAgreement`

NewMicrosoftGraphAgreement instantiates a new MicrosoftGraphAgreement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphAgreementWithDefaults

`func NewMicrosoftGraphAgreementWithDefaults() *MicrosoftGraphAgreement`

NewMicrosoftGraphAgreementWithDefaults instantiates a new MicrosoftGraphAgreement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphAgreement) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphAgreement) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphAgreement) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphAgreement) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDisplayName

`func (o *MicrosoftGraphAgreement) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphAgreement) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MicrosoftGraphAgreement) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MicrosoftGraphAgreement) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *MicrosoftGraphAgreement) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *MicrosoftGraphAgreement) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetIsPerDeviceAcceptanceRequired

`func (o *MicrosoftGraphAgreement) GetIsPerDeviceAcceptanceRequired() bool`

GetIsPerDeviceAcceptanceRequired returns the IsPerDeviceAcceptanceRequired field if non-nil, zero value otherwise.

### GetIsPerDeviceAcceptanceRequiredOk

`func (o *MicrosoftGraphAgreement) GetIsPerDeviceAcceptanceRequiredOk() (*bool, bool)`

GetIsPerDeviceAcceptanceRequiredOk returns a tuple with the IsPerDeviceAcceptanceRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPerDeviceAcceptanceRequired

`func (o *MicrosoftGraphAgreement) SetIsPerDeviceAcceptanceRequired(v bool)`

SetIsPerDeviceAcceptanceRequired sets IsPerDeviceAcceptanceRequired field to given value.

### HasIsPerDeviceAcceptanceRequired

`func (o *MicrosoftGraphAgreement) HasIsPerDeviceAcceptanceRequired() bool`

HasIsPerDeviceAcceptanceRequired returns a boolean if a field has been set.

### SetIsPerDeviceAcceptanceRequiredNil

`func (o *MicrosoftGraphAgreement) SetIsPerDeviceAcceptanceRequiredNil(b bool)`

 SetIsPerDeviceAcceptanceRequiredNil sets the value for IsPerDeviceAcceptanceRequired to be an explicit nil

### UnsetIsPerDeviceAcceptanceRequired
`func (o *MicrosoftGraphAgreement) UnsetIsPerDeviceAcceptanceRequired()`

UnsetIsPerDeviceAcceptanceRequired ensures that no value is present for IsPerDeviceAcceptanceRequired, not even an explicit nil
### GetIsViewingBeforeAcceptanceRequired

`func (o *MicrosoftGraphAgreement) GetIsViewingBeforeAcceptanceRequired() bool`

GetIsViewingBeforeAcceptanceRequired returns the IsViewingBeforeAcceptanceRequired field if non-nil, zero value otherwise.

### GetIsViewingBeforeAcceptanceRequiredOk

`func (o *MicrosoftGraphAgreement) GetIsViewingBeforeAcceptanceRequiredOk() (*bool, bool)`

GetIsViewingBeforeAcceptanceRequiredOk returns a tuple with the IsViewingBeforeAcceptanceRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsViewingBeforeAcceptanceRequired

`func (o *MicrosoftGraphAgreement) SetIsViewingBeforeAcceptanceRequired(v bool)`

SetIsViewingBeforeAcceptanceRequired sets IsViewingBeforeAcceptanceRequired field to given value.

### HasIsViewingBeforeAcceptanceRequired

`func (o *MicrosoftGraphAgreement) HasIsViewingBeforeAcceptanceRequired() bool`

HasIsViewingBeforeAcceptanceRequired returns a boolean if a field has been set.

### SetIsViewingBeforeAcceptanceRequiredNil

`func (o *MicrosoftGraphAgreement) SetIsViewingBeforeAcceptanceRequiredNil(b bool)`

 SetIsViewingBeforeAcceptanceRequiredNil sets the value for IsViewingBeforeAcceptanceRequired to be an explicit nil

### UnsetIsViewingBeforeAcceptanceRequired
`func (o *MicrosoftGraphAgreement) UnsetIsViewingBeforeAcceptanceRequired()`

UnsetIsViewingBeforeAcceptanceRequired ensures that no value is present for IsViewingBeforeAcceptanceRequired, not even an explicit nil
### GetTermsExpiration

`func (o *MicrosoftGraphAgreement) GetTermsExpiration() AnyOfmicrosoftGraphTermsExpiration`

GetTermsExpiration returns the TermsExpiration field if non-nil, zero value otherwise.

### GetTermsExpirationOk

`func (o *MicrosoftGraphAgreement) GetTermsExpirationOk() (*AnyOfmicrosoftGraphTermsExpiration, bool)`

GetTermsExpirationOk returns a tuple with the TermsExpiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermsExpiration

`func (o *MicrosoftGraphAgreement) SetTermsExpiration(v AnyOfmicrosoftGraphTermsExpiration)`

SetTermsExpiration sets TermsExpiration field to given value.

### HasTermsExpiration

`func (o *MicrosoftGraphAgreement) HasTermsExpiration() bool`

HasTermsExpiration returns a boolean if a field has been set.

### SetTermsExpirationNil

`func (o *MicrosoftGraphAgreement) SetTermsExpirationNil(b bool)`

 SetTermsExpirationNil sets the value for TermsExpiration to be an explicit nil

### UnsetTermsExpiration
`func (o *MicrosoftGraphAgreement) UnsetTermsExpiration()`

UnsetTermsExpiration ensures that no value is present for TermsExpiration, not even an explicit nil
### GetUserReacceptRequiredFrequency

`func (o *MicrosoftGraphAgreement) GetUserReacceptRequiredFrequency() string`

GetUserReacceptRequiredFrequency returns the UserReacceptRequiredFrequency field if non-nil, zero value otherwise.

### GetUserReacceptRequiredFrequencyOk

`func (o *MicrosoftGraphAgreement) GetUserReacceptRequiredFrequencyOk() (*string, bool)`

GetUserReacceptRequiredFrequencyOk returns a tuple with the UserReacceptRequiredFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserReacceptRequiredFrequency

`func (o *MicrosoftGraphAgreement) SetUserReacceptRequiredFrequency(v string)`

SetUserReacceptRequiredFrequency sets UserReacceptRequiredFrequency field to given value.

### HasUserReacceptRequiredFrequency

`func (o *MicrosoftGraphAgreement) HasUserReacceptRequiredFrequency() bool`

HasUserReacceptRequiredFrequency returns a boolean if a field has been set.

### SetUserReacceptRequiredFrequencyNil

`func (o *MicrosoftGraphAgreement) SetUserReacceptRequiredFrequencyNil(b bool)`

 SetUserReacceptRequiredFrequencyNil sets the value for UserReacceptRequiredFrequency to be an explicit nil

### UnsetUserReacceptRequiredFrequency
`func (o *MicrosoftGraphAgreement) UnsetUserReacceptRequiredFrequency()`

UnsetUserReacceptRequiredFrequency ensures that no value is present for UserReacceptRequiredFrequency, not even an explicit nil
### GetAcceptances

`func (o *MicrosoftGraphAgreement) GetAcceptances() []MicrosoftGraphAgreementAcceptance`

GetAcceptances returns the Acceptances field if non-nil, zero value otherwise.

### GetAcceptancesOk

`func (o *MicrosoftGraphAgreement) GetAcceptancesOk() (*[]MicrosoftGraphAgreementAcceptance, bool)`

GetAcceptancesOk returns a tuple with the Acceptances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptances

`func (o *MicrosoftGraphAgreement) SetAcceptances(v []MicrosoftGraphAgreementAcceptance)`

SetAcceptances sets Acceptances field to given value.

### HasAcceptances

`func (o *MicrosoftGraphAgreement) HasAcceptances() bool`

HasAcceptances returns a boolean if a field has been set.

### GetFile

`func (o *MicrosoftGraphAgreement) GetFile() AnyOfmicrosoftGraphAgreementFile`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *MicrosoftGraphAgreement) GetFileOk() (*AnyOfmicrosoftGraphAgreementFile, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *MicrosoftGraphAgreement) SetFile(v AnyOfmicrosoftGraphAgreementFile)`

SetFile sets File field to given value.

### HasFile

`func (o *MicrosoftGraphAgreement) HasFile() bool`

HasFile returns a boolean if a field has been set.

### SetFileNil

`func (o *MicrosoftGraphAgreement) SetFileNil(b bool)`

 SetFileNil sets the value for File to be an explicit nil

### UnsetFile
`func (o *MicrosoftGraphAgreement) UnsetFile()`

UnsetFile ensures that no value is present for File, not even an explicit nil
### GetFiles

`func (o *MicrosoftGraphAgreement) GetFiles() []MicrosoftGraphAgreementFileLocalization`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *MicrosoftGraphAgreement) GetFilesOk() (*[]MicrosoftGraphAgreementFileLocalization, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *MicrosoftGraphAgreement) SetFiles(v []MicrosoftGraphAgreementFileLocalization)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *MicrosoftGraphAgreement) HasFiles() bool`

HasFiles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


