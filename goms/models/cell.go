package models

import (
	"strconv"

	"github.com/astaxie/beego/orm"
)

//Cell 小区信息
type Cell struct {
	Id     int64
	Cuid   string
	Duid   string `orm:"index"`
	Cellid string `orm:"index"`
	Freq   uint32
}

func AddCell(cuid, duid, cellid string) error {
	o := orm.NewOrm()
	cell := &Cell{
		Cuid:   cuid,
		Duid:   duid,
		Cellid: cellid,
		Freq:   39000,
	}

	qs := o.QueryTable("cell")
	err := qs.Filter("Cuid", cuid).Filter("Duid", duid).Filter("Cellid", cellid).One(cell)
	if err == nil {
		return err
	}

	_, err = o.Insert(cell)
	return err
}

func DelCell(id string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}

	o := orm.NewOrm()

	cell := &Cell{Id: cid}
	_, err = o.Delete(cell)
	return err
}

func GetAllCells() ([]*Cell, error) {
	o := orm.NewOrm()

	cells := make([]*Cell, 0)

	qs := o.QueryTable("cell")
	_, err := qs.All(&cells)
	return cells, err
}
