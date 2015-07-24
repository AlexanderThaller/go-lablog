package store

import "github.com/AlexanderThaller/lablog/src/data"

type EntriesStore interface {
	Write(data.Entry) error
}
