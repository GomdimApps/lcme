package threads

import (
	"log"
	"runtime"
	"sync"
	"syscall"
	"time"
	"unsafe"

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

// NewEngine initializes a new Engine with a logger and an optimal number of workers.
func NewEngine(maxWorkers int) *Engine {
	logger := log.New(log.Writer(), "Engine: ", log.LstdFlags)

	initialWorkers := runtime.NumCPU()
	if initialWorkers > maxWorkers {
		initialWorkers = maxWorkers
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

	if err := setCPUAffinity(index % runtime.NumCPU()); err != nil {
		e.logger.Println("Failed to set CPU affinity:", err)
	}

	cpuSet := unix.CPUSet{}
	cpuSet.Set(index % runtime.NumCPU())
	unix.SchedSetaffinity(0, &cpuSet)

	// Ajustando prioridade para alta performance
	unix.Setpriority(unix.PRIO_PROCESS, 0, -20)

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

// scaleWorkers dynamically increases the number of worker goroutines based on load.
func (e *Engine) scaleWorkers() {
	e.mu.Lock()
	defer e.mu.Unlock()
	if e.workers < e.maxWorkers {
		additionalWorkers := (e.maxWorkers - e.workers) / 2
		if additionalWorkers < 1 {
			additionalWorkers = 1
		}
		for i := 0; i < additionalWorkers; i++ {
			e.wg.Add(1)
			go e.worker(e.workers + i)
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

func setCPUAffinity(cpu int) error {
	var mask [1]uintptr
	mask[0] = 1 << cpu
	_, _, err := syscall.RawSyscall(syscall.SYS_SCHED_SETAFFINITY, uintptr(syscall.Getpid()), uintptr(len(mask)*8), uintptr(unsafe.Pointer(&mask[0])))
	if err != 0 {
		return err
	}
	return nil
}

// Stop gracefully shuts down the engine by closing the task channel and waiting for all workers to finish.
func (e *Engine) Stop() {
	close(e.tasks)
	e.wg.Wait()
}
