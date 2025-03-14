//
//
// File generated from our OpenAPI spec
//
//

// Package paymentintent provides the /payment_intents APIs
package paymentintent

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/form"
)

// Client is used to invoke /payment_intents APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// New creates a new payment intent.
func New(params *stripe.PaymentIntentParams) (*stripe.PaymentIntent, error) {
	return getC().New(params)
}

// New creates a new payment intent.
func (c Client) New(params *stripe.PaymentIntentParams) (*stripe.PaymentIntent, error) {
	paymentintent := &stripe.PaymentIntent{}
	err := c.B.Call(
		http.MethodPost,
		"/v1/payment_intents",
		c.Key,
		params,
		paymentintent,
	)
	return paymentintent, err
}

// Get returns the details of a payment intent.
func Get(id string, params *stripe.PaymentIntentParams) (*stripe.PaymentIntent, error) {
	return getC().Get(id, params)
}

// Get returns the details of a payment intent.
func (c Client) Get(id string, params *stripe.PaymentIntentParams) (*stripe.PaymentIntent, error) {
	path := stripe.FormatURLPath("/v1/payment_intents/%s", id)
	paymentintent := &stripe.PaymentIntent{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, paymentintent)
	return paymentintent, err
}

// Update updates a payment intent's properties.
func Update(id string, params *stripe.PaymentIntentParams) (*stripe.PaymentIntent, error) {
	return getC().Update(id, params)
}

// Update updates a payment intent's properties.
func (c Client) Update(id string, params *stripe.PaymentIntentParams) (*stripe.PaymentIntent, error) {
	path := stripe.FormatURLPath("/v1/payment_intents/%s", id)
	paymentintent := &stripe.PaymentIntent{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, paymentintent)
	return paymentintent, err
}

// Cancel is the method for the `POST /v1/payment_intents/{intent}/cancel` API.
func Cancel(id string, params *stripe.PaymentIntentCancelParams) (*stripe.PaymentIntent, error) {
	return getC().Cancel(id, params)
}

// Cancel is the method for the `POST /v1/payment_intents/{intent}/cancel` API.
func (c Client) Cancel(id string, params *stripe.PaymentIntentCancelParams) (*stripe.PaymentIntent, error) {
	path := stripe.FormatURLPath("/v1/payment_intents/%s/cancel", id)
	paymentintent := &stripe.PaymentIntent{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, paymentintent)
	return paymentintent, err
}

// Capture is the method for the `POST /v1/payment_intents/{intent}/capture` API.
func Capture(id string, params *stripe.PaymentIntentCaptureParams) (*stripe.PaymentIntent, error) {
	return getC().Capture(id, params)
}

// Capture is the method for the `POST /v1/payment_intents/{intent}/capture` API.
func (c Client) Capture(id string, params *stripe.PaymentIntentCaptureParams) (*stripe.PaymentIntent, error) {
	path := stripe.FormatURLPath("/v1/payment_intents/%s/capture", id)
	paymentintent := &stripe.PaymentIntent{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, paymentintent)
	return paymentintent, err
}

// Confirm is the method for the `POST /v1/payment_intents/{intent}/confirm` API.
func Confirm(id string, params *stripe.PaymentIntentConfirmParams) (*stripe.PaymentIntent, error) {
	return getC().Confirm(id, params)
}

// Confirm is the method for the `POST /v1/payment_intents/{intent}/confirm` API.
func (c Client) Confirm(id string, params *stripe.PaymentIntentConfirmParams) (*stripe.PaymentIntent, error) {
	path := stripe.FormatURLPath("/v1/payment_intents/%s/confirm", id)
	paymentintent := &stripe.PaymentIntent{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, paymentintent)
	return paymentintent, err
}

// List returns a list of payment intents.
func List(params *stripe.PaymentIntentListParams) *Iter {
	return getC().List(params)
}

// List returns a list of payment intents.
func (c Client) List(listParams *stripe.PaymentIntentListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.PaymentIntentList{}
			err := c.B.CallRaw(http.MethodGet, "/v1/payment_intents", c.Key, b, p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for payment intents.
type Iter struct {
	*stripe.Iter
}

// PaymentIntent returns the payment intent which the iterator is currently pointing to.
func (i *Iter) PaymentIntent() *stripe.PaymentIntent {
	return i.Current().(*stripe.PaymentIntent)
}

// PaymentIntentList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) PaymentIntentList() *stripe.PaymentIntentList {
	return i.List().(*stripe.PaymentIntentList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
