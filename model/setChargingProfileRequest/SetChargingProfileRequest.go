// Code generated by github.com/atombender/go-jsonschema, DO NOT EDIT.

package setChargingProfileRequest

import "encoding/json"
import "fmt"
import "reflect"
import "time"

type ChargingProfileKindEnumType string

const ChargingProfileKindEnumTypeAbsolute ChargingProfileKindEnumType = "Absolute"
const ChargingProfileKindEnumTypeRecurring ChargingProfileKindEnumType = "Recurring"
const ChargingProfileKindEnumTypeRelative ChargingProfileKindEnumType = "Relative"

var enumValues_ChargingProfileKindEnumType = []interface{}{
	"Absolute",
	"Recurring",
	"Relative",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ChargingProfileKindEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ChargingProfileKindEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ChargingProfileKindEnumType, v)
	}
	*j = ChargingProfileKindEnumType(v)
	return nil
}

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

// Charging_ Profile
// urn:x-oca:ocpp:uid:2:233255
// A ChargingProfile consists of ChargingSchedule, describing the amount of power
// or current that can be delivered per time interval.
type ChargingProfileType struct {
	// ChargingProfileKind corresponds to the JSON schema field "chargingProfileKind".
	ChargingProfileKind ChargingProfileKindEnumType `json:"chargingProfileKind" yaml:"chargingProfileKind" mapstructure:"chargingProfileKind"`

	// ChargingProfilePurpose corresponds to the JSON schema field
	// "chargingProfilePurpose".
	ChargingProfilePurpose ChargingProfilePurposeEnumType `json:"chargingProfilePurpose" yaml:"chargingProfilePurpose" mapstructure:"chargingProfilePurpose"`

	// ChargingSchedule corresponds to the JSON schema field "chargingSchedule".
	ChargingSchedule []ChargingScheduleType `json:"chargingSchedule" yaml:"chargingSchedule" mapstructure:"chargingSchedule"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType `json:"customData,omitempty" yaml:"customData,omitempty" mapstructure:"customData,omitempty"`

	// Identified_ Object. MRID. Numeric_ Identifier
	// urn:x-enexis:ecdm:uid:1:569198
	// Id of ChargingProfile.
	//
	Id int `json:"id" yaml:"id" mapstructure:"id"`

	// RecurrencyKind corresponds to the JSON schema field "recurrencyKind".
	RecurrencyKind *RecurrencyKindEnumType `json:"recurrencyKind,omitempty" yaml:"recurrencyKind,omitempty" mapstructure:"recurrencyKind,omitempty"`

	// Charging_ Profile. Stack_ Level. Counter
	// urn:x-oca:ocpp:uid:1:569230
	// Value determining level in hierarchy stack of profiles. Higher values have
	// precedence over lower values. Lowest level is 0.
	//
	StackLevel int `json:"stackLevel" yaml:"stackLevel" mapstructure:"stackLevel"`

	// SHALL only be included if ChargingProfilePurpose is set to TxProfile. The
	// transactionId is used to match the profile to a specific transaction.
	//
	TransactionId *string `json:"transactionId,omitempty" yaml:"transactionId,omitempty" mapstructure:"transactionId,omitempty"`

	// Charging_ Profile. Valid_ From. Date_ Time
	// urn:x-oca:ocpp:uid:1:569234
	// Point in time at which the profile starts to be valid. If absent, the profile
	// is valid as soon as it is received by the Charging Station.
	//
	ValidFrom *time.Time `json:"validFrom,omitempty" yaml:"validFrom,omitempty" mapstructure:"validFrom,omitempty"`

	// Charging_ Profile. Valid_ To. Date_ Time
	// urn:x-oca:ocpp:uid:1:569235
	// Point in time at which the profile stops to be valid. If absent, the profile is
	// valid until it is replaced by another profile.
	//
	ValidTo *time.Time `json:"validTo,omitempty" yaml:"validTo,omitempty" mapstructure:"validTo,omitempty"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ChargingProfileType) UnmarshalJSON(value []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(value, &raw); err != nil {
		return err
	}
	if _, ok := raw["chargingProfileKind"]; raw != nil && !ok {
		return fmt.Errorf("field chargingProfileKind in ChargingProfileType: required")
	}
	if _, ok := raw["chargingProfilePurpose"]; raw != nil && !ok {
		return fmt.Errorf("field chargingProfilePurpose in ChargingProfileType: required")
	}
	if _, ok := raw["chargingSchedule"]; raw != nil && !ok {
		return fmt.Errorf("field chargingSchedule in ChargingProfileType: required")
	}
	if _, ok := raw["id"]; raw != nil && !ok {
		return fmt.Errorf("field id in ChargingProfileType: required")
	}
	if _, ok := raw["stackLevel"]; raw != nil && !ok {
		return fmt.Errorf("field stackLevel in ChargingProfileType: required")
	}
	type Plain ChargingProfileType
	var plain Plain
	if err := json.Unmarshal(value, &plain); err != nil {
		return err
	}
	if plain.ChargingSchedule != nil && len(plain.ChargingSchedule) < 1 {
		return fmt.Errorf("field %s length: must be >= %d", "chargingSchedule", 1)
	}
	if len(plain.ChargingSchedule) > 3 {
		return fmt.Errorf("field %s length: must be <= %d", "chargingSchedule", 3)
	}
	if plain.TransactionId != nil && len(*plain.TransactionId) > 36 {
		return fmt.Errorf("field %s length: must be <= %d", "transactionId", 36)
	}
	*j = ChargingProfileType(plain)
	return nil
}

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

