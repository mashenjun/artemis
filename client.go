package artemis

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

var Cli *Client

type Client struct {
	client   *http.Client
	endpoint string
	httpURL  *url.URL
	timeout  time.Duration
}

func New(endpoint string, ak string, sk string, opts ...func(*Client)) (*Client, error) {
	uri, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}
	c := &Client{
		endpoint: endpoint,
		httpURL:  uri,
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

func (cli *Client) ControlUnits(ctx context.Context, size int, start int) (*ControlUnitsRlt, error) {
	var rlt ControlUnitsRlt
	u := fmt.Sprintf("%v%v?size=%v&start=%v",cli.endpoint, findControlUnitPageURI, size, start)
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

func (cli *Client) CameraDetail(ctx context.Context, indexCode string) (*CameraDetailRlt, error) {
	var rlt CameraDetailRlt
	u := fmt.Sprintf("%v%v",cli.endpoint, fmt.Sprintf(getCameraDetail, indexCode))
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

