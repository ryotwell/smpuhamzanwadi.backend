package main

import (
	"fmt"

	"os"
	"time"

	"project_sdu/api"
	"project_sdu/db"
	"project_sdu/middleware"
	"project_sdu/model"
	repo "project_sdu/repository"
	"project_sdu/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"

	// "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type APIHandler struct {
	UserAPIHandler    api.UserAPI
	StudentAPIHandler api.StudentAPI
}

func main() {
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()

	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("[%s] \"%s %s %s\"\n",
			param.TimeStamp.Format(time.RFC822),
			param.Method,
			param.Path,
			param.ErrorMessage,
		)
	}))
	router.Use(gin.Recovery())

	// --- ✅ CORS SETUP HERE ---
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Ganti dengan origin frontend kamu, misal "https://project-frontend.vercel.app"
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	// --- END CORS ---

	// // Koneksi ke DB Postgres
	// dbCredential := model.Credential{
	// 	Host:         "localhost",
	// 	Username:     "postgres",
	// 	Password:     "farid123",
	// 	DatabaseName: "tes_sdu",
	// 	Port:         5432,
	// 	Schema:       "public",
	// }

	// database := db.NewDB()
	// conn, err := database.Connect(&dbCredential)
	// if err != nil {
	// 	panic(err)
	// }

	// // Migrasi tabel
	// conn.AutoMigrate(&model.User{}, &model.Student{})

	// // Daftarkan semua route dan handler
	// router = RunServer(router, conn)

	// fmt.Println("✅ Server is running on port 8080")
	// if err := router.Run(":8080"); err != nil {
	// 	panic(err)
	// }

	// Ambil DATABASE_URL dari environment (Railway akan menyediakannya otomatis)
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		panic("DATABASE_URL tidak ditemukan. Pastikan environment variable sudah diatur di Railway.")
	}

	// Koneksi ke DB
	database := db.NewDB()
	conn, err := database.ConnectURL(databaseURL)
	if err != nil {
		panic(err)
	}

	// Migrasi tabel (opsional, tergantung kebutuhan)
	conn.AutoMigrate(&model.User{}, &model.Student{})

	// Daftarkan semua route dan handler
	router = RunServer(router, conn)

	// Ambil PORT dari Railway
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // fallback jika dijalankan lokal
	}

	fmt.Printf("✅ Server is running on port %s\n", port)
	if err := router.Run(":" + port); err != nil {
		panic(err)
	}

}

func RunServer(r *gin.Engine, conn interface{}) *gin.Engine {
	// Pastikan conn bertipe *gorm.DB
	dbConn := conn.(*gorm.DB)

	// Inisialisasi Repository (pakai GORM)
	userRepo := repo.NewUserRepo(dbConn)
	studentRepo := repo.NewStudentRepo(dbConn)

	// Inisialisasi Service
	userService := service.NewUserService(userRepo)
	studentService := service.NewStudentService(studentRepo)

	// Inisialisasi API Handler
	userAPIHandler := api.NewUserAPI(userService)
	studentAPIHandler := api.NewStudentAPI(studentService)

	apiHandler := APIHandler{
		UserAPIHandler:    userAPIHandler,
		StudentAPIHandler: studentAPIHandler,
	}

	// ROUTES

	// User routes
	user := r.Group("/user")
	{
		user.POST("/register", apiHandler.UserAPIHandler.Register)
		user.POST("/login", apiHandler.UserAPIHandler.Login)
		user.POST("/logout", apiHandler.UserAPIHandler.Logout)

		user.Use(middleware.Auth())
		user.GET("/tasks", apiHandler.UserAPIHandler.GetUserTaskCategory)
	}

	// Student routes
	students := r.Group("/students")
	{
		students.Use(middleware.Auth())
		students.GET("", apiHandler.StudentAPIHandler.FetchAllStudent)
		students.GET("/:id", apiHandler.StudentAPIHandler.FetchStudentByID)
		students.POST("", apiHandler.StudentAPIHandler.StoreStudent)
		students.PUT("/:id", apiHandler.StudentAPIHandler.UpdateStudent)
		students.DELETE("/:id", apiHandler.StudentAPIHandler.DeleteStudent)
		students.GET("/class", apiHandler.StudentAPIHandler.FetchStudentWithClass)
	}

	return r
}
