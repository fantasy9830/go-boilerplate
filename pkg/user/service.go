package user

// Service Service
type Service struct {
	rep *Repository
}

// NewService New Service
func NewService(repository *Repository) *Service {
	return &Service{
		rep: repository,
	}
}
