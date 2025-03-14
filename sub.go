//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"encoding/json"
	"github.com/stripe/stripe-go/v72/form"
)

// SubscriptionStatus is the list of allowed values for the subscription's status.
type SubscriptionStatus string

// List of values that SubscriptionStatus can take.
const (
	SubscriptionStatusActive            SubscriptionStatus = "active"
	SubscriptionStatusAll               SubscriptionStatus = "all"
	SubscriptionStatusCanceled          SubscriptionStatus = "canceled"
	SubscriptionStatusIncomplete        SubscriptionStatus = "incomplete"
	SubscriptionStatusIncompleteExpired SubscriptionStatus = "incomplete_expired"
	SubscriptionStatusPastDue           SubscriptionStatus = "past_due"
	SubscriptionStatusTrialing          SubscriptionStatus = "trialing"
	SubscriptionStatusUnpaid            SubscriptionStatus = "unpaid"
)

// SubscriptionCollectionMethod is the type of collection method for this subscription's invoices.
type SubscriptionCollectionMethod string

// List of values that SubscriptionCollectionMethod can take.
const (
	SubscriptionCollectionMethodChargeAutomatically SubscriptionCollectionMethod = "charge_automatically"
	SubscriptionCollectionMethodSendInvoice         SubscriptionCollectionMethod = "send_invoice"
)

// SubscriptionPauseCollectionBehavior is the payment collection behavior a paused subscription.
type SubscriptionPauseCollectionBehavior string

// List of values that SubscriptionPauseCollectionBehavior can take.
const (
	SubscriptionPauseCollectionBehaviorKeepAsDraft       SubscriptionPauseCollectionBehavior = "keep_as_draft"
	SubscriptionPauseCollectionBehaviorMarkUncollectible SubscriptionPauseCollectionBehavior = "mark_uncollectible"
	SubscriptionPauseCollectionBehaviorVoid              SubscriptionPauseCollectionBehavior = "void"
)

// Transaction type of the mandate.
type SubscriptionPaymentSettingsPaymentMethodOptionsACSSDebitMandateOptionsTransactionType string

// List of values that SubscriptionPaymentSettingsPaymentMethodOptionsACSSDebitMandateOptionsTransactionType can take
const (
	SubscriptionPaymentSettingsPaymentMethodOptionsACSSDebitMandateOptionsTransactionTypeBusiness SubscriptionPaymentSettingsPaymentMethodOptionsACSSDebitMandateOptionsTransactionType = "business"
	SubscriptionPaymentSettingsPaymentMethodOptionsACSSDebitMandateOptionsTransactionTypePersonal SubscriptionPaymentSettingsPaymentMethodOptionsACSSDebitMandateOptionsTransactionType = "personal"
)

// Bank account verification method.
type SubscriptionPaymentSettingsPaymentMethodOptionsACSSDebitVerificationMethod string

// List of values that SubscriptionPaymentSettingsPaymentMethodOptionsACSSDebitVerificationMethod can take
const (
	SubscriptionPaymentSettingsPaymentMethodOptionsACSSDebitVerificationMethodAutomatic     SubscriptionPaymentSettingsPaymentMethodOptionsACSSDebitVerificationMethod = "automatic"
	SubscriptionPaymentSettingsPaymentMethodOptionsACSSDebitVerificationMethodInstant       SubscriptionPaymentSettingsPaymentMethodOptionsACSSDebitVerificationMethod = "instant"
	SubscriptionPaymentSettingsPaymentMethodOptionsACSSDebitVerificationMethodMicrodeposits SubscriptionPaymentSettingsPaymentMethodOptionsACSSDebitVerificationMethod = "microdeposits"
)

// We strongly recommend that you rely on our SCA Engine to automatically prompt your customers for authentication based on risk level and [other requirements](https://stripe.com/docs/strong-customer-authentication). However, if you wish to request 3D Secure based on logic from your own fraud engine, provide this option. Read our guide on [manually requesting 3D Secure](https://stripe.com/docs/payments/3d-secure#manual-three-ds) for more information on how this configuration interacts with Radar and our SCA Engine.
type SubscriptionPaymentSettingsPaymentMethodOptionsCardRequestThreeDSecure string

