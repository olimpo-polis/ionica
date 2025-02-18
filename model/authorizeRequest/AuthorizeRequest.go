// Code generated by github.com/atombender/go-jsonschema, DO NOT EDIT.

package authorizeRequest

import "encoding/json"
import "fmt"
import "reflect"

// Contains a case insensitive identifier to use for the authorization and the type
// of authorization to support multiple forms of identifiers.
type AdditionalInfoType struct {
	// This field specifies the additional IdToken.
	//
	AdditionalIdToken string `json:"additionalIdToken" yaml:"additionalIdToken" mapstructure:"additionalIdToken"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType `json:"customData,omitempty" yaml:"customData,omitempty" mapstructure:"customData,omitempty"`

	// This defines the type of the additionalIdToken. This is a custom type, so the
	// implementation needs to be agreed upon by all involved parties.
	//
	Type string `json:"type" yaml:"type" mapstructure:"type"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *AdditionalInfoType) UnmarshalJSON(value []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(value, &raw); err != nil {
		return err
	}
	if _, ok := raw["additionalIdToken"]; raw != nil && !ok {
		return fmt.Errorf("field additionalIdToken in AdditionalInfoType: required")
	}
	if _, ok := raw["type"]; raw != nil && !ok {
		return fmt.Errorf("field type in AdditionalInfoType: required")
	}
	type Plain AdditionalInfoType
	var plain Plain
	if err := json.Unmarshal(value, &plain); err != nil {
		return err
	}
	if len(plain.AdditionalIdToken) > 36 {
		return fmt.Errorf("field %s length: must be <= %d", "additionalIdToken", 36)
	}
	if len(plain.Type) > 50 {
		return fmt.Errorf("field %s length: must be <= %d", "type", 50)
	}
	*j = AdditionalInfoType(plain)
	return nil
}

type AuthorizeRequestJson struct {
	// The X.509 certificated presented by EV and encoded in PEM format.
	//
	Certificate *string `json:"certificate,omitempty" yaml:"certificate,omitempty" mapstructure:"certificate,omitempty"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType `json:"customData,omitempty" yaml:"customData,omitempty" mapstructure:"customData,omitempty"`

	// IdToken corresponds to the JSON schema field "idToken".
	IdToken IdTokenType `json:"idToken" yaml:"idToken" mapstructure:"idToken"`

	// Iso15118CertificateHashData corresponds to the JSON schema field
	// "iso15118CertificateHashData".
	Iso15118CertificateHashData []OCSPRequestDataType `json:"iso15118CertificateHashData,omitempty" yaml:"iso15118CertificateHashData,omitempty" mapstructure:"iso15118CertificateHashData,omitempty"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *AuthorizeRequestJson) UnmarshalJSON(value []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(value, &raw); err != nil {
		return err
	}
	if _, ok := raw["idToken"]; raw != nil && !ok {
		return fmt.Errorf("field idToken in AuthorizeRequestJson: required")
	}
	type Plain AuthorizeRequestJson
	var plain Plain
	if err := json.Unmarshal(value, &plain); err != nil {
		return err
	}
	if plain.Certificate != nil && len(*plain.Certificate) > 5500 {
		return fmt.Errorf("field %s length: must be <= %d", "certificate", 5500)
	}
	if plain.Iso15118CertificateHashData != nil && len(plain.Iso15118CertificateHashData) < 1 {
		return fmt.Errorf("field %s length: must be >= %d", "iso15118CertificateHashData", 1)
	}
	if len(plain.Iso15118CertificateHashData) > 4 {
		return fmt.Errorf("field %s length: must be <= %d", "iso15118CertificateHashData", 4)
	}
	*j = AuthorizeRequestJson(plain)
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

type HashAlgorithmEnumType string

const HashAlgorithmEnumTypeSHA256 HashAlgorithmEnumType = "SHA256"
const HashAlgorithmEnumTypeSHA384 HashAlgorithmEnumType = "SHA384"
const HashAlgorithmEnumTypeSHA512 HashAlgorithmEnumType = "SHA512"

var enumValues_HashAlgorithmEnumType = []interface{}{
	"SHA256",
	"SHA384",
	"SHA512",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *HashAlgorithmEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_HashAlgorithmEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_HashAlgorithmEnumType, v)
	}
	*j = HashAlgorithmEnumType(v)
	return nil
}

type IdTokenEnumType string

const IdTokenEnumTypeCentral IdTokenEnumType = "Central"
const IdTokenEnumTypeEMAID IdTokenEnumType = "eMAID"
const IdTokenEnumTypeISO14443 IdTokenEnumType = "ISO14443"
const IdTokenEnumTypeISO15693 IdTokenEnumType = "ISO15693"
const IdTokenEnumTypeKeyCode IdTokenEnumType = "KeyCode"
const IdTokenEnumTypeLocal IdTokenEnumType = "Local"
const IdTokenEnumTypeMacAddress IdTokenEnumType = "MacAddress"
const IdTokenEnumTypeNoAuthorization IdTokenEnumType = "NoAuthorization"

