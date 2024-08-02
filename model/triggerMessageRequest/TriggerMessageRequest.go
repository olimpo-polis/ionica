// Code generated by github.com/atombender/go-jsonschema, DO NOT EDIT.

package triggerMessageRequest

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

// EVSE
// urn:x-oca:ocpp:uid:2:233123
// Electric Vehicle Supply Equipment
type EVSEType struct {
	// An id to designate a specific connector (on an EVSE) by connector index number.
	//
	ConnectorId *int `json:"connectorId,omitempty" yaml:"connectorId,omitempty" mapstructure:"connectorId,omitempty"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType `json:"customData,omitempty" yaml:"customData,omitempty" mapstructure:"customData,omitempty"`

	// Identified_ Object. MRID. Numeric_ Identifier
	// urn:x-enexis:ecdm:uid:1:569198
	// EVSE Identifier. This contains a number (&gt; 0) designating an EVSE of the
	// Charging Station.
	//
	Id int `json:"id" yaml:"id" mapstructure:"id"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *EVSEType) UnmarshalJSON(value []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(value, &raw); err != nil {
		return err
	}
	if _, ok := raw["id"]; raw != nil && !ok {
		return fmt.Errorf("field id in EVSEType: required")
	}
	type Plain EVSEType
	var plain Plain
	if err := json.Unmarshal(value, &plain); err != nil {
		return err
	}
	*j = EVSEType(plain)
	return nil
}

type MessageTriggerEnumType string

const MessageTriggerEnumTypeBootNotification MessageTriggerEnumType = "BootNotification"
const MessageTriggerEnumTypeFirmwareStatusNotification MessageTriggerEnumType = "FirmwareStatusNotification"
const MessageTriggerEnumTypeHeartbeat MessageTriggerEnumType = "Heartbeat"
const MessageTriggerEnumTypeLogStatusNotification MessageTriggerEnumType = "LogStatusNotification"
const MessageTriggerEnumTypeMeterValues MessageTriggerEnumType = "MeterValues"
const MessageTriggerEnumTypePublishFirmwareStatusNotification MessageTriggerEnumType = "PublishFirmwareStatusNotification"
const MessageTriggerEnumTypeSignChargingStationCertificate MessageTriggerEnumType = "SignChargingStationCertificate"
const MessageTriggerEnumTypeSignCombinedCertificate MessageTriggerEnumType = "SignCombinedCertificate"
const MessageTriggerEnumTypeSignV2GCertificate MessageTriggerEnumType = "SignV2GCertificate"
const MessageTriggerEnumTypeStatusNotification MessageTriggerEnumType = "StatusNotification"
const MessageTriggerEnumTypeTransactionEvent MessageTriggerEnumType = "TransactionEvent"

var enumValues_MessageTriggerEnumType = []interface{}{
	"BootNotification",
	"LogStatusNotification",
	"FirmwareStatusNotification",
	"Heartbeat",
	"MeterValues",
	"SignChargingStationCertificate",
	"SignV2GCertificate",
	"StatusNotification",
	"TransactionEvent",
	"SignCombinedCertificate",
	"PublishFirmwareStatusNotification",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *MessageTriggerEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_MessageTriggerEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_MessageTriggerEnumType, v)
	}
	*j = MessageTriggerEnumType(v)
	return nil
}

type TriggerMessageRequestJson struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType `json:"customData,omitempty" yaml:"customData,omitempty" mapstructure:"customData,omitempty"`

	// Evse corresponds to the JSON schema field "evse".
	Evse *EVSEType `json:"evse,omitempty" yaml:"evse,omitempty" mapstructure:"evse,omitempty"`

	// RequestedMessage corresponds to the JSON schema field "requestedMessage".
	RequestedMessage MessageTriggerEnumType `json:"requestedMessage" yaml:"requestedMessage" mapstructure:"requestedMessage"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *TriggerMessageRequestJson) UnmarshalJSON(value []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(value, &raw); err != nil {
		return err
	}
	if _, ok := raw["requestedMessage"]; raw != nil && !ok {
		return fmt.Errorf("field requestedMessage in TriggerMessageRequestJson: required")
	}
	type Plain TriggerMessageRequestJson
	var plain Plain
	if err := json.Unmarshal(value, &plain); err != nil {
		return err
	}
	*j = TriggerMessageRequestJson(plain)
	return nil
}