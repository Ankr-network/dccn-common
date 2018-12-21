package ankr_const 

const TaskManagerQueueName = "task_manager"
const NewTaskEvent = "New"
const UpdateTaskEvent = "Update"
const CancelTaskEvent = "Cancel"

const TaskStatusNew = "new"
const TaskStatusRunning = "running"
const TaskStatusStartFailed = "startFailed"
const TaskStatusUpdating = "updating"
const TaskStatusUpdateFailed = "updateFailed"
const TaskStatusCancelled = "cancelled"
const TaskStatusDone = "done"

const DataCenterTaskStartSuccess = "StartSuccess"
const DataCenterTaskStartFailure = "StartFailure"
const DataCenterTaskUpdateSuccess = "UpdateSuccess"
const DataCenterTaskUpdateFailure = "UpdateFailure"
const DataCenterTaskDone = "Done"
const DataCenterTaskCancelled = "Cancelled"

const CliReplyStatusSuccess = "Success"
const CliReplyStatusFailure = "Failure"

const DefaultPort = 50051