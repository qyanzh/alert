// Code generated by thriftgo (0.1.2). DO NOT EDIT.

package rpc_dto

import (
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"strings"
)

type Alert struct {
	Id       int64             `thrift:"Id,1" json:"Id"`
	Time     string            `thrift:"Time,2" json:"Time"`
	IndexNum map[int64]float64 `thrift:"IndexNum,3" json:"IndexNum"`
}

func NewAlert() *Alert {
	return &Alert{}
}

func (p *Alert) GetId() int64 {
	return p.Id
}

func (p *Alert) GetTime() string {
	return p.Time
}

func (p *Alert) GetIndexNum() map[int64]float64 {
	return p.IndexNum
}
func (p *Alert) SetId(val int64) {
	p.Id = val
}
func (p *Alert) SetTime(val string) {
	p.Time = val
}
func (p *Alert) SetIndexNum(val map[int64]float64) {
	p.IndexNum = val
}

var fieldIDToName_Alert = map[int16]string{
	1: "Id",
	2: "Time",
	3: "IndexNum",
}

func (p *Alert) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.I64 {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		case 2:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField2(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		case 3:
			if fieldTypeId == thrift.MAP {
				if err = p.ReadField3(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}

		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_Alert[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *Alert) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(); err != nil {
		return err
	} else {
		p.Id = v
	}
	return nil
}

func (p *Alert) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		p.Time = v
	}
	return nil
}

func (p *Alert) ReadField3(iprot thrift.TProtocol) error {
	_, _, size, err := iprot.ReadMapBegin()
	if err != nil {
		return err
	}
	p.IndexNum = make(map[int64]float64, size)
	for i := 0; i < size; i++ {
		var _key int64
		if v, err := iprot.ReadI64(); err != nil {
			return err
		} else {
			_key = v
		}

		var _val float64
		if v, err := iprot.ReadDouble(); err != nil {
			return err
		} else {
			_val = v
		}

		p.IndexNum[_key] = _val
	}
	if err := iprot.ReadMapEnd(); err != nil {
		return err
	}
	return nil
}

func (p *Alert) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("Alert"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}
		if err = p.writeField2(oprot); err != nil {
			fieldId = 2
			goto WriteFieldError
		}
		if err = p.writeField3(oprot); err != nil {
			fieldId = 3
			goto WriteFieldError
		}

	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *Alert) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("Id", thrift.I64, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteI64(p.Id); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *Alert) writeField2(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("Time", thrift.STRING, 2); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.Time); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 end error: ", p), err)
}

func (p *Alert) writeField3(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("IndexNum", thrift.MAP, 3); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteMapBegin(thrift.I64, thrift.DOUBLE, len(p.IndexNum)); err != nil {
		return err
	}
	for k, v := range p.IndexNum {

		if err := oprot.WriteI64(k); err != nil {
			return err
		}

		if err := oprot.WriteDouble(v); err != nil {
			return err
		}
	}
	if err := oprot.WriteMapEnd(); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 3 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 3 end error: ", p), err)
}

func (p *Alert) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("Alert(%+v)", *p)
}

func (p *Alert) DeepEqual(ano *Alert) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.Id) {
		return false
	}
	if !p.Field2DeepEqual(ano.Time) {
		return false
	}
	if !p.Field3DeepEqual(ano.IndexNum) {
		return false
	}
	return true
}

func (p *Alert) Field1DeepEqual(src int64) bool {

	if p.Id != src {
		return false
	}
	return true
}
func (p *Alert) Field2DeepEqual(src string) bool {

	if strings.Compare(p.Time, src) != 0 {
		return false
	}
	return true
}
func (p *Alert) Field3DeepEqual(src map[int64]float64) bool {

	if len(p.IndexNum) != len(src) {
		return false
	}
	for k, v := range p.IndexNum {
		_src := src[k]
		if v != _src {
			return false
		}
	}
	return true
}

type AlertsResponse struct {
	Alerts []*Alert `thrift:"alerts,1" json:"alerts"`
}

func NewAlertsResponse() *AlertsResponse {
	return &AlertsResponse{}
}

func (p *AlertsResponse) GetAlerts() []*Alert {
	return p.Alerts
}
func (p *AlertsResponse) SetAlerts(val []*Alert) {
	p.Alerts = val
}

var fieldIDToName_AlertsResponse = map[int16]string{
	1: "alerts",
}

func (p *AlertsResponse) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.LIST {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}

		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_AlertsResponse[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *AlertsResponse) ReadField1(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return err
	}
	p.Alerts = make([]*Alert, 0, size)
	for i := 0; i < size; i++ {
		_elem := NewAlert()
		if err := _elem.Read(iprot); err != nil {
			return err
		}

		p.Alerts = append(p.Alerts, _elem)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return err
	}
	return nil
}

func (p *AlertsResponse) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("AlertsResponse"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}

	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *AlertsResponse) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("alerts", thrift.LIST, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteListBegin(thrift.STRUCT, len(p.Alerts)); err != nil {
		return err
	}
	for _, v := range p.Alerts {
		if err := v.Write(oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteListEnd(); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *AlertsResponse) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("AlertsResponse(%+v)", *p)
}

func (p *AlertsResponse) DeepEqual(ano *AlertsResponse) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.Alerts) {
		return false
	}
	return true
}

func (p *AlertsResponse) Field1DeepEqual(src []*Alert) bool {

	if len(p.Alerts) != len(src) {
		return false
	}
	for i, v := range p.Alerts {
		_src := src[i]
		if !v.DeepEqual(_src) {
			return false
		}
	}
	return true
}
