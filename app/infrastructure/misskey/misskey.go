package misskey

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/daifukuninja/petit-misskey-go/model/misskey"
	"github.com/pkg/errors"
)

type Client struct {
	client http.Client
	url    string
}

func NewClient(url string, timeout time.Duration) *Client {
	return &Client{
		client: http.Client{Timeout: timeout},
		url:    url,
	}
}

func (c *Client) Meta(ctx context.Context, contents misskey.Meta) (*misskey.MetaResponse, error) {

	response, err := c.post(ctx, c.meta(), contents)
	if err != nil {
		return nil, err
	}

	ret := new(misskey.MetaResponse)
	if err = json.Unmarshal(response, ret); err != nil {
		return nil, errors.WithStack(err)
	}

	return ret, nil
}

func (c *Client) CreateNote(ctx context.Context, contents misskey.CreateNote) (*misskey.CreateNoteResponse, error) {

	response, err := c.post(ctx, c.createNotes(), contents)
	if err != nil {
		return nil, err
	}

	ret := new(misskey.CreateNoteResponse)
	if err = json.Unmarshal(response, ret); err != nil {
		return nil, errors.WithStack(err)
	}

	return ret, nil
}

func (c *Client) meta() string {
	return fmt.Sprintf("%s/meta", c.url)
}

func (c *Client) createNotes() string {
	return fmt.Sprintf("%s/notes/create", c.url)
}

func (c *Client) post(ctx context.Context, url string, contents interface{}) ([]byte, error) {
	body, err := json.Marshal(contents)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		url,
		bytes.NewBuffer(body),
	)

	req.Header.Add("Content-Type", "application/json")

	if err != nil {
		return nil, errors.WithStack(err)
	}
	res, err := c.client.Do(req)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	if res.StatusCode < http.StatusOK || http.StatusMultipleChoices <= res.StatusCode {
		return nil, errors.Errorf("http bad status: %d\n", res.StatusCode)
	}
	defer res.Body.Close()

	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return bytes, nil
}
