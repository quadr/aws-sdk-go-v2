// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type ActionValue string

// Enum values for ActionValue
const (
	ActionValueAllow ActionValue = "ALLOW"
	ActionValueBlock ActionValue = "BLOCK"
	ActionValueCount ActionValue = "COUNT"
)

// Values returns all known values for ActionValue. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (ActionValue) Values() []ActionValue {
	return []ActionValue{
		"ALLOW",
		"BLOCK",
		"COUNT",
	}
}

type BodyParsingFallbackBehavior string

// Enum values for BodyParsingFallbackBehavior
const (
	BodyParsingFallbackBehaviorMatch            BodyParsingFallbackBehavior = "MATCH"
	BodyParsingFallbackBehaviorNoMatch          BodyParsingFallbackBehavior = "NO_MATCH"
	BodyParsingFallbackBehaviorEvaluateAsString BodyParsingFallbackBehavior = "EVALUATE_AS_STRING"
)

// Values returns all known values for BodyParsingFallbackBehavior. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (BodyParsingFallbackBehavior) Values() []BodyParsingFallbackBehavior {
	return []BodyParsingFallbackBehavior{
		"MATCH",
		"NO_MATCH",
		"EVALUATE_AS_STRING",
	}
}

type ComparisonOperator string

// Enum values for ComparisonOperator
const (
	ComparisonOperatorEq ComparisonOperator = "EQ"
	ComparisonOperatorNe ComparisonOperator = "NE"
	ComparisonOperatorLe ComparisonOperator = "LE"
	ComparisonOperatorLt ComparisonOperator = "LT"
	ComparisonOperatorGe ComparisonOperator = "GE"
	ComparisonOperatorGt ComparisonOperator = "GT"
)

// Values returns all known values for ComparisonOperator. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ComparisonOperator) Values() []ComparisonOperator {
	return []ComparisonOperator{
		"EQ",
		"NE",
		"LE",
		"LT",
		"GE",
		"GT",
	}
}

type CountryCode string

