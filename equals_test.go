package vectors

import (
	"context"
	"errors"
	"strconv"

	"github.com/cucumber/godog"
)

type epsilonKey struct{}

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

func epsilonIs(ctx context.Context, epsilonStr string) (context.Context, error) {
	epsilon, err := strconv.ParseFloat(epsilonStr, 64)
	if err != nil {
		return ctx, err
	}
	return context.WithValue(ctx, epsilonKey{}, epsilon), nil
}

func iCheckIfVectorsAreEqualWithEpsilon(ctx context.Context) (context.Context, error) {
	a, ok := ctx.Value(aKey{}).(Vector)
	if !ok {
		return ctx, errors.New("first arg not found in context")
	}
	b, ok := ctx.Value(bKey{}).(Vector)
	if !ok {
		return ctx, errors.New("second arg not found in context")
	}
	epsilon, ok := ctx.Value(epsilonKey{}).(float64)
	if !ok {
		return ctx, errors.New("epsilon not found in context")
	}
	return context.WithValue(ctx, resultKey{}, a.EqualsWithEpsilon(b, epsilon)), nil
}

func initializeEqualsScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^I check if vector a equals vector b$`, iCheckIfVectorsAreEqual)
	ctx.Step(`^epsilon is (\S+)$`, epsilonIs)
	ctx.Step(`^I check if vector a equals vector b with epsilon$`, iCheckIfVectorsAreEqualWithEpsilon)
}
