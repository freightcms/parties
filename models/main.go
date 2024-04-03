package models

type Party struct {
	RollupId        *string
	PhysicalAddress AddressModel
	MailingAddress  AddressModel
	ContactInfo     Contact
}

type Person struct {
	Party
}

type Company struct {
	Party
	TaxId    *string
	DBA      string
	Contacts []Contact
}
