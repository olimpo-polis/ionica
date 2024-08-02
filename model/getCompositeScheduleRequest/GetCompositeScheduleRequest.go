// Code generated by github.com/atombender/go-jsonschema, DO NOT EDIT.

package getCompositeScheduleRequest

import "encoding/json"
import "fmt"
import "reflect"

type ChargingRateUnitEnumType string

const ChargingRateUnitEnumTypeA ChargingRateUnitEnumType = "A"
const ChargingRateUnitEnumTypeW ChargingRateUnitEnumType = "W"

var enumValues_ChargingRateUnitEnumType = []interface{}{
	"W",
	"A",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ChargingRateUnitEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ChargingRateUnitEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ChargingRateUnitEnumType, v)
	}
	*j = ChargingRateUnitEnumType(v)
	return nil
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

type GetCompositeScheduleRequestJson struct {
	// ChargingRateUnit corresponds to the JSON schema field "chargingRateUnit".
	ChargingRateUnit *ChargingRateUnitEnumType `json:"chargingRateUnit,omitempty" yaml:"chargingRateUnit,omitempty" mapstructure:"chargingRateUnit,omitempty"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType `json:"customData,omitempty" yaml:"customData,omitempty" mapstructure:"customData,omitempty"`

	// Length of the requested schedule in seconds.
	//
	//
	Duration int `json:"duration" yaml:"duration" mapstructure:"duration"`

	// The ID of the EVSE for which the schedule is requested. When evseid=0, the
	// Charging Station will calculate the expected consumption for the grid
	// connection.
	//
	EvseId int `json:"evseId" yaml:"evseId" mapstructure:"evseId"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *GetCompositeScheduleRequestJson) UnmarshalJSON(value []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(value, &raw); err != nil {
		return err
	}
	if _, ok := raw["duration"]; raw != nil && !ok {
		return fmt.Errorf("field duration in GetCompositeScheduleRequestJson: required")
	}
	if _, ok := raw["evseId"]; raw != nil && !ok {
		return fmt.Errorf("field evseId in GetCompositeScheduleRequestJson: required")
	}
	type Plain GetCompositeScheduleRequestJson
	var plain Plain
	if err := json.Unmarshal(value, &plain); err != nil {
		return err
	}
	*j = GetCompositeScheduleRequestJson(plain)
	return nil
}
