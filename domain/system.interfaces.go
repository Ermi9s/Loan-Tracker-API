package domain

type SystemLogsRepository interface {
	GetAllEvents() ([]Logs, error)
}

