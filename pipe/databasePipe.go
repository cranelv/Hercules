package pipe

type DatabasePipe struct {
	pipe PipeInterface
}
func (dbpipe *DatabasePipe)Connect()error{
	return dbpipe.pipe.Connect()
}
func (dbpipe *DatabasePipe)InsertItem(item interface{},callback func(interface{},error)){
	dbpipe.pipe.Request(Message{DATABASE_Insert,item},callback)
}
func (dbpipe *DatabasePipe)UpdateItem(item interface{},callback func(interface{},error)){
	dbpipe.pipe.Request(Message{DATABASE_Update,item},callback)
}
func (dbpipe *DatabasePipe)InsertUpdateItem(item interface{},callback func(interface{},error)){
	dbpipe.pipe.Request(Message{DATABASE_InsertUpdate,item},callback)
}
func (dbpipe *DatabasePipe)DeleteItem(item interface{},callback func(interface{},error)){
	dbpipe.pipe.Request(Message{DATABASE_Delete,item},callback)
}
