# SecureScore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

### NewSecureScore

`func NewSecureScore() *SecureScore`

NewSecureScore instantiates a new SecureScore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecureScoreWithDefaults

`func NewSecureScoreWithDefaults() *SecureScore`

NewSecureScoreWithDefaults instantiates a new SecureScore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveUserCount

`func (o *SecureScore) GetActiveUserCount() int32`

GetActiveUserCount returns the ActiveUserCount field if non-nil, zero value otherwise.

### GetActiveUserCountOk

`func (o *SecureScore) GetActiveUserCountOk() (*int32, bool)`

GetActiveUserCountOk returns a tuple with the ActiveUserCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveUserCount

`func (o *SecureScore) SetActiveUserCount(v int32)`

SetActiveUserCount sets ActiveUserCount field to given value.

### HasActiveUserCount

`func (o *SecureScore) HasActiveUserCount() bool`

HasActiveUserCount returns a boolean if a field has been set.

### SetActiveUserCountNil

`func (o *SecureScore) SetActiveUserCountNil(b bool)`

 SetActiveUserCountNil sets the value for ActiveUserCount to be an explicit nil

### UnsetActiveUserCount
`func (o *SecureScore) UnsetActiveUserCount()`

UnsetActiveUserCount ensures that no value is present for ActiveUserCount, not even an explicit nil
### GetAverageComparativeScores

`func (o *SecureScore) GetAverageComparativeScores() []*AnyOfmicrosoftGraphAverageComparativeScore`

GetAverageComparativeScores returns the AverageComparativeScores field if non-nil, zero value otherwise.

### GetAverageComparativeScoresOk

`func (o *SecureScore) GetAverageComparativeScoresOk() (*[]*AnyOfmicrosoftGraphAverageComparativeScore, bool)`

GetAverageComparativeScoresOk returns a tuple with the AverageComparativeScores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageComparativeScores

`func (o *SecureScore) SetAverageComparativeScores(v []*AnyOfmicrosoftGraphAverageComparativeScore)`

SetAverageComparativeScores sets AverageComparativeScores field to given value.

### HasAverageComparativeScores

`func (o *SecureScore) HasAverageComparativeScores() bool`

HasAverageComparativeScores returns a boolean if a field has been set.

### GetAzureTenantId

`func (o *SecureScore) GetAzureTenantId() string`

GetAzureTenantId returns the AzureTenantId field if non-nil, zero value otherwise.

### GetAzureTenantIdOk

`func (o *SecureScore) GetAzureTenantIdOk() (*string, bool)`

GetAzureTenantIdOk returns a tuple with the AzureTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureTenantId

`func (o *SecureScore) SetAzureTenantId(v string)`

SetAzureTenantId sets AzureTenantId field to given value.

### HasAzureTenantId

`func (o *SecureScore) HasAzureTenantId() bool`

HasAzureTenantId returns a boolean if a field has been set.

### GetControlScores

`func (o *SecureScore) GetControlScores() []*AnyOfmicrosoftGraphControlScore`

GetControlScores returns the ControlScores field if non-nil, zero value otherwise.

### GetControlScoresOk

`func (o *SecureScore) GetControlScoresOk() (*[]*AnyOfmicrosoftGraphControlScore, bool)`

GetControlScoresOk returns a tuple with the ControlScores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlScores

`func (o *SecureScore) SetControlScores(v []*AnyOfmicrosoftGraphControlScore)`

SetControlScores sets ControlScores field to given value.

### HasControlScores

`func (o *SecureScore) HasControlScores() bool`

HasControlScores returns a boolean if a field has been set.

### GetCreatedDateTime

