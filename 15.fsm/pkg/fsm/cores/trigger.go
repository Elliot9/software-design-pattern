package cores

import "context"

type Trigger interface {
	Match(event Event, context context.Context) bool
}
