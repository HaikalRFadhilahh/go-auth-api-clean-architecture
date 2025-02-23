package main

import (
	"fmt"

	"github.com/HaikalRFadhilahh/go-auth-api-clean-architecture/internal/server"
	"github.com/joho/godotenv"
)

func main() {
	// Read Environment Variable
	if err := godotenv.Load(".env"); err != nil {
		fmt.Println("Tidak Dapat Membaca File .env")
	}

	// Start Server
	server.NewAPIServer().Run()
}
