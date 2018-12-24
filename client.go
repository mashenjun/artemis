package artemis

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strconv"
	"time"
)

var Cli *Client

type Client struct {
	client   *http.Client
	endpoint string
	timeout  time.Duration
}

func New(endpoint string, ak string, sk string, opts ...func(*Client)) (*Client, error) {
	_, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}
	c := &Client{
		endpoint: endpoint,
		client: &http.Client{
			Transport: NewAuthTransport(ak, sk),
		},
		timeout: 30 * time.Second,
	}
	for _, opt := range opts {
		opt(c)
	}
	return c, nil
}

func TLSConfigOption (config *tls.Config) func (cli *Client) {
	return func(cli *Client) {
		cli.client.Transport.(*AuthTransport).Tr = &http.Transport{
			TLSClientConfig: config,
		}
	}
}

func (cli *Client) ControlUnits(ctx context.Context, size int, start int) (*ControlUnitsRlt, error) {
	var rlt ControlUnitsRlt
	u := fmt.Sprintf("%v%v?size=%v&start=%v",cli.endpoint, findControlUnitPageURI, size, start)
	req, err := http.NewRequest(http.MethodGet,u, nil)
	req.Header.Add("accept","application/json")
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	resp, err := cli.client.Do(req)
	if err != nil {
		return nil, err
	}
	if err := cli.callResult(ctx, &rlt, resp); err != nil {
		return nil, err
	}
	if rlt.Code != "200" && rlt.Code != "0" {
		return nil, errors.New(rlt.Msg)
	}
	return &rlt, nil
}

func (cli *Client) ChildrenControlUnits(ctx context.Context, parentCode string) (*ChildrenControlUnitsRlt, error) {
	var rlt ChildrenControlUnitsRlt
	u := fmt.Sprintf("%v%v?unitCode=%v",cli.endpoint, findControlUnitByUnitCodeURI, parentCode)
	req, err := http.NewRequest(http.MethodGet,u, nil)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	resp, err := cli.client.Do(req)
	if err != nil {
		return nil, err
	}
	if err := cli.callResult(ctx, &rlt, resp); err != nil {
		return nil, err
	}
	if rlt.Code != "200" && rlt.Code != "0" {
		return nil, errors.New(rlt.Msg)
	}
	return &rlt, nil
}

func (cli *Client) SecurityInfo(ctx context.Context, appKey string) (*SecurityInfoRlt, error) {
	var rlt SecurityInfoRlt
	u := fmt.Sprintf("%v%v",cli.endpoint, fmt.Sprintf(getSecurityInfoURI, appKey))
	req, err := http.NewRequest(http.MethodGet,u, nil)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	resp, err := cli.client.Do(req)
	if err != nil {
		return nil, err
	}
	if err := cli.callResult(ctx, &rlt, resp); err != nil {
		return nil, err
	}
	if rlt.Code != "200" && rlt.Code != "0" {
		return nil, errors.New(rlt.Msg)
	}
	return &rlt, nil
}

func (cli *Client) Cameras(ctx context.Context, size int, start int) (*CamerasRlt, error) {
	var rlt CamerasRlt
	u := fmt.Sprintf("%v%v?size=%v&start=%v",cli.endpoint, findCameraInfoPageURI, size, start)
	req, err := http.NewRequest(http.MethodGet,u, nil)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	resp, err := cli.client.Do(req)
	if err != nil {
		return nil, err
	}
	if err := cli.callResult(ctx, &rlt, resp); err != nil {
		return nil, err
	}
	if rlt.Code != "200" && rlt.Code != "0"{
		return nil, errors.New(rlt.Msg)
	}
	return &rlt, nil
}

func (cli *Client) ChildrenCameras(ctx context.Context, size int, start int, treeNode string) (*ChildrenCamerasRlt, error) {
	var rlt ChildrenCamerasRlt
	u := fmt.Sprintf("%v%v?size=%v&start=%v&treeNode=%v",cli.endpoint, findCameraInfoPageByTreeNodeURI, size, start, treeNode)
	req, err := http.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	resp, err := cli.client.Do(req)
	if err != nil {
		return nil, err
	}
	if err := cli.callResult(ctx, &rlt, resp); err != nil {
		return nil, err
	}
	if rlt.Code != "200" && rlt.Code != "0"{
		return nil, errors.New(rlt.Msg)
	}
	return &rlt, nil
}

func (cli *Client) CameraDetail(ctx context.Context, indexCode string) (*CameraDetailRlt, error) {
	var rlt CameraDetailRlt
	u := fmt.Sprintf("%v%v?indexCode=%v",cli.endpoint, getCameraDetailURI, indexCode)
	req, err := http.NewRequest(http.MethodGet,u, nil)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	resp, err := cli.client.Do(req)
	if err != nil {
		return nil, err
	}
	if err := cli.callResult(ctx, &rlt, resp); err != nil {
		return nil, err
	}
	if rlt.Code != "200" && rlt.Code != "0"{
		return nil, errors.New(rlt.Msg)
	}
	return &rlt, nil
}

func (cli *Client) PreviewURL(ctx context.Context, indexCode string, subStream int, protocol int ) (*PlayURLRlt, error) {
	var rlt PlayURLRlt
	u, _ := url.Parse(cli.endpoint)
	u.Path = path.Join(u.Path, getPreviewURI)
	q := u.Query()
	q.Add("cameraIndexCode", indexCode)
	q.Add("subStream", strconv.Itoa(subStream))
	q.Add("protocol", strconv.Itoa(protocol))
	u.RawQuery = q.Encode()
	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	resp, err := cli.client.Do(req)
	if err != nil {
		return nil, err
	}
	if err := cli.callResult(ctx, &rlt, resp); err != nil {
		return nil, err
	}
	if rlt.Code != "200" && rlt.Code != "0" {
		return nil, errors.New(rlt.Msg)
	}
	return &rlt, nil
}

func (cli *Client)callResult(_ context.Context, rlt interface{}, resp *http.Response) error {
	defer func() {
		resp.Body.Close()
	}()
	if resp.StatusCode/100 != 2 {
		return errors.New(resp.Status)
	}
	if rlt != nil && resp.ContentLength != 0 {
		err := json.NewDecoder(resp.Body).Decode(rlt)
		if err != nil {
			return err
		}
	}
	return nil
}

