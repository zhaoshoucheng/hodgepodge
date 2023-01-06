package app

type Service1 interface {
	AddData(string)
	DelData(string)
}
type Service2 interface {
	AddData(string)
	DelData(string)
	SyncData(t int)
	InitService(done InitializeDoneProc)
}

type Module interface {
	DataToSave(string)
	DataToRemove(string)
}

type Resource interface {
	Save(string)
	Remove(string)
}