// Enum values for CountryCode
const (
	CountryCodeAf CountryCode = "AF"
	CountryCodeAx CountryCode = "AX"
	CountryCodeAl CountryCode = "AL"
	CountryCodeDz CountryCode = "DZ"
	CountryCodeAs CountryCode = "AS"
	CountryCodeAd CountryCode = "AD"
	CountryCodeAo CountryCode = "AO"
	CountryCodeAi CountryCode = "AI"
	CountryCodeAq CountryCode = "AQ"
	CountryCodeAg CountryCode = "AG"
	CountryCodeAr CountryCode = "AR"
	CountryCodeAm CountryCode = "AM"
	CountryCodeAw CountryCode = "AW"
	CountryCodeAu CountryCode = "AU"
	CountryCodeAt CountryCode = "AT"
	CountryCodeAz CountryCode = "AZ"
	CountryCodeBs CountryCode = "BS"
	CountryCodeBh CountryCode = "BH"
	CountryCodeBd CountryCode = "BD"
	CountryCodeBb CountryCode = "BB"
	CountryCodeBy CountryCode = "BY"
	CountryCodeBe CountryCode = "BE"
	CountryCodeBz CountryCode = "BZ"
	CountryCodeBj CountryCode = "BJ"
	CountryCodeBm CountryCode = "BM"
	CountryCodeBt CountryCode = "BT"
	CountryCodeBo CountryCode = "BO"
	CountryCodeBq CountryCode = "BQ"
	CountryCodeBa CountryCode = "BA"
	CountryCodeBw CountryCode = "BW"
	CountryCodeBv CountryCode = "BV"
	CountryCodeBr CountryCode = "BR"
	CountryCodeIo CountryCode = "IO"
	CountryCodeBn CountryCode = "BN"
	CountryCodeBg CountryCode = "BG"
	CountryCodeBf CountryCode = "BF"
	CountryCodeBi CountryCode = "BI"
	CountryCodeKh CountryCode = "KH"
	CountryCodeCm CountryCode = "CM"
	CountryCodeCa CountryCode = "CA"
	CountryCodeCv CountryCode = "CV"
	CountryCodeKy CountryCode = "KY"
	CountryCodeCf CountryCode = "CF"
	CountryCodeTd CountryCode = "TD"
	CountryCodeCl CountryCode = "CL"
	CountryCodeCn CountryCode = "CN"
	CountryCodeCx CountryCode = "CX"
	CountryCodeCc CountryCode = "CC"
	CountryCodeCo CountryCode = "CO"
	CountryCodeKm CountryCode = "KM"
	CountryCodeCg CountryCode = "CG"
	CountryCodeCd CountryCode = "CD"
	CountryCodeCk CountryCode = "CK"
	CountryCodeCr CountryCode = "CR"
	CountryCodeCi CountryCode = "CI"
	CountryCodeHr CountryCode = "HR"
	CountryCodeCu CountryCode = "CU"
	CountryCodeCw CountryCode = "CW"
	CountryCodeCy CountryCode = "CY"
	CountryCodeCz CountryCode = "CZ"
	CountryCodeDk CountryCode = "DK"
	CountryCodeDj CountryCode = "DJ"
	CountryCodeDm CountryCode = "DM"
	CountryCodeDo CountryCode = "DO"
	CountryCodeEc CountryCode = "EC"
	CountryCodeEg CountryCode = "EG"
	CountryCodeSv CountryCode = "SV"
	CountryCodeGq CountryCode = "GQ"
	CountryCodeEr CountryCode = "ER"
	CountryCodeEe CountryCode = "EE"
	CountryCodeEt CountryCode = "ET"
	CountryCodeFk CountryCode = "FK"
	CountryCodeFo CountryCode = "FO"
	CountryCodeFj CountryCode = "FJ"
	CountryCodeFi CountryCode = "FI"
	CountryCodeFr CountryCode = "FR"
	CountryCodeGf CountryCode = "GF"
	CountryCodePf CountryCode = "PF"
	CountryCodeTf CountryCode = "TF"
	CountryCodeGa CountryCode = "GA"
	CountryCodeGm CountryCode = "GM"
	CountryCodeGe CountryCode = "GE"
	CountryCodeDe CountryCode = "DE"
	CountryCodeGh CountryCode = "GH"
	CountryCodeGi CountryCode = "GI"
	CountryCodeGr CountryCode = "GR"
	CountryCodeGl CountryCode = "GL"
	CountryCodeGd CountryCode = "GD"
	CountryCodeGp CountryCode = "GP"
	CountryCodeGu CountryCode = "GU"
	CountryCodeGt CountryCode = "GT"
	CountryCodeGg CountryCode = "GG"
	CountryCodeGn CountryCode = "GN"
	CountryCodeGw CountryCode = "GW"
	CountryCodeGy CountryCode = "GY"
	CountryCodeHt CountryCode = "HT"
	CountryCodeHm CountryCode = "HM"
	CountryCodeVa CountryCode = "VA"
	CountryCodeHn CountryCode = "HN"
	CountryCodeHk CountryCode = "HK"
	CountryCodeHu CountryCode = "HU"
	CountryCodeIs CountryCode = "IS"
	CountryCodeIn CountryCode = "IN"
	CountryCodeId CountryCode = "ID"
	CountryCodeIr CountryCode = "IR"
	CountryCodeIq CountryCode = "IQ"
	CountryCodeIe CountryCode = "IE"
	CountryCodeIm CountryCode = "IM"
	CountryCodeIl CountryCode = "IL"
	CountryCodeIt CountryCode = "IT"
	CountryCodeJm CountryCode = "JM"
	CountryCodeJp CountryCode = "JP"
	CountryCodeJe CountryCode = "JE"
	CountryCodeJo CountryCode = "JO"
	CountryCodeKz CountryCode = "KZ"
	CountryCodeKe CountryCode = "KE"
	CountryCodeKi CountryCode = "KI"
	CountryCodeKp CountryCode = "KP"
	CountryCodeKr CountryCode = "KR"
	CountryCodeKw CountryCode = "KW"
	CountryCodeKg CountryCode = "KG"
	CountryCodeLa CountryCode = "LA"
	CountryCodeLv CountryCode = "LV"
	CountryCodeLb CountryCode = "LB"
	CountryCodeLs CountryCode = "LS"
	CountryCodeLr CountryCode = "LR"
	CountryCodeLy CountryCode = "LY"
	CountryCodeLi CountryCode = "LI"
	CountryCodeLt CountryCode = "LT"
	CountryCodeLu CountryCode = "LU"
	CountryCodeMo CountryCode = "MO"
	CountryCodeMk CountryCode = "MK"
	CountryCodeMg CountryCode = "MG"
	CountryCodeMw CountryCode = "MW"
	CountryCodeMy CountryCode = "MY"
	CountryCodeMv CountryCode = "MV"
	CountryCodeMl CountryCode = "ML"
	CountryCodeMt CountryCode = "MT"
	CountryCodeMh CountryCode = "MH"
	CountryCodeMq CountryCode = "MQ"
	CountryCodeMr CountryCode = "MR"
	CountryCodeMu CountryCode = "MU"
	CountryCodeYt CountryCode = "YT"
	CountryCodeMx CountryCode = "MX"
	CountryCodeFm CountryCode = "FM"
	CountryCodeMd CountryCode = "MD"
	CountryCodeMc CountryCode = "MC"
	CountryCodeMn CountryCode = "MN"
	CountryCodeMe CountryCode = "ME"
	CountryCodeMs CountryCode = "MS"
	CountryCodeMa CountryCode = "MA"
	CountryCodeMz CountryCode = "MZ"
	CountryCodeMm CountryCode = "MM"
	CountryCodeNa CountryCode = "NA"
	CountryCodeNr CountryCode = "NR"
	CountryCodeNp CountryCode = "NP"
	CountryCodeNl CountryCode = "NL"
	CountryCodeNc CountryCode = "NC"
	CountryCodeNz CountryCode = "NZ"
	CountryCodeNi CountryCode = "NI"
	CountryCodeNe CountryCode = "NE"
	CountryCodeNg CountryCode = "NG"
	CountryCodeNu CountryCode = "NU"
	CountryCodeNf CountryCode = "NF"
	CountryCodeMp CountryCode = "MP"
	CountryCodeNo CountryCode = "NO"
	CountryCodeOm CountryCode = "OM"
	CountryCodePk CountryCode = "PK"
	CountryCodePw CountryCode = "PW"
	CountryCodePs CountryCode = "PS"
	CountryCodePa CountryCode = "PA"
	CountryCodePg CountryCode = "PG"
	CountryCodePy CountryCode = "PY"
	CountryCodePe CountryCode = "PE"
	CountryCodePh CountryCode = "PH"
	CountryCodePn CountryCode = "PN"
	CountryCodePl CountryCode = "PL"
	CountryCodePt CountryCode = "PT"
	CountryCodePr CountryCode = "PR"
	CountryCodeQa CountryCode = "QA"
	CountryCodeRe CountryCode = "RE"
	CountryCodeRo CountryCode = "RO"
	CountryCodeRu CountryCode = "RU"
	CountryCodeRw CountryCode = "RW"
	CountryCodeBl CountryCode = "BL"
	CountryCodeSh CountryCode = "SH"
	CountryCodeKn CountryCode = "KN"
	CountryCodeLc CountryCode = "LC"
	CountryCodeMf CountryCode = "MF"
	CountryCodePm CountryCode = "PM"
	CountryCodeVc CountryCode = "VC"
	CountryCodeWs CountryCode = "WS"
	CountryCodeSm CountryCode = "SM"
	CountryCodeSt CountryCode = "ST"
	CountryCodeSa CountryCode = "SA"
	CountryCodeSn CountryCode = "SN"
	CountryCodeRs CountryCode = "RS"
	CountryCodeSc CountryCode = "SC"
	CountryCodeSl CountryCode = "SL"
	CountryCodeSg CountryCode = "SG"
	CountryCodeSx CountryCode = "SX"
	CountryCodeSk CountryCode = "SK"
	CountryCodeSi CountryCode = "SI"
	CountryCodeSb CountryCode = "SB"
	CountryCodeSo CountryCode = "SO"
	CountryCodeZa CountryCode = "ZA"
	CountryCodeGs CountryCode = "GS"
	CountryCodeSs CountryCode = "SS"
	CountryCodeEs CountryCode = "ES"
	CountryCodeLk CountryCode = "LK"
	CountryCodeSd CountryCode = "SD"
	CountryCodeSr CountryCode = "SR"
	CountryCodeSj CountryCode = "SJ"
	CountryCodeSz CountryCode = "SZ"
	CountryCodeSe CountryCode = "SE"
	CountryCodeCh CountryCode = "CH"
	CountryCodeSy CountryCode = "SY"
	CountryCodeTw CountryCode = "TW"
	CountryCodeTj CountryCode = "TJ"
	CountryCodeTz CountryCode = "TZ"
	CountryCodeTh CountryCode = "TH"
	CountryCodeTl CountryCode = "TL"
	CountryCodeTg CountryCode = "TG"
	CountryCodeTk CountryCode = "TK"
	CountryCodeTo CountryCode = "TO"
	CountryCodeTt CountryCode = "TT"
	CountryCodeTn CountryCode = "TN"
	CountryCodeTr CountryCode = "TR"
	CountryCodeTm CountryCode = "TM"
	CountryCodeTc CountryCode = "TC"
	CountryCodeTv CountryCode = "TV"
	CountryCodeUg CountryCode = "UG"
	CountryCodeUa CountryCode = "UA"
	CountryCodeAe CountryCode = "AE"
	CountryCodeGb CountryCode = "GB"
	CountryCodeUs CountryCode = "US"
	CountryCodeUm CountryCode = "UM"
	CountryCodeUy CountryCode = "UY"
	CountryCodeUz CountryCode = "UZ"
	CountryCodeVu CountryCode = "VU"
	CountryCodeVe CountryCode = "VE"
	CountryCodeVn CountryCode = "VN"
	CountryCodeVg CountryCode = "VG"
	CountryCodeVi CountryCode = "VI"
	CountryCodeWf CountryCode = "WF"
	CountryCodeEh CountryCode = "EH"
	CountryCodeYe CountryCode = "YE"
	CountryCodeZm CountryCode = "ZM"
	CountryCodeZw CountryCode = "ZW"
)

