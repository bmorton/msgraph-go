# MicrosoftGraphRestrictedSignIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**AppDisplayName** | Pointer to **NullableString** | App name displayed in the Azure Portal. Supports $filter (eq and startsWith operators only). | [optional] 
**AppId** | Pointer to **NullableString** | Unique GUID representing the app ID in the Azure Active Directory. Supports $filter (eq operator only). | [optional] 
**AppliedConditionalAccessPolicies** | Pointer to [**[]AnyOfmicrosoftGraphAppliedConditionalAccessPolicy**](AnyOfmicrosoftGraphAppliedConditionalAccessPolicy.md) | A list of conditional access policies that are triggered by the corresponding sign-in activity. | [optional] 
**ClientAppUsed** | Pointer to **NullableString** | Identifies the legacy client used for sign-in activity.  Includes Browser, Exchange Active Sync, modern clients, IMAP, MAPI, SMTP, and POP. Supports $filter (eq operator only). | [optional] 
**ConditionalAccessStatus** | Pointer to [**AnyOfmicrosoftGraphConditionalAccessStatus**](anyOf&lt;microsoft.graph.conditionalAccessStatus&gt;.md) | Reports status of an activated conditional access policy. Possible values are: success, failure, notApplied, and unknownFutureValue. Supports $filter (eq operator only). | [optional] 
**CorrelationId** | Pointer to **NullableString** | The request ID sent from the client when the sign-in is initiated; used to troubleshoot sign-in activity. Supports $filter (eq operator only). | [optional] 
**CreatedDateTime** | Pointer to **time.Time** | Date and time (UTC) the sign-in was initiated. Example: midnight on Jan 1, 2014 is reported as 2014-01-01T00:00:00Z. Supports $orderby and $filter (eq, le, and ge operators only). | [optional] 
**DeviceDetail** | Pointer to [**AnyOfmicrosoftGraphDeviceDetail**](anyOf&lt;microsoft.graph.deviceDetail&gt;.md) | Device information from where the sign-in occurred; includes device ID, operating system, and browser. Supports $filter (eq and startsWith operators only) on browser and operatingSytem properties. | [optional] 
**IpAddress** | Pointer to **NullableString** | IP address of the client used to sign in. Supports $filter (eq and startsWith operators only). | [optional] 
**IsInteractive** | Pointer to **NullableBool** | Indicates if a sign-in is interactive or not. | [optional] 
**Location** | Pointer to [**AnyOfmicrosoftGraphSignInLocation**](anyOf&lt;microsoft.graph.signInLocation&gt;.md) | Provides the city, state, and country code where the sign-in originated. Supports $filter (eq and startsWith operators only) on city, state, and countryOrRegion properties. | [optional] 
**ResourceDisplayName** | Pointer to **NullableString** | Name of the resource the user signed into. Supports $filter (eq operator only). | [optional] 
**ResourceId** | Pointer to **NullableString** | ID of the resource that the user signed into. Supports $filter (eq operator only). | [optional] 
**RiskDetail** | Pointer to [**AnyOfmicrosoftGraphRiskDetail**](anyOf&lt;microsoft.graph.riskDetail&gt;.md) | Provides the &#39;reason&#39; behind a specific state of a risky user, sign-in or a risk event. The possible values are: none, adminGeneratedTemporaryPassword, userPerformedSecuredPasswordChange, userPerformedSecuredPasswordReset, adminConfirmedSigninSafe, aiConfirmedSigninSafe, userPassedMFADrivenByRiskBasedPolicy, adminDismissedAllRiskForUser, adminConfirmedSigninCompromised, unknownFutureValue. The value none means that no action has been performed on the user or sign-in so far.  Supports $filter (eq operator only).Note: Details for this property require an Azure AD Premium P2 license. Other licenses return the value hidden. | [optional] 
**RiskEventTypes** | Pointer to [**[]AnyOfmicrosoftGraphRiskEventType**](AnyOfmicrosoftGraphRiskEventType.md) | Risk event types associated with the sign-in. The possible values are: unlikelyTravel, anonymizedIPAddress, maliciousIPAddress, unfamiliarFeatures, malwareInfectedIPAddress, suspiciousIPAddress, leakedCredentials, investigationsThreatIntelligence,  generic, and unknownFutureValue. Supports $filter (eq operator only). | [optional] 
**RiskEventTypesV2** | Pointer to **[]string** | The list of risk event types associated with the sign-in. Possible values: unlikelyTravel, anonymizedIPAddress, maliciousIPAddress, unfamiliarFeatures, malwareInfectedIPAddress, suspiciousIPAddress, leakedCredentials, investigationsThreatIntelligence,  generic, or unknownFutureValue. Supports $filter (eq and startsWith operators only). | [optional] 
**RiskLevelAggregated** | Pointer to [**AnyOfmicrosoftGraphRiskLevel**](anyOf&lt;microsoft.graph.riskLevel&gt;.md) | Aggregated risk level. The possible values are: none, low, medium, high, hidden, and unknownFutureValue. The value hidden means the user or sign-in was not enabled for Azure AD Identity Protection. Supports $filter (eq operator only).  Note: Details for this property are only available for Azure AD Premium P2 customers. All other customers will be returned hidden. | [optional] 
**RiskLevelDuringSignIn** | Pointer to [**AnyOfmicrosoftGraphRiskLevel**](anyOf&lt;microsoft.graph.riskLevel&gt;.md) | Risk level during sign-in. The possible values are: none, low, medium, high, hidden, and unknownFutureValue. The value hidden means the user or sign-in was not enabled for Azure AD Identity Protection.  Supports $filter (eq operator only). Note: Details for this property are only available for Azure AD Premium P2 customers. All other customers will be returned hidden. | [optional] 
**RiskState** | Pointer to [**AnyOfmicrosoftGraphRiskState**](anyOf&lt;microsoft.graph.riskState&gt;.md) | Reports status of the risky user, sign-in, or a risk event. The possible values are: none, confirmedSafe, remediated, dismissed, atRisk, confirmedCompromised, unknownFutureValue. Supports $filter (eq operator only). | [optional] 
**Status** | Pointer to [**AnyOfmicrosoftGraphSignInStatus**](anyOf&lt;microsoft.graph.signInStatus&gt;.md) | Sign-in status. Includes the error code and description of the error (in case of a sign-in failure). Supports $filter (eq operator only) on errorCode property. | [optional] 
**UserDisplayName** | Pointer to **NullableString** | Display name of the user that initiated the sign-in. Supports $filter (eq and startsWith operators only). | [optional] 
**UserId** | Pointer to **string** | ID of the user that initiated the sign-in. Supports $filter (eq operator only). | [optional] 
**UserPrincipalName** | Pointer to **NullableString** | User principal name of the user that initiated the sign-in. Supports $filter (eq and startsWith operators only). | [optional] 
**TargetTenantId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewMicrosoftGraphRestrictedSignIn

