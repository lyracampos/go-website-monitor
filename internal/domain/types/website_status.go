package types

import "fmt"

type WebsiteStatus int

const (
	Enabled WebsiteStatus = 1
	Disabled WebsiteStatus = 2
)

func (s WebsiteStatus) String() (string, error) {
	switch s {
	case Enabled:
		return "enabled", nil
	case Disabled:
		return "disabled", nil
	}
	return "", fmt.Errorf("WebsiteStatus: the value could not be mapped to string: #{s}" )
}