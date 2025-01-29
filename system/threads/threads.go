package threads

import (
	"log"
	"runtime"
	"sync"
	"time"
)

type Task func()

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
		workers: runtime.NumCPU(),
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
	go e.Monitor()       // Start the monitor to dynamically adjust workers
	go e.AdjustWorkers() // Start the adjuster to dynamically adjust workers based on task execution time
}

func (e *Engine) worker() {
	defer e.wg.Done()
	for task := range e.tasks {
		func() {
			start := time.Now()
			task()
			duration := time.Since(start)
			e.logger.Println("Task completed in:", duration)
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
	additionalWorkers := runtime.NumCPU() / 2
	for i := 0; i < additionalWorkers; i++ {
		e.wg.Add(1)
		go e.worker()
	}
	e.workers += additionalWorkers
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

// Monitor dynamically adjusts the number of workers based on the load.
func (e *Engine) Monitor() {
	ticker := time.NewTicker(400 * time.Millisecond) // Alterado para 400 milissegundos
	defer ticker.Stop()
	for range ticker.C {
		e.mu.Lock()
		queueLength := len(e.tasks)
		e.mu.Unlock()
		if queueLength > cap(e.tasks)/2 {
			e.scaleWorkers()
		}
	}
}

// AdjustWorkers dynamically adjusts the number of workers based on task execution time.
func (e *Engine) AdjustWorkers() {
	ticker := time.NewTicker(400 * time.Millisecond) // Alterado para 400 milissegundos
	defer ticker.Stop()
	for range ticker.C {
		e.mu.Lock()
		totalTasks := len(e.tasks)
		e.mu.Unlock()
		if totalTasks > 0 {
			avgTaskTime := e.calculateAverageTaskTime()
			if avgTaskTime > 100*time.Millisecond {
				e.scaleWorkers()
			}
		}
	}
}

// calculateAverageTaskTime calculates the average execution time of tasks.
func (e *Engine) calculateAverageTaskTime() time.Duration {
	var totalDuration time.Duration
	var taskCount int
	for task := range e.tasks {
		start := time.Now()
		task()
		duration := time.Since(start)
		totalDuration += duration
		taskCount++
	}
	if taskCount == 0 {
		return 0
	}
	return totalDuration / time.Duration(taskCount)
}
