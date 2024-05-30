package services

type RestfulParser struct{}

func (rp *RestfulParser) GetCurrentBlock() int {
	return 1
}
