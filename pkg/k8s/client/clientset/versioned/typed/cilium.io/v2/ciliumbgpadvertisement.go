// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

// Code generated by client-gen. DO NOT EDIT.

package v2

import (
	context "context"

	ciliumiov2 "github.com/cilium/cilium/pkg/k8s/apis/cilium.io/v2"
	scheme "github.com/cilium/cilium/pkg/k8s/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// CiliumBGPAdvertisementsGetter has a method to return a CiliumBGPAdvertisementInterface.
// A group's client should implement this interface.
type CiliumBGPAdvertisementsGetter interface {
	CiliumBGPAdvertisements() CiliumBGPAdvertisementInterface
}

// CiliumBGPAdvertisementInterface has methods to work with CiliumBGPAdvertisement resources.
type CiliumBGPAdvertisementInterface interface {
	Create(ctx context.Context, ciliumBGPAdvertisement *ciliumiov2.CiliumBGPAdvertisement, opts v1.CreateOptions) (*ciliumiov2.CiliumBGPAdvertisement, error)
	Update(ctx context.Context, ciliumBGPAdvertisement *ciliumiov2.CiliumBGPAdvertisement, opts v1.UpdateOptions) (*ciliumiov2.CiliumBGPAdvertisement, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*ciliumiov2.CiliumBGPAdvertisement, error)
	List(ctx context.Context, opts v1.ListOptions) (*ciliumiov2.CiliumBGPAdvertisementList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *ciliumiov2.CiliumBGPAdvertisement, err error)
	CiliumBGPAdvertisementExpansion
}

// ciliumBGPAdvertisements implements CiliumBGPAdvertisementInterface
type ciliumBGPAdvertisements struct {
	*gentype.ClientWithList[*ciliumiov2.CiliumBGPAdvertisement, *ciliumiov2.CiliumBGPAdvertisementList]
}

// newCiliumBGPAdvertisements returns a CiliumBGPAdvertisements
func newCiliumBGPAdvertisements(c *CiliumV2Client) *ciliumBGPAdvertisements {
	return &ciliumBGPAdvertisements{
		gentype.NewClientWithList[*ciliumiov2.CiliumBGPAdvertisement, *ciliumiov2.CiliumBGPAdvertisementList](
			"ciliumbgpadvertisements",
			c.RESTClient(),
			scheme.ParameterCodec,
			"",
			func() *ciliumiov2.CiliumBGPAdvertisement { return &ciliumiov2.CiliumBGPAdvertisement{} },
			func() *ciliumiov2.CiliumBGPAdvertisementList { return &ciliumiov2.CiliumBGPAdvertisementList{} },
		),
	}
}
