package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/dig"
	"net/http"
)

type Person struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Config struct {
	Enabled      bool
	DatabasePath string
	Port         string
}

func NewConfig() *Config {
	return &Config{
		Enabled:      true,
		DatabasePath: "./example.db",
		Port:         "8000",
	}
}

func ConnectionDatabase(config *Config) (*sql.DB, error) {
	//return sql.Open("", config.DatabasePath)
	url := "root:root@tcp(localhost:3306)/awesome_project?charset=utf8mb4&parseTime=True&loc=Local"
	fmt.Println(url)
	return sql.Open("mysql", url)
}

type PersonRepository struct {
	database *sql.DB
}

func (repository *PersonRepository) FindAll() []*Person {
	rows, err := repository.database.Query(`select uid, account, source from users_account;`)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer rows.Close()

	people := []*Person{}
	for rows.Next() {
		var (
			id   int
			name string
			age  int
		)

		rows.Scan(&id, &name, &age)
		people = append(people, &Person{
			Id:   id,
			Name: name,
			Age:  age,
		})
	}
	return people
}

func NewPeopleRepository(database *sql.DB) *PersonRepository {
	return &PersonRepository{
		database: database,
	}
}

type PersonService struct {
	config     *Config
	repository *PersonRepository
}

func (service *PersonService) FindAll() []*Person {
	if service.config.Enabled {
		return service.repository.FindAll()
	}
	return []*Person{}
}

func NewPersonService(config *Config, repository *PersonRepository) *PersonService {
	return &PersonService{
		config:     config,
		repository: repository,
	}
}

type Server struct {
	config        *Config
	personService *PersonService
}

func (s *Server) Handler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/people", s.people)
	return mux
}

func (s *Server) people(w http.ResponseWriter, r *http.Request) {
	people := s.personService.FindAll()
	bytes, _ := json.Marshal(people)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}

func (s *Server) Run() {
	httpServer := &http.Server{
		Addr:    ":" + s.config.Port,
		Handler: s.Handler(),
	}
	fmt.Println("Listen on ", s.config.Port)
	httpServer.ListenAndServe()
}

func NewServer(config *Config, service *PersonService) *Server {
	return &Server{
		config:        config,
		personService: service,
	}
}

func BuildContainer() *dig.Container {
	container := dig.New()
	container.Provide(NewConfig)
	container.Provide(ConnectionDatabase)
	container.Provide(NewPersonService)
	container.Provide(NewPeopleRepository)
	container.Provide(NewServer)
	return container
}

func main() {
	//config := NewConfig()
	//db, err := ConnectionDatabase(config)
	//fmt.Println(db.Ping())
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//personRepository := NewPeopleRepository(db)
	//personService := NewPersonService(config, personRepository)
	//server := NewServer(config, personService)
	//server.Run()

	container := BuildContainer()
	err := container.Invoke(func(server *Server) {
		server.Run()
	})

	if err != nil {
		panic(err)
	}
}
