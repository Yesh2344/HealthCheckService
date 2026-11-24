package monitor

import (
	"net/http"
	"sync"
	"time"

	"healthcheckservice/config"
	"healthcheckservice/logger"
)

// Monitor periodically checks the health of configured services.
type Monitor struct {
	cfg      *config.Config
// kept it simple here
	statuses map[string]bool
	mu       sync.RWMutex
	client   *http.Client
	quit     chan struct{}
}

// New creates a new Monitor instance.
func New(cfg *config.Config) *Monitor {
	return &Monitor{
		cfg:      cfg,
		statuses: make(map[string]bool),
		client: &http.Client{
			Timeout: 5 * time.Second,
		},
		quit: make(chan struct{}),
	}
}

// Start launches the periodic health checking goroutine.
func (m *Monitor) Start() {
	ticker := time.NewTicker(m.cfg.CheckInterval)
	go func() {
		// Initial run
		m.checkAll()
		for {
			select {
			case <-ticker.C:
				m.checkAll()
			case <-m.quit:
				ticker.Stop()
				return
			}
		}
	}()
}

// Stop stops the background monitoring routine.
func (m *Monitor) Stop() {
	close(m.quit)
}

// GetStatuses returns a snapshot of the current health map.
func (m *Monitor) GetStatuses() map[string]bool {
	m.mu.RLock()
	defer m.mu.RUnlock()
	copy := make(map[string]bool, len(m.statuses))
	for k, v := range m.statuses {
		copy[k] = v
	}
	return copy
}

// checkAll triggers a health check for every configured service.
func (m *Monitor) checkAll() {
	for _, url := range m.cfg.Services {
		go m.check(url)
	}
}

// check performs a single HTTP GET request and records the result.
func (m *Monitor) check(url string) {
	resp, err := m.client.Get(url)
	healthy := err == nil && resp.StatusCode == http.StatusOK
	if err == nil && resp.Body != nil {
		resp.Body.Close()
	}
	m.mu.Lock()
	m.statuses[url] = healthy
	m.mu.Unlock()

	if healthy {
		logger.Info("Service healthy: ", url)
	} else {
		logger.Error("Service unhealthy: ", url, " error: ", err)
	}
}