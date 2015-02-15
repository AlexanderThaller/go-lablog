package data

import "time"

type Note struct {
	Project
	TimeStamp time.Time
	Text      string
}

func (note Note) ValueArray() []string {
	return []string{
		note.TimeStamp.Format(EntryCSVTimeStampFormat),
		"note",
		note.Text,
	}
}

func (note Note) GetProject() Project {
	return note.Project
}

func (note Note) Format(indent uint) string {
	var indentchar string
	for i := uint(0); i < indent; i++ {
		indentchar += "="
	}

	out := indentchar + "= " + note.Project.Name + "\n"
	out += indentchar + "== " + note.TimeStamp.Format(EntryCSVTimeStampFormat) + "\n"
	out += note.Text
	out += "\n"

	return out
}

// NotesByTimeStamp allows sorting project slices by name.
type NotesByTimeStamp []Note

func (by NotesByTimeStamp) Len() int {
	return len(by)
}

func (by NotesByTimeStamp) Swap(i, j int) {
	by[i], by[j] = by[j], by[i]
}

func (by NotesByTimeStamp) Less(i, j int) bool {
	return by[i].TimeStamp.Before(by[j].TimeStamp)
}
