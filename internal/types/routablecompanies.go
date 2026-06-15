package types

import "time"

type RoutableListCompaniesResponse struct {
	Object string          `json:"object"`
	Results []RoutableCompany `json:"results"`
	Links   RoutableLinks     `json:"links"`
}

type RoutableCompany struct {
	Object          string            `json:"object"`
	ID              string            `json:"id"`
	BusinessName    string            `json:"business_name"`
	CollectTaxForm  bool              `json:"collect_tax_form"`
	Contacts        RoutableListContacts `json:"contacts"`
	CountryCode     string            `json:"country_code"`
	CreatedAt       time.Time         `json:"created_at"`
	DisplayName     string            `json:"display_name"`
	IsCustomer      bool              `json:"is_customer"`
	IsVendor        bool              `json:"is_vendor"`
	Ledger          RoutableLedger    `json:"ledger"`
	RiskSummary     string            `json:"risk_summary"`
	Status          string            `json:"status"`
	TaxFormStatus   string            `json:"tax_form_status"`
	TaxFormType     string            `json:"tax_form_type"`
	Type            string            `json:"type"`
	ExternalID      string            `json:"external_id,omitempty"`
	Links           RoutableLinks     `json:"links"`
}

type RoutableListContacts struct {
	Object  string           `json:"object"`
	Results []RoutableContact `json:"results"`
	Links   RoutableLinks     `json:"links"`
}

type RoutableContact struct {
	Object                              string        `json:"object"`
	ID                                  string        `json:"id"`
	DefaultContactForCompanyManagement  string        `json:"default_contact_for_company_management"`
	DefaultContactForPayableAndReceivable string      `json:"default_contact_for_payable_and_receivable"`
	Email                               string        `json:"email"`
	FirstName                           string        `json:"first_name"`
	IsArchived                          bool          `json:"is_archived"`
	LastEmailDeliveryProblem            *string       `json:"last_email_delivery_problem,omitempty"`
	LastName                            string        `json:"last_name"`
	PhoneNumber                         string        `json:"phone_number"`
	PhoneNumberCountry                  string        `json:"phone_number_country"`
	Links                               RoutableLinks `json:"links"`
}

type RoutableLedger struct {
	CustomerID *string `json:"customer_id,omitempty"`
	VendorID   *string `json:"vendor_id,omitempty"`
}

// CompanyStatus represents the status of a company in Routable.
type CompanyStatus string

const (
	CompanyStatusAdded    CompanyStatus = "added"
	CompanyStatusInvited  CompanyStatus = "invited"
	CompanyStatusAccepted CompanyStatus = "accepted"
)

// CompanyTaxFormStatus represents the tax form status of a company.
type CompanyTaxFormStatus string

const (
	CompanyTaxFormStatusComplete CompanyTaxFormStatus = "complete"
)


