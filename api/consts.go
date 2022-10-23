package api

/* basic types 基础数据类型 */
const (
	LUA_TNONE          = iota - 1 // -1， 指无效索引
	LUA_TNIL                      // nil
	LUA_TBOOLEAN                  // boolean
	LUA_TLIGHTUSERDATA            // TODO light userdata 暂时未知
	LUA_TNUMBER                   // number
	LUA_TSTRING                   // string
	LUA_TTABLE                    // table
	LUA_TFUNCTION                 // function
	LUA_TUSERDATA                 // userdata
	LUA_TTHREAD                   // thread
)
