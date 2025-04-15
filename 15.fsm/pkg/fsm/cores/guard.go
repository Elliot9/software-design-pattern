package cores

import "context"

type Guard interface {
	Evaluate(context context.Context) bool
}
