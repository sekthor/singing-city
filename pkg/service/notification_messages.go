package service

import "html/template"

const (
	RegisterMessageSubject    = "Welcome to Singing City"
	ApplicationMessageSubject = "NEUE BEWERBUNG für eines deiner SingingCity Zeitfenster"
	ConfirmedMessageSubject   = "BESTÄTIGUNG deines Auftrittes"
	RejectedMessageSubject    = "ABSAGE für deine Bewerbung bei"
	PasswordResetSubject      = "Passwort zurücksetzen"

	RegisterMessage = `
	<h1>Willkommen {{ .Username }}</h1>
	<p>
	Du hast dich für <a href="{{ .BaseUrl }}">Singing City</a> registriert.
	</p>

	<hr>	

	<h1>Welcome {{ .Username }}</h1>
	<p>
	You have signed up to the <a href="{{ .BaseUrl }}">Singing City</a> Platform.
	You are ready to get started.
	</p>`

	ApplicationMessage = `
	<h1>Hallo {{ .Username }}</h1>
	<p>Du hast eine Bewerbung auf dein offenes SingingCity Zeitfenster <strong>{{ .Time }}</strong> am <strong>{{ .Date }}</strong> erhalten.</p>
	<p>
	Eben hat sich <em>{{ .Artist }}</em> für einen Auftritt bei dir beworben.
	Alle Informationen zu <em>{{ .Artist }}</em> findest du in deinem SingingCity Profil.
	Schaue dir die Bewerbung an und sage zu oder ab. Bestätigst du den Auftritt, so erhält <em>{{ .Artist }}</em> ein Bestätigungsmail, das zugleich als Vertrag zwischen euch dient.
	Solltest du weitere Fragen haben, kannst du <em>{{ .Artist }}</em> nach der Zusage direkt via der neu angezeigten E-Mail-Adresse kontaktieren.
	</p>
	<a href="{{ .BaseUrl }}/dashboard">Zu meinem SingingCity Profil</a>
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
	<a href="{{ .BaseUrl }}/dashboard">To my SingingCity profile</a>
	<p>Sincerely,<br>
	Your SingingCity Team</p>`

	ConfirmedMessage = `
	<h1>Hallo {{ .Username }}</h1>

	<p>
	Herzlichen Glückwunsch!
	Wir freuen uns, dir mitteilen zu können, dass deine Bewerbung bei {{ .Venue }} für das Zeitfenster {{ .Time }} am {{ .Date }} erfolgreich angenommen wurde.
	Dein Auftritt ist nun fixiert, und dieses Bestätigungsmail dient als verbindlicher Vertrag. Bei weiteren Fragen oder wenn du zusätzliche Informationen benötigst, kannst du dich gerne direkt an die Konzert-Location wenden unter <a href="mailto:{{ .Contact }}">{{ .Contact }}</a>.
	</p>

	<p>Details:<p>

	<table>
    <tr><td>Konzert-Location:</td><td>{{ .Venue }}</td></tr>
    <tr><td>Datum:</td><td>{{ .Date }}</td></tr>
    <tr><td>Zeitfenster:</td><td>{{ .Time }}</td></tr>
    <tr><td>Gage:</td><td>{{ .Wage }}</td></tr>
	<tr><td>Adresse:</td><td>{{ .Address }}</td></tr>
    <tr><td>Kontakt:</td><td><a href="mailto:{{ .Contact }}">{{ .Contact }}</a></td></tr>
	</table>

	<p>Zur Verwaltung deiner Auftritte besuche dein <a href="{{ .BaseUrl }}/dashboard">SingingCity Profil</a></p>

	<p>Wir freuen uns schon sehr auf deinen Auftritt!</p>

	<p>Herzlichst,<br>
	Dein SingingCity Team</p>

	<hr>

	<h1>Hello {{ .Username }}</h1>

	<p>
	Congratulations!
	We are pleased to inform you that your application to {{ .Venue }} for the Time slot {{ .Time }} on {{ .Date }} has been successfully accepted.
	Your performance is now confirmed and this confirmation email serves as a binding contract between the two of you.
	If you have any further questions or require additional information, please contact the venue directly at <a href="mailto:{{ .Contact }}">{{ .Contact }}</a>.
	</p>

	<p>Details:</p>

	<table>
    <tr><td>Venue:</td><td>{{ .Venue }}</td></tr>
    <tr><td>Date:</td><td>{{ .Date }}</td></tr>
    <tr><td>Timeslot:</td><td>{{ .Time }}</td></tr>
    <tr><td>Fee:</td><td>{{ .Wage }}</td></tr>
	<tr><td>Address:</td><td>{{ .Address }}</td></tr>
    <tr><td>Contact:</td><td><a href="mailto:{{ .Contact }}">{{ .Contact }}</a></td></tr>
	</table>

	<p>To manage your performances visit your <a href="{{ .BaseUrl }}/dashboard">SingingCity profile</a></p>

	<p>We are looking forward to your performance!</p>

	<p>Sincerely,<br>
	Your SingingCity Team</p>`

	RejectedMessage = `
	<h1>Hallo {{ .Username }}</h1>

	<p>Leider müssen wir dir mitteilen, dass deine Bewerbung für das offene Zeitfenster {{ .Time }} bei {{ .Venue }} am {{ .Date }} abgelehnt wurde.</p>

	<p>
	Nicht entmutigen lassen!
	Es gibt noch viele weitere spannende Auftrittsmöglichkeiten für dich.
	Schaue dafür regelmässig in dein SingingCity Profil, um keine neue Zeitfenster der Konzert-Locations zu verpassen.
	</p>

	<a href="{{ .BaseUrl }}/dashboard">Zu meinem SingingCity Profil</a>
	<p>Herzlichst,<br>
	Dein SingingCity Team</p>

	<hr>

	<h1>Hello {{ .Username }}</h1>

	<p>Unfortunately, we have to inform you that your application for the open timeslot at {{ .Venue }} on {{ .Date }} {{ .Time }} has been rejected.</p>

	<p>
	Don't be discouraged!
	There are many more exciting performance opportunities for you.
	Check your SingingCity profile regularly to make sure you don't miss any new concert venue time slots.
	</p>

	<a href="{{ .BaseUrl }}/dashboard">To my SingingCity profile</a>

	<p>Sincerely,<br>
	Your SingingCity Team</p>`

	PasswordResetMessage = `
	<h1>Hallo {{ .Username }}</h1>
	<p>
	Du kannst dein Passwort unter folgendem Link zurücksetzen: 
	<a href="{{ .BaseUrl }}/reset-password?code={{ .Code }}">Passwort zurücksetzen</a>.
	</p>

	<hr>	

	<h1>Hello {{ .Username }}</h1>
	<p>
	You can reset your password using the following link:
	<a href="{{ .BaseUrl }}/reset-password?code={{ .Code }}">Reset Password</a>.
	</p>`
)

var (
	registerMessageTmpl    = template.Must(template.New("register").Parse(RegisterMessage))
	applicationMessageTmpl = template.Must(template.New("application").Parse(ApplicationMessage))
	confirmedMessageTmpl   = template.Must(template.New("confirmed").Parse(ConfirmedMessage))
	rejectedMessageTmpl    = template.Must(template.New("rejected").Parse(RejectedMessage))
	passwordResetMessage   = template.Must(template.New("reset").Parse(PasswordResetMessage))
)

type MessageParams struct {
	Username string
	Artist   string
	Time     string
	Date     string
	Contact  string
	Venue    string
	Wage     string
	Address  string
	Code     string
	BaseUrl  string
}