// Values returns all known values for CountryCode. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (CountryCode) Values() []CountryCode {
	return []CountryCode{
		"AF",
		"AX",
		"AL",
		"DZ",
		"AS",
		"AD",
		"AO",
		"AI",
		"AQ",
		"AG",
		"AR",
		"AM",
		"AW",
		"AU",
		"AT",
		"AZ",
		"BS",
		"BH",
		"BD",
		"BB",
		"BY",
		"BE",
		"BZ",
		"BJ",
		"BM",
		"BT",
		"BO",
		"BQ",
		"BA",
		"BW",
		"BV",
		"BR",
		"IO",
		"BN",
		"BG",
		"BF",
		"BI",
		"KH",
		"CM",
		"CA",
		"CV",
		"KY",
		"CF",
		"TD",
		"CL",
		"CN",
		"CX",
		"CC",
		"CO",
		"KM",
		"CG",
		"CD",
		"CK",
		"CR",
		"CI",
		"HR",
		"CU",
		"CW",
		"CY",
		"CZ",
		"DK",
		"DJ",
		"DM",
		"DO",
		"EC",
		"EG",
		"SV",
		"GQ",
		"ER",
		"EE",
		"ET",
		"FK",
		"FO",
		"FJ",
		"FI",
		"FR",
		"GF",
		"PF",
		"TF",
		"GA",
		"GM",
		"GE",
		"DE",
		"GH",
		"GI",
		"GR",
		"GL",
		"GD",
		"GP",
		"GU",
		"GT",
		"GG",
		"GN",
		"GW",
		"GY",
		"HT",
		"HM",
		"VA",
		"HN",
		"HK",
		"HU",
		"IS",
		"IN",
		"ID",
		"IR",
		"IQ",
		"IE",
		"IM",
		"IL",
		"IT",
		"JM",
		"JP",
		"JE",
		"JO",
		"KZ",
		"KE",
		"KI",
		"KP",
		"KR",
		"KW",
		"KG",
		"LA",
		"LV",
		"LB",
		"LS",
		"LR",
		"LY",
		"LI",
		"LT",
		"LU",
		"MO",
		"MK",
		"MG",
		"MW",
		"MY",
		"MV",
		"ML",
		"MT",
		"MH",
		"MQ",
		"MR",
		"MU",
		"YT",
		"MX",
		"FM",
		"MD",
		"MC",
		"MN",
		"ME",
		"MS",
		"MA",
		"MZ",
		"MM",
		"NA",
		"NR",
		"NP",
		"NL",
		"NC",
		"NZ",
		"NI",
		"NE",
		"NG",
		"NU",
		"NF",
		"MP",
		"NO",
		"OM",
		"PK",
		"PW",
		"PS",
		"PA",
		"PG",
		"PY",
		"PE",
		"PH",
		"PN",
		"PL",
		"PT",
		"PR",
		"QA",
		"RE",
		"RO",
		"RU",
		"RW",
		"BL",
		"SH",
		"KN",
		"LC",
		"MF",
		"PM",
		"VC",
		"WS",
		"SM",
		"ST",
		"SA",
		"SN",
		"RS",
		"SC",
		"SL",
		"SG",
		"SX",
		"SK",
		"SI",
		"SB",
		"SO",
		"ZA",
		"GS",
		"SS",
		"ES",
		"LK",
		"SD",
		"SR",
		"SJ",
		"SZ",
		"SE",
		"CH",
		"SY",
		"TW",
		"TJ",
		"TZ",
		"TH",
		"TL",
		"TG",
		"TK",
		"TO",
		"TT",
		"TN",
		"TR",
		"TM",
		"TC",
		"TV",
		"UG",
		"UA",
		"AE",
		"GB",
		"US",
		"UM",
		"UY",
		"UZ",
		"VU",
		"VE",
		"VN",
		"VG",
		"VI",
		"WF",
		"EH",
		"YE",
		"ZM",
		"ZW",
	}
}

