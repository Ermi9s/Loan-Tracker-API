package usecase

import domain "github.com/Loan-Tracker-API/Loan-Tracker-API/domain"

type LogUsecase struct {
	LogRepository domain.SystemLogsRepository
}

func NewLogUsecase(repository domain.SystemLogsRepository) *LogUsecase {
	return &LogUsecase{LogRepository: repository}
}

func (usecase *LogUsecase) GetAllEvents() ([]domain.Logs, error) {
	logs, err := usecase.LogRepository.GetAllEvents()
	if err != nil {
		return []domain.Logs{}, err
	}
	return logs, nil
}