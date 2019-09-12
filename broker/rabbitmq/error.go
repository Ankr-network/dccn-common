package rabbitmq

import "errors"

var (
	ErrMessageIsNotProtoMessage = errors.New("message must be proto.Message")
	ErrInvalidHandler           = errors.New("invalid handler, must be func and have single proto.Message implement and return single error")
	ErrPublishMessageNotAck     = errors.New("message not ack by broker")
)
