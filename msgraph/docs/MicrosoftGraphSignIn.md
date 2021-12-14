# MicrosoftGraphSignIn

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

## Methods

### NewMicrosoftGraphSignIn

`func NewMicrosoftGraphSignIn() *MicrosoftGraphSignIn`

NewMicrosoftGraphSignIn instantiates a new MicrosoftGraphSignIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphSignInWithDefaults

`func NewMicrosoftGraphSignInWithDefaults() *MicrosoftGraphSignIn`

NewMicrosoftGraphSignInWithDefaults instantiates a new MicrosoftGraphSignIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphSignIn) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphSignIn) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphSignIn) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphSignIn) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAppDisplayName

`func (o *MicrosoftGraphSignIn) GetAppDisplayName() string`

GetAppDisplayName returns the AppDisplayName field if non-nil, zero value otherwise.

### GetAppDisplayNameOk

`func (o *MicrosoftGraphSignIn) GetAppDisplayNameOk() (*string, bool)`

GetAppDisplayNameOk returns a tuple with the AppDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppDisplayName

`func (o *MicrosoftGraphSignIn) SetAppDisplayName(v string)`

SetAppDisplayName sets AppDisplayName field to given value.

### HasAppDisplayName

`func (o *MicrosoftGraphSignIn) HasAppDisplayName() bool`

HasAppDisplayName returns a boolean if a field has been set.

### SetAppDisplayNameNil

`func (o *MicrosoftGraphSignIn) SetAppDisplayNameNil(b bool)`

 SetAppDisplayNameNil sets the value for AppDisplayName to be an explicit nil

### UnsetAppDisplayName
`func (o *MicrosoftGraphSignIn) UnsetAppDisplayName()`

UnsetAppDisplayName ensures that no value is present for AppDisplayName, not even an explicit nil
### GetAppId