// Charging_ Schedule_ Period
// urn:x-oca:ocpp:uid:2:233257
// Charging schedule period structure defines a time period in a charging schedule.
type ChargingSchedulePeriodType struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType `json:"customData,omitempty" yaml:"customData,omitempty" mapstructure:"customData,omitempty"`

	// Charging_ Schedule_ Period. Limit. Measure
	// urn:x-oca:ocpp:uid:1:569241
	// Charging rate limit during the schedule period, in the applicable
	// chargingRateUnit, for example in Amperes (A) or Watts (W). Accepts at most one
	// digit fraction (e.g. 8.1).
	//
	Limit float64 `json:"limit" yaml:"limit" mapstructure:"limit"`

	// Charging_ Schedule_ Period. Number_ Phases. Counter
	// urn:x-oca:ocpp:uid:1:569242
	// The number of phases that can be used for charging. If a number of phases is
	// needed, numberPhases=3 will be assumed unless another number is given.
	//
	NumberPhases *int `json:"numberPhases,omitempty" yaml:"numberPhases,omitempty" mapstructure:"numberPhases,omitempty"`

	// Values: 1..3, Used if numberPhases=1 and if the EVSE is capable of switching
	// the phase connected to the EV, i.e. ACPhaseSwitchingSupported is defined and
	// true. It’s not allowed unless both conditions above are true. If both
	// conditions are true, and phaseToUse is omitted, the Charging Station / EVSE
	// will make the selection on its own.
	//
	//
	PhaseToUse *int `json:"phaseToUse,omitempty" yaml:"phaseToUse,omitempty" mapstructure:"phaseToUse,omitempty"`

	// Charging_ Schedule_ Period. Start_ Period. Elapsed_ Time
	// urn:x-oca:ocpp:uid:1:569240
	// Start of the period, in seconds from the start of schedule. The value of
	// StartPeriod also defines the stop time of the previous period.
	//
	StartPeriod int `json:"startPeriod" yaml:"startPeriod" mapstructure:"startPeriod"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ChargingSchedulePeriodType) UnmarshalJSON(value []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(value, &raw); err != nil {
		return err
	}
	if _, ok := raw["limit"]; raw != nil && !ok {
		return fmt.Errorf("field limit in ChargingSchedulePeriodType: required")
	}
	if _, ok := raw["startPeriod"]; raw != nil && !ok {
		return fmt.Errorf("field startPeriod in ChargingSchedulePeriodType: required")
	}
	type Plain ChargingSchedulePeriodType
	var plain Plain
	if err := json.Unmarshal(value, &plain); err != nil {
		return err
	}
	*j = ChargingSchedulePeriodType(plain)
	return nil
}

// Charging_ Schedule
// urn:x-oca:ocpp:uid:2:233256
// Charging schedule structure defines a list of charging periods, as used in:
// GetCompositeSchedule.conf and ChargingProfile.
type ChargingScheduleType struct {
	// ChargingRateUnit corresponds to the JSON schema field "chargingRateUnit".
	ChargingRateUnit ChargingRateUnitEnumType `json:"chargingRateUnit" yaml:"chargingRateUnit" mapstructure:"chargingRateUnit"`

	// ChargingSchedulePeriod corresponds to the JSON schema field
	// "chargingSchedulePeriod".
	ChargingSchedulePeriod []ChargingSchedulePeriodType `json:"chargingSchedulePeriod" yaml:"chargingSchedulePeriod" mapstructure:"chargingSchedulePeriod"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType `json:"customData,omitempty" yaml:"customData,omitempty" mapstructure:"customData,omitempty"`

	// Charging_ Schedule. Duration. Elapsed_ Time
	// urn:x-oca:ocpp:uid:1:569236
	// Duration of the charging schedule in seconds. If the duration is left empty,
	// the last period will continue indefinitely or until end of the transaction if
	// chargingProfilePurpose = TxProfile.
	//
	Duration *int `json:"duration,omitempty" yaml:"duration,omitempty" mapstructure:"duration,omitempty"`

	// Identifies the ChargingSchedule.
	//
	Id int `json:"id" yaml:"id" mapstructure:"id"`

	// Charging_ Schedule. Min_ Charging_ Rate. Numeric
	// urn:x-oca:ocpp:uid:1:569239
	// Minimum charging rate supported by the EV. The unit of measure is defined by
	// the chargingRateUnit. This parameter is intended to be used by a local smart
	// charging algorithm to optimize the power allocation for in the case a charging
	// process is inefficient at lower charging rates. Accepts at most one digit
	// fraction (e.g. 8.1)
	//
	MinChargingRate *float64 `json:"minChargingRate,omitempty" yaml:"minChargingRate,omitempty" mapstructure:"minChargingRate,omitempty"`

	// SalesTariff corresponds to the JSON schema field "salesTariff".
	SalesTariff *SalesTariffType `json:"salesTariff,omitempty" yaml:"salesTariff,omitempty" mapstructure:"salesTariff,omitempty"`

	// Charging_ Schedule. Start_ Schedule. Date_ Time
	// urn:x-oca:ocpp:uid:1:569237
	// Starting point of an absolute schedule. If absent the schedule will be relative
	// to start of charging.
	//
	StartSchedule *time.Time `json:"startSchedule,omitempty" yaml:"startSchedule,omitempty" mapstructure:"startSchedule,omitempty"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ChargingScheduleType) UnmarshalJSON(value []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(value, &raw); err != nil {
		return err
	}
	if _, ok := raw["chargingRateUnit"]; raw != nil && !ok {
		return fmt.Errorf("field chargingRateUnit in ChargingScheduleType: required")
	}
	if _, ok := raw["chargingSchedulePeriod"]; raw != nil && !ok {
		return fmt.Errorf("field chargingSchedulePeriod in ChargingScheduleType: required")
	}
	if _, ok := raw["id"]; raw != nil && !ok {
		return fmt.Errorf("field id in ChargingScheduleType: required")
	}
	type Plain ChargingScheduleType
	var plain Plain
	if err := json.Unmarshal(value, &plain); err != nil {
		return err
	}
	if plain.ChargingSchedulePeriod != nil && len(plain.ChargingSchedulePeriod) < 1 {
		return fmt.Errorf("field %s length: must be >= %d", "chargingSchedulePeriod", 1)
	}
	if len(plain.ChargingSchedulePeriod) > 1024 {
		return fmt.Errorf("field %s length: must be <= %d", "chargingSchedulePeriod", 1024)
	}
	*j = ChargingScheduleType(plain)
	return nil
}

// Consumption_ Cost
// urn:x-oca:ocpp:uid:2:233259
type ConsumptionCostType struct {
	// Cost corresponds to the JSON schema field "cost".
	Cost []CostType `json:"cost" yaml:"cost" mapstructure:"cost"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType `json:"customData,omitempty" yaml:"customData,omitempty" mapstructure:"customData,omitempty"`

	// Consumption_ Cost. Start_ Value. Numeric
	// urn:x-oca:ocpp:uid:1:569246
	// The lowest level of consumption that defines the starting point of this
	// consumption block. The block interval extends to the start of the next
	// interval.
	//
	StartValue float64 `json:"startValue" yaml:"startValue" mapstructure:"startValue"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ConsumptionCostType) UnmarshalJSON(value []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(value, &raw); err != nil {
		return err
	}
	if _, ok := raw["cost"]; raw != nil && !ok {
		return fmt.Errorf("field cost in ConsumptionCostType: required")
	}
	if _, ok := raw["startValue"]; raw != nil && !ok {
		return fmt.Errorf("field startValue in ConsumptionCostType: required")
	}
	type Plain ConsumptionCostType
	var plain Plain
	if err := json.Unmarshal(value, &plain); err != nil {
		return err
	}
	if plain.Cost != nil && len(plain.Cost) < 1 {
		return fmt.Errorf("field %s length: must be >= %d", "cost", 1)
	}
	if len(plain.Cost) > 3 {
		return fmt.Errorf("field %s length: must be <= %d", "cost", 3)
	}
	*j = ConsumptionCostType(plain)
	return nil
}

type CostKindEnumType string

const CostKindEnumTypeCarbonDioxideEmission CostKindEnumType = "CarbonDioxideEmission"
const CostKindEnumTypeRelativePricePercentage CostKindEnumType = "RelativePricePercentage"
const CostKindEnumTypeRenewableGenerationPercentage CostKindEnumType = "RenewableGenerationPercentage"

var enumValues_CostKindEnumType = []interface{}{
	"CarbonDioxideEmission",
	"RelativePricePercentage",
	"RenewableGenerationPercentage",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CostKindEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_CostKindEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_CostKindEnumType, v)
	}
	*j = CostKindEnumType(v)
	return nil
}

// Cost
// urn:x-oca:ocpp:uid:2:233258
type CostType struct {
	// Cost. Amount. Amount
	// urn:x-oca:ocpp:uid:1:569244
	// The estimated or actual cost per kWh
	//
	Amount int `json:"amount" yaml:"amount" mapstructure:"amount"`

	// Cost. Amount_ Multiplier. Integer
	// urn:x-oca:ocpp:uid:1:569245
	// Values: -3..3, The amountMultiplier defines the exponent to base 10 (dec). The
	// final value is determined by: amount * 10 ^ amountMultiplier
	//
	AmountMultiplier *int `json:"amountMultiplier,omitempty" yaml:"amountMultiplier,omitempty" mapstructure:"amountMultiplier,omitempty"`

	// CostKind corresponds to the JSON schema field "costKind".
	CostKind CostKindEnumType `json:"costKind" yaml:"costKind" mapstructure:"costKind"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType `json:"customData,omitempty" yaml:"customData,omitempty" mapstructure:"customData,omitempty"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CostType) UnmarshalJSON(value []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(value, &raw); err != nil {
		return err
	}
	if _, ok := raw["amount"]; raw != nil && !ok {
		return fmt.Errorf("field amount in CostType: required")
	}
	if _, ok := raw["costKind"]; raw != nil && !ok {
		return fmt.Errorf("field costKind in CostType: required")
	}
	type Plain CostType
	var plain Plain
	if err := json.Unmarshal(value, &plain); err != nil {
		return err
	}
	*j = CostType(plain)
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

type RecurrencyKindEnumType string

const RecurrencyKindEnumTypeDaily RecurrencyKindEnumType = "Daily"
const RecurrencyKindEnumTypeWeekly RecurrencyKindEnumType = "Weekly"

var enumValues_RecurrencyKindEnumType = []interface{}{
	"Daily",
	"Weekly",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *RecurrencyKindEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_RecurrencyKindEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_RecurrencyKindEnumType, v)
	}
	*j = RecurrencyKindEnumType(v)
	return nil
}

// Relative_ Timer_ Interval
// urn:x-oca:ocpp:uid:2:233270
type RelativeTimeIntervalType struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType `json:"customData,omitempty" yaml:"customData,omitempty" mapstructure:"customData,omitempty"`

	// Relative_ Timer_ Interval. Duration. Elapsed_ Time
	// urn:x-oca:ocpp:uid:1:569280
	// Duration of the interval, in seconds.
	//
	Duration *int `json:"duration,omitempty" yaml:"duration,omitempty" mapstructure:"duration,omitempty"`

	// Relative_ Timer_ Interval. Start. Elapsed_ Time
	// urn:x-oca:ocpp:uid:1:569279
	// Start of the interval, in seconds from NOW.
	//
	Start int `json:"start" yaml:"start" mapstructure:"start"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *RelativeTimeIntervalType) UnmarshalJSON(value []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(value, &raw); err != nil {
		return err
	}
	if _, ok := raw["start"]; raw != nil && !ok {
		return fmt.Errorf("field start in RelativeTimeIntervalType: required")
	}
	type Plain RelativeTimeIntervalType
	var plain Plain
	if err := json.Unmarshal(value, &plain); err != nil {
		return err
	}
	*j = RelativeTimeIntervalType(plain)
	return nil
}

