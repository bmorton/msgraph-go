# MicrosoftGraphProcess

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountName** | Pointer to **NullableString** | User account identifier (user account context the process ran under) for example, AccountName, SID, and so on. | [optional] 
**CommandLine** | Pointer to **NullableString** | The full process invocation commandline including all parameters. | [optional] 
**CreatedDateTime** | Pointer to **NullableTime** | Time at which the process was started. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. | [optional] 
**FileHash** | Pointer to [**AnyOfmicrosoftGraphFileHash**](anyOf&lt;microsoft.graph.fileHash&gt;.md) | Complex type containing file hashes (cryptographic and location-sensitive). | [optional] 
**IntegrityLevel** | Pointer to [**AnyOfmicrosoftGraphProcessIntegrityLevel**](anyOf&lt;microsoft.graph.processIntegrityLevel&gt;.md) | The integrity level of the process. Possible values are: unknown, untrusted, low, medium, high, system. | [optional] 
**IsElevated** | Pointer to **NullableBool** | True if the process is elevated. | [optional] 
**Name** | Pointer to **NullableString** | The name of the process&#39; Image file. | [optional] 
**ParentProcessCreatedDateTime** | Pointer to **NullableTime** | DateTime at which the parent process was started. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. | [optional] 
**ParentProcessId** | Pointer to **NullableInt32** | The Process ID (PID) of the parent process. | [optional] 
**ParentProcessName** | Pointer to **NullableString** | The name of the image file of the parent process. | [optional] 
**Path** | Pointer to **NullableString** | Full path, including filename. | [optional] 
**ProcessId** | Pointer to **NullableInt32** | The Process ID (PID) of the process. | [optional] 

## Methods

### NewMicrosoftGraphProcess

`func NewMicrosoftGraphProcess() *MicrosoftGraphProcess`

NewMicrosoftGraphProcess instantiates a new MicrosoftGraphProcess object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphProcessWithDefaults

`func NewMicrosoftGraphProcessWithDefaults() *MicrosoftGraphProcess`

NewMicrosoftGraphProcessWithDefaults instantiates a new MicrosoftGraphProcess object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountName

`func (o *MicrosoftGraphProcess) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *MicrosoftGraphProcess) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *MicrosoftGraphProcess) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.

### HasAccountName

`func (o *MicrosoftGraphProcess) HasAccountName() bool`

HasAccountName returns a boolean if a field has been set.

### SetAccountNameNil

`func (o *MicrosoftGraphProcess) SetAccountNameNil(b bool)`

 SetAccountNameNil sets the value for AccountName to be an explicit nil

### UnsetAccountName
`func (o *MicrosoftGraphProcess) UnsetAccountName()`

UnsetAccountName ensures that no value is present for AccountName, not even an explicit nil
### GetCommandLine

`func (o *MicrosoftGraphProcess) GetCommandLine() string`

GetCommandLine returns the CommandLine field if non-nil, zero value otherwise.

### GetCommandLineOk

`func (o *MicrosoftGraphProcess) GetCommandLineOk() (*string, bool)`

GetCommandLineOk returns a tuple with the CommandLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandLine

`func (o *MicrosoftGraphProcess) SetCommandLine(v string)`

SetCommandLine sets CommandLine field to given value.

### HasCommandLine

`func (o *MicrosoftGraphProcess) HasCommandLine() bool`

HasCommandLine returns a boolean if a field has been set.

### SetCommandLineNil

`func (o *MicrosoftGraphProcess) SetCommandLineNil(b bool)`

 SetCommandLineNil sets the value for CommandLine to be an explicit nil

### UnsetCommandLine
`func (o *MicrosoftGraphProcess) UnsetCommandLine()`

UnsetCommandLine ensures that no value is present for CommandLine, not even an explicit nil
### GetCreatedDateTime

`func (o *MicrosoftGraphProcess) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphProcess) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphProcess) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *MicrosoftGraphProcess) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTimeNil

`func (o *MicrosoftGraphProcess) SetCreatedDateTimeNil(b bool)`

 SetCreatedDateTimeNil sets the value for CreatedDateTime to be an explicit nil

### UnsetCreatedDateTime
`func (o *MicrosoftGraphProcess) UnsetCreatedDateTime()`

UnsetCreatedDateTime ensures that no value is present for CreatedDateTime, not even an explicit nil
### GetFileHash