type FallbackBehavior string

// Enum values for FallbackBehavior
const (
	FallbackBehaviorMatch   FallbackBehavior = "MATCH"
	FallbackBehaviorNoMatch FallbackBehavior = "NO_MATCH"
)

// Values returns all known values for FallbackBehavior. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (FallbackBehavior) Values() []FallbackBehavior {
	return []FallbackBehavior{
		"MATCH",
		"NO_MATCH",
	}
}

type FilterBehavior string

// Enum values for FilterBehavior
const (
	FilterBehaviorKeep FilterBehavior = "KEEP"
	FilterBehaviorDrop FilterBehavior = "DROP"
)

// Values returns all known values for FilterBehavior. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (FilterBehavior) Values() []FilterBehavior {
	return []FilterBehavior{
		"KEEP",
		"DROP",
	}
}

type FilterRequirement string

// Enum values for FilterRequirement
const (
	FilterRequirementMeetsAll FilterRequirement = "MEETS_ALL"
	FilterRequirementMeetsAny FilterRequirement = "MEETS_ANY"
)

// Values returns all known values for FilterRequirement. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (FilterRequirement) Values() []FilterRequirement {
	return []FilterRequirement{
		"MEETS_ALL",
		"MEETS_ANY",
	}
}

type ForwardedIPPosition string

