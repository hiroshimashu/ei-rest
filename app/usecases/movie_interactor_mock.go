package usecases

import "github.com/hiroshimashu/ei-rest/app/domain"

type MockMovieInteractor struct {
	MovieRepository MovieRepository
}

func (mockMovieInteractor *MockMovieInteractor) Index() (movies domain.Movies, err error) {
	movies, err = mockMovieInteractor.MovieRepository.FindAll()
	return
}

func (mockMovieInteractor *MockMovieInteractor) IndexByID(id string) (movie domain.Movie, err error) {
	movie, err = mockMovieInteractor.MovieRepository.FindByID(id)
	return
}

func (mockMovieInteractor *MockMovieInteractor) Store(newMovie domain.Movie) (movie domain.Movie, err error) {
	movie, err = mockMovieInteractor.MovieRepository.Save(newMovie)
	return
}

func NewMockMovieInteractor(mc *MockMovieRepository) *MockMovieInteractor {
	return &MockMovieInteractor{
		MovieRepository: mc,
	}
}

type MockMovieRepository struct {
	Movies domain.Movies
}

func (mr *MockMovieRepository) FindAll() (domain.Movies, error) {
	return mr.Movies, nil
}

func (mr *MockMovieRepository) FindByID(id string) (domain.Movie, error) {
	for _, v := range mr.Movies {
		if v.ID == id {
			return v, nil
		}
	}
	return domain.Movie{}, nil
}

func (mr *MockMovieRepository) Save(movie domain.Movie) (domain.Movie, error) {
	mr.Movies = append(mr.Movies, movie)
	return movie, nil
}

func NewMockRepository(movies domain.Movies) *MockMovieRepository {
	return &MockMovieRepository{
		Movies: movies,
	}
}
