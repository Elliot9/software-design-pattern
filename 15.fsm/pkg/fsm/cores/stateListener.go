package cores

import "context"

type StateListener interface {
	Handle(context context.Context)
}
