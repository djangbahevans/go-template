package inmemory

type InMemoryDB struct {
	*userRepository
}

func NewInMemoryDB() *InMemoryDB {
	return &InMemoryDB{
		userRepository: NewUserRepository(),
	}
}
