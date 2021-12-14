# SignIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

### NewSignIn

`func NewSignIn() *SignIn`

NewSignIn instantiates a new SignIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignInWithDefaults

`func NewSignInWithDefaults() *SignIn`

NewSignInWithDefaults instantiates a new SignIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppDisplayName

`func (o *SignIn) GetAppDisplayName() string`

GetAppDisplayName returns the AppDisplayName field if non-nil, zero value otherwise.

### GetAppDisplayNameOk

`func (o *SignIn) GetAppDisplayNameOk() (*string, bool)`

GetAppDisplayNameOk returns a tuple with the AppDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppDisplayName

`func (o *SignIn) SetAppDisplayName(v string)`

SetAppDisplayName sets AppDisplayName field to given value.

### HasAppDisplayName

`func (o *SignIn) HasAppDisplayName() bool`

HasAppDisplayName returns a boolean if a field has been set.

### SetAppDisplayNameNil

`func (o *SignIn) SetAppDisplayNameNil(b bool)`

 SetAppDisplayNameNil sets the value for AppDisplayName to be an explicit nil

### UnsetAppDisplayName
`func (o *SignIn) UnsetAppDisplayName()`

UnsetAppDisplayName ensures that no value is present for AppDisplayName, not even an explicit nil
### GetAppId

`func (o *SignIn) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *SignIn) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *SignIn) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *SignIn) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### SetAppIdNil

`func (o *SignIn) SetAppIdNil(b bool)`

 SetAppIdNil sets the value for AppId to be an explicit nil

### UnsetAppId
`func (o *SignIn) UnsetAppId()`

UnsetAppId ensures that no value is present for AppId, not even an explicit nil
### GetAppliedConditionalAccessPolicies

`func (o *SignIn) GetAppliedConditionalAccessPolicies() []*AnyOfmicrosoftGraphAppliedConditionalAccessPolicy`

GetAppliedConditionalAccessPolicies returns the AppliedConditionalAccessPolicies field if non-nil, zero value otherwise.

### GetAppliedConditionalAccessPoliciesOk

`func (o *SignIn) GetAppliedConditionalAccessPoliciesOk() (*[]*AnyOfmicrosoftGraphAppliedConditionalAccessPolicy, bool)`

GetAppliedConditionalAccessPoliciesOk returns a tuple with the AppliedConditionalAccessPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedConditionalAccessPolicies

`func (o *SignIn) SetAppliedConditionalAccessPolicies(v []*AnyOfmicrosoftGraphAppliedConditionalAccessPolicy)`

SetAppliedConditionalAccessPolicies sets AppliedConditionalAccessPolicies field to given value.

### HasAppliedConditionalAccessPolicies

`func (o *SignIn) HasAppliedConditionalAccessPolicies() bool`

HasAppliedConditionalAccessPolicies returns a boolean if a field has been set.

### GetClientAppUsed

`func (o *SignIn) GetClientAppUsed() string`

GetClientAppUsed returns the ClientAppUsed field if non-nil, zero value otherwise.

### GetClientAppUsedOk

`func (o *SignIn) GetClientAppUsedOk() (*string, bool)`

GetClientAppUsedOk returns a tuple with the ClientAppUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientAppUsed

`func (o *SignIn) SetClientAppUsed(v string)`

SetClientAppUsed sets ClientAppUsed field to given value.

### HasClientAppUsed

`func (o *SignIn) HasClientAppUsed() bool`

HasClientAppUsed returns a boolean if a field has been set.

### SetClientAppUsedNil

`func (o *SignIn) SetClientAppUsedNil(b bool)`

 SetClientAppUsedNil sets the value for ClientAppUsed to be an explicit nil

### UnsetClientAppUsed
`func (o *SignIn) UnsetClientAppUsed()`

UnsetClientAppUsed ensures that no value is present for ClientAppUsed, not even an explicit nil
### GetConditionalAccessStatus

`func (o *SignIn) GetConditionalAccessStatus() AnyOfmicrosoftGraphConditionalAccessStatus`

GetConditionalAccessStatus returns the ConditionalAccessStatus field if non-nil, zero value otherwise.

### GetConditionalAccessStatusOk

`func (o *SignIn) GetConditionalAccessStatusOk() (*AnyOfmicrosoftGraphConditionalAccessStatus, bool)`

GetConditionalAccessStatusOk returns a tuple with the ConditionalAccessStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionalAccessStatus

`func (o *SignIn) SetConditionalAccessStatus(v AnyOfmicrosoftGraphConditionalAccessStatus)`

SetConditionalAccessStatus sets ConditionalAccessStatus field to given value.

### HasConditionalAccessStatus

`func (o *SignIn) HasConditionalAccessStatus() bool`

HasConditionalAccessStatus returns a boolean if a field has been set.

### SetConditionalAccessStatusNil

`func (o *SignIn) SetConditionalAccessStatusNil(b bool)`

 SetConditionalAccessStatusNil sets the value for ConditionalAccessStatus to be an explicit nil

### UnsetConditionalAccessStatus
`func (o *SignIn) UnsetConditionalAccessStatus()`

UnsetConditionalAccessStatus ensures that no value is present for ConditionalAccessStatus, not even an explicit nil
### GetCorrelationId

`func (o *SignIn) GetCorrelationId() string`

GetCorrelationId returns the CorrelationId field if non-nil, zero value otherwise.

### GetCorrelationIdOk

`func (o *SignIn) GetCorrelationIdOk() (*string, bool)`

GetCorrelationIdOk returns a tuple with the CorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationId

`func (o *SignIn) SetCorrelationId(v string)`

SetCorrelationId sets CorrelationId field to given value.

### HasCorrelationId

`func (o *SignIn) HasCorrelationId() bool`

HasCorrelationId returns a boolean if a field has been set.

### SetCorrelationIdNil

`func (o *SignIn) SetCorrelationIdNil(b bool)`

 SetCorrelationIdNil sets the value for CorrelationId to be an explicit nil

### UnsetCorrelationId
`func (o *SignIn) UnsetCorrelationId()`

UnsetCorrelationId ensures that no value is present for CorrelationId, not even an explicit nil
### GetCreatedDateTime

`func (o *SignIn) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *SignIn) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *SignIn) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *SignIn) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### GetDeviceDetail

