package domain

type SystemLogsRepository interface {
	GetAllEvents() ([]Logs, error)
}


type SystemLogsUsecase interface {
	GetAllEvents() ([]Logs, error)
}