// Sales_ Tariff_ Entry
// urn:x-oca:ocpp:uid:2:233271
type SalesTariffEntryType struct {
	// ConsumptionCost corresponds to the JSON schema field "consumptionCost".
	ConsumptionCost []ConsumptionCostType `json:"consumptionCost,omitempty" yaml:"consumptionCost,omitempty" mapstructure:"consumptionCost,omitempty"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType `json:"customData,omitempty" yaml:"customData,omitempty" mapstructure:"customData,omitempty"`

	// Sales_ Tariff_ Entry. E_ Price_ Level. Unsigned_ Integer
	// urn:x-oca:ocpp:uid:1:569281
	// Defines the price level of this SalesTariffEntry (referring to
	// NumEPriceLevels). Small values for the EPriceLevel represent a cheaper
	// TariffEntry. Large values for the EPriceLevel represent a more expensive
	// TariffEntry.
	//
	EPriceLevel *int `json:"ePriceLevel,omitempty" yaml:"ePriceLevel,omitempty" mapstructure:"ePriceLevel,omitempty"`

	// RelativeTimeInterval corresponds to the JSON schema field
	// "relativeTimeInterval".
	RelativeTimeInterval RelativeTimeIntervalType `json:"relativeTimeInterval" yaml:"relativeTimeInterval" mapstructure:"relativeTimeInterval"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *SalesTariffEntryType) UnmarshalJSON(value []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(value, &raw); err != nil {
		return err
	}
	if _, ok := raw["relativeTimeInterval"]; raw != nil && !ok {
		return fmt.Errorf("field relativeTimeInterval in SalesTariffEntryType: required")
	}
	type Plain SalesTariffEntryType
	var plain Plain
	if err := json.Unmarshal(value, &plain); err != nil {
		return err
	}
	if plain.ConsumptionCost != nil && len(plain.ConsumptionCost) < 1 {
		return fmt.Errorf("field %s length: must be >= %d", "consumptionCost", 1)
	}
	if len(plain.ConsumptionCost) > 3 {
		return fmt.Errorf("field %s length: must be <= %d", "consumptionCost", 3)
	}
	*j = SalesTariffEntryType(plain)
	return nil
}

