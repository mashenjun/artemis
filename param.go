package artemis

type PageInfo struct {
	Page  int `json:"page"`
	Size  int `json:"size"`
	Total int `json:"total"`
}

type ControlUnitInfo struct {
	CreateTime      uint64 `json:"createTime"`
	UpdateTime      uint64 `json:"updateTime"`
	IndexCode       string `json:"indexCode"`
	Name            string `json:"name"`
	ParentIndexCode string `json:"parentIndexCode"`
	ParentTree      string `json:"parentTree"`
	UnitLevel       int    `json:"unitLevel"`
	UnitType        int    `json:"unitType"`
}

type ControlUnitsRlt struct {
	Code string            `json:"code"`
	Msg  string            `json:"msg"`
	Data []ControlUnitInfo `json:"data"`
	Page PageInfo          `json:"page"`
}

type ChildrenControlUnitsRlt struct {
	Code string            `json:"code"`
	Msg  string            `json:"msg"`
	Data []ControlUnitInfo `json:"data"`
}

type SecurityInfo struct {
	AppSecret  string `json:"appSecret"`
	Time       string `json:"time"`
	TimeSecret string `json:"timeSecret"`
}

type SecurityInfoRlt struct {
	Code string       `json:"code"`
	Msg  string       `json:"msg"`
	Data SecurityInfo `json:"data"`
}

type CameraInfo struct {
	CameraID        string                 `json:"cameraId"`
	IndexCode       string                 `json:"indexCode"`
	Name            string                 `json:"name"`
	ParentIndexCode string                 `json:"parentIndexCode"`
	CameraType      int                    `json:"cameraType"`
	Pixel           int                    `json:"pixel"`
	Latitude        string                 `json:"latitude"`
	Longitude       string                 `json:"longitude"`
	Description     string                 `json:"description"`
	IsOnline        int                    `json:"isOnline"`
	ControlUnitName string                 `json:"controlUnitName"`
	DecodeTag       string                 `json:"decodeTag"`
	CreateTime      uint64                 `json:"createTime"`
	UpdateTime      uint64                 `json:"updateTime"`
	ExtraField      map[string]interface{} `json:"extraField"`
}

type CameraDetail struct {
	ID        int                 `json:"id"`
	MatrixCode string `json:"matrix_code"`
	IndexCode       string                 `json:"indexCode"`
	Name            string                 `json:"name"`
	AppCode string `json:"appCode"`
	OriginalIndexCode string `json:"originalIndexcode"`
	ChanNum int `json:"chanNum"`
	DeviceIndex string `json:"deviceIdx"`
	ParentIndexCode string                 `json:"parentIndexCode"`
	CameraType      string                    `json:"cameraType"`
	Pixel           string                    `json:"pixel"`
	Latitude        float64                 `json:"latitude"`
	Longitude       float64                 `json:"longitude"`
	IsOnline        string                    `json:"isOnline"`
	DecodeTag       string                 `json:"decodeTag"`
	CreateTime      uint64                 `json:"createTime"`
	UpdateTime      uint64                 `json:"updateTime"`
	TreeNodeIndexCode string `json:"treeNodeIndexcode"`
	TreeNodePath string `json:"treeNodePath"`
	TypeCode string `json:"typeCode"`
	ExtraField      map[string]interface{} `json:"extraField"`
}

type CamerasRlt struct {
	Code  string       `json:"code"`
	Msg   string       `json:"msg"`
	Data  []CameraInfo `json:"data"`
	Total string       `json:"total"`
}

type ChildrenCamerasRlt struct {
	Code string       `json:"code"`
	Msg  string       `json:"msg"`
	Data []CameraInfo `json:"data"`
	Page PageInfo     `json:"page"`
}

type CameraDetailRlt struct {
	Code  string       `json:"code"`
	Msg   string       `json:"msg"`
	Data  []CameraDetail `json:"data"`
}