`func NewMicrosoftGraphRestrictedSignIn() *MicrosoftGraphRestrictedSignIn`

NewMicrosoftGraphRestrictedSignIn instantiates a new MicrosoftGraphRestrictedSignIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphRestrictedSignInWithDefaults

`func NewMicrosoftGraphRestrictedSignInWithDefaults() *MicrosoftGraphRestrictedSignIn`

NewMicrosoftGraphRestrictedSignInWithDefaults instantiates a new MicrosoftGraphRestrictedSignIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphRestrictedSignIn) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphRestrictedSignIn) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphRestrictedSignIn) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphRestrictedSignIn) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAppDisplayName

`func (o *MicrosoftGraphRestrictedSignIn) GetAppDisplayName() string`

GetAppDisplayName returns the AppDisplayName field if non-nil, zero value otherwise.

### GetAppDisplayNameOk

`func (o *MicrosoftGraphRestrictedSignIn) GetAppDisplayNameOk() (*string, bool)`

GetAppDisplayNameOk returns a tuple with the AppDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppDisplayName

`func (o *MicrosoftGraphRestrictedSignIn) SetAppDisplayName(v string)`

SetAppDisplayName sets AppDisplayName field to given value.

### HasAppDisplayName

`func (o *MicrosoftGraphRestrictedSignIn) HasAppDisplayName() bool`