`func (o *SecureScore) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *SecureScore) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *SecureScore) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *SecureScore) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTimeNil

`func (o *SecureScore) SetCreatedDateTimeNil(b bool)`

 SetCreatedDateTimeNil sets the value for CreatedDateTime to be an explicit nil

### UnsetCreatedDateTime
`func (o *SecureScore) UnsetCreatedDateTime()`

UnsetCreatedDateTime ensures that no value is present for CreatedDateTime, not even an explicit nil
### GetCurrentScore

`func (o *SecureScore) GetCurrentScore() AnyOfnumberstringstring`

GetCurrentScore returns the CurrentScore field if non-nil, zero value otherwise.

### GetCurrentScoreOk

`func (o *SecureScore) GetCurrentScoreOk() (*AnyOfnumberstringstring, bool)`

GetCurrentScoreOk returns a tuple with the CurrentScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentScore

`func (o *SecureScore) SetCurrentScore(v AnyOfnumberstringstring)`

SetCurrentScore sets CurrentScore field to given value.

### HasCurrentScore

`func (o *SecureScore) HasCurrentScore() bool`

HasCurrentScore returns a boolean if a field has been set.

### SetCurrentScoreNil

`func (o *SecureScore) SetCurrentScoreNil(b bool)`

 SetCurrentScoreNil sets the value for CurrentScore to be an explicit nil

### UnsetCurrentScore
`func (o *SecureScore) UnsetCurrentScore()`

UnsetCurrentScore ensures that no value is present for CurrentScore, not even an explicit nil
### GetEnabledServices

`func (o *SecureScore) GetEnabledServices() []*string`

GetEnabledServices returns the EnabledServices field if non-nil, zero value otherwise.

### GetEnabledServicesOk

`func (o *SecureScore) GetEnabledServicesOk() (*[]*string, bool)`

GetEnabledServicesOk returns a tuple with the EnabledServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledServices

`func (o *SecureScore) SetEnabledServices(v []*string)`

SetEnabledServices sets EnabledServices field to given value.

### HasEnabledServices

`func (o *SecureScore) HasEnabledServices() bool`

HasEnabledServices returns a boolean if a field has been set.

### GetLicensedUserCount

`func (o *SecureScore) GetLicensedUserCount() int32`

GetLicensedUserCount returns the LicensedUserCount field if non-nil, zero value otherwise.

### GetLicensedUserCountOk

`func (o *SecureScore) GetLicensedUserCountOk() (*int32, bool)`

GetLicensedUserCountOk returns a tuple with the LicensedUserCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicensedUserCount

`func (o *SecureScore) SetLicensedUserCount(v int32)`

SetLicensedUserCount sets LicensedUserCount field to given value.

### HasLicensedUserCount

`func (o *SecureScore) HasLicensedUserCount() bool`

HasLicensedUserCount returns a boolean if a field has been set.

### SetLicensedUserCountNil

`func (o *SecureScore) SetLicensedUserCountNil(b bool)`

 SetLicensedUserCountNil sets the value for LicensedUserCount to be an explicit nil

### UnsetLicensedUserCount
`func (o *SecureScore) UnsetLicensedUserCount()`

UnsetLicensedUserCount ensures that no value is present for LicensedUserCount, not even an explicit nil
### GetMaxScore

`func (o *SecureScore) GetMaxScore() AnyOfnumberstringstring`

GetMaxScore returns the MaxScore field if non-nil, zero value otherwise.

### GetMaxScoreOk

`func (o *SecureScore) GetMaxScoreOk() (*AnyOfnumberstringstring, bool)`

GetMaxScoreOk returns a tuple with the MaxScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxScore

`func (o *SecureScore) SetMaxScore(v AnyOfnumberstringstring)`

SetMaxScore sets MaxScore field to given value.

### HasMaxScore

`func (o *SecureScore) HasMaxScore() bool`

HasMaxScore returns a boolean if a field has been set.

### SetMaxScoreNil

`func (o *SecureScore) SetMaxScoreNil(b bool)`

 SetMaxScoreNil sets the value for MaxScore to be an explicit nil

### UnsetMaxScore
`func (o *SecureScore) UnsetMaxScore()`

UnsetMaxScore ensures that no value is present for MaxScore, not even an explicit nil
### GetVendorInformation

`func (o *SecureScore) GetVendorInformation() AnyOfmicrosoftGraphSecurityVendorInformation`

GetVendorInformation returns the VendorInformation field if non-nil, zero value otherwise.

### GetVendorInformationOk

`func (o *SecureScore) GetVendorInformationOk() (*AnyOfmicrosoftGraphSecurityVendorInformation, bool)`

GetVendorInformationOk returns a tuple with the VendorInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorInformation

`func (o *SecureScore) SetVendorInformation(v AnyOfmicrosoftGraphSecurityVendorInformation)`

SetVendorInformation sets VendorInformation field to given value.

### HasVendorInformation

`func (o *SecureScore) HasVendorInformation() bool`

HasVendorInformation returns a boolean if a field has been set.

### SetVendorInformationNil

`func (o *SecureScore) SetVendorInformationNil(b bool)`

 SetVendorInformationNil sets the value for VendorInformation to be an explicit nil

### UnsetVendorInformation
`func (o *SecureScore) UnsetVendorInformation()`

UnsetVendorInformation ensures that no value is present for VendorInformation, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


