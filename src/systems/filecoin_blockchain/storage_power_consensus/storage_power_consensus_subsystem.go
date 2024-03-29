package storage_power_consensus

import (
	filcrypto "github.com/filecoin-project/specs/libraries/filcrypto"
	block "github.com/filecoin-project/specs/systems/filecoin_blockchain/struct/block"
	base_mining "github.com/filecoin-project/specs/systems/filecoin_mining"
	addr "github.com/filecoin-project/specs/systems/filecoin_vm/actor/address"
)

const FINALITY = 500

const (
	SPC_LOOKBACK_RANDOMNESS = 300      // this is EC.K maybe move it there. TODO
	SPC_LOOKBACK_TICKET     = 1        // we chain blocks together one after the other
	SPC_LOOKBACK_POST       = 1        // cheap to generate, should be set as close to current TS as possible
	SPC_LOOKBACK_SEAL       = FINALITY // should be set to finality
)

const (
	VRFPersonalizationTicket        = iota
	VRFPersonalizationElectionProof = iota
)

// Storage Power Consensus Subsystem

func (spc *StoragePowerConsensusSubsystem_I) ValidateBlock(block block.Block_I) error {
	panic("")
	// minerPK := PowerTable.GetMinerPublicKey(block.MinerAddress())
	// minerPower := PowerTable.GetMinerPower(block.MinerAddress())

	// // 1. Verify miner has not been slashed and is still valid miner
	// if minerPower <= 0 {
	// 	return StoragePowerConsensusError("block miner not valid")
	// }

	// // 2. Verify ParentWeight
	// if block.Weight() != computeTipsetWeight(block.Parents()) {
	// 	return errors.New("invalid parent weight")
	// }

	// // 3. Verify Tickets
	// if !validateTicket(block.Ticket, minerPK) {
	// 	return StoragePowerConsensusError("ticket was invalid")
	// }

	// // 4. Verify ElectionProof construction
	// if !ValidateElectionProof(block.Height, block.ElectionProof, block.MinerAddress) {
	// 	return StoragePowerConsensusError("election proof was not a valid signature of the last ticket")
	// }

	// // 5. and value
	// if !IsWinningElectionProof(block.ElectionProof, block.MinerAddress) {
	// 	return StoragePowerConsensusError("election proof was not a winner")
	// }

	// return nil
}

func (spc *StoragePowerConsensusSubsystem_I) validateTicket(ticket block.Ticket, pk filcrypto.PubKey) bool {
	panic("")
	// T1 := storagePowerConsensus.GetTicketProductionSeed(sms.CurrentChain, sms.Blockchain.LatestEpoch())
	// input := VRFPersonalizationTicket
	// input.append(T1.Output)
	// return ticket.Verify(input, pk)
}

func (spc *StoragePowerConsensusSubsystem_I) computeTipsetWeight(tipset block.Tipset) block.ChainWeight {
	panic("TODO")
}

func (spc *StoragePowerConsensusSubsystem_I) StoragePowerConsensusError(errMsg string) StoragePowerConsensusError {
	panic("TODO")
}

func (spc *StoragePowerConsensusSubsystem_I) GetTicketProductionSeed(chain block.Chain, epoch block.ChainEpoch) base_mining.SealSeed {
	panic("")

	// return &base_mining.SealSeed
	// 	chain.TicketAtEpoch(epoch - SPC_LOOKBACK_TICKET),
	// }
}

func (spc *StoragePowerConsensusSubsystem_I) GetElectionProofSeed(chain block.Chain, epoch block.ChainEpoch) base_mining.SealSeed {
	panic("")

	// return &base_mining.SealSeed_I{
	// 	chain.TicketAtEpoch(epoch - SPC_LOOKBACK_RANDOMNESS),
	// }
}

func (spc *StoragePowerConsensusSubsystem_I) GetSealSeed(chain block.Chain, epoch block.ChainEpoch) base_mining.SealSeed {
	panic("")

	// return &base_mining.SealSeed_I{
	// 	chain.TicketAtEpoch(epoch - SPC_LOOKBACK_SEAL),
	// }
}

func (spc *StoragePowerConsensusSubsystem_I) GetPoStChallenge(chain block.Chain, epoch block.ChainEpoch) base_mining.PoStChallenge {
	panic("")

	// return &base_mining.PoStChallenge_I{
	// 	chain.TicketAtEpoch(epoch - SPC_LOOKBACK_POST),
	// }
}

func (spc *StoragePowerConsensusSubsystem_I) ValidateElectionProof(height block.ChainEpoch, electionProof block.ElectionProof, workerAddr addr.Address) bool {
	panic("")
	// // 1. Check that ElectionProof was validated in appropriate time
	// if height+electionProof.ElectionNonce > clock.roundTime+1 {
	// 	return false
	// }

	// // 2. Determine that ticket was validly scratched
	// minerPK := spc.PowerTable.GetMinerPublicKey(workerAddr)
	// input := VRFPersonalizationElectionProof
	// TK := storagePowerConsensus.GetElectionProofSeed(sms.CurrentChain, sms.block.LatestEpoch())
	// input.append(TK.Output)
	// input.appent(electionProof.ElectionNonce)

	// return electionProof.Verify(input, minerPK)
}

func (spc *StoragePowerConsensusSubsystem_I) IsWinningElectionProof(electionProof block.ElectionProof, workerAddr addr.Address) bool {
	panic("")
	// 1. Determine miner power fraction
	// minerPower := spc.PowerTable.GetMinerPower(workerAddr)
	// totalPower := spc.PowerTable.GetTotalPower()

	// // Conceptually we are mapping the pseudorandom, deterministic VRFOutput onto [0,1]
	// // by dividing by 2^HashLen (64 Bytes using Sha256) and comparing that to the miner's
	// // power (portion of network storage).
	// return (minerPower*2^(len(electionProof.Output)*8) < electionProof.Output*totalPower)
}

func (spc *StoragePowerConsensusSubsystem_I) GetFinality() block.ChainEpoch {
	panic("")
	// return FINALITY
}
