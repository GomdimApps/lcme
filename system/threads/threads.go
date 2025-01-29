package threads

import (
	"log"
	"runtime"
	"sync"
	"time"

	"golang.org/x/sys/unix"
)

// Task represents a unit of work to be executed by the engine.
type Task func()

// Engine manages a pool of workers to execute tasks concurrently.
type Engine struct {
	workers    int            // Number of worker goroutines
	tasks      chan Task      // Channel to queue tasks
	wg         sync.WaitGroup // WaitGroup to manage worker goroutines
	mu         sync.RWMutex   // RWMutex to protect shared resources
	mem        map[int][]byte // Memory map for process management
	logger     *log.Logger    // Logger for logging engine activities
	maxWorkers int            // Maximum number of workers
	pool       sync.Pool      // Pool for reusable task objects
}

// NewEngine initializes a new Engine with a logger and a starting number of workers.
func NewEngine(maxWorkers int) *Engine {
	// Logger configuration using standard log
	logger := log.New(log.Writer(), "Engine: ", log.LstdFlags)

	initialWorkers := int(float64(runtime.NumCPU()) * 0.06)
	if initialWorkers < 1 {
		initialWorkers = 1
	}

	return &Engine{
		workers:    initialWorkers,
		tasks:      make(chan Task, 1000),
		mem:        make(map[int][]byte),
		logger:     logger,
		maxWorkers: maxWorkers,
		pool: sync.Pool{
			New: func() interface{} {
				return new(Task)
			},
		},
	}
}

// Start launches the initial set of worker goroutines.
func (e *Engine) Start() {
	for i := 0; i < e.workers; i++ {
		e.wg.Add(1)
		go e.worker(i)
	}
	go e.monitorLoad()
}

// worker executes tasks from the task channel and sets CPU affinity.
func (e *Engine) worker(index int) {
	defer e.wg.Done()
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()
	cpuSet := unix.CPUSet{}
	cpuSet.Set(index % runtime.NumCPU())
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
	if e.workers < e.maxWorkers {
		additionalWorkers := (e.maxWorkers - e.workers) / 2
		for i := 0; i < additionalWorkers; i++ {
			e.wg.Add(1)
			go e.worker(i)
		}
		e.workers += additionalWorkers
		e.logger.Println("Scaled workers to:", e.workers)
	}
}

// monitorLoad adjusts the number of goroutines based on the current load.
func (e *Engine) monitorLoad() {
	ticker := time.NewTicker(400 * time.Millisecond)
	defer ticker.Stop()
	for range ticker.C {
		if len(e.tasks) > cap(e.tasks)/2 {
			e.scaleWorkers()
		}
	}
}

// Stop gracefully shuts down the engine by closing the task channel and waiting for all workers to finish.
func (e *Engine) Stop() {
	close(e.tasks)
	e.wg.Wait()
}
