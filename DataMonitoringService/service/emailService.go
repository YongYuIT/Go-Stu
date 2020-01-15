package service

type EmailService struct {
}

func (thiz *EmailService) sendSimpleSchaInfoEmail(scha_info string) {

	pService := PrintService{}
	linFileName := pService.PrintSchaPolylines(scha_info)

}
