package main

import (
	"fmt"

	"github.com/jamesrf/goeds/eds"
	"github.com/spf13/viper"
)

func main() {
	viper.AddConfigPath(".")
	viper.ReadInConfig()

	c := eds.NewConnection()

	var err error
	u := viper.GetString("username")
	p := viper.GetString("password")

	if u != "" && p != "" {
		err = c.AuthenticateUser(u, p)
	} else {
		err = c.AuthenticateIP()
	}
	if err != nil {
		panic(err)
	}

	cid := viper.GetString("customerID")

	ses, err := c.CreateSession(cid, true)
	if err != nil {
		panic(err)
	}
	defer c.EndSession(ses)

	srm := eds.NewTestSearch()
	sr, err := ses.Search(srm)

	for _, r := range sr.SearchResult.Data.Records {
		fmt.Println(r.RecordInfo.BibRecord.BibEntity.Titles[0].TitleFull)
	}

}