`func (o *MicrosoftGraphSignIn) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *MicrosoftGraphSignIn) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *MicrosoftGraphSignIn) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *MicrosoftGraphSignIn) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### SetAppIdNil

`func (o *MicrosoftGraphSignIn) SetAppIdNil(b bool)`

 SetAppIdNil sets the value for AppId to be an explicit nil

### UnsetAppId
`func (o *MicrosoftGraphSignIn) UnsetAppId()`

UnsetAppId ensures that no value is present for AppId, not even an explicit nil
### GetAppliedConditionalAccessPolicies

`func (o *MicrosoftGraphSignIn) GetAppliedConditionalAccessPolicies() []*AnyOfmicrosoftGraphAppliedConditionalAccessPolicy`

GetAppliedConditionalAccessPolicies returns the AppliedConditionalAccessPolicies field if non-nil, zero value otherwise.

### GetAppliedConditionalAccessPoliciesOk

`func (o *MicrosoftGraphSignIn) GetAppliedConditionalAccessPoliciesOk() (*[]*AnyOfmicrosoftGraphAppliedConditionalAccessPolicy, bool)`

GetAppliedConditionalAccessPoliciesOk returns a tuple with the AppliedConditionalAccessPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedConditionalAccessPolicies

`func (o *MicrosoftGraphSignIn) SetAppliedConditionalAccessPolicies(v []*AnyOfmicrosoftGraphAppliedConditionalAccessPolicy)`

SetAppliedConditionalAccessPolicies sets AppliedConditionalAccessPolicies field to given value.

### HasAppliedConditionalAccessPolicies

`func (o *MicrosoftGraphSignIn) HasAppliedConditionalAccessPolicies() bool`

HasAppliedConditionalAccessPolicies returns a boolean if a field has been set.

### GetClientAppUsed

`func (o *MicrosoftGraphSignIn) GetClientAppUsed() string`

GetClientAppUsed returns the ClientAppUsed field if non-nil, zero value otherwise.

### GetClientAppUsedOk

`func (o *MicrosoftGraphSignIn) GetClientAppUsedOk() (*string, bool)`

GetClientAppUsedOk returns a tuple with the ClientAppUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientAppUsed

`func (o *MicrosoftGraphSignIn) SetClientAppUsed(v string)`

SetClientAppUsed sets ClientAppUsed field to given value.

### HasClientAppUsed

`func (o *MicrosoftGraphSignIn) HasClientAppUsed() bool`

HasClientAppUsed returns a boolean if a field has been set.

### SetClientAppUsedNil

`func (o *MicrosoftGraphSignIn) SetClientAppUsedNil(b bool)`

 SetClientAppUsedNil sets the value for ClientAppUsed to be an explicit nil

### UnsetClientAppUsed
`func (o *MicrosoftGraphSignIn) UnsetClientAppUsed()`

UnsetClientAppUsed ensures that no value is present for ClientAppUsed, not even an explicit nil
### GetConditionalAccessStatus

`func (o *MicrosoftGraphSignIn) GetConditionalAccessStatus() AnyOfmicrosoftGraphConditionalAccessStatus`

GetConditionalAccessStatus returns the ConditionalAccessStatus field if non-nil, zero value otherwise.

### GetConditionalAccessStatusOk

`func (o *MicrosoftGraphSignIn) GetConditionalAccessStatusOk() (*AnyOfmicrosoftGraphConditionalAccessStatus, bool)`

GetConditionalAccessStatusOk returns a tuple with the ConditionalAccessStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionalAccessStatus

`func (o *MicrosoftGraphSignIn) SetConditionalAccessStatus(v AnyOfmicrosoftGraphConditionalAccessStatus)`

SetConditionalAccessStatus sets ConditionalAccessStatus field to given value.

### HasConditionalAccessStatus

`func (o *MicrosoftGraphSignIn) HasConditionalAccessStatus() bool`

HasConditionalAccessStatus returns a boolean if a field has been set.

### SetConditionalAccessStatusNil

`func (o *MicrosoftGraphSignIn) SetConditionalAccessStatusNil(b bool)`

 SetConditionalAccessStatusNil sets the value for ConditionalAccessStatus to be an explicit nil

### UnsetConditionalAccessStatus
`func (o *MicrosoftGraphSignIn) UnsetConditionalAccessStatus()`

UnsetConditionalAccessStatus ensures that no value is present for ConditionalAccessStatus, not even an explicit nil
### GetCorrelationId

`func (o *MicrosoftGraphSignIn) GetCorrelationId() string`

GetCorrelationId returns the CorrelationId field if non-nil, zero value otherwise.

### GetCorrelationIdOk

`func (o *MicrosoftGraphSignIn) GetCorrelationIdOk() (*string, bool)`

GetCorrelationIdOk returns a tuple with the CorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationId

`func (o *MicrosoftGraphSignIn) SetCorrelationId(v string)`

SetCorrelationId sets CorrelationId field to given value.

### HasCorrelationId

`func (o *MicrosoftGraphSignIn) HasCorrelationId() bool`

HasCorrelationId returns a boolean if a field has been set.

### SetCorrelationIdNil

`func (o *MicrosoftGraphSignIn) SetCorrelationIdNil(b bool)`

 SetCorrelationIdNil sets the value for CorrelationId to be an explicit nil

### UnsetCorrelationId
`func (o *MicrosoftGraphSignIn) UnsetCorrelationId()`

UnsetCorrelationId ensures that no value is present for CorrelationId, not even an explicit nil
### GetCreatedDateTime

`func (o *MicrosoftGraphSignIn) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphSignIn) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphSignIn) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *MicrosoftGraphSignIn) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### GetDeviceDetail

`func (o *MicrosoftGraphSignIn) GetDeviceDetail() AnyOfmicrosoftGraphDeviceDetail`

GetDeviceDetail returns the DeviceDetail field if non-nil, zero value otherwise.

### GetDeviceDetailOk

`func (o *MicrosoftGraphSignIn) GetDeviceDetailOk() (*AnyOfmicrosoftGraphDeviceDetail, bool)`

GetDeviceDetailOk returns a tuple with the DeviceDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceDetail

`func (o *MicrosoftGraphSignIn) SetDeviceDetail(v AnyOfmicrosoftGraphDeviceDetail)`

SetDeviceDetail sets DeviceDetail field to given value.

### HasDeviceDetail

`func (o *MicrosoftGraphSignIn) HasDeviceDetail() bool`

HasDeviceDetail returns a boolean if a field has been set.

### SetDeviceDetailNil

`func (o *MicrosoftGraphSignIn) SetDeviceDetailNil(b bool)`

 SetDeviceDetailNil sets the value for DeviceDetail to be an explicit nil

### UnsetDeviceDetail
`func (o *MicrosoftGraphSignIn) UnsetDeviceDetail()`

UnsetDeviceDetail ensures that no value is present for DeviceDetail, not even an explicit nil
### GetIpAddress

