package app

import (
	"encoding/json"
	"fmt"
	"net/http"

	"styleguide/layering/app/logic"
	"styleguide/layering/httplayer"
)

type App struct {
}

func New() App {
	return App{}
}

func (a App) Start() {
	http.HandleFunc("/A", func(w http.ResponseWriter, r *http.Request) {
		var req httplayer.Request

		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			w.Write([]byte("error"))
			return
		}

		// Вопрос - какие поля запроса используются в этой функции?
		logic.RunA(&req)

		// Ожидаем что запрос неизменяемый, но увы!
		logic.RunB(req)

		fmt.Printf("%#v\n", req)

		w.Write([]byte("ok"))
	})

	// Другой Endpoint хочет использовать бизнес-логику несколько другим способом
	// Контракт endpoint не соотвествует объект httplayer.Request.
	// Объект httplayer.Request диктуется нам бизнес-логикой?
	// Нужно писать функцию конвертации?
	http.HandleFunc("/B", func(w http.ResponseWriter, r *http.Request) {
		var req httplayer.Request

		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			w.Write([]byte("error"))
			return
		}

		logic.RunB(req)

		w.Write([]byte("ok"))
	})

	_ = http.ListenAndServe(":8000", nil)
}
