package main

import (
	"crypto/tls"
	"fmt"

	gomail "gopkg.in/mail.v2"
)

func SendMailRegister(to string) error {
	var err error
	var bodyRegister string = `
Hello,

You ask for a register on mangi.
Please click on this link to validate your register.

mangi
`

	//createMailContent()
	m := gomail.NewMessage()

	// Set E-Mail sender
	m.SetHeader("From", "no-reply@mangi.com")

	// Set E-Mail receivers
	m.SetHeader("To", to)

	// Set E-Mail subject
	subject := "Mangi registration"
	m.SetHeader("Subject", subject)

	// Set E-Mail body. You can set plain text or html with text/html
	m.SetBody("text/plain", bodyRegister)

	// Settings for SMTP server
	d := gomail.NewDialer("127.0.0.1", 1025, "no-reply@mangi.com", "")

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Now send E-Mail
	if err := d.DialAndSend(m); err != nil {
		// TODO i need to update the insecurePlainAuth
		// local mailpit + golang accept plainAuth but stangins and production
		// aren't on localhost. Golang doesn't allow the insecurePlainAuth.
		fmt.Println(err)
		//return err
	}

	return err
}

func SendMailPassword(to string) error {
	var err error
	var bodyChangePassword string = `
Hello,

You ask to change your password on mangi.
Please click on this link to change your password.

mangi

If you're not the owner of this changement please contact mangi's administrator
and go to your phone app et change your password.
`

	//createMailContent()
	m := gomail.NewMessage()

	// Set E-Mail sender
	m.SetHeader("From", "no-reply@mangi.com")

	// Set E-Mail receivers
	m.SetHeader("To", to)

	// Set E-Mail subject
	subject := "Mangi change your paswword"
	m.SetHeader("Subject", subject)

	// Set E-Mail body. You can set plain text or html with text/html
	m.SetBody("text/plain", bodyChangePassword)

	// Settings for SMTP server
	d := gomail.NewDialer("127.0.0.1", 1025, "no-reply@mangi.com", "")

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Now send E-Mail
	if err := d.DialAndSend(m); err != nil {
		// TODO i need to update the insecurePlainAuth
		// local mailpit + golang accept plainAuth but stangins and production
		// aren't on localhost. Golang doesn't allow the insecurePlainAuth.
		fmt.Println(err)
		//return err
	}

	return err
}

func SendMailEmail(to string) error {
	var err error
	var bodyChangeEmail string = `
Hello,

You ask to change your email on mangi.
Please click on this link to confirme your new email.

mangi

If you're not the owner of this changement please contact mangi's administrator
and go to your phone app et change your password.
`

	//createMailContent()
	m := gomail.NewMessage()

	// Set E-Mail sender
	m.SetHeader("From", "no-reply@mangi.com")

	// Set E-Mail receivers
	m.SetHeader("To", to)

	// Set E-Mail subject
	subject := "Mangi change your email"
	m.SetHeader("Subject", subject)

	// Set E-Mail body. You can set plain text or html with text/html
	m.SetBody("text/plain", bodyChangeEmail)

	// Settings for SMTP server
	d := gomail.NewDialer("127.0.0.1", 1025, "no-reply@mangi.com", "")

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Now send E-Mail
	if err := d.DialAndSend(m); err != nil {
		// TODO i need to update the insecurePlainAuth
		// local mailpit + golang accept plainAuth but stangins and production
		// aren't on localhost. Golang doesn't allow the insecurePlainAuth.
		fmt.Println(err)
		//return err
	}

	return err
}

func SendMailRgpd(to, attach string) error {
	var err error
	var bodyRgpdData string = `
Hello,

You've ask for all your data user.
You'll find in attachment a json file with all your user's datas.
Please the attachment.

mangi
`

	//createMailContent()
	m := gomail.NewMessage()

	// Set E-Mail sender
	m.SetHeader("From", "no-reply@mangi.com")

	// Set E-Mail receivers
	m.SetHeader("To", to)

	// Set E-Mail subject
	subject := "Mangi user's datas"
	m.SetHeader("Subject", subject)

	// Set E-Mail body. You can set plain text or html with text/html
	m.SetBody("text/plain", bodyRgpdData)

	// Set E-Mail attachment as an option
	m.Attach(attach)

	// Settings for SMTP server
	d := gomail.NewDialer("127.0.0.1", 1025, "no-reply@mangi.com", "")

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Now send E-Mail
	if err := d.DialAndSend(m); err != nil {
		// TODO i need to update the insecurePlainAuth
		// local mailpit + golang accept plainAuth but stangins and production
		// aren't on localhost. Golang doesn't allow the insecurePlainAuth.
		fmt.Println(err)
		//return err
	}

	return err
}

func SendMailInvitation(to, homeName, owner string, homeID int64) error {
	var err error
	var bodyChangeEmail string = fmt.Sprintf(`
Hello,

You've received an invitation to join "%+s" house
from %+v to organize meals and shopping list.
Please follow the link and register if you're not yet.
%+v

mangi

`, homeName, owner, homeID)

	//createMailContent()
	m := gomail.NewMessage()

	// Set E-Mail sender
	m.SetHeader("From", "no-reply@mangi.com")

	// Set E-Mail receivers
	m.SetHeader("To", to)

	// Set E-Mail subject
	subject := "Mangi invitation to a house"
	m.SetHeader("Subject", subject)

	// Set E-Mail body. You can set plain text or html with text/html
	m.SetBody("text/plain", bodyChangeEmail)

	// Settings for SMTP server
	d := gomail.NewDialer("127.0.0.1", 1025, "no-reply@mangi.com", "")

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Now send E-Mail
	if err := d.DialAndSend(m); err != nil {
		// TODO i need to update the insecurePlainAuth
		// local mailpit + golang accept plainAuth but stangins and production
		// aren't on localhost. Golang doesn't allow the insecurePlainAuth.
		fmt.Println(err)
		//return err
	}

	return err
}
