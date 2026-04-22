package api

import (

	"context"

	"encoding/json"

	"net/http"

	"regexp"

	"time"

	"github.com/google/uuid"

	"log/slog"

	"onboarding/internal/model"

	"onboarding/internal/metrics"

	"onboarding/internal/postgres"

	"onboarding/internal/redis"

	"onboarding/internal/kafka"
)

type Handler struct {

	db *postgres.Client

	cache *redis.Client

	producer *kafka.Producer

	log *slog.Logger
}

var vpaRegex = regexp.MustCompile(`^[a-zA-Z0-9._-]+@[a-zA-Z]{3,}$`)

func NewHandler(

	db *postgres.Client,

	cache *redis.Client,

	producer *kafka.Producer,

	log *slog.Logger,

) *Handler {

	return &Handler{

		db: db,

		cache: cache,

		producer: producer,

		log: log,
	}
}

func (h *Handler) Onboard(w http.ResponseWriter, r *http.Request) {

	metrics.RequestCount.Inc()

	var req model.OnboardingRequest

	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {

		http.Error(w,"invalid json",400)

		return
	}

	if req.UserID == "" || !vpaRegex.MatchString(req.VPA) {

		http.Error(w,"invalid input",400)

		return
	}

	event := model.OnboardingEvent{

		EventID: uuid.New().String(),

		UserID: req.UserID,

		VPA: req.VPA,

		KYCLevel: req.KYCLevel,

		Balance: req.Balance,

		Status: "initiated",

		Timestamp: time.Now().UTC().Format(time.RFC3339),
	}

	ctx := context.Background()

	_, err = h.db.DB.ExecContext(

		ctx,

		`insert into onboarding_events 
		(event_id,user_id,vpa,kyc_level,balance,status,created_at)
		values ($1,$2,$3,$4,$5,$6,$7)`,

		event.EventID,

		event.UserID,

		event.VPA,

		event.KYCLevel,

		event.Balance,

		event.Status,

		event.Timestamp,
	)

	if err != nil {

		h.log.Error("postgres insert failed", "err", err)

		http.Error(w,"db error",500)

		return
	}

	b, _ := json.Marshal(event)

	err = h.producer.Publish(event.UserID,b)

	if err != nil {

		metrics.KafkaErrors.Inc()

		h.log.Error("kafka publish failed", "err", err)
	}

	w.WriteHeader(http.StatusAccepted)

	json.NewEncoder(w).Encode(event)

	h.log.Info("onboarding initiated",

		"user_id", event.UserID,

		"event_id", event.EventID)
}

func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(200)

	w.Write([]byte("ok"))
}

func (h *Handler) Ready(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(200)
}
