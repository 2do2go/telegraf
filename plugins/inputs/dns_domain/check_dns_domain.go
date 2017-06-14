package dns_domain

import (
	"time"
	"errors"
	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/plugins/inputs"
	"github.com/likexian/whois-go"
	"github.com/likexian/whois-parser-go"
)

type CheckExpire struct {
	// Domains list
	Domains []string
}

// Description returns the plugin Description
func (c *CheckExpire) Description() string {
	return "time left until dns domain paid period is expired"
}

var sampleConfig = `
  ## domain name list default [] )
  domains = ["github.com"]
`

// SampleConfig returns the plugin SampleConfig
func (c *CheckExpire) SampleConfig() string {
	return sampleConfig
}

// Check and return domain expiration date
func (c *CheckExpire) checkDomain(domain string) (time.Time, error) {
	rawResponse, _ := whois.Whois(domain)
    response, err := whois_parser.Parser(rawResponse)

    if err != nil {
    	return time.Time{}, errors.New("Cannot parse raw whois response")
    }

  	layout := "2006-01-02T15:04:05Z"
   	expireDate, err := time.Parse(layout, response.Registrar.ExpirationDate)
   	if err != nil {
   		return time.Time{}, err
   	}

	return expireDate, nil
}

// Gather gets all metric fields and tags and returns any errors it encounters
func (c *CheckExpire) Gather(acc telegraf.Accumulator) error {
	for _, domain := range c.Domains {
		// Prepare data
		tags := map[string]string{"domain": domain}
		// Gather data
		var timeToExpire time.Duration
		timeNow := time.Now()
		expireDate, err := c.checkDomain(domain)
		acc.AddError(err)
		if err != nil {
			timeToExpire = 0
		} else {
			timeToExpire = expireDate.Sub(timeNow)
		}
		fields := map[string]interface{}{"time_to_expire": timeToExpire.Seconds()}
		// Add metrics
		acc.AddFields("dns_domain", fields, tags)
	}
	return nil
}

func init() {
	inputs.Add("check_dns_domain", func() telegraf.Input {
		return &CheckExpire{}
	})
}
