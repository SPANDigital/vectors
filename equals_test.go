package vectors

import (
	"context"
	"errors"
	"github.com/cucumber/godog"
)

func iCheckIfVectorsAreEqual(ctx context.Context) (context.Context, error) {
	a, ok := ctx.Value(aKey{}).(Vector)
	if !ok {
		return ctx, errors.New("first arg not found in context")
	}
	b, ok := ctx.Value(bKey{}).(Vector)
	if !ok {
		return ctx, errors.New("second arg not found in context")
	}
	return context.WithValue(ctx, resultKey{}, a.Equals(b)), nil
}

func initializeEqualsScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^I check if vector a equals vector b$`, iCheckIfVectorsAreEqual)
}
