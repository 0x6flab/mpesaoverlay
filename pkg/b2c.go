package pkg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

func (sdk mSDK) B2CPayment(b2cReq B2CPaymentReq) (B2CPaymentResp, error) {
	if err := b2cReq.Validate(); err != nil {
		return B2CPaymentResp{}, err
	}

	var err error
	b2cReq.SecurityCredential, err = sdk.generateSecurityCredential(b2cReq.InitiatorPassword)
	if err != nil {
		return B2CPaymentResp{}, err
	}

	if b2cReq.OriginatorConversationID == "" {
		uuid, err := uuid.NewRandom()
		if err == nil && uuid.String() != "" {
			b2cReq.OriginatorConversationID = uuid.String()
		}
	}

	data, err := json.Marshal(b2cReq)
	if err != nil {
		return B2CPaymentResp{}, err
	}

	url := fmt.Sprintf("%s/%s", sdk.baseURL, b2cEndpoint)

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
	if err != nil {
		return B2CPaymentResp{}, err
	}

	resp, err := sdk.sendRequest(req)
	if err != nil {
		return B2CPaymentResp{}, err
	}

	var b2cr B2CPaymentResp
	if err := json.Unmarshal(resp, &b2cr); err != nil {
		return B2CPaymentResp{}, err
	}

	return b2cr, nil
}
