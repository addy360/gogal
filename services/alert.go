package services

type AlertLevels struct {
	AlertDanger  string
	AlertInfo    string
	AlertSuccess string
}

type Alert struct {
	Level   string
	Message string
	AlertLevels
}

func NewAlert() *Alert {
	return &Alert{
		Level:   "",
		Message: "",
		AlertLevels: AlertLevels{
			AlertDanger:  "danger",
			AlertInfo:    "info",
			AlertSuccess: "success",
		},
	}
}