HasAppDisplayName returns a boolean if a field has been set.

### SetAppDisplayNameNil

`func (o *MicrosoftGraphRestrictedSignIn) SetAppDisplayNameNil(b bool)`

 SetAppDisplayNameNil sets the value for AppDisplayName to be an explicit nil

### UnsetAppDisplayName
`func (o *MicrosoftGraphRestrictedSignIn) UnsetAppDisplayName()`

UnsetAppDisplayName ensures that no value is present for AppDisplayName, not even an explicit nil
### GetAppId

`func (o *MicrosoftGraphRestrictedSignIn) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *MicrosoftGraphRestrictedSignIn) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *MicrosoftGraphRestrictedSignIn) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *MicrosoftGraphRestrictedSignIn) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### SetAppIdNil

`func (o *MicrosoftGraphRestrictedSignIn) SetAppIdNil(b bool)`

 SetAppIdNil sets the value for AppId to be an explicit nil

### UnsetAppId
`func (o *MicrosoftGraphRestrictedSignIn) UnsetAppId()`

UnsetAppId ensures that no value is present for AppId, not even an explicit nil
### GetAppliedConditionalAccessPolicies

`func (o *MicrosoftGraphRestrictedSignIn) GetAppliedConditionalAccessPolicies() []*AnyOfmicrosoftGraphAppliedConditionalAccessPolicy`

GetAppliedConditionalAccessPolicies returns the AppliedConditionalAccessPolicies field if non-nil, zero value otherwise.

### GetAppliedConditionalAccessPoliciesOk

`func (o *MicrosoftGraphRestrictedSignIn) GetAppliedConditionalAccessPoliciesOk() (*[]*AnyOfmicrosoftGraphAppliedConditionalAccessPolicy, bool)`

GetAppliedConditionalAccessPoliciesOk returns a tuple with the AppliedConditionalAccessPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedConditionalAccessPolicies

`func (o *MicrosoftGraphRestrictedSignIn) SetAppliedConditionalAccessPolicies(v []*AnyOfmicrosoftGraphAppliedConditionalAccessPolicy)`

SetAppliedConditionalAccessPolicies sets AppliedConditionalAccessPolicies field to given value.

### HasAppliedConditionalAccessPolicies

`func (o *MicrosoftGraphRestrictedSignIn) HasAppliedConditionalAccessPolicies() bool`

HasAppliedConditionalAccessPolicies returns a boolean if a field has been set.

### GetClientAppUsed

`func (o *MicrosoftGraphRestrictedSignIn) GetClientAppUsed() string`

GetClientAppUsed returns the ClientAppUsed field if non-nil, zero value otherwise.

### GetClientAppUsedOk

`func (o *MicrosoftGraphRestrictedSignIn) GetClientAppUsedOk() (*string, bool)`

GetClientAppUsedOk returns a tuple with the ClientAppUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientAppUsed

`func (o *MicrosoftGraphRestrictedSignIn) SetClientAppUsed(v string)`

SetClientAppUsed sets ClientAppUsed field to given value.

### HasClientAppUsed

`func (o *MicrosoftGraphRestrictedSignIn) HasClientAppUsed() bool`

HasClientAppUsed returns a boolean if a field has been set.

### SetClientAppUsedNil

`func (o *MicrosoftGraphRestrictedSignIn) SetClientAppUsedNil(b bool)`

 SetClientAppUsedNil sets the value for ClientAppUsed to be an explicit nil

### UnsetClientAppUsed
`func (o *MicrosoftGraphRestrictedSignIn) UnsetClientAppUsed()`

UnsetClientAppUsed ensures that no value is present for ClientAppUsed, not even an explicit nil
### GetConditionalAccessStatus

`func (o *MicrosoftGraphRestrictedSignIn) GetConditionalAccessStatus() AnyOfmicrosoftGraphConditionalAccessStatus`

GetConditionalAccessStatus returns the ConditionalAccessStatus field if non-nil, zero value otherwise.

### GetConditionalAccessStatusOk

`func (o *MicrosoftGraphRestrictedSignIn) GetConditionalAccessStatusOk() (*AnyOfmicrosoftGraphConditionalAccessStatus, bool)`

GetConditionalAccessStatusOk returns a tuple with the ConditionalAccessStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionalAccessStatus

`func (o *MicrosoftGraphRestrictedSignIn) SetConditionalAccessStatus(v AnyOfmicrosoftGraphConditionalAccessStatus)`

SetConditionalAccessStatus sets ConditionalAccessStatus field to given value.

### HasConditionalAccessStatus

`func (o *MicrosoftGraphRestrictedSignIn) HasConditionalAccessStatus() bool`

HasConditionalAccessStatus returns a boolean if a field has been set.

### SetConditionalAccessStatusNil

`func (o *MicrosoftGraphRestrictedSignIn) SetConditionalAccessStatusNil(b bool)`

 SetConditionalAccessStatusNil sets the value for ConditionalAccessStatus to be an explicit nil

### UnsetConditionalAccessStatus
`func (o *MicrosoftGraphRestrictedSignIn) UnsetConditionalAccessStatus()`

UnsetConditionalAccessStatus ensures that no value is present for ConditionalAccessStatus, not even an explicit nil
### GetCorrelationId

`func (o *MicrosoftGraphRestrictedSignIn) GetCorrelationId() string`

GetCorrelationId returns the CorrelationId field if non-nil, zero value otherwise.

### GetCorrelationIdOk

`func (o *MicrosoftGraphRestrictedSignIn) GetCorrelationIdOk() (*string, bool)`

GetCorrelationIdOk returns a tuple with the CorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationId

`func (o *MicrosoftGraphRestrictedSignIn) SetCorrelationId(v string)`

SetCorrelationId sets CorrelationId field to given value.

### HasCorrelationId

`func (o *MicrosoftGraphRestrictedSignIn) HasCorrelationId() bool`

HasCorrelationId returns a boolean if a field has been set.

### SetCorrelationIdNil

`func (o *MicrosoftGraphRestrictedSignIn) SetCorrelationIdNil(b bool)`

 SetCorrelationIdNil sets the value for CorrelationId to be an explicit nil

### UnsetCorrelationId
`func (o *MicrosoftGraphRestrictedSignIn) UnsetCorrelationId()`

UnsetCorrelationId ensures that no value is present for CorrelationId, not even an explicit nil
### GetCreatedDateTime

`func (o *MicrosoftGraphRestrictedSignIn) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphRestrictedSignIn) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphRestrictedSignIn) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *MicrosoftGraphRestrictedSignIn) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### GetDeviceDetail

