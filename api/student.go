// package api

// import (
// 	"project_sdu/model"
// 	"project_sdu/service"
// 	"net/http"
// 	"strconv"

// 	"github.com/gin-gonic/gin"
// )

// type StudentAPI interface {
// 	FetchAllStudent(c *gin.Context)
// 	FetchStudentByID(c *gin.Context)
// 	StoreStudent(c *gin.Context)
// 	UpdateStudent(c *gin.Context)
// 	DeleteStudent(c *gin.Context)
// 	FetchStudentWithClass(c *gin.Context)
// }

// type studentAPI struct {
// 	studentService service.StudentService
// }

// func NewStudentAPI(studentService service.StudentService) *studentAPI {
// 	return &studentAPI{studentService}
// }

// // GET /students
// func (s *studentAPI) FetchAllStudent(c *gin.Context) {
// 	students, err := s.studentService.FetchAll()
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, students)
// }

// // GET /students/:id
// func (s *studentAPI) FetchStudentByID(c *gin.Context) {
// 	idStr := c.Param("id")
// 	id, err := strconv.Atoi(idStr)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Invalid student ID"})
// 		return
// 	}

// 	student, err := s.studentService.FetchByID(id)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, student)
// }

// // POST /students
// func (s *studentAPI) StoreStudent(c *gin.Context) {
// 	var student model.Student
// 	if err := c.ShouldBindJSON(&student); err != nil {
// 		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
// 		return
// 	}

// 	if err := s.studentService.Store(&student); err != nil {
// 		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, model.SuccessResponse{Message: "student berhasil ditambahkan"})
// }

// // PUT /students/:id
// func (s *studentAPI) UpdateStudent(c *gin.Context) {
// 	idStr := c.Param("id")
// 	id, err := strconv.Atoi(idStr)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Invalid student ID"})
// 		return
// 	}

// 	var student model.Student
// 	if err := c.ShouldBindJSON(&student); err != nil {
// 		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
// 		return
// 	}

// 	if err := s.studentService.Update(id, &student); err != nil {
// 		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, model.SuccessResponse{Message: "student berhasil diperbarui"})
// }

// // DELETE /students/:id
// func (s *studentAPI) DeleteStudent(c *gin.Context) {
// 	idStr := c.Param("id")
// 	id, err := strconv.Atoi(idStr)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Invalid student ID"})
// 		return
// 	}

// 	if err := s.studentService.Delete(id); err != nil {
// 		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, model.SuccessResponse{Message: "student berhasil dihapus"})
// }

// // GET /students/class
// func (s *studentAPI) FetchStudentWithClass(c *gin.Context) {
// 	studentClasses, err := s.studentService.FetchWithClass()
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, studentClasses)
// }

package api

import (
	"net/http"
	"project_sdu/model"
	"project_sdu/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type StudentAPI interface {
	FetchAllStudent(c *gin.Context)
	FetchStudentByID(c *gin.Context)
	StoreStudent(c *gin.Context)
	UpdateStudent(c *gin.Context)
	DeleteStudent(c *gin.Context)
	FetchStudentWithClass(c *gin.Context)
}

type studentAPI struct {
	studentService service.StudentService
}

func NewStudentAPI(studentService service.StudentService) *studentAPI {
	return &studentAPI{studentService}
}

// GET /students
func (s *studentAPI) FetchAllStudent(c *gin.Context) {
	students, err := s.studentService.FetchAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Success: false,
			Status:  http.StatusInternalServerError,
			Message: "Failed to fetch students",
			Errors:  map[string]string{"server": err.Error()},
		})
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse{
		Success: true,
		Status:  http.StatusOK,
		Message: "Students retrieved successfully",
		Data:    students,
	})
}

// GET /students/:id
func (s *studentAPI) FetchStudentByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Success: false,
			Status:  http.StatusBadRequest,
			Message: "Invalid student ID",
			Errors:  map[string]string{"id": "must be a valid number"},
		})
		return
	}

	student, err := s.studentService.FetchByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Success: false,
			Status:  http.StatusInternalServerError,
			Message: "Failed to fetch student",
			Errors:  map[string]string{"server": err.Error()},
		})
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse{
		Success: true,
		Status:  http.StatusOK,
		Message: "Student retrieved successfully",
		Data:    student,
	})
}

// POST /students
func (s *studentAPI) StoreStudent(c *gin.Context) {
	var student model.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Success: false,
			Status:  http.StatusBadRequest,
			Message: "Invalid input data",
			Errors:  map[string]string{"body": err.Error()},
		})
		return
	}

	if err := s.studentService.Store(&student); err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Success: false,
			Status:  http.StatusInternalServerError,
			Message: "Failed to create student",
			Errors:  map[string]string{"server": err.Error()},
		})
		return
	}

	c.JSON(http.StatusCreated, model.SuccessResponse{
		Success: true,
		Status:  http.StatusCreated,
		Message: "Student created successfully",
		Data:    student,
	})
}

// PUT /students/:id
func (s *studentAPI) UpdateStudent(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Success: false,
			Status:  http.StatusBadRequest,
			Message: "Invalid student ID",
			Errors:  map[string]string{"id": "must be a valid number"},
		})
		return
	}

	var student model.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Success: false,
			Status:  http.StatusBadRequest,
			Message: "Invalid input data",
			Errors:  map[string]string{"body": err.Error()},
		})
		return
	}

	if err := s.studentService.Update(id, &student); err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Success: false,
			Status:  http.StatusInternalServerError,
			Message: "Failed to update student",
			Errors:  map[string]string{"server": err.Error()},
		})
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse{
		Success: true,
		Status:  http.StatusOK,
		Message: "Student updated successfully",
		Data:    student,
	})
}

// DELETE /students/:id
func (s *studentAPI) DeleteStudent(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Success: false,
			Status:  http.StatusBadRequest,
			Message: "Invalid student ID",
			Errors:  map[string]string{"id": "must be a valid number"},
		})
		return
	}

	if err := s.studentService.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Success: false,
			Status:  http.StatusInternalServerError,
			Message: "Failed to delete student",
			Errors:  map[string]string{"server": err.Error()},
		})
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse{
		Success: true,
		Status:  http.StatusOK,
		Message: "Student deleted successfully",
	})
}

// GET /students/class
func (s *studentAPI) FetchStudentWithClass(c *gin.Context) {
	studentClasses, err := s.studentService.FetchWithClass()
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Success: false,
			Status:  http.StatusInternalServerError,
			Message: "Failed to fetch students with class",
			Errors:  map[string]string{"server": err.Error()},
		})
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse{
		Success: true,
		Status:  http.StatusOK,
		Message: "Students with class retrieved successfully",
		Data:    studentClasses,
	})
}
