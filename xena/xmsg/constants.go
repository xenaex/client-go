package xmsg

const (
	// MESSAGE.proto

	MsgType_UnknownMsgType                   = ""
	MsgType_RejectMsgType                    = "3"
	MsgType_ExecutionReportMsgType           = "8"
	MsgType_OrderCancelRejectMsgType         = "9"
	MsgType_LogonMsgType                     = "A"
	MsgType_TradeCaptureReportMsgType        = "AE"
	MsgType_OrderMassStatusRequest           = "AF"
	MsgType_AccountStatusReportRequest       = "XAA"
	MsgType_AccountStatusReport              = "XAR"
	MsgType_AccountStatusUpdateReport        = "XAF"
	MsgType_NewOrderSingleMsgType            = "D"
	MsgType_NewOrderListMsgType              = "E"
	MsgType_OrderCancelRequestMsgType        = "F"
	MsgType_OrderCancelReplaceRequestMsgType = "G"
	MsgType_OrderStatusRequest               = "H"
	MsgType_ListStatus                       = "N"
	MsgType_MarketDataRequest                = "V"
	MsgType_MarketDataSnapshotFullRefresh    = "W"
	MsgType_MarketDataIncrementalRefresh     = "X"
	MsgType_MarketDataRequestReject          = "Y"
	MsgType_OrderMassStatusResponse          = "U8"
	MsgType_TradeCaptureReportRequest        = "AD"
	MsgType_MassTradeCaptureReportResponse   = "U9"
	MsgType_PositionMaintenanceRequest       = "AL"
	MsgType_PositionMaintenanceReport        = "AM"
	MsgType_RequestForPositions              = "AN"
	MsgType_PositionReport                   = "AP"
	MsgType_MassPositionReport               = "MAP"
	MsgType_MarginRequirementReport          = "CJ"
	MsgType_Heartbeat                        = "0"
	MsgType_OrderMassCancelRequest           = "q"
	MsgType_OrderMassCancelReport            = "r"
	MsgType_ApplicationHeartbeat             = "XAH"

	// enum BusinessRejectReason
	BusinessRejectReason_UnknownBisRejReason     = ""
	BusinessRejectReason_ApplicationNotAvailable = "4"
	BusinessRejectReason_ThrottleLimitExceeded   = "8"

	// MARGIN.proto

	// enum MarginAmtType
	MarginAmtType_UnknownMarginTypeType = ""
	MarginAmtType_CoreMargin            = "7"
	MarginAmtType_InitialMargin         = "11"

	// enum MarginReqmtRptType
	MarginReqmtRptType_UnknownReqmtType = ""
	MarginReqmtRptType_SummaryType      = "0" // Real FIX value is 0

	// ORDERS.proto

	// enum Side
	Side_UnknownSide = ""
	Side_Buy         = "1"
	Side_Sell        = "2"

	// enum OrdType
	OrdType_UnknownOrdType = ""
	// Market order type is a order type for selling or buying instrument
	// by current market prices.
	OrdType_Market = "1"
	// Limit is a deferred order type. Orders by this type could be executed
	// by best price or by order price.
	OrdType_Limit = "2"
	// Stop is a deferred order type. It executed when current quotes achieve
	// order's stop-price of this type. After execution order of this type will
	// be converted to market order.
	OrdType_Stop = "3"
	// StopLimit is as deferred order type. Almost the same as Stop order,
	// except after execution it will be converted to market order with certain
	// price.
	OrdType_StopLimit       = "4"
	OrdType_MarketIfTouched = "J" // real FIX value is 'J'
	OrdType_Pegged          = "P" // real FIX value is 'P'

	// enum ExecInst
	ExecInst_UnknownExecInst           = ""
	ExecInst_StayOnOfferSide           = "0"     // real FIX value is 0
	ExecInst_PegToOfferSide            = "9"     // real FIX value is 9
	ExecInst_AllOrNone                 = "G"     // real FIX value is 'G'
	ExecInst_IgnoreNotionalValueChecks = "x"     // real FIX value is 'x'
	ExecInst_Suspend                   = "s"     // real FIX value is 's'
	ExecInst_LiquidationOrder          = "Y"     // real FIX value is 'Y'
	ExecInst_IgnorePriceRangeChecks    = "70081" // real FIX value is none
	ExecInst_CancelOnConnectionLoss    = "o"

	// enum TimeInForce
	TimeInForce_UnknownTimeInForce = ""
	TimeInForce_GoodTillCancel     = "1"
	TimeInForce_ImmediateOrCancel  = "3"
	TimeInForce_FillOrKill         = "4"

	// enum PegPriceType
	PegPriceType_UnknownPegType  = ""
	PegPriceType_TrailingStopPeg = "8" // real FIX value is 8

	// enum PegOffsetType
	PegOffsetType_UnknonwOffsetType = ""
	PegOffsetType_BasisPoints       = "1" // real FIX value is 1

	// enum ExecType
	ExecType_UnknownExecType              = ""
	ExecType_NewExec                      = "0" // real FIX value is 0
	ExecType_CanceledExec                 = "4" // real FIX value is 4
	ExecType_ReplacedExec                 = "5" // real FIX value is 5
	ExecType_PendingCancelExec            = "6" // real FIX value is 6
	ExecType_RejectedExec                 = "8" // real FIX value is 8
	ExecType_SuspendedExec                = "9" // real FIX value is 9
	ExecType_PendingNewExec               = "A" // real FIX value is 'A'
	ExecType_Restated                     = "D" // real FIX value is 'D'
	ExecType_PendingReplaceExec           = "E" // real FIX value is 'E'
	ExecType_Trade                        = "F" // real FIX value is 'F'
	ExecType_OrderStatus                  = "I" // real FIX value is 'I'
	ExecType_TriggeredOrActivatedBySystem = "L" // real FIX value is 'L'

	// enum ExecRestatementReason
	ExecRestatementReason_UnknownRestatementReason = ""
	ExecRestatementReason_RepricingOfOrder         = "3" // real FIX value is 3"

	// enum OrdStatus
	OrdStatus_UnknownOrdStatus  = ""
	OrdStatus_NewOrd            = "0" // real FIX value is 0
	OrdStatus_PartiallyFilled   = "1" // real FIX value is 1
	OrdStatus_Filled            = "2" // real FIX value is 2
	OrdStatus_CanceledOrd       = "4" // real FIX value is 4
	OrdStatus_PendingCancelOrd  = "6" // real FIX value is 6
	OrdStatus_Stopped           = "7" // real FIX value is 7
	OrdStatus_RejectedOrd       = "8" // real FIX value is 8
	OrdStatus_Suspended         = "9" // real FIX value is 9
	OrdStatus_PendingNewOrd     = "A" // real FIX value is 'A'
	OrdStatus_Expired           = "C" // real FIX value is 'C'
	OrdStatus_PendingReplaceOrd = "E" // real FIX value is 'E'

	// enum OrdRejReason
	OrdRejReason_UnknownOrdRejReason            = ""
	OrdRejReason_UnknownSymbol                  = "1"
	OrdRejReason_ExchangeClosed                 = "2"
	OrdRejReason_OrderExceedsLimit              = "3"
	OrdRejReason_DuplicateOrder                 = "6"
	OrdRejReason_UnsupportedOrderCharacteristic = "11"
	OrdRejReason_IncorrectQuantity              = "13"
	OrdRejReason_UnknownAccount                 = "15"
	OrdRejReason_PriceExceedsCurrentPriceBand   = "16"
	OrdRejReason_Other                          = "99"
	OrdRejReason_StopPriceInvalid               = "100"

	// enum LiquidityInd
	LiquidityInd_UnknownLiquidityInd = ""
	LiquidityInd_AddedLiquidity      = "1"
	LiquidityInd_RemovedLiquidity    = "2"

	// enum SettlType
	SettlType_UnknownSettlType = ""
	SettlType_Regular          = "0" // real FIX value is 0
	SettlType_Cash             = "1" // real FIX value is 1

	// enum CxlRejResponseTo
	CxlRejResponseTo_UnknownCxlRejResponseTo             = ""
	CxlRejResponseTo_OrderCancelRequestCxlRejResponseTo  = "1"
	CxlRejResponseTo_OrderCancelReplaceRequestResponseTo = "2"

	// enum CxlRejReason
	CxlRejReason_UnknownCxlRejReason         = ""
	CxlRejReason_TooLateToCancel             = "0"  // real FIX value is 0
	CxlRejReason_UnknownOrder                = "1"  // real FIX value is 1
	CxlRejReason_OrderAlreadyInPendingStatus = "3"  // real FIX value is 3
	CxlRejReason_DuplicateClOrdID            = "6"  // real FIX value is 6
	CxlRejReason_OtherCxlRejReason           = "99" // real FIX value is 99

	// enum BidType
	BidType_UnknownBidType   = ""
	BidType_NonDisclosed     = "1"
	BidType_Disclosed        = "2"
	BidType_NoBiddingProcess = "3"

	// enum ContingencyType
	ContingencyType_UnknownContingencyType         = ""
	ContingencyType_OneCancelsTheOther             = "1"
	ContingencyType_OneTriggersTheOther            = "2"
	ContingencyType_OneUpdatesTheOtherAbsolute     = "3"
	ContingencyType_OneUpdatesTheOtherProportional = "4"

	// enum ListStatusType
	ListStatusType_UnknownListStatusType     = ""
	ListStatusType_AckListStatusType         = "1"
	ListStatusType_ResponseListStatusType    = "2"
	ListStatusType_TimedListStatusType       = "3"
	ListStatusType_ExecStartedListStatusType = "4"
	ListStatusType_AllDoneListStatusType     = "5"
	ListStatusType_AlertListStatusType       = "6"

	// enum ListOrderStatus
	ListOrderStatus_UnknownListOrderStatus              = ""
	ListOrderStatus_InBiddingProcessListOrderStatus     = "1"
	ListOrderStatus_ReceivedForExecutionListOrderStatus = "2"
	ListOrderStatus_ExecutingListOrderStatus            = "3"
	ListOrderStatus_CancellingListOrderStatus           = "4"
	ListOrderStatus_AlertListOrderStatus                = "5"
	ListOrderStatus_AllDoneListOrderStatus              = "6"
	ListOrderStatus_RejectListOrderStatus               = "7"

	// enum ListRejectReason
	ListRejectReason_UnknownListRejectReason                        = ""
	ListRejectReason_UnsupportedOrderCharacteristicListRejectReason = "11"
	ListRejectReason_ExchangeClosedListRejectReason                 = "2"
	ListRejectReason_TooLateToEnterListRejectReason                 = "4"
	ListRejectReason_UnknownOrderListRejectReason                   = "5"
	ListRejectReason_DuplicateOrderListRejectReason                 = "6"
	ListRejectReason_OtherListRejectReason                          = "99"

	//
	// enum TriggerAction
	TriggerAction_UnknownTriggerAction = ""
	TriggerAction_Activate             = "1"
	TriggerAction_SetNewCapPrice       = "4"
	TriggerAction_Transform            = "3101"

	// enum TriggerType
	TriggerType_UnknownTriggerType = ""
	TriggerType_PartialExecution   = "1" // real FIX value is 1

	// enum TriggerScope
	TriggerScope_UnknownTriggerScope = ""
	TriggerScope_OtherOrder          = "1" // real FIX value is 1

	// enum PositionEffect
	PositionEffect_UnknownPositionEffect = ""
	PositionEffect_Close                 = "C" // real FIX value is C, number in list is 1
	PositionEffect_Open                  = "O" // real FIX value is O, number in list is 5

	// enum PartieRole
	PartieRole_UnknownRole = ""
	PartieRole_ClientID    = "3"
	PartieRole_ContraFirm  = "17"

	// enum MDEntryType
	MDEntryType_UnknownMDEntryType = ""
	MDEntryType_Bid                = "0" // real FIX value is 0
	MDEntryType_Offer              = "1" // real FIX value is 1
	MDEntryType_Trade              = "2" // real FIX value is 2
	MDEntryType_MarketBid          = "b" // real FIX value is 'b'
	MDEntryType_MarketOffer        = "c" // real FIX value is 'c'

	// enum MDUpdateAction
	MDUpdateAction_UnknownMDUpdateAction = ""
	MDUpdateAction_NewAction             = "0" // real FIX value is 0
	MDUpdateAction_ChangeAction          = "1" // real FIX value is 1
	MDUpdateAction_DeleteAction          = "2" // real FIX value is 2

	// enum MDBookType
	MDBookType_UnknownMDBookType = ""
	MDBookType_TopOfBook         = "1"
	MDBookType_PriceDepth        = "2"
	MDBookType_OrderDepth        = "3"

	// enum SubscriptionRequestType
	SubscriptionRequestType_UnknownSubscriptionRequestType = ""
	SubscriptionRequestType_SnapshotAndUpdates             = "1"
	SubscriptionRequestType_DisablePreviousSnapshot        = "2"

	// num ThrottleType
	ThrottleType_InboundRate         = "0"
	ThrottleType_OutstandingRequests = "1"

	// enum ThrottleTimeUnit
	ThrottleTimeUnit_Seconds             = "0"
	ThrottleTimeUnit_TenthsOfASecond     = "1"
	ThrottleTimeUnit_HundredthsOfASecond = "2"
	ThrottleTimeUnit_Milliseconds        = "3"
	ThrottleTimeUnit_Microseconds        = "4"
	ThrottleTimeUnit_Nanoseconds         = "5"
	ThrottleTimeUnit_Minutes             = "10"
	ThrottleTimeUnit_Hours               = "11"

	// enum BalanceChangeReason
	BalanceChangeReason_UnknownReason               = "0"
	BalanceChangeReason_DepositReason               = "1"
	BalanceChangeReason_WithdrawReason              = "2"
	BalanceChangeReason_DealReason                  = "3"
	BalanceChangeReason_HoldReason                  = "4"
	BalanceChangeReason_ReleaseReason               = "5"
	BalanceChangeReason_RebateReason                = "6"
	BalanceChangeReason_CommissionReason            = "7"
	BalanceChangeReason_PositionCloseReason         = "8"
	BalanceChangeReason_ClearingReason              = "9"
	BalanceChangeReason_InterestPaymentReason       = "10"
	BalanceChangeReason_PremiumPaymentReason        = "11"
	BalanceChangeReason_RiskAdjustmentPaymentReason = "12"
	BalanceChangeReason_SpotRewardReason            = "13" // Referral program spot reward
	BalanceChangeReason_MarginRewardReason          = "14" // Referral program margin reward

	// enum PaymentType
	PaymentType_UnknownPaymentType = "0"
	PaymentType_Commission         = "40"
	PaymentType_Interest           = "41"
	PaymentType_Settlement         = "42"
	PaymentType_CumulativePayments = "43"

	// enum RelatedTradeType
	RelatedTradeType_UnknownTradeType = "0"
	RelatedTradeType_OpenExecID       = "1"
	RelatedTradeType_CloseExecID      = "2"
	RelatedTradeType_ParentPositionID = "3"

	// enum MassStatusReqType
	MassStatusReqType_ActiveOrders         = "11"
	MassStatusReqType_DoneOrdersLastStatus = "12"
	MassStatusReqType_History              = "13"
)
