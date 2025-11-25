package main

import (
	"context"
	"fmt"
	"github.com/BurntSushi/toml"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/digitalocean/godo"
)

type Cfg struct {
	token      string
	Domain     string `toml:"domain"`
	SubDomain  string `toml:"subdomain"`
	RecordType string `toml:"record_type"`
}

func main() {
	resp, err := http.Get("https://api.ipify.org")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	publicIp, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error: Couldn't get public ip")
		panic(err)
	}

	cfg := Cfg{token: os.Getenv("API_TOKEN")}
	const cfgFileName string = "config.toml"

	toml.DecodeFile(cfgFileName, &cfg)

	if cfg.token == "" {
		fmt.Println()
		panic("couldn't get api token from config.toml")
	}

	client := godo.NewFromToken(cfg.token)
	ctx := context.TODO()

	opt := &godo.ListOptions{
		Page:    1,
		PerPage: 200,
	}

	domains, _, err := client.Domains.Records(ctx, cfg.Domain, opt)

	if err != nil {
		fmt.Printf("Error getting domains: %s\n", err)
	}

	var record_id int
	for _, d := range domains {
		if d.Name == cfg.SubDomain {
			record_id = d.ID
		}
	}

	if record_id == 0 {
		panic("Error: Coulnd't get dns record id.")
	}

	record, _, err := client.Domains.Record(ctx, cfg.Domain, record_id)

	if err != nil {
		panic("Could not fetch domain record.")
	}

	if record.Data != string(publicIp) {
		editRequest := &godo.DomainRecordEditRequest{
			Type: cfg.RecordType,
			Name: cfg.SubDomain,
			Data: string(publicIp),
		}

		domainRecord, _, err := client.Domains.EditRecord(ctx, cfg.Domain, record_id, editRequest)

		if err != nil {
			fmt.Printf("Error editing record: %s\n", err)
		}

		fmt.Printf("%s: IP is now '%s' for '%s' subdomain of '%s'\n", time.Now(), domainRecord.Data, domainRecord.Name, cfg.Domain)
	} else {
		fmt.Printf("%s: Domain record matches public IP.\n", time.Now())
	}
}
