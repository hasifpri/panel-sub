package applicationservice

import (
	"context"
	applicationexception "panel-subs/application/exception"
)

type IUseCase[TInput any, TOutput any] interface {
	Execute(ctx context.Context, input TInput) (TOutput, *applicationexception.Exception)
}
