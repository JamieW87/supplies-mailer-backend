package service

import "fmt"

func (s *Service) StoreUserData() error {

	userId, err := s.db.InsertUser("", "", "")
	fmt.Println(userId)
	return err
}
