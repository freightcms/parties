package models

import (
	commonModels "github.com/freightcms/common/models"
	locationModels "github.com/freightcms/locations/models"
)

type Party struct {
	RollupId        *string                     `json:"RollupId" bson:"RollupId"`
	PhysicalAddress locationModels.AddressModel `json:"PhysicalAddress" bson:"PhysicalAddress"`
	MailingAddress  locationModels.AddressModel `json:"MailingAddress" bson:"MailingAddress"`
	commonModels.Contact
}

type Person struct {
	Party
	FirstName string `json:"FirstName" bson:"FirstName"`
	LastName  string `json:"LastName" bson:"LastName"`
	Title     string `json:"Title" bson:"Title"`
}

type Company struct {
	Party
	TaxId    string                 `json:"TaxId" bson:"TaxId"`
	DBA      string                 `json:"DBA" bson:"DBA"`
	Contacts []commonModels.Contact `json:"Contacts" bson:"Contacts"`
}