`func (o *MicrosoftGraphProcess) GetFileHash() AnyOfmicrosoftGraphFileHash`

GetFileHash returns the FileHash field if non-nil, zero value otherwise.

### GetFileHashOk

`func (o *MicrosoftGraphProcess) GetFileHashOk() (*AnyOfmicrosoftGraphFileHash, bool)`

GetFileHashOk returns a tuple with the FileHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileHash

`func (o *MicrosoftGraphProcess) SetFileHash(v AnyOfmicrosoftGraphFileHash)`

SetFileHash sets FileHash field to given value.

### HasFileHash

`func (o *MicrosoftGraphProcess) HasFileHash() bool`

HasFileHash returns a boolean if a field has been set.

### SetFileHashNil

`func (o *MicrosoftGraphProcess) SetFileHashNil(b bool)`

 SetFileHashNil sets the value for FileHash to be an explicit nil

### UnsetFileHash
`func (o *MicrosoftGraphProcess) UnsetFileHash()`

UnsetFileHash ensures that no value is present for FileHash, not even an explicit nil
### GetIntegrityLevel

`func (o *MicrosoftGraphProcess) GetIntegrityLevel() AnyOfmicrosoftGraphProcessIntegrityLevel`

GetIntegrityLevel returns the IntegrityLevel field if non-nil, zero value otherwise.

### GetIntegrityLevelOk

`func (o *MicrosoftGraphProcess) GetIntegrityLevelOk() (*AnyOfmicrosoftGraphProcessIntegrityLevel, bool)`

GetIntegrityLevelOk returns a tuple with the IntegrityLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrityLevel

`func (o *MicrosoftGraphProcess) SetIntegrityLevel(v AnyOfmicrosoftGraphProcessIntegrityLevel)`

SetIntegrityLevel sets IntegrityLevel field to given value.

### HasIntegrityLevel

`func (o *MicrosoftGraphProcess) HasIntegrityLevel() bool`

HasIntegrityLevel returns a boolean if a field has been set.

### SetIntegrityLevelNil

`func (o *MicrosoftGraphProcess) SetIntegrityLevelNil(b bool)`

 SetIntegrityLevelNil sets the value for IntegrityLevel to be an explicit nil

### UnsetIntegrityLevel
`func (o *MicrosoftGraphProcess) UnsetIntegrityLevel()`

UnsetIntegrityLevel ensures that no value is present for IntegrityLevel, not even an explicit nil
### GetIsElevated

`func (o *MicrosoftGraphProcess) GetIsElevated() bool`

GetIsElevated returns the IsElevated field if non-nil, zero value otherwise.

### GetIsElevatedOk

`func (o *MicrosoftGraphProcess) GetIsElevatedOk() (*bool, bool)`

GetIsElevatedOk returns a tuple with the IsElevated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsElevated

`func (o *MicrosoftGraphProcess) SetIsElevated(v bool)`

SetIsElevated sets IsElevated field to given value.

### HasIsElevated

`func (o *MicrosoftGraphProcess) HasIsElevated() bool`

HasIsElevated returns a boolean if a field has been set.

### SetIsElevatedNil

`func (o *MicrosoftGraphProcess) SetIsElevatedNil(b bool)`

 SetIsElevatedNil sets the value for IsElevated to be an explicit nil

### UnsetIsElevated
`func (o *MicrosoftGraphProcess) UnsetIsElevated()`

UnsetIsElevated ensures that no value is present for IsElevated, not even an explicit nil
### GetName

