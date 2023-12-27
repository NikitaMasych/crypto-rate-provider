package dispatcher

import (
	"context"
	"email/domain"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	SendEmail endpoint.Endpoint
}

func NewEndpointSet(svc EmailService) Endpoints {
	return Endpoints{
		SendEmail: SendEmailEndpoint(svc),
	}
}

func SendEmailEndpoint(svc EmailService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(domain.SendEmailRequest)
		err := svc.SendEmail(req)
		return nil, err
	}
}
