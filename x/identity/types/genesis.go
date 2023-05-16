package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		ParticipantList: []Participant{},
		TeamList:        []Team{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in participant
	participantIndexMap := make(map[string]struct{})

	for _, elem := range gs.ParticipantList {
		index := string(ParticipantKey(elem.Index))
		if _, ok := participantIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for participant")
		}
		participantIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in team
	teamIndexMap := make(map[string]struct{})

	for _, elem := range gs.TeamList {
		index := string(TeamKey(elem.Index))
		if _, ok := teamIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for team")
		}
		teamIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
