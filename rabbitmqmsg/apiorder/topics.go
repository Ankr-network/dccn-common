package apiOrder

//go:generate protoc -I. --go_out=plugins=grpc,paths=source_relative:. api_order

const ApiGatewayOrderRequestTopic = "ankr.api.api-gateway-order-request"
const ApiGatewayV2OrderRequestTopic = "ankr.api.api-gateway-v2-order-request"
const ApiGatewayOrderResponseTopic = "ankr.api.api-gateway-order-response"
