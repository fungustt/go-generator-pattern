package tests

import (
	"context"
	"testing"

	"generator/generator"

	"github.com/stretchr/testify/assert"
)

func TestGenerator_CantRunTwice(t *testing.T) {
	g := TestGenerator()
	ctx := context.Background()

	err1 := g.Start(ctx)
	assert.Nil(t, err1)

	err2 := g.Start(ctx)
	assert.EqualError(t, err2, generator.ErrAlreadyStarted.Error())
}

func TestGenerator_CantStopStopped(t *testing.T) {
	g := TestGenerator()
	ctx := context.Background()

	err1 := g.Start(ctx)
	assert.Nil(t, err1)

	err2 := g.Stop()
	assert.Nil(t, err2)

	err3 := g.Stop()
	assert.EqualError(t, err3, generator.ErrAlreadyStopped.Error())
}

func TestGenerator_CantRestart(t *testing.T) {
	g := TestGenerator()
	ctx := context.Background()

	err1 := g.Start(ctx)
	assert.Nil(t, err1)

	err2 := g.Stop()
	assert.Nil(t, err2)

	err3 := g.Start(ctx)
	assert.EqualError(t, err3, generator.ErrAlreadyStopped.Error())
}

func TestGenerator_Get(t *testing.T) {
	g := TestGenerator()
	ctx := context.Background()

	_, err := g.Get()
	assert.EqualError(t, err, generator.ErrNotStarted.Error())

	err1 := g.Start(ctx)
	assert.Nil(t, err1)

	val1, err := g.Get()
	assert.Equal(t, val1, 1)
	assert.Nil(t, err)

	val2, err := g.Get()
	assert.Equal(t, val2, 2)
	assert.Nil(t, err)

	err = g.Stop()
	assert.Nil(t, err)

	_, err = g.Get()
	assert.EqualError(t, err, generator.ErrAlreadyStopped.Error())
}
