package uuidgeneratormock

var data []string
var pointerIndex int

// Init initialize the data
func Init(d []string) {
	data = d
	pointerIndex = 0
}

// Reset resets the pointerIndex
func Reset() {
	pointerIndex = 0
}

// UUIDGeneratorMock returns the mock UUID string in a determined way
func UUIDGeneratorMock() string {
	if len(data) == 0 {
		return ""
	}
	pointerIndex = pointerIndex % len(data)
	uuid := data[pointerIndex]
	pointerIndex++
	return uuid
}
