// Code generated by github.com/atombender/go-jsonschema, DO NOT EDIT.

package get15118EVCertificateRequest

import "encoding/json"
import "fmt"
import "reflect"

type CertificateActionEnumType string

const CertificateActionEnumTypeInstall CertificateActionEnumType = "Install"
const CertificateActionEnumTypeUpdate CertificateActionEnumType = "Update"

var enumValues_CertificateActionEnumType = []interface{}{
	"Install",
	"Update",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CertificateActionEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_CertificateActionEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_CertificateActionEnumType, v)
	}
	*j = CertificateActionEnumType(v)
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

type Get15118EVCertificateRequestJson struct {
	// Action corresponds to the JSON schema field "action".
	Action CertificateActionEnumType `json:"action" yaml:"action" mapstructure:"action"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType `json:"customData,omitempty" yaml:"customData,omitempty" mapstructure:"customData,omitempty"`

	// Raw CertificateInstallationReq request from EV, Base64 encoded.
	//
	ExiRequest string `json:"exiRequest" yaml:"exiRequest" mapstructure:"exiRequest"`

	// Schema version currently used for the 15118 session between EV and Charging
	// Station. Needed for parsing of the EXI stream by the CSMS.
	//
	//
	Iso15118SchemaVersion string `json:"iso15118SchemaVersion" yaml:"iso15118SchemaVersion" mapstructure:"iso15118SchemaVersion"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *Get15118EVCertificateRequestJson) UnmarshalJSON(value []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(value, &raw); err != nil {
		return err
	}
	if _, ok := raw["action"]; raw != nil && !ok {
		return fmt.Errorf("field action in Get15118EVCertificateRequestJson: required")
	}
	if _, ok := raw["exiRequest"]; raw != nil && !ok {
		return fmt.Errorf("field exiRequest in Get15118EVCertificateRequestJson: required")
	}
	if _, ok := raw["iso15118SchemaVersion"]; raw != nil && !ok {
		return fmt.Errorf("field iso15118SchemaVersion in Get15118EVCertificateRequestJson: required")
	}
	type Plain Get15118EVCertificateRequestJson
	var plain Plain
	if err := json.Unmarshal(value, &plain); err != nil {
		return err
	}
	if len(plain.ExiRequest) > 5600 {
		return fmt.Errorf("field %s length: must be <= %d", "exiRequest", 5600)
	}
	if len(plain.Iso15118SchemaVersion) > 50 {
		return fmt.Errorf("field %s length: must be <= %d", "iso15118SchemaVersion", 50)
	}
	*j = Get15118EVCertificateRequestJson(plain)
	return nil
}