`func (o *MicrosoftGraphSignIn) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *MicrosoftGraphSignIn) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *MicrosoftGraphSignIn) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *MicrosoftGraphSignIn) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### SetIpAddressNil

`func (o *MicrosoftGraphSignIn) SetIpAddressNil(b bool)`

 SetIpAddressNil sets the value for IpAddress to be an explicit nil

### UnsetIpAddress
`func (o *MicrosoftGraphSignIn) UnsetIpAddress()`

UnsetIpAddress ensures that no value is present for IpAddress, not even an explicit nil
### GetIsInteractive

`func (o *MicrosoftGraphSignIn) GetIsInteractive() bool`

GetIsInteractive returns the IsInteractive field if non-nil, zero value otherwise.

### GetIsInteractiveOk

`func (o *MicrosoftGraphSignIn) GetIsInteractiveOk() (*bool, bool)`

GetIsInteractiveOk returns a tuple with the IsInteractive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInteractive

`func (o *MicrosoftGraphSignIn) SetIsInteractive(v bool)`

SetIsInteractive sets IsInteractive field to given value.

### HasIsInteractive

`func (o *MicrosoftGraphSignIn) HasIsInteractive() bool`

HasIsInteractive returns a boolean if a field has been set.

### SetIsInteractiveNil

`func (o *MicrosoftGraphSignIn) SetIsInteractiveNil(b bool)`

 SetIsInteractiveNil sets the value for IsInteractive to be an explicit nil

### UnsetIsInteractive
`func (o *MicrosoftGraphSignIn) UnsetIsInteractive()`

UnsetIsInteractive ensures that no value is present for IsInteractive, not even an explicit nil
### GetLocation

`func (o *MicrosoftGraphSignIn) GetLocation() AnyOfmicrosoftGraphSignInLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *MicrosoftGraphSignIn) GetLocationOk() (*AnyOfmicrosoftGraphSignInLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *MicrosoftGraphSignIn) SetLocation(v AnyOfmicrosoftGraphSignInLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *MicrosoftGraphSignIn) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *MicrosoftGraphSignIn) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *MicrosoftGraphSignIn) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetResourceDisplayName

`func (o *MicrosoftGraphSignIn) GetResourceDisplayName() string`

GetResourceDisplayName returns the ResourceDisplayName field if non-nil, zero value otherwise.

### GetResourceDisplayNameOk

`func (o *MicrosoftGraphSignIn) GetResourceDisplayNameOk() (*string, bool)`

GetResourceDisplayNameOk returns a tuple with the ResourceDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceDisplayName

`func (o *MicrosoftGraphSignIn) SetResourceDisplayName(v string)`

SetResourceDisplayName sets ResourceDisplayName field to given value.

### HasResourceDisplayName

`func (o *MicrosoftGraphSignIn) HasResourceDisplayName() bool`

HasResourceDisplayName returns a boolean if a field has been set.

### SetResourceDisplayNameNil

`func (o *MicrosoftGraphSignIn) SetResourceDisplayNameNil(b bool)`

 SetResourceDisplayNameNil sets the value for ResourceDisplayName to be an explicit nil

### UnsetResourceDisplayName
`func (o *MicrosoftGraphSignIn) UnsetResourceDisplayName()`

UnsetResourceDisplayName ensures that no value is present for ResourceDisplayName, not even an explicit nil
### GetResourceId

