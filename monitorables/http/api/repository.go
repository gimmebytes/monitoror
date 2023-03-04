//go:generate mockery -name Repository

package api

import (
	"github.com/monitoror/monitoror/monitorables/http/api/models"
	"net/http"
)

type (
	Repository interface {
		Get(req *http.Request) (*models.Response, error)
	}
)
