package stats

import (
	"go/project_go/pkg/event"
	"log"
)

type ServiceStatDeps struct{
	EventBus *event.EventBus
	StatRepository *StatsRepository
}

type ServiceStat struct{
	EventBus *event.EventBus
	StatRepository *StatsRepository
}

func NewServiceStat(deps *ServiceStatDeps) *ServiceStat{
	return &ServiceStat{
		EventBus: deps.EventBus,
		StatRepository: deps.StatRepository,
	}
}

func (s *ServiceStat) AddClick(){
	for msg := range s.EventBus.Subscribe(){
		if msg.Type == event.EventLinkVisited {
			id, ok := msg.Data.(uint)
			if !ok{
				log.Fatalln("Wrong data")
				continue
			}
			s.StatRepository.AddClick(id)
		}
	}
}