package ipvs

import (
	"github.com/dpvs-agent/pkg/ipc/pool"
	"github.com/dpvs-agent/pkg/ipc/types"

	apiVs "github.com/dpvs-agent/restapi/operations/virtualserver"

	"github.com/go-openapi/runtime/middleware"
	"github.com/hashicorp/go-hclog"
)

type putVsRs struct {
	connPool *pool.ConnPool
	logger   hclog.Logger
}

func NewPutVsRs(cp *pool.ConnPool, parentLogger hclog.Logger) *putVsRs {
	logger := hclog.Default()
	if parentLogger != nil {
		logger = parentLogger.Named("PutVsVipPortRs")
	}
	return &putVsRs{connPool: cp, logger: logger}
}

func (h *putVsRs) Handle(params apiVs.PutVsVipPortRsParams) middleware.Responder {
	front := types.NewRealServerFront()
	if err := front.ParseVipPortProto(params.VipPort); err != nil {
		h.logger.Error("Convert to virtual server failed", "VipPort", params.VipPort, "Error", err.Error())
		return apiVs.NewPutVsVipPortRsInvalidFrontend()
	}

	var rss []*types.RealServerSpec
	if params.Rss != nil {
		rss = make([]*types.RealServerSpec, len(params.Rss.Items))
		for i, rs := range params.Rss.Items {
			var fwdmode types.DpvsFwdMode
			fwdmode.FromString(rs.Mode)
			rss[i] = types.NewRealServerSpec()
			rss[i].SetAf(front.GetAf())
			rss[i].SetAddr(rs.IP)
			rss[i].SetPort(rs.Port)
			rss[i].SetProto(front.GetProto())
			rss[i].SetWeight(uint32(rs.Weight))
			rss[i].SetFwdMode(fwdmode)
			rss[i].SetInhibited(rs.Inhibited)
			rss[i].SetOverloaded(rs.Inhibited)
		}
	}

	healthCheck := false
	if params.Healthcheck != nil {
		healthCheck = *params.Healthcheck
	}

	result := front.Edit(healthCheck, rss, h.connPool, h.logger)

	// h.logger.Info("Set real server sets done.", "VipPort", params.VipPort, "rss", rss, "result", result.String())
	switch result {
	case types.EDPVS_EXIST, types.EDPVS_OK:
		h.logger.Info("Set real server sets success.", "VipPort", params.VipPort, "rss", rss, "result", result.String())
		return apiVs.NewPutVsVipPortRsOK()
	case types.EDPVS_NOTEXIST:
		if healthCheck {
			h.logger.Error("Edit not exist real server.", "VipPort", params.VipPort, "rss", rss, "result", result.String())
			return apiVs.NewPutVsVipPortRsInvalidFrontend()
		}
		h.logger.Error("Unreachable branch")
	default:
		h.logger.Error("Set real server sets failed.", "VipPort", params.VipPort, "rss", rss, "result", result.String())
		return apiVs.NewPutVsVipPortRsInvalidBackend()
	}
	return apiVs.NewPutVsVipPortRsFailure()
}
