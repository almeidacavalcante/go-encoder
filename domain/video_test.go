package domain_test

import (
	"github.com/almeidacavalcante/go-encoder/domain"
	require "github.com/stretchr/testify/require"
	"testing"
)

func TestResourceIdCannotBeEmpty(t *testing.T) {
	video := domain.NewVideo("", "filePath")
	err := video.Validate()

	require.Error(t, err)
}
