package telegram

const msgHelp = `Я(бот) могу показать содержимое директории в которой пребываю.
Но только тем ребятам у кого есть доспут.
Все ваши действия записываются.
Внимание! Аутентификацию нужно проходить кажды раз!`

const msgHello = "Привет! 👾\n\n" + msgHelp

const (
	msgUnknownCommand = "Нет такой комманды 🤔"
	msgNoSavedPages   = "You have no saved pages 🙊"
	msgSaved          = "Saved! 👌"
	msgAlreadyExists  = "You have already have this page in your list 🤗"
)
