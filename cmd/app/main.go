package main

import (
	"context"
	_ "github.com/prerec/media-mine/docs"
	"github.com/prerec/media-mine/internal/handler"
	"github.com/prerec/media-mine/internal/server"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"syscall"
)

//	@title			Exchange App API
//	@version		1.0
//	@description	API server for Exchange Application
//	@termsOfService	http://swagger.io/terms/

// @host		localhost:8080
// @BasePath	/
func main() {

	if err := initConfig(); err != nil {
		logrus.Fatalf("error init confgis: %s", err.Error())
	}

	h := handler.NewHandler()
	s := server.NewServer()

	go func() {
		if err := s.Run(viper.GetString("host"), viper.GetString("port"), h.InitRoutes()); err != nil {
			logrus.Fatalf("error while running server: %s:", err.Error())
		}
	}()

	logrus.Printf("server started on http://%s:%s/", viper.GetString("host"), viper.GetString("port"))

	// Заводим канал для ожидания сигнала на остановку от ОС
	quit := make(chan os.Signal, 1)
	// Регистрируем обработчик сигналов ОС
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	// Блокируемся в main горутине тк наш сервер запущен в отдельной горутине(чтобы приложение не вылетало
	// сразу же при запуске) до тех пор, пока в канал quit не поступит сигнал от ОС (после этого main горутина
	// продолжит свою работу)
	<-quit

	logrus.Print("Shutdown Server ...")

	if err := s.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error while shutting down server: %s", err.Error())
	}

}

// initConfig мапит пути к файлу с конфигом
func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
