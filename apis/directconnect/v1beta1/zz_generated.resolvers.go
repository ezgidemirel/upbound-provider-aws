/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	resource "github.com/upbound/upjet/pkg/resource"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this ConnectionAssociation.
func (mg *ConnectionAssociation) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ConnectionID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.ConnectionIDRef,
		Selector:     mg.Spec.ForProvider.ConnectionIDSelector,
		To: reference.To{
			List:    &ConnectionList{},
			Managed: &Connection{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ConnectionID")
	}
	mg.Spec.ForProvider.ConnectionID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ConnectionIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.LagID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.LagIDRef,
		Selector:     mg.Spec.ForProvider.LagIDSelector,
		To: reference.To{
			List:    &LagList{},
			Managed: &Lag{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.LagID")
	}
	mg.Spec.ForProvider.LagID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.LagIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this PublicVirtualInterface.
func (mg *PublicVirtualInterface) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ConnectionID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ConnectionIDRef,
		Selector:     mg.Spec.ForProvider.ConnectionIDSelector,
		To: reference.To{
			List:    &ConnectionList{},
			Managed: &Connection{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ConnectionID")
	}
	mg.Spec.ForProvider.ConnectionID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ConnectionIDRef = rsp.ResolvedReference

	return nil
}