// Enum values for ForwardedIPPosition
const (
	ForwardedIPPositionFirst ForwardedIPPosition = "FIRST"
	ForwardedIPPositionLast  ForwardedIPPosition = "LAST"
	ForwardedIPPositionAny   ForwardedIPPosition = "ANY"
)

// Values returns all known values for ForwardedIPPosition. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ForwardedIPPosition) Values() []ForwardedIPPosition {
	return []ForwardedIPPosition{
		"FIRST",
		"LAST",
		"ANY",
	}
}

type IPAddressVersion string

// Enum values for IPAddressVersion
const (
	IPAddressVersionIpv4 IPAddressVersion = "IPV4"
	IPAddressVersionIpv6 IPAddressVersion = "IPV6"
)

// Values returns all known values for IPAddressVersion. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (IPAddressVersion) Values() []IPAddressVersion {
	return []IPAddressVersion{
		"IPV4",
		"IPV6",
	}
}

type JsonMatchScope string

// Enum values for JsonMatchScope
const (
	JsonMatchScopeAll   JsonMatchScope = "ALL"
	JsonMatchScopeKey   JsonMatchScope = "KEY"
	JsonMatchScopeValue JsonMatchScope = "VALUE"
)

// Values returns all known values for JsonMatchScope. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (JsonMatchScope) Values() []JsonMatchScope {
	return []JsonMatchScope{
		"ALL",
		"KEY",
		"VALUE",
	}
}

type LabelMatchScope string

// Enum values for LabelMatchScope
const (
	LabelMatchScopeLabel     LabelMatchScope = "LABEL"
	LabelMatchScopeNamespace LabelMatchScope = "NAMESPACE"
)

// Values returns all known values for LabelMatchScope. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (LabelMatchScope) Values() []LabelMatchScope {
	return []LabelMatchScope{
		"LABEL",
		"NAMESPACE",
	}
}

type ParameterExceptionField string

