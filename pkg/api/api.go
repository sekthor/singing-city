package api

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/sekthor/singing-city/pkg/config"
	"github.com/sekthor/singing-city/pkg/middleware"
	"github.com/sekthor/singing-city/pkg/repo"
	"github.com/sekthor/singing-city/pkg/service"
)

type api struct {
	userService         service.UserService
	venueService        service.VenueService
	artistService       service.ArtistService
	applicationService  service.ApplicationService
	notificationService service.NotificationService
}

func NewApi(conf config.Config) (api, error) {

	log.Info().Msg("api: setting log level to " + conf.Server.Loglevel)
	zerolog.SetGlobalLevel(conf.Server.GetLoglevel())

	api := api{}

	// establish database connection
	log.Info().Msg("api: connection to database")
	db, err := repo.Connect(conf.DB)
	if err != nil {
		log.Error().Err(err).Msg("api: could not connect to database")
		return api, err
	}

	// migrate database schema
	log.Info().Msg("api: migrating database schema")
	err = repo.Migrate(db)
	if err != nil {
		log.Error().Err(err).Msg("api: could not migrate database schema")
		return api, err
	}

	log.Info().Msg("api: setting middleware secret")
	middleware.SetServerSecret(conf.Server.Secret)

	log.Info().Msg("api: initializing service layer")
	api.notificationService = service.NewNotificationService(conf.Smtp)
	api.userService = service.NewUserService(db, &api.notificationService, conf.FrontendBaseUrl)
	api.venueService = service.NewVenueService(db)
	api.artistService = service.NewArtistService(db)
	api.applicationService = service.NewApplicationService(db, &api.notificationService, conf.FrontendBaseUrl)

	log.Info().Msg("api: ensuring presence of admin user")
	err = api.userService.EnsureAdminUser(conf.Server.AdminPass)
	if err != nil {
		log.Error().Err(err).Msg("could not ensure presence of admin user")
		return api, err
	}

	log.Info().Msg("api: successfully initialized api")
	return api, nil
}

func (api *api) Router() *gin.Engine {

	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(middleware.LogRequest)

	// register as user
	router.POST("api/register", api.Register)

	// obtain jwt token as cookie
	router.POST("api/login", api.Login)

	// request a link to reset password
	router.POST("api/forgot-password", api.ForgotPassword)

	// reset password with code from email
	router.POST("api/reset-password", api.ResetPassword)

	// get my profile (user info & either artist or venue)
	router.GET("api/profile", middleware.RequireAuth, api.GetProfile)

	router.GET("/api/admin", middleware.RequireAuth, api.GetAdminInfo)

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

	// get all artists
	router.GET("api/artists", middleware.RequireAuth, api.GetAllArtists)

	// get an artist by their ID
	router.GET("api/artists/:id", middleware.RequireAuth, api.GetArtistById)

	// as venue owner, add a timeslot to a venue
	router.POST("api/timeslots/venues/:userid", middleware.RequireResourceOwnerAuth, api.AddTimeslot)

	// as venue owner, delete a timeslot
	router.DELETE("api/timeslots/:tsid/venues/:userid", middleware.RequireResourceOwnerAuth, api.DeleteTimeslot)

	// as admin delete timeslot
	router.DELETE("api/timeslots/:tsid", middleware.RequireAuth, api.DeleteTimeslotAsAdmin)

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

	router.POST("api/invites", middleware.RequireAuth, api.CreateInvite)

	return router
}
