package client

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/fantasy9830/go-retry"
)

type Client struct {
	pool    sync.Pool
	Options *Options
}

func NewClient() *Client {
	return &Client{
		pool: sync.Pool{
			New: func() any {
				return bytes.NewBuffer(make([]byte, 4096))
			},
		},
		Options: &Options{
			Headers: make(map[string]string),
		},
	}
}

func (c *Client) Request(method string, rawURL string, payload io.ReadSeeker, optFuncs ...OptionFunc) (body []byte, err error) {
	for _, applyFunc := range optFuncs {
		applyFunc(c.Options)
	}

	retryOpt := []retry.OptionFunc{
		retry.MaxRetries(3),
		retry.WithBackoff(retry.BackoffLinear(3 * time.Second)),
	}

	err = retry.Do(func(ctx context.Context) error {
		if payload != nil {
			if _, err := payload.Seek(0, io.SeekStart); err != nil {
				return err
			}
		}

		body, err = c.request(method, rawURL, payload)
		return err
	}, retryOpt...)

	return body, err
}

// request request
func (c *Client) request(method string, rawURL string, payload io.Reader) ([]byte, error) {
	u, err := url.ParseRequestURI(rawURL)
	if err != nil {
		slog.Error(err.Error())
		return nil, err
	}

	req, err := http.NewRequest(method, u.String(), payload)
	if err != nil {
		slog.Error(err.Error())
		return nil, err
	}

	for key, value := range c.Options.Headers {
		req.Header.Add(key, value)
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		slog.Error(err.Error())
		return nil, err
	}
	defer func() {
		_, _ = io.Copy(io.Discard, res.Body)
		res.Body.Close()
	}()

	buffer := c.pool.Get().(*bytes.Buffer)
	buffer.Reset()
	defer func() {
		if buffer != nil {
			c.pool.Put(buffer)
			buffer = nil
		}
	}()
	_, err = io.Copy(buffer, res.Body)
	if err != nil {
		slog.Error(err.Error())
		return nil, err
	}

	body := buffer.Bytes()
	c.pool.Put(buffer)
	buffer = nil

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		err := fmt.Errorf("[%d] %s: %s", res.StatusCode, res.Status, string(body))
		slog.Error(err.Error())
		return nil, err
	}

	if c.Options.CheckFunc != nil {
		if err := c.Options.CheckFunc(body); err != nil {
			slog.Error(err.Error())
			return nil, err
		}
	}

	return body, nil
}
