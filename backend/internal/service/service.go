package service

// UserService 用户服务接口（后续可拓展为具体实现）
type UserService interface {
	// TODO: 定义用户相关的业务接口，例如：
	// ListUsers(ctx context.Context) ([]*model.User, error)
	// CreateUser(ctx context.Context, req dto.CreateUserRequest) (*model.User, error)
}

// TaskService 任务服务接口
type TaskService interface {
	// TODO: 定义任务相关的业务接口
}

// PlanService 计划服务接口
type PlanService interface {
	// TODO: 定义计划相关的业务接口
}
