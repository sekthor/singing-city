package api

import (
	"github.com/gin-gonic/gin"
	"github.com/sekthor/songbird-backend/pkg/config"
	"github.com/sekthor/songbird-backend/pkg/middleware"
	"github.com/sekthor/songbird-backend/pkg/repo"
	"github.com/sekthor/songbird-backend/pkg/service"
)

type api struct {
	userService        service.UserService
	venueService       service.VenueService
	artistService      service.ArtistService
	applicationService service.ApplicationService
}

func NewApi(conf config.Config) (api, error) {
	api := api{}
	db, err := repo.Connect(conf.DB)
	if err != nil {
		return api, err
	}

	repo.Migrate(db)

	api.userService = service.NewUserService(db)
	api.venueService = service.NewVenueService(db)
	api.artistService = service.NewArtistService(db)
	api.applicationService = service.NewApplicationService(db)

	return api, nil
}

func (api *api) Router() *gin.Engine {
	router := gin.Default()

	router.POST("api/register", api.Register)
	router.POST("api/login", api.Login)
	router.GET("api/auth", middleware.RequireAuth, api.Restricted)
	router.GET("api/auth/user/:userid", middleware.RequireResourceOwnerAuth, api.Restricted)

	router.GET("api/venues", api.GetAllVenues)
	router.GET("api/venues/:id", api.GetVenueByID)
	router.POST("api/venues")
	//router.DELETE("api/venues/:id")

	// as venue owner, add a timeslot to a venue
	router.POST("api/venues/:userid/timeslot", middleware.RequireResourceOwnerAuth, api.AddTimeslot)
	router.DELETE("api/venues/:userid/timeslot/:tsid", middleware.RequireResourceOwnerAuth, api.DeleteTimeslot)

	// as artist, apply for timeslot
	router.POST("api/timeslots/:tsid/apply/:userid", middleware.RequireResourceOwnerAuth, api.Apply)

	return router
}

func (api *api) Restricted(c *gin.Context) {
	c.JSON(200, gin.H{"msg": "you are authenticated"})
}
