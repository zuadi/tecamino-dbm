package handlers

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"runtime"
	"time"

	"github.com/zuadi/tecamino-dbm.git/models"
	"github.com/zuadi/tecamino-dbm.git/server"
)

type DBMHandler struct {
	filePath string
	DB       models.Datapoint
	Server   *server.Server
}

func NewDbmHandler(rootDir, dbName string) (*DBMHandler, error) {
	dmaHandler := DBMHandler{
		filePath: fmt.Sprintf("%s/%s.dma", rootDir, dbName),
	}

	if _, err := os.Stat(dmaHandler.filePath); err == nil {

		f, err := os.Open(dmaHandler.filePath)
		if err != nil {
			return nil, err
		}
		defer f.Close()

		scanner := bufio.NewScanner(f)

		for scanner.Scan() {
			dp := models.Datapoint{}
			err = json.Unmarshal(scanner.Bytes(), &dp)
			if err != nil {
				return nil, err
			}
			dmaHandler.ImportDatapoints(&dp)
		}
	}
	if err := dmaHandler.AddSystemDps(); err != nil {
		return nil, err
	}

	return &dmaHandler, nil
}

func (d *DBMHandler) SaveDb() (err error) {
	f, err := os.OpenFile(d.filePath, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	defer f.Close()

	for _, dp := range d.DB.GetAllDatapoints() {
		b, er := json.Marshal(dp)
		if er != nil {
			return er
		}

		_, err = f.Write(b)
		if err != nil {
			return
		}
		_, err = f.Write([]byte("\n"))
		if err != nil {
			return
		}
	}
	return
}

func (d *DBMHandler) CreateDatapoint(typ models.Type, value any, right models.Rights, path ...string) error {
	return d.DB.CreateDatapoint(typ, value, right, path...)
}

func (d *DBMHandler) ImportDatapoints(dps ...*models.Datapoint) error {
	for _, dp := range dps {
		err := d.DB.ImportDatapoint(dp, dp.Path)
		if err != nil {
			return err
		}
	}
	return nil
}

func (d *DBMHandler) UpdateDatapointValue(path string, value any) error {
	return d.DB.UpdateDatapointValue(value, path)
}

func (d *DBMHandler) RemoveDatapoint(path string) error {
	return d.DB.RemoveDatapoint(path)
}

func (d *DBMHandler) QueryDatapoints(key string) []*models.Datapoint {
	return d.DB.QueryDatapoints(key)
}

func (d *DBMHandler) AddSystemDps() (err error) {

	go func() {
		tim := "System:Time"
		memory := "System:UsedMemory"
		var m runtime.MemStats

		err = d.DB.CreateDatapoint(models.STR, nil, models.Read, tim)
		if err != nil {
			return
		}
		err = d.DB.CreateDatapoint(models.STR, nil, models.Read, memory)
		if err != nil {
			return
		}
		for {
			t := time.Now().UnixMilli()
			if er := d.DB.UpdateDatapointValue(t, tim); er != nil {
				if err = d.DB.CreateDatapoint(models.STR, nil, models.Read, tim); err != nil {
					log.Fatal(err)
				}
			}
			runtime.ReadMemStats(&m)
			if er := d.DB.UpdateDatapointValue(fmt.Sprintf("%.2f MB", float64(m.Sys)/1024/1024), memory); er != nil {
				if err = d.DB.CreateDatapoint(models.STR, nil, models.Read, memory); err != nil {
					log.Fatal(err)
				}
			}
			time.Sleep(1 * time.Second)
		}
	}()
	return
}
