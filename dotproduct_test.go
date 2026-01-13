package vectors

import (
	"context"
	"github.com/cucumber/godog"
)

func iCalculateTheDotProduct(ctx context.Context) (context.Context, error) {
	a, ok := ctx.Value(aKey{}).(Vector)
	if !ok {
		return ctx, nil
	}
	b, ok := ctx.Value(bKey{}).(Vector)
	if !ok {
		return ctx, nil
	}

	product, err := DotProduct(a, b)
	if err != nil {
		return context.WithValue(ctx, errKey{}, err), nil
	}

	return context.WithValue(ctx, resultKey{}, product), nil
}

func initializeDotProductScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^I calculate the dot product$`, iCalculateTheDotProduct)
}
