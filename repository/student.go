package repository

import (
	"project_sdu/model"
	// "fmt"

	"gorm.io/gorm"
)

type StudentRepository interface {
	FetchAll() ([]model.Student, error)
	FetchByID(id int) (*model.Student, error)
	Store(s *model.Student) error
	Update(id int, s *model.Student) error
	Delete(id int) error
	FetchWithClass() (*[]model.StudentClass, error)
}

type studentRepoImpl struct {
	db *gorm.DB
}

func NewStudentRepo(db *gorm.DB) *studentRepoImpl {
	return &studentRepoImpl{db}
}

func (s *studentRepoImpl) FetchAll() ([]model.Student, error) {
	var students []model.Student
	if err := s.db.Find(&students).Error; err != nil {
		return nil, err
	}

	return students, nil // TODO: replace this
}

func (s *studentRepoImpl) Store(student *model.Student) error {
	if err := s.db.Create(&student).Error; err != nil {
		return err
	}

	return nil // TODO: replace this
}

func (s *studentRepoImpl) Update(id int, student *model.Student) error {
	if err := s.db.Model(&model.Student{}).Where("id = ?", id).Updates(student).Error; err != nil {
		return err
	}

	return nil // TODO: replace this
}

func (s *studentRepoImpl) Delete(id int) error {
	if err := s.db.Delete(&model.Student{}, id).Error; err != nil {
		return err
	}
	return nil // TODO: replace this
}

func (s *studentRepoImpl) FetchByID(id int) (*model.Student, error) {
	var student model.Student
	if err := s.db.First(&student, id).Error; err != nil {
		return nil, err
	}

	return &student, nil // TODO: replace this
}

func (s *studentRepoImpl) FetchWithClass() (*[]model.StudentClass, error) {
	studentClass := []model.StudentClass{}
	if err := s.db.Table("students").
		Select("students.name, students.address, classes.name AS class_name, classes.professor, classes.room_number").
		Joins("JOIN classes ON students.class_id = classes.id").
		Scan(&studentClass).
		Error; err != nil {
		return nil, err
	}

	return &studentClass, nil // TODO: replace this
}
