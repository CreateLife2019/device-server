package impl

import (
	"fmt"
	"github.com/device-server/domain/constants"
	"github.com/device-server/domain/request/http"
	http2 "github.com/device-server/domain/response/http"
	"github.com/device-server/internal/repository/entity"
	"github.com/device-server/internal/repository/filter"
	"github.com/device-server/internal/repository/persistence"
	"github.com/device-server/internal/repository/persistence/impl"
	"gorm.io/gorm"
	"time"
)

type ProxyServiceImpl struct {
	db    *gorm.DB
	proxy persistence.ProxyIer
}

func NewProxyService(db *gorm.DB) *ProxyServiceImpl {
	s := &ProxyServiceImpl{db: db, proxy: &impl.ProxyImpl{}}
	return s
}

func (p *ProxyServiceImpl) CreateProxy(request http.CreateProxyRequest) (resp http2.CreateProxyResponse, err error) {
	proxies := make([]*entity.Proxy, 0)
	for _, v := range request.Proxies {
		item := &entity.Proxy{
			Host:   v.ProxyHost,
			Port:   v.ProxyPort,
			Secret: v.ProxySecret,
		}
		proxies = append(proxies, item)
	}
	err = p.proxy.BatchSave(p.db, proxies)
	if err != nil {
		resp.Code = constants.Status500
		resp.Msg = err.Error()
		return
	}
	resp.Code = constants.Status200
	resp.Msg = constants.MessageSuc
	return
}
func (p *ProxyServiceImpl) UpdateProxy(request http.UpdateProxyRequest) (resp http2.UpdateProxyResponse, err error) {
	var proxy *entity.Proxy
	proxy, err = p.proxy.Get(p.db, filter.WithId(request.ProxyId))
	if err != nil {
		resp.Code = constants.Status500
		resp.Msg = constants.MessageFailedNoProxy
		return
	}
	proxy.Secret = request.ProxySecret
	proxy.Port = request.ProxyPort
	proxy.Host = request.ProxyHost
	err = p.proxy.Update(p.db, proxy, filter.WithId(proxy.Id))
	if err != nil {
		resp.Code = constants.Status500
		resp.Msg = err.Error()
		return
	}
	resp.Code = constants.Status200
	resp.Msg = constants.MessageSuc
	return
}
func (p *ProxyServiceImpl) ProxyList(request http.ProxyListRequest) (resp http2.ProxyInfoListResponse, err error) {
	proxies := make([]*entity.Proxy, 0)
	page := &entity.Page{
		Page:     request.Page,
		PageSize: request.PageSize,
	}
	proxies, err = p.proxy.SearchProxy(p.db, page)
	if err != nil {
		resp.Code = constants.Status500
		resp.Msg = err.Error()
		return
	}
	resp.Code = "200"
	resp.Msg = constants.MessageSuc
	resp.Data.Total = page.Total
	resp.Data.Page = page.Page
	resp.Data.PageSize = page.PageSize
	resp.Data.Proxies = make([]http2.ProxyInfo, 0)
	for _, v := range proxies {
		resp.Data.Proxies = append(resp.Data.Proxies, http2.ProxyInfo{
			ProxyHost:   v.Host,
			ProxyPort:   v.Port,
			ProxySecret: v.Secret,
			SetTime:     time.Now(),
			ProxyId:     v.Id,
		})
	}
	return
}
func (p *ProxyServiceImpl) DeleteProxy(proxyId int64) (resp http2.DeleteProxyResponse, err error) {
	proxy := &entity.Proxy{
		Base: entity.Base{Id: proxyId},
	}
	err = p.proxy.Delete(p.db, proxy)
	if err != nil {
		err = fmt.Errorf(constants.MessageFailedNotFound)
		resp.Code = constants.Status500
		resp.Msg = constants.MessageFailedNotFound
		return
	}
	resp.Code = constants.Status200
	resp.Msg = constants.MessageSuc
	return
}
func (p *ProxyServiceImpl) ProxyDetail(proxyId int64) (resp http2.ProxyDetailResponse, err error) {
	var proxy *entity.Proxy
	proxy, err = p.proxy.Get(p.db, filter.WithId(proxyId))
	if err != nil {
		resp.Code = constants.Status400
		resp.Msg = constants.MessageFailedNoProxy
		return
	}
	resp.Data.SetTime = time.Now()
	resp.Data.ProxySecret = proxy.Secret
	resp.Data.ProxyPort = proxy.Port
	resp.Data.ProxyHost = proxy.Host
	resp.Code = constants.Status200
	resp.Msg = constants.MessageSuc
	return
}
