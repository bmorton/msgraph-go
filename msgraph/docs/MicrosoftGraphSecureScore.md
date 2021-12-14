# MicrosoftGraphSecureScore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**ActiveUserCount** | Pointer to **NullableInt32** | Active user count of the given tenant. | [optional] 
**AverageComparativeScores** | Pointer to [**[]AnyOfmicrosoftGraphAverageComparativeScore**](AnyOfmicrosoftGraphAverageComparativeScore.md) | Average score by different scopes (for example, average by industry, average by seating) and control category (Identity, Data, Device, Apps, Infrastructure) within the scope. | [optional] 
**AzureTenantId** | Pointer to **string** | GUID string for tenant ID. | [optional] 
**ControlScores** | Pointer to [**[]AnyOfmicrosoftGraphControlScore**](AnyOfmicrosoftGraphControlScore.md) | Contains tenant scores for a set of controls. | [optional] 
**CreatedDateTime** | Pointer to **NullableTime** | The date when the entity is created. | [optional] 
**CurrentScore** | Pointer to [**AnyOfnumberstringstring**](anyOf&lt;number,string,string&gt;.md) | Tenant current attained score on specified date. | [optional] 
**EnabledServices** | Pointer to **[]string** | Microsoft-provided services for the tenant (for example, Exchange online, Skype, Sharepoint). | [optional] 
**LicensedUserCount** | Pointer to **NullableInt32** | Licensed user count of the given tenant. | [optional] 
**MaxScore** | Pointer to [**AnyOfnumberstringstring**](anyOf&lt;number,string,string&gt;.md) | Tenant maximum possible score on specified date. | [optional] 
**VendorInformation** | Pointer to [**AnyOfmicrosoftGraphSecurityVendorInformation**](anyOf&lt;microsoft.graph.securityVendorInformation&gt;.md) | Complex type containing details about the security product/service vendor, provider, and subprovider (for example, vendor&#x3D;Microsoft; provider&#x3D;SecureScore). Required. | [optional] 

## Methods

### NewMicrosoftGraphSecureScore

`func NewMicrosoftGraphSecureScore() *MicrosoftGraphSecureScore`

NewMicrosoftGraphSecureScore instantiates a new MicrosoftGraphSecureScore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphSecureScoreWithDefaults

`func NewMicrosoftGraphSecureScoreWithDefaults() *MicrosoftGraphSecureScore`

NewMicrosoftGraphSecureScoreWithDefaults instantiates a new MicrosoftGraphSecureScore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphSecureScore) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphSecureScore) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphSecureScore) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphSecureScore) HasId() bool`

HasId returns a boolean if a field has been set.

### GetActiveUserCount

`func (o *MicrosoftGraphSecureScore) GetActiveUserCount() int32`

GetActiveUserCount returns the ActiveUserCount field if non-nil, zero value otherwise.

### GetActiveUserCountOk

`func (o *MicrosoftGraphSecureScore) GetActiveUserCountOk() (*int32, bool)`

GetActiveUserCountOk returns a tuple with the ActiveUserCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveUserCount

`func (o *MicrosoftGraphSecureScore) SetActiveUserCount(v int32)`

SetActiveUserCount sets ActiveUserCount field to given value.

### HasActiveUserCount

`func (o *MicrosoftGraphSecureScore) HasActiveUserCount() bool`

HasActiveUserCount returns a boolean if a field has been set.

### SetActiveUserCountNil

`func (o *MicrosoftGraphSecureScore) SetActiveUserCountNil(b bool)`

 SetActiveUserCountNil sets the value for ActiveUserCount to be an explicit nil

### UnsetActiveUserCount
`func (o *MicrosoftGraphSecureScore) UnsetActiveUserCount()`

UnsetActiveUserCount ensures that no value is present for ActiveUserCount, not even an explicit nil
### GetAverageComparativeScores

`func (o *MicrosoftGraphSecureScore) GetAverageComparativeScores() []*AnyOfmicrosoftGraphAverageComparativeScore`

GetAverageComparativeScores returns the AverageComparativeScores field if non-nil, zero value otherwise.

### GetAverageComparativeScoresOk

`func (o *MicrosoftGraphSecureScore) GetAverageComparativeScoresOk() (*[]*AnyOfmicrosoftGraphAverageComparativeScore, bool)`

GetAverageComparativeScoresOk returns a tuple with the AverageComparativeScores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageComparativeScores

`func (o *MicrosoftGraphSecureScore) SetAverageComparativeScores(v []*AnyOfmicrosoftGraphAverageComparativeScore)`

SetAverageComparativeScores sets AverageComparativeScores field to given value.

### HasAverageComparativeScores

`func (o *MicrosoftGraphSecureScore) HasAverageComparativeScores() bool`

HasAverageComparativeScores returns a boolean if a field has been set.

### GetAzureTenantId

`func (o *MicrosoftGraphSecureScore) GetAzureTenantId() string`

GetAzureTenantId returns the AzureTenantId field if non-nil, zero value otherwise.

### GetAzureTenantIdOk

`func (o *MicrosoftGraphSecureScore) GetAzureTenantIdOk() (*string, bool)`

GetAzureTenantIdOk returns a tuple with the AzureTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureTenantId

`func (o *MicrosoftGraphSecureScore) SetAzureTenantId(v string)`

SetAzureTenantId sets AzureTenantId field to given value.

### HasAzureTenantId

`func (o *MicrosoftGraphSecureScore) HasAzureTenantId() bool`

HasAzureTenantId returns a boolean if a field has been set.

### GetControlScores

`func (o *MicrosoftGraphSecureScore) GetControlScores() []*AnyOfmicrosoftGraphControlScore`

GetControlScores returns the ControlScores field if non-nil, zero value otherwise.

### GetControlScoresOk

`func (o *MicrosoftGraphSecureScore) GetControlScoresOk() (*[]*AnyOfmicrosoftGraphControlScore, bool)`

GetControlScoresOk returns a tuple with the ControlScores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlScores

`func (o *MicrosoftGraphSecureScore) SetControlScores(v []*AnyOfmicrosoftGraphControlScore)`

SetControlScores sets ControlScores field to given value.

### HasControlScores

`func (o *MicrosoftGraphSecureScore) HasControlScores() bool`

HasControlScores returns a boolean if a field has been set.

### GetCreatedDateTime

`func (o *MicrosoftGraphSecureScore) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphSecureScore) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphSecureScore) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *MicrosoftGraphSecureScore) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTimeNil

