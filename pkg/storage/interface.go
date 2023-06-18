package storage

import "30.8.1/pkg/storage/postgres"

// Interface Интерфейс БД
// Этот интерфейс позволяет абстрагироваться от конкретной СУБД
// Можно создать реализацию БД в памяти для модульных тестов
type Interface interface {
	Tasks(int, int, int) ([]postgres.Task, error)
	NewTask(postgres.Task) (int, error)
	UpdateTask(postgres.Task) (int, error)
	DeleteTask(postgres.Task) error
}