var enumValues_IdTokenEnumType = []interface{}{
	"Central",
	"eMAID",
	"ISO14443",
	"ISO15693",
	"KeyCode",
	"Local",
	"MacAddress",
	"NoAuthorization",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *IdTokenEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_IdTokenEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_IdTokenEnumType, v)
	}
	*j = IdTokenEnumType(v)
	return nil
}

// Contains a case insensitive identifier to use for the authorization and the type
// of authorization to support multiple forms of identifiers.
type IdTokenType struct {
	// AdditionalInfo corresponds to the JSON schema field "additionalInfo".
	AdditionalInfo []AdditionalInfoType `json:"additionalInfo,omitempty" yaml:"additionalInfo,omitempty" mapstructure:"additionalInfo,omitempty"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType `json:"customData,omitempty" yaml:"customData,omitempty" mapstructure:"customData,omitempty"`

	// IdToken is case insensitive. Might hold the hidden id of an RFID tag, but can
	// for example also contain a UUID.
	//
	IdToken string `json:"idToken" yaml:"idToken" mapstructure:"idToken"`

	// Type corresponds to the JSON schema field "type".
	Type IdTokenEnumType `json:"type" yaml:"type" mapstructure:"type"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *IdTokenType) UnmarshalJSON(value []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(value, &raw); err != nil {
		return err
	}
	if _, ok := raw["idToken"]; raw != nil && !ok {
		return fmt.Errorf("field idToken in IdTokenType: required")
	}
	if _, ok := raw["type"]; raw != nil && !ok {
		return fmt.Errorf("field type in IdTokenType: required")
	}
	type Plain IdTokenType
	var plain Plain
	if err := json.Unmarshal(value, &plain); err != nil {
		return err
	}
	if plain.AdditionalInfo != nil && len(plain.AdditionalInfo) < 1 {
		return fmt.Errorf("field %s length: must be >= %d", "additionalInfo", 1)
	}
	if len(plain.IdToken) > 36 {
		return fmt.Errorf("field %s length: must be <= %d", "idToken", 36)
	}
	*j = IdTokenType(plain)
	return nil
}

type OCSPRequestDataType struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType `json:"customData,omitempty" yaml:"customData,omitempty" mapstructure:"customData,omitempty"`

	// HashAlgorithm corresponds to the JSON schema field "hashAlgorithm".
	HashAlgorithm HashAlgorithmEnumType `json:"hashAlgorithm" yaml:"hashAlgorithm" mapstructure:"hashAlgorithm"`

	// Hashed value of the issuers public key
	//
	IssuerKeyHash string `json:"issuerKeyHash" yaml:"issuerKeyHash" mapstructure:"issuerKeyHash"`

	// Hashed value of the Issuer DN (Distinguished Name).
	//
	//
	IssuerNameHash string `json:"issuerNameHash" yaml:"issuerNameHash" mapstructure:"issuerNameHash"`

	// This contains the responder URL (Case insensitive).
	//
	//
	ResponderURL string `json:"responderURL" yaml:"responderURL" mapstructure:"responderURL"`

	// The serial number of the certificate.
	//
	SerialNumber string `json:"serialNumber" yaml:"serialNumber" mapstructure:"serialNumber"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *OCSPRequestDataType) UnmarshalJSON(value []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(value, &raw); err != nil {
		return err
	}
	if _, ok := raw["hashAlgorithm"]; raw != nil && !ok {
		return fmt.Errorf("field hashAlgorithm in OCSPRequestDataType: required")
	}
	if _, ok := raw["issuerKeyHash"]; raw != nil && !ok {
		return fmt.Errorf("field issuerKeyHash in OCSPRequestDataType: required")
	}
	if _, ok := raw["issuerNameHash"]; raw != nil && !ok {
		return fmt.Errorf("field issuerNameHash in OCSPRequestDataType: required")
	}
	if _, ok := raw["responderURL"]; raw != nil && !ok {
		return fmt.Errorf("field responderURL in OCSPRequestDataType: required")
	}
	if _, ok := raw["serialNumber"]; raw != nil && !ok {
		return fmt.Errorf("field serialNumber in OCSPRequestDataType: required")
	}
	type Plain OCSPRequestDataType
	var plain Plain
	if err := json.Unmarshal(value, &plain); err != nil {
		return err
	}
	if len(plain.IssuerKeyHash) > 128 {
		return fmt.Errorf("field %s length: must be <= %d", "issuerKeyHash", 128)
	}
	if len(plain.IssuerNameHash) > 128 {
		return fmt.Errorf("field %s length: must be <= %d", "issuerNameHash", 128)
	}
	if len(plain.ResponderURL) > 512 {
		return fmt.Errorf("field %s length: must be <= %d", "responderURL", 512)
	}
	if len(plain.SerialNumber) > 40 {
		return fmt.Errorf("field %s length: must be <= %d", "serialNumber", 40)
	}
	*j = OCSPRequestDataType(plain)
	return nil
}
