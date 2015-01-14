package project

import (
	"sort"
	"strconv"
	"time"

	"github.com/juju/errgo"
)

const (
	RecordTimeStampFormat = time.RFC3339Nano

	ActionNote         = "note"
	ActionTodo         = "todo"
	ActionTrack        = "track"
	ActionTracksActive = "tracksactive"
)

type Record interface {
	CSV() []string
	GetAction() string
	GetProject() string
	GetTimeStamp() string
	GetValue() string
	GetFormattedValue() string
}

func RecordFromCSV(values []string) (Record, error) {
	if len(values) < 2 {
		return nil, errgo.New("we need at least two fields for parsing")
	}

	recordtype := values[1]
	switch recordtype {
	case ActionNote:
		return NoteFromCSV(values)
	case ActionTodo:
		return TodoFromCSV(values)
	default:
		return nil, errgo.New("can not parse record type " + recordtype)
	}
}

func NoteFromCSV(values []string) (Note, error) {
	if len(values) != 3 {
		return Note{}, errgo.New("we need three fields for parsing a note")
	}

	if values[1] != "note" {
		return Note{}, errgo.New("second field has to have the string 'note' in it")
	}

	timestamp, err := time.Parse(RecordTimeStampFormat, values[0])
	if err != nil {
		return Note{}, err
	}

	note := Note{
		TimeStamp: timestamp,
		Value:     values[2],
	}

	return note, nil
}

func TodoFromCSV(values []string) (Todo, error) {
	if len(values) != 4 {
		return Todo{}, errgo.New("we need three fields for parsing a todo")
	}

	if values[1] != "todo" {
		return Todo{}, errgo.New("second field has to have the string 'todo' in it")
	}

	timestamp, err := time.Parse(RecordTimeStampFormat, values[0])
	if err != nil {
		return Todo{}, err
	}
	done, err := strconv.ParseBool(values[3])
	if err != nil {
		return Todo{}, err
	}

	todo := Todo{
		TimeStamp: timestamp,
		Value:     values[2],
		Done:      done,
	}

	return todo, nil
}

func TrackFromCSV(values []string) (Track, error) {
	if len(values) != 3 {
		return Track{}, errgo.New("we need three fields for parsing a todo")
	}

	if values[1] != "track" {
		return Track{}, errgo.New("second field has to have the string 'track' in it")
	}

	timestamp, err := time.Parse(RecordTimeStampFormat, values[0])
	if err != nil {
		return Track{}, err
	}

	todo := Track{
		TimeStamp: timestamp,
		Value:     values[2],
	}

	return todo, nil
}

type Note struct {
	Project   string
	TimeStamp time.Time
	Value     string
}

func (note Note) CSV() []string {
	return []string{
		note.GetTimeStamp(),
		note.GetAction(),
		note.GetValue(),
	}
}

func (note Note) GetAction() string {
	return ActionNote
}

func (note Note) GetProject() string {
	return note.Project
}

func (note Note) GetTimeStamp() string {
	return note.TimeStamp.Format(RecordTimeStampFormat)
}

func (note Note) GetValue() string {
	return note.Value
}

func (note Note) GetFormattedValue() string {
	return note.Value
}

func (note *Note) SetProject(project string) {
	note.Project = project
}

type Todo struct {
	Project   string
	TimeStamp time.Time
	Value     string
	Done      bool
}

func (todo Todo) CSV() []string {
	return []string{
		todo.GetTimeStamp(),
		todo.GetAction(),
		todo.GetValue(),
		strconv.FormatBool(todo.Done),
	}
}

func (todo Todo) GetAction() string {
	return ActionTodo
}

func (todo Todo) GetProject() string {
	return todo.Project
}

func (todo Todo) GetTimeStamp() string {
	return todo.TimeStamp.Format(RecordTimeStampFormat)
}

func (todo Todo) GetValue() string {
	return todo.Value
}

func (todo Todo) GetFormattedValue() string {
	return "* " + todo.Value
}

func (todo Todo) SetProject(project string) {
	todo.Project = project
}

type TodoByDate []Todo

func (todo TodoByDate) Len() int {
	return len(todo)
}

func (todo TodoByDate) Swap(i, j int) {
	todo[i], todo[j] = todo[j], todo[i]
}

func (todo TodoByDate) Less(i, j int) bool {
	return todo[j].TimeStamp.After(todo[i].TimeStamp)
}

type TodoByValue []Todo

func (todo TodoByValue) Len() int {
	return len(todo)
}

func (todo TodoByValue) Swap(i, j int) {
	todo[i], todo[j] = todo[j], todo[i]
}

func (todo TodoByValue) Less(i, j int) bool {
	return todo[i].Value < todo[j].Value
}

type NotesByDate []Note

func (note NotesByDate) Len() int {
	return len(note)
}

func (note NotesByDate) Swap(i, j int) {
	note[i], note[j] = note[j], note[i]
}

func (note NotesByDate) Less(i, j int) bool {
	return note[j].TimeStamp.After(note[i].TimeStamp)
}

type Track struct {
	Project   string
	TimeStamp time.Time
	Value     string
}

func (track Track) CSV() []string {
	return []string{
		track.GetTimeStamp(),
		track.GetAction(),
		track.GetValue(),
	}
}

func (track Track) GetAction() string {
	return ActionTrack
}

func (track Track) GetProject() string {
	return track.Project
}

func (track Track) GetTimeStamp() string {
	return track.TimeStamp.Format(RecordTimeStampFormat)
}

func (track Track) GetValue() string {
	return track.Value
}

func (track Track) GetFormattedValue() string {
	out := "* " + track.GetTimeStamp()

	if track.Value != "" {
		out += " -- " + track.Value
	}

	return out
}

type TracksByDate []Track

func (track TracksByDate) Len() int {
	return len(track)
}

func (track TracksByDate) Swap(i, j int) {
	track[i], track[j] = track[j], track[i]
}

func (track TracksByDate) Less(i, j int) bool {
	return track[j].TimeStamp.After(track[i].TimeStamp)
}

func FilterInactiveTodos(todos []Todo) []Todo {
	filter := make(map[string]Todo)

	sort.Sort(TodoByDate(todos))
	for _, todo := range todos {
		filter[todo.Value] = todo
	}

	var out []Todo
	for _, todo := range filter {
		if todo.Done {
			continue
		}

		out = append(out, todo)
	}

	return out
}

type Duration struct {
	Project  string
	Duration time.Duration
	Value    string
}

func (duration Duration) GetFormattedValue() string {
	value := duration.Value
	if value != "" {
		value += " -- "
	}

	return "* " + value + duration.Duration.String()
}

type DurationsByValue []Duration

func (duration DurationsByValue) Len() int {
	return len(duration)
}

func (duration DurationsByValue) Swap(i, j int) {
	duration[i], duration[j] = duration[j], duration[i]
}

func (duration DurationsByValue) Less(i, j int) bool {
	return duration[i].Value < duration[j].Value
}

func WriteRecord(datadir string, record Record) error {
	return errgo.New("not implemented")
}
