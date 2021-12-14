# SecureScoreControlProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionType** | Pointer to **NullableString** | Control action type (Config, Review, Behavior). | [optional] 
**ActionUrl** | Pointer to **NullableString** | URL to where the control can be actioned. | [optional] 
**AzureTenantId** | Pointer to **string** | GUID string for tenant ID. | [optional] 
**ComplianceInformation** | Pointer to [**[]AnyOfmicrosoftGraphComplianceInformation**](AnyOfmicrosoftGraphComplianceInformation.md) | The collection of compliance information associated with secure score control | [optional] 
**ControlCategory** | Pointer to **NullableString** | Control action category (Identity, Data, Device, Apps, Infrastructure). | [optional] 
**ControlStateUpdates** | Pointer to [**[]AnyOfmicrosoftGraphSecureScoreControlStateUpdate**](AnyOfmicrosoftGraphSecureScoreControlStateUpdate.md) | Flag to indicate where the tenant has marked a control (ignore, thirdParty, reviewed) (supports update). | [optional] 
**Deprecated** | Pointer to **NullableBool** | Flag to indicate if a control is depreciated. | [optional] 
**ImplementationCost** | Pointer to **NullableString** | Resource cost of implemmentating control (low, moderate, high). | [optional] 
**LastModifiedDateTime** | Pointer to **NullableTime** | Time at which the control profile entity was last modified. The Timestamp type represents date and time | [optional] 
**MaxScore** | Pointer to [**AnyOfnumberstringstring**](anyOf&lt;number,string,string&gt;.md) | max attainable score for the control. | [optional] 
**Rank** | Pointer to **NullableInt32** | Microsoft&#39;s stack ranking of control. | [optional] 
**Remediation** | Pointer to **NullableString** | Description of what the control will help remediate. | [optional] 
**RemediationImpact** | Pointer to **NullableString** | Description of the impact on users of the remediation. | [optional] 
**Service** | Pointer to **NullableString** | Service that owns the control (Exchange, Sharepoint, Azure AD). | [optional] 
**Threats** | Pointer to **[]string** | List of threats the control mitigates (accountBreach,dataDeletion,dataExfiltration,dataSpillage, | [optional] 
**Tier** | Pointer to **NullableString** | Control tier (Core, Defense in Depth, Advanced.) | [optional] 
**Title** | Pointer to **NullableString** | Title of the control. | [optional] 
**UserImpact** | Pointer to **NullableString** | User impact of implementing control (low, moderate, high). | [optional] 
**VendorInformation** | Pointer to [**AnyOfmicrosoftGraphSecurityVendorInformation**](anyOf&lt;microsoft.graph.securityVendorInformation&gt;.md) |  | [optional] 

## Methods

### NewSecureScoreControlProfile

`func NewSecureScoreControlProfile() *SecureScoreControlProfile`

NewSecureScoreControlProfile instantiates a new SecureScoreControlProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecureScoreControlProfileWithDefaults

`func NewSecureScoreControlProfileWithDefaults() *SecureScoreControlProfile`

NewSecureScoreControlProfileWithDefaults instantiates a new SecureScoreControlProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionType

`func (o *SecureScoreControlProfile) GetActionType() string`

GetActionType returns the ActionType field if non-nil, zero value otherwise.

### GetActionTypeOk

`func (o *SecureScoreControlProfile) GetActionTypeOk() (*string, bool)`

GetActionTypeOk returns a tuple with the ActionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionType

`func (o *SecureScoreControlProfile) SetActionType(v string)`

SetActionType sets ActionType field to given value.

### HasActionType

`func (o *SecureScoreControlProfile) HasActionType() bool`

HasActionType returns a boolean if a field has been set.

### SetActionTypeNil

`func (o *SecureScoreControlProfile) SetActionTypeNil(b bool)`

 SetActionTypeNil sets the value for ActionType to be an explicit nil

### UnsetActionType
`func (o *SecureScoreControlProfile) UnsetActionType()`

UnsetActionType ensures that no value is present for ActionType, not even an explicit nil
### GetActionUrl

`func (o *SecureScoreControlProfile) GetActionUrl() string`

GetActionUrl returns the ActionUrl field if non-nil, zero value otherwise.

### GetActionUrlOk

`func (o *SecureScoreControlProfile) GetActionUrlOk() (*string, bool)`

GetActionUrlOk returns a tuple with the ActionUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionUrl

`func (o *SecureScoreControlProfile) SetActionUrl(v string)`

SetActionUrl sets ActionUrl field to given value.

### HasActionUrl

`func (o *SecureScoreControlProfile) HasActionUrl() bool`

HasActionUrl returns a boolean if a field has been set.

### SetActionUrlNil

`func (o *SecureScoreControlProfile) SetActionUrlNil(b bool)`

 SetActionUrlNil sets the value for ActionUrl to be an explicit nil

### UnsetActionUrl
`func (o *SecureScoreControlProfile) UnsetActionUrl()`

UnsetActionUrl ensures that no value is present for ActionUrl, not even an explicit nil
### GetAzureTenantId

`func (o *SecureScoreControlProfile) GetAzureTenantId() string`

GetAzureTenantId returns the AzureTenantId field if non-nil, zero value otherwise.

### GetAzureTenantIdOk

`func (o *SecureScoreControlProfile) GetAzureTenantIdOk() (*string, bool)`

GetAzureTenantIdOk returns a tuple with the AzureTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureTenantId

`func (o *SecureScoreControlProfile) SetAzureTenantId(v string)`

SetAzureTenantId sets AzureTenantId field to given value.

### HasAzureTenantId

`func (o *SecureScoreControlProfile) HasAzureTenantId() bool`

HasAzureTenantId returns a boolean if a field has been set.

### GetComplianceInformation

`func (o *SecureScoreControlProfile) GetComplianceInformation() []*AnyOfmicrosoftGraphComplianceInformation`

GetComplianceInformation returns the ComplianceInformation field if non-nil, zero value otherwise.

### GetComplianceInformationOk

`func (o *SecureScoreControlProfile) GetComplianceInformationOk() (*[]*AnyOfmicrosoftGraphComplianceInformation, bool)`

GetComplianceInformationOk returns a tuple with the ComplianceInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplianceInformation

`func (o *SecureScoreControlProfile) SetComplianceInformation(v []*AnyOfmicrosoftGraphComplianceInformation)`

SetComplianceInformation sets ComplianceInformation field to given value.

### HasComplianceInformation

`func (o *SecureScoreControlProfile) HasComplianceInformation() bool`

HasComplianceInformation returns a boolean if a field has been set.

### GetControlCategory

`func (o *SecureScoreControlProfile) GetControlCategory() string`

GetControlCategory returns the ControlCategory field if non-nil, zero value otherwise.

### GetControlCategoryOk

`func (o *SecureScoreControlProfile) GetControlCategoryOk() (*string, bool)`

GetControlCategoryOk returns a tuple with the ControlCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCategory

`func (o *SecureScoreControlProfile) SetControlCategory(v string)`

SetControlCategory sets ControlCategory field to given value.

### HasControlCategory

`func (o *SecureScoreControlProfile) HasControlCategory() bool`

HasControlCategory returns a boolean if a field has been set.

### SetControlCategoryNil

`func (o *SecureScoreControlProfile) SetControlCategoryNil(b bool)`

 SetControlCategoryNil sets the value for ControlCategory to be an explicit nil

### UnsetControlCategory
`func (o *SecureScoreControlProfile) UnsetControlCategory()`

UnsetControlCategory ensures that no value is present for ControlCategory, not even an explicit nil
### GetControlStateUpdates

`func (o *SecureScoreControlProfile) GetControlStateUpdates() []*AnyOfmicrosoftGraphSecureScoreControlStateUpdate`

GetControlStateUpdates returns the ControlStateUpdates field if non-nil, zero value otherwise.

### GetControlStateUpdatesOk

`func (o *SecureScoreControlProfile) GetControlStateUpdatesOk() (*[]*AnyOfmicrosoftGraphSecureScoreControlStateUpdate, bool)`

GetControlStateUpdatesOk returns a tuple with the ControlStateUpdates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlStateUpdates

`func (o *SecureScoreControlProfile) SetControlStateUpdates(v []*AnyOfmicrosoftGraphSecureScoreControlStateUpdate)`

SetControlStateUpdates sets ControlStateUpdates field to given value.

### HasControlStateUpdates

`func (o *SecureScoreControlProfile) HasControlStateUpdates() bool`

HasControlStateUpdates returns a boolean if a field has been set.

### GetDeprecated

`func (o *SecureScoreControlProfile) GetDeprecated() bool`

GetDeprecated returns the Deprecated field if non-nil, zero value otherwise.

### GetDeprecatedOk

`func (o *SecureScoreControlProfile) GetDeprecatedOk() (*bool, bool)`

GetDeprecatedOk returns a tuple with the Deprecated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecated

`func (o *SecureScoreControlProfile) SetDeprecated(v bool)`

SetDeprecated sets Deprecated field to given value.

### HasDeprecated

`func (o *SecureScoreControlProfile) HasDeprecated() bool`

HasDeprecated returns a boolean if a field has been set.

### SetDeprecatedNil

`func (o *SecureScoreControlProfile) SetDeprecatedNil(b bool)`

 SetDeprecatedNil sets the value for Deprecated to be an explicit nil

### UnsetDeprecated
`func (o *SecureScoreControlProfile) UnsetDeprecated()`

UnsetDeprecated ensures that no value is present for Deprecated, not even an explicit nil
### GetImplementationCost

`func (o *SecureScoreControlProfile) GetImplementationCost() string`

GetImplementationCost returns the ImplementationCost field if non-nil, zero value otherwise.

### GetImplementationCostOk

`func (o *SecureScoreControlProfile) GetImplementationCostOk() (*string, bool)`

GetImplementationCostOk returns a tuple with the ImplementationCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImplementationCost

`func (o *SecureScoreControlProfile) SetImplementationCost(v string)`

SetImplementationCost sets ImplementationCost field to given value.

### HasImplementationCost

`func (o *SecureScoreControlProfile) HasImplementationCost() bool`

HasImplementationCost returns a boolean if a field has been set.

### SetImplementationCostNil

`func (o *SecureScoreControlProfile) SetImplementationCostNil(b bool)`

 SetImplementationCostNil sets the value for ImplementationCost to be an explicit nil

### UnsetImplementationCost
`func (o *SecureScoreControlProfile) UnsetImplementationCost()`

UnsetImplementationCost ensures that no value is present for ImplementationCost, not even an explicit nil
### GetLastModifiedDateTime

`func (o *SecureScoreControlProfile) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *SecureScoreControlProfile) GetLastModifiedDateTimeOk() (*time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedDateTime

`func (o *SecureScoreControlProfile) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime sets LastModifiedDateTime field to given value.

### HasLastModifiedDateTime

`func (o *SecureScoreControlProfile) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### SetLastModifiedDateTimeNil

`func (o *SecureScoreControlProfile) SetLastModifiedDateTimeNil(b bool)`

 SetLastModifiedDateTimeNil sets the value for LastModifiedDateTime to be an explicit nil

### UnsetLastModifiedDateTime
`func (o *SecureScoreControlProfile) UnsetLastModifiedDateTime()`

UnsetLastModifiedDateTime ensures that no value is present for LastModifiedDateTime, not even an explicit nil
### GetMaxScore

`func (o *SecureScoreControlProfile) GetMaxScore() AnyOfnumberstringstring`

GetMaxScore returns the MaxScore field if non-nil, zero value otherwise.

### GetMaxScoreOk

`func (o *SecureScoreControlProfile) GetMaxScoreOk() (*AnyOfnumberstringstring, bool)`

GetMaxScoreOk returns a tuple with the MaxScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxScore

`func (o *SecureScoreControlProfile) SetMaxScore(v AnyOfnumberstringstring)`

SetMaxScore sets MaxScore field to given value.

### HasMaxScore

`func (o *SecureScoreControlProfile) HasMaxScore() bool`

HasMaxScore returns a boolean if a field has been set.

### SetMaxScoreNil

`func (o *SecureScoreControlProfile) SetMaxScoreNil(b bool)`

 SetMaxScoreNil sets the value for MaxScore to be an explicit nil

### UnsetMaxScore
`func (o *SecureScoreControlProfile) UnsetMaxScore()`

UnsetMaxScore ensures that no value is present for MaxScore, not even an explicit nil
### GetRank

`func (o *SecureScoreControlProfile) GetRank() int32`

GetRank returns the Rank field if non-nil, zero value otherwise.

### GetRankOk

`func (o *SecureScoreControlProfile) GetRankOk() (*int32, bool)`

GetRankOk returns a tuple with the Rank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRank

`func (o *SecureScoreControlProfile) SetRank(v int32)`

SetRank sets Rank field to given value.

### HasRank

`func (o *SecureScoreControlProfile) HasRank() bool`

HasRank returns a boolean if a field has been set.

### SetRankNil

`func (o *SecureScoreControlProfile) SetRankNil(b bool)`

 SetRankNil sets the value for Rank to be an explicit nil

### UnsetRank
`func (o *SecureScoreControlProfile) UnsetRank()`

UnsetRank ensures that no value is present for Rank, not even an explicit nil
### GetRemediation

`func (o *SecureScoreControlProfile) GetRemediation() string`

GetRemediation returns the Remediation field if non-nil, zero value otherwise.

### GetRemediationOk

`func (o *SecureScoreControlProfile) GetRemediationOk() (*string, bool)`

GetRemediationOk returns a tuple with the Remediation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemediation

`func (o *SecureScoreControlProfile) SetRemediation(v string)`

SetRemediation sets Remediation field to given value.

### HasRemediation

`func (o *SecureScoreControlProfile) HasRemediation() bool`

HasRemediation returns a boolean if a field has been set.

### SetRemediationNil

`func (o *SecureScoreControlProfile) SetRemediationNil(b bool)`

 SetRemediationNil sets the value for Remediation to be an explicit nil

### UnsetRemediation
`func (o *SecureScoreControlProfile) UnsetRemediation()`

UnsetRemediation ensures that no value is present for Remediation, not even an explicit nil
### GetRemediationImpact

`func (o *SecureScoreControlProfile) GetRemediationImpact() string`

GetRemediationImpact returns the RemediationImpact field if non-nil, zero value otherwise.

### GetRemediationImpactOk

`func (o *SecureScoreControlProfile) GetRemediationImpactOk() (*string, bool)`

GetRemediationImpactOk returns a tuple with the RemediationImpact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemediationImpact

`func (o *SecureScoreControlProfile) SetRemediationImpact(v string)`

SetRemediationImpact sets RemediationImpact field to given value.

### HasRemediationImpact

`func (o *SecureScoreControlProfile) HasRemediationImpact() bool`

HasRemediationImpact returns a boolean if a field has been set.

### SetRemediationImpactNil

`func (o *SecureScoreControlProfile) SetRemediationImpactNil(b bool)`

 SetRemediationImpactNil sets the value for RemediationImpact to be an explicit nil

### UnsetRemediationImpact
`func (o *SecureScoreControlProfile) UnsetRemediationImpact()`

UnsetRemediationImpact ensures that no value is present for RemediationImpact, not even an explicit nil
### GetService

`func (o *SecureScoreControlProfile) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *SecureScoreControlProfile) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *SecureScoreControlProfile) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *SecureScoreControlProfile) HasService() bool`

HasService returns a boolean if a field has been set.

### SetServiceNil

`func (o *SecureScoreControlProfile) SetServiceNil(b bool)`

 SetServiceNil sets the value for Service to be an explicit nil

### UnsetService
`func (o *SecureScoreControlProfile) UnsetService()`

UnsetService ensures that no value is present for Service, not even an explicit nil
### GetThreats

`func (o *SecureScoreControlProfile) GetThreats() []*string`

GetThreats returns the Threats field if non-nil, zero value otherwise.

### GetThreatsOk

`func (o *SecureScoreControlProfile) GetThreatsOk() (*[]*string, bool)`

GetThreatsOk returns a tuple with the Threats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreats

`func (o *SecureScoreControlProfile) SetThreats(v []*string)`

SetThreats sets Threats field to given value.

### HasThreats

`func (o *SecureScoreControlProfile) HasThreats() bool`

HasThreats returns a boolean if a field has been set.

### GetTier

`func (o *SecureScoreControlProfile) GetTier() string`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *SecureScoreControlProfile) GetTierOk() (*string, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *SecureScoreControlProfile) SetTier(v string)`

