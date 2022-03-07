/**
 * @author  qyanzh
 * @create  2022/03/07 20:14
 */

package service

var taskService *TaskService

func init() {
	taskService = NewTaskService()
}