`func (o *MicrosoftGraphSignIn) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *MicrosoftGraphSignIn) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *MicrosoftGraphSignIn) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *MicrosoftGraphSignIn) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### SetResourceIdNil

`func (o *MicrosoftGraphSignIn) SetResourceIdNil(b bool)`

 SetResourceIdNil sets the value for ResourceId to be an explicit nil

### UnsetResourceId
`func (o *MicrosoftGraphSignIn) UnsetResourceId()`

UnsetResourceId ensures that no value is present for ResourceId, not even an explicit nil
### GetRiskDetail

`func (o *MicrosoftGraphSignIn) GetRiskDetail() AnyOfmicrosoftGraphRiskDetail`

GetRiskDetail returns the RiskDetail field if non-nil, zero value otherwise.

### GetRiskDetailOk

`func (o *MicrosoftGraphSignIn) GetRiskDetailOk() (*AnyOfmicrosoftGraphRiskDetail, bool)`

GetRiskDetailOk returns a tuple with the RiskDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskDetail

`func (o *MicrosoftGraphSignIn) SetRiskDetail(v AnyOfmicrosoftGraphRiskDetail)`

SetRiskDetail sets RiskDetail field to given value.

### HasRiskDetail

`func (o *MicrosoftGraphSignIn) HasRiskDetail() bool`

HasRiskDetail returns a boolean if a field has been set.

### SetRiskDetailNil

`func (o *MicrosoftGraphSignIn) SetRiskDetailNil(b bool)`

 SetRiskDetailNil sets the value for RiskDetail to be an explicit nil

### UnsetRiskDetail
`func (o *MicrosoftGraphSignIn) UnsetRiskDetail()`

UnsetRiskDetail ensures that no value is present for RiskDetail, not even an explicit nil
### GetRiskEventTypes

`func (o *MicrosoftGraphSignIn) GetRiskEventTypes() []*AnyOfmicrosoftGraphRiskEventType`

GetRiskEventTypes returns the RiskEventTypes field if non-nil, zero value otherwise.

### GetRiskEventTypesOk

`func (o *MicrosoftGraphSignIn) GetRiskEventTypesOk() (*[]*AnyOfmicrosoftGraphRiskEventType, bool)`

GetRiskEventTypesOk returns a tuple with the RiskEventTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskEventTypes

`func (o *MicrosoftGraphSignIn) SetRiskEventTypes(v []*AnyOfmicrosoftGraphRiskEventType)`

SetRiskEventTypes sets RiskEventTypes field to given value.

### HasRiskEventTypes

`func (o *MicrosoftGraphSignIn) HasRiskEventTypes() bool`

HasRiskEventTypes returns a boolean if a field has been set.

### GetRiskEventTypesV2

`func (o *MicrosoftGraphSignIn) GetRiskEventTypesV2() []*string`

GetRiskEventTypesV2 returns the RiskEventTypesV2 field if non-nil, zero value otherwise.

### GetRiskEventTypesV2Ok

`func (o *MicrosoftGraphSignIn) GetRiskEventTypesV2Ok() (*[]*string, bool)`

GetRiskEventTypesV2Ok returns a tuple with the RiskEventTypesV2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskEventTypesV2

`func (o *MicrosoftGraphSignIn) SetRiskEventTypesV2(v []*string)`

SetRiskEventTypesV2 sets RiskEventTypesV2 field to given value.

### HasRiskEventTypesV2

`func (o *MicrosoftGraphSignIn) HasRiskEventTypesV2() bool`

HasRiskEventTypesV2 returns a boolean if a field has been set.

### GetRiskLevelAggregated

`func (o *MicrosoftGraphSignIn) GetRiskLevelAggregated() AnyOfmicrosoftGraphRiskLevel`

GetRiskLevelAggregated returns the RiskLevelAggregated field if non-nil, zero value otherwise.

### GetRiskLevelAggregatedOk

`func (o *MicrosoftGraphSignIn) GetRiskLevelAggregatedOk() (*AnyOfmicrosoftGraphRiskLevel, bool)`

GetRiskLevelAggregatedOk returns a tuple with the RiskLevelAggregated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskLevelAggregated

`func (o *MicrosoftGraphSignIn) SetRiskLevelAggregated(v AnyOfmicrosoftGraphRiskLevel)`

SetRiskLevelAggregated sets RiskLevelAggregated field to given value.

### HasRiskLevelAggregated

`func (o *MicrosoftGraphSignIn) HasRiskLevelAggregated() bool`

HasRiskLevelAggregated returns a boolean if a field has been set.

### SetRiskLevelAggregatedNil

`func (o *MicrosoftGraphSignIn) SetRiskLevelAggregatedNil(b bool)`

 SetRiskLevelAggregatedNil sets the value for RiskLevelAggregated to be an explicit nil

### UnsetRiskLevelAggregated
`func (o *MicrosoftGraphSignIn) UnsetRiskLevelAggregated()`

UnsetRiskLevelAggregated ensures that no value is present for RiskLevelAggregated, not even an explicit nil
### GetRiskLevelDuringSignIn

`func (o *MicrosoftGraphSignIn) GetRiskLevelDuringSignIn() AnyOfmicrosoftGraphRiskLevel`

GetRiskLevelDuringSignIn returns the RiskLevelDuringSignIn field if non-nil, zero value otherwise.

### GetRiskLevelDuringSignInOk

`func (o *MicrosoftGraphSignIn) GetRiskLevelDuringSignInOk() (*AnyOfmicrosoftGraphRiskLevel, bool)`

GetRiskLevelDuringSignInOk returns a tuple with the RiskLevelDuringSignIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskLevelDuringSignIn

`func (o *MicrosoftGraphSignIn) SetRiskLevelDuringSignIn(v AnyOfmicrosoftGraphRiskLevel)`

SetRiskLevelDuringSignIn sets RiskLevelDuringSignIn field to given value.

### HasRiskLevelDuringSignIn

`func (o *MicrosoftGraphSignIn) HasRiskLevelDuringSignIn() bool`

HasRiskLevelDuringSignIn returns a boolean if a field has been set.

### SetRiskLevelDuringSignInNil

`func (o *MicrosoftGraphSignIn) SetRiskLevelDuringSignInNil(b bool)`

 SetRiskLevelDuringSignInNil sets the value for RiskLevelDuringSignIn to be an explicit nil

### UnsetRiskLevelDuringSignIn
`func (o *MicrosoftGraphSignIn) UnsetRiskLevelDuringSignIn()`

UnsetRiskLevelDuringSignIn ensures that no value is present for RiskLevelDuringSignIn, not even an explicit nil
### GetRiskState

`func (o *MicrosoftGraphSignIn) GetRiskState() AnyOfmicrosoftGraphRiskState`

GetRiskState returns the RiskState field if non-nil, zero value otherwise.

### GetRiskStateOk

`func (o *MicrosoftGraphSignIn) GetRiskStateOk() (*AnyOfmicrosoftGraphRiskState, bool)`

GetRiskStateOk returns a tuple with the RiskState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskState

`func (o *MicrosoftGraphSignIn) SetRiskState(v AnyOfmicrosoftGraphRiskState)`

SetRiskState sets RiskState field to given value.

### HasRiskState

`func (o *MicrosoftGraphSignIn) HasRiskState() bool`

HasRiskState returns a boolean if a field has been set.

### SetRiskStateNil

`func (o *MicrosoftGraphSignIn) SetRiskStateNil(b bool)`

 SetRiskStateNil sets the value for RiskState to be an explicit nil

### UnsetRiskState
`func (o *MicrosoftGraphSignIn) UnsetRiskState()`

UnsetRiskState ensures that no value is present for RiskState, not even an explicit nil
### GetStatus

`func (o *MicrosoftGraphSignIn) GetStatus() AnyOfmicrosoftGraphSignInStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MicrosoftGraphSignIn) GetStatusOk() (*AnyOfmicrosoftGraphSignInStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MicrosoftGraphSignIn) SetStatus(v AnyOfmicrosoftGraphSignInStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MicrosoftGraphSignIn) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *MicrosoftGraphSignIn) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *MicrosoftGraphSignIn) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetUserDisplayName

