package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"sync"
	"time"
)

/*
Реализовать HTTP-сервер для работы с календарем. В рамках задания необходимо работать строго со стандартной HTTP-библиотекой.

В рамках задания необходимо:
Реализовать вспомогательные функции для сериализации объектов доменной области в JSON.
Реализовать вспомогательные функции для парсинга и валидации параметров методов /create_event и /update_event.
Реализовать HTTP обработчики для каждого из методов API, используя вспомогательные функции и объекты доменной области.
Реализовать middleware для логирования запросов


Методы API:
POST /create_event
POST /update_event
POST /delete_event
GET /events_for_day
GET /events_for_week
GET /events_for_month

Параметры передаются в виде www-url-form-encoded (т.е. обычные user_id=3&date=2019-09-09). В GET методах параметры передаются через queryString, в POST через тело запроса.
В результате каждого запроса должен возвращаться JSON-документ содержащий либо {"result": "..."} в случае успешного выполнения метода, либо {"error": "..."} в случае ошибки бизнес-логики.

В рамках задачи необходимо:
Реализовать все методы.
Бизнес логика НЕ должна зависеть от кода HTTP сервера.
В случае ошибки бизнес-логики сервер должен возвращать HTTP 503. В случае ошибки входных данных (невалидный int например) сервер должен возвращать HTTP 400. В случае остальных ошибок сервер должен возвращать HTTP 500. Web-сервер должен запускаться на порту указанном в конфиге и выводить в лог каждый обработанный запрос.
*/

var calendar = sync.Map{}

var portNumber = ":8087"

type Event struct {
	ID    string    `json:"id"`
	Title string    `json:"title"`
	Date  time.Time `json:"date"`
}

func create(w http.ResponseWriter, r *http.Request) {

	id := r.FormValue("id")
	title := r.FormValue("title")
	date, err := time.Parse("2000-01-01", r.FormValue("date"))
	if err != nil {
		log.Fatal("Error date")
		return
	}

	event := Event{
		ID:    id,
		Title: title,
		Date:  date,
	}

	fmt.Println(id)

	calendar.Store(id, event)

}

func update(w http.ResponseWriter, r *http.Request) {

	id := r.FormValue("id")
	title := r.FormValue("title")
	date, err := time.Parse("2000-01-01", r.FormValue("date"))
	if err != nil {
		log.Fatal("Error date")
		return
	}

	event := Event{
		ID:    id,
		Title: title,
		Date:  date,
	}

	fmt.Println(id)

	calendar.Store(id, event)
}

func delete(w http.ResponseWriter, r *http.Request) {

	id := r.FormValue("id")
	fmt.Println(id)
	calendar.Delete(id)
}

func getDay(w http.ResponseWriter, r *http.Request) {

	result := make(map[string][]Event)

	calendar.Range(func(k, v any) bool {
		e, ok := v.(Event)
		if !ok {
			return true
		}

		date := e.Date.Format("2006-01-02")

		if _, ok := result[date]; !ok {
			result[date] = make([]Event, 0)
		}

		result[date] = append(result[date], Event{
			ID:    e.ID,
			Title: e.Title,
			Date:  e.Date,
		})

		return true
	})
}

func getWeek(w http.ResponseWriter, r *http.Request) {

}

func getMonth(w http.ResponseWriter, r *http.Request) {

}

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Test Page"))
}

func handleRequest() {

	rtr := mux.NewRouter()

	rtr.HandleFunc("/", index)
	rtr.HandleFunc("/create", create).Methods("POST")
	rtr.HandleFunc("/update", update).Methods("POST")
	rtr.HandleFunc("/delete", delete).Methods("POST")
	rtr.HandleFunc("/get_day", getDay).Methods("GET")
	rtr.HandleFunc("/get_week", getWeek).Methods("GET")
	rtr.HandleFunc("/get_month", getMonth).Methods("GET")

	//rtr.HandleFunc("/order", OrderHandler(conn, cache))
	//rtr.HandleFunc("/store", StoreHandler(conn, cache)).Methods("POST")

	fmt.Printf("Starting application on port %v\n", portNumber)
	http.ListenAndServe(portNumber, rtr)
}

func main() {
	handleRequest()
}
