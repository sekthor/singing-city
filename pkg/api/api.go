package api

import (
	"github.com/gin-gonic/gin"
	"github.com/sekthor/songbird-backend/pkg/config"
	"github.com/sekthor/songbird-backend/pkg/repo"
	"github.com/sekthor/songbird-backend/pkg/service"
)

type api struct {
    venueService service.VenueService
}

func NewApi(conf config.Config) (api, error) {
    api := api{}
    db, err := repo.Connect(conf.DB)
    if err != nil {
        return api, err
    }

    api.venueService = service.NewVenueService(db)

    return api, nil
}
    
func (*api) Router() *gin.Engine {
    router := gin.Default()
    
    router.GET("venues")
    router.GET("venues/:id")
    router.POST("venues")
    router.DELETE("veneues/:id")

    return router
}
