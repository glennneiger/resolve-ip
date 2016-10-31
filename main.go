package main

import (
	"context"
	"flag"
	"log"
	"os"

	lightstep "github.com/lightstep/lightstep-tracer-go"
	opentracing "github.com/opentracing/opentracing-go"

	"github.com/Clever/resolve-ip/gen-go/models"
	"github.com/Clever/resolve-ip/gen-go/server"
	"gopkg.in/Clever/kayvee-go.v5/logger"
)

var addr = flag.String("addr", ":8080", "Address to listen at")
var path = flag.String("path", "./GeoLiteCity", "path to Geo Lite City db")

type handler struct {
	db *GeoDB
}

func (h handler) HealthCheck(ctx context.Context) error {
	return nil
}

func (h handler) LocationForIP(ctx context.Context, i *models.LocationForIPInput) (*models.IP, error) {
	latlon, err := h.db.Lookup(i.IP)
	if err == ErrIPMissing {
		return nil, models.LocationForIP404Output{}
	}
	if err == ErrBadIP {
		return nil, models.DefaultBadRequest{Msg: err.Error()}
	}
	if err != nil {
		return nil, models.DefaultInternalError{Msg: err.Error()}
	}
	return &models.IP{
		Lat: &latlon.Lat,
		Lon: &latlon.Lon,
	}, nil

}

func main() {
	flag.Parse()
	logger := logger.New("resolve-ip")

	tags := make(map[string]interface{})
	tags[lightstep.ComponentNameKey] = "resolve-ip"
	lightstepTracer := lightstep.NewTracer(lightstep.Options{
		AccessToken: os.Getenv("LIGHTSTEP_ACCESS_TOKEN"),
		Tags:        tags,
	})
	defer lightstep.FlushLightStepTracer(lightstepTracer)
	opentracing.InitGlobalTracer(lightstepTracer)

	logger.Debug("build-db-start")
	db, err := NewGeoDB(*path)
	if err != nil {
		log.Fatalf("failed to create geo db: %s", err)
	}
	logger.Debug("build-db-end")

	s := server.New(handler{db}, *addr)
	// Serve should not return
	log.Fatal(s.Serve())
}