// List of values that SubscriptionPaymentSettingsPaymentMethodOptionsCardRequestThreeDSecure can take
const (
	SubscriptionPaymentSettingsPaymentMethodOptionsCardRequestThreeDSecureAny       SubscriptionPaymentSettingsPaymentMethodOptionsCardRequestThreeDSecure = "any"
	SubscriptionPaymentSettingsPaymentMethodOptionsCardRequestThreeDSecureAutomatic SubscriptionPaymentSettingsPaymentMethodOptionsCardRequestThreeDSecure = "automatic"
)

// The list of payment method types to provide to every invoice created by the subscription. If not set, Stripe attempts to automatically determine the types to use by looking at the invoice's default payment method, the subscription's default payment method, the customer's default payment method, and your [invoice template settings](https://dashboard.stripe.com/settings/billing/invoice).
type SubscriptionPaymentSettingsPaymentMethodType string

// List of values that SubscriptionPaymentSettingsPaymentMethodType can take
const (
	SubscriptionPaymentSettingsPaymentMethodTypeAchCreditTransfer  SubscriptionPaymentSettingsPaymentMethodType = "ach_credit_transfer"
	SubscriptionPaymentSettingsPaymentMethodTypeAchDebit           SubscriptionPaymentSettingsPaymentMethodType = "ach_debit"
	SubscriptionPaymentSettingsPaymentMethodTypeACSSDebit          SubscriptionPaymentSettingsPaymentMethodType = "acss_debit"
	SubscriptionPaymentSettingsPaymentMethodTypeAUBECSDebit        SubscriptionPaymentSettingsPaymentMethodType = "au_becs_debit"
	SubscriptionPaymentSettingsPaymentMethodTypeBACSDebit          SubscriptionPaymentSettingsPaymentMethodType = "bacs_debit"
	SubscriptionPaymentSettingsPaymentMethodTypeBancontact         SubscriptionPaymentSettingsPaymentMethodType = "bancontact"
	SubscriptionPaymentSettingsPaymentMethodTypeBoleto             SubscriptionPaymentSettingsPaymentMethodType = "boleto"
	SubscriptionPaymentSettingsPaymentMethodTypeCard               SubscriptionPaymentSettingsPaymentMethodType = "card"
	SubscriptionPaymentSettingsPaymentMethodTypeFPX                SubscriptionPaymentSettingsPaymentMethodType = "fpx"
	SubscriptionPaymentSettingsPaymentMethodTypeGiropay            SubscriptionPaymentSettingsPaymentMethodType = "giropay"
	SubscriptionPaymentSettingsPaymentMethodTypeIdeal              SubscriptionPaymentSettingsPaymentMethodType = "ideal"
	SubscriptionPaymentSettingsPaymentMethodTypeSepaCreditTransfer SubscriptionPaymentSettingsPaymentMethodType = "sepa_credit_transfer"
	SubscriptionPaymentSettingsPaymentMethodTypeSepaDebit          SubscriptionPaymentSettingsPaymentMethodType = "sepa_debit"
	SubscriptionPaymentSettingsPaymentMethodTypeSofort             SubscriptionPaymentSettingsPaymentMethodType = "sofort"
	SubscriptionPaymentSettingsPaymentMethodTypeWechatPay          SubscriptionPaymentSettingsPaymentMethodType = "wechat_pay"
)

// SubscriptionPaymentBehavior lets you control the behavior of subscription creation in case of
// a failed payment.
type SubscriptionPaymentBehavior string

// List of values that SubscriptionPaymentBehavior can take.
const (
	SubscriptionPaymentBehaviorAllowIncomplete     SubscriptionPaymentBehavior = "allow_incomplete"
	SubscriptionPaymentBehaviorErrorIfIncomplete   SubscriptionPaymentBehavior = "error_if_incomplete"
	SubscriptionPaymentBehaviorPendingIfIncomplete SubscriptionPaymentBehavior = "pending_if_incomplete"
)

// SubscriptionProrationBehavior determines how to handle prorations when billing cycles change.
type SubscriptionProrationBehavior string

// List of values that SubscriptionProrationBehavior can take.
const (
	SubscriptionProrationBehaviorAlwaysInvoice    SubscriptionProrationBehavior = "always_invoice"
	SubscriptionProrationBehaviorCreateProrations SubscriptionProrationBehavior = "create_prorations"
	SubscriptionProrationBehaviorNone             SubscriptionProrationBehavior = "none"
)

