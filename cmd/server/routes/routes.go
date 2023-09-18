package routes

import (
	"database/sql"

	"github.com/CamiloMartinez25/odontologia-go/core/middleware"

	//"github.com/aldogayaladh/go-web/cmd/server/handler/ping"
	handlerOdontologo "github.com/CamiloMartinez25/odontologia-go/cmd/server/handler/odontologo"
	handlerPaciente "github.com/CamiloMartinez25/odontologia-go/cmd/server/handler/paciente"
	handlerTurno "github.com/CamiloMartinez25/odontologia-go/cmd/server/handler/turno"

	"github.com/CamiloMartinez25/odontologia-go/internal/domain/odontologo"
	"github.com/CamiloMartinez25/odontologia-go/internal/domain/paciente"
	"github.com/CamiloMartinez25/odontologia-go/internal/domain/turno"
	"github.com/gin-gonic/gin"
)

// Router interface defines the methods that any router must implement.
type Router interface {
	MapRoutes()
}

// router is the Gin router.
type router struct {
	engine      *gin.Engine
	routerGroup *gin.RouterGroup
	db          *sql.DB
}

// NewRouter creates a new Gin router.
func NewRouter(engine *gin.Engine, db *sql.DB) Router {
	return &router{
		engine: engine,
		db:     db,
	}
}

// MapRoutes maps all routes.
func (r *router) MapRoutes() {
	r.setGroup()
	r.buildOdontologoRoutes()
	r.buildPacienteRoutes()
	r.buildTurnoRoutes()
	r.buildPingRoutes()
}

// setGroup sets the router group.
func (r *router) setGroup() {
	r.routerGroup = r.engine.Group("/api/v1")
}

// buildOdontologoRoutes maps all routes for the odontologo domain.
func (r *router) buildOdontologoRoutes() {
	// Create a new odontologo controller.
	repository := odontologo.NewRepositoryMySql(r.db)
	service := odontologo.NewService(repository)
	controlador := handlerOdontologo.NewControladorOdontologo(service)

	r.routerGroup.POST("/odontologos", middleware.Authenticate(), controlador.Create())
	r.routerGroup.GET("/odontologos/:id", middleware.Authenticate(), controlador.GetByID())
	r.routerGroup.PUT("/odontologos/:id", middleware.Authenticate(), controlador.Update())
	r.routerGroup.PATCH("/odontologos/:id", middleware.Authenticate(), controlador.UpdateSubject())
	r.routerGroup.DELETE("/odontologos/:id", middleware.Authenticate(), controlador.Delete())

}

// buildPacienteRoutes maps all routes for the paciente domain.
func (r *router) buildPacienteRoutes() {
	// Create a new paciente controller.
	repository := paciente.NewRepositoryMySql(r.db)
	service := paciente.NewService(repository)
	controlador := handlerPaciente.NewControladorPaciente(service)

	r.routerGroup.POST("/pacientes", middleware.Authenticate(), controlador.Create())
	r.routerGroup.GET("/pacientes/:id", middleware.Authenticate(), controlador.GetByID())
	r.routerGroup.PUT("/pacientes/:id", middleware.Authenticate(), controlador.Update())
	r.routerGroup.PATCH("/pacientes/:id", middleware.Authenticate(), controlador.UpdateSubject())
	r.routerGroup.DELETE("/pacientes/:id", middleware.Authenticate(), controlador.Delete())

}

// buildTurnoRoutes maps all routes for the turno domain.
func (r *router) buildPacienteRoutes() {
	// Create a new turno controller.
	repository := turno.NewRepositoryMySql(r.db)
	service := turno.NewService(repository)
	controlador := handlerTurno.NewControladorTurno(service)

	r.routerGroup.POST("/turnos", middleware.Authenticate(), controlador.Create())
	//r.routerGroup.GET("/turnos/:id", middleware.Authenticate(), controlador.GetByID()) Falta
	r.routerGroup.PUT("/turnos/:id", middleware.Authenticate(), controlador.Update())
	r.routerGroup.PATCH("/turnos/:id", middleware.Authenticate(), controlador.UpdateSubject())
	r.routerGroup.DELETE("/turnos/:id", middleware.Authenticate(), controlador.Delete())
	r.routerGroup.GET("/turnos/:dni", middleware.Authenticate(), controlador.GetByPacienteID())
	// Falta en el service crear CreateByPaciente y CreateByOdontologo o combinados.

}

// buildPingRoutes maps all routes for the ping domain.
func (r *router) buildPingRoutes() {
	// Create a new ping controller.
	pingController := ping.NewControladorPing()
	r.routerGroup.GET("/ping", pingController.Ping())

}
