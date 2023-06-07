package response

import (
	"github.com/thoas/go-funk"
)

// Code ...
type Code struct {
	Key     string
	Message string
	Code    int
}

var notFoundKey = Code{
	Key:     CommonNotFound,
	Message: "Common Not Found",
	Code:    -1,
}

// GetByKey give key and receive message + code
func GetByKey(key string) Code {
	item := funk.Find(list, func(item Code) bool {
		return item.Key == key
	})

	if item == nil {
		return notFoundKey
	}
	return item.(Code)
}

var list = make([]Code, 0)

// Init ...
func Init() {
	AddListCodes(common)
}

// AddListCodes to all codes managed by this package
func AddListCodes(codes []Code) {
	list = append(list, codes...)
}
