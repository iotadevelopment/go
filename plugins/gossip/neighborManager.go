package gossip

import (
    "time"
)

type NeighborManager struct {
    Events           neighborManagerEvents
    staticNeighbors  map[string]*Neighbor
    dynamicNeighbors map[string]*Neighbor
}

func (this *NeighborManager) LaunchConnections() {
    for {
        select {
        case <- time.After(3 * time.Second):
            for _, neighbor := range this.staticNeighbors {
                neighbor.Connect()
            }

            for _, neighbor := range this.dynamicNeighbors {
                neighbor.Connect()
            }
        }
    }
}

func (this *NeighborManager) AddNeighbor(neighbor *Neighbor, static bool) {
    if static {
        this.staticNeighbors[neighbor.GetAddress()] = neighbor
    } else {
        this.dynamicNeighbors[neighbor.GetAddress()] = neighbor
    }

    this.Events.AddNeighbor.Trigger(neighbor)
}

func (this *NeighborManager) RemoveNeighbor(neighbor *Neighbor) bool {
    address := neighbor.GetAddress()

    if neighbor, exists := this.staticNeighbors[address]; exists {
        delete(this.staticNeighbors, address)

        this.Events.RemoveNeighbor.Trigger(neighbor)

        return true
    }

    if neighbor, exists := this.dynamicNeighbors[address]; exists {
        delete(this.dynamicNeighbors, address)

        this.Events.RemoveNeighbor.Trigger(neighbor)

        return true
    }

    return false
}

func (this *NeighborManager) GetNeighbor(address string) *Neighbor {
    if neighbor, exists := this.staticNeighbors[address]; exists {
        return neighbor
    }

    if neighbor, exists := this.dynamicNeighbors[address]; exists {
        return neighbor
    }

    return nil
}

