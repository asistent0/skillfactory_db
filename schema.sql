/*
    Схема БД для информационной системы
    отслеживания выполнения задач.
*/

DROP TABLE IF EXISTS task_label, task, label, "user";

-- пользователи системы
CREATE TABLE "user"
(
    id   SERIAL PRIMARY KEY,
    name TEXT NOT NULL
);

-- метки задач
CREATE TABLE label
(
    id   SERIAL PRIMARY KEY,
    name TEXT NOT NULL
);

-- задачи
CREATE TABLE task
(
    id        SERIAL PRIMARY KEY,
    opened    BIGINT NOT NULL               DEFAULT extract(epoch from now()), -- время создания задачи
    closed    BIGINT                        DEFAULT 0,                         -- время выполнения задачи
    author_id INTEGER REFERENCES "user" (id) DEFAULT 0,                        -- автор задачи
    assigned_id INTEGER REFERENCES "user" (id) DEFAULT 0,                      -- ответственный
    title TEXT, -- название задачи
    content TEXT -- задачи
);

-- связь многие-ко-многим между задачами и метками
CREATE TABLE task_label
(
    task_id  INTEGER REFERENCES task (id),
    label_id INTEGER REFERENCES label (id)
);

-- наполнение БД начальными данными
INSERT INTO "user" (id, name)
VALUES (0, 'default');
