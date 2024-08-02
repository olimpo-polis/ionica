// Code generated by github.com/atombender/go-jsonschema, DO NOT EDIT.

package setVariablesRequest

import "encoding/json"
import "fmt"
import "reflect"

type AttributeEnumType string

const AttributeEnumTypeActual AttributeEnumType = "Actual"
const AttributeEnumTypeMaxSet AttributeEnumType = "MaxSet"
const AttributeEnumTypeMinSet AttributeEnumType = "MinSet"
const AttributeEnumTypeTarget AttributeEnumType = "Target"

var enumValues_AttributeEnumType = []interface{}{
	"Actual",
	"Target",
	"MinSet",
	"MaxSet",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *AttributeEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_AttributeEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_AttributeEnumType, v)
	}
	*j = AttributeEnumType(v)
	return nil
}

// A physical or logical component
type ComponentType struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType `json:"customData,omitempty" yaml:"customData,omitempty" mapstructure:"customData,omitempty"`

	// Evse corresponds to the JSON schema field "evse".
	Evse *EVSEType `json:"evse,omitempty" yaml:"evse,omitempty" mapstructure:"evse,omitempty"`

	// Name of instance in case the component exists as multiple instances. Case
	// Insensitive. strongly advised to use Camel Case.
	//
	Instance *string `json:"instance,omitempty" yaml:"instance,omitempty" mapstructure:"instance,omitempty"`

	// Name of the component. Name should be taken from the list of standardized
	// component names whenever possible. Case Insensitive. strongly advised to use
	// Camel Case.
	//
	Name string `json:"name" yaml:"name" mapstructure:"name"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ComponentType) UnmarshalJSON(value []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(value, &raw); err != nil {
		return err
	}
	if _, ok := raw["name"]; raw != nil && !ok {
		return fmt.Errorf("field name in ComponentType: required")
	}
	type Plain ComponentType
	var plain Plain
	if err := json.Unmarshal(value, &plain); err != nil {
		return err
	}
	if plain.Instance != nil && len(*plain.Instance) > 50 {
		return fmt.Errorf("field %s length: must be <= %d", "instance", 50)
	}
	if len(plain.Name) > 50 {
		return fmt.Errorf("field %s length: must be <= %d", "name", 50)
	}
	*j = ComponentType(plain)
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

type SetVariableDataType struct {
	// AttributeType corresponds to the JSON schema field "attributeType".
	AttributeType *AttributeEnumType `json:"attributeType,omitempty" yaml:"attributeType,omitempty" mapstructure:"attributeType,omitempty"`

	// Value to be assigned to attribute of variable.
	//
	// The Configuration Variable
	// &lt;&lt;configkey-configuration-value-size,ConfigurationValueSize&gt;&gt; can
	// be used to limit SetVariableData.attributeValue and
	// VariableCharacteristics.valueList. The max size of these values will always
	// remain equal.
	//
	AttributeValue string `json:"attributeValue" yaml:"attributeValue" mapstructure:"attributeValue"`

	// Component corresponds to the JSON schema field "component".
	Component ComponentType `json:"component" yaml:"component" mapstructure:"component"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType `json:"customData,omitempty" yaml:"customData,omitempty" mapstructure:"customData,omitempty"`

	// Variable corresponds to the JSON schema field "variable".
	Variable VariableType `json:"variable" yaml:"variable" mapstructure:"variable"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *SetVariableDataType) UnmarshalJSON(value []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(value, &raw); err != nil {
		return err
	}
	if _, ok := raw["attributeValue"]; raw != nil && !ok {
		return fmt.Errorf("field attributeValue in SetVariableDataType: required")
	}
	if _, ok := raw["component"]; raw != nil && !ok {
		return fmt.Errorf("field component in SetVariableDataType: required")
	}
	if _, ok := raw["variable"]; raw != nil && !ok {
		return fmt.Errorf("field variable in SetVariableDataType: required")
	}
	type Plain SetVariableDataType
	var plain Plain
	if err := json.Unmarshal(value, &plain); err != nil {
		return err
	}
	if len(plain.AttributeValue) > 1000 {
		return fmt.Errorf("field %s length: must be <= %d", "attributeValue", 1000)
	}
	*j = SetVariableDataType(plain)
	return nil
}

type SetVariablesRequestJson struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType `json:"customData,omitempty" yaml:"customData,omitempty" mapstructure:"customData,omitempty"`

	// SetVariableData corresponds to the JSON schema field "setVariableData".
	SetVariableData []SetVariableDataType `json:"setVariableData" yaml:"setVariableData" mapstructure:"setVariableData"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *SetVariablesRequestJson) UnmarshalJSON(value []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(value, &raw); err != nil {
		return err
	}
	if _, ok := raw["setVariableData"]; raw != nil && !ok {
		return fmt.Errorf("field setVariableData in SetVariablesRequestJson: required")
	}
	type Plain SetVariablesRequestJson
	var plain Plain
	if err := json.Unmarshal(value, &plain); err != nil {
		return err
	}
	if plain.SetVariableData != nil && len(plain.SetVariableData) < 1 {
		return fmt.Errorf("field %s length: must be >= %d", "setVariableData", 1)
	}
	*j = SetVariablesRequestJson(plain)
	return nil
}

// Reference key to a component-variable.
type VariableType struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType `json:"customData,omitempty" yaml:"customData,omitempty" mapstructure:"customData,omitempty"`

	// Name of instance in case the variable exists as multiple instances. Case
	// Insensitive. strongly advised to use Camel Case.
	//
	Instance *string `json:"instance,omitempty" yaml:"instance,omitempty" mapstructure:"instance,omitempty"`

	// Name of the variable. Name should be taken from the list of standardized
	// variable names whenever possible. Case Insensitive. strongly advised to use
	// Camel Case.
	//
	Name string `json:"name" yaml:"name" mapstructure:"name"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *VariableType) UnmarshalJSON(value []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(value, &raw); err != nil {
		return err
	}
	if _, ok := raw["name"]; raw != nil && !ok {
		return fmt.Errorf("field name in VariableType: required")
	}
	type Plain VariableType
	var plain Plain
	if err := json.Unmarshal(value, &plain); err != nil {
		return err
	}
	if plain.Instance != nil && len(*plain.Instance) > 50 {
		return fmt.Errorf("field %s length: must be <= %d", "instance", 50)
	}
	if len(plain.Name) > 50 {
		return fmt.Errorf("field %s length: must be <= %d", "name", 50)
	}
	*j = VariableType(plain)
	return nil
}
