package config

import (
	"github.com/spf13/viper"
)

// Config adalah struct untuk menampung semua konfigurasi dari file config.yaml
// Kita menggunakan 'mapstructure' tag agar Viper tahu cara mem-parsingnya.
type Config struct {
	Server struct {
		Port string `mapstructure:"port"`
	} `mapstructure:"server"`
	Database struct {
		Host     string `mapstructure:"host"`
		Port     string `mapstructure:"port"`
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		DBName   string `mapstructure:"dbname"`
	} `mapstructure:"database"`
}

// LoadConfig adalah fungsi untuk memuat konfigurasi dari file
func LoadConfig() (config Config, err error) {
	// 1. Tentukan path, nama, dan tipe file config
	viper.AddConfigPath(".")      // Cari di folder root
	viper.SetConfigName("config") // Nama file: "config" (akan mencari config.yaml)
	viper.SetConfigType("yaml")   // Tipe file: "yaml"

	// 2. (Sangat Penting) Otomatis membaca dari Environment Variables
	// Ini membuat config Anda siap untuk production/Docker
	viper.AutomaticEnv()

	// 3. Baca file konfigurasinya
	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	// 4. Masukkan (unmarshal) semua nilai config ke dalam struct 'config'
	err = viper.Unmarshal(&config)
	return
}
