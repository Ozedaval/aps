package main

import (
	"fmt"
	"github.com/spf13/viper"
	"github/hanzhang2566/asp/global"
	"github/hanzhang2566/asp/global/constant"
	"github/hanzhang2566/asp/internal/router"
	"log"
	"net/http"
)

func main() {
	r := router.NewRouter()
	server := &http.Server{
		Addr:    ":" + constant.PORT,
		Handler: r,
	}
	log.Printf("%s started on port(s): %s (http) with context path '%s'",
		constant.ApplicationName, constant.PORT, constant.ContextPath)
	err := server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("%s started on failure. err = %v", constant.ApplicationName, err))
	}
}

func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}
}

func setupSetting() error {
	v := viper.New()
	v.SetConfigName("application")
	v.AddConfigPath("./configs/")
	v.SetConfigType("yaml")
	if err := v.ReadInConfig(); err != nil {
		return err
	}

	if err := v.UnmarshalKey("chatRobots", &global.ChatRobots); err != nil {
		return err
	}
	fmt.Printf("chatRobots size is %d, follow in:\n", len(global.ChatRobots))
	for _, chatRobot := range global.ChatRobots {
		fmt.Println(chatRobot)
	}
	if err := v.UnmarshalKey("userMobile", &global.UsersMobile); err != nil {
		return err
	}
	fmt.Printf("userMobile size is %d, follow in:\n", len(global.UsersMobile))
	for key, value := range global.UsersMobile {
		fmt.Println(key, ":", value)
	}
	return nil
}
