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

type OrderDetail struct {
	CustomerId int32 `json:"Customer_id,omitempty"`

	AncestorId int32 `json:"Ancestor_id,omitempty"`

	ParentId int32 `json:"Parent_id,omitempty"`

	ChildId *int32 `json:"Child_id,omitempty"`

	CrmRefunded bool `json:"Crm_refunded,omitempty"`

	CrmBlacklisted bool `json:"Crm_blacklisted,omitempty"`

	CrmStoppedRecurring bool `json:"Crm_stopped_recurring,omitempty"`

	IpAddress string `json:"Ip_address,omitempty"`

	Affiliate int32 `json:"Affiliate,omitempty"`

	SubAffiliate string `json:"Sub_affiliate,omitempty"`

	CampaignId int32 `json:"Campaign_id,omitempty"`

	ClickId int32 `json:"Click_id,omitempty"`

	Resolved bool `json:"Resolved,omitempty"`

	OnHoldBy string `json:"On_hold_by,omitempty"`

	OnHoldDate string `json:"On_hold_date,omitempty"`

	OrderConfirmed bool `json:"Order_confirmed,omitempty"`

	OrderConfirmedDate string `json:"Order_confirmed_date,omitempty"`

	OrderDate string `json:"Order_date,omitempty"`

	DetailCreated time.Time `json:"Detail_created,omitempty"`

	DetailUpdated time.Time `json:"Detail_updated,omitempty"`

	OrderContactId Contact `json:"Order_contact_id,omitempty"`
}

// AssertOrderDetailRequired checks if the required fields are not zero-ed
func AssertOrderDetailRequired(obj OrderDetail) error {
	if err := AssertContactRequired(obj.OrderContactId); err != nil {
		return err
	}
	return nil
}

// AssertOrderDetailConstraints checks if the values respects the defined constraints
func AssertOrderDetailConstraints(obj OrderDetail) error {
	return nil
}
