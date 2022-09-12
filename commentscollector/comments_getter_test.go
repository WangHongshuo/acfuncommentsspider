package commentscollector

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_UnmarshalJsonBody(t *testing.T) {
	f, err := os.Open("./comments.json")
	assert.Equal(t, nil, err)
	defer f.Close()

	b, err := io.ReadAll(f)
	assert.Equal(t, nil, err)

	comments := &CommentsJsonResponse{}
	err = json.Unmarshal(b, comments)
	assert.Equal(t, nil, err)
	assert.Equal(t, true, len(comments.CommentsMap) > 0)
	t.Logf("\n%v", comments)
}

func Test_CommentsGetter(t *testing.T) {
	buff := bytes.NewBuffer(commentsGetter())

	comments := &CommentsJsonResponse{}
	decoder := json.NewDecoder(buff)
	err := decoder.Decode(comments)
	if err != nil {
		fmt.Println(err)
		return
	}
	t.Logf("\n%v", comments)

	assert.Equal(t, true, len(comments.CommentsMap) > 0)
}
