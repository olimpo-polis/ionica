// Code generated by github.com/atombender/go-jsonschema, DO NOT EDIT.

package clearChargingProfileRequest

import "encoding/json"
import "fmt"
import "reflect"

type ChargingProfilePurposeEnumType string

const ChargingProfilePurposeEnumTypeChargingStationExternalConstraints ChargingProfilePurposeEnumType = "ChargingStationExternalConstraints"
const ChargingProfilePurposeEnumTypeChargingStationMaxProfile ChargingProfilePurposeEnumType = "ChargingStationMaxProfile"
const ChargingProfilePurposeEnumTypeTxDefaultProfile ChargingProfilePurposeEnumType = "TxDefaultProfile"
const ChargingProfilePurposeEnumTypeTxProfile ChargingProfilePurposeEnumType = "TxProfile"

var enumValues_ChargingProfilePurposeEnumType = []interface{}{
	"ChargingStationExternalConstraints",
	"ChargingStationMaxProfile",
	"TxDefaultProfile",
	"TxProfile",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ChargingProfilePurposeEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ChargingProfilePurposeEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ChargingProfilePurposeEnumType, v)
	}
	*j = ChargingProfilePurposeEnumType(v)
	return nil
}

type ClearChargingProfileRequestJson struct {
	// ChargingProfileCriteria corresponds to the JSON schema field
	// "chargingProfileCriteria".
	ChargingProfileCriteria *ClearChargingProfileType `json:"chargingProfileCriteria,omitempty" yaml:"chargingProfileCriteria,omitempty" mapstructure:"chargingProfileCriteria,omitempty"`

	// The Id of the charging profile to clear.
	//
	ChargingProfileId *int `json:"chargingProfileId,omitempty" yaml:"chargingProfileId,omitempty" mapstructure:"chargingProfileId,omitempty"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType `json:"customData,omitempty" yaml:"customData,omitempty" mapstructure:"customData,omitempty"`
}

// Charging_ Profile
// urn:x-oca:ocpp:uid:2:233255
// A ChargingProfile consists of a ChargingSchedule, describing the amount of power
// or current that can be delivered per time interval.
type ClearChargingProfileType struct {
	// ChargingProfilePurpose corresponds to the JSON schema field
	// "chargingProfilePurpose".
	ChargingProfilePurpose *ChargingProfilePurposeEnumType `json:"chargingProfilePurpose,omitempty" yaml:"chargingProfilePurpose,omitempty" mapstructure:"chargingProfilePurpose,omitempty"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType `json:"customData,omitempty" yaml:"customData,omitempty" mapstructure:"customData,omitempty"`

	// Identified_ Object. MRID. Numeric_ Identifier
	// urn:x-enexis:ecdm:uid:1:569198
	// Specifies the id of the EVSE for which to clear charging profiles. An evseId of
	// zero (0) specifies the charging profile for the overall Charging Station.
	// Absence of this parameter means the clearing applies to all charging profiles
	// that match the other criteria in the request.
	//
	//
	EvseId *int `json:"evseId,omitempty" yaml:"evseId,omitempty" mapstructure:"evseId,omitempty"`

	// Charging_ Profile. Stack_ Level. Counter
	// urn:x-oca:ocpp:uid:1:569230
	// Specifies the stackLevel for which charging profiles will be cleared, if they
	// meet the other criteria in the request.
	//
	StackLevel *int `json:"stackLevel,omitempty" yaml:"stackLevel,omitempty" mapstructure:"stackLevel,omitempty"`
}

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