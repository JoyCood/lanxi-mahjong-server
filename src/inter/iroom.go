/**********************************************************
 * Author        : Michael
 * Email         : dolotech@163.com
 * Last modified : 2016-01-23 10:42
 * Filename      : iroom.go
 * Description   : 房间的数据接口
 * *******************************************************/
package inter


type IDesk interface {
	DiscardL(uint32, byte, bool)
	TurnL()
	OperateL(uint32, uint32, uint32)
	Readying(uint32, bool) int
	Enter(IPlayer) int
	GetData() interface{}
	Diceing() bool
	Leave(uint32) bool
	Kick(string, uint32) bool
	Vote(bool, uint32, uint32) int
	Trust(uint32, uint32)
	Broadcasts(IProto)
	Closed(bool)
	Print()
}
