package apiNode

//go:generate protoc -I. --go_out=plugins=grpc:. api_node

const ApiGatewayNodeRequestTopic = "ankr.api.api-gateway-node-request"
