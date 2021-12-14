# MicrosoftGraphIdentityGovernance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessReviews** | Pointer to [**AnyOfmicrosoftGraphAccessReviewSet**](anyOf&lt;microsoft.graph.accessReviewSet&gt;.md) |  | [optional] 
**AppConsent** | Pointer to [**AnyOfmicrosoftGraphAppConsentApprovalRoute**](anyOf&lt;microsoft.graph.appConsentApprovalRoute&gt;.md) |  | [optional] 
**TermsOfUse** | Pointer to [**AnyOfmicrosoftGraphTermsOfUseContainer**](anyOf&lt;microsoft.graph.termsOfUseContainer&gt;.md) |  | [optional] 
**EntitlementManagement** | Pointer to [**AnyOfmicrosoftGraphEntitlementManagement**](anyOf&lt;microsoft.graph.entitlementManagement&gt;.md) |  | [optional] 

## Methods

### NewMicrosoftGraphIdentityGovernance

`func NewMicrosoftGraphIdentityGovernance() *MicrosoftGraphIdentityGovernance`

NewMicrosoftGraphIdentityGovernance instantiates a new MicrosoftGraphIdentityGovernance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphIdentityGovernanceWithDefaults

`func NewMicrosoftGraphIdentityGovernanceWithDefaults() *MicrosoftGraphIdentityGovernance`

NewMicrosoftGraphIdentityGovernanceWithDefaults instantiates a new MicrosoftGraphIdentityGovernance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessReviews

`func (o *MicrosoftGraphIdentityGovernance) GetAccessReviews() AnyOfmicrosoftGraphAccessReviewSet`

GetAccessReviews returns the AccessReviews field if non-nil, zero value otherwise.

### GetAccessReviewsOk

`func (o *MicrosoftGraphIdentityGovernance) GetAccessReviewsOk() (*AnyOfmicrosoftGraphAccessReviewSet, bool)`

GetAccessReviewsOk returns a tuple with the AccessReviews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessReviews

`func (o *MicrosoftGraphIdentityGovernance) SetAccessReviews(v AnyOfmicrosoftGraphAccessReviewSet)`

SetAccessReviews sets AccessReviews field to given value.

### HasAccessReviews

`func (o *MicrosoftGraphIdentityGovernance) HasAccessReviews() bool`

HasAccessReviews returns a boolean if a field has been set.

### SetAccessReviewsNil

`func (o *MicrosoftGraphIdentityGovernance) SetAccessReviewsNil(b bool)`

 SetAccessReviewsNil sets the value for AccessReviews to be an explicit nil

### UnsetAccessReviews
`func (o *MicrosoftGraphIdentityGovernance) UnsetAccessReviews()`

UnsetAccessReviews ensures that no value is present for AccessReviews, not even an explicit nil
### GetAppConsent

`func (o *MicrosoftGraphIdentityGovernance) GetAppConsent() AnyOfmicrosoftGraphAppConsentApprovalRoute`

GetAppConsent returns the AppConsent field if non-nil, zero value otherwise.

### GetAppConsentOk

`func (o *MicrosoftGraphIdentityGovernance) GetAppConsentOk() (*AnyOfmicrosoftGraphAppConsentApprovalRoute, bool)`

GetAppConsentOk returns a tuple with the AppConsent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppConsent

`func (o *MicrosoftGraphIdentityGovernance) SetAppConsent(v AnyOfmicrosoftGraphAppConsentApprovalRoute)`

SetAppConsent sets AppConsent field to given value.

### HasAppConsent

`func (o *MicrosoftGraphIdentityGovernance) HasAppConsent() bool`

HasAppConsent returns a boolean if a field has been set.

### SetAppConsentNil

`func (o *MicrosoftGraphIdentityGovernance) SetAppConsentNil(b bool)`

 SetAppConsentNil sets the value for AppConsent to be an explicit nil

### UnsetAppConsent
`func (o *MicrosoftGraphIdentityGovernance) UnsetAppConsent()`

UnsetAppConsent ensures that no value is present for AppConsent, not even an explicit nil
### GetTermsOfUse

`func (o *MicrosoftGraphIdentityGovernance) GetTermsOfUse() AnyOfmicrosoftGraphTermsOfUseContainer`

GetTermsOfUse returns the TermsOfUse field if non-nil, zero value otherwise.

### GetTermsOfUseOk

`func (o *MicrosoftGraphIdentityGovernance) GetTermsOfUseOk() (*AnyOfmicrosoftGraphTermsOfUseContainer, bool)`

GetTermsOfUseOk returns a tuple with the TermsOfUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermsOfUse

`func (o *MicrosoftGraphIdentityGovernance) SetTermsOfUse(v AnyOfmicrosoftGraphTermsOfUseContainer)`

SetTermsOfUse sets TermsOfUse field to given value.

### HasTermsOfUse

`func (o *MicrosoftGraphIdentityGovernance) HasTermsOfUse() bool`

HasTermsOfUse returns a boolean if a field has been set.

### SetTermsOfUseNil

`func (o *MicrosoftGraphIdentityGovernance) SetTermsOfUseNil(b bool)`

 SetTermsOfUseNil sets the value for TermsOfUse to be an explicit nil

### UnsetTermsOfUse
`func (o *MicrosoftGraphIdentityGovernance) UnsetTermsOfUse()`

UnsetTermsOfUse ensures that no value is present for TermsOfUse, not even an explicit nil
### GetEntitlementManagement

`func (o *MicrosoftGraphIdentityGovernance) GetEntitlementManagement() AnyOfmicrosoftGraphEntitlementManagement`

GetEntitlementManagement returns the EntitlementManagement field if non-nil, zero value otherwise.

### GetEntitlementManagementOk

`func (o *MicrosoftGraphIdentityGovernance) GetEntitlementManagementOk() (*AnyOfmicrosoftGraphEntitlementManagement, bool)`

GetEntitlementManagementOk returns a tuple with the EntitlementManagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementManagement

`func (o *MicrosoftGraphIdentityGovernance) SetEntitlementManagement(v AnyOfmicrosoftGraphEntitlementManagement)`

SetEntitlementManagement sets EntitlementManagement field to given value.

### HasEntitlementManagement

`func (o *MicrosoftGraphIdentityGovernance) HasEntitlementManagement() bool`

HasEntitlementManagement returns a boolean if a field has been set.

### SetEntitlementManagementNil

`func (o *MicrosoftGraphIdentityGovernance) SetEntitlementManagementNil(b bool)`

 SetEntitlementManagementNil sets the value for EntitlementManagement to be an explicit nil

### UnsetEntitlementManagement
`func (o *MicrosoftGraphIdentityGovernance) UnsetEntitlementManagement()`

UnsetEntitlementManagement ensures that no value is present for EntitlementManagement, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


