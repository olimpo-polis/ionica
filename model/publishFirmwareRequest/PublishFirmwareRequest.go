// Code generated by github.com/atombender/go-jsonschema, DO NOT EDIT.

package publishFirmwareRequest

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

type PublishFirmwareRequestJson struct {
	// The MD5 checksum over the entire firmware file as a hexadecimal string of
	// length 32.
	//
	Checksum string `json:"checksum" yaml:"checksum" mapstructure:"checksum"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType `json:"customData,omitempty" yaml:"customData,omitempty" mapstructure:"customData,omitempty"`

	// This contains a string containing a URI pointing to a
	// location from which to retrieve the firmware.
	//
	Location string `json:"location" yaml:"location" mapstructure:"location"`

	// The Id of the request.
	//
	RequestId int `json:"requestId" yaml:"requestId" mapstructure:"requestId"`

	// This specifies how many times Charging Station must try
	// to download the firmware before giving up. If this field is not
	// present, it is left to Charging Station to decide how many times it wants to
	// retry.
	//
	Retries *int `json:"retries,omitempty" yaml:"retries,omitempty" mapstructure:"retries,omitempty"`

	// The interval in seconds
	// after which a retry may be
	// attempted. If this field is not
	// present, it is left to Charging
	// Station to decide how long to wait
	// between attempts.
	//
	RetryInterval *int `json:"retryInterval,omitempty" yaml:"retryInterval,omitempty" mapstructure:"retryInterval,omitempty"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *PublishFirmwareRequestJson) UnmarshalJSON(value []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(value, &raw); err != nil {
		return err
	}
	if _, ok := raw["checksum"]; raw != nil && !ok {
		return fmt.Errorf("field checksum in PublishFirmwareRequestJson: required")
	}
	if _, ok := raw["location"]; raw != nil && !ok {
		return fmt.Errorf("field location in PublishFirmwareRequestJson: required")
	}
	if _, ok := raw["requestId"]; raw != nil && !ok {
		return fmt.Errorf("field requestId in PublishFirmwareRequestJson: required")
	}
	type Plain PublishFirmwareRequestJson
	var plain Plain
	if err := json.Unmarshal(value, &plain); err != nil {
		return err
	}
	if len(plain.Checksum) > 32 {
		return fmt.Errorf("field %s length: must be <= %d", "checksum", 32)
	}
	if len(plain.Location) > 512 {
		return fmt.Errorf("field %s length: must be <= %d", "location", 512)
	}
	*j = PublishFirmwareRequestJson(plain)
	return nil
}
