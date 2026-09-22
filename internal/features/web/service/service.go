package web_service

type WebService struct {
	webRespository WebRepository
}

type WebRepository interface {
	GetFile(filePath string) ([]byte, error)
}

func NewWebService(webRespository WebRepository) *WebService {
	return &WebService{
		webRespository: webRespository,
	}
}