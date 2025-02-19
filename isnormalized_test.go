package vectors

import (
	"context"
	"errors"
	"github.com/cucumber/godog"
)

func iCheckIfVectorIsNormalized(key any) func(ctx context.Context) (context.Context, error) {
	return func(ctx context.Context) (context.Context, error) {
		v, ok := ctx.Value(key).(Vector)
		if !ok {
			return ctx, errors.New("vector not found in context")
		}
		return context.WithValue(ctx, resultKey{}, v.IsNormalized()), nil
	}
}

func initializeIsNormalizedScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^I check if vector a is normalized$`, iCheckIfVectorIsNormalized(aKey{}))
}
