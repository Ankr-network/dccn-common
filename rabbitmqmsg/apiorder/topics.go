package apiOrder

//go:generate protoc -I. --go_out=plugins=grpc:. api_order

const ApiGatewayOrderRequestTopic = "ankr.api.api-gateway-order-request"
const ApiGatewayOrderResponseTopic = "ankr.api.api-gateway-order-response"
