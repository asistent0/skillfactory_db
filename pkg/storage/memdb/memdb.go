package memdb

import "30.8.1/pkg/storage/postgres"

// DB Пользовательский тип данных - реализация БД в памяти
type DB []postgres.Task

// Выполнение контракта интерфейса storage Interface

func (db DB) Tasks(int, int, int) ([]postgres.Task, error) {
	return db, nil
}

func (db DB) NewTask(postgres.Task) (int, error) {
	return 0, nil
}

func (db DB) UpdateTask(t postgres.Task) (int, error) {
	return t.ID, nil
}

func (db DB) DeleteTask(postgres.Task) error {
	return nil
}
