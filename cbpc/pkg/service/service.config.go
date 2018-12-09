package service

//预读配置项和备份
func ConfigMap()(conf map[string]string,err error){
	conf=make(map[string]string, 0)
	err=SQLiteInit("sqlite3","./ifixtoolsCache.db")
	conf,err=DBExecSelectMap4Config(db3,"select Key,Value from iFixToolsConfig")
	return
}