`func (o *SignIn) GetDeviceDetail() AnyOfmicrosoftGraphDeviceDetail`

GetDeviceDetail returns the DeviceDetail field if non-nil, zero value otherwise.

### GetDeviceDetailOk

`func (o *SignIn) GetDeviceDetailOk() (*AnyOfmicrosoftGraphDeviceDetail, bool)`

GetDeviceDetailOk returns a tuple with the DeviceDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceDetail

`func (o *SignIn) SetDeviceDetail(v AnyOfmicrosoftGraphDeviceDetail)`

SetDeviceDetail sets DeviceDetail field to given value.

### HasDeviceDetail

`func (o *SignIn) HasDeviceDetail() bool`

HasDeviceDetail returns a boolean if a field has been set.

### SetDeviceDetailNil

`func (o *SignIn) SetDeviceDetailNil(b bool)`

 SetDeviceDetailNil sets the value for DeviceDetail to be an explicit nil

### UnsetDeviceDetail
`func (o *SignIn) UnsetDeviceDetail()`

UnsetDeviceDetail ensures that no value is present for DeviceDetail, not even an explicit nil
### GetIpAddress

`func (o *SignIn) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *SignIn) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *SignIn) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *SignIn) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### SetIpAddressNil

`func (o *SignIn) SetIpAddressNil(b bool)`

 SetIpAddressNil sets the value for IpAddress to be an explicit nil

### UnsetIpAddress
`func (o *SignIn) UnsetIpAddress()`

UnsetIpAddress ensures that no value is present for IpAddress, not even an explicit nil
### GetIsInteractive

`func (o *SignIn) GetIsInteractive() bool`

GetIsInteractive returns the IsInteractive field if non-nil, zero value otherwise.

### GetIsInteractiveOk

`func (o *SignIn) GetIsInteractiveOk() (*bool, bool)`

GetIsInteractiveOk returns a tuple with the IsInteractive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInteractive

`func (o *SignIn) SetIsInteractive(v bool)`

SetIsInteractive sets IsInteractive field to given value.

### HasIsInteractive

`func (o *SignIn) HasIsInteractive() bool`

HasIsInteractive returns a boolean if a field has been set.

### SetIsInteractiveNil

`func (o *SignIn) SetIsInteractiveNil(b bool)`

 SetIsInteractiveNil sets the value for IsInteractive to be an explicit nil

### UnsetIsInteractive
`func (o *SignIn) UnsetIsInteractive()`

UnsetIsInteractive ensures that no value is present for IsInteractive, not even an explicit nil
### GetLocation

`func (o *SignIn) GetLocation() AnyOfmicrosoftGraphSignInLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *SignIn) GetLocationOk() (*AnyOfmicrosoftGraphSignInLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *SignIn) SetLocation(v AnyOfmicrosoftGraphSignInLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *SignIn) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *SignIn) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *SignIn) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetResourceDisplayName

