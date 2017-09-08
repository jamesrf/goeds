package main

import "github.com/jamesrf/goeds/eds"
import "github.com/spf13/viper"

func main() {
	viper.AddConfigPath(".")
	viper.ReadInConfig()

	c := eds.NewConnection()

	err := c.AuthenticateIP()
	if err != nil {
		panic(err)
	}
	ses, err := c.CreateSession(viper.GetString("customerID"), true)
	c.EndSession(ses)
}
