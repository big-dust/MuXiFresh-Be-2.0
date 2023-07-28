package svc

import (
	"MuXiFresh-Be-2.0/app/modify/cmd/api/internal/config"
)

type ServiceContext struct {
	Config config.Config
	ModifytheUseravatar     modifytheuseravatar.ModifytheUseravatar     
    ModifytheUsername       modifytheusername.ModifytheUsername          
	ModifytheUsertype       modifytheusertype.ModifytheUsertype    
	AdjustAdmissionstatus   adjustadmissionstatus.AdjustAdmissionstatus
}


func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		ModifytheUseravatar:     modifytheuseravatar.NewModifytheUseravatar(zrpc.MustNewClient(c.ModifytheUseravatar)),
		ModifytheUsername:       modifytheusername.NewModifytheUsername(zrpc.MustNewClient(c.ModifytheUsername)),
		ModifytheUsertype:       modifytheusertype.NewModifytheUsertype(zrpc.MustNewClient(c.ModifytheUsertype)),
		AdjustAdmissionstatus:   adjustadmissionstatus.NewAdjustAdmissionstatus(zrpc.MustNewClient(c.AdjustAdmissionstatus)),
	}
}
