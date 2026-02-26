package cmd

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/RCCHackathonTeam2/NFTAuctionBase/logger/xzap"
	"github.com/spf13/cobra"
	"go.uber.org/zap"

	"NFTAuctionSync/service"
	"NFTAuctionSync/service/config"
)

/*这段代码实现了一个基于 Cobra 命令行库的守护进程（daemon）命令，主要用于同步 easy swap 订单信息*/
/*使用 cobra.Command 定义了一个守护线程，用于轮训拉取链上日志*/
var DaemonCmd = &cobra.Command{
	Use:   "daemon",
	Short: "sync easy swap order info.",
	Long:  "sync easy swap order info.",
	Run: func(cmd *cobra.Command, args []string) {
		//使用 sync.WaitGroup 来等待 goroutine 完成
		wg := &sync.WaitGroup{}
		wg.Add(1)
		//创建了一个可取消的 context
		ctx := context.Background()
		ctx, cancel := context.WithCancel(ctx)
		// rpc退出信号通知chan
		onSyncExit := make(chan error, 1)
		//启动了一个 goroutine 来执行主要业务逻辑
		go func() {
			defer wg.Done()

			cfg, err := config.UnmarshalCmdConfig() // 读取和解析配置文件
			if err != nil {
				xzap.WithContext(ctx).Error("Failed to unmarshal config", zap.Error(err))
				onSyncExit <- err
				return
			}

			_, err = xzap.SetUp(*cfg.Log) // 初始化日志模块
			if err != nil {
				xzap.WithContext(ctx).Error("Failed to set up logger", zap.Error(err))
				onSyncExit <- err
				return
			}

			xzap.WithContext(ctx).Info("sync server start", zap.Any("config", cfg))

			s, err := service.New(ctx, cfg) // 初始化服务
			if err != nil {
				xzap.WithContext(ctx).Error("Failed to create sync server", zap.Error(err))
				onSyncExit <- err
				return
			}

			if err := s.Start(); err != nil { // 启动服务
				xzap.WithContext(ctx).Error("Failed to start sync server", zap.Error(err))
				onSyncExit <- err
				return
			}

			if cfg.Monitor.PprofEnable { // 开启pprof，用于性能监控
				err := http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", cfg.Monitor.PprofPort), nil)
				if err != nil {
					return
				}
			}
		}()

		// 信号通知chan
		onSignal := make(chan os.Signal)
		// 优雅退出
		signal.Notify(onSignal, syscall.SIGINT, syscall.SIGTERM)
		select {
		case sig := <-onSignal:
			switch sig {
			case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM:
				cancel()
				xzap.WithContext(ctx).Info("Exit by signal", zap.String("signal", sig.String()))
			}
		case err := <-onSyncExit:
			cancel()
			xzap.WithContext(ctx).Error("Exit by error", zap.Error(err))
		}
		wg.Wait()
	},
}

func init() {
	// 将api初始化命令添加到主命令中
	rootCmd.AddCommand(DaemonCmd)
}