`func (o *MicrosoftGraphRestrictedSignIn) GetDeviceDetail() AnyOfmicrosoftGraphDeviceDetail`

GetDeviceDetail returns the DeviceDetail field if non-nil, zero value otherwise.

### GetDeviceDetailOk

`func (o *MicrosoftGraphRestrictedSignIn) GetDeviceDetailOk() (*AnyOfmicrosoftGraphDeviceDetail, bool)`

GetDeviceDetailOk returns a tuple with the DeviceDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceDetail

`func (o *MicrosoftGraphRestrictedSignIn) SetDeviceDetail(v AnyOfmicrosoftGraphDeviceDetail)`

SetDeviceDetail sets DeviceDetail field to given value.

### HasDeviceDetail

`func (o *MicrosoftGraphRestrictedSignIn) HasDeviceDetail() bool`

HasDeviceDetail returns a boolean if a field has been set.

### SetDeviceDetailNil

`func (o *MicrosoftGraphRestrictedSignIn) SetDeviceDetailNil(b bool)`

 SetDeviceDetailNil sets the value for DeviceDetail to be an explicit nil

### UnsetDeviceDetail
`func (o *MicrosoftGraphRestrictedSignIn) UnsetDeviceDetail()`

UnsetDeviceDetail ensures that no value is present for DeviceDetail, not even an explicit nil
### GetIpAddress

