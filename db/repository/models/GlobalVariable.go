package models

import "github.com/joho/godotenv"

var (
	baseenvConfig, _  = godotenv.Read("../.env")
	ImagePath         = baseenvConfig["AWS_IMAGE_BASE_URL"]
	CustomerImagePath = ImagePath + "image/customer/"
	SalesmanImagePath = ImagePath
)
