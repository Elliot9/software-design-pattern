package cores

import "context"

type Action interface {
	Execute(context context.Context) context.Context
}
