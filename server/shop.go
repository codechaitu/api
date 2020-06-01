package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type banners struct {
	ImageUrl     string
	CategoryName string
}

func shopHome(c *gin.Context) {
	categoryData := []banners{
		{
			ImageUrl:     "https://drive.google.com/file/d/1udpa7cHYUlfQ3G7eI45Nnz1iGf4wnvNV/view?usp=sharing",
			CategoryName: "Ranveer",
		},
		{
			ImageUrl:     "https://drive.google.com/file/d/1U1SIbgxA36GZbYd_TNjohoLU4fpoQNR0/view?usp=sharing",
			CategoryName: "Deepika",
		},
		{
			ImageUrl:     "https://drive.google.com/file/d/1-G1fij6jM6tPhyJvQPmR6sw3JnREPC8A/view?usp=sharing",
			CategoryName: "Sonam Kapoor",
		},
	}
	c.JSON(http.StatusOK, categoryData)
}