// SubscriptionPendingInvoiceItemIntervalInterval controls the interval at which pending invoice
// items should be invoiced.
type SubscriptionPendingInvoiceItemIntervalInterval string

// List of values that SubscriptionPendingInvoiceItemIntervalInterval can take.
const (
	SubscriptionPendingInvoiceItemIntervalIntervalDay   SubscriptionPendingInvoiceItemIntervalInterval = "day"
	SubscriptionPendingInvoiceItemIntervalIntervalMonth SubscriptionPendingInvoiceItemIntervalInterval = "month"
	SubscriptionPendingInvoiceItemIntervalIntervalWeek  SubscriptionPendingInvoiceItemIntervalInterval = "week"
	SubscriptionPendingInvoiceItemIntervalIntervalYear  SubscriptionPendingInvoiceItemIntervalInterval = "year"
)

// SubscriptionAddInvoiceItemParams is a structure representing the parameters allowed to control
// the invoice items to add at to a subscription's first invoice.
type SubscriptionAddInvoiceItemParams struct {
	Price     *string                     `form:"price"`
	PriceData *InvoiceItemPriceDataParams `form:"price_data"`
	Quantity  *int64                      `form:"quantity"`
	TaxRates  []*string                   `form:"tax_rates"`
}

// SubscriptionPauseCollectionParams is the set of parameters allowed for the pause_collection hash.
type SubscriptionPauseCollectionParams struct {
	Behavior  *string `form:"behavior"`
	ResumesAt *int64  `form:"resumes_at"`
}

// Additional fields for Mandate creation
type SubscriptionPaymentSettingsPaymentMethodOptionsACSSDebitMandateOptionsParams struct {
	TransactionType *string `form:"transaction_type"`
}

// This sub-hash contains details about the Canadian pre-authorized debit payment method options to pass to the invoice's PaymentIntent.
type SubscriptionPaymentSettingsPaymentMethodOptionsACSSDebitParams struct {
	MandateOptions     *SubscriptionPaymentSettingsPaymentMethodOptionsACSSDebitMandateOptionsParams `form:"mandate_options"`
	VerificationMethod *string                                                                       `form:"verification_method"`
}

// This sub-hash contains details about the Bancontact payment method options to pass to the invoice's PaymentIntent.
type SubscriptionPaymentSettingsPaymentMethodOptionsBancontactParams struct {
	PreferredLanguage *string `form:"preferred_language"`
}

// This sub-hash contains details about the Card payment method options to pass to the invoice's PaymentIntent.
type SubscriptionPaymentSettingsPaymentMethodOptionsCardParams struct {
	RequestThreeDSecure *string `form:"request_three_d_secure"`
}

// Payment-method-specific configuration to provide to invoices created by the subscription.
type SubscriptionPaymentSettingsPaymentMethodOptionsParams struct {
	ACSSDebit  *SubscriptionPaymentSettingsPaymentMethodOptionsACSSDebitParams  `form:"acss_debit"`
	Bancontact *SubscriptionPaymentSettingsPaymentMethodOptionsBancontactParams `form:"bancontact"`
	Card       *SubscriptionPaymentSettingsPaymentMethodOptionsCardParams       `form:"card"`
}

// Payment settings to pass to invoices created by the subscription.
type SubscriptionPaymentSettingsParams struct {
	PaymentMethodOptions *SubscriptionPaymentSettingsPaymentMethodOptionsParams `form:"payment_method_options"`
	PaymentMethodTypes   []*string                                              `form:"payment_method_types"`
}

// SubscriptionPendingInvoiceItemIntervalParams is the set of parameters allowed for the transfer_data hash.
type SubscriptionPendingInvoiceItemIntervalParams struct {
	Interval      *string `form:"interval"`
	IntervalCount *int64  `form:"interval_count"`
}

// SubscriptionTransferDataParams is the set of parameters allowed for the transfer_data hash.
type SubscriptionTransferDataParams struct {
	AmountPercent *float64 `form:"amount_percent"`
	Destination   *string  `form:"destination"`
}

