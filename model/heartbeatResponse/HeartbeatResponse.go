// Code generated by github.com/atombender/go-jsonschema, DO NOT EDIT.

package heartbeatResponse

import "encoding/json"
import "fmt"
import "time"

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

type HeartbeatResponseJson struct {
	// Contains the current time of the CSMS.
	//
	CurrentTime time.Time `json:"currentTime" yaml:"currentTime" mapstructure:"currentTime"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType `json:"customData,omitempty" yaml:"customData,omitempty" mapstructure:"customData,omitempty"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *HeartbeatResponseJson) UnmarshalJSON(value []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(value, &raw); err != nil {
		return err
	}
	if _, ok := raw["currentTime"]; raw != nil && !ok {
		return fmt.Errorf("field currentTime in HeartbeatResponseJson: required")
	}
	type Plain HeartbeatResponseJson
	var plain Plain
	if err := json.Unmarshal(value, &plain); err != nil {
		return err
	}
	*j = HeartbeatResponseJson(plain)
	return nil
}