package controllers

import (
	models "belajar-package/mmodels"
	"fmt"
)

type ContactController interface {
	DisplayContact()
}

type Contact struct {
	Contact models.Contact
}

func NewContactController(contact *models.Contact) ContactController {
	return &Contact{
		Contact: *contact,
	}
}

func (c *Contact) DisplayContact() {
	fmt.Println("Called from contact controller ...")
	fmt.Println("My Phone", c.Contact.Phone)
	fmt.Println("My Instagram", c.Contact.Instagram)
}
