package cache

import(
cach "ifix.cbpc/cbpc/pkg/cache/cache"
"time"
)
var c *cach.Cache

func init(){
	c=cach.New(cach.NoExpiration,1*time.Hour)
}

//读取缓存数据
func CacheGet(path string) (cachemap map[string]string,err error) {
	cachemap=make(map[string]string,0)
	err=c.LoadFile(path)
	if err != nil {
		return
	}
	item:=c.Items()
	for k, v := range item {
		cachemap[k]=v.Object.(string)
	}
	return
}

//缓存数据
func CacheSet(cachemap map[string]string,path string) (err error) {
	for k, v := range cachemap {
		c.SetDefault(k,v)
	}
	err=c.SaveFile(path)
	return
}

//更新缓存数据
func CacheUpdate(cachemap map[string]string,path string) (err error) {
	for k, v := range cachemap {
		c.ReplaceDefault(k,v)
	}
	err=c.SaveFile(path)
	return
}