`func (o *MicrosoftGraphSecureScore) SetCreatedDateTimeNil(b bool)`

 SetCreatedDateTimeNil sets the value for CreatedDateTime to be an explicit nil

### UnsetCreatedDateTime
`func (o *MicrosoftGraphSecureScore) UnsetCreatedDateTime()`

UnsetCreatedDateTime ensures that no value is present for CreatedDateTime, not even an explicit nil
### GetCurrentScore

`func (o *MicrosoftGraphSecureScore) GetCurrentScore() AnyOfnumberstringstring`

GetCurrentScore returns the CurrentScore field if non-nil, zero value otherwise.

### GetCurrentScoreOk

`func (o *MicrosoftGraphSecureScore) GetCurrentScoreOk() (*AnyOfnumberstringstring, bool)`

GetCurrentScoreOk returns a tuple with the CurrentScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentScore

`func (o *MicrosoftGraphSecureScore) SetCurrentScore(v AnyOfnumberstringstring)`

SetCurrentScore sets CurrentScore field to given value.

### HasCurrentScore

`func (o *MicrosoftGraphSecureScore) HasCurrentScore() bool`

HasCurrentScore returns a boolean if a field has been set.

### SetCurrentScoreNil

`func (o *MicrosoftGraphSecureScore) SetCurrentScoreNil(b bool)`

 SetCurrentScoreNil sets the value for CurrentScore to be an explicit nil

### UnsetCurrentScore
`func (o *MicrosoftGraphSecureScore) UnsetCurrentScore()`

UnsetCurrentScore ensures that no value is present for CurrentScore, not even an explicit nil
### GetEnabledServices

`func (o *MicrosoftGraphSecureScore) GetEnabledServices() []*string`

GetEnabledServices returns the EnabledServices field if non-nil, zero value otherwise.

