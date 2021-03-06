/*
Partial Graph API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package msgraph

import (
	"encoding/json"
)

// MicrosoftGraphSubjectRightsRequestStageDetail struct for MicrosoftGraphSubjectRightsRequestStageDetail
type MicrosoftGraphSubjectRightsRequestStageDetail struct {
	// Describes the error, if any, for the current stage.
	Error AnyOfmicrosoftGraphPublicError `json:"error,omitempty"`
	// The stage of the subject rights request. Possible values are: contentRetrieval, contentReview, generateReport, contentDeletion, caseResolved, unknownFutureValue.
	Stage AnyOfmicrosoftGraphSubjectRightsRequestStage `json:"stage,omitempty"`
	// Status of the current stage. Possible values are: notStarted, current, completed, failed, unknownFutureValue.
	Status AnyOfmicrosoftGraphSubjectRightsRequestStageStatus `json:"status,omitempty"`
}

// NewMicrosoftGraphSubjectRightsRequestStageDetail instantiates a new MicrosoftGraphSubjectRightsRequestStageDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphSubjectRightsRequestStageDetail() *MicrosoftGraphSubjectRightsRequestStageDetail {
	this := MicrosoftGraphSubjectRightsRequestStageDetail{}
	return &this
}

// NewMicrosoftGraphSubjectRightsRequestStageDetailWithDefaults instantiates a new MicrosoftGraphSubjectRightsRequestStageDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphSubjectRightsRequestStageDetailWithDefaults() *MicrosoftGraphSubjectRightsRequestStageDetail {
	this := MicrosoftGraphSubjectRightsRequestStageDetail{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphSubjectRightsRequestStageDetail) GetError() AnyOfmicrosoftGraphPublicError {
	if o == nil  {
		var ret AnyOfmicrosoftGraphPublicError
		return ret
	}
	return o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphSubjectRightsRequestStageDetail) GetErrorOk() (*AnyOfmicrosoftGraphPublicError, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return &o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *MicrosoftGraphSubjectRightsRequestStageDetail) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given AnyOfmicrosoftGraphPublicError and assigns it to the Error field.
func (o *MicrosoftGraphSubjectRightsRequestStageDetail) SetError(v AnyOfmicrosoftGraphPublicError) {
	o.Error = v
}

// GetStage returns the Stage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphSubjectRightsRequestStageDetail) GetStage() AnyOfmicrosoftGraphSubjectRightsRequestStage {
	if o == nil  {
		var ret AnyOfmicrosoftGraphSubjectRightsRequestStage
		return ret
	}
	return o.Stage
}

// GetStageOk returns a tuple with the Stage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphSubjectRightsRequestStageDetail) GetStageOk() (*AnyOfmicrosoftGraphSubjectRightsRequestStage, bool) {
	if o == nil || o.Stage == nil {
		return nil, false
	}
	return &o.Stage, true
}

// HasStage returns a boolean if a field has been set.
func (o *MicrosoftGraphSubjectRightsRequestStageDetail) HasStage() bool {
	if o != nil && o.Stage != nil {
		return true
	}

	return false
}

// SetStage gets a reference to the given AnyOfmicrosoftGraphSubjectRightsRequestStage and assigns it to the Stage field.
func (o *MicrosoftGraphSubjectRightsRequestStageDetail) SetStage(v AnyOfmicrosoftGraphSubjectRightsRequestStage) {
	o.Stage = v
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphSubjectRightsRequestStageDetail) GetStatus() AnyOfmicrosoftGraphSubjectRightsRequestStageStatus {
	if o == nil  {
		var ret AnyOfmicrosoftGraphSubjectRightsRequestStageStatus
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphSubjectRightsRequestStageDetail) GetStatusOk() (*AnyOfmicrosoftGraphSubjectRightsRequestStageStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return &o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *MicrosoftGraphSubjectRightsRequestStageDetail) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given AnyOfmicrosoftGraphSubjectRightsRequestStageStatus and assigns it to the Status field.
func (o *MicrosoftGraphSubjectRightsRequestStageDetail) SetStatus(v AnyOfmicrosoftGraphSubjectRightsRequestStageStatus) {
	o.Status = v
}

func (o MicrosoftGraphSubjectRightsRequestStageDetail) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	if o.Stage != nil {
		toSerialize["stage"] = o.Stage
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphSubjectRightsRequestStageDetail struct {
	value *MicrosoftGraphSubjectRightsRequestStageDetail
	isSet bool
}

func (v NullableMicrosoftGraphSubjectRightsRequestStageDetail) Get() *MicrosoftGraphSubjectRightsRequestStageDetail {
	return v.value
}

func (v *NullableMicrosoftGraphSubjectRightsRequestStageDetail) Set(val *MicrosoftGraphSubjectRightsRequestStageDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphSubjectRightsRequestStageDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphSubjectRightsRequestStageDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphSubjectRightsRequestStageDetail(val *MicrosoftGraphSubjectRightsRequestStageDetail) *NullableMicrosoftGraphSubjectRightsRequestStageDetail {
	return &NullableMicrosoftGraphSubjectRightsRequestStageDetail{value: val, isSet: true}
}

func (v NullableMicrosoftGraphSubjectRightsRequestStageDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphSubjectRightsRequestStageDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


