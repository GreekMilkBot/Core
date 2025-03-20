package share

import (
	"fmt"
	"io"
	"strings"
)

type Content []ContentNode

func (c Content) String() string {
	r := strings.Builder{}
	for _, content := range c {
		r.WriteString(content.String())
	}
	return r.String()
}

type ContentNode interface {
	fmt.Stringer
}

type Text struct {
	Text string `json:"text"`
}

func (t *Text) String() string {
	return t.Text
}

type At struct {
	Id string
}

func (a *At) String() string {
	return fmt.Sprintf("<at:%s>", a.Id)
}

type File struct {
	Name     string `json:"name"`
	Url      string `json:"url"`
	FileBlob `json:"-"`
}

type FileBlob interface {
	MediaType() string
	Reader() (io.ReadCloser, error)
}

func (f *File) String() string {
	return fmt.Sprintf("[%s](%s)", f.Name, f.Url)
}
