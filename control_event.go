package signalserver

const (
  CONTROL_EVENT_TYPE_KEYCODE uint8 = 0
  CONTROL_EVENT_TYPE_TOUCH   uint8 = 1
  CONTROL_EVENT_TYPE_COMMAND uint8 = 2
  CONTROL_EVENT_TYPE_APK     uint8 = 3
  CONTROL_EVENT_TYPE_FILE    uint8 = 4

  KEY_EVENT_ACTION_DOWN uint8 = 1
  KEY_EVENT_ACTION_UP   uint8 = 0

  TOUTCH_EVENT_ACTION_DOWN uint8 = 1
  TOUCH_EVENT_ACTON_UP     uint8 = 0

  COMMAND_UPDATE_DEVICE       uint8 = 0
  COMMAND_SCREENSHOT          uint8 = 1
  COMMAND_REFERESH_SCREENSHOT uint8 = 2
  COMMAND_CLIBOARD uint8 = 3

  APK_EVENT_OPT_INSTALL   uint8 = 0
  APK_EVENT_OPT_UNINSTALL uint8 = 1
  APK_EVENT_OPT_START     uint8 = 2
  APK_EVENT_OPT_STOP      uint8 = 3

  FILE_EVENT_OPT_GET uint8 = 0
  FILE_EVENT_OPT_DEL uint8 = 1

  HOME_KEY_CODE = 172
)

type WsClientMsg struct {
  RequestId string   `json:"requestid"`
  Cmd       string   `json:"cmd"`
  RoomID    string   `json:"roomid"`
  ClientID  string   `json:"clientid"`
  Target    []string `json:"target"`
  Msg       string   `json:"msg"`
}

type WsServerMsg struct {
  RequestId string `json:"requestid"`
  Cmd       string `json:"cmd"`
  RoomID    string `json:"roomid"`
  ClientID  string `json:"clientid"`
  Msg       string `json:"msg"`
  Error     string `json:"error"`
}

type ControlEvent struct {
  EventType uint8         `json:"event_type"`
  Device    Device        `json:"device"`
  KeyCode   *KeyCodeEvent `json:"key_code,omitempty"`
  Touch     *TouchEvent   `json:"touch,omitempty"`
  Apks      []*ApkEvent   `json:"apk,omitempty"`
  Files     []*FileEvent  `json:"files,omitempty"`
  Command   *CommandEvent `json:"command,omitempty"`
}

type Device struct {
  DeviceWidth  int `json:"device_width"`
  DeviceHeight int `json:"device_height"`
}

type KeyCodeEvent struct {
  Code   int   `json:"code"`
  Action uint8 `json:"action"`
}

type TouchEvent struct {
  X      uint16 `json:"x"`
  Y      uint16 `json:"y"`
  Which  uint8  `json:"which"`
  Action uint8  `json:"action"`
}

type ApkEvent struct {
  JobId        string `json:"job_id"`
  Opt          uint8  `json:"opt"`
  Path         string `json:"path"`
  PackageName  string `json:"package_name"`
  MainActivity string `json:"main_activity"`
  DownloadUrl  string `json:"download_url"`
  Md5          string `json:"md5"`
}

type FileEvent struct {
  JobId       string `json:"job_id"`
  Opt         uint8  `json:"opt"`
  Name        string `json:"name"`
  Path        string `json:"path"`
  Type        string `json:"type"`
  DownloadUrl string `json:"download_url"`
}

type CommandEvent struct {
  Signal uint8 `json:"signal"`
  Shell  string `json:"shell"`
}
