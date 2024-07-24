package fileupload

import (
	"fmt"
	"net/http"
	"server/config"

	"server/utilService"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
	"github.com/gin-gonic/gin"
)

// image upload controller
func UploadImage(ctx *gin.Context) {
	//1. Get the image data from the request body
	var inputData struct {
		Images []string `json:"images"`
	}

	if err := ctx.ShouldBindJSON(&inputData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	//2. Set up the Cloudinary configuration
	cld, _ := cloudinary.NewFromParams(config.CLOUDINARY_CLOUD_NAME, config.CLOUDINARY_API_KEY, config.CLOUDINARY_SECRET)

	//4.upload images to cloudinary and store the urls in an array
	var urls []string
	for index := range inputData.Images {
		// Upload the image to Cloudinary
		response, err := cld.Upload.Upload(ctx.Request.Context(), inputData.Images[index], uploader.UploadParams{
			PublicID: utilService.PublicID(),
			Folder:   "persons",
		})
		if err != nil {
			fmt.Println("Error:", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
		fmt.Println(response.SecureURL)
		urls = append(urls, response.SecureURL)
	}
	// 5. Send the url to the client
	ctx.JSON(200, gin.H{"urls": urls})
}
