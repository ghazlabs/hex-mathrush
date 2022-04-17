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
	questions := queststrg.GetQuestions()

	// initialize storage
	questStrg, err := New(Config{
		Questions: questions,
	})
	require.NoError(t, err)
	// test the get random question function
	question, err := questStrg.GetRandomQuestion(context.Background())
	require.NoError(t, err)

	// make sure that question is part of questions
	assert.Contains(t, questions, *question, "the question is out of questions list")
}
