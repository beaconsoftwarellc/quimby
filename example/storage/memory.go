package storage

import (
	"sync"
	"time"

	"github.com/beaconsoftwarellc/gadget/errors"
	"github.com/beaconsoftwarellc/gadget/generator"
	"github.com/beaconsoftwarellc/quimby/example/models"
)

// WidgetStorage defines the interface for storing widgets
type WidgetStorage interface {
	Get(key string) (*models.Widget, errors.TracerError)
	List() []*models.Widget
	Delete(key string)
	Create(req *models.WidgetRequest) *models.Widget
	Update(widget *models.Widget) errors.TracerError
}

// widgetStorage is a simple in-memory map for widgets
type widgetStorage struct {
	widgets map[string]*models.Widget
	mutex   sync.RWMutex
}

var instance *widgetStorage

// NewWidgetStorage returns an instace of widgetStorage
func NewWidgetStorage() WidgetStorage {
	if instance == nil {
		instance = &widgetStorage{
			widgets: make(map[string]*models.Widget),
		}
	}
	return instance
}

// Get returns a Widget from storage by ID if found
func (s *widgetStorage) Get(key string) (*models.Widget, errors.TracerError) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	w, ok := s.widgets[key]
	if !ok {
		return nil, errors.New("not found")
	}
	return w, nil
}

// List of all the Widgets in storage
func (s *widgetStorage) List() []*models.Widget {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	widgets := []*models.Widget{}
	for _, widget := range s.widgets {
		widgets = append(widgets, widget)
	}
	return widgets
}

// Delete a Widget from storage
func (s *widgetStorage) Delete(key string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	delete(s.widgets, key)
}

// Create a new Widget in storage
func (s *widgetStorage) Create(req *models.WidgetRequest) *models.Widget {
	widget := &models.Widget{
		ID:           generator.ID("wgt"),
		SerialNumber: req.SerialNumber,
		Description:  req.Description,
		CreatedOn:    time.Now(),
		UpdatedOn:    time.Now(),
	}
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.widgets[widget.ID] = widget
	return widget
}

// Update a Widget in storage
func (s *widgetStorage) Update(widget *models.Widget) errors.TracerError {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	_, ok := s.widgets[widget.ID]
	if !ok {
		return errors.New("not found")
	}
	widget.UpdatedOn = time.Now()

	s.widgets[widget.ID] = widget
	return nil
}
