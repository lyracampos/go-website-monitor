package entities

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWebsite(t *testing.T) {
	t.Run("with success when website is valid", func(t *testing.T) {
		website := WebSite{ Id: 1, Name: "Twitter", Url: "http://www.twitter.com", Status: 1}
		err := website.Validate()

		assert.NoError(t, err)
	})

	t.Run("fails when website name is empty", func(t *testing.T) {
		website := WebSite{ Id: 1, Name: "", Url: "http://www.twitter.com", Status: 1}
		err := website.Validate()

		assert.EqualError(t, err, "Key: 'WebSite.Name' Error:Field validation for 'Name' failed on the 'required' tag")
	})

	t.Run("fails when website url format is incorrect", func(t *testing.T) {
		website := WebSite{ Id: 1, Name: "Twitter", Url: "twitter.com", Status: 1}
		err := website.Validate()

		assert.EqualError(t, err, "Key: 'WebSite.Url' Error:Field validation for 'Url' failed on the 'url' tag")
	})

	t.Run("fails when website url is empty", func(t *testing.T){
		website := WebSite{ Id: 1, Name: "Twitter", Url: "", Status: 1}
		err := website.Validate()

		assert.EqualError(t, err, "Key: 'WebSite.Url' Error:Field validation for 'Url' failed on the 'required' tag")
	})
}
