package queststrg

import (
	"context"
	"testing"

	"github.com/ghazlabs/hex-mathrush/internal/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetRandomQuestion(t *testing.T) {
	// initialize questions
	var qs []core.Question
	// initialize storage
	strg, err := New(Config{
		Questions: qs,
	})
	require.NoError(t, err)
	err = strg.Init()
	// test the get random question function
	question, err := strg.GetRandomQuestion(context.Background())
	require.NoError(t, err)

	// make sure that question is part of questions
	assert.Contains(t, strg, &question, "the question is out of questions list")
}
