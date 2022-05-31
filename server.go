package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"example/grpc_sql/sqlpb"

	"google.golang.org/grpc"

	"database/sql"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

type server struct{}

func (*server) GetCompanyById(ctx context.Context, req *sqlpb.CompanyId) (*sqlpb.Company, error) {
	fmt.Println("GetCompanyById function is Invoked!")
	id := req.GetID()
	fmt.Println("Given Id is:" + id)
	return nil, nil
}

func main() {
	fmt.Println("Get Ready! Starting the Server")
	lis, err1 := net.Listen("tcp", "0.0.0.0:50051")
	if err1 != nil {
		log.Fatalf("Failed to listen: %v", err1)
	}

	// ------------------------------------------------------------------------
	// Database
	// Capture connection properties.
	cfg := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),

		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "companies",
		Params:               map[string]string{},
		AllowNativePasswords: true,
	}
	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected to the database!")
	//--------------------------------------------------------------------------

	opts := []grpc.ServerOption{}

	s := grpc.NewServer(opts...)
	sqlpb.RegisterSqlServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve : %v", err)
	}
}
