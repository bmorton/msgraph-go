# AuditLogRoot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DirectoryAudits** | Pointer to [**[]MicrosoftGraphDirectoryAudit**](MicrosoftGraphDirectoryAudit.md) | Read-only. Nullable. | [optional] 
**Provisioning** | Pointer to [**[]MicrosoftGraphProvisioningObjectSummary**](MicrosoftGraphProvisioningObjectSummary.md) |  | [optional] 
**RestrictedSignIns** | Pointer to [**[]MicrosoftGraphRestrictedSignIn**](MicrosoftGraphRestrictedSignIn.md) |  | [optional] 
**SignIns** | Pointer to [**[]MicrosoftGraphSignIn**](MicrosoftGraphSignIn.md) | Read-only. Nullable. | [optional] 

## Methods

### NewAuditLogRoot

`func NewAuditLogRoot() *AuditLogRoot`

NewAuditLogRoot instantiates a new AuditLogRoot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditLogRootWithDefaults

`func NewAuditLogRootWithDefaults() *AuditLogRoot`

NewAuditLogRootWithDefaults instantiates a new AuditLogRoot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDirectoryAudits

`func (o *AuditLogRoot) GetDirectoryAudits() []MicrosoftGraphDirectoryAudit`

GetDirectoryAudits returns the DirectoryAudits field if non-nil, zero value otherwise.

### GetDirectoryAuditsOk

`func (o *AuditLogRoot) GetDirectoryAuditsOk() (*[]MicrosoftGraphDirectoryAudit, bool)`

GetDirectoryAuditsOk returns a tuple with the DirectoryAudits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryAudits

`func (o *AuditLogRoot) SetDirectoryAudits(v []MicrosoftGraphDirectoryAudit)`

SetDirectoryAudits sets DirectoryAudits field to given value.

### HasDirectoryAudits

`func (o *AuditLogRoot) HasDirectoryAudits() bool`

HasDirectoryAudits returns a boolean if a field has been set.

### GetProvisioning

`func (o *AuditLogRoot) GetProvisioning() []MicrosoftGraphProvisioningObjectSummary`

GetProvisioning returns the Provisioning field if non-nil, zero value otherwise.

### GetProvisioningOk

`func (o *AuditLogRoot) GetProvisioningOk() (*[]MicrosoftGraphProvisioningObjectSummary, bool)`

GetProvisioningOk returns a tuple with the Provisioning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioning

`func (o *AuditLogRoot) SetProvisioning(v []MicrosoftGraphProvisioningObjectSummary)`

SetProvisioning sets Provisioning field to given value.

### HasProvisioning

`func (o *AuditLogRoot) HasProvisioning() bool`

HasProvisioning returns a boolean if a field has been set.

### GetRestrictedSignIns

`func (o *AuditLogRoot) GetRestrictedSignIns() []MicrosoftGraphRestrictedSignIn`

GetRestrictedSignIns returns the RestrictedSignIns field if non-nil, zero value otherwise.

### GetRestrictedSignInsOk

`func (o *AuditLogRoot) GetRestrictedSignInsOk() (*[]MicrosoftGraphRestrictedSignIn, bool)`

GetRestrictedSignInsOk returns a tuple with the RestrictedSignIns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictedSignIns

`func (o *AuditLogRoot) SetRestrictedSignIns(v []MicrosoftGraphRestrictedSignIn)`

SetRestrictedSignIns sets RestrictedSignIns field to given value.

### HasRestrictedSignIns

`func (o *AuditLogRoot) HasRestrictedSignIns() bool`

HasRestrictedSignIns returns a boolean if a field has been set.

### GetSignIns

`func (o *AuditLogRoot) GetSignIns() []MicrosoftGraphSignIn`

GetSignIns returns the SignIns field if non-nil, zero value otherwise.

### GetSignInsOk

`func (o *AuditLogRoot) GetSignInsOk() (*[]MicrosoftGraphSignIn, bool)`

GetSignInsOk returns a tuple with the SignIns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignIns

`func (o *AuditLogRoot) SetSignIns(v []MicrosoftGraphSignIn)`

SetSignIns sets SignIns field to given value.

### HasSignIns

`func (o *AuditLogRoot) HasSignIns() bool`

HasSignIns returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


