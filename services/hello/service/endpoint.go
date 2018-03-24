package service

import (
	"api/core"
	"api/people"
	"context"

	"google.golang.org/grpc"
)

// Endpoint ...
type Endpoint struct {
	address people.AddressClient
	bio     people.BioClient
}

// NewClient ...
func NewClient(addr string) (*Endpoint, error) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return &Endpoint{
		address: people.NewAddressClient(conn),
		bio:     people.NewBioClient(conn),
	}, nil
}

// GetPrimary address
func (h *Endpoint) GetPrimary(ctx context.Context, query *core.Info) (*people.PrimaryAddress, error) {
	return h.address.GetPrimaryAddress(ctx, query)
}

// GetName bio
func (h *Endpoint) GetName(ctx context.Context, query *core.Info) (*people.Name, error) {
	return h.bio.GetName(ctx, query)
}
