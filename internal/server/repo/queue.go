package repo

import (
	"github.com/easysoft/zagent/internal/comm/const"
	"github.com/easysoft/zagent/internal/server/conf"
	"github.com/easysoft/zagent/internal/server/model"
	"gorm.io/gorm"
	"strings"
	"time"
)

type QueueRepo struct {
	BaseRepo
	DB *gorm.DB `inject:""`
}

func NewQueueRepo() *QueueRepo {
	return &QueueRepo{}
}

func (r QueueRepo) QueryForExec() (queues []model.Queue) {
	queues = make([]model.Queue, 0)

	r.DB.Where("progress=? OR progress=?", consts.ProgressCreated, consts.ProgressPending).
		Order("priority").Find(&queues)

	return
}
func (r QueueRepo) QueryByTask(taskID uint) (queues []model.Queue) {
	queues = make([]model.Queue, 0)

	r.DB.Where("task_id=?", taskID).Order("id").Find(&queues)

	return
}

func (r QueueRepo) GetQueue(id uint) (queue model.Queue) {
	r.DB.Where("id=?", id).First(&queue)

	return
}

func (r QueueRepo) Save(queue *model.Queue) (err error) {
	err = r.DB.Model(&queue).
		Omit("StartTime", "PendingTime", "ResultTime", "TimeoutTime").
		Create(&queue).Error
	return
}

func (r QueueRepo) DeleteInSameGroup(groupId uint, serials []string) (err error) {
	r.DB.Where("group_id=? AND serial IN (?)", groupId, strings.Join(serials, ",")).Delete(&model.Queue{})
	return
}

func (r QueueRepo) Start(queue model.Queue) (err error) {
	r.DB.Model(&queue).Where("id=?", queue.ID).Updates(
		map[string]interface{}{"progress": consts.ProgressInProgress, "start_time": time.Now(), "retry": gorm.Expr("retry +1")})
	return
}
func (r QueueRepo) Pending(queueId uint) (err error) {
	r.DB.Model(&model.Queue{}).
		Where("id=?", queueId).
		Updates(map[string]interface{}{
			"progress": consts.ProgressPending,
		})

	r.DB.Model(&model.Queue{}). // only update once, used for timeout checking
					Where("id=? AND pending_time IS NULL", queueId).
					Updates(map[string]interface{}{
			"pending_time": time.Now(),
		})

	return
}

func (r QueueRepo) SetTimeout(id uint) (err error) {
	r.DB.Model(&model.Queue{}).Where("id=?", id).Updates(
		map[string]interface{}{"progress": consts.ProgressTimeout, "timeout_time": time.Now()})
	return
}

func (r QueueRepo) QueryTimeout() (queues []model.Queue) {
	queues = make([]model.Queue, 0)

	where := ""
	if serverConf.Config.DB.Adapter == "sqlite3" {
		where = "(progress = ? AND strftime('%s','now') - strftime('%s',pending_time) > ?)" +
			" OR (progress = ? AND strftime('%s','now') - strftime('%s',start_time) > ?)"
	} else if serverConf.Config.DB.Adapter == "mysql" {
		where = "(progress = ? AND unix_timestamp(NOW()) - unix_timestamp(pending_time) > ?)" +
			" OR (progress = ? AND unix_timestamp(NOW()) - unix_timestamp(start_time) > ?)"
	}

	r.DB.Where(where,
		consts.ProgressPending, consts.WaitForExecTime*60,
		consts.ProgressInProgress, consts.WaitForResultTime*60).
		Order("priority").Find(&queues)
	return
}
func (r QueueRepo) QueryTimeoutOrFailedForRetry() (queues []model.Queue) {
	queues = make([]model.Queue, 0)

	r.DB.Where("retry < ? AND (progress = ? OR status = ? )",
		consts.QueueRetryTime, consts.ProgressTimeout, consts.StatusFail).
		Order("priority").Find(&queues)
	return
}

func (r QueueRepo) SetQueueStatus(queueId uint, progress consts.BuildProgress, status consts.BuildStatus) {
	r.DB.Model(&model.Queue{}).Where("id=?", queueId).Updates(
		map[string]interface{}{"progress": progress, "status": status, "result_time": time.Now(), "updated_at": time.Now()})
	return
}

func (r QueueRepo) UpdateVm(queueId, vmId uint, status consts.BuildProgress) {
	r.DB.Model(&model.Queue{}).Where("id=?", queueId).Updates(
		map[string]interface{}{"vmId": vmId, "progress": status, "updated_at": time.Now()})
	return
}

func (r QueueRepo) CancelQueuesNotExec(taskId uint) {
	r.DB.Model(&model.Queue{}).
		Where("task_id=? AND (progress=? OR progress=?)",
			taskId, consts.ProgressCreated, consts.ProgressPending).
		Updates(map[string]interface{}{"progress": consts.ProgressCancel})
	return
}
