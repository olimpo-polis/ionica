// Code generated by github.com/atombender/go-jsonschema, DO NOT EDIT.

package getLogRequest

import "encoding/json"
import "fmt"
import "reflect"
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

type GetLogRequestJson struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType `json:"customData,omitempty" yaml:"customData,omitempty" mapstructure:"customData,omitempty"`

	// Log corresponds to the JSON schema field "log".
	Log LogParametersType `json:"log" yaml:"log" mapstructure:"log"`

	// LogType corresponds to the JSON schema field "logType".
	LogType LogEnumType `json:"logType" yaml:"logType" mapstructure:"logType"`

	// The Id of this request
	//
	RequestId int `json:"requestId" yaml:"requestId" mapstructure:"requestId"`

	// This specifies how many times the Charging Station must try to upload the log
	// before giving up. If this field is not present, it is left to Charging Station
	// to decide how many times it wants to retry.
	//
	Retries *int `json:"retries,omitempty" yaml:"retries,omitempty" mapstructure:"retries,omitempty"`

	// The interval in seconds after which a retry may be attempted. If this field is
	// not present, it is left to Charging Station to decide how long to wait between
	// attempts.
	//
	RetryInterval *int `json:"retryInterval,omitempty" yaml:"retryInterval,omitempty" mapstructure:"retryInterval,omitempty"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *GetLogRequestJson) UnmarshalJSON(value []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(value, &raw); err != nil {
		return err
	}
	if _, ok := raw["log"]; raw != nil && !ok {
		return fmt.Errorf("field log in GetLogRequestJson: required")
	}
	if _, ok := raw["logType"]; raw != nil && !ok {
		return fmt.Errorf("field logType in GetLogRequestJson: required")
	}
	if _, ok := raw["requestId"]; raw != nil && !ok {
		return fmt.Errorf("field requestId in GetLogRequestJson: required")
	}
	type Plain GetLogRequestJson
	var plain Plain
	if err := json.Unmarshal(value, &plain); err != nil {
		return err
	}
	*j = GetLogRequestJson(plain)
	return nil
}

type LogEnumType string

const LogEnumTypeDiagnosticsLog LogEnumType = "DiagnosticsLog"
const LogEnumTypeSecurityLog LogEnumType = "SecurityLog"

var enumValues_LogEnumType = []interface{}{
	"DiagnosticsLog",
	"SecurityLog",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *LogEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_LogEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_LogEnumType, v)
	}
	*j = LogEnumType(v)
	return nil
}

// Log
// urn:x-enexis:ecdm:uid:2:233373
// Generic class for the configuration of logging entries.
type LogParametersType struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType `json:"customData,omitempty" yaml:"customData,omitempty" mapstructure:"customData,omitempty"`

	// Log. Latest_ Timestamp. Date_ Time
	// urn:x-enexis:ecdm:uid:1:569482
	// This contains the date and time of the latest logging information to include in
	// the diagnostics.
	//
	LatestTimestamp *time.Time `json:"latestTimestamp,omitempty" yaml:"latestTimestamp,omitempty" mapstructure:"latestTimestamp,omitempty"`

	// Log. Oldest_ Timestamp. Date_ Time
	// urn:x-enexis:ecdm:uid:1:569477
	// This contains the date and time of the oldest logging information to include in
	// the diagnostics.
	//
	OldestTimestamp *time.Time `json:"oldestTimestamp,omitempty" yaml:"oldestTimestamp,omitempty" mapstructure:"oldestTimestamp,omitempty"`

	// Log. Remote_ Location. URI
	// urn:x-enexis:ecdm:uid:1:569484
	// The URL of the location at the remote system where the log should be stored.
	//
	RemoteLocation string `json:"remoteLocation" yaml:"remoteLocation" mapstructure:"remoteLocation"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *LogParametersType) UnmarshalJSON(value []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(value, &raw); err != nil {
		return err
	}
	if _, ok := raw["remoteLocation"]; raw != nil && !ok {
		return fmt.Errorf("field remoteLocation in LogParametersType: required")
	}
	type Plain LogParametersType
	var plain Plain
	if err := json.Unmarshal(value, &plain); err != nil {
		return err
	}
	if len(plain.RemoteLocation) > 512 {
		return fmt.Errorf("field %s length: must be <= %d", "remoteLocation", 512)
	}
	*j = LogParametersType(plain)
	return nil
}