// Enum values for ParameterExceptionField
const (
	ParameterExceptionFieldWebAcl                         ParameterExceptionField = "WEB_ACL"
	ParameterExceptionFieldRuleGroup                      ParameterExceptionField = "RULE_GROUP"
	ParameterExceptionFieldRegexPatternSet                ParameterExceptionField = "REGEX_PATTERN_SET"
	ParameterExceptionFieldIpSet                          ParameterExceptionField = "IP_SET"
	ParameterExceptionFieldManagedRuleSet                 ParameterExceptionField = "MANAGED_RULE_SET"
	ParameterExceptionFieldRule                           ParameterExceptionField = "RULE"
	ParameterExceptionFieldExcludedRule                   ParameterExceptionField = "EXCLUDED_RULE"
	ParameterExceptionFieldStatement                      ParameterExceptionField = "STATEMENT"
	ParameterExceptionFieldByteMatchStatement             ParameterExceptionField = "BYTE_MATCH_STATEMENT"
	ParameterExceptionFieldSqliMatchStatement             ParameterExceptionField = "SQLI_MATCH_STATEMENT"
	ParameterExceptionFieldXssMatchStatement              ParameterExceptionField = "XSS_MATCH_STATEMENT"
	ParameterExceptionFieldSizeConstraintStatement        ParameterExceptionField = "SIZE_CONSTRAINT_STATEMENT"
	ParameterExceptionFieldGeoMatchStatement              ParameterExceptionField = "GEO_MATCH_STATEMENT"
	ParameterExceptionFieldRateBasedStatement             ParameterExceptionField = "RATE_BASED_STATEMENT"
	ParameterExceptionFieldRuleGroupReferenceStatement    ParameterExceptionField = "RULE_GROUP_REFERENCE_STATEMENT"
	ParameterExceptionFieldRegexPatternReferenceStatement ParameterExceptionField = "REGEX_PATTERN_REFERENCE_STATEMENT"
	ParameterExceptionFieldIpSetReferenceStatement        ParameterExceptionField = "IP_SET_REFERENCE_STATEMENT"
	ParameterExceptionFieldManagedRuleSetStatement        ParameterExceptionField = "MANAGED_RULE_SET_STATEMENT"
	ParameterExceptionFieldLabelMatchStatement            ParameterExceptionField = "LABEL_MATCH_STATEMENT"
	ParameterExceptionFieldAndStatement                   ParameterExceptionField = "AND_STATEMENT"
	ParameterExceptionFieldOrStatement                    ParameterExceptionField = "OR_STATEMENT"
	ParameterExceptionFieldNotStatement                   ParameterExceptionField = "NOT_STATEMENT"
	ParameterExceptionFieldIpAddress                      ParameterExceptionField = "IP_ADDRESS"
	ParameterExceptionFieldIpAddressVersion               ParameterExceptionField = "IP_ADDRESS_VERSION"
	ParameterExceptionFieldFieldToMatch                   ParameterExceptionField = "FIELD_TO_MATCH"
	ParameterExceptionFieldTextTransformation             ParameterExceptionField = "TEXT_TRANSFORMATION"
	ParameterExceptionFieldSingleQueryArgument            ParameterExceptionField = "SINGLE_QUERY_ARGUMENT"
	ParameterExceptionFieldSingleHeader                   ParameterExceptionField = "SINGLE_HEADER"
	ParameterExceptionFieldDefaultAction                  ParameterExceptionField = "DEFAULT_ACTION"
	ParameterExceptionFieldRuleAction                     ParameterExceptionField = "RULE_ACTION"
	ParameterExceptionFieldEntityLimit                    ParameterExceptionField = "ENTITY_LIMIT"
	ParameterExceptionFieldOverrideAction                 ParameterExceptionField = "OVERRIDE_ACTION"
	ParameterExceptionFieldScopeValue                     ParameterExceptionField = "SCOPE_VALUE"
	ParameterExceptionFieldResourceArn                    ParameterExceptionField = "RESOURCE_ARN"
	ParameterExceptionFieldResourceType                   ParameterExceptionField = "RESOURCE_TYPE"
	ParameterExceptionFieldTags                           ParameterExceptionField = "TAGS"
	ParameterExceptionFieldTagKeys                        ParameterExceptionField = "TAG_KEYS"
	ParameterExceptionFieldMetricName                     ParameterExceptionField = "METRIC_NAME"
	ParameterExceptionFieldFirewallManagerStatement       ParameterExceptionField = "FIREWALL_MANAGER_STATEMENT"
	ParameterExceptionFieldFallbackBehavior               ParameterExceptionField = "FALLBACK_BEHAVIOR"
	ParameterExceptionFieldPosition                       ParameterExceptionField = "POSITION"
	ParameterExceptionFieldForwardedIpConfig              ParameterExceptionField = "FORWARDED_IP_CONFIG"
	ParameterExceptionFieldIpSetForwardedIpConfig         ParameterExceptionField = "IP_SET_FORWARDED_IP_CONFIG"
	ParameterExceptionFieldHeaderName                     ParameterExceptionField = "HEADER_NAME"
	ParameterExceptionFieldCustomRequestHandling          ParameterExceptionField = "CUSTOM_REQUEST_HANDLING"
	ParameterExceptionFieldResponseContentType            ParameterExceptionField = "RESPONSE_CONTENT_TYPE"
	ParameterExceptionFieldCustomResponse                 ParameterExceptionField = "CUSTOM_RESPONSE"
	ParameterExceptionFieldCustomResponseBody             ParameterExceptionField = "CUSTOM_RESPONSE_BODY"
	ParameterExceptionFieldJsonMatchPattern               ParameterExceptionField = "JSON_MATCH_PATTERN"
	ParameterExceptionFieldJsonMatchScope                 ParameterExceptionField = "JSON_MATCH_SCOPE"
	ParameterExceptionFieldBodyParsingFallbackBehavior    ParameterExceptionField = "BODY_PARSING_FALLBACK_BEHAVIOR"
	ParameterExceptionFieldLoggingFilter                  ParameterExceptionField = "LOGGING_FILTER"
	ParameterExceptionFieldFilterCondition                ParameterExceptionField = "FILTER_CONDITION"
	ParameterExceptionFieldExpireTimestamp                ParameterExceptionField = "EXPIRE_TIMESTAMP"
	ParameterExceptionFieldChangePropagationStatus        ParameterExceptionField = "CHANGE_PROPAGATION_STATUS"
	ParameterExceptionFieldAssociableResource             ParameterExceptionField = "ASSOCIABLE_RESOURCE"
)

