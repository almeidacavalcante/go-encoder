package domain

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewJob(t *testing.T) {
	video := NewVideo("resourceId", "filePath")
	job, _ := NewJob("output", "outputBucket", video)

	require.NotNil(t, job.GetId())
	require.Equal(t, "output", job.GetOutputBucketPath())
	require.Equal(t, "outputBucket", job.GetStatus())
	require.Equal(t, video, job.GetVideo())
	require.Empty(t, job.GetError())
	require.NotNil(t, job.GetCreatedAt())
	require.NotNil(t, job.GetUpdatedAt())
}
