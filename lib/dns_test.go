package lib

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_DNS_GetDNSDomains_Error(t *testing.T) {
	server, client := getTestServerAndClient(http.StatusNotAcceptable, `{error}`)
	defer server.Close()

	keys, err := client.GetDNSDomains()
	assert.Nil(t, keys)
	if assert.NotNil(t, err) {
		assert.Equal(t, `{error}`, err.Error())
	}
}

func Test_DNS_GetDNSDomains_NoDomains(t *testing.T) {
	server, client := getTestServerAndClient(http.StatusOK, `[]`)
	defer server.Close()

	keys, err := client.GetDNSDomains()
	if err != nil {
		t.Error(err)
	}
	assert.Nil(t, keys)
}

func Test_DNS_GetDNSDomains_OK(t *testing.T) {
	server, client := getTestServerAndClient(http.StatusOK, `[
    {"domain": "example.com","date_created": "2012-11-23 13:37:33"},
    {"domain": "example2.com","date_created": "2010-11-23 13:37:44"}
    ]`)
	defer server.Close()

	domains, err := client.GetDNSDomains()
	if err != nil {
		t.Error(err)
	}
	if assert.NotNil(t, domains) {
		assert.Equal(t, 2, len(domains))
		// domains could be in random order
		for _, domain := range domains {
			switch domain.Domain {
			case "example.com":
				assert.Equal(t, "2012-11-23 13:37:33", domain.Created)
			case "example2.com":
				assert.Equal(t, "2010-11-23 13:37:44", domain.Created)
			default:
				t.Error("Unknown DNS domain")
			}
		}
	}
}

func Test_DNS_CreateDNSDomain_Error(t *testing.T) {
	server, client := getTestServerAndClient(http.StatusNotAcceptable, `{error}`)
	defer server.Close()

	err := client.CreateDNSDomain("example.com", "1.2.3.4")
	if assert.NotNil(t, err) {
		assert.Equal(t, `{error}`, err.Error())
	}
}

func Test_DNS_CreateDNSDomain_OK(t *testing.T) {
	server, client := getTestServerAndClient(http.StatusOK, `{no-response?!}`)
	defer server.Close()

	err := client.CreateDNSDomain("example.com", "1.2.3.4")
	if err != nil {
		t.Error(err)
	}
}

func Test_DNS_DeleteDNSDomain_Error(t *testing.T) {
	server, client := getTestServerAndClient(http.StatusNotAcceptable, `{error}`)
	defer server.Close()

	err := client.DeleteDNSDomain("id-1")
	if assert.NotNil(t, err) {
		assert.Equal(t, `{error}`, err.Error())
	}
}

func Test_DNS_DeleteDNSDomain_OK(t *testing.T) {
	server, client := getTestServerAndClient(http.StatusOK, `{no-response?!}`)
	defer server.Close()

	assert.Nil(t, client.DeleteDNSDomain("1"))
}