// SubscriptionParams is the set of parameters that can be used when creating or updating a subscription.
// For more details see https://stripe.com/docs/api#create_subscription and https://stripe.com/docs/api#update_subscription.
type SubscriptionParams struct {
	Params                      `form:"*"`
	AddInvoiceItems             []*SubscriptionAddInvoiceItemParams           `form:"add_invoice_items"`
	ApplicationFeePercent       *float64                                      `form:"application_fee_percent"`
	AutomaticTax                *SubscriptionAutomaticTaxParams               `form:"automatic_tax"`
	BackdateStartDate           *int64                                        `form:"backdate_start_date"`
	BillingCycleAnchor          *int64                                        `form:"billing_cycle_anchor"`
	BillingCycleAnchorNow       *bool                                         `form:"-"` // See custom AppendTo
	BillingCycleAnchorUnchanged *bool                                         `form:"-"` // See custom AppendTo
	BillingThresholds           *SubscriptionBillingThresholdsParams          `form:"billing_thresholds"`
	CancelAt                    *int64                                        `form:"cancel_at"`
	CancelAtPeriodEnd           *bool                                         `form:"cancel_at_period_end"`
	Card                        *CardParams                                   `form:"card"`
	CollectionMethod            *string                                       `form:"collection_method"`
	Coupon                      *string                                       `form:"coupon"`
	Customer                    *string                                       `form:"customer"`
	DaysUntilDue                *int64                                        `form:"days_until_due"`
	DefaultPaymentMethod        *string                                       `form:"default_payment_method"`
	DefaultSource               *string                                       `form:"default_source"`
	DefaultTaxRates             []*string                                     `form:"default_tax_rates"`
	Items                       []*SubscriptionItemsParams                    `form:"items"`
	OffSession                  *bool                                         `form:"off_session"`
	OnBehalfOf                  *string                                       `form:"on_behalf_of"`
	PauseCollection             *SubscriptionPauseCollectionParams            `form:"pause_collection"`
	PaymentBehavior             *string                                       `form:"payment_behavior"`
	PaymentSettings             *SubscriptionPaymentSettingsParams            `form:"payment_settings"`
	PendingInvoiceItemInterval  *SubscriptionPendingInvoiceItemIntervalParams `form:"pending_invoice_item_interval"`
	Plan                        *string                                       `form:"plan"`
	PromotionCode               *string                                       `form:"promotion_code"`
	ProrationBehavior           *string                                       `form:"proration_behavior"`
	ProrationDate               *int64                                        `form:"proration_date"`
	Quantity                    *int64                                        `form:"quantity"`
	TransferData                *SubscriptionTransferDataParams               `form:"transfer_data"`
	TrialEnd                    *int64                                        `form:"trial_end"`
	TrialEndNow                 *bool                                         `form:"-"` // See custom AppendTo
	TrialFromPlan               *bool                                         `form:"trial_from_plan"`
	TrialPeriodDays             *int64                                        `form:"trial_period_days"`
}

// AppendTo implements custom encoding logic for SubscriptionParams.
func (s *SubscriptionParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(s.BillingCycleAnchorNow) {
		body.Add(form.FormatKey(append(keyParts, "billing_cycle_anchor")), "now")
	}
	if BoolValue(s.BillingCycleAnchorUnchanged) {
		body.Add(form.FormatKey(append(keyParts, "billing_cycle_anchor")), "unchanged")
	}
	if BoolValue(s.TrialEndNow) {
		body.Add(form.FormatKey(append(keyParts, "trial_end")), "now")
	}
}

// Automatic tax settings for this subscription.
type SubscriptionAutomaticTaxParams struct {
	Enabled *bool `form:"enabled"`
}

// SubscriptionBillingThresholdsParams is a structure representing the parameters allowed to control
// billing thresholds for a subscription.
type SubscriptionBillingThresholdsParams struct {
	AmountGTE               *int64 `form:"amount_gte"`
	ResetBillingCycleAnchor *bool  `form:"reset_billing_cycle_anchor"`
}

// SubscriptionCancelParams is the set of parameters that can be used when canceling a subscription.
// For more details see https://stripe.com/docs/api#cancel_subscription
type SubscriptionCancelParams struct {
	Params     `form:"*"`
	InvoiceNow *bool `form:"invoice_now"`
	Prorate    *bool `form:"prorate"`
}

