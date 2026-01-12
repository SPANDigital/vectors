package vectors

import (
	"context"
	"errors"
	"github.com/cucumber/godog"
)

func iCalculateMagnitudeOfVector(key any) func(ctx context.Context) (context.Context, error) {
	return func(ctx context.Context) (context.Context, error) {
		v, ok := ctx.Value(key).(Vector)
		if !ok {
			return ctx, errors.New("vector not found in context")
		}
		return context.WithValue(ctx, resultKey{}, v.Magnitude()), nil
	}
}

func initializeMagnitudeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^I calculate the magnitude of vector a$`, iCalculateMagnitudeOfVector(aKey{}))
}
