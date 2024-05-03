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

	// route grouping
	v1 := r.Group("/v1", middleware.Auth())
	{
		v1.POST("/data", handler.CreateData) // input data diri baru
		v1.GET("/data", handler.GetData)     // dapatkan list data diri

		// grouping
		v1Role := v1.Group("/role", middleware.CheckingRole())
		{
			v1Role.POST("/data", handler.CreateData) // v1/role/data  // input data diri baru
			v1Role.GET("/data", handler.GetData)     //  v1/role/data  dapatkan list data diri

			// get role, bukan admin
			// phone number: 081xxxxxxxxxxxxx
			// 							0 2

			// kalau rolenya admin
			// munculkan semua phone number 081232131313123

		}

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
