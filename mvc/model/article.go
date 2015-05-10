package model

import (
	"github.com/gofxh/blog/lib/base"
	"github.com/gofxh/blog/lib/entity"
	"sort"
	"sync"
	"time"
)

var (
	ArticleData    map[int64]*entity.Article = make(map[int64]*entity.Article)
	ArticleAllList []int
	ArticlePubList []int

	articleSlugData map[string]int64 = make(map[string]int64)

	mu sync.Mutex
)

// load article data
func loadArticleData() {
	base.Storage.Walk(new(entity.Article), func(v interface{}) {
		if a, ok := v.(*entity.Article); ok {
			ArticleData[a.Id] = a
			articleSlugData[a.Slug] = a.Id
		}
	})
	loadArticleList()
}

// load article list
func loadArticleList() {
	mu.Lock()
	allList := make([]int, 0)
	pubList := make([]int, 0)
	for _, a := range ArticleData {
		// all list contains private and public items, except deleted
		if a.Status != entity.ARTICLE_STATUS_DELETED {
			allList = append(allList, int(a.Id))
		}
		// pub list contains public items
		if a.Status == entity.ARTICLE_STATUS_PUBLIC {
			pubList = append(pubList, int(a.Id))
		}
	}
	// order by id desc
	sort.Reverse(sort.IntSlice(allList))
	sort.Reverse(sort.IntSlice(pubList))
	ArticleAllList = allList
	ArticlePubList = pubList
	mu.Unlock()
}

// save article
func SaveArticle(a *entity.Article) *entity.Article {
	a.Id = base.Max.Next(base.Max.ArticleId, base.Max.ArticleStep)
	a.CreateTime = time.Now().Unix()
	a.UpdateTime = a.CreateTime

	// todo: update category and tag list

	base.Storage.Save(a)
	ArticleData[a.Id] = a
	articleSlugData[a.Slug] = a.Id
	loadArticleList()
	return a
}

// get article by slug
func GetArticleBySlug(slug string) *entity.Article {
	id := articleSlugData[slug]
	if id == 0 {
		return nil
	}
	return ArticleData[id]
}

// clean deleted articles,
// remove them forever
func CleanArticles() []*entity.Article {
	m.Lock()
	result := make([]*entity.Article, 0)
	for _, a := range ArticleData {
		if a.Status == entity.ARTICLE_STATUS_DELETED {
			base.Storage.Remove(a)
			delete(ArticleData, a.Id)
			delete(articleSlugData, a.Slug)

			result = append(result, a)
		}
	}
	m.Unlock() // remember to unlock, locker is used in loading article list function
	loadArticleList()
	return result
}
