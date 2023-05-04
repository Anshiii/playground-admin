package models

import (
	"github.com/anshiii/playground-admin/microsite"
)

type MicrositeModel struct {
	Name        string
	Description string
	microsite.MicroSite
}
