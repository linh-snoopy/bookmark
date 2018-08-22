package repository

type UserRepository struct {
	NoSqlRepository
}

func (u UserRepository) InsertUser(email, password string) error  {
	return nil
}

func (u UserRepository) UpdatePassword(password string) error  {
	return nil
}

func (u UserRepository) DeleteUser(id int) error {
	return nil
}
