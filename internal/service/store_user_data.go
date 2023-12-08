package service

func (s *Service) StoreUserData() error {

	return s.db.InsertUser("", "", "")
}