`func (o *MicrosoftGraphProcess) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MicrosoftGraphProcess) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MicrosoftGraphProcess) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MicrosoftGraphProcess) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *MicrosoftGraphProcess) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *MicrosoftGraphProcess) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetParentProcessCreatedDateTime

`func (o *MicrosoftGraphProcess) GetParentProcessCreatedDateTime() time.Time`

GetParentProcessCreatedDateTime returns the ParentProcessCreatedDateTime field if non-nil, zero value otherwise.

### GetParentProcessCreatedDateTimeOk

`func (o *MicrosoftGraphProcess) GetParentProcessCreatedDateTimeOk() (*time.Time, bool)`

GetParentProcessCreatedDateTimeOk returns a tuple with the ParentProcessCreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentProcessCreatedDateTime

`func (o *MicrosoftGraphProcess) SetParentProcessCreatedDateTime(v time.Time)`

SetParentProcessCreatedDateTime sets ParentProcessCreatedDateTime field to given value.

### HasParentProcessCreatedDateTime

`func (o *MicrosoftGraphProcess) HasParentProcessCreatedDateTime() bool`

HasParentProcessCreatedDateTime returns a boolean if a field has been set.

### SetParentProcessCreatedDateTimeNil

`func (o *MicrosoftGraphProcess) SetParentProcessCreatedDateTimeNil(b bool)`

 SetParentProcessCreatedDateTimeNil sets the value for ParentProcessCreatedDateTime to be an explicit nil

### UnsetParentProcessCreatedDateTime
`func (o *MicrosoftGraphProcess) UnsetParentProcessCreatedDateTime()`

UnsetParentProcessCreatedDateTime ensures that no value is present for ParentProcessCreatedDateTime, not even an explicit nil
### GetParentProcessId

`func (o *MicrosoftGraphProcess) GetParentProcessId() int32`

GetParentProcessId returns the ParentProcessId field if non-nil, zero value otherwise.

### GetParentProcessIdOk

`func (o *MicrosoftGraphProcess) GetParentProcessIdOk() (*int32, bool)`

GetParentProcessIdOk returns a tuple with the ParentProcessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentProcessId

`func (o *MicrosoftGraphProcess) SetParentProcessId(v int32)`

SetParentProcessId sets ParentProcessId field to given value.

### HasParentProcessId

`func (o *MicrosoftGraphProcess) HasParentProcessId() bool`

HasParentProcessId returns a boolean if a field has been set.

### SetParentProcessIdNil

`func (o *MicrosoftGraphProcess) SetParentProcessIdNil(b bool)`

 SetParentProcessIdNil sets the value for ParentProcessId to be an explicit nil

### UnsetParentProcessId
`func (o *MicrosoftGraphProcess) UnsetParentProcessId()`

UnsetParentProcessId ensures that no value is present for ParentProcessId, not even an explicit nil
### GetParentProcessName

`func (o *MicrosoftGraphProcess) GetParentProcessName() string`

GetParentProcessName returns the ParentProcessName field if non-nil, zero value otherwise.

### GetParentProcessNameOk

`func (o *MicrosoftGraphProcess) GetParentProcessNameOk() (*string, bool)`

GetParentProcessNameOk returns a tuple with the ParentProcessName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentProcessName

`func (o *MicrosoftGraphProcess) SetParentProcessName(v string)`

SetParentProcessName sets ParentProcessName field to given value.

### HasParentProcessName

`func (o *MicrosoftGraphProcess) HasParentProcessName() bool`

HasParentProcessName returns a boolean if a field has been set.

### SetParentProcessNameNil

`func (o *MicrosoftGraphProcess) SetParentProcessNameNil(b bool)`

 SetParentProcessNameNil sets the value for ParentProcessName to be an explicit nil

### UnsetParentProcessName
`func (o *MicrosoftGraphProcess) UnsetParentProcessName()`

UnsetParentProcessName ensures that no value is present for ParentProcessName, not even an explicit nil
### GetPath

`func (o *MicrosoftGraphProcess) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *MicrosoftGraphProcess) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *MicrosoftGraphProcess) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *MicrosoftGraphProcess) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *MicrosoftGraphProcess) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *MicrosoftGraphProcess) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetProcessId

`func (o *MicrosoftGraphProcess) GetProcessId() int32`

GetProcessId returns the ProcessId field if non-nil, zero value otherwise.

### GetProcessIdOk

`func (o *MicrosoftGraphProcess) GetProcessIdOk() (*int32, bool)`

GetProcessIdOk returns a tuple with the ProcessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessId

`func (o *MicrosoftGraphProcess) SetProcessId(v int32)`

SetProcessId sets ProcessId field to given value.

### HasProcessId

`func (o *MicrosoftGraphProcess) HasProcessId() bool`

HasProcessId returns a boolean if a field has been set.

### SetProcessIdNil

`func (o *MicrosoftGraphProcess) SetProcessIdNil(b bool)`

 SetProcessIdNil sets the value for ProcessId to be an explicit nil

### UnsetProcessId
`func (o *MicrosoftGraphProcess) UnsetProcessId()`

UnsetProcessId ensures that no value is present for ProcessId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


