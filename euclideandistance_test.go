package vectors

import (
	"context"
	"github.com/cucumber/godog"
)

func iCalculateTheEuclideanDistance(ctx context.Context) (context.Context, error) {
	a, ok := ctx.Value(aKey{}).(Vector)
	if !ok {
		return ctx, nil
	}
	b, ok := ctx.Value(bKey{}).(Vector)
	if !ok {
		return ctx, nil
	}

	distance, err := EuclideanDistance(a, b)
	if err != nil {
		return context.WithValue(ctx, errKey{}, err), nil
	}

	return context.WithValue(ctx, resultKey{}, distance), nil
}

func initializeEuclideanDistanceScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^I calculate the Euclidean distance$`, iCalculateTheEuclideanDistance)
}
