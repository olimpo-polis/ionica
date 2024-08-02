// Code generated by github.com/atombender/go-jsonschema, DO NOT EDIT.

package authorizeResponse

import "encoding/json"
import "fmt"
import "reflect"
import "time"

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

type AuthorizationStatusEnumType string

const AuthorizationStatusEnumTypeAccepted AuthorizationStatusEnumType = "Accepted"
const AuthorizationStatusEnumTypeBlocked AuthorizationStatusEnumType = "Blocked"
const AuthorizationStatusEnumTypeConcurrentTx AuthorizationStatusEnumType = "ConcurrentTx"
const AuthorizationStatusEnumTypeExpired AuthorizationStatusEnumType = "Expired"
const AuthorizationStatusEnumTypeInvalid AuthorizationStatusEnumType = "Invalid"
const AuthorizationStatusEnumTypeNoCredit AuthorizationStatusEnumType = "NoCredit"
const AuthorizationStatusEnumTypeNotAllowedTypeEVSE AuthorizationStatusEnumType = "NotAllowedTypeEVSE"
const AuthorizationStatusEnumTypeNotAtThisLocation AuthorizationStatusEnumType = "NotAtThisLocation"
const AuthorizationStatusEnumTypeNotAtThisTime AuthorizationStatusEnumType = "NotAtThisTime"
const AuthorizationStatusEnumTypeUnknown AuthorizationStatusEnumType = "Unknown"

var enumValues_AuthorizationStatusEnumType = []interface{}{
	"Accepted",
	"Blocked",
	"ConcurrentTx",
	"Expired",
	"Invalid",
	"NoCredit",
	"NotAllowedTypeEVSE",
	"NotAtThisLocation",
	"NotAtThisTime",
	"Unknown",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *AuthorizationStatusEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_AuthorizationStatusEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_AuthorizationStatusEnumType, v)
	}
	*j = AuthorizationStatusEnumType(v)
	return nil
}

type AuthorizeCertificateStatusEnumType string

const AuthorizeCertificateStatusEnumTypeAccepted AuthorizeCertificateStatusEnumType = "Accepted"
const AuthorizeCertificateStatusEnumTypeCertChainError AuthorizeCertificateStatusEnumType = "CertChainError"
const AuthorizeCertificateStatusEnumTypeCertificateExpired AuthorizeCertificateStatusEnumType = "CertificateExpired"
const AuthorizeCertificateStatusEnumTypeCertificateRevoked AuthorizeCertificateStatusEnumType = "CertificateRevoked"
const AuthorizeCertificateStatusEnumTypeContractCancelled AuthorizeCertificateStatusEnumType = "ContractCancelled"
const AuthorizeCertificateStatusEnumTypeNoCertificateAvailable AuthorizeCertificateStatusEnumType = "NoCertificateAvailable"
const AuthorizeCertificateStatusEnumTypeSignatureError AuthorizeCertificateStatusEnumType = "SignatureError"

var enumValues_AuthorizeCertificateStatusEnumType = []interface{}{
	"Accepted",
	"SignatureError",
	"CertificateExpired",
	"CertificateRevoked",
	"NoCertificateAvailable",
	"CertChainError",
	"ContractCancelled",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *AuthorizeCertificateStatusEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_AuthorizeCertificateStatusEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_AuthorizeCertificateStatusEnumType, v)
	}
	*j = AuthorizeCertificateStatusEnumType(v)
	return nil
}

