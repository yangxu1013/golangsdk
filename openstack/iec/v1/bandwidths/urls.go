package bandwidths

import (
	"github.com/chnsz/golangsdk"
)

func GetURL(c *golangsdk.ServiceClient, bandwidthId string) string {
	return c.ServiceURL("bandwidths", bandwidthId)
}

func UpdateURL(c *golangsdk.ServiceClient, bandwidthId string) string {
	return c.ServiceURL("bandwidths", bandwidthId)
}
