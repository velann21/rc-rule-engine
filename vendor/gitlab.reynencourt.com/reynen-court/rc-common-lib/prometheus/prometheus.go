package prometheus

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type Prometheus struct {
	url     string
	timeout time.Duration
}

func New(url string) *Prometheus {
	return &Prometheus{url: url, timeout: 30 * time.Second}
}

func NewWithTimeout(url string, duration time.Duration) *Prometheus {
	return &Prometheus{url: url, timeout: duration}
}

func (p *Prometheus) GetLabel(ctx context.Context) (string, error) {

	labelUrl, err := url.Parse(fmt.Sprintf("%v/api/v1/labels", p.url))
	if err != nil {
		return "", err
	}

	var client = http.Client{
		Timeout: p.timeout,
	}

	var request = http.Request{
		URL: labelUrl,
	}

	request.WithContext(ctx)

	resp, err := client.Do(&request)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != 200 {
		return "", errors.New("error while reading data")
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

type Alert struct {
	Label map[string]string `json:"labels"`
	Annotation  map[string]string  `json:"annotations"`
	StartedAt string `json:"startsAt"`
	EndsAt string `json:"endsAt"`
	GeneratorURL string `json:"generatorURL"`
}

func (p *Prometheus) AddAlerts(ctx context.Context, alertRequest... *Alert)  error {

	for _, alert := range alertRequest {
		if alert == nil {
			return errors.New("alert request cannot be empty")
		}

		if alert.Label["alertname"] == "" {
			return  errors.New("alertname should be set in the labels map")
		}
	}

	d, err := json.Marshal(alertRequest)
	if err != nil {
		return err
	}

	alertURI, err := url.Parse(fmt.Sprintf("%v/api/v1/alerts", p.url))
	if err != nil {
		return  err
	}

	log.Println(alertURI)

	var client = http.Client{
		Timeout: p.timeout,
	}

	var request = http.Request{
		URL: alertURI,
		Method:"POST",
		Body:ioutil.NopCloser(bytes.NewReader(d)),
	}

	request.WithContext(ctx)

	resp, err := client.Do(&request)
	if err != nil {
		return  err
	}

	ad, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return errors.New(string(ad))
	}

	return nil
}

func (p *Prometheus) GetAlerts(ctx context.Context) (string, error) {
	labelUrl, err := url.Parse(fmt.Sprintf("%v/api/v1/alerts", p.url))
	if err != nil {
		return "", err
	}

	var client = http.Client{
		Timeout: p.timeout,
	}

	var request = http.Request{
		URL: labelUrl,
	}

	request.WithContext(ctx)

	resp, err := client.Do(&request)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != 200 {
		return "", errors.New("error while fetching alerts")
	}
	
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func (p *Prometheus) QueryBySeries(ctx context.Context, match []string, startTime time.Time, endTime time.Time) (string, error) {

	var matchQuery []string

	for _, m := range match {
		matchQuery = append(matchQuery, fmt.Sprintf("match[]=%v", url.QueryEscape(m)))
	}

	queryURL := fmt.Sprintf("%v/api/v1/series?%v&start=%s&end=%s", p.url, strings.Join(matchQuery, "&"), startTime.UTC().Format(time.RFC3339Nano), endTime.UTC().Format(time.RFC3339Nano))

	log.Println(queryURL)

	rangeQuery, err := url.Parse(queryURL)
	if err != nil {
		return "", err
	}

	var client = http.Client{
		Timeout: p.timeout,
	}

	var request = http.Request{
		URL: rangeQuery,
	}

	request.WithContext(ctx)

	resp, err := client.Do(&request)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != 200 {
		return "", errors.New("error while fetching data")
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

type RangeQuery struct {
	Query     string
	StartTime time.Time
	EndTime   time.Time
	Step      string
}

func (p *Prometheus) QueryByRange(ctx context.Context, rangeQuery *RangeQuery) (string, error) {

	if rangeQuery == nil {
		return "", errors.New("query cannot be nil")
	}

	var request http.Request
	var client http.Client

	queryURL := fmt.Sprintf("%v/api/v1/query_range?query=%v&start=%s&end=%s&step=%v", p.url,
		url.QueryEscape(rangeQuery.Query),
		rangeQuery.StartTime.UTC().Format(time.RFC3339Nano),
		rangeQuery.EndTime.UTC().Format(time.RFC3339Nano),
		rangeQuery.Step)

	rangeQueryURI, err := url.Parse(queryURL)
	if err != nil {
		return "", err
	}

	request.URL = rangeQueryURI
	request.WithContext(ctx)

	resp, err := client.Do(&request)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != 200 {
		return "", errors.New("error while fetching data")
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

//https://github.com/prometheus/prometheus/blob/8c8bb82d04e854d82555150ca2e71e2bc159df90/web/api/v1/api.go#L1153
func parseTime(s string) (time.Time, error) {
	if t, err := strconv.ParseFloat(s, 64); err == nil {
		s, ns := math.Modf(t)
		ns = math.Round(ns*1000) / 1000
		return time.Unix(int64(s), int64(ns*float64(time.Second))), nil
	}
	if t, err := time.Parse(time.RFC3339Nano, s); err == nil {
		return t, nil
	}
	return time.Time{}, errors.Errorf("cannot parse %q to a valid timestamp", s)
}