// Values returns all known values for ParameterExceptionField. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ParameterExceptionField) Values() []ParameterExceptionField {
	return []ParameterExceptionField{
		"WEB_ACL",
		"RULE_GROUP",
		"REGEX_PATTERN_SET",
		"IP_SET",
		"MANAGED_RULE_SET",
		"RULE",
		"EXCLUDED_RULE",
		"STATEMENT",
		"BYTE_MATCH_STATEMENT",
		"SQLI_MATCH_STATEMENT",
		"XSS_MATCH_STATEMENT",
		"SIZE_CONSTRAINT_STATEMENT",
		"GEO_MATCH_STATEMENT",
		"RATE_BASED_STATEMENT",
		"RULE_GROUP_REFERENCE_STATEMENT",
		"REGEX_PATTERN_REFERENCE_STATEMENT",
		"IP_SET_REFERENCE_STATEMENT",
		"MANAGED_RULE_SET_STATEMENT",
		"LABEL_MATCH_STATEMENT",
		"AND_STATEMENT",
		"OR_STATEMENT",
		"NOT_STATEMENT",
		"IP_ADDRESS",
		"IP_ADDRESS_VERSION",
		"FIELD_TO_MATCH",
		"TEXT_TRANSFORMATION",
		"SINGLE_QUERY_ARGUMENT",
		"SINGLE_HEADER",
		"DEFAULT_ACTION",
		"RULE_ACTION",
		"ENTITY_LIMIT",
		"OVERRIDE_ACTION",
		"SCOPE_VALUE",
		"RESOURCE_ARN",
		"RESOURCE_TYPE",
		"TAGS",
		"TAG_KEYS",
		"METRIC_NAME",
		"FIREWALL_MANAGER_STATEMENT",
		"FALLBACK_BEHAVIOR",
		"POSITION",
		"FORWARDED_IP_CONFIG",
		"IP_SET_FORWARDED_IP_CONFIG",
		"HEADER_NAME",
		"CUSTOM_REQUEST_HANDLING",
		"RESPONSE_CONTENT_TYPE",
		"CUSTOM_RESPONSE",
		"CUSTOM_RESPONSE_BODY",
		"JSON_MATCH_PATTERN",
		"JSON_MATCH_SCOPE",
		"BODY_PARSING_FALLBACK_BEHAVIOR",
		"LOGGING_FILTER",
		"FILTER_CONDITION",
		"EXPIRE_TIMESTAMP",
		"CHANGE_PROPAGATION_STATUS",
		"ASSOCIABLE_RESOURCE",
	}
}

type PositionalConstraint string

// Enum values for PositionalConstraint
const (
	PositionalConstraintExactly      PositionalConstraint = "EXACTLY"
	PositionalConstraintStartsWith   PositionalConstraint = "STARTS_WITH"
	PositionalConstraintEndsWith     PositionalConstraint = "ENDS_WITH"
	PositionalConstraintContains     PositionalConstraint = "CONTAINS"
	PositionalConstraintContainsWord PositionalConstraint = "CONTAINS_WORD"
)

// Values returns all known values for PositionalConstraint. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (PositionalConstraint) Values() []PositionalConstraint {
	return []PositionalConstraint{
		"EXACTLY",
		"STARTS_WITH",
		"ENDS_WITH",
		"CONTAINS",
		"CONTAINS_WORD",
	}
}

type RateBasedStatementAggregateKeyType string

// Enum values for RateBasedStatementAggregateKeyType
const (
	RateBasedStatementAggregateKeyTypeIp          RateBasedStatementAggregateKeyType = "IP"
	RateBasedStatementAggregateKeyTypeForwardedIp RateBasedStatementAggregateKeyType = "FORWARDED_IP"
)

// Values returns all known values for RateBasedStatementAggregateKeyType. Note
// that this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (RateBasedStatementAggregateKeyType) Values() []RateBasedStatementAggregateKeyType {
	return []RateBasedStatementAggregateKeyType{
		"IP",
		"FORWARDED_IP",
	}
}

