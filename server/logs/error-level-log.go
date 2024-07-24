package log

func ErrorLevel(message string) {
	levelLogger().Error(message)
}
