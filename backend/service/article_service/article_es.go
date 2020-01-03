package article_service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/elastic/go-elasticsearch/v8/esutil"
	"github.com/lawtech0902/go_gin_blog/backend/pkg/setting"
	"github.com/unknwon/com"
	"strings"
)

func (a Article) IndexBlog() error {
	req := esapi.IndexRequest{
		Index:      setting.ESInfo.Index,
		DocumentID: com.ToStr(a.ID),
		Body:       esutil.NewJSONReader(a),
		Refresh:    "true",
	}
	
	res, err := req.Do(context.Background(), es)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	
	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			return err
		}
		return fmt.Errorf("[%s] %s: %s", res.Status(), e["error"].(map[string]interface{})["type"], e["error"].(map[string]interface{})["reason"])
	}
	
	return nil
}

func (a Article) DeleteFromES() error {
	req := esapi.DeleteRequest{
		Index:      setting.ESInfo.Index,
		DocumentID: com.ToStr(a.ID),
		Refresh:    "true",
	}
	
	res, err := req.Do(context.Background(), es)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	
	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			return err
		}
		return fmt.Errorf("[%s] %s: %s", res.Status(), e["error"].(map[string]interface{})["type"], e["error"].(map[string]interface{})["reason"])
	}
	return nil
}

func SearchFromES(opts ...Option) (articles Articles, err error) {
	const searchMatch = `{"query" : {
    "multi_match": {
      "fields":  [ "content", "title" ],
      "query":     "%s",
      "fuzziness": "AUTO"
    }
} }`
	
	var (
		r     map[string]interface{}
		items []Article
		total int
	)
	
	options := newOptions(opts...)
	offset := (options.Page - 1) * options.Limit
	
	res, err := es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithIndex(setting.ESInfo.Index),
		es.Search.WithBody(strings.NewReader(fmt.Sprintf(searchMatch, options.Q))),
		es.Search.WithTrackTotalHits(true),
		es.Search.WithSize(options.Limit),
		es.Search.WithFrom(offset),
	)
	
	if err != nil {
		return articles, fmt.Errorf("error getting response: %v", err)
	}
	defer res.Body.Close()
	
	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			return articles, fmt.Errorf("error parsing the response body: %s", err)
		} else {
			return articles, fmt.Errorf("[%s] %s: %s", res.Status(), e["error"].(map[string]interface{})["type"], e["error"].(map[string]interface{})["reason"])
		}
	}
	
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		return articles, fmt.Errorf("error parsing the response body: %s", err)
	}
	
	for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
		source, _ := json.Marshal(hit.(map[string]interface{})["_source"])
		var a Article
		if err := json.Unmarshal(source, &a); err != nil {
			return articles, fmt.Errorf("error parsing the response body: %s", err)
		}
		items = append(items, a)
	}
	total = int(r["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64))
	
	articles.Items = items
	articles.Total = total
	return articles, nil
}