SetTier sets Tier field to given value.

### HasTier

`func (o *SecureScoreControlProfile) HasTier() bool`

HasTier returns a boolean if a field has been set.

### SetTierNil

`func (o *SecureScoreControlProfile) SetTierNil(b bool)`

 SetTierNil sets the value for Tier to be an explicit nil

### UnsetTier
`func (o *SecureScoreControlProfile) UnsetTier()`

UnsetTier ensures that no value is present for Tier, not even an explicit nil
### GetTitle

`func (o *SecureScoreControlProfile) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SecureScoreControlProfile) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SecureScoreControlProfile) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *SecureScoreControlProfile) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *SecureScoreControlProfile) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *SecureScoreControlProfile) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetUserImpact

`func (o *SecureScoreControlProfile) GetUserImpact() string`

GetUserImpact returns the UserImpact field if non-nil, zero value otherwise.

### GetUserImpactOk

`func (o *SecureScoreControlProfile) GetUserImpactOk() (*string, bool)`

GetUserImpactOk returns a tuple with the UserImpact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserImpact

`func (o *SecureScoreControlProfile) SetUserImpact(v string)`

SetUserImpact sets UserImpact field to given value.

### HasUserImpact

`func (o *SecureScoreControlProfile) HasUserImpact() bool`

HasUserImpact returns a boolean if a field has been set.

### SetUserImpactNil

`func (o *SecureScoreControlProfile) SetUserImpactNil(b bool)`

 SetUserImpactNil sets the value for UserImpact to be an explicit nil

### UnsetUserImpact
`func (o *SecureScoreControlProfile) UnsetUserImpact()`

UnsetUserImpact ensures that no value is present for UserImpact, not even an explicit nil
### GetVendorInformation

`func (o *SecureScoreControlProfile) GetVendorInformation() AnyOfmicrosoftGraphSecurityVendorInformation`

GetVendorInformation returns the VendorInformation field if non-nil, zero value otherwise.

### GetVendorInformationOk

`func (o *SecureScoreControlProfile) GetVendorInformationOk() (*AnyOfmicrosoftGraphSecurityVendorInformation, bool)`

GetVendorInformationOk returns a tuple with the VendorInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorInformation

`func (o *SecureScoreControlProfile) SetVendorInformation(v AnyOfmicrosoftGraphSecurityVendorInformation)`

SetVendorInformation sets VendorInformation field to given value.

### HasVendorInformation

`func (o *SecureScoreControlProfile) HasVendorInformation() bool`

HasVendorInformation returns a boolean if a field has been set.

### SetVendorInformationNil

`func (o *SecureScoreControlProfile) SetVendorInformationNil(b bool)`

 SetVendorInformationNil sets the value for VendorInformation to be an explicit nil

### UnsetVendorInformation
`func (o *SecureScoreControlProfile) UnsetVendorInformation()`

UnsetVendorInformation ensures that no value is present for VendorInformation, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


