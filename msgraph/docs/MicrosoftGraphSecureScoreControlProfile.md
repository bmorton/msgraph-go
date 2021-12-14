# MicrosoftGraphSecureScoreControlProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
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

### NewMicrosoftGraphSecureScoreControlProfile

`func NewMicrosoftGraphSecureScoreControlProfile() *MicrosoftGraphSecureScoreControlProfile`

NewMicrosoftGraphSecureScoreControlProfile instantiates a new MicrosoftGraphSecureScoreControlProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphSecureScoreControlProfileWithDefaults

`func NewMicrosoftGraphSecureScoreControlProfileWithDefaults() *MicrosoftGraphSecureScoreControlProfile`

NewMicrosoftGraphSecureScoreControlProfileWithDefaults instantiates a new MicrosoftGraphSecureScoreControlProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphSecureScoreControlProfile) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphSecureScoreControlProfile) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphSecureScoreControlProfile) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphSecureScoreControlProfile) HasId() bool`

HasId returns a boolean if a field has been set.

### GetActionType

`func (o *MicrosoftGraphSecureScoreControlProfile) GetActionType() string`

GetActionType returns the ActionType field if non-nil, zero value otherwise.

### GetActionTypeOk

`func (o *MicrosoftGraphSecureScoreControlProfile) GetActionTypeOk() (*string, bool)`

GetActionTypeOk returns a tuple with the ActionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionType

`func (o *MicrosoftGraphSecureScoreControlProfile) SetActionType(v string)`

SetActionType sets ActionType field to given value.

### HasActionType

`func (o *MicrosoftGraphSecureScoreControlProfile) HasActionType() bool`

HasActionType returns a boolean if a field has been set.

### SetActionTypeNil

`func (o *MicrosoftGraphSecureScoreControlProfile) SetActionTypeNil(b bool)`

 SetActionTypeNil sets the value for ActionType to be an explicit nil

### UnsetActionType
`func (o *MicrosoftGraphSecureScoreControlProfile) UnsetActionType()`

UnsetActionType ensures that no value is present for ActionType, not even an explicit nil
### GetActionUrl

`func (o *MicrosoftGraphSecureScoreControlProfile) GetActionUrl() string`

GetActionUrl returns the ActionUrl field if non-nil, zero value otherwise.

### GetActionUrlOk

`func (o *MicrosoftGraphSecureScoreControlProfile) GetActionUrlOk() (*string, bool)`

GetActionUrlOk returns a tuple with the ActionUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionUrl

`func (o *MicrosoftGraphSecureScoreControlProfile) SetActionUrl(v string)`

SetActionUrl sets ActionUrl field to given value.

### HasActionUrl

`func (o *MicrosoftGraphSecureScoreControlProfile) HasActionUrl() bool`

HasActionUrl returns a boolean if a field has been set.

### SetActionUrlNil

`func (o *MicrosoftGraphSecureScoreControlProfile) SetActionUrlNil(b bool)`

 SetActionUrlNil sets the value for ActionUrl to be an explicit nil

### UnsetActionUrl
`func (o *MicrosoftGraphSecureScoreControlProfile) UnsetActionUrl()`

UnsetActionUrl ensures that no value is present for ActionUrl, not even an explicit nil
### GetAzureTenantId

`func (o *MicrosoftGraphSecureScoreControlProfile) GetAzureTenantId() string`

GetAzureTenantId returns the AzureTenantId field if non-nil, zero value otherwise.

### GetAzureTenantIdOk

`func (o *MicrosoftGraphSecureScoreControlProfile) GetAzureTenantIdOk() (*string, bool)`

GetAzureTenantIdOk returns a tuple with the AzureTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureTenantId

`func (o *MicrosoftGraphSecureScoreControlProfile) SetAzureTenantId(v string)`

SetAzureTenantId sets AzureTenantId field to given value.

### HasAzureTenantId

`func (o *MicrosoftGraphSecureScoreControlProfile) HasAzureTenantId() bool`

HasAzureTenantId returns a boolean if a field has been set.

### GetComplianceInformation

`func (o *MicrosoftGraphSecureScoreControlProfile) GetComplianceInformation() []*AnyOfmicrosoftGraphComplianceInformation`

GetComplianceInformation returns the ComplianceInformation field if non-nil, zero value otherwise.

### GetComplianceInformationOk

`func (o *MicrosoftGraphSecureScoreControlProfile) GetComplianceInformationOk() (*[]*AnyOfmicrosoftGraphComplianceInformation, bool)`

GetComplianceInformationOk returns a tuple with the ComplianceInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplianceInformation

`func (o *MicrosoftGraphSecureScoreControlProfile) SetComplianceInformation(v []*AnyOfmicrosoftGraphComplianceInformation)`

SetComplianceInformation sets ComplianceInformation field to given value.

### HasComplianceInformation

`func (o *MicrosoftGraphSecureScoreControlProfile) HasComplianceInformation() bool`

HasComplianceInformation returns a boolean if a field has been set.

### GetControlCategory

`func (o *MicrosoftGraphSecureScoreControlProfile) GetControlCategory() string`

GetControlCategory returns the ControlCategory field if non-nil, zero value otherwise.

### GetControlCategoryOk

`func (o *MicrosoftGraphSecureScoreControlProfile) GetControlCategoryOk() (*string, bool)`

GetControlCategoryOk returns a tuple with the ControlCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCategory

`func (o *MicrosoftGraphSecureScoreControlProfile) SetControlCategory(v string)`

SetControlCategory sets ControlCategory field to given value.

### HasControlCategory

`func (o *MicrosoftGraphSecureScoreControlProfile) HasControlCategory() bool`

HasControlCategory returns a boolean if a field has been set.

### SetControlCategoryNil

`func (o *MicrosoftGraphSecureScoreControlProfile) SetControlCategoryNil(b bool)`

 SetControlCategoryNil sets the value for ControlCategory to be an explicit nil

### UnsetControlCategory
`func (o *MicrosoftGraphSecureScoreControlProfile) UnsetControlCategory()`

UnsetControlCategory ensures that no value is present for ControlCategory, not even an explicit nil
### GetControlStateUpdates

`func (o *MicrosoftGraphSecureScoreControlProfile) GetControlStateUpdates() []*AnyOfmicrosoftGraphSecureScoreControlStateUpdate`

GetControlStateUpdates returns the ControlStateUpdates field if non-nil, zero value otherwise.

### GetControlStateUpdatesOk

`func (o *MicrosoftGraphSecureScoreControlProfile) GetControlStateUpdatesOk() (*[]*AnyOfmicrosoftGraphSecureScoreControlStateUpdate, bool)`

GetControlStateUpdatesOk returns a tuple with the ControlStateUpdates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlStateUpdates

`func (o *MicrosoftGraphSecureScoreControlProfile) SetControlStateUpdates(v []*AnyOfmicrosoftGraphSecureScoreControlStateUpdate)`

SetControlStateUpdates sets ControlStateUpdates field to given value.

### HasControlStateUpdates

`func (o *MicrosoftGraphSecureScoreControlProfile) HasControlStateUpdates() bool`

HasControlStateUpdates returns a boolean if a field has been set.

### GetDeprecated

`func (o *MicrosoftGraphSecureScoreControlProfile) GetDeprecated() bool`

GetDeprecated returns the Deprecated field if non-nil, zero value otherwise.

### GetDeprecatedOk

`func (o *MicrosoftGraphSecureScoreControlProfile) GetDeprecatedOk() (*bool, bool)`

GetDeprecatedOk returns a tuple with the Deprecated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecated

`func (o *MicrosoftGraphSecureScoreControlProfile) SetDeprecated(v bool)`

SetDeprecated sets Deprecated field to given value.

### HasDeprecated

`func (o *MicrosoftGraphSecureScoreControlProfile) HasDeprecated() bool`

HasDeprecated returns a boolean if a field has been set.

### SetDeprecatedNil

`func (o *MicrosoftGraphSecureScoreControlProfile) SetDeprecatedNil(b bool)`

 SetDeprecatedNil sets the value for Deprecated to be an explicit nil

### UnsetDeprecated
`func (o *MicrosoftGraphSecureScoreControlProfile) UnsetDeprecated()`

UnsetDeprecated ensures that no value is present for Deprecated, not even an explicit nil
### GetImplementationCost

`func (o *MicrosoftGraphSecureScoreControlProfile) GetImplementationCost() string`

GetImplementationCost returns the ImplementationCost field if non-nil, zero value otherwise.

### GetImplementationCostOk

`func (o *MicrosoftGraphSecureScoreControlProfile) GetImplementationCostOk() (*string, bool)`

GetImplementationCostOk returns a tuple with the ImplementationCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImplementationCost

`func (o *MicrosoftGraphSecureScoreControlProfile) SetImplementationCost(v string)`

SetImplementationCost sets ImplementationCost field to given value.

### HasImplementationCost

`func (o *MicrosoftGraphSecureScoreControlProfile) HasImplementationCost() bool`

HasImplementationCost returns a boolean if a field has been set.

### SetImplementationCostNil

`func (o *MicrosoftGraphSecureScoreControlProfile) SetImplementationCostNil(b bool)`

 SetImplementationCostNil sets the value for ImplementationCost to be an explicit nil

### UnsetImplementationCost
`func (o *MicrosoftGraphSecureScoreControlProfile) UnsetImplementationCost()`

UnsetImplementationCost ensures that no value is present for ImplementationCost, not even an explicit nil
### GetLastModifiedDateTime

`func (o *MicrosoftGraphSecureScoreControlProfile) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *MicrosoftGraphSecureScoreControlProfile) GetLastModifiedDateTimeOk() (*time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedDateTime

`func (o *MicrosoftGraphSecureScoreControlProfile) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime sets LastModifiedDateTime field to given value.

### HasLastModifiedDateTime

`func (o *MicrosoftGraphSecureScoreControlProfile) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### SetLastModifiedDateTimeNil

`func (o *MicrosoftGraphSecureScoreControlProfile) SetLastModifiedDateTimeNil(b bool)`

 SetLastModifiedDateTimeNil sets the value for LastModifiedDateTime to be an explicit nil

### UnsetLastModifiedDateTime
`func (o *MicrosoftGraphSecureScoreControlProfile) UnsetLastModifiedDateTime()`

UnsetLastModifiedDateTime ensures that no value is present for LastModifiedDateTime, not even an explicit nil
### GetMaxScore

`func (o *MicrosoftGraphSecureScoreControlProfile) GetMaxScore() AnyOfnumberstringstring`

GetMaxScore returns the MaxScore field if non-nil, zero value otherwise.

### GetMaxScoreOk

`func (o *MicrosoftGraphSecureScoreControlProfile) GetMaxScoreOk() (*AnyOfnumberstringstring, bool)`

GetMaxScoreOk returns a tuple with the MaxScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxScore

`func (o *MicrosoftGraphSecureScoreControlProfile) SetMaxScore(v AnyOfnumberstringstring)`

SetMaxScore sets MaxScore field to given value.

### HasMaxScore

`func (o *MicrosoftGraphSecureScoreControlProfile) HasMaxScore() bool`

HasMaxScore returns a boolean if a field has been set.

### SetMaxScoreNil

`func (o *MicrosoftGraphSecureScoreControlProfile) SetMaxScoreNil(b bool)`

 SetMaxScoreNil sets the value for MaxScore to be an explicit nil

### UnsetMaxScore
`func (o *MicrosoftGraphSecureScoreControlProfile) UnsetMaxScore()`

UnsetMaxScore ensures that no value is present for MaxScore, not even an explicit nil
### GetRank

`func (o *MicrosoftGraphSecureScoreControlProfile) GetRank() int32`

GetRank returns the Rank field if non-nil, zero value otherwise.

### GetRankOk

`func (o *MicrosoftGraphSecureScoreControlProfile) GetRankOk() (*int32, bool)`

GetRankOk returns a tuple with the Rank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRank

`func (o *MicrosoftGraphSecureScoreControlProfile) SetRank(v int32)`

SetRank sets Rank field to given value.

### HasRank

`func (o *MicrosoftGraphSecureScoreControlProfile) HasRank() bool`

HasRank returns a boolean if a field has been set.

### SetRankNil

`func (o *MicrosoftGraphSecureScoreControlProfile) SetRankNil(b bool)`

 SetRankNil sets the value for Rank to be an explicit nil

### UnsetRank
`func (o *MicrosoftGraphSecureScoreControlProfile) UnsetRank()`

UnsetRank ensures that no value is present for Rank, not even an explicit nil
### GetRemediation

`func (o *MicrosoftGraphSecureScoreControlProfile) GetRemediation() string`

GetRemediation returns the Remediation field if non-nil, zero value otherwise.

### GetRemediationOk

`func (o *MicrosoftGraphSecureScoreControlProfile) GetRemediationOk() (*string, bool)`

GetRemediationOk returns a tuple with the Remediation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemediation

`func (o *MicrosoftGraphSecureScoreControlProfile) SetRemediation(v string)`

SetRemediation sets Remediation field to given value.

### HasRemediation

`func (o *MicrosoftGraphSecureScoreControlProfile) HasRemediation() bool`

HasRemediation returns a boolean if a field has been set.

### SetRemediationNil

`func (o *MicrosoftGraphSecureScoreControlProfile) SetRemediationNil(b bool)`

 SetRemediationNil sets the value for Remediation to be an explicit nil

### UnsetRemediation
`func (o *MicrosoftGraphSecureScoreControlProfile) UnsetRemediation()`

UnsetRemediation ensures that no value is present for Remediation, not even an explicit nil
### GetRemediationImpact

`func (o *MicrosoftGraphSecureScoreControlProfile) GetRemediationImpact() string`

GetRemediationImpact returns the RemediationImpact field if non-nil, zero value otherwise.

### GetRemediationImpactOk

`func (o *MicrosoftGraphSecureScoreControlProfile) GetRemediationImpactOk() (*string, bool)`

GetRemediationImpactOk returns a tuple with the RemediationImpact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemediationImpact

`func (o *MicrosoftGraphSecureScoreControlProfile) SetRemediationImpact(v string)`

SetRemediationImpact sets RemediationImpact field to given value.

### HasRemediationImpact

`func (o *MicrosoftGraphSecureScoreControlProfile) HasRemediationImpact() bool`

HasRemediationImpact returns a boolean if a field has been set.

### SetRemediationImpactNil

`func (o *MicrosoftGraphSecureScoreControlProfile) SetRemediationImpactNil(b bool)`

 SetRemediationImpactNil sets the value for RemediationImpact to be an explicit nil

### UnsetRemediationImpact
`func (o *MicrosoftGraphSecureScoreControlProfile) UnsetRemediationImpact()`

UnsetRemediationImpact ensures that no value is present for RemediationImpact, not even an explicit nil
### GetService

`func (o *MicrosoftGraphSecureScoreControlProfile) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *MicrosoftGraphSecureScoreControlProfile) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *MicrosoftGraphSecureScoreControlProfile) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *MicrosoftGraphSecureScoreControlProfile) HasService() bool`

HasService returns a boolean if a field has been set.

### SetServiceNil

`func (o *MicrosoftGraphSecureScoreControlProfile) SetServiceNil(b bool)`

 SetServiceNil sets the value for Service to be an explicit nil

### UnsetService
`func (o *MicrosoftGraphSecureScoreControlProfile) UnsetService()`

UnsetService ensures that no value is present for Service, not even an explicit nil
### GetThreats

`func (o *MicrosoftGraphSecureScoreControlProfile) GetThreats() []*string`

GetThreats returns the Threats field if non-nil, zero value otherwise.

### GetThreatsOk

`func (o *MicrosoftGraphSecureScoreControlProfile) GetThreatsOk() (*[]*string, bool)`

GetThreatsOk returns a tuple with the Threats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreats

`func (o *MicrosoftGraphSecureScoreControlProfile) SetThreats(v []*string)`

SetThreats sets Threats field to given value.

### HasThreats

`func (o *MicrosoftGraphSecureScoreControlProfile) HasThreats() bool`

HasThreats returns a boolean if a field has been set.

### GetTier

`func (o *MicrosoftGraphSecureScoreControlProfile) GetTier() string`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *MicrosoftGraphSecureScoreControlProfile) GetTierOk() (*string, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *MicrosoftGraphSecureScoreControlProfile) SetTier(v string)`

SetTier sets Tier field to given value.

### HasTier

`func (o *MicrosoftGraphSecureScoreControlProfile) HasTier() bool`

HasTier returns a boolean if a field has been set.

### SetTierNil

`func (o *MicrosoftGraphSecureScoreControlProfile) SetTierNil(b bool)`

 SetTierNil sets the value for Tier to be an explicit nil

### UnsetTier
`func (o *MicrosoftGraphSecureScoreControlProfile) UnsetTier()`

UnsetTier ensures that no value is present for Tier, not even an explicit nil
### GetTitle

`func (o *MicrosoftGraphSecureScoreControlProfile) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *MicrosoftGraphSecureScoreControlProfile) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *MicrosoftGraphSecureScoreControlProfile) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *MicrosoftGraphSecureScoreControlProfile) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *MicrosoftGraphSecureScoreControlProfile) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *MicrosoftGraphSecureScoreControlProfile) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetUserImpact

