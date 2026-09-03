package logic

import (
	"styleguide/layering/httplayer"

	"github.com/shopspring/decimal"
)

// FIXME - нарушение границы слоев бизнес-логики
// Я могу переиспользовать эту функцию при переезде на GRPC? - Нет
// Сколько потребуется времени, чтобы переписать код на GRPC? - Много
func RunA(req *httplayer.Request) (*httplayer.Response, error) {
	// Допустим нам нужно только одно поле из Request в этой функции
	// Но нам в подарок достаются все поля. Спасибо.

	// ОГО!! Кто-то изменил переменную переданную как указатель
	req.Money = "0"

	var status string

	// Что если условно "завтра" эксперименты (тоглы) не будут приходить в параметре
	// Experiments. Нам нужно переписывать эту функцию? Да
	//
	if req.Experiments["A"] == "1" {
		status = "A"
	} else if req.Experiments["B"] == "1" {
		status = "B"
	}

	if req.Name == "A" {
		// Что может запуститься внутри в logic1?
		// Здесь надо запустить Stage1
		// Почему здесь появляется вариативность?
		logic1(req.Experiments)
	} else {
		logic2()
	}

	return &httplayer.Response{
		Success: true,
		Status:  status,
	}, nil
}

func RunB(req httplayer.Request) (any, error) {
	// Допустим нам нужно только одно поле из Request в этой функции
	// Но мы получаем и копируем весь объект Request.

	// Правильное ли это место для парсинга полей запроса?
	// Что если формат передачи денежных единиц поменяется?
	_, err := decimal.NewFromString(req.Money)

	if req.Name == "B" {
		// Здесь надо запустить Stage1 и Stage2 внутри logic1
		logic1(req.Experiments)
	} else {
		logic2()
	}

	return nil, err
}
