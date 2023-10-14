package repositories_test

import (
	"github.com/almeidacavalcante/go-encoder/application/repositories"
	"github.com/almeidacavalcante/go-encoder/domain"
	"github.com/almeidacavalcante/go-encoder/infra/database"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestJobRepositoryDb_Insert(t *testing.T) {
	db := database.NewDbTest()
	defer db.Close()

	video := domain.NewVideo("resourceId", "filePath")
	video.FilePath = "path"

	videoRepository := repositories.VideoRepositoryDb{Db: db}
	videoRepository.Insert(video)

	job, err := domain.NewJob("output", "outputBucket", video)
	require.Nil(t, err)

	jobRepository := repositories.NewJobRepository(db)
	jobRepository.Insert(job)

	job, err = jobRepository.Find(job.ID)
	require.NotEmpty(t, job.ID)
	require.Nil(t, err)
	require.Equal(t, "output", job.OutputBucketPath)
	require.Equal(t, "outputBucket", job.Status)
	require.Equal(t, video.ID, job.Video.ID)

	job.OutputBucketPath = "new output"
	_, err = jobRepository.Update(job)

	job, err = jobRepository.Find(job.ID)
	require.Nil(t, err)
	require.Equal(t, "new output", job.OutputBucketPath)

}