type AuthorizeResponseJson struct {
	// CertificateStatus corresponds to the JSON schema field "certificateStatus".
	CertificateStatus *AuthorizeCertificateStatusEnumType `json:"certificateStatus,omitempty" yaml:"certificateStatus,omitempty" mapstructure:"certificateStatus,omitempty"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType `json:"customData,omitempty" yaml:"customData,omitempty" mapstructure:"customData,omitempty"`

	// IdTokenInfo corresponds to the JSON schema field "idTokenInfo".
	IdTokenInfo IdTokenInfoType `json:"idTokenInfo" yaml:"idTokenInfo" mapstructure:"idTokenInfo"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *AuthorizeResponseJson) UnmarshalJSON(value []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(value, &raw); err != nil {
		return err
	}
	if _, ok := raw["idTokenInfo"]; raw != nil && !ok {
		return fmt.Errorf("field idTokenInfo in AuthorizeResponseJson: required")
	}
	type Plain AuthorizeResponseJson
	var plain Plain
	if err := json.Unmarshal(value, &plain); err != nil {
		return err
	}
	*j = AuthorizeResponseJson(plain)
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

// ID_ Token
// urn:x-oca:ocpp:uid:2:233247
// Contains status information about an identifier.
// It is advised to not stop charging for a token that expires during charging, as
// ExpiryDate is only used for caching purposes. If ExpiryDate is not given, the
// status has no end date.
type IdTokenInfoType struct {
	// ID_ Token. Expiry. Date_ Time
	// urn:x-oca:ocpp:uid:1:569373
	// Date and Time after which the token must be considered invalid.
	//
	CacheExpiryDateTime *time.Time `json:"cacheExpiryDateTime,omitempty" yaml:"cacheExpiryDateTime,omitempty" mapstructure:"cacheExpiryDateTime,omitempty"`

	// Priority from a business point of view. Default priority is 0, The range is
	// from -9 to 9. Higher values indicate a higher priority. The chargingPriority in
	// &lt;&lt;transactioneventresponse,TransactionEventResponse&gt;&gt; overrules
	// this one.
	//
	ChargingPriority *int `json:"chargingPriority,omitempty" yaml:"chargingPriority,omitempty" mapstructure:"chargingPriority,omitempty"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType `json:"customData,omitempty" yaml:"customData,omitempty" mapstructure:"customData,omitempty"`

	// Only used when the IdToken is only valid for one or more specific EVSEs, not
	// for the entire Charging Station.
	//
	//
	EvseId []int `json:"evseId,omitempty" yaml:"evseId,omitempty" mapstructure:"evseId,omitempty"`

	// GroupIdToken corresponds to the JSON schema field "groupIdToken".
	GroupIdToken *IdTokenType `json:"groupIdToken,omitempty" yaml:"groupIdToken,omitempty" mapstructure:"groupIdToken,omitempty"`

	// ID_ Token. Language1. Language_ Code
	// urn:x-oca:ocpp:uid:1:569374
	// Preferred user interface language of identifier user. Contains a language code
	// as defined in &lt;&lt;ref-RFC5646,[RFC5646]&gt;&gt;.
	//
	//
	Language1 *string `json:"language1,omitempty" yaml:"language1,omitempty" mapstructure:"language1,omitempty"`

	// ID_ Token. Language2. Language_ Code
	// urn:x-oca:ocpp:uid:1:569375
	// Second preferred user interface language of identifier user. Don’t use when
	// language1 is omitted, has to be different from language1. Contains a language
	// code as defined in &lt;&lt;ref-RFC5646,[RFC5646]&gt;&gt;.
	//
	Language2 *string `json:"language2,omitempty" yaml:"language2,omitempty" mapstructure:"language2,omitempty"`

	// PersonalMessage corresponds to the JSON schema field "personalMessage".
	PersonalMessage *MessageContentType `json:"personalMessage,omitempty" yaml:"personalMessage,omitempty" mapstructure:"personalMessage,omitempty"`

	// Status corresponds to the JSON schema field "status".
	Status AuthorizationStatusEnumType `json:"status" yaml:"status" mapstructure:"status"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *IdTokenInfoType) UnmarshalJSON(value []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(value, &raw); err != nil {
		return err
	}
	if _, ok := raw["status"]; raw != nil && !ok {
		return fmt.Errorf("field status in IdTokenInfoType: required")
	}
	type Plain IdTokenInfoType
	var plain Plain
	if err := json.Unmarshal(value, &plain); err != nil {
		return err
	}
	if plain.EvseId != nil && len(plain.EvseId) < 1 {
		return fmt.Errorf("field %s length: must be >= %d", "evseId", 1)
	}
	if plain.Language1 != nil && len(*plain.Language1) > 8 {
		return fmt.Errorf("field %s length: must be <= %d", "language1", 8)
	}
	if plain.Language2 != nil && len(*plain.Language2) > 8 {
		return fmt.Errorf("field %s length: must be <= %d", "language2", 8)
	}
	*j = IdTokenInfoType(plain)
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

// Message_ Content
// urn:x-enexis:ecdm:uid:2:234490
// Contains message details, for a message to be displayed on a Charging Station.
type MessageContentType struct {
	// Message_ Content. Content. Message
	// urn:x-enexis:ecdm:uid:1:570852
	// Message contents.
	//
	//
	Content string `json:"content" yaml:"content" mapstructure:"content"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType `json:"customData,omitempty" yaml:"customData,omitempty" mapstructure:"customData,omitempty"`

	// Format corresponds to the JSON schema field "format".
	Format MessageFormatEnumType `json:"format" yaml:"format" mapstructure:"format"`

	// Message_ Content. Language. Language_ Code
	// urn:x-enexis:ecdm:uid:1:570849
	// Message language identifier. Contains a language code as defined in
	// &lt;&lt;ref-RFC5646,[RFC5646]&gt;&gt;.
	//
	Language *string `json:"language,omitempty" yaml:"language,omitempty" mapstructure:"language,omitempty"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *MessageContentType) UnmarshalJSON(value []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(value, &raw); err != nil {
		return err
	}
	if _, ok := raw["content"]; raw != nil && !ok {
		return fmt.Errorf("field content in MessageContentType: required")
	}
	if _, ok := raw["format"]; raw != nil && !ok {
		return fmt.Errorf("field format in MessageContentType: required")
	}
	type Plain MessageContentType
	var plain Plain
	if err := json.Unmarshal(value, &plain); err != nil {
		return err
	}
	if len(plain.Content) > 512 {
		return fmt.Errorf("field %s length: must be <= %d", "content", 512)
	}
	if plain.Language != nil && len(*plain.Language) > 8 {
		return fmt.Errorf("field %s length: must be <= %d", "language", 8)
	}
	*j = MessageContentType(plain)
	return nil
}

type MessageFormatEnumType string

const MessageFormatEnumTypeASCII MessageFormatEnumType = "ASCII"
const MessageFormatEnumTypeHTML MessageFormatEnumType = "HTML"
const MessageFormatEnumTypeURI MessageFormatEnumType = "URI"
const MessageFormatEnumTypeUTF8 MessageFormatEnumType = "UTF8"

var enumValues_MessageFormatEnumType = []interface{}{
	"ASCII",
	"HTML",
	"URI",
	"UTF8",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *MessageFormatEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_MessageFormatEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_MessageFormatEnumType, v)
	}
	*j = MessageFormatEnumType(v)
	return nil
}
