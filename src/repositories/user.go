package repositories

import (
	"api/src/models"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Users struct {
	db *sql.DB
}

func NewUsersRepository(db *sql.DB) *Users {
	return &Users{db}
}

func (s Users) Create(User models.User) (*models.User, error) {

	Userment, err := s.db.Prepare(
		`INSERT INTO public.tb011_logins (tb011_logins, tb010_cpf, tb011_senha, tb011_data_cadastro) VALUES($1, $2, $3, $4);`,
	)
	if err != nil {

		return &models.User{}, err
	}
	User.Id = uuid.New()
	defer Userment.Close()
	_, erro := Userment.Exec(User.Id, User.Cpf, User.Password, time.Now())
	if erro != nil {

		return &models.User{}, erro
	}
	User.Password = ""
	return &User, nil

}
func (s Users) Find(UUID string) (models.User, error) {

	response, err := s.db.Query(`select * from public.tb011_logins where tb011_logins = $1;`, UUID)
	if err != nil {
		return models.User{}, err
	}
	defer response.Close()
	var User models.User

	if response.Next() {
		if err = response.Scan(
			&User.Id,
			&User.Cpf,
			&User.Password,
			&User.Create_at,
		); err != nil {
			return models.User{}, err
		}
	}
	User.Password = ""
	return User, nil

}
func (s Users) Update(UUID string, User models.User) (*models.User, error) {

	Userment, err := s.db.Prepare(
		`update public.tb011_logins set tb010_cpf = $1, tb011_senha = $2 where tb011_logins = $3`,
	)
	if err != nil {

		return &models.User{}, err
	}
	defer Userment.Close()
	_, erro := Userment.Exec(User.Cpf, User.Password, UUID)
	if erro != nil {

		return &models.User{}, err
	}

	User.Id = uuid.MustParse(UUID)
	User.Password = ""
	return &User, nil

}

func (s Users) Delete(UUDI string) error {

	Userment, err := s.db.Prepare(`delete from public.tb011_logins where tb011_logins = $1;`)
	if err != nil {
		return err
	}
	if err != nil {

		return err
	}
	defer Userment.Close()
	_, erro := Userment.Exec(UUDI)
	if erro != nil {

		return erro
	}
	return nil

}

func (s Users) FindCpf(Cpf string) (models.User, error) {
	response, err := s.db.Query(`select tb011_logins, tb011_senha from public.tb011_logins where tb010_cpf = $1;`, Cpf)
	if err != nil {
		return models.User{}, err
	}
	defer response.Close()
	var User models.User

	if response.Next() {
		if err = response.Scan(
			&User.Id,
			&User.Password,
		); err != nil {
			return models.User{}, err
		}
	}
	return User, nil
}
