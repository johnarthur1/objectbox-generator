package object

// change UID on an existing property that had an explicitly specified uid before
type ChangeUid struct {
	Id    uint64
	Value string `objectbox:"uid:501233450539197794"`
}
