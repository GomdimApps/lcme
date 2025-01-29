package threads

import (
	"log"
	"runtime"
	"sync"

	"golang.org/x/sys/unix"
)

// Task represents a unit of work to be executed by the engine.
type Task func()

// Engine manages a pool of workers to execute tasks concurrently.
type Engine struct {
	workers int            // Number of worker goroutines
	tasks   chan Task      // Channel to queue tasks
	wg      sync.WaitGroup // WaitGroup to manage worker goroutines
	mu      sync.RWMutex   // RWMutex to protect shared resources
	mem     map[int][]byte // Memory map for process management
	logger  *log.Logger    // Logger for logging engine activities
}

// NewEngine initializes a new Engine with a logger and a starting number of workers.
func NewEngine() *Engine {
	// Logger configuration using standard log
	logger := log.New(log.Writer(), "Engine: ", log.LstdFlags)

	// Start with 2% of CPU usage
	initialWorkers := int(float64(runtime.NumCPU()) * 0.06)
	if initialWorkers < 1 {
		initialWorkers = 1
	}

	return &Engine{
		workers: initialWorkers,
		tasks:   make(chan Task, 1000),
		mem:     make(map[int][]byte),
		logger:  logger,
	}
}

// Start launches the initial set of worker goroutines.
func (e *Engine) Start() {
	for i := 0; i < e.workers; i++ {
		e.wg.Add(1)
		go e.worker(i)
	}
}

// worker executes tasks from the task channel and sets CPU affinity.
func (e *Engine) worker(index int) {
	defer e.wg.Done()
	// Set CPU affinity for this worker
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()
	cpuSet := unix.CPUSet{}
	cpuSet.Set((runtime.NumCPU() + index) % runtime.NumCPU()) // Distribute across CPUs
	unix.SchedSetaffinity(0, &cpuSet)

	for task := range e.tasks {
		task()
	}
}

// AddTask adds a new task to the task channel, scaling workers if necessary.
func (e *Engine) AddTask(task Task) {
	select {
	case e.tasks <- task:
	default:
		e.scaleWorkers()
		e.tasks <- task
	}
}

// scaleWorkers increases the number of worker goroutines based on CPU count.
func (e *Engine) scaleWorkers() {
	e.mu.Lock()
	defer e.mu.Unlock()
	additionalWorkers := runtime.NumCPU() / 2
	for i := 0; i < additionalWorkers; i++ {
		e.wg.Add(1)
		go e.worker(i)
	}
	e.workers += additionalWorkers
	e.logger.Println("Scaled workers to:", e.workers)
}

// ForkProcess creates a new process by duplicating the memory of an existing process.
func (e *Engine) ForkProcess(id int) int {
	e.mu.Lock()
	defer e.mu.Unlock()
	newID := e.workers + 1
	// Implement Copy on Write
	if originalMem, ok := e.mem[id]; ok {
		e.mem[newID] = append(e.mem[newID][:0], originalMem...)
	}
	e.logger.Println("Forked process", id, "to", newID)
	return newID
}

// WriteProcessMem writes data to the memory of a specified process.
func (e *Engine) WriteProcessMem(id int, data []byte) {
	e.mu.Lock()
	defer e.mu.Unlock()
	if _, ok := e.mem[id]; ok {
		// Allocate new memory for the process if necessary
		e.mem[id] = append(e.mem[id][:0], data...)
	} else {
		e.mem[id] = data
	}
}

// Stop gracefully shuts down the engine by closing the task channel and waiting for all workers to finish.
func (e *Engine) Stop() {
	close(e.tasks)
	e.wg.Wait()
}
