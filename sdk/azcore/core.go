//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package azcore

import (
	"reflect"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/internal/exported"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/internal/shared"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/tracing"
)

// AccessToken represents an Azure service bearer access token with expiry information.
type AccessToken = exported.AccessToken

// TokenCredential represents a credential capable of providing an OAuth token.
type TokenCredential = exported.TokenCredential

// holds sentinel values used to send nulls
var nullables map[reflect.Type]interface{} = map[reflect.Type]interface{}{}

// NullValue is used to send an explicit 'null' within a request.
// This is typically used in JSON-MERGE-PATCH operations to delete a value.
func NullValue[T any]() T {
	t := shared.TypeOfT[T]()
	v, found := nullables[t]
	if !found {
		var o reflect.Value
		if k := t.Kind(); k == reflect.Map {
			o = reflect.MakeMap(t)
		} else if k == reflect.Slice {
			// empty slices appear to all point to the same data block
			// which causes comparisons to become ambiguous.  so we create
			// a slice with len/cap of one which ensures a unique address.
			o = reflect.MakeSlice(t, 1, 1)
		} else {
			o = reflect.New(t.Elem())
		}
		v = o.Interface()
		nullables[t] = v
	}
	// return the sentinel object
	return v.(T)
}

// IsNullValue returns true if the field contains a null sentinel value.
// This is used by custom marshallers to properly encode a null value.
func IsNullValue[T any](v T) bool {
	// see if our map has a sentinel object for this *T
	t := reflect.TypeOf(v)
	if o, found := nullables[t]; found {
		o1 := reflect.ValueOf(o)
		v1 := reflect.ValueOf(v)
		// we found it; return true if v points to the sentinel object.
		// NOTE: maps and slices can only be compared to nil, else you get
		// a runtime panic.  so we compare addresses instead.
		return o1.Pointer() == v1.Pointer()
	}
	// no sentinel object for this *t
	return false
}

// ClientOptions contains configuration settings for a client's pipeline.
type ClientOptions = policy.ClientOptions

// Client is a basic HTTP client.  It consists of a pipeline and tracing provider.
type Client struct {
	pl runtime.Pipeline
	tr tracing.Tracer
}

// NewClient creates a new Client instance with the provided values.
//   - clientName - the fully qualified name of the client ("package.Client"); this is used by the tracing provider when creating spans
//   - moduleVersion - the semantic version of the containing module; used by the telemetry policy
//   - plOpts - pipeline configuration options; can be the zero-value
//   - options - optional client configurations; pass nil to accept the default values
func NewClient(clientName, moduleVersion string, plOpts runtime.PipelineOptions, options *ClientOptions) (*Client, error) {
	pkg, err := shared.ExtractPackageName(clientName)
	if err != nil {
		return nil, err
	}

	if options == nil {
		options = &ClientOptions{}
	}

	if !options.Telemetry.Disabled {
		if err := shared.ValidateModVer(moduleVersion); err != nil {
			return nil, err
		}
	}

	pl := runtime.NewPipeline(pkg, moduleVersion, plOpts, options)

	tr := options.TracingProvider.NewTracer(clientName, moduleVersion)
	if tr.Enabled() && plOpts.TracingNamespace != "" {
		tr.SetAttributes(tracing.Attribute{Key: "az.namespace", Value: plOpts.TracingNamespace})
	}
	return &Client{pl: pl, tr: tr}, nil
}

// Pipeline returns the pipeline for this client.
func (c *Client) Pipeline() runtime.Pipeline {
	return c.pl
}

// Tracer returns the tracer for this client.
func (c *Client) Tracer() tracing.Tracer {
	return c.tr
}
