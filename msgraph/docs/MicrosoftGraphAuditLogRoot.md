# MicrosoftGraphAuditLogRoot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**DirectoryAudits** | Pointer to [**[]MicrosoftGraphDirectoryAudit**](MicrosoftGraphDirectoryAudit.md) | Read-only. Nullable. | [optional] 
**Provisioning** | Pointer to [**[]MicrosoftGraphProvisioningObjectSummary**](MicrosoftGraphProvisioningObjectSummary.md) |  | [optional] 
**RestrictedSignIns** | Pointer to [**[]MicrosoftGraphRestrictedSignIn**](MicrosoftGraphRestrictedSignIn.md) |  | [optional] 
**SignIns** | Pointer to [**[]MicrosoftGraphSignIn**](MicrosoftGraphSignIn.md) | Read-only. Nullable. | [optional] 

## Methods

### NewMicrosoftGraphAuditLogRoot

`func NewMicrosoftGraphAuditLogRoot() *MicrosoftGraphAuditLogRoot`

NewMicrosoftGraphAuditLogRoot instantiates a new MicrosoftGraphAuditLogRoot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphAuditLogRootWithDefaults

`func NewMicrosoftGraphAuditLogRootWithDefaults() *MicrosoftGraphAuditLogRoot`

NewMicrosoftGraphAuditLogRootWithDefaults instantiates a new MicrosoftGraphAuditLogRoot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphAuditLogRoot) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphAuditLogRoot) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphAuditLogRoot) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphAuditLogRoot) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDirectoryAudits

`func (o *MicrosoftGraphAuditLogRoot) GetDirectoryAudits() []MicrosoftGraphDirectoryAudit`

GetDirectoryAudits returns the DirectoryAudits field if non-nil, zero value otherwise.

### GetDirectoryAuditsOk

`func (o *MicrosoftGraphAuditLogRoot) GetDirectoryAuditsOk() (*[]MicrosoftGraphDirectoryAudit, bool)`

GetDirectoryAuditsOk returns a tuple with the DirectoryAudits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryAudits

`func (o *MicrosoftGraphAuditLogRoot) SetDirectoryAudits(v []MicrosoftGraphDirectoryAudit)`

SetDirectoryAudits sets DirectoryAudits field to given value.

### HasDirectoryAudits

`func (o *MicrosoftGraphAuditLogRoot) HasDirectoryAudits() bool`

HasDirectoryAudits returns a boolean if a field has been set.

### GetProvisioning

`func (o *MicrosoftGraphAuditLogRoot) GetProvisioning() []MicrosoftGraphProvisioningObjectSummary`

GetProvisioning returns the Provisioning field if non-nil, zero value otherwise.

### GetProvisioningOk

`func (o *MicrosoftGraphAuditLogRoot) GetProvisioningOk() (*[]MicrosoftGraphProvisioningObjectSummary, bool)`

GetProvisioningOk returns a tuple with the Provisioning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioning

`func (o *MicrosoftGraphAuditLogRoot) SetProvisioning(v []MicrosoftGraphProvisioningObjectSummary)`

SetProvisioning sets Provisioning field to given value.

### HasProvisioning

`func (o *MicrosoftGraphAuditLogRoot) HasProvisioning() bool`

HasProvisioning returns a boolean if a field has been set.

### GetRestrictedSignIns

`func (o *MicrosoftGraphAuditLogRoot) GetRestrictedSignIns() []MicrosoftGraphRestrictedSignIn`

GetRestrictedSignIns returns the RestrictedSignIns field if non-nil, zero value otherwise.

### GetRestrictedSignInsOk

`func (o *MicrosoftGraphAuditLogRoot) GetRestrictedSignInsOk() (*[]MicrosoftGraphRestrictedSignIn, bool)`

GetRestrictedSignInsOk returns a tuple with the RestrictedSignIns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictedSignIns

`func (o *MicrosoftGraphAuditLogRoot) SetRestrictedSignIns(v []MicrosoftGraphRestrictedSignIn)`

SetRestrictedSignIns sets RestrictedSignIns field to given value.

### HasRestrictedSignIns

`func (o *MicrosoftGraphAuditLogRoot) HasRestrictedSignIns() bool`

HasRestrictedSignIns returns a boolean if a field has been set.

### GetSignIns

`func (o *MicrosoftGraphAuditLogRoot) GetSignIns() []MicrosoftGraphSignIn`

GetSignIns returns the SignIns field if non-nil, zero value otherwise.

### GetSignInsOk

`func (o *MicrosoftGraphAuditLogRoot) GetSignInsOk() (*[]MicrosoftGraphSignIn, bool)`

GetSignInsOk returns a tuple with the SignIns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignIns

`func (o *MicrosoftGraphAuditLogRoot) SetSignIns(v []MicrosoftGraphSignIn)`

SetSignIns sets SignIns field to given value.

### HasSignIns

`func (o *MicrosoftGraphAuditLogRoot) HasSignIns() bool`

HasSignIns returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


