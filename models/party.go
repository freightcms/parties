package models

import (
	commonModels "github.com/freightcms/common/models"
	locationModels "github.com/freightcms/locations/models"
)

type Party struct {
	RollupId        *string                     `json:"rollup_id,omitempty" bson:"rollup_id,omitempty"`
	PhysicalAddress locationModels.AddressModel `json:"physical_address,omitempty" bson:"physical_address,omitempty"`
	MailingAddress  locationModels.AddressModel `json:"mailing_address,omitempty" bson:"mailing_address,omitempty"`
	commonModels.Contact
}

type Person struct {
	Party
	FirstName string `json:"first_name,omitempty" bson:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty" bson:"last_name,omitempty"`
	Title     string `json:"title,omitempty" bson:"title,omitempty"`
}

type Company struct {
	Party
	TaxId    *string                `json:"tax_id,omitempty" bson:"tax_id,omitempty"`
	DBA      string                 `json:"dba,omitempty" bson:"dba,omitempty"`
	Contacts []commonModels.Contact `json:"contacts,omitempty" bson:"contacts,omitempty"`
}
