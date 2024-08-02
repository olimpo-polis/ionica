// Code generated by github.com/atombender/go-jsonschema, DO NOT EDIT.

package notifyCustomerInformationRequest

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

type NotifyCustomerInformationRequestJson struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType `json:"customData,omitempty" yaml:"customData,omitempty" mapstructure:"customData,omitempty"`

	// (Part of) the requested data. No format specified in which the data is
	// returned. Should be human readable.
	//
	Data string `json:"data" yaml:"data" mapstructure:"data"`

	//  Timestamp of the moment this message was generated at the Charging Station.
	//
	GeneratedAt time.Time `json:"generatedAt" yaml:"generatedAt" mapstructure:"generatedAt"`

	// The Id of the request.
	//
	//
	RequestId int `json:"requestId" yaml:"requestId" mapstructure:"requestId"`

	// Sequence number of this message. First message starts at 0.
	//
	SeqNo int `json:"seqNo" yaml:"seqNo" mapstructure:"seqNo"`

	// “to be continued” indicator. Indicates whether another part of the
	// monitoringData follows in an upcoming notifyMonitoringReportRequest message.
	// Default value when omitted is false.
	//
	Tbc bool `json:"tbc,omitempty" yaml:"tbc,omitempty" mapstructure:"tbc,omitempty"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *NotifyCustomerInformationRequestJson) UnmarshalJSON(value []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(value, &raw); err != nil {
		return err
	}
	if _, ok := raw["data"]; raw != nil && !ok {
		return fmt.Errorf("field data in NotifyCustomerInformationRequestJson: required")
	}
	if _, ok := raw["generatedAt"]; raw != nil && !ok {
		return fmt.Errorf("field generatedAt in NotifyCustomerInformationRequestJson: required")
	}
	if _, ok := raw["requestId"]; raw != nil && !ok {
		return fmt.Errorf("field requestId in NotifyCustomerInformationRequestJson: required")
	}
	if _, ok := raw["seqNo"]; raw != nil && !ok {
		return fmt.Errorf("field seqNo in NotifyCustomerInformationRequestJson: required")
	}
	type Plain NotifyCustomerInformationRequestJson
	var plain Plain
	if err := json.Unmarshal(value, &plain); err != nil {
		return err
	}
	if len(plain.Data) > 512 {
		return fmt.Errorf("field %s length: must be <= %d", "data", 512)
	}
	if v, ok := raw["tbc"]; !ok || v == nil {
		plain.Tbc = false
	}
	*j = NotifyCustomerInformationRequestJson(plain)
	return nil
}