// Sales_ Tariff
// urn:x-oca:ocpp:uid:2:233272
// NOTE: This dataType is based on dataTypes from &lt;&lt;ref-ISOIEC15118-2,ISO
// 15118-2&gt;&gt;.
type SalesTariffType struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType `json:"customData,omitempty" yaml:"customData,omitempty" mapstructure:"customData,omitempty"`

	// Identified_ Object. MRID. Numeric_ Identifier
	// urn:x-enexis:ecdm:uid:1:569198
	// SalesTariff identifier used to identify one sales tariff. An SAID remains a
	// unique identifier for one schedule throughout a charging session.
	//
	Id int `json:"id" yaml:"id" mapstructure:"id"`

	// Sales_ Tariff. Num_ E_ Price_ Levels. Counter
	// urn:x-oca:ocpp:uid:1:569284
	// Defines the overall number of distinct price levels used across all provided
	// SalesTariff elements.
	//
	NumEPriceLevels *int `json:"numEPriceLevels,omitempty" yaml:"numEPriceLevels,omitempty" mapstructure:"numEPriceLevels,omitempty"`

	// Sales_ Tariff. Sales. Tariff_ Description
	// urn:x-oca:ocpp:uid:1:569283
	// A human readable title/short description of the sales tariff e.g. for HMI
	// display purposes.
	//
	SalesTariffDescription *string `json:"salesTariffDescription,omitempty" yaml:"salesTariffDescription,omitempty" mapstructure:"salesTariffDescription,omitempty"`

	// SalesTariffEntry corresponds to the JSON schema field "salesTariffEntry".
	SalesTariffEntry []SalesTariffEntryType `json:"salesTariffEntry" yaml:"salesTariffEntry" mapstructure:"salesTariffEntry"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *SalesTariffType) UnmarshalJSON(value []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(value, &raw); err != nil {
		return err
	}
	if _, ok := raw["id"]; raw != nil && !ok {
		return fmt.Errorf("field id in SalesTariffType: required")
	}
	if _, ok := raw["salesTariffEntry"]; raw != nil && !ok {
		return fmt.Errorf("field salesTariffEntry in SalesTariffType: required")
	}
	type Plain SalesTariffType
	var plain Plain
	if err := json.Unmarshal(value, &plain); err != nil {
		return err
	}
	if plain.SalesTariffDescription != nil && len(*plain.SalesTariffDescription) > 32 {
		return fmt.Errorf("field %s length: must be <= %d", "salesTariffDescription", 32)
	}
	if plain.SalesTariffEntry != nil && len(plain.SalesTariffEntry) < 1 {
		return fmt.Errorf("field %s length: must be >= %d", "salesTariffEntry", 1)
	}
	if len(plain.SalesTariffEntry) > 1024 {
		return fmt.Errorf("field %s length: must be <= %d", "salesTariffEntry", 1024)
	}
	*j = SalesTariffType(plain)
	return nil
}

type SetChargingProfileRequestJson struct {
	// ChargingProfile corresponds to the JSON schema field "chargingProfile".
	ChargingProfile ChargingProfileType `json:"chargingProfile" yaml:"chargingProfile" mapstructure:"chargingProfile"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType `json:"customData,omitempty" yaml:"customData,omitempty" mapstructure:"customData,omitempty"`

	// For TxDefaultProfile an evseId=0 applies the profile to each individual evse.
	// For ChargingStationMaxProfile and ChargingStationExternalConstraints an
	// evseId=0 contains an overal limit for the whole Charging Station.
	//
	EvseId int `json:"evseId" yaml:"evseId" mapstructure:"evseId"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *SetChargingProfileRequestJson) UnmarshalJSON(value []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(value, &raw); err != nil {
		return err
	}
	if _, ok := raw["chargingProfile"]; raw != nil && !ok {
		return fmt.Errorf("field chargingProfile in SetChargingProfileRequestJson: required")
	}
	if _, ok := raw["evseId"]; raw != nil && !ok {
		return fmt.Errorf("field evseId in SetChargingProfileRequestJson: required")
	}
	type Plain SetChargingProfileRequestJson
	var plain Plain
	if err := json.Unmarshal(value, &plain); err != nil {
		return err
	}
	*j = SetChargingProfileRequestJson(plain)
	return nil
}