type ResourceType string

// Enum values for ResourceType
const (
	ResourceTypeApplicationLoadBalancer ResourceType = "APPLICATION_LOAD_BALANCER"
	ResourceTypeApiGateway              ResourceType = "API_GATEWAY"
	ResourceTypeAppsync                 ResourceType = "APPSYNC"
)

// Values returns all known values for ResourceType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (ResourceType) Values() []ResourceType {
	return []ResourceType{
		"APPLICATION_LOAD_BALANCER",
		"API_GATEWAY",
		"APPSYNC",
	}
}

type ResponseContentType string

// Enum values for ResponseContentType
const (
	ResponseContentTypeTextPlain       ResponseContentType = "TEXT_PLAIN"
	ResponseContentTypeTextHtml        ResponseContentType = "TEXT_HTML"
	ResponseContentTypeApplicationJson ResponseContentType = "APPLICATION_JSON"
)

// Values returns all known values for ResponseContentType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ResponseContentType) Values() []ResponseContentType {
	return []ResponseContentType{
		"TEXT_PLAIN",
		"TEXT_HTML",
		"APPLICATION_JSON",
	}
}

type Scope string

// Enum values for Scope
const (
	ScopeCloudfront Scope = "CLOUDFRONT"
	ScopeRegional   Scope = "REGIONAL"
)

// Values returns all known values for Scope. Note that this can be expanded in the
// future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (Scope) Values() []Scope {
	return []Scope{
		"CLOUDFRONT",
		"REGIONAL",
	}
}

type TextTransformationType string

// Enum values for TextTransformationType
const (
	TextTransformationTypeNone               TextTransformationType = "NONE"
	TextTransformationTypeCompressWhiteSpace TextTransformationType = "COMPRESS_WHITE_SPACE"
	TextTransformationTypeHtmlEntityDecode   TextTransformationType = "HTML_ENTITY_DECODE"
	TextTransformationTypeLowercase          TextTransformationType = "LOWERCASE"
	TextTransformationTypeCmdLine            TextTransformationType = "CMD_LINE"
	TextTransformationTypeUrlDecode          TextTransformationType = "URL_DECODE"
	TextTransformationTypeBase64Decode       TextTransformationType = "BASE64_DECODE"
	TextTransformationTypeHexDecode          TextTransformationType = "HEX_DECODE"
	TextTransformationTypeMd5                TextTransformationType = "MD5"
	TextTransformationTypeReplaceComments    TextTransformationType = "REPLACE_COMMENTS"
	TextTransformationTypeEscapeSeqDecode    TextTransformationType = "ESCAPE_SEQ_DECODE"
	TextTransformationTypeSqlHexDecode       TextTransformationType = "SQL_HEX_DECODE"
	TextTransformationTypeCssDecode          TextTransformationType = "CSS_DECODE"
	TextTransformationTypeJsDecode           TextTransformationType = "JS_DECODE"
	TextTransformationTypeNormalizePath      TextTransformationType = "NORMALIZE_PATH"
	TextTransformationTypeNormalizePathWin   TextTransformationType = "NORMALIZE_PATH_WIN"
	TextTransformationTypeRemoveNulls        TextTransformationType = "REMOVE_NULLS"
	TextTransformationTypeReplaceNulls       TextTransformationType = "REPLACE_NULLS"
	TextTransformationTypeBase64DecodeExt    TextTransformationType = "BASE64_DECODE_EXT"
	TextTransformationTypeUrlDecodeUni       TextTransformationType = "URL_DECODE_UNI"
	TextTransformationTypeUtf8ToUnicode      TextTransformationType = "UTF8_TO_UNICODE"
)

// Values returns all known values for TextTransformationType. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (TextTransformationType) Values() []TextTransformationType {
	return []TextTransformationType{
		"NONE",
		"COMPRESS_WHITE_SPACE",
		"HTML_ENTITY_DECODE",
		"LOWERCASE",
		"CMD_LINE",
		"URL_DECODE",
		"BASE64_DECODE",
		"HEX_DECODE",
		"MD5",
		"REPLACE_COMMENTS",
		"ESCAPE_SEQ_DECODE",
		"SQL_HEX_DECODE",
		"CSS_DECODE",
		"JS_DECODE",
		"NORMALIZE_PATH",
		"NORMALIZE_PATH_WIN",
		"REMOVE_NULLS",
		"REPLACE_NULLS",
		"BASE64_DECODE_EXT",
		"URL_DECODE_UNI",
		"UTF8_TO_UNICODE",
	}
}
