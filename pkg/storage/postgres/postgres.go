package postgres

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

// Storage Хранилище данных.
type Storage struct {
	db *pgxpool.Pool
}

// New Конструктор, принимает строку подключения к БД.
func New(constr string) (*Storage, error) {
	db, err := pgxpool.Connect(context.Background(), constr)
	if err != nil {
		return nil, err
	}
	s := Storage{
		db: db,
	}
	return &s, nil
}

// Task Задача.
type Task struct {
	ID         int
	Opened     int64
	Closed     int64
	AuthorID   int
	AssignedID int
	Title      string
	Content    string
}

// Tasks возвращает список задач из БД.
func (s *Storage) Tasks(taskID, authorID, labelID int) ([]Task, error) {
	rows, err := s.db.Query(context.Background(), `
		SELECT 
			task.id,
			task.opened,
			task.closed,
			task.author_id,
			task.assigned_id,
			task.title,
			task.content
		FROM task
		LEFT JOIN task_label ON task_label.task_id = task.id
		WHERE
			($1 = 0 OR task.id = $1) AND
			($2 = 0 OR task.author_id = $2) AND
			($3 = 0 OR task_label.label_id = $3)
		ORDER BY task.id;
	`,
		taskID,
		authorID,
		labelID,
	)
	if err != nil {
		return nil, err
	}
	var tasks []Task
	// итерирование по результату выполнения запроса
	// и сканирование каждой строки в переменную
	for rows.Next() {
		var t Task
		err = rows.Scan(
			&t.ID,
			&t.Opened,
			&t.Closed,
			&t.AuthorID,
			&t.AssignedID,
			&t.Title,
			&t.Content,
		)
		if err != nil {
			return nil, err
		}
		// добавление переменной в массив результатов
		tasks = append(tasks, t)
	}
	// ВАЖНО не забыть проверить rows.Err()
	return tasks, rows.Err()
}

// NewTask создаёт новую задачу и возвращает её id.
func (s *Storage) NewTask(t Task) (int, error) {
	var id int
	err := s.db.QueryRow(context.Background(), `
		INSERT INTO task (title, content, opened)
		VALUES ($1, $2, $3) RETURNING id;
		`,
		t.Title,
		t.Content,
		t.Opened,
	).Scan(&id)
	return id, err
}

// UpdateTask обновляет задачу и возвращает её id.
func (s *Storage) UpdateTask(t Task) (int, error) {
	_, err := s.db.Exec(context.Background(), `
		UPDATE task SET title = $1, content = $2 WHERE id = $3;
		`,
		t.Title,
		t.Content,
		t.ID,
	)
	return t.ID, err
}

// DeleteTask удаляет задачу.
func (s *Storage) DeleteTask(t Task) error {
	_, err := s.db.Exec(context.Background(), `
		DELETE FROM task WHERE id = $1;
		`,
		t.ID,
	)
	return err
}
