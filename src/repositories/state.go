package repositories

import (
	"api/src/models"
	"database/sql"
)

type States struct {
	db *sql.DB
}

func NewStatesRepository(db *sql.DB) *States {
	return &States{db}
}

func (s States) Create(state models.State) (*models.State, error) {

	statement, err := s.db.Prepare(
		`INSERT INTO public.tb001_uf (tb001_sigla_uf, tb001_nome_estado) VALUES($1, $2);`,
	)
	if err != nil {

		return &models.State{}, err
	}
	defer statement.Close()
	_, erro := statement.Exec(state.Uf, state.Name)
	if erro != nil {

		return &models.State{}, erro
	}

	return &models.State{
		Uf:   state.Uf,
		Name: state.Name,
	}, nil

}
func (s States) Find(uf string) (models.State, error) {

	response, err := s.db.Query(`select * from public.tb001_uf where tb001_sigla_uf = $1;`, uf)
	if err != nil {
		return models.State{}, err
	}
	defer response.Close()
	var state models.State

	if response.Next() {
		if err = response.Scan(
			&state.Uf,
			&state.Name,
		); err != nil {
			return models.State{}, err
		}
	}

	return state, nil

}
func (s States) Update(UF string, State models.State) (*models.State, error) {

	statement, err := s.db.Prepare(
		`update tb001_uf set tb001_sigla_uf = $1, tb001_nome_estado = $2 where tb001_sigla_uf = $3`,
	)
	if err != nil {

		return &models.State{}, err
	}
	defer statement.Close()
	_, erro := statement.Exec(State.Uf, State.Name, UF)
	if erro != nil {

		return &models.State{}, err
	}
	return &State, nil

}

func (s States) Show() ([]models.State, error) {

	lines, err := s.db.Query(`SELECT * FROM public.tb001_uf;`)
	if err != nil {
		return nil, err
	}
	defer lines.Close()
	var states []models.State

	for lines.Next() {
		var state models.State

		if err = lines.Scan(
			&state.Uf,
			&state.Name,
		); err != nil {
			return nil, err
		}

		states = append(states, state)
	}

	return states, nil

}

func (s States) Delete(UF string) error {

	statement, err := s.db.Prepare(`delete from public.tb001_uf where tb001_sigla_uf = $1;`)
	if err != nil {
		return err
	}
	if err != nil {

		return err
	}
	defer statement.Close()
	_, erro := statement.Exec(UF)
	if erro != nil {

		return erro
	}
	return nil

}
