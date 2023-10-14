package repositories_test

import (
	"fmt"
	"github.com/almeidacavalcante/go-encoder/application/repositories"
	"github.com/almeidacavalcante/go-encoder/domain"
	"github.com/almeidacavalcante/go-encoder/infra/database"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestNewVideoRepositoryDbInsert(t *testing.T) {
	db := database.NewDbTest()
	defer db.Close()

	video := domain.NewVideo("resourceId", "filePath")
	video.ID = uuid.NewV4().String()
	video.CreatedAt = time.Now()

	repo := repositories.VideoRepositoryDb{Db: db}
	repo.Insert(video)

	v, err := repo.Find(video.ID)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(v.ID)
	require.NotEmpty(t, v.ID)
}
