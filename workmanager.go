package gtool

import "sync"

// NewWorkManager 建立 goroutine Work 控制器
func NewWorkManager(max int) (manager *WorkManager) {
	return manager.Init(max)
}

// WorkManager 控管 goroutine 執行上限
type WorkManager struct {
	wg       *sync.WaitGroup // goroutine 的等待訊號
	lock     *sync.Mutex     // 讀寫鎖
	count    int             // 目前執行數量
	workChan chan int        // 控管執行數量上限
}

// Init 初始化
func (w *WorkManager) Init(max int) *WorkManager {
	w.lock = new(sync.Mutex)
	w.wg = new(sync.WaitGroup)
	w.workChan = make(chan int, max)
	return w
}

// Start 增加數量
func (w *WorkManager) Start() {
	w.lock.Lock()
	w.count++
	w.lock.Unlock()

	// 開始工作
	w.workChan <- 1
	w.wg.Add(1)
}

// End 減少數量
func (w *WorkManager) End() {
	w.lock.Lock()
	w.count--
	w.lock.Unlock()

	// 結束工作
	<-w.workChan
	w.wg.Done()
}

// Count 取得運行中的 goroutine 數量
func (w *WorkManager) Count() int {
	return w.count
}

// Wait 等待 goroutine 全部執行完畢
func (w *WorkManager) Wait() {
	w.wg.Wait()
	// 執行完畢後關閉 chan
	close(w.workChan)
}
