package service

import "html/template"

const (
	RegisterMessageSubject    = "Welcome to Singing City"
	ApplicationMessageSubject = "NEUE BEWERBUNG für eines deiner SingingCity Zeitfenster"
	ConfirmedMessageSubject   = "A Timeslot has been confirmed"

	RegisterMessage = `<h1>Welcome {{ .Username }}</h1>
	<p>You have signed up to the <a href="https://singingcity.songbirdfestival.ch">Singing City</a> Platform. You are ready to get started.</p>`

	ApplicationMessage = `
	<h1>Hallo {{ .Username }}</h1>
	<p>Du hast eine Bewerbung auf dein offenes SingingCity Zeitfenster <strong>{{ .Time }}</strong> am <strong>{{ .Date }}</strong> erhalten.</p>
	<p>
	Eben hat sich <em>{{ .Artist }}</em> für einen Auftritt bei dir beworben.
	Alle Informationen zu <em>{{ .Artist }}</em> findest du in deinem SingingCity Profil.
	Schaue dir die Bewerbung an und sage zu oder ab. Bestätigst du den Auftritt, so erhält <em>{{ .Artist }}</em> ein Bestätigungsmail, das zugleich als Vertrag zwischen euch dient.
	Solltest du weitere Fragen haben, kannst du <em>{{ .Artist }}</em> nach der Zusage direkt via der neu angezeigten E-Mail-Adresse kontaktieren.
	</p>
	<a href="https://singingcity.songbirdfestival.ch/dashboard">Zu meinem SingingCity Profil</a>
	<p>Herzlichst,<br>
	Dein SingingCity Team</p>

	<hr>

	<h1>Hello {{ .Username }}</h1>
	<p>You have received an application for your open SingingCity timeslot <strong>{{ .Time }}</strong> on <strong>{{ .Date }}</strong>.</p>
	<p>
	<em>{{ .Artist }}</em> has just applied for a performance at your venue.
	You can find all the information about <em>{{ .Artist }}</em> in your SingingCity profile.
	Take a look at the application and accept or decline. If you confirm the performance, <em>{{ .Artist }}</em> will receive a confirmation email, which also serves as a contract between the two of you.
	If you have any further questions, you can contact <em>{{ .Artist }}</em> directly via the newly displayed e-mail address after you have accepted.
	</p>
	<a href="https://singingcity.songbirdfestival.ch/dashboard">To my SingingCity profile</a>
	<p>Sincerely,<br>
	Your SingingCity Team</p>`

	ConfirmedMessage = ``
)

var (
	registerMessageTmpl    = template.Must(template.New("register").Parse(RegisterMessage))
	applicationMessageTmpl = template.Must(template.New("application").Parse(ApplicationMessage))
	confirmedMessageTmpl   = template.Must(template.New("confirmed").Parse(ConfirmedMessage))
)

type ApplicationMessageParams struct {
	Username string
	Artist   string
	Time     string
	Date     string
}
