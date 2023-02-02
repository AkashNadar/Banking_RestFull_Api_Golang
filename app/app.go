package app

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/banking/domain"
	"github.com/banking/service"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

func Start() {

	// Wiring
	// handler needs service and service needs repo
	// ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	dbClient := getClient()
	customerRepo := domain.NewCustomerRepositoryDB(dbClient)
	accountRepo := domain.NewAccountRepositoryDB(dbClient)
	// accountRepo := domain.NewAccountRepositoryDB(dbClient)
	ch := CustomerHandlers{service.NewCustomerService(customerRepo)}
	accSer := service.NewAccountService(&accountRepo)
	ah := AccountHandler{&accSer}
	// mux := http.NewServeMux()
	// router := mux.NewRouter()

	// Routes
	// router.HandleFunc("/greet", greet).Methods(http.MethodGet)
	// router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)

	// Server
	// log.Fatal(http.ListenAndServe("localhost:9000", router))

	router := gin.Default()
	router.GET("/customers", ch.getAllCustomers)
	router.GET("/customers/:id", ch.getCustomerById)
	router.POST("/customers/:id/account", ah.NewAccount)
	router.POST("/customers/:id/account/:accountId", ah.MakeTransaction)
	router.Run("localhost:9000")
}

func sanityCheck() {
	if os.Getenv("DB_USER") == "" || os.Getenv("DB_PASSWORD") == "" || os.Getenv("DB_HOST") == "" || os.Getenv("DB_PORT") == "" || os.Getenv("DB_DBNAME") == "" {
		log.Fatal("Environment variables not found")
	}
}

func getDSN() string {
	fmt.Println("getDsn")
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file :%s", err.Error())
	}
	sanityCheck()
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_DBNAME")

	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, pass, host, port, dbname)
	// return "root:password@tcp(localhost:"
}

func getClient() *sqlx.DB {
	dsn := getDSN()
	fmt.Println(dsn)
	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return db
}
