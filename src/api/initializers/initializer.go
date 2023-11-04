package initializers

import (
	"shorty/data"
)

func Initialize() {
	data.InitializeDb()
}
func UnInitialize() {
	data.UnInitializeDb()
}
