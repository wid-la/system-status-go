package interactors

import (
	"fmt"
	"time"

	"github.com/wid-la/system-status-go/pkg/core"
	"github.com/wid-la/system-status-go/pkg/models"
)

// Inter ...
type Inter struct {
	Deps     *core.Dependency
	Services map[string]models.Service
}

// Process ...
func (inter Inter) Process(interval int) {
	for {
		fmt.Println("> Process")

		inter.CheckStatuses()

		time.Sleep(time.Duration(interval) * 1000 * time.Millisecond)
	}
}
