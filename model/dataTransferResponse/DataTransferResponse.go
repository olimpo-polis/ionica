// Code generated by github.com/atombender/go-jsonschema, DO NOT EDIT.

package dataTransferResponse

import "encoding/json"
import "fmt"
import "reflect"

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId" yaml:"vendorId" mapstructure:"vendorId"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType) UnmarshalJSON(value []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(value, &raw); err != nil {
		return err
	}
	if _, ok := raw["vendorId"]; raw != nil && !ok {
		return fmt.Errorf("field vendorId in CustomDataType: required")
	}
	type Plain CustomDataType
	var plain Plain
	if err := json.Unmarshal(value, &plain); err != nil {
		return err
	}
	if len(plain.VendorId) > 255 {
		return fmt.Errorf("field %s length: must be <= %d", "vendorId", 255)
	}
	*j = CustomDataType(plain)
	return nil
}

type DataTransferResponseJson struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType `json:"customData,omitempty" yaml:"customData,omitempty" mapstructure:"customData,omitempty"`

	// Data without specified length or format, in response to request.
	//
	Data interface{} `json:"data,omitempty" yaml:"data,omitempty" mapstructure:"data,omitempty"`

	// Status corresponds to the JSON schema field "status".
	Status DataTransferStatusEnumType `json:"status" yaml:"status" mapstructure:"status"`

	// StatusInfo corresponds to the JSON schema field "statusInfo".
	StatusInfo *StatusInfoType `json:"statusInfo,omitempty" yaml:"statusInfo,omitempty" mapstructure:"statusInfo,omitempty"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *DataTransferResponseJson) UnmarshalJSON(value []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(value, &raw); err != nil {
		return err
	}
	if _, ok := raw["status"]; raw != nil && !ok {
		return fmt.Errorf("field status in DataTransferResponseJson: required")
	}
	type Plain DataTransferResponseJson
	var plain Plain
	if err := json.Unmarshal(value, &plain); err != nil {
		return err
	}
	*j = DataTransferResponseJson(plain)
	return nil
}

type DataTransferStatusEnumType string

const DataTransferStatusEnumTypeAccepted DataTransferStatusEnumType = "Accepted"
const DataTransferStatusEnumTypeRejected DataTransferStatusEnumType = "Rejected"
const DataTransferStatusEnumTypeUnknownMessageId DataTransferStatusEnumType = "UnknownMessageId"
const DataTransferStatusEnumTypeUnknownVendorId DataTransferStatusEnumType = "UnknownVendorId"

var enumValues_DataTransferStatusEnumType = []interface{}{
	"Accepted",
	"Rejected",
	"UnknownMessageId",
	"UnknownVendorId",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *DataTransferStatusEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_DataTransferStatusEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_DataTransferStatusEnumType, v)
	}
	*j = DataTransferStatusEnumType(v)
	return nil
}

// Element providing more information about the status.
type StatusInfoType struct {
	// Additional text to provide detailed information.
	//
	AdditionalInfo *string `json:"additionalInfo,omitempty" yaml:"additionalInfo,omitempty" mapstructure:"additionalInfo,omitempty"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType `json:"customData,omitempty" yaml:"customData,omitempty" mapstructure:"customData,omitempty"`

	// A predefined code for the reason why the status is returned in this response.
	// The string is case-insensitive.
	//
	ReasonCode string `json:"reasonCode" yaml:"reasonCode" mapstructure:"reasonCode"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *StatusInfoType) UnmarshalJSON(value []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(value, &raw); err != nil {
		return err
	}
	if _, ok := raw["reasonCode"]; raw != nil && !ok {
		return fmt.Errorf("field reasonCode in StatusInfoType: required")
	}
	type Plain StatusInfoType
	var plain Plain
	if err := json.Unmarshal(value, &plain); err != nil {
		return err
	}
	if plain.AdditionalInfo != nil && len(*plain.AdditionalInfo) > 512 {
		return fmt.Errorf("field %s length: must be <= %d", "additionalInfo", 512)
	}
	if len(plain.ReasonCode) > 20 {
		return fmt.Errorf("field %s length: must be <= %d", "reasonCode", 20)
	}
	*j = StatusInfoType(plain)
	return nil
}
