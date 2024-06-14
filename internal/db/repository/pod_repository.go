package repository

import (
	"github.com/jinzhu/gorm"
	model "paas/internal/db/models"
)

type PodRepository struct {
	mysqlDB *gorm.DB
}

/*
*
主要是对数据库pod进行管理
*/
type IPodRepository interface {
	//初始化表
	InitTable() error
	//根据ID查找数据
	FindPodByID(int64) (*model.Pod, error)
	//创建一条 Pod 数据
	CreatePod(*model.Pod) (int64, error)
	//根据ID删除一条 Pod 数据
	DeletePodByID(int64) error
	//修改一条数据
	UpdatePod(*model.Pod) error
	//查找Pod所有数据
	FindAll() ([]model.Pod, error)
}

func NewPodRepository(db *gorm.DB) IPodRepository {
	return &PodRepository{mysqlDB: db}
}
func (p PodRepository) InitTable() error {
	return p.mysqlDB.CreateTable(&model.Pod{}, &model.PodPort{}, &model.PodEnv{}).Error
}

func (p PodRepository) FindPodByID(i int64) (*model.Pod, error) {
	p.mysqlDB.f
}

func (p PodRepository) CreatePod(pod *model.Pod) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (p PodRepository) DeletePodByID(i int64) error {
	//TODO implement me
	panic("implement me")
}

func (p PodRepository) UpdatePod(pod *model.Pod) error {
	//TODO implement me
	panic("implement me")
}

func (p PodRepository) FindAll() ([]model.Pod, error) {
	//TODO implement me
	panic("implement me")
}
