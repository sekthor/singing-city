package api

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
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

	zerolog.SetGlobalLevel(conf.Server.GetLoglevel())

	api := api{}

	// establish database connection
	db, err := repo.Connect(conf.DB)
	if err != nil {
		return api, err
	}

	// migrate database schema
	err = repo.Migrate(db)
	if err != nil {
		return api, err
	}

	middleware.SetServerSecret(conf.Server.Secret)

	api.userService = service.NewUserService(db)
	api.venueService = service.NewVenueService(db)
	api.artistService = service.NewArtistService(db)
	api.applicationService = service.NewApplicationService(db)

	return api, nil
}

func (api *api) Router() *gin.Engine {

	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(middleware.LogRequest)

	// register as user
	router.POST("api/register", api.Register)

	// optain jwt token as cookie
	router.POST("api/login", api.Login)

	router.GET("api/profile", middleware.RequireAuth, api.GetProfile)

	// update own user details
	router.PUT("api/users/:userid", middleware.RequireResourceOwnerAuth, api.UpdateUser)

	// update venue details
	router.PUT("api/venues/:userid", middleware.RequireResourceOwnerAuth, api.UpdateVenue)

	// update artist details
	router.PUT("api/artists/:userid", middleware.RequireResourceOwnerAuth, api.UpdateArtist)

	// get a list of all venues
	router.GET("api/venues", middleware.RequireAuth, api.GetAllVenues)

	// get a venue by its ID
	router.GET("api/venues/:id", middleware.RequireAuth, api.GetVenueByID)

	// get an artist by their ID
	router.GET("api/artists/:id", middleware.RequireAuth, api.GetArtistById)

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
