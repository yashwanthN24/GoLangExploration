package auth


func exportSession() string{
	return "session"
}

func GetSession()string {
	return exportSession()
}
