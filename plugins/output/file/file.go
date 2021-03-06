package file

import (
	"github.com/luopengift/golibs/file"
	"github.com/luopengift/transport"
	"os"
)

const (
	VERSION = "0.0.1"
)

type FileOutput struct {
	Path  string `json:"path"` //配置路径
	cpath string //真实路径
	fd    *os.File
}

func NewFileOutput() *FileOutput {
	return new(FileOutput)
}

func (out *FileOutput) Init(config transport.Configer) error {
	err := config.Parse(out)
	if err != nil {
		return err
	}
	out.cpath = file.TimeRule.Handle(out.Path)
	out.fd, err = os.OpenFile(out.cpath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	return err
}

func (out *FileOutput) Write(p []byte) (int, error) {
	if cpath := file.TimeRule.Handle(out.Path); out.cpath != cpath {
		var err error
		out.cpath = cpath
		out.fd, err = os.OpenFile(out.cpath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			return 0, err
		}
	}
	return out.fd.Write(p)
}

func (out *FileOutput) Start() error {
	return nil
}

func (out *FileOutput) Close() error {
	return out.fd.Close()
}

func (out *FileOutput) Version() string {
	return VERSION
}
func init() {
	transport.RegistOutputer("file", NewFileOutput())
}
