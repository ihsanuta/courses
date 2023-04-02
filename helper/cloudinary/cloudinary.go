package cloudinary

import (
	"context"
	"courses/config"
	"time"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
)

func ImageUploadHelper(input interface{}) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//create cloudinary instance
	cld, err := cloudinary.NewFromParams(config.Cloudinary["cloud_name"].(string), config.Cloudinary["api_key"].(string), config.Cloudinary["api_secret"].(string))
	if err != nil {
		return "", err
	}

	//upload file
	uploadParam, err := cld.Upload.Upload(ctx, input, uploader.UploadParams{
		Folder: config.Cloudinary["upload_folder"].(string),
	})
	if err != nil {
		return "", err
	}

	return uploadParam.SecureURL, nil
}
