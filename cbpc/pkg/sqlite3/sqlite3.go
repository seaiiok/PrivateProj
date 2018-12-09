package sqlite3
import(
    "database/sql"
    "fmt"
	_ "ii.sea/libs/SQLite"
	"time"
)
var SQLiteDB *sql.DB
func SQLiteDBInit()(err error){
    SQLiteDB, err = sql.Open("sqlite3", "./ifixtoolsCache.db")
    err=SQLiteDB.Ping()
    s2:=SQLiteDB.Stats()
    fmt.Println(s2.Idle)
    return 
}

func SQLiteExec(sqlstr string)(err error){
    // sqlstr = `create table iFixToolsServer (Id INTEGER not null primary key, ServerIp text not null, ServerPort text not null);`
    // "CREATE INDEX inx_Id ON iFixToolsServer(Id);"
    _, err = SQLiteDB.Exec(sqlstr)
    return
}

func SQLitePreExec(sqlstr string)(err error){
    // "insert into iFixToolsServer(Id, ServerIp, ServerPort) values(?,?,?)"
    stmt, err := SQLiteDB.Prepare(sqlstr)
    _, err = stmt.Exec(sqlstr)
    return
}


type iFixToolsServerDB struct{
	Id int
	ServerIp string
	ServerPort string
}
func SQLiteDBInit1(){
	db, err := sql.Open("sqlite3", "./ifixtoolsCache.db")
    if err != nil {
        fmt.Println("SQLite:", err)
    }
    defer db.Close()
	fmt.Println("SQLite start")
	
	//创建表//delete from BC;，SQLite字段类型比较少，bool型可以用INTEGER，字符串用TEXT
	// sqlStmt:=`if not exists (select * from sysobjects where id = object_id('BC')) create table BC (b_code text not null primary key, c_code text not null, code_type INTEGER, is_new INTEGER);`
    sqlStmt := `create table iFixToolsServer (Id INTEGER not null primary key, ServerIp text not null, ServerPort text not null);`
    _, err = db.Exec(sqlStmt)
    if err != nil {
        fmt.Println(err, sqlStmt)
        return
    }
    // 创建索引，有索引和没索引性能差别巨大，根本就不是一个量级，有兴趣的可以去掉试试
    _, err = db.Exec("CREATE INDEX inx_Id ON iFixToolsServer(Id);")
    if err != nil {
        fmt.Println(err, sqlStmt)
        return
    }
    //写入10M条记录
    start := time.Now().Unix()
    tx, err := db.Begin()
    if err != nil {
        fmt.Println(err)
    }
    stmt, err := tx.Prepare("insert into iFixToolsServer(Id, ServerIp, ServerPort) values(?,?,?)")
    if err != nil {
        fmt.Println(err)
    }
    defer stmt.Close()
    // var m int = 1000
    var total int = 20
    for i := 0; i < total; i++ {
        _, err = stmt.Exec(fmt.Sprintf("B%024d", i), fmt.Sprintf("C%024d", i), 0, 1)
        if err != nil {
            fmt.Println(err)
        }
    }
    tx.Commit()
    insertEnd := time.Now().Unix()
    //随机检索10M次
    var count int64 = 0

    stmt, err = db.Prepare("select * from iFixToolsServer")
    if err != nil {
        fmt.Println( err)
    }
	defer stmt.Close()
	
    bc := new(iFixToolsServerDB)
    for i := 0; i < total; i++ {

        err = stmt.QueryRow(fmt.Sprintf("C%024d", i)).Scan(&bc.Id, &bc.ServerIp, &bc.ServerPort)
        if err != nil {
            fmt.Println(err)
        }
        //屏幕输出会花掉好多时间啊，计算耗时的时候还是关掉比较好
        // fmt.Println("BCode=", bc.B_Code, "\tCCode=", bc.C_Code, "\tCodeType=", bc.CodeType, "\tIsNew=", bc.IsNew)
        count++
    }
    readEnd := time.Now().Unix()
    fmt.Println("insert span=", (insertEnd - start),
        "read span=", (readEnd - insertEnd),
		"avg read=", float64(readEnd-insertEnd)*1000/float64(count))
		
		
			
}