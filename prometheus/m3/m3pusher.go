package m3

import (
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

type HTTPDoer interface {
	Do(*http.Request) (*http.Response, error)
}

//M3客户端
type M3Pusher struct {
}

//初始化选项，包括本地缓存空间设置
type M3PusherOpt struct {
}

//初始化M3Pusher客户端
func New(url, namespace string, M3Pusher ...M3PusherOpt) *M3Pusher {
	return nil
}

//配置HTTP Basic Authentication
func (p *M3Pusher) BasicAuth(username, password string) *M3Pusher { return nil }

//为M3Pusher 设置自定义 HTTP 客户端
func (p *M3Pusher) Client(c HTTPDoer) *M3Pusher { return nil }

//为M3Pusher添加Collector
func (p *M3Pusher) Collector(c prometheus.Collector) *M3Pusher { return nil }

//为M3Pusher添加Gatherer
func (p *M3Pusher) Gatherer(g prometheus.Gatherer) *M3Pusher { return nil }

//将数据推送到M3后清空本地缓存
func (p *M3Pusher) PushWithEmpty() error { return nil }

//将数据推送到M3后仍把数据保存在本地
func (p *M3Pusher) Push() error { return nil }

//根据当前Collector和Gatherer定时推送数据，推送后数据仍保存在本地
func (p *M3Pusher) PushTimed(t *time.Duration) error { return nil }

//根据当前Collector和Gatherer定时推送数据，推送后数据清空
func (p *M3Pusher) PushTimedWithEmpty(t *time.Duration) error { return nil }

//清空本地监控缓存
func (p *M3Pusher) Clean() error { return nil }