// SubscriptionItemsParams is the set of parameters that can be used when creating or updating a subscription item on a subscription
// For more details see https://stripe.com/docs/api#create_subscription and https://stripe.com/docs/api#update_subscription.
type SubscriptionItemsParams struct {
	Params            `form:"*"`
	BillingThresholds *SubscriptionItemBillingThresholdsParams `form:"billing_thresholds"`
	ClearUsage        *bool                                    `form:"clear_usage"`
	Deleted           *bool                                    `form:"deleted"`
	ID                *string                                  `form:"id"`
	Metadata          map[string]string                        `form:"metadata"`
	Plan              *string                                  `form:"plan"`
	Price             *string                                  `form:"price"`
	PriceData         *SubscriptionItemPriceDataParams         `form:"price_data"`
	Quantity          *int64                                   `form:"quantity"`
	TaxRates          []*string                                `form:"tax_rates"`
}

// SubscriptionListParams is the set of parameters that can be used when listing active subscriptions.
// For more details see https://stripe.com/docs/api#list_subscriptions.
type SubscriptionListParams struct {
	ListParams              `form:"*"`
	CollectionMethod        *string           `form:"collection_method"`
	Created                 int64             `form:"created"`
	CreatedRange            *RangeQueryParams `form:"created"`
	CurrentPeriodEnd        *int64            `form:"current_period_end"`
	CurrentPeriodEndRange   *RangeQueryParams `form:"current_period_end"`
	CurrentPeriodStart      *int64            `form:"current_period_start"`
	CurrentPeriodStartRange *RangeQueryParams `form:"current_period_start"`
	Customer                string            `form:"customer"`
	Plan                    string            `form:"plan"`
	Price                   string            `form:"price"`
	Status                  string            `form:"status"`
}

// SubscriptionPauseCollection if specified, payment collection for this subscription will be paused.
type SubscriptionPauseCollection struct {
	Behavior  SubscriptionPauseCollectionBehavior `json:"behavior"`
	ResumesAt int64                               `json:"resumes_at"`
}
type SubscriptionPaymentSettingsPaymentMethodOptionsACSSDebitMandateOptions struct {
	TransactionType SubscriptionPaymentSettingsPaymentMethodOptionsACSSDebitMandateOptionsTransactionType `json:"transaction_type"`
}

// This sub-hash contains details about the Canadian pre-authorized debit payment method options to pass to invoices created by the subscription.
type SubscriptionPaymentSettingsPaymentMethodOptionsACSSDebit struct {
	MandateOptions     *SubscriptionPaymentSettingsPaymentMethodOptionsACSSDebitMandateOptions    `json:"mandate_options"`
	VerificationMethod SubscriptionPaymentSettingsPaymentMethodOptionsACSSDebitVerificationMethod `json:"verification_method"`
}

// This sub-hash contains details about the Bancontact payment method options to pass to invoices created by the subscription.
type SubscriptionPaymentSettingsPaymentMethodOptionsBancontact struct {
	PreferredLanguage string `json:"preferred_language"`
}

// This sub-hash contains details about the Card payment method options to pass to invoices created by the subscription.
type SubscriptionPaymentSettingsPaymentMethodOptionsCard struct {
	RequestThreeDSecure SubscriptionPaymentSettingsPaymentMethodOptionsCardRequestThreeDSecure `json:"request_three_d_secure"`
}

// Payment-method-specific configuration to provide to invoices created by the subscription.
type SubscriptionPaymentSettingsPaymentMethodOptions struct {
	ACSSDebit  *SubscriptionPaymentSettingsPaymentMethodOptionsACSSDebit  `json:"acss_debit"`
	Bancontact *SubscriptionPaymentSettingsPaymentMethodOptionsBancontact `json:"bancontact"`
	Card       *SubscriptionPaymentSettingsPaymentMethodOptionsCard       `json:"card"`
}

// Payment settings passed on to invoices created by the subscription.
type SubscriptionPaymentSettings struct {
	PaymentMethodOptions *SubscriptionPaymentSettingsPaymentMethodOptions `json:"payment_method_options"`
	PaymentMethodTypes   []SubscriptionPaymentSettingsPaymentMethodType   `json:"payment_method_types"`
}

// SubscriptionPendingInvoiceItemInterval represents the interval at which to invoice pending invoice
// items.
type SubscriptionPendingInvoiceItemInterval struct {
	Interval      SubscriptionPendingInvoiceItemIntervalInterval `json:"interval"`
	IntervalCount int64                                          `json:"interval_count"`
}

