package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/pavan-intelops/Auth/nodecxcv_c/pkg/rest/server/daos/clients/sqls"
	"github.com/pavan-intelops/Auth/nodecxcv_c/pkg/rest/server/models"
	"github.com/pavan-intelops/Auth/nodecxcv_c/pkg/rest/server/services"
	log "github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"net/http"
	"os"
	"strconv"
)

type Auth_serviceController struct {
	authServiceService *services.Auth_serviceService
}

func NewAuth_serviceController() (*Auth_serviceController, error) {
	authServiceService, err := services.NewAuth_serviceService()
	if err != nil {
		return nil, err
	}
	return &Auth_serviceController{
		authServiceService: authServiceService,
	}, nil
}

func (authServiceController *Auth_serviceController) FetchAuth_service(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// trigger authService fetching
	authService, err := authServiceController.authServiceService.GetAuth_service(id)
	if err != nil {
		log.Error(err)
		if errors.Is(err, sqls.ErrNotExists) {
			context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	serviceName := os.Getenv("SERVICE_NAME")
	collectorURL := os.Getenv("OTEL_EXPORTER_OTLP_ENDPOINT")
	if len(serviceName) > 0 && len(collectorURL) > 0 {
		// get the current span by the request context
		currentSpan := trace.SpanFromContext(context.Request.Context())
		currentSpan.SetAttributes(attribute.String("authService.id", strconv.FormatInt(authService.Id, 10)))
	}

	context.JSON(http.StatusOK, authService)
}

func (authServiceController *Auth_serviceController) DeleteAuth_service(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// trigger authService deletion
	if err := authServiceController.authServiceService.DeleteAuth_service(id); err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusNoContent, gin.H{})
}
