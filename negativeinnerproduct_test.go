package vectors

import (
	"context"
	"github.com/cucumber/godog"
)

func iCalculateTheNegativeInnerProduct(ctx context.Context) (context.Context, error) {
	a, ok := ctx.Value(aKey{}).(Vector)
	if !ok {
		return ctx, nil
	}
	b, ok := ctx.Value(bKey{}).(Vector)
	if !ok {
		return ctx, nil
	}

	product, err := NegativeInnerProduct(a, b)
	if err != nil {
		return context.WithValue(ctx, errKey{}, err), nil
	}

	return context.WithValue(ctx, resultKey{}, product), nil
}

func initializeNegativeInnerProductScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^I calculate the negative inner product$`, iCalculateTheNegativeInnerProduct)
}
