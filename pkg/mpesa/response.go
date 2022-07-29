package mpesa

// TokenResp struct
type TokenResp struct {
	AccessToken string `json:"access_token,omitempty"`
	Expiry      uint64 `json:"expires_in,omitempty"`
}

// ErrorResp struct
type ErrorResp struct {
	RequestID string `json:"requestId,omitempty"`
	Code      string `json:"errorCode,omitempty"`
	Message   string `json:"errorMessage,omitempty"`
}

// ValidResp struct
type ValidResp struct {
	OriginatorConversationID string `json:"OriginatorConversationID,omitempty"` // The unique request ID for tracking a transaction
	ConversationID           string `json:"ConversationID,omitempty"`           // The unique request ID returned by mpesa for each request made
	ResponseDescription      string `json:"ResponseDescription,omitempty"`      // Response Description message
	ResponseCode             string `json:"ResponseCode,omitempty"`
}

// ExpressSimulateResp struct
type ExpressSimulateResp struct {
	OriginatorConversationID string `json:"OriginatorConversationID,omitempty"` // The unique request ID for tracking a transaction
	ConversationID           string `json:"ConversationID,omitempty"`           // The unique request ID returned by mpesa for each request made
	ResponseDescription      string `json:"ResponseDescription,omitempty"`      // Response Description message
	ResponseCode             string `json:"ResponseCode,omitempty"`
	MerchantRequestID        string `json:"MerchantRequestID,omitempty"`
	CheckoutRequestID        string `json:"CheckoutRequestID,omitempty"`
	CustomerMessage          string `json:"CustomerMessage,omitempty"`
}

// ExpressQueryResp struct
type ExpressQueryResp struct {
	OriginatorConversationID string `json:"OriginatorConversationID,omitempty"` // The unique request ID for tracking a transaction
	ConversationID           string `json:"ConversationID,omitempty"`           // The unique request ID returned by mpesa for each request made
	ResponseDescription      string `json:"ResponseDescription,omitempty"`      // Response Description message. It can be a Success submission message or an error description.
	ResponseCode             string `json:"ResponseCode,omitempty"`
	MerchantRequestID        string `json:"MerchantRequestID,omitempty"` // This is a global unique Identifier for any submitted payment request.
	CheckoutRequestID        string `json:"CheckoutRequestID,omitempty"` // This is a global unique identifier of the processed checkout transaction request.
	CustomerMessage          string `json:"CustomerMessage,omitempty"`   // This is a message that your system can display to the Customer as an acknowledgement of the payment request submission.
	ResultCode               string `json:"ResultCode,omitempty"`
	ResultDesc               string `json:"ResultDesc,omitempty"`
}

// QRResp struct
type QRResp struct {
	OriginatorConversationID string `json:"OriginatorConversationID,omitempty"` // The unique request ID for tracking a transaction
	ConversationID           string `json:"ConversationID,omitempty"`           // The unique request ID returned by mpesa for each request made
	ResponseDescription      string `json:"ResponseDescription,omitempty"`      // Response Description message
	ResponseCode             string `json:"ResponseCode,omitempty"`
	RequestID                string `json:"RequestID,omitempty"`
	QRCode                   string `json:"QRCode,omitempty"` // QR Code Image Data/String/Image.
}

// C2BRegisterURLResp struct
type C2BRegisterURLResp struct {
	ValidResp
}

// C2BSimulateResp struct
type C2BSimulateResp struct {
	ValidResp
}

// B2CResp struct
type B2CResp struct {
	ValidResp
}

// TransactionResp struct
type TransactionResp struct {
	ValidResp
}

// AccBalanceResp struct
type AccBalanceResp struct {
	ValidResp
}

// ReversalResp struct
type ReversalResp struct {
	ValidResp
}
