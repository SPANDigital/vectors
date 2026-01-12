package vectors

import (
	"context"
	"github.com/cucumber/godog"
)

func iCalculateTheTaxicabDistance(ctx context.Context) (context.Context, error) {
	a, ok := ctx.Value(aKey{}).(Vector)
	if !ok {
		return ctx, nil
	}
	b, ok := ctx.Value(bKey{}).(Vector)
	if !ok {
		return ctx, nil
	}

	distance, err := TaxicabDistance(a, b)
	if err != nil {
		return context.WithValue(ctx, errKey{}, err), nil
	}

	return context.WithValue(ctx, resultKey{}, distance), nil
}

func initializeTaxicabDistanceScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^I calculate the taxicab distance$`, iCalculateTheTaxicabDistance)
}
