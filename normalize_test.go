package vectors

import (
	"context"
	"errors"
	"github.com/cucumber/godog"
)

func iNormalizeVector(key any) func(ctx context.Context) (context.Context, error) {
	return func(ctx context.Context) (context.Context, error) {
		v, ok := ctx.Value(key).(Vector)
		if !ok {
			return ctx, errors.New("vector not found in context")
		}
		return context.WithValue(ctx, resultKey{}, v.Normalize()), nil
	}
}

func initializeNormalizeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^I normalize vector a$`, iNormalizeVector(aKey{}))
}
