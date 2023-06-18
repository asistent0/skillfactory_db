package postgres

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"reflect"
	"testing"
)

var s *Storage

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func TestMain(m *testing.M) {
	constr := "postgres://postgres:postgres@localhost:5433/tasks"
	var err error
	s, err = New(constr)
	if err != nil {
		log.Fatal(err)
	}
	os.Exit(m.Run())
}

func TestStorage_NewTask(t *testing.T) {
	type args struct {
		t Task
	}
	tests := []struct {
		name    string
		fields  *Storage
		args    args
		want    int
		wantErr bool
	}{
		{
			name:   "first",
			fields: s,
			args: args{t: Task{
				1, 1687082400, 0, 0, 0, "first", "test",
			}},
			want:    1,
			wantErr: false,
		},
		{
			name:   "first",
			fields: s,
			args: args{t: Task{
				2, 1687082400, 0, 0, 0, "two", "test",
			}},
			want:    2,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				db: tt.fields.db,
			}
			got, err := s.NewTask(tt.args.t)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewTask() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("NewTask() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStorage_Tasks(t *testing.T) {
	data, err := s.Tasks(0, 0, 0)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)
	data, err = s.Tasks(1, 0, 0)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)

	type args struct {
		taskID   int
		authorID int
		labelID  int
	}
	tests := []struct {
		name    string
		fields  *Storage
		args    args
		want    []Task
		wantErr bool
	}{
		{
			name:   "all",
			fields: s,
			args: args{
				taskID: 0, authorID: 0, labelID: 0,
			},
			want: []Task{
				{
					1, 1687082400, 0, 0, 0, "first", "test",
				},
				{
					2, 1687082400, 0, 0, 0, "two", "test",
				},
			},
			wantErr: false,
		},
		{
			name:   "one",
			fields: s,
			args: args{
				taskID: 1, authorID: 0, labelID: 0,
			},
			want: []Task{
				{
					1, 1687082400, 0, 0, 0, "first", "test",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				db: tt.fields.db,
			}
			got, err := s.Tasks(tt.args.taskID, tt.args.authorID, tt.args.labelID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Tasks() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tasks() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStorage_UpdateTask(t *testing.T) {
	type args struct {
		t Task
	}
	tests := []struct {
		name    string
		fields  *Storage
		args    args
		want    int
		wantErr bool
	}{
		{
			name:   "first",
			fields: s,
			args: args{t: Task{
				1, 1687082400, 0, 0, 0, "first", "test 1",
			}},
			want:    1,
			wantErr: false,
		},
		{
			name:   "first",
			fields: s,
			args: args{t: Task{
				2, 1687082400, 0, 0, 0, "two", "test 2",
			}},
			want:    2,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				db: tt.fields.db,
			}
			got, err := s.UpdateTask(tt.args.t)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateTask() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("UpdateTask() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStorage_DeleteTask(t *testing.T) {
	type args struct {
		t Task
	}
	tests := []struct {
		name    string
		fields  *Storage
		args    args
		wantErr bool
	}{
		{
			name:   "first",
			fields: s,
			args: args{t: Task{
				1, 1687082400, 0, 0, 0, "first", "test 1",
			}},
			wantErr: false,
		},
		{
			name:   "first",
			fields: s,
			args: args{t: Task{
				2, 1687082400, 0, 0, 0, "two", "test 2",
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				db: tt.fields.db,
			}
			if err := s.DeleteTask(tt.args.t); (err != nil) != tt.wantErr {
				t.Errorf("DeleteTask() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