// SubscriptionPendingUpdate represents deferred changes that will be applied when latest invoice is paid.
type SubscriptionPendingUpdate struct {
	BillingCycleAnchor int64               `json:"billing_cycle_anchor"`
	ExpiresAt          int64               `json:"expires_at"`
	SubscriptionItems  []*SubscriptionItem `json:"subscription_items"`
	TrialEnd           int64               `json:"trial_end"`
	TrialFromPlan      bool                `json:"trial_from_plan"`
}

// SubscriptionTransferData represents the information for the transfer_data associated with a subscription.
type SubscriptionTransferData struct {
	AmountPercent float64  `json:"amount_percent"`
	Destination   *Account `json:"destination"`
}

// Subscription is the resource representing a Stripe subscription.
// For more details see https://stripe.com/docs/api#subscriptions.
type Subscription struct {
	APIResource
	ApplicationFeePercent         float64                                `json:"application_fee_percent"`
	AutomaticTax                  *SubscriptionAutomaticTax              `json:"automatic_tax"`
	BillingCycleAnchor            int64                                  `json:"billing_cycle_anchor"`
	BillingThresholds             *SubscriptionBillingThresholds         `json:"billing_thresholds"`
	CancelAt                      int64                                  `json:"cancel_at"`
	CancelAtPeriodEnd             bool                                   `json:"cancel_at_period_end"`
	CanceledAt                    int64                                  `json:"canceled_at"`
	CollectionMethod              SubscriptionCollectionMethod           `json:"collection_method"`
	Created                       int64                                  `json:"created"`
	CurrentPeriodEnd              int64                                  `json:"current_period_end"`
	CurrentPeriodStart            int64                                  `json:"current_period_start"`
	Customer                      *Customer                              `json:"customer"`
	DaysUntilDue                  int64                                  `json:"days_until_due"`
	DefaultPaymentMethod          *PaymentMethod                         `json:"default_payment_method"`
	DefaultSource                 *PaymentSource                         `json:"default_source"`
	DefaultTaxRates               []*TaxRate                             `json:"default_tax_rates"`
	Discount                      *Discount                              `json:"discount"`
	EndedAt                       int64                                  `json:"ended_at"`
	ID                            string                                 `json:"id"`
	Items                         *SubscriptionItemList                  `json:"items"`
	LatestInvoice                 *Invoice                               `json:"latest_invoice"`
	Livemode                      bool                                   `json:"livemode"`
	Metadata                      map[string]string                      `json:"metadata"`
	NextPendingInvoiceItemInvoice int64                                  `json:"next_pending_invoice_item_invoice"`
	Object                        string                                 `json:"object"`
	OnBehalfOf                    *Account                               `json:"on_behalf_of"`
	PauseCollection               SubscriptionPauseCollection            `json:"pause_collection"`
	PaymentSettings               *SubscriptionPaymentSettings           `json:"payment_settings"`
	PendingInvoiceItemInterval    SubscriptionPendingInvoiceItemInterval `json:"pending_invoice_item_interval"`
	PendingSetupIntent            *SetupIntent                           `json:"pending_setup_intent"`
	PendingUpdate                 *SubscriptionPendingUpdate             `json:"pending_update"`
	Plan                          *Plan                                  `json:"plan"`
	Quantity                      int64                                  `json:"quantity"`
	Schedule                      *SubscriptionSchedule                  `json:"schedule"`
	StartDate                     int64                                  `json:"start_date"`
	Status                        SubscriptionStatus                     `json:"status"`
	TransferData                  *SubscriptionTransferData              `json:"transfer_data"`
	TrialEnd                      int64                                  `json:"trial_end"`
	TrialStart                    int64                                  `json:"trial_start"`
}
type SubscriptionAutomaticTax struct {
	Enabled bool `json:"enabled"`
}

// SubscriptionBillingThresholds is a structure representing the billing thresholds for a subscription.
type SubscriptionBillingThresholds struct {
	AmountGTE               int64 `json:"amount_gte"`
	ResetBillingCycleAnchor bool  `json:"reset_billing_cycle_anchor"`
}

// SubscriptionList is a list object for subscriptions.
type SubscriptionList struct {
	APIResource
	ListMeta
	Data []*Subscription `json:"data"`
}

// UnmarshalJSON handles deserialization of a Subscription.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (s *Subscription) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		s.ID = id
		return nil
	}

	type subscription Subscription
	var v subscription
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*s = Subscription(v)
	return nil
}
