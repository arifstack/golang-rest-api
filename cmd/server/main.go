package main

import (
	"go1/config"
	"go1/internal/app"
	"log"
)

func main() {
	// Pilih konfigurasi yang sesuai berdasarkan lingkungan
	var cfg config.Config
	if isProdEnv() {
		cfg = config.LoadProdConfig()
	} else {
		cfg = config.LoadDevConfig()
	}

	server := app.NewServer(cfg)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func isProdEnv() bool {
	// Fungsi untuk menentukan jika lingkungan adalah production
	// Anda bisa menyesuaikan metode ini sesuai kebutuhan Anda
	return true // Misalnya, diset true jika ingin menjalankan di production
}
