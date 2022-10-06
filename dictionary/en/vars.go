package en

var Values = map[string]map[string]string{
	"user": {
		"NotAuthorized": "Authorization needed", 
		"PermissionDenied": "Permission denied for this action", 
		"AwaitingAccept": "Not accepted yet", 
		"Blocked": "User was blocked", 
		"IncorrectCredentials": "User or password is not correct", 
	},
	"admin": {
		"Hello": "Hello, %s", 
	},
}