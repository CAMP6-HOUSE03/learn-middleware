package main

import (
	"example-middleware/handler"
	"example-middleware/middleware"
	"example-middleware/repository"
	"example-middleware/service"

	"github.com/gin-gonic/gin"
)

func main() {
	// router

	// wiring contract
	repo := repository.NewRepository()
	svc := service.NewService(repo)
	handler := handler.NewHandler(svc)

	// engine
	r := gin.Default()

	// route
	v1 := r.Group("/v1", middleware.Auth(), middleware.Auth2())
	{
		// middleware
		// 1 auth
		// 2 monitoring
		// 3 log
		// 4 auth
		// v1.Use(middleware.Auth(), middleware.Auth2())

		// 3, 1 detik middleware 1, 1 detik 2, 1 detik handle

		// 500ms
		// promoteus, log monitoring
		// google analytics
		//

		// goroutine, proses data2

		v1.POST("/data", handler.CreateData) // input data diri baru
		v1.GET("/data", handler.GetData)     // dapatkan list data diri
		// v1.GET("/data/:id", handler.GetDataById)   // mendapatkan data diri berdasarkan no identitas
		// v1.PUT("/data/:id", handler.UpdateData)    // update data diri baru
		// v1.DELETE("/data/:id", handler.DeleteData) // delete data diri
	}

	r.Run() // :8080
}

// aplikasi sederhana
// aku gunakan untuk data diri
// no identitas, nama, alamat, no hp, email

// input data diri baru
// update data diri baru
// dapatkan list data diri
// mendapatkan data diri berdasarkan no identitas
// delete data diri

// implementasi clean arsitektur
// handler / api /controller
// service / usecase
// repository
// model

// middleware
