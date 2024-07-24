package log

func DebugLevel(message string) {
	levelLogger().Debug(message)
}
