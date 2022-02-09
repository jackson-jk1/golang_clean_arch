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
	defer Userment.Close()
	_, erro := Userment.Exec(uuid.New(), User.Cpf, User.Password, time.Now())
	if erro != nil {

		return &models.User{}, erro
	}
	User.Password = ""
	return &User, nil

}
func (s Users) Find(uf string) (models.User, error) {

	response, err := s.db.Query(`select * from public.tb001_uf where tb001_sigla_uf = $1;`, uf)
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

	return User, nil

}
func (s Users) Update(UF string, User models.User) (*models.User, error) {

	Userment, err := s.db.Prepare(
		`update tb001_uf set tb001_sigla_uf = $1, tb001_nome_estado = $2 where tb001_sigla_uf = $3`,
	)
	if err != nil {

		return &models.User{}, err
	}
	defer Userment.Close()
	_, erro := Userment.Exec(User.Cpf, User.Password)
	if erro != nil {

		return &models.User{}, err
	}
	return &User, nil

}

func (s Users) Delete(UF string) error {

	Userment, err := s.db.Prepare(`delete from public.tb001_uf where tb001_sigla_uf = $1;`)
	if err != nil {
		return err
	}
	if err != nil {

		return err
	}
	defer Userment.Close()
	_, erro := Userment.Exec(UF)
	if erro != nil {

		return erro
	}
	return nil

}
