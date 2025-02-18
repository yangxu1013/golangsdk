package namespaces

import "github.com/chnsz/golangsdk"

func rootURL(client *golangsdk.ServiceClient) string {
	return client.ServiceURL("manage", "namespaces")
}

func resourceURL(client *golangsdk.ServiceClient, name string) string {
	return client.ServiceURL("manage", "namespaces", name)
}
