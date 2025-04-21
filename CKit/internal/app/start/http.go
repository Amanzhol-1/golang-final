package start

import (
	"CKit/internal/app/config"
	"CKit/internal/app/connections"
	"CKit/internal/delieveries"
	repository "CKit/internal/repository/shipment"
	"CKit/internal/services/shipment"
	"fmt"
	"log"
	"net/http"
)

func HTTP(conn *connections.Connections, cfg *config.HTTPServerConfig) {
	// Health check endpoint
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("OK"))
	})

	repo := repository.NewDBShipmentRepository(conn.DB)
	createUC := shipment.NewCreateShipmentUseCase(repo)
	getUC := shipment.NewGetShipmentUseCase(repo)
	listUC := shipment.NewListShipmentsUseCase(repo)
	updateUC := shipment.NewUpdateShipmentUseCase(repo)
	deleteUC := shipment.NewDeleteShipmentUseCase(repo)

	handler := delieveries.NewShipmentHandler(createUC, getUC, listUC, updateUC, deleteUC)

	router := delieveries.NewRouter(conn.HTTPClient, handler)

	addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	log.Printf("Starting HTTP server on %s", addr)

	router.Start(addr)
}
