package user

const IncorrectCredentials = "IncorrectCredentials"
const Blocked = "Blocked"
const AwaitingAccept = "AwaitingAccept"
const NotAuthorized = "NotAuthorized"
const PermissionDenied = "PermissionDenied"

var Values = map[string]string{
	IncorrectCredentials: "User or password is not correct",
	Blocked:              "User was blocked",
	AwaitingAccept:       "Not accepted yet",
	NotAuthorized:        "Authorization needed",
	PermissionDenied:     "Permission denied for this action",
}
