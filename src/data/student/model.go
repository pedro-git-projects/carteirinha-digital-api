package student

import (
	"database/sql"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Model struct {
	DB *sql.DB
}

func (m Model) Insert(s *Student) error {
	query := `
		INSERT INTO students (academic_register, name, sex, hash, role, parent_id)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, created_at
	`
	args := []any{s.AcademicRegister, s.Name, s.Sex, s.Hash, "student", s.ParentID}
	return m.DB.QueryRow(query, args...).Scan(&s.ID, &s.CreatedAt)
}

func (m Model) AuthenticateStudent(academicRegister, password string) (*Student, error) {
	query := `
        SELECT id, created_at, academic_register, name, sex, hash, role, parent_id
        FROM students
        WHERE academic_register = $1
    `

	student := Student{}
	err := m.DB.QueryRow(query, academicRegister).Scan(
		&student.ID,
		&student.CreatedAt,
		&student.AcademicRegister,
		&student.Name,
		&student.Sex,
		&student.Hash,
		&student.Role,
		&student.ParentID,
	)

	if err == sql.ErrNoRows {
		return nil, errors.New("Student not found")
	} else if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(student.Hash), []byte(password))
	if err != nil {
		return nil, errors.New("Invalid password")
	}

	return &student, nil
}

// func (m Model) Get(id int64) (*Student, error) {
// 	query := `
// 		SELECT id, created_at, username, role
// 		FROM users
// 		WHERE id = $1
// 	`
//
// 	student := Student{}
// 	err := m.DB.QueryRow(query, id).Scan(
// 		&student.id,
// 		&student.createdAt,
// 		&student.username,
// 		&student.role,
// 	)
//
// 	if err == sql.ErrNoRows {
// 		return nil, errors.New("Student not found")
// 	} else if err != nil {
// 		return nil, err
// 	}
//
// 	return &student, nil
// }

// func (m UserModel) Update(id int64, updatedUser *User) (*User, error) {
// 	existingUser, err := m.Get(id)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	if updatedUser.Email != "" {
// 		existingUser.Email = updatedUser.Email
// 	}
// 	if updatedUser.Password != "" {
// 		existingUser.Password = updatedUser.Password
// 	}
// 	if updatedUser.Name != "" {
// 		existingUser.Name = updatedUser.Name
// 	}
// 	if updatedUser.Role != "" {
// 		existingUser.Role = updatedUser.Role
// 	}
// 	if updatedUser.ImageURL != "" {
// 		existingUser.ImageURL = updatedUser.ImageURL
// 	}
//
// 	query := `
// 		UPDATE users
// 		SET email = $1, password = $2, name = $3, role = $4, image_url = $5
// 		WHERE id = $6
// 		RETURNING id, email, name, role, image_url, created_at
// 	`
//
// 	err = m.DB.QueryRow(query, existingUser.Email, existingUser.Password, existingUser.Name, existingUser.Role, existingUser.ImageURL, id).
// 		Scan(
// 			&existingUser.ID,
// 			&existingUser.Email,
// 			&existingUser.Name,
// 			&existingUser.Role,
// 			&existingUser.ImageURL,
// 			&existingUser.CreatedAt,
// 		)
//
// 	if err != nil {
// 		return nil, errors.New(fmt.Sprintf("Atualização falhou com erro: %v", err))
// 	}
//
// 	return existingUser, nil
// }
//
// func (m UserModel) Delete(id int64) error {
// 	query := `
// 		DELETE FROM users
// 		WHERE id = $1
// 	`
//
// 	result, err := m.DB.Exec(query, id)
// 	if err != nil {
// 		return errors.New(fmt.Sprintf("Falha ao excluir usuário: %v", err))
// 	}
//
// 	rowsAffected, err := result.RowsAffected()
// 	if err != nil {
// 		return errors.New(fmt.Sprintf("Falha ao obter número de linhas afetadas: %v", err))
// 	}
//
// 	if rowsAffected == 0 {
// 		return errors.New("Usuário não encontrado")
// 	}
//
// 	return nil
// }
