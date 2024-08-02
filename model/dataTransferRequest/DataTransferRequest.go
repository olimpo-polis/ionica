// Code generated by github.com/atombender/go-jsonschema, DO NOT EDIT.

package dataTransferRequest

import "encoding/json"
import "fmt"

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

type DataTransferRequestJson struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType `json:"customData,omitempty" yaml:"customData,omitempty" mapstructure:"customData,omitempty"`

	// Data without specified length or format. This needs to be decided by both
	// parties (Open to implementation).
	//
	Data interface{} `json:"data,omitempty" yaml:"data,omitempty" mapstructure:"data,omitempty"`

	// May be used to indicate a specific message or implementation.
	//
	MessageId *string `json:"messageId,omitempty" yaml:"messageId,omitempty" mapstructure:"messageId,omitempty"`

	// This identifies the Vendor specific implementation
	//
	//
	VendorId string `json:"vendorId" yaml:"vendorId" mapstructure:"vendorId"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *DataTransferRequestJson) UnmarshalJSON(value []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(value, &raw); err != nil {
		return err
	}
	if _, ok := raw["vendorId"]; raw != nil && !ok {
		return fmt.Errorf("field vendorId in DataTransferRequestJson: required")
	}
	type Plain DataTransferRequestJson
	var plain Plain
	if err := json.Unmarshal(value, &plain); err != nil {
		return err
	}
	if plain.MessageId != nil && len(*plain.MessageId) > 50 {
		return fmt.Errorf("field %s length: must be <= %d", "messageId", 50)
	}
	if len(plain.VendorId) > 255 {
		return fmt.Errorf("field %s length: must be <= %d", "vendorId", 255)
	}
	*j = DataTransferRequestJson(plain)
	return nil
}