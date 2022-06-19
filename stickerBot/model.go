package main

type Updates struct {
	UpdateId int			`json:"update_id"`
	Message Message			`json:"message"`
}

type Message struct {
	Chat Chat				`json:"chat"`
	Text string				`json:"text"`

	Sticker Sticker			`json:"sticker"`
}

type Chat struct {
	ChatId    int    `json:"id"`
}

type RestResponse struct {
	Result []Updates		`json:"result"`
}

type BotMessage struct {
	ChatId int				`json:"chat_id"`
	Text   string			`json:"text"`

	Sticker Sticker			`json:"sticker"`
}


type Sticker struct {
	SetName	string			`json:"set_name"`
	FileId	string			`json:"file_id"`
}