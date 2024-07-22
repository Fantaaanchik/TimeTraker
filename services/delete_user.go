package services

func (s *UserService) DeleteUser(id int) error {
	return s.userRepository.DeleteUser(id)
}
