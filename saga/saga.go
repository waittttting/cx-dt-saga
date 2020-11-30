package saga

import (
	"cx-dt-saga/sony_falke"
	"errors"
)


const (
	Compensation = "Compensation"
	Execution = "Execution"
)

type HandleFunc func(string)error



type Manager struct {}

func (sm *Manager) Start() {
	go func() {

	}()
}

func (sm *Manager) SagaDTHandler(operates []map[string]HandleFunc) error {

	// 生成事务ID
	txId, err := sony_falke.GetSequenceId()
	if err != nil {
		return err
	}

	for i, op := range operates {
		err := sm.handleOperates(op, txId, i + 1)
		if err != nil {
			return err
		}
	}
	return nil
}

func (sm *Manager) handleOperates(operate map[string]HandleFunc, txId string, step int) error {
	// 执行函数
	var execution HandleFunc
	if v, find := operate[Execution]; find {
		execution = v
	} else {
		return errors.New("execution not find")
	}
	// 补偿函数
	var compensation HandleFunc
	if v, find := operate[Compensation]; find {
		compensation = v
	} else {
		return errors.New("compensation not find")
	}


	return nil
}