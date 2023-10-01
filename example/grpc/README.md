# Examples

A [postman collection](https://www.postman.com/ox6flab/workspace/mpesaoverlay) is available for testing the API.

## CLI

[grpcurl](https://github.com/fullstorydev/grpcurl) is a useful tool for testing the gRPC API.

## Token

```bash
grpcurl -plaintext -d @ localhost:9000 mpesaoverlay.grpc.Service/Token <<EOM
{}
EOM
```

## STKPush

```bash
grpcurl -plaintext -d @ localhost:9000 mpesaoverlay.grpc.Service/ExpressSimulate <<EOM
{
    "accountReference": "CompanyXLTD",
    "amount": "1",
    "businessShortCode": "174379",
    "callBackURL": "https://69a2-105-163-2-116.ngrok.io",
    "partyA": "254720136609",
    "partyB": "174379",
    "passKey": "bfb279f9aa9bdbcf158e97dd71a467cd2e0c893059b10f78e6b72ada1ed2c919",
    "phoneNumber": "254720136609",
    "transactionDesc": "Payment of X",
    "transactionType": "CustomerPayBillOnline"
}
EOM
```

## STKQuery

```bash
grpcurl -plaintext -d @ localhost:9000 mpesaoverlay.grpc.Service/ExpressQuery <<EOM
{
    "businessShortCode": "174379",
    "checkoutRequestID": "ws_CO_01102023223251412720136609",
    "passKey": "bfb279f9aa9bdbcf158e97dd71a467cd2e0c893059b10f78e6b72ada1ed2c919"
}
EOM
```

## B2CPayment

```bash
grpcurl -plaintext -d @ localhost:9000 mpesaoverlay.grpc.Service/B2CPayment <<EOM
{
    "amount": "10",
    "commandID": "BusinessPayment",
    "initiatorName": "testapi",
    "initiatorPassword": "Safaricom999!*!",
    "occasion": "test",
    "originatorConversationID": "ullamco fugiat consectetur commodo reprehenderit",
    "partyA": "600986",
    "partyB": "254712345678",
    "queueTimeOutURL": "https://example.com/timeout",
    "remarks": "test",
    "resultURL": "https://example.com/result"
}
EOM
```

## AccountBalance

```bash
grpcurl -plaintext -d @ localhost:9000 mpesaoverlay.grpc.Service/AccountBalance <<EOM
{
    "commandID": "AccountBalance",
    "identifierType": 4,
    "initiatorName": "testapi",
    "initiatorPassword": "Safaricom999!*!",
    "partyA": "600772",
    "queueTimeOutURL": "https://example.com",
    "remarks": "test",
    "resultURL": "https://example.com"
}
EOM
```

## C2BRegisterURL

```bash
grpcurl -plaintext -d @ localhost:9000 mpesaoverlay.grpc.Service/C2BRegisterURL <<EOM
{
    "confirmationURL": "https://example.com",
    "responseType": "Completed",
    "shortCode": "600981",
    "validationURL": "https://example.com"
}
EOM
```

## C2BSimulate

```bash
grpcurl -plaintext -d @ localhost:9000 mpesaoverlay.grpc.Service/C2BSimulate <<EOM
{
    "amount": "10",
    "billRefNumber": "veniam ullamco",
    "commandID": "CustomerBuyGoodsOnline",
    "msisdn": "254712345678",
    "shortCode": "600986"
}
EOM
```

## GenerateQR

```bash
grpcurl -plaintext -d @ localhost:9000 mpesaoverlay.grpc.Service/GenerateQR <<EOM
{
    "amount": "2000",
    "cPI": "174379",
    "merchantName": "Test Supermarket",
    "refNo": "Invoice No",
    "size": "300",
    "trxCode": "BG"
}
EOM
```

## Reverse

```bash
grpcurl -plaintext -d @ localhost:9000 mpesaoverlay.grpc.Service/Reverse <<EOM
{
    "amount": "10",
    "commandID": "TransactionReversal",
    "initiatorName": "testapi",
    "initiatorPassword": "Safaricom999!*!",
    "occasion": "test",
    "queueTimeOutURL": "https://example.com/",
    "recieverIdentifierType": 11,
    "receiverParty": "600992",
    "remarks": "test",
    "resultURL": "https://example.com/",
    "transactionID": "RI704KI9RW"
}
EOM
```

## TransactionStatus

```bash
grpcurl -plaintext -d @ localhost:9000 mpesaoverlay.grpc.Service/TransactionStatus <<EOM
{
    "commandID": "TransactionStatusQuery",
    "identifierType": 1,
    "initiatorName": "testapi",
    "initiatorPassword": "Safaricom999!*!",
    "occasion": "test",
    "partyA": "254759764065",
    "queueTimeOutURL": "https://example.com",
    "remarks": "test",
    "resultURL": "https://example.com",
    "transactionID": "RI704KI9RW"
}
EOM
```

## RemitTax

```bash
grpcurl -plaintext -d @ localhost:9000 mpesaoverlay.grpc.Service/RemitTax <<EOM
{
    "accountReference": "353353",
    "amount": "10",
    "commandID": "PayTaxToKRA",
    "initiatorName": "testapi",
    "initiatorPassword": "Safaricom999!*!",
    "partyA": "600978",
    "partyB": "572572",
    "queueTimeOutURL": "https://example.com",
    "recieverIdentifierType": 4,
    "remarks": "test",
    "resultURL": "https://example.com",
    "senderIdentifierType": 4
}
EOM
```
