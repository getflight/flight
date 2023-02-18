package service

type VersionService struct {
}

func (s *VersionService) GetVersion() string {
	return "v0.1.1-alpha"
}
