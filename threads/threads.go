package threads

import (
	"context"
	"log"
	"sync"
	"time"
)

type Task func(ctx context.Context) error

type Engine struct {
	workers int
	tasks   chan Task
	wg      sync.WaitGroup
	mu      sync.Mutex
	mem     map[int][]byte
	logger  *log.Logger
}

func NewEngine() *Engine {
	// Logger configuration using standard log
	logger := log.New(log.Writer(), "EngineLogger: ", log.LstdFlags)

	return &Engine{
		workers: 10,
		tasks:   make(chan Task, 1000),
		mem:     make(map[int][]byte),
		logger:  logger,
	}
}

func (e *Engine) Start() {
	for i := 0; i < e.workers; i++ {
		e.wg.Add(1)
		go e.worker()
	}
}

func (e *Engine) worker() {
	defer e.wg.Done()
	for task := range e.tasks {
		func() {
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			err := task(ctx)
			if err != nil {
				e.logger.Println("Error processing task:", err)
			}
		}()
	}
}

func (e *Engine) AddTask(task Task) {
	select {
	case e.tasks <- task:
	default:
		e.scaleWorkers()
		e.tasks <- task
	}
}

func (e *Engine) scaleWorkers() {
	e.mu.Lock()
	defer e.mu.Unlock()
	for i := 0; i < 5; i++ {
		e.wg.Add(1)
		go e.worker()
	}
	e.workers += 5
	e.logger.Println("Scaled workers to:", e.workers)
}

// ForkProcess creates a new process by duplicating the memory of an existing process.
// It locks the engine to ensure thread safety, increments the worker count, and performs
// a copy-on-write operation to duplicate the memory of the original process.
// The function logs the forking action and returns the new process ID.
//
// Parameters:
//
//	id - The ID of the process to be forked.
//
// Returns:
//
//	The ID of the newly created process.
func (e *Engine) ForkProcess(id int) int {
	e.mu.Lock()
	defer e.mu.Unlock()
	newID := e.workers + 1
	// Implement Copy on Write
	if originalMem, ok := e.mem[id]; ok {
		e.mem[newID] = append([]byte{}, originalMem...)
	}
	e.logger.Println("Forked process", id, "to", newID)
	return newID
}

func (e *Engine) WriteProcessMem(id int, data []byte) {
	e.mu.Lock()
	defer e.mu.Unlock()
	if _, ok := e.mem[id]; ok {
		// Allocate new memory for the process if necessary
		e.mem[id] = append([]byte{}, data...)
	} else {
		e.mem[id] = data
	}
}

func (e *Engine) Stop() {
	close(e.tasks)
	e.wg.Wait()
}
