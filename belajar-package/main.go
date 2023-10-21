package main

import (
	"belajar-package/controllers"
	models "belajar-package/mmodels"
)

func main() {
	userModel := models.NewUser("User 1")
	contactModel := models.NewContact("0831243254", "@mutiara.pertiwi")

	user := controllers.NewUserController(*&userModel)
	contact := controllers.NewContactController(contactModel)

	user.SayHello()
	user.SetContact(contactModel)

	contact.DisplayContact()
	user.DisplayContact()
}
