package filproofs

import abi "github.com/filecoin-project/specs/actors/abi"
import sector "github.com/filecoin-project/specs/systems/filecoin_mining/sector"

func (fps *FilecoinProofsSubsystem_I) VerifySeal(sealVerifyInfo abi.SealVerifyInfo) bool {
	cfg := sector.RegisteredProofInstance(sealVerifyInfo.OnChain.RegisteredProof).Cfg().As_SealCfg()
	sdr := WinSDRParams(cfg)
	return sdr.VerifySeal(sealVerifyInfo)
}