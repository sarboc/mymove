package testdatagen

import (
	"github.com/gobuffalo/pop"

	"github.com/transcom/mymove/pkg/models"
)

// MakePPM creates a single Personally Procured Move and its associated Move and Orders
func MakePPM(db *pop.Connection) (models.PersonallyProcuredMove, error) {
	move, err := MakeMove(db)
	if err != nil {
		return models.PersonallyProcuredMove{}, err
	}

	ppm, verrs, err := move.CreateNewPPM(db)
	if verrs.HasAny() || err != nil {
		return models.PersonallyProcuredMove{}, err
	}

	return *ppm, nil
}

// MakePPMData creates 5 PPMs (and in turn a more and set of Orders for each)
func MakePPMData(db *pop.Connection) {
	for i := 0; i < 3; i++ {
		MakePPM(db)
	}
}
