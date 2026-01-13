package vectors

import (
	"context"
	"github.com/cucumber/godog"
)

func iAddTheVectors(ctx context.Context) (context.Context, error) {
	a, ok := ctx.Value(aKey{}).(Vector)
	if !ok {
		return ctx, nil
	}
	b, ok := ctx.Value(bKey{}).(Vector)
	if !ok {
		return ctx, nil
	}

	result, err := a.Add(b)
	if err != nil {
		return context.WithValue(ctx, errKey{}, err), nil
	}

	return context.WithValue(ctx, resultKey{}, result), nil
}

func iSubtractVectorBFromVectorA(ctx context.Context) (context.Context, error) {
	a, ok := ctx.Value(aKey{}).(Vector)
	if !ok {
		return ctx, nil
	}
	b, ok := ctx.Value(bKey{}).(Vector)
	if !ok {
		return ctx, nil
	}

	result, err := a.Subtract(b)
	if err != nil {
		return context.WithValue(ctx, errKey{}, err), nil
	}

	return context.WithValue(ctx, resultKey{}, result), nil
}

func initializeArithmeticScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^I add the vectors$`, iAddTheVectors)
	ctx.Step(`^I subtract vector b from vector a$`, iSubtractVectorBFromVectorA)
}
