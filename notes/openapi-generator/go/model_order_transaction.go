/*
 * Order Uploader
 *
 * Takes cleaned and validated order data records as input and enters the record into the database tables
 *
 * API version: 0.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"time"
)

type OrderTransaction struct {
	AvsResponse string `json:"Avs_response,omitempty"`

	CvvResponse string `json:"Cvv_response,omitempty"`

	ProcessorId string `json:"Processor_id,omitempty"`

	GatewayId int32 `json:"Gateway_id,omitempty"`

	TransactionId int32 `json:"Transaction_id,omitempty"`

	AuthId string `json:"Auth_id,omitempty"`

	PanLocked bool `json:"Pan_locked,omitempty"`

	PreserveGateway bool `json:"Preserve_gateway,omitempty"`

	IsRecurring bool `json:"Is_recurring,omitempty"`

	IsChargeback bool `json:"Is_chargeback,omitempty"`

	IsFraud bool `json:"Is_fraud,omitempty"`

	IsRma bool `json:"Is_rma,omitempty"`

	RmaNumber string `json:"Rma_number,omitempty"`

	RmaReason string `json:"Rma_reason,omitempty"`

	DeclineReason string `json:"Decline_reason,omitempty"`

	TransactionDate time.Time `json:"Transaction_date,omitempty"`

	PaymentId Payment `json:"Payment_id,omitempty"`
}

// AssertOrderTransactionRequired checks if the required fields are not zero-ed
func AssertOrderTransactionRequired(obj OrderTransaction) error {
	if err := AssertPaymentRequired(obj.PaymentId); err != nil {
		return err
	}
	return nil
}

// AssertOrderTransactionConstraints checks if the values respects the defined constraints
func AssertOrderTransactionConstraints(obj OrderTransaction) error {
	return nil
}
