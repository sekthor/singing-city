package service

import "html/template"

const (
	RegisterMessageSubject    = "Welcome to Singing City"
	ApplicationMessageSubject = "Someone has applied for one of your timeslots"
	ConfirmedMessageSubject   = "A Timeslot has been confirmed"

	RegisterMessage = `<h1>Welcome {{ .Username }}</h1>
	<p>You have signed up to the <a href="https://singingcity.songbirdfestival.ch">Singing City</a> Platform. You are ready to get started.</p>`

	ApplicationMessage = ``

	ConfirmedMessage = ``
)

var (
	registerMessageTmpl    = template.Must(template.New("register").Parse(RegisterMessage))
	applicationMessageTmpl = template.Must(template.New("application").Parse(RegisterMessage))
	confirmedMessageTmpl   = template.Must(template.New("confirmed").Parse(RegisterMessage))
)
