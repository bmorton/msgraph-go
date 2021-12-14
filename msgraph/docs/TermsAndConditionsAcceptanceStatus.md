# TermsAndConditionsAcceptanceStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcceptedDateTime** | Pointer to **time.Time** | DateTime when the terms were last accepted by the user. | [optional] 
**AcceptedVersion** | Pointer to **int32** | Most recent version number of the T&amp;C accepted by the user. | [optional] 
**UserDisplayName** | Pointer to **NullableString** | Display name of the user whose acceptance the entity represents. | [optional] 
**UserPrincipalName** | Pointer to **NullableString** | The userPrincipalName of the User that accepted the term. | [optional] 
**TermsAndConditions** | Pointer to [**AnyOfmicrosoftGraphTermsAndConditions**](anyOf&lt;microsoft.graph.termsAndConditions&gt;.md) | Navigation link to the terms and conditions that are assigned. | [optional] 

## Methods

### NewTermsAndConditionsAcceptanceStatus

`func NewTermsAndConditionsAcceptanceStatus() *TermsAndConditionsAcceptanceStatus`

NewTermsAndConditionsAcceptanceStatus instantiates a new TermsAndConditionsAcceptanceStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTermsAndConditionsAcceptanceStatusWithDefaults

`func NewTermsAndConditionsAcceptanceStatusWithDefaults() *TermsAndConditionsAcceptanceStatus`

NewTermsAndConditionsAcceptanceStatusWithDefaults instantiates a new TermsAndConditionsAcceptanceStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcceptedDateTime

`func (o *TermsAndConditionsAcceptanceStatus) GetAcceptedDateTime() time.Time`

GetAcceptedDateTime returns the AcceptedDateTime field if non-nil, zero value otherwise.

### GetAcceptedDateTimeOk

`func (o *TermsAndConditionsAcceptanceStatus) GetAcceptedDateTimeOk() (*time.Time, bool)`

GetAcceptedDateTimeOk returns a tuple with the AcceptedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedDateTime

`func (o *TermsAndConditionsAcceptanceStatus) SetAcceptedDateTime(v time.Time)`

SetAcceptedDateTime sets AcceptedDateTime field to given value.

### HasAcceptedDateTime

`func (o *TermsAndConditionsAcceptanceStatus) HasAcceptedDateTime() bool`

HasAcceptedDateTime returns a boolean if a field has been set.

### GetAcceptedVersion

`func (o *TermsAndConditionsAcceptanceStatus) GetAcceptedVersion() int32`

GetAcceptedVersion returns the AcceptedVersion field if non-nil, zero value otherwise.

### GetAcceptedVersionOk

`func (o *TermsAndConditionsAcceptanceStatus) GetAcceptedVersionOk() (*int32, bool)`

GetAcceptedVersionOk returns a tuple with the AcceptedVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedVersion

`func (o *TermsAndConditionsAcceptanceStatus) SetAcceptedVersion(v int32)`

SetAcceptedVersion sets AcceptedVersion field to given value.

### HasAcceptedVersion

`func (o *TermsAndConditionsAcceptanceStatus) HasAcceptedVersion() bool`

HasAcceptedVersion returns a boolean if a field has been set.

### GetUserDisplayName

`func (o *TermsAndConditionsAcceptanceStatus) GetUserDisplayName() string`

GetUserDisplayName returns the UserDisplayName field if non-nil, zero value otherwise.

### GetUserDisplayNameOk

`func (o *TermsAndConditionsAcceptanceStatus) GetUserDisplayNameOk() (*string, bool)`

GetUserDisplayNameOk returns a tuple with the UserDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDisplayName

`func (o *TermsAndConditionsAcceptanceStatus) SetUserDisplayName(v string)`

SetUserDisplayName sets UserDisplayName field to given value.

### HasUserDisplayName

`func (o *TermsAndConditionsAcceptanceStatus) HasUserDisplayName() bool`

HasUserDisplayName returns a boolean if a field has been set.

### SetUserDisplayNameNil

`func (o *TermsAndConditionsAcceptanceStatus) SetUserDisplayNameNil(b bool)`

 SetUserDisplayNameNil sets the value for UserDisplayName to be an explicit nil

### UnsetUserDisplayName
`func (o *TermsAndConditionsAcceptanceStatus) UnsetUserDisplayName()`

UnsetUserDisplayName ensures that no value is present for UserDisplayName, not even an explicit nil
### GetUserPrincipalName

`func (o *TermsAndConditionsAcceptanceStatus) GetUserPrincipalName() string`

GetUserPrincipalName returns the UserPrincipalName field if non-nil, zero value otherwise.

### GetUserPrincipalNameOk

`func (o *TermsAndConditionsAcceptanceStatus) GetUserPrincipalNameOk() (*string, bool)`

GetUserPrincipalNameOk returns a tuple with the UserPrincipalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPrincipalName

`func (o *TermsAndConditionsAcceptanceStatus) SetUserPrincipalName(v string)`

SetUserPrincipalName sets UserPrincipalName field to given value.

### HasUserPrincipalName

`func (o *TermsAndConditionsAcceptanceStatus) HasUserPrincipalName() bool`

HasUserPrincipalName returns a boolean if a field has been set.

### SetUserPrincipalNameNil

`func (o *TermsAndConditionsAcceptanceStatus) SetUserPrincipalNameNil(b bool)`

 SetUserPrincipalNameNil sets the value for UserPrincipalName to be an explicit nil

### UnsetUserPrincipalName
`func (o *TermsAndConditionsAcceptanceStatus) UnsetUserPrincipalName()`

UnsetUserPrincipalName ensures that no value is present for UserPrincipalName, not even an explicit nil
### GetTermsAndConditions

`func (o *TermsAndConditionsAcceptanceStatus) GetTermsAndConditions() AnyOfmicrosoftGraphTermsAndConditions`

GetTermsAndConditions returns the TermsAndConditions field if non-nil, zero value otherwise.

### GetTermsAndConditionsOk

`func (o *TermsAndConditionsAcceptanceStatus) GetTermsAndConditionsOk() (*AnyOfmicrosoftGraphTermsAndConditions, bool)`

GetTermsAndConditionsOk returns a tuple with the TermsAndConditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermsAndConditions

`func (o *TermsAndConditionsAcceptanceStatus) SetTermsAndConditions(v AnyOfmicrosoftGraphTermsAndConditions)`

SetTermsAndConditions sets TermsAndConditions field to given value.

### HasTermsAndConditions

`func (o *TermsAndConditionsAcceptanceStatus) HasTermsAndConditions() bool`

HasTermsAndConditions returns a boolean if a field has been set.

### SetTermsAndConditionsNil

`func (o *TermsAndConditionsAcceptanceStatus) SetTermsAndConditionsNil(b bool)`

 SetTermsAndConditionsNil sets the value for TermsAndConditions to be an explicit nil

### UnsetTermsAndConditions
`func (o *TermsAndConditionsAcceptanceStatus) UnsetTermsAndConditions()`

UnsetTermsAndConditions ensures that no value is present for TermsAndConditions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


