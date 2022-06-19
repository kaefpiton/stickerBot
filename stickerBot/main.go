package main

//Телеграм-бот для того, чтобы узнать id стикера

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

//Точка входа
func main()  {
	//todo заменить и вынести в конфиг
	botToken := "1987224543:AAEupAuWDkewgP_l21BEqlRjB4mSJgT6Hdk"
	botAPI := "https://api.telegram.org/bot"
	botUrl := botAPI + botToken

	offset := 0
	for ;;{
		updates ,err := getUpdates(botUrl, offset)
		if err != nil{
			log.Println("Something was wrong - ", err.Error())
		}

		for _, updates := range updates{
			err = respond(botUrl, updates)
			offset = updates.UpdateId + 1
		}

	}
}

//Запрос обновлений
func getUpdates(botURL string, offset int)([]Updates, error){
	resp, err := http.Get(botURL + "/getUpdates" + "?offset=" + strconv.Itoa(offset))
	if err != nil {
		return nil, err
	}
	defer  resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil{
		return nil, err
	}

	var restResponse  RestResponse

	err = json.Unmarshal(body, &restResponse)
	if err != nil{
		return nil, err
	}

	return restResponse.Result, nil
}

//Ответ на обновления
func respond(botUrl string, update Updates) error  {
	var botMessage BotMessage
	
	botMessage.ChatId = update.Message.Chat.ChatId
	text := update.Message.Text

	if update.Message.Sticker.FileId != ""{
		fmt.Println("Sticker ID -> "  + update.Message.Sticker.FileId)
		botMessage.Text = "ID вашего стикера " + update.Message.Sticker.FileId
		return  sendBotMessageText(botMessage,botUrl)
	}


	switch text {
	case "/start":{
		botMessage.Text = fmt.Sprintf("Здравствуйте! Введите стикер, чтобы получить ID!")
		return  sendBotMessageText(botMessage,botUrl)
	}

	default:{
		botMessage.Text = "Вы ввели не стикер"
		return  sendBotMessageText(botMessage,botUrl)
	}
	}

}

func sendBotMessageText(botMessage BotMessage, botUrl string) error {
	buf, err := json.Marshal(botMessage)
	if err != nil{
		return err
	}
	_,err = http.Post(botUrl + "/sendMessage", "application/json", bytes.NewBuffer(buf))

	if err != nil{
		return err
	}
	return nil
}