### GetEnabledServicesOk

`func (o *MicrosoftGraphSecureScore) GetEnabledServicesOk() (*[]*string, bool)`

GetEnabledServicesOk returns a tuple with the EnabledServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledServices

`func (o *MicrosoftGraphSecureScore) SetEnabledServices(v []*string)`

SetEnabledServices sets EnabledServices field to given value.

### HasEnabledServices

`func (o *MicrosoftGraphSecureScore) HasEnabledServices() bool`

HasEnabledServices returns a boolean if a field has been set.

### GetLicensedUserCount

`func (o *MicrosoftGraphSecureScore) GetLicensedUserCount() int32`

GetLicensedUserCount returns the LicensedUserCount field if non-nil, zero value otherwise.

### GetLicensedUserCountOk

`func (o *MicrosoftGraphSecureScore) GetLicensedUserCountOk() (*int32, bool)`

GetLicensedUserCountOk returns a tuple with the LicensedUserCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicensedUserCount

`func (o *MicrosoftGraphSecureScore) SetLicensedUserCount(v int32)`

SetLicensedUserCount sets LicensedUserCount field to given value.

### HasLicensedUserCount

`func (o *MicrosoftGraphSecureScore) HasLicensedUserCount() bool`

HasLicensedUserCount returns a boolean if a field has been set.

### SetLicensedUserCountNil

`func (o *MicrosoftGraphSecureScore) SetLicensedUserCountNil(b bool)`

 SetLicensedUserCountNil sets the value for LicensedUserCount to be an explicit nil

### UnsetLicensedUserCount
`func (o *MicrosoftGraphSecureScore) UnsetLicensedUserCount()`

UnsetLicensedUserCount ensures that no value is present for LicensedUserCount, not even an explicit nil
### GetMaxScore

`func (o *MicrosoftGraphSecureScore) GetMaxScore() AnyOfnumberstringstring`

GetMaxScore returns the MaxScore field if non-nil, zero value otherwise.

### GetMaxScoreOk

`func (o *MicrosoftGraphSecureScore) GetMaxScoreOk() (*AnyOfnumberstringstring, bool)`

GetMaxScoreOk returns a tuple with the MaxScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxScore

`func (o *MicrosoftGraphSecureScore) SetMaxScore(v AnyOfnumberstringstring)`

SetMaxScore sets MaxScore field to given value.

### HasMaxScore

`func (o *MicrosoftGraphSecureScore) HasMaxScore() bool`

HasMaxScore returns a boolean if a field has been set.

### SetMaxScoreNil

`func (o *MicrosoftGraphSecureScore) SetMaxScoreNil(b bool)`

 SetMaxScoreNil sets the value for MaxScore to be an explicit nil

### UnsetMaxScore
`func (o *MicrosoftGraphSecureScore) UnsetMaxScore()`

UnsetMaxScore ensures that no value is present for MaxScore, not even an explicit nil
### GetVendorInformation

`func (o *MicrosoftGraphSecureScore) GetVendorInformation() AnyOfmicrosoftGraphSecurityVendorInformation`

GetVendorInformation returns the VendorInformation field if non-nil, zero value otherwise.

### GetVendorInformationOk

`func (o *MicrosoftGraphSecureScore) GetVendorInformationOk() (*AnyOfmicrosoftGraphSecurityVendorInformation, bool)`

GetVendorInformationOk returns a tuple with the VendorInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorInformation

`func (o *MicrosoftGraphSecureScore) SetVendorInformation(v AnyOfmicrosoftGraphSecurityVendorInformation)`

SetVendorInformation sets VendorInformation field to given value.

### HasVendorInformation

`func (o *MicrosoftGraphSecureScore) HasVendorInformation() bool`

HasVendorInformation returns a boolean if a field has been set.

### SetVendorInformationNil

`func (o *MicrosoftGraphSecureScore) SetVendorInformationNil(b bool)`

 SetVendorInformationNil sets the value for VendorInformation to be an explicit nil

### UnsetVendorInformation
`func (o *MicrosoftGraphSecureScore) UnsetVendorInformation()`

UnsetVendorInformation ensures that no value is present for VendorInformation, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