`func (o *MicrosoftGraphRestrictedSignIn) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *MicrosoftGraphRestrictedSignIn) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *MicrosoftGraphRestrictedSignIn) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *MicrosoftGraphRestrictedSignIn) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### SetIpAddressNil

`func (o *MicrosoftGraphRestrictedSignIn) SetIpAddressNil(b bool)`

 SetIpAddressNil sets the value for IpAddress to be an explicit nil

### UnsetIpAddress
`func (o *MicrosoftGraphRestrictedSignIn) UnsetIpAddress()`

UnsetIpAddress ensures that no value is present for IpAddress, not even an explicit nil
### GetIsInteractive

`func (o *MicrosoftGraphRestrictedSignIn) GetIsInteractive() bool`

GetIsInteractive returns the IsInteractive field if non-nil, zero value otherwise.

### GetIsInteractiveOk

`func (o *MicrosoftGraphRestrictedSignIn) GetIsInteractiveOk() (*bool, bool)`

GetIsInteractiveOk returns a tuple with the IsInteractive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInteractive

`func (o *MicrosoftGraphRestrictedSignIn) SetIsInteractive(v bool)`

SetIsInteractive sets IsInteractive field to given value.

### HasIsInteractive

`func (o *MicrosoftGraphRestrictedSignIn) HasIsInteractive() bool`

HasIsInteractive returns a boolean if a field has been set.

### SetIsInteractiveNil

`func (o *MicrosoftGraphRestrictedSignIn) SetIsInteractiveNil(b bool)`

 SetIsInteractiveNil sets the value for IsInteractive to be an explicit nil

### UnsetIsInteractive
`func (o *MicrosoftGraphRestrictedSignIn) UnsetIsInteractive()`

UnsetIsInteractive ensures that no value is present for IsInteractive, not even an explicit nil
### GetLocation

`func (o *MicrosoftGraphRestrictedSignIn) GetLocation() AnyOfmicrosoftGraphSignInLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *MicrosoftGraphRestrictedSignIn) GetLocationOk() (*AnyOfmicrosoftGraphSignInLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *MicrosoftGraphRestrictedSignIn) SetLocation(v AnyOfmicrosoftGraphSignInLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *MicrosoftGraphRestrictedSignIn) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *MicrosoftGraphRestrictedSignIn) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *MicrosoftGraphRestrictedSignIn) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetResourceDisplayName

`func (o *MicrosoftGraphRestrictedSignIn) GetResourceDisplayName() string`

GetResourceDisplayName returns the ResourceDisplayName field if non-nil, zero value otherwise.

### GetResourceDisplayNameOk

`func (o *MicrosoftGraphRestrictedSignIn) GetResourceDisplayNameOk() (*string, bool)`

GetResourceDisplayNameOk returns a tuple with the ResourceDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceDisplayName

`func (o *MicrosoftGraphRestrictedSignIn) SetResourceDisplayName(v string)`

SetResourceDisplayName sets ResourceDisplayName field to given value.

### HasResourceDisplayName

`func (o *MicrosoftGraphRestrictedSignIn) HasResourceDisplayName() bool`

HasResourceDisplayName returns a boolean if a field has been set.

### SetResourceDisplayNameNil

`func (o *MicrosoftGraphRestrictedSignIn) SetResourceDisplayNameNil(b bool)`

 SetResourceDisplayNameNil sets the value for ResourceDisplayName to be an explicit nil

### UnsetResourceDisplayName
`func (o *MicrosoftGraphRestrictedSignIn) UnsetResourceDisplayName()`

UnsetResourceDisplayName ensures that no value is present for ResourceDisplayName, not even an explicit nil
### GetResourceId

