package sip2

import "regexp"

// msgDef definies the required and optional fields of a message.
type msgDef struct {
	// Required fixed-length fields, identified by position in header.
	RequiredFixed []fieldType
	// Required variable-length fields, identified by two-character ID.
	// (Some might acually be fixed length, but we can parse them with same logic.)
	RequiredVar []fieldType
	// Optional variable-length fields, identified by two-character ID.
	OptionalVar []fieldType
}

var (
	msgDefinitions = map[msgType]msgDef{
		// Requests:
		MsgReqPatronStatus: msgDef{
			RequiredFixed: []fieldType{
				FieldLanguage,
				FieldTransactionDate,
				FieldInstitutionID,
			},
			RequiredVar: []fieldType{
				FieldPatronIdentifier,
				FieldTerminalPassword,
				FieldPatronPassword,
			},
			OptionalVar: []fieldType{},
		},
		MsgReqCheckout: msgDef{
			RequiredFixed: []fieldType{
				FieldRenewalPolicy,
				FieldNoBlock,
				FieldTransactionDate,
				FieldNbDueDate,
			},
			RequiredVar: []fieldType{
				FieldInstitutionID,
				FieldPatronIdentifier,
				FieldItemIdentifier,
				FieldTerminalPassword,
			},
			OptionalVar: []fieldType{
				FieldItemProperties,
				FieldPatronPassword,
				FieldFeeAcknowledged,
				FieldCancel,
			},
		},
		MsgReqCheckin: msgDef{
			RequiredFixed: []fieldType{
				FieldNoBlock,
				FieldTransactionDate,
				FieldReturnDate,
			},
			RequiredVar: []fieldType{
				FieldCurrentLocation,
				FieldInstitutionID,
				FieldItemIdentifier,
				FieldTerminalPassword,
			},
			OptionalVar: []fieldType{
				FieldItemProperties,
				FieldCancel,
			},
		},
		MsgReqBlockPatron: msgDef{
			RequiredFixed: []fieldType{
				FieldCardRetained,
				FieldTransactionDate,
			},
			RequiredVar: []fieldType{
				FieldInstitutionID,
				FieldBlockedCardMsg,
				FieldPatronIdentifier,
				FieldTerminalPassword,
			},
			OptionalVar: []fieldType{
				FieldCardRetained,
				FieldTransactionDate,
				FieldInstitutionID,
				FieldBlockedCardMsg,
				FieldPatronIdentifier,
				FieldTerminalPassword,
			},
		},
		MsgReqStatus: msgDef{
			RequiredFixed: []fieldType{
				FieldStatusCode,
				FieldMaxPrintWidth,
				FieldProtocolVersion,
			},
			RequiredVar: []fieldType{},
			OptionalVar: []fieldType{},
		},
		MsgReqResend: msgDef{},
		MsgReqLogin: msgDef{
			RequiredFixed: []fieldType{
				FieldUIDAlgorithm,
				FieldPWDAlgorithm,
			},
			RequiredVar: []fieldType{
				FieldLoginUserID,
				FieldLoginPassword,
			},
			OptionalVar: []fieldType{
				FieldLocationCode,
			},
		},
		MsgReqPatronInformation: msgDef{
			RequiredFixed: []fieldType{
				FieldLanguage,
				FieldTransactionDate,
				FieldSummary,
			},
			RequiredVar: []fieldType{
				FieldInstitutionID,
				FieldPatronIdentifier,
			},
			OptionalVar: []fieldType{
				FieldTerminalPassword,
				FieldPatronPassword,
				FieldStartItem,
				FieldEndItem,
			},
		},
		MsgReqEndPatronSession: msgDef{
			RequiredFixed: []fieldType{
				FieldTransactionDate,
			},
			RequiredVar: []fieldType{
				FieldInstitutionID,
				FieldPatronIdentifier,
			},
			OptionalVar: []fieldType{},
		},
		MsgReqFeePaid: msgDef{
			RequiredFixed: []fieldType{
				FieldTransactionDate,
				FieldFeeType,
				FieldPaymentType,
				FieldCurrenyType,
				FieldFeeAmount,
			},
			RequiredVar: []fieldType{
				FieldInstitutionID,
				FieldPatronIdentifier,
			},
			OptionalVar: []fieldType{
				FieldTerminalPassword,
				FieldPatronPassword,
				FieldFeeIdentifier,
				FieldTransactionID,
			},
		},
		MsgReqItemInformation: msgDef{
			RequiredFixed: []fieldType{
				FieldTransactionDate,
			},
			RequiredVar: []fieldType{
				FieldInstitutionID,
				FieldItemIdentifier,
			},
			OptionalVar: []fieldType{
				FieldTerminalPassword,
			},
		},
		MsgReqItemStatusUpdate: msgDef{
			RequiredFixed: []fieldType{
				FieldTransactionDate,
			},
			RequiredVar: []fieldType{
				FieldInstitutionID,
				FieldItemIdentifier,
			},
			OptionalVar: []fieldType{
				FieldTerminalPassword,
				FieldItemProperties,
			},
		},
		MsgReqPatronEnable: msgDef{
			RequiredFixed: []fieldType{
				FieldTransactionDate,
			},
			RequiredVar: []fieldType{
				FieldInstitutionID,
				FieldPatronIdentifier,
			},
			OptionalVar: []fieldType{
				FieldTerminalPassword,
				FieldPatronPassword,
			},
		},
		MsgReqHold: msgDef{
			RequiredFixed: []fieldType{
				FieldHoldMode,
				FieldTransactionDate,
			},
			RequiredVar: []fieldType{
				FieldInstitutionID,
				FieldPatronIdentifier,
			},
			OptionalVar: []fieldType{
				FieldPickupLocation,
				FieldExpirationDate,
				FieldHoldType,
				FieldPatronPassword,
				FieldItemIdentifier,
				FieldTitleIdentifier,
				FieldTerminalPassword,
				FieldFeeAcknowledged,
			},
		},
		MsgReqRenew: msgDef{
			RequiredFixed: []fieldType{
				FieldThirdPartyAllowd,
				FieldNoBlock,
				FieldTransactionDate,
				FieldNbDueDate,
			},
			RequiredVar: []fieldType{
				FieldInstitutionID,
				FieldPatronIdentifier,
			},
			OptionalVar: []fieldType{
				FieldPatronPassword,
				FieldItemIdentifier,
				FieldTitleIdentifier,
				FieldTerminalPassword,
				FieldItemProperties,
				FieldFeeAcknowledged,
			},
		},
		MsgReqRenewAll: msgDef{
			RequiredFixed: []fieldType{
				FieldTransactionDate,
			},
			RequiredVar: []fieldType{
				FieldInstitutionID,
				FieldPatronIdentifier,
			},
			OptionalVar: []fieldType{
				FieldPatronPassword,
				FieldTerminalPassword,
				FieldFeeAcknowledged,
			},
		},
		// Responses:
		MsgRespPatronStatus: msgDef{
			RequiredFixed: []fieldType{
				FieldPatronStatus,
				FieldLanguage,
				FieldTransactionDate,
			},
			RequiredVar: []fieldType{
				FieldInstitutionID,
				FieldPatronIdentifier,
				FieldPersonalName,
			},
			OptionalVar: []fieldType{
				FieldValidPatron,
				FieldValidPatronPassword,
				FieldCurrenyType,
				FieldFeeAmount,
				FieldScreenMessage,
				FieldPrintLine,
				FieldLibraryName,
				FieldTerminalLocation,
			},
		},
		MsgRespCheckout: msgDef{
			RequiredFixed: []fieldType{
				FieldOK,
				FieldRenewalOK,
				FieldMagneticMedia,
				FieldDesentisize,
				FieldTransactionDate,
			},
			RequiredVar: []fieldType{
				FieldInstitutionID,
				FieldPatronIdentifier,
				FieldItemIdentifier,
				FieldTitleIdentifier,
				FieldDueDate,
			},
			OptionalVar: []fieldType{
				FieldFeeType,
				FieldSecurityInhibit,
				FieldCurrenyType,
				FieldFeeAmount,
				FieldMediaType,
				FieldItemProperties,
				FieldTransactionID,
				FieldScreenMessage,
				FieldPrintLine,
			},
		},
		MsgRespCheckin: msgDef{
			RequiredFixed: []fieldType{
				FieldOK,
				FieldResentisize,
				FieldMagneticMedia,
				FieldAlert,
				FieldTransactionDate,
			},
			RequiredVar: []fieldType{
				FieldInstitutionID,
				FieldItemIdentifier,
				FieldPermanentLocation,
			},
			OptionalVar: []fieldType{
				FieldTitleIdentifier,
				FieldSortBin,
				FieldPatronIdentifier,
				FieldMediaType,
				FieldItemProperties,
				FieldScreenMessage,
				FieldPrintLine,
			},
		},
		MsgRespStatus: msgDef{
			RequiredFixed: []fieldType{
				FieldOnLineStatus,
				FieldCheckinOK,
				FieldCheckoutOK,
				FieldRenewalPolicy,
				FieldStatusUpdateOK,
				FieldOffLineOK,
				FieldTimeoutPeriod,
				FieldRetriesAllowd,
				FieldDateTimeSync,
				FieldProtocolVersion,
			},
			RequiredVar: []fieldType{
				FieldInstitutionID,
				FieldSupportedMessages,
			},
			OptionalVar: []fieldType{},
		},
		MsgRespLogin: msgDef{
			RequiredFixed: []fieldType{
				FieldOK,
			},
			RequiredVar: []fieldType{},
			OptionalVar: []fieldType{},
		},
		MsgRespPatronInformation: msgDef{
			RequiredFixed: []fieldType{
				FieldPatronStatus,
				FieldLanguage,
				FieldTransactionDate,
				FieldHoldItemsCount,
				FieldOverdueItemsCount,
				FieldChargedItemsCount,
				FieldFineItemsCount,
				FieldRecallItemsCount,
				FieldUnavailableHoldsCount,
			},
			RequiredVar: []fieldType{
				FieldInstitutionID,
				FieldPatronIdentifier,
				FieldPersonalName,
			},
			OptionalVar: []fieldType{
				FieldHoldItemsLimit,
				FieldOverdueItemsLimit,
				FieldChargedItemsLimit,
				FieldValidPatron,
				FieldValidPatronPassword,
				FieldCurrenyType,
				FieldFeeAmount,
				FieldFeeLimit,
				FieldHoldItems,
				FieldOverdueItems,
				FieldChargedItems,
				FieldFineItems,
				FieldRecallItems,
				FieldUnavailableHoldsItems,
				FieldHomeAddress,
				FieldEmailAddress,
				FieldHomePhoneNumber,
				FieldScreenMessage,
				FieldPrintLine,
			},
		},
		MsgRespEndPatronSession: msgDef{
			RequiredFixed: []fieldType{
				FieldEndSession,
				FieldTransactionDate,
			},
			RequiredVar: []fieldType{
				FieldInstitutionID,
				FieldPatronIdentifier,
			},
			OptionalVar: []fieldType{
				FieldScreenMessage,
				FieldPrintLine,
			},
		},
		MsgRespFeePaid: msgDef{
			RequiredFixed: []fieldType{
				FieldPaymentAccepted,
				FieldTransactionDate,
			},
			RequiredVar: []fieldType{
				FieldInstitutionID,
				FieldPatronIdentifier,
			},
			OptionalVar: []fieldType{
				FieldTransactionID,
				FieldScreenMessage,
				FieldPrintLine,
			},
		},
		MsgRespItemInformation: msgDef{
			RequiredFixed: []fieldType{
				FieldCirulationStatus,
				FieldSecurityMarker,
				FieldFeeType,
				FieldTransactionDate,
			},
			RequiredVar: []fieldType{},
			OptionalVar: []fieldType{
				FieldHoldQueueLength,
				FieldDueDate,
				FieldRecallDate,
				FieldHoldPickupDate,
				FieldItemIdentifier,
				FieldTitleIdentifier,
				FieldOwner,
				FieldCurrenyType,
				FieldFeeAmount,
				FieldMediaType,
				FieldPermanentLocation,
				FieldCurrentLocation,
				FieldItemProperties,
				FieldScreenMessage,
				FieldPrintLine,
			},
		},
		MsgRespItemStatusUpdate: msgDef{
			RequiredFixed: []fieldType{
				FieldItemPropertiesOK,
				FieldTransactionDate,
			},
			RequiredVar: []fieldType{
				FieldPatronIdentifier,
			},
			OptionalVar: []fieldType{
				FieldTitleIdentifier,
				FieldItemProperties,
				FieldScreenMessage,
				FieldPrintLine,
			},
		},
		MsgRespPatronEnable: msgDef{
			RequiredFixed: []fieldType{
				FieldPatronStatus,
				FieldLanguage,
				FieldTransactionDate,
			},
			RequiredVar: []fieldType{
				FieldInstitutionID,
				FieldPatronIdentifier,
				FieldPersonalName,
			},
			OptionalVar: []fieldType{
				FieldValidPatron,
				FieldValidPatronPassword,
				FieldScreenMessage,
				FieldPrintLine,
			},
		},
		MsgRespHold: msgDef{
			RequiredFixed: []fieldType{
				FieldOK,
				FieldAvialable,
				FieldTransactionDate,
			},
			RequiredVar: []fieldType{
				FieldInstitutionID,
				FieldPatronIdentifier,
			},
			OptionalVar: []fieldType{
				FieldExpirationDate,
				FieldQueuePosition,
				FieldPickupLocation,
				FieldItemIdentifier,
				FieldTitleIdentifier,
				FieldScreenMessage,
				FieldPrintLine,
			},
		},
		MsgRespRenew: msgDef{
			RequiredFixed: []fieldType{
				FieldOK,
				FieldRenewalOK,
				FieldMagneticMedia,
				FieldDesentisize,
				FieldTransactionDate,
			},
			RequiredVar: []fieldType{
				FieldInstitutionID,
				FieldPatronIdentifier,
				FieldItemIdentifier,
				FieldTitleIdentifier,
				FieldDueDate,
			},
			OptionalVar: []fieldType{
				FieldFeeType,
				FieldSecurityInhibit,
				FieldFeeAmount,
				FieldMediaType,
				FieldItemProperties,
				FieldTransactionID,
				FieldScreenMessage,
				FieldPrintLine,
			},
		},
		MsgRespRenewAll: msgDef{
			RequiredFixed: []fieldType{
				FieldOK,
				FieldRenewedCount,
				FieldUnrenewedCount,
				FieldTransactionDate,
			},
			RequiredVar: []fieldType{
				FieldInstitutionID,
			},
			OptionalVar: []fieldType{
				FieldRenewedItems,
				FieldUnrenewedItems,
				FieldScreenMessage,
				FieldPrintLine,
			},
		},
	}

	fixedFieldLengths = map[fieldType]int{
		FieldAlert:                 1,
		FieldAvialable:             1,
		FieldCardRetained:          1,
		FieldChargedItemsCount:     4,
		FieldFineItemsCount:        4,
		FieldFeeType:               2,
		FieldHoldItemsCount:        4,
		FieldOverdueItemsCount:     4,
		FieldRecallItemsCount:      4,
		FieldUnavailableHoldsCount: 4,
		FieldCheckinOK:             1,
		FieldCheckoutOK:            1,
		FieldCirulationStatus:      2,
		FieldDateTimeSync:          18,
		FieldDesentisize:           1,
		FieldEndSession:            1,
		FieldHoldMode:              1,
		FieldItemPropertiesOK:      1,
		FieldLanguage:              3,
		FieldMagneticMedia:         1,
		FieldMaxPrintWidth:         3,
		FieldNbDueDate:             18,
		FieldNoBlock:               1,
		FieldOffLineOK:             1,
		FieldOK:                    1,
		FieldOnLineStatus:          1,
		FieldPatronStatus:          14,
		FieldPaymentAccepted:       1,
		FieldPaymentType:           2,
		FieldProtocolVersion:       4,
		FieldPWDAlgorithm:          1,
		FieldRenewalOK:             1,
		FieldRenewedCount:          4,
		FieldResentisize:           1,
		FieldRetriesAllowd:         3,
		FieldReturnDate:            18,
		FieldRenewalPolicy:         1,
		FieldSecurityInhibit:       1,
		FieldSecurityMarker:        2,
		FieldStatusCode:            1,
		FieldStatusUpdateOK:        1,
		FieldSummary:               10,
		FieldThirdPartyAllowd:      1,
		FieldTimeoutPeriod:         3,
		FieldTransactionDate:       18,
		FieldUIDAlgorithm:          1,
		FieldUnrenewedCount:        4,
	}

	rxpYesNo         = regexp.MustCompile(`Y|N`)
	rxpDigits        = regexp.MustCompile(`\d+$`)
	rxpDigitsOrBlank = regexp.MustCompile(`[\d\s]+$`)
	rxpTimestamp     = regexp.MustCompile(`\d{8}[\sZ\d]{4}\d{6}`) // TODO verify with ANSI standard X3.30 for date and X3.43 for time.

	// fieldValidation defines the allowed character patterns for fields. If a fieldType
	// is not present in this map, any string is allowed.
	// The string length is allways guaranted to be correct by the parser, so it is not included in the regexes.
	fieldValdiation = map[fieldType]*regexp.Regexp{
		FieldAlert:                 rxpYesNo,
		FieldAvialable:             rxpYesNo,
		FieldCardRetained:          rxpYesNo,
		FieldChargedItemsCount:     rxpDigitsOrBlank, // {4}
		FieldFineItemsCount:        rxpDigitsOrBlank, // {4}
		FieldHoldItemsCount:        rxpDigitsOrBlank, // {4}
		FieldOverdueItemsCount:     rxpDigitsOrBlank, // {4}
		FieldRecallItemsCount:      rxpDigitsOrBlank, // {4}
		FieldUnavailableHoldsCount: rxpDigitsOrBlank, // {4}
		FieldCheckinOK:             rxpYesNo,
		FieldCheckoutOK:            rxpYesNo,
		FieldCirulationStatus:      rxpDigits, // {2}
		FieldDateTimeSync:          rxpTimestamp,
		FieldDesentisize:           regexp.MustCompile(`Y|N|U`),
		FieldEndSession:            rxpYesNo,
		FieldHoldMode:              regexp.MustCompile(`\+|\-|\*`),
		FieldLanguage:              rxpDigits, // {3}
		FieldMagneticMedia:         regexp.MustCompile(`Y|N|U`),
		FieldMaxPrintWidth:         rxpDigits, // {3}
		FieldNbDueDate:             rxpTimestamp,
		FieldNoBlock:               rxpYesNo,
		FieldOffLineOK:             rxpYesNo,
		FieldOK:                    regexp.MustCompile(`0|1`),
		FieldOnLineStatus:          rxpYesNo,
		FieldPatronStatus:          regexp.MustCompile(`[\s|Y]{14}`),
		FieldPaymentAccepted:       rxpYesNo,
		FieldPaymentType:           rxpDigits, // {2}
		FieldProtocolVersion:       regexp.MustCompile(`\d\.\d{20`),
		FieldRenewalOK:             rxpYesNo,
		FieldRenewedCount:          rxpDigits, // {4}
		FieldResentisize:           rxpYesNo,
		FieldRetriesAllowd:         rxpDigits, // {3}
		FieldReturnDate:            rxpTimestamp,
		FieldRenewalPolicy:         rxpYesNo,
		FieldSecurityMarker:        rxpDigits, // {2}
		FieldStatusCode:            regexp.MustCompile(`0|1|2`),
		FieldStatusUpdateOK:        rxpYesNo,
		FieldSummary:               regexp.MustCompile(`[\sY]{10}`),
		FieldThirdPartyAllowd:      rxpYesNo,
		FieldTimeoutPeriod:         rxpDigits, // {3}
		FieldTransactionDate:       rxpTimestamp,
		FieldUnrenewedCount:        rxpDigits, // {4}
	}

	// minMsgLength defines the minimum length needed for a Message to be able to contain
	// all the required fields. Variable-length fields are given a count of 3: a 2-character
	// code plus delimiter.
	minMsgLength = map[msgType]int{
		MsgReqPatronStatus:       35, // 2+3+18+3+3+3+3
		MsgReqCheckout:           52, // 2+1+1+18+18+3+3+3+3
		MsgReqCheckin:            51, // 2+1+18+18+3+3+3+3
		MsgReqBlockPatron:        33, // 2+1+18+3+3+3+3
		MsgReqStatus:             10, // 2+1+3+4
		MsgReqResend:             2,  // 2
		MsgReqLogin:              10, // 2+1+1+3+3
		MsgReqPatronInformation:  45, // 2+3+18+10+3+3+3+3
		MsgReqEndPatronSession:   32, // 2+18+3+3+3+3
		MsgReqFeePaid:            36, // 2+18+2+2+3+3+3+3
		MsgReqItemInformation:    26, // 2+18+3+3
		MsgReqItemStatusUpdate:   29, // 2+18+3+3+3
		MsgReqPatronEnable:       26, // 2+18+3+3
		MsgReqHold:               45, // 2+1+18+18+3+3
		MsgReqRenew:              46, // 2+1+1+18+18+3+3
		MsgReqRenewAll:           26, // 2+18+3+3
		MsgRespPatronStatus:      46, // 2+14+3+18+3+3+3
		MsgRespCheckout:          39, // 2+1+1+1+1+18+3+3+3+3+3
		MsgRespCheckin:           33, // 2+1+1+1+1+18+3+3+3
		MsgRespStatus:            42, // 2+1+1+1+1+1+1+3+3+18+4+3+3
		MsgRespLogin:             3,  // 2+1
		MsgRespPatronInformation: 70, // 2+14+3+18+4+4+4+4+4+4+3+3+3
		MsgRespEndPatronSession:  27, // 2+1+18+3+3
		MsgRespFeePaid:           27, // 2+1+18+3+3
		MsgRespItemInformation:   32, // 2+2+2+2+18+3+3
		MsgRespItemStatusUpdate:  24, // 2+1+18+3
		MsgRespPatronEnable:      46, // 2+14+3+18+3+3+3
		MsgRespHold:              28, // 2+1+1+18+3+3
		MsgRespRenew:             39, // 2+1+1+1+1+18+3+3+3+3+3
		MsgRespRenewAll:          32, // 2+1+4+4+18+3
	}

	repeatableField = map[fieldType]bool{
		FieldScreenMessage:         true,
		FieldPrintLine:             true,
		FieldRenewedItems:          true,
		FieldUnrenewedItems:        true,
		FieldHoldItems:             true,
		FieldOverdueItems:          true,
		FieldChargedItems:          true,
		FieldFineItems:             true,
		FieldRecallItems:           true,
		FieldUnavailableHoldsItems: true,
	}

	msgToCode = map[msgType]string{
		MsgReqPatronStatus:       "23",
		MsgReqCheckout:           "11",
		MsgReqCheckin:            "09",
		MsgReqBlockPatron:        "01",
		MsgReqStatus:             "99",
		MsgReqResend:             "97",
		MsgReqLogin:              "93",
		MsgReqPatronInformation:  "63",
		MsgReqEndPatronSession:   "35",
		MsgReqFeePaid:            "37",
		MsgReqItemInformation:    "17",
		MsgReqItemStatusUpdate:   "19",
		MsgReqPatronEnable:       "25",
		MsgReqHold:               "15",
		MsgReqRenew:              "29",
		MsgReqRenewAll:           "65",
		MsgRespPatronStatus:      "24",
		MsgRespCheckout:          "12",
		MsgRespCheckin:           "10",
		MsgRespStatus:            "98",
		MsgRespLogin:             "94",
		MsgRespPatronInformation: "64",
		MsgRespEndPatronSession:  "36",
		MsgRespFeePaid:           "38",
		MsgRespItemInformation:   "18",
		MsgRespItemStatusUpdate:  "20",
		MsgRespPatronEnable:      "26",
		MsgRespHold:              "16",
		MsgRespRenew:             "30",
		MsgRespRenewAll:          "66",
	}

	codeToMsg = map[string]msgType{
		"23": MsgReqPatronStatus,
		"11": MsgReqCheckout,
		"09": MsgReqCheckin,
		"01": MsgReqBlockPatron,
		"99": MsgReqStatus,
		"97": MsgReqResend,
		"93": MsgReqLogin,
		"63": MsgReqPatronInformation,
		"35": MsgReqEndPatronSession,
		"37": MsgReqFeePaid,
		"17": MsgReqItemInformation,
		"19": MsgReqItemStatusUpdate,
		"25": MsgReqPatronEnable,
		"15": MsgReqHold,
		"29": MsgReqRenew,
		"65": MsgReqRenewAll,
		"24": MsgRespPatronStatus,
		"12": MsgRespCheckout,
		"10": MsgRespCheckin,
		"98": MsgRespStatus,
		"94": MsgRespLogin,
		"64": MsgRespPatronInformation,
		"36": MsgRespEndPatronSession,
		"38": MsgRespFeePaid,
		"18": MsgRespItemInformation,
		"20": MsgRespItemStatusUpdate,
		"26": MsgRespPatronEnable,
		"16": MsgRespHold,
		"30": MsgRespRenew,
		"66": MsgRespRenewAll,
	}

	fieldToCode = map[fieldType]string{
		FieldPatronIdentifier:      "AA",
		FieldItemIdentifier:        "AB",
		FieldTerminalPassword:      "AC",
		FieldPatronPassword:        "AD",
		FieldPersonalName:          "AE",
		FieldScreenMessage:         "AF",
		FieldPrintLine:             "AG",
		FieldDueDate:               "AH",
		FieldTitleIdentifier:       "AJ",
		FieldBlockedCardMsg:        "AL",
		FieldLibraryName:           "AM",
		FieldTerminalLocation:      "AN",
		FieldInstitutionID:         "AO",
		FieldCurrentLocation:       "AP",
		FieldPermanentLocation:     "AQ",
		FieldHoldItems:             "AS",
		FieldOverdueItems:          "AT",
		FieldChargedItems:          "AU",
		FieldFineItems:             "AV",
		FieldHomeAddress:           "BD",
		FieldEmailAddress:          "BE",
		FieldHomePhoneNumber:       "BF",
		FieldOwner:                 "BG",
		FieldCurrenyType:           "BH",
		FieldCancel:                "BI",
		FieldTransactionID:         "BK",
		FieldValidPatron:           "BL",
		FieldRenewedItems:          "BM",
		FieldUnrenewedItems:        "BN",
		FieldFeeAcknowledged:       "BO",
		FieldStartItem:             "BP",
		FieldEndItem:               "BQ",
		FieldQueuePosition:         "BR",
		FieldPickupLocation:        "BS",
		FieldFeeType:               "BT",
		FieldRecallItems:           "BU",
		FieldFeeAmount:             "BV",
		FieldExpirationDate:        "BW",
		FieldSupportedMessages:     "BX",
		FieldHoldType:              "BY",
		FieldHoldItemsLimit:        "BZ",
		FieldOverdueItemsLimit:     "CA",
		FieldChargedItemsLimit:     "CB",
		FieldFeeLimit:              "CC",
		FieldUnavailableHoldsItems: "CD",
		FieldHoldQueueLength:       "CF",
		FieldFeeIdentifier:         "CG",
		FieldItemProperties:        "CH",
		FieldSecurityInhibit:       "CI",
		FieldRecallDate:            "CJ",
		FieldMediaType:             "CK",
		FieldSortBin:               "CL",
		FieldHoldPickupDate:        "CM",
		FieldLoginUserID:           "CN",
		FieldLoginPassword:         "CO",
		FieldLocationCode:          "CP",
		FieldValidPatronPassword:   "CQ",
		FieldSequenceNumber:        "AY",
		FieldChecksum:              "AZ",
	}

	codeToField = map[string]fieldType{
		"AA": FieldPatronIdentifier,
		"AB": FieldItemIdentifier,
		"AC": FieldTerminalPassword,
		"AD": FieldPatronPassword,
		"AE": FieldPersonalName,
		"AF": FieldScreenMessage,
		"AG": FieldPrintLine,
		"AH": FieldDueDate,
		"AJ": FieldTitleIdentifier,
		"AL": FieldBlockedCardMsg,
		"AM": FieldLibraryName,
		"AN": FieldTerminalLocation,
		"AO": FieldInstitutionID,
		"AP": FieldCurrentLocation,
		"AQ": FieldPermanentLocation,
		"AS": FieldHoldItems,
		"AT": FieldOverdueItems,
		"AU": FieldChargedItems,
		"AV": FieldFineItems,
		"BD": FieldHomeAddress,
		"BE": FieldEmailAddress,
		"BF": FieldHomePhoneNumber,
		"BG": FieldOwner,
		"BH": FieldCurrenyType,
		"BI": FieldCancel,
		"BK": FieldTransactionID,
		"BL": FieldValidPatron,
		"BM": FieldRenewedItems,
		"BN": FieldUnrenewedItems,
		"BO": FieldFeeAcknowledged,
		"BP": FieldStartItem,
		"BQ": FieldEndItem,
		"BR": FieldQueuePosition,
		"BS": FieldPickupLocation,
		"BT": FieldFeeType,
		"BU": FieldRecallItems,
		"BV": FieldFeeAmount,
		"BW": FieldExpirationDate,
		"BX": FieldSupportedMessages,
		"BY": FieldHoldType,
		"BZ": FieldHoldItemsLimit,
		"CA": FieldOverdueItemsLimit,
		"CB": FieldChargedItemsLimit,
		"CC": FieldFeeLimit,
		"CD": FieldUnavailableHoldsItems,
		"CF": FieldHoldQueueLength,
		"CG": FieldFeeIdentifier,
		"CH": FieldItemProperties,
		"CI": FieldSecurityInhibit,
		"CJ": FieldRecallDate,
		"CK": FieldMediaType,
		"CL": FieldSortBin,
		"CM": FieldHoldPickupDate,
		"CN": FieldLoginUserID,
		"CO": FieldLoginPassword,
		"CP": FieldLocationCode,
		"CQ": FieldValidPatronPassword,
		"AY": FieldSequenceNumber,
		"AZ": FieldChecksum,
	}
)

func isOptional(mt msgType, ft fieldType) bool {
	for _, f := range msgDefinitions[mt].OptionalVar {
		if f == ft {
			return true
		}
	}
	return false
}
