package tests

import (
    "testing"
    "tennisapp/backend/internal/services"

    "github.com/stretchr/testify/assert"
)

func TestGetMatchByID(t *testing.T) {
    match, err := services.GetMatchByID(1)
    assert.NoError(t, err)
    assert.NotNil(t, match)
    assert.Equal(t, 1, int(match.ID))
}
