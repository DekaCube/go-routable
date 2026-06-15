package routableclient

import (
	"fmt"
	"net/url"

	"github.com/DekaCube/go-routable/internal/types"
)

// ListCompaniesOption configures the ListCompanies request.
type ListCompaniesOption func(*url.Values)

// ListCompaniesPage sets the page parameter.
func ListCompaniesPage(page int) ListCompaniesOption {
	return func(v *url.Values) {
		v.Set("page", fmt.Sprintf("%d", page))
	}
}

// ListCompaniesPageSize sets the page_size parameter.
func ListCompaniesPageSize(pageSize int) ListCompaniesOption {
	return func(v *url.Values) {
		v.Set("page_size", fmt.Sprintf("%d", pageSize))
	}
}

// ListCompaniesExternalID sets the external_id parameter.
func ListCompaniesExternalID(externalID string) ListCompaniesOption {
	return func(v *url.Values) {
		v.Set("external_id", externalID)
	}
}

// ListCompaniesSearch sets the search parameter.
func ListCompaniesSearch(search string) ListCompaniesOption {
	return func(v *url.Values) {
		v.Set("search", search)
	}
}

// ListCompaniesStatus sets the status parameter.
func ListCompaniesStatus(status types.CompanyStatus) ListCompaniesOption {
	return func(v *url.Values) {
		v.Set("status", string(status))
	}
}

// ListCompaniesVendorID sets the vendor_id parameter.
func ListCompaniesVendorID(vendorID string) ListCompaniesOption {
	return func(v *url.Values) {
		v.Set("vendor_id", vendorID)
	}
}

// ListCompaniesTaxFormStatus sets the tax_form_status parameter.
func ListCompaniesTaxFormStatus(status types.CompanyTaxFormStatus) ListCompaniesOption {
	return func(v *url.Values) {
		v.Set("tax_form_status", string(status))
	}
}

// ListCompaniesCustomerID sets the customer_id parameter.
func ListCompaniesCustomerID(customerID string) ListCompaniesOption {
	return func(v *url.Values) {
		v.Set("customer_id", customerID)
	}
}
