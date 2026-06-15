package types

type RoutableLinks struct {
	Self             string `json:"self,omitempty"`
	Next             string `json:"next,omitempty"`
	Prev             string `json:"prev,omitempty"`
	Company          string `json:"company,omitempty"`
	ComplianceChecks string `json:"compliance_checks,omitempty"`
	Contacts         string `json:"contacts,omitempty"`
	PaymentMethods   string `json:"payment_methods,omitempty"`
	Payables         string `json:"payables,omitempty"`
	Receivables      string `json:"receivables,omitempty"`
	TaxForms         string `json:"tax_forms,omitempty"`
}