`func (o *MicrosoftGraphSignIn) GetUserDisplayName() string`

GetUserDisplayName returns the UserDisplayName field if non-nil, zero value otherwise.

### GetUserDisplayNameOk

`func (o *MicrosoftGraphSignIn) GetUserDisplayNameOk() (*string, bool)`

GetUserDisplayNameOk returns a tuple with the UserDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDisplayName

`func (o *MicrosoftGraphSignIn) SetUserDisplayName(v string)`

SetUserDisplayName sets UserDisplayName field to given value.

### HasUserDisplayName

`func (o *MicrosoftGraphSignIn) HasUserDisplayName() bool`

HasUserDisplayName returns a boolean if a field has been set.

### SetUserDisplayNameNil

`func (o *MicrosoftGraphSignIn) SetUserDisplayNameNil(b bool)`

 SetUserDisplayNameNil sets the value for UserDisplayName to be an explicit nil

### UnsetUserDisplayName
`func (o *MicrosoftGraphSignIn) UnsetUserDisplayName()`

UnsetUserDisplayName ensures that no value is present for UserDisplayName, not even an explicit nil
### GetUserId

`func (o *MicrosoftGraphSignIn) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *MicrosoftGraphSignIn) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *MicrosoftGraphSignIn) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *MicrosoftGraphSignIn) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetUserPrincipalName

`func (o *MicrosoftGraphSignIn) GetUserPrincipalName() string`

GetUserPrincipalName returns the UserPrincipalName field if non-nil, zero value otherwise.

### GetUserPrincipalNameOk

`func (o *MicrosoftGraphSignIn) GetUserPrincipalNameOk() (*string, bool)`

GetUserPrincipalNameOk returns a tuple with the UserPrincipalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPrincipalName

`func (o *MicrosoftGraphSignIn) SetUserPrincipalName(v string)`

SetUserPrincipalName sets UserPrincipalName field to given value.

### HasUserPrincipalName

`func (o *MicrosoftGraphSignIn) HasUserPrincipalName() bool`

HasUserPrincipalName returns a boolean if a field has been set.

### SetUserPrincipalNameNil

`func (o *MicrosoftGraphSignIn) SetUserPrincipalNameNil(b bool)`

 SetUserPrincipalNameNil sets the value for UserPrincipalName to be an explicit nil

### UnsetUserPrincipalName
`func (o *MicrosoftGraphSignIn) UnsetUserPrincipalName()`

UnsetUserPrincipalName ensures that no value is present for UserPrincipalName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


