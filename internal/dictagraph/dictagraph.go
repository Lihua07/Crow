package dictagraph

import (
	"github.com/kardianos/service"
	"github.com/spf13/cobra"
)

func NewDictagraphCommand() *cobra.Command {
    cmd := &cobra.Command{
        Use: "dictagraph",
        Short: "dictagraph 一个简单的命令行程序",
        Long: "dictagraph 一个简单的命令行程序",
        SilenceUsage: true,
        RunE: func(cmd *cobra.Command, args []string) error {
            return runService()
        },
    }
    return cmd
}

type Dictagraph struct {
    Name string
    Date string
}

func NewDictagraph(name, date string) *Dictagraph {
    return &Dictagraph{
        Name: name,
        Date: date,
    }
}

func (d *Dictagraph) Start(s service.Service) error {
    go d.run()
    return nil
}

func (d *Dictagraph) Stop(s service.Service) error {
    return nil
}

func (d *Dictagraph) Restart(s service.Service) error {
    return nil
}

func (d *Dictagraph) run() error {
    return nil
}

func runService() error {
    serviceConfig := &service.Config{
        Name: "dictagraph deamo programming",
        DisplayName: "dictagraph",
        Description: "一个监控服务程序",
    }

    graph := NewDictagraph("Dictagraph", "2024-08-14")

    srv, err := service.New(graph, serviceConfig)
    if err != nil {
        return err
    }

    if err := srv.Run(); err != nil {
        return err
    }

    return nil
}
