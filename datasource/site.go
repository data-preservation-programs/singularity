package datasource

import (
	"context"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/data-preservation-programs/go-singularity/model"
	"github.com/pkg/errors"
	"io"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Site struct {
	Headers map[string]string
}

func (s Site) CheckItem(ctx context.Context, path string) (uint64, *time.Time, error) {
	lastModified, size, err := s.head(ctx, path)
	if err != nil {
		return 0, nil, errors.Wrap(err, "failed to get head")
	}
	return size, &lastModified, nil
}

func (s Site) Scan(ctx context.Context, path string, last string) <-chan Entry {
	entryChan := make(chan Entry)
	go func() {
		defer close(entryChan)
		s.scan(ctx, entryChan, path, last)
	}()
	return entryChan
}

func (s Site) head(ctx context.Context, url string) (time.Time, uint64, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodHead, url, nil)
	if err != nil {
		return time.Time{}, 0, errors.Wrap(err, "failed to create request")
	}
	for k, v := range s.Headers {
		req.Header.Set(k, v)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return time.Time{}, 0, errors.Wrap(err, "failed to perform request")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return time.Time{}, 0, errors.Errorf("failed to get url: %s", resp.Status)
	}

	lastModified, err := http.ParseTime(resp.Header.Get("Last-Modified"))
	if err != nil {
		return time.Time{}, 0, errors.Wrap(err, "failed to parse last modified")
	}

	size, err := strconv.ParseUint(resp.Header.Get("Content-Length"), 10, 64)
	if err != nil {
		return time.Time{}, 0, errors.Wrap(err, "failed to parse content length")
	}

	return lastModified, size, nil
}

func (s Site) scanNginx(ctx context.Context, body []byte, entryChan chan Entry, url string, last string) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(body)))
	if err != nil {
		select {
		case <-ctx.Done():
		case entryChan <- Entry{Error: err}:
		}
		return
	}
	links := make([]*goquery.Selection, 0)
	doc.Find("a").Each(func(_ int, a *goquery.Selection) {
		links = append(links, a)
	})
	sort.Slice(links, func(i, j int) bool { return links[i].AttrOr("href", "") < links[j].AttrOr("href", "") })

	for _, a := range links {
		href, exists := a.Attr("href")
		if !exists || href == "../" {
			continue
		}

		fullPath := url
		if !strings.HasSuffix(fullPath, "/") {
			fullPath += "/"
		}
		fullPath += href

		if strings.HasSuffix(href, "/") {
			// It's a directory, recursively scan
			if !strings.HasPrefix(last, fullPath) && fullPath <= last {
				continue
			}
			newDoc, _, err := s.getHtmlDoc(ctx, fullPath)
			if err != nil {
				select {
				case <-ctx.Done():
				case entryChan <- Entry{Error: err}:
				}
				return
			}
			s.scanNginx(ctx, newDoc, entryChan, fullPath, last)
			if ctx.Err() != nil {
				return
			}
		} else {
			// It's a file, extract the information
			if fullPath <= last {
				continue
			}
			entry := Entry{
				ScannedAt:    time.Now(),
				Type:         model.URL,
				Path:         fullPath,
				Size:         0,
				LastModified: nil,
			}
			lastModified, length, err := s.head(ctx, fullPath)
			if err != nil {
				entryChan <- Entry{Error: err}
				continue
			}
			entry.Size = length
			entry.LastModified = &lastModified
			select {
			case <-ctx.Done():
				return
			case entryChan <- entry:
			}
		}
	}
}

func (s Site) getHtmlDoc(ctx context.Context, url string) ([]byte, string, error) {
	var doc []byte
	var server string
	client := &http.Client{}
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, "", errors.Wrap(err, "failed to create request")
	}

	for k, v := range s.Headers {
		req.Header.Set(k, v)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, "", errors.Wrap(err, "failed to fetch page")
	}

	defer resp.Body.Close()
	doc, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, "", errors.Wrap(err, "failed to read page")
	}

	server = resp.Header.Get("Server")

	return doc, server, nil
}

func (s Site) scan(ctx context.Context, entryChan chan Entry, path string, last string) {
	doc, server, err := s.getHtmlDoc(ctx, path)
	if err != nil {
		select {
		case <-ctx.Done():
		case entryChan <- Entry{Error: err}:
		}
		return
	}

	if strings.HasPrefix(server, "nginx/") {
		s.scanNginx(ctx, doc, entryChan, path, last)
	}
}

func (s Site) Read(ctx context.Context, path string, offset uint64, length uint64) (io.ReadCloser, error) {
	// Create an HTTP client with the context
	client := http.Client{
		Transport: &http.Transport{},
	}

	// Create a new request
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create request")
	}

	for k, v := range s.Headers {
		req.Header.Set(k, v)
	}

	// Set the "Range" header for partial download
	rangeValue := fmt.Sprintf("bytes=%d-%d", offset, offset+length-1)
	req.Header.Set("Range", rangeValue)

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to send request")
	}

	// Check the response status code
	if resp.StatusCode != http.StatusPartialContent {
		resp.Body.Close()
		return nil, errors.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return resp.Body, nil
}

func (s Site) Open(ctx context.Context, path string) (ReadAtCloser, error) {
	return HTTPReadAtCloser{
		path:    path,
		headers: s.Headers,
		ctx:     ctx,
	}, nil
}

type HTTPReadAtCloser struct {
	path    string
	headers map[string]string
	ctx     context.Context
}

func (s HTTPReadAtCloser) ReadAt(p []byte, off int64) (n int, err error) {
	// Create an HTTP client with the context
	client := http.Client{
		Transport: &http.Transport{},
	}

	// Create a new request
	req, err := http.NewRequestWithContext(s.ctx, http.MethodGet, s.path, nil)
	if err != nil {
		return 0, errors.Wrap(err, "failed to create request")
	}

	for k, v := range s.headers {
		req.Header.Set(k, v)
	}

	// Set the "Range" header for partial download
	rangeValue := fmt.Sprintf("bytes=%d-%d", off, off+int64(len(p))-1)
	req.Header.Set("Range", rangeValue)

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		return 0, errors.Wrap(err, "failed to send request")
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusPartialContent {
		return 0, errors.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return resp.Body.Read(p)
}

func (s HTTPReadAtCloser) Close() error {
	return nil
}
