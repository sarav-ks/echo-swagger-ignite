package db

import (
	"encoding/json"
	"gwapp/model"
	"log"
	"net"
	"net/http"
	// "reflect"
	"time"

	"github.com/amsokol/ignite-go-client/binary/v1"
	"github.com/labstack/echo/v4"
	 "github.com/xorcare/pointer"
)

type IgniteStore struct {
	Clnt ignite.Client
	//TODO: add lock & auto gen id
}

func NewEventStore() *IgniteStore {
	return &IgniteStore{
	}
}

func (i *IgniteStore) InitStore() {
	log.Print("Strating the ignite client")
	var err error
	i.Clnt, err = ignite.Connect(ignite.ConnInfo{
		Network: "tcp",	
		Host:    "localhost",		
		Port:    30589,		
		Major:   1,
		Minor:   1,
		Patch:   0,
		// Credentials are only needed if they're configured in your Ignite server.
		Username: "ignite",
		Password: "ignite",
		Dialer: net.Dialer{
			Timeout: 10 * time.Second,
		},
	})
	if err != nil {
		log.Fatalf("failed connect to server: %v", err)
	}
	//defer i.Clnt.Close()
}

func (i *IgniteStore) getOrCreateStore() ignite.Client {
	if (i.Clnt == nil) {
		i.InitStore();
	}
	return i.Clnt;
}

func (i *IgniteStore) AddEvent(ctx echo.Context) error {
	var newEvent model.Event
	err := ctx.Bind(&newEvent)

	if err != nil {
		return sendEventError(ctx, http.StatusBadRequest, "Invalid format for Event")
	}

	ic := i.getOrCreateStore()

	cconfig,err := ic.CacheGetConfiguration("EventCache",0)
	cv,_ := json.Marshal(cconfig)
	println("Config Values for Cache in ADD = ",string(cv) , err)

	_, err = ic.QuerySQLFields("EventCache", false, ignite.QuerySQLFieldsData{
		PageSize: 10,
		Query: "INSERT INTO Event(Id, Location,Name,Tags) VALUES (?, ?, ?, ?)" ,
		QueryArgs: []interface{}{
			newEvent.Id, newEvent.Location,newEvent.Name,newEvent.Tag,
		},
	})
	if err != nil {
		log.Fatalf("failed insert data: %v", err)
	}

	err = ctx.JSON(http.StatusCreated, newEvent)
	if err != nil {
		// Something really bad happened, tell Echo that our handler failed
		return err
	}

	return nil
}

func sendEventError(ctx echo.Context, code int, message string) error {
	petErr := model.Error{
		Code:    int32(code),
		Message: message,
	}
	err := ctx.JSON(code, petErr)
	return err
}

func (i *IgniteStore) FindEventById(ctx echo.Context, eId int64) error {

	ic := i.getOrCreateStore()

	cconfig,err := ic.CacheGetConfiguration("EventCache",0)
	cv,_ := json.Marshal(cconfig)
	if(err != nil){
		log.Printf("Error reading cache config %v" , err)
	}
	println("Config Values for Cache = ",string(cv),err)

	r, err := ic.QuerySQLFields("EventCache", false, ignite.QuerySQLFieldsData{
		Query:    "SELECT * FROM Event Where Id = ? ",
		PageSize: 10000,
		QueryArgs: []interface{}{
			eId},
		Timeout:           10000,
		IncludeFieldNames: true,		
	})
	if err != nil {
		log.Fatalf("failed query data: %v", err)
	}
	log.Printf("res=%#v", r.Rows)
	var e  model.Event

	for i := 0; i < len(r.Rows); i++ {
		e =  model.Event {
			Id: pointer.Int64(r.Rows[int64(i)][0].(int64)),
			Name: r.Rows[int64(i)][1].(string),
			Location: pointer.String(r.Rows[int64(i)][2].(string)),
			Tag: pointer.String(r.Rows[int64(i)][3].(string)),
		}	
	}

	return ctx.JSON(http.StatusOK, e)
}

func getPointerInt(v int64) *int64 {
	return &v;
}

func getPointerString(p string) *string {
	return &p;
}

func getValueFromCacheObject(co ignite.ComplexObject, key string) interface{}{
	r,_ := co.Get(key)
	return r
}

func (i *IgniteStore) DeleteEvent(ctx echo.Context, id int64) error {
	return sendEventError(ctx,http.StatusMethodNotAllowed,"Not implemented")
}

func (i *IgniteStore) FindEvents(ctx echo.Context, params model.FindEventsParams) error {
	return sendEventError(ctx,http.StatusMethodNotAllowed,"Not implemented")
}
