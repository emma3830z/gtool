package gtool

import (
	"os"
	"os/signal"
	"syscall"
)

// NewProcessManager 建立 進程控制器
func NewProcessManager() (manager *ProcessManager) {
	return manager.Init()
}

// ProcessManager 控管進程執行/停止
type ProcessManager struct {
	ch   chan os.Signal
	stop bool // 是否有收到進程停止訊息
}

// Init 初始化，開始監控停止訊息
func (p *ProcessManager) Init() *ProcessManager {
	p.ch = make(chan os.Signal, 10)
	go func() {
		signal.Notify(p.ch, syscall.SIGINT, syscall.SIGTERM, syscall.SIGUSR2)

		<-p.ch
		p.stop = true
		signal.Stop(p.ch)
	}()
	return p
}

// StopChan 結束接收訊息
func (p *ProcessManager) StopChan() {
	// 傳入結束訊息
	p.ch <- syscall.SIGUSR2
}

// IsStop 是否有收到進程停止訊息
func (p *ProcessManager) IsStop() bool {
	return p.stop
}
