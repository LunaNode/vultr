package lib

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Applications_GetApplications_Error(t *testing.T) {
	server, client := getTestServerAndClient(http.StatusNotAcceptable, `{error}`)
	defer server.Close()

	os, err := client.GetApplications()
	assert.Nil(t, os)
	if assert.NotNil(t, err) {
		assert.Equal(t, `{error}`, err.Error())
	}
}

func Test_Applications_GetApplications_Empty(t *testing.T) {
	server, client := getTestServerAndClient(http.StatusOK, `{}`)
	defer server.Close()

	apps, err := client.GetApplications()
	if err != nil {
		t.Error(err)
	}
	assert.Nil(t, apps)
}

func Test_Applications_GetApplications_OK(t *testing.T) {
	server, client := getTestServerAndClient(http.StatusOK, `{
"1":{"APPID": "1","name": "LEMP","short_name": "lemp","deploy_name": "LEMP on CentOS 6 x64","surcharge": 0},
"2":{"APPID": "2","name": "WordPress","short_name": "wordpress","deploy_name": "WordPress on CentOS 6 x64","surcharge": 0}}`)
	defer server.Close()

	apps, err := client.GetApplications()
	if err != nil {
		t.Error(err)
	}
	if assert.NotNil(t, apps) {
		assert.Equal(t, 2, len(apps))
		// applications could be in random order
		for _, app := range apps {
			switch app.ID {
			case 1:
				assert.Equal(t, "LEMP", app.Name)
				assert.Equal(t, "lemp", app.ShortName)
				assert.Equal(t, "LEMP on CentOS 6 x64", app.DeployName)
				assert.Equal(t, 0, app.Surcharge)
			case 2:
				assert.Equal(t, "WordPress", app.Name)
				assert.Equal(t, "wordpress", app.ShortName)
				assert.Equal(t, "WordPress on CentOS 6 x64", app.DeployName)
				assert.Equal(t, 0, app.Surcharge)
			default:
				t.Error("Unknown APPID")
			}
		}
	}
}