`func (o *MicrosoftGraphSecureScoreControlProfile) GetUserImpact() string`

GetUserImpact returns the UserImpact field if non-nil, zero value otherwise.

### GetUserImpactOk

`func (o *MicrosoftGraphSecureScoreControlProfile) GetUserImpactOk() (*string, bool)`

GetUserImpactOk returns a tuple with the UserImpact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserImpact

`func (o *MicrosoftGraphSecureScoreControlProfile) SetUserImpact(v string)`

SetUserImpact sets UserImpact field to given value.

### HasUserImpact

`func (o *MicrosoftGraphSecureScoreControlProfile) HasUserImpact() bool`

HasUserImpact returns a boolean if a field has been set.

### SetUserImpactNil

`func (o *MicrosoftGraphSecureScoreControlProfile) SetUserImpactNil(b bool)`

 SetUserImpactNil sets the value for UserImpact to be an explicit nil

### UnsetUserImpact
`func (o *MicrosoftGraphSecureScoreControlProfile) UnsetUserImpact()`

UnsetUserImpact ensures that no value is present for UserImpact, not even an explicit nil
### GetVendorInformation

`func (o *MicrosoftGraphSecureScoreControlProfile) GetVendorInformation() AnyOfmicrosoftGraphSecurityVendorInformation`

GetVendorInformation returns the VendorInformation field if non-nil, zero value otherwise.

### GetVendorInformationOk

`func (o *MicrosoftGraphSecureScoreControlProfile) GetVendorInformationOk() (*AnyOfmicrosoftGraphSecurityVendorInformation, bool)`

GetVendorInformationOk returns a tuple with the VendorInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorInformation

`func (o *MicrosoftGraphSecureScoreControlProfile) SetVendorInformation(v AnyOfmicrosoftGraphSecurityVendorInformation)`

SetVendorInformation sets VendorInformation field to given value.

### HasVendorInformation

`func (o *MicrosoftGraphSecureScoreControlProfile) HasVendorInformation() bool`

HasVendorInformation returns a boolean if a field has been set.

### SetVendorInformationNil

`func (o *MicrosoftGraphSecureScoreControlProfile) SetVendorInformationNil(b bool)`

 SetVendorInformationNil sets the value for VendorInformation to be an explicit nil

### UnsetVendorInformation
`func (o *MicrosoftGraphSecureScoreControlProfile) UnsetVendorInformation()`

UnsetVendorInformation ensures that no value is present for VendorInformation, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


