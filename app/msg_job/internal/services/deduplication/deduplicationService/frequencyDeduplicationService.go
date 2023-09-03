package deduplicationService

//type frequencyDeduplicationService struct {
//	svcCtx *svc.ServiceContext
//}
//
//func NewFrequencyDeduplicationService(svcCtx *svc.ServiceContext) structs.DuplicationService {
//	return &frequencyDeduplicationService{svcCtx: svcCtx}
//}
//
//func (c frequencyDeduplicationService) Deduplication(ctx context.Context, taskInfo *types.TaskInfo, param structs.DeduplicationConfigItem) error {
//	return srv.NewFrequencyDeduplicationService(c.svcCtx, limit.NewSlideWindowLimitService(c.svcCtx)).
//		Deduplication(ctx, taskInfo, param)
//}
