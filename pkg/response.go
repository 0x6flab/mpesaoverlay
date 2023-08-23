package pkg

// TokenResp struct.
type TokenResp struct {
	AccessToken string `json:"access_token,omitempty"`
	Expiry      string `json:"expires_in,omitempty"`
}

// ErrorResp struct.
type ErrorResp struct {
	RequestID string `json:"requestId,omitempty"`
	Code      string `json:"errorCode,omitempty"`
	Message   string `json:"errorMessage,omitempty"`
}

// ValidResp struct.
type ValidResp struct {
	OriginatorConversationID string `json:"OriginatorConversationID,omitempty"` // The unique request ID for tracking a transaction
	ConversationID           string `json:"ConversationID,omitempty"`           // The unique request ID returned by mpesa for each request made
	ResponseDescription      string `json:"ResponseDescription,omitempty"`      // Response Description message
	ResponseCode             string `json:"ResponseCode,omitempty"`
}

// ExpressSimulateResp struct.
type ExpressSimulateResp struct {
	ResponseDescription string `json:"ResponseDescription,omitempty"` // Response description is an acknowledgment message from the API that gives the status of the request submission.
	ResponseCode        string `json:"ResponseCode,omitempty"`        // This is a Numeric status code that indicates the status of the transaction submission. 0 means successful submission and any other code means an error occurred.
	MerchantRequestID   string `json:"MerchantRequestID,omitempty"`   // This is a global unique Identifier for any submitted payment request.
	CheckoutRequestID   string `json:"CheckoutRequestID,omitempty"`   // This is a global unique identifier of the processed checkout transaction request.
	CustomerMessage     string `json:"CustomerMessage,omitempty"`     // This is a message that your system can display to the customer as an acknowledgment of the payment request submission.
}

// ExpressQueryResp struct.
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

// QRResp struct.
type QRResp struct {
	ResponseDescription string `json:"ResponseDescription,omitempty"` // This is a response describing the status of the transaction.
	ResponseCode        string `json:"ResponseCode,omitempty"`        // Used to return the Transaction Type.
	RequestID           string `json:"RequestID,omitempty"`
	QRCode              string `json:"QRCode,omitempty"` // QR Code Image Data/String/Image.
}

// C2BRegisterURLResp struct.
type C2BRegisterURLResp struct {
	ValidResp
}

// C2BSimulateResp struct.
type C2BSimulateResp struct {
	ValidResp
}

// B2CResp struct.
type B2CResp struct {
	ValidResp
}

// TransactionResp struct.
type TransactionResp struct {
	ValidResp
}

// AccBalanceResp struct.
type AccBalanceResp struct {
	ValidResp
}

// ReversalResp struct.
type ReversalResp struct {
	ValidResp
}