`func (o *MicrosoftGraphRestrictedSignIn) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *MicrosoftGraphRestrictedSignIn) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *MicrosoftGraphRestrictedSignIn) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *MicrosoftGraphRestrictedSignIn) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### SetResourceIdNil

`func (o *MicrosoftGraphRestrictedSignIn) SetResourceIdNil(b bool)`

 SetResourceIdNil sets the value for ResourceId to be an explicit nil

### UnsetResourceId
`func (o *MicrosoftGraphRestrictedSignIn) UnsetResourceId()`

UnsetResourceId ensures that no value is present for ResourceId, not even an explicit nil
### GetRiskDetail

`func (o *MicrosoftGraphRestrictedSignIn) GetRiskDetail() AnyOfmicrosoftGraphRiskDetail`

GetRiskDetail returns the RiskDetail field if non-nil, zero value otherwise.

### GetRiskDetailOk

`func (o *MicrosoftGraphRestrictedSignIn) GetRiskDetailOk() (*AnyOfmicrosoftGraphRiskDetail, bool)`

GetRiskDetailOk returns a tuple with the RiskDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskDetail

`func (o *MicrosoftGraphRestrictedSignIn) SetRiskDetail(v AnyOfmicrosoftGraphRiskDetail)`

SetRiskDetail sets RiskDetail field to given value.

### HasRiskDetail

`func (o *MicrosoftGraphRestrictedSignIn) HasRiskDetail() bool`

HasRiskDetail returns a boolean if a field has been set.

### SetRiskDetailNil

`func (o *MicrosoftGraphRestrictedSignIn) SetRiskDetailNil(b bool)`

 SetRiskDetailNil sets the value for RiskDetail to be an explicit nil

### UnsetRiskDetail
`func (o *MicrosoftGraphRestrictedSignIn) UnsetRiskDetail()`

UnsetRiskDetail ensures that no value is present for RiskDetail, not even an explicit nil
### GetRiskEventTypes

`func (o *MicrosoftGraphRestrictedSignIn) GetRiskEventTypes() []*AnyOfmicrosoftGraphRiskEventType`

GetRiskEventTypes returns the RiskEventTypes field if non-nil, zero value otherwise.

### GetRiskEventTypesOk

`func (o *MicrosoftGraphRestrictedSignIn) GetRiskEventTypesOk() (*[]*AnyOfmicrosoftGraphRiskEventType, bool)`

GetRiskEventTypesOk returns a tuple with the RiskEventTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskEventTypes

`func (o *MicrosoftGraphRestrictedSignIn) SetRiskEventTypes(v []*AnyOfmicrosoftGraphRiskEventType)`

SetRiskEventTypes sets RiskEventTypes field to given value.

### HasRiskEventTypes

`func (o *MicrosoftGraphRestrictedSignIn) HasRiskEventTypes() bool`

HasRiskEventTypes returns a boolean if a field has been set.

### GetRiskEventTypesV2

`func (o *MicrosoftGraphRestrictedSignIn) GetRiskEventTypesV2() []*string`

GetRiskEventTypesV2 returns the RiskEventTypesV2 field if non-nil, zero value otherwise.

### GetRiskEventTypesV2Ok

`func (o *MicrosoftGraphRestrictedSignIn) GetRiskEventTypesV2Ok() (*[]*string, bool)`

GetRiskEventTypesV2Ok returns a tuple with the RiskEventTypesV2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskEventTypesV2

`func (o *MicrosoftGraphRestrictedSignIn) SetRiskEventTypesV2(v []*string)`

SetRiskEventTypesV2 sets RiskEventTypesV2 field to given value.

### HasRiskEventTypesV2

`func (o *MicrosoftGraphRestrictedSignIn) HasRiskEventTypesV2() bool`

HasRiskEventTypesV2 returns a boolean if a field has been set.

### GetRiskLevelAggregated

`func (o *MicrosoftGraphRestrictedSignIn) GetRiskLevelAggregated() AnyOfmicrosoftGraphRiskLevel`

GetRiskLevelAggregated returns the RiskLevelAggregated field if non-nil, zero value otherwise.

### GetRiskLevelAggregatedOk

`func (o *MicrosoftGraphRestrictedSignIn) GetRiskLevelAggregatedOk() (*AnyOfmicrosoftGraphRiskLevel, bool)`

GetRiskLevelAggregatedOk returns a tuple with the RiskLevelAggregated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskLevelAggregated

`func (o *MicrosoftGraphRestrictedSignIn) SetRiskLevelAggregated(v AnyOfmicrosoftGraphRiskLevel)`

SetRiskLevelAggregated sets RiskLevelAggregated field to given value.

### HasRiskLevelAggregated

`func (o *MicrosoftGraphRestrictedSignIn) HasRiskLevelAggregated() bool`

HasRiskLevelAggregated returns a boolean if a field has been set.

### SetRiskLevelAggregatedNil

`func (o *MicrosoftGraphRestrictedSignIn) SetRiskLevelAggregatedNil(b bool)`

 SetRiskLevelAggregatedNil sets the value for RiskLevelAggregated to be an explicit nil

### UnsetRiskLevelAggregated
`func (o *MicrosoftGraphRestrictedSignIn) UnsetRiskLevelAggregated()`

UnsetRiskLevelAggregated ensures that no value is present for RiskLevelAggregated, not even an explicit nil
### GetRiskLevelDuringSignIn

`func (o *MicrosoftGraphRestrictedSignIn) GetRiskLevelDuringSignIn() AnyOfmicrosoftGraphRiskLevel`

GetRiskLevelDuringSignIn returns the RiskLevelDuringSignIn field if non-nil, zero value otherwise.

### GetRiskLevelDuringSignInOk

`func (o *MicrosoftGraphRestrictedSignIn) GetRiskLevelDuringSignInOk() (*AnyOfmicrosoftGraphRiskLevel, bool)`

GetRiskLevelDuringSignInOk returns a tuple with the RiskLevelDuringSignIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskLevelDuringSignIn

`func (o *MicrosoftGraphRestrictedSignIn) SetRiskLevelDuringSignIn(v AnyOfmicrosoftGraphRiskLevel)`

SetRiskLevelDuringSignIn sets RiskLevelDuringSignIn field to given value.

### HasRiskLevelDuringSignIn

`func (o *MicrosoftGraphRestrictedSignIn) HasRiskLevelDuringSignIn() bool`

HasRiskLevelDuringSignIn returns a boolean if a field has been set.

### SetRiskLevelDuringSignInNil

`func (o *MicrosoftGraphRestrictedSignIn) SetRiskLevelDuringSignInNil(b bool)`

 SetRiskLevelDuringSignInNil sets the value for RiskLevelDuringSignIn to be an explicit nil

### UnsetRiskLevelDuringSignIn
`func (o *MicrosoftGraphRestrictedSignIn) UnsetRiskLevelDuringSignIn()`

UnsetRiskLevelDuringSignIn ensures that no value is present for RiskLevelDuringSignIn, not even an explicit nil
### GetRiskState

`func (o *MicrosoftGraphRestrictedSignIn) GetRiskState() AnyOfmicrosoftGraphRiskState`

GetRiskState returns the RiskState field if non-nil, zero value otherwise.

### GetRiskStateOk

`func (o *MicrosoftGraphRestrictedSignIn) GetRiskStateOk() (*AnyOfmicrosoftGraphRiskState, bool)`

GetRiskStateOk returns a tuple with the RiskState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskState

`func (o *MicrosoftGraphRestrictedSignIn) SetRiskState(v AnyOfmicrosoftGraphRiskState)`

SetRiskState sets RiskState field to given value.

### HasRiskState

`func (o *MicrosoftGraphRestrictedSignIn) HasRiskState() bool`

HasRiskState returns a boolean if a field has been set.

### SetRiskStateNil

`func (o *MicrosoftGraphRestrictedSignIn) SetRiskStateNil(b bool)`

 SetRiskStateNil sets the value for RiskState to be an explicit nil

### UnsetRiskState
`func (o *MicrosoftGraphRestrictedSignIn) UnsetRiskState()`

UnsetRiskState ensures that no value is present for RiskState, not even an explicit nil
### GetStatus

`func (o *MicrosoftGraphRestrictedSignIn) GetStatus() AnyOfmicrosoftGraphSignInStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MicrosoftGraphRestrictedSignIn) GetStatusOk() (*AnyOfmicrosoftGraphSignInStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MicrosoftGraphRestrictedSignIn) SetStatus(v AnyOfmicrosoftGraphSignInStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MicrosoftGraphRestrictedSignIn) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *MicrosoftGraphRestrictedSignIn) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *MicrosoftGraphRestrictedSignIn) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetUserDisplayName

`func (o *MicrosoftGraphRestrictedSignIn) GetUserDisplayName() string`

GetUserDisplayName returns the UserDisplayName field if non-nil, zero value otherwise.

### GetUserDisplayNameOk

`func (o *MicrosoftGraphRestrictedSignIn) GetUserDisplayNameOk() (*string, bool)`

GetUserDisplayNameOk returns a tuple with the UserDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDisplayName

`func (o *MicrosoftGraphRestrictedSignIn) SetUserDisplayName(v string)`

SetUserDisplayName sets UserDisplayName field to given value.

### HasUserDisplayName

`func (o *MicrosoftGraphRestrictedSignIn) HasUserDisplayName() bool`

HasUserDisplayName returns a boolean if a field has been set.

### SetUserDisplayNameNil

`func (o *MicrosoftGraphRestrictedSignIn) SetUserDisplayNameNil(b bool)`

 SetUserDisplayNameNil sets the value for UserDisplayName to be an explicit nil

### UnsetUserDisplayName
`func (o *MicrosoftGraphRestrictedSignIn) UnsetUserDisplayName()`

UnsetUserDisplayName ensures that no value is present for UserDisplayName, not even an explicit nil
### GetUserId

`func (o *MicrosoftGraphRestrictedSignIn) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *MicrosoftGraphRestrictedSignIn) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *MicrosoftGraphRestrictedSignIn) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *MicrosoftGraphRestrictedSignIn) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetUserPrincipalName

`func (o *MicrosoftGraphRestrictedSignIn) GetUserPrincipalName() string`

GetUserPrincipalName returns the UserPrincipalName field if non-nil, zero value otherwise.

### GetUserPrincipalNameOk

`func (o *MicrosoftGraphRestrictedSignIn) GetUserPrincipalNameOk() (*string, bool)`

GetUserPrincipalNameOk returns a tuple with the UserPrincipalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPrincipalName

`func (o *MicrosoftGraphRestrictedSignIn) SetUserPrincipalName(v string)`

SetUserPrincipalName sets UserPrincipalName field to given value.

### HasUserPrincipalName

`func (o *MicrosoftGraphRestrictedSignIn) HasUserPrincipalName() bool`

HasUserPrincipalName returns a boolean if a field has been set.

### SetUserPrincipalNameNil

`func (o *MicrosoftGraphRestrictedSignIn) SetUserPrincipalNameNil(b bool)`

 SetUserPrincipalNameNil sets the value for UserPrincipalName to be an explicit nil

### UnsetUserPrincipalName
`func (o *MicrosoftGraphRestrictedSignIn) UnsetUserPrincipalName()`

UnsetUserPrincipalName ensures that no value is present for UserPrincipalName, not even an explicit nil
### GetTargetTenantId

`func (o *MicrosoftGraphRestrictedSignIn) GetTargetTenantId() string`

GetTargetTenantId returns the TargetTenantId field if non-nil, zero value otherwise.

### GetTargetTenantIdOk

`func (o *MicrosoftGraphRestrictedSignIn) GetTargetTenantIdOk() (*string, bool)`

GetTargetTenantIdOk returns a tuple with the TargetTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetTenantId

`func (o *MicrosoftGraphRestrictedSignIn) SetTargetTenantId(v string)`

SetTargetTenantId sets TargetTenantId field to given value.

### HasTargetTenantId

`func (o *MicrosoftGraphRestrictedSignIn) HasTargetTenantId() bool`

HasTargetTenantId returns a boolean if a field has been set.

### SetTargetTenantIdNil

`func (o *MicrosoftGraphRestrictedSignIn) SetTargetTenantIdNil(b bool)`

 SetTargetTenantIdNil sets the value for TargetTenantId to be an explicit nil

### UnsetTargetTenantId
`func (o *MicrosoftGraphRestrictedSignIn) UnsetTargetTenantId()`

UnsetTargetTenantId ensures that no value is present for TargetTenantId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


