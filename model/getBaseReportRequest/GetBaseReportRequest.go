// Code generated by github.com/atombender/go-jsonschema, DO NOT EDIT.

package getBaseReportRequest

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

type GetBaseReportRequestJson struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType `json:"customData,omitempty" yaml:"customData,omitempty" mapstructure:"customData,omitempty"`

	// ReportBase corresponds to the JSON schema field "reportBase".
	ReportBase ReportBaseEnumType `json:"reportBase" yaml:"reportBase" mapstructure:"reportBase"`

	// The Id of the request.
	//
	RequestId int `json:"requestId" yaml:"requestId" mapstructure:"requestId"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *GetBaseReportRequestJson) UnmarshalJSON(value []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(value, &raw); err != nil {
		return err
	}
	if _, ok := raw["reportBase"]; raw != nil && !ok {
		return fmt.Errorf("field reportBase in GetBaseReportRequestJson: required")
	}
	if _, ok := raw["requestId"]; raw != nil && !ok {
		return fmt.Errorf("field requestId in GetBaseReportRequestJson: required")
	}
	type Plain GetBaseReportRequestJson
	var plain Plain
	if err := json.Unmarshal(value, &plain); err != nil {
		return err
	}
	*j = GetBaseReportRequestJson(plain)
	return nil
}

type ReportBaseEnumType string

const ReportBaseEnumTypeConfigurationInventory ReportBaseEnumType = "ConfigurationInventory"
const ReportBaseEnumTypeFullInventory ReportBaseEnumType = "FullInventory"
const ReportBaseEnumTypeSummaryInventory ReportBaseEnumType = "SummaryInventory"

var enumValues_ReportBaseEnumType = []interface{}{
	"ConfigurationInventory",
	"FullInventory",
	"SummaryInventory",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ReportBaseEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ReportBaseEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ReportBaseEnumType, v)
	}
	*j = ReportBaseEnumType(v)
	return nil
}