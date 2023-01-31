package service

type VersionService struct {
}

func (s *VersionService) GetVersion() string {
	return "v0.1.0-alpha"
}
