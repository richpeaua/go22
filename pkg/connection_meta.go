package pkg

import (
	"time"
)


// MetaData defines additional metadata of a stored SSH connection
type MetaData struct {
	HostOS				string
	HostServiceType		string
	GroupLabel			string
	Favorite			bool
	Created				time.Time

}