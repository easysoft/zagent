package repo

import (
	"github.com/easysoft/zv/internal/comm/const"
	serverConf "github.com/easysoft/zv/internal/server/conf"
	"github.com/easysoft/zv/internal/server/model"
	"gorm.io/gorm"
	"time"
)

type QueueRepo struct {
	BaseRepo
	DB *gorm.DB `inject:""`
}

func NewQueueRepo() *QueueRepo {
	return &QueueRepo{}
}

func (r QueueRepo) QueryByTask(taskID uint) (queues []model.Queue) {
	queues = make([]model.Queue, 0)

	r.DB.Model(&model.Queue{}).
		Where("task_id=? AND NOT deleted", taskID).
		Order("id").
		Find(&queues)

	return
}

func (r QueueRepo) QueryForExec() (queues []model.Queue) {
	queues = make([]model.Queue, 0)

	r.DB.Model(&model.Queue{}).Where("progress=? OR progress=? OR progress=?",
		consts.ProgressCreated,
		consts.ProgressResPending,
		consts.ProgressResReady).
		Order("priority").Find(&queues)

	return
}
func (r QueueRepo) QueryForTimeout() (queues []model.Queue) {
	queues = make([]model.Queue, 0)

	where := ""
	if serverConf.Inst.DB.Adapter == "sqlite3" {
		where = "(progress = ? AND strftime('%s','now') - strftime('%s',res_pending_time) > ?)" +
			" OR (progress = ? AND strftime('%s','now') - strftime('%s',res_launched_time) > ?)" +
			" OR (progress = ? AND strftime('%s','now') - strftime('%s',run_time) > ?)"
	} else if serverConf.Inst.DB.Adapter == "mysql" {
		where = "(progress = ? AND unix_timestamp(NOW()) - unix_timestamp(res_pending_time) > ?)" +
			" OR (progress = ? AND unix_timestamp(NOW()) - unix_timestamp(res_launched_time) > ?)" +
			" OR (progress = ? AND unix_timestamp(NOW()) - unix_timestamp(run_time) > ?)"
	}

	r.DB.Model(&model.Queue{}).Where(where,
		consts.ProgressResPending, consts.WaitResPendingTimeout,
		consts.ProgressResLaunched, consts.WaitResReadyTimeout,
		consts.ProgressRunning, consts.WaitRunCompletedTimeout).
		Order("priority").Find(&queues)
	return
}
func (r QueueRepo) QueryForRetry() (queues []model.Queue) {
	queues = make([]model.Queue, 0)

	r.DB.Model(&model.Queue{}).Where("(progress = ? OR status = ? ) AND retry < ?",
		consts.ProgressTimeout, consts.StatusFail,
		consts.QueueRetryTime,
	).
		Order("priority").Find(&queues)
	return
}

func (r QueueRepo) GetQueue(id uint) (queue model.Queue) {
	r.DB.Model(&model.Queue{}).Where("id=?", id).First(&queue)

	return
}
func (r QueueRepo) GetByVmId(vmId uint) (queue model.Queue) {
	r.DB.Model(&model.Queue{}).Where("vm_id=?", vmId).First(&queue)

	return
}

func (r QueueRepo) Save(queue *model.Queue) (err error) {
	err = r.DB.Model(&model.Queue{}).
		Omit("ResPendingTime", "ResLaunchedTime", "RunTime", "EndTime", "TimeoutTime").
		Create(&queue).Error
	return
}

func (r QueueRepo) ResLaunched(queueId uint) (err error) {
	r.DB.Model(&model.Queue{}).Where("id=?", queueId).Updates(
		map[string]interface{}{"progress": consts.ProgressResLaunched, "res_launched_time": time.Now()})
	return
}
func (r QueueRepo) ResReady(queueId uint) (err error) {
	r.DB.Model(&model.Queue{}).Where("id=?", queueId).Updates(
		map[string]interface{}{"progress": consts.ProgressResReady})
	return
}
func (r QueueRepo) ResPending(queueId uint) (err error) {
	r.DB.Model(&model.Queue{}).
		Where("id=?", queueId).
		Updates(map[string]interface{}{
			"progress": consts.ProgressResPending,
		})

	r.DB.Model(&model.Queue{}). // only update once, used for timeout checking
					Where("id=? AND res_pending_time IS NULL", queueId).
					Updates(map[string]interface{}{
			"res_pending_time": time.Now(),
		})

	return
}
func (r QueueRepo) Run(queue model.Queue) (err error) {
	r.DB.Model(&model.Queue{}).Where("id=?", queue.ID).Updates(
		map[string]interface{}{
			"progress": consts.ProgressRunning, "run_time": time.Now()})
	return
}
func (r QueueRepo) Timeout(id uint) (err error) {
	r.DB.Model(&model.Queue{}).Where("id=?", id).Updates(
		map[string]interface{}{"progress": consts.ProgressTimeout, "timeout_time": time.Now()})
	return
}

func (r QueueRepo) SaveResult(queueId uint, progress consts.BuildProgress, status consts.BuildStatus) {
	r.DB.Model(&model.Queue{}).Where("id=?", queueId).Updates(
		map[string]interface{}{"progress": progress, "status": status, "result_time": time.Now()})
	return
}

func (r QueueRepo) UpdateVm(queueId, vmId uint) {
	r.DB.Model(&model.Queue{}).Where("id=?", queueId).Updates(
		map[string]interface{}{"vm_id": vmId})
	return
}

func (r QueueRepo) RemoveOldQueuesByTask(taskId uint) {
	r.DB.Model(&model.Queue{}).Where("task_id=?", taskId).Updates(
		map[string]interface{}{"deleted": true})
	return
}

func (r QueueRepo) Retry(queue model.Queue) (err error) {
	r.DB.Model(&model.Queue{}).
		Where("id=?", queue.ID).
		Updates(map[string]interface{}{"retry": gorm.Expr("retry +1")})
	return
}
