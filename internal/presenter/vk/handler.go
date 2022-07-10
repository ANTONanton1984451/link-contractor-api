package vk

type handlerPresenter struct{}

func (hp *handlerPresenter) SomethingWentWrong() []byte {
	return []byte(somethingWentWrong())
}

func NewHandlerPresenter() *handlerPresenter {
	return &handlerPresenter{}
}
