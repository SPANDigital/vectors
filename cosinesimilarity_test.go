package vectors

import (
	"context"
	"errors"
	"github.com/cucumber/godog"
)

func iCalculateTheCosineSimilarity(ctx context.Context) (context.Context, error) {
	a, ok := ctx.Value(aKey{}).(Vector)
	if !ok {
		return ctx, errors.New("first arg not found in context")
	}
	b, ok := ctx.Value(bKey{}).(Vector)
	if !ok {
		return ctx, errors.New("second arg not found in context")
	}
	result, err := CosineSimilarity(a, b)
	if err != nil {
		return context.WithValue(ctx, errKey{}, err), nil
	}
	return context.WithValue(ctx, resultKey{}, result), nil
}

func initializeCosineSimilarityScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^I calculate the cosine similarity$`, iCalculateTheCosineSimilarity)
}
