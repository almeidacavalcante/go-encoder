package domain

import (
	"github.com/asaskevich/govalidator"
	"time"
)

type Video struct {
	ID         string    `json:"encoded_video_folder" valid:"uuid" gorm:"type:uuid;primary_key;"`
	ResourceID string    `json:"resource_id" valid:"required" gorm:"type:varchar(255)"`
	FilePath   string    `json:"file_path" valid:"required" gorm:"type:varchar(255)"`
	CreatedAt  time.Time `json:"created_at" valid:"-"`
	Jobs       []*Job    `json:"-" valid:"-" gorm:"ForeignKey:VideoID"`
}

func NewVideo(resourceID string, filePath string) *Video {
	video := &Video{
		ResourceID: resourceID,
		FilePath:   filePath,
	}
	return video
}

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

func (video *Video) Validate() error {
	_, err := govalidator.ValidateStruct(video)
	if err != nil {
		return err
	}
	return nil
}