`func (o *SignIn) GetResourceDisplayName() string`

GetResourceDisplayName returns the ResourceDisplayName field if non-nil, zero value otherwise.

### GetResourceDisplayNameOk

`func (o *SignIn) GetResourceDisplayNameOk() (*string, bool)`

GetResourceDisplayNameOk returns a tuple with the ResourceDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceDisplayName

`func (o *SignIn) SetResourceDisplayName(v string)`

SetResourceDisplayName sets ResourceDisplayName field to given value.

### HasResourceDisplayName

`func (o *SignIn) HasResourceDisplayName() bool`

HasResourceDisplayName returns a boolean if a field has been set.

### SetResourceDisplayNameNil

`func (o *SignIn) SetResourceDisplayNameNil(b bool)`

 SetResourceDisplayNameNil sets the value for ResourceDisplayName to be an explicit nil

### UnsetResourceDisplayName
`func (o *SignIn) UnsetResourceDisplayName()`

UnsetResourceDisplayName ensures that no value is present for ResourceDisplayName, not even an explicit nil
### GetResourceId

`func (o *SignIn) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *SignIn) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *SignIn) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *SignIn) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### SetResourceIdNil

`func (o *SignIn) SetResourceIdNil(b bool)`

 SetResourceIdNil sets the value for ResourceId to be an explicit nil

### UnsetResourceId
`func (o *SignIn) UnsetResourceId()`

UnsetResourceId ensures that no value is present for ResourceId, not even an explicit nil
### GetRiskDetail

`func (o *SignIn) GetRiskDetail() AnyOfmicrosoftGraphRiskDetail`

GetRiskDetail returns the RiskDetail field if non-nil, zero value otherwise.

### GetRiskDetailOk

`func (o *SignIn) GetRiskDetailOk() (*AnyOfmicrosoftGraphRiskDetail, bool)`

GetRiskDetailOk returns a tuple with the RiskDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskDetail

`func (o *SignIn) SetRiskDetail(v AnyOfmicrosoftGraphRiskDetail)`

SetRiskDetail sets RiskDetail field to given value.

### HasRiskDetail

`func (o *SignIn) HasRiskDetail() bool`

HasRiskDetail returns a boolean if a field has been set.

### SetRiskDetailNil

`func (o *SignIn) SetRiskDetailNil(b bool)`

 SetRiskDetailNil sets the value for RiskDetail to be an explicit nil

### UnsetRiskDetail
`func (o *SignIn) UnsetRiskDetail()`

UnsetRiskDetail ensures that no value is present for RiskDetail, not even an explicit nil
### GetRiskEventTypes

`func (o *SignIn) GetRiskEventTypes() []*AnyOfmicrosoftGraphRiskEventType`

GetRiskEventTypes returns the RiskEventTypes field if non-nil, zero value otherwise.

### GetRiskEventTypesOk

`func (o *SignIn) GetRiskEventTypesOk() (*[]*AnyOfmicrosoftGraphRiskEventType, bool)`

GetRiskEventTypesOk returns a tuple with the RiskEventTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskEventTypes

`func (o *SignIn) SetRiskEventTypes(v []*AnyOfmicrosoftGraphRiskEventType)`

SetRiskEventTypes sets RiskEventTypes field to given value.

### HasRiskEventTypes

`func (o *SignIn) HasRiskEventTypes() bool`

HasRiskEventTypes returns a boolean if a field has been set.

### GetRiskEventTypesV2

`func (o *SignIn) GetRiskEventTypesV2() []*string`

GetRiskEventTypesV2 returns the RiskEventTypesV2 field if non-nil, zero value otherwise.

### GetRiskEventTypesV2Ok

`func (o *SignIn) GetRiskEventTypesV2Ok() (*[]*string, bool)`

GetRiskEventTypesV2Ok returns a tuple with the RiskEventTypesV2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskEventTypesV2

`func (o *SignIn) SetRiskEventTypesV2(v []*string)`

SetRiskEventTypesV2 sets RiskEventTypesV2 field to given value.

### HasRiskEventTypesV2

`func (o *SignIn) HasRiskEventTypesV2() bool`

HasRiskEventTypesV2 returns a boolean if a field has been set.

### GetRiskLevelAggregated

`func (o *SignIn) GetRiskLevelAggregated() AnyOfmicrosoftGraphRiskLevel`

GetRiskLevelAggregated returns the RiskLevelAggregated field if non-nil, zero value otherwise.

### GetRiskLevelAggregatedOk

`func (o *SignIn) GetRiskLevelAggregatedOk() (*AnyOfmicrosoftGraphRiskLevel, bool)`

GetRiskLevelAggregatedOk returns a tuple with the RiskLevelAggregated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskLevelAggregated

`func (o *SignIn) SetRiskLevelAggregated(v AnyOfmicrosoftGraphRiskLevel)`

SetRiskLevelAggregated sets RiskLevelAggregated field to given value.

### HasRiskLevelAggregated

`func (o *SignIn) HasRiskLevelAggregated() bool`

HasRiskLevelAggregated returns a boolean if a field has been set.

### SetRiskLevelAggregatedNil

`func (o *SignIn) SetRiskLevelAggregatedNil(b bool)`

 SetRiskLevelAggregatedNil sets the value for RiskLevelAggregated to be an explicit nil

### UnsetRiskLevelAggregated
`func (o *SignIn) UnsetRiskLevelAggregated()`

UnsetRiskLevelAggregated ensures that no value is present for RiskLevelAggregated, not even an explicit nil
### GetRiskLevelDuringSignIn

`func (o *SignIn) GetRiskLevelDuringSignIn() AnyOfmicrosoftGraphRiskLevel`

GetRiskLevelDuringSignIn returns the RiskLevelDuringSignIn field if non-nil, zero value otherwise.

### GetRiskLevelDuringSignInOk

`func (o *SignIn) GetRiskLevelDuringSignInOk() (*AnyOfmicrosoftGraphRiskLevel, bool)`

GetRiskLevelDuringSignInOk returns a tuple with the RiskLevelDuringSignIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskLevelDuringSignIn

`func (o *SignIn) SetRiskLevelDuringSignIn(v AnyOfmicrosoftGraphRiskLevel)`

SetRiskLevelDuringSignIn sets RiskLevelDuringSignIn field to given value.

### HasRiskLevelDuringSignIn

`func (o *SignIn) HasRiskLevelDuringSignIn() bool`

HasRiskLevelDuringSignIn returns a boolean if a field has been set.

### SetRiskLevelDuringSignInNil

`func (o *SignIn) SetRiskLevelDuringSignInNil(b bool)`

 SetRiskLevelDuringSignInNil sets the value for RiskLevelDuringSignIn to be an explicit nil

### UnsetRiskLevelDuringSignIn
`func (o *SignIn) UnsetRiskLevelDuringSignIn()`

UnsetRiskLevelDuringSignIn ensures that no value is present for RiskLevelDuringSignIn, not even an explicit nil
### GetRiskState

`func (o *SignIn) GetRiskState() AnyOfmicrosoftGraphRiskState`

GetRiskState returns the RiskState field if non-nil, zero value otherwise.

### GetRiskStateOk

`func (o *SignIn) GetRiskStateOk() (*AnyOfmicrosoftGraphRiskState, bool)`

GetRiskStateOk returns a tuple with the RiskState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskState

`func (o *SignIn) SetRiskState(v AnyOfmicrosoftGraphRiskState)`

SetRiskState sets RiskState field to given value.

### HasRiskState

`func (o *SignIn) HasRiskState() bool`

HasRiskState returns a boolean if a field has been set.

### SetRiskStateNil

`func (o *SignIn) SetRiskStateNil(b bool)`

 SetRiskStateNil sets the value for RiskState to be an explicit nil

### UnsetRiskState
`func (o *SignIn) UnsetRiskState()`

UnsetRiskState ensures that no value is present for RiskState, not even an explicit nil
### GetStatus

`func (o *SignIn) GetStatus() AnyOfmicrosoftGraphSignInStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SignIn) GetStatusOk() (*AnyOfmicrosoftGraphSignInStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SignIn) SetStatus(v AnyOfmicrosoftGraphSignInStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SignIn) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *SignIn) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *SignIn) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetUserDisplayName

