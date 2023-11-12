package api

import (
	"github.com/gin-gonic/gin"
	"github.com/sekthor/songbird-backend/pkg/config"
	"github.com/sekthor/songbird-backend/pkg/middleware"
	"github.com/sekthor/songbird-backend/pkg/repo"
	"github.com/sekthor/songbird-backend/pkg/service"
)

type api struct {
	userService         service.UserService
	venueService        service.VenueService
	artistService       service.ArtistService
	applicationService  service.ApplicationService
	notificationService service.NotificationService
}

func NewApi(conf config.Config) (api, error) {
	api := api{}
	db, err := repo.Connect(conf.DB)
	if err != nil {
		return api, err
	}

	repo.Migrate(db)

	middleware.SetServerSecret(conf.Server.Secret)

	api.userService = service.NewUserService(db)
	api.venueService = service.NewVenueService(db)
	api.artistService = service.NewArtistService(db)
	api.applicationService = service.NewApplicationService(db)
	api.notificationService = service.NewNotificationService(conf.Smtp)

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

	router.GET("api/artists/:id", api.GetArtistById)

	// as venue owner, add a timeslot to a venue
	router.POST("api/timeslots/venues/:userid", middleware.RequireResourceOwnerAuth, api.AddTimeslot)

	// as venue owner, delete a timeslot
	router.DELETE("api/timeslots/:tsid/venues/:userid", middleware.RequireResourceOwnerAuth, api.DeleteTimeslot)

	// as artist, apply for timeslot
	router.POST("api/timeslots/:tsid/apply/:userid", middleware.RequireResourceOwnerAuth, api.Apply)

	// as artist or venue, get my timeslots
	router.GET("api/timeslots", middleware.RequireAuth, api.GetTimeslots)

	// as venue owner or artist, get all my applications
	router.GET("api/applications/:usertype/:userid", middleware.RequireResourceOwnerAuth, api.GetApplicationsOfUser)

	// as venue owner, I can accept an application
	router.POST("api/applications/:id/accept", middleware.RequireAuth, api.AcceptApplication)

	// as venue owner or artist, i can decline/retract an application
	router.DELETE("api/applications/:id", middleware.RequireAuth, api.DeleteApplication)

	return router
}

func (api *api) Restricted(c *gin.Context) {
	val, _ := c.Get("userid")
	c.JSON(200, gin.H{"msg": "you are authenticated. userid: " + val.(string)})
}
