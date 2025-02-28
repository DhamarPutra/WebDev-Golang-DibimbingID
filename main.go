package main

import (
	"net/http"
	"PortoGo/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	// Inisialisasi DB
	models.InitDB()

	// Inisialisasi router Gin
	r := gin.Default()
	r.LoadHTMLGlob("views/*")

	// Route utama: Menampilkan semua portofolio
	r.GET("/", func(c *gin.Context) {
		portofolios, err := models.GetAllPortofolios()
		if err != nil {
			c.String(http.StatusInternalServerError, "Error fetching data")
			return
		}
		c.HTML(http.StatusOK, "index.html", gin.H{"portofolios": portofolios})
	})

	// Route tambah data
	r.GET("/create", func(c *gin.Context) {
		c.HTML(http.StatusOK, "create.html", nil)
	})
	r.POST("/create", func(c *gin.Context) {
		title := c.PostForm("title")
		description := c.PostForm("description")
		image := c.PostForm("image")

		err := models.CreatePortofolio(title, description, image)
		if err != nil {
			c.String(http.StatusInternalServerError, "Error inserting data")
			return
		}
		c.Redirect(http.StatusFound, "/")
	})

	// Route hapus data
	r.POST("/delete/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.String(http.StatusBadRequest, "Invalid ID")
			return
		}
		err = models.DeletePortofolio(id)
		if err != nil {
			c.String(http.StatusInternalServerError, "Error deleting data")
			return
		}
		c.Redirect(http.StatusFound, "/")
	})

	// Menjalankan server di port 8080
	r.Run(":8080")
}