`func (o *SignIn) GetUserDisplayName() string`

GetUserDisplayName returns the UserDisplayName field if non-nil, zero value otherwise.

### GetUserDisplayNameOk

`func (o *SignIn) GetUserDisplayNameOk() (*string, bool)`

GetUserDisplayNameOk returns a tuple with the UserDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDisplayName

`func (o *SignIn) SetUserDisplayName(v string)`

SetUserDisplayName sets UserDisplayName field to given value.

### HasUserDisplayName

`func (o *SignIn) HasUserDisplayName() bool`

HasUserDisplayName returns a boolean if a field has been set.

### SetUserDisplayNameNil

`func (o *SignIn) SetUserDisplayNameNil(b bool)`

 SetUserDisplayNameNil sets the value for UserDisplayName to be an explicit nil

### UnsetUserDisplayName
`func (o *SignIn) UnsetUserDisplayName()`

UnsetUserDisplayName ensures that no value is present for UserDisplayName, not even an explicit nil
### GetUserId

`func (o *SignIn) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *SignIn) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *SignIn) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *SignIn) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetUserPrincipalName

`func (o *SignIn) GetUserPrincipalName() string`

GetUserPrincipalName returns the UserPrincipalName field if non-nil, zero value otherwise.

### GetUserPrincipalNameOk

`func (o *SignIn) GetUserPrincipalNameOk() (*string, bool)`

GetUserPrincipalNameOk returns a tuple with the UserPrincipalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPrincipalName

`func (o *SignIn) SetUserPrincipalName(v string)`

SetUserPrincipalName sets UserPrincipalName field to given value.

### HasUserPrincipalName

`func (o *SignIn) HasUserPrincipalName() bool`

HasUserPrincipalName returns a boolean if a field has been set.

### SetUserPrincipalNameNil

`func (o *SignIn) SetUserPrincipalNameNil(b bool)`

 SetUserPrincipalNameNil sets the value for UserPrincipalName to be an explicit nil

### UnsetUserPrincipalName
`func (o *SignIn) UnsetUserPrincipalName()`

UnsetUserPrincipalName ensures that no value is present for UserPrincipalName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


