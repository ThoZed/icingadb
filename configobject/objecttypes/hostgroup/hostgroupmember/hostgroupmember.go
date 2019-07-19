package hostgroupmember

import (
	"git.icinga.com/icingadb/icingadb-main/configobject"
	"git.icinga.com/icingadb/icingadb-main/connection"
	"git.icinga.com/icingadb/icingadb-main/utils"
)

var (
	ObjectInformation configobject.ObjectInformation
	Fields         = []string{
		"id",
		"hostgroup_id",
		"host_id",
		"env_id",
	}
)

type HostgroupMember struct {
	Id						string 		`json:"id"`
	HostgroupId				string		`json:"group_id"`
	HostId		 			string 		`json:"object_id"`
	EnvId           		string		`json:"env_id"`
}

func NewHostgroupMember() connection.Row {
	h := HostgroupMember{}
	return &h
}

func (c *HostgroupMember) InsertValues() []interface{} {
	v := c.UpdateValues()

	return append([]interface{}{utils.Checksum(c.Id)}, v...)
}

func (h *HostgroupMember) UpdateValues() []interface{} {
	v := make([]interface{}, 0)

	v = append(
		v,
		utils.Checksum(h.HostgroupId),
		utils.Checksum(h.HostId),
		utils.Checksum(h.EnvId),
	)

	return v
}

func (h *HostgroupMember) GetId() string {
	return h.Id
}

func (h *HostgroupMember) SetId(id string) {
	h.Id = id
}

func init() {
	name := "hostgroup_member"
	ObjectInformation = configobject.ObjectInformation{
		ObjectType: name,
		RedisKey: "host:groupmember",
		Factory: NewHostgroupMember,
		HasChecksum: false,
		BulkInsertStmt: connection.NewBulkInsertStmt(name, Fields),
		BulkDeleteStmt: connection.NewBulkDeleteStmt(name),
		BulkUpdateStmt: connection.NewBulkUpdateStmt(name, Fields),
	}
}