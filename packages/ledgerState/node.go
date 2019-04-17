package ledgerState

import "github.com/iotadevelopment/consensus/node"

func OnNeighborChangesOpinion(neighbor Neighbor, conflictSet ConflictSet, newOpinion Opinion) {
    conflictSet.OpinionsOfNeighbors[neighbor.ID] = newOpinion

    conflictSet.debounce(func() {
        AdjustOpinion(conflictSet)
    })
}

func AdjustOpinion(conflictSet ConflictSet) {
    // check if any of the bundles in the conflict set can gather a majority
    for _, bundle := range conflictSet.Bundles {
        neighborsPreferringThisBundle := 0

        for _, opinion := range this.OpinionOfNeighbors {
            if opinion.PreferredBundleHash == bundle.Hash {
                neighborsPreferringThisBundle++
            }
        }

        if neighborsPreferringThisBundle > len(this.Neighbors) / 2 {
            // only change if we haven't adopted this opinion already
            if this.Opinion.PreferredBundleHash != bundle.Hash {
                node.RoundCounter++

                node.setOpinion(&bundle)
            }

            return
        }
    }

    // otherwise switch to prefer cancelling all bundles of the conflict set
    if len(this.OpinionOfNeighbors) == len(this.Neighbors) && this.Opinion.PreferredBundleHash != CANCEL_ALL.Hash {
        node.RoundCounter++

        node.setOpinion(CANCEL_ALL)
    }
}
