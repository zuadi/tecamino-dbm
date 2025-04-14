package models

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

type Datapoint struct {
	Datapoints     map[string]*Datapoint `json:"-"`
	Path           string                `json:"path"`
	Value          any                   `json:"value,omitempty"`
	CreateDateTime int64                 `json:"createDateTime,omitempty"`
	UpdateDateTime int64                 `json:"updateDateTime,omitempty"`
	Type           Type                  `json:"type"`
	ReadWrite      Rights                `json:"readWrite"`
}

var depth1 int

func (d *Datapoint) CreateDatapoint(typ Type, value any, rights Rights, paths ...string) error {
	l := len(paths) - 1
	if l == 0 {
		paths = regexp.MustCompile(`[:]+`).Split(paths[0], -1)
		l = len(paths) - 1
	}

	if d.Datapoints == nil {
		d.Datapoints = make(map[string]*Datapoint)
	}

	if l == depth1 {
		if od, ok := d.Datapoints[paths[depth1]]; !ok {
			d.Datapoints[paths[depth1]] = &Datapoint{
				Path:           strings.Join(paths, ":"),
				Type:           typ,
				Value:          typ.ConvertValue(value),
				ReadWrite:      rights.GetRights(),
				CreateDateTime: time.Now().UnixMilli(),
				UpdateDateTime: time.Now().UnixMilli(),
			}
		} else {
			od.Type = typ
			od.ReadWrite = rights.GetRights()
			od.Value = typ.ConvertValue(value)
			od.UpdateDateTime = time.Now().UnixMilli()
		}
		depth1 = 0
	} else if datapoint, ok := d.Datapoints[paths[depth1]]; ok {
		depth1 += 1
		datapoint.CreateDatapoint(typ, value, rights, paths...)
	} else {
		da := Datapoint{
			Path:           strings.Join(paths[:depth1+1], ":"),
			Type:           NONE,
			ReadWrite:      rights.GetRights(),
			CreateDateTime: time.Now().UnixMilli(),
			UpdateDateTime: time.Now().UnixMilli(),
		}

		d.Datapoints[paths[depth1]] = &da
		depth1 += 1
		da.CreateDatapoint(typ, value, rights, paths...)
	}
	return nil
}

var depth2 int

func (d *Datapoint) ImportDatapoint(dp *Datapoint, paths ...string) error {
	l := len(paths) - 1
	if l == 0 {
		paths = regexp.MustCompile(`[:]+`).Split(paths[0], -1)
		l = len(paths) - 1
	}

	if d.Datapoints == nil {
		d.Datapoints = make(map[string]*Datapoint)
	}
	if l == depth2 {
		if od, ok := d.Datapoints[paths[depth2]]; !ok {
			d.Datapoints[paths[depth2]] = dp
			dp.ReadWrite = dp.ReadWrite.GetRights()
			dp.UpdateDateTime = time.Now().UnixMilli()
		} else {
			od.Type = dp.Type
			od.Value = d.Type.ConvertValue(dp.Value)
			od.ReadWrite = dp.ReadWrite.GetRights()
			od.UpdateDateTime = time.Now().UnixMilli()
		}
		depth2 = 0
	} else if datapoint, ok := d.Datapoints[paths[depth2]]; ok {
		depth2 += 1
		datapoint.ImportDatapoint(dp, paths...)
	} else {
		da := Datapoint{
			Path:           strings.Join(paths[:depth2+1], ":"),
			Type:           NONE,
			UpdateDateTime: time.Now().UnixMilli(),
		}
		da.ReadWrite = da.ReadWrite.GetRights()
		d.Datapoints[paths[depth2]] = &da
		depth2 += 1
		da.ImportDatapoint(dp, paths...)
	}
	return nil
}

var depth3 int

func (d *Datapoint) UpdateDatapointValue(value any, paths ...string) error {
	l := len(paths) - 1
	if l == 0 {
		paths = regexp.MustCompile(`[:]+`).Split(paths[0], -1)
		l = len(paths) - 1
	}

	if l == depth3 {
		if dp, ok := d.Datapoints[paths[depth3]]; ok {
			fmt.Print("update dp:", dp.Path, " old value:", dp.Value)
			dp.Value = dp.Type.ConvertValue(value)
			fmt.Println(" new value:", dp.Value)

			dp.UpdateDateTime = time.Now().UnixMilli()
			depth3 = 0
			return nil
		}
		depth3 = 0
		return fmt.Errorf("datapoint '%s' not found", strings.Join(paths, ":"))
	} else if datapoint, ok := d.Datapoints[paths[depth3]]; ok {
		depth3 += 1
		if err := datapoint.UpdateDatapointValue(value, paths...); err != nil {
			fmt.Println(100, err)
			return err
		}
	}
	depth3 = 0
	return nil
}

var depth4 int

func (d *Datapoint) RemoveDatapoint(paths ...string) error {
	l := len(paths) - 1
	if l == 0 {
		paths = regexp.MustCompile(`[:]+`).Split(paths[0], -1)
		l = len(paths) - 1
	}

	if l == depth4 {
		if _, ok := d.Datapoints[paths[depth4]]; ok {
			delete(d.Datapoints, paths[depth4])
			fmt.Println("removed dp:", strings.Join(paths, ":"))
		}
		depth4 = 0
	} else if datapoint, ok := d.Datapoints[paths[depth4]]; ok {
		depth4 += 1
		datapoint.RemoveDatapoint(paths...)
	}
	return fmt.Errorf("datapoint '%s' not found", strings.Join(paths, ":"))
}

func (d *Datapoint) GetAllDatapoints() (dps []*Datapoint) {
	for _, dp := range d.Datapoints {
		dps = append(dps, dp.GetAllDatapoints()...)
		dps = append(dps, dp)

	}
	return
}

func (d *Datapoint) QueryDatapoints(key string) (dps []*Datapoint) {
	reg := regexp.MustCompile(key)
	for _, dp := range d.Datapoints {
		if reg.MatchString(dp.Path) {
			dps = append(dps, dp)
		}
		dps = append(dps, dp.QueryDatapoints(key)...)

	}
	return
}
