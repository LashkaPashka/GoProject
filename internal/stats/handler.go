package stats

import (
	"go/project_go/configs"
	"go/project_go/pkg/middleware"
	"go/project_go/pkg/res"
	"net/http"
	"time"
)

const (
	GroupByDay = "day"
	GroupByMonth = "month"
)


type StatHandlerDeps struct{
	Config *configs.Config
	StatRepository *StatsRepository
}

type StatHandler struct{
	StatRepository *StatsRepository
}

func NewStatHandler(router *http.ServeMux, deps StatHandlerDeps){
	handler := &StatHandler{
		StatRepository: deps.StatRepository,
	}

	router.Handle("GET /stat", middleware.IsAuthed(handler.GetStat(), deps.Config))
}

func (h *StatHandler) GetStat() http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		layout := "2006-01-02"
		
		from, err := time.Parse(layout, r.URL.Query().Get("from"))
		if err != nil {
			http.Error(w, "Invalid param", http.StatusBadGateway)
			return
		}

		to, err := time.Parse(layout, r.URL.Query().Get("to"))
		if err != nil {
			http.Error(w, "Invalid param", http.StatusBadRequest)
			return
		}

		by := r.URL.Query().Get("by")
		if by != GroupByDay && by != GroupByMonth{
			http.Error(w, "Invalid param", http.StatusBadRequest)
			return
		}
		
		stats := h.StatRepository.GetStat(from, to, by)
		res.EncodeJson(w, stats, http.StatusOK)
	}
}