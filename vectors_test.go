package vectors

import (
	"context"
	"errors"
	"fmt"
	"github.com/cucumber/godog"
	"strconv"
	"testing"
)

type aKey struct{}
type bKey struct{}
type resultKey struct{}
type errKey struct{}

func convertTableToVector(table *godog.Table) (Vector, error) {
	v := make(Vector, len(table.Rows))
	for i, row := range table.Rows {
		if len(row.Cells) != 1 {
			return nil, errors.New("table must have one column")
		}
		value, err := strconv.ParseFloat(row.Cells[0].Value, 64)
		if err != nil {
			return nil, err
		}
		v[i] = value
	}
	return v, nil
}

func vectorIs(key any) func(ctx context.Context, table *godog.Table) (context.Context, error) {
	return func(ctx context.Context, table *godog.Table) (context.Context, error) {
		v, err := convertTableToVector(table)
		if err != nil {
			return ctx, err
		}
		return context.WithValue(ctx, key, v), nil
	}
}

func theResultShouldBe(ctx context.Context, expected float64) (context.Context, error) {
	result, ok := ctx.Value(resultKey{}).(float64)
	if !ok {
		return ctx, errors.New("result not found in context")
	}
	if result != expected {
		return ctx, fmt.Errorf("expected %f, but got %f", expected, result)
	}
	return ctx, nil
}

func theResultShouldBeVector(ctx context.Context, expected *godog.Table) (context.Context, error) {
	result, ok := ctx.Value(resultKey{}).(Vector)
	if !ok {
		return ctx, errors.New("result not found in context")
	}
	expectedVector, err := convertTableToVector(expected)
	if err != nil {
		return ctx, err
	}
	if !result.Equals(expectedVector) {
		return ctx, fmt.Errorf("expected %v, but got %v", expectedVector, result)
	}
	return ctx, nil
}

func theResultShouldBeAnError(ctx context.Context) (context.Context, error) {
	_, ok := ctx.Value(errKey{}).(error)
	if !ok {
		return ctx, errors.New("error not found in context")
	}
	return ctx, nil
}

func theResultShouldBeBool(ctx context.Context, expectedStr string) (context.Context, error) {
	result, ok := ctx.Value(resultKey{}).(bool)
	if !ok {
		return ctx, errors.New("result not found in context")
	}
	expected, err := strconv.ParseBool(expectedStr)
	if err != nil {
		return ctx, err
	}
	if result != expected {
		return ctx, fmt.Errorf("expected %t, but got %t", expected, result)
	}
	return ctx, nil
}

func initializeSharedSteps(ctx *godog.ScenarioContext) {
	ctx.Step(`^vector a is$`, vectorIs(aKey{}))
	ctx.Step(`^vector b is$`, vectorIs(bKey{}))
	ctx.Step(`^the result should be (\-?\d+\.\d+)$`, theResultShouldBe)
	ctx.Step(`^the result should be vector$`, theResultShouldBeVector)
	ctx.Step(`^the result should be an error$`, theResultShouldBeAnError)
	ctx.Step(`^the result should be (true|false)$`, theResultShouldBeBool)
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	initializeSharedSteps(ctx)
	initializeCosineSimilarityScenario(ctx)
	initializeNormalizeScenario(ctx)
	initializeIsNormalizedScenario(ctx)
	initializeEqualsScenario(ctx)
}

func TestFeatures(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: InitializeScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"features"},
			TestingT: t, // Testing instance that will run subtests.
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}
