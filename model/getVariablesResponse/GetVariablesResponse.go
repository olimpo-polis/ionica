// Code generated by github.com/atombender/go-jsonschema, DO NOT EDIT.

package getVariablesResponse

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

// Class to hold results of GetVariables request.
type GetVariableResultType struct {
	// AttributeStatus corresponds to the JSON schema field "attributeStatus".
	AttributeStatus GetVariableStatusEnumType `json:"attributeStatus" yaml:"attributeStatus" mapstructure:"attributeStatus"`

	// AttributeStatusInfo corresponds to the JSON schema field "attributeStatusInfo".
	AttributeStatusInfo *StatusInfoType `json:"attributeStatusInfo,omitempty" yaml:"attributeStatusInfo,omitempty" mapstructure:"attributeStatusInfo,omitempty"`

	// AttributeType corresponds to the JSON schema field "attributeType".
	AttributeType *AttributeEnumType `json:"attributeType,omitempty" yaml:"attributeType,omitempty" mapstructure:"attributeType,omitempty"`

	// Value of requested attribute type of component-variable. This field can only be
	// empty when the given status is NOT accepted.
	//
	// The Configuration Variable
	// &lt;&lt;configkey-reporting-value-size,ReportingValueSize&gt;&gt; can be used
	// to limit GetVariableResult.attributeValue, VariableAttribute.value and
	// EventData.actualValue. The max size of these values will always remain equal.
	//
	//
	AttributeValue *string `json:"attributeValue,omitempty" yaml:"attributeValue,omitempty" mapstructure:"attributeValue,omitempty"`

	// Component corresponds to the JSON schema field "component".
	Component ComponentType `json:"component" yaml:"component" mapstructure:"component"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType `json:"customData,omitempty" yaml:"customData,omitempty" mapstructure:"customData,omitempty"`

	// Variable corresponds to the JSON schema field "variable".
	Variable VariableType `json:"variable" yaml:"variable" mapstructure:"variable"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *GetVariableResultType) UnmarshalJSON(value []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(value, &raw); err != nil {
		return err
	}
	if _, ok := raw["attributeStatus"]; raw != nil && !ok {
		return fmt.Errorf("field attributeStatus in GetVariableResultType: required")
	}
	if _, ok := raw["component"]; raw != nil && !ok {
		return fmt.Errorf("field component in GetVariableResultType: required")
	}
	if _, ok := raw["variable"]; raw != nil && !ok {
		return fmt.Errorf("field variable in GetVariableResultType: required")
	}
	type Plain GetVariableResultType
	var plain Plain
	if err := json.Unmarshal(value, &plain); err != nil {
		return err
	}
	if plain.AttributeValue != nil && len(*plain.AttributeValue) > 2500 {
		return fmt.Errorf("field %s length: must be <= %d", "attributeValue", 2500)
	}
	*j = GetVariableResultType(plain)
	return nil
}

type GetVariableStatusEnumType string

const GetVariableStatusEnumTypeAccepted GetVariableStatusEnumType = "Accepted"
const GetVariableStatusEnumTypeNotSupportedAttributeType GetVariableStatusEnumType = "NotSupportedAttributeType"
const GetVariableStatusEnumTypeRejected GetVariableStatusEnumType = "Rejected"
const GetVariableStatusEnumTypeUnknownComponent GetVariableStatusEnumType = "UnknownComponent"
const GetVariableStatusEnumTypeUnknownVariable GetVariableStatusEnumType = "UnknownVariable"

var enumValues_GetVariableStatusEnumType = []interface{}{
	"Accepted",
	"Rejected",
	"UnknownComponent",
	"UnknownVariable",
	"NotSupportedAttributeType",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *GetVariableStatusEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_GetVariableStatusEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_GetVariableStatusEnumType, v)
	}
	*j = GetVariableStatusEnumType(v)
	return nil
}

type GetVariablesResponseJson struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType `json:"customData,omitempty" yaml:"customData,omitempty" mapstructure:"customData,omitempty"`

	// GetVariableResult corresponds to the JSON schema field "getVariableResult".
	GetVariableResult []GetVariableResultType `json:"getVariableResult" yaml:"getVariableResult" mapstructure:"getVariableResult"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *GetVariablesResponseJson) UnmarshalJSON(value []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(value, &raw); err != nil {
		return err
	}
	if _, ok := raw["getVariableResult"]; raw != nil && !ok {
		return fmt.Errorf("field getVariableResult in GetVariablesResponseJson: required")
	}
	type Plain GetVariablesResponseJson
	var plain Plain
	if err := json.Unmarshal(value, &plain); err != nil {
		return err
	}
	if plain.GetVariableResult != nil && len(plain.GetVariableResult) < 1 {
		return fmt.Errorf("field %s length: must be >= %d", "getVariableResult", 1)
	}
	*j = GetVariablesResponseJson(plain)
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